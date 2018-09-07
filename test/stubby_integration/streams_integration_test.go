package stubbyintegration

import (
	"encoding/json"
	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var testPipelineID = "TEST_PIPELINE_01"
var testPipelineName = "TEST_PIPELINE"
var testNodeID1 = "TEST_NODE_01"
var testNodeID2 = "TEST_NODE_02"
var testNodeID3 = "TEST_NODE_03"
var testNodeID4 = "TEST_NODE_04"
var testRootNodeID = "ROOT_NODE_01"
var testJobID = "JOB_ID_01"

// Stubby test for GetPipelines() streams service endpoint
func TestCompileDslToUpl(t *testing.T) {

	// Create a test UPL JSON from a test DSL and verify the data
	result, err := getClient(t).StreamsService.CompileDslToUpl(&model.DslCompilationRequest{Dsl: "kafka-brokers=\"localhost:9092\";input-topic = \"intopic\";output-topic-1 = \"output-topic-1\";events = deserialize-events(read-kafka(kafka-brokers, input-topic, {}));write-kafka(serialize-events(events, output-topic-1), kafka-brokers, {});"})
	require.Empty(t, err)
	require.NotEmpty(t, result)

	require.NotEmpty(t, result.Nodes)
	assert.Equal(t, 4, len(result.Nodes))

	dataNode1 := result.Nodes[0].(map[string]interface{})
	assert.NotEmpty(t, dataNode1["id"])
	assert.Equal(t, "read-kafka", dataNode1["op"])
	assert.Equal(t, "localhost:9092", dataNode1["brokers"])
	assert.Equal(t, "intopic", dataNode1["topic"])

	dataNode2 := result.Nodes[1].(map[string]interface{})
	assert.NotEmpty(t, dataNode2["id"])
	assert.Equal(t, "deserialize-events", dataNode2["op"])
	assert.Empty(t, dataNode2["attributes"])

	dataNode3 := result.Nodes[2].(map[string]interface{})
	assert.NotEmpty(t, dataNode3["id"])
	assert.Equal(t, "serialize-events", dataNode3["op"])
	assert.Equal(t, "output-topic-1", dataNode3["topic"])

	dataNode4 := result.Nodes[3].(map[string]interface{})
	assert.NotEmpty(t, dataNode4["id"])
	assert.Equal(t, "write-kafka", dataNode4["op"])
	assert.Equal(t, "localhost:9092", dataNode4["brokers"])
	assert.Empty(t, dataNode4["producer-properties"])

	require.NotEmpty(t, result.Edges)
	assert.Equal(t, 3, len(result.Edges))
	assert.NotEmpty(t, result.Edges[0].SourceNode)
	assert.NotEmpty(t, result.Edges[0].TargetNode)
	assert.NotEmpty(t, result.Edges[1].SourceNode)
	assert.NotEmpty(t, result.Edges[1].TargetNode)
	assert.NotEmpty(t, result.Edges[2].SourceNode)
	assert.NotEmpty(t, result.Edges[2].TargetNode)

	assert.NotEmpty(t, result.RootNode)
}

// Stubby test for GetPipelines() streams service endpoint
func TestGetPipelines(t *testing.T) {
	result, err := getClient(t).StreamsService.GetPipelines(nil)
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, 1, len(result.Items))
	assert.Equal(t, int64(1), result.Total)
	assert.Equal(t, testPipelineID, result.Items[0].ID)
	assert.Equal(t, testutils.TestTenantID, result.Items[0].CreateUserID)
}

// Stubby test for CreatePipeline() streams service endpoint
func TestCreatePipeline(t *testing.T) {

	result, err := getClient(t).StreamsService.CreatePipeline(CreatePipelineRequest(t, "TEST_PIPELINE", "Stubby Test Pipeline"))
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, testPipelineID, result.ID)
	assert.Equal(t, testPipelineName, result.Name)
	assert.Equal(t, testutils.TestTenantID, result.CreateUserID)

	require.NotEmpty(t, result.Data)
	require.NotEmpty(t, result.Data.Nodes)
	assert.Equal(t, 4, len(result.Data.Nodes))

	dataNode1 := result.Data.Nodes[0].(map[string]interface{})
	assert.Equal(t, testNodeID1, dataNode1["id"])
	assert.Equal(t, "read-kafka", dataNode1["op"])
	assert.Equal(t, "localhost:9092", dataNode1["brokers"])
	assert.Equal(t, "intopic", dataNode1["topic"])

	dataNode2 := result.Data.Nodes[1].(map[string]interface{})
	assert.Equal(t, testNodeID2, dataNode2["id"])
	assert.Equal(t, "deserialize-events", dataNode2["op"])
	assert.Empty(t, dataNode2["attributes"])

	dataNode3 := result.Data.Nodes[2].(map[string]interface{})
	assert.Equal(t, testNodeID3, dataNode3["id"])
	assert.Equal(t, "serialize-events", dataNode3["op"])
	assert.Equal(t, "output-topic-1", dataNode3["topic"])

	dataNode4 := result.Data.Nodes[3].(map[string]interface{})
	assert.Equal(t, testNodeID4, dataNode4["id"])
	assert.Equal(t, "write-kafka", dataNode4["op"])
	assert.Equal(t, "localhost:9092", dataNode4["brokers"])
	assert.Empty(t, dataNode4["producer-properties"])

	require.NotEmpty(t, result.Data.Edges)
	assert.Equal(t, 3, len(result.Data.Edges))
	assert.Equal(t, testNodeID1, result.Data.Edges[0].SourceNode)
	assert.Equal(t, testNodeID2, result.Data.Edges[0].TargetNode)
	assert.Equal(t, testNodeID2, result.Data.Edges[1].SourceNode)
	assert.Equal(t, testNodeID3, result.Data.Edges[1].TargetNode)
	assert.Equal(t, testNodeID3, result.Data.Edges[2].SourceNode)
	assert.Equal(t, testNodeID4, result.Data.Edges[2].TargetNode)

	assert.Equal(t, []string{testNodeID4}, result.Data.RootNode)
}

// Stubby test for UpdatePipeline() streams service endpoint
func TestUpdatePipeline(t *testing.T) {

	result, err := getClient(t).StreamsService.UpdatePipeline(testPipelineID, CreatePipelineRequest(t, "UPDATED_TEST_PIPELINE", "Updated Stubby Test Pipeline"))
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, testPipelineID, result.ID)
	assert.Equal(t, "UPDATED_TEST_PIPELINE", result.Name)
	assert.Equal(t, "Updated Stubby Test Pipeline", result.Description)
	assert.Equal(t, testutils.TestTenantID, result.CreateUserID)

	require.NotEmpty(t, result.Data)
	require.NotEmpty(t, result.Data.Nodes)
	assert.Equal(t, 4, len(result.Data.Nodes))

	dataNode1 := result.Data.Nodes[0].(map[string]interface{})
	assert.Equal(t, testNodeID1, dataNode1["id"])
	assert.Equal(t, "read-kafka", dataNode1["op"])
	assert.Equal(t, "localhost:9092", dataNode1["brokers"])
	assert.Equal(t, "intopic", dataNode1["topic"])

	dataNode2 := result.Data.Nodes[1].(map[string]interface{})
	assert.Equal(t, testNodeID2, dataNode2["id"])
	assert.Equal(t, "deserialize-events", dataNode2["op"])
	assert.Empty(t, dataNode2["attributes"])

	dataNode3 := result.Data.Nodes[2].(map[string]interface{})
	assert.Equal(t, testNodeID3, dataNode3["id"])
	assert.Equal(t, "serialize-events", dataNode3["op"])
	assert.Equal(t, "output-topic-1", dataNode3["topic"])

	dataNode4 := result.Data.Nodes[3].(map[string]interface{})
	assert.Equal(t, testNodeID4, dataNode4["id"])
	assert.Equal(t, "write-kafka", dataNode4["op"])
	assert.Equal(t, "localhost:9092", dataNode4["brokers"])
	assert.Empty(t, dataNode4["producer-properties"])

	require.NotEmpty(t, result.Data.Edges)
	assert.Equal(t, 3, len(result.Data.Edges))
	assert.Equal(t, testNodeID1, result.Data.Edges[0].SourceNode)
	assert.Equal(t, testNodeID2, result.Data.Edges[0].TargetNode)
	assert.Equal(t, testNodeID2, result.Data.Edges[1].SourceNode)
	assert.Equal(t, testNodeID3, result.Data.Edges[1].TargetNode)
	assert.Equal(t, testNodeID3, result.Data.Edges[2].SourceNode)
	assert.Equal(t, testNodeID4, result.Data.Edges[2].TargetNode)

	assert.Equal(t, []string{testNodeID4}, result.Data.RootNode)
}

// Stubby test for ActivatePipeline() streams service endpoint
func TestActivatePipeline(t *testing.T) {
	ids := []string{testPipelineID}
	result, err := getClient(t).StreamsService.ActivatePipeline(ids)
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, []string{testPipelineID}, result["activated"])
}

// Stubby test for DeactivatePipeline() streams service endpoint
func TestDeactivatePipeline(t *testing.T) {
	ids := []string{testPipelineID}
	result, err := getClient(t).StreamsService.DeactivatePipeline(ids)
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, []string{testPipelineID}, result["deactivated"])
}

// Stubby test for GetPipeline() streams service endpoint
func TestGetPipeline(t *testing.T) {
	result, err := getClient(t).StreamsService.GetPipeline(testPipelineID)
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, testPipelineID, result.ID)
	assert.Equal(t, testutils.TestTenantID, result.CreateUserID)
	assert.Equal(t, 2, len(result.Data.Nodes))
	assert.Equal(t, testNodeID1, result.Data.Nodes[0].(map[string]interface{})["id"])
	assert.Equal(t, testNodeID2, result.Data.Nodes[1].(map[string]interface{})["id"])

	assert.Equal(t, 2, len(result.Data.Edges))
	assert.Equal(t, testNodeID1, result.Data.Edges[0].SourceNode)
	assert.Equal(t, testRootNodeID, result.Data.Edges[0].TargetNode)
	assert.Equal(t, testNodeID2, result.Data.Edges[1].SourceNode)
	assert.Equal(t, testRootNodeID, result.Data.Edges[1].TargetNode)

	assert.Equal(t, []string{testRootNodeID}, result.Data.RootNode)
	assert.Equal(t, testJobID, result.JobID)
}

// Stubby test for DeletePipeline() streams service endpoint
func TestDeletePipeline(t *testing.T) {
	result, err := getClient(t).StreamsService.DeletePipeline(testPipelineID)
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, true, result.CouldDeactivate)
	assert.Equal(t, true, result.Running)
}

// Creates a pipeline request
func CreatePipelineRequest(t *testing.T, name string, description string) *model.PipelineRequest {
	uplPipelineData := []byte(`{
                 "nodes": [
                   {
                     "op": "read-kafka",
                     "id": "TEST_NODE_01",
                     "attributes": null,
                     "brokers": "localhost:9092",
                     "consumer-properties": {},
                     "topic": "intopic"
                   },
                   {
                     "op": "deserialize-events",
                     "id": "TEST_NODE_02",
                     "attributes": null
                   },
                   {
                     "op": "serialize-events",
                     "id": "TEST_NODE_03",
                     "attributes": null,
                     "topic": "output-topic-1"
                   },
                   {
                     "op": "write-kafka",
                     "id": "TEST_NODE_04",
                     "attributes": null,
                     "producer-properties": {},
                     "brokers": "localhost:9092"
                   }
                 ],
                 "edges": [
                   {
                     "attributes": null,
                     "sourceNode": "TEST_NODE_01",
                     "sourcePort": "output",
                     "targetNode": "TEST_NODE_02",
                     "targetPort": "input"
                   },
                   {
                     "attributes": null,
                     "sourceNode": "TEST_NODE_02",
                     "sourcePort": "output",
                     "targetNode": "TEST_NODE_03",
                     "targetPort": "input"
                   },
                   {
                     "attributes": null,
                     "sourceNode": "TEST_NODE_03",
                     "sourcePort": "output",
                     "targetNode": "TEST_NODE_04",
                     "targetPort": "input"
                   }
                 ],
                 "root-node": [
                   "TEST_NODE_04"
                 ],
                 "version": 3
               }`)

	var uplPipeline model.UplPipeline
	err := json.Unmarshal(uplPipelineData, &uplPipeline)
	if err != nil {
		t.Errorf("Unable to unmarshal upl pipeline json object, the error message is %v", err)
	}
	return &model.PipelineRequest{
		BypassValidation: true,
		Name:             name,
		Description:      description,
		CreateUserID:     testutils.TestTenantID,
		Data:             &uplPipeline,
	}
}
