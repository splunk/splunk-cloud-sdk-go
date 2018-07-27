package stubbyintegration

import (
	"testing"

	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
)

// Stubby test for Create Action service endpoint
func TestCreateAction(t *testing.T) {
	actionData := model.Action{Name: "test10", Kind: "webhook", WebhookURL: "http://webhook.site/40fa145b-43d7-48f9-a391-aa7558042fa6", Message: "{{ .name }} is a {{ .species }}"}
	result, err := getClient(t).ActionService.CreateAction(actionData)
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, "be7ab21a-995c-4392-9834-66f4a2aec48a", result.ID)
}

// Stubby test for Post trigger action
func TestTriggerWebHookAction(t *testing.T) {
	actionNotificationData := model.ActionNotification{Kind: "rawJSON", UserID: "unused", Tenant: "tenantId", Payload: "{ \"name\": \"bean bag\", \"species\": \"cat\"}"}
	u, err := getClient(t).ActionService.TriggerAction("test10", actionNotificationData)
	assert.Empty(t, err)
	assert.Empty(t, u)
}


// Stubby test for GetAction service endpoint
func TestGetAction(t *testing.T) {
	result, err := getClient(t).ActionService.GetAction("test10")
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

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
	actionData := model.Action{Message: "updated message"}
	err := getClient(t).ActionService.UpdateAction("test10", actionData)
	assert.Empty(t, err)
}

