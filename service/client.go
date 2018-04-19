/*
Package service implements a service client which is used to communicate
with Search Service endpoints
*/
package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
	"fmt"

	"github.com/splunk/ssc-client-go/util"
)

// Declare constants for service package
const (
	MethodGet         = "GET"
	MethodPost        = "POST"
	MethodPut         = "PUT"
	MethodPatch       = "PATCH"
	MethodDelete      = "DELETE"
	AuthorizationType = "Bearer"
)

// A Client is used to communicate with service endpoints
type Client struct {
	// Authorization token
	token string
	// Url string
	URL url.URL
	// HTTP Client used to interact with endpoints
	httpClient *http.Client
	// Services designed to talk to different parts of Splunk
	SearchService *SearchService
	// CatalogService is to talk to catalog service of Splunk
	CatalogService *CatalogService
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

// BuildURL creates full Splunk URL
func (c *Client) BuildURL(urlPathParts ...string) url.URL {
	var buildPath = ""
	for _, pathPart := range urlPathParts {
		buildPath = path.Join(buildPath, url.PathEscape(pathPart))
	}

	return url.URL{
		Scheme: c.URL.Scheme,
		Host:   c.URL.Host,
		Path:   buildPath,
	}
}

// Do sends out request and returns HTTP response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Get implements HTTP Get call
func (c *Client) Get(getURL url.URL) (*http.Response, error) {
	return c.DoRequest(MethodGet, getURL, nil)
}

// Post implements HTTP POST call
func (c *Client) Post(postURL url.URL, body interface{}) (*http.Response, error) {
	return c.DoRequest(MethodPost, postURL, body)
}

// Put implements HTTP PUT call
func (c *Client) Put(putURL url.URL, body interface{}) (*http.Response, error) {
	return c.DoRequest(MethodPut, putURL, body)
}

// Delete implements HTTP DELETE call
func (c *Client) Delete(deleteURL url.URL) (*http.Response, error) {
	return c.DoRequest(MethodDelete, deleteURL, nil)
}

// Patch implements HTTP Patch call
func (c *Client) Patch(patchURL url.URL, body interface{}) (*http.Response, error) {
	return c.DoRequest(MethodPatch, patchURL, body)
}

// DoRequest creates and execute a new request
func (c *Client) DoRequest(method string, requestURL url.URL, body interface{}) (*http.Response, error) {

	content, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := c.NewRequest(method, requestURL.String(), bytes.NewBuffer(content))
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
func NewClient(token, URL string, timeout time.Duration, skipValidateTLS bool) *Client {
	httpClient := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: skipValidateTLS},
		},
	}
	parsed, _ := url.Parse(URL)
	c := &Client{token: token, URL: *parsed, httpClient: httpClient}
	c.SearchService = &SearchService{client: c}
	c.CatalogService = &CatalogService{client: c}
	return c
}
