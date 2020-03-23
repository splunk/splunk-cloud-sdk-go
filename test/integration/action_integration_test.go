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
	"net/http"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services/action"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test consts
var (
	body           = "<html><h1>The HTML</h1></html>"
	bodyPlainText  = "This is a plain text body."
	subject        = "The Subject"
	title          = "The Title."
	webhookURL     = "https://webhook.site/test"
	webhookPayload = "{{ .name }} is a {{ .species }}"
)

// Test vars
var (
	addresses    = []string{"success@simulator.amazonses.com", "success@simulator.amazonses.com"}
	triggerEvent = genTriggerEvent()
)

func genTriggerEvent() action.TriggerEvent {
	payload := action.RawJsonPayload{"name": "beanbag", "species": "cat"}
	var kind action.TriggerEventKind
	kind = action.TriggerEventKindTrigger
	return action.TriggerEvent{Kind: &kind, Payload: &payload}
}

func genEmailAction() action.Action {
	emailActionName := fmt.Sprintf("e_cr_%d", testutils.RunSuffix)
	emailAction := action.EmailAction{Kind: action.ActionKindEmail, Name: emailActionName, Title: &title, Body: &body, BodyPlainText: &bodyPlainText, Subject: &subject}
	return action.MakeActionFromEmailAction(emailAction)
}

func genWebhookAction() action.Action {
	webhookActionName := fmt.Sprintf("w_cr_%d", testutils.RunSuffix)
	webhookAction := action.WebhookAction{Kind: action.ActionKindWebhook, Name: webhookActionName, Title: &title, WebhookPayload: webhookPayload, WebhookUrl: webhookURL}
	return action.MakeActionFromWebhookAction(webhookAction)
}

func cleanupAction(client *sdk.Client, name string) {
	err := client.ActionService.DeleteAction(name)
	if err != nil {
		fmt.Printf("WARN: error deleting action: %s, err: %s", name, err)
	}
}

func validateUnauthenticatedActionError(t *testing.T, err error) {
	require.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok, fmt.Sprintf("error casting err to HTTPError, err: %+v", err))
	assert.Equal(t, 401, httpErr.HTTPStatusCode)
	assert.Equal(t, "401 Unauthorized", httpErr.HTTPStatus)
}

func validateNotFoundActionError(t *testing.T, err error) {
	require.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok, fmt.Sprintf("error casting err to HTTPError, err: %+v", err))
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
	assert.Equal(t, "404 Not Found", httpErr.HTTPStatus)
}

// Test GetActions which returns the list of all actions for the tenant
func TestIntegrationGetActions(t *testing.T) {
	client := getSdkClient(t)

	// Get Actions
	actions, err := client.ActionService.ListActions()
	require.NoError(t, err)
	assert.True(t, len(actions) >= 0)
}

// Test CreateAction / GetAction for EmailAction
func TestGetCreateActionEmail(t *testing.T) {
	client := getSdkClient(t)
	act := genEmailAction()
	var resp http.Response // provide a response pointer for response info
	_, err := client.ActionService.CreateAction(act, &resp)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	in := act.EmailAction()
	require.NotNil(t, in)
	defer cleanupAction(client, in.Name)
	actout, err := client.ActionService.GetAction(in.Name)
	require.NoError(t, err)
	out := actout.EmailAction()
	require.NotNil(t, out)
	assert.Equal(t, in.Kind, out.Kind)
	assert.Equal(t, in.Name, out.Name)
	assert.Equal(t, in.Subject, out.Subject)
	assert.Equal(t, in.Title, out.Title)
	assert.Equal(t, in.Body, out.Body)
	assert.Equal(t, in.BodyPlainText, out.BodyPlainText)
	assert.Equal(t, in.Addresses, out.Addresses)
}

// Test CreateAction / GetAction for WebhookAction
func TestGetCreateActionWebhook(t *testing.T) {
	client := getSdkClient(t)
	act := genWebhookAction()
	in := act.WebhookAction()
	require.NotNil(t, in)
	actout, err := client.ActionService.CreateAction(act)
	defer cleanupAction(client, in.Name)
	require.NoError(t, err)
	out := actout.WebhookAction()
	require.NotNil(t, out)
	assert.EqualValues(t, in.Kind, out.Kind)
	assert.EqualValues(t, in.Name, out.Name)
	assert.EqualValues(t, *in.Title, *out.Title)
	assert.EqualValues(t, in.WebhookUrl, out.WebhookUrl)
	assert.EqualValues(t, in.WebhookPayload, out.WebhookPayload)
	a, err := client.ActionService.GetAction(in.Name)
	require.NoError(t, err)
	out = a.WebhookAction()
	require.NotNil(t, out)
	assert.EqualValues(t, in.Kind, out.Kind)
	assert.EqualValues(t, in.Name, out.Name)
	assert.EqualValues(t, *in.Title, *out.Title)
	assert.EqualValues(t, in.WebhookUrl, out.WebhookUrl)
	assert.EqualValues(t, in.WebhookPayload, out.WebhookPayload)
}

// Invalid action name should result in 400 Bad Request
func TestCreateActionFailInvalidAction(t *testing.T) {
	client := getSdkClient(t)
	// Get Invalid Action
	_, err := client.ActionService.GetAction("NoCapitals")

	require.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok, fmt.Sprintf("error casting err to HTTPError, err: %+v", err))
	assert.Equal(t, 400, httpErr.HTTPStatusCode)
	assert.Equal(t, "400 Bad Request", httpErr.HTTPStatus)
}

// Getting non-existent action should result in 404
func TestCreateActionFailNonExistentAction(t *testing.T) {
	client := getSdkClient(t)
	// Get Invalid Action
	_, err := client.ActionService.GetAction("dontexist")

	require.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok, fmt.Sprintf("error casting err to HTTPError, err: %+v", err))
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Create Existing action should result in 409 Conflict
func TestCreateActionFailExistingAction(t *testing.T) {
	client := getSdkClient(t)
	act := genEmailAction()
	_, err := client.ActionService.CreateAction(act)
	require.NoError(t, err)
	in := act.EmailAction()
	require.NotNil(t, in)
	defer cleanupAction(client, in.Name)
	actout, err := client.ActionService.GetAction(in.Name)
	require.NoError(t, err)
	out := actout.EmailAction()
	require.NotNil(t, out)
	assert.EqualValues(t, in.Name, out.Name)

	_, err = client.ActionService.CreateAction(act)
	require.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok, fmt.Sprintf("error casting err to HTTPError, err: %+v", err))
	assert.Equal(t, 409, httpErr.HTTPStatusCode)
	assert.Equal(t, "409 Conflict", httpErr.HTTPStatus)
}

// Access action endpoints using an Unauthenticated client results in a 401 Unauthenticated error
func TestActionFailUnauthenticatedClient(t *testing.T) {
	client := getSdkClient(t)
	wact := genWebhookAction()
	webhookAction := wact.WebhookAction()
	require.NotNil(t, webhookAction)
	_, err := client.ActionService.CreateAction(wact)
	require.NoError(t, err)
	defer cleanupAction(client, webhookAction.Name)

	eact := genEmailAction()
	// This shouldn't be needed since the CreateAction should fail for 401:
	// defer cleanupAction(client, emailAction.EmailAction.Name)
	invalidClient := getInvalidClient(t)

	_, err = invalidClient.ActionService.CreateAction(eact)
	if err == nil {
		// We aren't expecting this to succeed, but if it does then clean up the created action
		defer cleanupAction(client, eact.EmailAction().Name)
	}
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.GetAction(webhookAction.Name)
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.ListActions()
	validateUnauthenticatedActionError(t, err)

	err = invalidClient.ActionService.TriggerAction(webhookAction.Name,
		genTriggerEvent())
	validateUnauthenticatedActionError(t, err)

	str := "new payload is a {{ .species }}"
	_, err = invalidClient.ActionService.UpdateAction(webhookAction.Name,
		action.MakeActionMutableFromWebhookActionMutable(action.WebhookActionMutable{WebhookPayload: &str}))
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.GetActionStatus("Action123", "statusID")
	validateUnauthenticatedActionError(t, err)

	err = invalidClient.ActionService.DeleteAction(webhookAction.Name)
	validateUnauthenticatedActionError(t, err)
}

// Test UpdateAction updates with the new fields in the action
func TestUpdateAction(t *testing.T) {
	client := getSdkClient(t)
	act := genEmailAction()
	_, err := client.ActionService.CreateAction(act)
	require.NoError(t, err)
	emailAction := act.EmailAction()
	require.NotNil(t, emailAction)
	defer cleanupAction(client, emailAction.Name)
	var newText = "updated email text"
	var newTitle = "I am a new title"
	actout, err := client.ActionService.UpdateAction(emailAction.Name,
		action.MakeActionMutableFromEmailActionMutable(action.EmailActionMutable{
			Title:   &newTitle,
			Subject: &newText,
		}))
	require.NoError(t, err)
	require.NotNil(t, actout)
	require.NotNil(t, actout.EmailAction)
	assert.Equal(t, newText, *actout.EmailAction().Subject)
	assert.Equal(t, newTitle, *actout.EmailAction().Title)
}

// Test DeleteAction deletes the action specified
func TestDeleteAction(t *testing.T) {
	client := getSdkClient(t)
	act := genEmailAction()
	_, err := client.ActionService.CreateAction(act)
	require.NoError(t, err)
	err = client.ActionService.DeleteAction(act.EmailAction().Name)
	assert.NoError(t, err)
}

// Get non-existent Action results in a 404 NotFound
func TestActionFailNotFoundAction(t *testing.T) {
	client := getSdkClient(t)
	// Get Invalid Action
	_, err := client.ActionService.GetAction("action123")
	validateNotFoundActionError(t, err)

	_, err = client.ActionService.GetActionStatus("action123", "statusID")
	validateNotFoundActionError(t, err)

	_, err = client.ActionService.UpdateAction("action123", action.ActionMutable{})
	validateNotFoundActionError(t, err)

	err = client.ActionService.DeleteAction("action123")
	validateNotFoundActionError(t, err)
}

// TestGetActionStatus gets the status of the action after it is triggered
func TestGetActionStatus(t *testing.T) {
	client := getSdkClient(t)
	wact := genWebhookAction()
	webhookAction := wact.WebhookAction()
	require.NotNil(t, webhookAction)
	act, err := client.ActionService.CreateAction(wact)
	require.NoError(t, err)
	defer cleanupAction(client, webhookAction.Name)
	require.NotNil(t, act)
	resp, err := client.ActionService.TriggerActionWithStatus(webhookAction.Name,
		genTriggerEvent())
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotEmpty(t, *resp.StatusID)

	stat, err := client.ActionService.GetActionStatus(webhookAction.Name, *resp.StatusID)
	assert.NoError(t, err)
	assert.NotNil(t, stat)
}
