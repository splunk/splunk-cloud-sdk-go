// +build !integration

package playgroundintegration

import (
	"os"
	"testing"

	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
)

//Expired token
var TestAuthenticationToken = os.Getenv("EXPIRED_BEARER_TOKEN")

// CRUD tenant and add/delete user to the tenant
func TestIntegrationRefreshTokenWorkflow(t *testing.T) {
	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost
	client, _ := service.NewClient(testutils.TestTenantID, TestAuthenticationToken, url, testutils.TestTimeOut)
	testTenantID := testutils.TestTenantID
	//get user profile
	user, err := client.IdentityService.GetUserProfile(testTenantID)
	if err != nil {
		t.FailNow()
	}
	assert.Nil(t, err)
	assert.Equal(t, "test1@splunk.com", user.ID)
	assert.Equal(t, "test1@splunk.com", user.Email)
	assert.Equal(t, "Test1", user.FirstName)
	assert.Equal(t, "Splunk", user.LastName)
	assert.Equal(t, "Test1 Splunk", user.Name)
	assert.Equal(t, "en-US", user.Locale)
}
