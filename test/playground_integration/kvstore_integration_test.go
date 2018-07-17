package playgroundintegration

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/splunk/ssc-client-go/util"
)

// Collection and Namespace test variables
var testIndex = "integtestindex"

// Test GetCollectionStatus against nova playground
func TestIntegrationGetCollectionStatus(t *testing.T) {
	// Create the test collection and test namespace
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	response, err := getClient(t).KVStoreService.GetCollectionStats(testutils.TestNamespace, testutils.TestCollection)
	assert.Empty(t, err)
	assert.NotEmpty(t, response)
}

// Test GetServiceHealthStatus against nova playground
func TestIntegrationGetServiceHealth(t *testing.T) {
	response, err := getClient(t).KVStoreService.GetServiceHealthStatus()
	assert.Empty(t, err)
	assert.NotEmpty(t, response)
}

// Test CreateIndex, ListIndexes and DeleteIndex kvstore endpoints
func TestIntegrationIndexEndpoints(t *testing.T) {
	// Create the test collection and test namespace for all the index operations
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
	err = getClient(t).KVStoreService.CreateIndex(model.IndexDescription{
		Name:       testIndex,
		Collection: testutils.TestCollection,
		Namespace:  testutils.TestNamespace,
		Fields:     fields[:]},
		testutils.TestNamespace,
		testutils.TestCollection)
	assert.Nil(t, err)

	// Validate if the index was created
	result, err := getClient(t).KVStoreService.ListIndexes(testutils.TestNamespace, testutils.TestCollection)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0].Name, testIndex)

	// Delete the test index
	err = getClient(t).KVStoreService.DeleteIndex(testutils.TestNamespace, testutils.TestCollection, testIndex)
	assert.Nil(t, err)

	// Validate if the index was deleted
	result, err = getClient(t).KVStoreService.ListIndexes(testutils.TestNamespace, testutils.TestCollection)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 0)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test CreateIndex for 422 Unprocessable Entity error
func TestIntegrationCreateIndexUnprocessableEntityError(t *testing.T) {
	// Create the test collection and test namespace
	_, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// Create Index
	err = getClient(t).KVStoreService.CreateIndex(model.IndexDescription{Name: testIndex, Collection: testutils.TestCollection, Namespace: testutils.TestNamespace, Fields: nil}, testutils.TestNamespace, testutils.TestCollection)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 422, "Expected error code 422")
	assert.True(t, err.(*util.HTTPError).Message == "422 Unprocessable Entity", "Expected error message should be 422 Unprocessable Entity")
}

// Test CreateIndex for 500 Internal server error
func TestIntegrationCreateIndexNonExistingCollection(t *testing.T) {
	// Create Index
	var fields [1]model.IndexFieldDefinition
	fields[0] = model.IndexFieldDefinition{Direction: -1, Field: "integ_testField1"}
	err := getClient(t).KVStoreService.CreateIndex(model.IndexDescription{Name: testIndex, Collection: testutils.TestCollection, Namespace: testutils.TestNamespace, Fields: fields[:]}, testutils.TestNamespace, testutils.TestCollection)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 500, "Expected error code 500")
	assert.True(t, err.(*util.HTTPError).Message == "500 Internal Server Error", "Expected error message should be 500 Internal Server Error")
}

// Test DeleteIndex for 404 Index not found error
func TestIntegrationDeleteNonExitingIndex(t *testing.T) {
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// DeleteIndex
	err := getClient(t).KVStoreService.DeleteIndex(testutils.TestNamespace, testutils.TestCollection, testIndex)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
	assert.True(t, err.(*util.HTTPError).Message == "404 Not Found", "Expected error message should be 404 Not Found")
}

// Test InsertRecords() kvstore service endpoint against nova playground
func TestCreateRecords(t *testing.T) {
	// Create the test collection and test namespace
	_, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	CreateTestRecord(err, t)
}

// Test getRecordByKey() kvstore service endpoint against the nova playground
func TestGetRecordByKey(t *testing.T) {
	// Create the test collection and test namespace
	_, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	keys := CreateTestRecord(err, t)

	result, err := getClient(t).KVStoreService.GetRecordByKey(testutils.TestNamespace, testutils.TestCollection, keys[0])

	assert.Nil(t, err)
	assert.NotNil(t, result["_key"])
	assert.Equal(t, result["capacity_gb"], float64(8))
	assert.Equal(t, result["description"], "This is a tiny amount of GB")
	assert.Equal(t, result["size"], "tiny")
}

// Test DeleteRecords() kvstore service endpoint based on a key against the nova playground
func TestDeleteRecordByKey(t *testing.T) {
	// Create the test collection and test namespace
	_, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	keys := CreateTestRecord(err, t)

	// Delete record by key
	err = getClient(t).KVStoreService.DeleteRecordByKey(testutils.TestNamespace, testutils.TestCollection, keys[0])
	assert.Nil(t, err)

	// Validate that the record has been deleted
	retrievedRecordsByKey, err := getClient(t).KVStoreService.GetRecordByKey(testutils.TestNamespace, testutils.TestCollection, keys[0])
	assert.Nil(t, retrievedRecordsByKey)

	retrievedRecords, err := getClient(t).KVStoreService.QueryRecords(testutils.TestNamespace, testutils.TestCollection, nil)
	assert.Equal(t, len(retrievedRecords), 2)
}

// Test DeleteRecords() kvstore service endpoint based on a query against the nova playground
func TestDeleteRecord(t *testing.T) {
	// Create the test collection and test namespace
	_, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// Create records
	CreateTestRecord(err, t)

	// Create query to test delete operation
	var integrationTestQuery = `{"capacity_gb": 16}`
	outerQuery := make(url.Values)
	outerQuery.Add("query", integrationTestQuery)
	outerQuery.Encode()

	err = getClient(t).KVStoreService.DeleteRecords(outerQuery, testutils.TestNamespace, testutils.TestCollection)
	assert.Nil(t, err)

	// Validate that the record has been deleted
	retrievedRecords, err := getClient(t).KVStoreService.QueryRecords(testutils.TestNamespace, testutils.TestCollection, nil)
	assert.Equal(t, len(retrievedRecords), 2)
}

// Create test record
func CreateTestRecord(err error, t *testing.T) []string {
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
	err = json.Unmarshal([]byte(integrationTestRecord), &res)
	assert.Nil(t, err)

	keys, err := getClient(t).KVStoreService.InsertRecords(testutils.TestNamespace, testutils.TestCollection, res)
	assert.Nil(t, err)
	assert.Equal(t, len(keys), 3)

	return keys
}
