package stubbyintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIdentityService_GetUserProfile(t *testing.T) {
	user, err := getClient(t).IdentityService.GetUserProfile()
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
	err := getClient(t).IdentityService.CreateTenant(model.Tenant{TenantID: "devtestTenant"})
	assert.Nil(t, err)
}

func TestIdentityService_DeleteTenant(t *testing.T) {
	err := getClient(t).IdentityService.DeleteTenant("devtestTenant")
	assert.Nil(t, err)
}

func TestIdentityService_GetTenantUsers(t *testing.T) {
	users, err := getClient(t).IdentityService.GetTenantUsers("devtestTenant")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "devtest1@splunk.com", users[0].ID)
}

func TestIdentityService_ReplaceTenantUsers(t *testing.T) {
	err := getClient(t).IdentityService.ReplaceTenantUsers("devtestTenant", []model.User{
		{ID: "devtest2@splunk.com"},
		{ID: "devtest3@splunk.com"},
		{ID: "devtest4@splunk.com"},
		{ID: "devtest5@splunk.com"},
	})
	assert.Nil(t, err)
}

func TestIdentityService_AddTenantUsers(t *testing.T) {
	err := getClient(t).IdentityService.AddTenantUsers("devtestTenant", []model.User{
		{ID: "devtest7@splunk.com"},
		{ID: "devtest8@splunk.com"},
	})
	assert.Nil(t, err)
}

func TestIdentityService_DeleteTenantUsers(t *testing.T) {
	err := getClient(t).IdentityService.DeleteTenantUsers("devtestTenant", []model.User{
		{ID: "devtest4@splunk.com"},
		{ID: "devtest5@splunk.com"},
	})
	assert.Nil(t, err)
}
