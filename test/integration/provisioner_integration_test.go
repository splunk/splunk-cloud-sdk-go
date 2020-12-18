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
