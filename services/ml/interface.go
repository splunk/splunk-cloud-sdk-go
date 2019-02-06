// Code generated by gen_interface.go. DO NOT EDIT.

package ml

// Servicer ...
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
	//GetWorkflowBuilds Get list of workflow builds
	GetWorkflowBuilds(id string) ([]WorkflowBuild, error)
	//GetWorkflowRun Get status of a workflow run
	GetWorkflowRun(id string, buildID string, runID string) (*WorkflowRun, error)
	//GetWorkflowRuns Get list of workflow runs
	GetWorkflowRuns(id string, buildID string) ([]WorkflowRun, error)
	//GetWorkflows Get the list of workflow configurations
	GetWorkflows() ([]WorkflowsGetResponse, error)
}
