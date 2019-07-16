# ml
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/ml"


## Usage

#### type ClassificationReport

```go
type ClassificationReport map[string]interface{}
```


#### type ClusteringReport

```go
type ClusteringReport map[string]interface{}
```


#### type CrossValidation

```go
type CrossValidation struct {
	// Number of folds in the K-fold cross validation.
	KFold *int32 `json:"kFold,omitempty"`
	// Random state to shuffle and partition data.
	RandomSeed *int32 `json:"randomSeed,omitempty"`
	// Determine whether stratification is used in partitioning the data.
	Stratified *bool `json:"stratified,omitempty"`
}
```


#### type CrossValidationScore

```go
type CrossValidationScore []map[string]interface{}
```


#### type DeploymentSpec

```go
type DeploymentSpec struct {
	// CPU Resource limit for each container in a deployment.
	CpuLimit *string `json:"cpuLimit,omitempty"`
	// CPU Resource limit for serving requests.
	CpuRequest *string `json:"cpuRequest,omitempty"`
	// Memory Resource limit for each container in a deployment.
	MemoryLimit *string `json:"memoryLimit,omitempty"`
	// Memory Resource limit for serving requests.
	MemoryRequest *string `json:"memoryRequest,omitempty"`
	// Create replicated pods in a deployment.
	Replicas *int32 `json:"replicas,omitempty"`
}
```


#### type Error

```go
type Error struct {
	Code    string                   `json:"code"`
	Message string                   `json:"message"`
	Details []map[string]interface{} `json:"details,omitempty"`
}
```


#### type Events

```go
type Events struct {
	// Specifies a JSON object that contains explicit custom fields to be defined at index time.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// Splunk host field.
	Host *string `json:"host,omitempty"`
	// Splunk source field.
	Source *string `json:"source,omitempty"`
	// Splunk sourcetype field.
	Sourcetype *string `json:"sourcetype,omitempty"`
}
```

Output events to the Ingest /events endpoint.

#### type ExecutorErrors

```go
type ExecutorErrors struct {
	Message *string `json:"message,omitempty"`
}
```

Executor errors.

#### type ExecutorLogs

```go
type ExecutorLogs struct {
	Level   *string `json:"level,omitempty"`
	Message *string `json:"message,omitempty"`
}
```

Executor logs.

#### type Fields

```go
type Fields struct {
	// Fields necessary for task.
	Features []string `json:"features"`
	// Fields produced by task.
	Created []string `json:"created,omitempty"`
	// Target field necessary for task.
	Target *string `json:"target,omitempty"`
}
```


#### type FitTask

```go
type FitTask struct {
	Algorithm string       `json:"algorithm"`
	Fields    Fields       `json:"fields"`
	Kind      *FitTaskKind `json:"kind,omitempty"`
	// The name has to be unique in the same workflow, it is optional, can be used to identify that task artifact.
	Name              *string                `json:"name,omitempty"`
	OutputTransformer *string                `json:"outputTransformer,omitempty"`
	Parameters        map[string]interface{} `json:"parameters,omitempty"`
}
```

Fit task does an estimation/training/fitting, fit task outputs a transformer.

#### type FitTaskKind

```go
type FitTaskKind string
```


```go
const (
	FitTaskKindFit FitTaskKind = "fit"
)
```
List of FitTaskKind

#### type ForecastingReport

```go
type ForecastingReport map[string]interface{}
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
	InputDataKindSpl     InputDataKind = "SPL"
	InputDataKindRawData InputDataKind = "RawData"
)
```
List of InputDataKind

#### type InputDataSource

```go
type InputDataSource struct {
}
```

InputDataSource is RawData, Spl, (or interface{} if no matches are found)

#### func  MakeInputDataSourceFromRawData

```go
func MakeInputDataSourceFromRawData(f RawData) InputDataSource
```
MakeInputDataSourceFromRawData creates a new InputDataSource from an instance of
RawData

#### func  MakeInputDataSourceFromRawInterface

```go
func MakeInputDataSourceFromRawInterface(f interface{}) InputDataSource
```
MakeInputDataSourceFromRawInterface creates a new InputDataSource from a raw
interface{}

#### func  MakeInputDataSourceFromSpl

```go
func MakeInputDataSourceFromSpl(f Spl) InputDataSource
```
MakeInputDataSourceFromSpl creates a new InputDataSource from an instance of Spl

#### func (InputDataSource) IsRawData

```go
func (m InputDataSource) IsRawData() bool
```
IsRawData checks if the InputDataSource is a RawData

#### func (InputDataSource) IsRawInterface

```go
func (m InputDataSource) IsRawInterface() bool
```
IsRawInterface checks if the InputDataSource is an interface{} (unknown type)

#### func (InputDataSource) IsSpl

```go
func (m InputDataSource) IsSpl() bool
```
IsSpl checks if the InputDataSource is a Spl

#### func (InputDataSource) MarshalJSON

```go
func (m InputDataSource) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals InputDataSource using InputDataSource.InputDataSource

#### func (InputDataSource) RawData

```go
func (m InputDataSource) RawData() *RawData
```
RawData returns RawData if IsRawData() is true, nil otherwise

#### func (InputDataSource) RawInterface

```go
func (m InputDataSource) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (InputDataSource) Spl

```go
func (m InputDataSource) Spl() *Spl
```
Spl returns Spl if IsSpl() is true, nil otherwise

#### func (*InputDataSource) UnmarshalJSON

```go
func (m *InputDataSource) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals InputDataSource into RawData, Spl, or interface{} if no
matches are found

#### type OutputData

```go
type OutputData struct {
	Destination *OutputDataDestination `json:"destination,omitempty"`
	Kind        *OutputDataKind        `json:"kind,omitempty"`
}
```


#### type OutputDataDestination

```go
type OutputDataDestination struct {
}
```

OutputDataDestination is Events, (or interface{} if no matches are found)

#### func  MakeOutputDataDestinationFromEvents

```go
func MakeOutputDataDestinationFromEvents(f Events) OutputDataDestination
```
MakeOutputDataDestinationFromEvents creates a new OutputDataDestination from an
instance of Events

#### func  MakeOutputDataDestinationFromRawInterface

```go
func MakeOutputDataDestinationFromRawInterface(f interface{}) OutputDataDestination
```
MakeOutputDataDestinationFromRawInterface creates a new OutputDataDestination
from a raw interface{}

#### func (OutputDataDestination) Events

```go
func (m OutputDataDestination) Events() *Events
```
Events returns Events if IsEvents() is true, nil otherwise

#### func (OutputDataDestination) IsEvents

```go
func (m OutputDataDestination) IsEvents() bool
```
IsEvents checks if the OutputDataDestination is a Events

#### func (OutputDataDestination) IsRawInterface

```go
func (m OutputDataDestination) IsRawInterface() bool
```
IsRawInterface checks if the OutputDataDestination is an interface{} (unknown
type)

#### func (OutputDataDestination) MarshalJSON

```go
func (m OutputDataDestination) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals OutputDataDestination using
OutputDataDestination.OutputDataDestination

#### func (OutputDataDestination) RawInterface

```go
func (m OutputDataDestination) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (*OutputDataDestination) UnmarshalJSON

```go
func (m *OutputDataDestination) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals OutputDataDestination into Events, or interface{} if no
matches are found

#### type OutputDataKind

```go
type OutputDataKind string
```


```go
const (
	OutputDataKindEvents OutputDataKind = "Events"
)
```
List of OutputDataKind

#### type RawData

```go
type RawData struct {
	// A base-64 encoded CSV string.
	Data *string `json:"data,omitempty"`
}
```

Send data directly via the reqest body as a base-64 encoded CSV string.

#### type RegressionReport

```go
type RegressionReport map[string]interface{}
```


#### type Score

```go
type Score struct {
	Kind   ScoreKind   `json:"kind"`
	Report ScoreReport `json:"report"`
}
```


#### type ScoreKind

```go
type ScoreKind string
```


```go
const (
	ScoreKindRegression     ScoreKind = "regression"
	ScoreKindClassification ScoreKind = "classification"
	ScoreKindClustering     ScoreKind = "clustering"
	ScoreKindForecasting    ScoreKind = "forecasting"
)
```
List of ScoreKind

#### type ScoreReport

```go
type ScoreReport struct {
}
```

ScoreReport is ClassificationReport, ClusteringReport, ForecastingReport,
RegressionReport, (or interface{} if no matches are found)

#### func  MakeScoreReportFromClassificationReport

```go
func MakeScoreReportFromClassificationReport(f ClassificationReport) ScoreReport
```
MakeScoreReportFromClassificationReport creates a new ScoreReport from an
instance of ClassificationReport

#### func  MakeScoreReportFromClusteringReport

```go
func MakeScoreReportFromClusteringReport(f ClusteringReport) ScoreReport
```
MakeScoreReportFromClusteringReport creates a new ScoreReport from an instance
of ClusteringReport

#### func  MakeScoreReportFromForecastingReport

```go
func MakeScoreReportFromForecastingReport(f ForecastingReport) ScoreReport
```
MakeScoreReportFromForecastingReport creates a new ScoreReport from an instance
of ForecastingReport

#### func  MakeScoreReportFromRawInterface

```go
func MakeScoreReportFromRawInterface(f interface{}) ScoreReport
```
MakeScoreReportFromRawInterface creates a new ScoreReport from a raw interface{}

#### func  MakeScoreReportFromRegressionReport

```go
func MakeScoreReportFromRegressionReport(f RegressionReport) ScoreReport
```
MakeScoreReportFromRegressionReport creates a new ScoreReport from an instance
of RegressionReport

#### func (ScoreReport) ClassificationReport

```go
func (m ScoreReport) ClassificationReport() *ClassificationReport
```
ClassificationReport returns ClassificationReport if IsClassificationReport() is
true, nil otherwise

#### func (ScoreReport) ClusteringReport

```go
func (m ScoreReport) ClusteringReport() *ClusteringReport
```
ClusteringReport returns ClusteringReport if IsClusteringReport() is true, nil
otherwise

#### func (ScoreReport) ForecastingReport

```go
func (m ScoreReport) ForecastingReport() *ForecastingReport
```
ForecastingReport returns ForecastingReport if IsForecastingReport() is true,
nil otherwise

#### func (ScoreReport) IsClassificationReport

```go
func (m ScoreReport) IsClassificationReport() bool
```
IsClassificationReport checks if the ScoreReport is a ClassificationReport

#### func (ScoreReport) IsClusteringReport

```go
func (m ScoreReport) IsClusteringReport() bool
```
IsClusteringReport checks if the ScoreReport is a ClusteringReport

#### func (ScoreReport) IsForecastingReport

```go
func (m ScoreReport) IsForecastingReport() bool
```
IsForecastingReport checks if the ScoreReport is a ForecastingReport

#### func (ScoreReport) IsRawInterface

```go
func (m ScoreReport) IsRawInterface() bool
```
IsRawInterface checks if the ScoreReport is an interface{} (unknown type)

#### func (ScoreReport) IsRegressionReport

```go
func (m ScoreReport) IsRegressionReport() bool
```
IsRegressionReport checks if the ScoreReport is a RegressionReport

#### func (ScoreReport) MarshalJSON

```go
func (m ScoreReport) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals ScoreReport using ScoreReport.ScoreReport

#### func (ScoreReport) RawInterface

```go
func (m ScoreReport) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (ScoreReport) RegressionReport

```go
func (m ScoreReport) RegressionReport() *RegressionReport
```
RegressionReport returns RegressionReport if IsRegressionReport() is true, nil
otherwise

#### func (*ScoreReport) UnmarshalJSON

```go
func (m *ScoreReport) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals ScoreReport into ClassificationReport,
ClusteringReport, ForecastingReport, RegressionReport, or interface{} if no
matches are found

#### type Service

```go
type Service services.BaseService
```


#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new ml service client from the given Config

#### func (*Service) CreateWorkflow

```go
func (s *Service) CreateWorkflow(workflow Workflow, resp ...*http.Response) (*Workflow, error)
```
CreateWorkflow - Creates a workflow configuration. Parameters:

    workflow: Workflow configuration to be created.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateWorkflowBuild

```go
func (s *Service) CreateWorkflowBuild(id string, workflowBuild WorkflowBuild, resp ...*http.Response) (*WorkflowBuild, error)
```
CreateWorkflowBuild - Creates a workflow build. Parameters:

    id: The workflow ID.
    workflowBuild: Input data used to build the workflow.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateWorkflowDeployment

```go
func (s *Service) CreateWorkflowDeployment(id string, buildId string, workflowDeployment WorkflowDeployment, resp ...*http.Response) (*WorkflowDeployment, error)
```
CreateWorkflowDeployment - Creates a workflow deployment. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    workflowDeployment: Input data used to build the workflow deployment.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateWorkflowInference

```go
func (s *Service) CreateWorkflowInference(id string, buildId string, deploymentId string, workflowInference WorkflowInference, resp ...*http.Response) (*WorkflowInference, error)
```
CreateWorkflowInference - Creates a workflow inference request. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    deploymentId: The workflow deployment ID.
    workflowInference: Input data to the inference request.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateWorkflowRun

```go
func (s *Service) CreateWorkflowRun(id string, buildId string, workflowRun WorkflowRun, resp ...*http.Response) (*WorkflowRun, error)
```
CreateWorkflowRun - Creates a workflow run. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    workflowRun: Input data used to build the workflow.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteWorkflow

```go
func (s *Service) DeleteWorkflow(id string, resp ...*http.Response) error
```
DeleteWorkflow - Removes a workflow configuration. Parameters:

    id: The workflow ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteWorkflowBuild

```go
func (s *Service) DeleteWorkflowBuild(id string, buildId string, resp ...*http.Response) error
```
DeleteWorkflowBuild - Removes a workflow build. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteWorkflowDeployment

```go
func (s *Service) DeleteWorkflowDeployment(id string, buildId string, deploymentId string, resp ...*http.Response) error
```
DeleteWorkflowDeployment - Removes a workflow deployment. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    deploymentId: The workflow deployment ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteWorkflowRun

```go
func (s *Service) DeleteWorkflowRun(id string, buildId string, runId string, resp ...*http.Response) error
```
DeleteWorkflowRun - Removes a workflow run. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    runId: The workflow run ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflow

```go
func (s *Service) GetWorkflow(id string, resp ...*http.Response) (*Workflow, error)
```
GetWorkflow - Returns a workflow configuration. Parameters:

    id: The workflow ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowBuild

```go
func (s *Service) GetWorkflowBuild(id string, buildId string, resp ...*http.Response) (*WorkflowBuild, error)
```
GetWorkflowBuild - Returns the status of a workflow build. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowBuildError

```go
func (s *Service) GetWorkflowBuildError(id string, buildId string, resp ...*http.Response) (*WorkflowBuildError, error)
```
GetWorkflowBuildError - Returns a list of workflow errors. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowBuildLog

```go
func (s *Service) GetWorkflowBuildLog(id string, buildId string, resp ...*http.Response) (*WorkflowBuildLog, error)
```
GetWorkflowBuildLog - Returns the logs from a workflow build. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowDeployment

```go
func (s *Service) GetWorkflowDeployment(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeployment, error)
```
GetWorkflowDeployment - Returns the status of a workflow deployment. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    deploymentId: The workflow deployment ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowDeploymentError

```go
func (s *Service) GetWorkflowDeploymentError(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeploymentError, error)
```
GetWorkflowDeploymentError - Returns a list of workflow deployment errors.
Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    deploymentId: The workflow deployment ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowDeploymentLog

```go
func (s *Service) GetWorkflowDeploymentLog(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeploymentLog, error)
```
GetWorkflowDeploymentLog - Returns the logs from a workflow deployment.
Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    deploymentId: The workflow deployment ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowRun

```go
func (s *Service) GetWorkflowRun(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRun, error)
```
GetWorkflowRun - Returns the status of a workflow run. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    runId: The workflow run ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowRunError

```go
func (s *Service) GetWorkflowRunError(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRunError, error)
```
GetWorkflowRunError - Returns the errors for a workflow run. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    runId: The workflow run ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowRunLog

```go
func (s *Service) GetWorkflowRunLog(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRunLog, error)
```
GetWorkflowRunLog - Returns the logs for a workflow run. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    runId: The workflow run ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListWorkflowBuilds

```go
func (s *Service) ListWorkflowBuilds(id string, resp ...*http.Response) ([]WorkflowBuild, error)
```
ListWorkflowBuilds - Returns a list of workflow builds. Parameters:

    id: The workflow ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListWorkflowDeployments

```go
func (s *Service) ListWorkflowDeployments(id string, buildId string, resp ...*http.Response) ([]WorkflowDeployment, error)
```
ListWorkflowDeployments - Returns a list of workflow deployments. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListWorkflowRuns

```go
func (s *Service) ListWorkflowRuns(id string, buildId string, resp ...*http.Response) ([]WorkflowRun, error)
```
ListWorkflowRuns - Returns a list of workflow runs. Parameters:

    id: The workflow ID.
    buildId: The workflow build ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListWorkflows

```go
func (s *Service) ListWorkflows(resp ...*http.Response) ([]WorkflowsGetResponse, error)
```
ListWorkflows - Returns a list of workflow configurations. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### type Servicer

```go
type Servicer interface {
	/*
		CreateWorkflow - Creates a workflow configuration.
		Parameters:
			workflow: Workflow configuration to be created.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflow(workflow Workflow, resp ...*http.Response) (*Workflow, error)
	/*
		CreateWorkflowBuild - Creates a workflow build.
		Parameters:
			id: The workflow ID.
			workflowBuild: Input data used to build the workflow.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflowBuild(id string, workflowBuild WorkflowBuild, resp ...*http.Response) (*WorkflowBuild, error)
	/*
		CreateWorkflowDeployment - Creates a workflow deployment.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			workflowDeployment: Input data used to build the workflow deployment.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflowDeployment(id string, buildId string, workflowDeployment WorkflowDeployment, resp ...*http.Response) (*WorkflowDeployment, error)
	/*
		CreateWorkflowInference - Creates a workflow inference request.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			deploymentId: The workflow deployment ID.
			workflowInference: Input data to the inference request.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflowInference(id string, buildId string, deploymentId string, workflowInference WorkflowInference, resp ...*http.Response) (*WorkflowInference, error)
	/*
		CreateWorkflowRun - Creates a workflow run.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			workflowRun: Input data used to build the workflow.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflowRun(id string, buildId string, workflowRun WorkflowRun, resp ...*http.Response) (*WorkflowRun, error)
	/*
		DeleteWorkflow - Removes a workflow configuration.
		Parameters:
			id: The workflow ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflow(id string, resp ...*http.Response) error
	/*
		DeleteWorkflowBuild - Removes a workflow build.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflowBuild(id string, buildId string, resp ...*http.Response) error
	/*
		DeleteWorkflowDeployment - Removes a workflow deployment.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			deploymentId: The workflow deployment ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflowDeployment(id string, buildId string, deploymentId string, resp ...*http.Response) error
	/*
		DeleteWorkflowRun - Removes a workflow run.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			runId: The workflow run ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflowRun(id string, buildId string, runId string, resp ...*http.Response) error
	/*
		GetWorkflow - Returns a workflow configuration.
		Parameters:
			id: The workflow ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflow(id string, resp ...*http.Response) (*Workflow, error)
	/*
		GetWorkflowBuild - Returns the status of a workflow build.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowBuild(id string, buildId string, resp ...*http.Response) (*WorkflowBuild, error)
	/*
		GetWorkflowBuildError - Returns a list of workflow errors.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowBuildError(id string, buildId string, resp ...*http.Response) (*WorkflowBuildError, error)
	/*
		GetWorkflowBuildLog - Returns the logs from a workflow build.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowBuildLog(id string, buildId string, resp ...*http.Response) (*WorkflowBuildLog, error)
	/*
		GetWorkflowDeployment - Returns the status of a workflow deployment.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			deploymentId: The workflow deployment ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowDeployment(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeployment, error)
	/*
		GetWorkflowDeploymentError - Returns a list of workflow deployment errors.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			deploymentId: The workflow deployment ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowDeploymentError(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeploymentError, error)
	/*
		GetWorkflowDeploymentLog - Returns the logs from a workflow deployment.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			deploymentId: The workflow deployment ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowDeploymentLog(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeploymentLog, error)
	/*
		GetWorkflowRun - Returns the status of a workflow run.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			runId: The workflow run ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowRun(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRun, error)
	/*
		GetWorkflowRunError - Returns the errors for a workflow run.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			runId: The workflow run ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowRunError(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRunError, error)
	/*
		GetWorkflowRunLog - Returns the logs for a workflow run.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			runId: The workflow run ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowRunLog(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRunLog, error)
	/*
		ListWorkflowBuilds - Returns a list of workflow builds.
		Parameters:
			id: The workflow ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListWorkflowBuilds(id string, resp ...*http.Response) ([]WorkflowBuild, error)
	/*
		ListWorkflowDeployments - Returns a list of workflow deployments.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListWorkflowDeployments(id string, buildId string, resp ...*http.Response) ([]WorkflowDeployment, error)
	/*
		ListWorkflowRuns - Returns a list of workflow runs.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListWorkflowRuns(id string, buildId string, resp ...*http.Response) ([]WorkflowRun, error)
	/*
		ListWorkflows - Returns a list of workflow configurations.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListWorkflows(resp ...*http.Response) ([]WorkflowsGetResponse, error)
}
```

Servicer represents the interface for implementing all endpoints for this
service

#### type Spl

```go
type Spl struct {
	Query string `json:"query"`
	// Determine whether the Search service extracts all available fields in the data, including fields not mentioned in the SPL for the search job. Set to 'false' for better search performance.
	ExtractAllFields *bool `json:"extractAllFields,omitempty"`
	// The number of seconds to run this search before finalizing.
	MaxTime *int32 `json:"maxTime,omitempty"`
	// The module to run the search in. The default module is used if a module is not specified.
	Module *string `json:"module,omitempty"`
	// Represents parameters on the search job such as 'earliest' and 'latest'.
	QueryParameters map[string]interface{} `json:"queryParameters,omitempty"`
}
```


#### type Task

```go
type Task struct {
}
```


#### func  MakeTaskFromFitTask

```go
func MakeTaskFromFitTask(f FitTask) Task
```
MakeTaskFromFitTask creates a new Task from an instance of FitTask

#### func  MakeTaskFromRawInterface

```go
func MakeTaskFromRawInterface(f interface{}) Task
```
MakeTaskFromRawInterface creates a new Task from a raw interface{}

#### func (Task) FitTask

```go
func (m Task) FitTask() *FitTask
```
FitTask returns FitTask if IsFitTask() is true, nil otherwise

#### func (Task) IsFitTask

```go
func (m Task) IsFitTask() bool
```
IsFitTask checks if the Task is a FitTask

#### func (Task) IsRawInterface

```go
func (m Task) IsRawInterface() bool
```
IsRawInterface checks if the Task is an interface{} (unknown type)

#### func (Task) MarshalJSON

```go
func (m Task) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals Task using the appropriate struct field

#### func (Task) RawInterface

```go
func (m Task) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (*Task) UnmarshalJSON

```go
func (m *Task) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals Task using the "kind" property

#### type TaskCommon

```go
type TaskCommon struct {
	Kind *TaskCommonKind `json:"kind,omitempty"`
	// The name has to be unique in the same workflow, it is optional, can be used to identify that task artifact.
	Name *string `json:"name,omitempty"`
}
```


#### type TaskCommonKind

```go
type TaskCommonKind string
```


```go
const (
	TaskCommonKindFit TaskCommonKind = "fit"
)
```
List of TaskCommonKind

#### type TaskKind

```go
type TaskKind string
```


```go
const (
	TaskKindFit TaskKind = "fit"
)
```
List of TaskKind

#### type TaskSummary

```go
type TaskSummary struct {
	Algorithm *string `json:"algorithm,omitempty"`
	Name      *string `json:"name,omitempty"`
	// Summary of the task, including but not limited to learned parameters and statistics.
	Summary interface{} `json:"summary,omitempty"`
}
```


#### type TrainTestScore

```go
type TrainTestScore struct {
	TestScore     *Score `json:"testScore,omitempty"`
	TrainingScore *Score `json:"trainingScore,omitempty"`
}
```


#### type TrainTestSplit

```go
type TrainTestSplit struct {
	// Random state to shuffle and partition data.
	RandomSeed *int32 `json:"randomSeed,omitempty"`
	// Ratio to split out training set and test set. For example, 0.8 means 80% of data for the training set and 20% for the test set.
	Ratio *float32 `json:"ratio,omitempty"`
	// Determine whether stratification is used in partitioning the data.
	Stratified *bool `json:"stratified,omitempty"`
}
```


#### type TrainingParameters

```go
type TrainingParameters map[string]interface{}
```


#### type Workflow

```go
type Workflow struct {
	Tasks        []Task  `json:"tasks"`
	CreationTime *string `json:"creationTime,omitempty"`
	Id           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
}
```


#### type WorkflowBuild

```go
type WorkflowBuild struct {
	Input           InputData            `json:"input"`
	CreationTime    *string              `json:"creationTime,omitempty"`
	EndTime         *string              `json:"endTime,omitempty"`
	Id              *string              `json:"id,omitempty"`
	Name            *string              `json:"name,omitempty"`
	Output          *OutputData          `json:"output,omitempty"`
	PipelineSummary []TaskSummary        `json:"pipelineSummary,omitempty"`
	StartTime       *string              `json:"startTime,omitempty"`
	Status          *WorkflowBuildStatus `json:"status,omitempty"`
	// Number of seconds before a workflow build times out.
	TimeoutSecs      *int32                         `json:"timeoutSecs,omitempty"`
	TrainingScore    *Score                         `json:"trainingScore,omitempty"`
	ValidationOption *WorkflowBuildValidationOption `json:"validationOption,omitempty"`
	ValidationScore  *WorkflowBuildValidationScore  `json:"validationScore,omitempty"`
	Workflow         *Workflow                      `json:"workflow,omitempty"`
}
```


#### type WorkflowBuildError

```go
type WorkflowBuildError struct {
	Id                    string                  `json:"id"`
	WorkflowId            string                  `json:"workflowId"`
	ExecutorErrors        []ExecutorErrors        `json:"executorErrors,omitempty"`
	RequestId             *string                 `json:"requestId,omitempty"`
	WorkflowManagerErrors []WorkflowManagerErrors `json:"workflowManagerErrors,omitempty"`
}
```


#### type WorkflowBuildLog

```go
type WorkflowBuildLog struct {
	Id                  string                `json:"id"`
	WorkflowId          string                `json:"workflowId"`
	ExecutorLogs        []ExecutorLogs        `json:"executorLogs,omitempty"`
	RequestId           *string               `json:"requestId,omitempty"`
	WorkflowManagerLogs []WorkflowManagerLogs `json:"workflowManagerLogs,omitempty"`
}
```


#### type WorkflowBuildStatus

```go
type WorkflowBuildStatus string
```


```go
const (
	WorkflowBuildStatusRunning   WorkflowBuildStatus = "running"
	WorkflowBuildStatusFailed    WorkflowBuildStatus = "failed"
	WorkflowBuildStatusSuccess   WorkflowBuildStatus = "success"
	WorkflowBuildStatusScheduled WorkflowBuildStatus = "scheduled"
)
```
List of WorkflowBuildStatus

#### type WorkflowBuildValidationOption

```go
type WorkflowBuildValidationOption struct {
	Kind   WorkflowBuildValidationOptionKind `json:"kind"`
	Option *WorkflowValidationOption         `json:"option,omitempty"`
}
```

Represents which type of validation to use in the workflow along with any
parameters if specified. If this is not included, no validation is done (all
data is used for training). Default parameter values are used if no `option` is
specified.

#### type WorkflowBuildValidationOptionKind

```go
type WorkflowBuildValidationOptionKind string
```


```go
const (
	WorkflowBuildValidationOptionKindTrainTest       WorkflowBuildValidationOptionKind = "TrainTest"
	WorkflowBuildValidationOptionKindCrossValidation WorkflowBuildValidationOptionKind = "CrossValidation"
)
```
List of WorkflowBuildValidationOptionKind

#### type WorkflowBuildValidationScore

```go
type WorkflowBuildValidationScore struct {
	Kind  WorkflowBuildValidationScoreKind `json:"kind"`
	Score WorkflowValidationScore          `json:"score"`
}
```

The validation score whose type is specified by the user in `validationOption`.

#### type WorkflowBuildValidationScoreKind

```go
type WorkflowBuildValidationScoreKind string
```


```go
const (
	WorkflowBuildValidationScoreKindTrainTest       WorkflowBuildValidationScoreKind = "TrainTest"
	WorkflowBuildValidationScoreKindCrossValidation WorkflowBuildValidationScoreKind = "CrossValidation"
)
```
List of WorkflowBuildValidationScoreKind

#### type WorkflowDeployment

```go
type WorkflowDeployment struct {
	Spec          DeploymentSpec            `json:"spec"`
	CreationTime  *string                   `json:"creationTime,omitempty"`
	EndTime       *string                   `json:"endTime,omitempty"`
	Id            *string                   `json:"id,omitempty"`
	Name          *string                   `json:"name,omitempty"`
	StartTime     *string                   `json:"startTime,omitempty"`
	Status        *WorkflowDeploymentStatus `json:"status,omitempty"`
	WorkflowBuild *WorkflowBuild            `json:"workflowBuild,omitempty"`
}
```


#### type WorkflowDeploymentError

```go
type WorkflowDeploymentError struct {
	BuildId               string                  `json:"buildId"`
	Id                    string                  `json:"id"`
	WorkflowId            string                  `json:"workflowId"`
	ExecutorErrors        []ExecutorErrors        `json:"executorErrors,omitempty"`
	RequestId             *string                 `json:"requestId,omitempty"`
	WorkflowManagerErrors []WorkflowManagerErrors `json:"workflowManagerErrors,omitempty"`
}
```


#### type WorkflowDeploymentLog

```go
type WorkflowDeploymentLog struct {
	BuildId             string                `json:"buildId"`
	Id                  string                `json:"id"`
	WorkflowId          string                `json:"workflowId"`
	ExecutorLogs        []ExecutorLogs        `json:"executorLogs,omitempty"`
	RequestId           *string               `json:"requestId,omitempty"`
	WorkflowManagerLogs []WorkflowManagerLogs `json:"workflowManagerLogs,omitempty"`
}
```


#### type WorkflowDeploymentStatus

```go
type WorkflowDeploymentStatus string
```


```go
const (
	WorkflowDeploymentStatusRunning   WorkflowDeploymentStatus = "running"
	WorkflowDeploymentStatusFailed    WorkflowDeploymentStatus = "failed"
	WorkflowDeploymentStatusSuccess   WorkflowDeploymentStatus = "success"
	WorkflowDeploymentStatusScheduled WorkflowDeploymentStatus = "scheduled"
)
```
List of WorkflowDeploymentStatus

#### type WorkflowInference

```go
type WorkflowInference struct {
	Input  string  `json:"input"`
	Output *string `json:"output,omitempty"`
}
```


#### type WorkflowManagerErrors

```go
type WorkflowManagerErrors struct {
	Message *string `json:"message,omitempty"`
}
```

Workflow manager errors.

#### type WorkflowManagerLogs

```go
type WorkflowManagerLogs struct {
	Level   *string `json:"level,omitempty"`
	Message *string `json:"message,omitempty"`
}
```

Workflow manager logs.

#### type WorkflowRun

```go
type WorkflowRun struct {
	Input        InputData  `json:"input"`
	Output       OutputData `json:"output"`
	CreationTime *string    `json:"creationTime,omitempty"`
	EndTime      *string    `json:"endTime,omitempty"`
	// Determine whether to evaluate the prediction.
	Evaluate        *bool              `json:"evaluate,omitempty"`
	Id              *string            `json:"id,omitempty"`
	Name            *string            `json:"name,omitempty"`
	PredictionScore *Score             `json:"predictionScore,omitempty"`
	StartTime       *string            `json:"startTime,omitempty"`
	Status          *WorkflowRunStatus `json:"status,omitempty"`
	// Number of seconds before a workflow run times out.
	TimeoutSecs   *int32         `json:"timeoutSecs,omitempty"`
	WorkflowBuild *WorkflowBuild `json:"workflowBuild,omitempty"`
}
```


#### type WorkflowRunError

```go
type WorkflowRunError struct {
	BuildId               string                  `json:"buildId"`
	Id                    string                  `json:"id"`
	WorkflowId            string                  `json:"workflowId"`
	ExecutorErrors        []ExecutorErrors        `json:"executorErrors,omitempty"`
	RequestId             *string                 `json:"requestId,omitempty"`
	WorkflowManagerErrors []WorkflowManagerErrors `json:"workflowManagerErrors,omitempty"`
}
```


#### type WorkflowRunLog

```go
type WorkflowRunLog struct {
	BuildId             string                `json:"buildId"`
	Id                  string                `json:"id"`
	WorkflowId          string                `json:"workflowId"`
	ExecutorLogs        []ExecutorLogs        `json:"executorLogs,omitempty"`
	RequestId           *string               `json:"requestId,omitempty"`
	WorkflowManagerLogs []WorkflowManagerLogs `json:"workflowManagerLogs,omitempty"`
}
```


#### type WorkflowRunStatus

```go
type WorkflowRunStatus string
```


```go
const (
	WorkflowRunStatusRunning   WorkflowRunStatus = "running"
	WorkflowRunStatusFailed    WorkflowRunStatus = "failed"
	WorkflowRunStatusSuccess   WorkflowRunStatus = "success"
	WorkflowRunStatusScheduled WorkflowRunStatus = "scheduled"
)
```
List of WorkflowRunStatus

#### type WorkflowValidationOption

```go
type WorkflowValidationOption struct {
}
```

WorkflowValidationOption is CrossValidation, TrainTestSplit, (or interface{} if
no matches are found)

#### func  MakeWorkflowValidationOptionFromCrossValidation

```go
func MakeWorkflowValidationOptionFromCrossValidation(f CrossValidation) WorkflowValidationOption
```
MakeWorkflowValidationOptionFromCrossValidation creates a new
WorkflowValidationOption from an instance of CrossValidation

#### func  MakeWorkflowValidationOptionFromRawInterface

```go
func MakeWorkflowValidationOptionFromRawInterface(f interface{}) WorkflowValidationOption
```
MakeWorkflowValidationOptionFromRawInterface creates a new
WorkflowValidationOption from a raw interface{}

#### func  MakeWorkflowValidationOptionFromTrainTestSplit

```go
func MakeWorkflowValidationOptionFromTrainTestSplit(f TrainTestSplit) WorkflowValidationOption
```
MakeWorkflowValidationOptionFromTrainTestSplit creates a new
WorkflowValidationOption from an instance of TrainTestSplit

#### func (WorkflowValidationOption) CrossValidation

```go
func (m WorkflowValidationOption) CrossValidation() *CrossValidation
```
CrossValidation returns CrossValidation if IsCrossValidation() is true, nil
otherwise

#### func (WorkflowValidationOption) IsCrossValidation

```go
func (m WorkflowValidationOption) IsCrossValidation() bool
```
IsCrossValidation checks if the WorkflowValidationOption is a CrossValidation

#### func (WorkflowValidationOption) IsRawInterface

```go
func (m WorkflowValidationOption) IsRawInterface() bool
```
IsRawInterface checks if the WorkflowValidationOption is an interface{} (unknown
type)

#### func (WorkflowValidationOption) IsTrainTestSplit

```go
func (m WorkflowValidationOption) IsTrainTestSplit() bool
```
IsTrainTestSplit checks if the WorkflowValidationOption is a TrainTestSplit

#### func (WorkflowValidationOption) MarshalJSON

```go
func (m WorkflowValidationOption) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals WorkflowValidationOption using
WorkflowValidationOption.WorkflowValidationOption

#### func (WorkflowValidationOption) RawInterface

```go
func (m WorkflowValidationOption) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (WorkflowValidationOption) TrainTestSplit

```go
func (m WorkflowValidationOption) TrainTestSplit() *TrainTestSplit
```
TrainTestSplit returns TrainTestSplit if IsTrainTestSplit() is true, nil
otherwise

#### func (*WorkflowValidationOption) UnmarshalJSON

```go
func (m *WorkflowValidationOption) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals WorkflowValidationOption into CrossValidation,
TrainTestSplit, or interface{} if no matches are found

#### type WorkflowValidationScore

```go
type WorkflowValidationScore struct {
}
```

WorkflowValidationScore is CrossValidationScore, TrainTestScore, (or interface{}
if no matches are found)

#### func  MakeWorkflowValidationScoreFromCrossValidationScore

```go
func MakeWorkflowValidationScoreFromCrossValidationScore(f CrossValidationScore) WorkflowValidationScore
```
MakeWorkflowValidationScoreFromCrossValidationScore creates a new
WorkflowValidationScore from an instance of CrossValidationScore

#### func  MakeWorkflowValidationScoreFromRawInterface

```go
func MakeWorkflowValidationScoreFromRawInterface(f interface{}) WorkflowValidationScore
```
MakeWorkflowValidationScoreFromRawInterface creates a new
WorkflowValidationScore from a raw interface{}

#### func  MakeWorkflowValidationScoreFromTrainTestScore

```go
func MakeWorkflowValidationScoreFromTrainTestScore(f TrainTestScore) WorkflowValidationScore
```
MakeWorkflowValidationScoreFromTrainTestScore creates a new
WorkflowValidationScore from an instance of TrainTestScore

#### func (WorkflowValidationScore) CrossValidationScore

```go
func (m WorkflowValidationScore) CrossValidationScore() *CrossValidationScore
```
CrossValidationScore returns CrossValidationScore if IsCrossValidationScore() is
true, nil otherwise

#### func (WorkflowValidationScore) IsCrossValidationScore

```go
func (m WorkflowValidationScore) IsCrossValidationScore() bool
```
IsCrossValidationScore checks if the WorkflowValidationScore is a
CrossValidationScore

#### func (WorkflowValidationScore) IsRawInterface

```go
func (m WorkflowValidationScore) IsRawInterface() bool
```
IsRawInterface checks if the WorkflowValidationScore is an interface{} (unknown
type)

#### func (WorkflowValidationScore) IsTrainTestScore

```go
func (m WorkflowValidationScore) IsTrainTestScore() bool
```
IsTrainTestScore checks if the WorkflowValidationScore is a TrainTestScore

#### func (WorkflowValidationScore) MarshalJSON

```go
func (m WorkflowValidationScore) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals WorkflowValidationScore using
WorkflowValidationScore.WorkflowValidationScore

#### func (WorkflowValidationScore) RawInterface

```go
func (m WorkflowValidationScore) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (WorkflowValidationScore) TrainTestScore

```go
func (m WorkflowValidationScore) TrainTestScore() *TrainTestScore
```
TrainTestScore returns TrainTestScore if IsTrainTestScore() is true, nil
otherwise

#### func (*WorkflowValidationScore) UnmarshalJSON

```go
func (m *WorkflowValidationScore) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals WorkflowValidationScore into CrossValidationScore,
TrainTestScore, or interface{} if no matches are found

#### type WorkflowsGetResponse

```go
type WorkflowsGetResponse struct {
	CreationTime *string `json:"creationTime,omitempty"`
	Id           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
}
```
