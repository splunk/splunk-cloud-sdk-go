package streams

// ActivatePipelineRequest contains the request to activate the pipeline
type ActivatePipelineRequest struct {
	IDs           []string `json:"ids"`
	SkipSavePoint bool     `json:"skipSavepoint"`
}

// AdditionalProperties contain the properties in an activate/deactivate response
type AdditionalProperties map[string][]string

// DslCompilationRequest contains the DSL that needs to be compiled into a valid UPL JSON
type DslCompilationRequest struct {
	Dsl string `json:"dsl"`
}

// Pipeline defines a pipeline object
type Pipeline struct {
	ActivatedDate            int64          `json:"activatedDate"`
	ActivatedUserID          string         `json:"activatedUserId"`
	ActivatedVersion         int64          `json:"activatedVersion"`
	CreateDate               int64          `json:"createDate"`
	CreateUserID             string         `json:"createUserId"`
	CurrentVersion           int64          `json:"currentVersion"`
	Data                     UplPipeline    `json:"data"`
	Description              string         `json:"description"`
	ID                       string         `json:"id"`
	JobID                    string         `json:"jobId"`
	LastUpdateDate           int64          `json:"lastUpdateDate"`
	LastUpdateUserID         string         `json:"lastUpdateUserId"`
	Name                     string         `json:"name"`
	Status                   PipelineStatus `json:"status"`
	StatusMessage            string         `json:"statusMessage"`
	StreamingConfigurationID int64          `json:"streamingConfigurationId"`
	TenantID                 string         `json:"tenantId"`
	ValidationMessages       []string       `json:"validationMessages"`
	Version                  int64          `json:"version"`
}

// PaginatedPipelineResponse contains the pipeline response
type PaginatedPipelineResponse struct {
	Items []Pipeline `json:"items"`
	Total int64      `json:"total"`
}

// PipelineDeleteResponse contains the response returned as a result of a delete pipeline call
type PipelineDeleteResponse struct {
	CouldDeactivate bool `json:"couldDeactivate"`
	Running         bool `json:"running"`
}

// PipelineQueryParams contains the query parameters that can be provided by the user to fetch specific pipelines
type PipelineQueryParams struct {
	Offset       *int32  `json:"offset,omitempty"`
	PageSize     *int32  `json:"pageSize,omitempty"`
	SortField    *string `json:"sortField,omitempty"`
	SortDir      *string `json:"sortDir,omitempty"`
	Activated    *bool   `json:"activated,omitempty"`
	CreateUserID *string `json:"createUserId,omitempty"`
	Name         *string `json:"name,omitempty"`
	IncludeData  *bool   `json:"includeData,omitempty"`
}

// PipelineRequest contains the pipeline data
type PipelineRequest struct {
	BypassValidation         bool         `json:"bypassValidation"`
	CreateUserID             string       `json:"createUserId"`
	Data                     *UplPipeline `json:"data"`
	Description              string       `json:"description"`
	Name                     string       `json:"name"`
	StreamingConfigurationID *int64       `json:"streamingConfigurationId,omitempty"`
}

// PipelineStatus reflects the status of a pipeline
type PipelineStatus string

const (
	// Created status
	Created PipelineStatus = "CREATED"
	// Activated status
	Activated PipelineStatus = "ACTIVATED"
)

// UplPipeline contains the pipeline data
type UplPipeline struct {
	Edges    []UplEdge `json:"edges"`
	Nodes    []UplNode `json:"nodes"`
	RootNode []string  `json:"root-node"`
	Version  int32     `json:"version"`
}

// UplNode defines the nodes forming a pipeline
type UplNode interface{}

// UplEdge contains information on the edges between two pipeline nodes
type UplEdge struct {
	Attributes interface{} `json:"attributes"`
	SourceNode string      `json:"sourceNode"`
	SourcePort string      `json:"sourcePort"`
	TargetNode string      `json:"targetNode"`
	TargetPort string      `json:"targetPort"`
}
