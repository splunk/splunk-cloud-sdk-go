// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package identity

import (
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const servicePrefix = "identity"
const serviceVersion = "v1"
const serviceCluster = "api"

// Service talks to the IAC service
type Service services.BaseService

// NewService creates a new identity service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

// CreateTenant creates a tenant
func (s *Service) CreateTenant(name string) (*Tenant, error) {
	url, err := s.Client.BuildURLWithTenant("system", nil, serviceCluster, servicePrefix, serviceVersion,
		"tenants")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: map[string]string{"name": name}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result Tenant
	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetTenants returns the list of tenants in the system
func (s *Service) GetTenants() ([]string, error) {
	var result []string

	url, err := s.Client.BuildURLWithTenant("system", nil, serviceCluster, servicePrefix, serviceVersion,
		"tenants")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return result, err
}

// GetTenant returns  the tenant details
func (s *Service) GetTenant(name string) (*Tenant, error) {
	var result Tenant

	url, err := s.Client.BuildURLWithTenant("system", nil, serviceCluster, servicePrefix, serviceVersion,
		"tenants", name)

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return &result, err
}

// Validate validates the access token obtained from authorization header and returns the principal name and tenant memberships
func (s *Service) Validate() (*ValidateInfo, error) {
	var result ValidateInfo

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"validate")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return &result, err
}

// DeleteTenant deletes a tenant by name
func (s *Service) DeleteTenant(name string) error {
	url, err := s.Client.BuildURLWithTenant("system", nil, serviceCluster, servicePrefix, serviceVersion,
		"tenants", name)
	if err != nil {
		return err
	}

	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetPrincipals returns the list of principals known to IAC
func (s *Service) GetPrincipals() ([]string, error) {
	var result []string

	url, err := s.Client.BuildURLWithTenant("system", nil, serviceCluster, servicePrefix, serviceVersion,
		"principals")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return result, err
}

// GetPrincipal returns the principal details
func (s *Service) GetPrincipal(name string) (*Principal, error) {
	var result Principal

	url, err := s.Client.BuildURLWithTenant("system", nil, serviceCluster, servicePrefix, serviceVersion,
		"principals", name)

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return &result, err
}

// DeletePrincipal deletes a principal by name
func (s *Service) DeletePrincipal(name string) error {
	url, err := s.Client.BuildURLWithTenant("system", nil, serviceCluster, servicePrefix, serviceVersion,
		"tenants", name)
	if err != nil {
		return err
	}

	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetMembers returns the list of members in the given tenant
func (s *Service) GetMembers() ([]string, error) {
	var result []string

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"members")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return result, err
}

// GetMember gets a member of the given tenant
func (s *Service) GetMember(name string) (*Member, error) {
	var result Member

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"members", name)

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetMemberGroups returns the list of groups a member belongs to within a tenant
func (s *Service) GetMemberGroups(memberName string) ([]string, error) {
	var result []string

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"members", memberName, "groups")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return result, err
}

// GetMemberRoles returns the set of roles thet member posesses within the tenant
func (s *Service) GetMemberRoles(memberName string) ([]string, error) {
	var result []string

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"members", memberName, "roles")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return result, err
}

// GetMemberPermissions returns the set of permissions granted to the member within the tenant
func (s *Service) GetMemberPermissions(memberName string) ([]string, error) {
	var result []string

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"members", memberName, "permissions")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return result, err
}

// GetRoles get all roles for the given tenant
func (s *Service) GetRoles() ([]string, error) {
	var result []string

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"roles")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return result, err
}

// GetRole get a role for the given tenant
func (s *Service) GetRole(name string) (*Role, error) {
	var result Role

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"roles", name)

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetRolePermissions gets permissions for a role in this tenant
func (s *Service) GetRolePermissions(roleName string) ([]string, error) {
	var result []string

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"roles", roleName, "permissions")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return result, err
}

// GetRolePermission gets permissions for a role in this tenant
func (s *Service) GetRolePermission(roleName string, permissionName string) (*RolePermission, error) {
	var result RolePermission

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"roles", roleName, "permissions", permissionName)

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetGroups list groups that exist int he tenant
func (s *Service) GetGroups() ([]string, error) {
	var result []string

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"groups")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return result, err
}

// GetGroup gets a group in the given tenant
func (s *Service) GetGroup(name string) (*Group, error) {
	var result Group

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"groups", name)

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetGroupRoles lists the roles attached to the group
func (s *Service) GetGroupRoles(groupName string) ([]string, error) {
	var result []string

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"groups", groupName, "roles")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return result, err
}

// GetGroupRole returns group-role relationship details
func (s *Service) GetGroupRole(groupName string, roleName string) (*GroupRole, error) {
	var result GroupRole

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"groups", groupName, "roles", roleName)

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetGroupMembers lists the members attached to the group
func (s *Service) GetGroupMembers(groupName string) ([]string, error) {
	var result []string

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"groups", groupName, "members")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return result, err
}

// GetGroupMember returns group-member relationship details
func (s *Service) GetGroupMember(groupName string, memberName string) (*GroupMember, error) {
	var result GroupMember

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion,
		"groups", groupName, "members", memberName)

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreatePrincipal creates a new principal
func (s *Service) CreatePrincipal(name string, kind string) (*Principal, error) {
	url, err := s.Client.BuildURLWithTenant("system", nil, serviceCluster, servicePrefix, serviceVersion, "principals")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: map[string]string{"name": name, "kind": kind}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result Principal
	err = util.ParseResponse(&result, response)
	return &result, err
}

// AddMember adds a member to the given tenant
func (s *Service) AddMember(name string) (*Member, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "members")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: map[string]string{"name": name}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result Member
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateRole creates a new authorization role in the given tenant
func (s *Service) CreateRole(name string) (*Role, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "roles")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: map[string]string{"name": name}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result Role
	err = util.ParseResponse(&result, response)
	return &result, err
}

// AddPermissionToRole Adds permission to a role in this tenant
func (s *Service) AddPermissionToRole(roleName string, permissionName string) (*RolePermission, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "roles", roleName, "permissions")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: permissionName})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result RolePermission
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateGroup creates a new group in the given tenant
func (s *Service) CreateGroup(name string) (*Group, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "groups")
	if err != nil {
		return nil, err
	}

	response, err := s.Client.Post(services.RequestParams{URL: url, Body: map[string]string{"name": name}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result Group
	err = util.ParseResponse(&result, response)
	return &result, err
}

// AddRoleToGroup adds a role to the group
func (s *Service) AddRoleToGroup(groupName string, roleName string) (*GroupRole, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "groups", groupName, "roles")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: map[string]string{"name": roleName}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result GroupRole
	err = util.ParseResponse(&result, response)
	return &result, err
}

// AddMemberToGroup adds a member to the group
func (s *Service) AddMemberToGroup(groupName string, memberName string) (*GroupMember, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "groups", groupName, "members")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: map[string]string{"name": memberName}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result GroupMember
	err = util.ParseResponse(&result, response)
	return &result, err
}

// RemoveMember removes a member from the given tenant
func (s *Service) RemoveMember(name string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "members", name)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}

// DeleteRole deletes a defined role for the given tenant
func (s *Service) DeleteRole(name string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "roles", name)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}

// RemoveRolePermission removes a permission from the role
func (s *Service) RemoveRolePermission(roleName string, permissionName string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "roles", roleName, "permissions", permissionName)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}

// DeleteGroup deletes a group in the given tenant
func (s *Service) DeleteGroup(name string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "groups", name)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}

// RemoveGroupRole removes the role from the group
func (s *Service) RemoveGroupRole(groupName string, roleName string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "groups", groupName, "roles", roleName)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}

// RemoveGroupMember removes the member from the group
func (s *Service) RemoveGroupMember(groupName string, memberName string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "groups", groupName, "members", memberName)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}
