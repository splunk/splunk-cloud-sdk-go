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

package action

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/utils"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services/action"
	"net/http"
	"strings"
)

//const (
//	ActionServiceVersion = "v1beta2"
//)

//
// ./scloud action <command> [command-options]
//

//var createActionService = func() *action.Service {
//	return apiClient().ActionService
//}
//
//type ActionCommand struct {
//	actionService *sdk.Client
//}
//
//func newActionCommand() *ActionCommand {
//	return &ActionCommand{
//		actionService: createActionService(),
//	}
//}
//
//func  parse(cmd *sdk.Client,args []string) []string {
//	flags := flag.NewFlagSet("action command", flag.ExitOnError)
//	flags.Parse(args) //nolint:errcheck
//	return flags.Args()
//}
//
//func  Dispatch(cmd *sdk.Client,args []string) (result interface{}, err error) {
//	arg, args := utils.Head(args)
//	args = cmd.parse(args)
//	switch arg {
//	case "":
//		utils.Eusage("too few arguments")
//	case "create-action":
//		result, err = cmd.createAction(args)
//	case "delete-action":
//		err = cmd.deleteAction(args)
//	case "get-action":
//		result, err = cmd.getAction(args)
//	case "get-action-status":
//		result, err = cmd.getActionStatus(args)
//	case "list-actions":
//		result, err = cmd.getActions(args)
//	case "get-spec-json":
//		result, err = cmd.getSpecJSON(args)
//	case "get-spec-yaml":
//		result, err = cmd.getSpecYaml(args)
//	case "help":
//		err = utils.Help("action.txt")
//	case "trigger-action":
//		result, err = cmd.triggerAction(args)
//	case "update-action":
//		result, err = cmd.updateAction(args)
//	default:
//		utils.Fatal("unknown sub-command: '%s'", arg)
//	}
//	return
//}

func CreateAction(cmd *sdk.Client,args []string) (interface{}, error) {

	// Required args
	name, args := utils.Head(args)
	kindString, args := utils.Head(args)
	flags := flag.NewFlagSet("create-action", flag.ExitOnError)

	var addressesInput, body, subject, title, webhookPayload, webhookURL string
	var addresses []string

	flags.StringVar(&title, "title", "", "The action title")

	// email action
	flags.StringVar(&body, "body", "", "Email action: The email body")
	flags.StringVar(&subject, "subject", "", "Email action: The email subject")
	flags.StringVar(
		&addressesInput,
		"addresses",
		"",
		"Email action: Addresses to send to when the email action is triggered",
	)

	// webhook action
	flags.StringVar(&webhookURL, "webhookURL", "", "Webhook action: WebhookURL to trigger Webhook action")
	flags.StringVar(&webhookPayload, "webhookPayload", "", "Webhook action: Payload to send over Webhook action")

	if len(addressesInput) > 0 {
		addresses = strings.Split(addressesInput, ",")
	}

	err := flags.Parse(args)

	if err != nil {
		utils.Fatal(err.Error())
	}

	if strings.ToLower(kindString) == string(action.ActionKindEmail) {
		act := action.EmailAction{
			Name:      name,
			Kind:      action.ActionKindEmail,
			Body:      &body,
			Subject:   &subject,
			Title:     &title,
			Addresses: addresses,
		}
		return cmd.ActionService.CreateAction(action.MakeActionFromEmailAction(act))

	} else if strings.ToLower(kindString) == string(action.ActionKindWebhook) {
		act1 := action.WebhookAction{
			Name:           name,
			Kind:           action.ActionKindWebhook,
			WebhookUrl:     webhookURL,
			WebhookPayload: webhookPayload,
		}
		return cmd.ActionService.CreateAction(action.MakeActionFromWebhookAction(act1))
	}

	return nil, errors.New("Kind value is not supported")
}

func  DeleteAction(cmd *sdk.Client,args []string) (interface{}, error) {
	name := utils.Head1(args)
	return nil,cmd.ActionService.DeleteAction(name)
}

func  GetAction(cmd *sdk.Client, args []string) (interface{}, error) {
	name := utils.Head1(args)
	return cmd.ActionService.GetAction(name)
}

func  GetActionStatus(cmd *sdk.Client,args []string) (interface{}, error) {
	name, statusID := utils.Head2(args)
	return cmd.ActionService.GetActionStatus(name, statusID)
}

func  ListActions(cmd *sdk.Client,args []string) (interface{}, error) {
	return cmd.ActionService.ListActions()
}

func  TriggerAction(cmd *sdk.Client,args []string) (interface{}, error) {
	name, args := utils.Head(args)
	kindString, args := utils.Head(args)

	notificationPayloadJSON, args := utils.Head(args)
	utils.CheckEmpty(args)

	var notificationPayload action.RawJsonPayload

	err := json.Unmarshal([]byte(notificationPayloadJSON), &notificationPayload)
	if err != nil {
		return nil, err
	}

	tenant := utils.GetTenantName()

	notification := action.TriggerEvent{
		Kind:    action.TriggerEventKind(kindString),
		Tenant:  &tenant,
		Payload: &notificationPayload,
	}

	var resp http.Response
	err = cmd.ActionService.TriggerAction(name, notification, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Header.Get("Location"), cmd.ActionService.TriggerAction(name, notification)
}

func  UpdateAction(cmd *sdk.Client,args []string) (interface{}, error) {
	name, args := utils.Head(args)

	flags := flag.NewFlagSet("update-action", flag.ExitOnError)

	act, err := cmd.ActionService.GetAction(name)
	if err != nil {
		return nil, err
	}

	var addressesInput, body, subject, title, webhookPayload, webhookURL string
	var addresses []string

	flags.StringVar(&title, "title", "", "The action title")

	// email action
	flags.StringVar(&body, "body", "", "Email action: The email body")
	flags.StringVar(&subject, "subject", "", "Email action: The email subject")
	flags.StringVar(
		&addressesInput,
		"addresses",
		"",
		"Email action: Addresses to send to when the email action is triggered",
	)

	// webhook action
	flags.StringVar(&webhookURL, "webhookURL", "", "Webhook action: WebhookURL to trigger Webhook action")
	flags.StringVar(&webhookPayload, "webhookPayload", "", "Webhook action: Payload to send over Webhook action")

	if len(addressesInput) > 0 {
		addresses = strings.Split(addressesInput, ",")
	}

	err = flags.Parse(args)

	if err != nil {
		utils.Fatal(err.Error())
	}

	if len(addressesInput) > 0 {
		addresses = strings.Split(addressesInput, ",")
	}

	if act.IsEmailAction() {

		au := action.EmailActionMutable{
			Body:      parseStrInput(body),
			Subject:   parseStrInput(subject),
			Title:     parseStrInput(title),
			Addresses: addresses,
		}
		return cmd.ActionService.UpdateAction(name, action.MakeActionMutableFromEmailActionMutable(au))

	}

	if act.IsWebhookAction() {
		au := action.WebhookActionMutable{
			Title:          parseStrInput(title),
			WebhookUrl:     parseStrInput(webhookURL),
			WebhookPayload: parseStrInput(webhookPayload),
		}

		return cmd.ActionService.UpdateAction(name, action.MakeActionMutableFromWebhookActionMutable(au))
	}

	return nil, errors.New("Action kind and updated fields didn't match")
}

func  GetActionStatusDetails(cmd *sdk.Client,args []string) (interface{}, error) {
	fmt.Println("Not implemented yet")
	return nil,nil
}

func  GetPublicWebhookKeys(cmd *sdk.Client,args []string) (interface{}, error) {
	fmt.Println("Not implemented yet")
	return nil,nil
}

//
//func  GetSpecJSON(cmd *sdk.Client,args []string) (interface{}, error) {
//	utils.CheckEmpty(args)
//	return GetSpecJSON("api", ActionServiceVersion, "action", cmd.ActionService.Client)
//}
//
//func  GetSpecYaml(cmd *sdk.Client,args []string) (interface{}, error) {
//	utils.CheckEmpty(args)
//	return GetSpecYaml("api", ActionServiceVersion, "action", cmd.ActionService.Client)
//}

func  parseStrInput(input string) *string {
	if input == "" {
		return nil
	}
	return &input
}
