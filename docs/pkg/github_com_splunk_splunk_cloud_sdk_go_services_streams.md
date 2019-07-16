# streams
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/streams"


## Usage

#### type ActivatePipelineRequest

```go
type ActivatePipelineRequest struct {
	// Set to true to activate the latest version of the pipeline. Set to false to use the previously activated version of the pipeline. Defaults to true.
	ActivateLatestVersion *bool `json:"activateLatestVersion,omitempty"`
	// Set to true to allow the pipeline to ignore any unused progress states. In some cases, when a data pipeline is changed, the progress state will be stored for functions that no longer exist, so this must be set to activate a pipeline in this state. Defaults to false.
	AllowNonRestoredState *bool `json:"allowNonRestoredState,omitempty"`
	// Set to true to start reading from the latest input rather than from where the pipeline's previous run left off, which can cause data loss. Defaults to false.
	SkipRestoreState *bool `json:"skipRestoreState,omitempty"`
}
```


#### type AdditionalProperties

```go
type AdditionalProperties map[string][]string
```

AdditionalProperties contain the properties in an activate/deactivate response

#### type ConnectionPatchRequest

```go
type ConnectionPatchRequest struct {
	// The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.
	Data map[string]interface{} `json:"data,omitempty"`
	// The description of the connection.
	Description *string `json:"description,omitempty"`
	// The name of the connection.
	Name *string `json:"name,omitempty"`
}
```


#### type ConnectionPutRequest

```go
type ConnectionPutRequest struct {
	// The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.
	Data map[string]interface{} `json:"data"`
	// The description of the connection.
	Description string `json:"description"`
	// The name of the connection.
	Name string `json:"name"`
}
```


#### type ConnectionRequest

```go
type ConnectionRequest struct {
	// The ID of the parent connector.
	ConnectorId string `json:"connectorId"`
	// The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.
	Data map[string]interface{} `json:"data"`
	// The description of the connection.
	Description string `json:"description"`
	// The name of the connection.
	Name string `json:"name"`
}
```


#### type ConnectionResponse

```go
type ConnectionResponse struct {
	ActivePipelinesUsing []map[string]interface{}    `json:"activePipelinesUsing,omitempty"`
	ConnectorId          *string                     `json:"connectorId,omitempty"`
	ConnectorName        *string                     `json:"connectorName,omitempty"`
	CreateDate           *int64                      `json:"createDate,omitempty"`
	CreateUserId         *string                     `json:"createUserId,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	LastUpdateDate       *int64                      `json:"lastUpdateDate,omitempty"`
	LastUpdateUserId     *string                     `json:"lastUpdateUserId,omitempty"`
	Versions             []ConnectionVersionResponse `json:"versions,omitempty"`
}
```


#### type ConnectionSaveResponse

```go
type ConnectionSaveResponse struct {
	ConnectorId  *string     `json:"connectorId,omitempty"`
	CreateDate   *int64      `json:"createDate,omitempty"`
	CreateUserId *string     `json:"createUserId,omitempty"`
	Data         *ObjectNode `json:"data,omitempty"`
	Description  *string     `json:"description,omitempty"`
	Id           *string     `json:"id,omitempty"`
	Name         *string     `json:"name,omitempty"`
	Version      *int64      `json:"version,omitempty"`
}
```


#### type ConnectionVersionResponse

```go
type ConnectionVersionResponse struct {
	CreateDate   *int64      `json:"createDate,omitempty"`
	CreateUserId *string     `json:"createUserId,omitempty"`
	Data         *ObjectNode `json:"data,omitempty"`
	Description  *string     `json:"description,omitempty"`
	Name         *string     `json:"name,omitempty"`
	Version      *int64      `json:"version,omitempty"`
}
```


#### type ConnectorResponse

```go
type ConnectorResponse struct {
	Config        *ObjectNode              `json:"config,omitempty"`
	ConnectorType *string                  `json:"connectorType,omitempty"`
	Description   *string                  `json:"description,omitempty"`
	Functions     []map[string]interface{} `json:"functions,omitempty"`
	Hidden        *bool                    `json:"hidden,omitempty"`
	Id            *string                  `json:"id,omitempty"`
	Name          *string                  `json:"name,omitempty"`
	PanelUrl      *string                  `json:"panelUrl,omitempty"`
	Tag           *string                  `json:"tag,omitempty"`
}
```


#### type DeactivatePipelineRequest

```go
type DeactivatePipelineRequest struct {
	// Set to true to skip saving the state of a deactivated pipeline. When the pipeline is later activated, it will start with the newest data and skip any data that arrived after this deactivation, which can cause data loss. Defaults to false.
	SkipSavepoint *bool `json:"skipSavepoint,omitempty"`
}
```


#### type DslCompilationRequest

```go
type DslCompilationRequest struct {
	// The Streams DSL representation of a pipeline.
	Dsl string `json:"dsl"`
}
```


#### type GetInputSchemaRequest

```go
type GetInputSchemaRequest struct {
	// The function ID.
	NodeUuid string `json:"nodeUuid"`
	// The name of the input port.
	TargetPortName string      `json:"targetPortName"`
	UplJson        UplPipeline `json:"uplJson"`
}
```


#### type GetOutputSchemaRequest

```go
type GetOutputSchemaRequest struct {
	UplJson UplPipeline `json:"uplJson"`
	// The function ID. If omitted, returns the output schema for all functions.
	NodeUuid *string `json:"nodeUuid,omitempty"`
	// The name of the output port. Deprecated.
	SourcePortName *string `json:"sourcePortName,omitempty"`
}
```


#### type GetPipelineQueryParams

```go
type GetPipelineQueryParams struct {
	// Version : version
	Version string `key:"version"`
}
```

GetPipelineQueryParams represents valid query parameters for the GetPipeline
operation For convenience GetPipelineQueryParams can be formed in a single
statement, for example:

    `v := GetPipelineQueryParams{}.SetVersion(...)`

#### func (GetPipelineQueryParams) SetVersion

```go
func (q GetPipelineQueryParams) SetVersion(v string) GetPipelineQueryParams
```

#### type GetPipelinesStatusQueryParams

```go
type GetPipelinesStatusQueryParams struct {
	// Activated : activated
	Activated *bool `key:"activated"`
	// CreateUserId : createUserId
	CreateUserId string `key:"createUserId"`
	// Name : name
	Name string `key:"name"`
	// Offset : offset
	Offset *int32 `key:"offset"`
	// PageSize : pageSize
	PageSize *int32 `key:"pageSize"`
	// SortDir : sortDir
	SortDir string `key:"sortDir"`
	// SortField : sortField
	SortField string `key:"sortField"`
}
```

GetPipelinesStatusQueryParams represents valid query parameters for the
GetPipelinesStatus operation For convenience GetPipelinesStatusQueryParams can
be formed in a single statement, for example:

    `v := GetPipelinesStatusQueryParams{}.SetActivated(...).SetCreateUserId(...).SetName(...).SetOffset(...).SetPageSize(...).SetSortDir(...).SetSortField(...)`

#### func (GetPipelinesStatusQueryParams) SetActivated

```go
func (q GetPipelinesStatusQueryParams) SetActivated(v bool) GetPipelinesStatusQueryParams
```

#### func (GetPipelinesStatusQueryParams) SetCreateUserId

```go
func (q GetPipelinesStatusQueryParams) SetCreateUserId(v string) GetPipelinesStatusQueryParams
```

#### func (GetPipelinesStatusQueryParams) SetName

```go
func (q GetPipelinesStatusQueryParams) SetName(v string) GetPipelinesStatusQueryParams
```

#### func (GetPipelinesStatusQueryParams) SetOffset

```go
func (q GetPipelinesStatusQueryParams) SetOffset(v int32) GetPipelinesStatusQueryParams
```

#### func (GetPipelinesStatusQueryParams) SetPageSize

```go
func (q GetPipelinesStatusQueryParams) SetPageSize(v int32) GetPipelinesStatusQueryParams
```

#### func (GetPipelinesStatusQueryParams) SetSortDir

```go
func (q GetPipelinesStatusQueryParams) SetSortDir(v string) GetPipelinesStatusQueryParams
```

#### func (GetPipelinesStatusQueryParams) SetSortField

```go
func (q GetPipelinesStatusQueryParams) SetSortField(v string) GetPipelinesStatusQueryParams
```

#### type GetRegistryQueryParams

```go
type GetRegistryQueryParams struct {
	// Local : local
	Local *bool `key:"local"`
}
```

GetRegistryQueryParams represents valid query parameters for the GetRegistry
operation For convenience GetRegistryQueryParams can be formed in a single
statement, for example:

    `v := GetRegistryQueryParams{}.SetLocal(...)`

#### func (GetRegistryQueryParams) SetLocal

```go
func (q GetRegistryQueryParams) SetLocal(v bool) GetRegistryQueryParams
```

#### type GetTemplateQueryParams

```go
type GetTemplateQueryParams struct {
	// Version : version of the template
	Version *int64 `key:"version"`
}
```

GetTemplateQueryParams represents valid query parameters for the GetTemplate
operation For convenience GetTemplateQueryParams can be formed in a single
statement, for example:

    `v := GetTemplateQueryParams{}.SetVersion(...)`

#### func (GetTemplateQueryParams) SetVersion

```go
func (q GetTemplateQueryParams) SetVersion(v int64) GetTemplateQueryParams
```

#### type GroupArgumentsNode

```go
type GroupArgumentsNode struct {
	// The argument name for your group function.
	GroupArg string `json:"groupArg"`
	// Group function argument position number.
	Position int32 `json:"position"`
	// The group function's data type.
	Type string `json:"type"`
}
```


#### type GroupExpandRequest

```go
type GroupExpandRequest struct {
	// Function arguments for the given id. Overrides default values.
	Arguments map[string]interface{} `json:"arguments"`
	// The ID associated with your group function in the pipeline Streams JSON
	Id string `json:"id"`
}
```


#### type GroupFunctionArgsMappingNode

```go
type GroupFunctionArgsMappingNode struct {
	// List of mappings from group function argument to function argument.
	Arguments []GroupFunctionArgsNode `json:"arguments"`
	// The function id to map to a group function argument.
	FunctionId string `json:"functionId"`
}
```


#### type GroupFunctionArgsNode

```go
type GroupFunctionArgsNode struct {
	// Function argument name.
	FunctionArg string `json:"functionArg"`
	// The argument name for your group function.
	GroupArg string `json:"groupArg"`
}
```


#### type GroupPatchRequest

```go
type GroupPatchRequest struct {
	// Group function arguments list.
	Arguments []GroupArgumentsNode `json:"arguments,omitempty"`
	Ast       *UplPipeline         `json:"ast,omitempty"`
	// Attributes map for function.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Categories for this function.
	Categories []int64 `json:"categories,omitempty"`
	// Group function mappings list.
	Mappings []GroupFunctionArgsMappingNode `json:"mappings,omitempty"`
	// The name for the group function.
	Name *string `json:"name,omitempty"`
	// The data type of the output of your function.
	OutputType *string `json:"outputType,omitempty"`
	Scalar     *bool   `json:"scalar,omitempty"`
	Variadic   *bool   `json:"variadic,omitempty"`
}
```


#### type GroupPutRequest

```go
type GroupPutRequest struct {
	// The group function arguments list.
	Arguments []GroupArgumentsNode `json:"arguments"`
	Ast       UplPipeline          `json:"ast"`
	// The attributes map for function.
	Attributes map[string]interface{} `json:"attributes"`
	// The categories for this function.
	Categories []int64 `json:"categories"`
	// The group function mappings list.
	Mappings []GroupFunctionArgsMappingNode `json:"mappings"`
	// The group function name.
	Name string `json:"name"`
	// The data type of the function's output.
	OutputType string `json:"outputType"`
	Scalar     *bool  `json:"scalar,omitempty"`
	Variadic   *bool  `json:"variadic,omitempty"`
}
```


#### type GroupRequest

```go
type GroupRequest struct {
	// The group function arguments list.
	Arguments []GroupArgumentsNode `json:"arguments"`
	Ast       UplPipeline          `json:"ast"`
	// The attributes map for function.
	Attributes map[string]interface{} `json:"attributes"`
	// The categories for this function.
	Categories []int64 `json:"categories"`
	// The group function mappings list.
	Mappings []GroupFunctionArgsMappingNode `json:"mappings"`
	// The group function name.
	Name string `json:"name"`
	// The data type of the function's output.
	OutputType string `json:"outputType"`
	Scalar     *bool  `json:"scalar,omitempty"`
	Variadic   *bool  `json:"variadic,omitempty"`
}
```


#### type GroupResponse

```go
type GroupResponse struct {
	Arguments        []GroupArgumentsNode           `json:"arguments,omitempty"`
	Ast              *UplPipeline                   `json:"ast,omitempty"`
	Attributes       map[string]interface{}         `json:"attributes,omitempty"`
	Categories       []int64                        `json:"categories,omitempty"`
	CreateDate       *int64                         `json:"createDate,omitempty"`
	CreateUserId     *string                        `json:"createUserId,omitempty"`
	GroupId          *string                        `json:"groupId,omitempty"`
	LastUpdateDate   *int64                         `json:"lastUpdateDate,omitempty"`
	LastUpdateUserId *string                        `json:"lastUpdateUserId,omitempty"`
	Mappings         []GroupFunctionArgsMappingNode `json:"mappings,omitempty"`
	Name             *string                        `json:"name,omitempty"`
	OutputType       *string                        `json:"outputType,omitempty"`
	Scalar           *bool                          `json:"scalar,omitempty"`
	TenantId         *string                        `json:"tenantId,omitempty"`
	Variadic         *bool                          `json:"variadic,omitempty"`
}
```


#### type ListConnectionsQueryParams

```go
type ListConnectionsQueryParams struct {
	ConnectorId     string `key:"connectorId"`
	CreateUserId    string `key:"createUserId"`
	FunctionId      string `key:"functionId"`
	Name            string `key:"name"`
	Offset          *int32 `key:"offset"`
	PageSize        *int32 `key:"pageSize"`
	ShowSecretNames string `key:"showSecretNames"`
	// SortDir : Specify either ascending (&#39;asc&#39;) or descending (&#39;desc&#39;) sort order for a given field (sortField), which must be set for sortDir to apply. Defaults to &#39;asc&#39;.
	SortDir   string `key:"sortDir"`
	SortField string `key:"sortField"`
}
```

ListConnectionsQueryParams represents valid query parameters for the
ListConnections operation For convenience ListConnectionsQueryParams can be
formed in a single statement, for example:

    `v := ListConnectionsQueryParams{}.SetConnectorId(...).SetCreateUserId(...).SetFunctionId(...).SetName(...).SetOffset(...).SetPageSize(...).SetShowSecretNames(...).SetSortDir(...).SetSortField(...)`

#### func (ListConnectionsQueryParams) SetConnectorId

```go
func (q ListConnectionsQueryParams) SetConnectorId(v string) ListConnectionsQueryParams
```

#### func (ListConnectionsQueryParams) SetCreateUserId

```go
func (q ListConnectionsQueryParams) SetCreateUserId(v string) ListConnectionsQueryParams
```

#### func (ListConnectionsQueryParams) SetFunctionId

```go
func (q ListConnectionsQueryParams) SetFunctionId(v string) ListConnectionsQueryParams
```

#### func (ListConnectionsQueryParams) SetName

```go
func (q ListConnectionsQueryParams) SetName(v string) ListConnectionsQueryParams
```

#### func (ListConnectionsQueryParams) SetOffset

```go
func (q ListConnectionsQueryParams) SetOffset(v int32) ListConnectionsQueryParams
```

#### func (ListConnectionsQueryParams) SetPageSize

```go
func (q ListConnectionsQueryParams) SetPageSize(v int32) ListConnectionsQueryParams
```

#### func (ListConnectionsQueryParams) SetShowSecretNames

```go
func (q ListConnectionsQueryParams) SetShowSecretNames(v string) ListConnectionsQueryParams
```

#### func (ListConnectionsQueryParams) SetSortDir

```go
func (q ListConnectionsQueryParams) SetSortDir(v string) ListConnectionsQueryParams
```

#### func (ListConnectionsQueryParams) SetSortField

```go
func (q ListConnectionsQueryParams) SetSortField(v string) ListConnectionsQueryParams
```

#### type ListPipelinesQueryParams

```go
type ListPipelinesQueryParams struct {
	// Activated : activated
	Activated *bool `key:"activated"`
	// CreateUserId : createUserId
	CreateUserId string `key:"createUserId"`
	// IncludeData : includeData
	IncludeData *bool `key:"includeData"`
	// Name : name
	Name string `key:"name"`
	// Offset : offset
	Offset *int32 `key:"offset"`
	// PageSize : pageSize
	PageSize *int32 `key:"pageSize"`
	// SortDir : sortDir
	SortDir string `key:"sortDir"`
	// SortField : sortField
	SortField string `key:"sortField"`
}
```

ListPipelinesQueryParams represents valid query parameters for the ListPipelines
operation For convenience ListPipelinesQueryParams can be formed in a single
statement, for example:

    `v := ListPipelinesQueryParams{}.SetActivated(...).SetCreateUserId(...).SetIncludeData(...).SetName(...).SetOffset(...).SetPageSize(...).SetSortDir(...).SetSortField(...)`

#### func (ListPipelinesQueryParams) SetActivated

```go
func (q ListPipelinesQueryParams) SetActivated(v bool) ListPipelinesQueryParams
```

#### func (ListPipelinesQueryParams) SetCreateUserId

```go
func (q ListPipelinesQueryParams) SetCreateUserId(v string) ListPipelinesQueryParams
```

#### func (ListPipelinesQueryParams) SetIncludeData

```go
func (q ListPipelinesQueryParams) SetIncludeData(v bool) ListPipelinesQueryParams
```

#### func (ListPipelinesQueryParams) SetName

```go
func (q ListPipelinesQueryParams) SetName(v string) ListPipelinesQueryParams
```

#### func (ListPipelinesQueryParams) SetOffset

```go
func (q ListPipelinesQueryParams) SetOffset(v int32) ListPipelinesQueryParams
```

#### func (ListPipelinesQueryParams) SetPageSize

```go
func (q ListPipelinesQueryParams) SetPageSize(v int32) ListPipelinesQueryParams
```

#### func (ListPipelinesQueryParams) SetSortDir

```go
func (q ListPipelinesQueryParams) SetSortDir(v string) ListPipelinesQueryParams
```

#### func (ListPipelinesQueryParams) SetSortField

```go
func (q ListPipelinesQueryParams) SetSortField(v string) ListPipelinesQueryParams
```

#### type ListTemplatesQueryParams

```go
type ListTemplatesQueryParams struct {
	// Offset : offset
	Offset *int32 `key:"offset"`
	// PageSize : pageSize
	PageSize *int32 `key:"pageSize"`
	// SortDir : sortDir
	SortDir string `key:"sortDir"`
	// SortField : sortField
	SortField string `key:"sortField"`
}
```

ListTemplatesQueryParams represents valid query parameters for the ListTemplates
operation For convenience ListTemplatesQueryParams can be formed in a single
statement, for example:

    `v := ListTemplatesQueryParams{}.SetOffset(...).SetPageSize(...).SetSortDir(...).SetSortField(...)`

#### func (ListTemplatesQueryParams) SetOffset

```go
func (q ListTemplatesQueryParams) SetOffset(v int32) ListTemplatesQueryParams
```

#### func (ListTemplatesQueryParams) SetPageSize

```go
func (q ListTemplatesQueryParams) SetPageSize(v int32) ListTemplatesQueryParams
```

#### func (ListTemplatesQueryParams) SetSortDir

```go
func (q ListTemplatesQueryParams) SetSortDir(v string) ListTemplatesQueryParams
```

#### func (ListTemplatesQueryParams) SetSortField

```go
func (q ListTemplatesQueryParams) SetSortField(v string) ListTemplatesQueryParams
```

#### type MapOfstringAndobject

```go
type MapOfstringAndobject map[string]interface{}
```


#### type MetricsResponse

```go
type MetricsResponse struct {
	Nodes map[string]NodeMetrics `json:"nodes,omitempty"`
}
```


#### type NodeMetrics

```go
type NodeMetrics struct {
	Metrics map[string]interface{} `json:"metrics,omitempty"`
}
```


#### type ObjectNode

```go
type ObjectNode struct {
	Array               *bool               `json:"array,omitempty"`
	BigDecimal          *bool               `json:"bigDecimal,omitempty"`
	BigInteger          *bool               `json:"bigInteger,omitempty"`
	Binary              *bool               `json:"binary,omitempty"`
	Boolean             *bool               `json:"boolean,omitempty"`
	ContainerNode       *bool               `json:"containerNode,omitempty"`
	Double              *bool               `json:"double,omitempty"`
	Float               *bool               `json:"float,omitempty"`
	FloatingPointNumber *bool               `json:"floatingPointNumber,omitempty"`
	Int                 *bool               `json:"int,omitempty"`
	IntegralNumber      *bool               `json:"integralNumber,omitempty"`
	Long                *bool               `json:"long,omitempty"`
	MissingNode         *bool               `json:"missingNode,omitempty"`
	NodeType            *ObjectNodeNodeType `json:"nodeType,omitempty"`
	Null                *bool               `json:"null,omitempty"`
	Number              *bool               `json:"number,omitempty"`
	Object              *bool               `json:"object,omitempty"`
	Pojo                *bool               `json:"pojo,omitempty"`
	Short               *bool               `json:"short,omitempty"`
	Textual             *bool               `json:"textual,omitempty"`
	ValueNode           *bool               `json:"valueNode,omitempty"`
}
```


#### type ObjectNodeNodeType

```go
type ObjectNodeNodeType string
```


```go
const (
	ObjectNodeNodeTypeArray   ObjectNodeNodeType = "ARRAY"
	ObjectNodeNodeTypeBinary  ObjectNodeNodeType = "BINARY"
	ObjectNodeNodeTypeBoolean ObjectNodeNodeType = "BOOLEAN"
	ObjectNodeNodeTypeMissing ObjectNodeNodeType = "MISSING"
	ObjectNodeNodeTypeNull    ObjectNodeNodeType = "NULL"
	ObjectNodeNodeTypeNumber  ObjectNodeNodeType = "NUMBER"
	ObjectNodeNodeTypeObject  ObjectNodeNodeType = "OBJECT"
	ObjectNodeNodeTypePojo    ObjectNodeNodeType = "POJO"
	ObjectNodeNodeTypeString  ObjectNodeNodeType = "STRING"
)
```
List of ObjectNodeNodeType

#### type PaginatedResponseOfConnectionResponse

```go
type PaginatedResponseOfConnectionResponse struct {
	Items []ConnectionResponse `json:"items,omitempty"`
	Total *int64               `json:"total,omitempty"`
}
```


#### type PaginatedResponseOfConnectorResponse

```go
type PaginatedResponseOfConnectorResponse struct {
	Items []ConnectorResponse `json:"items,omitempty"`
	Total *int64              `json:"total,omitempty"`
}
```


#### type PaginatedResponseOfPipelineJobStatus

```go
type PaginatedResponseOfPipelineJobStatus struct {
	Items []PipelineJobStatus `json:"items,omitempty"`
	Total *int64              `json:"total,omitempty"`
}
```


#### type PaginatedResponseOfPipelineResponse

```go
type PaginatedResponseOfPipelineResponse struct {
	Items []PipelineResponse `json:"items,omitempty"`
	Total *int64             `json:"total,omitempty"`
}
```


#### type PaginatedResponseOfTemplateResponse

```go
type PaginatedResponseOfTemplateResponse struct {
	Items []TemplateResponse `json:"items,omitempty"`
	Total *int64             `json:"total,omitempty"`
}
```


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

#### type PipelineDeleteResponse

```go
type PipelineDeleteResponse struct {
	CouldDeactivate *bool `json:"couldDeactivate,omitempty"`
	Running         *bool `json:"running,omitempty"`
}
```


#### type PipelineJobStatus

```go
type PipelineJobStatus struct {
	JobId      *string `json:"jobId,omitempty"`
	JobStatus  *string `json:"jobStatus,omitempty"`
	PipelineId *string `json:"pipelineId,omitempty"`
}
```


#### type PipelinePatchRequest

```go
type PipelinePatchRequest struct {
	// Set to true to bypass initial pipeline validation upon creation. The pipeline still needs to be validated before activation. Defaults to false.
	BypassValidation *bool `json:"bypassValidation,omitempty"`
	// The user that created the pipeline. Deprecated.
	CreateUserId *string      `json:"createUserId,omitempty"`
	Data         *UplPipeline `json:"data,omitempty"`
	// The description of the pipeline. Defaults to null.
	Description *string `json:"description,omitempty"`
	// The name of the pipeline.
	Name *string `json:"name,omitempty"`
}
```


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
	CurrentlyActiveVersion     *int64                                                `json:"currentlyActiveVersion,omitempty"`
	PipelineId                 *string                                               `json:"pipelineId,omitempty"`
	PipelineReactivationStatus *PipelineReactivateResponsePipelineReactivationStatus `json:"pipelineReactivationStatus,omitempty"`
}
```


#### type PipelineReactivateResponsePipelineReactivationStatus

```go
type PipelineReactivateResponsePipelineReactivationStatus string
```


```go
const (
	PipelineReactivateResponsePipelineReactivationStatusActivated                          PipelineReactivateResponsePipelineReactivationStatus = "activated"
	PipelineReactivateResponsePipelineReactivationStatusAlreadyActivatedWithCurrentVersion PipelineReactivateResponsePipelineReactivationStatus = "alreadyActivatedWithCurrentVersion"
	PipelineReactivateResponsePipelineReactivationStatusCurrentVersionInvalid              PipelineReactivateResponsePipelineReactivationStatus = "currentVersionInvalid"
	PipelineReactivateResponsePipelineReactivationStatusFailedToDeactivateCurrentVersion   PipelineReactivateResponsePipelineReactivationStatus = "failedToDeactivateCurrentVersion"
	PipelineReactivateResponsePipelineReactivationStatusRolledBack                         PipelineReactivateResponsePipelineReactivationStatus = "rolledBack"
	PipelineReactivateResponsePipelineReactivationStatusRolledBackError                    PipelineReactivateResponsePipelineReactivationStatus = "rolledBackError"
)
```
List of PipelineReactivateResponsePipelineReactivationStatus

#### type PipelineRequest

```go
type PipelineRequest struct {
	Data UplPipeline `json:"data"`
	// The name of the pipeline.
	Name string `json:"name"`
	// Set to true to bypass initial pipeline validation upon creation. The pipeline still needs to be validated before activation. Defaults to false.
	BypassValidation *bool `json:"bypassValidation,omitempty"`
	// The user that created the pipeline. Deprecated.
	CreateUserId *string `json:"createUserId,omitempty"`
	// The description of the pipeline. Defaults to null.
	Description *string `json:"description,omitempty"`
}
```


#### type PipelineResponse

```go
type PipelineResponse struct {
	ActivatedDate            *int64                  `json:"activatedDate,omitempty"`
	ActivatedUserId          *string                 `json:"activatedUserId,omitempty"`
	ActivatedVersion         *int64                  `json:"activatedVersion,omitempty"`
	CreateDate               *int64                  `json:"createDate,omitempty"`
	CreateUserId             *string                 `json:"createUserId,omitempty"`
	CurrentVersion           *int64                  `json:"currentVersion,omitempty"`
	Data                     *UplPipeline            `json:"data,omitempty"`
	Description              *string                 `json:"description,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	LastUpdateDate           *int64                  `json:"lastUpdateDate,omitempty"`
	LastUpdateUserId         *string                 `json:"lastUpdateUserId,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	Status                   *PipelineResponseStatus `json:"status,omitempty"`
	StatusMessage            *string                 `json:"statusMessage,omitempty"`
	StreamingConfigurationId *int64                  `json:"streamingConfigurationId,omitempty"`
	TenantId                 *string                 `json:"tenantId,omitempty"`
	ValidationMessages       []string                `json:"validationMessages,omitempty"`
	Version                  *int64                  `json:"version,omitempty"`
}
```


#### type PipelineResponseStatus

```go
type PipelineResponseStatus string
```


```go
const (
	PipelineResponseStatusCreated   PipelineResponseStatus = "CREATED"
	PipelineResponseStatusActivated PipelineResponseStatus = "ACTIVATED"
)
```
List of PipelineResponseStatus

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
	InputTree UplPipeline `json:"inputTree"`
	MainTree  UplPipeline `json:"mainTree"`
	// The function ID of the merge target in the main pipeline.
	TargetNode string `json:"targetNode"`
	// The input port of the merge target in the main pipeline.
	TargetPort string `json:"targetPort"`
}
```


#### type PreviewData

```go
type PreviewData struct {
	CurrentNumberOfRecords *int32                 `json:"currentNumberOfRecords,omitempty"`
	Nodes                  map[string]PreviewNode `json:"nodes,omitempty"`
	PreviewId              *string                `json:"previewId,omitempty"`
	RecordsPerPipeline     *int32                 `json:"recordsPerPipeline,omitempty"`
	TenantId               *string                `json:"tenantId,omitempty"`
}
```


#### type PreviewNode

```go
type PreviewNode struct {
	NodeName *string      `json:"nodeName,omitempty"`
	Records  []ObjectNode `json:"records,omitempty"`
}
```


#### type PreviewSessionStartRequest

```go
type PreviewSessionStartRequest struct {
	Upl UplPipeline `json:"upl"`
	// The maximum number of events per function. Defaults to 100.
	RecordsLimit *int32 `json:"recordsLimit,omitempty"`
	// The maximum number of events per pipeline. Defaults to 10000.
	RecordsPerPipeline *int32 `json:"recordsPerPipeline,omitempty"`
	// The maximum lifetime of a session, in milliseconds. Defaults to 300,000.
	SessionLifetimeMs *int64 `json:"sessionLifetimeMs,omitempty"`
	// Deprecated. Must be null if set.
	StreamingConfigurationId *int64 `json:"streamingConfigurationId,omitempty"`
	// Deprecated. Must be true if set.
	UseNewData *bool `json:"useNewData,omitempty"`
}
```


#### type PreviewStartResponse

```go
type PreviewStartResponse struct {
	PreviewId *int64 `json:"previewId,omitempty"`
}
```


#### type PreviewState

```go
type PreviewState struct {
	ActivatedDate          *int64  `json:"activatedDate,omitempty"`
	CreatedDate            *int64  `json:"createdDate,omitempty"`
	CurrentNumberOfRecords *int32  `json:"currentNumberOfRecords,omitempty"`
	JobId                  *string `json:"jobId,omitempty"`
	PreviewId              *int64  `json:"previewId,omitempty"`
	RecordsPerPipeline     *int32  `json:"recordsPerPipeline,omitempty"`
}
```


#### type Response

```go
type Response struct {
	// Only set for /activate endpoint
	Activated *string `json:"activated,omitempty"`
	// Only set for /deactivate endpoint
	Deactivated *string `json:"deactivated,omitempty"`
}
```


#### type Service

```go
type Service services.BaseService
```


#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new streams service client from the given Config

#### func (*Service) ActivatePipeline

```go
func (s *Service) ActivatePipeline(id string, activatePipelineRequest ActivatePipelineRequest, resp ...*http.Response) (*Response, error)
```
ActivatePipeline - Activates an existing pipeline. Parameters:

    id: id of the pipeline to activate
    activatePipelineRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CompileDSL

```go
func (s *Service) CompileDSL(dslCompilationRequest DslCompilationRequest, resp ...*http.Response) (*UplPipeline, error)
```
CompileDSL - Compiles the Streams DSL and returns Streams JSON. Parameters:

    dslCompilationRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateConnection

```go
func (s *Service) CreateConnection(connectionRequest ConnectionRequest, resp ...*http.Response) (*ConnectionSaveResponse, error)
```
CreateConnection - Create a new DSP connection. Parameters:

    connectionRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateGroup

```go
func (s *Service) CreateGroup(groupRequest GroupRequest, resp ...*http.Response) (*GroupResponse, error)
```
CreateGroup - Create a new group function by combining the Streams JSON of two
or more functions. Parameters:

    groupRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreatePipeline

```go
func (s *Service) CreatePipeline(pipelineRequest PipelineRequest, resp ...*http.Response) (*PipelineResponse, error)
```
CreatePipeline - Creates a pipeline. Parameters:

    pipelineRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateTemplate

```go
func (s *Service) CreateTemplate(templateRequest TemplateRequest, resp ...*http.Response) (*TemplateResponse, error)
```
CreateTemplate - Creates a template for a tenant. Parameters:

    templateRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeactivatePipeline

```go
func (s *Service) DeactivatePipeline(id string, deactivatePipelineRequest DeactivatePipelineRequest, resp ...*http.Response) (*Response, error)
```
DeactivatePipeline - Deactivates an existing pipeline. Parameters:

    id: id of the pipeline to deactivate
    deactivatePipelineRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteConnection

```go
func (s *Service) DeleteConnection(connectionId string, resp ...*http.Response) error
```
DeleteConnection - Delete all versions of a connection by its id. Parameters:

    connectionId: ID of the connection
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteGroup

```go
func (s *Service) DeleteGroup(groupId string, resp ...*http.Response) error
```
DeleteGroup - Removes an existing group. Parameters:

    groupId: The group function's ID from the function registry
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeletePipeline

```go
func (s *Service) DeletePipeline(id string, resp ...*http.Response) (*PipelineDeleteResponse, error)
```
DeletePipeline - Removes a pipeline. Parameters:

    id: id of the pipeline to delete
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteTemplate

```go
func (s *Service) DeleteTemplate(templateId string, resp ...*http.Response) error
```
DeleteTemplate - Removes a template with a specific ID. Parameters:

    templateId: ID of the template to delete
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ExpandGroup

```go
func (s *Service) ExpandGroup(groupId string, groupExpandRequest GroupExpandRequest, resp ...*http.Response) (*UplPipeline, error)
```
ExpandGroup - Creates and returns the expanded version of a group. Parameters:

    groupId: The group function's ID from the function registry
    groupExpandRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ExpandPipeline

```go
func (s *Service) ExpandPipeline(uplPipeline UplPipeline, resp ...*http.Response) (*UplPipeline, error)
```
ExpandPipeline - Returns the entire Streams JSON, including the expanded Streams
JSON of any group functions in the pipeline. Parameters:

    uplPipeline: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetGroup

```go
func (s *Service) GetGroup(groupId string, resp ...*http.Response) (*GroupResponse, error)
```
GetGroup - Returns the full Streams JSON of a group. Parameters:

    groupId: The group function's ID from the function registry
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetInputSchema

```go
func (s *Service) GetInputSchema(getInputSchemaRequest GetInputSchemaRequest, resp ...*http.Response) (*UplType, error)
```
GetInputSchema - Returns the input schema for a function in a pipeline.
Parameters:

    getInputSchemaRequest: Input Schema Request
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetOutputSchema

```go
func (s *Service) GetOutputSchema(getOutputSchemaRequest GetOutputSchemaRequest, resp ...*http.Response) (map[string]UplType, error)
```
GetOutputSchema - Returns the output schema for a specified function in a
pipeline. If no function ID is specified, the request returns the output schema
for all functions in a pipeline. Parameters:

    getOutputSchemaRequest: Output Schema Request
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetPipeline

```go
func (s *Service) GetPipeline(id string, query *GetPipelineQueryParams, resp ...*http.Response) (*PipelineResponse, error)
```
GetPipeline - Returns an individual pipeline by version. Parameters:

    id: id of the pipeline to get
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetPipelineLatestMetrics

```go
func (s *Service) GetPipelineLatestMetrics(id string, resp ...*http.Response) (*MetricsResponse, error)
```
GetPipelineLatestMetrics - Returns the latest metrics for a single pipeline.
Parameters:

    id: ID of the pipeline to get metrics for
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetPipelinesStatus

```go
func (s *Service) GetPipelinesStatus(query *GetPipelinesStatusQueryParams, resp ...*http.Response) (*PaginatedResponseOfPipelineJobStatus, error)
```
GetPipelinesStatus - Returns the status of pipelines from the underlying
streaming system. Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetPreviewData

```go
func (s *Service) GetPreviewData(previewSessionId int64, resp ...*http.Response) (*PreviewData, error)
```
GetPreviewData - Returns the preview data for a session. Parameters:

    previewSessionId: ID of the preview session
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetPreviewSession

```go
func (s *Service) GetPreviewSession(previewSessionId int64, resp ...*http.Response) (*PreviewState, error)
```
GetPreviewSession - Returns information from a preview session. Parameters:

    previewSessionId: ID of the preview session
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetPreviewSessionLatestMetrics

```go
func (s *Service) GetPreviewSessionLatestMetrics(previewSessionId int64, resp ...*http.Response) (*MetricsResponse, error)
```
GetPreviewSessionLatestMetrics - Returns the latest metrics for a preview
session. Parameters:

    previewSessionId: ID of the preview session
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetRegistry

```go
func (s *Service) GetRegistry(query *GetRegistryQueryParams, resp ...*http.Response) (*UplRegistry, error)
```
GetRegistry - Returns all functions in JSON format. Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetTemplate

```go
func (s *Service) GetTemplate(templateId string, query *GetTemplateQueryParams, resp ...*http.Response) (*TemplateResponse, error)
```
GetTemplate - Returns an individual template by version. Parameters:

    templateId: ID of the template
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListConnections

```go
func (s *Service) ListConnections(query *ListConnectionsQueryParams, resp ...*http.Response) (*PaginatedResponseOfConnectionResponse, error)
```
ListConnections - Returns a list of connections (latest versions only) by tenant
ID. Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListConnectors

```go
func (s *Service) ListConnectors(resp ...*http.Response) (*PaginatedResponseOfConnectorResponse, error)
```
ListConnectors - Returns a list of the available connectors. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListPipelines

```go
func (s *Service) ListPipelines(query *ListPipelinesQueryParams, resp ...*http.Response) (*PaginatedResponseOfPipelineResponse, error)
```
ListPipelines - Returns all pipelines. Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListTemplates

```go
func (s *Service) ListTemplates(query *ListTemplatesQueryParams, resp ...*http.Response) (*PaginatedResponseOfTemplateResponse, error)
```
ListTemplates - Returns a list of all templates. Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) MergePipelines

```go
func (s *Service) MergePipelines(pipelinesMergeRequest PipelinesMergeRequest, resp ...*http.Response) (*UplPipeline, error)
```
MergePipelines - Combines two Streams JSON programs. Parameters:

    pipelinesMergeRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) PutConnection

```go
func (s *Service) PutConnection(connectionId string, connectionPutRequest ConnectionPutRequest, resp ...*http.Response) (*ConnectionSaveResponse, error)
```
PutConnection - Modifies an existing DSP connection. Parameters:

    connectionId: ID of the connection
    connectionPutRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) PutGroup

```go
func (s *Service) PutGroup(groupId string, groupPutRequest GroupPutRequest, resp ...*http.Response) (*GroupResponse, error)
```
PutGroup - Update a group function combining the Streams JSON of two or more
functions. Parameters:

    groupId: The group function's ID from the function registry
    groupPutRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) PutTemplate

```go
func (s *Service) PutTemplate(templateId string, templatePutRequest TemplatePutRequest, resp ...*http.Response) (*TemplateResponse, error)
```
PutTemplate - Updates an existing template. Parameters:

    templateId: ID of the template
    templatePutRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ReactivatePipeline

```go
func (s *Service) ReactivatePipeline(id string, resp ...*http.Response) (*PipelineReactivateResponse, error)
```
ReactivatePipeline - Reactivate a pipeline Parameters:

    id: Pipeline UUID to reactivate
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) StartPreview

```go
func (s *Service) StartPreview(previewSessionStartRequest PreviewSessionStartRequest, resp ...*http.Response) (*PreviewStartResponse, error)
```
StartPreview - Creates a preview session for a pipeline. Parameters:

    previewSessionStartRequest: Parameters to start a new Preview session
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) StopPreview

```go
func (s *Service) StopPreview(previewSessionId int64, resp ...*http.Response) (*string, error)
```
StopPreview - Stops a preview session. Parameters:

    previewSessionId: ID of the preview session
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateConnection

```go
func (s *Service) UpdateConnection(connectionId string, connectionPatchRequest ConnectionPatchRequest, resp ...*http.Response) (*ConnectionSaveResponse, error)
```
UpdateConnection - Partially modifies an existing DSP connection. Parameters:

    connectionId: ID of the connection
    connectionPatchRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateGroup

```go
func (s *Service) UpdateGroup(groupId string, groupPatchRequest GroupPatchRequest, resp ...*http.Response) (*GroupResponse, error)
```
UpdateGroup - Modify a group function by combining the Streams JSON of two or
more functions. Parameters:

    groupId: The group function's ID from the function registry
    groupPatchRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdatePipeline

```go
func (s *Service) UpdatePipeline(id string, pipelinePatchRequest PipelinePatchRequest, resp ...*http.Response) (*PipelineResponse, error)
```
UpdatePipeline - Partially modifies an existing pipeline. Parameters:

    id: id of the pipeline to update
    pipelinePatchRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateTemplate

```go
func (s *Service) UpdateTemplate(templateId string, templatePatchRequest TemplatePatchRequest, resp ...*http.Response) (*TemplateResponse, error)
```
UpdateTemplate - Partially modifies an existing template. Parameters:

    templateId: ID of the template
    templatePatchRequest: Request JSON
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ValidatePipeline

```go
func (s *Service) ValidatePipeline(validateRequest ValidateRequest, resp ...*http.Response) (*ValidateResponse, error)
```
ValidatePipeline - Verifies whether the Streams JSON is valid. Parameters:

    validateRequest: JSON UPL to validate
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### type Servicer

```go
type Servicer interface {
	/*
		ActivatePipeline - Activates an existing pipeline.
		Parameters:
			id: id of the pipeline to activate
			activatePipelineRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ActivatePipeline(id string, activatePipelineRequest ActivatePipelineRequest, resp ...*http.Response) (*Response, error)
	/*
		CompileDSL - Compiles the Streams DSL and returns Streams JSON.
		Parameters:
			dslCompilationRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CompileDSL(dslCompilationRequest DslCompilationRequest, resp ...*http.Response) (*UplPipeline, error)
	/*
		CreateConnection - Create a new DSP connection.
		Parameters:
			connectionRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateConnection(connectionRequest ConnectionRequest, resp ...*http.Response) (*ConnectionSaveResponse, error)
	/*
		CreateGroup - Create a new group function by combining the Streams JSON of two or more functions.
		Parameters:
			groupRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateGroup(groupRequest GroupRequest, resp ...*http.Response) (*GroupResponse, error)
	/*
		CreatePipeline - Creates a pipeline.
		Parameters:
			pipelineRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreatePipeline(pipelineRequest PipelineRequest, resp ...*http.Response) (*PipelineResponse, error)
	/*
		CreateTemplate - Creates a template for a tenant.
		Parameters:
			templateRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateTemplate(templateRequest TemplateRequest, resp ...*http.Response) (*TemplateResponse, error)
	/*
		DeactivatePipeline - Deactivates an existing pipeline.
		Parameters:
			id: id of the pipeline to deactivate
			deactivatePipelineRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeactivatePipeline(id string, deactivatePipelineRequest DeactivatePipelineRequest, resp ...*http.Response) (*Response, error)
	/*
		DeleteConnection - Delete all versions of a connection by its id.
		Parameters:
			connectionId: ID of the connection
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteConnection(connectionId string, resp ...*http.Response) error
	/*
		DeleteGroup - Removes an existing group.
		Parameters:
			groupId: The group function's ID from the function registry
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteGroup(groupId string, resp ...*http.Response) error
	/*
		DeletePipeline - Removes a pipeline.
		Parameters:
			id: id of the pipeline to delete
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeletePipeline(id string, resp ...*http.Response) (*PipelineDeleteResponse, error)
	/*
		DeleteTemplate - Removes a template with a specific ID.
		Parameters:
			templateId: ID of the template to delete
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteTemplate(templateId string, resp ...*http.Response) error
	/*
		ExpandGroup - Creates and returns the expanded version of a group.
		Parameters:
			groupId: The group function's ID from the function registry
			groupExpandRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ExpandGroup(groupId string, groupExpandRequest GroupExpandRequest, resp ...*http.Response) (*UplPipeline, error)
	/*
		ExpandPipeline - Returns the entire Streams JSON, including the expanded Streams JSON of any group functions in the pipeline.
		Parameters:
			uplPipeline: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ExpandPipeline(uplPipeline UplPipeline, resp ...*http.Response) (*UplPipeline, error)
	/*
		GetGroup - Returns the full Streams JSON of a group.
		Parameters:
			groupId: The group function's ID from the function registry
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetGroup(groupId string, resp ...*http.Response) (*GroupResponse, error)
	/*
		GetInputSchema - Returns the input schema for a function in a pipeline.
		Parameters:
			getInputSchemaRequest: Input Schema Request
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetInputSchema(getInputSchemaRequest GetInputSchemaRequest, resp ...*http.Response) (*UplType, error)
	/*
		GetOutputSchema - Returns the output schema for a specified function in a pipeline. If no function ID is  specified, the request returns the output schema for all functions in a pipeline.
		Parameters:
			getOutputSchemaRequest: Output Schema Request
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetOutputSchema(getOutputSchemaRequest GetOutputSchemaRequest, resp ...*http.Response) (map[string]UplType, error)
	/*
		GetPipeline - Returns an individual pipeline by version.
		Parameters:
			id: id of the pipeline to get
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPipeline(id string, query *GetPipelineQueryParams, resp ...*http.Response) (*PipelineResponse, error)
	/*
		GetPipelineLatestMetrics - Returns the latest metrics for a single pipeline.
		Parameters:
			id: ID of the pipeline to get metrics for
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPipelineLatestMetrics(id string, resp ...*http.Response) (*MetricsResponse, error)
	/*
		GetPipelinesStatus - Returns the status of pipelines from the underlying streaming system.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPipelinesStatus(query *GetPipelinesStatusQueryParams, resp ...*http.Response) (*PaginatedResponseOfPipelineJobStatus, error)
	/*
		GetPreviewData - Returns the preview data for a session.
		Parameters:
			previewSessionId: ID of the preview session
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPreviewData(previewSessionId int64, resp ...*http.Response) (*PreviewData, error)
	/*
		GetPreviewSession - Returns information from a preview session.
		Parameters:
			previewSessionId: ID of the preview session
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPreviewSession(previewSessionId int64, resp ...*http.Response) (*PreviewState, error)
	/*
		GetPreviewSessionLatestMetrics - Returns the latest metrics for a preview session.
		Parameters:
			previewSessionId: ID of the preview session
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPreviewSessionLatestMetrics(previewSessionId int64, resp ...*http.Response) (*MetricsResponse, error)
	/*
		GetRegistry - Returns all functions in JSON format.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetRegistry(query *GetRegistryQueryParams, resp ...*http.Response) (*UplRegistry, error)
	/*
		GetTemplate - Returns an individual template by version.
		Parameters:
			templateId: ID of the template
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetTemplate(templateId string, query *GetTemplateQueryParams, resp ...*http.Response) (*TemplateResponse, error)
	/*
		ListConnections - Returns a list of connections (latest versions only) by tenant ID.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListConnections(query *ListConnectionsQueryParams, resp ...*http.Response) (*PaginatedResponseOfConnectionResponse, error)
	/*
		ListConnectors - Returns a list of the available connectors.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListConnectors(resp ...*http.Response) (*PaginatedResponseOfConnectorResponse, error)
	/*
		ListPipelines - Returns all pipelines.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListPipelines(query *ListPipelinesQueryParams, resp ...*http.Response) (*PaginatedResponseOfPipelineResponse, error)
	/*
		ListTemplates - Returns a list of all templates.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListTemplates(query *ListTemplatesQueryParams, resp ...*http.Response) (*PaginatedResponseOfTemplateResponse, error)
	/*
		MergePipelines - Combines two Streams JSON programs.
		Parameters:
			pipelinesMergeRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	MergePipelines(pipelinesMergeRequest PipelinesMergeRequest, resp ...*http.Response) (*UplPipeline, error)
	/*
		PutConnection - Modifies an existing DSP connection.
		Parameters:
			connectionId: ID of the connection
			connectionPutRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PutConnection(connectionId string, connectionPutRequest ConnectionPutRequest, resp ...*http.Response) (*ConnectionSaveResponse, error)
	/*
		PutGroup - Update a group function combining the Streams JSON of two or more functions.
		Parameters:
			groupId: The group function's ID from the function registry
			groupPutRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PutGroup(groupId string, groupPutRequest GroupPutRequest, resp ...*http.Response) (*GroupResponse, error)
	/*
		PutTemplate - Updates an existing template.
		Parameters:
			templateId: ID of the template
			templatePutRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PutTemplate(templateId string, templatePutRequest TemplatePutRequest, resp ...*http.Response) (*TemplateResponse, error)
	/*
		ReactivatePipeline - Reactivate a pipeline
		Parameters:
			id: Pipeline UUID to reactivate
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ReactivatePipeline(id string, resp ...*http.Response) (*PipelineReactivateResponse, error)
	/*
		StartPreview - Creates a preview session for a pipeline.
		Parameters:
			previewSessionStartRequest: Parameters to start a new Preview session
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	StartPreview(previewSessionStartRequest PreviewSessionStartRequest, resp ...*http.Response) (*PreviewStartResponse, error)
	/*
		StopPreview - Stops a preview session.
		Parameters:
			previewSessionId: ID of the preview session
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	StopPreview(previewSessionId int64, resp ...*http.Response) (*string, error)
	/*
		UpdateConnection - Partially modifies an existing DSP connection.
		Parameters:
			connectionId: ID of the connection
			connectionPatchRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateConnection(connectionId string, connectionPatchRequest ConnectionPatchRequest, resp ...*http.Response) (*ConnectionSaveResponse, error)
	/*
		UpdateGroup - Modify a group function by combining the Streams JSON of two or more functions.
		Parameters:
			groupId: The group function's ID from the function registry
			groupPatchRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateGroup(groupId string, groupPatchRequest GroupPatchRequest, resp ...*http.Response) (*GroupResponse, error)
	/*
		UpdatePipeline - Partially modifies an existing pipeline.
		Parameters:
			id: id of the pipeline to update
			pipelinePatchRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdatePipeline(id string, pipelinePatchRequest PipelinePatchRequest, resp ...*http.Response) (*PipelineResponse, error)
	/*
		UpdateTemplate - Partially modifies an existing template.
		Parameters:
			templateId: ID of the template
			templatePatchRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateTemplate(templateId string, templatePatchRequest TemplatePatchRequest, resp ...*http.Response) (*TemplateResponse, error)
	/*
		ValidatePipeline - Verifies whether the Streams JSON is valid.
		Parameters:
			validateRequest: JSON UPL to validate
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ValidatePipeline(validateRequest ValidateRequest, resp ...*http.Response) (*ValidateResponse, error)
}
```

Servicer represents the interface for implementing all endpoints for this
service

#### type TemplatePatchRequest

```go
type TemplatePatchRequest struct {
	Data *UplPipeline `json:"data,omitempty"`
	// Template description
	Description *string `json:"description,omitempty"`
	// Template name
	Name *string `json:"name,omitempty"`
}
```


#### type TemplatePutRequest

```go
type TemplatePutRequest struct {
	Data UplPipeline `json:"data"`
	// Template description
	Description string `json:"description"`
	// Template name
	Name string `json:"name"`
}
```


#### type TemplateRequest

```go
type TemplateRequest struct {
	Data UplPipeline `json:"data"`
	// Template description
	Description string `json:"description"`
	// Template name
	Name string `json:"name"`
}
```


#### type TemplateResponse

```go
type TemplateResponse struct {
	CreateDate    *int64       `json:"createDate,omitempty"`
	CreateUserId  *string      `json:"createUserId,omitempty"`
	Data          *UplPipeline `json:"data,omitempty"`
	Description   *string      `json:"description,omitempty"`
	Name          *string      `json:"name,omitempty"`
	OwnerTenantId *string      `json:"ownerTenantId,omitempty"`
	TemplateId    *string      `json:"templateId,omitempty"`
	Version       *int64       `json:"version,omitempty"`
}
```


#### type UplArgument

```go
type UplArgument struct {
	Type        string      `json:"type"`
	ElementType interface{} `json:"elementType,omitempty"`
}
```


#### type UplCategory

```go
type UplCategory struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
```


#### type UplEdge

```go
type UplEdge struct {
	// The source function's (node's) id
	SourceNode string `json:"sourceNode"`
	// The source function's (node's) port
	SourcePort string `json:"sourcePort"`
	// The target function's (node's) id
	TargetNode string `json:"targetNode"`
	// The target function's (node's) port
	TargetPort string                 `json:"targetPort"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}
```


#### type UplFunction

```go
type UplFunction struct {
	Arguments  map[string]UplArgument `json:"arguments,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Categories []int64                `json:"categories,omitempty"`
	Id         *string                `json:"id,omitempty"`
	IsVariadic *bool                  `json:"isVariadic,omitempty"`
	Op         *string                `json:"op,omitempty"`
	Output     *UplArgument           `json:"output,omitempty"`
	ResolvedId *string                `json:"resolvedId,omitempty"`
}
```


#### type UplNode

```go
type UplNode struct {
	// The function's (node's) UUID
	Id string `json:"id"`
	// The function's ID or its API name
	Op string `json:"op"`
	// Optional key-value pair for a function (node)
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	ResolvedId *string                `json:"resolvedId,omitempty"`
}
```


#### type UplPipeline

```go
type UplPipeline struct {
	// A list of links or connections between the output of one pipeline function and the input of another pipeline function
	Edges []UplEdge `json:"edges"`
	// The functions (or nodes) in your entire pipeline, including each function's operations, attributes, and properties
	Nodes interface{} `json:"nodes"`
	// The UUIDs of all sink functions in a given pipeline
	RootNode []string `json:"rootNode"`
}
```


#### type UplRegistry

```go
type UplRegistry struct {
	Categories []UplCategory `json:"categories,omitempty"`
	Functions  []UplFunction `json:"functions,omitempty"`
	Types      []UplType     `json:"types,omitempty"`
}
```


#### type UplType

```go
type UplType struct {
	FieldName  *string   `json:"fieldName,omitempty"`
	Parameters []UplType `json:"parameters,omitempty"`
	Type       *string   `json:"type,omitempty"`
}
```


#### type ValidateRequest

```go
type ValidateRequest struct {
	Upl UplPipeline `json:"upl"`
}
```


#### type ValidateResponse

```go
type ValidateResponse struct {
	Success            *bool    `json:"success,omitempty"`
	ValidationMessages []string `json:"validationMessages,omitempty"`
}
```
