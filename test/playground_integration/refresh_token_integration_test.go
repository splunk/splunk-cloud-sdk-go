// +build !integration

package playgroundintegration

import (
	"testing"
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"os"
)

//Expired token
var TestAuthenticationToken = os.Getenv("EXPIRED_BEARER_TOKEN")

// CRUD tenant and add/delete user to the tenant
func TestIntegrationRefreshTokenWorkflow(t *testing.T) {
	testTenantID := "6fecc441-fb60-4992-a822-069feb622fab"

	var url = testutils.TestURLProtocol + "://" + testutils.TestSSCHost
	client,_ := service.NewClient(testutils.TestTenantID, TestAuthenticationToken, url, testutils.TestTimeOut)

	defer client.IdentityService.DeleteTenant(testTenantID)

	//get user profile
	user, err := client.IdentityService.GetUserProfile()
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

	//create tenant
	err = client.IdentityService.CreateTenant(model.Tenant{TenantID: testTenantID})
	assert.Nil(t, err)

	//add tenant user
	addedUserName := "newUser@splunk.com"
	err = client.IdentityService.AddTenantUsers(testTenantID, []model.User{{ID: addedUserName}})
	assert.Nil(t, err)

	//get tenant users
	users, err := client.IdentityService.GetTenantUsers(testTenantID)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))

	found := false
	for _, v := range users {
		if v.ID == addedUserName {
			found = true
			break
		}
	}
	assert.True(t, found)

	//delete tenant user
	err = client.IdentityService.DeleteTenantUsers(testTenantID, []model.User{{ID: addedUserName}})
	assert.Nil(t, err)

	users, err = client.IdentityService.GetTenantUsers(testTenantID)
	found = false
	for _, v := range users {
		if v.ID == addedUserName {
			fmt.Println(v.ID)
			found = true
			break
		}
	}
	assert.False(t, found)

	//replace tenant users
	err = client.IdentityService.ReplaceTenantUsers(testTenantID, []model.User{
		{ID: "devtest2@splunk.com"},
		{ID: "devtest3@splunk.com"}})

	users, err = client.IdentityService.GetTenantUsers(testTenantID)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))

	//delete tenant
	err = client.IdentityService.DeleteTenant(testTenantID)
	assert.Nil(t, err)
}

