/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package integration

import (
	"fmt"
	"testing"

	"time"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/identity"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestIdentityClientInit tests initializing an identity service-specific Splunk Cloud client and validating the token provided
func TestIdentityClientInit(t *testing.T) {
	identityClient, err := identity.NewService(&services.Config{
		Token:  testutils.TestAuthenticationToken,
		Host:   testutils.TestSplunkCloudHost,
		Tenant: testutils.TestTenant,
	})
	require.Emptyf(t, err, "error calling services.NewService(): %s", err)
	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}
	info, err := identityClient.ValidateToken(&input)
	assert.Emptyf(t, err, "error calling identityClient.Validate(): %s", err)
	assert.NotNil(t, info)
}

func TestCRUDGroups(t *testing.T) {
	client := getClient(t)

	_, err := client.IdentityService.ListGroups(nil)
	require.Nil(t, err)

	groupName := fmt.Sprintf("grouptest%d", testutils.TimeSec)

	// create/get/delete group and groups
	resultgroup, err := client.IdentityService.CreateGroup(identity.CreateGroupBody{Name: groupName})
	require.Nil(t, err)
	defer client.IdentityService.DeleteGroup(groupName)
	assert.Equal(t, groupName, resultgroup.Name)
	assert.Equal(t, testutils.TestUsername, resultgroup.CreatedBy)
	assert.Equal(t, testutils.TestTenant, resultgroup.Tenant)

	time.Sleep(2 * time.Second)
	resultgroup1, err := client.IdentityService.GetGroup(groupName)
	require.Nil(t, err)
	assert.Equal(t, groupName, resultgroup1.Name)
	assert.Equal(t, testutils.TestUsername, resultgroup1.CreatedBy)
	assert.Equal(t, testutils.TestTenant, resultgroup1.Tenant)

	resultgroup2, err := client.IdentityService.ListGroups(nil)
	require.Nil(t, err)
	assert.Contains(t, resultgroup2, groupName)

	// group-roles
	roleName := fmt.Sprintf("grouptestrole%d", testutils.TimeSec)
	_, err = client.IdentityService.ListGroupRoles(groupName)
	require.Nil(t, err)

	resultrole, err := client.IdentityService.CreateRole(identity.CreateRoleBody{Name: roleName})
	require.Nil(t, err)
	defer client.IdentityService.DeleteRole(roleName)
	assert.Equal(t, roleName, resultrole.Name)
	assert.Equal(t, testutils.TestUsername, resultrole.CreatedBy)
	assert.Equal(t, testutils.TestTenant, resultrole.Tenant)

	time.Sleep(2 * time.Second)
	resultrole1, err := client.IdentityService.AddGroupRole(groupName, identity.AddGroupRoleBody{Name: roleName})
	require.Nil(t, err)
	defer client.IdentityService.RemoveGroupRole(groupName, roleName)
	assert.Equal(t, roleName, resultrole1.Role)
	assert.Equal(t, groupName, resultrole1.Group)
	assert.Equal(t, testutils.TestTenant, resultrole1.Tenant)

	time.Sleep(2 * time.Second)
	resultrole2, err := client.IdentityService.ListGroupRoles(groupName)
	require.Nil(t, err)
	assert.Contains(t, resultrole2, roleName)

	time.Sleep(2 * time.Second)
	roleGroups, err := client.IdentityService.ListRoleGroups(roleName)
	require.Nil(t, err)
	assert.Contains(t, roleGroups, groupName)

	//group-members
	memberName := "test1@splunk.com"
	_, err = client.IdentityService.ListGroupMembers(groupName)
	require.Nil(t, err)

	//add group member
	_, err = client.IdentityService.AddMember(identity.AddMemberBody{Name: memberName})
	require.Nil(t, err)
	defer client.IdentityService.RemoveMember(memberName)

	resultmember1, err := client.IdentityService.AddGroupMember(groupName, identity.AddGroupMemberBody{Name: memberName})
	require.Nil(t, err)
	defer client.IdentityService.RemoveGroupMember(groupName, memberName)
	assert.Equal(t, memberName, resultmember1.Principal)
	assert.Equal(t, groupName, resultmember1.Group)
	assert.Equal(t, testutils.TestTenant, resultmember1.Tenant)

	time.Sleep(2 * time.Second)
	resultmember2, err := client.IdentityService.ListGroupMembers(groupName)
	require.Nil(t, err)
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

	_, err := client.IdentityService.ListRoles()
	require.Nil(t, err)

	roleName := fmt.Sprintf("roletest%d", testutils.TimeSec)

	// create/get/delete role and roles
	resultrole, err := client.IdentityService.CreateRole(identity.CreateRoleBody{Name: roleName})
	require.Nil(t, err)
	defer client.IdentityService.DeleteRole(roleName)
	assert.Equal(t, roleName, resultrole.Name)
	assert.Equal(t, testutils.TestUsername, resultrole.CreatedBy)
	assert.Equal(t, testutils.TestTenant, resultrole.Tenant)

	time.Sleep(2 * time.Second)
	resultrole1, err := client.IdentityService.GetRole(roleName)
	require.Nil(t, err)
	assert.Equal(t, roleName, resultrole1.Name)
	assert.Equal(t, testutils.TestUsername, resultrole1.CreatedBy)
	assert.Equal(t, testutils.TestTenant, resultrole1.Tenant)

	resultrole2, err := client.IdentityService.ListRoles()
	require.Nil(t, err)
	assert.Contains(t, resultrole2, roleName)

	// role-permissions
	_, err = client.IdentityService.ListRolePermissions(roleName)
	require.Nil(t, err)

	permissionName := fmt.Sprintf("%v:all:perm1.%d", testutils.TestTenant, testutils.TimeSec)
	resultroleperm, err := client.IdentityService.AddRolePermission(roleName, permissionName)
	require.Nil(t, err)
	defer client.IdentityService.RemoveRolePermission(roleName, permissionName)
	assert.Equal(t, roleName, resultroleperm.Role)
	assert.Equal(t, permissionName, resultroleperm.Permission)
	assert.Equal(t, testutils.TestUsername, resultroleperm.AddedBy)
	assert.Equal(t, testutils.TestTenant, resultroleperm.Tenant)

	time.Sleep(2 * time.Second)
	resultroleperm1, err := client.IdentityService.GetRolePermission(roleName, permissionName)
	require.Nil(t, err)
	assert.Equal(t, roleName, resultroleperm1.Role)
	assert.Equal(t, permissionName, resultroleperm1.Permission)
	assert.Equal(t, testutils.TestUsername, resultroleperm1.AddedBy)
	assert.Equal(t, testutils.TestTenant, resultroleperm1.Tenant)

	resultroleperm2, err := client.IdentityService.ListRolePermissions(roleName)
	require.Nil(t, err)
	assert.Contains(t, resultroleperm2, permissionName)

	// delete
	err = client.IdentityService.RemoveRolePermission(roleName, permissionName)
	require.Nil(t, err)
	err = client.IdentityService.DeleteRole(roleName)
	require.Nil(t, err)
}

func TestCRUDMembers(t *testing.T) {
	client := getClient(t)

	_, err := client.IdentityService.ListMembers()
	require.Nil(t, err)

	memberName := "test1@splunk.com"

	// create/get/delete member and members
	result, err := client.IdentityService.AddMember(identity.AddMemberBody{Name: memberName})
	require.Nil(t, err)
	defer client.IdentityService.RemoveMember(memberName)
	assert.Equal(t, memberName, result.Name)
	assert.Equal(t, testutils.TestTenant, result.Tenant)

	time.Sleep(2 * time.Second)
	result1, err := client.IdentityService.ListMembers()
	require.Nil(t, err)
	assert.Contains(t, result1, memberName)

	result2, err := client.IdentityService.GetMember(memberName)
	require.Nil(t, err)
	assert.Equal(t, memberName, result2.Name)
	assert.Equal(t, testutils.TestTenant, result2.Tenant)

	groupName := fmt.Sprintf("grouptest%d", testutils.TimeSec)

	// create a group
	resultgroup, err := client.IdentityService.CreateGroup(identity.CreateGroupBody{Name: groupName})
	require.Nil(t, err)
	defer client.IdentityService.DeleteGroup(groupName)
	assert.Equal(t, groupName, resultgroup.Name)
	assert.Equal(t, testutils.TestUsername, resultgroup.CreatedBy)
	assert.Equal(t, testutils.TestTenant, resultgroup.Tenant)

	// add member to group
	result3, err := client.IdentityService.AddGroupMember(groupName, identity.AddGroupMemberBody{Name: memberName})
	require.Nil(t, err)
	defer client.IdentityService.RemoveGroupMember(groupName, memberName)
	assert.Equal(t, groupName, result3.Group)

	time.Sleep(2 * time.Second)
	result4, err := client.IdentityService.ListMemberGroups(memberName)
	require.Nil(t, err)
	assert.Contains(t, result4, groupName)

	// group-role
	roleName := fmt.Sprintf("grouptestrole%d", testutils.TimeSec)

	// create a test role
	resultrole, err := client.IdentityService.CreateRole(identity.CreateRoleBody{Name: roleName})
	require.Nil(t, err)
	defer client.IdentityService.DeleteRole(roleName)
	assert.Equal(t, roleName, resultrole.Name)
	assert.Equal(t, testutils.TestUsername, resultrole.CreatedBy)
	assert.Equal(t, testutils.TestTenant, resultrole.Tenant)

	// add role to group
	time.Sleep(2 * time.Second)
	resultrole1, err := client.IdentityService.AddGroupRole(groupName, identity.AddGroupRoleBody{Name: roleName})
	require.Nil(t, err)
	defer client.IdentityService.RemoveGroupRole(groupName, roleName)
	assert.Equal(t, roleName, resultrole1.Role)
	assert.Equal(t, groupName, resultrole1.Group)
	assert.Equal(t, testutils.TestTenant, resultrole1.Tenant)

	//get group role relationship
	time.Sleep(2 * time.Second)
	groupRole, err := client.IdentityService.GetGroupRole(groupName, roleName)
	require.Nil(t, err)
	assert.Equal(t, roleName, groupRole.Role)
	assert.Equal(t, groupName, groupRole.Group)
	assert.Equal(t, testutils.TestTenant, groupRole.Tenant)

	time.Sleep(2 * time.Second)
	result5, err := client.IdentityService.ListMemberRoles(memberName)
	require.Nil(t, err)
	assert.Contains(t, result5, roleName)

	// add permission to role
	permissionName := fmt.Sprintf("%v:%v:myperm.%d", testutils.TestTenant, groupName, testutils.TimeSec)
	result6, err := client.IdentityService.AddRolePermission(roleName, permissionName)
	require.Nil(t, err)
	defer client.IdentityService.RemoveRolePermission(roleName, permissionName)
	assert.Equal(t, roleName, result6.Role)
	assert.Equal(t, permissionName, result6.Permission)

	time.Sleep(2 * time.Second)
	permissionName1 := fmt.Sprintf("%v:%v:identity.groups.read", testutils.TestTenant, groupName)
	permissionName2 := fmt.Sprintf("%v:%v:identity.members.read", testutils.TestTenant, memberName)
	result7, err := client.IdentityService.ListMemberPermissions(memberName)
	require.Nil(t, err)
	assert.Contains(t, result7, permissionName)
	assert.Contains(t, result7, permissionName1)
	assert.Contains(t, result7, permissionName2)

	// delete the test member
	err = client.IdentityService.RemoveMember(memberName)
	require.Nil(t, err)
}

func TestPrincipals(t *testing.T) {
	client := getClient(t)

	principals, err := client.IdentityService.ListPrincipals()
	require.Nil(t, err)
	assert.Equal(t, principals[0], testutils.TestUsername)

	principal, err := client.IdentityService.GetPrincipal(testutils.TestUsername)
	require.Nil(t, err)
	assert.Equal(t, principal.Name, testutils.TestUsername)
}

func TestValidate(t *testing.T) {
	client := getClient(t)
	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}
	res, err := client.IdentityService.ValidateToken(&input)
	require.Nil(t, err)
	assert.Equal(t, testutils.TestUsername, res.Name)
}
