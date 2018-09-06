// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"net/url"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/service"
	"github.com/splunk/splunk-cloud-sdk-go/testutils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"fmt"
)

// Test Rule variables
var ruleName = "goSdkTestrRule1"
var ruleModule = "catalog"
var ruleMatch = "sourcetype::integration_test_match"
var owner = "splunk"

// Test Dataset variables
var datasetOwner = "Splunk"
var datasetCapabilities = "1101-00000:11010"
var datasetName = "integ_dataset_1000"
var externalKind = "kvcollection"
var externalName = "test_externalName"

func cleanupDatasets(t *testing.T) {
	client := getClient(t)
	result, err := client.CatalogService.GetDatasets()
	assert.Emptyf(t, err, "Error retrieving the datasets: %s", err)

	for _, item := range result {
		if item.Name == "main" || item.Name == "metrics" {
			continue
		}
		err = client.CatalogService.DeleteDataset(item.ID)
		assert.Emptyf(t, err, "Error deleting dataset: %s", err)
	}
}

func cleanupRules(t *testing.T) {
	client := getClient(t)
	result, err := client.CatalogService.GetRules()
	assert.Emptyf(t, err, "Error retrieving the rules: %s", err)

	for _, item := range result {
		err := client.CatalogService.DeleteRule(item.ID)
		assert.Emptyf(t, err, "Error deleting dataset: %s", err)
	}
}

// createDatastoreKVCollection - Helper function for creating a valid KV Collection in Catalog
func createLookupDataset(t *testing.T, namespaceName string, collectionName string, datasetOwner string, capabilities string, externalKind string, externalName string) (*model.DatasetInfo, error) {
	createLookupDatasetInfo := model.DatasetInfo{
		Name:         collectionName,
		Kind:         model.LOOKUP,
		Owner:        datasetOwner,
		Module:       namespaceName,
		Capabilities: capabilities,
		ExternalKind: externalKind,
		ExternalName: externalName,
	}

	datasetInfo, err := getClient(t).CatalogService.CreateDataset(createLookupDatasetInfo)
	require.NotNil(t, datasetInfo)
	require.IsType(t, model.DatasetInfo{}, *datasetInfo)
	require.Nil(t, err)
	require.Equal(t, model.LOOKUP, datasetInfo.Kind)

	return datasetInfo, err
}

func createKVCollectionDataset(t *testing.T, namespaceName string, collectionName string, datasetOwner string, capabilities string) (*model.DatasetInfo, error) {
	createKVCollectionDatasetInfo := model.DatasetInfo{
		Name:         collectionName,
		Kind:         model.KVCOLLECTION,
		Owner:        datasetOwner,
		Module:       namespaceName,
		Capabilities: capabilities,
	}

	datasetInfo, err := getClient(t).CatalogService.CreateDataset(createKVCollectionDatasetInfo)
	require.NotNil(t, datasetInfo)
	require.IsType(t, model.DatasetInfo{}, *datasetInfo)
	require.Nil(t, err)
	require.Equal(t, model.KVCOLLECTION, datasetInfo.Kind)

	return datasetInfo, err
}

// Test CreateDataset
func TestIntegrationCreateDataset(t *testing.T) {
	defer cleanupDatasets(t)

	createLookupDataset(t, testutils.TestNamespace, "integ_dataset_1000", datasetOwner, datasetCapabilities, externalKind, externalName)
	createLookupDataset(t, testutils.TestNamespace, "integ_dataset_2000", datasetOwner, datasetCapabilities, externalKind, externalName)
	createLookupDataset(t, testutils.TestNamespace, "integ_dataset_3000", datasetOwner, datasetCapabilities, externalKind, externalName)
}

// Test CreateDataset for 409 DatasetInfo already present error
func TestIntegrationCreateDatasetDataAlreadyPresentError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// create dataset
	createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)

	_, err := client.CatalogService.CreateDataset(
		model.DatasetInfo{
			Name:         testutils.TestCollection,
			Kind:         model.LOOKUP,
			Owner:        datasetOwner,
			Module:       testutils.TestNamespace,
			Capabilities: datasetCapabilities,
			ExternalKind: externalKind,
			ExternalName: externalName,
		})
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 409, "Expected error code 409")
}

// Test CreateDataset for 401 Unauthorized operation error
func TestIntegrationCreateDatasetUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	invalidClient := getInvalidClient(t)

	_, err := invalidClient.CatalogService.CreateDataset(
		model.DatasetInfo{Name: datasetName, Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: externalKind, ExternalName: externalName})
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test CreateDataset for 400 Invalid DatasetInfo error
func TestIntegrationCreateDatasetInvalidDatasetInfoError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	_, err := client.CatalogService.CreateDataset(
		model.DatasetInfo{Name: "integ_dataset_4000", Kind: model.LOOKUP})
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 400, "Expected error code 400")
}

// Test GetDatasets
func TestIntegrationGetAllDatasets(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	_, errOne := createLookupDataset(t, testutils.TestNamespace, "integ_dataset_1000", datasetOwner, datasetCapabilities, externalKind, externalName)
	assert.Emptyf(t, errOne, "Error creating dataset: %s", errOne)

	_, errTwo := createLookupDataset(t, testutils.TestNamespace, "integ_dataset_2000", datasetOwner, datasetCapabilities, externalKind, externalName)
	assert.Emptyf(t, errTwo, "Error creating dataset: %s", errTwo)

	_, errThree := createLookupDataset(t, testutils.TestNamespace, "integ_dataset_3000", datasetOwner, datasetCapabilities, externalKind, externalName)
	assert.Emptyf(t, errThree, "Error creating dataset: %s", errThree)

	datasets, err := client.CatalogService.GetDatasets()
	assert.Emptyf(t, err, "Error retrieving the datasets: %s", err)
	assert.NotNil(t, len(datasets))
}

// Test GetDatasets for 401 Unauthorized operation error
func TestIntegrationGetAllDatasetsUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	invalidClient := getInvalidClient(t)

	_, err := invalidClient.CatalogService.GetDatasets()
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test GetDataset by ID
func TestIntegrationGetDatasetByID(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	assert.Emptyf(t, err, "Error creating dataset: %s", err)

	datasetByID, err := client.CatalogService.GetDataset(dataset.ID)
	assert.Emptyf(t, err, "Error getting dataset: %s", err)
	assert.NotNil(t, datasetByID)
}

// Test GetDataset for 401 Unauthorized operation error
func TestIntegrationGetDatasetByIDUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: externalKind, ExternalName: externalName})
	assert.Emptyf(t, err, "Error creating dataset: %s", err)

	_, err = invalidClient.CatalogService.GetDataset(dataset.ID)
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test GetDataset for 404 DatasetInfo not found error
func TestIntegrationGetDatasetByIDDatasetNotFoundError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	_, err := client.CatalogService.GetDataset("123")
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test UpdateDataset
func TestIntegrationUpdateExistingDataset(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// create dataset
	updateVersion := 6
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)

	updatedDataset, err := client.CatalogService.UpdateDataset(model.PartialDatasetInfo{Version: updateVersion}, dataset.ID)
	assert.Emptyf(t, err, "Error updating dataset: %s", err)
	assert.NotNil(t, updatedDataset)
	assert.IsType(t, &(model.DatasetInfo{}), updatedDataset)

	// validate the update operation
	datasetByID, err := client.CatalogService.GetDataset(dataset.ID)
	require.Nil(t, err)
	assert.Emptyf(t, err, "Error retrieving dataset: %s", err)
	assert.Equal(t, updateVersion, datasetByID.Version)
	assert.NotNil(t, datasetByID.ID)
	assert.IsType(t, &(model.DatasetInfo{}), datasetByID)
}

// Test UpdateDataset for 404 DatasetInfo not found error
func TestIntegrationUpdateExistingDatasetDataNotFoundError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	_, err := client.CatalogService.UpdateDataset(model.PartialDatasetInfo{Name: "goSdkDataset6", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: externalKind, ExternalName: externalName, Version: 2}, "123")
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test DeleteDataset
func TestIntegrationDeleteDataset(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)

	err = client.CatalogService.DeleteDataset(dataset.ID)
	require.Nil(t, err)
	assert.Emptyf(t, err, "Error deleting dataset: %s", err)

	_, err = client.CatalogService.GetDataset(dataset.ID)
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test DeleteDataset for 401 Unauthorized operation error
func TestIntegrationDeleteDatasetUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	invalidClient := getInvalidClient(t)

	// create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	assert.NotNil(t, dataset.ID)

	err = invalidClient.CatalogService.DeleteDataset(dataset.ID)
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test DeleteDataset for 404 DatasetInfo not found error
func TestIntegrationDeleteDatasetDataNotFoundError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	err := client.CatalogService.DeleteDataset("123")
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// todo (Parul): 405 DatasetInfo cannot be deleted because of dependencies error case

// Test CreateRules
func TestIntegrationCreateRules(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, ruleMatch, rule.Match)

	_, err = client.CatalogService.CreateRule(model.Rule{Name: "anotherone", Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.Nil(t, err)

	_, err = client.CatalogService.CreateRule(model.Rule{Name: "thirdone", Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.Nil(t, err)
}

// Test CreateRule for 409 Rule already present error
func TestIntegrationCreateRuleDataAlreadyPresent(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, ruleMatch, rule.Match)

	_, err = client.CatalogService.CreateRule(model.Rule{ID: rule.ID, Name: ruleName, Module: ruleModule, Owner: owner, Match: ruleMatch})
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 409, "Expected error code 409")
}

// Test CreateRule for 401 Unauthorized operation error
func TestIntegrationCreateRuleUnauthorizedOperationError(t *testing.T) {
	defer cleanupRules(t)

	invalidClient := getInvalidClient(t)

	// create rule
	_, err := invalidClient.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Owner: owner, Match: ruleMatch})
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
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
	require.Nil(t, err)
	assert.Emptyf(t, err, "Error retrieving rules: %s", err)
	assert.NotNil(t, len(rules))
}

// Test GetRules for 401 Unauthorized operation error
func TestIntegrationGetAllRulesUnauthorizedOperationError(t *testing.T) {
	defer cleanupRules(t)

	invalidClient := getInvalidClient(t)

	_, err := invalidClient.CatalogService.GetRules()
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test GetRule By ID
func TestIntegrationGetRuleByID(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.NotNil(t, rule.ID)

	ruleByID, err := client.CatalogService.GetRule(rule.ID)
	assert.Emptyf(t, err, "Error retrieving rule by ID: %s", err)
	assert.NotNil(t, ruleByID)
}

// Test GetRules for 401 Unauthorized operation error
func TestIntegrationGetRuleByIDUnauthorizedOperationError(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.NotNil(t, rule.ID)

	_, err = invalidClient.CatalogService.GetRule(rule.ID)
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test GetRules for 404 Rule not found error
func TestIntegrationGetRuleByIDRuleNotFoundError(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	_, err := client.CatalogService.GetRule("123")
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test DeleteRule by ID
func TestIntegrationDeleteRule(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.NotNil(t, rule.ID)

	err = client.CatalogService.DeleteRule(rule.ID)
	assert.Emptyf(t, err, "Error deleting a rule by ID: %s", err)
}

// Test DeleteRule for 401 Unauthorized operation error
func TestIntegrationDeleteRuleByIDUnauthorizedOperationError(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.NotNil(t, rule.ID)

	err = invalidClient.CatalogService.DeleteRule(rule.ID)
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
	assert.True(t, err.(*util.HTTPError).Message == "401 Unauthorized", "Expected error message should be 401 Unauthorized")
}

// Test DeleteRule for 404 Rule not found error
func TestIntegrationDeleteRulebyIDRuleNotFoundError(t *testing.T) {
	defer cleanupRules(t)

	client := getClient(t)

	err := client.CatalogService.DeleteRule("123")
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test GetDatasetFields
func TestIntegrationGetDatasetFields(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)

	// create new fields in the dataset
	testField1 := model.Field{Name: "integ_test_field1", DatasetID: dataset.ID, DataType: "S", FieldType: "D", Prevalence: "A"}
	testField2 := model.Field{Name: "integ_test_field2", DatasetID: dataset.ID, DataType: "N", FieldType: "U", Prevalence: "S"}
	_, err = client.CatalogService.CreateDatasetField(dataset.ID, testField1)
	_, err = client.CatalogService.CreateDatasetField(dataset.ID, testField2)

	// Validate the creation of new dataset fields
	result, err := client.CatalogService.GetDatasetFields(dataset.ID, nil)
	require.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Emptyf(t, err, "Error retrieving dataset fields: %s", err)
	assert.Equal(t, 2, len(result))
}

// Test GetDatasetFields based on filter
func TestIntegrationGetDatasetFieldsOnFilter(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	require.Nil(t, err)
	require.Emptyf(t, err, "Error creating test Dataset: %s", err)

	// create new fields in the dataset
	testField1 := model.Field{Name: "integ_test_field1", DatasetID: dataset.ID, DataType: "S", FieldType: "D", Prevalence: "A"}
	testField2 := model.Field{Name: "integ_test_field2", DatasetID: dataset.ID, DataType: "N", FieldType: "U", Prevalence: "S"}
	_, err = client.CatalogService.CreateDatasetField(dataset.ID, testField1)
	_, err = client.CatalogService.CreateDatasetField(dataset.ID, testField2)

	filter := make(url.Values)
	filter.Add("filter", "name==\"integ_test_field2\"")

	// Validate the creation of new dataset fields
	result, err := client.CatalogService.GetDatasetFields(dataset.ID, nil)
	require.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Emptyf(t, err, "Error retrieving dataset fields: %s", err)
	assert.Equal(t, 2, len(result))
}

// Test PostDatasetField
func TestIntegrationPostDatasetField(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	assert.Emptyf(t, err, "Error creating dataset: %s", err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(dataset.ID, resultField.ID)
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Emptyf(t, err, "Error retrieving dataset field : %s", err)
}

// Test PatchDatasetField
func TestIntegrationPatchDatasetField(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(dataset.ID, resultField.ID)
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Emptyf(t, err, "Error retrieving dataset field: %s", err)

	// Update the existing dataset field
	resultField, err = client.CatalogService.UpdateDatasetField(dataset.ID, resultField.ID, model.Field{DataType: "O"})
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Equal(t, "integ_test_field", resultField.Name)
	assert.Equal(t, model.OBJECTID, resultField.DataType)
	assert.Equal(t, model.DIMENSION, resultField.FieldType)
	assert.Equal(t, model.ALL, resultField.Prevalence)
	assert.Emptyf(t, err, "Error updating dataset field: %s", err)
}

// Test DeleteDatasetField
func TestIntegrationDeleteDatasetField(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Delete dataset field
	err = client.CatalogService.DeleteDatasetField(dataset.ID, resultField.ID)
	require.Nil(t, err)
	assert.Emptyf(t, err, "Error deleting dataset field: %s", err)

	// Validate the deletion of the dataset field
	result, err := client.CatalogService.GetDatasetField(dataset.ID, resultField.ID)
	require.NotNil(t, err)
	assert.Empty(t, result)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404)
}

// Test PostDatasetField for 401 error
func TestIntegrationPostDatasetFieldUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	invalidClient := getInvalidClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)

	// Create a new dataset field
	testField := model.Field{Name: "integ_test_field", DatasetID: dataset.ID, DataType: "N", FieldType: "U", Prevalence: "S"}
	resultField, err := invalidClient.CatalogService.CreateDatasetField(dataset.ID, testField)
	require.NotNil(t, err)
	assert.Empty(t, resultField)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
}

// Test PostDatasetField for 409 error
func TestIntegrationPostDatasetFieldDataAlreadyPresent(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)

	// Create a new dataset field
	PostDatasetField(dataset, client, t)

	// Post an already created dataset field
	duplicateTestField := model.Field{Name: "integ_test_field", DataType: "S", FieldType: "D", Prevalence: "A"}
	resultField, err := client.CatalogService.CreateDatasetField(dataset.ID, duplicateTestField)
	require.NotNil(t, err)
	assert.Empty(t, resultField)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 409, "Expected error code 409")
}

// Test PostDatasetField for 500 error
func TestIntegrationPostDatasetFieldInvalidDataFormat(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)

	// Create a new dataset field
	testField := model.Field{Name: "integ_test_field"}
	resultField, err := client.CatalogService.CreateDatasetField(dataset.ID, testField)
	require.NotNil(t, err)
	assert.Empty(t, resultField)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 500, "Expected error code 500")
}

// Test GetDatasetFields for 401 Unauthorized operation error
func TestIntegrationGetDatasetFieldsUnauthorizedOperation(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)

	// Create new fields in the dataset
	PostDatasetField(dataset, client, t)

	// Validate the creation of new dataset fields
	result, err := invalidClient.CatalogService.GetDatasetFields(dataset.ID, nil)
	require.NotNil(t, err)
	assert.Empty(t, result)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
}

// Test PatchDatasetField for 401 error
func TestIntegrationPatchDatasetFieldUnauthorizedOperation(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(dataset.ID, resultField.ID)
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Emptyf(t, err, "Error retrieving dataset field: %s", err)

	// Update the existing dataset field
	resultField, err = invalidClient.CatalogService.UpdateDatasetField(dataset.ID, resultField.ID, model.Field{DataType: "O"})
	require.NotNil(t, err)
	assert.Empty(t, resultField)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
}

// Test PatchDatasetField for 404 error
func TestIntegrationPatchDatasetFieldDataNotFound(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Ceate dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(dataset.ID, resultField.ID)
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Emptyf(t, err, "Error retrieving dataset field: %s", err)

	// Update the existing dataset field
	resultField, err = client.CatalogService.UpdateDatasetField(dataset.ID, "123", model.Field{DataType: "O"})
	require.NotNil(t, err)
	assert.Empty(t, resultField)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test DeleteDatasetField for 401 error
func TestIntegrationDeleteDatasetFieldUnauthorizedOperation(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)
	invalidClient := getInvalidClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Delete dataset field
	err = invalidClient.CatalogService.DeleteDatasetField(dataset.ID, resultField.ID)
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
}

// Test DeleteDatasetField for 404 error
func TestIntegrationDeleteDatasetFieldDataNotFound(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)

	// Delete dataset field
	err = client.CatalogService.DeleteDatasetField(dataset.ID, "123")
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test rule actions endpoints
func TestRuleActions(t *testing.T) {
	defer cleanupDatasets(t)

	client := getClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(model.DatasetInfo{Name: "integ_dataset_1000", Kind: model.LOOKUP, Owner: datasetOwner, Capabilities: datasetCapabilities, ExternalKind: "kvcollection", ExternalName: "test_externalName"})
	require.Nil(t, err)
	defer client.CatalogService.DeleteDataset("integ_dataset_1000")

	// create new fields in the dataset
	testField1 := model.Field{Name: "integ_test_field1", DatasetID: dataset.ID, DataType: "S", FieldType: "D", Prevalence: "A"}
	field, err := client.CatalogService.CreateDatasetField(dataset.ID, testField1)
	require.Nil(t, err)
	defer client.CatalogService.DeleteDatasetField(dataset.ID, field.ID)

	// Create rule and rule action
	rule, err := client.CatalogService.CreateRule(model.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	defer client.CatalogService.DeleteRule(fmt.Sprintf("%v.%v", ruleModule, ruleName))
	require.Nil(t, err)

	// Create rule action
	action1, err := client.CatalogService.CreateRuleAction(rule.ID, *model.NewAliasAction(field.Name, "myfieldalias", "", 1))
	require.Nil(t, err)
	defer client.CatalogService.DeleteRuleAction(rule.ID, action1.ID)

	action2, err := client.CatalogService.CreateRuleAction(rule.ID, *model.NewAutoKVAction("mymode", "owner1",  1))
	require.Nil(t, err)
	defer client.CatalogService.DeleteRuleAction(rule.ID, action2.ID)

	action3, err := client.CatalogService.CreateRuleAction(rule.ID, *model.NewEvalAction(field.Name, "some expression", "", 1))
	require.Nil(t, err)
	defer client.CatalogService.DeleteRuleAction(rule.ID, action3.ID)

	action4, err := client.CatalogService.CreateRuleAction(rule.ID, *model.NewLookupAction( "myexpression2", "", 1))
	require.Nil(t, err)
	defer client.CatalogService.DeleteRuleAction(rule.ID, action4.ID)

	//Get rule actions
	actions, err := client.CatalogService.GetRuleActions(rule.ID)
	require.NotNil(t, actions)
	assert.Equal(t, 4, len(actions))

	action5, err := client.CatalogService.GetRuleAction(rule.ID, actions[0].ID)
	require.NotNil(t, action5)

	// Delete action
	err = client.CatalogService.DeleteRuleAction(rule.ID, action1.ID)
	require.Nil(t, err)
	err = client.CatalogService.DeleteRuleAction(rule.ID, action2.ID)
	require.Nil(t, err)
	err = client.CatalogService.DeleteRuleAction(rule.ID, action3.ID)
	require.Nil(t, err)
	err = client.CatalogService.DeleteRuleAction(rule.ID, action4.ID)
	require.Nil(t, err)

	// Delete rule with resource name
	err = client.CatalogService.DeleteRule(fmt.Sprintf("%v.%v", ruleModule, ruleName))
	require.Nil(t, err)

	// Delete dataset field
	err = client.CatalogService.DeleteDatasetField(dataset.ID, field.ID)
	require.Nil(t, err)

	// Delete dataset
	err = client.CatalogService.DeleteDataset(dataset.ID)
	require.Nil(t, err)
}


func PostDatasetField(dataset *model.DatasetInfo, client *service.Client, t *testing.T) *model.Field {
	testField := model.Field{Name: "integ_test_field", DatasetID: dataset.ID, DataType: "S", FieldType: "D", Prevalence: "A"}

	resultField, err := client.CatalogService.CreateDatasetField(dataset.ID, testField)
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Equal(t, "integ_test_field", resultField.Name)
	assert.Equal(t, model.STRING, resultField.DataType)
	assert.Equal(t, model.DIMENSION, resultField.FieldType)
	assert.Equal(t, model.ALL, resultField.Prevalence)
	assert.Emptyf(t, err, "Error creating dataset field: %s", err)

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
