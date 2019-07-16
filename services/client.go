/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

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
	"net/url"
	"path"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

//go:generate go run ../util/gen_interface.go -svc=action -p=action -sf=service_generated.go -sf=service_sdk.go
//go:generate go run ../util/gen_interface.go -svc=appregistry -p=appregistry -sf=service_generated.go
//go:generate go run ../util/gen_interface.go -svc=catalog -p=catalog -sf=service_generated.go
//go:generate go run ../util/gen_interface.go -svc=forwarders -p=forwarders -sf=service_generated.go
//go:generate go run ../util/gen_interface.go -svc=identity -p=identity -sf=service_generated.go
//go:generate go run ../util/gen_interface.go -svc=ingest -p=ingest -sf=service.go -sf=service_generated.go
//go:generate go run ../util/gen_interface.go -svc=kvstore -p=kvstore -sf=service_generated.go
//go:generate go run ../util/gen_interface.go -svc=ml -p=ml -sf=service_generated.go
//go:generate go run ../util/gen_interface.go -svc=search -p=search -sf=service.go -sf=service_generated.go
//go:generate go run ../util/gen_interface.go -svc=streams -p=streams -sf=service_generated.go
//go:generate go run ../util/gen_interface.go -svc=provisioner -p=provisioner -sf=service_generated.go

// Declare constants for service package
const (
	AuthorizationType = "Bearer"
)

// A BaseClient for communicating with Splunk Cloud
type BaseClient struct {
	// defaultTenant is the Splunk Cloud tenant to use to form requests
	defaultTenant string
	// rootDomain is the Splunk Cloud rootDomain or rootDomain:port used to form requests, `"scp.splunk.com"` by default.
	// Note that requests would be made to `"<scheme>://api.<rootDomain>/..."` for most services by default.
	rootDomain string
	// overrideHost if set would override rootDomain and service cluster settings when forming the host such that requests would be made to `"<scheme>://<overrideHost>/..."` for all services.
	overrideHost string
	// scheme is the HTTP scheme used to form requests, `"https"` by default
	scheme string
	// tokenContext is the access token to include in `"Authorization: Bearer"` headers and related context information
	tokenContext *idp.Context
	// HTTP Client used to interact with endpoints
	httpClient *http.Client
	// responseHandlers is a slice of handlers to call after a response has been received in the client
	responseHandlers []ResponseHandler
	// tokenRetriever to gather access tokens to be sent in the Authorization: Bearer header on client initialization and upon encountering an expired token
	tokenRetriever idp.TokenRetriever
	// tokenExpireWindow is the (optional) window within which a new token gets retreieved before the existing token expires. Default to 1 minute
	tokenExpireWindow time.Duration
	// tokenMux is used to lock resource when a new token is being fetched
	tokenMux sync.Mutex
}

// Request extends net/http.Request to track number of total attempts and error
// counts by type of error
type Request struct {
	*http.Request
	NumAttempts     uint
	NumErrorsByType map[string]uint
}

//ConfigurableRetryConfig that will accept a user configurable RetryNumber and Interval between retries
type ConfigurableRetryConfig struct {
	RetryNum uint
	Interval int
}

//DefaultRetryConfig that will use a default RetryNumber and a default Interval between retries
type DefaultRetryConfig struct {
}

//RetryStrategyConfig to be specified while creating a NewClient
type RetryStrategyConfig struct {
	DefaultRetryConfig      *DefaultRetryConfig
	ConfigurableRetryConfig *ConfigurableRetryConfig
}

// GetNumErrorsByResponseCode returns number of attempts for a given response code >= 400
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
	// Host is the (optional) default host or host:port used to form requests, `"scp.splunk.com"` by default.
	// NOTE: This is really a root domain, most requests will be formed using `<config.Scheme>://api.<config.Host>/<tenant>/<service>/<version>/...` where `api` could vary by service
	Host string
	// OverrideHost if set would override the Splunk Cloud root domain (`"scp.splunk.com"` by default) and service settings when forming the host such that requests would be made to `"<scheme>://<overrideHost>/..."` for all services.
	// NOTE: Providing a Host and OverrideHost is not valid.
	OverrideHost string
	// Scheme is the (optional) default HTTP Scheme used to form requests, `"https"` by default
	Scheme string
	// Timeout is the (optional) default request-level timeout to use, 5 seconds by default
	Timeout time.Duration
	// ResponseHandlers is an (optional) slice of handlers to call after a response has been received in the client
	ResponseHandlers []ResponseHandler
	//RetryRequests Knob that will turn on and off retrying incoming service requests when they result in the service returning a 429 TooManyRequests Error
	RetryRequests bool
	//RetryStrategyConfig
	RetryConfig RetryStrategyConfig
	// RoundTripper
	RoundTripper http.RoundTripper
	// TokenExpireWindow is the (optional) window within which a new token gets retreieved before the existing token expires. Default to 1 minute
	TokenExpireWindow time.Duration
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

// BuildHost returns host including serviceCluster
func (c *BaseClient) BuildHost(serviceCluster string) string {
	// If overrideHost is specified, always use that
	if c.overrideHost != "" {
		return c.overrideHost
	}
	// Otherwise form using <serviceCluster>.<rootDomain>
	if serviceCluster != "" {
		return fmt.Sprintf("%s.%s", serviceCluster, c.rootDomain)
	}
	return fmt.Sprintf("api.%s", c.rootDomain)
}

// BuildURL creates full Splunk Cloud URL using the client's defaultTenant
func (c *BaseClient) BuildURL(queryValues url.Values, serviceCluster string, urlPathParts ...string) (url.URL, error) {
	return c.BuildURLWithTenant(c.defaultTenant, queryValues, serviceCluster, urlPathParts...)
}

// BuildURLWithTenant creates full Splunk Cloud URL with tenant
func (c *BaseClient) BuildURLWithTenant(tenant string, queryValues url.Values, serviceCluster string, urlPathParts ...string) (url.URL, error) {
	var u url.URL
	if len(tenant) == 0 {
		return u, errors.New("a non-empty tenant must be specified")
	}
	if queryValues == nil {
		queryValues = url.Values{}
	}
	host := c.BuildHost(serviceCluster)
	pathWithTenant := path.Join(append([]string{tenant}, urlPathParts...)...)

	u = url.URL{
		Scheme:   c.scheme,
		Host:     host,
		Path:     pathWithTenant,
		RawQuery: queryValues.Encode(),
	}
	return u, nil
}

// BuildURLFromPathParams creates full Splunk Cloud URL from path template and path params
func (c *BaseClient) BuildURLFromPathParams(queryValues url.Values, serviceCluster string, templ string, pathParams interface{}) (url.URL, error) {
	var u url.URL
	t, err := template.New("path").Parse(templ)
	if err != nil {
		return u, err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, pathParams)
	if err != nil {
		return u, err
	}
	path := buf.String()
	if !strings.HasPrefix(path, "/system/") {
		// for non-system-namespace endpoints, add tenant namespace
		path = "/" + c.defaultTenant + path
	}
	if queryValues == nil {
		queryValues = url.Values{}
	}
	host := c.BuildHost(serviceCluster)
	u = url.URL{
		Scheme:   c.scheme,
		Host:     host,
		Path:     path,
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
	var request *Request
	var err error
	now := time.Now().Add(c.tokenExpireWindow)
	curEpoch := now.Unix()
	// renew token if it's about to expire
	if curEpoch >= c.tokenContext.StartTime+int64(c.tokenContext.ExpiresIn) {
		c.tokenMux.Lock()
		ctx, err := c.tokenRetriever.GetTokenContext()
		if err != nil {
			c.tokenMux.Unlock()
			return nil, err
		}
		// Update the client such that future requests will use the new access token and retain context information
		c.UpdateTokenContext(ctx)
		c.tokenMux.Unlock()
	}
	if requestParams.Body != nil {
		var buffer *bytes.Buffer

		if contentBytes, ok := requestParams.Body.([]byte); ok {
			buffer = bytes.NewBuffer(contentBytes)
		} else {
			bodyMarshaler, ok := requestParams.Body.(util.MethodMarshaler)
			var marshalErr error
			var content []byte
			if ok {
				content, marshalErr = bodyMarshaler.MarshalJSONByMethod(requestParams.Method)
			} else {
				content, marshalErr = json.Marshal(requestParams.Body)
			}
			if marshalErr != nil {
				return nil, marshalErr
			}
			buffer = bytes.NewBuffer(content)
		}
		request, err = c.NewRequest(requestParams.Method, requestParams.URL.String(), buffer, requestParams.Headers)
		if err != nil {
			return nil, err
		}

	} else {
		request, err = c.NewRequest(requestParams.Method, requestParams.URL.String(), nil, requestParams.Headers)
		if err != nil {
			return nil, err
		}
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

// GetDefaultTenant returns the tenant used to form most request URIs
func (c *BaseClient) GetDefaultTenant() string {
	return c.defaultTenant
}

// SetDefaultTenant updates the tenant used to form most request URIs
func (c *BaseClient) SetDefaultTenant(tenant string) {
	c.defaultTenant = tenant
}

// SetOverrideHost updates the host to force all requests to be made to `<scheme>://<overrideHost>/...` ignoring Config.Host and serviceCluster values
func (c *BaseClient) SetOverrideHost(host string) {
	c.overrideHost = host
}

// GetURL returns the Splunk Cloud scheme/host formed as URL
func (c *BaseClient) GetURL(serviceCluster string) *url.URL {
	host := c.BuildHost(serviceCluster)
	return &url.URL{
		Scheme: c.scheme,
		Host:   host,
	}
}

// NewClient creates a Client with config values passed in
func NewClient(config *Config) (*BaseClient, error) {
	// Enforce that at most one of Host or OverrideHost may be specified, not both
	if config.Host != "" && config.OverrideHost != "" {
		return nil, errors.New("either config.Host or config.OverrideHost may be set, setting both is invalid")
	}
	rootDomain := "scp.splunk.com"
	if config.Host != "" {
		rootDomain = config.Host
	}
	overrideHost := ""
	if config.OverrideHost != "" {
		overrideHost = config.OverrideHost
	}

	scheme := "https"
	if config.Scheme != "" {
		scheme = config.Scheme
	}
	timeout := 5 * time.Second
	if config.Timeout != 0 {
		timeout = config.Timeout
	}
	tokenExpireWindow := 1 * time.Minute
	if config.TokenExpireWindow != 0 {
		tokenExpireWindow = config.TokenExpireWindow
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
	}
	if config.RetryRequests {
		//if knob to RetryRequests is on, Retry Response Handler is created to handle the 429 response and retry the incoming requests that are being throttled based on the retry strategy specified in the config
		if config.RetryConfig.ConfigurableRetryConfig == nil {
			defaultStrategyHandler := DefaultRetryResponseHandler{DefaultRetryConfig{}}
			handlers = append([]ResponseHandler{ResponseHandler(defaultStrategyHandler)}, config.ResponseHandlers...)
		} else {
			configStrategyHandler := ConfigurableRetryResponseHandler{ConfigurableRetryConfig{config.RetryConfig.ConfigurableRetryConfig.RetryNum, config.RetryConfig.ConfigurableRetryConfig.Interval}}
			handlers = append([]ResponseHandler{ResponseHandler(configStrategyHandler)}, config.ResponseHandlers...)
		}
	}
	// Start by retrieving the access token
	ctx, err := config.TokenRetriever.GetTokenContext()
	if err != nil {
		return nil, fmt.Errorf("service.NewClient: error retrieving token: %s", err)
	}

	// Finally, initialize the Client
	c := &BaseClient{
		rootDomain:        rootDomain,
		overrideHost:      overrideHost,
		scheme:            scheme,
		defaultTenant:     config.Tenant,
		httpClient:        &http.Client{Timeout: timeout},
		tokenRetriever:    config.TokenRetriever,
		tokenContext:      ctx,
		responseHandlers:  handlers,
		tokenExpireWindow: tokenExpireWindow,
	}

	if config.RoundTripper != nil {
		c.httpClient = &http.Client{Timeout: timeout, Transport: config.RoundTripper}
	}

	return c, nil
}
