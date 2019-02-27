# streams
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/streams"


## Usage

#### type ActivatePipelineRequest

```go
type ActivatePipelineRequest struct {
	IDs           []string `json:"ids"`
	SkipSavePoint bool     `json:"skipSavepoint"`
}
```

ActivatePipelineRequest contains the request to activate the pipeline

#### type AdditionalProperties

```go
type AdditionalProperties map[string][]string
```

AdditionalProperties contain the properties in an activate/deactivate response

#### type Connection

```go
type Connection struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
```

Connection represents a single connection

#### type Connections

```go
type Connections struct {
	Connections []Connection `json:"connections,omitempty"`
}
```

Connections for a specific connector

#### type Connector

```go
type Connector struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}
```

Connector represents a single connector

#### type Connectors

```go
type Connectors struct {
	Connectors []Connector `json:"connectors,omitempty"`
}
```

Connectors contains all the available connectors

#### type DslCompilationRequest

```go
type DslCompilationRequest struct {
	Dsl string `json:"dsl"`
}
```

DslCompilationRequest contains the DSL that needs to be compiled into a valid
UPL JSON

#### type GetInputSchemaRequest

```go
type GetInputSchemaRequest struct {
	NodeUUID       *string      `json:"nodeUuid,omitempty"`
	TargetPortName *string      `json:"targetPortName,omitempty"`
	UplJSON        *UplPipeline `json:"uplJson,omitempty"`
}
```

GetInputSchemaRequest contains the request for the input schema for a function
in the pipeline

#### type GetOutputSchemaRequest

```go
type GetOutputSchemaRequest struct {
	NodeUUID       *string      `json:"nodeUuid,omitempty"`
	SourcePortName *string      `json:"sourcePortName,omitempty"`
	UplJSON        *UplPipeline `json:"uplJson,omitempty"`
}
```

GetOutputSchemaRequest contains the request for the output schema for a function
in the pipeline

#### type GroupExpandRequest

```go
type GroupExpandRequest struct {
	// Function arguments for the given id. Overrides default values.
	Arguments map[string]interface{} `json:"arguments"`
	// The ID associated with your group function in the pipeline Streams JSON
	ID string `json:"id"`
}
```

GroupExpandRequest contains request to create expanded group

#### type GroupFunctionArgsMappingNode

```go
type GroupFunctionArgsMappingNode struct {
	Arguments  *[]GroupFunctionArgsNode `json:"arguments,omitempty"`
	FunctionID *string                  `json:"function_id,omitempty"`
}
```

GroupFunctionArgsMappingNode Group Arguments Mapping

#### type GroupFunctionArgsNode

```go
type GroupFunctionArgsNode struct {
	FunctionArg *string `json:"function_arg,omitempty"`
	GroupArg    *string `json:"group_arg,omitempty"`
}
```

GroupFunctionArgsNode contains arguments specific to a function and group

#### type GroupResponse

```go
type GroupResponse struct {
	Ast             *UplPipeline                   `json:"ast,omitempty"`
	Attributes      *map[string]interface{}        `json:"attributes,omitempty"`
	Categories      *[]int64                       `json:"categories,omitempty"`
	CreateDate      *int64                         `json:"createDate,omitempty"`
	CreateUserID    *string                        `json:"createUserId,omitempty"`
	GroupID         *string                        `json:"groupId,omitempty"`
	LastUpdateDate  *int64                         `json:"lastUpdateDate,omitempty"`
	LastUpdateUserD *string                        `json:"lastUpdateUserId,omitempty"`
	Mappings        []GroupFunctionArgsMappingNode `json:"mappings,omitempty"`
	Name            *string                        `json:"name,omitempty"`
	OutputType      *string                        `json:"outputType,omitempty"`
	Scalar          *bool                          `json:"scalar,omitempty"`
	TenantID        *string                        `json:"tenantId,omitempty"`
	Variadic        *bool                          `json:"variadic,omitempty"`
}
```

GroupResponse contains full streams response of a group

#### type MetricsResponse

```go
type MetricsResponse struct {
	Nodes map[string]NodeMetrics `json:"nodes,omitempty"`
}
```

MetricsResponse contains metrics for a single pipeline.

#### type NodeMetrics

```go
type NodeMetrics struct {
	Metrics map[string]interface{} `json:"metrics,omitempty"`
}
```

NodeMetrics contains metrics corresponding to each node

#### type NodeType

```go
type NodeType string
```

NodeType lists different node types

```go
const (
	// Array nodetype
	Array NodeType = "ARRAY"
	// Binary nodetype
	Binary NodeType = "BINARY"
	// Boolean nodetype
	Boolean NodeType = "BOOLEAN"
	// Missing nodetype
	Missing NodeType = "MISSING"
	// Null nodetype
	Null NodeType = "NULL"
	// Number nodetype
	Number NodeType = "NUMBER"
	// Object nodetype
	Object NodeType = "OBJECT"
	// Pojo nodetype
	Pojo NodeType = "POJO"
	// String nodetype
	String NodeType = "STRING"
)
```

#### type ObjectNode

```go
type ObjectNode struct {
	Array               *bool    `json:"array,omitempty"`
	BigDecimal          *bool    `json:"bigDecimal,omitempty"`
	BigInteger          *bool    `json:"bigInteger,omitempty"`
	Binary              *bool    `json:"binary,omitempty"`
	Boolean             *bool    `json:"boolean,omitempty"`
	ContainerNode       *bool    `json:"containerNode,omitempty"`
	Double              *bool    `json:"double,omitempty"`
	Float               *bool    `json:"float,omitempty"`
	FloatingPointNumber *bool    `json:"floatingPointNumber,omitempty"`
	Int                 *bool    `json:"int,omitempty"`
	IntegralNumber      *bool    `json:"integralNumber,omitempty"`
	Long                *bool    `json:"long,omitempty"`
	MissingNode         *bool    `json:"missingNode,omitempty"`
	NodeType            NodeType `json:"nodeType,omitempty"`
	Null                *bool    `json:"null,omitempty"`
	Number              *bool    `json:"number,omitempty"`
	Object              *bool    `json:"object,omitempty"`
	Pojo                *bool    `json:"pojo,omitempty"`
	Short               *bool    `json:"short,omitempty"`
	Textual             *bool    `json:"textual,omitempty"`
	ValueNode           *bool    `json:"valueNode,omitempty"`
}
```

ObjectNode contains different object types

#### type PaginatedPipelineResponse

```go
type PaginatedPipelineResponse struct {
	Items []Pipeline `json:"items"`
	Total int64      `json:"total"`
}
```

PaginatedPipelineResponse contains the pipeline response

#### type PaginatedPipelineStatusResponse

```go
type PaginatedPipelineStatusResponse struct {
	Items []PipelineJob `json:"items"`
	Total *int64        `json:"total"`
}
```

PaginatedPipelineStatusResponse contains a list of pipeline job statuses and the
total count of pipeline jobs

#### type PaginatedTemplateResponse

```go
type PaginatedTemplateResponse struct {
	Items []TemplateResponse `json:"items"`
	Total *int64             `json:"total"`
}
```

PaginatedTemplateResponse contains a list of templates and the total count of
templates

#### type Parameters

```go
type Parameters struct {
	Parameters []UplType `json:"parameters"`
}
```

Parameters represents all the UplTypes in the pipeline

#### type PartialTemplateRequest

```go
type PartialTemplateRequest struct {
	Data        *UplPipeline `json:"data,omitempty"`
	Description *string      `json:"description,omitempty"`
	Name        *string      `json:"name,omitempty"`
}
```

PartialTemplateRequest contains the template request data for partial update
operation

#### type Pipeline

```go
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
```

Pipeline defines a pipeline object

#### type PipelineDeleteResponse

```go
type PipelineDeleteResponse struct {
	CouldDeactivate bool `json:"couldDeactivate"`
	Running         bool `json:"running"`
}
```

PipelineDeleteResponse contains the response returned as a result of a delete
pipeline call

#### type PipelineJob

```go
type PipelineJob struct {
	JobID      string `json:"jobId"`
	JobStatus  string `json:"jobStatus"`
	PipelineID string `json:"createUserId,omitempty"`
}
```

PipelineJob contains pipeline job data from the underlying streaming system

#### type PipelineQueryParams

```go
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
```

PipelineQueryParams contains the query parameters that can be provided by the
user to fetch specific pipelines

#### type PipelineReactivateResponse

```go
type PipelineReactivateResponse struct {
	CurrentlyActiveVersion     int                        `json:"currentlyActiveVersion"`
	PipelineID                 string                     `json:"pipelineId"`
	PipelineReactivationStatus PipelineReactivationStatus `json:"pipelineReactivationStatus"`
}
```

PipelineReactivateResponse contains the response returned as a result of a
reactivate pipeline call

#### type PipelineReactivationStatus

```go
type PipelineReactivationStatus string
```

PipelineReactivationStatus reflects the possible states of a pipeline that are
returned when reactivation request is sent

```go
const (
	// ReactivationActivated status
	ReactivationActivated PipelineReactivationStatus = "activated"
	// AlreadyActivatedWithCurrentVersion status
	AlreadyActivatedWithCurrentVersion PipelineReactivationStatus = "alreadyActivatedWithCurrentVersion"
	// CurrentVersionInvalid status
	CurrentVersionInvalid PipelineReactivationStatus = "currentVersionInvalid"
	// FailedToDeactivateCurrentVersion status
	FailedToDeactivateCurrentVersion PipelineReactivationStatus = "failedToDeactivateCurrentVersion"
	// RolledBack status
	RolledBack PipelineReactivationStatus = "rolledBack"
	// RolledBackError status
	RolledBackError PipelineReactivationStatus = "rolledBackError"
)
```

#### type PipelineRequest

```go
type PipelineRequest struct {
	BypassValidation         bool         `json:"bypassValidation"`
	CreateUserID             string       `json:"createUserId"`
	Data                     *UplPipeline `json:"data"`
	Description              string       `json:"description"`
	Name                     string       `json:"name"`
	StreamingConfigurationID *int64       `json:"streamingConfigurationId,omitempty"`
}
```

PipelineRequest contains the pipeline data

#### type PipelineStatus

```go
type PipelineStatus string
```

PipelineStatus reflects the status of a pipeline

```go
const (
	// Created status
	Created PipelineStatus = "CREATED"
	// Activated status
	Activated PipelineStatus = "ACTIVATED"
)
```

#### type PipelineStatusQueryParams

```go
type PipelineStatusQueryParams struct {
	Offset       *int32  `json:"offset,omitempty"`
	PageSize     *int32  `json:"pageSize,omitempty"`
	SortField    *string `json:"sortField,omitempty"`
	SortDir      *string `json:"sortDir,omitempty"`
	Activated    *bool   `json:"activated,omitempty"`
	CreateUserID *string `json:"createUserId,omitempty"`
	Name         *string `json:"name,omitempty"`
}
```

PipelineStatusQueryParams contains the query parameters that can be provided by
the user to fetch specific pipeline job statuses

#### type PipelinesMergeRequest

```go
type PipelinesMergeRequest struct {
	InputTree  *UplPipeline `json:"inputTree"`
	MainTree   *UplPipeline `json:"mainTree"`
	TargetNode *string      `json:"targetNode"`
	TargetPort *string      `json:"targetPort"`
}
```

PipelinesMergeRequest contains pipelines merge request data

#### type PreviewData

```go
type PreviewData struct {
	CurrentNumberOfRecords *int64                 `json:"currentNumberOfRecords"`
	Nodes                  map[string]PreviewNode `json:"nodes"`
	PipelineID             *string                `json:"pipelineId"`
	PreviewID              *string                `json:"previewId"`
	RecordsPerPipeline     *int64                 `json:"recordsPerPipeline"`
}
```

PreviewData contains the preview data response

#### type PreviewNode

```go
type PreviewNode struct {
	NodeName *string      `json:"nodeName"`
	Records  []ObjectNode `json:"records"`
}
```

PreviewNode contains Preview node data

#### type PreviewSessionStartRequest

```go
type PreviewSessionStartRequest struct {
	RecordsLimit             *int64       `json:"recordsLimit,omitempty"`
	RecordsPerPipeline       *int64       `json:"recordsPerPipeline,omitempty"`
	SessionLifetimeMs        *int64       `json:"sessionLifetimeMs,omitempty"`
	StreamingConfigurationID *int64       `json:"streamingConfigurationId,omitempty"`
	Upl                      *UplPipeline `json:"upl"`
	UseNewData               *bool        `json:"useNewData,omitempty"`
}
```

PreviewSessionStartRequest contains the preview session start request data

#### type PreviewStartResponse

```go
type PreviewStartResponse struct {
	PipelineID *string `json:"pipelineId"`
	PreviewID  *int64  `json:"previewId"`
}
```

PreviewStartResponse contains the preview start response

#### type PreviewState

```go
type PreviewState struct {
	ActivatedDate          *int64  `json:"activatedDate"`
	CreatedDate            *int64  `json:"createdDate"`
	CurrentNumberOfRecords *int64  `json:"currentNumberOfRecords"`
	JobID                  *string `json:"jobId"`
	PreviewID              *int64  `json:"previewId"`
	RecordsPerPipeline     *int64  `json:"recordsPerPipeline"`
}
```

PreviewState contains the preview session data

#### type Service

```go
type Service services.BaseService
```

Service - A service that deals with pipelines

#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new streams service client from the given Config

#### func (*Service) ActivatePipeline

```go
func (s *Service) ActivatePipeline(ids []string) (AdditionalProperties, error)
```
ActivatePipeline activates an existing pipeline

#### func (*Service) CompileDslToUpl

```go
func (s *Service) CompileDslToUpl(dsl *DslCompilationRequest) (*UplPipeline, error)
```
CompileDslToUpl creates a Upl Json from DSL

#### func (*Service) CreateExpandedGroup

```go
func (s *Service) CreateExpandedGroup(groupID string, args map[string]interface{}, functionID string) (*UplPipeline, error)
```
CreateExpandedGroup creates and returns the expanded version of a group

#### func (*Service) CreatePipeline

```go
func (s *Service) CreatePipeline(pipeline *PipelineRequest) (*Pipeline, error)
```
CreatePipeline creates a new pipeline

#### func (*Service) CreateTemplate

```go
func (s *Service) CreateTemplate(previewSession *TemplateRequest) (*TemplateResponse, error)
```
CreateTemplate creates a new template for a tenant

#### func (*Service) DeactivatePipeline

```go
func (s *Service) DeactivatePipeline(ids []string) (AdditionalProperties, error)
```
DeactivatePipeline deactivates an existing pipeline

#### func (*Service) DeletePipeline

```go
func (s *Service) DeletePipeline(id string) (*PipelineDeleteResponse, error)
```
DeletePipeline deletes a pipeline

#### func (*Service) DeletePreviewSession

```go
func (s *Service) DeletePreviewSession(id string) error
```
DeletePreviewSession stops a preview session

#### func (*Service) DeleteTemplate

```go
func (s *Service) DeleteTemplate(id string) error
```
DeleteTemplate deletes a template based on the provided template id

#### func (*Service) GetConnections

```go
func (s *Service) GetConnections(connectorID string) (*Connections, error)
```
GetConnections gets the connections for a specific connector

#### func (*Service) GetConnectors

```go
func (s *Service) GetConnectors() (*Connectors, error)
```
GetConnectors gets all the available connectors

#### func (*Service) GetGroupByID

```go
func (s *Service) GetGroupByID(groupID string) (*GroupResponse, error)
```
GetGroupByID retrieves the full streams JSON of a group

#### func (*Service) GetInputSchema

```go
func (s *Service) GetInputSchema(nodeUUID *string, targetPortName *string, upl *UplPipeline) (*Parameters, error)
```
GetInputSchema returns the input schema for a function in the pipeline

#### func (*Service) GetLatestPipelineMetrics

```go
func (s *Service) GetLatestPipelineMetrics(pipelineID string) (*MetricsResponse, error)
```
GetLatestPipelineMetrics gets latest Pipeline metrics

#### func (*Service) GetLatestPreviewSessionMetrics

```go
func (s *Service) GetLatestPreviewSessionMetrics(previewSessionID string) (*MetricsResponse, error)
```
GetLatestPreviewSessionMetrics gets latest Preview session metrics

#### func (*Service) GetOutputSchema

```go
func (s *Service) GetOutputSchema(nodeUUID *string, sourcePortName *string, upl *UplPipeline) (*Parameters, error)
```
GetOutputSchema returns the output schema for the specified function in the
pipeline.

#### func (*Service) GetPipeline

```go
func (s *Service) GetPipeline(id string) (*Pipeline, error)
```
GetPipeline gets an individual pipeline

#### func (*Service) GetPipelineStatus

```go
func (s *Service) GetPipelineStatus(queryParams PipelineStatusQueryParams) (*PaginatedPipelineStatusResponse, error)
```
GetPipelineStatus gets status of pipelines from the underlying streaming system

#### func (*Service) GetPipelines

```go
func (s *Service) GetPipelines(queryParams PipelineQueryParams) (*PaginatedPipelineResponse, error)
```
GetPipelines gets all the pipelines

#### func (*Service) GetPreviewData

```go
func (s *Service) GetPreviewData(id string) (*PreviewData, error)
```
GetPreviewData gets preview data for a session

#### func (*Service) GetPreviewSession

```go
func (s *Service) GetPreviewSession(id string) (*PreviewState, error)
```
GetPreviewSession gets an individual pipeline

#### func (*Service) GetRegistry

```go
func (s *Service) GetRegistry(local url.Values) (*UplRegistry, error)
```
GetRegistry returns all functions in JSON format in the registry

#### func (*Service) GetTemplate

```go
func (s *Service) GetTemplate(id string) (*TemplateResponse, error)
```
GetTemplate gets an individual template by template id

#### func (*Service) GetTemplates

```go
func (s *Service) GetTemplates() (*PaginatedTemplateResponse, error)
```
GetTemplates gets a list of latest templates

#### func (*Service) MergePipelines

```go
func (s *Service) MergePipelines(mergeRequest *PipelinesMergeRequest) (*UplPipeline, error)
```
MergePipelines merges two Upl pipelines into one Upl pipeline

#### func (*Service) ReactivatePipeline

```go
func (s *Service) ReactivatePipeline(id string) (*PipelineReactivateResponse, error)
```
ReactivatePipeline reactivates an existing pipeline

#### func (*Service) StartPreviewSession

```go
func (s *Service) StartPreviewSession(previewSession *PreviewSessionStartRequest) (*PreviewStartResponse, error)
```
StartPreviewSession starts a preview session for an existing pipeline

#### func (*Service) UpdatePipeline

```go
func (s *Service) UpdatePipeline(id string, pipeline *PipelineRequest) (*Pipeline, error)
```
UpdatePipeline updates an existing pipeline

#### func (*Service) UpdateTemplate

```go
func (s *Service) UpdateTemplate(id string, template *TemplateRequest) (*TemplateResponse, error)
```
UpdateTemplate updates an existing template (requires all fields)

#### func (*Service) UpdateTemplatePartially

```go
func (s *Service) UpdateTemplatePartially(id string, template *PartialTemplateRequest) (*TemplateResponse, error)
```
UpdateTemplatePartially partially updates an existing template (able to send
partial data)

#### func (*Service) ValidateUplResponse

```go
func (s *Service) ValidateUplResponse(upl *UplPipeline) (*ValidateResponse, error)
```
ValidateUplResponse validates if the Streams JSON is valid

#### type Servicer

```go
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
```

Servicer ...

#### type TemplateRequest

```go
type TemplateRequest struct {
	Data        *UplPipeline `json:"data"`
	Description *string      `json:"description"`
	Name        *string      `json:"name"`
}
```

TemplateRequest contains the create/update template request data

#### type TemplateResponse

```go
type TemplateResponse struct {
	CreateDate    *int64       `json:"createDate,omitempty"`
	CreateUserID  *string      `json:"createUserId,omitempty"`
	Data          *UplPipeline `json:"data"`
	Description   *string      `json:"description,omitempty"`
	Name          *string      `json:"name,omitempty"`
	OwnerTenantID *string      `json:"ownerTenantId,omitempty"`
	TemplateID    *string      `json:"templateId,omitempty"`
	Version       *int64       `json:"version,omitempty"`
}
```

TemplateResponse contains the create template response data

#### type UplArgument

```go
type UplArgument struct {
	//ElementType *map[string]interface{} `json:"element-type,omitempty"`
	ElementType interface{} `json:"element-type,omitempty"`
	Type        string      `json:"type"`
}
```

UplArgument are arguments to UplFunctions, UplFunctions have one or more of the
arguments.

#### type UplCategory

```go
type UplCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
```

UplCategory represents a category in the Upl registry

#### type UplEdge

```go
type UplEdge struct {
	Attributes interface{} `json:"attributes"`
	SourceNode string      `json:"sourceNode"`
	SourcePort string      `json:"sourcePort"`
	TargetNode string      `json:"targetNode"`
	TargetPort string      `json:"targetPort"`
}
```

UplEdge contains information on the edges between two pipeline nodes

#### type UplFunction

```go
type UplFunction struct {
	Arguments  map[string]UplArgument `json:"arguments,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Categories *[]int64               `json:"categories,omitempty"`
	ID         *string                `json:"id,omitempty"`
	IsVariadic *bool                  `json:"isVariadic,omitempty"`
	Op         *string                `json:"op,omitempty"`
	Output     UplArgument            `json:"output,omitempty"`
	ResolvedID *string                `json:"resolvedId,omitempty"`
}
```

UplFunction contains the basic building block of a UPL pipeline

#### type UplNode

```go
type UplNode interface{}
```

UplNode defines the nodes forming a pipeline

#### type UplPipeline

```go
type UplPipeline struct {
	Edges    []UplEdge `json:"edges"`
	Nodes    []UplNode `json:"nodes"`
	RootNode []string  `json:"root-node"`
	Version  int32     `json:"version"`
}
```

UplPipeline contains the pipeline data

#### type UplRegistry

```go
type UplRegistry struct {
	Categories []UplCategory `json:"categories,omitempty"`
	Functions  []UplFunction `json:"functions,omitempty"`
	Types      []UplType     `json:"types,omitempty"`
}
```

UplRegistry contains all functions and types

#### type UplType

```go
type UplType struct {
	FieldName  *string   `json:"fieldName,omitempty"`
	Parameters []UplType `json:"parameters,omitempty"`
	Type       *string   `json:"type,omitempty"`
}
```

UplType contains UplTypes stored in the registry

#### type ValidateRequest

```go
type ValidateRequest struct {
	Upl *UplPipeline `json:"upl,omitempty"`
}
```

ValidateRequest contains the request with the UplPipeline to validate

#### type ValidateResponse

```go
type ValidateResponse struct {
	Success            *bool     `json:"success,omitempty"`
	ValidationMessages *[]string `json:"validationMessages,omitempty"`
}
```

ValidateResponse contains the Validation response after validating the
UplPipeline
