package integration

import (
	"net/http"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/provisioner"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestInvite tests invitation functions in provisioner
func TestInvite(t *testing.T) {
	config := &services.Config{
		Token:  testutils.TestAuthenticationToken,
		Host:   testutils.TestSplunkCloudHost,
		Tenant: testutils.TestTenant,
	}
	client, err := services.NewClient(config)
	require.Emptyf(t, err, "error calling services.NewClient(config): %s", err)

	provClient := provisioner.NewService(client)
	tenant := testutils.TestTenant

	t.Run("pre: clean up invites", func(t *testing.T) {
		invites, err := provClient.ListInvites()
		require.Emptyf(t, err, "error cleaning up invites with provisioner.ListInvites(): %s", err)

		for _, i := range *invites {
			err = provClient.DeleteInvite(i.InviteID)
			assert.Emptyf(t, err, "error cleaning up invites with provisioner.DeleteInvite(): %s", err)
		}
	})

	var invite *provisioner.InviteInfo
	var inviteID string

	t.Run("create invite", func(t *testing.T) {
		comment := "Go SDK test invite"
		email := "bounce@simulator.amazonses.com"
		var resp http.Response
		invite, err = provClient.CreateInvite(provisioner.InviteBody{
			Email:   email,
			Comment: &comment,
			Groups:  []string{"group1", "group2"},
		}, &resp)
		require.Emptyf(t, err, "error calling provisioner.CreateInvite(): %s", err)
		assert.Equal(t, tenant, invite.Tenant)
		assert.Equal(t, comment, invite.Comment)
		assert.Equal(t, email, invite.Email)
		assert.Equal(t, 2, len(invite.Groups))
	})

	time.Sleep(time.Second * 5)

	require.NotEmpty(t, invite, "invite test flow failed")

	t.Run("get invite", func(t *testing.T) {
		inviteID = invite.InviteID
		invite, err = provClient.GetInvite(inviteID)
		require.Emptyf(t, err, "error calling provisioner.GetInvite(): %s", err)
		assert.Equal(t, inviteID, invite.InviteID)
		assert.Equal(t, tenant, invite.Tenant)
	})

	t.Run("list invites", func(t *testing.T) {
		invites, err := provClient.ListInvites()
		require.Emptyf(t, err, "error calling provisioner.ListInvites(): %s", err)
		found := false
		for _, i := range *invites {
			if i.InviteID == inviteID {
				found = true
			}
		}
		assert.True(t, found)
	})

	t.Run("update invite", func(t *testing.T) {
		invite, err = provClient.UpdateInvite(inviteID, provisioner.UpdateInviteBody{
			Action: provisioner.UpdateInviteBodyActionResend,
		})
		assert.Equal(t, http.StatusLocked, err.(*util.HTTPError).HTTPStatusCode, "error calling provisioner.UpdateInvite(): %s", err)
	})

	t.Run("delete invite", func(t *testing.T) {
		err = provClient.DeleteInvite(inviteID)
		assert.Emptyf(t, err, "error calling provisioner.DeleteInvite(): %s", err)
	})
}

// TestProvisioner tests provisioner service-specific Splunk Cloud client
func TestProvisioner(t *testing.T) {
	config := &services.Config{
		Token:  testutils.TestAuthenticationToken,
		Host:   testutils.TestSplunkCloudHost,
		Tenant: "system",
	}
	client, err := services.NewClient(config)
	assert.Emptyf(t, err, "error calling services.NewClient(config): %s", err)

	provclient := provisioner.NewService(client)
	tenant := testutils.TestTenant

	t.Run("create provision job", func(t *testing.T) {
		bannedName := "splunk" // this is a banned word and should fail with a 403
		pJob, err := provclient.CreateProvisionJob(provisioner.CreateProvisionJobBody{
			Apps:   []string{},
			Tenant: &bannedName,
		})
		assert.Equal(t, http.StatusForbidden, err.(*util.HTTPError).HTTPStatusCode, "error calling provisioner.CreateProvisionJob(): %s", err)
		assert.Nil(t, pJob)
	})

	t.Run("get provision job", func(t *testing.T) {
		pJobInfo, err := provclient.GetProvisionJob("-1")
		assert.Equal(t, http.StatusNotFound, err.(*util.HTTPError).HTTPStatusCode, "error calling provisioner.GetProvisionJob(): %s", err)
		assert.Nil(t, pJobInfo)
	})

	t.Run("list provision jobs", func(t *testing.T) {
		pJobs, err := provclient.ListProvisionJobs()
		assert.Emptyf(t, err, "error calling provisioner.ListProvisionJobs(): %s", err)
		assert.NotNil(t, pJobs)
	})

	t.Run("get tenant", func(t *testing.T) {
		ten, err := provclient.GetTenant(tenant)
		assert.Emptyf(t, err, "error calling provisioner.GetTenant(): %s", err)
		assert.Equal(t, tenant, ten.Name)
	})

	t.Run("list tenants", func(t *testing.T) {
		tenants, err := provclient.ListTenants()
		assert.Emptyf(t, err, "error calling provisioner.ListTenants(): %s", err)
		assert.NotNil(t, tenants)
		found := false
		for _, i := range *tenants {
			if i.Name == tenant {
				found = true
			}
		}
		assert.True(t, found)
	})
}
