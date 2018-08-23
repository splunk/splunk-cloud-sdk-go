// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package stubbyintegration

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIdentityService_CreateTenant(t *testing.T) {
	result, err := getClient(t).IdentityService.CreateTenant("devtestTenant")
	assert.Nil(t, err)
	assert.Equal(t, "devtestTenant", result.Name)
	assert.Equal(t, "provisioning", result.Status)
}

func TestIdentityService_DeleteTenant(t *testing.T) {
	err := getClient(t).IdentityService.DeleteTenant("devtestTenant")
	assert.Nil(t, err)
}

func TestIdentityService_GetTenant(t *testing.T) {
	result, err := getClient(t).IdentityService.GetTenant("devtestTenant")
	assert.Nil(t, err)
	assert.Equal(t, "devtestTenant", result.Name)
}

func TestIdentityService_GetTenants(t *testing.T) {
	result, err := getClient(t).IdentityService.GetTenants()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "devtestTenant", result[0])

}

func TestIdentityService_GetValidate(t *testing.T) {
	result, err := getClient(t).IdentityService.ValidateTenant()
	assert.Nil(t, err)
	assert.Equal(t, "TEST_TENANT", result.Name)
	assert.Equal(t, 2, len(result.Tenants))
}

func TestIdentityService_CreatePrincipal(t *testing.T) {
	result, err := getClient(t).IdentityService.CreatePrincipal("mem1", "user")
	assert.Nil(t, err)
	assert.Equal(t, "mem1", result.Name)
	assert.Equal(t, "user", result.Kind)
	assert.NotNil(t, result.Profile)
}

func TestIdentityService_DeletePrincipal(t *testing.T) {
	err := getClient(t).IdentityService.DeletePrincipal("devtestTenant")
	assert.Nil(t, err)
}

func TestIdentityService_GetPrincipal(t *testing.T) {
	result, err := getClient(t).IdentityService.GetPrincipal("mem1")
	assert.Nil(t, err)
	assert.Equal(t, "mem1", result.Name)
}

func TestIdentityService_GetPrincipals(t *testing.T) {
	result, err := getClient(t).IdentityService.GetPrincipals()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "mem2", result[1])
}

func TestIdentityService_GetMember(t *testing.T) {
	result, err := getClient(t).IdentityService.GetMember("mem1")
	assert.Nil(t, err)
	assert.Equal(t, "mem1", result.Name)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_GetMembers(t *testing.T) {
	result, err := getClient(t).IdentityService.GetMembers()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "mem1", result[0])
}

func TestIdentityService_GetMemberGroups(t *testing.T) {
	result, err := getClient(t).IdentityService.GetMemberGroups("mem1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "group1", result[0])
}

func TestIdentityService_GetMemberPermissions(t *testing.T) {
	result, err := getClient(t).IdentityService.GetMemberPermissions("mem1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "perm2", result[1])
}

func TestIdentityService_GetMemberRoles(t *testing.T) {
	result, err := getClient(t).IdentityService.GetMemberRoles("mem1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "role1", result[0])
}

func TestIdentityService_GetRoles(t *testing.T) {
	result, err := getClient(t).IdentityService.GetRoles()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "role1", result[0])
}

func TestIdentityService_GetRole(t *testing.T) {
	result, err := getClient(t).IdentityService.GetRole("role2")
	assert.Nil(t, err)
	assert.Equal(t, "role2", result.Name)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_GetRolePermissions(t *testing.T) {
	result, err := getClient(t).IdentityService.GetRolePermissions("role2")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "perm1", result[0])
}

func TestIdentityService_GetRolePermission(t *testing.T) {
	result, err := getClient(t).IdentityService.GetRolePermission("role2", "perm2")
	assert.Nil(t, err)
	assert.Equal(t, "role2", result.Role)
	assert.Equal(t, "perm2", result.Permission)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_GetGroups(t *testing.T) {
	result, err := getClient(t).IdentityService.GetGroups()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "grp1", result[0])
}

func TestIdentityService_GetGroup(t *testing.T) {
	result, err := getClient(t).IdentityService.GetGroup("grp1")
	assert.Nil(t, err)
	assert.Equal(t, "grp1", result.Name)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_GetGroupRoles(t *testing.T) {
	result, err := getClient(t).IdentityService.GetGroupRoles("grp1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "role1", result[0])
}

func TestIdentityService_GetGroupRole(t *testing.T) {
	result, err := getClient(t).IdentityService.GetGroupRole("grp1", "role1")
	assert.Nil(t, err)
	assert.Equal(t, "grp1", result.Group)
	assert.Equal(t, "role1", result.Role)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_GetGroupMembers(t *testing.T) {
	result, err := getClient(t).IdentityService.GetGroupMembers("grp1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "mem1", result[0])
}

func TestIdentityService_GetGroupMember(t *testing.T) {
	result, err := getClient(t).IdentityService.GetGroupMember("grp1", "mem1")
	assert.Nil(t, err)
	assert.Equal(t, "mem1", result.Principal)
	assert.Equal(t, "grp1", result.Group)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_CreateGroup(t *testing.T) {
	result, err := getClient(t).IdentityService.CreateGroup("sdk-group")
	assert.Nil(t, err)
	assert.Equal(t, "sdk-group", result.Name)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_CreateRole(t *testing.T) {
	result, err := getClient(t).IdentityService.CreateRole("roles.sdk-test")
	assert.Nil(t, err)
	assert.Equal(t, "roles.sdk-test", result.Name)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_AddMember(t *testing.T) {
	result, err := getClient(t).IdentityService.AddMember("mem1")
	assert.Nil(t, err)
	assert.Equal(t, "mem1", result.Name)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_AddMemberToGroup(t *testing.T) {
	result, err := getClient(t).IdentityService.AddMemberToGroup("sdk-group", "sdk-int-test@splunk.com")
	assert.Nil(t, err)
	assert.Equal(t, "sdk-int-test@splunk.com", result.Principal)
	assert.Equal(t, "sdk-group", result.Group)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_AddPermissionToRole(t *testing.T) {
	result, err := getClient(t).IdentityService.AddPermissionToRole("roles.sdk-test", "TEST_TENANT%3A%2A%3Akvstore.%2A")
	assert.Nil(t, err)
	assert.Equal(t, "TEST_TENANT%3A%2A%3Akvstore.%2A", result.Permission)
	assert.Equal(t, "roles.sdk-test", result.Role)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_AddRoleToGroup(t *testing.T) {
	result, err := getClient(t).IdentityService.AddRoleToGroup("sdk-group", "sdk-test-role")
	assert.Nil(t, err)
	assert.Equal(t, "sdk-test-role", result.Role)
	assert.Equal(t, "sdk-group", result.Group)
	assert.Equal(t, "TEST_TENANT", result.Tenant)
}

func TestIdentityService_DeleteGroup(t *testing.T) {
	err := getClient(t).IdentityService.DeleteGroup("sdk-group")
	assert.Nil(t, err)
}

func TestIdentityService_DeleteMember(t *testing.T) {
	err := getClient(t).IdentityService.DeleteMember("mem1")
	assert.Nil(t, err)
}

func TestIdentityService_DeleteRole(t *testing.T) {
	err := getClient(t).IdentityService.DeleteRole("roles.sdk-test")
	assert.Nil(t, err)
}

func TestIdentityService_RemoveGroupMember(t *testing.T) {
	err := getClient(t).IdentityService.RemoveGroupMember("sdk-group", "sdk-int-test@splunk.com")
	assert.Nil(t, err)
}

func TestIdentityService_RemoveGroupRole(t *testing.T) {
	err := getClient(t).IdentityService.RemoveGroupRole("sdk-group", "roles.sdk-test")
	assert.Nil(t, err)
}

func TestIdentityService_RemoveRolePermission(t *testing.T) {
	err := getClient(t).IdentityService.RemoveRolePermission("roles.sdk-test", "TEST_TENANT:*:kvstore.*")
	assert.Nil(t, err)
}
