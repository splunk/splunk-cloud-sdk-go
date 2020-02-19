package action

import (
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/auth"
	model "github.com/splunk/splunk-cloud-sdk-go/services/action"
)

// PostEvents Sends events.
func TriggerActionOverride(actionName string, body model.TriggerEvent) (*model.TriggerResponse, error) {

	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	return client.ActionService.TriggerActionWithStatus(actionName, body)
}
