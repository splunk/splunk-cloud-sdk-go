// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package services

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	defaultAPIClusterURL = "https://api.splunkbeta.com"
	defaultAPPClusterURL = "https://app.splunkbeta.com"
)

func TestNewClientDefaultProductionClient(t *testing.T) {
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	client, err := NewClient(&Config{
		Token: token,
	})
	require.Nil(t, err)
	require.Equal(t, client.GetURL("api").String(), defaultAPIClusterURL)
	require.Equal(t, client.GetURL("app").String(), defaultAPPClusterURL)
}

func TestBuildURLDefaultProductionClient(t *testing.T) {
	token := "EXAMPLE_AUTHENTICATION_TOKEN"
	tenant := "system"
	client, err := NewClient(&Config{
		Token:  token,
		Tenant: tenant,
	})
	require.Nil(t, err)
	testService := "myservice"
	testVersion := "v1beta1"
	testEndpoint := "widgets"
	testURL, err := client.BuildURL(nil, "", testService, testVersion, testEndpoint)
	require.Nil(t, err)
	expectedURL := fmt.Sprintf("%s/%s/%s/%s/%s", defaultAPIClusterURL, tenant, testService, testVersion, testEndpoint)
	require.Equal(t, testURL.String(), expectedURL)
}

func TestBuildAPIClusterCustomURL(t *testing.T) {
	var apiURLProtocol = "http"
	var apiHostname = "apiexample.com"
	var apiPort = "8882"
	var apiHostWithPort = fmt.Sprintf("%s%s%s", apiHostname, ":", apiPort)
	var urls = map[string]string{
		"api": apiHostWithPort,
	}
	var tenant = "EXAMPLE_TENANT"
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	client, err := NewClient(&Config{
		Token:  token,
		Scheme: apiURLProtocol,
		URLs:   urls,
		Tenant: tenant,
	})
	require.Nil(t, err)
	assert.Equal(t, client.httpClient.Timeout, time.Second*5, "default timeout should be 5 seconds")
	testURL, err := client.BuildURL(nil, "api", "services", "search", "jobs")
	require.Nil(t, err)
	assert.Equal(t, apiHostname, testURL.Hostname())
	assert.Equal(t, apiURLProtocol, testURL.Scheme)
	assert.Equal(t, apiPort, testURL.Port())
	assert.Equal(t, apiHostWithPort, testURL.Host)
	assert.Equal(t, fmt.Sprintf("%s%s", tenant, "/services/search/jobs"), testURL.Path)
	assert.Empty(t, testURL.Fragment)
}

func TestBuildAPPClusterCustomURL(t *testing.T) {
	var appURLProtocol = "http"
	var appHostname = "appexample.com"
	var appPort = "8882"
	var appHostWithPort = fmt.Sprintf("%s%s%s", appHostname, ":", appPort)
	var urls = map[string]string{
		"app": appHostWithPort,
	}
	var tenant = "EXAMPLE_TENANT"
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	client, err := NewClient(&Config{
		Token:  token,
		Scheme: appURLProtocol,
		URLs:   urls,
		Tenant: tenant,
	})
	require.Nil(t, err)
	assert.Equal(t, client.httpClient.Timeout, time.Second*5, "default timeout should be 5 seconds")
	testURL, err := client.BuildURL(nil, "app", "services", "search", "jobs")
	require.Nil(t, err)
	assert.Equal(t, appHostname, testURL.Hostname())
	assert.Equal(t, appURLProtocol, testURL.Scheme)
	assert.Equal(t, appPort, testURL.Port())
	assert.Equal(t, appHostWithPort, testURL.Host)
	assert.Equal(t, fmt.Sprintf("%s%s", tenant, "/services/search/jobs"), testURL.Path)
	assert.Empty(t, testURL.Fragment)
}

func TestBuildURLWithHostAndURLs(t *testing.T) {
	var hostname = "appexample.com"
	var port = "8882"
	var host = fmt.Sprintf("%s%s%s", hostname, ":", port)
	var urls = map[string]string{
		"app": host,
	}
	var tenant = "EXAMPLE_TENANT"
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	_, err := NewClient(&Config{
		URLs:   urls,
		Host:   "localhost.com",
		Token:  token,
		Tenant: tenant,
	})
	assert.NotNil(t, err)
	assert.Equal(t, "either URLs or Host must be set, not both. URLs are prefferred since Host will be depreciated", err.Error())
}

func TestBuildURLEscapedCharacters(t *testing.T) {
	client, err := NewClient(&Config{
		Token:  "TEST_TOKEN",
		Tenant: "mytenant",
	})
	require.Nil(t, err)

	query := url.Values{}
	query.Set("filter", `kind=="import"`)
	query.Set("email", "user@example.com")

	testURL, err := client.BuildURL(query, "api", "permissions", "mytenant:*:*write")
	require.Nil(t, err)
	assert.Equal(t, "https://api.splunkbeta.com/mytenant/permissions/mytenant:%2A:%2Awrite?email=user%40example.com&filter=kind%3D%3D%22import%22", testURL.String())
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

func TestHostOnlyClient(t *testing.T) {
	var apiURLProtocol = "http"
	var apiPort = "8882"
	var apiHostname = "example.com"
	var clusterAPIHostname = "api." + apiHostname
	var apiHost = apiHostname + ":" + apiPort
	var clusterAPIHost = "api." + apiHost
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

	NoClusterMatchFromURLs := client.GetURL("")
	assert.Equal(t, clusterAPIHostname, NoClusterMatchFromURLs.Hostname())
	assert.Equal(t, apiURLProtocol, NoClusterMatchFromURLs.Scheme)
	assert.Equal(t, apiPort, NoClusterMatchFromURLs.Port())
	assert.Equal(t, clusterAPIHost, NoClusterMatchFromURLs.Host)

	FoundClusterMatchFromURLs := client.GetURL("api")
	assert.Equal(t, "api.splunkbeta.com", FoundClusterMatchFromURLs.Hostname())
	assert.Equal(t, apiURLProtocol, FoundClusterMatchFromURLs.Scheme)
	assert.Equal(t, "", FoundClusterMatchFromURLs.Port())
	assert.Equal(t, "api.splunkbeta.com", FoundClusterMatchFromURLs.Host)

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
