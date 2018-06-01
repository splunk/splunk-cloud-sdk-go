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
	// TenantID used for ssc service
	TenantID string
	// Authorization token
	token string
	// Url string
	URL url.URL
	// HTTP Client used to interact with endpoints
	httpClient *http.Client
	// SearchService talks to the SSC search service
	SearchService *SearchService
	// CatalogService talks to the SSC catalog service
	CatalogService *CatalogService
	// HecService talks to the SSC hec service
	HecService *HecService
	// IdentityService talks to the IAC service
	IdentityService *IdentityService
}

// service provides the interface between client and services
type service struct {
	client *Client
}

// NewRequest creates a new HTTP Request and set proper header
func (c *Client) NewRequest(httpMethod, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}
	if len(c.token) > 0 {
		request.Header.Set("Authorization", fmt.Sprintf("%s %s", AuthorizationType, c.token))
	}
	request.Header.Set("Content-Type", "application/json")
	return request, nil
}

// BuildURL creates full SSC URL with the client cached tenantID
func (c *Client) BuildURL(urlPathParts ...string) (url.URL, error) {
	var buildPath = ""
	for _, pathPart := range urlPathParts {
		buildPath = path.Join(buildPath, url.PathEscape(pathPart))
	}

	var u url.URL
	if len(c.TenantID) == 0 {
		return u, errors.New("A non-empty tenant ID must be set on client")
	}

	u = url.URL{
		Scheme: c.URL.Scheme,
		Host:   c.URL.Host,
		Path:   path.Join(c.TenantID, buildPath),
	}
	return u, nil
}

// BuildURLWithTenantID creates full SSC URL with tenantID
func (c *Client) BuildURLWithTenantID(tenantID string, urlPathParts ...string) (url.URL, error) {
	var buildPath = ""
	for _, pathPart := range urlPathParts {
		buildPath = path.Join(buildPath, url.PathEscape(pathPart))
	}

	var u url.URL
	if len(tenantID) == 0 {
		return u, errors.New("A non-empty tenant ID must be passed in for BuildURLWithTenantID")
	}

	u = url.URL{
		Scheme: c.URL.Scheme,
		Host:   c.URL.Host,
		Path:   path.Join(tenantID, buildPath),
	}
	return u, nil
}

// Do sends out request and returns HTTP response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
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
	c.token = token
}

// NewClient creates a Client with custom values passed in
func NewClient(tenantID, token, URL string, timeout time.Duration) *Client {
	httpClient := &http.Client{
		Timeout: timeout,
	}
	parsed, _ := url.Parse(URL)
	c := &Client{TenantID: tenantID, token: token, URL: *parsed, httpClient: httpClient}
	c.SearchService = &SearchService{client: c}
	c.CatalogService = &CatalogService{client: c}
	c.IdentityService = &IdentityService{client: c}
	c.HecService = &HecService{client: c}
	return c
}

// NewBatchEventsSender used to initialize dependencies and set values
func (c *Client) NewBatchEventsSender(batchSize int, interval int64) (*BatchEventsSender, error) {
	// Rather than return a super general error for both it will block on batchSize first
	if batchSize == 0 {
		return nil, errors.New("batchSize cannot be 0")
	}
	if interval == 0 {
		return nil, errors.New("interval cannot be 0")
	}

	eventsChan := make(chan model.HecEvent, batchSize)
	eventsQueue := make([]model.HecEvent, 0, batchSize)
	quit := make(chan struct{}, 1)
	ticker := model.NewTicker(time.Duration(interval) * time.Millisecond)
	var wg sync.WaitGroup

	batchEventsSender := &BatchEventsSender{
		BatchSize:    batchSize,
		EventsChan:   eventsChan,
		EventsQueue:  eventsQueue,
		EventService: c.HecService,
		QuitChan:     quit,
		HecTicker:    ticker,
		WaitGroup:    &wg,
	}

	return batchEventsSender, nil
}
