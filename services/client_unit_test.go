// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package services

import (
	"fmt"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildURLDefaultTenant(t *testing.T) {
	var apiURLProtocol = "http"
	var apiPort = "8882"
	var host = "example.com"
	var hostWithPort = host + ":" + apiPort
	var tenant = "EXAMPLE_TENANT"
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	client, err := NewClient(&Config{
		Token:  token,
		Scheme: apiURLProtocol,
		Host:   hostWithPort,
		Tenant: tenant,
	})
	require.Nil(t, err)
	assert.Equal(t, client.httpClient.Timeout, time.Second*5, "default timeout should be 5 seconds")
	testURL, err := client.BuildURL(nil, "api", "services", "search", "jobs")

	require.Nil(t, err)
	apiHostName := fmt.Sprintf("%s%s%s", "api", ".", host)
	apiHostWithPort := fmt.Sprintf("%s%s%s", apiHostName, ":", apiPort)
	assert.Equal(t, apiHostName, testURL.Hostname())
	assert.Equal(t, apiURLProtocol, testURL.Scheme)
	assert.Equal(t, apiPort, testURL.Port())
	assert.Equal(t, apiHostWithPort, testURL.Host)
	assert.Equal(t, fmt.Sprintf("%s%s", tenant, "/services/search/jobs"), testURL.Path)
	assert.Empty(t, testURL.Fragment)
}

func TestBuildURLSetDefaultTenant(t *testing.T) {
	var tenant = "EXAMPLE_TENANT"
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	client, err := NewClient(&Config{
		Token:  token,
		Tenant: tenant,
	})
	require.Nil(t, err)
	testURL, err := client.BuildURL(nil, "api", "services", "search", "jobs")
	require.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%s%s", tenant, "/services/search/jobs"), testURL.Path)
	assert.Empty(t, testURL.Fragment)
	// Set to new tenant
	tenant = "NEW_TENANT"
	client.SetDefaultTenant(tenant)
	testURL, err = client.BuildURL(nil, "api", "services", "search", "jobs")
	require.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%s%s", tenant, "/services/search/jobs"), testURL.Path)
	assert.Empty(t, testURL.Fragment)
}

func TestNewTokenClient(t *testing.T) {
	var apiURLProtocol = "http"
	var apiPort = "8882"
	var apiHostname = "example.com"
	var apiHost = apiHostname + ":" + apiPort
	var tenant = "EXAMPLE_TENANT"
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	var timeout = time.Second * 8
	var client, err = NewClient(&Config{
		Token:   token,
		Scheme:  apiURLProtocol,
		Host:    apiHost,
		Tenant:  tenant,
		Timeout: timeout,
	})
	require.Nil(t, err)
	assert.Equal(t, token, client.tokenContext.AccessToken)

	testURL := client.GetURL()
	assert.Equal(t, apiHostname, testURL.Hostname())
	assert.Equal(t, apiURLProtocol, testURL.Scheme)
	assert.Equal(t, apiPort, testURL.Port())
	assert.Equal(t, apiHost, testURL.Host)
}

type tRet struct{}

const xyzToken = "X.Y.Z"

func (tr *tRet) GetTokenContext() (*idp.Context, error) {
	return &idp.Context{AccessToken: xyzToken}, nil
}

func TestNewTokenRetrieverClient(t *testing.T) {
	var tokenRetriever = &tRet{}
	var client, err = NewClient(&Config{TokenRetriever: tokenRetriever})
	require.Nil(t, err)
	assert.Equal(t, client.tokenContext.AccessToken, xyzToken, "access token should have been initialized to X.Y.Z")
}

func TestNewClientTokenAndTokenRetriever(t *testing.T) {
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	var tokenRetriever = &tRet{}
	var _, err = NewClient(&Config{Token: token, TokenRetriever: tokenRetriever})
	// This should fail, users should specify Token or TokenRetriever, not both
	assert.NotNil(t, err)
}
