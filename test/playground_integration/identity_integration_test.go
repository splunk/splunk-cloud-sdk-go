// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"fmt"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCRUDGroups(t *testing.T) {
	client := getClient(t)

	res, err := client.IdentityService.GetGroups()
	require.Nil(t, err)
	groupNum := len(res)

	groupName := fmt.Sprintf("grouptest%d", timeSec)

	// create/get/delete group and groups
	resultgroup, err := client.IdentityService.CreateGroup(groupName)
	defer client.IdentityService.DeleteGroup(groupName)
	require.Nil(t, err)
	assert.Equal(t, groupName, resultgroup.Name)
	assert.Equal(t, "test1@splunk.com", resultgroup.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultgroup.Tenant)

	resultgroup1, err := client.IdentityService.GetGroup(groupName)
	require.Nil(t, err)
	assert.Equal(t, groupName, resultgroup1.Name)
	assert.Equal(t, "test1@splunk.com", resultgroup1.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultgroup1.Tenant)

	resultgroup2, err := client.IdentityService.GetGroups()
	require.Nil(t, err)
	assert.Equal(t, groupNum+2, len(resultgroup2))
	assert.Contains(t, resultgroup2, groupName)

	// group-roles
	roleName := fmt.Sprintf("grouptestrole%d", timeSec)
	res2, err := client.IdentityService.GetGroupRoles(groupName)
	require.Nil(t, err)
	roleNum := len(res2)

	resultrole, err := client.IdentityService.CreateRole(roleName)
	defer client.IdentityService.DeleteRole(roleName)
	require.Nil(t, err)
	assert.Equal(t, roleName, resultrole.Name)
	assert.Equal(t, "test1@splunk.com", resultrole.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultrole.Tenant)

	resultrole1, err := client.IdentityService.AddRoleToGroup(groupName, roleName)
	defer client.IdentityService.RemoveGroupRole(groupName, roleName)
	require.Nil(t, err)
	assert.Equal(t, roleName, resultrole1.Role)
	assert.Equal(t, groupName, resultrole1.Group)
	assert.Equal(t, testutils.TestTenantID, resultrole1.Tenant)

	resultrole2, err := client.IdentityService.GetGroupRoles(groupName)
	require.Nil(t, err)
	assert.Equal(t, roleNum+1, len(resultrole2))
	assert.Contains(t, resultrole2, roleName)

	//group-members TODO: To replace with a generic user account
	memberName := "pupadhyaya@splunk.com"
	res3, err := client.IdentityService.GetGroupMembers(groupName)
	require.Nil(t, err)
	memberNum := len(res3)

	resultmember1, err := client.IdentityService.AddMemberToGroup(groupName, memberName)
	defer client.IdentityService.RemoveGroupMember(groupName, memberName)
	require.Nil(t, err)
	assert.Equal(t, memberName, resultmember1.Principal)
	assert.Equal(t, groupName, resultmember1.Group)
	assert.Equal(t, testutils.TestTenantID, resultmember1.Tenant)

	resultmember2, err := client.IdentityService.GetGroupMembers(groupName)
	require.Nil(t, err)
	assert.Equal(t, memberNum+2, len(resultmember2))
	assert.Contains(t, resultmember2, memberName)

	resultmember3, err := client.IdentityService.GetGroupMember(groupName, memberName)
	require.Nil(t, err)
	assert.Equal(t, memberName, resultmember3.Principal)

	//delete
	err = client.IdentityService.RemoveGroupMember(groupName, memberName)
	require.Nil(t, err)

	err = client.IdentityService.RemoveGroupRole(groupName, roleName)
	require.Nil(t, err)

	err = client.IdentityService.DeleteRole(roleName)
	require.Nil(t, err)

	err = client.IdentityService.DeleteGroup(groupName)
	require.Nil(t, err)
}

func TestCRUDRoles(t *testing.T) {
	client := getClient(t)

	res, err := client.IdentityService.GetRoles()
	require.Nil(t, err)
	roleNum := len(res)

	roleName := fmt.Sprintf("roletest%d", timeSec)

	// create/get/delete role and roles
	resultrole, err := client.IdentityService.CreateRole(roleName)
	defer client.IdentityService.DeleteRole(roleName)
	require.Nil(t, err)
	assert.Equal(t, roleName, resultrole.Name)
	assert.Equal(t, "test1@splunk.com", resultrole.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultrole.Tenant)

	resultrole1, err := client.IdentityService.GetRole(roleName)
	require.Nil(t, err)
	assert.Equal(t, roleName, resultrole1.Name)
	assert.Equal(t, "test1@splunk.com", resultrole1.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultrole1.Tenant)

	resultrole2, err := client.IdentityService.GetRoles()
	require.Nil(t, err)
	assert.Equal(t, roleNum+2, len(resultrole2))
	assert.Contains(t, resultrole2, roleName)

	// role-permissions
	permissionName := fmt.Sprintf("perm1-%d", timeSec)
	result1, err := client.IdentityService.GetRolePermissions(roleName)
	require.Nil(t, err)
	permNum := len(result1)

	resultroleperm, err := client.IdentityService.AddPermissionToRole(roleName, permissionName)
	defer client.IdentityService.RemoveRolePermission(roleName, permissionName)
	require.Nil(t, err)
	assert.Equal(t, roleName, resultroleperm.Role)
	assert.Equal(t, permissionName, resultroleperm.Permission)
	assert.Equal(t, "test1@splunk.com", resultroleperm.AddedBy)
	assert.Equal(t, testutils.TestTenantID, resultroleperm.Tenant)

	resultroleperm1, err := client.IdentityService.GetRolePermission(roleName, permissionName)
	require.Nil(t, err)
	assert.Equal(t, roleName, resultroleperm1.Role)
	assert.Equal(t, permissionName, resultroleperm1.Permission)
	assert.Equal(t, "test1@splunk.com", resultroleperm1.AddedBy)
	assert.Equal(t, testutils.TestTenantID, resultroleperm1.Tenant)

	resultroleperm2, err := client.IdentityService.GetRolePermissions(roleName)
	require.Nil(t, err)
	assert.Equal(t, permNum+1, len(resultroleperm2))
	assert.Contains(t, resultroleperm2, permissionName)

	// delete
	err = client.IdentityService.RemoveRolePermission(roleName, permissionName)
	require.Nil(t, err)
	err = client.IdentityService.DeleteRole(roleName)
	require.Nil(t, err)
}

func TestCRUDMembers(t *testing.T) {
	client := getClient(t)

	res, err := client.IdentityService.GetMembers()
	require.Nil(t, err)
	memNum := len(res)

	memberName := "ljiang@splunk.com"

	// create/get/delete member and members
	result, err := client.IdentityService.AddMember(memberName)
	defer client.IdentityService.RemoveMember(memberName)
	require.Nil(t, err)
	assert.Equal(t, memberName, result.Name)
	assert.Equal(t, testutils.TestTenantID, result.Tenant)

	result1, err := client.IdentityService.GetMembers()
	require.Nil(t, err)
	assert.Equal(t, memNum+1, len(result1))
	assert.Contains(t, result1, memberName)

	result2, err := client.IdentityService.GetMember(memberName)
	require.Nil(t, err)
	assert.Equal(t, memberName, result2.Name)
	assert.Equal(t, testutils.TestTenantID, result2.Tenant)

	// add member to group
	groupName := "users" //pre-defined group
	result3, err := client.IdentityService.AddMemberToGroup(groupName, memberName)
	defer client.IdentityService.RemoveGroupMember(groupName, memberName)
	require.Nil(t, err)
	assert.Equal(t, groupName, result3.Group)

	result4, err := client.IdentityService.GetMemberGroups(memberName)
	require.Nil(t, err)
	assert.Equal(t, 1, len(result4))
	assert.Contains(t, result4, groupName)

	roleName := "user" //pre-defind role
	result5, err := client.IdentityService.GetMemberRoles(memberName)
	require.Nil(t, err)
	assert.Equal(t, 1, len(result5))
	assert.Contains(t, result5, roleName)

	// add permission to role
	permissionName := "myperm"
	result6, err := client.IdentityService.AddPermissionToRole(roleName, permissionName)
	defer client.IdentityService.RemoveRolePermission(roleName, permissionName)
	require.Nil(t, err)
	assert.Equal(t, roleName, result6.Role)
	assert.Equal(t, permissionName, result6.Permission)

	permissionName1 := fmt.Sprintf("%v:users:identity.groups.read", testutils.TestTenantID)
	permissionName2 := fmt.Sprintf("%v:%v:identity.members.read", testutils.TestTenantID, memberName)
	result7, err := client.IdentityService.GetMemberPermissions(memberName)
	require.Nil(t, err)
	assert.Equal(t, 3, len(result7))
	assert.Contains(t, result7, permissionName)
	assert.Contains(t,result7, permissionName1)
	assert.Contains(t,result7, permissionName2)

	// delete
	err = client.IdentityService.RemoveMember(memberName)
	require.Nil(t, err)
}

func TestValidate(t *testing.T) {
	client := getClient(t)

	res, err := client.IdentityService.Validate()
	require.Nil(t, err)
	assert.Equal(t, "test1@splunk.com", res.Name)
}
