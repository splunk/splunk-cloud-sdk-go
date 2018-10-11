// DO NOT EDIT
package action

type ActionIface interface {
	// GetActions get all actions
	GetActions() ([]Action, error)
	// CreateAction creates an action
	CreateAction(action Action) (*Action, error)
	// GetAction get an action by name
	GetAction(name string) (*Action, error)
	// TriggerAction triggers an action from a notification
	TriggerAction(name string, notification Notification) (*TriggerResponse, error)
	// UpdateAction updates and action by name
	UpdateAction(name string, action UpdateFields) (*Action, error)
	// DeleteAction deletes an action by name
	DeleteAction(name string) error
	// GetActionStatus returns an action's status by name
	GetActionStatus(name string, statusID string) (*Status, error)
}
