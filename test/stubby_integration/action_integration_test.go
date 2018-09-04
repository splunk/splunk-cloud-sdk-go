// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package stubbyintegration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/testutils"
)

// Stubby test for Create Action service endpoint
func TestCreateAction(t *testing.T) {
	actionData := model.NewWebhookAction("test10", "http://webhook.site/40fa145b-43d7-48f9-a391-aa7558042fa6", "{{ .name }} is a {{ .species }}")
	result, err := getClient(t).ActionService.CreateAction(*actionData)
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.Equal(t, "test10", result.Name)
	assert.Equal(t, "http://webhook.site/40fa145b-43d7-48f9-a391-aa7558042fa6", result.WebhookURL)
	assert.Equal(t, "{{ .name }} is a {{ .species }}", result.Message)
}

// Stubby test for Post trigger action
func TestTriggerWebHookAction(t *testing.T) {
	var payloadweb = &map[string]interface{}{"name": "bean bag"}
	actionNotificationData := model.ActionNotification{Kind: model.RawJSONPayloadKind, Tenant: testutils.TestTenantID, Payload: payloadweb}
	u, err := getClient(t).ActionService.TriggerAction("test10", actionNotificationData)
	assert.Empty(t, err)
	assert.NotEmpty(t, u)
}

// Stubby test for GetAction service endpoint
func TestGetAction(t *testing.T) {
	result, err := getClient(t).ActionService.GetAction("test10")
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.Equal(t, "test10", result.Name)
	assert.Equal(t, "http://webhook.site/40fa145b-43d7-48f9-a391-aa7558042fa6", result.WebhookURL)
	assert.Equal(t, "{{ .name }} is a {{ .species }}", result.Message)
}

// Stubby test for Get All actions Action service endpoint
func TestGetAllActions(t *testing.T) {
	result, err := getClient(t).ActionService.GetActions()
	assert.Empty(t, err)
	assert.NotEmpty(t, result)
}

// Stubby test for Get Action Status Action service endpoint
func TestGetActionStatus(t *testing.T) {
	result, err := getClient(t).ActionService.GetActionStatus("test10", "5f718aaf-f205-4af6-995f-54a3ba059b59")
	assert.Empty(t, err)
	assert.NotEmpty(t, result)
}

// Stubby test for Delete action service endpoint
func TestDeleteAction(t *testing.T) {
	err := getClient(t).ActionService.DeleteAction("test10")
	assert.Empty(t, err)
}

//Stubby test for Update action service endpoint
func TestUpdateAction(t *testing.T) {
	actionData := model.ActionUpdateFields{Message: "updated message"}
	result, err := getClient(t).ActionService.UpdateAction("test10", actionData)
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.Equal(t, "updated message", result.Message)
}
