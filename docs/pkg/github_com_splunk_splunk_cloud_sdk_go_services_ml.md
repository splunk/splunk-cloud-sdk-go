# ml
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/ml"


## Usage

#### type Fields

```go
type Fields struct {
	// Fields necessary for task.
	Features []string `json:"features"`
	// Fields produced by task.
	Created *[]string `json:"created,omitempty"`
	// Target field necessary for task.
	Target *string `json:"target,omitempty"`
}
```


#### type FitTask

```go
type FitTask struct {
	// The name has to be unique in the same workflow, it is optional, can be used to identify that task artifact
	Name              *string   `json:"name,omitempty"`
	Kind              *TaskKind `json:"kind,omitempty"`
	Algorithm         string    `json:"algorithm"`
	Fields            Fields    `json:"fields"`
	OutputTransformer *string   `json:"outputTransformer,omitempty"`
	// time out in seconds
	TimeoutSecs int32              `json:"timeoutSecs"`
	Parameters  *map[string]string `json:"parameters,omitempty"`
}
```

Fit task does an estimation/training/fitting, fit task outputs a transformer

#### type Health

```go
type Health struct {
	Healthy bool   `json:"healthy"`
	Version string `json:"version"`
}
```


#### type Hec

```go
type Hec struct {
	// Specifies a JSON object that contains explicit custom fields to be defined at index time
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// Splunk source field
	Source *string `json:"source,omitempty"`
	// Splunk sourcetype field
	Sourcetype *string `json:"sourcetype,omitempty"`
	// Splunk host field
	Host *string `json:"host,omitempty"`
}
```


#### type InputData

```go
type InputData struct {
	Kind   InputDataKind   `json:"kind"`
	Source InputDataSource `json:"source"`
}
```


#### type InputDataKind

```go
type InputDataKind string
```


```go
const (
	S3InputKind      InputDataKind = "S3"
	SPLInputKind     InputDataKind = "SPL"
	RawDataInputKind InputDataKind = "RawData"
)
```
List of InputDataKind

#### type InputDataSource

```go
type InputDataSource struct {
	Key   *string `json:"key,omitempty"`
	Query string  `json:"query"`
	// Determine whether the Search service extracts all available fields in the data, including fields not mentioned in the SPL for the search job. Set to 'false' for better search performance.
	ExtractAllFields *bool `json:"extractAllFields,omitempty"`
	// The number of seconds to run this search before finalizing
	MaxTime *int32 `json:"maxTime,omitempty"`
	// Represents parameters on the search job such as 'earliest' and 'latest'.
	QueryParameters *map[string]interface{} `json:"queryParameters,omitempty"`
	Data            *string                 `json:"data,omitempty"`
}
```


#### type OutputData

```go
type OutputData struct {
	Kind        *OutputDataKind        `json:"kind,omitempty"`
	Destination *OutputDataDestination `json:"destination,omitempty"`
}
```


#### type OutputDataDestination

```go
type OutputDataDestination struct {
	Key *string `json:"key,omitempty"`
	// Specifies a JSON object that contains explicit custom fields to be defined at index time
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// Splunk source field
	Source *string `json:"source,omitempty"`
	// Splunk sourcetype field
	Sourcetype *string `json:"sourcetype,omitempty"`
	// Splunk host field
	Host *string `json:"host,omitempty"`
}
```


#### type OutputDataKind

```go
type OutputDataKind string
```


```go
const (
	S3OutputKind  OutputDataKind = "S3"
	HecOutputKind OutputDataKind = "HEC"
)
```
List of OutputDataKind

#### type RawData

```go
type RawData struct {
	Data *string `json:"data,omitempty"`
}
```


#### type S3Key

```go
type S3Key struct {
	Key *string `json:"key,omitempty"`
}
```


#### type Service

```go
type Service services.BaseService
```

Service talks to the Splunk Cloud machine learning service

#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new machine learning service client from the given Config

#### func (*Service) CreateWorkflow

```go
func (s *Service) CreateWorkflow(workflow Workflow) (*Workflow, error)
```
CreateWorkflow Create a workflow configuration

#### func (*Service) CreateWorkflowBuild

```go
func (s *Service) CreateWorkflowBuild(id string, workflowBuild WorkflowBuild) (*WorkflowBuild, error)
```
CreateWorkflowBuild Create a workflow build

#### func (*Service) CreateWorkflowRun

```go
func (s *Service) CreateWorkflowRun(id string, buildID string, workflowRun WorkflowRun) (*WorkflowRun, error)
```
CreateWorkflowRun Create a workflow Run

#### func (*Service) DeleteWorkflow

```go
func (s *Service) DeleteWorkflow(id string) error
```
DeleteWorkflow Delete a workflow configuration

#### func (*Service) DeleteWorkflowBuild

```go
func (s *Service) DeleteWorkflowBuild(id string, buildID string) error
```
DeleteWorkflowBuild Delete workflow build

#### func (*Service) DeleteWorkflowRun

```go
func (s *Service) DeleteWorkflowRun(id string, buildID string, runID string) error
```
DeleteWorkflowRun Delete a workflow run

#### func (*Service) GetWorkflow

```go
func (s *Service) GetWorkflow(id string) (*Workflow, error)
```
GetWorkflow Get a workflow configuration

#### func (*Service) GetWorkflowBuild

```go
func (s *Service) GetWorkflowBuild(id string, buildID string) (*WorkflowBuild, error)
```
GetWorkflowBuild Get status of a workflow build

#### func (*Service) GetWorkflowRun

```go
func (s *Service) GetWorkflowRun(id string, buildID string, runID string) (*WorkflowRun, error)
```
GetWorkflowRun Get status of a workflow run

#### func (*Service) ListWorkflowBuilds

```go
func (s *Service) ListWorkflowBuilds(id string) ([]WorkflowBuild, error)
```
ListWorkflowBuilds Get list of workflow builds

#### func (*Service) ListWorkflowRuns

```go
func (s *Service) ListWorkflowRuns(id string, buildID string) ([]WorkflowRun, error)
```
ListWorkflowRuns Get list of workflow runs

#### func (*Service) ListWorkflows

```go
func (s *Service) ListWorkflows() ([]WorkflowsGetResponse, error)
```
ListWorkflows Get the list of workflow configurations

#### type Servicer

```go
type Servicer interface {
	//CreateWorkflow Create a workflow configuration
	CreateWorkflow(workflow Workflow) (*Workflow, error)
	//CreateWorkflowBuild Create a workflow build
	CreateWorkflowBuild(id string, workflowBuild WorkflowBuild) (*WorkflowBuild, error)
	//CreateWorkflowRun Create a workflow Run
	CreateWorkflowRun(id string, buildID string, workflowRun WorkflowRun) (*WorkflowRun, error)
	//DeleteWorkflow Delete a workflow configuration
	DeleteWorkflow(id string) error
	//DeleteWorkflowBuild Delete workflow build
	DeleteWorkflowBuild(id string, buildID string) error
	//DeleteWorkflowRun Delete a workflow run
	DeleteWorkflowRun(id string, buildID string, runID string) error
	//GetWorkflow Get a workflow configuration
	GetWorkflow(id string) (*Workflow, error)
	//GetWorkflowBuild Get status of a workflow build
	GetWorkflowBuild(id string, buildID string) (*WorkflowBuild, error)
	//ListWorkflowBuilds Get list of workflow builds
	ListWorkflowBuilds(id string) ([]WorkflowBuild, error)
	//GetWorkflowRun Get status of a workflow run
	GetWorkflowRun(id string, buildID string, runID string) (*WorkflowRun, error)
	//ListWorkflowRuns Get list of workflow runs
	ListWorkflowRuns(id string, buildID string) ([]WorkflowRun, error)
	//ListWorkflows Get the list of workflow configurations
	ListWorkflows() ([]WorkflowsGetResponse, error)
}
```

Servicer ...

#### type Spl

```go
type Spl struct {
	Query string `json:"query"`
	// Determine whether the Search service extracts all available fields in the data, including fields not mentioned in the SPL for the search job. Set to 'false' for better search performance.
	ExtractAllFields *bool `json:"extractAllFields,omitempty"`
	// The number of seconds to run this search before finalizing
	MaxTime *int32 `json:"maxTime,omitempty"`
	// Represents parameters on the search job such as 'earliest' and 'latest'.
	QueryParameters *map[string]interface{} `json:"queryParameters,omitempty"`
}
```


#### type Task

```go
type Task struct {
	// The name has to be unique in the same workflow, it is optional, can be used to identify that task artifact
	Name              *string   `json:"name,omitempty"`
	Kind              *TaskKind `json:"kind,omitempty"`
	Algorithm         string    `json:"algorithm"`
	Fields            Fields    `json:"fields"`
	OutputTransformer *string   `json:"outputTransformer,omitempty"`
	// time out in seconds
	TimeoutSecs int32                   `json:"timeoutSecs"`
	Parameters  *map[string]interface{} `json:"parameters,omitempty"`
}
```


#### type TaskCommon

```go
type TaskCommon struct {
	// The name has to be unique in the same workflow, it is optional, can be used to identify that task artifact
	Name *string   `json:"name,omitempty"`
	Kind *TaskKind `json:"kind,omitempty"`
}
```


#### type TaskKind

```go
type TaskKind string
```


```go
const (
	FitTaskKind TaskKind = "fit"
)
```
List of TaskKind

#### type Workflow

```go
type Workflow struct {
	ID           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Tasks        []Task  `json:"tasks"`
	CreationTime *string `json:"creationTime,omitempty"`
}
```


#### type WorkflowBuild

```go
type WorkflowBuild struct {
	ID           *string              `json:"id,omitempty"`
	Name         *string              `json:"name,omitempty"`
	Status       *WorkflowBuildStatus `json:"status,omitempty"`
	Input        InputData            `json:"input"`
	Output       *OutputData          `json:"output,omitempty"`
	Workflow     *Workflow            `json:"workflow,omitempty"`
	CreationTime *string              `json:"creationTime,omitempty"`
	StartTime    *string              `json:"startTime,omitempty"`
	EndTime      *string              `json:"endTime,omitempty"`
}
```


#### type WorkflowBuildStatus

```go
type WorkflowBuildStatus string
```


```go
const (
	RunningWorkflowBuildStatus   WorkflowBuildStatus = "running"
	FailedWorkflowBuildStatus    WorkflowBuildStatus = "failed"
	SuccessWorkflowBuildStatus   WorkflowBuildStatus = "success"
	ScheduledWorkflowBuildStatus WorkflowBuildStatus = "scheduled"
)
```
List of WorkflowBuildStatus

#### type WorkflowRun

```go
type WorkflowRun struct {
	ID            *string            `json:"id,omitempty"`
	Name          *string            `json:"name,omitempty"`
	Status        *WorkflowRunStatus `json:"status,omitempty"`
	WorkflowBuild *WorkflowBuild     `json:"workflowBuild,omitempty"`
	Input         InputData          `json:"input"`
	Output        OutputData         `json:"output"`
	CreationTime  *string            `json:"creationTime,omitempty"`
	StartTime     *string            `json:"startTime,omitempty"`
	EndTime       *string            `json:"endTime,omitempty"`
}
```


#### type WorkflowRunStatus

```go
type WorkflowRunStatus string
```


```go
const (
	RunningWorkflowRunStatus   WorkflowRunStatus = "running"
	FailedWorkflowRunStatus    WorkflowRunStatus = "failed"
	SuccessWorkflowRunStatus   WorkflowRunStatus = "success"
	ScheduledWorkflowRunStatus WorkflowRunStatus = "scheduled"
)
```
List of WorkflowRunStatus

#### type WorkflowsGetResponse

```go
type WorkflowsGetResponse struct {
	ID           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	CreationTime *string `json:"creationTime,omitempty"`
}
```
