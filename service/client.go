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
)

// Declare constants for service package
const (
	AuthorizationType = "Bearer"
	API               = "api"
)

const RefreshToken = "tobefilled"

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
	// 1. Check for 401
	// 2. Attempt to get new token from refresh via Okta
	//   a. 400 from Okta - return nil, err
	//   b. 200 from Okta with new token
	// 3. Retry request
	//   a. 401 (again) - return nil, err
	//   b. 200 - all good
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
		Path:   path.Join(API, c.TenantID, buildPath),
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
		Path:   path.Join(API, tenantID, buildPath),
	}
	return u, nil
}

// Do sends out request and returns HTTP response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	response, err := c.httpClient.Do(req)
	//if 401 ->
	fmt.Print(c.token)
	//Refresh access token with refresh token
	c.RefreshToken(RefreshToken, "0oa12zcrqk8jXGIDZ2p7")
	fmt.Println(c.token)
	//retry request with new access token
	//response, err := c.httpClient.Do(req)
	//if 401 return error
	//else return response, nil

	if err != nil {
		return nil, err
	}
	return response, nil
}

type refreshDat struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpireIn int `json:"expires_in"`
	Scope string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

func (c *Client) RefreshToken(refreshToken string, clientId string) (*http.Response, error){

	client := http.Client{}
	var urlPath= ""
	urlPath = path.Join("splunk-ciam.okta.com/oauth2/default/v1/token")

	var u *url.URL

	u, _ = url.Parse(urlPath)
	u.Path = urlPath
	u.Scheme = "https"

	data := url.Values{}
	data.Set("refresh_token",refreshToken)
	data.Add("grant_type", "refresh_token")
	data.Add("client_id", clientId)

	u.RawQuery = data.Encode()

	urlStr := fmt.Sprintf("%v", u)

	req, err := http.NewRequest("POST", urlStr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(req)

	var accessToken string

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err == nil {
			s, _ := parseRefreshData([]byte(body))
			accessToken = s.AccessToken
		}
	}
	c.UpdateToken(accessToken)

	return response, err
}


func parseRefreshData(body []byte) (*refreshDat, error) {
	var refreshJson = new(refreshDat)
	err := json.Unmarshal(body, &refreshJson)
	if err != nil {
		fmt.Println("error:", err)
	}
	return refreshJson, err
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
	var err error
	if request, err := c.NewRequest(method, requestURL.String(), buffer); err == nil {
		if response, err := c.Do(request); err == nil {
			return util.ParseHTTPStatusCodeInResponse(response)
		}
	}
	return nil, err
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
