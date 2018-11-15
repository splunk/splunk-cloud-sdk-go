// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"net/url"
	"testing"

	"strings"

	"encoding/json"
	"fmt"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test Rule variables
var ruleName = "goSdkTestrRule1"
var ruleModule = "catalog"
var ruleMatch = "sourcetype::integration_test_match"
var owner = "splunk"

// Test Dataset variables
var datasetOwner = "Splunk"
var datasetCapabilities = "1101-00000:11010"
var datasetName1 = fmt.Sprintf("go_integ_dataset_1000_%d", timeSec)
var datasetName2 = fmt.Sprintf("go_integ_dataset_2000_%d", timeSec)
var datasetName3 = fmt.Sprintf("go_integ_dataset_3000_%d", timeSec)
var datasetName4 = fmt.Sprintf("go_integ_dataset_index_%d", timeSec)
var datasetName5 = fmt.Sprintf("go_integ_dataset_view_%d", timeSec)
var datasetName6 = fmt.Sprintf("go_integ_dataset_metric_%d", timeSec)
var externalKind = "kvcollection"
var externalName = "test_externalName"
var frozenTimePeriodInSecs = 8
var disabled = false
var searchString = "search index=main|stats count()"

func cleanupDatasets(t *testing.T) {
	client := getSdkClient(t)
	result, err := client.CatalogService.ListDatasets(nil)
	assert.Emptyf(t, err, "Error retrieving the datasets: %s", err)

	// deletes all the datasets with name containing a 'go' prefix
	for _, item := range result {
		if strings.HasPrefix(item.Name, "go") {
			err = client.CatalogService.DeleteDataset(item.ID)
			assert.Emptyf(t, err, "Error deleting dataset: %s", err)
		}
	}
}

func cleanupRules(t *testing.T) {
	client := getSdkClient(t)
	result, err := client.CatalogService.GetRules()
	assert.Emptyf(t, err, "Error retrieving the rules: %s", err)

	for _, item := range result {
		err := client.CatalogService.DeleteRule(item.ID)
		assert.Emptyf(t, err, "Error deleting dataset: %s", err)
	}
}

// createDatastoreKVCollection - Helper function for creating a valid KV Collection in Catalog
func createLookupDataset(t *testing.T, namespaceName string, collectionName string, datasetOwner string, capabilities string, externalKind string, externalName string) (catalog.Dataset, error) {
	createLookupDataset := catalog.LookupDataset{
		Name:         collectionName,
		Kind:         "lookup",
		Module:       namespaceName,
		Capabilities: capabilities,
		ExternalKind: &externalKind,
		ExternalName: &externalName,
	}

	datasetInfo, err := getSdkClient(t).CatalogService.CreateDataset(createLookupDataset)
	require.NotNil(t, datasetInfo)
	require.IsType(t, catalog.LookupDataset{}, datasetInfo)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(datasetInfo)
	require.Nil(t, err)
	require.Equal(t, "lookup", datasetMap["kind"].(string))

	return datasetInfo, err
}

func createLookupDatasets(t *testing.T) {
	_, errOne := createLookupDataset(t, testutils.TestNamespace, datasetName1, datasetOwner, datasetCapabilities, externalKind, externalName)
	assert.Emptyf(t, errOne, "Error creating dataset: %s", errOne)

	_, errTwo := createLookupDataset(t, testutils.TestNamespace, datasetName2, datasetOwner, datasetCapabilities, externalKind, externalName)
	assert.Emptyf(t, errTwo, "Error creating dataset: %s", errTwo)

	_, errThree := createLookupDataset(t, testutils.TestNamespace, datasetName3, datasetOwner, datasetCapabilities, externalKind, externalName)
	assert.Emptyf(t, errThree, "Error creating dataset: %s", errThree)
}

func createKVCollectionDataset(t *testing.T, namespaceName string, collectionName string, datasetOwner string, capabilities string) (catalog.Dataset, error) {
	createKVCollectionDatasetInfo := catalog.KVCollectionDataset{
		Name:         collectionName,
		Kind:         "kvcollection",
		Module:       namespaceName,
		Capabilities: capabilities,
	}

	datasetInfo, err := getSdkClient(t).CatalogService.CreateDataset(createKVCollectionDatasetInfo)
	require.NotNil(t, datasetInfo)
	require.IsType(t, catalog.KVCollectionDataset{}, datasetInfo)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(datasetInfo)
	require.Nil(t, err)
	require.Equal(t, "kvcollection", datasetMap["kind"].(string))

	return datasetInfo, err
}

func createMetricDataset(t *testing.T, namespaceName string, collectionName string, datasetOwner string, capabilities string, isDisabled bool) (catalog.Dataset, error) {
	createMetricDatasetInfo := catalog.MetricDataset{
		Name:         collectionName,
		Kind:         "metric",
		Module:       namespaceName,
		Capabilities: capabilities,
		Disabled:     &isDisabled,
	}

	datasetInfo, err := getSdkClient(t).CatalogService.CreateDataset(&createMetricDatasetInfo)
	require.NotNil(t, datasetInfo)
	require.IsType(t, catalog.MetricDataset{}, datasetInfo)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(datasetInfo)
	require.Nil(t, err)
	require.Equal(t, "metric", datasetMap["kind"].(string))

	return datasetInfo, err
}

func createViewDataset(t *testing.T, namespaceName string, collectionName string, datasetOwner string, capabilities string, search string) (catalog.Dataset, error) {
	createViewDatasetInfo := catalog.ViewDataset{
		Name:         collectionName,
		Kind:         "view",
		Module:       namespaceName,
		Capabilities: capabilities,
		Search:       &search,
	}

	datasetInfo, err := getSdkClient(t).CatalogService.CreateDataset(&createViewDatasetInfo)
	require.NotNil(t, datasetInfo)
	require.IsType(t, catalog.ViewDataset{}, datasetInfo)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(datasetInfo)
	require.Nil(t, err)
	require.Equal(t, "view", datasetMap["kind"].(string))

	return datasetInfo, err
}

func createIndexDataset(t *testing.T, namespaceName string, collectionName string, datasetOwner string, capabilities string, frozenTimePeriodInSecs int, isDisabled bool) (catalog.Dataset, error) {
	createIndexDatasetInfo := catalog.IndexDataset{
		Name:                   collectionName,
		Kind:                   "index",
		Module:                 namespaceName,
		Capabilities:           capabilities,
		FrozenTimePeriodInSecs: &frozenTimePeriodInSecs,
		Disabled:               &isDisabled,
	}

	datasetInfo, err := getSdkClient(t).CatalogService.CreateDataset(&createIndexDatasetInfo)
	require.NotNil(t, datasetInfo)
	require.IsType(t, catalog.IndexDataset{}, datasetInfo)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(datasetInfo)
	require.Nil(t, err)
	require.Equal(t, "index", datasetMap["kind"].(string))

	return datasetInfo, err
}

func createImportDataset(t *testing.T, namespaceName string, collectionName string, datasetOwner string, capabilities string, sourceName string, sourceModule string) (catalog.Dataset, error) {
	createImportDatasetInfo := catalog.ImportDataset{
		Name:         collectionName,
		Kind:         "import",
		Module:       namespaceName,
		Capabilities: capabilities,
		SourceName:   &sourceName,
		SourceModule: &sourceModule,
	}

	datasetInfo, err := getSdkClient(t).CatalogService.CreateDataset(&createImportDatasetInfo)
	require.NotNil(t, datasetInfo)
	require.IsType(t, catalog.ImportDataset{}, datasetInfo)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(datasetInfo)
	require.Nil(t, err)
	require.Equal(t, "import", datasetMap["kind"].(string))

	return datasetInfo, err
}

// Test CreateDataset
func TestIntegrationCreateDataset(t *testing.T) {
	defer cleanupDatasets(t)

	createLookupDataset(t, testutils.TestNamespace, datasetName1, datasetOwner, datasetCapabilities, externalKind, externalName)
	createLookupDataset(t, testutils.TestNamespace, datasetName2, datasetOwner, datasetCapabilities, externalKind, externalName)
	createLookupDataset(t, testutils.TestNamespace, datasetName3, datasetOwner, datasetCapabilities, externalKind, externalName)
	createIndexDataset(t, testutils.TestNamespace, datasetName4, datasetOwner, datasetCapabilities, frozenTimePeriodInSecs, disabled)
	createViewDataset(t, testutils.TestNamespace, datasetName5, datasetOwner, datasetCapabilities, searchString)
	createMetricDataset(t, testutils.TestNamespace, datasetName6, datasetOwner, datasetCapabilities, disabled)
}

// Test CreateDataset for 409 DatasetInfo already present error
func TestIntegrationCreateDatasetDataAlreadyPresentError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// create dataset
	createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)

	_, err := client.CatalogService.CreateDataset(
		&catalog.LookupDataset{
			Name:         testutils.TestCollection,
			Kind:         "lookup",
			Module:       testutils.TestNamespace,
			Capabilities: datasetCapabilities,
			ExternalKind: &externalKind,
			ExternalName: &externalName,
		})
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 409, "Expected error code 409")
}

// Test CreateDataset for 401 Unauthorized operation error
func TestIntegrationCreateDatasetUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	invalidClient := getInvalidClient(t)

	_, err := invalidClient.CatalogService.CreateDataset(
		&catalog.LookupDataset{Name: datasetName1, Kind: "lookup", Capabilities: datasetCapabilities, ExternalKind: &externalKind, ExternalName: &externalName})
	require.NotNil(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "Error validating request", err.(*util.HTTPError).Message)
}

// Test CreateDataset for 400 Invalid DatasetInfo error
func TestIntegrationCreateDatasetInvalidDatasetInfoError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	_, err := client.CatalogService.CreateDataset(
		&catalog.LookupDataset{Name: "go_integ_dataset_4000", Kind: "lookup"})
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 400, "Expected error code 400")
}

// Test GetDatasets
func TestIntegrationGetAllDatasets(t *testing.T) {
	defer cleanupDatasets(t)
	createLookupDatasets(t)

	datasets, err := getSdkClient(t).CatalogService.ListDatasets(nil)
	assert.Emptyf(t, err, "Error retrieving the datasets: %s", err)
	assert.NotNil(t, len(datasets))
}

// Test GetDatasets for 401 Unauthorized operation error
func TestIntegrationGetAllDatasetsUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	invalidClient := getInvalidClient(t)

	_, err := invalidClient.CatalogService.ListDatasets(nil)
	require.NotNil(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "Error validating request", err.(*util.HTTPError).Message)
}

// Test ListDatasetsNil
func TestListDatasetsNil(t *testing.T) {
	defer cleanupDatasets(t)
	createLookupDatasets(t)

	datasets, err := getSdkClient(t).CatalogService.ListDatasets(nil)
	assert.Emptyf(t, err, "Error retrieving the datasets: %s", err)
	assert.NotNil(t, len(datasets))
}

// Test TestListDatasetsFilter
func TestListDatasetsFilter(t *testing.T) {
	defer cleanupDatasets(t)
	createLookupDatasets(t)

	values := make(url.Values)
	values.Set("filter", "kind==\"kvcollection\"")

	datasets, err := getSdkClient(t).CatalogService.ListDatasets(values)
	assert.Emptyf(t, err, "Error retrieving the datasets: %s", err)
	assert.NotNil(t, len(datasets))
}

// Test TestListDatasetsCount
func TestListDatasetsCount(t *testing.T) {
	defer cleanupDatasets(t)
	createLookupDatasets(t)

	values := make(url.Values)
	values.Set("count", "1")

	datasets, err := getSdkClient(t).CatalogService.ListDatasets(values)
	assert.Emptyf(t, err, "Error retrieving the datasets: %s", err)
	assert.NotNil(t, len(datasets))
}

// Test TestListDatasetsOrderBy
func TestListDatasetsOrderBy(t *testing.T) {
	defer cleanupDatasets(t)
	createLookupDatasets(t)

	values := make(url.Values)
	values.Set("orderby", "id Descending")

	datasets, err := getSdkClient(t).CatalogService.ListDatasets(values)
	assert.Emptyf(t, err, "Error retrieving the datasets: %s", err)
	assert.NotNil(t, len(datasets))
}

// Test TestListDatasetsAll with filter, count, and orderby
func TestListDatasetsAll(t *testing.T) {
	defer cleanupDatasets(t)
	createLookupDatasets(t)

	values := make(url.Values)
	values.Set("filter", "kind==\"kvcollection\"")
	values.Set("count", "1")
	values.Set("orderby", "id Descending")

	datasets, err := getSdkClient(t).CatalogService.ListDatasets(values)
	assert.Emptyf(t, err, "Error retrieving the datasets: %s", err)
	assert.NotNil(t, len(datasets))
}

// Test GetDataset by ID
func TestIntegrationGetDatasetByID(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	assert.Emptyf(t, err, "Error creating dataset: %s", err)

	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	datasetByID, err := client.CatalogService.GetDataset(datasetMap["id"].(string))

	assert.Emptyf(t, err, "Error getting dataset: %s", err)
	assert.NotNil(t, datasetByID)
}

// Test GetDataset for 401 Unauthorized operation error
func TestIntegrationGetDatasetByIDUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)
	invalidClient := getInvalidClient(t)

	// create dataset
	dataset, err := client.CatalogService.CreateDataset(&catalog.LookupDataset{Name: datasetName1, Kind: "lookup", Capabilities: datasetCapabilities, ExternalKind: &externalKind, ExternalName: &externalName})
	assert.Emptyf(t, err, "Error creating dataset: %s", err)

	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	_, err = invalidClient.CatalogService.GetDataset(datasetMap["id"].(string))
	require.NotNil(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "Error validating request", err.(*util.HTTPError).Message)
}

// Test GetDataset for 404 DatasetInfo not found error
func TestIntegrationGetDatasetByIDDatasetNotFoundError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	_, err := client.CatalogService.GetDataset("123")
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test UpdateDataset
func TestIntegrationUpdateExistingDataset(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// create dataset
	updateVersion := 6
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	updateDatasetInfo := catalog.UpdateDataset{
		Version: &updateVersion,
	}
	updateLookupDataset := catalog.UpdateLookup{
		UpdateDataset: updateDatasetInfo,
	}
	updatedDataset, err := client.CatalogService.UpdateLookupDataset(&updateLookupDataset, datasetMap["id"].(string))
	assert.Emptyf(t, err, "Error updating dataset: %s", err)
	assert.NotNil(t, updatedDataset)
	assert.IsType(t, &(catalog.DatasetInfo{}), updatedDataset)

	// validate the update operation
	datasetByID, err := client.CatalogService.GetDataset(datasetMap["id"].(string))
	datasetByIDMap, err := convertDatasetInterfaceToMap(datasetByID)
	require.Nil(t, err)
	assert.Emptyf(t, err, "Error retrieving dataset: %s", err)
	assert.Equal(t, float64(updateVersion), datasetByIDMap["version"].(float64))
	assert.NotNil(t, datasetByIDMap["id"].(string))
	assert.IsType(t, catalog.LookupDataset{}, datasetByID)
}

// Test UpdateDataset for 404 DatasetInfo not found error
func TestIntegrationUpdateExistingDatasetDataNotFoundError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)
	version := 2

	updateDatasetInfo := catalog.UpdateDataset{
		Name:         "goSdkDataset6",
		Kind:         catalog.Lookup,
		Owner:        datasetOwner,
		Capabilities: datasetCapabilities,
		Version:      &version,
	}

	updateLookupDataset := catalog.UpdateLookup{
		UpdateDataset: updateDatasetInfo,
		ExternalKind:  externalKind,
		ExternalName:  externalName,
	}

	_, err := client.CatalogService.UpdateLookupDataset(&updateLookupDataset, "123")
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test DeleteDataset
func TestIntegrationDeleteDataset(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)

	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	err = client.CatalogService.DeleteDataset(datasetMap["id"].(string))
	require.Nil(t, err)
	assert.Emptyf(t, err, "Error deleting dataset: %s", err)

	_, err = client.CatalogService.GetDataset(datasetMap["id"].(string))
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test DeleteDataset for 401 Unauthorized operation error
func TestIntegrationDeleteDatasetUnauthorizedOperationError(t *testing.T) {
	defer cleanupDatasets(t)

	invalidClient := getInvalidClient(t)

	// create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)
	assert.NotNil(t, datasetMap["id"].(string))

	err = invalidClient.CatalogService.DeleteDataset(datasetMap["id"].(string))
	require.NotNil(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "Error validating request", err.(*util.HTTPError).Message)
}

// Test DeleteDataset for 404 DatasetInfo not found error
func TestIntegrationDeleteDatasetDataNotFoundError(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	err := client.CatalogService.DeleteDataset("123")
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// todo (Parul): 405 DatasetInfo cannot be deleted because of dependencies error case

// Test CreateRules
func TestIntegrationCreateRules(t *testing.T) {
	defer cleanupRules(t)

	client := getSdkClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, ruleMatch, rule.Match)

	_, err = client.CatalogService.CreateRule(catalog.Rule{Name: "anotherone", Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.Nil(t, err)

	_, err = client.CatalogService.CreateRule(catalog.Rule{Name: "thirdone", Module: ruleModule, Match: ruleMatch, Owner: owner})
	assert.Nil(t, err)
}

// Test CreateRule for 409 Rule already present error
func TestIntegrationCreateRuleDataAlreadyPresent(t *testing.T) {
	defer cleanupRules(t)

	client := getSdkClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.Equal(t, ruleName, rule.Name)
	assert.Equal(t, ruleMatch, rule.Match)

	_, err = client.CatalogService.CreateRule(catalog.Rule{ID: rule.ID, Name: ruleName, Module: ruleModule, Owner: owner, Match: ruleMatch})
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 409, "Expected error code 409")
}

// Test CreateRule for 401 Unauthorized operation error
func TestIntegrationCreateRuleUnauthorizedOperationError(t *testing.T) {
	defer cleanupRules(t)

	invalidClient := getInvalidClient(t)

	// create rule
	_, err := invalidClient.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Owner: owner, Match: ruleMatch})
	require.NotNil(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "Error validating request", err.(*util.HTTPError).Message)
}

// Test GetRules
func TestIntegrationGetAllRules(t *testing.T) {
	defer cleanupRules(t)

	client := getSdkClient(t)

	// create rule
	_, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	_, err = client.CatalogService.CreateRule(catalog.Rule{Name: "anotherone", Module: ruleModule, Match: ruleMatch, Owner: owner})
	_, err = client.CatalogService.CreateRule(catalog.Rule{Name: "thirdone", Module: ruleModule, Match: ruleMatch, Owner: owner})

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
	assert.Equal(t, 401, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "Error validating request", err.(*util.HTTPError).Message)
}

// Test GetRule By ID
func TestIntegrationGetRuleByID(t *testing.T) {
	defer cleanupRules(t)

	client := getSdkClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.NotNil(t, rule.ID)

	ruleByID, err := client.CatalogService.GetRule(rule.ID)
	assert.Emptyf(t, err, "Error retrieving rule by ID: %s", err)
	assert.NotNil(t, ruleByID)
}

// Test GetRules for 401 Unauthorized operation error
func TestIntegrationGetRuleByIDUnauthorizedOperationError(t *testing.T) {
	defer cleanupRules(t)

	client := getSdkClient(t)
	invalidClient := getInvalidClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.NotNil(t, rule.ID)

	_, err = invalidClient.CatalogService.GetRule(rule.ID)
	require.NotNil(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "Error validating request", err.(*util.HTTPError).Message)
}

// Test GetRules for 404 Rule not found error
func TestIntegrationGetRuleByIDRuleNotFoundError(t *testing.T) {
	defer cleanupRules(t)

	client := getSdkClient(t)

	_, err := client.CatalogService.GetRule("123")
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test DeleteRule by ID
func TestIntegrationDeleteRule(t *testing.T) {
	defer cleanupRules(t)

	client := getSdkClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.NotNil(t, rule.ID)

	err = client.CatalogService.DeleteRule(rule.ID)
	assert.Emptyf(t, err, "Error deleting a rule by ID: %s", err)
}

// Test DeleteRule for 401 Unauthorized operation error
func TestIntegrationDeleteRuleByIDUnauthorizedOperationError(t *testing.T) {
	defer cleanupRules(t)

	client := getSdkClient(t)
	invalidClient := getInvalidClient(t)

	// create rule
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	require.Nil(t, err)
	assert.NotNil(t, rule.ID)

	err = invalidClient.CatalogService.DeleteRule(rule.ID)
	require.NotNil(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "Error validating request", err.(*util.HTTPError).Message)
}

// Test DeleteRule for 404 Rule not found error
func TestIntegrationDeleteRulebyIDRuleNotFoundError(t *testing.T) {
	defer cleanupRules(t)

	client := getSdkClient(t)

	err := client.CatalogService.DeleteRule("123")
	require.NotNil(t, err)
	assert.Equal(t, 404, err.(*util.HTTPError).HTTPStatusCode)
}

// Test GetDatasetFields
func TestIntegrationGetDatasetFields(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// create new fields in the dataset
	testField1 := catalog.Field{Name: "integ_test_field1", DataType: "S", FieldType: "D", Prevalence: "A"}
	testField2 := catalog.Field{Name: "integ_test_field2", DataType: "N", FieldType: "U", Prevalence: "S"}
	_, err = client.CatalogService.CreateDatasetField(datasetMap["id"].(string), &testField1)
	_, err = client.CatalogService.CreateDatasetField(datasetMap["id"].(string), &testField2)

	// Validate the creation of new dataset fields
	result, err := client.CatalogService.GetDatasetFields(datasetMap["id"].(string), nil)
	require.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Emptyf(t, err, "Error retrieving dataset fields: %s", err)
	assert.Equal(t, 2, len(result))
}

// Test GetDatasetFields based on filter
func TestIntegrationGetDatasetFieldsOnFilter(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(&catalog.LookupDataset{Name: datasetName1, Kind: "lookup", Capabilities: datasetCapabilities, ExternalKind: &externalKind, ExternalName: &externalName})
	require.Nil(t, err)
	require.Emptyf(t, err, "Error creating test Dataset: %s", err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// create new fields in the dataset
	testField1 := catalog.Field{Name: "integ_test_field1", DataType: "S", FieldType: "D", Prevalence: "A"}
	testField2 := catalog.Field{Name: "integ_test_field2", DataType: "N", FieldType: "U", Prevalence: "S"}
	_, err = client.CatalogService.CreateDatasetField(datasetMap["id"].(string), &testField1)
	_, err = client.CatalogService.CreateDatasetField(datasetMap["id"].(string), &testField2)

	filter := make(url.Values)
	filter.Add("filter", "name==\"integ_test_field2\"")

	// Validate the creation of new dataset fields
	result, err := client.CatalogService.GetDatasetFields(datasetMap["id"].(string), nil)
	require.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.Emptyf(t, err, "Error retrieving dataset fields: %s", err)
	assert.Equal(t, 2, len(result))
}

// Test PostDatasetField
func TestIntegrationPostDatasetField(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	assert.Emptyf(t, err, "Error creating dataset: %s", err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(datasetMap["id"].(string), resultField.ID)
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Emptyf(t, err, "Error retrieving dataset field : %s", err)
}

// Test PatchDatasetField
func TestIntegrationPatchDatasetField(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(datasetMap["id"].(string), resultField.ID)
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Emptyf(t, err, "Error retrieving dataset field: %s", err)

	// Update the existing dataset field
	resultField, err = client.CatalogService.UpdateDatasetField(datasetMap["id"].(string), resultField.ID, &catalog.Field{DataType: "O"})
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Equal(t, "integ_test_field", resultField.Name)
	assert.Equal(t, catalog.ObjectID, resultField.DataType)
	assert.Equal(t, catalog.Dimension, resultField.FieldType)
	assert.Equal(t, catalog.All, resultField.Prevalence)
	assert.Emptyf(t, err, "Error updating dataset field: %s", err)
}

// Test DeleteDatasetField
func TestIntegrationDeleteDatasetField(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Delete dataset field
	err = client.CatalogService.DeleteDatasetField(datasetMap["id"].(string), resultField.ID)
	require.Nil(t, err)
	assert.Emptyf(t, err, "Error deleting dataset field: %s", err)

	// Validate the deletion of the dataset field
	result, err := client.CatalogService.GetDatasetField(datasetMap["id"].(string), resultField.ID)
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
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// Create a new dataset field
	testField := catalog.Field{Name: "integ_test_field", DataType: "N", FieldType: "U", Prevalence: "S"}
	resultField, err := invalidClient.CatalogService.CreateDatasetField(datasetMap["id"].(string), &testField)
	require.NotNil(t, err)
	assert.Empty(t, resultField)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
}

// Test PostDatasetField for 409 error
func TestIntegrationPostDatasetFieldDataAlreadyPresent(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// Create a new dataset field
	PostDatasetField(dataset, client, t)

	// Post an already created dataset field
	duplicateTestField := catalog.Field{Name: "integ_test_field", DataType: "S", FieldType: "D", Prevalence: "A"}
	resultField, err := client.CatalogService.CreateDatasetField(datasetMap["id"].(string), &duplicateTestField)
	require.NotNil(t, err)
	assert.Empty(t, resultField)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 409, "Expected error code 409")
}

// Test PostDatasetField for 500 error
func TestIntegrationPostDatasetFieldInvalidDataFormat(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// Create a new dataset field
	testField := catalog.Field{}
	resultField, err := client.CatalogService.CreateDatasetField(datasetMap["id"].(string), &testField)
	require.NotNil(t, err)
	assert.Empty(t, resultField)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 400, "Expected error code 400")
}

// Test GetDatasetFields for 401 Unauthorized operation error
func TestIntegrationGetDatasetFieldsUnauthorizedOperation(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)
	invalidClient := getInvalidClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// Create new fields in the dataset
	PostDatasetField(dataset, client, t)

	// Validate the creation of new dataset fields
	result, err := invalidClient.CatalogService.GetDatasetFields(datasetMap["id"].(string), nil)
	require.NotNil(t, err)
	assert.Empty(t, result)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
}

// Test PatchDatasetField for 401 error
func TestIntegrationPatchDatasetFieldUnauthorizedOperation(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)
	invalidClient := getInvalidClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(datasetMap["id"].(string), resultField.ID)
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Emptyf(t, err, "Error retrieving dataset field: %s", err)

	// Update the existing dataset field
	resultField, err = invalidClient.CatalogService.UpdateDatasetField(datasetMap["id"].(string), resultField.ID, &catalog.Field{DataType: "O"})
	require.NotNil(t, err)
	assert.Empty(t, resultField)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
}

// Test PatchDatasetField for 404 error
func TestIntegrationPatchDatasetFieldDataNotFound(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// Ceate dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Validate the creation of a new dataset field
	resultField, err = client.CatalogService.GetDatasetField(datasetMap["id"].(string), resultField.ID)
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Emptyf(t, err, "Error retrieving dataset field: %s", err)

	// Update the existing dataset field
	resultField, err = client.CatalogService.UpdateDatasetField(datasetMap["id"].(string), "123", &catalog.Field{DataType: "O"})
	require.NotNil(t, err)
	assert.Empty(t, resultField)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test DeleteDatasetField for 401 error
func TestIntegrationDeleteDatasetFieldUnauthorizedOperation(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)
	invalidClient := getInvalidClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// Create a new dataset field
	resultField := PostDatasetField(dataset, client, t)

	// Delete dataset field
	err = invalidClient.CatalogService.DeleteDatasetField(datasetMap["id"].(string), resultField.ID)
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 401, "Expected error code 401")
}

// Test DeleteDatasetField for 404 error
func TestIntegrationDeleteDatasetFieldDataNotFound(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// Create dataset
	dataset, err := createLookupDataset(t, testutils.TestNamespace, testutils.TestCollection, datasetOwner, datasetCapabilities, externalKind, externalName)
	require.Nil(t, err)
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	// Delete dataset field
	err = client.CatalogService.DeleteDatasetField(datasetMap["id"].(string), "123")
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).HTTPStatusCode == 404, "Expected error code 404")
}

// Test getfield(s) endpoints
func TestGetFields(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(
		&catalog.KVCollectionDataset{
			Name:         testutils.TestCollection,
			Kind:         "kvcollection",
			Module:       testutils.TestNamespace,
			Capabilities: datasetCapabilities,
		})
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)
	defer client.CatalogService.DeleteDataset(datasetName1)

	// create new fields in the dataset
	testField1 := catalog.Field{Name: "integ_test_field1", DataType: "S", FieldType: "D", Prevalence: "A"}
	field, err := client.CatalogService.CreateDatasetField(datasetMap["id"].(string), &testField1)
	require.Nil(t, err)
	defer client.CatalogService.DeleteDatasetField(datasetMap["id"].(string), field.ID)

	// get fields
	fields, err := client.CatalogService.GetFields()
	require.Nil(t, err)
	assert.True(t, len(fields) > 0)

	field1, err := client.CatalogService.GetField(field.ID)
	require.Nil(t, err)
	assert.Equal(t, field.Name, field1.Name)
	assert.Equal(t, field.ID, field1.ID)

	// Delete dataset field
	err = client.CatalogService.DeleteDatasetField(datasetMap["id"].(string), field.ID)
	require.Nil(t, err)

	// Delete dataset
	err = client.CatalogService.DeleteDataset(datasetMap["id"].(string))
	require.Nil(t, err)
}

// Test rule actions endpoints
func TestRuleActions(t *testing.T) {
	defer cleanupDatasets(t)

	client := getSdkClient(t)

	// Create dataset
	dataset, err := client.CatalogService.CreateDataset(&catalog.KVCollectionDataset{
		Name:         testutils.TestCollection,
		Kind:         "kvcollection",
		Module:       testutils.TestNamespace,
		Capabilities: datasetCapabilities,
	})
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)
	defer client.CatalogService.DeleteDataset(datasetName1)

	// create new fields in the dataset
	testField1 := catalog.Field{Name: "integ_test_field1", DataType: "S", FieldType: "D", Prevalence: "A"}
	field, err := client.CatalogService.CreateDatasetField(datasetMap["id"].(string), &testField1)
	require.Nil(t, err)
	defer client.CatalogService.DeleteDatasetField(datasetMap["id"].(string), field.ID)

	// Create rule and rule action
	rule, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName, Module: ruleModule, Match: ruleMatch, Owner: owner})
	defer client.CatalogService.DeleteRule(fmt.Sprintf("%v.%v", ruleModule, ruleName))
	require.Nil(t, err)

	// Create rule action
	action1, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewAliasAction(field.Name, "myfieldalias", ""))
	require.Nil(t, err)
	defer client.CatalogService.DeleteRuleAction(rule.ID, action1.ID)

	//update rule action
	tmpstr := "newaliasi"
	updateact, err := client.CatalogService.UpdateRuleAction(rule.ID, action1.ID, catalog.NewUpdateAliasAction(nil, &tmpstr))
	require.NotNil(t, updateact)
	assert.Equal(t, tmpstr, updateact.Alias)

	action2, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewAutoKVAction("auto", "owner1"))
	require.Nil(t, err)
	defer client.CatalogService.DeleteRuleAction(rule.ID, action2.ID)

	//update rule action
	tmpstr = "auto"
	updateact, err = client.CatalogService.UpdateRuleAction(rule.ID, action2.ID, catalog.NewUpdateAutoKVAction(&tmpstr))
	require.NotNil(t, updateact)
	assert.Equal(t, tmpstr, updateact.Mode)

	action3, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewEvalAction(field.Name, "some expression", ""))
	require.Nil(t, err)
	defer client.CatalogService.DeleteRuleAction(rule.ID, action3.ID)

	//update rule action
	tmpstr = "newField"
	updateact, err = client.CatalogService.UpdateRuleAction(rule.ID, action3.ID, catalog.NewUpdateEvalAction(&tmpstr, nil))
	require.NotNil(t, updateact)
	assert.Equal(t, tmpstr, updateact.Field)

	action4, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewLookupAction("myexpression2", ""))
	require.Nil(t, err)
	defer client.CatalogService.DeleteRuleAction(rule.ID, action4.ID)

	//update rule action
	tmpstr = "newexpr"
	updateact, err = client.CatalogService.UpdateRuleAction(rule.ID, action4.ID, catalog.NewUpdateLookupAction(&tmpstr))
	require.NotNil(t, updateact)
	assert.Equal(t, tmpstr, updateact.Expression)

	limit := 5
	action5, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewRegexAction(field.Name, "some pattern", &limit, ""))
	require.Nil(t, err)
	assert.Equal(t, 5, *action5.Limit)
	defer client.CatalogService.DeleteRuleAction(rule.ID, action5.ID)

	action6, err := client.CatalogService.CreateRuleAction(rule.ID, catalog.NewRegexAction(field.Name, "some pattern", nil, ""))
	require.Nil(t, err)
	assert.Equal(t, (*int)(nil), action6.Limit)
	defer client.CatalogService.DeleteRuleAction(rule.ID, action6.ID)

	//update rule action
	tmpstr = "newpattern"
	limit = 9
	updateact, err = client.CatalogService.UpdateRuleAction(rule.ID, action6.ID, catalog.NewUpdateRegexAction(nil, &tmpstr, &limit))
	require.NotNil(t, updateact)
	assert.Equal(t, tmpstr, updateact.Pattern)
	assert.Equal(t, limit, *updateact.Limit)

	//Get rule actions
	actions, err := client.CatalogService.GetRuleActions(rule.ID)
	require.NotNil(t, actions)
	assert.Equal(t, 6, len(actions))

	action7, err := client.CatalogService.GetRuleAction(rule.ID, actions[0].ID)
	require.NotNil(t, action7)

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
	err = client.CatalogService.DeleteDatasetField(datasetMap["id"].(string), field.ID)
	require.Nil(t, err)

	// Delete dataset
	err = client.CatalogService.DeleteDataset(datasetMap["id"].(string))
	require.Nil(t, err)
}

func PostDatasetField(dataset catalog.Dataset, client *sdk.Client, t *testing.T) *catalog.Field {
	testField := catalog.Field{Name: "integ_test_field", DataType: "S", FieldType: "D", Prevalence: "A"}
	datasetMap, err := convertDatasetInterfaceToMap(dataset)
	require.Nil(t, err)

	resultField, err := client.CatalogService.CreateDatasetField(datasetMap["id"].(string), &testField)
	require.Nil(t, err)
	assert.NotEmpty(t, resultField)
	assert.Equal(t, "integ_test_field", resultField.Name)
	assert.Equal(t, catalog.String, resultField.DataType)
	assert.Equal(t, catalog.Dimension, resultField.FieldType)
	assert.Equal(t, catalog.All, resultField.Prevalence)
	assert.Emptyf(t, err, "Error creating dataset field: %s", err)

	return resultField
}

// Test list modules
func TestIntegrationGetModules(t *testing.T) {
	client := getSdkClient(t)

	// test using NO filter
	modules, err := client.CatalogService.GetModules(nil)
	require.Nil(t, err)
	assert.True(t, len(modules) > 0)

	// test using filter
	filter := make(url.Values)
	filter.Add("filter", "module==\"\"")
	modules, err = client.CatalogService.GetModules(filter)
	require.Nil(t, err)
	assert.Equal(t, 1, len(modules))
	assert.Equal(t, "", modules[0].Name)
}

/*
/ Currently unable to generate a bad rule
func TestIntegrationCreateRuleInvalidRuleError(t *testing.T)  {
	defer cleanupRules(t)

	client := getSdkClient()

	// testing CreateRule for 400 Invalid Rule error
	ruleName := "goSdkTestrRule1"
	_, err := client.CatalogService.CreateRule(catalog.Rule{Name: ruleName})
	assert.NotNil(t, err)
   assert.True(t, err.(*util.HTTPError).Status == 400, "Expected error code 400")
}*/

// todo (Parul): 405 Rule cannot be deleted because of dependencies error case

// Helper function to convert custom interface type to map
func convertDatasetInterfaceToMap(dataset catalog.Dataset) (map[string]interface{}, error) {
	datasetByte, _ := json.Marshal(dataset)
	datasetMap := make(map[string]interface{})
	err := json.Unmarshal(datasetByte, &datasetMap)
	if err != nil {
		return nil, err
	}

	return datasetMap, err
}
