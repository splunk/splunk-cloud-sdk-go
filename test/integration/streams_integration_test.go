package integration

import (
	"fmt"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/service"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test variables
var testPipelineDescription = "integration test pipeline"

// Test GetPipelines streams endpoint
func TestIntegrationGetAllPipelines(t *testing.T) {
	pipelineName1 := fmt.Sprintf("testPipeline01%d", testutils.TimeSec)
	pipelineName2 := fmt.Sprintf("testPipeline02%d", testutils.TimeSec)

	// Create two test pipelines
	pipeline1, err := getClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName1, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getClient(t), pipeline1.ID, pipeline1.Name)
	require.NotEmpty(t, pipeline1)
	assert.Equal(t, model.Created, pipeline1.Status)
	assert.Equal(t, pipelineName1, pipeline1.Name)
	assert.Equal(t, testPipelineDescription, pipeline1.Description)

	pipeline2, err := getClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName2, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getClient(t), pipeline2.ID, pipeline2.Name)
	require.NotEmpty(t, pipeline2)
	assert.Equal(t, model.Created, pipeline2.Status)
	assert.Equal(t, pipelineName2, pipeline2.Name)
	assert.Equal(t, testPipelineDescription, pipeline2.Description)

	// Get all the pipelines
	result, err := getClient(t).StreamsService.GetPipelines(model.PipelineQueryParams{})
	require.Empty(t, err)
	require.NotEmpty(t, result)

	// Activate the second test pipeline
	ids := []string{pipeline2.ID}
	activatePipelineResponse, err := getClient(t).StreamsService.ActivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline2.ID}, activatePipelineResponse["activated"])
	assert.Empty(t, activatePipelineResponse["notActivated"])

	// Get and verify the pipelines based on filters
	result, err = getClient(t).StreamsService.GetPipelines(model.PipelineQueryParams{Name: &pipelineName2})
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
	pipeline, err := getClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getClient(t), pipeline.ID, pipeline.Name)
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
	pipeline, err := getClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getClient(t), pipeline.ID, pipeline.Name)
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

	// Get the pipeline and verify that the pipeline status is 'activated'
	pipeline, err = getClient(t).StreamsService.GetPipeline(pipeline.ID)
	require.Empty(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Activated, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)
}

// TODO (Parul): Known bug - BLAM-4340, until the fix is ready, setting a workaround field - skipSavePoint (=true)
// Test DeactivatePipeline streams endpoint
func TestIntegrationDeactivatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getClient(t), pipeline.ID, pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Activate the newly created test pipeline
	ids := []string{pipeline.ID}
	activatePipelineResponse, err := getClient(t).StreamsService.ActivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, activatePipelineResponse["activated"])

	// Deactivate the active test pipeline
	deactivatePipelineResponse, err := getClient(t).StreamsService.DeactivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, deactivatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, deactivatePipelineResponse["deactivated"])
	assert.Empty(t, deactivatePipelineResponse["notDeactivated"])

	// Get the test pipeline and verify that the status is 'deactivated'
	pipeline, err = getClient(t).StreamsService.GetPipeline(pipeline.ID)
	require.Empty(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, "Deactivated", pipeline.StatusMessage)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)
}

// Test UpdatePipeline streams endpoint
func TestIntegrationUpdatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getClient(t), pipeline.ID, pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Update the newly created test pipeline
	updatedPipelineName := fmt.Sprintf("updated%v", pipelineName)
	updatedPipeline, err := getClient(t).StreamsService.UpdatePipeline(pipeline.ID, makePipelineRequest(t, updatedPipelineName, "Updated Integration Test Pipeline"))
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
	pipeline, err := getClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Delete the test pipeline
	deletePipelineResponse, err := getClient(t).StreamsService.DeletePipeline(pipeline.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)

	// Get the test pipeline and verify that its deleted
	pipeline, err = getClient(t).StreamsService.GetPipeline(pipeline.ID)
	require.NotEmpty(t, err)
	require.Empty(t, pipeline)
}

// makePipelineRequest is a helper function to make a PipelineRequest model
func makePipelineRequest(t *testing.T, name string, description string) *model.PipelineRequest {
	// Create a test UPL JSON from a test DSL
	var dsl = "events = read-splunk-firehose(); write-splunk-index(events);"
	result, err := getClient(t).StreamsService.CompileDslToUpl(&model.DslCompilationRequest{Dsl: dsl})
	require.Empty(t, err)
	require.NotEmpty(t, result)

	return &model.PipelineRequest{
		BypassValidation: true,
		Name:             name,
		Description:      description,
		CreateUserID:     testutils.TestTenant,
		Data:             result,
	}
}

// Deletes the test pipeline
func cleanupPipeline(client *service.Client, id string, name string) {
	_, err := client.StreamsService.DeletePipeline(id)
	if err != nil {
		fmt.Printf("WARN: error deleting pipeline: name:%s, err: %s", name, err)
	}
}
