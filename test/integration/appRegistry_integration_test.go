/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package integration

import (
	"fmt"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/services/appregistry"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newAppTitle creates standardized app title for tests
func newAppTitle(title string) string {
	ret := fmt.Sprintf("gsdk-%s-%d", title, testutils.RunSuffix)
	return ret
}

// Test Create/Get/Update/Delete app in app-registry service
func TestCRUDApp(t *testing.T) {
	client := getSdkClient(t)
	appName := fmt.Sprintf("g.c%d", testutils.RunSuffix)

	// Create app

	app := appregistry.MakeCreateAppRequestFromWebAppPost(appregistry.WebAppPost{
		Kind:  appregistry.AppResourceKindWeb,
		Name:  appName,
		Title: newAppTitle("testtitle"),

		RedirectUrls: []string{
			"https://localhost",
		},
	})

	_, err := client.AppRegistryService.CreateApp(app)
	require.NoError(t, err)
	defer client.AppRegistryService.DeleteApp(appName)

	// List all apps
	apps, err := client.AppRegistryService.ListApps()
	require.NoError(t, err)
	// At least the app we just created should be present
	assert.NotZero(t, len(apps))
	//app-reg service bug https://jira.splunk.com/browse/APPLAT-5043
	// assert.EqualValues(t, apps[0], app)

	// Get the app
	app_ret, err := client.AppRegistryService.GetApp(appName)
	require.NoError(t, err)
	require.Equal(t, app.WebAppPost().Name, app_ret.WebApp().Name)
	require.Equal(t, app.WebAppPost().Title, app_ret.WebApp().Title)
	require.Equal(t, app.WebAppPost().RedirectUrls, app_ret.WebApp().RedirectUrls)
	require.Equal(t, string(app.WebAppPost().Kind), string(app_ret.WebApp().Kind))

	// Update the app. TODO: title and redirecturl should not needed once patch method is implemented.
	description := "new Description"
	title := newAppTitle("newtitle")
	redirecturl := []string{"https://newlocalhost"}
	updateApp := appregistry.UpdateAppRequest{
		Description:  &description,
		RedirectUrls: redirecturl,
		Title:        &title,
	}
	app_update_ret, err := client.AppRegistryService.UpdateApp(appName, updateApp)
	require.NoError(t, err)
	require.Equal(t, app.WebAppPost().Name, app_update_ret.WebApp().Name)
	require.Equal(t, title, app_update_ret.WebApp().Title)
	require.Equal(t, description, *app_update_ret.WebApp().Description)
	require.Equal(t, string(app.WebAppPost().Kind), string(app_update_ret.WebApp().Kind))

	// Delete the app
	err = client.AppRegistryService.DeleteApp(appName)
	require.NoError(t, err)
}

// Test RotateSecret in app-registry service
func TestAppRotateSecret(t *testing.T) {
	client := getSdkClient(t)
	appName := fmt.Sprintf("g.r%d", testutils.RunSuffix)

	// Create app
	app := appregistry.MakeCreateAppRequestFromWebAppPost(appregistry.WebAppPost{
		Kind:  appregistry.AppResourceKindWeb,
		Name:  appName,
		Title: newAppTitle("testtitle"),

		RedirectUrls: []string{
			"https://localhost",
		},
	})

	app_created, err := client.AppRegistryService.CreateApp(app)
	require.NoError(t, err)
	defer client.AppRegistryService.DeleteApp(appName)

	// rotate secret
	app_ret, err := client.AppRegistryService.RotateSecret(appName)
	require.NoError(t, err)
	require.NotEmpty(t, app_created.WebApp().ClientId, app_ret.WebApp().ClientSecret)
}

// Test Create/Get/List/Delete subscriptions and get apps/subscriptions in app-registry service
func TestSubscriptions(t *testing.T) {
	client := getSdkClient(t)
	appName := fmt.Sprintf("g.s1%d", testutils.RunSuffix)

	// Create app
	app := appregistry.MakeCreateAppRequestFromWebAppPost(appregistry.WebAppPost{
		Kind:  appregistry.AppResourceKindWeb,
		Name:  appName,
		Title: newAppTitle("testtitle"),

		RedirectUrls: []string{
			"https://localhost",
		},
	})
	_, err := client.AppRegistryService.CreateApp(app)
	require.NoError(t, err)
	defer client.AppRegistryService.DeleteApp(appName)

	// create subscription
	err = client.AppRegistryService.CreateSubscription(appregistry.AppName{AppName: appName})
	require.NoError(t, err)
	defer client.AppRegistryService.DeleteSubscription(appName)

	// get app subscription
	appsubs, err := client.AppRegistryService.ListAppSubscriptions(appName)
	require.NoError(t, err)
	require.Equal(t, 1, len(appsubs))

	// Get a subscription of an app
	subs, err := client.AppRegistryService.GetSubscription(appName)
	require.NoError(t, err)
	require.NotEmpty(t, subs)

	// Get a subscription from non-exist-app
	_, err = client.AppRegistryService.GetSubscription("notExistApp")
	require.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	require.Equal(t, 404, httpErr.HTTPStatusCode)

	// List all subscriptions
	// create the 2nd subscription
	appName2 := fmt.Sprintf("g.s2%d", testutils.RunSuffix)
	perms := []string{"*:action.*"}
	permFilter := []string{"*:*.*"}
	app2 := appregistry.MakeCreateAppRequestFromWebAppPost(appregistry.WebAppPost{
		Kind:                    appregistry.AppResourceKindService,
		Name:                    appName2,
		Title:                   newAppTitle("testtitle-2"),
		AppPrincipalPermissions: perms,
		UserPermissionsFilter:   permFilter,
		RedirectUrls: []string{
			"https://localhost",
		},
	})

	_, err = client.AppRegistryService.CreateApp(app2)
	require.NoError(t, err)
	defer client.AppRegistryService.DeleteApp(appName2)
	err = client.AppRegistryService.CreateSubscription(appregistry.AppName{AppName: appName2})
	require.NoError(t, err)
	defer client.AppRegistryService.DeleteSubscription(appName2)

	query := appregistry.ListSubscriptionsQueryParams{}.SetKind(appregistry.AppResourceKindService)
	all_subs, err := client.AppRegistryService.ListSubscriptions(&query)
	found := 0
	for _, element := range all_subs {
		if element.AppName == appName2 {
			found++
		}
	}
	require.NoError(t, err)
	require.Equal(t, 1, found)

	// Delete the subscription
	err = client.AppRegistryService.DeleteSubscription(appName)
	require.NoError(t, err)
}
