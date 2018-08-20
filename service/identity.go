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
func (c *IdentityService) CreateTenant(tenant model.Tenant) error {
	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"tenants")
	if err != nil {
		return err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: tenant})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// DeleteTenant deletes a tenant by tenantID
func (c *IdentityService) DeleteTenant(tenantID string) error {
	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"tenants", tenantID)
	if err != nil {
		return err
	}

	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetUserProfile retrieves the user profile associated with the current cached auth token
func (c *IdentityService) GetUserProfile(tenantID string) (*model.User, error) {
	var user model.User
	var scope = tenantID

	if len(tenantID) == 0 {
		scope = "system"
	}
	url, err := c.client.BuildURLWithTenantID(scope, nil, identityServicePrefix, identityServiceVersion,
		"userprofile")

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
	err = util.ParseResponse(&user, response)
	return &user, err
}

// GetTenantUsers returns users registered with the tenant
func (c *IdentityService) GetTenantUsers(tenantID string) ([]model.User, error) {
	var users []model.User
	url, err := c.client.BuildURLWithTenantID(tenantID, nil, identityServicePrefix, identityServiceVersion,
		"users")
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
	err = util.ParseResponse(&users, response)
	return users, err
}

// ReplaceTenantUsers replaces existing users on a tenant with the provided user list
func (c *IdentityService) ReplaceTenantUsers(tenantID string, users []model.User) error {
	url, err := c.client.BuildURLWithTenantID(tenantID, nil, identityServicePrefix, identityServiceVersion,
		"users")
	if err != nil {
		return err
	}

	response, err := c.client.Put(RequestParams{URL: url, Body: users})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// AddTenantUsers adds users to a tenant
func (c *IdentityService) AddTenantUsers(tenantID string, users []model.User) error {
	url, err := c.client.BuildURLWithTenantID(tenantID, nil, identityServicePrefix, identityServiceVersion,
		"users")
	if err != nil {
		return err
	}

	response, err := c.client.Patch(RequestParams{URL: url, Body: users})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// DeleteTenantUsers deletes users from a tenant
func (c *IdentityService) DeleteTenantUsers(tenantID string, users []model.User) error {
	url, err := c.client.BuildURLWithTenantID(tenantID, nil, identityServicePrefix, identityServiceVersion,
		"users")
	if err != nil {
		return err
	}

	response, err := c.client.Delete(RequestParams{URL: url, Body: users})
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
func (c *IdentityService) GetMemberGroups(name string) ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"members", name, "groups")

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
func (c *IdentityService) GetMemberRoles(name string) ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"members", name, "roles")

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
func (c *IdentityService) GetMemberPermissions(name string) ([]string, error) {
	var result []string

	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion,
		"members", name, "permissions")

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

// AddMember adds a member to the given tenant
func (c *IdentityService) AddMember(memberName string) (*model.Member, error) {
	url, err := c.client.BuildURL(nil, identityServicePrefix, identityServiceVersion, "members")
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

// DeleteMember removes a member fromt the given tenant
func (c *IdentityService) DeleteMember(name string) error {
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

// RemoveRolePermission Removes a permission from the role
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

// RemoveGroupMember removes the memeber from the group
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
