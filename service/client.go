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
	"io/ioutil"
	"os"
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
	// IdentityService talks to SSC IAC service
	IdentityService *IdentityService
	// KVStoreService talks to SSC kvstore service
	KVStoreService *KVStoreService
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
func (c *Client) BuildURL(queryValues url.Values, urlPathParts ...string) (url.URL, error) {
	var buildPath = ""
	for _, pathPart := range urlPathParts {
		buildPath = path.Join(buildPath, url.PathEscape(pathPart))
	}
	if queryValues == nil {
		queryValues = url.Values{}
	}
	var u url.URL
	if len(c.TenantID) == 0 {
		return u, errors.New("A non-empty tenant ID must be set on client")
	}
	u = url.URL{
		Scheme:   c.URL.Scheme,
		Host:     c.URL.Host,
		Path:     path.Join(c.TenantID, buildPath),
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

	u = url.URL{
		Scheme:   c.URL.Scheme,
		Host:     c.URL.Host,
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
	body := req.Body

	//Refresh access token with refresh token
	var accessToken string
	var err error
	accessToken, err = c.GetNewAccessToken()
	if err != nil || len(accessToken) == 0 {
		return nil, err
	}
	// Update the client with the newly obtained access token
	c.UpdateToken(accessToken)
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
func NewClient(tenantID, token, URL string, timeout time.Duration) (*Client, error) {
	if tenantID == "" || token == "" || URL == "" {
		return nil, errors.New("tenantID or token or url can't be empty")
	}

	httpClient := &http.Client{
		Timeout: timeout,
	}
	parsed, err := url.Parse(URL)
	if err != nil {
		return nil, errors.New("Url is not correct")
	}

	c := &Client{TenantID: tenantID, token: token, URL: *parsed, httpClient: httpClient}
	c.SearchService = &SearchService{client: c}
	c.CatalogService = &CatalogService{client: c}
	c.IdentityService = &IdentityService{client: c}
	c.HecService = &HecService{client: c}
	c.KVStoreService = &KVStoreService{client: c}
	return c, nil
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
