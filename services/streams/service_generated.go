/*
 * Copyright © 2020 Splunk, Inc.
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
 * Data Stream Processing REST API
 *
 * Use the Streams service to perform create, read, update, and delete (CRUD) operations on your data pipeline. The Streams service also has metrics and preview session endpoints and gives you full control over your data pipeline.
 *
 * API version: v3beta1.1 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package streams

import (
	"net/http"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const serviceCluster = "api"

type Service services.BaseService

// NewService creates a new streams service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

/*
	ActivatePipeline - Activates an existing pipeline.
	Parameters:
		id: Pipeline ID
		activatePipelineRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ActivatePipeline(id string, activatePipelineRequest ActivatePipelineRequest, resp ...*http.Response) (*Response, error) {
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/{{.Id}}/activate`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: activatePipelineRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb Response
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	Compile - Compiles SPL2 and returns streams JSON.
	Parameters:
		splCompileRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) Compile(splCompileRequest SplCompileRequest, resp ...*http.Response) (*Pipeline, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/compile`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: splCompileRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb Pipeline
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreateConnection - Create a new DSP connection.
	Parameters:
		connectionRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateConnection(connectionRequest ConnectionRequest, resp ...*http.Response) (*ConnectionSaveResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/connections`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: connectionRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb ConnectionSaveResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreatePipeline - Creates a pipeline.
	Parameters:
		pipelineRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreatePipeline(pipelineRequest PipelineRequest, resp ...*http.Response) (*PipelineResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: pipelineRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PipelineResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreateTemplate - Creates a template for a tenant.
	Parameters:
		templateRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateTemplate(templateRequest TemplateRequest, resp ...*http.Response) (*TemplateResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/templates`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: templateRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb TemplateResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	DeactivatePipeline - Deactivates an existing pipeline.
	Parameters:
		id: Pipeline ID
		deactivatePipelineRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeactivatePipeline(id string, deactivatePipelineRequest DeactivatePipelineRequest, resp ...*http.Response) (*Response, error) {
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/{{.Id}}/deactivate`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: deactivatePipelineRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb Response
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	Decompile - Decompiles UPL and returns SPL.
	Parameters:
		decompileRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) Decompile(decompileRequest DecompileRequest, resp ...*http.Response) (*DecompileResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/decompile`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: decompileRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb DecompileResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	DeleteConnection - Delete all versions of a connection by its id.
	Parameters:
		connectionId: Connection ID
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteConnection(connectionId string, resp ...*http.Response) error {
	pp := struct {
		ConnectionId string
	}{
		ConnectionId: connectionId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/connections/{{.ConnectionId}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	DeleteFile - Delete file.
	Parameters:
		fileId: File ID
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteFile(fileId string, resp ...*http.Response) error {
	pp := struct {
		FileId string
	}{
		FileId: fileId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/files/{{.FileId}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	DeletePipeline - Removes a pipeline.
	Parameters:
		id: Pipeline ID
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeletePipeline(id string, resp ...*http.Response) error {
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/{{.Id}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	DeleteTemplate - Removes a template with a specific ID.
	Parameters:
		templateId: Template ID
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteTemplate(templateId string, resp ...*http.Response) error {
	pp := struct {
		TemplateId string
	}{
		TemplateId: templateId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/templates/{{.TemplateId}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	GetFilesMetadata - Returns files metadata.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetFilesMetadata(resp ...*http.Response) (*FilesMetaDataResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/files`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb FilesMetaDataResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetInputSchema - Returns the input schema for a function in a pipeline.
	Parameters:
		getInputSchemaRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetInputSchema(getInputSchemaRequest GetInputSchemaRequest, resp ...*http.Response) (*UplType, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/input-schema`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: getInputSchemaRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb UplType
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetLookupTable - Returns lookup table results.
	Parameters:
		connectionId: Connection ID
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetLookupTable(connectionId string, query *GetLookupTableQueryParams, resp ...*http.Response) (*LookupTableResponse, error) {
	values := util.ParseURLParams(query)
	pp := struct {
		ConnectionId string
	}{
		ConnectionId: connectionId,
	}
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/streams/v3beta1/lookups/{{.ConnectionId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb LookupTableResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetOutputSchema - Returns the output schema for a specified function in a pipeline.
	Parameters:
		getOutputSchemaRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetOutputSchema(getOutputSchemaRequest GetOutputSchemaRequest, resp ...*http.Response) (map[string]UplType, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/output-schema`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: getOutputSchemaRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb map[string]UplType
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	GetPipeline - Returns an individual pipeline by version.
	Parameters:
		id: Pipeline ID
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetPipeline(id string, query *GetPipelineQueryParams, resp ...*http.Response) (*PipelineResponse, error) {
	values := util.ParseURLParams(query)
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/streams/v3beta1/pipelines/{{.Id}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PipelineResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetPipelineLatestMetrics - Returns the latest metrics for a single pipeline.
	Parameters:
		id: Pipeline ID
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetPipelineLatestMetrics(id string, resp ...*http.Response) (*MetricsResponse, error) {
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/{{.Id}}/metrics/latest`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb MetricsResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetPipelinesStatus - Returns the status of pipelines from the underlying streaming system.
	Parameters:
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetPipelinesStatus(query *GetPipelinesStatusQueryParams, resp ...*http.Response) (*PaginatedResponseOfPipelineJobStatus, error) {
	values := util.ParseURLParams(query)
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/streams/v3beta1/pipelines/status`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PaginatedResponseOfPipelineJobStatus
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetPreviewData - Returns the preview data for a session.
	Parameters:
		previewSessionId: Preview Session ID
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetPreviewData(previewSessionId int64, resp ...*http.Response) (*PreviewData, error) {
	pp := struct {
		PreviewSessionId int64
	}{
		PreviewSessionId: previewSessionId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/preview-data/{{.PreviewSessionId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PreviewData
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetPreviewSession - Returns information from a preview session.
	Parameters:
		previewSessionId: Preview Session ID
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetPreviewSession(previewSessionId int64, resp ...*http.Response) (*PreviewState, error) {
	pp := struct {
		PreviewSessionId int64
	}{
		PreviewSessionId: previewSessionId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/preview-session/{{.PreviewSessionId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PreviewState
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetPreviewSessionLatestMetrics - Returns the latest metrics for a preview session.
	Parameters:
		previewSessionId: Preview Session ID
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetPreviewSessionLatestMetrics(previewSessionId int64, resp ...*http.Response) (*MetricsResponse, error) {
	pp := struct {
		PreviewSessionId int64
	}{
		PreviewSessionId: previewSessionId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/preview-session/{{.PreviewSessionId}}/metrics/latest`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb MetricsResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetRegistry - Returns all functions in JSON format.
	Parameters:
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetRegistry(query *GetRegistryQueryParams, resp ...*http.Response) (*RegistryModel, error) {
	values := util.ParseURLParams(query)
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/streams/v3beta1/pipelines/registry`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb RegistryModel
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetTemplate - Returns an individual template by version.
	Parameters:
		templateId: Template ID
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetTemplate(templateId string, query *GetTemplateQueryParams, resp ...*http.Response) (*TemplateResponse, error) {
	values := util.ParseURLParams(query)
	pp := struct {
		TemplateId string
	}{
		TemplateId: templateId,
	}
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/streams/v3beta1/templates/{{.TemplateId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb TemplateResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListConnections - Returns a list of connections (latest versions only) by tenant ID.
	Parameters:
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListConnections(query *ListConnectionsQueryParams, resp ...*http.Response) (*PaginatedResponseOfConnectionResponse, error) {
	values := util.ParseURLParams(query)
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/streams/v3beta1/connections`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PaginatedResponseOfConnectionResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListConnectors - Returns a list of the available connectors.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListConnectors(resp ...*http.Response) (*PaginatedResponseOfConnectorResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/connectors`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PaginatedResponseOfConnectorResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListPipelines - Returns all pipelines.
	Parameters:
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListPipelines(query *ListPipelinesQueryParams, resp ...*http.Response) (*PaginatedResponseOfPipelineResponse, error) {
	values := util.ParseURLParams(query)
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/streams/v3beta1/pipelines`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PaginatedResponseOfPipelineResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListTemplates - Returns a list of all templates.
	Parameters:
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListTemplates(query *ListTemplatesQueryParams, resp ...*http.Response) (*PaginatedResponseOfTemplateResponse, error) {
	values := util.ParseURLParams(query)
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/streams/v3beta1/templates`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PaginatedResponseOfTemplateResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	PatchPipeline - Patches an existing pipeline.
	Parameters:
		id: Pipeline ID
		pipelinePatchRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PatchPipeline(id string, pipelinePatchRequest PipelinePatchRequest, resp ...*http.Response) (*PipelineResponse, error) {
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/{{.Id}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: u, Body: pipelinePatchRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PipelineResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	PutConnection - Updates an existing DSP connection.
	Parameters:
		connectionId: Connection ID
		connectionPutRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PutConnection(connectionId string, connectionPutRequest ConnectionPutRequest, resp ...*http.Response) (*ConnectionSaveResponse, error) {
	pp := struct {
		ConnectionId string
	}{
		ConnectionId: connectionId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/connections/{{.ConnectionId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Put(services.RequestParams{URL: u, Body: connectionPutRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb ConnectionSaveResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	PutTemplate - Updates an existing template.
	Parameters:
		templateId: Template ID
		templatePutRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PutTemplate(templateId string, templatePutRequest TemplatePutRequest, resp ...*http.Response) (*TemplateResponse, error) {
	pp := struct {
		TemplateId string
	}{
		TemplateId: templateId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/templates/{{.TemplateId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Put(services.RequestParams{URL: u, Body: templatePutRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb TemplateResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ReactivatePipeline - Reactivate a pipeline
	Parameters:
		id: Pipeline ID
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ReactivatePipeline(id string, resp ...*http.Response) (*PipelineReactivateResponse, error) {
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/{{.Id}}/reactivate`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PipelineReactivateResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	StartPreview - Creates a preview session for a pipeline.
	Parameters:
		previewSessionStartRequest: Parameters to start a new Preview session
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) StartPreview(previewSessionStartRequest PreviewSessionStartRequest, resp ...*http.Response) (*PreviewStartResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/preview-session`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: previewSessionStartRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PreviewStartResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	StopPreview - Stops a preview session.
	Parameters:
		previewSessionId: Preview Session ID
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) StopPreview(previewSessionId int64, resp ...*http.Response) error {
	pp := struct {
		PreviewSessionId int64
	}{
		PreviewSessionId: previewSessionId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/preview-session/{{.PreviewSessionId}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	UpdateConnection - Patches an existing DSP connection.
	Parameters:
		connectionId: Connection ID
		connectionPatchRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UpdateConnection(connectionId string, connectionPatchRequest ConnectionPatchRequest, resp ...*http.Response) (*ConnectionSaveResponse, error) {
	pp := struct {
		ConnectionId string
	}{
		ConnectionId: connectionId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/connections/{{.ConnectionId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: u, Body: connectionPatchRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb ConnectionSaveResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	UpdatePipeline - Updates an existing pipeline.
	Parameters:
		id: Pipeline ID
		pipelineRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UpdatePipeline(id string, pipelineRequest PipelineRequest, resp ...*http.Response) (*PipelineResponse, error) {
	pp := struct {
		Id string
	}{
		Id: id,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/{{.Id}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Put(services.RequestParams{URL: u, Body: pipelineRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb PipelineResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	UpdateTemplate - Patches an existing template.
	Parameters:
		templateId: Template ID
		templatePatchRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UpdateTemplate(templateId string, templatePatchRequest TemplatePatchRequest, resp ...*http.Response) (*TemplateResponse, error) {
	pp := struct {
		TemplateId string
	}{
		TemplateId: templateId,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/templates/{{.TemplateId}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: u, Body: templatePatchRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb TemplateResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ValidatePipeline - Verifies whether the Streams JSON is valid.
	Parameters:
		validateRequest: Request JSON
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ValidatePipeline(validateRequest ValidateRequest, resp ...*http.Response) (*ValidateResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/streams/v3beta1/pipelines/validate`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: validateRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb ValidateResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}
