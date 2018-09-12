// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"encoding/json"
	"fmt"
	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"net/url"
)

// streams service url prefix
const streamsServicePrefix = "streams"
const streamsServiceVersion = "v1"

// StreamsService - A service that deals with pipelines
type StreamsService service

// CompileDslToUpl creates a Upl Json from DSL
func (c *StreamsService) CompileDslToUpl(dsl *model.DslCompilationRequest) (*model.UplPipeline, error) {
	url, err := c.client.BuildURL(nil, streamsServicePrefix, streamsServiceVersion, "pipelines", "compile-dsl")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: dsl})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.UplPipeline
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPipelines gets all the pipelines
func (c *StreamsService) GetPipelines(queryParams model.PipelineQueryParams) (*model.PaginatedPipelineResponse, error) {
	queryValues, err := convertToURLQueryValues(queryParams)

	url, err := c.client.BuildURL(queryValues, streamsServicePrefix, streamsServiceVersion, "pipelines")
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
	var result model.PaginatedPipelineResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePipeline creates a new pipeline
func (c *StreamsService) CreatePipeline(pipeline *model.PipelineRequest) (*model.Pipeline, error) {
	url, err := c.client.BuildURL(nil, streamsServicePrefix, streamsServiceVersion, "pipelines")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: pipeline})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.Pipeline
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ActivatePipeline activates an existing pipeline
func (c *StreamsService) ActivatePipeline(ids []string) (model.AdditionalProperties, error) {
	url, err := c.client.BuildURL(nil, streamsServicePrefix, streamsServiceVersion, "pipelines", "activate")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: model.ActivatePipelineRequest{IDs: ids}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.AdditionalProperties
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeactivatePipeline deactivates an existing pipeline
func (c *StreamsService) DeactivatePipeline(ids []string) (model.AdditionalProperties, error) {
	url, err := c.client.BuildURL(nil, streamsServicePrefix, streamsServiceVersion, "pipelines", "deactivate")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: model.ActivatePipelineRequest{IDs: ids}})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.AdditionalProperties
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPipeline gets an individual pipeline
func (c *StreamsService) GetPipeline(id string) (*model.Pipeline, error) {
	url, err := c.client.BuildURL(nil, streamsServicePrefix, streamsServiceVersion, "pipelines", id)
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
	var result model.Pipeline
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePipeline updates an existing pipeline
func (c *StreamsService) UpdatePipeline(id string, pipeline *model.PipelineRequest) (*model.Pipeline, error) {
	url, err := c.client.BuildURL(nil, streamsServicePrefix, streamsServiceVersion, "pipelines", id)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Put(RequestParams{URL: url, Body: pipeline})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.Pipeline
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeletePipeline deletes a pipeline
func (c *StreamsService) DeletePipeline(id string) (*model.PipelineDeleteResponse, error) {
	url, err := c.client.BuildURL(nil, streamsServicePrefix, streamsServiceVersion, "pipelines", id)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.PipelineDeleteResponse
	err = util.ParseResponse(&result, response)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// TODO: Change input parameters to take in generic struct
// Converts the Pipeline query parameters to url.Values type
func convertToURLQueryValues(queryParams model.PipelineQueryParams) (url.Values, error) {
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
