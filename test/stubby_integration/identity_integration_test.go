// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package stubbyintegration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func TestIdentityService_GetUserProfile_TenantId(t *testing.T) {
//	user, err := getClient(t).IdentityService.GetUserProfile("devtestTenant")
//	assert.Nil(t, err)
//	assert.Equal(t, "devtest@splunk.com", user.ID)
//	assert.Equal(t, "devtest@splunk.com", user.Email)
//	assert.Equal(t, "Dev", user.FirstName)
//	assert.Equal(t, "Test", user.LastName)
//	assert.Equal(t, "Dev Test", user.Name)
//	assert.Equal(t, "en-US", user.Locale)
//	assert.Equal(t, 1, len(user.TenantMemberships))
//	assert.Equal(t, "devtestTenant", user.TenantMemberships[0])
//}
//
//func TestIdentityService_GetUserProfile_System(t *testing.T) {
//	user, err := getClient(t).IdentityService.GetUserProfile("")
//	assert.Nil(t, err)
//	assert.Equal(t, "devtest@splunk.com", user.ID)
//	assert.Equal(t, "devtest@splunk.com", user.Email)
//	assert.Equal(t, "Dev", user.FirstName)
//	assert.Equal(t, "Test", user.LastName)
//	assert.Equal(t, "Dev Test", user.Name)
//	assert.Equal(t, "en-US", user.Locale)
//	assert.Equal(t, 1, len(user.TenantMemberships))
//	assert.Equal(t, "devtestTenant", user.TenantMemberships[0])
//}
//
//func TestIdentityService_CreateTenant(t *testing.T) {
//	err := getClient(t).IdentityService.CreateTenant(model.Tenant{TenantID: "devtestTenant"})
//	assert.Nil(t, err)
//}
//
//func TestIdentityService_DeleteTenant(t *testing.T) {
//	err := getClient(t).IdentityService.DeleteTenant("devtestTenant")
//	assert.Nil(t, err)
//}
//
//func TestIdentityService_GetTenantUsers(t *testing.T) {
//	users, err := getClient(t).IdentityService.GetTenantUsers("devtestTenant")
//	assert.Nil(t, err)
//	assert.Equal(t, 1, len(users))
//	assert.Equal(t, "devtest1@splunk.com", users[0].ID)
//}
//
//func TestIdentityService_ReplaceTenantUsers(t *testing.T) {
//	err := getClient(t).IdentityService.ReplaceTenantUsers("devtestTenant", []model.User{
//		{ID: "devtest2@splunk.com"},
//		{ID: "devtest3@splunk.com"},
//		{ID: "devtest4@splunk.com"},
//		{ID: "devtest5@splunk.com"},
//	})
//	assert.Nil(t, err)
//}
//
//func TestIdentityService_AddTenantUsers(t *testing.T) {
//	err := getClient(t).IdentityService.AddTenantUsers("devtestTenant", []model.User{
//		{ID: "devtest7@splunk.com"},
//		{ID: "devtest8@splunk.com"},
//	})
//	assert.Nil(t, err)
//}
//
//func TestIdentityService_DeleteTenantUsers(t *testing.T) {
//	err := getClient(t).IdentityService.DeleteTenantUsers("devtestTenant", []model.User{
//		{ID: "devtest4@splunk.com"},
//		{ID: "devtest5@splunk.com"},
//	})
//	assert.Nil(t, err)
//}

func TestIdentityService_GetMember(t *testing.T) {
	result, err := getClient(t).IdentityService.GetMember("mem1")
	assert.Nil(t, err)
	assert.Equal(t, "mem1", result.Name)
	assert.Equal(t, "devtestTenant", result.Tenant)
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
	assert.Equal(t, "devtestTenant", result.Tenant)
}

func TestIdentityService_GetRolePermissions(t *testing.T) {
	result, err := getClient(t).IdentityService.GetRolePermissions("role2")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "perm1", result[0])
}

func TestIdentityService_GetRolePermission(t *testing.T) {
	result, err := getClient(t).IdentityService.GetRolePermission("role2","perm2")
	assert.Nil(t, err)
	assert.Equal(t, "role2", result.Role)
	assert.Equal(t, "perm2", result.Permission)
	assert.Equal(t, "devtestTenant", result.Tenant)
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
	assert.Equal(t, "devtestTenant", result.Tenant)
}

func TestIdentityService_GetGroupRoles(t *testing.T) {
	result, err := getClient(t).IdentityService.GetGroupRoles("grp1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "role1", result[0])
}

func TestIdentityService_GetGroupRole(t *testing.T) {
	result, err := getClient(t).IdentityService.GetGroupRole("grp1","role1")
	assert.Nil(t, err)
	assert.Equal(t, "grp1", result.Group)
	assert.Equal(t, "role1", result.Role)
	assert.Equal(t, "devtestTenant", result.Tenant)
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
	assert.Equal(t, "devtestTenant", result.Tenant)
}
