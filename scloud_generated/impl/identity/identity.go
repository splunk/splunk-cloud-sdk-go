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

package identity

import (
	"fmt"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/utils"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	sdkIdentity "github.com/splunk/splunk-cloud-sdk-go/services/identity"
)

//const (
//	IdentityServiceVersion = "v2beta1"
//)
//
//var createIdentityService = func() *sdkIdentity.Service {
//	return apiClient().IdentityService
//}
//
//type IdentityCommand struct {
//	identityService *sdkIdentity.Service
//}
//
//func newIdentityCommand() *IdentityCommand {
//	return &IdentityCommand{
//		identityService: createIdentityService(),
//	}
//}
//func (cmd *IdentityCommand) Dispatch(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (result interface{}, err error) {
//	arg, argv := head(argv)
//	switch arg {
//	case "":
//		eusage("too few arguments")
//	case "add-group-member":
//		result, err = addGroupMember(argv)
//	case "add-group-role":
//		result, err = addGroupRole(argv)
//	case "add-member":
//		result, err = addMember(argv)
//	case "add-role-permission":
//		result, err = addRolePermission(argv)
//	case "create-group":
//		result, err = createGroup(argv)
//	case "create-role":
//		result, err = createRole(argv)
//	case "delete-group":
//		err = deleteGroup(argv)
//	case "delete-role":
//		err = deleteRole(argv)
//	case "get-group":
//		result, err = getGroup(argv)
//	case "get-group-member":
//		result, err = getGroupMember(argv)
//	case "get-group-role":
//		result, err = getGroupRole(argv)
//	case "get-member":
//		result, err = getMember(argv)
//	case "get-principal":
//		result, err = getPrincipal(argv)
//	case "get-role":
//		result, err = getRole(argv)
//	case "get-role-permission":
//		result, err = getRolePermission(argv)
//	case "get-spec-json":
//		result, err = cmd.getSpecJSON(argv)
//	case "get-spec-yaml":
//		result, err = cmd.getSpecYaml(argv)
//	case "help":
//		err = help("identity.txt")
//	case "list-groups":
//		result, err = listGroups(argv)
//	case "list-group-members":
//		result, err = listGroupMembers(argv)
//	case "list-group-roles":
//		result, err = listGroupRoles(argv)
//	case "list-members":
//		result, err = listMembers(argv)
//	case "list-member-groups":
//		result, err = listMemberGroups(argv)
//	case "list-member-permissions":
//		result, err = listMemberPermissions(argv)
//	case "list-member-roles":
//		result, err = listMemberRoles(argv)
//	case "list-principals":
//		result, err = listPrincipals(argv)
//	case "list-roles":
//		result, err = listRoles(argv)
//	case "list-role-permissions":
//		result, err = listRolePermissions(argv)
//	case "remove-group-member":
//		err = removeGroupMember(argv)
//	case "remove-member":
//		err = removeMember(argv)
//	case "remove-group-role":
//		err = removeGroupRole(argv)
//	case "remove-role-permission":
//		err = removeRolePermission(argv)
//	case "validate-token":
//		result, err = validateToken(argv)
//	default:
//		fatal("unknown command: '%s'", arg)
//	}
//	return
//}

//// identity returns an identity service client with a tenant set by the environment/user
//func identity() *sdkIdentity.Service {
//	return apiClient().IdentityService
//}
//
//// identitySystem returns an identity service client using the "system" tenant for system-wide calls
//// that don't require a tenant
//func identitySystem() *sdkIdentity.Service {
//	return apiClientWithTenant("system").IdentityService
//}

func AddGroupMember(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string) (interface{}, error) {
	group, member :=utils.Head2(argv)
	return cmd.IdentityService.AddGroupMember(group, sdkIdentity.AddGroupMemberBody{Name: member})
}

func AddGroupRole(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	group, role :=utils.Head2(argv)
	return cmd.IdentityService.AddGroupRole(group, sdkIdentity.AddGroupRoleBody{Name: role})
}

func AddMember(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	member := utils.Head1(argv)
	return cmd.IdentityService.AddMember(sdkIdentity.AddMemberBody{Name: member})
}

func AddRolePermission(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	role, perm :=utils.Head2(argv)
	return cmd.IdentityService.AddRolePermission(role, perm)
}

func CreateGroup(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	group := utils.Head1(argv)
	return cmd.IdentityService.CreateGroup(sdkIdentity.CreateGroupBody{Name: group})
}

func CreateRole(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	role := utils.Head1(argv)
	return cmd.IdentityService.CreateRole(sdkIdentity.CreateRoleBody{Name: role})
}

func DeleteGroup(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)    (interface{}, error) {
	group := utils.Head1(argv)
	return nil, cmd.IdentityService.DeleteGroup(group)
}

func DeleteRole(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string) (interface{}, error) {
	role := utils.Head1(argv)
	return nil, cmd.IdentityService.DeleteRole(role)
}

func GetGroup(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	group := utils.Head1(argv)
	return cmd.IdentityService.GetGroup(group)
}

func GetGroupMember(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	group, member :=utils.Head2(argv)
	return cmd.IdentityService.GetGroupMember(group, member)
}

func GetGroupRole(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	group, role :=utils.Head2(argv)
	return cmd.IdentityService.GetGroupRole(group, role)
}

func GetMember(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	member := utils.Head1(argv)
	return cmd.IdentityService.GetMember(member)
}

func GetPrincipal(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	principal := utils.Head1(argv)
	return cmdSystem.IdentityService.GetPrincipal(principal)
}

func GetRole(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	role := utils.Head1(argv)
	return cmd.IdentityService.GetRole(role)
}

func GetRolePermission(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	role, perm :=utils.Head2(argv)
	return cmd.IdentityService.GetRolePermission(role, perm)
}

func ListGroups(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	return cmd.IdentityService.ListGroups()
}

func ListGroupMembers(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	group := utils.Head1(argv)
	return cmd.IdentityService.ListGroupMembers(group)
}

func ListGroupRoles(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	group := utils.Head1(argv)
	return cmd.IdentityService.ListGroupRoles(group)
}

func ListMembers(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	return cmd.IdentityService.ListMembers()
}

func ListMemberGroups(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	member := utils.Head1(argv)
	return cmd.IdentityService.ListMemberGroups(member)
}

func ListMemberPermissions(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	member := utils.Head1(argv)
	return cmd.IdentityService.ListMemberPermissions(member)
}

func ListMemberRoles(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	member := utils.Head1(argv)
	return cmd.IdentityService.ListMemberRoles(member)
}

func ListPrincipals(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	utils.CheckEmpty(argv)
	return cmdSystem.IdentityService.ListPrincipals()
}

func ListRoles(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	return cmd.IdentityService.ListRoles()
}

func ListRolePermissions(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	role := utils.Head1(argv)
	return cmd.IdentityService.ListRolePermissions(role)
}

func RemoveGroupMember(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string) (interface{}, error) {
	group, member :=utils.Head2(argv)
	return nil, cmd.IdentityService.RemoveGroupMember(group, member)
}

func RemoveMember(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string) (interface{}, error) {
	member := utils.Head1(argv)
	return nil, cmd.IdentityService.RemoveMember(member)
}

func RemoveGroupRole(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string) (interface{}, error) {
	group, role :=utils.Head2(argv)
	return nil, cmd.IdentityService.RemoveGroupRole(group, role)
}

func RemoveRolePermission(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string) (interface{}, error) {
	role, perm :=utils.Head2(argv)
	return nil, cmd.IdentityService.RemoveRolePermission(role, perm)
}

func ValidateToken(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
	tokenParams := sdkIdentity.ValidateTokenQueryParams{Include: []string{}}
	return cmd.IdentityService.ValidateToken(&tokenParams)
}

func  ListRoleGroups(cmd *sdk.Client,cmdSystem *sdk.Client,args []string) (interface{}, error) {
	fmt.Println("Not implemented yet")
	return nil, nil
}
//func (cmd *IdentityCommand) getSpecJSON(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
//	utils.checkEmpty(argv)
//	return GetSpecJSON("api", IdentityServiceVersion, "identity", cmd.identityService.Client)
//}
//
//func (cmd *IdentityCommand) getSpecYaml(cmd *sdk.Client, cmdSystem *sdk.Client, argv []string)  (interface{}, error) {
//	utils.checkEmpty(argv)
//	return GetSpecYaml("api", IdentityServiceVersion, "identity", cmd.identityService.Client)
//}
