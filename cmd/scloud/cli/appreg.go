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

package main

import (
	"fmt"
	"strings"

	"flag"

	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/argx"
	"github.com/splunk/splunk-cloud-sdk-go/services/appregistry"
)

const (
	AppRegServiceVersion = "v1beta2"
)

var createAppRegistryService = func() *appregistry.Service {
	return apiClient().AppRegistryService
}

type AppRegistryCommand struct {
	appRegistryService *appregistry.Service
}

func newAppRegistryCommand() *AppRegistryCommand {
	return &AppRegistryCommand{
		appRegistryService: createAppRegistryService(),
	}
}

func (appRegCommand *AppRegistryCommand) Dispatch(argv []string) (result interface{}, err error) {
	arg, argv := head(argv)
	switch arg {
	case "":
		eusage("too few arguments")
	case "create-app":
		result, err = appRegCommand.createApp(argv)
	case "create-subscription":
		err = appRegCommand.createSubscription(argv)
	case "delete-app":
		err = appRegCommand.deleteApp(argv)
	case "delete-subscription":
		err = appRegCommand.deleteSubscription(argv)
	case "get-app":
		result, err = appRegCommand.getApp(argv)
	case "get-spec-json":
		result, err = appRegCommand.getSpecJSON(argv)
	case "get-spec-yaml":
		result, err = appRegCommand.getSpecYaml(argv)
	case "get-subscription":
		result, err = appRegCommand.getSubscription(argv)
	case "help":
		err = help("appreg.txt")
	case "list-apps":
		result, err = appRegCommand.listApps(argv)
	case "list-subscriptions":
		result, err = appRegCommand.listSubscriptions(argv)
	case "rotate-secret":
		result, err = appRegCommand.rotateSecret(argv)
	case "update-app":
		result, err = appRegCommand.updateApp(argv)
	default:
		fatal("unknown command: '%s'", arg)
	}
	return
}

type appArgs struct {
	Name                    string `arg:"0"`
	Kind                    string `arg:"1"`
	Description             string `arg:"description"`
	LoginURL                string `arg:"login-url"`
	LogoURL                 string `arg:"logo-url"`
	RedirectURLs            string `arg:"redirect-urls"`
	Title                   string `arg:"title"`
	SetupURL                string `arg:"setup-url"`
	AppPrincipalPermissions string `arg:"app-principal-permissions"`
	UserPermissionsFilter   string `arg:"user-permissions-filter"`
	WebhookURL              string `arg:"webhook-url"`
}

type appUpdateArgs struct {
	Name                    string `arg:"0"`
	Description             string `arg:"description"`
	LoginURL                string `arg:"login-url"`
	LogoURL                 string `arg:"logo-url"`
	RedirectURLs            string `arg:"redirect-urls"`
	Title                   string `arg:"title"`
	SetupURL                string `arg:"setup-url"`
	AppPrincipalPermissions string `arg:"app-principal-permissions"`
	UserPermissionsFilter   string `arg:"user-permissions-filter"`
	WebhookURL              string `arg:"webhook-url"`
}

func (appRegCommand *AppRegistryCommand) createApp(argv []string) (interface{}, error) {
	var args appArgs
	var err error
	_, err = argx.Parse(argv, &args)
	if err != nil {
		return nil, err
	}

	if args.Kind == "" || args.Name == "" {
		etoofew()
	}
	var appKind appregistry.AppResourceKind

	var nativeAppPost appregistry.NativeAppPost
	var serviceAppPost appregistry.ServiceAppPost
	var webAppPost appregistry.WebAppPost

	switch args.Kind {
	case "web":
		appKind = appregistry.AppResourceKindWeb
		webAppPost = appregistry.WebAppPost{
			Name:         args.Name,
			Kind:         appKind,
			Title:        args.Title,
			Description:  &args.Description,
			LogoUrl:      &args.LogoURL,
			LoginUrl:     &args.LoginURL,
			WebhookUrl:   &args.WebhookURL,
			RedirectUrls: strings.Split(args.RedirectURLs, ","),
			SetupUrl:     &args.SetupURL,
		}
		app, err := appRegCommand.appRegistryService.CreateApp(appregistry.MakeCreateAppRequestFromWebAppPost(webAppPost))
		if err != nil {
			return nil, err
		}
		if args.AppPrincipalPermissions != "" {
			appPrincipalPermissions := strings.Split(args.AppPrincipalPermissions, ",")
			app.WebApp().AppPrincipalPermissions = appPrincipalPermissions
		}
		if args.UserPermissionsFilter != "" {
			userPermissionsFilter := strings.Split(args.UserPermissionsFilter, ",")
			app.WebApp().UserPermissionsFilter = userPermissionsFilter
		}
		return app, err
	case "native":
		appKind = appregistry.AppResourceKindNative
		nativeAppPost = appregistry.NativeAppPost{
			Name:         args.Name,
			Kind:         appKind,
			Title:        args.Title,
			Description:  &args.Description,
			LogoUrl:      &args.LogoURL,
			LoginUrl:     &args.LoginURL,
			WebhookUrl:   &args.WebhookURL,
			RedirectUrls: strings.Split(args.RedirectURLs, ","),
			SetupUrl:     &args.SetupURL,
		}
		app, err := appRegCommand.appRegistryService.CreateApp(appregistry.MakeCreateAppRequestFromNativeAppPost(nativeAppPost))
		if err != nil {
			return nil, err
		}
		if args.AppPrincipalPermissions != "" {
			appPrincipalPermissions := strings.Split(args.AppPrincipalPermissions, ",")
			app.NativeApp().AppPrincipalPermissions = appPrincipalPermissions
		}
		if args.UserPermissionsFilter != "" {
			userPermissionsFilter := strings.Split(args.UserPermissionsFilter, ",")
			app.NativeApp().UserPermissionsFilter = userPermissionsFilter
		}
		return app, err
	case "service":
		appKind = appregistry.AppResourceKindService
		serviceAppPost = appregistry.ServiceAppPost{
			Name:         args.Name,
			Kind:         appKind,
			Title:        args.Title,
			Description:  &args.Description,
			LogoUrl:      &args.LogoURL,
			LoginUrl:     &args.LoginURL,
			WebhookUrl:   &args.WebhookURL,
			RedirectUrls: strings.Split(args.RedirectURLs, ","),
			SetupUrl:     &args.SetupURL,
		}
		app, err := appRegCommand.appRegistryService.CreateApp(appregistry.MakeCreateAppRequestFromServiceAppPost(serviceAppPost))
		if err != nil {
			return nil, err
		}
		if args.AppPrincipalPermissions != "" {
			appPrincipalPermissions := strings.Split(args.AppPrincipalPermissions, ",")
			app.NativeApp().AppPrincipalPermissions = appPrincipalPermissions
		}
		if args.UserPermissionsFilter != "" {
			userPermissionsFilter := strings.Split(args.UserPermissionsFilter, ",")
			app.NativeApp().UserPermissionsFilter = userPermissionsFilter
		}
		return app, err
	default:
		msg := fmt.Sprintf("'%v' was passed, use subcommand 'web', 'native', or 'service'", args.Kind)
		fatal(msg)
	}
	return appregistry.WebAppPost{}, fmt.Errorf("appregistry create app failed. please refer to help text for usage")
}

func (appRegCommand *AppRegistryCommand) createSubscription(args []string) error {
	app := head1(args)
	appName := appregistry.AppName{AppName: app}
	return appRegCommand.appRegistryService.CreateSubscription(appName)
}

func (appRegCommand *AppRegistryCommand) deleteApp(args []string) error {
	app := head1(args)
	return appRegCommand.appRegistryService.DeleteApp(app)
}

func (appRegCommand *AppRegistryCommand) deleteSubscription(args []string) error {
	app := head1(args)
	return appRegCommand.appRegistryService.DeleteSubscription(app)
}

func (appRegCommand *AppRegistryCommand) getApp(args []string) (interface{}, error) {
	app := head1(args)
	return appRegCommand.appRegistryService.GetApp(app)
}

func (appRegCommand *AppRegistryCommand) getSubscription(argv []string) (interface{}, error) {
	app := head1(argv)
	return appRegCommand.appRegistryService.GetSubscription(app)
}

func (appRegCommand *AppRegistryCommand) listApps(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return appRegCommand.appRegistryService.ListApps()
}

func (appRegCommand *AppRegistryCommand) listSubscriptions(argv []string) (interface{}, error) {
	// Optional flags
	flags := flag.NewFlagSet("list-subscriptions", flag.ExitOnError)
	var appKind string

	flags.StringVar(&appKind, "app-kind", "", "Action kind")

	err := flags.Parse(argv) //nolint:errcheck
	if err != nil {
		return nil, err
	}
	if appKind != "" {
		switch appKind {
		case "web":
			appModelKind := appregistry.AppResourceKindWeb
			return appRegCommand.appRegistryService.ListSubscriptions(&appregistry.ListSubscriptionsQueryParams{Kind: &appModelKind})
		case "native":
			appModelKind := appregistry.AppResourceKindNative
			return appRegCommand.appRegistryService.ListSubscriptions(&appregistry.ListSubscriptionsQueryParams{Kind: &appModelKind})
		case "service":
			appModelKind := appregistry.AppResourceKindService
			return appRegCommand.appRegistryService.ListSubscriptions(&appregistry.ListSubscriptionsQueryParams{Kind: &appModelKind})
		default:
			msg := fmt.Sprintf("'%v' was passed, use subcommand 'web', 'native', 'service' or default of empty string to list all subscriptions", appKind)
			fatal(msg)
		}
	}
	return appRegCommand.appRegistryService.ListSubscriptions(&appregistry.ListSubscriptionsQueryParams{})
}

func (appRegCommand *AppRegistryCommand) rotateSecret(argv []string) (interface{}, error) {
	app := head1(argv)
	return appRegCommand.appRegistryService.RotateSecret(app)
}

func (appRegCommand *AppRegistryCommand) updateApp(argv []string) (interface{}, error) {
	var args appUpdateArgs
	argv, err := argx.Parse(argv, &args)
	if err != nil {
		return nil, err
	}
	checkEmpty(argv)

	app, err := appRegCommand.appRegistryService.GetApp(args.Name)
	if err != nil {
		return nil, err
	}

	var redirectURLs, appPrincipalPermissions, userPermissionsFilter []string

	if args.RedirectURLs != "" {
		redirectURLs = strings.Split(args.RedirectURLs, ",")
		for index, ele := range redirectURLs {
			redirectURLs[index] = strings.TrimSpace(ele)
		}
	}
	if args.AppPrincipalPermissions != "" {
		appPrincipalPermissions = strings.Split(args.AppPrincipalPermissions, ",")
	}
	if args.UserPermissionsFilter != "" {
		userPermissionsFilter = strings.Split(args.UserPermissionsFilter, ",")
	}

	if app.IsWebApp() {
		return appRegCommand.appRegistryService.UpdateApp(args.Name, appregistry.MakeUpdateAppRequestFromWebAppPut(appregistry.WebAppPut{Title: args.Title,
			Description:             &args.Description,
			LogoUrl:                 &args.LogoURL,
			LoginUrl:                &args.LoginURL,
			SetupUrl:                &args.SetupURL,
			WebhookUrl:              &args.WebhookURL,
			RedirectUrls:            redirectURLs,
			AppPrincipalPermissions: appPrincipalPermissions,
			UserPermissionsFilter:   userPermissionsFilter}))
	}
	return appRegCommand.appRegistryService.UpdateApp(args.Name, appregistry.MakeUpdateAppRequestFromCommonAppPut(appregistry.CommonAppPut{Title: args.Title,
		Description:             &args.Description,
		LogoUrl:                 &args.LogoURL,
		LoginUrl:                &args.LoginURL,
		SetupUrl:                &args.SetupURL,
		WebhookUrl:              &args.WebhookURL,
		RedirectUrls:            redirectURLs,
		AppPrincipalPermissions: appPrincipalPermissions,
		UserPermissionsFilter:   userPermissionsFilter}))
}

func (appRegCommand *AppRegistryCommand) getSpecJSON(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecJSON("api", AppRegServiceVersion, "app-registry", appRegCommand.appRegistryService.Client)
}

func (appRegCommand *AppRegistryCommand) getSpecYaml(argv []string) (interface{}, error) {
	checkEmpty(argv)
	return GetSpecYaml("api", AppRegServiceVersion, "app-registry", appRegCommand.appRegistryService.Client)
}
