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
	webhookMsg   = "Webhook Message"
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
	webhookAction     = model.NewWebhookAction(webhookActionName, webhookUrl, webhookMsg)
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

// Test TriggerAction
func TestTriggerAction(t *testing.T) {
	client := getClient(t)
	defer cleanupAction(client, emailAction.Name)
	_, err := client.ActionService.CreateAction(*emailAction)
	require.Nil(t, err)
	err = client.ActionService.TriggerAction(emailActionName,
		model.ActionNotification{
			Kind:    model.RawJSONPayloadKind,
			Tenant:  testutils.TestTenantID,
			UserID:  "sdk_tester",
			Payload: "some data",
		})
	require.Nil(t, err)
}

// Test UpdateAction
func TestUpdateAction(t *testing.T) {
	// TODO
}

// Test DeleteAction
func TestDeleteAction(t *testing.T) {
	// TODO
}

// Test GetActionStatus
func TestGetActionStatus(t *testing.T) {
	// TODO
}
