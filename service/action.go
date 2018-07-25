package service

import (
	"fmt"
	"net/http"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

// action service url prefix
const actionServicePrefix = "action"
const actionServiceVersion = "v1"

// ActionService talks to the SSC action service
type ActionService service

// GetActions get all actions
func (c *ActionService) GetActions() ([]interface{}, error) {
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
	var results []interface{}
	err = util.ParseResponse(&results, response)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// CreateAction creates an action
func (c *ActionService) CreateAction(action interface{}) error {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions")
	if err != nil {
		return err
	}
	response, err := c.client.Post(url, action)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	// TODO: parse response/header
	return nil
}

// GetAction get an action by name
func (c *ActionService) GetAction(name string) (interface{}, error) {
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
	var base model.ActionBase
	err = util.ParseResponse(&base, response)
	if err != nil {
		return nil, err
	}
	return UnmarshallAction(base.Kind, response)
}

// UnmarshallAction parses the response into the appropriate Action struct
func UnmarshallAction(kind model.ActionKind, response *http.Response) (interface{}, error) {
	switch kind {
	case model.EmailKind:
		var email model.EmailAction
		err := util.ParseResponse(&email, response)
		if err != nil {
			return nil, err
		}
		return email, nil
	case model.SNSKind:
		var sns model.SNSAction
		err := util.ParseResponse(&sns, response)
		if err != nil {
			return nil, err
		}
		return sns, nil
	case model.WebhookKind:
		var webhook model.EmailAction
		err := util.ParseResponse(&webhook, response)
		if err != nil {
			return nil, err
		}
		return webhook, nil
	default:
		return nil, fmt.Errorf("unrecognized ActionKind: %s", kind)
	}
}

// TriggerAction triggers an action from a notification
func (c *ActionService) TriggerAction(name string, notification model.ActionNotification) error {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions", name)
	if err != nil {
		return err
	}
	response, err := c.client.Post(url, notification)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	//
	return nil
}

// UpdateAction triggers an action from a notification
func (c *ActionService) UpdateAction(name string, action interface{}) error {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions", name)
	if err != nil {
		return err
	}
	response, err := c.client.Patch(url, action)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
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
