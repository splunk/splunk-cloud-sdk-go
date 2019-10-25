package integration

import (
	"fmt"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/v2/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/v2/services/streams"
	testutils "github.com/splunk/splunk-cloud-sdk-go/v2/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/v2/util"
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
	defer cleanupPipeline(getSdkClient(t), *pipeline1.Id, *pipeline1.Name)
	require.NotEmpty(t, pipeline1)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline1.Status)
	assert.Equal(t, pipelineName1, *pipeline1.Name)
	assert.Equal(t, testPipelineDescription, *pipeline1.Description)

	pipeline2, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName2, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline2.Id, *pipeline2.Name)
	require.NotEmpty(t, pipeline2)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline2.Status)
	assert.Equal(t, pipelineName2, *pipeline2.Name)
	assert.Equal(t, testPipelineDescription, *pipeline2.Description)

	// Get all the pipelines
	result, err := getSdkClient(t).StreamsService.ListPipelines(nil)
	require.Empty(t, err)
	require.NotEmpty(t, result)

	// Activate the second test pipeline
	boolvar := true
	activatePipelineRequest := streams.ActivatePipelineRequest{
		SkipRestoreState: &boolvar,
	}

	activatePipelineResponse, err := getSdkClient(t).StreamsService.ActivatePipeline(*pipeline2.Id, activatePipelineRequest)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, pipeline2.Id, activatePipelineResponse.Activated)
	assert.Empty(t, activatePipelineResponse.Deactivated)

	// Get and verify the pipelines based on filters
	query := streams.ListPipelinesQueryParams{}.SetName(pipelineName2)
	result, err = getSdkClient(t).StreamsService.ListPipelines(&query)
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.Equal(t, int64(1), *result.Total)
	require.NotEmpty(t, result.Items)
	assert.Equal(t, pipelineName2, *result.Items[0].Name)
}

// Test CreatePipeline streams endpoint
func TestIntegrationCreatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipelinea%d", testutils.TimeSec)

	// Create a test pipeline and verify that the pipeline was created
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline.Id, *pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline.Status)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)

	require.NotEmpty(t, pipeline.Data)
	require.NotEmpty(t, pipeline.Data.Edges)
	require.Equal(t, 1, len(pipeline.Data.Edges))
	assert.NotEmpty(t, pipeline.Data.Edges[0].SourceNode)
	assert.NotEmpty(t, pipeline.Data.Edges[0].TargetNode)

	require.NotEmpty(t, pipeline.Data.Nodes)
	//require.Equal(t, 2, len(pipeline.Data.Nodes))

	dataNode1 := pipeline.Data.Nodes[0]
	assert.NotEmpty(t, dataNode1["id"])
	assert.Equal(t, "read-splunk-firehose", dataNode1["op"])
	assert.Equal(t, "read-splunk-firehose", dataNode1["resolvedId"])

	dataNode2 := pipeline.Data.Nodes[1]
	assert.NotEmpty(t, dataNode2["id"])
	assert.Equal(t, "write-index", dataNode2["op"])
	assert.Empty(t, dataNode2["attributes"])
	assert.Equal(t, "module:index", dataNode2["module"])
	assert.Equal(t, "dataset:main", dataNode2["dataset"])
}

// Test ActivatePipeline streams endpoint
func TestIntegrationActivatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipelineb%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline.Id, *pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline.Status)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)

	// Activate the test pipeline
	boolvar := true
	activatePipelineRequest := streams.ActivatePipelineRequest{
		SkipRestoreState: &boolvar,
	}

	activatePipelineResponse, err := getSdkClient(t).StreamsService.ActivatePipeline(*pipeline.Id, activatePipelineRequest)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, pipeline.Id, activatePipelineResponse.Activated)
	assert.Empty(t, activatePipelineResponse.Deactivated)

	// Get the pipeline and verify that the pipeline status is 'activated'
	pipeline, err = getSdkClient(t).StreamsService.GetPipeline(*pipeline.Id, nil)
	require.Empty(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, streams.PipelineResponseStatusActivated, *pipeline.Status)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)
}

// Test DeactivatePipeline streams endpoint
func TestIntegrationDeactivatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipelinec%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline.Id, *pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline.Status)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)

	// Activate the newly created test pipeline
	boolvar := true
	activatePipelineRequest := streams.ActivatePipelineRequest{

		SkipRestoreState: &boolvar,
	}

	activatePipelineResponse, err := getSdkClient(t).StreamsService.ActivatePipeline(*pipeline.Id, activatePipelineRequest)

	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, pipeline.Id, activatePipelineResponse.Activated)
	time.Sleep(3 * time.Second)

	// Deactivate the active test pipeline
	deactivatePipelineRequest := streams.DeactivatePipelineRequest{
		SkipSavepoint: &boolvar,
	}

	deactivatePipelineResponse, err := getSdkClient(t).StreamsService.DeactivatePipeline(*pipeline.Id, deactivatePipelineRequest)
	require.Nil(t, err)
	require.NotEmpty(t, deactivatePipelineResponse)
	assert.Equal(t, pipeline.Id, deactivatePipelineResponse.Deactivated)
	assert.Empty(t, deactivatePipelineResponse.Activated)

	// Get the test pipeline and verify that the status is 'deactivated'
	pipeline, err = getSdkClient(t).StreamsService.GetPipeline(*pipeline.Id, nil)
	require.Empty(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, "Deactivated", *pipeline.StatusMessage)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)
}

// Test ReactivatePipeline streams endpoint
func TestIntegrationReactivatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipelined%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline.Id, *pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline.Status)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)

	// Activate the newly created test pipeline
	boolvar := true
	activatePipelineRequest := streams.ActivatePipelineRequest{

		SkipRestoreState: &boolvar,
	}
	activatePipelineResponse, err := getSdkClient(t).StreamsService.ActivatePipeline(*pipeline.Id, activatePipelineRequest)

	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, pipeline.Id, activatePipelineResponse.Activated)
	time.Sleep(5 * time.Second)

	// Deactivate the active test pipeline
	deactivatePipelineRequest := streams.DeactivatePipelineRequest{
		SkipSavepoint: &boolvar,
	}

	deactivatePipelineResponse, err := getSdkClient(t).StreamsService.DeactivatePipeline(*pipeline.Id, deactivatePipelineRequest)
	require.Nil(t, err)
	require.NotEmpty(t, deactivatePipelineResponse)
	assert.Equal(t, *pipeline.Id, *deactivatePipelineResponse.Deactivated)
	assert.Empty(t, deactivatePipelineResponse.Activated)

	// Reactivate the deactivated test pipeline
	reactivatePipelineResponse, err := getSdkClient(t).StreamsService.ReactivatePipeline(*pipeline.Id)
	require.Nil(t, err)
	require.NotEmpty(t, reactivatePipelineResponse)
	assert.Equal(t, *pipeline.Id, *reactivatePipelineResponse.PipelineId)
	assert.Equal(t, streams.PipelineReactivateResponsePipelineReactivationStatusAlreadyActivatedWithCurrentVersion, *reactivatePipelineResponse.PipelineReactivationStatus)
}

// Test GetPipelinesStatus streams endpoint
func TestIntegrationGetPipelinesStatus(t *testing.T) {
	pipelineName1 := fmt.Sprintf("testPipelineab%d", testutils.TimeSec)
	pipelineName2 := fmt.Sprintf("testPipelinecd%d", testutils.TimeSec)

	// Create two test pipelines
	pipeline1, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName1, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline1.Id, *pipeline1.Name)
	require.NotEmpty(t, pipeline1)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline1.Status)
	assert.Equal(t, pipelineName1, *pipeline1.Name)
	assert.Equal(t, testPipelineDescription, *pipeline1.Description)

	pipeline2, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName2, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline2.Id, *pipeline2.Name)
	require.NotEmpty(t, pipeline2)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline2.Status)
	assert.Equal(t, pipelineName2, *pipeline2.Name)
	assert.Equal(t, testPipelineDescription, *pipeline2.Description)

	// Get and verify the status of the pipelines
	result, err := getSdkClient(t).StreamsService.GetPipelinesStatus(nil)
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.True(t, *result.Total >= 2)
	require.NotEmpty(t, result.Items)

	// Get and verify the status of the pipelines based on filters (query parameters)
	query := streams.GetPipelinesStatusQueryParams{}.SetName(*pipeline2.Name)
	result, err = getSdkClient(t).StreamsService.GetPipelinesStatus(&query)
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.Equal(t, int64(1), *result.Total)
	require.NotEmpty(t, result.Items)
}

// Test MergePipelines streams endpoint
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
		InputTree:  pipeline1,
		MainTree:   pipeline2,
		TargetPort: pipeline2.Edges[0].TargetPort,
		TargetNode: pipeline2.Edges[0].TargetNode,
	}

	// Merge and verify the status of the merged UPL pipelines
	result, err := getSdkClient(t).StreamsService.MergePipelines(mergeRequest)
	require.Nil(t, err)
	require.NotEmpty(t, result)
	require.NotEmpty(t, result.Edges)
	require.Equal(t, 3, len(result.Edges))
	require.NotEmpty(t, result.Nodes)
	require.Equal(t, 4, len(result.Nodes))
}

// Test UpdatePipeline streams endpoint
func TestIntegrationUpdatePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipelinfe%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline.Id, *pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)

	// Update the newly created test pipeline
	updatedPipelineName := fmt.Sprintf("updated%v", pipelineName)
	desc := fmt.Sprintf("Updated Integration Test Pipeline %v", pipelineName)
	updatedPipeline, err := getSdkClient(t).StreamsService.UpdatePipeline(*pipeline.Id,
		streams.PipelinePatchRequest{Name: &updatedPipelineName, Description: &desc})
	require.Nil(t, err)
	require.NotEmpty(t, updatedPipeline)
	assert.Equal(t, updatedPipelineName, *updatedPipeline.Name)
	assert.Equal(t, desc, *updatedPipeline.Description)
	assert.Equal(t, *pipeline.CurrentVersion+1, *updatedPipeline.CurrentVersion)
}

// Test DeletePipeline streams endpoint
func TestIntegrationDeletePipeline(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipelineg%d", testutils.TimeSec)

	// Create a test pipeline
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline.Status)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)

	// Delete the test pipeline
	deletePipelineResponse, err := getSdkClient(t).StreamsService.DeletePipeline(*pipeline.Id)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)

	// Get the test pipeline and verify that its deleted
	pipeline, err = getSdkClient(t).StreamsService.GetPipeline(*pipeline.Id, nil)
	require.NotEmpty(t, err)
	require.Empty(t, pipeline)
}

// Test Get Input Schema streams endpoint
func TestIntegrationGetInputSchema(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipelineh%d", testutils.TimeSec)
	uplPipeline := createTestUplPipeline(t)
	require.NotEmpty(t, uplPipeline)

	nodeUID := uplPipeline.Edges[0].TargetNode
	port := uplPipeline.Edges[0].TargetPort

	// Create a test pipeline
	boolvar := true
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(streams.PipelineRequest{
		BypassValidation: &boolvar,
		Name:             pipelineName,
		Description:      &testPipelineDescription,
		CreateUserId:     &testutils.TestTenant,
		Data:             uplPipeline,
	})
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline.Id, pipelineName)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline.Status)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)

	// Activate the test pipeline
	activatePipelineRequest := streams.ActivatePipelineRequest{

		SkipRestoreState: &boolvar,
	}
	activatePipelineResponse, err := getSdkClient(t).StreamsService.ActivatePipeline(*pipeline.Id, activatePipelineRequest)

	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, pipeline.Id, activatePipelineResponse.Activated)
	assert.Empty(t, activatePipelineResponse.Deactivated)

	//Get input schema
	result1, err1 := getClient(t).StreamsService.GetInputSchema(streams.GetInputSchemaRequest{NodeUuid: nodeUID, TargetPortName: port, UplJson: uplPipeline})
	require.Empty(t, err1)
	require.NotEmpty(t, result1)
	assert.Equal(t, *result1.Parameters[0].Type, "field")
	assert.Equal(t, *result1.Parameters[0].FieldName, "timestamp")

	assert.NotNil(t, result1.Parameters[0].Parameters)

	assert.Equal(t, *result1.Parameters[0].Parameters[0].Type, "long")

}

// Test Get Output Schema streams endpoint
func TestIntegrationGetOutputSchema(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipelinei%d", testutils.TimeSec)
	uplPipeline := createTestUplPipeline(t)
	require.NotEmpty(t, uplPipeline)

	//nodeUID := uplPipeline.Edges[0].SourceNode
	//port := uplPipeline.Edges[0].SourcePort

	// Create a test pipeline
	boolvar := true
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(streams.PipelineRequest{
		BypassValidation: &boolvar,
		Name:             pipelineName,
		Description:      &testPipelineDescription,
		CreateUserId:     &testutils.TestTenant,
		Data:             uplPipeline,
	})
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline.Id, pipelineName)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline.Status)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)

	// Activate the test pipeline
	activatePipelineRequest := streams.ActivatePipelineRequest{

		SkipRestoreState: &boolvar,
	}
	activatePipelineResponse, err := getSdkClient(t).StreamsService.ActivatePipeline(*pipeline.Id, activatePipelineRequest)

	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, pipeline.Id, activatePipelineResponse.Activated)
	assert.Empty(t, activatePipelineResponse.Deactivated)

	//Get output schema
	//_, err1 := getClient(t).StreamsService.GetOutputSchema(streams.GetOutputSchemaRequest{&nodeUID, &port, uplPipeline})
	//require.Empty(t, err1)
	////TODO(shilpa) Follow up when INGEST-8089 is investigated. Currently the output from this call could be empty sometimes
	////require.NotEmpty(t, result1)
	////assert.Equal(t, *result1.Parameters[0].Type, "field")
	////assert.Equal(t, *result1.Parameters[0].FieldName, "timestamp")
	////assert.Equal(t, *result1.Parameters[0].Parameters[0].Type, "long")
}

// Test Get Registry endpoint
func TestIntegrationGetRegistry(t *testing.T) {
	//Set local query parameter
	local := make(url.Values)
	local.Add("local", `true`)

	query := streams.GetRegistryQueryParams{}.SetLocal(true)
	result, err := getSdkClient(t).StreamsService.GetRegistry(&query)
	require.Empty(t, err)
	require.NotEmpty(t, result)
	assert.NotEmpty(t, (result.Functions)[0].Categories)
	assert.True(t, len((result.Functions)[0].Categories) > 0)
	assert.NotEmpty(t, (result.Categories)[0].Id)
	assert.NotEmpty(t, (result.Types)[0].Type)
}

//Test Get Latest pipeline metrics endpoint
func TestIntegrationGetLatestPipelineMetrics(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipelinej%d", testutils.TimeSec)

	uplPipeline := createTestUplPipeline(t)
	require.NotEmpty(t, uplPipeline)

	// Create a test pipeline
	boolvar := true
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(streams.PipelineRequest{
		BypassValidation: &boolvar,
		Name:             pipelineName,
		Description:      &testPipelineDescription,
		CreateUserId:     &testutils.TestTenant,
		Data:             uplPipeline,
	})
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline.Id, pipelineName)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline.Status)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)

	// Activate the test pipeline
	activatePipelineRequest := streams.ActivatePipelineRequest{

		SkipRestoreState: &boolvar,
	}
	activatePipelineResponse, err := getSdkClient(t).StreamsService.ActivatePipeline(*pipeline.Id, activatePipelineRequest)

	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, pipeline.Id, activatePipelineResponse.Activated)
	assert.Empty(t, activatePipelineResponse.Deactivated)

	//Get latest pipeline metrics
	//Validation of the metrics output is not reliable since its real-time data, no guarantees if metric data will be populated at that instant of time
	//Attempt the call to get metrics 5 times and validate if there is data returned.
	cnt := 0
	for cnt < 5 {
		result1, err1 := getClient(t).StreamsService.GetPipelineLatestMetrics(*pipeline.Id)
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

}

//Test Latest Preview Session Metrics
func TestIntegrationGetLatestPreviewSessionMetrics(t *testing.T) {
	// Create and start a test Ge session
	response, err := getSdkClient(t).StreamsService.StartPreview(createPreviewSessionStartRequest(t))
	require.Nil(t, err)
	require.NotEmpty(t, response)
	assert.NotEmpty(t, response.PreviewId)

	response1, err := getSdkClient(t).StreamsService.StopPreview(*response.PreviewId)
	require.Nil(t, err)
	assert.NotEmpty(t, response1)

	////Get latest preview session metrics
	////Validation of the metrics output is not reliable since its real-time data, no guarantees if metric data will be populated at that instant of time
	////Attempt the call to get metrics 5 times and validate if there is data returned.
	//cnt := 0
	//for cnt < 5 {
	//	result1, err1 := getClient(t).StreamsService.GetPipelineLatestMetrics(strconv.FormatInt(*response.PreviewId, 10))
	//	require.Empty(t, err1)
	//	require.NotEmpty(t, result1)
	//	if len(result1.Nodes) > 0 {
	//		for key, element := range result1.Nodes {
	//			assert.NotEmpty(t, key)
	//			assert.NotEmpty(t, element.Metrics)
	//		}
	//	}
	//	time.Sleep(20 * time.Second)
	//	cnt++
	//}
}

// Test Get Connectors
func TestIntegrationListConnectors(t *testing.T) {
	response, err := getSdkClient(t).StreamsService.ListConnectors()
	require.Nil(t, err)
	require.NotEmpty(t, response)
}

// Test CRUD Connections
func TestIntegrationCRUEConnections(t *testing.T) {
	connectorId := "879837b0-cabf-4bc2-8589-fcc4dad753e7" //Splunk Enterprise Connector
	data := make(map[string]interface{})
	data["splunk-url"] = "https://hostname.port"
	data["token"] = "mytoken"
	connectionName := fmt.Sprintf("testConnection%d", testutils.TimeSec)
	connection, err := getSdkClient(t).StreamsService.CreateConnection(streams.ConnectionRequest{ConnectorId: connectorId, Data: data, Name: connectionName})
	require.Nil(t, err)
	assert.NotEmpty(t, connection)
	defer getSdkClient(t).StreamsService.DeleteConnection(*connection.Id)

	connections, err := getSdkClient(t).StreamsService.ListConnections(nil)
	require.Nil(t, err)
	require.NotEmpty(t, connections)

	query := streams.ListConnectionsQueryParams{}.SetName(connectionName)
	response1, err := getSdkClient(t).StreamsService.ListConnections(&query)
	require.Nil(t, err)
	require.NotEmpty(t, response1)
	require.Equal(t, 1, len(response1.Items))
	assert.Equal(t, connectorId, *response1.Items[0].ConnectorId)
	assert.Equal(t, *connection.Name, *response1.Items[0].Versions[0].Name)
	assert.Equal(t, *connection.Id, *response1.Items[0].Id)

	err = getSdkClient(t).StreamsService.DeleteConnection(*connection.Id)
	assert.Nil(t, err)
}

// Test Validate Upl Response streams endpoint
func TestIntegrationValidateResponse(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipelinek%d", testutils.TimeSec)

	uplPipeline := createTestUplPipeline(t)
	require.NotEmpty(t, uplPipeline)

	// Create a test pipeline
	boolvar := true
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(streams.PipelineRequest{
		BypassValidation: &boolvar,
		Name:             pipelineName,
		Description:      &testPipelineDescription,
		CreateUserId:     &testutils.TestTenant,
		Data:             uplPipeline,
	})
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline.Id, *pipeline.Name)
	require.NotEmpty(t, pipeline)
	assert.Equal(t, streams.PipelineResponseStatusCreated, *pipeline.Status)
	assert.Equal(t, pipelineName, *pipeline.Name)
	assert.Equal(t, testPipelineDescription, *pipeline.Description)

	//Validate Upl response
	result1, err1 := getClient(t).StreamsService.ValidatePipeline(streams.ValidateRequest{Upl: uplPipeline})
	require.Empty(t, err1)
	require.NotEmpty(t, result1)
	assert.Equal(t, *result1.Success, true)

}

// Test StartPreviewSession streams endpoint
func TestIntegrationStartPreviewSession(t *testing.T) {
	// Create and start a test preview session
	response, err := getSdkClient(t).StreamsService.StartPreview(createPreviewSessionStartRequest(t))
	require.Nil(t, err)
	require.NotEmpty(t, response)
	//previewIDStringVal := strconv.FormatInt(*response.PreviewId, 10)
	//defer cleanupPreview(t, previewIDStringVal)
	assert.NotEmpty(t, response.PreviewId)

	// Verify that the test preview session is created
	previewState, err := getSdkClient(t).StreamsService.GetPreviewSession(*response.PreviewId)
	require.Nil(t, err)
	require.NotEmpty(t, previewState)
	assert.NotEmpty(t, response.PreviewId, previewState.PreviewId)
	assert.NotEmpty(t, previewState.JobId)
}

// Test DeletePreviewSession streams endpoint
func TestIntegrationDeletePreviewSession(t *testing.T) {
	// Create and start a test preview session
	response, err := getSdkClient(t).StreamsService.StartPreview(createPreviewSessionStartRequest(t))
	require.Nil(t, err)
	require.NotEmpty(t, response)
	assert.NotEmpty(t, response.PreviewId)

	// Delete the test preview session
	_, err = getSdkClient(t).StreamsService.StopPreview(*response.PreviewId)
	require.Nil(t, err)

	// Verify that the test preview session is deleted
	_, err = getSdkClient(t).StreamsService.GetPreviewSession(*response.PreviewId)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
	assert.Equal(t, "preview-id-not-found", httpErr.Code)
}

// Test GetPreviewData streams endpoint
func TestIntegrationGetPreviewData(t *testing.T) {
	// Create and start a test preview session
	response, err := getSdkClient(t).StreamsService.StartPreview(createPreviewSessionStartRequest(t))
	require.Nil(t, err)
	require.NotEmpty(t, response)
	//previewIDStringVal := strconv.FormatInt(*response.PreviewId, 10)
	//defer cleanupPreview(t, *response.PreviewId)
	assert.NotEmpty(t, response.PreviewId)

	// Verify that the preview data is generated
	previewData, err := getSdkClient(t).StreamsService.GetPreviewData(*response.PreviewId)
	require.Nil(t, err)
	require.NotEmpty(t, previewData)
	assert.NotEmpty(t, response.PreviewId, previewData.PreviewId)
}

// Test CreateTemplate streams endpoint
func TestIntegrationCreateTemplate(t *testing.T) {
	templateName := fmt.Sprintf("testTemplate%d", testutils.TimeSec)

	// Create a test template and verify that the template was created
	template, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, *template.TemplateId)
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

	assert.NotEmpty(t, (*template.Data).Nodes[0]["id"])
	assert.Equal(t, "read-splunk-firehose", (*template.Data).Nodes[0]["op"])

	dataNode2 := (*template.Data).Nodes[1]
	assert.NotEmpty(t, dataNode2["id"])
	assert.Equal(t, "write-index", dataNode2["op"])
	assert.Empty(t, dataNode2["attributes"])
}

// Test GetTemplates streams endpoint
func TestIntegrationGetAllTemplates(t *testing.T) {
	templateName1 := fmt.Sprintf("testTemplate01%d", testutils.TimeSec)
	templateName2 := fmt.Sprintf("testTemplate02%d", testutils.TimeSec)

	// Create two test templates
	template1, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName1, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, *template1.TemplateId)
	require.NotEmpty(t, template1)
	assert.Equal(t, templateName1, *template1.Name)
	assert.Equal(t, testTemplateDescription, *template1.Description)

	template2, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName2, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, *template2.TemplateId)
	require.NotEmpty(t, template2)
	assert.Equal(t, templateName2, *template2.Name)
	assert.Equal(t, testTemplateDescription, *template2.Description)

	// Get all the templates
	result, err := getSdkClient(t).StreamsService.ListTemplates(nil)
	require.Empty(t, err)
	require.NotEmpty(t, result)
}

// Test UpdateTemplate streams endpoint
func TestIntegrationUpdateTemplate(t *testing.T) {
	templateName := fmt.Sprintf("testTemplate%d", testutils.TimeSec)

	// Create a test template and verify that the template was created
	template, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, *template.TemplateId)
	require.NotEmpty(t, template)
	assert.Equal(t, templateName, *template.Name)
	assert.Equal(t, testTemplateDescription, *template.Description)

	// Update the newly created test template (partial update data is provided)
	updatedTemplateName := fmt.Sprintf("updated%v", templateName)
	desc := "Updated Integration Test Template"
	updatedTemplate, err := getSdkClient(t).StreamsService.UpdateTemplate(*template.TemplateId,
		streams.TemplatePatchRequest{Description: &desc, Name: &updatedTemplateName})
	require.Nil(t, err)
	require.NotEmpty(t, updatedTemplate)
	assert.Equal(t, updatedTemplateName, *updatedTemplate.Name)
	assert.Equal(t, "Updated Integration Test Template", *updatedTemplate.Description)
	assert.Equal(t, *template.Version+1, *updatedTemplate.Version)
}

// Test PutTemplate streams endpoint
func TestIntegrationPutTemplate(t *testing.T) {
	templateName := fmt.Sprintf("testTemplate%d", testutils.TimeSec)

	// Create a test template and verify that the template was created
	template, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, *template.TemplateId)
	require.NotEmpty(t, template)
	assert.Equal(t, templateName, *template.Name)
	assert.Equal(t, testTemplateDescription, *template.Description)

	// Update the newly created test template with PUT
	updatedDescription := "Updated Integration Test Template"
	result := createTestUplPipeline(t)

	updatedTemplate, err := getSdkClient(t).StreamsService.PutTemplate(*template.TemplateId,
		streams.TemplatePutRequest{Name: "new" + templateName, Description: updatedDescription, Data: result})
	require.Nil(t, err)
	require.NotEmpty(t, updatedTemplate)
	assert.Equal(t, "new"+templateName, *updatedTemplate.Name)
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
	err = getSdkClient(t).StreamsService.DeleteTemplate(*template.TemplateId)
	require.Nil(t, err)

	// Verify that the test template is deleted
	_, err = getSdkClient(t).StreamsService.GetTemplate(*template.TemplateId, nil)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
	assert.Equal(t, "template-id-not-found", httpErr.Code)
}

// Test Get Groups endpoint
func TestIntegrationGetGroups(t *testing.T) {
	request := streams.GetRegistryQueryParams{}.SetLocal(false)
	result, err := getSdkClient(t).StreamsService.GetRegistry(&request)
	require.Empty(t, err)
	require.NotEmpty(t, result)
	//assert.NotEmpty(t, (result.Functions)[0].Id)
	//assert.NotEmpty(t, (result.Categories)[0].Id)
	//assert.NotEmpty(t, (result.Types)[0].Type)

	pipelineName := fmt.Sprintf("testPipelineef%d", testutils.TimeSec)
	pipeline, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, pipelineName, testPipelineDescription))
	require.Nil(t, err)
	defer cleanupPipeline(getSdkClient(t), *pipeline.Id, *pipeline.Name)

	group, err := getSdkClient(t).StreamsService.CreateGroup(streams.GroupRequest{
		Arguments: []streams.GroupArgumentsNode{{GroupArg: "time", Type: "long", Position: 0}},
		Name:      "test-fn",
		Ast:       *pipeline.Data,
	})
	require.Nil(t, err)

	//cnt := 0
	//temp := 0
	//for cnt < len(result.Functions) {
	//	if result.Functions[cnt].ResolvedId != nil && strings.Contains(*result.Functions[cnt].ResolvedId, "receive-from-ingest-rest-api") {
	//		temp = cnt
	//		break
	//
	//	}
	//	cnt++
	//}

	//assert.True(t, cnt < len(result.Functions))

	//applicationData := ((result.Functions)[temp].Attributes)["application"]
	//application, ok := ((result.Functions)[temp].Attributes)["application"].(map[string]interface{})
	//assert.NotEmpty(t, applicationData)
	//require.True(t, ok)
	//groupId, ok := application["groupId"].(string)
	//require.True(t, ok)

	test, err := getSdkClient(t).StreamsService.GetGroup(*group.GroupId)
	require.Empty(t, err)
	require.NotEmpty(t, test)
	assert.NotEmpty(t, *test.Name)
	assert.NotEmpty(t, *test.CreateUserId)
	assert.NotEmpty(t, *test.OutputType)
}

//
////Test the Create Expanded version of the group Endpoint
//func TestIntegrationCreateExpandedGroup(t *testing.T) {
//	local := make(url.Values)
//	local.Add("local", `true`)
//	//GetRegistry to retrieve the groupID
//	pipeline1, err := getSdkClient(t).StreamsService.CreatePipeline(makePipelineRequest(t, "mypipeline", testPipelineDescription))
//	require.Nil(t, err)
//	defer cleanupPipeline(getSdkClient(t), *pipeline1.Id, *pipeline1.Name)
//
//	result, err := getSdkClient(t).StreamsService.GetUPLRegistry((&streams.GetUPLRegistryQueryParams{}.SetLocal(true))
//	require.Empty(t, err)
//	require.NotEmpty(t, result)
//	assert.NotEmpty(t, (result.Functions)[0].Categories)
//	assert.NotEmpty(t, (result.Categories)[0].Id)
//	assert.NotEmpty(t, (result.Types)[0].Type)
//
//	cnt := 0
//	temp := 0
//	for cnt < len(result.Functions) {
//		fmt.Println(*result.Functions[cnt].ResolvedId)
//		if result.Functions[cnt].ResolvedId != nil && strings.Contains(*result.Functions[cnt].ResolvedId, "read-splunk-firehose()") {
//			temp = cnt
//			break
//
//		}
//		cnt++
//	}
//
//	assert.True(t, cnt < len(result.Functions))
//	application, ok := ((result.Functions)[temp].Attributes)["application"].(map[string]interface{})
//	require.True(t, ok)
//	groupId, ok := application["groupId"].(string)
//	require.True(t, ok)
//
//	assert.NotEmpty(t, groupId)
//
//	//GetGroupID to get the Group Function ID
//
//	result1, err := getSdkClient(t).StreamsService.GetGroupById(groupId)
//	require.Empty(t, err)
//	require.NotEmpty(t, result1)
//	assert.NotEmpty(t, *result1.Name)
//	assert.NotEmpty(t, *result1.CreateUserId)
//	assert.NotEmpty(t, *result1.OutputType)
//
//	//functionID := (result1.Mappings)[0].FunctionId
//	//
//	//type argumentsMap map[string]interface{}
//	//arguments := argumentsMap{"group_arg": "connection", "function_arg": "right"}
//
//	//result2, err := getClient(t).StreamsService.CreateExpandedGroup(groupId, arguments, functionID)
//	//require.Empty(t, err)
//	//require.NotEmpty(t, result2)
//	//assert.NotEmpty(t, result2.Version)
//	//assert.NotEmpty(t, result2.RootNode)
//	//
//	//assert.NotEmpty(t, (result2.Nodes)[0].Id)
//	//
//	//assert.NotEmpty(t, (result2.Nodes)[0].Id)
//	//
//	//assert.Empty(t, (result2.Nodes)[0].Attributes)
//	//assert.NotEmpty(t, result2.Edges[0].SourceNode)
//	//assert.NotEmpty(t, result2.Edges[0].TargetNode)
//}

//makePipelineRequest is a helper function to make a PipelineRequest model
func makePipelineRequest(t *testing.T, name string, description string) streams.PipelineRequest {
	result := createTestUplPipeline(t)

	boolvar := true
	return streams.PipelineRequest{
		BypassValidation: &boolvar,
		Name:             name,
		Description:      &description,
		CreateUserId:     &testutils.TestTenant,
		Data:             result,
	}
}

// createTestUplPipeline is a helper function to create a test UPL JSON from a test DSL.
func createTestUplPipeline(t *testing.T) streams.UplPipeline {
	var dsl = "events = read-splunk-firehose(); write-index(events, \"module:index\", \"dataset:main\");"
	result, err := getSdkClient(t).StreamsService.CompileDSL(streams.DslCompilationRequest{Dsl: dsl})
	require.Empty(t, err)
	require.NotEmpty(t, result)

	resp, err := getSdkClient(t).StreamsService.ValidatePipeline(streams.ValidateRequest{Upl: *result})
	assert.NotNil(t, resp)
	assert.Nil(t, err)

	return *result
}

// createPreviewSessionStartRequest is a helper function to create a test PreviewSessionStartRequest model
func createPreviewSessionStartRequest(t *testing.T) streams.PreviewSessionStartRequest {
	result := createTestUplPipeline(t)
	recordsLimit := int32(100)
	recordsPerPipeline := int32(2)
	sessionLifetimeMs := int64(300000)
	useNewData := true

	return streams.PreviewSessionStartRequest{
		RecordsLimit:       &recordsLimit,
		RecordsPerPipeline: &recordsPerPipeline,
		SessionLifetimeMs:  &sessionLifetimeMs,
		Upl:                result,
		UseNewData:         &useNewData,
	}
}

// makeTemplateRequest is a helper function to make a TemplateRequest model
func makeTemplateRequest(t *testing.T, name string, description string) streams.TemplateRequest {
	result := createTestUplPipeline(t)

	return streams.TemplateRequest{
		Data:        result,
		Description: description,
		Name:        name,
	}
}

//Deletes the test pipeline
func cleanupPipeline(client *sdk.Client, id string, name string) {
	_, err := client.StreamsService.DeletePipeline(id)
	if err != nil {
		fmt.Printf("WARN: error deleting pipeline: name:%s, err: %s", name, err)
	}
}

// Deletes the test preview-session
func cleanupPreview(t *testing.T, id int64) {
	// First check if the preview session exists

	//n, err := strconv.ParseInt(id, 10, 64)
	//if err == nil {
	//	fmt.Printf("%d of type %T", n, n)
	//}

	previewState, err := getSdkClient(t).StreamsService.GetPreviewSession(id)
	if err == nil && previewState != nil {
		_, err = getSdkClient(t).StreamsService.DeletePipeline(strconv.FormatInt(id, 10))
		if err != nil {
			fmt.Printf("WARN: error deleting preview session: id:%v, err: %s", id, err)
		}
	}
	fmt.Println(previewState)
	fmt.Println(err)
}

// Deletes the test template
func cleanupTemplate(t *testing.T, id string) {
	err := getSdkClient(t).StreamsService.DeleteTemplate(id)
	if err != nil {
		fmt.Printf("WARN: error deleting template: id:%s, err: %s", id, err)
	}
}
