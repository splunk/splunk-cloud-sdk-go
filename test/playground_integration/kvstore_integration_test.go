package playgroundintegration

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
)

// Collection and Namespace test variables
var testCollection = "integtestcollection"
var testNamespace = "integtestnamespace"
var testIndex = "integtestindex"

// TODO (Logan): circle back and align the kvcollection creation on catalog with the other kvcollection integration tests
// Test GetCollectionStatus against nova playground
func TestIntegrationGetCollectionStatus(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: testCollection, Kind: "kvcollection", Owner: "integ_test", Module: testNamespace, Capabilities: "1100-11111:00000"})

	response, err := getClient(t).KVStoreService.GetCollectionStats(testNamespace, testCollection)
	assert.Empty(t, err)
	assert.NotEmpty(t, response)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
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
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: testCollection, Kind: "kvcollection", Owner: "integ_test", Module: testNamespace, Capabilities: "1100-11111:00000"})

	// Create Index
	var fields [1]model.IndexFieldDefinition
	fields[0] = model.IndexFieldDefinition{Direction: -1, Field: "integ_testField1"}
	err = getClient(t).KVStoreService.CreateIndex(model.IndexDescription{Name: testIndex, Collection: testCollection, Namespace: testNamespace, Fields: fields[:]}, testNamespace, testCollection)
	assert.Nil(t, err)

	// Validate if the index was created
	result, err := getClient(t).KVStoreService.ListIndexes(testNamespace, testCollection)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0].Name, testIndex)

	// Delete the test index
	err = getClient(t).KVStoreService.DeleteIndex(testNamespace, testCollection, testIndex)
	assert.Nil(t, err)

	// Validate if the index was deleted
	result, err = getClient(t).KVStoreService.ListIndexes(testNamespace, testCollection)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 0)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test CreateIndex for 422 Unprocessable Entity error
func TestIntegrationCreateIndexUnprocessableEntityError(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: testCollection, Kind: "kvcollection", Owner: "integ_test", Module: testNamespace, Capabilities: "1100-11111:00000"})

	// Create Index
	err = getClient(t).KVStoreService.CreateIndex(model.IndexDescription{Name: testIndex, Collection: testCollection, Namespace: testNamespace, Fields: nil}, testNamespace, testCollection)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 422, "Expected error code 422")
	assert.True(t, err.(*util.HTTPError).Message == "422 Unprocessable Entity", "Expected error message should be 422 Unprocessable Entity")

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test CreateIndex for 500 Internal server error
func TestIntegrationCreateIndexNonExistingCollection(t *testing.T) {
	// Create Index
	var fields [1]model.IndexFieldDefinition
	fields[0] = model.IndexFieldDefinition{Direction: -1, Field: "integ_testField1"}
	err := getClient(t).KVStoreService.CreateIndex(model.IndexDescription{Name: testIndex, Collection: testCollection, Namespace: testNamespace, Fields: fields[:]}, testNamespace, testCollection)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 500, "Expected error code 500")
	assert.True(t, err.(*util.HTTPError).Message == "500 Internal Server Error", "Expected error message should be 500 Internal Server Error")
}

// Test DeleteIndex for 404 Index not found error
func TestIntegrationDeleteNonExitingIndex(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: testCollection, Kind: "kvcollection", Owner: "integ_test", Module: testNamespace, Capabilities: "1100-11111:00000"})

	// DeleteIndex
	err = getClient(t).KVStoreService.DeleteIndex(testNamespace, testCollection, testIndex)
	assert.NotNil(t, err)
	assert.True(t, err.(*util.HTTPError).Status == 404, "Expected error code 404")
	assert.True(t, err.(*util.HTTPError).Message == "404 Not Found", "Expected error message should be 404 Not Found")

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test InsertRecords() kvstore service endpoint against nova playground
func TestCreateRecords(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: testCollection, Kind: "kvcollection", Owner: "integ_test", Module: testCollection, Capabilities: "1100-11111:00000"})

	CreateTestRecord(err, t)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test QueryRecords() kvstore service endpoint based on a query against nova playground
func TestGetRecordsWithQuery(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: testCollection, Kind: "kvcollection", Owner: "integ_test", Module: testCollection, Capabilities: "1100-11111:00000"})

	CreateTestRecord(err, t)

	// Get records
	query := make(url.Values)
	query.Add("size", "tiny")
	query.Add("capacity_gb", "8")
	result, err := getClient(t).KVStoreService.QueryRecords(query, testCollection, testCollection)

	assert.Nil(t, err)
	assert.Equal(t, len(result), 1)
	assert.NotNil(t, result[0]["_key"])
	assert.Equal(t, result[0]["capacity_gb"], float64(8))
	assert.Equal(t, result[0]["description"], "This is a tiny amount of GB")
	assert.Equal(t, result[0]["size"], "tiny")

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test QueryRecords() kvstore service endpoint against the nova playground
func TestGetAllRecords(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: testCollection, Kind: "kvcollection", Owner: "integ_test", Module: testCollection, Capabilities: "1100-11111:00000"})

	CreateTestRecord(err, t)

	// Get all the records for validation
	result, err := getClient(t).KVStoreService.QueryRecords(nil, testCollection, testCollection)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 3)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test getRecordByKey() kvstore service endpoint against the nova playground
func TestGetRecordByKey(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: testCollection, Kind: "kvcollection", Owner: "integ_test", Module: testCollection, Capabilities: "1100-11111:00000"})

	keys := CreateTestRecord(err, t)

	result, err := getClient(t).KVStoreService.GetRecordByKey(testCollection, testCollection, keys[0])

	assert.Nil(t, err)
	assert.NotNil(t, result["_key"])
	assert.Equal(t, result["capacity_gb"], float64(8))
	assert.Equal(t, result["description"], "This is a tiny amount of GB")
	assert.Equal(t, result["size"], "tiny")

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test DeleteRecords() kvstore service endpoint based on a key against the nova playground
func TestDeleteRecordByKey(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: testCollection, Kind: "kvcollection", Owner: "integ_test", Module: testCollection, Capabilities: "1100-11111:00000"})

	keys := CreateTestRecord(err, t)

	// Delete record by key
	err = getClient(t).KVStoreService.DeleteRecordByKey(testCollection, testCollection, keys[0])
	assert.Nil(t, err)

	// Validate that the record has been deleted
	retrievedRecordsByKey, err := getClient(t).KVStoreService.GetRecordByKey(testCollection, testCollection, keys[0])
	assert.Nil(t, retrievedRecordsByKey)

	retrievedRecords, err := getClient(t).KVStoreService.QueryRecords(nil, testCollection, testCollection)
	assert.Equal(t, len(retrievedRecords), 2)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test DeleteRecords() kvstore service endpoint based on a query against the nova playground
func TestDeleteRecord(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: testCollection, Kind: "kvcollection", Owner: "integ_test", Module: testCollection, Capabilities: "1100-11111:00000"})

	// Create records
	CreateTestRecord(err, t)

	// Create query to test delete operation
	var integrationTestQuery = `{"capacity_gb": 16}`
	outerQuery := make(url.Values)
	outerQuery.Add("query", integrationTestQuery)
	outerQuery.Encode()

	err = getClient(t).KVStoreService.DeleteRecords(outerQuery, testCollection, testCollection)
	assert.Nil(t, err)

	// Validate that the record has been deleted
	retrievedRecords, err := getClient(t).KVStoreService.QueryRecords(nil, testCollection, testCollection)
	assert.Equal(t, len(retrievedRecords), 2)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
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

	keys, err := getClient(t).KVStoreService.InsertRecords(testCollection, testCollection, res)
	assert.Nil(t, err)
	assert.Equal(t, len(keys), 3)

	return keys
}
