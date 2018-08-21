// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package idp

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Supported authentication flows
//
//   Authorization code with PKCE (pkce) -- known/trusted app
//      client_id + code challenge + redirect_uri + username + password =>
//          access, id_token, refresh_token[*]
//
//   Client credentials (client) -- private service
//      client_id + client_secret + custom scope =>
//          access
//
//   * refresh_token is available if enabled and if the offline_access scope
//     is requested.
//
// Note: pkce is normally browser based and involves redirection. Support is
// also provided for IdPs with extensions to the standard OIDC flows that allow
// client code to first authenticate with user credentials against a "primary"
// endpoint (/authn) and retrieve a one time session token, which when used with
// these flows, will result in the requested grants being returned directly in
// the redirect url.

// OIDCScope defines scopes that are OpenID Connect compatible, see: https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims
type OIDCScope string

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

// Read and decode JSON data from given reader and return as a map.
func load(r io.Reader) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// Returns the requested string value from the given map.
func gets(m map[string]interface{}, key string) (string, error) {
	value, ok := m[key]
	if !ok {
		return "", fmt.Errorf("key error: '%s'", key)
	}
	result, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("type error: %T", value)
	}
	return result, nil
}

// HTTPError Represents an error response
type HTTPError struct {
	StatusCode int                    `json:"status,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty"`
}

// Error handles marshalling of the HttpError to error type
func (e *HTTPError) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// Returns a golang error corresponding to the given http response.
func httpError(response *http.Response) error {
	var result = &HTTPError{StatusCode: response.StatusCode}

	// ignore if we cant read body details
	_ = json.NewDecoder(response.Body).Decode(&result.Body)

	return result
}

// Context Represents an authentication "context", which is the result of a
// successful OAuth authentication flow.
type Context struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

const (
	defaultAuthnPath     = "api/v1/authn"
	defaultAuthorizePath = "oauth2/default/v1/authorize"
	defaultKeysPath      = "oauth2/default/v1/keys"
	defaultTokenPath     = "oauth2/default/v1/token"
)

// Client captures host and route information for the IdP endpoints
type Client struct {
	Host          string
	PathAuthn     string
	PathAuthorize string
	PathKeys      string
	PathToken     string
}

// NewClient Returns a new IdP client object.
func NewClient(host string, authnPath string, authorizePath string, keysPath string, tokenPath string) *Client {
	if authnPath == "" {
		authnPath = defaultAuthnPath
	}
	if authorizePath == "" {
		authorizePath = defaultAuthorizePath
	}
	if keysPath == "" {
		keysPath = defaultKeysPath
	}
	if tokenPath == "" {
		tokenPath = defaultTokenPath
	}
	return &Client{
		Host:          host,
		PathAuthn:     authnPath,
		PathAuthorize: authorizePath,
		PathKeys:      keysPath,
		PathToken:     tokenPath,
	}
}

// NewDefaultClient returns a new IdP client object with default routes.
func NewDefaultClient(host string) *Client {
	return NewClient(host, "", "", "", "")
}

// Returns a new HTTP client object with redirects disabled.
func newHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		}}
}

// Encode the given value and return its reader.
func encode(value interface{}) (*strings.Reader, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(data)), nil
}

// Decode and return the given response body.
func decode(response *http.Response) (*Context, error) {
	context := &Context{}
	if err := json.NewDecoder(response.Body).Decode(&context); err != nil {
		return nil, err
	}
	return context, nil
}

func newGet(url string, params url.Values) (*http.Request, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.URL.RawQuery = params.Encode()
	return request, nil
}

func newPost(url string, body interface{}) (*http.Request, error) {
	reader, err := encode(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	return request, nil
}

func newFormPost(url string, data url.Values) (*http.Request, error) {
	reader := strings.NewReader(data.Encode())
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return request, nil
}

func get(url string, params url.Values) (*http.Response, error) {
	request, err := newGet(url, params)
	if err != nil {
		return nil, err
	}
	return newHTTPClient().Do(request)
}

func post(url string, body interface{}) (*http.Response, error) {
	request, err := newPost(url, body)
	if err != nil {
		return nil, err
	}
	return newHTTPClient().Do(request)
}

func formPost(url string, data url.Values) (*http.Response, error) {
	request, err := newFormPost(url, data)
	if err != nil {
		return nil, err
	}
	return newHTTPClient().Do(request)
}

// Returns a synthetic state value.
func state() string {
	result, _ := time.Now().MarshalText()
	return string(result)
}

// Return a full URL basd on the given path template.
func (c *Client) url(path string) string {
	return fmt.Sprintf("%s%s", c.Host, path)
}

// ClientFlow will authenticate using the "client credentials" flow.
func (c *Client) ClientFlow(clientID, clientSecret, scope string) (*Context, error) {
	form := url.Values{
		"grant_type": {"client_credentials"},
		"scope":      {scope}}
	request, err := newFormPost(c.url(c.PathToken), form)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(clientID, clientSecret)
	response, err := newHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, httpError(response)
	}
	return decode(response)
}

// GetSessionToken Returns a one-time session token by authenticating using a
// "primary" endpoint (/authn).
func (c *Client) GetSessionToken(username, password string) (string, error) {
	body := map[string]string{
		"username": username,
		"password": password}
	response, err := post(c.url(c.PathAuthn), body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", errors.New(response.Status)
	}
	data, err := load(response.Body)
	if err != nil {
		return "", err
	}
	status, err := gets(data, "status")
	if err != nil {
		return "", err
	}
	if status != "SUCCESS" { // eg: LOCKED_OUT
		return "", errors.New(status)
	}
	sessionToken, err := gets(data, "sessionToken")
	if err != nil {
		return "", err
	}
	return sessionToken, nil
}

// Returns a codeVerfier and codeChallenge for use in a PKCE flow.
func createCodeChallenge(n int) (string, string, error) {
	const safe = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-._~"
	if n < 43 || n > 128 {
		return "", "", errors.New("invalid argument")
	}
	buff := make([]byte, n)
	if _, err := rand.Read(buff); err != nil {
		return "", "", err
	}
	nsafe := byte(len(safe))
	for i, b := range buff {
		b = b % nsafe
		buff[i] = safe[b]
	}
	cv := base64.RawURLEncoding.EncodeToString(buff)
	s256 := sha256.Sum256([]byte(cv))
	cc := base64.RawURLEncoding.EncodeToString(s256[:])
	return cv, cc, nil
}

// PKCEFlow will authenticate using the "proof key for code exchange" flow.
func (c *Client) PKCEFlow(clientID, redirectURI, scope, username, password string) (*Context, error) {
	// retrieve one-time session token
	sessionToken, err := c.GetSessionToken(username, password)
	if err != nil {
		return nil, err
	}

	cv, cc, err := createCodeChallenge(50)

	// request authorization code
	params := url.Values{
		"client_id":             {clientID},
		"code_challenge":        {cc},
		"code_challenge_method": {"S256"},
		"nonce":                 {"none"},
		"redirect_uri":          {redirectURI},
		"response_type":         {"code"},
		"scope":                 {scope},
		"sessionToken":          {sessionToken},
		"state":                 {state()}}
	response, err := get(c.url(c.PathAuthorize), params)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusFound {
		return nil, httpError(response)
	}

	// retrieve the authorization code from the redirect url query string
	location := response.Header.Get("Location")
	locationURL, err := url.Parse(location)
	if err != nil {
		return nil, err
	}
	code := locationURL.Query().Get("code")

	// exchange authorization code for token(s)
	form := url.Values{
		"client_id":     {clientID},
		"code":          {code},
		"code_verifier": {cv},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {redirectURI}}
	response, err = formPost(c.url(c.PathToken), form)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, httpError(response)
	}

	return decode(response)
}

// Decode the given implicit flow location url fragment, and return values
// as an authentication Context.
func decodeFragment(fragment string) (*Context, error) {
	context := &Context{}
	pairs := strings.Split(fragment, "&")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid fragment property: '%s", kv)
		}
		k := kv[0]
		v := kv[1]
		switch k {
		case "access_token":
			context.AccessToken = v
		case "expires_in":
			expiresIn, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			context.ExpiresIn = expiresIn
		case "id_token":
			context.IDToken = v
		case "refresh_token":
			context.RefreshToken = v
		case "scope":
			context.Scope = v
		case "state":
			break // ignore
		case "token_type":
			context.TokenType = v
		default:
			return nil, fmt.Errorf("invalid fragment key: '%s'", k)
		}
	}
	return context, nil
}

// Refresh will authenticate using a refresh token.
func (c *Client) Refresh(clientID, scope, refreshToken string) (*Context, error) {
	form := url.Values{
		"client_id":     {clientID},
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
		"scope":         {scope}}
	response, err := formPost(c.url(c.PathToken), form)
	if err != nil {
		return nil, err
	}
	return decode(response)
}
