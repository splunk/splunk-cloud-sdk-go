// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

// +build !integration

package playgroundintegration

import (
	"os"
	"testing"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/service/handler"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//Expired token
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

//Test ingesting event with invalid access token then retrying after obtaining new access token with refresh token
func TestIntegrationRefreshTokenWorkflow(t *testing.T) {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	client, err := service.NewClient(&service.Config{
		Token:    ExpiredAuthenticationToken,
		URL:      url,
		TenantID: testutils.TestTenantID,
		Timeout:  testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)
	rh := handler.NewRefreshTokenAuthnResponseHandler(IDPHost, NativeClientID, handler.DefaultOIDCScopes, RefreshToken)
	client.SetResponseHandler(rh)

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

//Test ingesting event with invalid access token then retrying after obtaining new access token with client credentials flow
func TestIntegrationClientCredentialsWorkflow(t *testing.T) {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	client, err := service.NewClient(&service.Config{
		Token:    ExpiredAuthenticationToken,
		URL:      url,
		TenantID: testutils.TestTenantID,
		Timeout:  testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)
	rh := handler.NewClientCredentialsAuthnResponseHandler(IDPHost, BackendClientID, BackendClientSecret, BackendServiceScope)
	client.SetResponseHandler(rh)

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

//Test ingesting event with invalid access token then retrying after obtaining new access token with PKCE flow
func TestIntegrationPKCEWorkflow(t *testing.T) {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	client, err := service.NewClient(&service.Config{
		Token:    ExpiredAuthenticationToken,
		URL:      url,
		TenantID: testutils.TestTenantID,
		Timeout:  testutils.TestTimeOut,
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)
	rh := handler.NewPKCEAuthnResponseHandler(IDPHost, NativeClientID, NativeAppRedirectURI, handler.DefaultOIDCScopes, TestUsername, TestPassword)
	client.SetResponseHandler(rh)

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
