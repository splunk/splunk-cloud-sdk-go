// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/testutils"
)

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
