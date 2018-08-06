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
var IdPHost = "https://splunk-ciam.okta.com/" //os.Getenv("IDP_HOST")

// ClientID - Okta app Client Id for SDK
var ClientID = os.Getenv("CLIENT_ID")

//Test ingesting event with invalid access token then retrying after refreshing token
func TestIntegrationRefreshTokenWorkflow(t *testing.T) {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost

	client, err := service.NewClient(&service.Config{
		Token:           TestAuthenticationToken,
		URL:             url,
		TenantID:        testutils.TestTenantID,
		Timeout:         testutils.TestTimeOut,
		ResponseHandler: handler.NewRefreshTokenAuthnResponseHandler(IdPHost, ClientID, "openid email profile", RefreshToken),
	})
	require.Emptyf(t, err, "Error initializing client: %s", err)

	clientURL, err := client.GetURL()
	assert.Emptyf(t, err, "Error retrieving client URL: %s", err)

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
