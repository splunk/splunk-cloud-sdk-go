// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"fmt"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/service"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test consts
const (
	htmlPart     = "<html><h1>The HTML</h1></html>"
	subjectPArt  = "The Subject"
	textPart     = "The Text"
	templateName = "template1000"
	snsTopic     = "myTopic"
	snsMsg       = "SNS Message"
	webhookURL   = "https://webhook.site/test"
	webhookMsg   = "{{ .name }} is a {{ .species }}"
	actionUserID = "sdk_tester"
)

// Test vars
var (
	timeSec           = time.Now().Unix()
	snsActionName     = fmt.Sprintf("sact_%d", timeSec)
	webhookActionName = fmt.Sprintf("wact_%d", timeSec)
	addresses         = []string{"test1@splunk.com", "test2@splunk.com"}
	webhookPayload    = &map[string]interface{}{"name": "bean bag", "species": "cat"}
)

func cleanupAction(client *service.Client, name string) {
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
	client := getClient(t)

	// Get Actions
	actions, err := client.ActionService.GetActions()
	require.Nil(t, err)
	assert.True(t, len(actions) >= 0)
}

// Test CreateAction / GetAction for EmailAction
func TestGetCreateActionEmail(t *testing.T) {
	client := getClient(t)
	emailActionName := fmt.Sprintf("e_cr_%d", timeSec)
	emailAction := model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
	defer cleanupAction(client, emailAction.Name)
	_, err := client.ActionService.CreateAction(*emailAction)
	require.Nil(t, err)
	action, err := client.ActionService.GetAction(emailAction.Name)
	assert.EqualValues(t, action, emailAction)
}

// Test CreateAction / GetAction for SNSAction
func TestGetCreateActionSNS(t *testing.T) {
	client := getClient(t)
	snsActionName := fmt.Sprintf("s_cr_%d", timeSec)
	snsAction := model.NewSNSAction(snsActionName, snsTopic, snsMsg)
	defer cleanupAction(client, snsAction.Name)
	_, err := client.ActionService.CreateAction(*snsAction)
	require.Nil(t, err)
	action, err := client.ActionService.GetAction(snsAction.Name)
	assert.EqualValues(t, action, snsAction)
}

// Test CreateAction / GetAction for WebhookAction
func TestGetCreateActionWebhook(t *testing.T) {
	client := getClient(t)
	webhookActionName := fmt.Sprintf("w_cr_%d", timeSec)
	webhookAction := model.NewWebhookAction(webhookActionName, webhookURL, webhookMsg)
	defer cleanupAction(client, webhookAction.Name)
	_, err := client.ActionService.CreateAction(*webhookAction)
	require.Nil(t, err)
	action, err := client.ActionService.GetAction(webhookAction.Name)
	assert.EqualValues(t, action, webhookAction)
}

// Get Non-Existent Action should result in 404 Not Found
func TestCreateActionFailInvalidAction(t *testing.T) {
	client := getClient(t)
	// Get Invalid Action
	_, err := client.ActionService.GetAction("Dontexist")

	require.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok, fmt.Sprintf("error casting err to HTTPError, err: %+v", err))
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
	assert.Equal(t, "404 Not Found", httpErr.HTTPStatus)
}

// Create Existing action should result in 409 Conflict
func TestCreateActionFailExistingAction(t *testing.T) {
	client := getClient(t)
	emailActionName := fmt.Sprintf("e_confl_%d", timeSec)
	emailAction := model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
	defer cleanupAction(client, emailAction.Name)
	_, err := client.ActionService.CreateAction(*emailAction)
	require.Nil(t, err)
	action, err := client.ActionService.GetAction(emailAction.Name)
	assert.EqualValues(t, action, emailAction)

	_, err = client.ActionService.CreateAction(*emailAction)
	require.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok, fmt.Sprintf("error casting err to HTTPError, err: %+v", err))
	assert.Equal(t, 409, httpErr.HTTPStatusCode)
	assert.Equal(t, "409 Conflict", httpErr.HTTPStatus)
}

// Access action endpoints using an Unauthenticated client results in a 401 Unauthenticated error
func TestActionFailUnauthenticatedClient(t *testing.T) {
	client := getClient(t)
	webhookActionName := fmt.Sprintf("w_unauth_%d", timeSec)
	webhookAction := model.NewWebhookAction(webhookActionName, webhookURL, webhookMsg)
	defer cleanupAction(client, webhookAction.Name)
	_, err := client.ActionService.CreateAction(*webhookAction)
	require.Nil(t, err)

	emailActionName := fmt.Sprintf("e_unauth_%d", timeSec)
	emailAction := model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
	// This shouldn't be needed since the CreateAction should fail for 401:
	// defer cleanupAction(client, emailAction.Name)
	invalidClient := getInvalidClient(t)

	_, err = invalidClient.ActionService.CreateAction(*emailAction)
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.GetAction(webhookAction.Name)
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.GetActions()
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.TriggerAction(webhookAction.Name,
		model.ActionNotification{
			Kind:    model.RawJSONPayloadKind,
			Tenant:  testutils.TestTenant,
			Payload: webhookPayload,
		})
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.UpdateAction(webhookActionName, model.ActionUpdateFields{TextPart: "updated email text"})
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.GetActionStatus("Action123", "statusID")
	validateUnauthenticatedActionError(t, err)

	err = invalidClient.ActionService.DeleteAction(webhookActionName)
	validateUnauthenticatedActionError(t, err)
}

// Trigger action with invalid fields results in a 422 Unprocessable Entity error
func TestTriggerActionFailInvalidFields(t *testing.T) {
	client := getClient(t)
	webhookActionName := fmt.Sprintf("w_unproc_%d", timeSec)
	webhookAction := model.NewWebhookAction(webhookActionName, webhookURL, webhookMsg)
	defer cleanupAction(client, webhookAction.Name)
	_, err := client.ActionService.CreateAction(*webhookAction)
	require.Nil(t, err)
	_, err = client.ActionService.TriggerAction(webhookAction.Name,
		model.ActionNotification{
			Kind:    model.RawJSONPayloadKind,
			Tenant:  "",
			Payload: webhookPayload,
		})

	require.NotEmpty(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok, fmt.Sprintf("error casting err to HTTPError, err: %+v", err))
	assert.Equal(t, 422, httpErr.HTTPStatusCode)
	assert.Equal(t, "validation-failed", httpErr.Code)
}

// Test UpdateAction updates with the new fields in the action
func TestUpdateAction(t *testing.T) {
	client := getClient(t)
	emailActionName := fmt.Sprintf("e_up_%d", timeSec)
	emailAction := model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
	defer cleanupAction(client, emailAction.Name)
	_, err := client.ActionService.CreateAction(*emailAction)
	require.Nil(t, err)
	const newText = "updated email text"
	result, err := client.ActionService.UpdateAction(emailActionName, model.ActionUpdateFields{TextPart: newText})
	assert.Equal(t, result.TextPart, newText)
	assert.Nil(t, err)
}

// Test DeleteAction deletes the action specified
func TestDeleteAction(t *testing.T) {
	client := getClient(t)
	emailActionName := fmt.Sprintf("e_del_%d", timeSec)
	emailAction := model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
	_, err := client.ActionService.CreateAction(*emailAction)
	require.Nil(t, err)
	err = client.ActionService.DeleteAction(emailActionName)
	assert.Nil(t, err)
}

// Access action endpoints with a non-existent Action results in a 404 NotFound
func TestActionFailNotFoundAction(t *testing.T) {
	client := getClient(t)
	// Get Invalid Action
	_, err := client.ActionService.GetAction("Action123")
	validateNotFoundActionError(t, err)

	_, err = client.ActionService.GetActionStatus("Action123", "statusID")
	validateNotFoundActionError(t, err)

	_, err = client.ActionService.UpdateAction("Action123", model.ActionUpdateFields{TextPart: "updated email text"})
	validateNotFoundActionError(t, err)

	err = client.ActionService.DeleteAction("Action123")
	validateNotFoundActionError(t, err)
}

// TestGetActionStatus gets the status of the action after it is triggered
func TestGetActionStatus(t *testing.T) {
	client := getClient(t)
	webhookActionName := fmt.Sprintf("w_stat_%d", timeSec)
	webhookAction := model.NewWebhookAction(webhookActionName, webhookURL, webhookMsg)
	defer cleanupAction(client, webhookAction.Name)
	action, err := client.ActionService.CreateAction(*webhookAction)
	require.Nil(t, err)
	require.NotNil(t, action)
	resp, err := client.ActionService.TriggerAction(webhookAction.Name,
		model.ActionNotification{
			Kind:    model.RawJSONPayloadKind,
			Tenant:  testutils.TestTenant,
			Payload: webhookPayload,
		})
	require.Nil(t, err)
	require.NotNil(t, resp.StatusID)

	stat, err := client.ActionService.GetActionStatus(webhookAction.Name, *resp.StatusID)
	assert.Nil(t, err)
	assert.NotNil(t, stat)
}

// TestTriggerActionTenantMismatch triggers an action with tenant not matching the URL
func TestTriggerActionTenantMismatch(t *testing.T) {
	client := getClient(t)
	webhookActionName := fmt.Sprintf("w_badten_%d", timeSec)
	webhookAction := model.NewWebhookAction(webhookActionName, webhookURL, webhookMsg)
	defer cleanupAction(client, webhookAction.Name)
	action, err := client.ActionService.CreateAction(*webhookAction)
	require.Nil(t, err)
	require.NotNil(t, action)
	_, err = client.ActionService.TriggerAction(webhookAction.Name,
		model.ActionNotification{
			Kind:    model.RawJSONPayloadKind,
			Tenant:  "INCORRECT_TENANT",
			Payload: webhookPayload,
		})
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok, fmt.Sprintf("error casting err to HTTPError, err: %+v", err))
	assert.Equal(t, 403, httpErr.HTTPStatusCode)
	assert.Equal(t, "403 Forbidden", httpErr.HTTPStatus)
}
