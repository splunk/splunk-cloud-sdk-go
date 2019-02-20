// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/services/appregistry"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test Create/Get/Update/Delete app in app-registry service
func TestCRUDApp(t *testing.T) {
	client := getSdkClient(t)
	appName := "gotestapp"

	// Create app
	app := appregistry.CreateAppRequest{
		Kind:  appregistry.WEB,
		Name:  appName,
		Title: "testtitle",
		RedirectURLs: []string{
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
	require.Equal(t, app.RedirectURLs, app_ret.RedirectURLs)
	require.Equal(t, app.Kind, app_ret.Kind)

	// Update the app. TODO: title and redirecturl should not needed once patch method is implemented.
	description := "new Description"
	title := "new title"
	redirecturl := []string{"http://newlocalhost"}
	updateApp := appregistry.UpdateAppRequest{
		Description:  &description,
		RedirectURLs: &redirecturl,
		Title:        &title,
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
	appName := "gotestrotatesecret"

	// Create app
	app := appregistry.CreateAppRequest{
		Kind:  appregistry.WEB,
		Name:  appName,
		Title: "testtitle",
		RedirectURLs: []string{
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
}

// Test Create/Get/List/Delete subscriptions and get apps/subscriptions in app-registry service
func TestSubscriptions(t *testing.T) {
	client := getSdkClient(t)
	appName := "gotestsubscriptions"

	// Create app
	app := appregistry.CreateAppRequest{
		Kind:  appregistry.NATIVE,
		Name:  appName,
		Title: "testtitle",
		RedirectURLs: []string{
			"https://localhost",
		},
	}
	_, err := client.AppRegistryService.CreateApp(&app)
	require.Nil(t, err)
	defer client.AppRegistryService.DeleteApp(appName)

	// create subscription
	err = client.AppRegistryService.CreateSubscription(appName)
	require.Nil(t, err)
	defer client.AppRegistryService.DeleteSubscription(appName)

	// get app subscription
	appsubs, err := client.AppRegistryService.GetAppSubscriptions(appName)
	require.Nil(t, err)
	require.Equal(t, 1, len(appsubs))

	// Get a subscription of an app
	subs, err := client.AppRegistryService.GetSubscription(appName)
	require.Nil(t, err)
	require.NotEmpty(t, subs)

	// Get a subscription from non-exist-app
	_, err = client.AppRegistryService.GetSubscription("notExistApp")
	require.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	require.Equal(t, 404, httpErr.HTTPStatusCode)

	// List all subscriptions
	// create the 2nd subscription
	appName2 := "gotestsubscriptions_2"
	app2 := appregistry.CreateAppRequest{
		Kind:  appregistry.SERVICE,
		Name:  appName2,
		Title: "testtitle2",
		RedirectURLs: []string{
			"https://localhost",
		},
	}
	_, err = client.AppRegistryService.CreateApp(&app2)
	require.Nil(t, err)
	defer client.AppRegistryService.DeleteApp(appName2)
	err = client.AppRegistryService.CreateSubscription(appName2)
	require.Nil(t, err)
	defer client.AppRegistryService.DeleteSubscription(appName2)

	all_subs, err := client.AppRegistryService.ListSubscriptions()
	require.Nil(t, err)
	require.Equal(t, 2, len(all_subs))

	// Delete the subscription
	err = client.AppRegistryService.DeleteSubscription(appName)
	require.Nil(t, err)
}
