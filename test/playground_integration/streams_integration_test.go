package playgroundintegration

import (
	"fmt"
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// Test variables
var testPipelineDescription = "integration test pipeline"
var testNodeID1 = "TEST_NODE_01"
var testNodeID2 = "TEST_NODE_02"
var testNodeID3 = "TEST_NODE_03"
var testNodeID4 = "TEST_NODE_04"
var testInputPort = "input"
var testOutputPort = "output"

// Test CreatePipeline streams endpoint
func TestIntegrationGetAllPipelines(t *testing.T) {
	pipelineName1 := fmt.Sprintf("testPipeline01%d", timeSec)
	pipelineName2 := fmt.Sprintf("testPipeline02%d", timeSec)

	// Create two test pipelines
	pipeline1, err := getClient(t).StreamsService.CreatePipeline(CreatePipelineRequest(pipelineName1, testPipelineDescription))
	require.Nil(t, err)
	require.NotEmpty(t, pipeline1)
	assert.Equal(t, model.Created, pipeline1.Status)
	assert.Equal(t, pipelineName1, pipeline1.Name)
	assert.Equal(t, testPipelineDescription, pipeline1.Description)

	pipeline2, err := getClient(t).StreamsService.CreatePipeline(CreatePipelineRequest(pipelineName2, testPipelineDescription))
	require.Nil(t, err)
	require.NotEmpty(t, pipeline2)
	assert.Equal(t, model.Created, pipeline2.Status)
	assert.Equal(t, pipelineName2, pipeline2.Name)
	assert.Equal(t, testPipelineDescription, pipeline2.Description)

	// Retrieve all the pipelines
	result, err := getClient(t).StreamsService.GetPipelines(nil)
	require.Empty(t, err)
	require.NotEmpty(t, result)
	// TODO Check count

	// Delete the test pipelines
	deletePipelineResponse, err := getClient(t).StreamsService.DeletePipeline(pipeline1.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)

	deletePipelineResponse, err = getClient(t).StreamsService.DeletePipeline(pipeline2.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)
}

// Test CreatePipeline streams endpoint
func TestIntegrationCreatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", timeSec)

	// Create a test pipeline
	pipeline, err := getClient(t).StreamsService.CreatePipeline(CreatePipelineRequest(pipelineName, testPipelineDescription))
	require.Nil(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	assert.Equal(t, 3, len(pipeline.Data.Edges))
	assert.Equal(t, testNodeID1, pipeline.Data.Edges[0].SourceNode)
	assert.Equal(t, testNodeID2, pipeline.Data.Edges[0].TargetNode)
	assert.Equal(t, testNodeID2, pipeline.Data.Edges[1].SourceNode)
	assert.Equal(t, testNodeID3, pipeline.Data.Edges[1].TargetNode)
	assert.Equal(t, testNodeID3, pipeline.Data.Edges[2].SourceNode)
	assert.Equal(t, testNodeID4, pipeline.Data.Edges[2].TargetNode)

	assert.Equal(t, 4, len(pipeline.Data.Nodes))

	dataNode1 := pipeline.Data.Nodes[0].(map[string]interface{})
	assert.Equal(t, testNodeID1, dataNode1["id"])
	assert.Equal(t, "read-kafka", dataNode1["op"])
	assert.Equal(t, "localhost:9092", dataNode1["brokers"])
	assert.Equal(t, "intopic", dataNode1["topic"])

	dataNode2 := pipeline.Data.Nodes[1].(map[string]interface{})
	assert.Equal(t, testNodeID2, dataNode2["id"])
	assert.Equal(t, "deserialize-events", dataNode2["op"])
	assert.Empty(t, dataNode2["attributes"])

	dataNode3 := pipeline.Data.Nodes[2].(map[string]interface{})
	assert.Equal(t, testNodeID3, dataNode3["id"])
	assert.Equal(t, "serialize-events", dataNode3["op"])
	assert.Equal(t, "output-topic-1", dataNode3["topic"])

	dataNode4 := pipeline.Data.Nodes[3].(map[string]interface{})
	assert.Equal(t, testNodeID4, dataNode4["id"])
	assert.Equal(t, "write-kafka", dataNode4["op"])
	assert.Equal(t, "localhost:9092", dataNode4["brokers"])
	assert.Empty(t, dataNode4["producer-properties"])

	// Delete the test pipeline
	deletePipelineResponse, err := getClient(t).StreamsService.DeletePipeline(pipeline.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)
}

// Test ActivatePipeline streams endpoint
func TestIntegrationActivatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", timeSec)

	// Create a test pipeline
	pipeline, err := getClient(t).StreamsService.CreatePipeline(CreatePipelineRequest(pipelineName, testPipelineDescription))
	require.Nil(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Activate Pipeline
	ids := []string{pipeline.ID}
	activatePipelineResponse, err := getClient(t).StreamsService.ActivatePipeline(model.ActivatePipelineRequest{IDs: ids})
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, activatePipelineResponse["activated"])
	assert.Empty(t, activatePipelineResponse["notActivated"])

	// Retrieve the pipeline and verify that the status is 'activated'
	pipeline, err = getClient(t).StreamsService.GetPipeline(pipeline.ID)
	require.Empty(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Activated, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Delete the test pipeline
	deletePipelineResponse, err := getClient(t).StreamsService.DeletePipeline(pipeline.ID)
	fmt.Println(deletePipelineResponse)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)
}

// Test DeactivatePipeline streams endpoint
func TestIntegrationDeactivatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", timeSec)

	// Create test pipeline
	pipeline, err := getClient(t).StreamsService.CreatePipeline(CreatePipelineRequest(pipelineName, testPipelineDescription))
	require.Nil(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, model.Created, pipeline.Status)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Activate the newly created test pipeline
	ids := []string{pipeline.ID}
	activatePipelineResponse, err := getClient(t).StreamsService.ActivatePipeline(model.ActivatePipelineRequest{IDs: ids})
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, activatePipelineResponse["activated"])

	// Deactivate the active test pipeline
	deactivatePipelineResponse, err := getClient(t).StreamsService.DeactivatePipeline(model.ActivatePipelineRequest{IDs: ids})
	require.Nil(t, err)
	require.NotEmpty(t, deactivatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, deactivatePipelineResponse["deactivated"])
	assert.Empty(t, deactivatePipelineResponse["notDeactivated"])

	// Retrieve the pipeline and verify that the status is 'deactivated'
	pipeline, err = getClient(t).StreamsService.GetPipeline(pipeline.ID)
	require.Empty(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, "Deactivated", pipeline.StatusMessage)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Delete the test pipeline
	deletePipelineResponse, err := getClient(t).StreamsService.DeletePipeline(pipeline.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)
}

// Test UpdatePipeline streams endpoint
func TestIntegrationUpdatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", timeSec)

	// Create a test pipeline
	pipeline, err := getClient(t).StreamsService.CreatePipeline(CreatePipelineRequest(pipelineName, testPipelineDescription))
	require.Nil(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, pipelineName, pipeline.Name)
	assert.Equal(t, testPipelineDescription, pipeline.Description)

	// Update the newly created test pipeline
	updatedPipelineName := fmt.Sprintf("updated%v", pipelineName)
	updatedPipeline, err := getClient(t).StreamsService.UpdatePipeline(pipeline.ID, CreatePipelineRequest(updatedPipelineName, "Updated Integration Test Pipeline"))
	require.Nil(t, err)
	require.NotEmpty(t, updatedPipeline)
	assert.Equal(t, updatedPipelineName, updatedPipeline.Name)
	assert.Equal(t, "Updated Integration Test Pipeline", updatedPipeline.Description)
	assert.Equal(t, pipeline.CurrentVersion+1, updatedPipeline.CurrentVersion)

	// Delete the test pipeline
	deletePipelineResponse, err := getClient(t).StreamsService.DeletePipeline(pipeline.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)
}

// Creates a pipeline request
func CreatePipelineRequest(name string, description string) model.PipelineRequest {
	var producerProperties struct{}

	uplNode1 := model.UplNodeCommon{Op: "read-kafka", ID: testNodeID1}
	node1 := model.KafkaReader{UplNodeCommon: uplNode1, Brokers: "localhost:9092", Topic: "intopic"}

	node2 := model.UplNodeCommon{Op: "deserialize-events", ID: testNodeID2}

	uplNode3 := model.UplNodeCommon{Op: "serialize-events", ID: testNodeID3}
	node3 := model.SerializeEvents{UplNodeCommon: uplNode3, Topic: "output-topic-1"}

	uplNode4 := model.UplNodeCommon{Op: "write-kafka", ID: testNodeID4}
	node4 := model.KafkaWriter{UplNodeCommon: uplNode4, Brokers: "localhost:9092", ProducerProperties: producerProperties}

	nodes := []model.UplNode{node1, node2, node3, node4}

	edges := []model.UplEdge{{SourceNode: testNodeID1, SourcePort: testOutputPort, TargetNode: testNodeID2, TargetPort: testInputPort},
		{SourceNode: testNodeID2, SourcePort: testOutputPort, TargetNode: testNodeID3, TargetPort: testInputPort},
		{SourceNode: testNodeID3, SourcePort: testOutputPort, TargetNode: testNodeID4, TargetPort: testInputPort}}

	data := model.UplPipeline{Version: 3, RootNode: []string{"TEST_NODE_04"}, Nodes: nodes, Edges: edges}

	return model.PipelineRequest{
		BypassValidation: true,
		Name:             name,
		Description:      description,
		CreateUserID:     testutils.TestTenantID,
		Data:             data,
	}
}

// TODO: Add error scenarios and tests for getPipeline queries after one code review
