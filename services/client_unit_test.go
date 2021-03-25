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
	"net/url"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const defaultCloudURL = "https://api.scp.splunk.com"

func TestNewClientDefaultProductionClient(t *testing.T) {
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	client, err := NewClient(&Config{
		Token: token,
	})
	require.NoError(t, err)
	baseURL := client.GetURL("")
	require.Equal(t, baseURL.String(), defaultCloudURL)
	apiURL := client.GetURL("api")
	require.Equal(t, apiURL.String(), defaultCloudURL)
}

func TestNewClientAppClusterClient(t *testing.T) {
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	client, err := NewClient(&Config{
		Token: token,
	})
	require.NoError(t, err)
	appURL := client.GetURL("app")
	require.Equal(t, appURL.String(), "https://app.scp.splunk.com")
}

func TestBuildURLDefaultProductionClient(t *testing.T) {
	token := "EXAMPLE_AUTHENTICATION_TOKEN"
	tenant := "system"
	client, err := NewClient(&Config{
		Token:  token,
		Tenant: tenant,
	})
	require.NoError(t, err)
	testService := "myservice"
	testVersion := "v1beta1"
	testEndpoint := "widgets"
	testURL, err := client.BuildURL(nil, "", testService, testVersion, testEndpoint)
	require.NoError(t, err)
	expectedURL := fmt.Sprintf("%s/%s/%s/%s/%s", defaultCloudURL, tenant, testService, testVersion, testEndpoint)
	require.Equal(t, testURL.String(), expectedURL)
}

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
	require.NoError(t, err)
	assert.Equal(t, client.httpClient.Timeout, time.Second*5, "default timeout should be 5 seconds")
	testURL, err := client.BuildURL(nil, "api", "services", "search", "jobs")

	require.NoError(t, err)
	apiHostName := fmt.Sprintf("api.%s", host)
	apiHostWithPort := fmt.Sprintf("%s:%s", apiHostName, apiPort)
	assert.Equal(t, apiHostName, testURL.Hostname())
	assert.Equal(t, apiURLProtocol, testURL.Scheme)
	assert.Equal(t, apiPort, testURL.Port())
	assert.Equal(t, apiHostWithPort, testURL.Host)
	assert.Equal(t, fmt.Sprintf("%s/services/search/jobs", tenant), testURL.Path)
	assert.Empty(t, testURL.Fragment)

	client.SetOverrideHost("localhost")
	assert.Equal(t, "localhost", client.overrideHost)
	testURL, err = client.BuildURL(nil, "api", "services", "search", "jobs")
	require.NoError(t, err)
	assert.Equal(t, "localhost", testURL.Hostname())
	assert.Equal(t, apiURLProtocol, testURL.Scheme)
	// "localhost" has no port specified, so port is now "" rather than 8882
	assert.Equal(t, "", testURL.Port())
	assert.Equal(t, "localhost", testURL.Host)
	assert.Equal(t, fmt.Sprintf("%s/services/search/jobs", tenant), testURL.Path)
	assert.Empty(t, testURL.Fragment)

	client.SetOverrideHost("127.0.0.1:8080")
	assert.Equal(t, "127.0.0.1:8080", client.overrideHost)
	testURL, err = client.BuildURL(nil, "api", "services", "search", "jobs")
	require.NoError(t, err)
	assert.Equal(t, "127.0.0.1", testURL.Hostname())
	assert.Equal(t, apiURLProtocol, testURL.Scheme)
	assert.Equal(t, "8080", testURL.Port())
	assert.Equal(t, "127.0.0.1:8080", testURL.Host)
	assert.Equal(t, fmt.Sprintf("%s/services/search/jobs", tenant), testURL.Path)
	assert.Empty(t, testURL.Fragment)
}

func TestBuildURLPathParams(t *testing.T) {
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
	require.NoError(t, err)
	pp := struct {
		WidgetID     int
		SprocketName string
	}{
		WidgetID:     1234,
		SprocketName: "spok",
	}
	u, err := client.BuildURLFromPathParams(nil, "api", `/myservice/v1beta3/widgets/{{.WidgetID}}/sprockets/{{.SprocketName}}`, pp)
	require.NoError(t, err)
	require.NotEmpty(t, u)
	assert.Equal(t, `http://api.example.com:8882/EXAMPLE_TENANT/myservice/v1beta3/widgets/1234/sprockets/spok`, u.String())
}

func TestNewRequestClientVersion(t *testing.T) {
	var clientVersion = "exampleClient/2.0.0"
	var httpMethod = "GET"
	var url = "https://api.staging.scp.splunk.com:443/mytenant/catalog/v2beta1/rules?count=2"
	var expectedHTTPSplunkClient = fmt.Sprintf("%s/%s,%s", UserAgent, Version, clientVersion)
	client, err := NewClient(&Config{
		Token:         "MY TOKEN",
		Tenant:        "mytenant",
		ClientVersion: clientVersion,
	})
	require.NoError(t, err)
	request, _ := client.NewRequest(httpMethod, url, nil, nil)
	require.Equal(t, request.Header.Get("Splunk-Client"), expectedHTTPSplunkClient)
}

func TestNewClientOverrideHost(t *testing.T) {
	client, err := NewClient(&Config{
		Token:        "MY TOKEN",
		OverrideHost: "localhost:8080",
		Tenant:       "mytenant",
	})
	assert.NoError(t, err)
	assert.Equal(t, "localhost:8080", client.overrideHost)
	testURL, err := client.BuildURL(nil, "api", "services", "search", "jobs")
	require.NoError(t, err)
	assert.Equal(t, "localhost", testURL.Hostname())
	assert.Equal(t, "https", testURL.Scheme)
	assert.Equal(t, "8080", testURL.Port())
	assert.Equal(t, "localhost:8080", testURL.Host)
	assert.Equal(t, "mytenant/services/search/jobs", testURL.Path)
	assert.Empty(t, testURL.Fragment)
}

func TestNewClientHostAndOverrideHost(t *testing.T) {
	_, err := NewClient(&Config{
		Token:        "MY TOKEN",
		Host:         "scp.splunk.com",
		OverrideHost: "localhost:8080",
	})
	assert.Equal(t, "either config.Host or config.OverrideHost may be set, setting both is invalid", err.Error())
}

func TestBuildURLEscapedCharacters(t *testing.T) {
	client, err := NewClient(&Config{
		Token:  "TEST_TOKEN",
		Tenant: "mytenant",
	})
	require.NoError(t, err)

	query := url.Values{}
	query.Set("filter", `kind=="import"`)
	query.Set("email", "user@example.com")

	testURL, err := client.BuildURL(query, "api", "permissions", "mytenant:*:*write")
	require.NoError(t, err)
	assert.Equal(t, "https://api.scp.splunk.com/mytenant/permissions/mytenant:%2A:%2Awrite?email=user%40example.com&filter=kind%3D%3D%22import%22", testURL.String())
}

func TestBuildURLSetDefaultTenant(t *testing.T) {
	var tenant = "EXAMPLE_TENANT"
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	client, err := NewClient(&Config{
		Token:  token,
		Tenant: tenant,
	})
	require.NoError(t, err)
	testURL, err := client.BuildURL(nil, "api", "services", "search", "jobs")
	require.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s%s", tenant, "/services/search/jobs"), testURL.Path)
	assert.Empty(t, testURL.Fragment)
	// Set to new tenant
	tenant = "NEW_TENANT"
	client.SetDefaultTenant(tenant)
	testURL, err = client.BuildURL(nil, "api", "services", "search", "jobs")
	require.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%s%s", tenant, "/services/search/jobs"), testURL.Path)
	assert.Empty(t, testURL.Fragment)
}

func TestNewTokenClient(t *testing.T) {
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
	require.NoError(t, err)
	assert.Equal(t, token, client.tokenContext.AccessToken)

	testURL := client.GetURL("")
	assert.Equal(t, clusterAPIHostname, testURL.Hostname())
	assert.Equal(t, apiURLProtocol, testURL.Scheme)
	assert.Equal(t, apiPort, testURL.Port())
	assert.Equal(t, clusterAPIHost, testURL.Host)
}

func TestBuildURLWithTenantScopedEndpointClient(t *testing.T) {
	token := "EXAMPLE_AUTHENTICATION_TOKEN"
	host := "myenv.scs.splunk.com"
	tenant := "non-system"
	cluster := "api"
	client, err := NewClient(&Config{
		Token:        token,
		Tenant:       tenant,
		TenantScoped: true,
		Host:         host,
	})
	require.NoError(t, err)
	testService := "myservice"
	testVersion := "v2beta1"
	testEndpoint := "widgets"
	testURL, err := client.BuildURLFromPathParams(nil, "api", `/myservice/v2beta1/widgets`, nil)
	require.NoError(t, err)
	expectedURL := fmt.Sprintf("https://%s.%s.%s/%s/%s/%s/%s", tenant, cluster, host, tenant, testService, testVersion, testEndpoint)
	fmt.Println("Actual URL")
	fmt.Println(testURL.String())
	fmt.Println("Expected URL")
	fmt.Println(expectedURL)
	require.Equal(t, testURL.String(), expectedURL)
}

func TestBuildURLWithRegionScopedSystemEndpointClient(t *testing.T) {
	token := "EXAMPLE_AUTHENTICATION_TOKEN"
	host := "myenv.scs.splunk.com"
	tenant := "system"
	cluster := "api"
	region := "region10"
	client, err := NewClient(&Config{
		Token:        token,
		Tenant:       tenant,
		TenantScoped: true,
		Host:         host,
		Region:       region,
	})
	require.NoError(t, err)
	testService := "myservice"
	testVersion := "v2beta1"
	testEndpoint := "widgets"
	testURL, err := client.BuildURLFromPathParams(nil, "api", `/system/myservice/v2beta1/widgets`, nil)
	require.NoError(t, err)
	expectedURL := fmt.Sprintf("https://region-%s.%s.%s/%s/%s/%s/%s", region, cluster, host, tenant, testService, testVersion, testEndpoint)
	fmt.Println("Actual URL")
	fmt.Println(testURL.String())
	fmt.Println("Expected URL")
	fmt.Println(expectedURL)
	require.Equal(t, testURL.String(), expectedURL)
}

func TestBuildURLWithMissingRegionSystemEndpointClient(t *testing.T) {
	token := "EXAMPLE_AUTHENTICATION_TOKEN"
	host := "myenv.scs.splunk.com"
	tenant := "system"
	client, err := NewClient(&Config{
		Token:        token,
		Tenant:       tenant,
		TenantScoped: true,
		Host:         host,
	})
	require.NoError(t, err)
	_, err = client.BuildURLFromPathParams(nil, "api", `/system/myservice/v2beta1/widgets`, nil)
	fmt.Print(err)
	require.Error(t, err)
}

func TestBuildURLWitOverrideHostClient(t *testing.T) {
	token := "EXAMPLE_AUTHENTICATION_TOKEN"
	host := "custom.scp.splunk.com"
	tenant := "non-system"
	client, err := NewClient(&Config{
		Token:        token,
		Tenant:       tenant,
		TenantScoped: true,
		OverrideHost: host,
	})
	require.NoError(t, err)
	testService := "myservice"
	testVersion := "v2beta1"
	testEndpoint := "widgets"
	testURL, err := client.BuildURLFromPathParams(nil, "api", `/myservice/v2beta1/widgets`, nil)
	require.NoError(t, err)
	expectedURL := fmt.Sprintf("https://%s/%s/%s/%s/%s", host, tenant, testService, testVersion, testEndpoint)
	fmt.Println("Actual URL")
	fmt.Println(testURL.String())
	fmt.Println("Expected URL")
	fmt.Println(expectedURL)
	require.Equal(t, testURL.String(), expectedURL)
}

func TestBuildURLWithDefaultTenantScopedClient(t *testing.T) {
	token := "EXAMPLE_AUTHENTICATION_TOKEN"
	tenant := "non-system"
	client, err := NewClient(&Config{
		Token:        token,
		Tenant:       tenant,
		TenantScoped: true,
	})
	require.NoError(t, err)
	testService := "myservice"
	testVersion := "v2beta1"
	testEndpoint := "widgets"
	testURL, err := client.BuildURLFromPathParams(nil, "api", `/myservice/v2beta1/widgets`, nil)
	require.NoError(t, err)
	expectedURL := fmt.Sprintf("https://%s.%s/%s/%s/%s/%s", tenant, "api.scp.splunk.com", tenant, testService, testVersion, testEndpoint)
	fmt.Println("Actual URL")
	fmt.Println(testURL.String())
	fmt.Println("Expected URL")
	fmt.Println(expectedURL)
	require.Equal(t, testURL.String(), expectedURL)
}

type tRet struct{}

const xyzToken = "X.Y.Z"

func (tr *tRet) GetTokenContext() (*idp.Context, error) {
	return &idp.Context{AccessToken: xyzToken}, nil
}

func TestNewTokenRetrieverClient(t *testing.T) {
	var tokenRetriever = &tRet{}
	var client, err = NewClient(&Config{TokenRetriever: tokenRetriever})
	require.NoError(t, err)
	assert.Equal(t, client.tokenContext.AccessToken, xyzToken, "access token should have been initialized to X.Y.Z")
}

func TestNewClientTokenAndTokenRetriever(t *testing.T) {
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	var tokenRetriever = &tRet{}
	var _, err = NewClient(&Config{Token: token, TokenRetriever: tokenRetriever})
	// This should fail, users should specify Token or TokenRetriever, not both
	assert.NotNil(t, err)
}
