// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/splunk/splunk-cloud-sdk-go/services/appreg"
)

// Test GetActions which returns the list of all actions for the tenant
func TestIntegrationGetApps(t *testing.T) {
	client := getSdkClient(t)

	// Get Actions
	appregs, err := client.AppRegService.ListApps()
	require.Nil(t, err)
	assert.True(t, len(appregs) >= 0)
}

// Test Create/Get/Delete app in app-registry service
func TestCRUDApp(t *testing.T) {
	client := getSdkClient(t)
	appName := "testname"
	app := appreg.CreateAppRequest{
		Kind:                 "web",
		Name:                 appName,
		Title:                "testtitle",
		RedirectUrls: []string{
			"https://localhost",
		},
	}
	_, err := client.AppRegService.CreateApp(&app)
	require.Nil(t, err)

	defer client.AppRegService.DeleteApp(appName)
	apps, err := client.AppRegService.ListApps()
	require.Nil(t, err)
	require.Equal(t, 1, len(apps))
	//app-reg service bug https://jira.splunk.com/browse/APPLAT-5043
	// assert.EqualValues(t, apps[0], app)

	app_ret, err := client.AppRegService.GetApp(appName)
	require.Nil(t, err)
	require.Equal(t, app.Name, app_ret.Name)
	require.Equal(t, app.Title, app_ret.Title)
	require.Equal(t, app.RedirectUrls, app_ret.RedirectUrls)
	require.Equal(t, app.Kind, app_ret.Kind)

	err = client.AppRegService.DeleteApp(appName)
	require.Nil(t, err)

	apps, err = client.AppRegService.ListApps()
	require.Nil(t, err)
	assert.Equal(t, 0, len(apps))
}
