// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test variables
var testIndex = "integtestindex"
var kvCollection = testutils.TestNamespace + "." + testutils.TestCollection

// --------------------------------------------------------------------------------
// Admin Endpoints
// --------------------------------------------------------------------------------

// Test GetServiceHealthStatus against nova playground
func TestIntegrationGetServiceHealth(t *testing.T) {
	response, err := getClient(t).KVStoreService.GetServiceHealthStatus()
	require.Empty(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, model.PingOKBodyStatusHealthy, response.Status)
}

// --------------------------------------------------------------------------------
// Index Endpoints
// --------------------------------------------------------------------------------
// /TENANT_NAME/kvstore/v2/collections/COLLECTION_NAME/indexes

// Test CreateIndex, ListIndexes and DeleteIndex kvstore endpoints
func TestIntegrationIndexEndpoints(t *testing.T) {
	// Create the test collection
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// Create Index
	var fields [1]model.IndexFieldDefinition
	fields[0] = model.IndexFieldDefinition{Direction: -1, Field: "integ_testField1"}
	indexDescription, err := getClient(t).KVStoreService.CreateIndex(kvCollection,
		model.IndexDefinition{
			Name:   testIndex,
			Fields: fields[:]})
	require.Nil(t, err)
	require.NotEmpty(t, indexDescription)
	assert.Equal(t, indexDescription.Collection, kvCollection)

	// Validate if the index was created
	indexes, err := getClient(t).KVStoreService.ListIndexes(kvCollection)
	require.Nil(t, err)
	require.NotNil(t, indexes)
	assert.Equal(t, len(indexes), 1)
	assert.Equal(t, indexes[0].Name, testIndex)

	// Delete the test index
	err = getClient(t).KVStoreService.DeleteIndex(kvCollection, testIndex)
	assert.Nil(t, err)

	// Validate if the index was deleted
	result, err := getClient(t).KVStoreService.ListIndexes(kvCollection)
	assert.Nil(t, err)
	require.NotNil(t, result)
	assert.Equal(t, len(result), 0)
}

// Test CreateIndex for 422 Unprocessable Entity error
func TestIntegrationCreateIndexUnprocessableEntityError(t *testing.T) {
	// Create the test collection
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// Create Index
	_, err := getClient(t).KVStoreService.CreateIndex(kvCollection, model.IndexDefinition{Name: testIndex, Fields: nil})
	require.NotNil(t, err)
	assert.Equal(t, 422, err.(*util.HTTPError).HTTPStatusCode)
	assert.Equal(t, "fields in body is required", err.(*util.HTTPError).Message)
}

// Test CreateIndex for 404 Not Found error TODO: Change name of non existing collection
func TestIntegrationCreateIndexNonExistingCollection(t *testing.T) {
	// Create the test collection
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// Create Index
	var fields [1]model.IndexFieldDefinition
	fields[0] = model.IndexFieldDefinition{Direction: -1, Field: "integ_testField1"}
	_, err := getClient(t).KVStoreService.CreateIndex(testutils.TestCollection, model.IndexDefinition{Name: testIndex, Fields: fields[:]})
	require.NotNil(t, err)
	assert.EqualValues(t, 404, err.(*util.HTTPError).HTTPStatusCode)
	// Known bug: should actually provide collection name - see https://jira.splunk.com/browse/SSC-5084
	assert.EqualValues(t, "collection not found: ", err.(*util.HTTPError).Message)
}

// Test DeleteIndex for 404 Index not found error
func TestIntegrationDeleteNonExistingIndex(t *testing.T) {
	// Create the test collection
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// DeleteIndex
	err := getClient(t).KVStoreService.DeleteIndex(kvCollection, testIndex)
	require.Nil(t, err)
}

// --------------------------------------------------------------------------------
// Record Endpoints
// --------------------------------------------------------------------------------
// /TENANT_NAME/kvstore/v2/collections/COLLECTION_NAME

// Test InsertRecords() kvstore service endpoint against nova playground
func TestCreateRecords(t *testing.T) {
	// Create the test collection
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	CreateTestRecord(t)
}

// Test InsertRecords() kvstore service endpoint against nova playground
func TestPutRecords(t *testing.T) {
	// Create the test collection
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	keys := CreateTestRecord(t)

	record := `
	{
		"capacity_gb": 8,
		"size": "notbig",
		"description": "this is a notbig amount of GB",
		"_raw": ""
	}`

	var res model.Record
	err := json.Unmarshal([]byte(record), &res)
	require.Nil(t, err)

	// test replace record
	key, created, err := getClient(t).KVStoreService.PutRecord(kvCollection, keys[0], res)
	require.Nil(t, err)
	require.NotNil(t, key)
	assert.Equal(t, key["_key"], keys[0])
	assert.False(t, created)

	// test insert record
	recordID := "recordID"
	key, created, err = getClient(t).KVStoreService.PutRecord(kvCollection, recordID, res)
	require.Nil(t, err)
	require.NotNil(t, key)
	assert.Equal(t, key["_key"], recordID)
	assert.True(t, created)
}

// Test getRecordByKey() kvstore service endpoint against the nova playground
func TestGetRecordByKey(t *testing.T) {
	// Create the test collection
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	keys := CreateTestRecord(t)

	result, err := getClient(t).KVStoreService.GetRecordByKey(kvCollection, keys[0])

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.NotNil(t, result["_key"])
	assert.Equal(t, result["capacity_gb"], float64(8))
	assert.Equal(t, result["description"], "This is a tiny amount of GB")
	assert.Equal(t, result["size"], "tiny")
}

// Test DeleteRecords() kvstore service endpoint based on a key against the nova playground
func TestDeleteRecordByKey(t *testing.T) {
	// Create the test collection
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	keys := CreateTestRecord(t)

	// Delete record by key
	err := getClient(t).KVStoreService.DeleteRecordByKey(kvCollection, keys[0])
	require.Nil(t, err)

	// Validate that the record has been deleted
	retrievedRecordsByKey, err := getClient(t).KVStoreService.GetRecordByKey(kvCollection, keys[0])
	assert.Nil(t, retrievedRecordsByKey)

	retrievedRecords, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	require.NotNil(t, retrievedRecords)
	assert.Equal(t, len(retrievedRecords), 2)
}

// Test DeleteRecords() kvstore service endpoint based on a query against the nova playground
func TestDeleteRecord(t *testing.T) {
	// Create the test collection
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// Create records
	CreateTestRecord(t)

	// Create query to test delete operation
	var integrationTestQuery = `{"capacity_gb": 16}`
	outerQuery := make(url.Values)
	outerQuery.Add("query", integrationTestQuery)
	outerQuery.Encode()

	err := getClient(t).KVStoreService.DeleteRecords(kvCollection, outerQuery)
	require.Nil(t, err)

	// Validate that the record has been deleted
	retrievedRecords, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	require.NotNil(t, retrievedRecords)
	assert.Equal(t, len(retrievedRecords), 2)
}

// Create test record
func CreateTestRecord(t *testing.T) []string {
	var integrationTestRecord = `[
         {
          "capacity_gb": 8,
          "size": "tiny",
          "description": "This is a tiny amount of GB",
          "_raw": ""
         },
         {
          "capacity_gb": 16,
          "size": "small",
          "description": "This is a small amount of GB",
          "_raw": ""
         },
         {
          "type": "A",
          "name": "test_record",
          "count_of_fields": 3
         }
        ]`
	var res []model.Record
	err := json.Unmarshal([]byte(integrationTestRecord), &res)
	require.Nil(t, err)

	keys, err := getClient(t).KVStoreService.InsertRecords(kvCollection, res)
	require.Nil(t, err)
	require.NotNil(t, keys)
	assert.Equal(t, len(keys), 3)

	return keys
}
