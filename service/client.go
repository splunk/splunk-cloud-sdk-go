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

	"io/ioutil"
	"os"

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

// RefreshToken - RefreshToken to refresh the bearer token if expired
var RefreshToken = os.Getenv("REFRESH_TOKEN")

// RefreshTokenEndpoint - Okta end point to hit to retrieve the bearer token
var RefreshTokenEndpoint = os.Getenv("REFRESH_TOKEN_ENDPOINT")

// ClientID - Okta app Client Id for SDK
var ClientID = os.Getenv("CLIENT_ID")

// service provides the interface between client and services
type service struct {
	client *Client
}

// NewRequest creates a new HTTP Request and set proper header
func (c *Client) NewRequest(httpMethod, url string, body io.Reader, headers map[string]string) (*http.Request, error) {
	request, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}
	if len(c.config.Token) > 0 {
		request.Header.Set("Authorization", fmt.Sprintf("%s %s", AuthorizationType, c.config.Token))
	}
	request.Header.Set("Content-Type", "application/json")
	if len(headers) != 0 {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}
	return request, nil
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
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	response, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	//If bearer token results in a 401 and there is a refresh token available, get a new bearer token
	if response != nil && response.StatusCode == 401 && len(RefreshToken) != 0 {
		response, err = c.onUnauthorizedRequest(req)
		return response, err
	}
	return response, err
}

func (c *Client) onUnauthorizedRequest(req *http.Request) (*http.Response, error) {
	// refresh and retry request here
	httpMethod := req.Method

	//Refresh access token with refresh token
	var accessToken string
	var err error
	accessToken, err = c.GetNewAccessToken()
	if err != nil || len(accessToken) == 0 {
		return nil, err
	}
	// Update the client with the newly obtained access token
	c.UpdateToken(accessToken)
	body, err := req.GetBody()
	request, err := http.NewRequest(httpMethod, req.URL.String(), body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", AuthorizationType, accessToken))
	request.Header.Set("Content-Type", "application/json")

	// retry request with new access token
	response, err := c.httpClient.Do(request)

	return response, err
}

type refreshData struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpireIn     int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

// GetNewAccessToken gets a new bearer token from the okta token endpoint given the refresh token
func (c *Client) GetNewAccessToken() (string, error) {
	var accessToken = ""
	client := http.Client{}
	var urlPath = ""
	urlPath = path.Join(RefreshTokenEndpoint)

	data := url.Values{}
	data.Set("refresh_token", RefreshToken)
	data.Add("grant_type", "refresh_token")
	data.Add("client_id", ClientID)
	data.Add("scope", "openid email profile")

	tokenURL := url.URL{
		Scheme:   "https",
		Path:     urlPath,
		RawQuery: data.Encode(),
	}

	req, err := http.NewRequest("POST", tokenURL.String(), nil)
	if err != nil {
		return accessToken, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(req)

	if response != nil && response.StatusCode == 200 {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)

		if err == nil {
			s, err := parseRefreshData([]byte(body))
			if err != nil {
				return accessToken, err
			}
			accessToken = s.AccessToken
			return accessToken, err
		}

		return accessToken, err
	}
	return accessToken, err
}

func parseRefreshData(body []byte) (*refreshData, error) {
	var refreshJSON = new(refreshData)
	err := json.Unmarshal(body, &refreshJSON)

	return refreshJSON, err
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