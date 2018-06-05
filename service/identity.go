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
	response, err := c.client.Post(url, tenant)
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

	response, err := c.client.Delete(url)
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetUserProfile retrieves the user profile associated with the current cached auth token
func (c *IdentityService) GetUserProfile() (*model.User, error) {
	var user model.User
	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"userprofile")
	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(url)
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
	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"tenants", tenantID, "users")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(url)
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
	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"tenants", tenantID, "users")
	if err != nil {
		return err
	}

	response, err := c.client.Put(url, users)
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// AddTenantUsers adds users to a tenant
func (c *IdentityService) AddTenantUsers(tenantID string, users []model.User) error {
	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"tenants", tenantID, "users")
	if err != nil {
		return err
	}

	response, err := c.client.Patch(url, users)
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// DeleteTenantUsers deletes users from a tenant
func (c *IdentityService) DeleteTenantUsers(tenantID string, users []model.User) error {
	url, err := c.client.BuildURLWithTenantID("system", nil, identityServicePrefix, identityServiceVersion,
		"tenants", tenantID, "users")
	if err != nil {
		return err
	}

	response, err := c.client.DeleteWithBody(url, users)
	if response != nil {
		defer response.Body.Close()
	}
	return err
}
