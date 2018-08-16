// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package handler

import (
	"fmt"
	"testing"

	"github.com/splunk/ssc-client-go/idp"
	"github.com/stretchr/testify/assert"
)

const retrievedToken = "MY.RETRIEVED.TOKEN"

type testTokenRetriever struct{}

func (tr *testTokenRetriever) GetAccessToken() (token string, err error) {
	return retrievedToken, nil
}

type testErrTokenRetriever struct{}

func (tr *testErrTokenRetriever) GetAccessToken() (token string, err error) {
	return "", fmt.Errorf("no luck with that token")
}

func TestAuthnResponseHandlerGoodToken(t *testing.T) {
	rh := AuthnResponseHandler{
		IdpClient:      idp.NewDefaultClient("myhost"),
		TokenRetriever: &testTokenRetriever{},
	}
	tok, err := rh.TokenRetriever.GetAccessToken()
	assert.Nil(t, err)
	assert.Equal(t, tok, retrievedToken)
}

func TestAuthnResponseHandlerErrorToken(t *testing.T) {
	rh := AuthnResponseHandler{
		IdpClient:      idp.NewDefaultClient("myhost"),
		TokenRetriever: &testErrTokenRetriever{},
	}
	tok, err := rh.TokenRetriever.GetAccessToken()
	assert.NotNil(t, err)
	assert.Empty(t, tok)
}

func TestRefreshTokenAuthnResponseHandler(t *testing.T) {
	rh := NewRefreshTokenAuthnResponseHandler("myhost", "myclientid", "myscope", "myrefreshtoken")
	rh.TokenRetriever = &testTokenRetriever{}
	tok, err := rh.TokenRetriever.GetAccessToken()
	assert.Nil(t, err)
	assert.Equal(t, tok, retrievedToken)
}

func TestClientCredentialsAuthnResponseHandler(t *testing.T) {
	rh := NewClientCredentialsAuthnResponseHandler("myhost", "myclientid", "myclientsecret", "myscope")
	rh.TokenRetriever = &testTokenRetriever{}
	tok, err := rh.TokenRetriever.GetAccessToken()
	assert.Nil(t, err)
	assert.Equal(t, tok, retrievedToken)
}

func TestPKCEAuthnResponseHandler(t *testing.T) {
	rh := NewPKCEAuthnResponseHandler("myhost", "myclientid", "myredirect", "myscope", "myusername", "mypassword")
	rh.TokenRetriever = &testTokenRetriever{}
	tok, err := rh.TokenRetriever.GetAccessToken()
	assert.Nil(t, err)
	assert.Equal(t, tok, retrievedToken)
}
