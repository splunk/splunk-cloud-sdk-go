package playgroundintegration

import (
	"testing"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
	"net/url"
)

// Test Rule variables
var ruleName = "goSdkTestrRule1"
var ruleModule = "catalog"
var ruleMatch = "host::integration_test_match"
var owner = "splunk"

// Test Dataset variables
var datasetOwner = "Splunk"
var datasetCapabilities = "1101-00000:11010"
var datasetName = "integ_dataset_1000"

func cleanupDatasets(t *testing.T) {
	client := getClient(t)
	result, err := client.CatalogService.GetDatasets()
	assert.Nil(t, err)

	for _, item := range result {
		err = client.CatalogService.DeleteDataset(item.ID)
		assert.Nil(t, err)
	}
}

func cleanupRules(t *testing.T) {
	client := getClient(t)
	result, err := client.CatalogService.GetRules()
	assert.Nil(t, err)

	for _, item := range result {
		err := client.CatalogService.DeleteRule(item.ID)
		assert.Nil(t, err)
	}
}

// Test CreateDataset
func TestIntegrationCreateDataset(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: datasetName, Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	assert.Nil(t, err)
	assert.Equal(t, datasetName, dataset.Name)
	assert.Equal(t, model.LOOKUP, dataset.Kind)
	_, err = client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "integ_dataset_2000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	assert.Nil(t, err)
	_, err = client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "integ_dataset_3000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	assert.Nil(t, err)
}

// Test CreateDataset for 409 DatasetInfo already present error
func TestIntegrationCreateDatasetDataAlreadyPresentError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: datasetName, Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	_, err = client.CatalogService.CreateDataset(
		model.DatasetInfo{ID: dataset.ID, Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 409, "Expected error code 409")
}

// Test CreateDataset for 401 Unauthorized operation error
func TestIntegrationCreateDatasetUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	invalidClient := getInvalidClient(t)

	_, err := invalidClient.CatalogService.CreateDataset(
		model.DatasetInfo{Name: datasetName, Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test CreateDataset for 400 Invalid DatasetInfo error
func TestIntegrationCreateDatasetInvalidDatasetInfoError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	_, err := client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "integ_dataset_4000", Kind: model.LOOKUP})
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 400, "Expected error code 400")
}

// Test GetDatasets
func TestIntegrationGetAllDatasets(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// create dataset
	_, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	_, err = client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_2000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	_, err = client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_3000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	datasets, err := client.CatalogService.GetDatasets()
	assert.Nil(t, err)
	assert.NotNil(t, len(datasets))
}

// Test GetDatasets for 401 Unauthorized operation error
func TestIntegrationGetAllDatasetsUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	invalidClient := getInvalidClient(t)

	_, err := invalidClient.CatalogService.GetDatasets()
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test GetDataset by ID
func TestIntegrationGetDatasetByID(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	datasetByID, err := client.CatalogService.GetDataset(dataset.ID)
	assert.Nil(t, err)
	assert.NotNil(t, datasetByID)
}

// Test GetDataset for 401 Unauthorized operation error
func TestIntegrationGetDatasetByIDUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	_, err = invalidClient.CatalogService.GetDataset(dataset.ID)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test GetDataset for 404 DatasetInfo not found error
func TestIntegrationGetDatasetByIDDatasetNotFoundError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	_, err := client.CatalogService.GetDataset("123")
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
}

// Test UpdateDataset
func TestIntegrationUpdateExistingDataset(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// create dataset
	updateVersion := 6
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	updatedDataset, err := client.CatalogService.UpdateDataset(model.PartialDatasetInfo{Version: updateVersion}, dataset.ID)
	assert.Nil(t, err)
	assert.NotNil(t, updatedDataset)

	// validate the update operation
	datasetByID, err := client.CatalogService.GetDataset(dataset.ID)
	assert.Nil(t, err)
	assert.Equal(t, updateVersion, datasetByID.Version)
}

// Test UpdateDataset for 404 DatasetInfo not found error
func TestIntegrationUpdateExistingDatasetDataNotFoundError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	_, err := client.CatalogService.UpdateDataset(model.PartialDatasetInfo{Name: "goSdkDataset6", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName", Version: 2}, "123")
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
}

// Test DeleteDataset
func TestIntegrationDeleteDataset(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	assert.NotNil(t, dataset.ID)

	err = client.CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)

	_, err = client.CatalogService.GetDataset(dataset.ID)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
}

// Test DeleteDataset for 401 Unauthorized operation error
func TestIntegrationDeleteDatasetUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	assert.NotNil(t, dataset.ID)

	err = invalidClient.CatalogService.DeleteDataset(dataset.ID)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test DeleteDataset for 404 DatasetInfo not found error
func TestIntegrationDeleteDatasetDataNotFoundError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	err := client.CatalogService.DeleteDataset("123")
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
}

// todo (Parul): 405 DatasetInfo cannot be deleted because of dependencies error case

// Test CreateRules
func TestIntegrationCreateRules(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(
		model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.Nil(t, err)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, ruleMatch, rule.Match)

	_, err = client.CatalogService.CreateRule(
		model.Rule{Name: "anotherone", Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.Nil(t, err)

	_, err = client.CatalogService.CreateRule(
		model.Rule{Name: "thirdone", Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.Nil(t, err)
}

// Test CreateRule for 409 Rule already present error
func TestIntegrationCreateRuleDataAlreadyPresent(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(
		model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.Nil(t, err)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, ruleMatch, rule.Match)

	_, err = client.CatalogService.CreateRule(model.Rule{ID: rule.ID, Name: ruleName, Module: ruleModule, Owner: owner, Match: ruleMatch})
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 409, "Expected error code 409")
}

// Test CreateRule for 401 Unauthorized operation error
func TestIntegrationCreateRuleUnauthorizedOperationError(t *testing.T) {
	defer cleanupRules(t)

	invalidClient := getInvalidClient(t)

	// create rule
	_, err := invalidClient.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Owner: owner, Match: ruleMatch})
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test GetRules
func TestIntegrationGetAllRules(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	// create rule
	_, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	_, err = client.CatalogService.CreateRule(model.Rule{Name: "anotherone", Module: ruleModule, Match: ruleMatch, Owner: owner})
	_, err = client.CatalogService.CreateRule(model.Rule{Name: "thirdone", Module: ruleModule, Match: ruleMatch, Owner: owner})

	rules, err := client.CatalogService.GetRules()
	assert.Nil(t, err)
	assert.NotNil(t, len(rules))
}

// Test GetRules for 401 Unauthorized operation error
func TestIntegrationGetAllRulesUnauthorizedOperationError(t *testing.T) {
	defer cleanupRules(t)

	invalidClient := getInvalidClient(t)

	_, err := invalidClient.CatalogService.GetRules()
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test GetRule By ID
func TestIntegrationGetRuleByID(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.NotNil(t, rule.ID)

	ruleByID, err := client.CatalogService.GetRule(rule.ID)
	assert.Nil(t, err)
	assert.NotNil(t, ruleByID)
}

// Test GetRules for 401 Unauthorized operation error
func TestIntegrationGetRuleByIDUnauthorizedOperationError(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.NotNil(t, rule.ID)

	_, err = invalidClient.CatalogService.GetRule(rule.ID)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test GetRules for 404 Rule not found error
func TestIntegrationGetRuleByIDRuleNotFoundError(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	_, err := client.CatalogService.GetRule("123")
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
}

// Test DeleteRule by ID
func TestIntegrationDeleteRule(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.NotNil(t, rule.ID)

	err = client.CatalogService.DeleteRule(rule.ID)
	assert.Nil(t, err)
}

// Test DeleteRule for 401 Unauthorized operation error
func TestIntegrationDeleteRuleByIDUnauthorizedOperationError(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.NotNil(t, rule.ID)

	err = invalidClient.CatalogService.DeleteRule(rule.ID)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test DeleteRule for 404 Rule not found error
func TestIntegrationDeleteRulebyIDRuleNotFoundError(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	err := client.CatalogService.DeleteRule("123")
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
}

// Test GetDatasetFields
func TestIntegrationGetDatasetFields(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// create new fields in the dataset
	testField1 := model.Field{Name: "integ_test_field1", DatasetID: dataset.ID, DataType: "S", FieldType: "D", Prevalence: "A"}
	testField2 := model.Field{Name: "integ_test_field2", DatasetID: dataset.ID, DataType: "N", FieldType: "U", Prevalence: "S"}
	_, err = client.CatalogService.PostDatasetField(dataset.ID, testField1)
	_, err = client.CatalogService.PostDatasetField(dataset.ID, testField2)

	// Validate the creation of new dataset fields
	result, err := client.CatalogService.GetDatasetFields(dataset.ID, nil)
	assert.NotEmpty(t, result)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
}

// Test GetDatasetFields based on filter
func TestIntegrationGetDatasetFieldsOnFilter(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// create new fields in the dataset
	testField1 := model.Field{Name: "integ_test_field1", DatasetID: dataset.ID, DataType: "S", FieldType: "D", Prevalence: "A"}
	testField2 := model.Field{Name: "integ_test_field2", DatasetID: dataset.ID, DataType: "N", FieldType: "U", Prevalence: "S"}
	_, err = client.CatalogService.PostDatasetField(dataset.ID, testField1)
	_, err = client.CatalogService.PostDatasetField(dataset.ID, testField2)

	filter := make(url.Values)
	filter.Add("filter", "name==\"integ_test_field2\"")

	// Validate the creation of new dataset fields
	result, err := client.CatalogService.GetDatasetFields(dataset.ID, filter)
	assert.NotEmpty(t, result)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, result[0].Name, testField2.Name)
}

// Test PostDatasetField
func TestIntegrationPostDatasetField(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(dataset.ID, resultField.ID)
	assert.NotEmpty(t, resultField)
	assert.Nil(t, err)
}

// Test PatchDatasetField
func TestIntegrationPatchDatasetField(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(dataset.ID, resultField.ID)
	assert.NotEmpty(t, resultField)
	assert.Nil(t, err)

	// Update the existing dataset field
	resultField, err = client.CatalogService.PatchDatasetField(dataset.ID, resultField.ID, model.Field{DataType: "O"})
	assert.NotEmpty(t, resultField)
	assert.Equal(t, "integ_test_field", resultField.Name)
	assert.Equal(t, model.OBJECTID, resultField.DataType)
	assert.Equal(t, model.DIMENSION, resultField.FieldType)
	assert.Equal(t, model.ALL, resultField.Prevalence)
	assert.Nil(t, err)
}

// Test DeleteDatasetField
func TestIntegrationDeleteDatasetField(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Delete dataset field
	err = client.CatalogService.DeleteDatasetField(dataset.ID, resultField.ID)
	assert.Nil(t, err)

	// Validate the deletion of the dataset field
	result, err := client.CatalogService.GetDatasetField(dataset.ID, resultField.ID)
	assert.Empty(t, result)
	assert.True(t, err.(*util.HTTPError).Status == 404)
	assert.NotNil(t, err)
}

// Test PostDatasetField for 401 error
func TestIntegrationPostDatasetFieldUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// Create a new dataset field
	testField := model.Field{Name: "integ_test_field", DatasetID: dataset.ID, DataType: "N", FieldType: "U", Prevalence: "S"}
	resultField, err := invalidClient.CatalogService.PostDatasetField(dataset.ID, testField)
	assert.Empty(t, resultField)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
}

// Test PostDatasetField for 409 error
func TestIntegrationPostDatasetFieldDataAlreadyPresent(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// Create a new dataset field
	PostDatasetField(dataset, client, t)

	// Post an already created dataset field
	duplicateTestField := model.Field{Name: "integ_test_field", DataType: "S", FieldType: "D", Prevalence: "A"}
	resultField, err := client.CatalogService.PostDatasetField(dataset.ID, duplicateTestField)
	assert.Empty(t, resultField)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 409, "Expected error code 409")
}

// Test PostDatasetField for 500 error
func TestIntegrationPostDatasetFieldInvalidDataFormat(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// Create a new dataset field
	testField := model.Field{Name: "integ_test_field"}
	resultField, err := client.CatalogService.PostDatasetField(dataset.ID, testField)
	assert.Empty(t, resultField)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 500, "Expected error code 500")
}

// Test GetDatasetFields for 401 Unauthorized operation error
func TestIntegrationGetDatasetFieldsUnauthorizedOperation(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// Create new fields in the dataset
	PostDatasetField(dataset, client, t)

	// Validate the creation of new dataset fields
	result, err := invalidClient.CatalogService.GetDatasetFields(dataset.ID, nil)
	assert.Empty(t, result)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
}

// Test PatchDatasetField for 401 error
func TestIntegrationPatchDatasetFieldUnauthorizedOperation(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(dataset.ID, resultField.ID)
	assert.NotEmpty(t, resultField)
	assert.Nil(t, err)

	// Update the existing dataset field
	resultField, err = invalidClient.CatalogService.PatchDatasetField(dataset.ID, resultField.ID, model.Field{DataType: "O"})
	assert.Empty(t, resultField)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
}

// Test PatchDatasetField for 404 error
func TestIntegrationPatchDatasetFieldDataNotFound(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Ceate dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(dataset.ID, resultField.ID)
	assert.NotEmpty(t, resultField)
	assert.Nil(t, err)

	// Update the existing dataset field
	resultField, err = client.CatalogService.PatchDatasetField(dataset.ID, "123", model.Field{DataType: "O"})
	assert.Empty(t, resultField)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
}

// Test DeleteDatasetField for 401 error
func TestIntegrationDeleteDatasetFieldUnauthorizedOperation(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Delete dataset field
	err = invalidClient.CatalogService.DeleteDatasetField(dataset.ID, resultField.ID)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 401, "Expected error code 401")
}

// Test DeleteDatasetField for 404 error
func TestIntegrationDeleteDatasetFieldDataNotFound(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})

	// Delete dataset field
	err = client.CatalogService.DeleteDatasetField(dataset.ID, "123")
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
}

func PostDatasetField(dataset *model.DatasetInfo, client *service.Client, t *testing.T) *model.Field {
	testField := model.Field{Name: "integ_test_field", DatasetID: dataset.ID, DataType: "S", FieldType: "D", Prevalence: "A"}

	resultField, err := client.CatalogService.PostDatasetField(dataset.ID, testField)
	assert.NotEmpty(t, resultField)
	assert.Equal(t, "integ_test_field", resultField.Name)
	assert.Equal(t, model.STRING, resultField.DataType)
	assert.Equal(t, model.DIMENSION, resultField.FieldType)
	assert.Equal(t, model.ALL, resultField.Prevalence)
	assert.Nil(t, err)

	return resultField
}

/*// Currently unable to generate a bad rule
func TestIntegrationCreateRuleInvalidRuleError(t *testing.T)  {
	defer cleanupRules(t)

	client := getClient()

	// testing CreateRule for 400 Invalid Rule error
	ruleName := "goSdkTestrRule1"
	_, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName})
	assert.NotNil(t, err)
   assert.True(t, err.(*util.HTTPError).Status == 400, "Expected error code 400")
}*/

// todo (Parul): 405 Rule cannot be deleted because of dependencies error case
