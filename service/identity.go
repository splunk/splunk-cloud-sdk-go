// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

const identityServicePrefix = "identity"
const identityServiceVersion = "v1"

// IdentityService talks to the IAC service
type IdentityService service

// CreateTenant creates a tenant
func (c *IdentityService) CreateTenant(name string) (*model.Tenant, error) {
	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"tenants")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: map[string]string{"name": name}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result model.Tenant
	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetTenants returns the list of tenants in the system
func (c *IdentityService) GetTenants() ([]string, error) {
	var result []string

	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"tenants")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetTenant(name string) (*model.Tenant, error) {
	var result model.Tenant

	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"tenants", name)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) Validate() (*model.ValidateInfo, error) {
	var result model.ValidateInfo

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"validate")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) DeleteTenant(name string) error {
	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"tenants", name)
	if err != nil {
		return err
	}

	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetPrincipals returns the list of principals known to IAC
func (c *IdentityService) GetPrincipals() ([]string, error) {
	var result []string

	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"principals")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetPrincipal(name string) (*model.Principal, error) {
	var result model.Principal

	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"principals", name)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) DeletePrincipal(name string) error {
	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"tenants", name)
	if err != nil {
		return err
	}

	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetMembers returns the list of members in the given tenant
func (c *IdentityService) GetMembers() ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"members")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetMember(name string) (*model.Member, error) {
	var result model.Member

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"members", name)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetMemberGroups(memberName string) ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"members", memberName, "groups")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetMemberRoles(memberName string) ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"members", memberName, "roles")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetMemberPermissions(memberName string) ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"members", memberName, "permissions")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetRoles() ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"roles")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetRole(name string) (*model.Role, error) {
	var result model.Role

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"roles", name)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetRolePermissions(roleName string) ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"roles", roleName, "permissions")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetRolePermission(roleName string, permissionName string) (*model.RolePermission, error) {
	var result model.RolePermission

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"roles", roleName, "permissions", permissionName)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetGroups() ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"groups")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetGroup(name string) (*model.Group, error) {
	var result model.Group

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"groups", name)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetGroupRoles(groupName string) ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"groups", groupName, "roles")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetGroupRole(groupName string, roleName string) (*model.GroupRole, error) {
	var result model.GroupRole

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"groups", groupName, "roles", roleName)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetGroupMembers(groupName string) ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"groups", groupName, "members")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) GetGroupMember(groupName string, memberName string) (*model.GroupMember, error) {
	var result model.GroupMember

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"groups", groupName, "members", memberName)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *IdentityService) CreatePrincipal(name string, kind string) (*model.Principal, error) {
	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion, "principals")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: map[string]string{"name": name, "kind": kind}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result model.Principal
	err = util.ParseResponse(&result, response)
	return &result, err
}

// AddMember adds a member to the given tenant
func (c *IdentityService) AddMember(name string) (*model.Member, error) {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "members")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: map[string]string{"name": name}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result model.Member
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateRole creates a new authorization role in the given tenant
func (c *IdentityService) CreateRole(name string) (*model.Role, error) {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "roles")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: map[string]string{"name": name}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result model.Role
	err = util.ParseResponse(&result, response)
	return &result, err
}

// AddPermissionToRole Adds permission to a role in this tenant
func (c *IdentityService) AddPermissionToRole(roleName string, permissionName string) (*model.RolePermission, error) {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "roles", roleName, "permissions")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: permissionName})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result model.RolePermission
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateGroup creates a new group in the given tenant
func (c *IdentityService) CreateGroup(name string) (*model.Group, error) {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "groups")
	if err != nil {
		return nil, err
	}

	response, err := c.client.Post(RequestParams{URL: url, Body: map[string]string{"name": name}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result model.Group
	err = util.ParseResponse(&result, response)
	return &result, err
}

// AddRoleToGroup adds a role to the group
func (c *IdentityService) AddRoleToGroup(groupName string, roleName string) (*model.GroupRole, error) {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "groups", groupName, "roles")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: map[string]string{"name": roleName}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result model.GroupRole
	err = util.ParseResponse(&result, response)
	return &result, err
}

// AddMemberToGroup adds a member to the group
func (c *IdentityService) AddMemberToGroup(groupName string, memberName string) (*model.GroupMember, error) {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "groups", groupName, "members")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: map[string]string{"name": memberName}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result model.GroupMember
	err = util.ParseResponse(&result, response)
	return &result, err
}

// RemoveMember removes a member from the given tenant
func (c *IdentityService) RemoveMember(name string) error {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "members", name)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}

// DeleteRole deletes a defined role for the given tenant
func (c *IdentityService) DeleteRole(name string) error {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "roles", name)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}

// RemoveRolePermission removes a permission from the role
func (c *IdentityService) RemoveRolePermission(roleName string, permissionName string) error {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "roles", roleName, "permissions", permissionName)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}

// DeleteGroup deletes a group in the given tenant
func (c *IdentityService) DeleteGroup(name string) error {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "groups", name)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}

// RemoveGroupRole removes the role from the group
func (c *IdentityService) RemoveGroupRole(groupName string, roleName string) error {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "groups", groupName, "roles", roleName)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}

// RemoveGroupMember removes the member from the group
func (c *IdentityService) RemoveGroupMember(groupName string, memberName string) error {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "groups", groupName, "members", memberName)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}
