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
	"github.com/pkg/errors"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const (
	// SplunkCloudIdpHost is the default identity provider host for Splunk Cloud
	SplunkCloudIdpHost = "https://auth.scp.splunk.com"
)

// TokenRetriever retrieves an access token with context
type TokenRetriever interface {
	GetTokenContext() (*Context, error)
}

// NoOpTokenRetriever just returns the same static Context
type NoOpTokenRetriever struct {
	Context *Context
}

// GetTokenContext just returns the same static Context
func (tr *NoOpTokenRetriever) GetTokenContext() (*Context, error) {
	return tr.Context, nil
}

// makeClient creates an *idp.Client
func makeClient(idpHost string, insecure bool) *Client {
	if idpHost == "" {
		idpHost = SplunkCloudIdpHost
	}
	return NewClient(
		idpHost,
		defaultAuthnPath,
		defaultAuthorizePath,
		defaultTokenPath,
		insecure)
}

// RefreshTokenRetriever retries a request after getting a new access token from the identity provider using a RefreshToken
type RefreshTokenRetriever struct {
	*Client
	// ClientID which corresponds to an Refresh Token ("offline_access" scope) supported IdP client
	ClientID string
	// Scope(s) to request, separated by spaces -- "openid email profile" is recommended for individual users
	Scope string
	// RefreshToken to use to authenticate in order to generate an access token
	RefreshToken *util.Credential
}

// NewRefreshTokenRetriever initializes a new token context retriever
//   idpURL: should be of the form https://example.com or optionally https://example.com:port
//     - if "" is specified then SplunkCloudIdpURL will be used.
func NewRefreshTokenRetriever(clientID string, scope string, refreshToken string, idpHost string) *RefreshTokenRetriever {
	return &RefreshTokenRetriever{
		Client:       makeClient(idpHost, false),
		ClientID:     clientID,
		Scope:        scope,
		RefreshToken: util.NewCredential(refreshToken),
	}
}

// GetTokenContext gets a new access token context from the identity provider
func (tr *RefreshTokenRetriever) GetTokenContext() (*Context, error) {
	ctx, err := tr.Refresh(tr.ClientID, tr.Scope, tr.RefreshToken.ClearText())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token in refresh token flow")
	}
	return ctx, nil
}

// ClientCredentialsRetriever retries a request after gettting a new access token from the identity provider using the Client Credentials flow
type ClientCredentialsRetriever struct {
	*Client
	// ClientID to authenticate as which corresponds to a Client Credentials flow supported IdP client
	ClientID string
	// ClientSecret corresponding to the ClientID above
	ClientSecret *util.Credential
	// Scope(s) to request, separated by spaces -- this will be a custom scope, for example: "backend_service"
	Scope string
}

// NewClientCredentialsRetriever initializes a new token context retriever
//   idpURL: should be of the form https://example.com or optionally https://example.com:port
//     - if "" is specified then SplunkCloudIdpURL will be used.
func NewClientCredentialsRetriever(clientID string, clientSecret string, scope string, idpHost string) *ClientCredentialsRetriever {
	return &ClientCredentialsRetriever{
		Client:       makeClient(idpHost, false),
		ClientID:     clientID,
		ClientSecret: util.NewCredential(clientSecret),
		Scope:        scope,
	}
}

// GetTokenContext gets a new access token context from the identity provider
func (tr *ClientCredentialsRetriever) GetTokenContext() (*Context, error) {
	ctx, err := tr.ClientFlow(tr.ClientID, tr.ClientSecret.ClearText(), tr.Scope)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token in client credentials flow")
	}
	return ctx, nil
}

// PKCERetriever retries a request after gettting a new access token from the identity provider using the Proof Key for Code Exchange (PKCE) flow
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

// NewPKCERetriever initializes a new token context retriever
//   idpURL: should be of the form https://example.com or optionally https://example.com:port
//     - if "" is specified then SplunkCloudIdpURL will be used.
func NewPKCERetriever(clientID string, redirectURI string, scope string, username string, password string, idpHost string) *PKCERetriever {
	return &PKCERetriever{
		Client:      makeClient(idpHost, false),
		ClientID:    clientID,
		RedirectURI: redirectURI,
		Scope:       scope,
		Username:    username,
		Password:    util.NewCredential(password),
	}
}

// GetTokenContext gets a new access token context from the identity provider
func (tr *PKCERetriever) GetTokenContext() (*Context, error) {
	ctx, err := tr.PKCEFlow(tr.ClientID, tr.RedirectURI, tr.Scope, tr.Username, tr.Password.ClearText())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token in PKCE flow")
	}
	return ctx, nil
}
