// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/testutils"
)

// CRUD tenant and add/delete user to the tenant
func TestIntegrationCRUDTenant(t *testing.T) {
	client := getClient(t)

	//get user profile
	userSystem, err := client.IdentityService.GetUserProfile("")
	assert.Nil(t, err)
	assert.Equal(t, "test1@splunk.com", userSystem.ID)
	assert.Equal(t, "test1@splunk.com", userSystem.Email)
	assert.Equal(t, "test1@splunk.com", userSystem.FirstName)
	assert.Equal(t, "test1@splunk.com", userSystem.LastName)
	assert.Equal(t, "test1@splunk.com", userSystem.Name)
	assert.Equal(t, "US", userSystem.Locale)

	userTenant, errTenant := client.IdentityService.GetUserProfile(testutils.TestTenantID)
	assert.Nil(t, errTenant)
	assert.Equal(t, "test1@splunk.com", userTenant.ID)
	assert.Equal(t, "test1@splunk.com", userTenant.Email)
	assert.Equal(t, "test1@splunk.com", userTenant.FirstName)
	assert.Equal(t, "test1@splunk.com", userTenant.LastName)
	assert.Equal(t, "test1@splunk.com", userTenant.Name)
	assert.Equal(t, "US", userTenant.Locale)

	if testutils.TenantCreationOn {
		//prepare a temp tenant that will be deleted
		testTenantID := fmt.Sprintf("%d-sdk-integration", time.Now().Unix())

		defer client.IdentityService.DeleteTenant(testTenantID)

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
	} else {
		t.Skip("Tenant creation tests were skipped, to turn them on set the TENANT_CREATION env to 1")
	}
}

// test Errors with auth endpoints
func TestIntegrationTenantErrors(t *testing.T) {
	client := getClient(t)

	//integration test tenant should already exist
	testTenantID := testutils.TestTenantID

	//create duplicate tenant should return 409
	if testutils.TenantCreationOn {
		err := client.IdentityService.CreateTenant(model.Tenant{TenantID: testTenantID})
		assert.True(t, strings.Contains(err.Error(), "409 Conflict"))
	} else {
		t.Skip("Tenant creation tests were skipped, to turn them on set the TENANT_CREATION env to 1")
	}
}
