package playgroundintegration

import (
	"testing"

	"encoding/json"
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"net/url"
)

// Collection and Namespace test variables
var namespaceName = "ns100"
var collectionName = "collection100"

// Test GetCollectionStatus against nova playground
func TestIntegrationGetCollectionStatus(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: collectionName, Kind: "kvcollection", Owner: "integ_test", Module: namespaceName, Capabilities: "1100-11111:00000"})

	response, err := getClient(t).KVStoreService.GetCollectionStats(namespaceName, collectionName)
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

// Test CreateRecords() kvstore service endpoint against nova playground
func TestCreateRecords(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: collectionName, Kind: "kvcollection", Owner: "integ_test", Module: namespaceName, Capabilities: "1100-11111:00000"})

	var integrationTestRecord = `[{ "capacity_gb": 8, "size": "tiny", "description": "This is a tiny amount of GB", "_raw": ""} ,{"capacity_gb": 16,"size": "small","description": "This is a small amount of GB","_raw": ""}, {"type": "A","name": "test_record","count_of_fields": 3}]`
	var res []model.Record
	err = json.Unmarshal([]byte(integrationTestRecord), &res)
	assert.Nil(t, err)

	result, err := getClient(t).KVStoreService.CreateRecords(namespaceName, collectionName, res)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 3)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test GetRecords() kvstore service endpoint based on a query against nova playground
func TestGetRecordsWithQuery(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: collectionName, Kind: "kvcollection", Owner: "integ_test", Module: namespaceName, Capabilities: "1100-11111:00000"})

	// Create records
	var integrationTestRecord = `[{ "capacity_gb": 8, "size": "tiny", "description": "This is a tiny amount of GB", "_raw": ""} ,{"capacity_gb": 16,"size": "small","description": "This is a small amount of GB","_raw": ""}, {"type": "A","name": "test_record","count_of_fields": 3}]`
	var res []model.Record
	err = json.Unmarshal([]byte(integrationTestRecord), &res)
	assert.Nil(t, err)

	keys, err := getClient(t).KVStoreService.CreateRecords(namespaceName, collectionName, res)
	assert.Nil(t, err)
	assert.Equal(t, len(keys), 3)

	// Get records
	query := make(url.Values)
	query.Add("size", "tiny")
	query.Add("capacity_gb", "8")
	result, err := getClient(t).KVStoreService.GetRecords(query, namespaceName, collectionName)

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

// Test GetRecords() kvstore service endpoint against the nova playground
func TestGetAllRecords(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: collectionName, Kind: "kvcollection", Owner: "integ_test", Module: namespaceName, Capabilities: "1100-11111:00000"})

	// Create records
	var integrationTestRecord = `[{ "capacity_gb": 8, "size": "tiny", "description": "This is a tiny amount of GB", "_raw": ""} ,{"capacity_gb": 16,"size": "small","description": "This is a small amount of GB","_raw": ""}, {"type": "A","name": "test_record","count_of_fields": 3}]`
	var res []model.Record
	err = json.Unmarshal([]byte(integrationTestRecord), &res)
	assert.Nil(t, err)

	keys, err := getClient(t).KVStoreService.CreateRecords(namespaceName, collectionName, res)
	assert.Nil(t, err)
	assert.Equal(t, len(keys), 3)

	// Get all the records for validation
	result, err := getClient(t).KVStoreService.GetRecords(nil, namespaceName, collectionName)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 3)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test GetRecords() kvstore service endpoint based on a key against the nova playground
func TestGetRecordByKey(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: collectionName, Kind: "kvcollection", Owner: "integ_test", Module: namespaceName, Capabilities: "1100-11111:00000"})

	// Create records
	var integrationTestRecord = `[{ "capacity_gb": 8, "size": "tiny", "description": "This is a tiny amount of GB", "_raw": ""} ,{"capacity_gb": 16,"size": "small","description": "This is a small amount of GB","_raw": ""}, {"type": "A","name": "test_record","count_of_fields": 3}]`
	var res []model.Record
	err = json.Unmarshal([]byte(integrationTestRecord), &res)
	assert.Nil(t, err)

	keys, err := getClient(t).KVStoreService.CreateRecords(namespaceName, collectionName, res)
	assert.Nil(t, err)
	assert.Equal(t, len(keys), 3)

	result, err := getClient(t).KVStoreService.GetRecordByKey(namespaceName, collectionName, keys[0])

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
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: collectionName, Kind: "kvcollection", Owner: "integ_test", Module: namespaceName, Capabilities: "1100-11111:00000"})

	// Create records
	var integrationTestRecord = `[{ "capacity_gb": 8, "size": "tiny", "description": "This is a tiny amount of GB", "_raw": ""} ,{"capacity_gb": 16,"size": "small","description": "This is a small amount of GB","_raw": ""}, {"type": "A","name": "test_record","count_of_fields": 3}]`
	var res []model.Record
	err = json.Unmarshal([]byte(integrationTestRecord), &res)
	assert.Nil(t, err)

	keys, err := getClient(t).KVStoreService.CreateRecords(namespaceName, collectionName, res)
	assert.Nil(t, err)
	assert.Equal(t, len(keys), 3)

	// Delete record by key
	err = getClient(t).KVStoreService.DeleteRecordByKey(namespaceName, collectionName, keys[0])
	assert.Nil(t, err)

	// Validate that the record has been deleted
	retrievedRecordsByKey, err := getClient(t).KVStoreService.GetRecordByKey(namespaceName, collectionName, keys[0])
	assert.Nil(t, retrievedRecordsByKey)

	retrievedRecords, err := getClient(t).KVStoreService.GetRecords(nil, namespaceName, collectionName)
	assert.Equal(t, len(retrievedRecords), 2)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test DeleteRecords() kvstore service endpoint based on a query against the nova playground
func TestDeleteRecord(t *testing.T) {
	// Create the test collection and test namespace
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: collectionName, Kind: "kvcollection", Owner: "integ_test", Module: namespaceName, Capabilities: "1100-11111:00000"})

	// Create records
	var integrationTestRecord = `[{ "capacity_gb": 8, "size": "tiny", "description": "This is a tiny amount of GB", "_raw": ""} ,{"capacity_gb": 16,"size": "small","description": "This is a small amount of GB","_raw": ""}, {"type": "A","name": "test_record","count_of_fields": 3}]`
	var res []model.Record
	err = json.Unmarshal([]byte(integrationTestRecord), &res)
	assert.Nil(t, err)

	keys, err := getClient(t).KVStoreService.CreateRecords(namespaceName, collectionName, res)
	assert.Nil(t, err)
	assert.Equal(t, len(keys), 3)

	// Create query to test delete operation
	var integrationTestQuery = `{"capacity_gb": 16}`
	outerQuery := make(url.Values)
	outerQuery.Add("query", integrationTestQuery)
	outerQuery.Encode()

	err = getClient(t).KVStoreService.DeleteRecords(outerQuery, namespaceName, collectionName)
	assert.Nil(t, err)

	// Validate that the record has been deleted
	retrievedRecords, err := getClient(t).KVStoreService.GetRecords(nil, namespaceName, collectionName)
	assert.Equal(t, len(retrievedRecords), 2)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}
