/*
 * Copyright © 2019 Splunk, Inc.
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
 *
 * Identity and Access Control
 *
 * With the Splunk Cloud Identity and Access Control (IAC) Service, you can authenticate and authorize Splunk API users.
 *
 * API version: v2beta1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package identity

import (
	"net/http"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const serviceCluster = "api"

type Service services.BaseService

// NewService creates a new identity service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

/*
	AddGroupMember - identity service endpoint
	Adds a member to a given group.
	Parameters:
		group: The group name.
		addGroupMemberBody: The member to add to a group.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) AddGroupMember(group string, addGroupMemberBody AddGroupMemberBody, resp ...*http.Response) (*GroupMember, error) {
	pp := struct {
		Group string
	}{
		Group: group,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups/{{.Group}}/members`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: addGroupMemberBody})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb GroupMember
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	AddGroupRole - identity service endpoint
	Adds a role to a given group.
	Parameters:
		group: The group name.
		addGroupRoleBody: The role to add to a group.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) AddGroupRole(group string, addGroupRoleBody AddGroupRoleBody, resp ...*http.Response) (*GroupRole, error) {
	pp := struct {
		Group string
	}{
		Group: group,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups/{{.Group}}/roles`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: addGroupRoleBody})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb GroupRole
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	AddMember - identity service endpoint
	Adds a member to a given tenant.
	Parameters:
		addMemberBody: The member to associate with a tenant.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) AddMember(addMemberBody AddMemberBody, resp ...*http.Response) (*Member, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/members`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: addMemberBody})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb Member
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	AddRolePermission - identity service endpoint
	Adds permissions to a role in a given tenant.
	Parameters:
		role: The role name.
		body: The permission to add to a role.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) AddRolePermission(role string, body string, resp ...*http.Response) (*RolePermission, error) {
	pp := struct {
		Role string
	}{
		Role: role,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/roles/{{.Role}}/permissions`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: body})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb RolePermission
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreateGroup - identity service endpoint
	Creates a new group in a given tenant.
	Parameters:
		createGroupBody: The group definition.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateGroup(createGroupBody CreateGroupBody, resp ...*http.Response) (*Group, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: createGroupBody})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb Group
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreateRole - identity service endpoint
	Creates a new authorization role in a given tenant.
	Parameters:
		createRoleBody: Role definition
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateRole(createRoleBody CreateRoleBody, resp ...*http.Response) (*Role, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/roles`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: createRoleBody})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb Role
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	DeleteGroup - identity service endpoint
	Deletes a group in a given tenant.
	Parameters:
		group: The group name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteGroup(group string, resp ...*http.Response) error {
	pp := struct {
		Group string
	}{
		Group: group,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups/{{.Group}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	DeleteRole - identity service endpoint
	Deletes a defined role for a given tenant.
	Parameters:
		role: The role name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteRole(role string, resp ...*http.Response) error {
	pp := struct {
		Role string
	}{
		Role: role,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/roles/{{.Role}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	GetGroup - identity service endpoint
	Returns information about a given group within a tenant.
	Parameters:
		group: The group name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetGroup(group string, resp ...*http.Response) (*Group, error) {
	pp := struct {
		Group string
	}{
		Group: group,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups/{{.Group}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb Group
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetGroupMember - identity service endpoint
	Returns information about a given member within a given group.
	Parameters:
		group: The group name.
		member: The member name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetGroupMember(group string, member string, resp ...*http.Response) (*GroupMember, error) {
	pp := struct {
		Group  string
		Member string
	}{
		Group:  group,
		Member: member,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups/{{.Group}}/members/{{.Member}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb GroupMember
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetGroupRole - identity service endpoint
	Returns information about a given role within a given group.
	Parameters:
		group: The group name.
		role: The role name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetGroupRole(group string, role string, resp ...*http.Response) (*GroupRole, error) {
	pp := struct {
		Group string
		Role  string
	}{
		Group: group,
		Role:  role,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups/{{.Group}}/roles/{{.Role}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb GroupRole
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetMember - identity service endpoint
	Returns a member of a given tenant.
	Parameters:
		member: The member name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetMember(member string, resp ...*http.Response) (*Member, error) {
	pp := struct {
		Member string
	}{
		Member: member,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/members/{{.Member}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb Member
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetPrincipal - identity service endpoint
	Returns the details of a principal, including its tenant membership.
	Parameters:
		principal: The principal name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetPrincipal(principal string, resp ...*http.Response) (*Principal, error) {
	pp := struct {
		Principal string
	}{
		Principal: principal,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/system/identity/v2beta1/principals/{{.Principal}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb Principal
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetRole - identity service endpoint
	Returns a role for a given tenant.
	Parameters:
		role: The role name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetRole(role string, resp ...*http.Response) (*Role, error) {
	pp := struct {
		Role string
	}{
		Role: role,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/roles/{{.Role}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb Role
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetRolePermission - identity service endpoint
	Gets a permission for the specified role.
	Parameters:
		role: The role name.
		permission: The permission string.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetRolePermission(role string, permission string, resp ...*http.Response) (*RolePermission, error) {
	pp := struct {
		Role       string
		Permission string
	}{
		Role:       role,
		Permission: permission,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/roles/{{.Role}}/permissions/{{.Permission}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb RolePermission
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListGroupMembers - identity service endpoint
	Returns a list of the members within a given group.
	Parameters:
		group: The group name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListGroupMembers(group string, resp ...*http.Response) ([]string, error) {
	pp := struct {
		Group string
	}{
		Group: group,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups/{{.Group}}/members`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListGroupRoles - identity service endpoint
	Returns a list of the roles that are attached to a group within a given tenant.
	Parameters:
		group: The group name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListGroupRoles(group string, resp ...*http.Response) ([]string, error) {
	pp := struct {
		Group string
	}{
		Group: group,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups/{{.Group}}/roles`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListGroups - identity service endpoint
	List the groups that exist in a given tenant.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListGroups(resp ...*http.Response) ([]string, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListMemberGroups - identity service endpoint
	Returns a list of groups that a member belongs to within a tenant.
	Parameters:
		member: The member name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListMemberGroups(member string, resp ...*http.Response) ([]string, error) {
	pp := struct {
		Member string
	}{
		Member: member,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/members/{{.Member}}/groups`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListMemberPermissions - identity service endpoint
	Returns a set of permissions granted to the member within the tenant.
	Parameters:
		member: The member name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListMemberPermissions(member string, resp ...*http.Response) ([]string, error) {
	pp := struct {
		Member string
	}{
		Member: member,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/members/{{.Member}}/permissions`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListMemberRoles - identity service endpoint
	Returns a set of roles that a given member holds within the tenant.
	Parameters:
		member: The member name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListMemberRoles(member string, resp ...*http.Response) ([]string, error) {
	pp := struct {
		Member string
	}{
		Member: member,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/members/{{.Member}}/roles`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListMembers - identity service endpoint
	Returns a list of members in a given tenant.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListMembers(resp ...*http.Response) ([]string, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/members`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListPrincipals - identity service endpoint
	Returns the list of principals known to IAC.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListPrincipals(resp ...*http.Response) ([]string, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/system/identity/v2beta1/principals`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListRoleGroups - identity service endpoint
	Gets a list of groups for a role in a given tenant.
	Parameters:
		role: The role name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListRoleGroups(role string, resp ...*http.Response) ([]string, error) {
	pp := struct {
		Role string
	}{
		Role: role,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/roles/{{.Role}}/groups`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListRolePermissions - identity service endpoint
	Gets the permissions for a role in a given tenant.
	Parameters:
		role: The role name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListRolePermissions(role string, resp ...*http.Response) ([]string, error) {
	pp := struct {
		Role string
	}{
		Role: role,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/roles/{{.Role}}/permissions`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListRoles - identity service endpoint
	Returns all roles for a given tenant.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListRoles(resp ...*http.Response) ([]string, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/roles`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []string
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	RemoveGroupMember - identity service endpoint
	Removes the member from a given group.
	Parameters:
		group: The group name.
		member: The member name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) RemoveGroupMember(group string, member string, resp ...*http.Response) error {
	pp := struct {
		Group  string
		Member string
	}{
		Group:  group,
		Member: member,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups/{{.Group}}/members/{{.Member}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	RemoveGroupRole - identity service endpoint
	Removes a role from a given group.
	Parameters:
		group: The group name.
		role: The role name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) RemoveGroupRole(group string, role string, resp ...*http.Response) error {
	pp := struct {
		Group string
		Role  string
	}{
		Group: group,
		Role:  role,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/groups/{{.Group}}/roles/{{.Role}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	RemoveMember - identity service endpoint
	Removes a member from a given tenant
	Parameters:
		member: The member name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) RemoveMember(member string, resp ...*http.Response) error {
	pp := struct {
		Member string
	}{
		Member: member,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/members/{{.Member}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	RemoveRolePermission - identity service endpoint
	Removes a permission from the role.
	Parameters:
		role: The role name.
		permission: The permission string.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) RemoveRolePermission(role string, permission string, resp ...*http.Response) error {
	pp := struct {
		Role       string
		Permission string
	}{
		Role:       role,
		Permission: permission,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/identity/v2beta1/roles/{{.Role}}/permissions/{{.Permission}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	ValidateToken - identity service endpoint
	Validates the access token obtained from the authorization header and returns the principal name and tenant memberships.
	Parameters:
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ValidateToken(query *ValidateTokenQueryParams, resp ...*http.Response) (*ValidateInfo, error) {
	values := util.ParseURLParams(query)
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/identity/v2beta1/validate`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb ValidateInfo
	err = util.ParseResponse(&rb, response)
	return &rb, err
}