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

package services

import (
	"fmt"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/v2/idp"
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
