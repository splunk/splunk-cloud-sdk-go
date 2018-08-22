// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

// +build !integration

package playgroundintegration

import (
	"os"
)

//Expired token
var ExpiredAuthenticationToken = os.Getenv("EXPIRED_BEARER_TOKEN")

// RefreshToken - RefreshToken to refresh the bearer token if expired
var RefreshToken = os.Getenv("REFRESH_TOKEN")

// IDPHost - host to retrieve access token from
var IDPHost = os.Getenv("IDP_HOST")

// RefreshClientID - Okta app Client Id for SDK Native App
var RefreshClientID = os.Getenv("REFRESH_TOKEN_CLIENT_ID")

// RefreshTokenScope - scope for obtaining access token using refresh
const RefreshTokenScope = "openid email profile"

// BackendClientID - Okta app Client id for client credentials flow
var BackendClientID = os.Getenv("BACKEND_CLIENT_ID")

// BackendClientSecret - Okta app Client secret for client credentials flow
var BackendClientSecret = os.Getenv("BACKEND_CLIENT_SECRET")

// BackendServiceScope - scope for obtaining access token for client credentials flow
const BackendServiceScope = "backend_service"

////Test ingesting event with invalid access token then retrying after obtaining new access token with refresh token
//func TestIntegrationRefreshTokenWorkflow(t *testing.T) {
//	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost
//
//	client, err := service.NewClient(&service.Config{
//		Token:    ExpiredAuthenticationToken,
//		URL:      url,
//		TenantID: testutils.TestTenantID,
//		Timeout:  testutils.TestTimeOut,
//	})
//	require.Emptyf(t, err, "Error initializing client: %s", err)
//	rh := handler.NewRefreshTokenAuthnResponseHandler(IDPHost, RefreshClientID, RefreshTokenScope, RefreshToken)
//	client.SetResponseHandler(rh)
//
//	clientURL, err := client.GetURL()
//	require.Emptyf(t, err, "Error retrieving client URL: %s", err)
//
//	// Make sure the backend client id has been added to the tenant, err is ignored - if this fails (e.g. for 405 duplicate) we are probably still OK
//	_ = getClient(t).IdentityService.AddTenantUsers(testutils.TestTenantID, []model.User{{ID: BackendClientID}})
//
//	timeValue := float64(1529945178)
//	testIngestEvent := model.Event{
//		Host:       clientURL.RequestURI(),
//		Index:      "main",
//		Event:      "refreshtokentest",
//		Sourcetype: "sourcetype:refreshtokentest",
//		Source:     "manual-events",
//		Time:       &timeValue,
//		Fields:     map[string]string{"testKey": "testValue"}}
//
//	err = client.IngestService.CreateEvent(testIngestEvent)
//	assert.Emptyf(t, err, "Error ingesting test event using refresh logic: %s", err)
//}
//
////Test ingesting event with invalid access token then retrying after obtaining new access token with client credentials flow
//func TestIntegrationClientCredentialsWorkflow(t *testing.T) {
//	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost
//
//	client, err := service.NewClient(&service.Config{
//		Token:    ExpiredAuthenticationToken,
//		URL:      url,
//		TenantID: testutils.TestTenantID,
//		Timeout:  testutils.TestTimeOut,
//	})
//	require.Emptyf(t, err, "Error initializing client: %s", err)
//	rh := handler.NewClientCredentialsAuthnResponseHandler(IDPHost, BackendClientID, BackendClientSecret, BackendServiceScope)
//	client.SetResponseHandler(rh)
//
//	clientURL, err := client.GetURL()
//	require.Emptyf(t, err, "Error retrieving client URL: %s", err)
//
//	// Make sure the backend client id has been added to the tenant, err is ignored - if this fails (e.g. for 405 duplicate) we are probably still OK
//	_ = client.IdentityService.AddTenantUsers(testutils.TestTenantID, []model.User{{ID: BackendClientID}})
//
//	timeValue := float64(1529945178)
//	testIngestEvent := model.Event{
//		Host:       clientURL.RequestURI(),
//		Index:      "main",
//		Event:      "clientcredentialstest",
//		Sourcetype: "sourcetype:clientcredentialstest",
//		Source:     "manual-events",
//		Time:       &timeValue,
//		Fields:     map[string]string{"testKey": "testValue"}}
//
//	err = client.IngestService.CreateEvent(testIngestEvent)
//	assert.Emptyf(t, err, "Error ingesting test event using client credentials logic error: %s", err)
//}
