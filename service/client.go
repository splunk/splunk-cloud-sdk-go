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

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

// Declare constants for service package
const (
	AuthorizationType = "Bearer"
)

// A Client is used to communicate with service endpoints
type Client struct {
	// config
	config *Config
	// HTTP Client used to interact with endpoints
	httpClient *http.Client
	// SearchService talks to the SSC search service
	SearchService *SearchService
	// CatalogService talks to the SSC catalog service
	CatalogService *CatalogService
	// IngestService talks to the SSC ingest service
	IngestService *IngestService
	// IdentityService talks to SSC IAC service
	IdentityService *IdentityService
	// KVStoreService talks to SSC kvstore service
	KVStoreService *KVStoreService
	// ActionService talks to SSC action service
	ActionService *ActionService
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

// ResponseHandler defines the interface for implementing custom response
// handling logic
type ResponseHandler interface {
	HandleResponse(client *Client, request *Request, response *http.Response) (*http.Response, error)
}

// Config is used to set the client specific attributes
type Config struct {
	// Authorization token
	Token string
	// Url string
	URL string
	// TenantID
	TenantID string
	// Timeout
	Timeout time.Duration
	// ResponseHandler to support additional response handling logic
	ResponseHandler ResponseHandler
}

// service provides the interface between client and services
type service struct {
	client *Client
}

// NewRequest creates a new HTTP Request and set proper header
func (c *Client) NewRequest(httpMethod, url string, body io.Reader) (*Request, error) {
	request, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}
	if len(c.config.Token) > 0 {
		request.Header.Set("Authorization", fmt.Sprintf("%s %s", AuthorizationType, c.config.Token))
	}
	request.Header.Set("Content-Type", "application/json")
	retryRequest := &Request{request, 0, make(map[string]uint)}
	return retryRequest, nil
}

// BuildURL creates full SSC URL with the client cached tenantID
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

// BuildURLWithTenantID creates full SSC URL with tenantID
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
		// TODO: This could be extended to include SSC error fields in addition
		// to response code
		code := fmt.Sprintf("%d", response.StatusCode)
		if _, ok := req.NumErrorsByType[code]; ok {
			req.NumErrorsByType[code]++
		} else {
			req.NumErrorsByType[code] = 1
		}
	}
	if c.config.ResponseHandler != nil {
		response, err = c.config.ResponseHandler.HandleResponse(c, req, response)
	}
	return response, err
}

// Get implements HTTP Get call
func (c *Client) Get(getURL url.URL) (*http.Response, error) {
	return c.DoRequest(http.MethodGet, getURL, nil)
}

// Post implements HTTP POST call
func (c *Client) Post(postURL url.URL, body interface{}) (*http.Response, error) {
	return c.DoRequest(http.MethodPost, postURL, body)
}

// Put implements HTTP PUT call
func (c *Client) Put(putURL url.URL, body interface{}) (*http.Response, error) {
	return c.DoRequest(http.MethodPut, putURL, body)
}

// Delete implements HTTP DELETE call
func (c *Client) Delete(deleteURL url.URL) (*http.Response, error) {
	return c.DoRequest(http.MethodDelete, deleteURL, nil)
}

// DeleteWithBody implements HTTP DELETE call with a request body
// RFC2616 does not explicitly forbid it but in practice some versions of server implementations (tomcat,
// netty etc) ignore bodies in DELETE requests
func (c *Client) DeleteWithBody(deleteURL url.URL, body interface{}) (*http.Response, error) {
	return c.DoRequest(http.MethodDelete, deleteURL, body)
}

// Patch implements HTTP Patch call
func (c *Client) Patch(patchURL url.URL, body interface{}) (*http.Response, error) {
	return c.DoRequest(http.MethodPatch, patchURL, body)
}

// DoRequest creates and execute a new request
func (c *Client) DoRequest(method string, requestURL url.URL, body interface{}) (*http.Response, error) {
	var buffer *bytes.Buffer
	if contentBytes, ok := body.([]byte); ok {
		buffer = bytes.NewBuffer(contentBytes)
	} else {
		if content, err := json.Marshal(body); err == nil {
			buffer = bytes.NewBuffer(content)
		} else {
			return nil, err
		}
	}
	request, err := c.NewRequest(method, requestURL.String(), buffer)
	if err != nil {
		return nil, err
	}
	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	return util.ParseHTTPStatusCodeInResponse(response)
}

// UpdateToken updates the authorization token
func (c *Client) UpdateToken(token string) {
	c.config.Token = token
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
	if config.TenantID == "" || config.Token == "" || config.URL == "" {
		return nil, errors.New("at least one of tenantID, token, or url must be set")
	}

	c := &Client{config: config, httpClient: &http.Client{Timeout: config.Timeout}}
	c.SearchService = &SearchService{client: c}
	c.CatalogService = &CatalogService{client: c}
	c.IdentityService = &IdentityService{client: c}
	c.IngestService = &IngestService{client: c}
	c.KVStoreService = &KVStoreService{client: c}
	c.ActionService = &ActionService{client: c}
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
	}

	return batchEventsSender, nil
}

// NewBatchEventsSender used to initialize dependencies and set values
func (c *Client) NewBatchEventsSender(batchSize int, interval int64) (*BatchEventsSender, error) {
	return c.NewBatchEventsSenderWithMaxAllowedError(batchSize, interval, 1)
}
