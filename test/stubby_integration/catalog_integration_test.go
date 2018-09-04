// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package stubbyintegration

import (
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test variables
var testDatasetID = "TEST_DATASET_ID"
var testFieldID1 = "TEST_FIELD_ID_01"

// Stubby test for GetDataset() catalog service endpoint
func TestGetDataset(t *testing.T) {
	result, err := getClient(t).CatalogService.GetDataset("ds1")

	require.Empty(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, "ds1", result.ID)
	assert.Equal(t, model.INDEX, result.Kind)
}

// Stubby test for GetDatasets() catalog service endpoint
func TestGetDatasets(t *testing.T) {
	result, err := getClient(t).CatalogService.GetDatasets()

	require.Empty(t, err)
	assert.Equal(t, 2, len(result))
}

// Stubby test for CreateDataset() catalog service endpoint
func TestPostDataset(t *testing.T) {
	result, err := getClient(t).CatalogService.CreateDataset(
		model.DatasetInfo{Name: "stubby_dataset_1", Kind: model.INDEX, Owner: "Splunk", Capabilities: "1101-00000:11010", Disabled: true})

	require.Empty(t, err)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "stubby_dataset_1", result.Name)
	assert.Equal(t, model.INDEX, result.Kind)
}

// Stubby test for UpdateDataset() catalog service endpoint
func TestUpdateDataset(t *testing.T) {
	result, err := getClient(t).CatalogService.UpdateDataset(
		model.PartialDatasetInfo{Disabled: true, Version: 5}, "ds1")
	require.Empty(t, err)
	assert.NotEmpty(t, result)
	assert.IsType(t, &(model.DatasetInfo{}), result)
	assert.NotNil(t, result.ID)
	assert.Equal(t, "stubby_dataset_1", result.Name)
	assert.Equal(t, model.INDEX, result.Kind)
}

// Stubby test for DeleteDataset() catalog service endpoint
func TestDeleteDataset(t *testing.T) {
	err := getClient(t).CatalogService.DeleteDataset("ds1")
	require.Empty(t, err)
}

// Stubby test for DeleteRule() catalog service endpoint
func TestDeleteRule(t *testing.T) {
	err := getClient(t).CatalogService.DeleteRule("rule1")
	require.Empty(t, err)
}

// Stubby test for GetRules() catalog service endpoint
func TestGetRules(t *testing.T) {
	result, err := getClient(t).CatalogService.GetRules()

	require.Empty(t, err)
	assert.Equal(t, 2, len(result))
}

// Stubby test for GetRule() catalog service endpoint
func TestGetRule(t *testing.T) {
	result, err := getClient(t).CatalogService.GetRule("rule1")
	require.Empty(t, err)

	assert.NotNil(t, "rule1", result.ID)
	assert.Equal(t, "_internal", result.Name)
}

// Stubby test for GetDatasetFields() catalog service endpoint
func TestGetDatasetFields(t *testing.T) {
	result, err := getClient(t).CatalogService.GetDatasetFields(testDatasetID, nil)

	require.Empty(t, err)
	assert.Equal(t, 3, len(result))
}

// Stubby test for GetDatasetField() catalog service endpoint
func TestGetDatasetField(t *testing.T) {
	result, err := getClient(t).CatalogService.GetDatasetField(testDatasetID, testFieldID1)

	require.Empty(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, "date_second", result.Name)
	assert.Equal(t, model.NUMBER, result.DataType)
	assert.Equal(t, model.DIMENSION, result.FieldType)
	assert.Equal(t, model.ALL, result.Prevalence)
}

// Stubby test for PostDatasetField() catalog service endpoint
func TestPostDatasetField(t *testing.T) {
	testField := model.Field{Name: "test_data", DatasetID: testDatasetID, DataType: "S", FieldType: "D", Prevalence: "A"}
	result, err := getClient(t).CatalogService.PostDatasetField(testDatasetID, testField)

	require.Empty(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, testFieldID1, result.ID)
	assert.Equal(t, "test_data", result.Name)
	assert.Equal(t, model.STRING, result.DataType)
	assert.Equal(t, model.DIMENSION, result.FieldType)
	assert.Equal(t, model.ALL, result.Prevalence)
}

// Stubby test for PatchDatasetField() catalog service endpoint
func TestPatchDatasetField(t *testing.T) {
	testField := model.Field{Name: "test_data", DatasetID: testDatasetID, DataType: "N", FieldType: "D", Prevalence: "A"}
	result, err := getClient(t).CatalogService.PatchDatasetField(testDatasetID, testFieldID1, testField)

	require.Empty(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, testFieldID1, result.ID)
	assert.Equal(t, "test_data", result.Name)
	assert.Equal(t, model.NUMBER, result.DataType)
	assert.Equal(t, model.DIMENSION, result.FieldType)
	assert.Equal(t, model.ALL, result.Prevalence)
}

// Stubby test for DeleteDatasetField() catalog service endpoint
func TestDeleteDatasetField(t *testing.T) {
	err := getClient(t).CatalogService.DeleteDatasetField(testDatasetID, testFieldID1)

	require.Empty(t, err)
}

// Stubby test for CreateRule() catalog service endpoint
func TestPostRule(t *testing.T) {
	var actions [3]model.CatalogAction
	actions[0] = CreateAction("AUTOKV", "Splunk", 0, "", model.NONE, "", "", "", 0)
	actions[1] = CreateAction("EVAL", "Splunk", 0, "Splunk", "", "string", "", "", 0)
	actions[2] = CreateAction("LOOKUP", "Splunk", 0, "", "", "string", "", "", 0)

	result, err := getClient(t).CatalogService.CreateRule(CreateRule("_internal", "test_match", "splunk", "Splunk", actions[:]))

	require.Empty(t, err)
	assert.Equal(t, "_internal", result.Name)
	assert.Equal(t, "test_match", result.Match)
	assert.Equal(t, 3, len(result.Actions))
}

// creates a rule to post
func CreateRule(name string, match string, module string, owner string, actions []model.CatalogAction) model.Rule {
	return model.Rule{
		Name:    name,
		Match:   match,
		Module:  module,
		Owner:   owner,
		Actions: actions,
	}
}

// creates an action for rule to post
func CreateAction(kind model.CatalogActionKind, owner string, version int, field string, mode model.AutoMode, expression string, pattern string, alias string, limit int) model.CatalogAction {
	return model.CatalogAction{
		Kind:       kind,
		Owner:      owner,
		Version:    version,
		Field:      field,
		Mode:       mode,
		Expression: expression,
		Pattern:    pattern,
		Alias:      alias,
		Limit:      limit,
	}
}
