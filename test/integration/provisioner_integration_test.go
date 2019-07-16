package integration

import (
	"net/http"
	"testing"

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
	require.Emptyf(t, err, "error calling services.NewService(): %s", err)

	bannedName := "splunk" // this is a banned word and should fail with a 422
	pJob, err := client.CreateProvisionJob(provisioner.CreateProvisionJobBody{
		Apps:   []string{},
		Tenant: &bannedName,
	})
	assert.Equal(t, http.StatusUnprocessableEntity, err.(*util.HTTPError).HTTPStatusCode, "error calling provisioner.CreateProvisionJob(): %s", err)
	assert.Nil(t, pJob)

	pJobInfo, err := client.GetProvisionJob("-1")
	assert.Equal(t, http.StatusNotFound, err.(*util.HTTPError).HTTPStatusCode, "error calling provisioner.GetProvisionJob(): %s", err)
	assert.Nil(t, pJobInfo)

	pJobs, err := client.ListProvisionJobs()
	assert.Emptyf(t, err, "error calling provisioner.ListProvisionJobs(): %s", err)
	assert.NotNil(t, pJobs)

	tenant := testutils.TestProvisionerTenant
	ten, err := client.GetTenant(tenant)
	assert.Emptyf(t, err, "error calling provisioner.GetTenant(): %s", err)
	assert.Equal(t, tenant, ten.Name)

	tenants, err := client.ListTenants()
	assert.Emptyf(t, err, "error calling provisioner.ListTenants(): %s", err)
	assert.NotNil(t, tenants)
	found := false
	for _, i := range *tenants {
		if i.Name == tenant {
			found = true
		}
	}
	assert.True(t, found)
}
