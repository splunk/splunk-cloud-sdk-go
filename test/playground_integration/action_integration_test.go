package playgroundintegration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test vars

func cleanupActions(t *testing.T) {
	client := getClient(t)
	result, err := client.ActionService.GetActions()
	assert.Nil(t, err)

	for _, item := range result {
		err = client.ActionService.DeleteAction(item.Name)
		assert.Nil(t, err)
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

// Test CreateAction
func TestCreateAction(t *testing.T) {
	// TODO
}

// Test GetAction
func TestGetAction(t *testing.T) {
	// TODO
}

// Test TriggerAction
func TestTriggerAction(t *testing.T) {
	// TODO
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
