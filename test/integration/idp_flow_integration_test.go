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

package integration

import (
	"os"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"

	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/services/identity"
	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//Tenantscoped host
var IdpHostTenantScoped = os.Getenv("IDP_HOST_TENANT_SCOPED")

// NativeClientID - App Registry Client Id for SDK Native App
var NativeClientID = os.Getenv("REFRESH_TOKEN_CLIENT_ID")

// NativeAppRedirectURI is one of the redirect uris configured for the app
const NativeAppRedirectURI = "https://localhost:8000"

// BackendClientID - App Registry Client id for client credentials flow tenant scoped
var BackendClientIDTenantScoped = os.Getenv("BACKEND_CLIENT_ID_TENANT_SCOPED")

// BackendClientSecret - App Registry Client secret for client credentials flow tenant scoped
var BackendClientSecretTenantScoped = os.Getenv("BACKEND_CLIENT_SECRET_TENANT_SCOPED")

// BackendServiceScope - scope for obtaining access token for client credentials flow
const BackendServiceScope = ""

// TestUsername corresponds to the test user for integration testing
var TestUsername = os.Getenv("TEST_USERNAME")

// TestPassword corresponds to the test user's password for integration testing
var TestPassword = os.Getenv("TEST_PASSWORD")

type retryTokenRetriever struct {
	TR idp.TokenRetriever
	n  int
}

func (r *retryTokenRetriever) GetTokenContext() (*idp.Context, error) {
	r.n++
	// Return a bad access token the first time for testing 401 retry logic
	if r.n == 1 {
		return &idp.Context{AccessToken: testutils.ExpiredAuthenticationToken}, nil
	}
	// For subsequent requests get the real token using the real TokenRetriever
	return r.TR.GetTokenContext()
}

type badTokenRetriever struct {
	N int
}

func (r *badTokenRetriever) GetTokenContext() (*idp.Context, error) {
	r.N++
	// Return a bad access token every time
	return &idp.Context{AccessToken: testutils.ExpiredAuthenticationToken}, nil
}

// TestIntegrationRefreshTokenInitWorkflow tests initializing the client with a TokenRetriever impleme
func TestIntegrationRefreshTokenInitWorkflow(t *testing.T) {
	hostURLConfig := idp.HostURLConfig{TenantScoped: true, Tenant: testutils.TestTenant, Region: testutils.TestRegion}
	// get a new refresh token
	tr := idp.NewPKCERetriever(NativeClientID, NativeAppRedirectURI, idp.DefaultRefreshScope, TestUsername, TestPassword, IdpHostTenantScoped, "", hostURLConfig)
	ctx, err := tr.GetTokenContext()
	require.Emptyf(t, err, "Error validating using access token generated from PKCE flow: %s", err)
	require.NotNil(t, ctx)

	tr_refresh := idp.NewRefreshTokenRetriever(NativeClientID, idp.DefaultRefreshScope, ctx.RefreshToken, IdpHostTenantScoped, "", hostURLConfig)
	client, err := sdk.NewClient(&services.Config{
		TokenRetriever: tr_refresh,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}
	info, err := client.IdentityService.ValidateToken(&input)
	assert.Emptyf(t, err, "Error validating using access token generated from refresh token: %s", err)
	assert.NotNil(t, info)
}

// TestIntegrationRefreshTokenRetryWorkflow tests ingesting event with invalid access token then retrying after obtaining new access token with refresh token
func TestIntegrationRefreshTokenRetryWorkflow(t *testing.T) {
	// get a new refresh token
	hostURLConfig := idp.HostURLConfig{TenantScoped: false, Tenant: testutils.TestTenant, Region: testutils.TestRegion}
	tr := idp.NewPKCERetriever(NativeClientID, NativeAppRedirectURI, idp.DefaultRefreshScope, TestUsername, TestPassword, IdpHostTenantScoped, "", hostURLConfig)
	ctx, err := tr.GetTokenContext()
	require.Emptyf(t, err, "Error validating using access token generated from PKCE flow: %s", err)
	require.NotNil(t, ctx)

	tr_refresh := &retryTokenRetriever{TR: idp.NewRefreshTokenRetriever(NativeClientID, idp.DefaultRefreshScope, ctx.RefreshToken, IdpHostTenantScoped, "", hostURLConfig)}
	client, err := sdk.NewClient(&services.Config{
		TokenRetriever: tr_refresh,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	sourcetype := "sourcetype:refreshtokentest"
	source := "manual-events"
	host := client.GetURL("").RequestURI()
	body := make(map[string]interface{})
	body["event"] = "refreshtokentest"
	timeValue := int64(1529945001)
	attributes := make(map[string]interface{})
	attributes1 := make(map[string]interface{})
	attributes1["testKey"] = "testValue"
	attributes["testkey2"] = attributes1

	testIngestEvent := ingest.Event{
		Host:       &host,
		Body:       body,
		Sourcetype: &sourcetype,
		Source:     &source,
		Timestamp:  &timeValue,
		Attributes: attributes}

	_, err = client.IngestService.PostEvents([]ingest.Event{testIngestEvent})
	assert.Emptyf(t, err, "Error ingesting test event using refresh token: %s", err)
}

// TestIntegrationClientCredentialsInitWorkflow tests initializing the client with a TokenRetriever impleme
func TestIntegrationClientCredentialsInitWorkflow(t *testing.T) {
	hostURLConfig := idp.HostURLConfig{TenantScoped: false, Tenant: testutils.TestTenant, Region: testutils.TestRegion}
	tr := idp.NewClientCredentialsRetriever(BackendClientIDTenantScoped, BackendClientSecretTenantScoped, BackendServiceScope, IdpHostTenantScoped, "", hostURLConfig)
	client, err := sdk.NewClient(&services.Config{
		TokenRetriever: tr,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}
	_, err = client.IdentityService.ValidateToken(&input)
	assert.Emptyf(t, err, "Error validating using access token generated from client credentials: %s", err)
}

// TestIntegrationClientCredentialsInitWorkflow tests with tenantScoped true
func TestIntegrationClientCredentialsInitWorkflowTenantScoped(t *testing.T) {
	hostURLConfig := idp.HostURLConfig{TenantScoped: true, Tenant: testutils.TestTenant, Region: testutils.TestRegion}
	tr := idp.NewClientCredentialsRetriever(BackendClientIDTenantScoped, BackendClientSecretTenantScoped, BackendServiceScope, IdpHostTenantScoped, "", hostURLConfig)
	client, err := sdk.NewClient(&services.Config{
		TokenRetriever: tr,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
		TenantScoped:   false,
		Region:         testutils.TestRegion,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)
	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}
	_, err = client.IdentityService.ValidateToken(&input)
	assert.Emptyf(t, err, "Error validating using access token generated from client credentials: %s", err)
}

// TestIntegrationClientCredentialsInitWorkflow tests with Empty HostURL config
func TestIntegrationClientCredentialsInitWorkflowEmptyHostUrlConfig(t *testing.T) {
	tr := idp.NewClientCredentialsRetriever(BackendClientIDTenantScoped, BackendClientSecretTenantScoped, BackendServiceScope, IdpHostTenantScoped, "", idp.HostURLConfig{})
	client, err := sdk.NewClient(&services.Config{
		TokenRetriever: tr,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)
	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}
	_, err = client.IdentityService.ValidateToken(&input)
	assert.Emptyf(t, err, "Error validating using access token generated from client credentials: %s", err)
}

// TestIntegrationClientCredentialsInitWorkflow tests with OverrideAuthURL
func TestIntegrationClientCredentialsInitWorkflowOverrideAuthURL(t *testing.T) {
	overrideAuthURL := IdpHostTenantScoped
	hostURLConfig := idp.HostURLConfig{TenantScoped: true, Tenant: testutils.TestTenant, Region: testutils.TestRegion}
	tr := idp.NewClientCredentialsRetriever(BackendClientIDTenantScoped, BackendClientSecretTenantScoped, BackendServiceScope, IdpHostTenantScoped, overrideAuthURL, hostURLConfig)
	client, err := sdk.NewClient(&services.Config{
		TokenRetriever: tr,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
		TenantScoped:   false,
		Region:         "region",
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)
	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}
	_, err = client.IdentityService.ValidateToken(&input)
	assert.Emptyf(t, err, "Error validating using access token generated from client credentials: %s", err)
}

// TestIntegrationClientCredentialsRetryWorkflow tests ingesting event with invalid access token then retrying after obtaining new access token with client credentials flow
func TestIntegrationClientCredentialsRetryWorkflow(t *testing.T) {
	hostURLConfig := idp.HostURLConfig{TenantScoped: true, Tenant: testutils.TestTenant, Region: testutils.TestRegion}
	tr := &retryTokenRetriever{TR: idp.NewClientCredentialsRetriever(BackendClientIDTenantScoped, BackendClientSecretTenantScoped, BackendServiceScope, IdpHostTenantScoped, "", hostURLConfig)}
	client, err := sdk.NewClient(&services.Config{
		TokenRetriever: tr,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	// Make sure the backend client id has been added to the tenant as an admin (authz is needed for ingest), errs are ignored - if either fails (e.g. for 405 duplicate) we are probably still OK
	_, _ = getClient(t).IdentityService.AddMember(identity.AddMemberBody{Name: BackendClientIDTenantScoped})
	_, _ = getClient(t).IdentityService.AddGroupMember("tenant.admins", identity.AddGroupMemberBody{Name: BackendClientIDTenantScoped})

	sourcetype := "sourcetype:clientcredentialstest"
	source := "manual-events"
	host := client.GetURL("").RequestURI()
	body := make(map[string]interface{})
	body["event"] = "clientcredentialstest"
	timeValue := int64(1529945001)
	attributes := make(map[string]interface{})
	attributes1 := make(map[string]interface{})
	attributes1["testKey"] = "testValue"
	attributes["testkey2"] = attributes1

	testIngestEvent := ingest.Event{
		Host:       &host,
		Body:       body,
		Sourcetype: &sourcetype,
		Source:     &source,
		Timestamp:  &timeValue,
		Attributes: attributes}

	_, err = client.IngestService.PostEvents([]ingest.Event{testIngestEvent})
	assert.Emptyf(t, err, "Error ingesting test event using client credentials flow error: %s", err)
}

// TestIntegrationPKCEInitWorkflow tests initializing the client with a TokenRetriever which obtains a new access token with PKCE flow
func TestIntegrationPKCEInitWorkflow(t *testing.T) {
	hostURLConfig := idp.HostURLConfig{TenantScoped: true, Tenant: testutils.TestTenant, Region: testutils.TestRegion}
	tr := idp.NewPKCERetriever(NativeClientID, NativeAppRedirectURI, idp.DefaultOIDCScopes, TestUsername, TestPassword, IdpHostTenantScoped, "", hostURLConfig)
	client, err := sdk.NewClient(&services.Config{
		TokenRetriever: tr,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}
	_, err = client.IdentityService.ValidateToken(&input)
	assert.Emptyf(t, err, "Error validating using access token generated from PKCE flow: %s", err)
}

// TestIntegrationPKCERetryWorkflow tests ingesting event with invalid access token then retrying after obtaining new access token with PKCE flow
func TestIntegrationPKCERetryWorkflow(t *testing.T) {
	hostURLConfig := idp.HostURLConfig{TenantScoped: true, Tenant: testutils.TestTenant, Region: testutils.TestRegion}
	tr := &retryTokenRetriever{TR: idp.NewPKCERetriever(NativeClientID, NativeAppRedirectURI, idp.DefaultOIDCScopes, TestUsername, TestPassword, IdpHostTenantScoped, "", hostURLConfig)}

	client, err := sdk.NewClient(&services.Config{
		TokenRetriever: tr,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	sourcetype := "sourcetype:clientcredentialstest"
	source := "manual-events"
	host := client.GetURL("").RequestURI()
	body := make(map[string]interface{})
	body["event"] = "clientcredentialstest"
	timeValue := int64(1529945001)
	attributes := make(map[string]interface{})
	attributes1 := make(map[string]interface{})
	attributes1["testKey"] = "testValue"
	attributes["testkey2"] = attributes1

	testIngestEvent := ingest.Event{
		Host:       &host,
		Body:       body,
		Sourcetype: &sourcetype,
		Source:     &source,
		Timestamp:  &timeValue,
		Attributes: attributes}

	_, err = client.IngestService.PostEvents([]ingest.Event{testIngestEvent})
	assert.Emptyf(t, err, "Error ingesting test event using PKCE flow error: %s", err)
}

// TestBadTokenRetryWorkflow tests to make sure that a 401 is returned to the end user when a bad token is retrieved and requests are re-tried exactly once
func TestBadTokenRetryWorkflow(t *testing.T) {
	tr := &badTokenRetriever{}

	client, err := sdk.NewClient(&services.Config{
		TokenRetriever: tr,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	sourcetype := "sourcetype:badtokentest"
	source := "manual-events"
	host := client.GetURL("").RequestURI()
	body := make(map[string]interface{})
	body["event"] = "badtokentest"
	timeValue := int64(1529945001)
	attributes := make(map[string]interface{})
	attributes1 := make(map[string]interface{})
	attributes1["testKey"] = "testValue"
	attributes["testkey2"] = attributes1

	testIngestEvent := ingest.Event{
		Host:       &host,
		Body:       body,
		Sourcetype: &sourcetype,
		Source:     &source,
		Timestamp:  &timeValue,
		Attributes: attributes}

	_, err = client.IngestService.PostEvents([]ingest.Event{testIngestEvent})
	assert.Equal(t, tr.N, 2, "Expected exactly two calls to TokenRetriever.GetTokenContext(): 1) at client initialization and 2) after 401 is encountered when client.IngestService.CreateEvent is called")
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok, "Expected err to be util.HTTPError")
	assert.True(t, httpErr.HTTPStatusCode == 401, "Expected error code 401 for multiple attempts with expired access tokens")
}

// TestIntegrationDeviceWorkflow tests getting an access token with device flow and failing because of invalid device code
// TODO more test cases in SCP-33667
func TestIntegrationDeviceWorkflowInvalidCode(t *testing.T) {
	t.Skip("Skip pending TODO mentioned above")
	hostURLConfig := idp.HostURLConfig{TenantScoped: true, Tenant: testutils.TestTenant, Region: testutils.TestRegion}
	tr := idp.NewDeviceFlowRetriever(NativeClientID, IdpHostTenantScoped, "", hostURLConfig)
	result, err := tr.Client.GetDeviceCodes(NativeClientID, "offline_access profile email")
	require.Nil(t, err)
	tr.DeviceCode = result.DeviceCode + "invalid"
	tr.ExpiresIn = result.ExpiresIn
	tr.Interval = result.Interval
	ctx, err := tr.GetTokenContext()
	assert.Nil(t, ctx)
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to get token in Device flow: code expired: 400")
}

// TestIntegrationDeviceWorkflow tests getting an access token with device flow and failing because of timeout
// TODO more test cases in SCP-33667
func TestIntegrationDeviceWorkflowPolling(t *testing.T) {
	t.Skip("Skip pending TODO mentioned above")
	hostURLConfig := idp.HostURLConfig{TenantScoped: true, Tenant: testutils.TestTenant, Region: testutils.TestRegion}
	tr := idp.NewDeviceFlowRetriever(NativeClientID, IdpHostTenantScoped, "", hostURLConfig)
	result, err := tr.Client.GetDeviceCodes(NativeClientID, "offline_access profile email")
	require.Nil(t, err)
	tr.DeviceCode = result.DeviceCode
	ctx, err := tr.GetTokenContext()
	require.Nil(t, ctx)
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to get token in Device flow: code expired: 400")
}
