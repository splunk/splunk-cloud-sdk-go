# services
--
    import "github.com/splunk/splunk-cloud-sdk-go/services"

Package services implements a service client which is used to communicate with
Splunk Cloud endpoints, each service being split into its own package.

## Usage

```go
const (
	AuthorizationType = "Bearer"
)
```
Declare constants for service package

```go
const (
	// DefaultMaxAuthnAttempts defines the maximum number of retries that will be performed for a request encountering an authentication issue
	DefaultMaxAuthnAttempts = 1
)
```

```go
const UserAgent = "client-go"
```
UserAgent SDK Client Identifier

```go
const Version = "1.0.0-beta.0"
```
Version the released version of the SDK

#### type AuthnResponseHandler

```go
type AuthnResponseHandler struct {
	TokenRetriever idp.TokenRetriever
}
```

AuthnResponseHandler handles logic for updating the client access token in
response to 401 errors

#### func (AuthnResponseHandler) HandleResponse

```go
func (rh AuthnResponseHandler) HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error)
```
HandleResponse will retry a request once after re-authenticating if a 401
response code is encountered

#### type BaseClient

```go
type BaseClient struct {
}
```

A BaseClient for communicating with Splunk Cloud

#### func  NewClient

```go
func NewClient(config *Config) (*BaseClient, error)
```
NewClient creates a Client with config values passed in

#### func (*BaseClient) BuildHost

```go
func (c *BaseClient) BuildHost(serviceCluster string) string
```
BuildHost returns host including serviceCluster

#### func (*BaseClient) BuildURL

```go
func (c *BaseClient) BuildURL(queryValues url.Values, serviceCluster string, urlPathParts ...string) (url.URL, error)
```
BuildURL creates full Splunk Cloud URL using the client's defaultTenant

#### func (*BaseClient) BuildURLFromPathParams

```go
func (c *BaseClient) BuildURLFromPathParams(queryValues url.Values, serviceCluster string, templ string, pathParams interface{}) (url.URL, error)
```
BuildURLFromPathParams creates full Splunk Cloud URL from path template and path
params

#### func (*BaseClient) BuildURLWithTenant

```go
func (c *BaseClient) BuildURLWithTenant(tenant string, queryValues url.Values, serviceCluster string, urlPathParts ...string) (url.URL, error)
```
BuildURLWithTenant creates full Splunk Cloud URL with tenant

#### func (*BaseClient) Delete

```go
func (c *BaseClient) Delete(requestParams RequestParams) (*http.Response, error)
```
Delete implements HTTP DELETE call RFC2616 does not explicitly forbid it but in
practice some versions of server implementations (tomcat, netty etc) ignore
bodies in DELETE requests

#### func (*BaseClient) Do

```go
func (c *BaseClient) Do(req *Request) (*http.Response, error)
```
Do sends out request and returns HTTP response

#### func (*BaseClient) DoRequest

```go
func (c *BaseClient) DoRequest(requestParams RequestParams) (*http.Response, error)
```
DoRequest creates and execute a new request

#### func (*BaseClient) Get

```go
func (c *BaseClient) Get(requestParams RequestParams) (*http.Response, error)
```
Get implements HTTP Get call

#### func (*BaseClient) GetDefaultTenant

```go
func (c *BaseClient) GetDefaultTenant() string
```
GetDefaultTenant returns the tenant used to form most request URIs

#### func (*BaseClient) GetURL

```go
func (c *BaseClient) GetURL(serviceCluster string) *url.URL
```
GetURL returns the Splunk Cloud scheme/host formed as URL

#### func (*BaseClient) NewRequest

```go
func (c *BaseClient) NewRequest(httpMethod, url string, body io.Reader, headers map[string]string) (*Request, error)
```
NewRequest creates a new HTTP Request and set proper header

#### func (*BaseClient) Patch

```go
func (c *BaseClient) Patch(requestParams RequestParams) (*http.Response, error)
```
Patch implements HTTP Patch call

#### func (*BaseClient) Post

```go
func (c *BaseClient) Post(requestParams RequestParams) (*http.Response, error)
```
Post implements HTTP POST call

#### func (*BaseClient) Put

```go
func (c *BaseClient) Put(requestParams RequestParams) (*http.Response, error)
```
Put implements HTTP PUT call

#### func (*BaseClient) SetDefaultTenant

```go
func (c *BaseClient) SetDefaultTenant(tenant string)
```
SetDefaultTenant updates the tenant used to form most request URIs

#### func (*BaseClient) SetOverrideHost

```go
func (c *BaseClient) SetOverrideHost(host string)
```
SetOverrideHost updates the host to force all requests to be made to
`<scheme>://<overrideHost>/...` ignoring Config.Host and serviceCluster values

#### func (*BaseClient) UpdateTokenContext

```go
func (c *BaseClient) UpdateTokenContext(ctx *idp.Context)
```
UpdateTokenContext the access token in the Authorization: Bearer header and
retains related context information

#### type BaseService

```go
type BaseService struct {
	Client *BaseClient
}
```

BaseService provides the interface between client and services

#### type Config

```go
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
```

Config is used to set the client specific attributes

#### type ConfigurableRetryConfig

```go
type ConfigurableRetryConfig struct {
	RetryNum uint
	Interval int
}
```

ConfigurableRetryConfig that will accept a user configurable RetryNumber and
Interval between retries

#### type ConfigurableRetryResponseHandler

```go
type ConfigurableRetryResponseHandler struct {
	ConfigurableRetryConfig ConfigurableRetryConfig
}
```

ConfigurableRetryResponseHandler handles logic for retrying requests with user
configurable settings for Retry number and interval

#### func (ConfigurableRetryResponseHandler) HandleResponse

```go
func (configRh ConfigurableRetryResponseHandler) HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error)
```
HandleResponse will retry a request once a 429 is encountered using a
Configurable exponential BackOff Retry Strategy

#### type DefaultRetryConfig

```go
type DefaultRetryConfig struct {
}
```

DefaultRetryConfig that will use a default RetryNumber and a default Interval
between retries

#### type DefaultRetryResponseHandler

```go
type DefaultRetryResponseHandler struct {
	DefaultRetryConfig DefaultRetryConfig
}
```

DefaultRetryResponseHandler handles logic for retrying requests with default
settings for Retry number and interval

#### func (DefaultRetryResponseHandler) HandleResponse

```go
func (defRh DefaultRetryResponseHandler) HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error)
```
HandleResponse will retry a request once a 429 is encountered using a Default
exponential BackOff Retry Strategy

#### type Request

```go
type Request struct {
	*http.Request
	NumAttempts     uint
	NumErrorsByType map[string]uint
}
```

Request extends net/http.Request to track number of total attempts and error
counts by type of error

#### func (*Request) GetNumErrorsByResponseCode

```go
func (r *Request) GetNumErrorsByResponseCode(respCode int) uint
```
GetNumErrorsByResponseCode returns number of attempts for a given response code
>= 400

#### func (*Request) UpdateToken

```go
func (r *Request) UpdateToken(accessToken string)
```
UpdateToken replaces the access token in the `Authorization: Bearer` header

#### type RequestParams

```go
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
```

RequestParams contains all the optional request URL parameters

#### type ResponseHandler

```go
type ResponseHandler interface {
	HandleResponse(client *BaseClient, request *Request, response *http.Response) (*http.Response, error)
}
```

ResponseHandler defines the interface for implementing custom response handling
logic

#### type RetryStrategyConfig

```go
type RetryStrategyConfig struct {
	DefaultRetryConfig      *DefaultRetryConfig
	ConfigurableRetryConfig *ConfigurableRetryConfig
}
```

RetryStrategyConfig to be specified while creating a NewClient
