// Code generated by gen_interface.go. DO NOT EDIT.

package streams

import (
	"net/url"
)

// Servicer ...
type Servicer interface {
	// CompileDslToUpl creates a Upl Json from DSL
	CompileDslToUpl(dsl *DslCompilationRequest) (*UplPipeline, error)
	// GetPipelineStatus gets status of pipelines from the underlying streaming system
	GetPipelineStatus(queryParams PipelineStatusQueryParams) (*PaginatedPipelineStatusResponse, error)
	// GetPipelines gets all the pipelines
	GetPipelines(queryParams PipelineQueryParams) (*PaginatedPipelineResponse, error)
	// CreatePipeline creates a new pipeline
	CreatePipeline(pipeline *PipelineRequest) (*Pipeline, error)
	// ActivatePipeline activates an existing pipeline
	ActivatePipeline(ids []string) (AdditionalProperties, error)
	// DeactivatePipeline deactivates an existing pipeline
	DeactivatePipeline(ids []string) (AdditionalProperties, error)
	// ReactivatePipeline reactivates an existing pipeline
	ReactivatePipeline(id string) (*PipelineReactivateResponse, error)
	// GetPipeline gets an individual pipeline
	GetPipeline(id string) (*Pipeline, error)
	// UpdatePipeline updates an existing pipeline
	UpdatePipeline(id string, pipeline *PipelineRequest) (*Pipeline, error)
	// MergePipelines merges two Upl pipelines into one Upl pipeline
	MergePipelines(mergeRequest *PipelinesMergeRequest) (*UplPipeline, error)
	// DeletePipeline deletes a pipeline
	DeletePipeline(id string) (*PipelineDeleteResponse, error)
	// GetInputSchema returns the input schema for a function in the pipeline
	GetInputSchema(nodeUUID *string, targetPortName *string, upl *UplPipeline) (*Parameters, error)
	// GetOutputSchema returns the output schema for the specified function in the pipeline.
	GetOutputSchema(nodeUUID *string, sourcePortName *string, upl *UplPipeline) (*Parameters, error)
	// GetRegistry returns all functions in JSON format in the registry
	GetRegistry(local url.Values) (*UplRegistry, error)
	// GetLatestPreviewSessionMetrics gets latest Preview session metrics
	GetLatestPreviewSessionMetrics(previewSessionID string) (*MetricsResponse, error)
	// GetLatestPipelineMetrics gets latest Pipeline metrics
	GetLatestPipelineMetrics(pipelineID string) (*MetricsResponse, error)
	// ValidateUplResponse validates if the Streams JSON is valid
	ValidateUplResponse(upl *UplPipeline) (*ValidateResponse, error)
	// GetConnectors gets all the available connectors
	GetConnectors() (*Connectors, error)
	// GetConnections gets the connections for a specific connector
	GetConnections(connectorID string) (*Connections, error)
	// StartPreviewSession starts a preview session for an existing pipeline
	StartPreviewSession(previewSession *PreviewSessionStartRequest) (*PreviewStartResponse, error)
	// GetPreviewSession gets an individual pipeline
	GetPreviewSession(id string) (*PreviewState, error)
	// DeletePreviewSession stops a preview session
	DeletePreviewSession(id string) error
	// GetPreviewData gets preview data for a session
	GetPreviewData(id string) (*PreviewData, error)
	// CreateTemplate creates a new template for a tenant
	CreateTemplate(previewSession *TemplateRequest) (*TemplateResponse, error)
	// GetTemplates gets a list of latest templates
	GetTemplates() (*PaginatedTemplateResponse, error)
	// GetTemplate gets an individual template by template id
	GetTemplate(id string) (*TemplateResponse, error)
	// UpdateTemplate updates an existing template (requires all fields)
	UpdateTemplate(id string, template *TemplateRequest) (*TemplateResponse, error)
	// UpdateTemplatePartially partially updates an existing template (able to send partial data)
	UpdateTemplatePartially(id string, template *PartialTemplateRequest) (*TemplateResponse, error)
	// DeleteTemplate deletes a template based on the provided template id
	DeleteTemplate(id string) error
	// GetGroupByID retrieves the full streams JSON of a group
	GetGroupByID(groupID string) (*GroupResponse, error)
	// CreateExpandedGroup creates and returns the expanded version of a group
	CreateExpandedGroup(groupID string, args map[string]interface{}, functionID string) (*UplPipeline, error)
}
