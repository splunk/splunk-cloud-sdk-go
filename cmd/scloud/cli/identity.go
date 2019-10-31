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

package main

import (
	"fmt"

	sdkIdentity "github.com/splunk/splunk-cloud-sdk-go/services/identity"
)

const (
	IdentityServiceVersion = "v2beta1"
)

var createIdentityService = func() *sdkIdentity.Service {
	return apiClient().IdentityService
}

type IdentityCommand struct {
	identityService *sdkIdentity.Service
}

func newIdentityCommand() *IdentityCommand {
	return &IdentityCommand{
		identityService: createIdentityService(),
	}
}
func (cmd *IdentityCommand) Dispatch(argv []string) (result interface{}, err error) {
	arg, argv := head(argv)
	switch arg {
	case "":
		eusage("too few arguments")
	case "add-group-member":
		result, err = addGroupMember(argv)
	case "add-group-role":
		result, err = addGroupRole(argv)
	case "add-member":
		result, err = addMember(argv)
	case "add-role-permission":
		result, err = addRolePermission(argv)
	case "create-group":
		result, err = createGroup(argv)
	case "create-role":
		result, err = createRole(argv)
	case "delete-group":
		err = deleteGroup(argv)
	case "delete-role":
		err = deleteRole(argv)
	case "get-group":
		result, err = getGroup(argv)
	case "get-group-member":
		result, err = getGroupMember(argv)
	case "get-group-role":
		result, err = getGroupRole(argv)
	case "get-member":
		result, err = getMember(argv)
	case "get-principal":
		result, err = getPrincipal(argv)
	case "get-role":
		result, err = getRole(argv)
	case "get-role-permission":
		result, err = getRolePermission(argv)
	case "get-spec-json":
		result, err = cmd.getSpecJSON(argv)
	case "get-spec-yaml":
		result, err = cmd.getSpecYaml(argv)
	case "help":
		result, err := getHelp("identity.txt")
		if err == nil {
			fmt.Println(result)
		}
	case "list-groups":
		result, err = listGroups(argv)
	case "list-group-members":
		result, err = listGroupMembers(argv)
	case "list-group-roles":
		result, err = listGroupRoles(argv)
	case "list-members":
		result, err = listMembers(argv)
	case "list-member-groups":
		result, err = listMemberGroups(argv)
	case "list-member-permissions":
		result, err = listMemberPermissions(argv)
	case "list-member-roles":
		result, err = listMemberRoles(argv)
	case "list-principals":
		result, err = listPrincipals(argv)
	case "list-roles":
		result, err = listRoles(argv)
	case "list-role-permissions":
		result, err = listRolePermissions(argv)
	case "remove-group-member":
		err = removeGroupMember(argv)
	case "remove-member":
		err = removeMember(argv)
	case "remove-group-role":
		err = removeGroupRole(argv)
	case "remove-role-permission":
		err = removeRolePermission(argv)
	case "validate-token":
		result, err = validateToken(argv)
	default:
		fatal("unknown command: '%s'", arg)
	}
	return
}

// identity returns an identity service client with a tenant set by the environment/user
func identity() *sdkIdentity.Service {
	return apiClient().IdentityService
}

// identitySystem returns an identity service client using the "system" tenant for system-wide calls
// that don't require a tenant
func identitySystem() *sdkIdentity.Service {
	return apiClientWithTenant("system").IdentityService
}

func addGroupMember(argv []string) (interface{}, error) {
	group, member := head2(argv)
	return identity().AddGroupMember(group, sdkIdentity.AddGroupMemberBody{Name: member})
}

func addGroupRole(argv []string) (interface{}, error) {
	group, role := head2(argv)
	return identity().AddGroupRole(group, sdkIdentity.AddGroupRoleBody{Name: role})
}

func addMember(argv []string) (interface{}, error) {
	member := head1(argv)
	return identity().AddMember(sdkIdentity.AddMemberBody{Name: member})
}

func addRolePermission(argv []string) (interface{}, error) {
	role, perm := head2(argv)
	return identity().AddRolePermission(role, perm)
}

func createGroup(argv []string) (interface{}, error) {
	group := head1(argv)
	return identity().CreateGroup(sdkIdentity.CreateGroupBody{Name: group})
}

func createRole(argv []string) (interface{}, error) {
	role := head1(argv)
	return identity().CreateRole(sdkIdentity.CreateRoleBody{Name: role})
}

func deleteGroup(argv []string) error {
	group := head1(argv)
	return identity().DeleteGroup(group)
}

func deleteRole(argv []string) error {
	role := head1(argv)
	return identity().DeleteRole(role)
}

func getGroup(argv []string) (interface{}, error) {
	group := head1(argv)
	return identity().GetGroup(group)
}

func getGroupMember(argv []string) (interface{}, error) {
	group, member := head2(argv)
	return identity().GetGroupMember(group, member)
}

func getGroupRole(argv []string) (interface{}, error) {
	group, role := head2(argv)
	return identity().GetGroupRole(group, role)
}

func getMember(argv []string) (interface{}, error) {
	member := head1(argv)
	return identity().GetMember(member)
}

func getPrincipal(argv []string) (interface{}, error) {
	principal := head1(argv)
	return identitySystem().GetPrincipal(principal)
}

func getRole(argv []string) (interface{}, error) {
	role := head1(argv)
	return identity().GetRole(role)
}

func getRolePermission(argv []string) (interface{}, error) {
	role, perm := head2(argv)
	return identity().GetRolePermission(role, perm)
}

func listGroups(argv []string) (interface{}, error) {
	return identity().ListGroups()
}

func listGroupMembers(argv []string) (interface{}, error) {
	group := head1(argv)
	return identity().ListGroupMembers(group)
}

func listGroupRoles(argv []string) (interface{}, error) {
	group := head1(argv)
	return identity().ListGroupRoles(group)
}

func listMembers(argv []string) (interface{}, error) {
	return identity().ListMembers()
}

func listMemberGroups(argv []string) (interface{}, error) {
	member := head1(argv)
	return identity().ListMemberGroups(member)
}

func listMemberPermissions(argv []string) (interface{}, error) {
	member := head1(argv)
	return identity().ListMemberPermissions(member)
}

func listMemberRoles(argv []string) (interface{}, error) {
	member := head1(argv)
	return identity().ListMemberRoles(member)
}

func listPrincipals(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return identitySystem().ListPrincipals()
}

func listRoles(argv []string) (interface{}, error) {
	return identity().ListRoles()
}

func listRolePermissions(argv []string) (interface{}, error) {
	role := head1(argv)
	return identity().ListRolePermissions(role)
}

func removeGroupMember(argv []string) error {
	group, member := head2(argv)
	return identity().RemoveGroupMember(group, member)
}

func removeMember(argv []string) error {
	member := head1(argv)
	return identity().RemoveMember(member)
}

func removeGroupRole(argv []string) error {
	group, role := head2(argv)
	return identity().RemoveGroupRole(group, role)
}

func removeRolePermission(argv []string) error {
	role, perm := head2(argv)
	return identity().RemoveRolePermission(role, perm)
}

func validateToken(argv []string) (interface{}, error) {
	tokenParams := sdkIdentity.ValidateTokenQueryParams{Include: []string{}}
	return identity().ValidateToken(&tokenParams)
}

func (cmd *IdentityCommand) getSpecJSON(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecJSON("api", IdentityServiceVersion, "identity", cmd.identityService.Client)
}

func (cmd *IdentityCommand) getSpecYaml(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecYaml("api", IdentityServiceVersion, "identity", cmd.identityService.Client)
}
