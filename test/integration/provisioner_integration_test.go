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

// TestProvisioner tests provisioner service-specific Splunk Cloud client
func TestProvisioner(t *testing.T) {
	client, err := provisioner.NewService(&services.Config{
		Token:  testutils.TestAuthenticationToken,
		Host:   testutils.TestSplunkCloudHost,
		Tenant: "system",
	})
	provClient, err := provisioner.NewService(&services.Config{
		Token:  testutils.TestAuthenticationToken,
		Host:   testutils.TestSplunkCloudHost,
		Tenant: testutils.TestTenant,
	})

	require.Emptyf(t, err, "error calling services.NewService(): %s", err)

	tenant := testutils.TestTenant

	comment := "Go SDK test invite"
	email := "bounce@simulator.amazonses.com"
	var resp http.Response
	invite, err := provClient.CreateInvite(provisioner.InviteBody{
		Email:   email,
		Comment: &comment,
		Groups:  []string{"group1", "group2"},
	}, &resp)
	require.Emptyf(t, err, "error calling provisioner.CreateInvite(): %s", err)
	assert.Equal(t, tenant, invite.Tenant)
	assert.Equal(t, comment, invite.Comment)
	assert.Equal(t, email, invite.Email)
	assert.Equal(t, 2, len(invite.Groups))

	time.Sleep(time.Second * 5)

	inviteID := invite.InviteID
	invite, err = provClient.GetInvite(inviteID)
	require.Emptyf(t, err, "error calling provisioner.GetInvite(): %s", err)
	assert.Equal(t, inviteID, invite.InviteID)
	assert.Equal(t, tenant, invite.Tenant)

	invites, err := provClient.ListInvites()
	require.Emptyf(t, err, "error calling provisioner.ListInvites(): %s", err)
	found := false
	for _, i := range *invites {
		if i.InviteID == inviteID {
			found = true
		}
	}
	assert.True(t, found)

	invite, err = provClient.UpdateInvite(inviteID, provisioner.UpdateInviteBody{
		Action: provisioner.UpdateInviteBodyActionResend,
	})
	assert.Equal(t, http.StatusLocked, err.(*util.HTTPError).HTTPStatusCode, "error calling provisioner.UpdateInvite(): %s", err)

	err = provClient.DeleteInvite(inviteID)
	assert.Emptyf(t, err, "error calling provisioner.DeleteInvite(): %s", err)

	bannedName := "splunk" // this is a banned word and should fail with a 403
	pJob, err := client.CreateProvisionJob(provisioner.CreateProvisionJobBody{
		Apps:   []string{},
		Tenant: &bannedName,
	})
	assert.Equal(t, http.StatusForbidden, err.(*util.HTTPError).HTTPStatusCode, "error calling provisioner.CreateProvisionJob(): %s", err)
	assert.Nil(t, pJob)

	pJobInfo, err := client.GetProvisionJob("-1")
	assert.Equal(t, http.StatusNotFound, err.(*util.HTTPError).HTTPStatusCode, "error calling provisioner.GetProvisionJob(): %s", err)
	assert.Nil(t, pJobInfo)

	pJobs, err := client.ListProvisionJobs()
	assert.Emptyf(t, err, "error calling provisioner.ListProvisionJobs(): %s", err)
	assert.NotNil(t, pJobs)

	ten, err := client.GetTenant(tenant)
	assert.Emptyf(t, err, "error calling provisioner.GetTenant(): %s", err)
	assert.Equal(t, tenant, ten.Name)

	tenants, err := client.ListTenants()
	assert.Emptyf(t, err, "error calling provisioner.ListTenants(): %s", err)
	assert.NotNil(t, tenants)
	found = false
	for _, i := range *tenants {
		if i.Name == tenant {
			found = true
		}
	}
	assert.True(t, found)
}
