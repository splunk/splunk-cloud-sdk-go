// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package action

import (
	"strings"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

// action service url prefix
const servicePrefix = "action"
const serviceVersion = "v1beta1"

// Service - A service the receives incoming notifications and uses
// pre-defined templates to turn those notifications into meaningful actions
type Service services.BaseService

// NewService creates a new action service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

// GetActions get all actions
func (s *Service) GetActions() ([]Action, error) {
	url, err := s.Client.BuildURL(nil, servicePrefix, serviceVersion, "actions")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var results []Action
	err = util.ParseResponse(&results, response)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// CreateAction creates an action
func (s *Service) CreateAction(action Action) (*Action, error) {
	url, err := s.Client.BuildURL(nil, servicePrefix, serviceVersion, "actions")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: action})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Action
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAction get an action by name
func (s *Service) GetAction(name string) (*Action, error) {
	url, err := s.Client.BuildURL(nil, servicePrefix, serviceVersion, "actions", name)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Action
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// TriggerAction triggers an action from a notification
func (s *Service) TriggerAction(name string, notification Notification) (*TriggerResponse, error) {
	url, err := s.Client.BuildURL(nil, servicePrefix, serviceVersion, "actions", name)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: notification})
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
		return &TriggerResponse{StatusID: &parts[l-1], StatusURL: u}, nil
	}
	// If format doesn't match what we expect just return url
	return &TriggerResponse{StatusURL: u}, nil
}

// UpdateAction updates and action by name
func (s *Service) UpdateAction(name string, action UpdateFields) (*Action, error) {
	url, err := s.Client.BuildURL(nil, servicePrefix, serviceVersion, "actions", name)
	if err != nil {
		return nil, err
	}

	response, err := s.Client.Patch(services.RequestParams{URL: url, Body: action})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result Action
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteAction deletes an action by name
func (s *Service) DeleteAction(name string) error {
	url, err := s.Client.BuildURL(nil, servicePrefix, serviceVersion, "actions", name)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// GetActionStatus returns an action's status by name
func (s *Service) GetActionStatus(name string, statusID string) (*Status, error) {
	url, err := s.Client.BuildURL(nil, servicePrefix, serviceVersion, "actions", name, "status", statusID)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Status
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
