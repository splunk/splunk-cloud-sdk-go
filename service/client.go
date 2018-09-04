// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

/*
Package service implements a service client which is used to communicate
with Search Service endpoints
*/
package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"sync"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

// Declare constants for service package
const (
	AuthorizationType = "Bearer"
)

// A Client is used to communicate with service endpoints
type Client struct {
	// config
	config *Config
	// tokenContext is the access token to include in "Authorization: Bearer" headers and related context information
	tokenContext *idp.Context
	// HTTP Client used to interact with endpoints
	httpClient *http.Client
	// responseHandlers is a slice of handlers to call after a response has been received in the client
	responseHandlers []ResponseHandler
	// SearchService talks to the Splunk Cloud search service
	SearchService *SearchService
	// CatalogService talks to the Splunk Cloud catalog service
	CatalogService *CatalogService
	// IngestService talks to the Splunk Cloud ingest service
	IngestService *IngestService
	// IdentityService talks to Splunk Cloud IAC service
	IdentityService *IdentityService
	// KVStoreService talks to Splunk Cloud kvstore service
	KVStoreService *KVStoreService
	// ActionService talks to Splunk Cloud action service
	ActionService *ActionService
	// StreamsService talks to SSC streams service
	StreamsService *StreamsService
}

// Request extends net/http.Request to track number of total attempts and error
// counts by type of error
type Request struct {
	*http.Request
	NumAttempts     uint
	NumErrorsByType map[string]uint
}

// GetNumErrorsByResponseCode returns number of attemps for a given response code >= 400
func (r *Request) GetNumErrorsByResponseCode(respCode int) uint {
	code := fmt.Sprintf("%d", respCode)
	if val, ok := r.NumErrorsByType[code]; ok {
		return val
	}
	return 0
}

// UpdateToken replaces the access token in the `Authorization: Bearer` header
func (r *Request) UpdateToken(accessToken string) {
	r.Header.Set("Authorization", fmt.Sprintf("%s %s", AuthorizationType, accessToken))
}

// Config is used to set the client specific attributes
type Config struct {
	// TokenRetriever to gather access tokens to be sent in the Authorization: Bearer header on client initialization and upon encountering a 401 response
	TokenRetriever idp.TokenRetriever
	// Token to be sent in the Authorization: Bearer header (not required if TokenRetriever is specified)
	Token string
	// Url string
	URL string
	// TenantID
	TenantID string
	// Timeout
	Timeout time.Duration
	// ResponseHandlers is a slice of handlers to call after a response has been received in the client
	ResponseHandlers []ResponseHandler
}

// RequestParams contains all the optional request URL parameters
type RequestParams struct {
	// Http method name
	Method string
	// Http url
	URL url.URL
	// Body parameter
	Body interface{}
	// Additional headers
	Headers map[string]string
}

// service provides the interface between client and services
type service struct {
	client *Client
}

// NewRequest creates a new HTTP Request and set proper header
func (c *Client) NewRequest(httpMethod, url string, body io.Reader, headers map[string]string) (*Request, error) {
	request, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}
	if c.tokenContext != nil && len(c.tokenContext.AccessToken) > 0 {
		request.Header.Set("Authorization", fmt.Sprintf("%s %s", AuthorizationType, c.tokenContext.AccessToken))
		request.Header.Set("Splunk-Client", fmt.Sprintf("%s/%s", UserAgent, Version))
	}
	request.Header.Set("Content-Type", "application/json")
	if len(headers) != 0 {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}
	retryRequest := &Request{request, 0, make(map[string]uint)}
	return retryRequest, nil
}

// BuildURL creates full Splunk Cloud URL with the client cached tenantID
func (c *Client) BuildURL(queryValues url.Values, urlPathParts ...string) (url.URL, error) {
	var buildPath = ""
	for _, pathPart := range urlPathParts {
		buildPath = path.Join(buildPath, url.PathEscape(pathPart))
	}
	if queryValues == nil {
		queryValues = url.Values{}
	}
	var u url.URL
	if len(c.config.TenantID) == 0 {
		return u, errors.New("A non-empty tenant ID must be set on client")
	}

	clientURL, err := c.GetURL()
	if err != nil {
		return u, err
	}

	u = url.URL{
		Scheme:   clientURL.Scheme,
		Host:     clientURL.Host,
		Path:     path.Join(c.config.TenantID, buildPath),
		RawQuery: queryValues.Encode(),
	}
	return u, nil
}

// BuildURLWithTenantID creates full Splunk Cloud URL with tenantID
func (c *Client) BuildURLWithTenantID(tenantID string, queryValues url.Values, urlPathParts ...string) (url.URL, error) {
	var buildPath = ""
	for _, pathPart := range urlPathParts {
		buildPath = path.Join(buildPath, url.PathEscape(pathPart))
	}
	if queryValues == nil {
		queryValues = url.Values{}
	}
	var u url.URL
	if len(tenantID) == 0 {
		return u, errors.New("A non-empty tenant ID must be passed in for BuildURLWithTenantID")
	}

	clientURL, err := c.GetURL()
	if err != nil {
		return u, err
	}

	u = url.URL{
		Scheme:   clientURL.Scheme,
		Host:     clientURL.Host,
		Path:     path.Join(tenantID, buildPath),
		RawQuery: queryValues.Encode(),
	}
	return u, nil
}

// Do sends out request and returns HTTP response
func (c *Client) Do(req *Request) (*http.Response, error) {
	req.NumAttempts++
	response, err := c.httpClient.Do(req.Request)
	if err != nil {
		return nil, err
	}
	// If error response found, record number of errors by response code
	if response.StatusCode >= 400 {
		// TODO: This could be extended to include specific Splunk Cloud error fields in
		// addition to response code
		code := fmt.Sprintf("%d", response.StatusCode)
		if _, ok := req.NumErrorsByType[code]; ok {
			req.NumErrorsByType[code]++
		} else {
			req.NumErrorsByType[code] = 1
		}
	}
	for _, hr := range c.responseHandlers {
		response, err = hr.HandleResponse(c, req, response)
		if err != nil {
			return response, err
		}
	}
	return response, err
}

// Get implements HTTP Get call
func (c *Client) Get(requestParams RequestParams) (*http.Response, error) {
	requestParams.Method = http.MethodGet
	return c.DoRequest(requestParams)
}

// Post implements HTTP POST call
func (c *Client) Post(requestParams RequestParams) (*http.Response, error) {
	requestParams.Method = http.MethodPost
	return c.DoRequest(requestParams)
}

// Put implements HTTP PUT call
func (c *Client) Put(requestParams RequestParams) (*http.Response, error) {
	requestParams.Method = http.MethodPut
	return c.DoRequest(requestParams)
}

// Delete implements HTTP DELETE call
// RFC2616 does not explicitly forbid it but in practice some versions of server implementations (tomcat,
// netty etc) ignore bodies in DELETE requests
func (c *Client) Delete(requestParams RequestParams) (*http.Response, error) {
	requestParams.Method = http.MethodDelete
	return c.DoRequest(requestParams)
}

// Patch implements HTTP Patch call
func (c *Client) Patch(requestParams RequestParams) (*http.Response, error) {
	requestParams.Method = http.MethodPatch
	return c.DoRequest(requestParams)
}

// DoRequest creates and execute a new request
func (c *Client) DoRequest(requestParams RequestParams) (*http.Response, error) {
	var buffer *bytes.Buffer
	if contentBytes, ok := requestParams.Body.([]byte); ok {
		buffer = bytes.NewBuffer(contentBytes)
	} else {
		if content, err := json.Marshal(requestParams.Body); err == nil {
			buffer = bytes.NewBuffer(content)
		} else {
			return nil, err
		}
	}
	request, err := c.NewRequest(requestParams.Method, requestParams.URL.String(), buffer, requestParams.Headers)
	if err != nil {
		return nil, err
	}
	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	return util.ParseHTTPStatusCodeInResponse(response)
}

// UpdateTokenContext the access token in the Authorization: Bearer header and retains related context information
func (c *Client) UpdateTokenContext(ctx *idp.Context) {
	c.tokenContext = ctx
}

// GetURL returns the client config url string as a url.URL
func (c *Client) GetURL() (*url.URL, error) {
	parsed, err := url.Parse(c.config.URL)
	if c.config.URL == "" || err != nil {
		return nil, errors.New("url is not correct")
	}
	return parsed, nil
}

// NewClient creates a Client with config values passed in
func NewClient(config *Config) (*Client, error) {
	if config.TenantID == "" || config.URL == "" {
		return nil, errors.New("tenantID and url must be set")
	}

	// TODO: Token will eventually be fully replaced by TokenRetriever, it is retained for now for backwards compatibility
	if (config.TokenRetriever != nil && config.Token != "") || (config.TokenRetriever == nil && config.Token == "") {
		return nil, errors.New("either config.TokenRetriever or config.Token must be set, not both")
	}

	var handlers []ResponseHandler
	if config.Token != "" {
		// If static Token is provided then set the token retriever to no-op (just return static token)
		config.TokenRetriever = &idp.NoOpTokenRetriever{Context: &idp.Context{AccessToken: config.Token}}
		handlers = config.ResponseHandlers
	} else {
		// If TokenRetriever is provided, create an AuthnHandler to retry 401 requests using this interface and prepend before any custom handlers from the config
		authnHandler := AuthnResponseHandler{TokenRetriever: config.TokenRetriever}
		handlers = append([]ResponseHandler{ResponseHandler(authnHandler)}, config.ResponseHandlers...)
	}

	// Start by retrieving the access token
	ctx, err := config.TokenRetriever.GetTokenContext()
	if err != nil {
		return nil, fmt.Errorf("service.NewClient: error retrieving token: %s", err)
	}

	// Finally, initialize the Client
	c := &Client{
		config:           config,
		httpClient:       &http.Client{Timeout: config.Timeout},
		tokenContext:     ctx,
		responseHandlers: handlers,
	}
	c.SearchService = &SearchService{client: c}
	c.CatalogService = &CatalogService{client: c}
	c.IdentityService = &IdentityService{client: c}
	c.IngestService = &IngestService{client: c}
	c.KVStoreService = &KVStoreService{client: c}
	c.ActionService = &ActionService{client: c}
	c.StreamsService = &StreamsService{client: c}
	return c, nil
}

// NewBatchEventsSenderWithMaxAllowedError used to initialize dependencies and set values, the maxErrorsAllowed is the max number of errors allowed before the eventsender quit
func (c *Client) NewBatchEventsSenderWithMaxAllowedError(batchSize int, interval int64, maxErrorsAllowed int) (*BatchEventsSender, error) {
	// Rather than return a super general error for both it will block on batchSize first
	if batchSize == 0 {
		return nil, errors.New("batchSize cannot be 0")
	}
	if interval == 0 {
		return nil, errors.New("interval cannot be 0")
	}

	if maxErrorsAllowed < 0 {
		maxErrorsAllowed = 1
	}

	eventsChan := make(chan model.Event, batchSize)
	eventsQueue := make([]model.Event, 0, batchSize)
	quit := make(chan struct{}, 1)
	ticker := model.NewTicker(time.Duration(interval) * time.Millisecond)
	var wg sync.WaitGroup
	errorChan := make(chan string, maxErrorsAllowed)

	batchEventsSender := &BatchEventsSender{
		BatchSize:        batchSize,
		EventsChan:       eventsChan,
		EventsQueue:      eventsQueue,
		EventService:     c.IngestService,
		QuitChan:         quit,
		IngestTicker:     ticker,
		WaitGroup:        &wg,
		ErrorChan:        errorChan,
		IsRunning:        false,
		chanWaitInMilSec: 300,
		callbackFunc:     nil,
	}

	return batchEventsSender, nil
}

// NewBatchEventsSender used to initialize dependencies and set values
func (c *Client) NewBatchEventsSender(batchSize int, interval int64) (*BatchEventsSender, error) {
	return c.NewBatchEventsSenderWithMaxAllowedError(batchSize, interval, 1)
}
