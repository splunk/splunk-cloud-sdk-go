// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

// +build !integration

package playgroundintegration

import (
	"os"
	"testing"

	"github.com/splunk/ssc-client-go/idp"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ExpiredAuthenticationToken - to test authentication retries
var ExpiredAuthenticationToken = os.Getenv("EXPIRED_BEARER_TOKEN")

// RefreshToken - RefreshToken to refresh the bearer token if expired
var RefreshToken = os.Getenv("REFRESH_TOKEN")

// IDPHost - host to retrieve access token from
var IDPHost = os.Getenv("IDP_HOST")

// NativeClientID - Okta app Client Id for SDK Native App
var NativeClientID = os.Getenv("REFRESH_TOKEN_CLIENT_ID")

// NativeAppRedirectURI is one of the redirect uris configured for the app
const NativeAppRedirectURI = "http://localhost:9090"

// BackendClientID - Okta app Client id for client credentials flow
var BackendClientID = os.Getenv("BACKEND_CLIENT_ID")

// BackendClientSecret - Okta app Client secret for client credentials flow
var BackendClientSecret = os.Getenv("BACKEND_CLIENT_SECRET")

// BackendServiceScope - scope for obtaining access token for client credentials flow
const BackendServiceScope = "backend_service"

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
		return &idp.Context{AccessToken: ExpiredAuthenticationToken}, nil
	}
	// For subsequent requests get the real token using the real TokenRetriever
	return r.TR.GetTokenContext()
}

// TestIntegrationRefreshTokenInitWorkflow tests initializing the client with a TokenRetriever impleme
func TestIntegrationRefreshTokenInitWorkflow(t *testing.T) {
	url := testutils.TestURLProtocol + "://" + testutils.TestSSCHost
	tr := idp.NewRefreshTokenRetriever(IDPHost, NativeClientID, idp.DefaultOIDCScopes, RefreshToken)
	client, err := service.NewClient(&service.Config{
		TokenRetriever: tr,
		URL:            url,
		TenantID:       testutils.TestTenantID,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	_, err = client.SearchService.GetJobs(nil)
	assert.Emptyf(t, err, "Error searching using access token generated from refresh token: %s", err)
}

// TestIntegrationRefreshTokenRetryWorkflow tests ingesting event with invalid access token then retrying after obtaining new access token with refresh token
func TestIntegrationRefreshTokenRetryWorkflow(t *testing.T) {
	url := testutils.TestURLProtocol + "://" + testutils.TestSSCHost
	tr := &retryTokenRetriever{TR: idp.NewRefreshTokenRetriever(IDPHost, NativeClientID, idp.DefaultOIDCScopes, RefreshToken)}
	client, err := service.NewClient(&service.Config{
		TokenRetriever: tr,
		URL:            url,
		TenantID:       testutils.TestTenantID,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	clientURL, err := client.GetURL()
	require.Emptyf(t, err, "Error retrieving client URL: %s", err)

	// Make sure the backend client id has been added to the tenant, err is ignored - if this fails (e.g. for 405 duplicate) we are probably still OK
	_ = getClient(t).IdentityService.AddTenantUsers(testutils.TestTenantID, []model.User{{ID: BackendClientID}})

	timeValue := float64(1529945001)
	testIngestEvent := model.Event{
		Host:       clientURL.RequestURI(),
		Index:      "main",
		Event:      "refreshtokentest",
		Sourcetype: "sourcetype:refreshtokentest",
		Source:     "manual-events",
		Time:       &timeValue,
		Fields:     map[string]string{"testKey": "testValue"}}

	err = client.IngestService.CreateEvent(testIngestEvent)
	assert.Emptyf(t, err, "Error ingesting test event using refresh token: %s", err)
}

// TestIntegrationClientCredentialsInitWorkflow tests initializing the client with a TokenRetriever impleme
func TestIntegrationClientCredentialsInitWorkflow(t *testing.T) {
	url := testutils.TestURLProtocol + "://" + testutils.TestSSCHost
	tr := idp.NewClientCredentialsRetriever(IDPHost, BackendClientID, BackendClientSecret, BackendServiceScope)
	client, err := service.NewClient(&service.Config{
		TokenRetriever: tr,
		URL:            url,
		TenantID:       testutils.TestTenantID,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	_, err = client.SearchService.GetJobs(nil)
	assert.Emptyf(t, err, "Error searching using access token generated from refresh token: %s", err)
}

// TestIntegrationClientCredentialsRetryWorkflow tests ingesting event with invalid access token then retrying after obtaining new access token with client credentials flow
func TestIntegrationClientCredentialsRetryWorkflow(t *testing.T) {
	url := testutils.TestURLProtocol + "://" + testutils.TestSSCHost
	tr := &retryTokenRetriever{TR: idp.NewClientCredentialsRetriever(IDPHost, BackendClientID, BackendClientSecret, BackendServiceScope)}
	client, err := service.NewClient(&service.Config{
		TokenRetriever: tr,
		URL:            url,
		TenantID:       testutils.TestTenantID,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	clientURL, err := client.GetURL()
	require.Emptyf(t, err, "Error retrieving client URL: %s", err)

	// Make sure the backend client id has been added to the tenant, err is ignored - if this fails (e.g. for 405 duplicate) we are probably still OK
	_ = client.IdentityService.AddTenantUsers(testutils.TestTenantID, []model.User{{ID: BackendClientID}})

	timeValue := float64(1529945002)
	testIngestEvent := model.Event{
		Host:       clientURL.RequestURI(),
		Index:      "main",
		Event:      "clientcredentialstest",
		Sourcetype: "sourcetype:clientcredentialstest",
		Source:     "manual-events",
		Time:       &timeValue,
		Fields:     map[string]string{"testKey": "testValue"}}

	err = client.IngestService.CreateEvent(testIngestEvent)
	assert.Emptyf(t, err, "Error ingesting test event using client credentials flow error: %s", err)
}

// TestIntegrationPKCEInitWorkflow tests initializing the client with a TokenRetriever which obtains a new access token with PKCE flow
func TestIntegrationPKCEInitWorkflow(t *testing.T) {
	url := testutils.TestURLProtocol + "://" + testutils.TestSSCHost
	tr := idp.NewPKCERetriever(IDPHost, NativeClientID, NativeAppRedirectURI, idp.DefaultOIDCScopes, TestUsername, TestPassword)
	client, err := service.NewClient(&service.Config{
		TokenRetriever: tr,
		URL:            url,
		TenantID:       testutils.TestTenantID,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	_, err = client.SearchService.GetJobs(nil)
	assert.Emptyf(t, err, "Error searching using access token generated from refresh token: %s", err)
}

// TestIntegrationPKCERetryWorkflow tests ingesting event with invalid access token then retrying after obtaining new access token with PKCE flow
func TestIntegrationPKCERetryWorkflow(t *testing.T) {
	url := testutils.TestURLProtocol + "://" + testutils.TestSSCHost
	tr := &retryTokenRetriever{TR: idp.NewPKCERetriever(IDPHost, NativeClientID, NativeAppRedirectURI, idp.DefaultOIDCScopes, TestUsername, TestPassword)}

	client, err := service.NewClient(&service.Config{
		TokenRetriever: tr,
		URL:            url,
		TenantID:       testutils.TestTenantID,
		Timeout:        testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	clientURL, err := client.GetURL()
	require.Emptyf(t, err, "Error retrieving client URL: %s", err)

	timeValue := float64(1529945003)
	testIngestEvent := model.Event{
		Host:       clientURL.RequestURI(),
		Index:      "main",
		Event:      "pkcetest",
		Sourcetype: "sourcetype:pkcetest",
		Source:     "manual-events",
		Time:       &timeValue,
		Fields:     map[string]string{"testKey": "testValue"}}

	err = client.IngestService.CreateEvent(testIngestEvent)
	assert.Emptyf(t, err, "Error ingesting test event using PKCE flow error: %s", err)
}
