// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package streams

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

// streams service url prefix
const servicePrefix = "streams"
const serviceVersion = "v1"
const serviceCluster = "api"

// Service - A service that deals with pipelines
type Service services.BaseService

// NewService creates a new streams service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

// CompileDslToUpl creates a Upl Json from DSL
func (s *Service) CompileDslToUpl(dsl *DslCompilationRequest) (*UplPipeline, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", "compile-dsl")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: dsl})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result UplPipeline
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPipelineStatus gets status of pipelines from the underlying streaming system
func (s *Service) GetPipelineStatus(queryParams PipelineStatusQueryParams) (*PaginatedPipelineStatusResponse, error) {
	queryValues, err := convertToURLQueryValues(queryParams)
	if err != nil {
		return nil, err
	}

	url, err := s.Client.BuildURL(queryValues, serviceCluster, servicePrefix, serviceVersion, "pipelines", "status")
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
	var result PaginatedPipelineStatusResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPipelines gets all the pipelines
func (s *Service) GetPipelines(queryParams PipelineQueryParams) (*PaginatedPipelineResponse, error) {
	queryValues, err := convertToURLQueryValues(queryParams)
	if err != nil {
		return nil, err
	}

	url, err := s.Client.BuildURL(queryValues, serviceCluster, servicePrefix, serviceVersion, "pipelines")
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
	var result PaginatedPipelineResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePipeline creates a new pipeline
func (s *Service) CreatePipeline(pipeline *PipelineRequest) (*Pipeline, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: pipeline})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Pipeline
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ActivatePipeline activates an existing pipeline
func (s *Service) ActivatePipeline(ids []string) (AdditionalProperties, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", "activate")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: ActivatePipelineRequest{IDs: ids, SkipSavePoint: true}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result AdditionalProperties
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeactivatePipeline deactivates an existing pipeline
func (s *Service) DeactivatePipeline(ids []string) (AdditionalProperties, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", "deactivate")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: ActivatePipelineRequest{IDs: ids, SkipSavePoint: true}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result AdditionalProperties
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ReactivatePipeline reactivates an existing pipeline
func (s *Service) ReactivatePipeline(id string) (*PipelineReactivateResponse, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", id, "reactivate")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result PipelineReactivateResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPipeline gets an individual pipeline
func (s *Service) GetPipeline(id string) (*Pipeline, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", id)
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
	var result Pipeline
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePipeline updates an existing pipeline
func (s *Service) UpdatePipeline(id string, pipeline *PipelineRequest) (*Pipeline, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", id)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Put(services.RequestParams{URL: url, Body: pipeline})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Pipeline
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// MergePipelines merges two Upl pipelines into one Upl pipeline
func (s *Service) MergePipelines(mergeRequest *PipelinesMergeRequest) (*UplPipeline, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", "merge")
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
	var result UplPipeline
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeletePipeline deletes a pipeline
func (s *Service) DeletePipeline(id string) (*PipelineDeleteResponse, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", id)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result PipelineDeleteResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//Returns the input schema for a function in the pipeline
func (s *Service) GetInputSchema(request *GetInputSchemaRequest) (*Parameters, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", "input-schema")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: GetInputSchemaRequest{request.NodeUUID, request.TargetPortName, request.UplJSON}})

	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result *Parameters
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//Returns the output schema for the specified function in the pipeline.
func (s *Service) GetOutputSchema(request *GetOutputSchemaRequest) (*Parameters, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", "output-schema")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: GetOutputSchemaRequest{request.NodeUUID, request.SourcePortName, request.UplJSON}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result *Parameters
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//Returns all functions in JSON format in the registry
func (s *Service) GetRegistry(local url.Values) (*UplRegistry, error) {

	url, err := s.Client.BuildURL(local, serviceCluster, servicePrefix, serviceVersion, "pipelines", "registry")
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
	var result *UplRegistry
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//Get latest Preview session metrics
func (s *Service) GetLatestPreviewSessionMetrics(previewSessionID string) (*MetricsResponse, error) {

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "preview", previewSessionID, "metrics", "latest")

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
	var result *MetricsResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//Get latest Pipeline metrics
func (s *Service) GetLatestPipelineMetrics(pipelineID string) (*MetricsResponse, error) {

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", pipelineID, "metrics", "latest")
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
	var result *MetricsResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//Validate if the Streams JSON is valid
func (s *Service) ValidateUplResponse(request *ValidateRequest) (*ValidateResponse, error) {

	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "pipelines", "validate")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: ValidateRequest{request.Upl}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result *ValidateResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//Get all the available connectors
func (s *Service) GetConnectors() (*Connectors, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "connectors")
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
	var result Connectors
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil

}

//Get the connections for a specific connector
func (s *Service) GetConnections(connectorID string) (*Connections, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "connections", connectorID)
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
	var result *Connections
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StartPreviewSession starts a preview session for an existing pipeline
func (s *Service) StartPreviewSession(previewSession *PreviewSessionStartRequest) (*PreviewStartResponse, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "preview-session")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: previewSession})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result PreviewStartResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPreviewSession gets an individual pipeline
func (s *Service) GetPreviewSession(id string) (*PreviewState, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "preview", id)
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
	var result PreviewState
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeletePreviewSession stops a preview session
func (s *Service) DeletePreviewSession(id string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "preview-session", id)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetPreviewData gets preview data for a session
func (s *Service) GetPreviewData(id string) (*PreviewData, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "preview-data", id)
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
	var result PreviewData
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateTemplate creates a new template for a tenant
func (s *Service) CreateTemplate(previewSession *TemplateRequest) (*TemplateResponse, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "templates")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: previewSession})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result TemplateResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTemplates gets a list of latest templates
func (s *Service) GetTemplates() (*PaginatedTemplateResponse, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "templates")
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
	var result PaginatedTemplateResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTemplate gets an individual template by template id
func (s *Service) GetTemplate(id string) (*TemplateResponse, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "templates", id)
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
	var result TemplateResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateTemplate updates an existing template (requires all fields)
func (s *Service) UpdateTemplate(id string, template *TemplateRequest) (*TemplateResponse, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "templates", id)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Put(services.RequestParams{URL: url, Body: template})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result TemplateResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateTemplatePartially partially updates an existing template (able to send partial data)
func (s *Service) UpdateTemplatePartially(id string, template *PartialTemplateRequest) (*TemplateResponse, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "templates", id)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: url, Body: template})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result TemplateResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteTemplate deletes a template based on the provided template id
func (s *Service) DeleteTemplate(id string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "templates", id)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// Retrieve the full streams JSON of a group
func (s *Service) GetGroupByID(groupID string) (*GroupResponse, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "groups", groupID)
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
	var result GroupResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create and return the expanded version of a group
func (s *Service) CreateExpandedGroup(groupID string, groupExpandedRequest *GroupExpandRequest) (*UplPipeline, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "groups", groupID, "expand")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: groupExpandedRequest})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result UplPipeline
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// TODO: Change input parameters to take in generic struct
// Converts the Pipeline query parameters to url.Values type
func convertToURLQueryValues(queryParams interface{}) (url.Values, error) {
	jsonQueryParams, err := json.Marshal(queryParams)
	if err != nil {
		return nil, err
	}

	queryValues := url.Values{}
	interfaceMap := make(map[string]interface{})
	err = json.Unmarshal(jsonQueryParams, &interfaceMap)
	if err != nil {
		return nil, err
	}

	for key, value := range interfaceMap {
		queryValues.Set(key, fmt.Sprintf("%v", value))
	}
	return queryValues, err
}
