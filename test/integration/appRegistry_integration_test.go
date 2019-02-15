// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/splunk/splunk-cloud-sdk-go/services/appRegistry"
)

// Test GetActions which returns the list of all actions for the tenant
func TestIntegrationGetApps(t *testing.T) {
	client := getSdkClient(t)

	// Get Actions
	apps, err := client.AppRegistryService.ListApps()
	require.Nil(t, err)
	assert.True(t, len(apps) >= 0)
}

// Test Create/Get/Update/Delete app in app-registry service
func TestCRUDApp(t *testing.T) {
	client := getSdkClient(t)
	appName := "testname"

	// Create app
	app := appRegistry.CreateAppRequest{
		Kind:                 "web",
		Name:                 appName,
		Title:                "testtitle",
		RedirectUrls: []string{
			"https://localhost",
		},
	}
	_, err := client.AppRegistryService.CreateApp(&app)
	require.Nil(t, err)

	defer client.AppRegistryService.DeleteApp(appName)

	// List all apps
	apps, err := client.AppRegistryService.ListApps()
	require.Nil(t, err)
	require.Equal(t, 1, len(apps))
	//app-reg service bug https://jira.splunk.com/browse/APPLAT-5043
	// assert.EqualValues(t, apps[0], app)

	// Get the app
	app_ret, err := client.AppRegistryService.GetApp(appName)
	require.Nil(t, err)
	require.Equal(t, app.Name, app_ret.Name)
	require.Equal(t, app.Title, app_ret.Title)
	require.Equal(t, app.RedirectUrls, app_ret.RedirectUrls)
	require.Equal(t, app.Kind, app_ret.Kind)

	// Update the app. TODO: title and redirecturl should not needed once patch method is implemented.
	description:="new Description"
	title:= "new title"
	redirecturl:=[]string{"http://newlocalhost"}
	updateApp:=appRegistry.UpdateAppRequest{
		Description: &description,
		RedirectUrls:&redirecturl,
		Title:&title,
	}
	app_ret, err = client.AppRegistryService.UpdateApp(appName, &updateApp)
	require.Nil(t, err)
	require.Equal(t, app.Name, app_ret.Name)
	require.Equal(t, title, app_ret.Title)
	require.Equal(t, description, app_ret.Description)
	require.Equal(t, app.Kind, app_ret.Kind)


	// Delete the app
	err = client.AppRegistryService.DeleteApp(appName)
	require.Nil(t, err)

	apps, err = client.AppRegistryService.ListApps()
	require.Nil(t, err)
	assert.Equal(t, 0, len(apps))
}

// Test RotateSecret in app-registry service
func TestAppRotateSecret(t *testing.T) {
	client := getSdkClient(t)
	appName := "testrotatesecret"

	// Create app
	app := appRegistry.CreateAppRequest{
		Kind:                 "web",
		Name:                 appName,
		Title:                "testtitle",
		RedirectUrls: []string{
			"https://localhost",
		},
	}
	app_created, err := client.AppRegistryService.CreateApp(&app)
	require.Nil(t, err)

	defer client.AppRegistryService.DeleteApp(appName)

	// rotate secret
	app_ret, err := client.AppRegistryService.RotateSecret(appName)
	require.Nil(t, err)
	require.NotEmpty(t, app_created.ClientSecret, app_ret.ClientSecret)

	// get subscription
	subs, err := client.AppRegistryService.GetAppSubscriptions(appName)
	require.Nil(t, err)
	require.NotEmpty(t, 0, len(subs));
}

