// Copyright © 2019 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"

	"github.com/splunk/splunk-cloud-sdk-go/services/search"

	"github.com/splunk/splunk-cloud-sdk-go/sdk"

	"github.com/splunk/splunk-cloud-sdk-go/services/ml"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var workflowName = fmt.Sprintf("go_sdk_wf_%d", testutils.TimeSec)
var source = "sdk_ml_csv_import"
var inputQuerySPL = fmt.Sprintf("| from main where source=\"%s\"", source)

func TestMain(t *testing.M) {
	code := t.Run()
	cleanup()
	os.Exit(code)
}

// Try to delete all workflows that were created in this test run, ignoring errors
func cleanup() {
	client, _ := MakeSdkClient()
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

	//source := fmt.Sprintf("go_sdk_wfb_data_%d", testutils.TimeSec)
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

	buildName := fmt.Sprintf("go_sdk_wfb_%d", testutils.TimeSec)

	// TODO: need to ingest before this runs
	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.ID)

	assert.NotEmpty(t, workflowBuild.ID)
	assert.Equal(t, buildName, *workflowBuild.Name)
	assert.NotEmpty(t, workflowBuild.CreationTime)
	assert.NotEqual(t, ml.FailedWorkflowBuildStatus, workflowBuild.Status)

	// Deleting a workflow will delete the builds also
	cleanupWorkflow(t, client, workflow.ID)
}

func TestCreateWorkflowRun(t *testing.T) {
	client := getSdkClient(t)

	runName := fmt.Sprintf("go_sdk_wfb_%d", testutils.TimeSec)

	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.ID)
	workflowRun := newWorkflowRun(t, client, workflow.ID, workflowBuild.ID)

	assert.NotNil(t, workflowRun.ID)
	assert.Equal(t, runName, *workflowRun.Name)
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

	var name = fmt.Sprintf(workflowName)
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
	// TODO: this will be necessary for each new tenant, and will make tests more robust, will revisit
	//ensureWorkflowCSVData(t, client)

	var buildName = fmt.Sprintf("go_sdk_wfb_%d", testutils.TimeSec)
	extract := true
	queryParams := map[string]interface{}{
		"earliest": "0",
		"latest":   "now",
	}
	build := ml.WorkflowBuild{
		Name: &buildName,
		Input: ml.InputData{
			Kind: ml.SPLInputKind,
			Source: ml.InputDataSource{
				Query:            fmt.Sprintf("| from main where source=\"%s\"", "sdk_ml_csv_import"),
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
	outputSource := fmt.Sprintf("sdk_ml_csv_export")

	outputDestination := ml.OutputDataDestination{
		Source: &outputSource,
	}
	queryParams := map[string]interface{}{
		"earliest": "0",
		"latest":   "now",
	}
	runName := fmt.Sprintf("go_sdk_wfb_%d", testutils.TimeSec)
	var run = ml.WorkflowRun{
		Name: &runName,
		Input: ml.InputData{
			Kind: ml.SPLInputKind,
			Source: ml.InputDataSource{
				Query:            fmt.Sprintf(inputQuerySPL),
				ExtractAllFields: &extract,
				QueryParameters:  &queryParams,
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

func ensureWorkflowCSVData(t *testing.T, client *sdk.Client) {
	// First, check if we already have data
	searchRequest := search.CreateJobRequest{
		Query: inputQuerySPL,
	}
	job, err := client.SearchService.CreateJob(&searchRequest)
	require.Nil(t, err)
	require.NotNil(t, job)

	state, err := client.SearchService.WaitForJob(job.ID, time.Second)
	require.NotNil(t, err)
	require.Equal(t, search.Done, state)

	rawResults, err := client.SearchService.GetResults(job.ID, 0, 0)
	require.Nil(t, err)
	require.NotNil(t, rawResults)
	require.IsType(t, search.Results{}, rawResults)

	results := rawResults.(*search.Results)

	// TODO: this might be not be right way to check for no results
	// If there's no data, ingest it
	if len(results.Results) == 0 {
		ingestWorkflowCSVData(t, client)
		// TODO: sleep & verify that data has been ingested and ready to build a workflow
	}
}

func ingestWorkflowCSVData(t *testing.T, client *sdk.Client) {
	file, err := os.Open("../data/iris.csv")
	require.Nil(t, err)
	require.NotNil(t, file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	require.Nil(t, err)
	require.NotNil(t, records)

	headers := records[0]

	events := make([]ingest.Event, 0, len(records))
	for i := 1; i < len(records); i++ {
		// Convert the CSV rows to a header->value map
		body := make(map[string]interface{})
		for j := 0; j < len(headers); j++ {
			body[headers[j]] = records[i][j]
		}

		event := ingest.Event{
			Body:   body,
			Source: source,
		}
		events = append(events, event)
	}

	err = client.IngestService.PostEvents(events)
	require.Nil(t, err)
}
