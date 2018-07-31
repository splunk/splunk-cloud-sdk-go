package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
	"net/url"
)

// action service url prefix
const actionServicePrefix = "action"
const actionServiceVersion = "v1"

// ActionService talks to the SSC action service
type ActionService service

// GetActions get all actions
func (c *ActionService) GetActions() ([]model.Action, error) {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(url)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var results []model.Action
	err = util.ParseResponse(&results, response)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// CreateAction creates an action
func (c *ActionService) CreateAction(action model.Action) (*model.Action, error) {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(url, action)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.Action
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAction get an action by name
func (c *ActionService) GetAction(name string) (*model.Action, error) {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions", name)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(url)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.Action
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// TriggerAction triggers an action from a notification
func (c *ActionService) TriggerAction(name string, notification model.ActionNotification) (*url.URL, error) {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions", name)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(url, notification)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	u, err := response.Location()
	return u, nil
}

// UpdateAction updates and action by name
func (c *ActionService) UpdateAction(name string, action model.Action) (*model.Action, error) {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions", name)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Patch(url, action)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result model.Action
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAction deletes an action by name
func (c *ActionService) DeleteAction(name string) error {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions", name)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(url)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// GetActionStatus returns an action's status by name
func (c *ActionService) GetActionStatus(name string, statusID string) (*model.ActionStatus, error) {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions", name, "status", statusID)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(url)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.ActionStatus
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
