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
	"encoding/json"
	"errors"
	"flag"
	"net/http"
	"strings"

	"github.com/splunk/splunk-cloud-sdk-go/services/action"
)

const (
	ActionServiceVersion = "v1beta2"
)

//
// ./scloud action <command> [command-options]
//

var createActionService = func() *action.Service {
	return apiClient().ActionService
}

type ActionCommand struct {
	actionService *action.Service
}

func newActionCommand() *ActionCommand {
	return &ActionCommand{
		actionService: createActionService(),
	}
}

func (cmd *ActionCommand) parse(args []string) []string {
	flags := flag.NewFlagSet("action command", flag.ExitOnError)
	flags.Parse(args) //nolint:errcheck
	return flags.Args()
}

func (cmd *ActionCommand) Dispatch(args []string) (result interface{}, err error) {
	arg, args := head(args)
	args = cmd.parse(args)
	switch arg {
	case "":
		eusage("too few arguments")
	case "create-action":
		result, err = cmd.createAction(args)
	case "delete-action":
		err = cmd.deleteAction(args)
	case "get-action":
		result, err = cmd.getAction(args)
	case "get-action-status":
		result, err = cmd.getActionStatus(args)
	case "list-actions":
		result, err = cmd.getActions(args)
	case "get-spec-json":
		result, err = cmd.getSpecJSON(args)
	case "get-spec-yaml":
		result, err = cmd.getSpecYaml(args)
	case "help":
		err = help("action.txt")
	case "trigger-action":
		result, err = cmd.triggerAction(args)
	case "update-action":
		result, err = cmd.updateAction(args)
	default:
		fatal("unknown sub-command: '%s'", arg)
	}
	return
}

func (cmd *ActionCommand) createAction(args []string) (interface{}, error) {

	// Required args
	name, args := head(args)
	kindString, args := head(args)
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
		fatal(err.Error())
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
		return cmd.actionService.CreateAction(action.MakeActionFromEmailAction(act))

	} else if strings.ToLower(kindString) == string(action.ActionKindWebhook) {
		act1 := action.WebhookAction{
			Name:           name,
			Kind:           action.ActionKindWebhook,
			WebhookUrl:     webhookURL,
			WebhookPayload: webhookPayload,
		}
		return cmd.actionService.CreateAction(action.MakeActionFromWebhookAction(act1))
	}

	return nil, errors.New("Kind value is not supported")
}

func (cmd *ActionCommand) deleteAction(args []string) error {
	name := head1(args)
	return cmd.actionService.DeleteAction(name)
}

func (cmd *ActionCommand) getAction(args []string) (interface{}, error) {
	name := head1(args)
	return cmd.actionService.GetAction(name)
}

func (cmd *ActionCommand) getActionStatus(args []string) (interface{}, error) {
	name, statusID := head2(args)
	return cmd.actionService.GetActionStatus(name, statusID)
}

func (cmd *ActionCommand) getActions(args []string) (interface{}, error) {
	return cmd.actionService.ListActions()
}

func (cmd *ActionCommand) triggerAction(args []string) (interface{}, error) {
	name, args := head(args)
	kindString, args := head(args)

	notificationPayloadJSON, args := head(args)
	checkEmpty(args)

	var notificationPayload action.RawJsonPayload

	err := json.Unmarshal([]byte(notificationPayloadJSON), &notificationPayload)
	if err != nil {
		return nil, err
	}

	tenant := getTenantName()

	kind := action.TriggerEventKind(kindString)
	notification := action.TriggerEvent{
		Kind:    &kind,
		Tenant:  &tenant,
		Payload: &notificationPayload,
	}

	var resp http.Response
	err = cmd.actionService.TriggerAction(name, notification, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Header.Get("Location"), cmd.actionService.TriggerAction(name, notification)
}

func (cmd *ActionCommand) updateAction(args []string) (interface{}, error) {
	name, args := head(args)

	flags := flag.NewFlagSet("update-action", flag.ExitOnError)

	act, err := cmd.actionService.GetAction(name)
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
		fatal(err.Error())
	}

	if len(addressesInput) > 0 {
		addresses = strings.Split(addressesInput, ",")
	}

	if act.IsEmailAction() {

		au := action.EmailActionMutable{
			Body:      cmd.parseStrInput(body),
			Subject:   cmd.parseStrInput(subject),
			Title:     cmd.parseStrInput(title),
			Addresses: addresses,
		}
		return cmd.actionService.UpdateAction(name, action.MakeActionMutableFromEmailActionMutable(au))

	}

	if act.IsWebhookAction() {
		au := action.WebhookActionMutable{
			Title:          cmd.parseStrInput(title),
			WebhookUrl:     cmd.parseStrInput(webhookURL),
			WebhookPayload: cmd.parseStrInput(webhookPayload),
		}

		return cmd.actionService.UpdateAction(name, action.MakeActionMutableFromWebhookActionMutable(au))
	}

	return nil, errors.New("Action kind and updated fields didn't match")
}

func (cmd *ActionCommand) getSpecJSON(args []string) (interface{}, error) {
	checkEmpty(args)
	return GetSpecJSON("api", ActionServiceVersion, "action", cmd.actionService.Client)
}

func (cmd *ActionCommand) getSpecYaml(args []string) (interface{}, error) {
	checkEmpty(args)
	return GetSpecYaml("api", ActionServiceVersion, "action", cmd.actionService.Client)
}

func (cmd *ActionCommand) parseStrInput(input string) *string {
	if input == "" {
		return nil
	}
	return &input
}
