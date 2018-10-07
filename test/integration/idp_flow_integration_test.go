// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"os"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/service"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RefreshToken - RefreshToken to refresh the bearer token if expired
var RefreshToken = os.Getenv("REFRESH_TOKEN")

// IdpURL - host to retrieve access token from
var IdpURL = os.Getenv("IDP_URL")

// IDPHost - host to retrieve access token from
var IdpAuthz = os.Getenv("IDP_AUTHZ_SERVER")

// NativeClientID - Okta app Client Id for SDK Native App
var NativeClientID = os.Getenv("REFRESH_TOKEN_CLIENT_ID")

// NativeAppRedirectURI is one of the redirect uris configured for the app
const NativeAppRedirectURI = "http://localhost:9090"

// BackendClientID - Okta app Client id for client credentials flow
var BackendClientID = os.Getenv("BACKEND_CLIENT_ID")

// BackendClientSecret - Okta app Client secret for client credentials flow
var BackendClientSecret = os.Getenv("BACKEND_CLIENT_SECRET")

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
// func TestIntegrationRefreshTokenInitWorkflow(t *testing.T) {
// 	tr := idp.NewRefreshTokenRetriever(NativeClientID, idp.DefaultOIDCScopes, RefreshToken, IdpURL, IdpAuthz)
// 	client, err := service.NewClient(&service.Config{
// 		TokenRetriever: tr,
// 		Scheme:         testutils.TestURLProtocol,
// 		Host:           testutils.TestSplunkCloudHost,
// 		Tenant:         "system",
// 		Timeout:        testutils.TestTimeOut,
// 	})
// 	require.Emptyf(t, err, "Error initializing client: %s", err)

// 	_, err = client.IdentityService.Validate()
// 	assert.Emptyf(t, err, "Error validating using access token generated from refresh token: %s", err)
// }

// TestIntegrationRefreshTokenRetryWorkflow tests ingesting event with invalid access token then retrying after obtaining new access token with refresh token
func TestIntegrationRefreshTokenRetryWorkflow(t *testing.T) {
	tr := &retryTokenRetriever{TR: idp.NewRefreshTokenRetriever(NativeClientID, idp.DefaultOIDCScopes, RefreshToken, IdpURL, IdpAuthz)}
	client, err := service.NewClient(&service.Config{
		TokenRetriever: tr,
		Scheme:         testutils.TestURLProtocol,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	// Make sure the backend client id has been added to the tenant, err is ignored - if this fails (e.g. for 405 duplicate) we are probably still OK
	// _, _ = getClient(t).IdentityService.AddMember(BackendClientID)

	timeValue := int64(1529945001)
	testIngestEvent := model.Event{
		Host:       client.GetURL().RequestURI(),
		Body:       "refreshtokentest",
		Sourcetype: "sourcetype:refreshtokentest",
		Source:     "manual-events",
		Timestamp:  timeValue,
		Attributes: map[string]interface{}{"testKey": "testValue"}}

	err = client.IngestService.PostEvents([]model.Event{testIngestEvent})
	assert.Emptyf(t, err, "Error ingesting test event using refresh token: %s", err)
}

// TestIntegrationClientCredentialsInitWorkflow tests initializing the client with a TokenRetriever impleme
// func TestIntegrationClientCredentialsInitWorkflow(t *testing.T) {
// 	tr := idp.NewClientCredentialsRetriever(BackendClientID, BackendClientSecret, BackendServiceScope, IdpURL, IdpAuthz)
// 	client, err := service.NewClient(&service.Config{
// 		TokenRetriever: tr,
// 		Scheme:         testutils.TestURLProtocol,
// 		Host:           testutils.TestSplunkCloudHost,
// 		Tenant:         "system",
// 		Timeout:        testutils.TestTimeOut,
// 	})
// 	require.Emptyf(t, err, "Error initializing client: %s", err)

// 	_, err = client.IdentityService.Validate()
// 	assert.Emptyf(t, err, "Error validating using access token generated from client credentials: %s", err)
// }

// TestIntegrationClientCredentialsRetryWorkflow tests ingesting event with invalid access token then retrying after obtaining new access token with client credentials flow
func TestIntegrationClientCredentialsRetryWorkflow(t *testing.T) {
	tr := &retryTokenRetriever{TR: idp.NewClientCredentialsRetriever(BackendClientID, BackendClientSecret, BackendServiceScope, IdpURL, IdpAuthz)}
	client, err := service.NewClient(&service.Config{
		TokenRetriever: tr,
		Scheme:         testutils.TestURLProtocol,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	// Make sure the backend client id has been added to the tenant, err is ignored - if this fails (e.g. for 405 duplicate) we are probably still OK
	// _, _ = client.IdentityService.AddMember(BackendClientID)

	timeValue := int64(1529945002)
	testIngestEvent := model.Event{
		Host:       client.GetURL().RequestURI(),
		Body:       "clientcredentialstest",
		Sourcetype: "sourcetype:clientcredentialstest",
		Source:     "manual-events",
		Timestamp:  timeValue,
		Attributes: map[string]interface{}{"testKey": "testValue"}}

	err = client.IngestService.PostEvents([]model.Event{testIngestEvent})
	assert.Emptyf(t, err, "Error ingesting test event using client credentials flow error: %s", err)
}

// TestIntegrationPKCEInitWorkflow tests initializing the client with a TokenRetriever which obtains a new access token with PKCE flow
// func TestIntegrationPKCEInitWorkflow(t *testing.T) {
// 	tr := idp.NewPKCERetriever(NativeClientID, NativeAppRedirectURI, idp.DefaultOIDCScopes, TestUsername, TestPassword, IdpURL, IdpAuthz)
// 	client, err := service.NewClient(&service.Config{
// 		TokenRetriever: tr,
// 		Scheme:         testutils.TestURLProtocol,
// 		Host:           testutils.TestSplunkCloudHost,
// 		Tenant:         "system",
// 		Timeout:        testutils.TestTimeOut,
// 	})
// 	require.Emptyf(t, err, "Error initializing client: %s", err)

// 	_, err = client.IdentityService.Validate()
// 	assert.Emptyf(t, err, "Error validating using access token generated from PKCE flow: %s", err)
// }

// TestIntegrationPKCERetryWorkflow tests ingesting event with invalid access token then retrying after obtaining new access token with PKCE flow
func TestIntegrationPKCERetryWorkflow(t *testing.T) {
	tr := &retryTokenRetriever{TR: idp.NewPKCERetriever(NativeClientID, NativeAppRedirectURI, idp.DefaultOIDCScopes, TestUsername, TestPassword, IdpURL, IdpAuthz)}

	client, err := service.NewClient(&service.Config{
		TokenRetriever: tr,
		Scheme:         testutils.TestURLProtocol,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	timeValue := int64(1529945003)
	testIngestEvent := model.Event{
		Host:       client.GetURL().RequestURI(),
		Body:       "pkcetest",
		Sourcetype: "sourcetype:pkcetest",
		Source:     "manual-events",
		Timestamp:  timeValue,
		Attributes: map[string]interface{}{"testKey": "testValue"}}

	err = client.IngestService.PostEvents([]model.Event{testIngestEvent})
	assert.Emptyf(t, err, "Error ingesting test event using PKCE flow error: %s", err)
}

// TestBadTokenRetryWorkflow tests to make sure that a 401 is returned to the end user when a bad token is retrieved and requests are re-tried exactly once
func TestBadTokenRetryWorkflow(t *testing.T) {
	tr := &badTokenRetriever{}

	client, err := service.NewClient(&service.Config{
		TokenRetriever: tr,
		Scheme:         testutils.TestURLProtocol,
		Host:           testutils.TestSplunkCloudHost,
		Tenant:         testutils.TestTenant,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	timeValue := int64(1529945004)
	testIngestEvent := model.Event{
		Host:       client.GetURL().RequestURI(),
		Body:       "badtokentest",
		Sourcetype: "sourcetype:badtokentest",
		Source:     "manual-events",
		Timestamp:  timeValue,
		Attributes: map[string]interface{}{"testKey": "testValue"}}

	err = client.IngestService.PostEvents([]model.Event{testIngestEvent})
	assert.Equal(t, tr.N, 2, "Expected exactly two calls to TokenRetriever.GetTokenContext(): 1) at client initialization and 2) after 401 is encountered when client.IngestService.CreateEvent is called")
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok, "Expected err to be util.HTTPError")
	assert.True(t, httpErr.HTTPStatusCode == 401, "Expected error code 401 for multiple attempts with expired access tokens")
}
