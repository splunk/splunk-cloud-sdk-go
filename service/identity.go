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
		"members", name,"groups")

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
		"members", name,"roles")

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
		"members", name,"permissions")

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
		"roles",name)

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

