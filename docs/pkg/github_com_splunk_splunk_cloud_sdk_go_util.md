# util
--
    import "github.com/splunk/splunk-cloud-sdk-go/util"


## Usage

#### func  ParseHTTPStatusCodeInResponse

```go
func ParseHTTPStatusCodeInResponse(response *http.Response) (*http.Response, error)
```
ParseHTTPStatusCodeInResponse returns http response and HTTPError struct based
on response status code

#### func  ParseResponse

```go
func ParseResponse(model interface{}, response *http.Response) error
```
ParseResponse parses json-formatted http response and decodes it into
pre-defined models

#### func  ParseURLParams

```go
func ParseURLParams(model interface{}) url.Values
```
ParseURLParams parses a struct into url params based on its "key" tag It parses
basic values and slices, and will parse structs recursively

#### type Credential

```go
type Credential struct {
}
```

Credential is a simple string whose value is redacted when converted to a string
via the Stringer interface in order to prevent accidental logging or other
unintentional disclosure - value is retrieved using ClearText() method

#### func  NewCredential

```go
func NewCredential(s string) *Credential
```
NewCredential creates a Credential from a simple string

#### func (*Credential) ClearText

```go
func (c *Credential) ClearText() string
```
ClearText returns the actual cleartext string value

#### func (*Credential) String

```go
func (c *Credential) String() string
```
String returns a redacted string

#### type HTTPError

```go
type HTTPError struct {
	HTTPStatusCode int
	HTTPStatus     string
	Message        string      `json:"message,omitempty"`
	Code           string      `json:"code,omitempty"`
	MoreInfo       string      `json:"moreInfo,omitempty"`
	Details        interface{} `json:"details,omitempty"`
}
```

HTTPError is raised when status code is not 2xx

#### func (*HTTPError) Error

```go
func (he *HTTPError) Error() string
```
This allows HTTPError to satisfy the error interface

#### type Logger

```go
type Logger interface {
	Print(v ...interface{})
}
```

Logger compatible with standard "log" library

#### type SdkTransport

```go
type SdkTransport struct {
}
```

SdkTransport is to define a transport RoundTripper with user-defined logger

#### func  CreateRoundTripperWithLogger

```go
func CreateRoundTripperWithLogger(logger Logger) *SdkTransport
```
CreateRoundTripperWithLogger Creates a RoundTripper with user defined logger

#### func (*SdkTransport) RoundTrip

```go
func (st *SdkTransport) RoundTrip(request *http.Request) (*http.Response, error)
```
RoundTrip implements the RoundTripper interface

#### type Ticker

```go
type Ticker struct {
}
```

Ticker is a wrapper of time.Ticker with additional functionality

#### func  NewTicker

```go
func NewTicker(duration time.Duration) *Ticker
```
NewTicker spits out a pointer to Ticker model. It sets ticker to stop state by
default

#### func (*Ticker) GetChan

```go
func (t *Ticker) GetChan() <-chan time.Time
```
GetChan returns the channel from ticker

#### func (*Ticker) IsRunning

```go
func (t *Ticker) IsRunning() bool
```
IsRunning returns bool indicating whether or not ticker is running

#### func (*Ticker) Reset

```go
func (t *Ticker) Reset()
```
Reset resets ticker

#### func (*Ticker) Start

```go
func (t *Ticker) Start()
```
Start starts a new ticker and set property running to true

#### func (*Ticker) Stop

```go
func (t *Ticker) Stop()
```
Stop stops ticker and set property running to false
