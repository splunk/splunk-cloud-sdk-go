package playgroundintegration

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/require"
)

// Test variables
var testIndex = "integtestindex"
var kvCollection = testutils.TestNamespace + "." + testutils.TestCollection

// --------------------------------------------------------------------------------
// Admin Endpoints
// --------------------------------------------------------------------------------

// Test GetCollectionStatus against nova playground
func TestIntegrationGetCollectionStatus(t *testing.T) {
	// Create the test collection
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	response, err := getClient(t).KVStoreService.GetCollectionStats(kvCollection)
	require.Empty(t, err)
	assert.NotEmpty(t, response)
}

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
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// Create Index
	var fields [1]model.IndexFieldDefinition
	fields[0] = model.IndexFieldDefinition{Direction: -1, Field: "integ_testField1"}
	indexDescription, err := getClient(t).KVStoreService.CreateIndex(model.IndexDefinition{
		Name:   testIndex,
		Fields: fields[:]},
		kvCollection)
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

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test CreateIndex for 422 Unprocessable Entity error
func TestIntegrationCreateIndexUnprocessableEntityError(t *testing.T) {
	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// Create Index
	_, err = getClient(t).KVStoreService.CreateIndex(model.IndexDefinition{Name: testIndex, Fields: nil}, kvCollection)
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 422, "Expected error code 422")
	assert.True(t, err.(*util.HTTPError).Message == "422 Unprocessable Entity", "Expected error message should be 422 Unprocessable Entity")

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test CreateIndex for 404 Not Found error TODO: Change name of non existing collection
func TestIntegrationCreateIndexNonExistingCollection(t *testing.T) {
	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// Create Index
	var fields [1]model.IndexFieldDefinition
	fields[0] = model.IndexFieldDefinition{Direction: -1, Field: "integ_testField1"}
	_, err = getClient(t).KVStoreService.CreateIndex(model.IndexDefinition{Name: testIndex, Fields: fields[:]}, testutils.TestCollection)
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
	assert.True(t, err.(*util.HTTPError).Message == "404 Not Found", "Expected error message should be 404 Not Found")

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test DeleteIndex for 404 Index not found error
func TestIntegrationDeleteNonExitingIndex(t *testing.T) {
	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// DeleteIndex
	err = getClient(t).KVStoreService.DeleteIndex(kvCollection, testIndex)
	require.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
	assert.True(t, err.(*util.HTTPError).Message == "404 Not Found", "Expected error message should be 404 Not Found")

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// --------------------------------------------------------------------------------
// Record Endpoints
// --------------------------------------------------------------------------------
// /TENANT_NAME/kvstore/v2/collections/COLLECTION_NAME

// Test InsertRecords() kvstore service endpoint against nova playground
func TestCreateRecords(t *testing.T) {
	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	CreateTestRecord(t)

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test getRecordByKey() kvstore service endpoint against the nova playground
func TestGetRecordByKey(t *testing.T) {
	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
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

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test DeleteRecords() kvstore service endpoint based on a key against the nova playground
func TestDeleteRecordByKey(t *testing.T) {
	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	keys := CreateTestRecord(t)

	// Delete record by key
	err = getClient(t).KVStoreService.DeleteRecordByKey(kvCollection, keys[0])
	require.Nil(t, err)

	// Validate that the record has been deleted
	retrievedRecordsByKey, err := getClient(t).KVStoreService.GetRecordByKey(kvCollection, keys[0])
	assert.Nil(t, retrievedRecordsByKey)

	retrievedRecords, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	require.NotNil(t, retrievedRecords)
	assert.Equal(t, len(retrievedRecords), 2)

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test DeleteRecords() kvstore service endpoint based on a query against the nova playground
func TestDeleteRecord(t *testing.T) {
	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
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

	err = getClient(t).KVStoreService.DeleteRecords(outerQuery, kvCollection)
	require.Nil(t, err)

	// Validate that the record has been deleted
	retrievedRecords, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	require.NotNil(t, retrievedRecords)
	assert.Equal(t, len(retrievedRecords), 2)

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
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