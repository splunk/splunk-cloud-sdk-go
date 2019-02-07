// Copyright © 2019 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"fmt"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"

	"github.com/splunk/splunk-cloud-sdk-go/services/ml"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testTime = testutils.TimeSec

func TestCreateWorkflow(t *testing.T) {
	client := getSdkClient(t)

	var name = fmt.Sprintf("gosdk_wf_%d", testTime)

	createdWorkFlow := newWorkflow(t, client)

	assert.NotEmpty(t, createdWorkFlow.ID)
	assert.Equal(t, name, *createdWorkFlow.Name)
	assert.NotEmpty(t, createdWorkFlow.CreationTime)
	assert.NotEmpty(t, createdWorkFlow.Tasks)
}

func TestCreateWorkflowBuild(t *testing.T) {
	client := getSdkClient(t)

	//source := fmt.Sprintf("gosdk_wfb_data_%d", testTime)
	//f, err := os.Open("../data/iris.csv")
	//require.Nil(t, err)
	//require.NotNil(t, f)
	//
	//bytes, err := ioutil.ReadAll(f)
	//require.Nil(t, err)
	//require.NotNil(t, bytes)
	//
	//events := []ingest.Event{
	//	{
	//		Body:   string(bytes),
	//		Source: source,
	//	},
	//}
	//
	//err = client.IngestService.PostEvents(events)
	//require.Nil(t, err)

	// Wait for events to be ingested
	//time.Sleep(10 * time.Second)

	createdWorkFlow := newWorkflow(t, client)

	var buildName = fmt.Sprintf("gosdk_wfb_%d", testTime)

	// TODO: need to ingest before this runs

	createdWorkflowBuild := newWorkflowBuild(t, client, createdWorkFlow.ID)

	assert.NotEmpty(t, createdWorkflowBuild.ID)
	assert.Equal(t, buildName, *createdWorkflowBuild.Name)
	assert.NotEmpty(t, createdWorkflowBuild.CreationTime)
	assert.NotEqual(t, ml.FailedWorkflowBuildStatus, createdWorkflowBuild.Status)
}

// TODO: revisit workflow runs with batter workflow build data
func TestCreateWorkflowRun(t *testing.T) {
	client := getSdkClient(t)

	createdWorkFlow := newWorkflow(t, client)
	createdWorkflowBuild := newWorkflowBuild(t, client, createdWorkFlow.ID)

	createdWorkflowRun := newWorkflowRun(t, client, createdWorkFlow.ID, createdWorkflowBuild.ID)
	assert.NotNil(t, createdWorkflowRun.ID)
}

func TestDeleteWorkflow(t *testing.T) {
	client := getSdkClient(t)

	createdWorkFlow := newWorkflow(t, client)

	err := client.MachineLearningService.DeleteWorkflow(*createdWorkFlow.ID)
	require.Nil(t, err)
}

func TestDeleteWorkflowBuild(t *testing.T) {
	client := getSdkClient(t)

	workFlow := newWorkflow(t, client)
	workFlowBuild := newWorkflowBuild(t, client, workFlow.ID)

	err := client.MachineLearningService.DeleteWorkflowBuild(*workFlow.ID, *workFlowBuild.ID)
	require.Nil(t, err)

	// Cleanup
	err = client.MachineLearningService.DeleteWorkflow(*workFlow.ID)
	require.Nil(t, err)
}

func TestDeleteWorkflowRun(t *testing.T) {
	client := getSdkClient(t)

	workFlow := newWorkflow(t, client)
	workFlowBuild := newWorkflowBuild(t, client, workFlow.ID)
	workFlowRun := newWorkflowRun(t, client, workFlow.ID, workFlowBuild.ID)

	err := client.MachineLearningService.DeleteWorkflowRun(*workFlow.ID, *workFlowBuild.ID, *workFlowRun.ID)
	require.Nil(t, err)

	// Cleanup
	err = client.MachineLearningService.DeleteWorkflowBuild(*workFlow.ID, *workFlowBuild.ID)
	require.Nil(t, err)
	err = client.MachineLearningService.DeleteWorkflow(*workFlow.ID)
	require.Nil(t, err)
}

func TestListWorkflows(t *testing.T) {
	client := getSdkClient(t)

	workflows, err := client.MachineLearningService.ListWorkflows()
	require.Nil(t, err)
	require.NotNil(t, workflows)
	assert.NotEmpty(t, workflows)
}

func TestListWorkflowBuilds(t *testing.T) {
	client := getSdkClient(t)

	workflow := newWorkflow(t, client)

	workflowBuilds, err := client.MachineLearningService.ListWorkflowBuilds(*workflow.ID)
	require.Nil(t, err)
	require.NotNil(t, workflowBuilds)
	assert.Empty(t, workflowBuilds)
}

func TestListWorkflowRuns(t *testing.T) {
	client := getSdkClient(t)

	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.ID)

	workflowRuns, err := client.MachineLearningService.ListWorkflowRuns(*workflow.ID, *workflowBuild.ID)
	require.Nil(t, err)
	require.NotNil(t, workflowRuns)
	assert.Empty(t, workflowRuns)
}

func TestGetWorkflow(t *testing.T) {
	client := getSdkClient(t)

	workflow := newWorkflow(t, client)

	workflowRetrieved, err := client.MachineLearningService.GetWorkflow(*workflow.ID)
	require.Nil(t, err)
	require.NotNil(t, workflow)
	assert.Equal(t, workflowRetrieved.ID, workflow.ID)
}

func TestGetWorkflowBuild(t *testing.T) {
	client := getSdkClient(t)

	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.ID)

	workflowBuildRetrieved, err := client.MachineLearningService.GetWorkflowBuild(*workflow.ID, *workflowBuild.ID)
	require.Nil(t, err)
	require.NotNil(t, workflow)
	assert.Equal(t, workflowBuildRetrieved.ID, workflowBuild.ID)
}

func TestGetWorkflowRun(t *testing.T) {
	client := getSdkClient(t)

	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.ID)
	workflowRun := newWorkflowRun(t, client, workflow.ID, workflowBuild.ID)

	workflowRunRetrieved, err := client.MachineLearningService.GetWorkflowRun(*workflow.ID, *workflowBuild.ID, *workflowRun.ID)
	require.Nil(t, err)
	require.NotNil(t, workflow)
	assert.Equal(t, workflowRunRetrieved.ID, workflowRun.ID)
}

// TODO: once done with implementation, clean up all 3 levels of entities
// TODO: audit all tests for resource cleanup

// Util functions

func newWorkflow(t *testing.T, client *sdk.Client) *ml.Workflow {
	var taskName = "fitTask"
	var kind = ml.FitTaskKind
	var task = ml.Task{
		Name:      &taskName,
		Kind:      &kind,
		Algorithm: "PCA",
		Fields: ml.Fields{
			Features: []string{"host", "index"},
		},
		TimeoutSecs: 2,
	}

	var name = fmt.Sprintf("gosdk_wf_%d", testTime)
	var workflow = ml.Workflow{
		Name:  &name,
		Tasks: []ml.Task{task},
	}

	createdWorkFlow, err := client.MachineLearningService.CreateWorkflow(workflow)
	require.Nil(t, err)
	require.NotEmpty(t, createdWorkFlow)

	return createdWorkFlow
}

func newWorkflowBuild(t *testing.T, client *sdk.Client, workflowID *string) *ml.WorkflowBuild {
	// TODO: need to ingest before this runs
	var buildName = fmt.Sprintf("gosdk_wfb_%d", testTime)
	extract := true
	build := ml.WorkflowBuild{
		Name: &buildName,
		Input: ml.InputData{
			Kind: ml.SPLInputKind,
			Source: ml.InputDataSource{
				Query:            fmt.Sprintf("| from main where source=\"%s\"", "sdk_ml_csv_import"),
				ExtractAllFields: &extract,
			},
		},
	}

	createdWorkflowBuild, err := client.MachineLearningService.CreateWorkflowBuild(*workflowID, build)
	require.Nil(t, err)
	require.NotEmpty(t, createdWorkflowBuild)

	return createdWorkflowBuild
}

func newWorkflowRun(t *testing.T, client *sdk.Client, workflowID *string, workflowBuildID *string) *ml.WorkflowRun {
	// Wait for the build to be available, smaller delay is probably possible
	time.Sleep(45 * time.Second)

	extract := true
	outputKind := ml.HecOutputKind
	outputSource := fmt.Sprintf("sdk_ml_csv_export")

	outputDestination := ml.OutputDataDestination{
		Source: &outputSource,
	}

	var run = ml.WorkflowRun{
		Input: ml.InputData{
			Kind: ml.SPLInputKind,
			Source: ml.InputDataSource{
				Query:            fmt.Sprintf("| from main where source=\"%s\"", "sdk_ml_csv_import"),
				ExtractAllFields: &extract,
			},
		},
		Output: ml.OutputData{
			Kind:        &outputKind,
			Destination: &outputDestination,
		},
	}

	createdWorkflowRun, err := client.MachineLearningService.CreateWorkflowRun(*workflowID, *workflowBuildID, run)
	require.Nil(t, err)
	require.NotEmpty(t, createdWorkflowRun)

	return createdWorkflowRun
}
