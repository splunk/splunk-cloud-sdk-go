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
)

// Test consts
const (
	htmlPart     = "<html><h1>The HTML</h1></html>"
	subjectPArt  = "The Subject"
	textPart     = "The Text"
	templateName = "template1000"
	snsTopic     = "myTopic"
	snsMsg       = "SNS Message"
	webhookUrl   = "https://locahost:9999/test"
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
	webhookAction := model.NewWebhookAction(webhookActionName, webhookUrl, webhookMsg)
	defer cleanupAction(client, webhookAction.Name)
	_, err := client.ActionService.CreateAction(*webhookAction)
	require.Nil(t, err)
	action, err := client.ActionService.GetAction(webhookAction.Name)
	assert.EqualValues(t, action, webhookAction)
}

// Test TriggerAction
func TestTriggerAction(t *testing.T) {
	client := getClient(t)
	webhookActionName := fmt.Sprintf("w-trig-%d", timeSec)
	webhookAction := model.NewWebhookAction(webhookActionName, webhookUrl, webhookMsg)
	defer cleanupAction(client, webhookAction.Name)
	action, err := client.ActionService.CreateAction(*webhookAction)
	require.Nil(t, err)
	require.NotNil(t, action)
	url, err := client.ActionService.TriggerAction(webhookAction.Name,
		model.ActionNotification{
			Kind:    model.RawJSONPayloadKind,
			Tenant:  testutils.TestTenantID,
			Payload: webhookPayload,
		})
	assert.Nil(t, err)
	assert.NotEmpty(t, url)
}

// Test UpdateAction
func TestUpdateAction(t *testing.T) {
	client := getClient(t)
	emailActionName := fmt.Sprintf("e-up-%d", timeSec)
	emailAction := model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
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
	emailActionName := fmt.Sprintf("e-del-%d", timeSec)
	emailAction := model.NewEmailAction(emailActionName, htmlPart, subjectPArt, textPart, templateName, addresses)
	_, err := client.ActionService.CreateAction(*emailAction)
	require.Nil(t, err)
	err = client.ActionService.DeleteAction(emailActionName)
	assert.Nil(t, err)
}

// Test GetActionStatus
func TestGetActionStatus(t *testing.T) {
	client := getClient(t)
	webhookActionName := fmt.Sprintf("w-stat-%d", timeSec)
	webhookAction := model.NewWebhookAction(webhookActionName, webhookUrl, webhookMsg)
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
