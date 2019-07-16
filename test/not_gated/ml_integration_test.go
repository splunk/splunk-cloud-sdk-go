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

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var hostTrain = "server_power_train_ef5wlcd4njiovmdl"
var hostTest = "server_power_test_ef5wlcd4njiovmdl"
var hostOut = "server_power_out_ef5wlcd4njiovmdl"
var module = "mlapishowcase"

var workflowName = "PredictServerPowerConsumption"
var buildSpl = fmt.Sprintf("| from mlapishowcase.mlapishowcase where host=\"%s\"", hostTrain)
var runSpl = fmt.Sprintf("| from mlapishowcase.mlapishowcase where host=\"%s\"", hostTest)

var PkceTokenRetriver = idp.NewPKCERetriever(testutils.PkceClientID, testutils.NativeAppRedirectURI,
	idp.DefaultOIDCScopes, testutils.Username, testutils.Password, testutils.IdpHost)

func TestMain(t *testing.M) {
	code := t.Run()
	cleanup()
	os.Exit(code)
}

func getSdkClient(t *testing.T) *sdk.Client {
	client, err := testutils.MakeSdkClient(PkceTokenRetriver, testutils.TestMLTenant)
	require.Emptyf(t, err, "error calling sdk.NewClient(): %s", err)
	return client
}

// Try to delete all workflows that were created in this test run, ignoring errors
func cleanup() {
	client, _ := testutils.MakeSdkClient(PkceTokenRetriver, testutils.TestMLTenant)
	workflows, _ := client.MachineLearningService.ListWorkflows()
	for i := 0; i < len(workflows); i++ {
		if strings.HasPrefix(*workflows[i].Name, workflowName) {
			_ = client.MachineLearningService.DeleteWorkflow(*workflows[i].Id)
		}
	}
}

func TestCreateWorkflow(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	assert.NotEmpty(t, workflow.Id)
	assert.Equal(t, workflowName, *workflow.Name)
	assert.NotEmpty(t, workflow.CreationTime)
	assert.NotEmpty(t, workflow.Tasks)
	cleanupWorkflow(t, client, workflow.Id)
}

func TestCreateWorkflowBuild(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.Id)
	assert.NotEmpty(t, workflowBuild.Id)
	assert.NotEmpty(t, workflowBuild.CreationTime)
	assert.NotEqual(t, ml.WorkflowBuildStatusFailed, workflowBuild.Status)

	// Deleting a workflow will delete the builds also
	cleanupWorkflow(t, client, workflow.Id)
}

func TestCreateWorkflowRun(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.Id)
	workflowRun := newWorkflowRun(t, client, workflow.Id, workflowBuild.Id)
	assert.NotNil(t, workflowRun)
	assert.NotEqual(t, ml.WorkflowBuildStatusFailed, workflowRun.Id)

	// Deleting a workflow will delete the builds and runs also
	cleanupWorkflow(t, client, workflow.Id)
}

func TestDeleteWorkflow(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	cleanupWorkflow(t, client, workflow.Id)
}

func TestDeleteWorkflowBuild(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workFlowBuild := newWorkflowBuild(t, client, workflow.Id)
	cleanupWorkflowBuild(t, client, workflow.Id, workFlowBuild.Id)
	cleanupWorkflow(t, client, workflow.Id)
}

func TestDeleteWorkflowRun(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.Id)
	workflowRun := newWorkflowRun(t, client, workflow.Id, workflowBuild.Id)
	cleanupWorkflowRun(t, client, workflow.Id, workflowBuild.Id, workflowRun.Id)
	cleanupWorkflowBuild(t, client, workflow.Id, workflowBuild.Id)
	cleanupWorkflow(t, client, workflow.Id)
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
	workflowBuilds, err := client.MachineLearningService.ListWorkflowBuilds(*workflow.Id)
	require.Nil(t, err)
	require.NotNil(t, workflowBuilds)
	assert.Empty(t, workflowBuilds)
}

func TestListWorkflowRuns(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.Id)
	workflowRuns, err := client.MachineLearningService.ListWorkflowRuns(*workflow.Id, *workflowBuild.Id)
	require.Nil(t, err)
	require.NotNil(t, workflowRuns)
	assert.Empty(t, workflowRuns)
}

func TestGetWorkflow(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowRetrieved, err := client.MachineLearningService.GetWorkflow(*workflow.Id)
	require.Nil(t, err)
	require.NotNil(t, workflow)
	assert.Equal(t, workflowRetrieved.Id, workflow.Id)
}

func TestGetWorkflowBuild(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.Id)
	workflowBuildRetrieved, err := client.MachineLearningService.GetWorkflowBuild(*workflow.Id, *workflowBuild.Id)
	require.Nil(t, err)
	require.NotNil(t, workflow)
	assert.Equal(t, workflowBuildRetrieved.Id, workflowBuild.Id)
}

func TestGetWorkflowRun(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.Id)
	workflowRun := newWorkflowRun(t, client, workflow.Id, workflowBuild.Id)
	workflowRunRetrieved, err := client.MachineLearningService.GetWorkflowRun(*workflow.Id, *workflowBuild.Id, *workflowRun.Id)
	require.Nil(t, err)
	require.NotNil(t, workflow)
	assert.Equal(t, workflowRunRetrieved.Id, workflowRun.Id)
}

// Util functions

// Maps to data/01_ml_workflow.json
func newWorkflow(t *testing.T, client *sdk.Client) *ml.Workflow {
	var taskName = "linearregression"
	var target = "ac_power"
	var outputTransformer = "example_server_power"
	var parameters = map[string]interface{}{
		"fit_intercept": "true",
		"normalize":     "false",
	}

	kind := ml.FitTaskKindFit
	var fitTask = ml.FitTask{
		Kind:      &kind,
		Name:      &taskName,
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
			Target: &target,
		},
		OutputTransformer: &outputTransformer,
		Parameters:        parameters,
	}

	var task = ml.MakeTaskFromFitTask(fitTask)
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

	var splDataSource = ml.Spl{ExtractAllFields: &extract, Query: buildSpl, QueryParameters: queryParams, Module: &module}

	build := ml.WorkflowBuild{
		Input: ml.InputData{
			Kind:   ml.InputDataKindSpl,
			Source: ml.MakeInputDataSourceFromSpl(splDataSource),
		},
	}
	createdWorfklow, err := client.MachineLearningService.CreateWorkflowBuild(*workflowID, build)
	require.Nil(t, err)
	require.NotEmpty(t, createdWorfklow)

	buildID := createdWorfklow.Id

	// Check every 5 seconds for the workflow build to complete or fail
	for i := 1; err == nil; time.Sleep(5 * time.Second) {
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
		if *workflowBuild.Status == ml.WorkflowBuildStatusFailed || *workflowBuild.Status == ml.WorkflowBuildStatusSuccess {
			return workflowBuild
		}
		require.Nil(t, err)
	}
	return nil
}

func newWorkflowRun(t *testing.T, client *sdk.Client, workflowID *string, workflowBuildID *string) *ml.WorkflowRun {
	extract := true
	outputKind := ml.OutputDataKindEvents
	outputSource := "mlapi-showcase"

	queryParams := map[string]interface{}{
		"earliest": "0",
		"latest":   "now",
	}

	var splDataSource = ml.Spl{ExtractAllFields: &extract, Query: runSpl, QueryParameters: queryParams, Module: &module}

	events := ml.Events{
		Attributes: map[string]interface{}{
			"index":  "mlapishowcase",
			"module": "mlapishowcase",
		},
		Source: &outputSource,
		Host:   &hostOut,
	}

	des := ml.MakeOutputDataDestinationFromEvents(events)
	var run = ml.WorkflowRun{
		Input: ml.InputData{
			Kind:   ml.InputDataKindSpl,
			Source: ml.MakeInputDataSourceFromSpl(splDataSource),
		},
		Output: ml.OutputData{
			Kind:        &outputKind,
			Destination: &des,
		},
	}
	createdWorkflowRun, err := client.MachineLearningService.CreateWorkflowRun(*workflowID, *workflowBuildID, run)
	require.Nil(t, err)
	require.NotEmpty(t, createdWorkflowRun)
	return createdWorkflowRun
}

func cleanupWorkflow(t *testing.T, client *sdk.Client, workflowID *string) {
	err := client.MachineLearningService.DeleteWorkflow(*workflowID)
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
