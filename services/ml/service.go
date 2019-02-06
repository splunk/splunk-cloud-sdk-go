// Copyright © 2019 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package ml

import (
	"encoding/json"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const servicePrefix = "ml"
const serviceVersion = "v1"
const serviceCluster = "api"

// Service talks to the Splunk Cloud machine learning service
type Service services.BaseService

// NewService creates a new machine learning service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

//CreateWorkflow Create a workflow configuration
func (s *Service) CreateWorkflow(workflow Workflow) (*Workflow, error) {
	var createdWorkflow Workflow

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows")
	if err != nil {
		return &createdWorkflow, err
	}
	jsonBytes, err := json.Marshal(workflow)
	if err != nil {
		return &createdWorkflow, err
	}

	response, err := s.Client.Post(services.RequestParams{URL: url, Body: jsonBytes})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return &createdWorkflow, err
	}

	// TODO: something wrong with the umarshalling here, the model isn't quite right
	//bodyBytes, _ := ioutil.ReadAll(response.Body)
	//bodyString := string(bodyBytes)
	//fmt.Println(bodyString)

	err = util.ParseResponse(createdWorkflow, response)

	return &createdWorkflow, err
}

//CreateWorkflowBuild Create a workflow build
func (s *Service) CreateWorkflowBuild(id string, workflowBuild WorkflowBuild) (*WorkflowBuild, error) {
	var createdWorkflowBuild WorkflowBuild

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows", id, "builds")
	if err != nil {
		return &createdWorkflowBuild, err
	}
	jsonBytes, err := json.Marshal(workflowBuild)
	if err != nil {
		return &createdWorkflowBuild, err
	}

	response, err := s.Client.Post(services.RequestParams{URL: url, Body: jsonBytes})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return &createdWorkflowBuild, err
	}

	err = util.ParseResponse(createdWorkflowBuild, response)

	return &createdWorkflowBuild, nil
}

//CreateWorkflowRun Create a workflow Run
func (s *Service) CreateWorkflowRun(id string, buildID string, workflowRun WorkflowRun) (*WorkflowRun, error) {
	var createdWorkflowRun WorkflowRun

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows", id, "builds", buildID, "runs")
	if err != nil {
		return &createdWorkflowRun, err
	}
	jsonBytes, err := json.Marshal(workflowRun)
	if err != nil {
		return &createdWorkflowRun, err
	}

	response, err := s.Client.Post(services.RequestParams{URL: url, Body: jsonBytes})
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return &createdWorkflowRun, err
	}

	err = util.ParseResponse(createdWorkflowRun, response)

	return &createdWorkflowRun, nil
}

//DeleteWorkflow Delete a workflow configuration
func (s *Service) DeleteWorkflow(id string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows", id)
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

//DeleteWorkflowBuild Delete workflow build
func (s *Service) DeleteWorkflowBuild(id string, buildID string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows", id, "builds", buildID)
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

//DeleteWorkflowRun Delete a workflow run
func (s *Service) DeleteWorkflowRun(id string, buildID string, runID string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows", id, "builds", buildID, "runs", runID)
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

//GetWorkflow Get a workflow configuration
func (s *Service) GetWorkflow(id string) (*Workflow, error) {
	var workflow Workflow

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows", id)
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
	err = util.ParseResponse(&workflow, response)

	return &workflow, nil
}

//GetWorkflowBuild Get status of a workflow build
func (s *Service) GetWorkflowBuild(id string, buildID string) (*WorkflowBuild, error) {
	var workflowBuild WorkflowBuild

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows", id, "builds", buildID)
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
	err = util.ParseResponse(&workflowBuild, response)

	return &workflowBuild, nil
}

//GetWorkflowBuilds Get list of workflow builds
func (s *Service) GetWorkflowBuilds(id string) ([]WorkflowBuild, error) {
	var workflowBuilds []WorkflowBuild

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows", id, "builds")
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
	err = util.ParseResponse(&workflowBuilds, response)

	return workflowBuilds, nil
}

//GetWorkflowRun Get status of a workflow run
func (s *Service) GetWorkflowRun(id string, buildID string, runID string) (*WorkflowRun, error) {
	var workflowRun WorkflowRun

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows", id, "builds", buildID, "runs", runID)
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
	err = util.ParseResponse(&workflowRun, response)

	return &workflowRun, nil
}

//GetWorkflowRuns Get list of workflow runs
func (s *Service) GetWorkflowRuns(id string, buildID string) ([]WorkflowRun, error) {
	var workflowRuns []WorkflowRun

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows", id, "builds", buildID, "runs")
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
	err = util.ParseResponse(&workflowRuns, response)

	return workflowRuns, nil
}

//GetWorkflows Get the list of workflow configurations
func (s *Service) GetWorkflows() ([]WorkflowsGetResponse, error) {
	var workflows []WorkflowsGetResponse

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "workflows")
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
	err = util.ParseResponse(&workflows, response)

	return workflows, nil
}
