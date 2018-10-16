// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

/*
Package services implements a service client which is used to communicate
with Splunk Cloud endpoints, each service being split into its own package.
*/
package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

// Declare constants for service package
const (
	AuthorizationType = "Bearer"
)

// A BaseClient for communicating with Splunk Cloud
type BaseClient struct {
	// defaultTenant is the Splunk Cloud tenant to use to form requests
	defaultTenant string
	// host is the Splunk Cloud host or host:port used to form requests, `"api.splunkbeta.com"` by default
	host string
	// scheme is the HTTP scheme used to form requests, `"https"` by default
	scheme string
	// tokenContext is the access token to include in `"Authorization: Bearer"` headers and related context information
	tokenContext *idp.Context
	// HTTP Client used to interact with endpoints
	httpClient *http.Client
	// responseHandlers is a slice of handlers to call after a response has been received in the client
	responseHandlers []ResponseHandler
}

// Request extends net/http.Request to track number of total attempts and error
// counts by type of error
type Request struct {
	*http.Request
	NumAttempts     uint
	NumErrorsByType map[string]uint
}

// Logger define logger interface for roundtripper
type Logger interface {
	Debug(string)
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
	// Tenant is the default Tenant used to form requests
	Tenant string
	// Host is the (optional) default host or host:port used to form requests, `"api.splunkbeta.com"` by default
	Host string
	// Scheme is the (optional) default HTTP Scheme used to form requests, `"https"` by default
	Scheme string
	// Timeout is the (optional) default request-level timeout to use, 5 seconds by default
	Timeout time.Duration
	// ResponseHandlers is an (optional) slice of handlers to call after a response has been received in the client
	ResponseHandlers []ResponseHandler
	// RoundTripper
	RoundTripper http.RoundTripper
}

// RequestParams contains all the optional request URL parameters
type RequestParams struct {
	// Method is the HTTP method of the request
	Method string
	// URL is the URL of the request
	URL url.URL
	// Body is the body of the request
	Body interface{}
	// Headers are additional headers to add to the request
	Headers map[string]string
}

// BaseService provides the interface between client and services
type BaseService struct {
	Client *BaseClient
}

// NewRequest creates a new HTTP Request and set proper header
func (c *BaseClient) NewRequest(httpMethod, url string, body io.Reader, headers map[string]string) (*Request, error) {
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

// BuildURL creates full Splunk Cloud URL using the client's defaultTenant
func (c *BaseClient) BuildURL(queryValues url.Values, urlPathParts ...string) (url.URL, error) {
	return c.BuildURLWithTenant(c.defaultTenant, queryValues, urlPathParts...)
}

// BuildURLWithTenant creates full Splunk Cloud URL with tenant
func (c *BaseClient) BuildURLWithTenant(tenant string, queryValues url.Values, urlPathParts ...string) (url.URL, error) {
	var u url.URL
	if len(tenant) == 0 {
		return u, errors.New("a non-empty tenant must be specified")
	}
	var buildPath = ""
	for _, pathPart := range urlPathParts {
		buildPath = path.Join(buildPath, url.PathEscape(pathPart))
	}
	if queryValues == nil {
		queryValues = url.Values{}
	}
	u = url.URL{
		Scheme:   c.scheme,
		Host:     c.host,
		Path:     path.Join(tenant, buildPath),
		RawQuery: queryValues.Encode(),
	}
	return u, nil
}

// Do sends out request and returns HTTP response
func (c *BaseClient) Do(req *Request) (*http.Response, error) {
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
func (c *BaseClient) Get(requestParams RequestParams) (*http.Response, error) {
	requestParams.Method = http.MethodGet
	return c.DoRequest(requestParams)
}

// Post implements HTTP POST call
func (c *BaseClient) Post(requestParams RequestParams) (*http.Response, error) {
	requestParams.Method = http.MethodPost
	return c.DoRequest(requestParams)
}

// Put implements HTTP PUT call
func (c *BaseClient) Put(requestParams RequestParams) (*http.Response, error) {
	requestParams.Method = http.MethodPut
	return c.DoRequest(requestParams)
}

// Delete implements HTTP DELETE call
// RFC2616 does not explicitly forbid it but in practice some versions of server implementations (tomcat,
// netty etc) ignore bodies in DELETE requests
func (c *BaseClient) Delete(requestParams RequestParams) (*http.Response, error) {
	requestParams.Method = http.MethodDelete
	return c.DoRequest(requestParams)
}

// Patch implements HTTP Patch call
func (c *BaseClient) Patch(requestParams RequestParams) (*http.Response, error) {
	requestParams.Method = http.MethodPatch
	return c.DoRequest(requestParams)
}

// DoRequest creates and execute a new request
func (c *BaseClient) DoRequest(requestParams RequestParams) (*http.Response, error) {
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
func (c *BaseClient) UpdateTokenContext(ctx *idp.Context) {
	c.tokenContext = ctx
}

// SetDefaultTenant updates the tenant used to form most request URIs
func (c *BaseClient) SetDefaultTenant(tenant string) {
	c.defaultTenant = tenant
}

// GetURL returns the Splunk Cloud scheme/host formed as URL
func (c *BaseClient) GetURL() *url.URL {
	return &url.URL{
		Scheme: c.scheme,
		Host:   c.host,
	}
}

// NewClient creates a Client with config values passed in
func NewClient(config *Config) (*BaseClient, error) {
	host := "api.splunkbeta.com"
	if config.Host != "" {
		host = config.Host
	}
	scheme := "https"
	if config.Scheme != "" {
		scheme = config.Scheme
	}
	timeout := 5 * time.Second
	if config.Timeout != 0 {
		timeout = config.Timeout
	}

	// Enforce that exactly one of TokenRetriever or Token must be specified
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
	c := &BaseClient{
		host:             host,
		scheme:           scheme,
		defaultTenant:    config.Tenant,
		httpClient:       &http.Client{Timeout: timeout},
		tokenContext:     ctx,
		responseHandlers: handlers,
	}

	if config.RoundTripper != nil {
		c.httpClient = &http.Client{Timeout: timeout, Transport: config.RoundTripper}
	}


	return c, nil
}


// SdkTransport is to define a transport RoundTripper with user-defined logger
type SdkTransport struct {
	transport http.RoundTripper
	logger    Logger
}

// RoundTrip implements the RoundTripper interface
func (st *SdkTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		return nil, err
	}

	st.logger.Debug(fmt.Sprintf("===Request:\n%s\n", string(requestDump)))

	response, err := st.transport.RoundTrip(request)

	responseDump, err := httputil.DumpResponse(response, true)
	if err != nil {
		return response, err
	}

	st.logger.Debug(fmt.Sprintf("===Response:\n%s\n", string(responseDump)))

	return response, err
}

// CreateRoundTripperWithLogger Creates a RoundTripper with user defined logger
func CreateRoundTripperWithLogger(logger Logger) *SdkTransport {
	return &SdkTransport{transport: &http.Transport{}, logger: logger}
}
