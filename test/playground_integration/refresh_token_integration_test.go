// +build !integration

package playgroundintegration

import (
	"os"
	"testing"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//Expired token
var TestAuthenticationToken = os.Getenv("EXPIRED_BEARER_TOKEN")

//Test ingesting event with invalid access token then retrying after refreshing token
func TestIntegrationRefreshTokenWorkflow(t *testing.T) {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost
	client, err := service.NewClient(testutils.TestTenantID, TestAuthenticationToken, url, testutils.TestTimeOut)

	require.Emptyf(t, err, "Error initializing client: %s", err)

	timeValue := float64(1529945178)
	testHecEvent := model.HecEvent{
		Host:       client.URL.RequestURI(),
		Index:      "main",
		Event:      "refreshtokentest",
		Sourcetype: "sourcetype:refreshtokentest",
		Source:     "manual-events",
		Time:       &timeValue,
		Fields:     map[string]string{"testKey": "testValue"}}

	err = client.HecService.CreateEvent(testHecEvent)
	assert.Emptyf(t, err, "Error ingesting test event using refresh logic: %s", err)
}
