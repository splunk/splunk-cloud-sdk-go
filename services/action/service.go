// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package action

import (
	"github.com/splunk/splunk-cloud-sdk-go/services"
)

// action service url prefix
const actionServicePrefix = "action"
const actionServiceVersion = "v1beta1"

// Service - A service the receives incoming notifications and uses
// pre-defined templates to turn those notifications into meaningful actions
type Service services.BaseService

// NewService creates a new action service with client
func NewService(client *services.Client) *Service {
	return &Service{Client: client}
}

/*
// GetActions get all actions
func (c *ActionService) GetActions() ([]model.Action, error) {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(RequestParams{URL: url})
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
	response, err := c.client.Post(RequestParams{URL: url, Body: action})
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
	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *ActionService) TriggerAction(name string, notification model.ActionNotification) (*model.ActionTriggerResponse, error) {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions", name)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: notification})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	u, err := response.Location()
	parts := strings.Split(u.Path, "/")
	l := len(parts)
	if l >= 2 && parts[l-2] == "status" {
		// Parse the Location url for the user looking for .../status/{statusid}
		// at the end of the URL,
		return &model.ActionTriggerResponse{StatusID: &parts[l-1], StatusURL: u}, nil
	}
	// If format doesn't match what we expect just return url
	return &model.ActionTriggerResponse{StatusURL: u}, nil
}

// UpdateAction updates and action by name
func (c *ActionService) UpdateAction(name string, action model.ActionUpdateFields) (*model.Action, error) {
	url, err := c.client.BuildURL(nil, actionServicePrefix, actionServiceVersion, "actions", name)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Patch(RequestParams{URL: url, Body: action})
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
	response, err := c.client.Delete(RequestParams{URL: url})
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
	response, err := c.client.Get(RequestParams{URL: url})
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
*/
