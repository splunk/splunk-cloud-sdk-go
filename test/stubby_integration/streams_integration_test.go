package stubbyintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/testutils"
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
var testInputPort = "input"
var testOutputPort = "output"

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

	result, err := getClient(t).StreamsService.CreatePipeline(CreatePipelineRequest("TEST_PIPELINE", "Stubby Test Pipeline"))
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, testPipelineID, result.ID)
	assert.Equal(t, testPipelineName, result.Name)
	assert.Equal(t, testutils.TestTenantID, result.CreateUserID)

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

	result, err := getClient(t).StreamsService.UpdatePipeline(testPipelineID, CreatePipelineRequest("UPDATED_TEST_PIPELINE", "Updated Stubby Test Pipeline"))
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, testPipelineID, result.ID)
	assert.Equal(t, "UPDATED_TEST_PIPELINE", result.Name)
	assert.Equal(t, "Updated Stubby Test Pipeline", result.Description)
	assert.Equal(t, testutils.TestTenantID, result.CreateUserID)

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
	result, err := getClient(t).StreamsService.ActivatePipeline(model.ActivatePipelineRequest{IDs: ids})
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, []string{testPipelineID}, result["activated"])
}

// Stubby test for DeactivatePipeline() streams service endpoint
func TestDeactivatePipeline(t *testing.T) {
	ids := []string{testPipelineID}
	result, err := getClient(t).StreamsService.DeactivatePipeline(model.ActivatePipelineRequest{IDs: ids})
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
func CreatePipelineRequest(name string, description string) model.PipelineRequest {
	/*var nodeAttributesJson1 = `"attributes": {
	          "inputs": [],
	          "outputs": [
	            "output"
	          ],
	          "isValid": true,
	          "name": "Kafka Source",
	          "dsl": {},
	          "isSource": true,
	          "isSink": false,
	          "arguments": {
	            "brokers": {
	              "type": "string"
	            },
	            "topic": {
	              "type": "string"
	            },
	            "consumer-properties": {
	              "type": "Map",
	              "element-type": {
	                "key-type": "string",
	                "value-type": "string"
	              }
	            }
	          },
	          "formData": {
	            "brokers": "localhost:8884",
	            "topic": "intopic",
	            "consumer-properties": {}
	          },
	          "isActive": true
	        }`

		var nodeAttributes1 map[string]interface{}
		err := json.Unmarshal([]byte(nodeAttributesJson1), &nodeAttributes1)

		var nodeAttributesJson2 = `{
	          "inputs": [
	            "input"
	          ],
	          "outputs": [
	            "output"
	          ],
	          "isValid": true,
	          "name": "Parse Events",
	          "dsl": {},
	          "isSource": false,
	          "isSink": false,
	          "arguments": {},
	          "formData": {},
	          "isActive": false
	        }`

		var nodeAttributes2 map[string]interface{}
		err = json.Unmarshal([]byte(nodeAttributesJson2), &nodeAttributes2)

		var nodeAttributesJson3 = `{
	          "inputs": [
	            "input"
	          ],
	          "outputs": [
	            "output"
	          ],
	          "name": "To Events",
	          "isValid": true,
	          "dsl": {},
	          "isSource": false,
	          "isSink": false,
	          "arguments": {
	            "topic": {
	              "type": "string"
	            }
	          },
	          "formData": {
	            "topic": "output-topic-1"
	          },
	          "isActive": false
	        }`

		var nodeAttributes3 map[string]interface{}
		err = json.Unmarshal([]byte(nodeAttributesJson3), &nodeAttributes3)

		var nodeAttributesJson4 = `{
	          "inputs": [
	            "input"
	          ],
	          "outputs": [],
	          "isValid": true,
	          "name": "Kafka Write",
	          "dsl": {},
	          "isSource": false,
	          "isSink": true,
	          "arguments": {
	            "brokers": {
	              "type": "string"
	            },
	            "producer-properties": {
	              "type": "Map",
	              "element-type": {
	                "key-type": "string",
	                "value-type": "string"
	              }
	            }
	          },
	          "formData": {
	            "brokers": "localhost:8884",
	            "producer-properties": {}
	          },
	          "isActive": false
	        }`

		var nodeAttributes4 map[string]interface{}
		err = json.Unmarshal([]byte(nodeAttributesJson4), &nodeAttributes4)
		fmt.Println(err)*/
	var producerProperties struct{}
	uplNode1 := model.UplNodeCommon{Op: "read-kafka", ID: testNodeID1}
	node1 := model.KafkaReader{UplNodeCommon: uplNode1, Brokers: "localhost:9092", Topic: "intopic"}

	node2 := model.UplNodeCommon{Op: "deserialize-events", ID: testNodeID2}

	uplNode3 := model.UplNodeCommon{Op: "serialize-events", ID: testNodeID3}
	node3 := model.SerializeEvents{UplNodeCommon: uplNode3, Topic: "output-topic-1"}

	uplNode4 := model.UplNodeCommon{Op: "write-kafka", ID: testNodeID4}
	node4 := model.KafkaWriter{UplNodeCommon: uplNode4, Brokers: "localhost:9092", ProducerProperties: producerProperties}

	nodes := []model.UplNode{node1, node2, node3, node4}
	/*nodes := []model.KafkaReader{{UplNode: model.UplNode{Op: "read-kafka", ID: testNodeID1}, Brokers: "localhost:8884", Topic: "intopic"},
	{Op: "deserialize-events", ID: testNodeID2},
	{Op: "serialize-events", ID: testNodeID3},
	{Op: "write-kafka", ID: testNodeID4}}*/

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
