// Copyright © 2019 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package ml

// TODO: add HTTP methods annotations to all models here

type Fields struct {
	// Fields necessary for task.
	Features []string `json:"features"`
	// Fields produced by task.
	Created *[]string `json:"created,omitempty"`
	// Target field necessary for task.
	Target *string `json:"target,omitempty"`
}

// Fit task does an estimation/training/fitting, fit task outputs a transformer
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

type Health struct {
	Healthy bool   `json:"healthy"`
	Version string `json:"version"`
}

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

type InputData struct {
	Kind   InputDataKind   `json:"kind"`
	Source InputDataSource `json:"source"`
}

type InputDataKind string

// List of InputDataKind
const (
	S3InputKind      InputDataKind = "S3"
	SPLInputKind     InputDataKind = "SPL"
	RawDataInputKind InputDataKind = "RawData"
)

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

type OutputData struct {
	Kind        *OutputDataKind        `json:"kind,omitempty"`
	Destination *OutputDataDestination `json:"destination,omitempty"`
}

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

type OutputDataKind string

// List of OutputDataKind
const (
	S3OutputKind  OutputDataKind = "S3"
	HecOutputKind OutputDataKind = "HEC"
)

type RawData struct {
	Data *string `json:"data,omitempty"`
}

type S3Key struct {
	Key *string `json:"key,omitempty"`
}

type Spl struct {
	Query string `json:"query"`
	// Determine whether the Search service extracts all available fields in the data, including fields not mentioned in the SPL for the search job. Set to 'false' for better search performance.
	ExtractAllFields *bool `json:"extractAllFields,omitempty"`
	// The number of seconds to run this search before finalizing
	MaxTime *int32 `json:"maxTime,omitempty"`
	// Represents parameters on the search job such as 'earliest' and 'latest'.
	QueryParameters *map[string]interface{} `json:"queryParameters,omitempty"`
}

type TaskCommon struct {
	// The name has to be unique in the same workflow, it is optional, can be used to identify that task artifact
	Name *string   `json:"name,omitempty"`
	Kind *TaskKind `json:"kind,omitempty"`
}

type TaskKind string

// List of TaskKind
const (
	FitTaskKind TaskKind = "fit"
)

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

type Workflow struct {
	ID           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Tasks        []Task  `json:"tasks"`
	CreationTime *string `json:"creationTime,omitempty"`
}

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

type WorkflowBuildStatus string

// List of WorkflowBuildStatus
const (
	RunningWorkflowBuildStatus   WorkflowBuildStatus = "running"
	FailedWorkflowBuildStatus    WorkflowBuildStatus = "failed"
	SuccessWorkflowBuildStatus   WorkflowBuildStatus = "success"
	ScheduledWorkflowBuildStatus WorkflowBuildStatus = "scheduled"
)

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

type WorkflowRunStatus string

// List of WorkflowRunStatus
const (
	RunningWorkflowRunStatus   WorkflowRunStatus = "running"
	FailedWorkflowRunStatus    WorkflowRunStatus = "failed"
	SuccessWorkflowRunStatus   WorkflowRunStatus = "success"
	ScheduledWorkflowRunStatus WorkflowRunStatus = "scheduled"
)

type WorkflowsGetResponse struct {
	ID           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	CreationTime *string `json:"creationTime,omitempty"`
}
