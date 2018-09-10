# idp
--
    import "github.com/splunk/splunk-cloud-sdk-go/idp"


## Usage

```go
var (
	// DefaultOIDCScopes defines the default OpenID Connect scopes to use in authn requests - "openid email profile"
	DefaultOIDCScopes = fmt.Sprintf("%s %s %s", ScopeOpenID, ScopeEmail, ScopeProfile)
)
```

#### type Client

```go
type Client struct {
	Host          string
	PathAuthn     string
	PathAuthorize string
	PathKeys      string
	PathToken     string
}
```

Client captures host and route information for the IdP endpoints

#### func  NewClient

```go
func NewClient(host string, authnPath string, authorizePath string, keysPath string, tokenPath string) *Client
```
NewClient Returns a new IdP client object.

#### func  NewDefaultClient

```go
func NewDefaultClient(host string) *Client
```
NewDefaultClient returns a new IdP client object with default routes.

#### func (*Client) ClientFlow

```go
func (c *Client) ClientFlow(clientID, clientSecret, scope string) (*Context, error)
```
ClientFlow will authenticate using the "client credentials" flow.

#### func (*Client) GetSessionToken

```go
func (c *Client) GetSessionToken(username, password string) (string, error)
```
GetSessionToken Returns a one-time session token by authenticating using a
"primary" endpoint (/authn).

#### func (*Client) PKCEFlow

```go
func (c *Client) PKCEFlow(clientID, redirectURI, scope, username, password string) (*Context, error)
```
PKCEFlow will authenticate using the "proof key for code exchange" flow.

#### func (*Client) Refresh

```go
func (c *Client) Refresh(clientID, scope, refreshToken string) (*Context, error)
```
Refresh will authenticate using a refresh token.

#### type ClientCredentialsRetriever

```go
type ClientCredentialsRetriever struct {
	*Client
	// ClientID to authenticate as which corresponds to a Client Credentials flow supported IdP client
	ClientID string
	// ClientSecret corresponding to the ClientID above
	ClientSecret *util.Credential
	// Scope(s) to request, separated by spaces -- this will be a custom scope, for example: "backend_service"
	Scope string
}
```

ClientCredentialsRetriever retries a request after gettting a new access token
from the identity provider using the Client Credentials flow

#### func  NewClientCredentialsRetriever

```go
func NewClientCredentialsRetriever(idpHost string, clientID string, clientSecret string, scope string) *ClientCredentialsRetriever
```
NewClientCredentialsRetriever initializes a new token context retriever

#### func (*ClientCredentialsRetriever) GetTokenContext

```go
func (tr *ClientCredentialsRetriever) GetTokenContext() (*Context, error)
```
GetTokenContext gets a new access token context from the identity provider

#### type Context

```go
type Context struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
```

Context Represents an authentication "context", which is the result of a
successful OAuth authentication flow.

#### type HTTPError

```go
type HTTPError struct {
	StatusCode int                    `json:"status,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty"`
}
```

HTTPError Represents an error response

#### func (*HTTPError) Error

```go
func (e *HTTPError) Error() string
```
Error handles marshalling of the HttpError to error type

#### type NoOpTokenRetriever

```go
type NoOpTokenRetriever struct {
	Context *Context
}
```

NoOpTokenRetriever just returns the same static Context

#### func (*NoOpTokenRetriever) GetTokenContext

```go
func (tr *NoOpTokenRetriever) GetTokenContext() (*Context, error)
```
GetTokenContext just returns the same static Context

#### type OIDCScope

```go
type OIDCScope string
```

OIDCScope defines scopes that are OpenID Connect compatible, see:
https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims

```go
const (
	// ScopeOpenID - The basic (and required) scope for OpenID Connect
	ScopeOpenID OIDCScope = "openid"
	// ScopeEmail - This scope value requests access to the email and email_verified Claims
	ScopeEmail OIDCScope = "email"
	// ScopeProfile - This scope value requests access to the End-User's default profile Claims, which are: name, family_name,
	// given_name, middle_name, nickname, preferred_username, profile, picture, website, gender, birthdate, zoneinfo,
	// locale, and updated_at
	ScopeProfile OIDCScope = "profile"
	// ScopeAddress - This scope value requests access to the address Claim
	ScopeAddress OIDCScope = "address"
	// ScopePhone - This scope value requests access to the phone_number and phone_number_verified Claims
	ScopePhone OIDCScope = "phone"
)
```

#### type PKCERetriever

```go
type PKCERetriever struct {
	*Client
	// ClientID corresponding to a PKCE flow supported IdP client
	ClientID string
	// RedirectURI that has been whitelisted according to the ClientID (NOTE: redirection is not actually needed for this implementation but this URI must match one specified by the IdP)
	RedirectURI string
	// Scope(s) to request, separated by spaces -- "openid email profile" is recommended for individual users
	Scope string
	// Username to authenticate as which must be registered to the ClientID in the IdP
	Username string
	// Password corresponding to the Username above
	Password *util.Credential
}
```

PKCERetriever retries a request after gettting a new access token from the
identity provider using the Proof Key for Code Exchange (PKCE) flow

#### func  NewPKCERetriever

```go
func NewPKCERetriever(idpHost string, clientID string, redirectURI string, scope string, username string, password string) *PKCERetriever
```
NewPKCERetriever initializes a new token context retriever

#### func (*PKCERetriever) GetTokenContext

```go
func (tr *PKCERetriever) GetTokenContext() (*Context, error)
```
GetTokenContext gets a new access token context from the identity provider

#### type RefreshTokenRetriever

```go
type RefreshTokenRetriever struct {
	*Client
	// ClientID which corresponds to an Refresh Token ("offline_access" scope) supported IdP client
	ClientID string
	// Scope(s) to request, separated by spaces -- "openid email profile" is recommended for individual users
	Scope string
	// RefreshToken to use to authenticate in order to generate an access token
	RefreshToken *util.Credential
}
```

RefreshTokenRetriever retries a request after gettting a new access token from
the identity provider using a RefreshToken

#### func  NewRefreshTokenRetriever

```go
func NewRefreshTokenRetriever(idpHost string, clientID string, scope string, refreshToken string) *RefreshTokenRetriever
```
NewRefreshTokenRetriever initializes a new token context retriever

#### func (*RefreshTokenRetriever) GetTokenContext

```go
func (tr *RefreshTokenRetriever) GetTokenContext() (*Context, error)
```
GetTokenContext gets a new access token context from the identity provider

#### type TokenRetriever

```go
type TokenRetriever interface {
	GetTokenContext() (*Context, error)
}
```

TokenRetriever retrieves an access token with context
