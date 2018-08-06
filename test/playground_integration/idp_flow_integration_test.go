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
var TestAuthenticationToken = os.Getenv("EXPIRED_BEARER_TOKEN")

// RefreshToken - RefreshToken to refresh the bearer token if expired
var RefreshToken = os.Getenv("REFRESH_TOKEN")

// IdPHost - host to retrieve access token from
var IdPHost = os.Getenv("IDP_HOST")

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

//Test ingesting event with invalid access token then retrying after obtaining new access token with refresh token
func TestIntegrationRefreshTokenWorkflow(t *testing.T) {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	client, err := service.NewClient(&service.Config{
		Token:           TestAuthenticationToken,
		URL:             url,
		TenantID:        testutils.TestTenantID,
		Timeout:         testutils.TestTimeOut,
		ResponseHandler: handler.NewRefreshTokenAuthnResponseHandler(IdPHost, RefreshClientID, RefreshTokenScope, RefreshToken),
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	clientURL, err := client.GetURL()
	require.Emptyf(t, err, "Error retrieving client URL: %s", err)

	timeValue := float64(1529945178)
	testIngestEvent := model.Event{
		Host:       clientURL.RequestURI(),
		Index:      "main",
		Event:      "refreshtokentest",
		Sourcetype: "sourcetype:refreshtokentest",
		Source:     "manual-events",
		Time:       &timeValue,
		Fields:     map[string]string{"testKey": "testValue"}}

	err = client.IngestService.CreateEvent(testIngestEvent)
	assert.Emptyf(t, err, "Error ingesting test event using refresh logic: %s", err)
}

//Test ingesting event with invalid access token then retrying after obtaining new access token with client credentials flow
func TestIntegrationClientCredentialsWorkflow(t *testing.T) {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	client, err := service.NewClient(&service.Config{
		Token:           TestAuthenticationToken,
		URL:             url,
		TenantID:        testutils.TestTenantID,
		Timeout:         testutils.TestTimeOut,
		ResponseHandler: handler.NewClientCredentialsAuthnResponseHandler(IdPHost, BackendClientID, BackendClientSecret, BackendServiceScope),
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	clientURL, err := client.GetURL()
	require.Emptyf(t, err, "Error retrieving client URL: %s", err)

	timeValue := float64(1529945178)
	testIngestEvent := model.Event{
		Host:       clientURL.RequestURI(),
		Index:      "main",
		Event:      "clientcredentialstest",
		Sourcetype: "sourcetype:clientcredentialstest",
		Source:     "manual-events",
		Time:       &timeValue,
		Fields:     map[string]string{"testKey": "testValue"}}

	err = client.IngestService.CreateEvent(testIngestEvent)
	assert.Emptyf(t, err, "Error ingesting test event using client credentials logic. NOTE: If 401 encountered this may be becaue the backend service app=%s needs to be added to SDK testing tenant=%s via PATCH to identity/{}/users error: %s", BackendClientID, testutils.TestTenantID, err)
}
