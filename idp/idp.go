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

package idp

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
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
	// ScopeOffline - This scope value requests that an OAuth 2.0 Refresh Token be issued that can be used to obtain an Access Token that grants access to the End-User's UserInfo Endpoint even when the End-User is not present (not logged in).
	ScopeOffline OIDCScope = "offline_access"
	// RequestIDHeader uniquely identifies a request in the platform, found in response headers
	RequestIDHeader string = "X-Request-Id"
)

var (
	// DefaultOIDCScopes defines the default OpenID Connect scopes to use in authn requests - "openid email profile"
	DefaultOIDCScopes   = fmt.Sprintf("%s %s %s", ScopeOpenID, ScopeEmail, ScopeProfile)
	DefaultRefreshScope = fmt.Sprintf("%s %s %s %s", ScopeOpenID, ScopeEmail, ScopeOffline, ScopeProfile)
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

// Context Represents an authentication "context", which is the result of a
// successful OAuth authentication flow.
type Context struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	StartTime    int64
}

// DeviceCodeInfo captures codes, verification URI and polling parameters for device flow.
type DeviceCodeInfo struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
	VerificationURI string `json:"verification_uri"`
}

const (
	defaultAuthnPath           = "authn"
	defaultAuthorizePath       = "authorize"
	defaultTokenPath           = "token"
	defaultTenantTokenPath     = "system/token"
	defaultTenantTokenTemplate = "%s/token"
	defaultCsrfTokenPath       = "csrfToken"
	defaultDevicePath          = "system/device"
	defaultDevicePathTemplate  = "%s/device"
)

// Client captures url and route information for the IdP endpoints
type Client struct {
	ProviderHost    string
	OverrideAuthURL string
	AuthnPath       string
	AuthorizePath   string
	TokenPath       string
	TenantTokenPath string
	DevicePath      string
	CsrfTokenPath   string
	Insecure        bool
	hostURLConfig   HostURLConfig
}

// NewClient Returns a new IdP client object.
//   providerURL: should be of the form https://example.com or optionally https://example.com:port
func NewClient(providerURL string,
	overrideAuthURL string,
	authnPath string,
	authorizePath string,
	tokenPath string,
	tenantTokenPath string,
	csrfTokenPath string,
	devicePath string,
	insecure bool,
	hostURLConfig HostURLConfig) *Client {
	// Add a trailing slash if none
	if overrideAuthURL != "" && overrideAuthURL[len(overrideAuthURL)-1:] != "/" {
		overrideAuthURL = overrideAuthURL + "/"
	}
	// Add a trailing slash if none
	if providerURL[len(providerURL)-1:] != "/" {
		providerURL = providerURL + "/"
	}
	if authnPath == "" {
		authnPath = defaultAuthnPath
	}
	if authorizePath == "" {
		authorizePath = defaultAuthorizePath
	}
	if tokenPath == "" {
		tokenPath = defaultTokenPath
	}
	if tenantTokenPath == "" {
		tenantTokenPath = defaultTenantTokenPath
	}
	if csrfTokenPath == "" {
		csrfTokenPath = defaultCsrfTokenPath
	}
	if devicePath == "" {
		devicePath = defaultDevicePath
	}
	return &Client{
		ProviderHost:    providerURL,
		OverrideAuthURL: overrideAuthURL,
		AuthnPath:       authnPath,
		AuthorizePath:   authorizePath,
		TokenPath:       tokenPath,
		TenantTokenPath: tenantTokenPath,
		CsrfTokenPath:   csrfTokenPath,
		DevicePath:      devicePath,
		Insecure:        insecure,
		hostURLConfig:   hostURLConfig,
	}
}

// Returns a new HTTP client object with redirects disabled.
func newHTTPClient(insecure bool) *http.Client {
	return &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure}},
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		}}
}

// Decode and return the given response body.
func decode(response *http.Response) (*Context, error) {
	context := &Context{}
	if err := json.NewDecoder(response.Body).Decode(&context); err != nil {
		return nil, err
	}
	now := time.Now()
	context.StartTime = now.Unix()
	return context, nil
}

func newGet(reqURL string, params url.Values) (*http.Request, error) {
	request, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.URL.RawQuery = params.Encode()
	return request, nil
}

func newFormPost(reqURL string, data url.Values) (*http.Request, error) {
	reader := strings.NewReader(data.Encode())
	request, err := http.NewRequest("POST", reqURL, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return request, nil
}

func get(reqURL string, params url.Values, cookies []*http.Cookie, insecure bool) (*http.Response, error) {
	request, err := newGet(reqURL, params)
	if err != nil {
		return nil, err
	}
	if cookies != nil {
		for _, cookie := range cookies {
			request.AddCookie(cookie)
		}
	}
	return newHTTPClient(insecure).Do(request)
}

// Encode the given value and return its reader.
func encode(value interface{}) (*strings.Reader, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(data)), nil
}

func newPost(reqURL string, body interface{}, cookies ...*http.Cookie) (*http.Request, error) {
	reader, err := encode(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", reqURL, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	if cookies != nil {
		for _, cookie := range cookies {
			request.Header.Add("Cookie", cookie.String())
		}
	}

	return request, nil
}

func post(reqURL string, body interface{}, insecure bool, cookies ...*http.Cookie) (*http.Response, error) {
	request, err := newPost(reqURL, body, cookies...)
	if err != nil {
		return nil, err
	}
	return newHTTPClient(insecure).Do(request)
}

func formPost(reqURL string, data url.Values, insecure bool) (*http.Response, error) {
	request, err := newFormPost(reqURL, data)
	if err != nil {
		return nil, err
	}
	return newHTTPClient(insecure).Do(request)
}

// Returns a synthetic state value.
func state() string {
	result, _ := time.Now().MarshalText()
	return string(result)
}

// Return a full URL based on the given host and path template.
func (c *Client) makeURL(hostURL string, path string) string {
	return fmt.Sprintf("%s%s", hostURL, path)
}

func getCookie(cookies []*http.Cookie, name string) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie
		}
	}
	return nil
}

// If tenantScoped is true, return tenant.hostname
func getTenantScopedHost(tenantScoped bool, tenant string, hostURL string) (string, error) {
	//check tenant scoped, tenant is not empty
	var err error
	if tenantScoped == true && tenant == "" {
		return "", errors.New("Tenant needs to be specified to get token")
	}
	//check tenant scoped, append tenant to idp host, so hostname is tenant scoped
	if tenantScoped == true {
		u, err := url.Parse(hostURL)
		if err != nil {
			return "", errors.New("Host url specified is not in the correct format")
		}
		hostURL = fmt.Sprintf("%s://%s.%s/", u.Scheme, tenant, u.Host)
	}

	return hostURL, err
}

// If tenantScoped is true, return region-region.hostname
func getRegionScopedHost(tenantScoped bool, region string, hostURL string) (string, error) {
	//check tenant scoped, region is not empty
	var err error
	if tenantScoped == true && region == "" {
		return "", errors.New("Region needs to be specified to get token")
	}
	//check tenant scoped, append region to idp host
	if tenantScoped == true {
		u, err := url.Parse(hostURL)
		if err != nil {
			return "", errors.New("Host url specified is not in the correct format")
		}
		hostURL = fmt.Sprintf("%s://region-%s.%s/", u.Scheme, region, u.Host)
	}
	return hostURL, err
}

// getRequestID retrieves request id from response header if exists
func getRequestID(response *http.Response) string {
	requestID := "00000000-0000-0000-0000-000000000000"
	if response == nil {
		return requestID
	}
	if rid := response.Header.Get(RequestIDHeader); rid != "" {
		requestID = rid
	}
	return requestID
}

// ClientFlow will authenticate using the "client credentials" flow.
func (c *Client) ClientFlow(clientID, clientSecret, scope string) (*Context, error) {
	form := url.Values{
		"grant_type": {"client_credentials"},
		"scope":      {scope}}

	var err error
	hostURL := c.ProviderHost

	if c.OverrideAuthURL != "" {
		hostURL = c.OverrideAuthURL
	}
	//Tenant scoped or region scoped hostname based on non-system or system tenant
	if c.hostURLConfig.TenantScoped && c.OverrideAuthURL == "" {
		if c.hostURLConfig.Tenant != "system" {
			hostURL, err = getTenantScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Tenant, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating tenant scoped url")
			}
			c.TokenPath = fmt.Sprintf(defaultTenantTokenTemplate, c.hostURLConfig.Tenant)
		} else if c.hostURLConfig.Tenant == "system" {
			hostURL, err = getRegionScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Region, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating region scoped url")
			}
			c.TokenPath = defaultTenantTokenPath
		}
	}

	tokenURL := c.makeURL(hostURL, c.TokenPath)
	request, err := newFormPost(tokenURL, form)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to create request to token endpoint url: %s", tokenURL))
	}
	request.SetBasicAuth(clientID, clientSecret)
	response, err := newHTTPClient(c.Insecure).Do(request)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get response from token endpoint url: %s", tokenURL))
	}
	requestID := getRequestID(response)
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		data, err := load(response.Body)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to parse response body from token endpoint url: %s request id: %s", tokenURL, requestID))
		}
		msg := response.Status
		if data != nil {
			msg += "\nResponse body: \n"
			for k, v := range data {
				msg += fmt.Sprintf("%s:%s \n", k, v)
			}
		}
		return nil, errors.Wrap(errors.New(msg), fmt.Sprintf("failed to get a successful response from token endpoint url: %s request id: %s", tokenURL, requestID))
	}
	return decode(response)
}

func (c *Client) GetCsrfToken() (string, []*http.Cookie, error) {
	var err error
	hostURL := c.ProviderHost

	if c.OverrideAuthURL != "" {
		hostURL = c.OverrideAuthURL
	}

	if c.hostURLConfig.TenantScoped && c.OverrideAuthURL == "" {
		hostURL, err = getRegionScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Region, c.ProviderHost)
		if err != nil {
			return "", nil, errors.Wrap(err, "error in creating region scoped url")
		}
	}

	tokenURL := c.makeURL(hostURL, c.CsrfTokenPath)
	response, err := get(tokenURL, nil, nil, c.Insecure)
	if err != nil {
		return "", nil, errors.Wrap(err, fmt.Sprintf("failed to get valid response from csrfToken endpoint url: %s", tokenURL))
	}
	requestID := getRequestID(response)
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", nil, errors.Wrap(errors.New(response.Status), fmt.Sprintf("unexpected status response from csrfToken endpoint url: %s request id: %s", tokenURL, requestID))
	}

	csrfTokenCookie := getCookie(response.Cookies(), "csrf")
	if csrfTokenCookie == nil {
		return "", nil, errors.Wrap(errors.New("missing cookies"), fmt.Sprintf("failed to get successful response from csrfToken endpoint url: %s request id: %s", tokenURL, requestID))
	}

	return csrfTokenCookie.Value, response.Cookies(), nil
}

// GetSessionToken Returns a one-time session token by authenticating using a
// "primary" endpoint (/authn).
func (c *Client) GetSessionToken(username, password string) (string, []*http.Cookie, error) {
	csrfToken, cookies, err := c.GetCsrfToken()
	if err != nil {
		return "", nil, errors.Wrap(err, "failed to get valid response from csrfToken endpoint")
	}

	body := map[string]string{
		"username":  username,
		"password":  password,
		"csrfToken": csrfToken}

	hostURL := c.ProviderHost

	if c.OverrideAuthURL != "" {
		hostURL = c.OverrideAuthURL
	}

	if c.hostURLConfig.TenantScoped && c.OverrideAuthURL == "" {
		hostURL, err = getRegionScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Region, c.ProviderHost)
		if err != nil {
			return "", nil, errors.Wrap(err, "error in creating region scoped url")
		}
	}

	authnURL := c.makeURL(hostURL, c.AuthnPath)
	response, err := post(authnURL, body, c.Insecure, cookies...)
	if err != nil {
		return "", nil, errors.Wrap(err, fmt.Sprintf("failed to get valid response from authn endpoint url: %s", authnURL))
	}
	requestID := getRequestID(response)
	defer response.Body.Close()
	data, err := load(response.Body)
	if err != nil {
		return "", nil, errors.Wrap(err, fmt.Sprintf("failed to parse response body from authn endpoint url: %s request id: %s", authnURL, requestID))
	}
	if response.StatusCode != http.StatusOK {
		msg := response.Status
		if data != nil {
			msg += "\nResponse body: \n"
			for k, v := range data {
				msg += fmt.Sprintf("%s:%s \n", k, v)
			}
		}
		return "", nil, errors.Wrap(errors.New(msg), fmt.Sprintf("unexpected status response from authn endpoint url: %s request id: %s", authnURL, requestID))
	}
	status, err := gets(data, "status")
	if err != nil {
		return "", nil, errors.Wrap(err, fmt.Sprintf("unable to get status data from authn endpoint url: %s request id: %s", authnURL, requestID))
	}
	if status != "SUCCESS" { // eg: LOCKED_OUT
		return "", nil, errors.Wrap(errors.New(status), fmt.Sprintf("unexpected status data from authn endpoint url: %s request id: %s", authnURL, requestID))
	}
	// get session token from the cookies (Keycloak session tokens are passed as cookies)
	var sessionToken string
	var sessionCookies []*http.Cookie
	for _, cookie := range response.Cookies() {
		if cookie.Name == "sessionToken" {
			sessionCookies = []*http.Cookie{cookie}
			break
		}
	}
	if sessionCookies == nil {
		// try data (OKTA session token)
		sessionToken, err = gets(data, "sessionToken")
		if err != nil {
			return "", nil, errors.Wrap(err, fmt.Sprintf("failed to parse session token data from authn endpoint response url: %s request id: %s", authnURL, requestID))
		}
	}
	return sessionToken, sessionCookies, nil
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
	sessionToken, sessionCookies, err := c.GetSessionToken(username, password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get session token")
	}

	cv, cc, err := createCodeChallenge(50)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get challenge code")
	}

	// request authorization code
	params := url.Values{
		"client_id":             {clientID},
		"code_challenge":        {cc},
		"code_challenge_method": {"S256"},
		"nonce":                 {"none"},
		"redirect_uri":          {redirectURI},
		"response_type":         {"code"},
		"scope":                 {scope},
		"session_token":         {sessionToken},
		"state":                 {state()}}

	hostURL := c.ProviderHost

	if c.OverrideAuthURL != "" {
		hostURL = c.OverrideAuthURL
	}

	if c.hostURLConfig.TenantScoped && c.OverrideAuthURL == "" {
		if c.hostURLConfig.Tenant != "system" {
			hostURL, err = getTenantScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Tenant, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating tenant scoped url")
			}
		} else if c.hostURLConfig.Tenant == "system" {
			hostURL, err = getRegionScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Region, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating region scoped url")
			}
		}
	}

	authzURL := c.makeURL(hostURL, c.AuthorizePath)
	response, err := get(authzURL, params, sessionCookies, c.Insecure)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get valid response from authorize endpoint url: %s", authzURL))
	}
	requestID := getRequestID(response)
	if response.StatusCode != http.StatusFound {
		return nil, errors.Wrap(errors.New(fmt.Sprintf("%v", response.StatusCode)), fmt.Sprintf("failed to get successful response from authorize endpoint url: %s request id: %s", authzURL, requestID))
	}

	// retrieve the authorization code from the redirect url query string
	location := response.Header.Get("Location")
	locationURL, err := url.Parse(location)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to parse location header: '%s' url: %s request id: %s", location, authzURL, requestID))
	}
	code := locationURL.Query().Get("code")
	if code == "" {
		err := errors.New("")
		return nil, errors.Wrap(err, fmt.Sprintf("failed to retrieve valid authorization code from the redirect url: %s request url: %s request id: %s", locationURL, authzURL, requestID))
	}

	// exchange authorization code for token(s)
	form := url.Values{
		"client_id":     {clientID},
		"code":          {code},
		"code_verifier": {cv},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {redirectURI}}

	hostURL = c.ProviderHost

	if c.OverrideAuthURL != "" {
		hostURL = c.OverrideAuthURL
	}
	if c.hostURLConfig.TenantScoped && c.OverrideAuthURL == "" {
		if c.hostURLConfig.Tenant != "system" {
			hostURL, err = getTenantScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Tenant, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating tenant scoped url")
			}
			c.TokenPath = fmt.Sprintf(defaultTenantTokenTemplate, c.hostURLConfig.Tenant)
		} else if c.hostURLConfig.Tenant == "system" {
			hostURL, err = getRegionScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Region, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating region scoped url")
			}
			c.TokenPath = defaultTenantTokenPath
		}
	}

	tokenURL := c.makeURL(hostURL, c.TokenPath)
	response, err = formPost(tokenURL, form, c.Insecure)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get valid response from token endpoint url: %s", tokenURL))
	}
	requestID = getRequestID(response)
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		data, err := load(response.Body)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to parse response body from token endpoint url: %s request id: %s", tokenURL, requestID))
		}
		msg := response.Status
		if data != nil {
			msg += "\nResponse body: \n"
			for k, v := range data {
				msg += fmt.Sprintf("%s:%s \n", k, v)
			}
		}
		return nil, errors.Wrap(errors.New(msg), fmt.Sprintf("failed to get a successful response from token endpoint url: %s request id: %s", tokenURL, requestID))

	}
	return decode(response)
}

// Refresh will authenticate using a refresh token.
func (c *Client) Refresh(clientID, scope, refreshToken string) (*Context, error) {
	form := url.Values{
		"client_id":     {clientID},
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
		"scope":         {scope}}

	var err error
	hostURL := c.ProviderHost

	if c.OverrideAuthURL != "" {
		hostURL = c.OverrideAuthURL
	}

	if c.hostURLConfig.TenantScoped == true && c.OverrideAuthURL == "" {
		if c.hostURLConfig.Tenant != "system" {
			hostURL, err = getTenantScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Tenant, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating tenant scoped url")
			}
		} else if c.hostURLConfig.Tenant == "system" {
			hostURL, err = getRegionScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Region, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating region scoped url")
			}
		}
	}
	// tenant/token path if non-system tenant, else system/token
	if c.hostURLConfig.Tenant != "system" {
		c.TokenPath = fmt.Sprintf(defaultTenantTokenTemplate, c.hostURLConfig.Tenant)
	} else if c.hostURLConfig.Tenant == "system" {
		c.TokenPath = defaultTenantTokenPath
	}
	tokenURL := c.makeURL(hostURL, c.TokenPath)
	response, err := formPost(tokenURL, form, c.Insecure)

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get valid response from token endpoint url: %s", tokenURL))
	}
	requestID := getRequestID(response)
	if response.StatusCode != http.StatusOK {
		return nil, errors.Wrap(errors.New(fmt.Sprintf("%v", response.StatusCode)), fmt.Sprintf("failed to get successful response from token endpoint url: %s request id: %s", tokenURL, requestID))
	}
	return decode(response)
}

// GetDeviceCodes will get info for the device flow.
func (c *Client) GetDeviceCodes(clientID, scope string) (*DeviceCodeInfo, error) {
	form := url.Values{
		"client_id": {clientID},
		"scope":     {scope}}

	var err error
	hostURL := c.ProviderHost

	if c.OverrideAuthURL != "" {
		hostURL = c.OverrideAuthURL
	}

	//Tenant scoped or region scoped hostname based on non-system or system tenant
	if c.hostURLConfig.TenantScoped == true && c.OverrideAuthURL == "" {
		if c.DevicePath == defaultDevicePath && c.hostURLConfig.Tenant != "system" {
			hostURL, err = getTenantScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Tenant, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating tenant scoped url")
			}
		} else if c.hostURLConfig.Tenant == "system" {
			hostURL, err = getRegionScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Region, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating region scoped url")
			}
		}
	}
	// tenant/token path if non-system tenant, else system/token
	if c.hostURLConfig.Tenant != "system" {
		c.DevicePath = fmt.Sprintf(defaultDevicePathTemplate, c.hostURLConfig.Tenant)
	} else if c.hostURLConfig.Tenant == "system" {
		c.DevicePath = defaultDevicePath
	}

	deviceURL := c.makeURL(hostURL, c.DevicePath)
	response, err := formPost(deviceURL, form, c.Insecure)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to get valid response from device endpoint url: %s", deviceURL))
	}
	requestID := getRequestID(response)
	if response.StatusCode != http.StatusOK {
		return nil, errors.Wrap(errors.New(fmt.Sprintf("%v", response.StatusCode)), fmt.Sprintf("failed to get successful response from device endpoint url: %s request id: %s", deviceURL, requestID))
	}

	var info DeviceCodeInfo
	if err := json.NewDecoder(response.Body).Decode(&info); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to decode response from device endpoint url: %s request id: %s", deviceURL, requestID))
	}
	return &info, nil
}

// DeviceFlow will authenticate using the device flow.
func (c *Client) DeviceFlow(clientID, deviceCode string, expiresIn, interval int) (*Context, error) {
	var response *http.Response
	codeExpiration := time.Now().Add(time.Duration(expiresIn) * time.Second)
	pollingInterval := time.Duration(interval) * time.Second
	form := url.Values{
		"client_id":   {clientID},
		"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
		"device_code": {deviceCode}}

	var err error
	hostURL := c.ProviderHost

	if c.OverrideAuthURL != "" {
		hostURL = c.OverrideAuthURL
	}
	//Tenant scoped or region scoped hostname based on non-system or system tenant
	if c.hostURLConfig.TenantScoped == true && c.OverrideAuthURL == "" {
		if c.hostURLConfig.Tenant != "system" {
			hostURL, err = getTenantScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Tenant, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating tenant scoped url")
			}
		} else if c.hostURLConfig.Tenant == "system" {
			hostURL, err = getRegionScopedHost(c.hostURLConfig.TenantScoped, c.hostURLConfig.Region, c.ProviderHost)
			if err != nil {
				return nil, errors.Wrap(err, "error in creating region scoped url")
			}
		}
	}

	// tenant/token path if non-system tenant, else system/token
	if c.hostURLConfig.Tenant != "system" {
		c.TenantTokenPath = fmt.Sprintf(defaultTenantTokenTemplate, c.hostURLConfig.Tenant)
	} else if c.hostURLConfig.Tenant == "system" {
		c.TenantTokenPath = defaultTenantTokenPath
	}
	for time.Now().Before(codeExpiration) {
		tokenURL := c.makeURL(hostURL, c.TenantTokenPath)
		response, err = formPost(tokenURL, form, c.Insecure)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to get valid response from tenant token endpoint url: %s", tokenURL))
		}
		requestID := getRequestID(response)
		if response.StatusCode == http.StatusBadRequest {
			defer response.Body.Close()
			data, err := load(response.Body)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("failed to parse response body from tenant token endpoint url: %s request id: %s", tokenURL, requestID))
			}
			switch data["error_description"] {
			case "authorization_pending":
				time.Sleep(pollingInterval)
				continue
			case "slow_down":
				// add 5 seconds to polling interval on "slow_down" error
				pollingInterval += 5
				time.Sleep(pollingInterval)
				continue
			case "expired_token":
				return nil, errors.Wrap(errors.New(fmt.Sprintf("%v", response.StatusCode)), fmt.Sprintf("code expired %s url: request id: %s", tokenURL, requestID))
			case "access_denied":
				return nil, errors.Wrap(errors.New(fmt.Sprintf("%v", response.StatusCode)), fmt.Sprintf("access denied %s url: request id: %s", tokenURL, requestID))
			default:
				return nil, errors.Wrap(errors.New(fmt.Sprintf("%v", response.StatusCode)), fmt.Sprintf("failed to get successful response from tenant token endpoint %s url: request id: %s", tokenURL, requestID))
			}
		} else if response.StatusCode != http.StatusOK {
			return nil, errors.Wrap(errors.New(fmt.Sprintf("%v", response.StatusCode)), fmt.Sprintf("failed to get successful response from tenant token endpoint %s url: request id: %s", tokenURL, requestID))
		} else {
			// return decoded response if response.StatusCode = 200
			return decode(response)
		}
	}
	// return "code expired" on timeout
	return nil, errors.Wrap(errors.New(fmt.Sprintf("%v", http.StatusBadRequest)), "code expired")
}
