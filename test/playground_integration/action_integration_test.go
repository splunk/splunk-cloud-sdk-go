package playgroundintegration

import (
	"fmt"
	"testing"
	"time"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/splunk/ssc-client-go/util"
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
	emailActionName   = fmt.Sprintf("eact-%d", timeSec)
	snsActionName     = fmt.Sprintf("sact-%d", timeSec)
	webhookActionName = fmt.Sprintf("wact-%d", timeSec)
	addresses         = []string{"test1@splunk.com", "test2@splunk.com"}
	emailAction       = model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
	snsAction         = model.NewSNSAction(snsActionName, snsTopic, snsMsg)
	webhookAction     = model.NewWebhookAction(webhookActionName, webhookURL, webhookMsg)
	webhookPayload    = &map[string]interface{}{"name": "bean bag", "species": "cat"}
)

func cleanupActions(client *service.Client) {
	cleanupAction(client, emailActionName)
	cleanupAction(client, snsActionName)
	cleanupAction(client, webhookActionName)
}

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
	defer cleanupAction(client, emailAction.Name)
	_, err := client.ActionService.CreateAction(*emailAction)
	require.Nil(t, err)
	action, err := client.ActionService.GetAction(emailAction.Name)
	assert.EqualValues(t, action, emailAction)
}

// Test CreateAction / GetAction for SNSAction
func TestGetCreateActionSNS(t *testing.T) {
	client := getClient(t)
	defer cleanupAction(client, snsAction.Name)
	_, err := client.ActionService.CreateAction(*snsAction)
	require.Nil(t, err)
	action, err := client.ActionService.GetAction(snsAction.Name)
	assert.EqualValues(t, action, snsAction)
}

// Test CreateAction / GetAction for WebhookAction
func TestGetCreateActionWebhook(t *testing.T) {
	client := getClient(t)
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
	defer cleanupAction(client, webhookAction.Name)

	_, err := client.ActionService.CreateAction(*webhookAction)
	require.Nil(t, err)

	_, err = invalidClient.ActionService.CreateAction(*emailAction)
	validateUnauthenticatedActionError(t , err)

	_, err = invalidClient.ActionService.GetAction(webhookAction.Name)
	validateUnauthenticatedActionError(t , err)

	_, err = invalidClient.ActionService.GetActions()
	validateUnauthenticatedActionError(t , err)

	_, err = invalidClient.ActionService.TriggerAction(webhookAction.Name,
		model.ActionNotification{
			Kind:    model.RawJSONPayloadKind,
			Tenant:  testutils.TestTenantID,
			Payload: webhookPayload,
		})
	validateUnauthenticatedActionError(t , err)

	_, err = invalidClient.ActionService.UpdateAction(webhookActionName, model.Action{TextPart: "updated email text"})
	validateUnauthenticatedActionError(t , err)

	err = invalidClient.ActionService.DeleteAction(webhookActionName)
	validateUnauthenticatedActionError(t , err)
}

// Trigger action with invalid fields
func TestTriggerActionFailInvalidFields(t *testing.T) {
	client := getClient(t)
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
	defer cleanupAction(client, emailAction.Name)
	_, err := client.ActionService.CreateAction(*emailAction)
	require.Nil(t, err)
	result, err := client.ActionService.UpdateAction(emailActionName, model.Action{TextPart: "updated email text"})
	assert.NotEmpty(t, result)
	assert.Nil(t, err)
}

// Test DeleteAction
func TestDeleteAction(t *testing.T) {
	client := getClient(t)
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

	_, err = client.ActionService.UpdateAction("Action123", model.Action{TextPart: "updated email text"})
	validateNotFoundActionError(t, err)

	err = client.ActionService.DeleteAction("Action123")
	validateNotFoundActionError(t, err)
}

// Test GetActionStatus
func TestGetActionStatus(t *testing.T) {
	// TODO
}
