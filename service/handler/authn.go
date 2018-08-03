package handler

import (
	"net/http"

	"github.com/splunk/ssc-client-go/idp"
	"github.com/splunk/ssc-client-go/service"
)

const (
	// DefaultMaxAuthnAttempts defines the maximum number of retries that will be performed for a request encountering an authentication issue
	DefaultMaxAuthnAttempts = 1
)

// TokenRetriever retrieves an access token
type TokenRetriever interface {
	GetAccessToken() (token string, err error)
}

// AuthnResponseHandler handles logic for updating the client access token in
// response to 401 errors
type AuthnResponseHandler struct {
	IdpClient      *idp.Client
	TokenRetriever TokenRetriever
}

// HandleResponse will retry a request once after re-authenticating if a 401 response code is encountered
func (rh *AuthnResponseHandler) HandleResponse(client *service.Client, request *service.Request, response *http.Response) (*http.Response, error) {
	if response.StatusCode != 401 || rh.TokenRetriever == nil || request.GetNumErrorsByResponseCode(401) > DefaultMaxAuthnAttempts {
		return response, nil
	}
	token, err := rh.TokenRetriever.GetAccessToken()
	if err != nil {
		return response, err
	}
	client.UpdateToken(token)
	// Retry the request with the updated token
	return client.Do(request)
}

// RefreshTokenAuthnResponseHandler retries a request after gettting a new access token from the identity provider using a RefreshToken
type RefreshTokenAuthnResponseHandler struct {
	*AuthnResponseHandler
	ClientID     string
	Scope        string
	RefreshToken string
}

// NewRefreshTokenAuthnResponseHandler initializes a new response handler
func NewRefreshTokenAuthnResponseHandler(idpHost string, clientID string, scope string, refreshToken string) *RefreshTokenAuthnResponseHandler {
	handler := &RefreshTokenAuthnResponseHandler{
		AuthnResponseHandler: &AuthnResponseHandler{IdpClient: idp.NewClient(idpHost, "", "", "", "")},
		ClientID:             clientID,
		Scope:                scope,
		RefreshToken:         refreshToken,
	}
	handler.TokenRetriever = handler
	return handler
}

// GetAccessToken gets a new access token from the identity provider
func (rh *RefreshTokenAuthnResponseHandler) GetAccessToken() (token string, err error) {
	ctx, err := rh.IdpClient.Refresh(rh.ClientID, rh.Scope, rh.RefreshToken)
	if err != nil {
		return "", err
	}
	return ctx.AccessToken, nil
}

// ClientCredentialsAuthnResponseHandler retries a request after gettting a new access token from the identity provider using the Client Credentials flow
type ClientCredentialsAuthnResponseHandler struct {
	*AuthnResponseHandler
	ClientID     string
	ClientSecret string
	Scope        string
}

// NewClientCredentialsAuthnResponseHandler initializes a new response handler
func NewClientCredentialsAuthnResponseHandler(idpHost string, clientID string, clientSecret string, scope string) *ClientCredentialsAuthnResponseHandler {
	handler := &ClientCredentialsAuthnResponseHandler{
		AuthnResponseHandler: &AuthnResponseHandler{
			IdpClient: idp.NewClient(idpHost, "", "", "", ""),
		},
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}
	handler.TokenRetriever = handler
	return handler
}

// GetAccessToken gets a new access token from the identity provider
func (rh *ClientCredentialsAuthnResponseHandler) GetAccessToken() (token string, err error) {
	ctx, err := rh.IdpClient.ClientFlow(rh.ClientID, rh.ClientSecret, rh.Scope)
	if err != nil {
		return "", err
	}
	return ctx.AccessToken, nil
}
