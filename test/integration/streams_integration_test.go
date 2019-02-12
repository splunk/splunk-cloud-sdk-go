package integration

import (
	"fmt"
	"testing"

	"net/url"
	"strconv"
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

// Test Get Input Schema streams endpoint
func TestIntegrationGetInputSchema(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	var dsl = "events = read-splunk-firehose(); write-splunk-index(events);"
	uplPipeline, err := getClient(t).StreamsService.CompileDslToUpl(&model.DslCompilationRequest{Dsl: dsl})
	require.Empty(t, err)
	require.NotEmpty(t, uplPipeline)

	nodeUid := uplPipeline.Edges[0].TargetNode
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

	//Get input Schema
	result1, err1 := getClient(t).StreamsService.GetInputSchema(&streams.GetInputSchemaRequest{NodeUUID: &nodeUid, TargetPortName: &port, UplJSON: uplPipeline})
	require.Empty(t, err1)
	require.NotEmpty(t, result1)
	assert.Equal(t, *result1.Parameters[0].Type, "field")
	assert.Equal(t, *result1.Parameters[0].FieldName, "timestamp")
	assert.Equal(t, *result1.Parameters[0].Parameters[0].Type, "long")

	// Delete the test pipeline
	deletePipelineResponse, err := getClient(t).StreamsService.DeletePipeline(pipeline.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)

}

// Test Get Input Schema streams endpoint
func TestIntegrationGetOutputSchema(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	var dsl = "events = read-splunk-firehose(); write-splunk-index(events);"
	uplPipeline, err := getClient(t).StreamsService.CompileDslToUpl(&model.DslCompilationRequest{Dsl: dsl})
	require.Empty(t, err)
	require.NotEmpty(t, uplPipeline)

	nodeUid := uplPipeline.Edges[0].SourceNode
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

	//Get input Schema
	_, err1 := getClient(t).StreamsService.GetOutputSchema(&streams.GetOutputSchemaRequest{NodeUUID: &nodeUid, SourcePortName: &port, UplJSON: uplPipeline})
	require.Empty(t, err1)
	//TODO(shilpa) Follow up why output schema result is empty
	//require.NotEmpty(t, result1)
	//assert.Equal(t, *result1.Parameters[0].Type, "field")
	//assert.Equal(t, *result1.Parameters[0].FieldName, "timestamp")
	//assert.Equal(t, *result1.Parameters[0].Parameters[0].Type, "long")

	// Delete the test pipeline
	deletePipelineResponse, err := getClient(t).StreamsService.DeletePipeline(pipeline.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)
}

// Test Get Registry endpoint
func TestIntegrationGetRegistry(t *testing.T) {
	//Set local query parameter
	local := make(url.Values)
	local.Add("local", `true`)
	//ToDo(shilpa) Follow up Element-Type collision during json unmarshalling
	result, err := getClient(t).StreamsService.GetRegistry(local)
	require.NotEmpty(t, err)
	require.NotEmpty(t, result)
}

//Test Get Latest pipeline metrics endpoint
//TODO(shilpa) Follow up potential timing issue for the metrics to return
func TestIntegrationGetLatestPipelineMetrics(t *testing.T) {
	pipelineName := fmt.Sprintf("testPipeline%d", testutils.TimeSec)

	var dsl = "events = read-splunk-firehose(); write-splunk-index(events);"
	uplPipeline, err := getClient(t).StreamsService.CompileDslToUpl(&model.DslCompilationRequest{Dsl: dsl})
	require.Empty(t, err)
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

	// Activate the test pipeline
	ids := []string{pipeline.ID}
	activatePipelineResponse, err := getClient(t).StreamsService.ActivatePipeline(ids)
	require.Nil(t, err)
	require.NotEmpty(t, activatePipelineResponse)
	assert.Equal(t, []string{pipeline.ID}, activatePipelineResponse["activated"])
	assert.Empty(t, activatePipelineResponse["notActivated"])

	//Get latest pipeline metrics
	result1, err1 := getClient(t).StreamsService.GetLatestPipelineMetrics(pipeline.ID)
	time.Sleep(time.Duration(60) * time.Second)
	require.Empty(t, err1)
	require.NotEmpty(t, result1)
	assert.NotEmpty(t, result1.Nodes)

	for key, element := range result1.Nodes {
		assert.NotEmpty(t, key)
		assert.NotEmpty(t, element.Metrics)
	}
	// Delete the test pipeline
	deletePipelineResponse, err := getClient(t).StreamsService.DeletePipeline(pipeline.ID)
	require.Nil(t, err)
	require.NotNil(t, deletePipelineResponse)

}

//Test Latest Preview Session Metrics
//TODO(shilpa) Follow up potential timing issue for the metrics to return
func TestIntegrationGetLatestPreviewSessionMetrics(t *testing.T) {

	// Create and start a test preview session
	response, err := getSdkClient(t).StreamsService.StartPreviewSession(createPreviewSessionStartRequest(t))
	require.Nil(t, err)
	require.NotEmpty(t, response)
	previewIdStringVal := strconv.FormatInt(response.PreviewID, 10)
	defer cleanupPreview(t, previewIdStringVal)
	assert.NotEmpty(t, response.PipelineID)
	assert.NotEmpty(t, response.PreviewID)

	//Get latest pipeline metrics
	result1, err1 := getClient(t).StreamsService.GetLatestPreviewSessionMetrics(previewIdStringVal)
	time.Sleep(time.Duration(60) * time.Second)
	require.Empty(t, err1)
	require.NotEmpty(t, result1)
	assert.NotEmpty(t, result1.Nodes)

	for key, element := range result1.Nodes {
		assert.NotEmpty(t, key)
		assert.NotEmpty(t, element.Metrics)
	}
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

	var dsl = "events = read-splunk-firehose(); write-splunk-index(events);"
	uplPipeline, err := getClient(t).StreamsService.CompileDslToUpl(&model.DslCompilationRequest{Dsl: dsl})
	require.Empty(t, err)
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

	//Get input Schema
	result1, err1 := getClient(t).StreamsService.ValidateUplResponse(&streams.ValidateRequest{Upl: uplPipeline})
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
	previewIdStringVal := strconv.FormatInt(response.PreviewID, 10)
	defer cleanupPreview(t, previewIdStringVal)
	assert.NotEmpty(t, response.PipelineID)
	assert.NotEmpty(t, response.PreviewID)

	// Verify that the test preview session is created
	previewState, err := getSdkClient(t).StreamsService.GetPreviewSession(previewIdStringVal)
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
	previewIdStringVal := strconv.FormatInt(response.PreviewID, 10)
	assert.NotEmpty(t, response.PipelineID)
	assert.NotEmpty(t, response.PreviewID)

	// Delete the test preview session
	err = getSdkClient(t).StreamsService.DeletePreviewSession(previewIdStringVal)
	require.Nil(t, err)

	// Verify that the test preview session is deleted
	_, err = getSdkClient(t).StreamsService.GetPreviewSession(previewIdStringVal)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
	assert.Equal(t, "preview-id-not-found", httpErr.Code)
}

// Test CreateTemplate streams endpoint
func TestIntegrationCreateTemplate(t *testing.T) {
	templateName := fmt.Sprintf("testTemplate%d", testutils.TimeSec)

	// Create a test template and verify that the template was created
	template, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, template.TemplateID)
	require.NotEmpty(t, template)
	assert.Equal(t, templateName, template.Name)
	assert.Equal(t, testTemplateDescription, template.Description)

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
	defer cleanupTemplate(t, template1.TemplateID)
	require.NotEmpty(t, template1)
	assert.Equal(t, templateName1, template1.Name)
	assert.Equal(t, testTemplateDescription, template1.Description)

	template2, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName2, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, template2.TemplateID)
	require.NotEmpty(t, template2)
	assert.Equal(t, templateName2, template2.Name)
	assert.Equal(t, testTemplateDescription, template2.Description)

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
	defer cleanupTemplate(t, template.TemplateID)
	require.NotEmpty(t, template)
	assert.Equal(t, templateName, template.Name)
	assert.Equal(t, testTemplateDescription, template.Description)

	// Update the newly created test template
	updatedTemplateName := fmt.Sprintf("updated%v", templateName)
	updatedTemplate, err := getSdkClient(t).StreamsService.UpdateTemplate(template.TemplateID, makeTemplateRequest(t, updatedTemplateName, "Updated Integration Test Template"))
	require.Nil(t, err)
	require.NotEmpty(t, updatedTemplate)
	assert.Equal(t, updatedTemplateName, updatedTemplate.Name)
	assert.Equal(t, "Updated Integration Test Template", updatedTemplate.Description)
	assert.Equal(t, template.Version+1, updatedTemplate.Version)
}

// Test PartialUpdateTemplate streams endpoint
func TestIntegrationPartialUpdateTemplate(t *testing.T) {
	templateName := fmt.Sprintf("testTemplate%d", testutils.TimeSec)

	// Create a test template and verify that the template was created
	template, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, template.TemplateID)
	require.NotEmpty(t, template)
	assert.Equal(t, templateName, template.Name)
	assert.Equal(t, testTemplateDescription, template.Description)

	// Update the newly created test template (partial update data is provided)
	updatedTemplate, err := getSdkClient(t).StreamsService.UpdateTemplatePartially(template.TemplateID, &streams.PartialTemplateRequest{Description: "Updated Integration Test Template"})
	require.Nil(t, err)
	require.NotEmpty(t, updatedTemplate)
	assert.Equal(t, templateName, updatedTemplate.Name)
	assert.Equal(t, "Updated Integration Test Template", updatedTemplate.Description)
	assert.Equal(t, template.Version+1, updatedTemplate.Version)
}

// Test DeleteTemplate streams endpoint
func TestIntegrationDeleteTemplate(t *testing.T) {
	templateName := fmt.Sprintf("testTemplate%d", testutils.TimeSec)

	// Create a test template and verify that the template was created
	template, err := getSdkClient(t).StreamsService.CreateTemplate(makeTemplateRequest(t, templateName, testTemplateDescription))
	require.Nil(t, err)
	defer cleanupTemplate(t, template.TemplateID)
	require.NotEmpty(t, template)
	assert.Equal(t, templateName, template.Name)
	assert.Equal(t, testTemplateDescription, template.Description)

	// Delete the test template
	err = getSdkClient(t).StreamsService.DeleteTemplate(template.TemplateID)
	require.Nil(t, err)

	// Verify that the test template is deleted
	_, err = getSdkClient(t).StreamsService.GetTemplate(template.TemplateID)
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
	assert.Equal(t, "template-id-not-found", httpErr.Code)
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

	return &streams.PreviewSessionStartRequest{
		RecordsLimit:       100,
		RecordsPerPipeline: 2,
		SessionLifetimeMs:  10000,
		Upl:                result,
		UseNewData:         false,
	}
}

// makeTemplateRequest is a helper function to make a TemplateRequest model
func makeTemplateRequest(t *testing.T, name string, description string) *streams.TemplateRequest {
	result := createTestUplPipeline(t)

	return &streams.TemplateRequest{
		Data:        result,
		Description: description,
		Name:        name,
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
