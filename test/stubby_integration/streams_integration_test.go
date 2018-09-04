package stubbyintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var testPipelineID = "TEST_PIPELINE_01"
var testNodeID1 = "TEST_NODE_01"
var testNodeID2 = "TEST_NODE_02"
var testRootNodeID = "ROOT_NODE_01"
var testJobID = "JOB_ID_01"

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

// Stubby test for ActivatePipeline() streams service endpoint
func TestActivatePipeline(t *testing.T) {
	ids := []string{testPipelineID}
	ids[0] = testPipelineID
	result, err := getClient(t).StreamsService.ActivatePipeline(model.ActivatePipelineRequest{IDs: ids})
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, []string{testPipelineID}, result["activated"])
}

// Stubby test for DeactivatePipeline() streams service endpoint
func TestDeactivatePipeline(t *testing.T) {
	ids := []string{testPipelineID}
	ids[0] = testPipelineID
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
	assert.Equal(t, testNodeID1, result.Data.Nodes[0].ID)
	assert.Equal(t, testNodeID2, result.Data.Nodes[1].ID)

	assert.Equal(t, 2, len(result.Data.Edges))
	assert.Equal(t, testNodeID1, result.Data.Edges[0].SourceNode)
	assert.Equal(t, testRootNodeID, result.Data.Edges[0].TargetNode)

	assert.Equal(t, 2, len(result.Data.Edges))
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