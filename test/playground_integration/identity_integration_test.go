// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"strings"
	"testing"

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
	assert.Equal(t, "Test1", userSystem.FirstName)
	assert.Equal(t, "Splunk", userSystem.LastName)
	assert.Equal(t, "Test1 Splunk", userSystem.Name)
	assert.Equal(t, "en-US", userSystem.Locale)

	userTenant, errTenant := client.IdentityService.GetUserProfile(testutils.TestTenantID)
	assert.Nil(t, errTenant)
	assert.Equal(t, "test1@splunk.com", userTenant.ID)
	assert.Equal(t, "test1@splunk.com", userTenant.Email)
	assert.Equal(t, "Test1", userTenant.FirstName)
	assert.Equal(t, "Splunk", userTenant.LastName)
	assert.Equal(t, "Test1 Splunk", userTenant.Name)
	assert.Equal(t, "en-US", userTenant.Locale)

	// //prepare a temp tenant that will be deleted
	// testTenantID := fmt.Sprintf("%d-sdk-integration", time.Now().Unix())

	// defer client.IdentityService.DeleteTenant(testTenantID)

	// //create tenant
	// err = client.IdentityService.CreateTenant(model.Tenant{TenantID: testTenantID})
	// assert.Nil(t, err)

	// //add tenant user
	// addedUserName := "newUser@splunk.com"
	// err = client.IdentityService.AddTenantUsers(testTenantID, []model.User{{ID: addedUserName}})
	// assert.Nil(t, err)

	// //get tennant users
	// users, err := client.IdentityService.GetTenantUsers(testTenantID)
	// assert.Nil(t, err)
	// assert.Equal(t, 2, len(users))

	// found := false
	// for _, v := range users {
	// 	if v.ID == addedUserName {
	// 		found = true
	// 		break
	// 	}
	// }
	// assert.True(t, found)

	// //delete tenant user
	// err = client.IdentityService.DeleteTenantUsers(testTenantID, []model.User{{ID: addedUserName}})
	// assert.Nil(t, err)

	// users, err = client.IdentityService.GetTenantUsers(testTenantID)
	// found = false
	// for _, v := range users {
	// 	if v.ID == addedUserName {
	// 		found = true
	// 		break
	// 	}
	// }
	// assert.False(t, found)

	// //replace tenant users
	// err = client.IdentityService.ReplaceTenantUsers(testTenantID, []model.User{
	// 	{ID: "devtest2@splunk.com"},
	// 	{ID: "devtest3@splunk.com"}})

	// users, err = client.IdentityService.GetTenantUsers(testTenantID)
	// assert.Nil(t, err)
	// assert.Equal(t, 3, len(users))

	// //delete tenant
	// err = client.IdentityService.DeleteTenant(testTenantID)
	// assert.Nil(t, err)
}

// test Erros with auth endpoints
func TestIntegrationTenantErrors(t *testing.T) {
	client := getClient(t)

	//integration test tenant should already exist
	testTenantID := testutils.TestTenantID

	//create duplicate tenant should return 409
	// TODO: uncomment when tenant Maestro gets better at handling tenant creation/deletion
	// err = client.IdentityService.CreateTenant(model.Tenant{TenantID: testTenantID})
	// assert.True(t, strings.Contains(err.Error(), "409 Conflict"))

	//add duplicate tenant user
	addedUserName := "newUser@splunk.com"
	_ = client.IdentityService.AddTenantUsers(testTenantID, []model.User{{ID: addedUserName}})
	err := client.IdentityService.AddTenantUsers(testTenantID, []model.User{{ID: addedUserName}})
	assert.True(t, strings.Contains(err.Error(), "405 Method Not Allowed"))
}

func TestCRUDGroups(t *testing.T) {
	client := getClient(t)

	res,err:= client.IdentityService.GetGroups()
	assert.Nil(t,err)
	groupNum:=len(res)

	groupName := "testcrudgroup"

	// create/get/delete group and groups
	resultgroup, err := client.IdentityService.CreateGroup(groupName)
	defer client.IdentityService.DeleteGroup(groupName)
	assert.Nil(t, err)
	assert.Equal(t, groupName, resultgroup.Name)
	assert.Equal(t, "test1@splunk.com", resultgroup.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultgroup.Tenant)

	resultgroup1, err := client.IdentityService.GetGroup(groupName)
	assert.Nil(t, err)
	assert.Equal(t, groupName, resultgroup1.Name)
	assert.Equal(t, "test1@splunk.com", resultgroup1.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultgroup1.Tenant)

	resultgroup2, err := client.IdentityService.GetGroups()
	assert.Nil(t, err)
	assert.Equal(t, groupNum+1,len(resultgroup2))
	assert.Contains(t,resultgroup2,groupName )


	// group-roles
	roleName:="testcrudgrouprole"
	res2, err := client.IdentityService.GetGroupRoles(groupName)
	assert.Nil(t, err)
	roleNum:=len(res2)

	resultrole, err := client.IdentityService.CreateRole(roleName)
	defer client.IdentityService.DeleteRole(roleName)
	assert.Nil(t, err)
	assert.Equal(t, roleName, resultrole.Name)
	assert.Equal(t, "test1@splunk.com", resultrole.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultrole.Tenant)

	resultrole1, err := client.IdentityService.AddRoleToGroup(groupName,roleName)
	defer client.IdentityService.RemoveGroupRole(groupName,roleName)
	assert.Nil(t, err)
	assert.Equal(t, roleName, resultrole1.Role)
	assert.Equal(t, groupName, resultrole1.Group)
	assert.Equal(t, testutils.TestTenantID, resultrole1.Tenant)

	resultrole2, err := client.IdentityService.GetGroupRoles(groupName)
	assert.Nil(t, err)
	assert.Equal(t,roleNum+1,len(resultrole2))
	assert.Contains(t,resultrole2, roleName)

	// todo group-members

	//delete
	err = client.IdentityService.RemoveGroupRole(groupName,roleName)
	assert.Nil(t, err)

	err = client.IdentityService.DeleteRole(roleName)
	assert.Nil(t, err)

	err = client.IdentityService.DeleteGroup(groupName)
	assert.Nil(t, err)

}

func TestCRUDRoles(t *testing.T) {
	client := getClient(t)

	res,err:= client.IdentityService.GetRoles()
	assert.Nil(t,err)
	roleNum:=len(res)

	roleName := "testcrudgroup"

	// create/get/delete role and roles
	resultrole, err := client.IdentityService.CreateRole(roleName)
	defer client.IdentityService.DeleteRole(roleName)
	assert.Nil(t, err)
	assert.Equal(t, roleName, resultrole.Name)
	assert.Equal(t, "test1@splunk.com", resultrole.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultrole.Tenant)

	resultrole1, err := client.IdentityService.GetRole(roleName)
	assert.Nil(t, err)
	assert.Equal(t, roleName, resultrole1.Name)
	assert.Equal(t, "test1@splunk.com", resultrole1.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultrole1.Tenant)

	resultrole2, err := client.IdentityService.GetRoles()
	assert.Nil(t, err)
	assert.Equal(t, roleNum+1,len(resultrole2))
	assert.Contains(t,resultrole2,roleName )

	// role-permissions
	permissionName:="perm1"
	result1, err := client.IdentityService.GetRolePermissions(roleName)
	assert.Nil(t, err)
	permNum:=len(result1)

	resultroleperm, err := client.IdentityService.AddPermissionToRole(roleName,permissionName)
	defer client.IdentityService.RemoveRolePermission(roleName,permissionName)
	assert.Nil(t, err)
	assert.Equal(t, roleName, resultroleperm.Role)
	assert.Equal(t, permissionName, resultroleperm.Permission)
	assert.Equal(t, "test1@splunk.com", resultroleperm.AddedBy)
	assert.Equal(t, testutils.TestTenantID, resultroleperm.Tenant)

	resultroleperm1, err := client.IdentityService.GetRolePermission(roleName,permissionName)
	assert.Nil(t, err)
	assert.Equal(t, roleName, resultroleperm1.Role)
	assert.Equal(t, permissionName, resultroleperm1.Permission)
	assert.Equal(t, "test1@splunk.com", resultroleperm1.AddedBy)
	assert.Equal(t, testutils.TestTenantID, resultroleperm1.Tenant)

	resultroleperm2, err := client.IdentityService.GetRolePermissions(roleName)
	assert.Nil(t, err)
	assert.Equal(t, permNum+1, len(resultroleperm2))
	assert.Contains(t, resultroleperm2,permissionName)


	// delete
	err = client.IdentityService.RemoveRolePermission(roleName,permissionName)
	assert.Nil(t, err)
	err = client.IdentityService.DeleteRole(roleName)
	assert.Nil(t, err)
}
