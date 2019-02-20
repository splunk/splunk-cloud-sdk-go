package integration

import (
	"fmt"
	"testing"

	"strconv"

	"net/url"

	"time"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/service"
	"github.com/splunk/splunk-cloud-sdk-go/services/streams"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test variables
var testPipelineDescription = "integration test pipeline"
var testTemplateDescription = "integration test template"

// Test GetPipelines streams endpoint
func TestIntegrationGetAllPipelines(t *testing.T) {
	pipelineName1 := fmt.Sprintf("testPipeline01%d", testutils.TimeSec)
	pipelineName2 := fmt.Sprintf("testPipeline02%d", testutils.TimeSec)

	// Create two test pipelines
	pipeline1, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName1, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), pipeline1.ID, pipeline1.Name)
	require.NotEmpty(t, pipeline1)
	assert.Equal(t, model.Created, pipeline1.Status)
	assert.Equal(t, pipelineName1, pipeline1.Name)
	assert.Equal(t, testPipelineDescription, pipeline1.Description)

	pipeline2, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName2, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), pipeline2.ID, pipeline2.Name)
	require.NotEmpty(t, pipeline2)
	assert.Equal(t, model.Created, pipeline2.Status)
	assert.Equal(t, pipelineName2, pipeline2.Name)
	assert.Equal(t, testPipelineDescription, pipeline2.Description)

	// Get all the pipelines
	result, err := getSdkClient(t).StreamsService.GetPipelines(model.PipelineQueryParams{})
	require.Empty(t, err)
	require.NotEmpty(t, result)

	// Activate the second test pipeline
	ids := []string{pipeline2.ID}
	activatePipelineResponse, err := getSdkClient(t).StreamsService.ActivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline2.ID}, activatePipelineResponse["activated"])
	assert.Empty(t, activatePipelineResponse["notActivated"])

	// Get and verify the pipelines based on filters
	result, err = getSdkClient(t).StreamsService.GetPipelines(model.PipelineQueryParams{Name: &pipelineName2})
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.Equal(t, int64(1), result.Total)
	require.NotEmpty(t, result.Items)
	assert.Equal(t, pipelineName2, result.Items[0].Name)
}

// Test CreatePipeline streams endpoint
func TestIntegrationCreatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	// Create a test pipeline and verify that the pipeline was created
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), pipeline.ID, pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	require.NotEmpty(t, pipeline.Data)
	require.NotEmpty(t, pipeline.Data.Edges)
	require.Equal(t, 1, len(pipeline.Data.Edges))
	assert.NotEmpty(t, pipeline.Data.Edges[0].SourceNode)
	assert.NotEmpty(t, pipeline.Data.Edges[0].TargetNode)

	require.NotEmpty(t, pipeline.Data.Nodes)
	require.Equal(t, 2, len(pipeline.Data.Nodes))

	dataNode1, ok := pipeline.Data.Nodes[0].(map[string]interface{})
	require.True(t, ok)
	assert.NotEmpty(t, dataNode1["id"])
	assert.Equal(t, "read-splunk-firehose", dataNode1["op"])

	dataNode2, ok := pipeline.Data.Nodes[1].(map[string]interface{})
	require.True(t, ok)
	assert.NotEmpty(t, dataNode2["id"])
	assert.Equal(t, "write-splunk-index", dataNode2["op"])
	assert.Empty(t, dataNode2["attributes"])
}

// Test ActivatePipeline streams endpoint
func TestIntegrationActivatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), pipeline.ID, pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Activate the test pipeline
	ids := []string{pipeline.ID}
	activatePipelineResponse, err := getSdkClient(t).StreamsService.ActivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, activatePipelineResponse["activated"])
	assert.Empty(t, activatePipelineResponse["notActivated"])

	// Get the pipeline and verify that the pipeline status is 'activated'
	pipeline, err = getSdkClient(t).StreamsService.GetPipeline(pipeline.ID)
	require.Empty(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Activated, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)
}

// Test DeactivatePipeline streams endpoint
func TestIntegrationDeactivatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), pipeline.ID, pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Activate the newly created test pipeline
	ids := []string{pipeline.ID}
	activatePipelineResponse, err := getSdkClient(t).StreamsService.ActivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, activatePipelineResponse["activated"])

	// Deactivate the active test pipeline
	deactivatePipelineResponse, err := getSdkClient(t).StreamsService.DeactivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, deactivatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, deactivatePipelineResponse["deactivated"])
	assert.Empty(t, deactivatePipelineResponse["notDeactivated"])

	// Get the test pipeline and verify that the status is 'deactivated'
	pipeline, err = getSdkClient(t).StreamsService.GetPipeline(pipeline.ID)
	require.Empty(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, "Deactivated", pipeline.StatusMessage)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)
}

// Test ReactivatePipeline streams endpoint
func TestIntegrationReactivatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), pipeline.ID, pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Activate the newly created test pipeline
	ids := []string{pipeline.ID}
	activatePipelineResponse, err := getSdkClient(t).StreamsService.ActivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, activatePipelineResponse["activated"])

	// Deactivate the active test pipeline
	deactivatePipelineResponse, err := getSdkClient(t).StreamsService.DeactivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, deactivatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, deactivatePipelineResponse["deactivated"])
	assert.Empty(t, deactivatePipelineResponse["notDeactivated"])

	// Reactivate the deactivated test pipeline
	reactivatePipelineResponse, err := getSdkClient(t).StreamsService.ReactivatePipeline(pipeline.ID)
	require.Nil(t, err)
	require.NotEmpty(t, reactivatePipelineResponse)
	assert.Equal(t, pipeline.ID, reactivatePipelineResponse.PipelineID)
	assert.Equal(t, "activated", reactivatePipelineResponse.PipelineReactivationStatus)
}

// Test GetPipelinesStatus streams endpoint
func TestIntegrationGetPipelinesStatus(t *testing.T) {
	pipelineName1 := fmt.Sprintf("testPipeline01%d", testutils.TimeSec)
	pipelineName2 := fmt.Sprintf("testPipeline02%d", testutils.TimeSec)

	// Create two test pipelines
	pipeline1, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName1, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), pipeline1.ID, pipeline1.Name)
	require.NotEmpty(t, pipeline1)
	assert.Equal(t, model.Created, pipeline1.Status)
	assert.Equal(t, pipelineName1, pipeline1.Name)
	assert.Equal(t, testPipelineDescription, pipeline1.Description)

	pipeline2, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName2, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), pipeline2.ID, pipeline2.Name)
	require.NotEmpty(t, pipeline2)
	assert.Equal(t, model.Created, pipeline2.Status)
	assert.Equal(t, pipelineName2, pipeline2.Name)
	assert.Equal(t, testPipelineDescription, pipeline2.Description)

	// Get and verify the status of the pipelines
	result, err := getSdkClient(t).StreamsService.GetPipelineStatus(streams.PipelineStatusQueryParams{})
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.True(t, *result.Total >= 2)
	require.NotEmpty(t, result.Items)

	/*// Get and verify the status of the pipelines based on filters (query parameters) TODO (Parul): Verify specs with ingest team
	result, err = getSdkClient(t).StreamsService.GetPipelineStatus(streams.PipelineStatusQueryParams{s: p})
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.Equal(t, int64(2), result.Total)
	require.NotEmpty(t, result.Items)
	assert.Equal(t, pipelineName2, result.Items[0].PipelineId)*/
}

/*// Test MergePipelines streams endpoint TODO(Parul): Fix the failing test
func TestIntegrationMergePipelines(t *testing.T) {

	// Create two test upl pipelines
	pipeline1 := createTestUplPipeline(t)
	require.NotEmpty(t, pipeline1)

	pipeline2 := createTestUplPipeline(t)
	require.NotEmpty(t, pipeline2)
	require.NotEmpty(t, pipeline2.Edges)
	require.Equal(t, 1, len(pipeline2.Edges))
	assert.NotEmpty(t, pipeline2.Edges[0].TargetPort)
	assert.NotEmpty(t, pipeline2.Edges[0].TargetNode)

	mergeRequest := streams.PipelinesMergeRequest{
		InputTree: pipeline1,
		MainTree: pipeline2,
		TargetPort: pipeline2.Edges[0].TargetPort,
		TargetNode: pipeline2.Edges[0].TargetNode,
	}

	fmt.Println(mergeRequest)

	// Merge and verify the status of the merged UPL pipelines
	result, err := getSdkClient(t).StreamsService.MergePipelines(&mergeRequest)
	require.Nil(t, err)
	require.NotEmpty(t, result)
	require.NotEmpty(t, result.Edges)
	require.Equal(t, 3, len(pipeline2.Edges))
	require.NotEmpty(t, result.Nodes)
	require.Equal(t, 4, len(pipeline2.Nodes))
}*/

// Test UpdatePipeline streams endpoint
func TestIntegrationUpdatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), pipeline.ID, pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Update the newly created test pipeline
	updatedPipelineName := fmt.Sprintf("updated%v", pipelineName)
	updatedPipeline, err := getSdkClient(t).StreamsService.UpdatePipeline(pipeline.ID, makePipelineRequest(t, updatedPipelineName, "Updated Integration Test Pipeline"))
	require.Nil(t, err)
	require.NotEmpty(t, updatedPipeline)
	assert.Equal(t, updatedPipelineName, updatedPipeline.Name)
	assert.Equal(t, "Updated Integration Test Pipeline", updatedPipeline.Description)
	assert.Equal(t, pipeline.CurrentVersion+1, updatedPipeline.CurrentVersion)
}

// Test DeletePipeline streams endpoint
func TestIntegrationDeletePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Delete the test pipeline
	deletePipelineResponse, err := getSdkClient(t).StreamsService.DeletePipeline(pipeline.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)

	// Get the test pipeline and verify that its deleted
	pipeline, err = getSdkClient(t).StreamsService.GetPipeline(pipeline.ID)
	require.NotEmpty(t, err)
	require.Empty(t, pipeline)
}

// Test Get Input Schema streams endpoint
func TestIntegrationGetInputSchema(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)
	uplPipeline := createTestUplPipeline(t)
	require.NotEmpty(t, uplPipeline)

	nodeUID := uplPipeline.Edges[0].TargetNode
	port := uplPipeline.Edges[0].TargetPort

	// Create a test pipeline
	pipeline, err := getClient(t).StreamsService.CreatePipeline(&model.PipelineRequest{
		BypassValidation: true,
		Name:             pipelineName,
		Description:      testPipelineDescription,
		CreateUserID:     testutils.TestTenant,
		Data:             uplPipeline,
	})
	require.Nil(t, err)
	defer cleanupPipeline(getClient(t), pipeline.ID, pipelineName)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Activate the test pipeline
	ids := []string{pipeline.ID}
	activatePipelineResponse, err := getClient(t).StreamsService.ActivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, activatePipelineResponse["activated"])
	assert.Empty(t, activatePipelineResponse["notActivated"])

	//Get input Schema
	result1, err1 := getClient(t).StreamsService.GetInputSchema(&nodeUID, &port, uplPipeline)
	require.Empty(t, err1)
	require.NotEmpty(t, result1)
	assert.Equal(t, *result1.Parameters[0].Type, "field")
	assert.Equal(t, *result1.Parameters[0].FieldName, "timestamp")
	assert.Equal(t, *result1.Parameters[0].Parameters[0].Type, "long")

}

// Test Get Output Schema streams endpoint
func TestIntegrationGetOutputSchema(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)
	uplPipeline := createTestUplPipeline(t)
	require.NotEmpty(t, uplPipeline)

	nodeUID := uplPipeline.Edges[0].SourceNode
	port := uplPipeline.Edges[0].SourcePort

	// Create a test pipeline
	pipeline, err := getClient(t).StreamsService.CreatePipeline(&model.PipelineRequest{
		BypassValidation: true,
		Name:             pipelineName,
		Description:      testPipelineDescription,
		CreateUserID:     testutils.TestTenant,
		Data:             uplPipeline,
	})
	require.Nil(t, err)
	defer cleanupPipeline(getClient(t), pipeline.ID, pipelineName)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Activate the test pipeline
	ids := []string{pipeline.ID}
	activatePipelineResponse, err := getClient(t).StreamsService.ActivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, activatePipelineResponse["activated"])
	assert.Empty(t, activatePipelineResponse["notActivated"])

	//Get input Schema
	_, err1 := getClient(t).StreamsService.GetOutputSchema(&nodeUID, &port, uplPipeline)
	require.Empty(t, err1)
	//TODO(shilpa) Follow up when INGEST-8089 is investigated. Currently the output from this call could be empty sometimes
	//require.NotEmpty(t, result1)
	//assert.Equal(t, *result1.Parameters[0].Type, "field")
	//assert.Equal(t, *result1.Parameters[0].FieldName, "timestamp")
	//assert.Equal(t, *result1.Parameters[0].Parameters[0].Type, "long")
}

// Test Get Registry endpoint
func TestIntegrationGetRegistry(t *testing.T) {
	//Set local query parameter
	local := make(url.Values)
	local.Add("local", `true`)
	result, err := getClient(t).StreamsService.GetRegistry(local)
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.NotEmpty(t, *result.Functions[0].ID)
	assert.NotEmpty(t, result.Categories[0].ID)
	assert.NotEmpty(t, *result.Types[0].Type)
}

//Test Get Latest pipeline metrics endpoint
func TestIntegrationGetLatestPipelineMetrics(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	uplPipeline := createTestUplPipeline(t)
	require.NotEmpty(t, uplPipeline)

	// Create a test pipeline
	pipeline, err := getClient(t).StreamsService.CreatePipeline(&model.PipelineRequest{
		BypassValidation: true,
		Name:             pipelineName,
		Description:      testPipelineDescription,
		CreateUserID:     testutils.TestTenant,
		Data:             uplPipeline,
	})
	require.Nil(t, err)
	//defer cleanupPipeline(getClient(t), pipeline.ID, pipelineName)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Activate the test pipeline
	ids := []string{pipeline.ID}
	activatePipelineResponse, err := getClient(t).StreamsService.ActivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, activatePipelineResponse["activated"])
	assert.Empty(t, activatePipelineResponse["notActivated"])

	//Get latest pipeline metrics
	//Validation of the metrics output is not reliable since its real-time data, no guarantees if metric data will be populated at that instant of time
	//Attempt the call to get metrics 5 times and validate if there is data returned.
	cnt := 0
	for cnt < 5 {
		result1, err1 := getClient(t).StreamsService.GetLatestPipelineMetrics(pipeline.ID)
		require.Empty(t, err1)
		require.NotEmpty(t, result1)
		if len(result1.Nodes) > 0 {
			for key, element := range result1.Nodes {
				assert.NotEmpty(t, key)
				assert.NotEmpty(t, element.Metrics)
			}
		}
		time.Sleep(20 * time.Second)
		cnt++
	}
	// Delete the test pipeline
	deletePipelineResponse, err := getClient(t).StreamsService.DeletePipeline(pipeline.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)
}

//Test Latest Preview Session Metrics
func TestIntegrationGetLatestPreviewSessionMetrics(t *testing.T) {
	// Create and start a test Ge session
	response, err := getSdkClient(t).StreamsService.StartPreviewSession(createPreviewSessionStartRequest(t))
	require.Nil(t, err)
	require.NotEmpty(t, response)
	previewIDStringVal := strconv.FormatInt(*response.PreviewID, 10)
	defer cleanupPreview(t, previewIDStringVal)
	assert.NotEmpty(t, response.PipelineID)
	assert.NotEmpty(t, response.PreviewID)

	//Get latest preview session metrics
	//Validation of the metrics output is not reliable since its real-time data, no guarantees if metric data will be populated at that instant of time
	//Attempt the call to get metrics 5 times and validate if there is data returned.
	cnt := 0
	for cnt < 5 {
		result1, err1 := getClient(t).StreamsService.GetLatestPreviewSessionMetrics(previewIDStringVal)
		require.Empty(t, err1)
		require.NotEmpty(t, result1)
		if len(result1.Nodes) > 0 {
			for key, element := range result1.Nodes {
				assert.NotEmpty(t, key)
				assert.NotEmpty(t, element.Metrics)
			}
		}
		time.Sleep(20 * time.Second)
		cnt++
	}
	// Delete the test preview session
	err = getSdkClient(t).StreamsService.DeletePreviewSession(previewIDStringVal)
	require.Nil(t, err)
}

// Test Get Connectors
func TestIntegrationGetConnectors(t *testing.T) {
	response, err := getSdkClient(t).StreamsService.GetConnectors()
	require.Nil(t, err)
	require.NotEmpty(t, response)
	assert.Equal(t, *response.Connectors[0].ID, "read-splunk-firehose")
	assert.Equal(t, *response.Connectors[0].Name, "Splunk Firehose")
}

// Test Get Connections
func TestIntegrationGetConnections(t *testing.T) {
	response, err := getSdkClient(t).StreamsService.GetConnectors()
	require.Nil(t, err)
	require.NotEmpty(t, response)

	response1, err := getSdkClient(t).StreamsService.GetConnections(*response.Connectors[0].ID)
	require.Nil(t, err)
	require.NotEmpty(t, response1)
	assert.Equal(t, *response1.Connections[0].ID, "splunk-firehose:all")
	assert.Equal(t, *response1.Connections[0].Name, "Splunk Firehose")
}

// Test Validate Upl Response streams endpoint
func TestIntegrationValidateResponse(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	uplPipeline := createTestUplPipeline(t)
	require.NotEmpty(t, uplPipeline)

	// Create a test pipeline
	pipeline, err := getClient(t).StreamsService.CreatePipeline(&model.PipelineRequest{
		BypassValidation: true,
		Name:             pipelineName,
		Description:      testPipelineDescription,
		CreateUserID:     testutils.TestTenant,
		Data:             uplPipeline,
	})
	require.Nil(t, err)
	defer cleanupPipeline(getClient(t), pipeline.ID, pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	//Validate Upl response
	result1, err1 := getClient(t).StreamsService.ValidateUplResponse(uplPipeline)
	require.Empty(t, err1)
	require.NotEmpty(t, result1)
	assert.Equal(t, *result1.Success, true)

	// Delete the test pipeline
	deletePipelineResponse, err := getClient(t).StreamsService.DeletePipeline(pipeline.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)
}

// Test StartPreviewSession streams endpoint TODO: The pipelineID returned is currently equal to previewID and is incorrect and will be soon removed by the ingest team.
func TestIntegrationStartPreviewSession(t *testing.T) {
	// Create and start a test preview session
	response, err := getSdkClient(t).StreamsService.StartPreviewSession(createPreviewSessionStartRequest(t))
	require.Nil(t, err)
	require.NotEmpty(t, response)
	previewIDStringVal := strconv.FormatInt(*response.PreviewID, 10)
	defer cleanupPreview(t, previewIDStringVal)
	assert.NotEmpty(t, response.PipelineID)
	assert.NotEmpty(t, response.PreviewID)

	// Verify that the test preview session is created
	previewState, err := getSdkClient(t).StreamsService.GetPreviewSession(previewIDStringVal)
	require.Nil(t, err)
	require.NotEmpty(t, previewState)
	assert.NotEmpty(t, response.PreviewID, previewState.PreviewID)
	assert.NotEmpty(t, previewState.JobID)
}

// Test DeletePreviewSession streams endpoint
func TestIntegrationDeletePreviewSession(t *testing.T) {
	// Create and start a test preview session
	response, err := getSdkClient(t).StreamsService.StartPreviewSession(createPreviewSessionStartRequest(t))
	require.Nil(t, err)
	require.NotEmpty(t, response)
	previewIDStringVal := strconv.FormatInt(*response.PreviewID, 10)
	assert.NotEmpty(t, response.PipelineID)
	assert.NotEmpty(t, response.PreviewID)

	// Delete the test preview session
	err = getSdkClient(t).StreamsService.DeletePreviewSession(previewIDStringVal)
	require.Nil(t, err)

	// Verify that the test preview session is deleted
	_, err = getSdkClient(t).StreamsService.GetPreviewSession(previewIDStringVal)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
	assert.Equal(t, "preview-id-not-found", httpErr.Code)
}

// Test GetPreviewData streams endpoint
func TestIntegrationGetPreviewData(t *testing.T) {
	// Create and start a test preview session
	response, err := getSdkClient(t).StreamsService.StartPreviewSession(createPreviewSessionStartRequest(t))
	require.Nil(t, err)
	require.NotEmpty(t, response)
	previewIDStringVal := strconv.FormatInt(*response.PreviewID, 10)
	defer cleanupPreview(t, previewIDStringVal)
	assert.NotEmpty(t, response.PipelineID)
	assert.NotEmpty(t, response.PreviewID)

	// Verify that the preview data is generated
	previewData, err := getSdkClient(t).StreamsService.GetPreviewData(previewIDStringVal)
	require.Nil(t, err)
	require.NotEmpty(t, previewData)
	assert.NotEmpty(t, response.PreviewID, previewData.PreviewID)
}

// Test CreateTemplate streams endpoint
func TestIntegrationCreateTemplate(t *testing.T) {
	templateName := fmt.Sprintf("testTemplate%d", testutils.TimeSec)

	// Create a test template and verify that the template was created
	template, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, *template.TemplateID)
	require.NotEmpty(t, template)
	assert.Equal(t, templateName, *template.Name)
	assert.Equal(t, testTemplateDescription, *template.Description)

	require.NotEmpty(t, template.Data)
	require.NotEmpty(t, template.Data.Edges)
	require.Equal(t, 1, len(template.Data.Edges))
	assert.NotEmpty(t, template.Data.Edges[0].SourceNode)
	assert.NotEmpty(t, template.Data.Edges[0].TargetNode)

	require.NotEmpty(t, template.Data.Nodes)
	require.Equal(t, 2, len(template.Data.Nodes))

	dataNode1, ok := template.Data.Nodes[0].(map[string]interface{})
	require.True(t, ok)
	assert.NotEmpty(t, dataNode1["id"])
	assert.Equal(t, "read-splunk-firehose", dataNode1["op"])

	dataNode2, ok := template.Data.Nodes[1].(map[string]interface{})
	require.True(t, ok)
	assert.NotEmpty(t, dataNode2["id"])
	assert.Equal(t, "write-splunk-index", dataNode2["op"])
	assert.Empty(t, dataNode2["attributes"])
}

// Test GetTemplates streams endpoint
func TestIntegrationGetAllTemplates(t *testing.T) {
	templateName1 := fmt.Sprintf("testTemplate01%d", testutils.TimeSec)
	templateName2 := fmt.Sprintf("testTemplate02%d", testutils.TimeSec)

	// Create two test templates
	template1, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName1, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, *template1.TemplateID)
	require.NotEmpty(t, template1)
	assert.Equal(t, templateName1, *template1.Name)
	assert.Equal(t, testTemplateDescription, *template1.Description)

	template2, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName2, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, *template2.TemplateID)
	require.NotEmpty(t, template2)
	assert.Equal(t, templateName2, *template2.Name)
	assert.Equal(t, testTemplateDescription, *template2.Description)

	// Get all the templates
	result, err := getSdkClient(t).StreamsService.GetTemplates()
	require.Empty(t, err)
	require.NotEmpty(t, result)
}

// Test UpdateTemplate streams endpoint
func TestIntegrationUpdateTemplate(t *testing.T) {
	templateName := fmt.Sprintf("testTemplate%d", testutils.TimeSec)

	// Create a test template and verify that the template was created
	template, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, *template.TemplateID)
	require.NotEmpty(t, template)
	assert.Equal(t, templateName, *template.Name)
	assert.Equal(t, testTemplateDescription, *template.Description)

	// Update the newly created test template
	updatedTemplateName := fmt.Sprintf("updated%v", templateName)
	updatedTemplate, err := getSdkClient(t).StreamsService.UpdateTemplate(*template.TemplateID, makeTemplateRequest(t, updatedTemplateName, "Updated Integration Test Template"))
	require.Nil(t, err)
	require.NotEmpty(t, updatedTemplate)
	assert.Equal(t, updatedTemplateName, *updatedTemplate.Name)
	assert.Equal(t, "Updated Integration Test Template", *updatedTemplate.Description)
	assert.Equal(t, *template.Version+1, *updatedTemplate.Version)
}

// Test PartialUpdateTemplate streams endpoint
func TestIntegrationPartialUpdateTemplate(t *testing.T) {
	templateName := fmt.Sprintf("testTemplate%d", testutils.TimeSec)

	// Create a test template and verify that the template was created
	template, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, *template.TemplateID)
	require.NotEmpty(t, template)
	assert.Equal(t, templateName, *template.Name)
	assert.Equal(t, testTemplateDescription, *template.Description)

	// Update the newly created test template (partial update data is provided)
	updatedDescription := "Updated Integration Test Template"
	updatedTemplate, err := getSdkClient(t).StreamsService.UpdateTemplatePartially(*template.TemplateID, &streams.PartialTemplateRequest{Description: &updatedDescription})
	require.Nil(t, err)
	require.NotEmpty(t, updatedTemplate)
	assert.Equal(t, templateName, *updatedTemplate.Name)
	assert.Equal(t, "Updated Integration Test Template", *updatedTemplate.Description)
	assert.Equal(t, *template.Version+1, *updatedTemplate.Version)
}

// Test DeleteTemplate streams endpoint
func TestIntegrationDeleteTemplate(t *testing.T) {
	templateName := fmt.Sprintf("testTemplate%d", testutils.TimeSec)

	// Create a test template and verify that the template was created
	template, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName, testTemplateDescription))
	require.Nil(t, err)
	require.NotEmpty(t, template)
	assert.Equal(t, templateName, *template.Name)
	assert.Equal(t, testTemplateDescription, *template.Description)

	// Delete the test template
	err = getSdkClient(t).StreamsService.DeleteTemplate(*template.TemplateID)
	require.Nil(t, err)

	// Verify that the test template is deleted
	_, err = getSdkClient(t).StreamsService.GetTemplate(*template.TemplateID)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
	assert.Equal(t, "template-id-not-found", httpErr.Code)
}

// Test Get Groups endpoint
func TestIntegrationGetGroups(t *testing.T) {
	local := make(url.Values)
	local.Add("local", `true`)
	result, err := getClient(t).StreamsService.GetRegistry(local)
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.NotEmpty(t, *result.Functions[0].ID)
	assert.NotEmpty(t, result.Categories[0].ID)
	assert.NotEmpty(t, *result.Types[0].Type)

	cnt := 0
	temp := 0
	for cnt < len(result.Functions) {
		if *result.Functions[cnt].ID == "receive-from-ingest-rest-api" {
			temp = cnt
		}
		cnt++
	}
	applicationData, _ := result.Functions[temp].Attributes["application"].(map[string]interface{})
	groupId := applicationData["groupId"].(string)
	assert.NotEmpty(t, groupId)

	test, err := getClient(t).StreamsService.GetGroupByID(groupId)
	require.Empty(t, err)
	require.NotEmpty(t, test)
	assert.NotEmpty(t, *test.Name)
	assert.NotEmpty(t, *test.CreateUserID)
	assert.NotEmpty(t, *test.OutputType)
}

//Test the Create Expanded version of the group Endpoint
func TestIntegrationCreateExpandedGroup(t *testing.T) {
	local := make(url.Values)
	local.Add("local", `true`)
	//GetRegistry to retrieve the groupID
	result, err := getClient(t).StreamsService.GetRegistry(local)
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.NotEmpty(t, *result.Functions[0].ID)
	assert.NotEmpty(t, result.Categories[0].ID)
	assert.NotEmpty(t, *result.Types[0].Type)

	cnt := 0
	temp := 0
	for cnt < len(result.Functions) {
		if *result.Functions[cnt].ID == "receive-from-ingest-rest-api" {
			temp = cnt
		}
		cnt++
	}
	applicationData, _ := result.Functions[temp].Attributes["application"].(map[string]interface{})
	groupId := applicationData["groupId"].(string)
	assert.NotEmpty(t, groupId)

	//GetGroupID to get the Group Function ID
	result1, err := getClient(t).StreamsService.GetGroupByID(groupId)
	require.Empty(t, err)
	require.NotEmpty(t, result1)
	assert.NotEmpty(t, *result1.Name)
	assert.NotEmpty(t, *result1.CreateUserID)
	assert.NotEmpty(t, *result1.OutputType)

	functionID := *result1.Mappings[0].FunctionID

	type argumentsMap map[string]interface{}
	arguments := argumentsMap{"group_arg": "connection", "function_arg": "right"}

	result2, err := getClient(t).StreamsService.CreateExpandedGroup(groupId, arguments, functionID)
	require.Empty(t, err)
	require.NotEmpty(t, result2)
	assert.NotEmpty(t, result2.Version)
	assert.NotEmpty(t, result2.RootNode)

	dataNode, ok := result2.Nodes[0].(map[string]interface{})
	require.True(t, ok)
	assert.NotEmpty(t, dataNode["id"])

	dataNode2, ok := result2.Nodes[1].(map[string]interface{})
	require.True(t, ok)
	assert.NotEmpty(t, dataNode2["id"])

	assert.Empty(t, dataNode2["attributes"])
	assert.NotEmpty(t, result2.Edges[0].SourceNode)
	assert.NotEmpty(t, result2.Edges[0].TargetNode)
}

// makePipelineRequest is a helper function to make a PipelineRequest model
func makePipelineRequest(t *testing.T, name string, description string) *model.PipelineRequest {
	result := createTestUplPipeline(t)

	return &model.PipelineRequest{
		BypassValidation: true,
		Name:             name,
		Description:      description,
		CreateUserID:     testutils.TestTenant,
		Data:             result,
	}
}

// createTestUplPipeline is a helper function to create a test UPL JSON from a test DSL.
func createTestUplPipeline(t *testing.T) *streams.UplPipeline {
	var dsl = "events = read-splunk-firehose(); write-splunk-index(events);"
	result, err := getSdkClient(t).StreamsService.CompileDslToUpl(&model.DslCompilationRequest{Dsl: dsl})
	require.Empty(t, err)
	require.NotEmpty(t, result)

	return result
}

// createPreviewSessionStartRequest is a helper function to create a test PreviewSessionStartRequest model
func createPreviewSessionStartRequest(t *testing.T) *streams.PreviewSessionStartRequest {
	result := createTestUplPipeline(t)
	recordsLimit := int64(100)
	recordsPerPipeline := int64(2)
	sessionLifetimeMs := int64(10000)
	useNewData := false

	return &streams.PreviewSessionStartRequest{
		RecordsLimit:       &recordsLimit,
		RecordsPerPipeline: &recordsPerPipeline,
		SessionLifetimeMs:  &sessionLifetimeMs,
		Upl:                result,
		UseNewData:         &useNewData,
	}
}

// makeTemplateRequest is a helper function to make a TemplateRequest model
func makeTemplateRequest(t *testing.T, name string, description string) *streams.TemplateRequest {
	result := createTestUplPipeline(t)

	return &streams.TemplateRequest{
		Data:        result,
		Description: &description,
		Name:        &name,
	}
}

// Deletes the test pipeline
func cleanupPipeline(client *service.Client, id string, name string) {
	_, err := client.StreamsService.DeletePipeline(id)
	if err != nil {
		fmt.Printf("WARN: error deleting pipeline: name:%s, err: %s", name, err)
	}
}

// Deletes the test preview-session
func cleanupPreview(t *testing.T, id string) {
	err := getSdkClient(t).StreamsService.DeletePreviewSession(id)
	if err != nil {
		fmt.Printf("WARN: error deleting preview session: id:%s, err: %s", id, err)
	}
}

// Deletes the test template
func cleanupTemplate(t *testing.T, id string) {
	err := getSdkClient(t).StreamsService.DeleteTemplate(id)
	if err != nil {
		fmt.Printf("WARN: error deleting template: id:%s, err: %s", id, err)
	}
}
