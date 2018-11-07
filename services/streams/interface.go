// Code generated by gen-interface.go. DO NOT EDIT.
package streams

type Servicer interface {
	// CompileDslToUpl creates a Upl Json from DSL
	CompileDslToUpl(dsl *DslCompilationRequest) (*UplPipeline, error)
	// GetPipelines gets all the pipelines
	GetPipelines(queryParams PipelineQueryParams) (*PaginatedPipelineResponse, error)
	// CreatePipeline creates a new pipeline
	CreatePipeline(pipeline *PipelineRequest) (*Pipeline, error)
	// ActivatePipeline activates an existing pipeline
	ActivatePipeline(ids []string) (AdditionalProperties, error)
	// DeactivatePipeline deactivates an existing pipeline
	DeactivatePipeline(ids []string) (AdditionalProperties, error)
	// GetPipeline gets an individual pipeline
	GetPipeline(id string) (*Pipeline, error)
	// UpdatePipeline updates an existing pipeline
	UpdatePipeline(id string, pipeline *PipelineRequest) (*Pipeline, error)
	// DeletePipeline deletes a pipeline
	DeletePipeline(id string) (*PipelineDeleteResponse, error)
}
