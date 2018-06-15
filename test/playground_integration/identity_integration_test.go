package playgroundintegration

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
)

// CRUD tenant and add/delete user to the tenant
func TestIntegrationCRUDTenant(t *testing.T) {
	testTenantID := "sscsdk-06152018"
	client := getClient(t)

	defer client.IdentityService.DeleteTenant(testTenantID)

	//get user profile
	user, err := client.IdentityService.GetUserProfile()
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

	//get tennant users
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

// test Erros with auth endpoints
func TestIntegrationTenantErrors(t *testing.T) {
	testTenantID := "sscsdk-06152018"
	client := getClient(t)

	defer client.IdentityService.DeleteTenant(testTenantID)

	//create duplicate tenant should return 409
	err := client.IdentityService.CreateTenant(model.Tenant{TenantID: testTenantID})
	assert.Nil(t, err)

	err = client.IdentityService.CreateTenant(model.Tenant{TenantID: testTenantID})
	assert.True(t, strings.Contains(err.Error(), "409 Conflict"))

	//add duplicate tenant user
	addedUserName := "newUser@splunk.com"
	err = client.IdentityService.AddTenantUsers(testTenantID, []model.User{{ID: addedUserName}})
	assert.Nil(t, err)
	err = client.IdentityService.AddTenantUsers(testTenantID, []model.User{{ID: addedUserName}})
	assert.True(t, strings.Contains(err.Error(), "405 Method Not Allowed"))

	//delete tenant
	err = client.IdentityService.DeleteTenant(testTenantID)
	assert.Nil(t, err)
}
