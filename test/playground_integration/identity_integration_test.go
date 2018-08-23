// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"fmt"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCRUDGroups(t *testing.T) {
	client := getClient(t)

	res, err := client.IdentityService.GetGroups()
	assert.Emptyf(t, err, "Error:%s", err)
	groupNum := len(res)

	groupName := fmt.Sprintf("grouptest%d", timeSec)

	// create/get/delete group and groups
	resultgroup, err := client.IdentityService.CreateGroup(groupName)
	defer client.IdentityService.DeleteGroup(groupName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, groupName, resultgroup.Name)
	assert.Equal(t, "test1@splunk.com", resultgroup.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultgroup.Tenant)

	resultgroup1, err := client.IdentityService.GetGroup(groupName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, groupName, resultgroup1.Name)
	assert.Equal(t, "test1@splunk.com", resultgroup1.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultgroup1.Tenant)

	resultgroup2, err := client.IdentityService.GetGroups()
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, groupNum+1, len(resultgroup2))
	assert.Contains(t, resultgroup2, groupName)

	// group-roles
	roleName := fmt.Sprintf("grouptestrole%d", timeSec)
	res2, err := client.IdentityService.GetGroupRoles(groupName)
	assert.Emptyf(t, err, "Error:%s", err)
	roleNum := len(res2)

	resultrole, err := client.IdentityService.CreateRole(roleName)
	defer client.IdentityService.DeleteRole(roleName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, roleName, resultrole.Name)
	assert.Equal(t, "test1@splunk.com", resultrole.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultrole.Tenant)

	resultrole1, err := client.IdentityService.AddRoleToGroup(groupName, roleName)
	defer client.IdentityService.RemoveGroupRole(groupName, roleName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, roleName, resultrole1.Role)
	assert.Equal(t, groupName, resultrole1.Group)
	assert.Equal(t, testutils.TestTenantID, resultrole1.Tenant)

	resultrole2, err := client.IdentityService.GetGroupRoles(groupName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, roleNum+1, len(resultrole2))
	assert.Contains(t, resultrole2, roleName)

	//group-members
	memberName := "test1@splunk.com"
	res3, err := client.IdentityService.GetGroupMembers(groupName)
	assert.Emptyf(t, err, "Error:%s", err)
	memberNum := len(res3)

	resultmember1, err := client.IdentityService.AddMemberToGroup(groupName, memberName)
	defer client.IdentityService.RemoveGroupMember(groupName, memberName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, memberName, resultmember1.Principal)
	assert.Equal(t, groupName, resultmember1.Group)
	assert.Equal(t, testutils.TestTenantID, resultmember1.Tenant)

	resultmember2, err := client.IdentityService.GetGroupMembers(groupName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, memberNum+1, len(resultmember2))
	assert.Contains(t, resultmember2, memberName)

	resultmember3, err := client.IdentityService.GetGroupMember(groupName, memberName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, memberName, resultmember3.Principal)

	//delete
	err = client.IdentityService.RemoveGroupMember(groupName, memberName)
	assert.Emptyf(t, err, "Error:%s", err)

	err = client.IdentityService.RemoveGroupRole(groupName, roleName)
	assert.Emptyf(t, err, "Error:%s", err)

	err = client.IdentityService.DeleteRole(roleName)
	assert.Emptyf(t, err, "Error:%s", err)

	err = client.IdentityService.DeleteGroup(groupName)
	assert.Emptyf(t, err, "Error:%s", err)
}

func TestCRUDRoles(t *testing.T) {
	client := getClient(t)

	res, err := client.IdentityService.GetRoles()
	assert.Emptyf(t, err, "Error:%s", err)
	roleNum := len(res)

	roleName := fmt.Sprintf("roletest%d", timeSec)

	// create/get/delete role and roles
	resultrole, err := client.IdentityService.CreateRole(roleName)
	defer client.IdentityService.DeleteRole(roleName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, roleName, resultrole.Name)
	assert.Equal(t, "test1@splunk.com", resultrole.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultrole.Tenant)

	resultrole1, err := client.IdentityService.GetRole(roleName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, roleName, resultrole1.Name)
	assert.Equal(t, "test1@splunk.com", resultrole1.CreatedBy)
	assert.Equal(t, testutils.TestTenantID, resultrole1.Tenant)

	resultrole2, err := client.IdentityService.GetRoles()
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, roleNum+1, len(resultrole2))
	assert.Contains(t, resultrole2, roleName)

	// role-permissions
	permissionName := fmt.Sprintf("perm1-%d", timeSec)
	result1, err := client.IdentityService.GetRolePermissions(roleName)
	assert.Emptyf(t, err, "Error:%s", err)
	permNum := len(result1)

	resultroleperm, err := client.IdentityService.AddPermissionToRole(roleName, permissionName)
	defer client.IdentityService.RemoveRolePermission(roleName, permissionName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, roleName, resultroleperm.Role)
	assert.Equal(t, permissionName, resultroleperm.Permission)
	assert.Equal(t, "test1@splunk.com", resultroleperm.AddedBy)
	assert.Equal(t, testutils.TestTenantID, resultroleperm.Tenant)

	resultroleperm1, err := client.IdentityService.GetRolePermission(roleName, permissionName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, roleName, resultroleperm1.Role)
	assert.Equal(t, permissionName, resultroleperm1.Permission)
	assert.Equal(t, "test1@splunk.com", resultroleperm1.AddedBy)
	assert.Equal(t, testutils.TestTenantID, resultroleperm1.Tenant)

	resultroleperm2, err := client.IdentityService.GetRolePermissions(roleName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, permNum+1, len(resultroleperm2))
	assert.Contains(t, resultroleperm2, permissionName)

	// delete
	err = client.IdentityService.RemoveRolePermission(roleName, permissionName)
	assert.Emptyf(t, err, "Error:%s", err)
	err = client.IdentityService.DeleteRole(roleName)
	assert.Emptyf(t, err, "Error:%s", err)
}

func TestCRUDMembers(t *testing.T) {
	client := getClient(t)

	res, err := client.IdentityService.GetMembers()
	assert.Emptyf(t, err, "Error:%s", err)
	memNum := len(res)

	memberName := "ljiang@splunk.com"

	// create/get/delete member and members
	result, err := client.IdentityService.AddMember(memberName)
	defer client.IdentityService.DeleteMember(memberName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, memberName, result.Name)
	assert.Equal(t, testutils.TestTenantID, result.Tenant)

	result1, err := client.IdentityService.GetMembers()
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, memNum+1, len(result1))
	assert.Contains(t, result1, memberName)

	result3, err := client.IdentityService.GetMember(memberName)
	assert.Emptyf(t, err, "Error:%s", err)
	assert.Equal(t, memberName, result3.Name)
	assert.Equal(t, testutils.TestTenantID, result3.Tenant)
	// delete
	err = client.IdentityService.DeleteMember(memberName)
	assert.Emptyf(t, err, "Error:%s", err)
}
