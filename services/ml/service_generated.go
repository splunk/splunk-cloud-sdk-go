/*
 * Copyright © 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * Machine Learning Service (ML API)
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2beta1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package ml

import (
	"net/http"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const serviceCluster = "api"

type Service services.BaseService

// NewService creates a new ml service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

/*
	CreateWorkflow - Creates a workflow configuration.
	Parameters:
		workflow: Workflow configuration to be created.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateWorkflow(workflow Workflow, resp ...*http.Response) (*Workflow, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: workflow})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb Workflow
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreateWorkflowBuild - Creates a workflow build.
	Parameters:
		id: The workflow ID.
		workflowBuild: Input data used to build the workflow.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateWorkflowBuild(id string, workflowBuild WorkflowBuild, resp ...*http.Response) (*WorkflowBuild, error) {
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: workflowBuild})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowBuild
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreateWorkflowDeployment - Creates a workflow deployment.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		workflowDeployment: Input data used to build the workflow deployment.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateWorkflowDeployment(id string, buildId string, workflowDeployment WorkflowDeployment, resp ...*http.Response) (*WorkflowDeployment, error) {
	pp := struct {
		Id      string
		BuildId string
	}{
		Id:      id,
		BuildId: buildId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/deployments`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: workflowDeployment})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowDeployment
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreateWorkflowInference - Creates a workflow inference request.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		deploymentId: The workflow deployment ID.
		workflowInference: Input data to the inference request.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateWorkflowInference(id string, buildId string, deploymentId string, workflowInference WorkflowInference, resp ...*http.Response) (*WorkflowInference, error) {
	pp := struct {
		Id           string
		BuildId      string
		DeploymentId string
	}{
		Id:           id,
		BuildId:      buildId,
		DeploymentId: deploymentId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/deployments/{{.DeploymentId}}/inference`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: workflowInference})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowInference
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreateWorkflowRun - Creates a workflow run.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		workflowRun: Input data used to build the workflow.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateWorkflowRun(id string, buildId string, workflowRun WorkflowRun, resp ...*http.Response) (*WorkflowRun, error) {
	pp := struct {
		Id      string
		BuildId string
	}{
		Id:      id,
		BuildId: buildId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/runs`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: workflowRun})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowRun
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	DeleteWorkflow - Removes a workflow configuration.
	Parameters:
		id: The workflow ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteWorkflow(id string, resp ...*http.Response) error {
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	DeleteWorkflowBuild - Removes a workflow build.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteWorkflowBuild(id string, buildId string, resp ...*http.Response) error {
	pp := struct {
		Id      string
		BuildId string
	}{
		Id:      id,
		BuildId: buildId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	DeleteWorkflowDeployment - Removes a workflow deployment.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		deploymentId: The workflow deployment ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteWorkflowDeployment(id string, buildId string, deploymentId string, resp ...*http.Response) error {
	pp := struct {
		Id           string
		BuildId      string
		DeploymentId string
	}{
		Id:           id,
		BuildId:      buildId,
		DeploymentId: deploymentId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/deployments/{{.DeploymentId}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	DeleteWorkflowRun - Removes a workflow run.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		runId: The workflow run ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteWorkflowRun(id string, buildId string, runId string, resp ...*http.Response) error {
	pp := struct {
		Id      string
		BuildId string
		RunId   string
	}{
		Id:      id,
		BuildId: buildId,
		RunId:   runId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/runs/{{.RunId}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	return err
}

/*
	GetWorkflow - Returns a workflow configuration.
	Parameters:
		id: The workflow ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetWorkflow(id string, resp ...*http.Response) (*Workflow, error) {
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb Workflow
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetWorkflowBuild - Returns the status of a workflow build.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetWorkflowBuild(id string, buildId string, resp ...*http.Response) (*WorkflowBuild, error) {
	pp := struct {
		Id      string
		BuildId string
	}{
		Id:      id,
		BuildId: buildId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowBuild
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetWorkflowBuildError - Returns a list of workflow errors.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetWorkflowBuildError(id string, buildId string, resp ...*http.Response) (*WorkflowBuildError, error) {
	pp := struct {
		Id      string
		BuildId string
	}{
		Id:      id,
		BuildId: buildId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/errors`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowBuildError
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetWorkflowBuildLog - Returns the logs from a workflow build.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetWorkflowBuildLog(id string, buildId string, resp ...*http.Response) (*WorkflowBuildLog, error) {
	pp := struct {
		Id      string
		BuildId string
	}{
		Id:      id,
		BuildId: buildId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/logs`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowBuildLog
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetWorkflowDeployment - Returns the status of a workflow deployment.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		deploymentId: The workflow deployment ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetWorkflowDeployment(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeployment, error) {
	pp := struct {
		Id           string
		BuildId      string
		DeploymentId string
	}{
		Id:           id,
		BuildId:      buildId,
		DeploymentId: deploymentId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/deployments/{{.DeploymentId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowDeployment
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetWorkflowDeploymentError - Returns a list of workflow deployment errors.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		deploymentId: The workflow deployment ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetWorkflowDeploymentError(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeploymentError, error) {
	pp := struct {
		Id           string
		BuildId      string
		DeploymentId string
	}{
		Id:           id,
		BuildId:      buildId,
		DeploymentId: deploymentId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/deployments/{{.DeploymentId}}/errors`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowDeploymentError
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetWorkflowDeploymentLog - Returns the logs from a workflow deployment.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		deploymentId: The workflow deployment ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetWorkflowDeploymentLog(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeploymentLog, error) {
	pp := struct {
		Id           string
		BuildId      string
		DeploymentId string
	}{
		Id:           id,
		BuildId:      buildId,
		DeploymentId: deploymentId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/deployments/{{.DeploymentId}}/logs`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowDeploymentLog
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetWorkflowRun - Returns the status of a workflow run.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		runId: The workflow run ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetWorkflowRun(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRun, error) {
	pp := struct {
		Id      string
		BuildId string
		RunId   string
	}{
		Id:      id,
		BuildId: buildId,
		RunId:   runId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/runs/{{.RunId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowRun
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetWorkflowRunError - Returns the errors for a workflow run.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		runId: The workflow run ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetWorkflowRunError(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRunError, error) {
	pp := struct {
		Id      string
		BuildId string
		RunId   string
	}{
		Id:      id,
		BuildId: buildId,
		RunId:   runId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/runs/{{.RunId}}/errors`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowRunError
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetWorkflowRunLog - Returns the logs for a workflow run.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		runId: The workflow run ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetWorkflowRunLog(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRunLog, error) {
	pp := struct {
		Id      string
		BuildId string
		RunId   string
	}{
		Id:      id,
		BuildId: buildId,
		RunId:   runId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/runs/{{.RunId}}/logs`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb WorkflowRunLog
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListWorkflowBuilds - Returns a list of workflow builds.
	Parameters:
		id: The workflow ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListWorkflowBuilds(id string, resp ...*http.Response) ([]WorkflowBuild, error) {
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []WorkflowBuild
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListWorkflowDeployments - Returns a list of workflow deployments.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListWorkflowDeployments(id string, buildId string, resp ...*http.Response) ([]WorkflowDeployment, error) {
	pp := struct {
		Id      string
		BuildId string
	}{
		Id:      id,
		BuildId: buildId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/deployments`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []WorkflowDeployment
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListWorkflowRuns - Returns a list of workflow runs.
	Parameters:
		id: The workflow ID.
		buildId: The workflow build ID.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListWorkflowRuns(id string, buildId string, resp ...*http.Response) ([]WorkflowRun, error) {
	pp := struct {
		Id      string
		BuildId string
	}{
		Id:      id,
		BuildId: buildId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows/{{.Id}}/builds/{{.BuildId}}/runs`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []WorkflowRun
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListWorkflows - Returns a list of workflow configurations.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListWorkflows(resp ...*http.Response) ([]WorkflowsGetResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ml/v2beta1/workflows`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	// populate input *http.Response if provided
	if len(resp) > 0 && resp[0] != nil {
		*resp[0] = *response
	}
	if err != nil {
		return nil, err
	}
	var rb []WorkflowsGetResponse
	err = util.ParseResponse(&rb, response)
	return rb, err
}