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

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/services/ml"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"bufio"
	"encoding/json"

	"io"
	"path/filepath"
	"strconv"
)

const mlData = "./data/iris.csv"

var workflowName = "test_go_workflow"
var Base64DATA = "LHNlcGFsX2xlbmd0aCxzZXBhbF93aWR0aCxwZXRhbF9sZW5ndGgscGV0YWxfd2lkdGgsc3BlY2llcw0KMCw1LjEsMy41LDEuNCwwLjIsSXJpcyBTZXRvc2ENCjUwLDcuMCwzLjIsNC43LDEuNCxJcmlzIFZlcnNpY29sb3INCjEwMCw2LjMsMy4zLDYuMCwyLjUsSXJpcyBWaXJnaW5pY2ENCg=="

var source = "go-tests"
var sourceType = "json"

var PkceTokenRetriver = idp.NewPKCERetriever(testutils.PkceClientID, testutils.NativeAppRedirectURI,
	idp.DefaultOIDCScopes, testutils.Username, testutils.Password, testutils.IdpHost)

var sdkClient, sdkErr = testutils.MakeSdkClient(PkceTokenRetriver, testutils.TestMLTenant)

func TestMain(t *testing.M) {
	code := t.Run()
	cleanup()
	os.Exit(code)
}

func getSdkClient(t *testing.T) *sdk.Client {
	require.Emptyf(t, sdkErr, "error calling sdk.NewClient(): %s", sdkErr)
	return sdkClient
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

func TestIngestData(t *testing.T) {
	// read in iris.csv file and populate data array
	var data = make([]string, 0)
	fp, _ := filepath.Abs(mlData)
	f, err := os.Open(fp)
	defer f.Close()
	if err != nil {
		assert.Error(t, err)
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		data = append(data, line)
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			assert.Error(t, err)
		}
	}
	// transform to array of maps {header => column}
	headers := strings.Split(data[0], ",")
	rawData := data[1:]
	var maps = make([]map[string]interface{}, 0)
	var events = make([]ingest.Event, 0)
	for _, e := range rawData {
		cols := strings.Split(e, ",")
		for i, c := range cols {
			// parse cols at idx 0-3 to floats
			if i < 4 {
				f, err := strconv.ParseFloat(c, 64)
				maps = append(maps, map[string]interface{}{headers[i]: f})
				if err != nil {
					assert.Error(t, err)
				}
				// last col at idx 4 is a string
			} else {
				maps = append(maps, map[string]interface{}{headers[i]: c})
			}
		}
		payload, err := json.Marshal(maps)
		if err != nil {
			assert.Error(t, err)
		}
		events = append(events, ingest.Event{Body: string(payload), Sourcetype: &sourceType,
			Source: &source, Attributes: map[string]interface{}{"index": "main"}})
	}
	// send data to ingest API
	client := getSdkClient(t)
	_, err = client.IngestService.PostEvents(events)
	if err != nil {
		assert.Error(t, err)
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
	require.NoError(t, err)
	require.NotNil(t, workflows)
	assert.IsType(t, []ml.WorkflowsGetResponse{}, workflows)
}

func TestListWorkflowBuilds(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowBuilds, err := client.MachineLearningService.ListWorkflowBuilds(*workflow.Id)
	require.NoError(t, err)
	require.NotNil(t, workflowBuilds)
	assert.Empty(t, workflowBuilds)
}

func TestListWorkflowRuns(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.Id)
	workflowRuns, err := client.MachineLearningService.ListWorkflowRuns(*workflow.Id, *workflowBuild.Id)
	require.NoError(t, err)
	require.NotNil(t, workflowRuns)
	assert.Empty(t, workflowRuns)
}

func TestGetWorkflow(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowRetrieved, err := client.MachineLearningService.GetWorkflow(*workflow.Id)
	require.NoError(t, err)
	require.NotNil(t, workflow)
	assert.Equal(t, workflowRetrieved.Id, workflow.Id)
}

func TestGetWorkflowBuild(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.Id)
	workflowBuildRetrieved, err := client.MachineLearningService.GetWorkflowBuild(*workflow.Id, *workflowBuild.Id)
	require.NoError(t, err)
	require.NotNil(t, workflow)
	assert.Equal(t, workflowBuildRetrieved.Id, workflowBuild.Id)
}

func TestGetWorkflowRun(t *testing.T) {
	client := getSdkClient(t)
	workflow := newWorkflow(t, client)
	workflowBuild := newWorkflowBuild(t, client, workflow.Id)
	workflowRun := newWorkflowRun(t, client, workflow.Id, workflowBuild.Id)
	workflowRunRetrieved, err := client.MachineLearningService.GetWorkflowRun(*workflow.Id, *workflowBuild.Id, *workflowRun.Id)
	require.NoError(t, err)
	require.NotNil(t, workflow)
	assert.Equal(t, workflowRunRetrieved.Id, workflowRun.Id)
}

// Util functions

// Maps to data/01_ml_workflow.json
func newWorkflow(t *testing.T, client *sdk.Client) *ml.Workflow {
	var taskName1 = "PCA"
	var target1 = ""
	var outputTransformer1 = "PCA_model"
	var parameters1 = map[string]interface{}{
		"k": "3",
	}

	kind := ml.FitTaskKindFit
	var fitTask1 = ml.FitTask{
		Kind:      &kind,
		Name:      &taskName1,
		Algorithm: "PCA",
		Fields: ml.Fields{
			Features: []string{
				"petal_length",
				"petal_width",
				"sepal_length",
				"sepal_width"},
			Target:  &target1,
			Created: []string{"PC_1", "PC_2", "PC_3"},
		},
		OutputTransformer: &outputTransformer1,
		Parameters:        parameters1,
	}
	var task1 = ml.MakeTaskFromFitTask(fitTask1)

	var taskName2 = "RandomForestClassifier"
	var target2 = "species"
	var outputTransformer2 = "RFC_model"
	var parameters2 = map[string]interface{}{
		"n_estimators":      25,
		"max_depth":         10,
		"min_samples_split": 5,
		"max_features":      "auto",
		"criterion":         "gini"}

	var fitTask2 = ml.FitTask{
		Kind:      &kind,
		Name:      &taskName2,
		Algorithm: "RandomForestClassifier",
		Fields: ml.Fields{
			Features: []string{
				"petal_length",
				"petal_width",
				"sepal_length",
				"sepal_width"},
			Target: &target2,
		},
		OutputTransformer: &outputTransformer2,
		Parameters:        parameters2,
	}
	var task2 = ml.MakeTaskFromFitTask(fitTask2)

	var workflow = ml.Workflow{
		Name:  &workflowName,
		Tasks: []ml.Task{task1, task2},
	}

	createdWorkFlow, err := client.MachineLearningService.CreateWorkflow(workflow)
	require.NoError(t, err)
	require.NotEmpty(t, createdWorkFlow)
	return createdWorkFlow
}

// Maps to data/02_ml_build.json
func newWorkflowBuild(t *testing.T, client *sdk.Client, workflowID *string) *ml.WorkflowBuild {
	build := ml.WorkflowBuild{
		Input: ml.InputData{
			Kind:   ml.InputDataKindRawData,
			Source: ml.MakeInputDataSourceFromRawData(ml.RawData{Data: &Base64DATA}),
		},
	}
	createdWorfklow, err := client.MachineLearningService.CreateWorkflowBuild(*workflowID, build)
	require.NoError(t, err)
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
		require.NoError(t, err)
	}
	return nil
}

func newWorkflowRun(t *testing.T, client *sdk.Client, workflowID *string, workflowBuildID *string) *ml.WorkflowRun {
	var outputKind ml.OutputDataKind = "S3"

	des := ml.MakeOutputDataDestinationFromRawInterface(&map[string]interface{}{"key": "iris.csv"})
	var run = ml.WorkflowRun{
		Input: ml.InputData{
			Kind:   ml.InputDataKindRawData,
			Source: ml.MakeInputDataSourceFromRawData(ml.RawData{Data: &Base64DATA}),
		},
		Output: ml.OutputData{
			Kind:        &outputKind,
			Destination: &des,
		},
	}
	createdWorkflowRun, err := client.MachineLearningService.CreateWorkflowRun(*workflowID, *workflowBuildID, run)
	require.NoError(t, err)
	require.NotEmpty(t, createdWorkflowRun)
	return createdWorkflowRun
}

func cleanupWorkflow(t *testing.T, client *sdk.Client, workflowID *string) {
	err := client.MachineLearningService.DeleteWorkflow(*workflowID)
	if t != nil {
		require.NoError(t, err)
	}
}

func cleanupWorkflowBuild(t *testing.T, client *sdk.Client, workflowID *string, workflowBuildID *string) {
	err := client.MachineLearningService.DeleteWorkflowBuild(*workflowID, *workflowBuildID)
	require.NoError(t, err)
}

func cleanupWorkflowRun(t *testing.T, client *sdk.Client, workflowID *string, workflowBuildID *string, workflowRunID *string) {
	err := client.MachineLearningService.DeleteWorkflowRun(*workflowID, *workflowBuildID, *workflowRunID)
	require.NoError(t, err)
}
