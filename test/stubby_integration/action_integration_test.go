package stubbyintegration

import (
	"testing"

	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
)

// Stubby test for GetCollectionStats() kvstore service endpoint
func TestCreateAction(t *testing.T) {
	actionData := model.Action{Name: "test10", Kind: "webhook", WebhookURL: "http://webhook.site/40fa145b-43d7-48f9-a391-aa7558042fa6", Message: "{{ .name }} is a {{ .species }}"}
	result, err := getClient(t).ActionService.CreateAction(actionData)
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, "be7ab21a-995c-4392-9834-66f4a2aec48a", result.ID)
}


// Stubby test for GetCollectionStats() kvstore service endpoint
func TestGetAction(t *testing.T) {
	result, err := getClient(t).ActionService.GetAction("test10")
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

}

