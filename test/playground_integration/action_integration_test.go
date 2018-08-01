package playgroundintegration

import (
	"fmt"
	"testing"
	"time"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/splunk/ssc-client-go/util"
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
	webhookURL   = "https://locahost:9999/test"
	webhookMsg   = "{{ .name }} is a {{ .species }}"
	actionUserID = "sdk_tester"
)

// Test vars
var (
	timeSec           = time.Now().Unix()
	snsActionName     = fmt.Sprintf("sact-%d", timeSec)
	webhookActionName = fmt.Sprintf("wact-%d", timeSec)
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
	assert.NotEmpty(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).Status)
	assert.Equal(t, "401 Unauthorized", err.(*util.HTTPError).Message)
}

func validateNotFoundActionError(t *testing.T, err error) {
	assert.NotEmpty(t, err)
	assert.Equal(t, 404, err.(*util.HTTPError).Status)
	assert.Equal(t, "404 Not Found", err.(*util.HTTPError).Message)
}

// Test GetActions
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
	emailActionName := fmt.Sprintf("e-cr-%d", timeSec)
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
	snsActionName := fmt.Sprintf("s-cr-%d", timeSec)
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
	webhookActionName := fmt.Sprintf("w-cr-%d", timeSec)
	webhookAction := model.NewWebhookAction(webhookActionName, webhookURL, webhookMsg)
	defer cleanupAction(client, webhookAction.Name)
	_, err := client.ActionService.CreateAction(*webhookAction)
	require.Nil(t, err)
	action, err := client.ActionService.GetAction(webhookAction.Name)
	assert.EqualValues(t, action, webhookAction)
}

// Get Non-Existent Action
func TestCreateActionFailInvalidAction(t *testing.T) {
	client := getClient(t)
	// Get Invalid Action
	_, err := client.ActionService.GetAction("Dontexist")

	assert.NotEmpty(t, err)
	assert.Equal(t, 404, err.(*util.HTTPError).Status)
	assert.Equal(t, "404 Not Found", err.(*util.HTTPError).Message)
}

// Create Existing action
func TestCreateActionFailExistingAction(t *testing.T) {
	client := getClient(t)
	emailActionName := fmt.Sprintf("e-confl-%d", timeSec)
	emailAction := model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
	defer cleanupAction(client, emailAction.Name)
	_, err := client.ActionService.CreateAction(*emailAction)
	require.Nil(t, err)
	action, err := client.ActionService.GetAction(emailAction.Name)
	assert.EqualValues(t, action, emailAction)

	_, err = client.ActionService.CreateAction(*emailAction)
	assert.NotEmpty(t, err)
	assert.Equal(t, 409, err.(*util.HTTPError).Status)
	assert.Equal(t, "409 Conflict", err.(*util.HTTPError).Message)
}

// Access action endpoints using an Unauthenticated client
func TestActionFailUnauthenticatedClient(t *testing.T) {
	invalidClient := getInvalidClient(t)
	client := getClient(t)
	webhookActionName := fmt.Sprintf("w-unauth-%d", timeSec)
	webhookAction := model.NewWebhookAction(webhookActionName, webhookURL, webhookMsg)
	defer cleanupAction(client, webhookAction.Name)

	_, err := client.ActionService.CreateAction(*webhookAction)
	require.Nil(t, err)

	emailActionName := fmt.Sprintf("e-unauth-%d", timeSec)
	emailAction := model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
	// This shouldn't be needed since the CreateAction should fail for 401:
	// defer cleanupAction(client, emailAction.Name)

	_, err = invalidClient.ActionService.CreateAction(*emailAction)
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.GetAction(webhookAction.Name)
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.GetActions()
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.TriggerAction(webhookAction.Name,
		model.ActionNotification{
			Kind:    model.RawJSONPayloadKind,
			Tenant:  testutils.TestTenantID,
			Payload: webhookPayload,
		})
	validateUnauthenticatedActionError(t, err)

	_, err = invalidClient.ActionService.UpdateAction(webhookActionName, model.ActionUpdateFields{TextPart: "updated email text"})
	validateUnauthenticatedActionError(t, err)

	err = invalidClient.ActionService.DeleteAction(webhookActionName)
	validateUnauthenticatedActionError(t, err)
}

// Trigger action with invalid fields
func TestTriggerActionFailInvalidFields(t *testing.T) {
	client := getClient(t)
	webhookActionName := fmt.Sprintf("w-unproc-%d", timeSec)
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

	assert.NotEmpty(t, err)
	assert.Equal(t, 422, err.(*util.HTTPError).Status)
	assert.Equal(t, "422 Unprocessable Entity", err.(*util.HTTPError).Message)
}

// Test UpdateAction
func TestUpdateAction(t *testing.T) {
	client := getClient(t)
	emailActionName := fmt.Sprintf("e-up-%d", timeSec)
	emailAction := model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
	defer cleanupAction(client, emailAction.Name)
	_, err := client.ActionService.CreateAction(*emailAction)
	require.Nil(t, err)
	const newText = "updated email text"
	result, err := client.ActionService.UpdateAction(emailActionName, model.ActionUpdateFields{TextPart: newText})
	assert.Equal(t, result.TextPart, newText)
	assert.Nil(t, err)
}

// Test DeleteAction
func TestDeleteAction(t *testing.T) {
	client := getClient(t)
	emailActionName := fmt.Sprintf("e-del-%d", timeSec)
	emailAction := model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
	_, err := client.ActionService.CreateAction(*emailAction)
	require.Nil(t, err)
	err = client.ActionService.DeleteAction(emailActionName)
	assert.Nil(t, err)
}

// Access action endpoints with a non-existent Action
func TestActionFailNotFoundAction(t *testing.T) {
	client := getClient(t)
	// Get Invalid Action
	_, err := client.ActionService.GetAction("Action123")
	validateNotFoundActionError(t, err)

	_, err = client.ActionService.UpdateAction("Action123", model.ActionUpdateFields{TextPart: "updated email text"})
	validateNotFoundActionError(t, err)

	err = client.ActionService.DeleteAction("Action123")
	validateNotFoundActionError(t, err)
}

// Test GetActionStatus
func TestGetActionStatus(t *testing.T) {
	client := getClient(t)
	webhookActionName := fmt.Sprintf("w-stat-%d", timeSec)
	webhookAction := model.NewWebhookAction(webhookActionName, webhookURL, webhookMsg)
	defer cleanupAction(client, webhookAction.Name)
	action, err := client.ActionService.CreateAction(*webhookAction)
	require.Nil(t, err)
	require.NotNil(t, action)
	resp, err := client.ActionService.TriggerAction(webhookAction.Name,
		model.ActionNotification{
			Kind:    model.RawJSONPayloadKind,
			Tenant:  testutils.TestTenantID,
			Payload: webhookPayload,
		})
	require.Nil(t, err)
	require.NotNil(t, resp.StatusID)

	stat, err := client.ActionService.GetActionStatus(webhookAction.Name, *resp.StatusID)
	assert.Nil(t, err)
	assert.NotNil(t, stat)
}
