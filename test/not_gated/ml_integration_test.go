// Copyright © 2019 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package not_gated

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services/ml"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/splunk/splunk-cloud-sdk-go/services"
)

var hostTrain = "server_power_train_ef5wlcd4njiovmdl"
var hostTest = "server_power_test_ef5wlcd4njiovmdl"
var hostOut = "server_power_out_ef5wlcd4njiovmdl"

var workflowName = "PredictServerPowerConsumption"
var buildSpl = fmt.Sprintf("| from mlapishowcase.mlapishowcase where host=\"%s\"", hostTrain)
var runSpl = fmt.Sprintf("| from mlapishowcase.mlapishowcase where host=\"%s\"", hostTest)

func TestMain(t *testing.M) {
	code := t.Run()
	cleanup()
	os.Exit(code)
}


func getSdkClient(t *testing.T) *sdk.Client {
	client, err := makeSdkClient()
	require.Emptyf(t, err, "error calling sdk.NewClient(): %s", err)
	return client
}

// Get an client without the testing interface
func makeSdkClient() (*sdk.Client, error) {
	return sdk.NewClient(&services.Config{
		Token:   testutils.TestAuthenticationToken,
		Host:    testutils.TestSplunkCloudHost,
		Tenant:  testutils.TestMLTenant,  // testsdksml
		Timeout: testutils.TestTimeOut,
	})
}

// Try to delete all workflows that were created in this test run, ignoring errors
func cleanup() {
	client, _ :=  makeSdkClient()
	workflows, _ := client.MachineLearningService.ListWorkflows()
	for i := 0; i < len(workflows); i++ {
		if strings.HasPrefix(*workflows[i].Name, workflowName) {
			_ = client.MachineLearningService.DeleteWorkflow(*workflows[i].ID)
		}
	}
}

func TestCreateWorkflow(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	assert.NotEmpty(t, workflow.ID)
	assert.Equal(t, workflowName, *workflow.Name)
	assert.NotEmpty(t, workflow.CreationTime)
	assert.NotEmpty(t, workflow.Tasks)
	cleanupWorkflow(t, client, workflow.ID)
}

func TestCreateWorkflowBuild(t *testing.T) {
	client := getSdkClient(t)

	//buildName := fmt.Sprintf("go_sdk_wfb_%d", testutils.TimeSec)

	// TODO: need to ingest before this runs
	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.ID)

	assert.NotEmpty(t, workflowBuild.ID)
	//assert.Equal(t, buildName, *workflowBuild.Name)
	assert.NotEmpty(t, workflowBuild.CreationTime)
	assert.NotEqual(t, ml.FailedWorkflowBuildStatus, workflowBuild.Status)

	// Deleting a workflow will delete the builds also
	cleanupWorkflow(t, client, workflow.ID)
}

func TestCreateWorkflowRun(t *testing.T) {
	client := getSdkClient(t)

	//runName := fmt.Sprintf("go_sdk_wfb_%d", testutils.TimeSec)

	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.ID)
	workflowRun := newWorkflowRun(t, client, workflow.ID, workflowBuild.ID)

	assert.NotNil(t, workflowRun.ID)
	//assert.Equal(t, runName, *workflowRun.Name)
	assert.NotEqual(t, ml.FailedWorkflowRunStatus, workflowRun.Status)

	// Deleting a workflow will delete the builds and runs also
	cleanupWorkflow(t, client, workflow.ID)
}

func TestDeleteWorkflow(t *testing.T) {
	client := getSdkClient(t)

	workflow := newWorkflow(t, client)

	cleanupWorkflow(t, client, workflow.ID)
}

func TestDeleteWorkflowBuild(t *testing.T) {
	client := getSdkClient(t)

	workflow := newWorkflow(t, client)
	workFlowBuild := newWorkflowBuild(t, client, workflow.ID)

	cleanupWorkflowBuild(t, client, workflow.ID, workFlowBuild.ID)
	cleanupWorkflow(t, client, workflow.ID)
}

func TestDeleteWorkflowRun(t *testing.T) {
	client := getSdkClient(t)

	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.ID)
	workflowRun := newWorkflowRun(t, client, workflow.ID, workflowBuild.ID)

	cleanupWorkflowRun(t, client, workflow.ID, workflowBuild.ID, workflowRun.ID)
	cleanupWorkflowBuild(t, client, workflow.ID, workflowBuild.ID)
	cleanupWorkflow(t, client, workflow.ID)
}

func TestListWorkflows(t *testing.T) {
	client := getSdkClient(t)

	workflows, err := client.MachineLearningService.ListWorkflows()

	require.Nil(t, err)
	require.NotNil(t, workflows)
	assert.IsType(t, []ml.WorkflowsGetResponse{}, workflows)
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

// Util functions

// Maps to data/01_ml_workflow.json
func newWorkflow(t *testing.T, client *sdk.Client) *ml.Workflow {
	var taskName = "linearregression"
	var kind = ml.FitTaskKind
	var target = "ac_power"
	var outputTransformer = "example_server_power"
	var parameters = map[string]interface{}{
		"fit_intercept": true,
		"normalize": false,
	}
	var task = ml.Task{
		Name:      &taskName,
		Kind:      &kind,
		Algorithm: "LinearRegression",
		Fields: ml.Fields{
			Features: []string{
				"total-unhalted_core_cycles",
				"total-instructions_retired",
				"total-last_level_cache_references",
				"total-memory_bus_transactions",
				"total-cpu-utilization",
				"total-disk-accesses",
				"total-disk-blocks",
				"total-disk-utilization"},
			Target:   &target,
		},
		OutputTransformer: &outputTransformer,
		Parameters : &parameters,
		TimeoutSecs: 600,
	}
	name := "PredictServerPowerConsumption"
	var workflow = ml.Workflow{
		Name:  &name,
		Tasks: []ml.Task{task},
	}
	createdWorkFlow, err := client.MachineLearningService.CreateWorkflow(workflow)
	require.Nil(t, err)
	require.NotEmpty(t, createdWorkFlow)
	return createdWorkFlow
}

// Maps to data/02_ml_build.json
func newWorkflowBuild(t *testing.T, client *sdk.Client, workflowID *string) *ml.WorkflowBuild {
	extract := true
	queryParams := map[string]interface{}{
		"earliest": "0",
		"latest":   "now",
	}
	build := ml.WorkflowBuild{
		Input: ml.InputData{
			Kind: ml.SPLInputKind,
			Source: ml.InputDataSource{
				Query:            buildSpl,
				ExtractAllFields: &extract,
				QueryParameters:  &queryParams,
			},
		},
	}
	createdWorfklow, err := client.MachineLearningService.CreateWorkflowBuild(*workflowID, build)
	require.Nil(t, err)
	require.NotEmpty(t, createdWorfklow)

	buildID := createdWorfklow.ID

	// Check every 5 seconds for the workflow build to complete or fail
	for i := 1; err == nil; time.Sleep(5 * time.Second) {
		//fmt.Printf("Waiting for workflowID=%s buildID=%s to complete, time waited (s): %d\n", *workflowID, *buildID, 5*i)
		if workflowID == nil {
			fmt.Println("workflow ID is nil")
		}
		if buildID == nil {
			fmt.Println("buildID is nil")
		}
		fmt.Printf("Waiting for workflowID=%s buildID=%s to complete\n", *workflowID, *buildID)
		i++

		workflowBuild, err := client.MachineLearningService.GetWorkflowBuild(*workflowID, *buildID)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Printf("\tworkflowID=%s buildID=%s: %s\n", *workflowID, *buildID, string(*workflowBuild.Status))
		if *workflowBuild.Status == ml.FailedWorkflowBuildStatus || *workflowBuild.Status == ml.SuccessWorkflowBuildStatus {
			return workflowBuild
		}

		require.Nil(t, err)
	}

	return nil
}

func newWorkflowRun(t *testing.T, client *sdk.Client, workflowID *string, workflowBuildID *string) *ml.WorkflowRun {
	extract := true
	outputKind := ml.HecOutputKind
	outputSource := "mlapi-showcase"

	queryParams := map[string]interface{}{
		"earliest": "0",
		"latest":   "now",
	}
	//runName := fmt.Sprintf("go_sdk_wfb_%d", testutils.TimeSec)
	var run = ml.WorkflowRun{
		Input: ml.InputData{
			Kind: ml.SPLInputKind,
			Source: ml.InputDataSource{
				Query:            runSpl,
				ExtractAllFields: &extract,
				QueryParameters:  &queryParams,
			},
		},
		Output: ml.OutputData{
			Kind:        &outputKind,
			Destination: &ml.OutputDataDestination{
				Attributes: &map[string]interface{}{
					"index": "mlapishowcase",
					"module": "mlapishowcase",
	},
				Source: &outputSource,
				Host: &hostOut,
			},
		},
	}

	createdWorkflowRun, err := client.MachineLearningService.CreateWorkflowRun(*workflowID, *workflowBuildID, run)
	require.Nil(t, err)
	require.NotEmpty(t, createdWorkflowRun)

	return createdWorkflowRun
}

func cleanupWorkflow(t *testing.T, client *sdk.Client, workflowID *string) {
	err := client.MachineLearningService.DeleteWorkflow(*workflowID)
	// This condition allows us to reuse the function for cleanup()
	if t != nil {
		require.Nil(t, err)
	}
}

func cleanupWorkflowBuild(t *testing.T, client *sdk.Client, workflowID *string, workflowBuildID *string) {
	err := client.MachineLearningService.DeleteWorkflowBuild(*workflowID, *workflowBuildID)
	require.Nil(t, err)
}

func cleanupWorkflowRun(t *testing.T, client *sdk.Client, workflowID *string, workflowBuildID *string, workflowRunID *string) {
	err := client.MachineLearningService.DeleteWorkflowRun(*workflowID, *workflowBuildID, *workflowRunID)
	require.Nil(t, err)
}


