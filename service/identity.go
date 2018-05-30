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
	url, err := c.client.BuildURLWithTenantID("system", identityServicePrefix, identityServiceVersion,
		"tenants")
	if err != nil {
		return err
	}

	response, err := c.client.Post(url, tenant)
	return util.ParseError(response, err)
}

// DeleteTenant deletes a tenant by tenantID
func (c *IdentityService) DeleteTenant(tenantID string) error {
	url, err := c.client.BuildURLWithTenantID("system", identityServicePrefix, identityServiceVersion,
		"tenants", tenantID)
	if err != nil {
		return err
	}

	response, err := c.client.Delete(url)
	return util.ParseError(response, err)
}

// GetUserProfile retrieves the user profile associated with the current cached auth token
func (c *IdentityService) GetUserProfile() (*model.User, error) {
	var user model.User
	url, err := c.client.BuildURLWithTenantID("system", identityServicePrefix, identityServiceVersion,
		"userprofile")
	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(url)
	err = util.ParseResponse(&user, response, err)
	return &user, err
}

// GetTenantUsers returns users registered with the tenant
func (c *IdentityService) GetTenantUsers(tenantID string) ([]model.User, error) {
	var users []model.User
	url, err := c.client.BuildURLWithTenantID("system", identityServicePrefix, identityServiceVersion,
		"tenants", tenantID, "users")
	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(url)
	err = util.ParseResponse(&users, response, err)
	return users, err
}

// ReplaceTenantUsers replaces existing users on a tenant with the provided user list
func (c *IdentityService) ReplaceTenantUsers(tenantID string, users []model.User) error {
	url, err := c.client.BuildURLWithTenantID("system", identityServicePrefix, identityServiceVersion,
		"tenants", tenantID, "users")
	if err != nil {
		return err
	}

	response, err := c.client.Put(url, users)
	return util.ParseError(response, err)
}

// AddTenantUsers adds users to a tenant
func (c *IdentityService) AddTenantUsers(tenantID string, users []model.User) error {
	url, err := c.client.BuildURLWithTenantID("system", identityServicePrefix, identityServiceVersion,
		"tenants", tenantID, "users")
	if err != nil {
		return err
	}

	response, err := c.client.Patch(url, users)
	return util.ParseError(response, err)
}

// DeleteTenantUsers deletes users from a tenant
func (c *IdentityService) DeleteTenantUsers(tenantID string, users []model.User) error {
	url, err := c.client.BuildURLWithTenantID("system", identityServicePrefix, identityServiceVersion,
		"tenants", tenantID, "users")
	if err != nil {
		return err
	}

	response, err := c.client.DeleteWithBody(url, users)
	return util.ParseError(response, err)
}
