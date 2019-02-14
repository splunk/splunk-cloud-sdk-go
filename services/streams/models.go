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

// PipelineReactivateResponse contains the response returned as a result of a reactivate pipeline call
type PipelineReactivateResponse struct {
	CurrentlyActiveVersion     int    `json:"currentlyActiveVersion"`
	PipelineId                 string `json:"pipelineId"`
	PipelineReactivationStatus string `json:"pipelineReactivationStatus"`
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

//Input schema request contains the request for the input schema for a function in the pipeline
type GetInputSchemaRequest struct {
	NodeUUID       *string      `json:"nodeUuid,omitempty"`
	TargetPortName *string      `json:"targetPortName,omitempty"`
	UplJSON        *UplPipeline `json:"uplJson,omitempty"`
}

//Output schema request contains the request for the output schema for a function in the pipeline
type GetOutputSchemaRequest struct {
	NodeUUID       *string      `json:"nodeUuid,omitempty"`
	SourcePortName *string      `json:"sourcePortName,omitempty"`
	UplJSON        *UplPipeline `json:"uplJson,omitempty"`
}

//UplType contains UplTypes stored in the registry
type UplType struct {
	FieldName  *string   `json:"fieldName,omitempty"`
	Parameters []UplType `json:"parameters,omitempty"`
	Type       *string   `json:"type,omitempty"`
}

//Parameters represents all the UplTypes in the pipeline
type Parameters struct {
	Parameters []UplType `json:"parameters"`
}

//UplFunction contains the basic building block of a UPL pipeline
type UplFunction struct {
	Arguments  map[string]interface{} `json:"arguments,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Categories *[]int64               `json:"categories,omitempty"`
	ID         *string                `json:"id,omitempty"`
	IsVariadic *bool                  `json:"isVariadic,omitempty"`
	Op         *string                `json:"op,omitempty"`
	Output     *UplArgument           `json:"output,omitempty"`
	ResolvedId *string                `json:"resolvedId,omitempty"`
}

//UplArgument are arguments to UplFunctions, UplFunctions have one or more of the arguments.
type UplArgument struct {
	ElementType *map[string]interface{} `json:"element-type,omitempty"`
	Type        string                  `json:"type"`
}

//UplRegistry contains all functions and types
type UplRegistry struct {
	Categories []UplCategory  `json:"categories,omitempty"`
	Functions  *[]UplFunction `json:"functions,omitempty"`
	Types      *[]UplType     `json:"types,omitempty"`
}

//UplCategory represents a category in the Upl registry
type UplCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//Contains metrics for a single pipeline.
type MetricsResponse struct {
	Nodes map[string]NodeMetrics `json:"nodes,omitempty"`
}

//NodeMetrics contains metrics corresponding to each node
type NodeMetrics struct {
	Metrics map[string]interface{} `json:"metrics,omitempty"`
}

//ValidateRequest contains the request with the UplPipeline to validate
type ValidateRequest struct {
	Upl *UplPipeline `json:"upl,omitempty"`
}

//ValidateRequest contains the Validation response after validating the UplPipeline
type ValidateResponse struct {
	Success            *bool     `json:"success,omitempty"`
	ValidationMessages *[]string `json:"validationMessages,omitempty"`
}

// PreviewSessionStartRequest contains the preview session start request data
type PreviewSessionStartRequest struct {
	RecordsLimit             int64        `json:"recordsLimit,omitempty"`
	RecordsPerPipeline       int64        `json:"recordsPerPipeline,omitempty"`
	SessionLifetimeMs        int64        `json:"sessionLifetimeMs,omitempty"`
	StreamingConfigurationID int64        `json:"streamingConfigurationId,omitempty"`
	Upl                      *UplPipeline `json:"upl"`
	UseNewData               bool         `json:"useNewData,omitempty"`
}

// PreviewStartResponse contains the preview start response
type PreviewStartResponse struct {
	PipelineID string `json:"pipelineId"`
	PreviewID  int64  `json:"previewId"`
}

// Connector represents a single connector
type Connector struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Connections for a specific connector
type Connections struct {
	Connections []Connection `json:"connections,omitempty"`
}

// Connection represents a single connection
type Connection struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Contains all the available connectors
type Connectors struct {
	Connectors []Connector `json:"connectors,omitempty"`
}

// PreviewState contains the preview session data
type PreviewState struct {
	ActivatedDate          int64  `json:"activatedDate"`
	CreatedDate            int64  `json:"createdDate"`
	CurrentNumberOfRecords int64  `json:"currentNumberOfRecords"`
	JobID                  string `json:"jobId"`
	PreviewID              int64  `json:"previewId"`
	RecordsPerPipeline     int64  `json:"recordsPerPipeline"`
}

// PaginatedPipelineStatusResponse contains a list of pipeline job statuses and the total count of pipeline jobs
type PaginatedPipelineStatusResponse struct {
	Items []PipelineJob `json:"items"`
	Total int64         `json:"total"`
}

// PipelineJob contains pipeline job data from the underlying streaming system
type PipelineJob struct {
	JobID      string `json:"jobId"`
	JobStatus  string `json:"jobStatus"`
	PipelineID string `json:"createUserId,omitempty"`
}

// PipelineStatusQueryParams contains the query parameters that can be provided by the user to fetch specific pipeline job statuses
type PipelineStatusQueryParams struct {
	Offset       *int32  `json:"offset,omitempty"`
	PageSize     *int32  `json:"pageSize,omitempty"`
	SortField    *string `json:"sortField,omitempty"`
	SortDir      *string `json:"sortDir,omitempty"`
	Activated    *bool   `json:"activated,omitempty"`
	CreateUserID *string `json:"createUserId,omitempty"`
	Name         *string `json:"name,omitempty"`
}

// TemplateRequest contains the create/update template request data
type TemplateRequest struct {
	Data        *UplPipeline `json:"data"`
	Description string       `json:"description"`
	Name        string       `json:"name"`
}

// PartialTemplateRequest contains the template request data for partial update operation
type PartialTemplateRequest struct {
	Data        *UplPipeline `json:"data,omitempty"`
	Description string       `json:"description,omitempty"`
	Name        string       `json:"name,omitempty"`
}

// TemplateResponse contains the create template response data
type TemplateResponse struct {
	CreateDate    int64        `json:"createDate,omitempty"`
	CreateUserID  string       `json:"createUserId,omitempty"`
	Data          *UplPipeline `json:"data"`
	Description   string       `json:"description,omitempty"`
	Name          string       `json:"name,omitempty"`
	OwnerTenantID string       `json:"ownerTenantId,omitempty"`
	TemplateID    string       `json:"templateId,omitempty"`
	Version       int64        `json:"version,omitempty"`
}

// PaginatedTemplateResponse contains a list of templates and the total count of templates
type PaginatedTemplateResponse struct {
	Items []TemplateResponse `json:"items"`
	Total int64              `json:"total"`
}
