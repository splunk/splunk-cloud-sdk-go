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

#### type DslCompilationRequest

```go
type DslCompilationRequest struct {
	Dsl string `json:"dsl"`
}
```

DslCompilationRequest contains the DSL that needs to be compiled into a valid
UPL JSON

#### type PaginatedPipelineResponse

```go
type PaginatedPipelineResponse struct {
	Items []Pipeline `json:"items"`
	Total int64      `json:"total"`
}
```

PaginatedPipelineResponse contains the pipeline response

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

#### func (*Service) CreatePipeline

```go
func (s *Service) CreatePipeline(pipeline *PipelineRequest) (*Pipeline, error)
```
CreatePipeline creates a new pipeline

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

#### func (*Service) GetPipeline

```go
func (s *Service) GetPipeline(id string) (*Pipeline, error)
```
GetPipeline gets an individual pipeline

#### func (*Service) GetPipelines

```go
func (s *Service) GetPipelines(queryParams PipelineQueryParams) (*PaginatedPipelineResponse, error)
```
GetPipelines gets all the pipelines

#### func (*Service) UpdatePipeline

```go
func (s *Service) UpdatePipeline(id string, pipeline *PipelineRequest) (*Pipeline, error)
```
UpdatePipeline updates an existing pipeline

#### type Servicer

```go
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
```

Servicer ...

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
