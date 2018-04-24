package service

import (
	"testing"

	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
)

func TestIdentityService_GetUserProfile(t *testing.T) {
	user, err := getSplunkClient().IdentityService.GetUserProfile()
	assert.Nil(t, err)
	assert.Equal(t, "devtest@splunk.com", user.ID)
	assert.Equal(t, "devtest@splunk.com", user.Email)
	assert.Equal(t, "Dev", user.FirstName)
	assert.Equal(t, "Test", user.LastName)
	assert.Equal(t, "Dev Test", user.Name)
	assert.Equal(t, "en-US", user.Locale)
	assert.Equal(t, 1, len(user.TenantMemberships))
	assert.Equal(t, "devtestTenant", user.TenantMemberships[0])
}

func TestIdentityService_CreateTenant(t *testing.T) {
	err := getSplunkClient().IdentityService.CreateTenant(model.Tenant{TenantID: "devtestTenant"})
	assert.Nil(t, err)
}

func TestIdentityService_DeleteTenant(t *testing.T) {
	err := getSplunkClient().IdentityService.DeleteTenant("devtestTenant")
	assert.Nil(t, err)
}

func TestIdentityService_GetTenantUsers(t *testing.T) {
	users, err := getSplunkClient().IdentityService.GetTenantUsers("devtestTenant")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "devtest1@splunk.com", users[0].ID)
}

func TestIdentityService_ReplaceTenantUsers(t *testing.T) {
	err := getSplunkClient().IdentityService.ReplaceTenantUsers("devtestTenant", []model.User{
		{ID: "devtest2@splunk.com"},
		{ID: "devtest3@splunk.com"},
		{ID: "devtest4@splunk.com"},
		{ID: "devtest5@splunk.com"},
	})
	assert.Nil(t, err)
}

func TestIdentityService_AddTenantUsers(t *testing.T) {
	err := getSplunkClient().IdentityService.AddTenantUsers("devtestTenant", []model.User{
		{ID: "devtest7@splunk.com"},
		{ID: "devtest8@splunk.com"},
	})
	assert.Nil(t, err)
}

func TestIdentityService_DeleteTenantUsers(t *testing.T) {
	err := getSplunkClient().IdentityService.DeleteTenantUsers("devtestTenant", []model.User{
		{ID: "devtest5@splunk.com"},
		{ID: "devtest6@splunk.com"},
	})
	assert.Nil(t, err)
}
