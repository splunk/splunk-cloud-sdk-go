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

// GetPipelines gets all the pipelines
func (s *Service) GetPipelines(queryParams PipelineQueryParams) (*PaginatedPipelineResponse, error) {
	queryValues, err := convertToURLQueryValues(queryParams)

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

// TODO: Change input parameters to take in generic struct
// Converts the Pipeline query parameters to url.Values type
func convertToURLQueryValues(queryParams PipelineQueryParams) (url.Values, error) {
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
