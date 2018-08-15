// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package handler

import (
	"net/http"

	"github.com/splunk/ssc-client-go/idp"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/util"
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
	// Replace the access token in the request's Authorization: Bearer header
	request.UpdateToken(token)
	// Re-initialize body (otherwise body is empty)
	body, err := request.GetBody()
	request.Body = body
	// Update the client such that future requests will use the new access token
	client.UpdateToken(token)
	// Retry the request with the updated token
	return client.Do(request)
}

// RefreshTokenAuthnResponseHandler retries a request after gettting a new access token from the identity provider using a RefreshToken
type RefreshTokenAuthnResponseHandler struct {
	*AuthnResponseHandler
	// ClientID which corresponds to an Refresh Token ("offline_access" scope) supported IdP client
	ClientID string
	// Scope(s) to request, separated by spaces -- "openid email profile" is recommended for individual users
	Scope string
	// RefreshToken to use to authenticate in order to generate an access token
	RefreshToken *util.Credential
}

// NewRefreshTokenAuthnResponseHandler initializes a new response handler
func NewRefreshTokenAuthnResponseHandler(idpHost string, clientID string, scope string, refreshToken string) *RefreshTokenAuthnResponseHandler {
	handler := &RefreshTokenAuthnResponseHandler{
		AuthnResponseHandler: &AuthnResponseHandler{IdpClient: idp.NewDefaultClient(idpHost)},
		ClientID:             clientID,
		Scope:                scope,
		RefreshToken:         &util.Credential{refreshToken},
	}
	handler.TokenRetriever = handler
	return handler
}

// GetAccessToken gets a new access token from the identity provider
func (rh *RefreshTokenAuthnResponseHandler) GetAccessToken() (token string, err error) {
	ctx, err := rh.IdpClient.Refresh(rh.ClientID, rh.Scope, rh.RefreshToken.ClearText())
	if err != nil {
		return "", err
	}
	return ctx.AccessToken, nil
}

// ClientCredentialsAuthnResponseHandler retries a request after gettting a new access token from the identity provider using the Client Credentials flow
type ClientCredentialsAuthnResponseHandler struct {
	*AuthnResponseHandler
	// ClientID to authenticate as which corresponds to a Client Credentials flow supported IdP client
	ClientID string
	// ClientSecret corresponding to the ClientID above
	ClientSecret *util.Credential
	// Scope(s) to request, separated by spaces -- this will be a custom scope, for example: "backend_service"
	Scope string
}

// NewClientCredentialsAuthnResponseHandler initializes a new response handler
func NewClientCredentialsAuthnResponseHandler(idpHost string, clientID string, clientSecret string, scope string) *ClientCredentialsAuthnResponseHandler {
	handler := &ClientCredentialsAuthnResponseHandler{
		AuthnResponseHandler: &AuthnResponseHandler{
			IdpClient: idp.NewDefaultClient(idpHost),
		},
		ClientID:     clientID,
		ClientSecret: &util.Credential{clientSecret},
		Scope:        scope,
	}
	handler.TokenRetriever = handler
	return handler
}

// GetAccessToken gets a new access token from the identity provider
func (rh *ClientCredentialsAuthnResponseHandler) GetAccessToken() (token string, err error) {
	ctx, err := rh.IdpClient.ClientFlow(rh.ClientID, rh.ClientSecret.ClearText(), rh.Scope)
	if err != nil {
		return "", err
	}
	return ctx.AccessToken, nil
}

// PKCEAuthnResponseHandler retries a request after gettting a new access token from the identity provider using the Proof Key for Code Exchange (PKCE) flow
type PKCEAuthnResponseHandler struct {
	*AuthnResponseHandler
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

// NewPKCEAuthnResponseHandler initializes a new response handler
func NewPKCEAuthnResponseHandler(idpHost string, clientID string, redirectURI string, scope string, username string, password string) *PKCEAuthnResponseHandler {
	handler := &PKCEAuthnResponseHandler{
		AuthnResponseHandler: &AuthnResponseHandler{
			IdpClient: idp.NewDefaultClient(idpHost),
		},
		ClientID:    clientID,
		RedirectURI: redirectURI,
		Scope:       scope,
		Username:    username,
		Password:    &util.Credential{password},
	}
	handler.TokenRetriever = handler
	return handler
}

// GetAccessToken gets a new access token from the identity provider
func (rh *PKCEAuthnResponseHandler) GetAccessToken() (token string, err error) {
	ctx, err := rh.IdpClient.PKCEFlow(rh.ClientID, rh.RedirectURI, rh.Scope, rh.Username, rh.Password.ClearText())
	if err != nil {
		return "", err
	}
	return ctx.AccessToken, nil
}
