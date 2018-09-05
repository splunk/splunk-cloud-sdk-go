// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"fmt"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/stretchr/testify/assert"
)

const retrievedToken = "MY.RETRIEVED.TOKEN"

type testTokenRetriever struct{}

func (tr *testTokenRetriever) GetTokenContext() (*idp.Context, error) {
	return &idp.Context{AccessToken: retrievedToken}, nil
}

type testErrTokenRetriever struct{}

func (tr *testErrTokenRetriever) GetTokenContext() (*idp.Context, error) {
	return nil, fmt.Errorf("no luck with that token")
}

func TestAuthnResponseHandlerGoodToken(t *testing.T) {
	rh := AuthnResponseHandler{
		TokenRetriever: &testTokenRetriever{},
	}
	ctx, err := rh.TokenRetriever.GetTokenContext()
	assert.Nil(t, err)
	assert.NotNil(t, ctx)
	assert.Equal(t, ctx.AccessToken, retrievedToken)
}

func TestAuthnResponseHandlerErrorToken(t *testing.T) {
	rh := AuthnResponseHandler{
		TokenRetriever: &testErrTokenRetriever{},
	}
	ctx, err := rh.TokenRetriever.GetTokenContext()
	assert.NotNil(t, err)
	assert.Nil(t, ctx)
}
