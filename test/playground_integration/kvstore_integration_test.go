package playgroundintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Collection and Namespace test variables
var testCollection = "integ_test_collection"
var testNamespace = "integ_test_namespace"
var testIndex = "integ_test_index"

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

// Test CreateIndex, GetIndexes and DeleteIndex kvstore endpoints
func TestIntegrationIndexEndpoints(t *testing.T) {
	// Create the test collection and test namespace for all the index operations
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: testCollection, Kind: "kvcollection", Owner: "integ_test", Module: testNamespace, Capabilities: "1100-11111:00000"})

	// Create Index
	var fields [1]model.IndexFieldDefinition
	fields[0] = model.IndexFieldDefinition{Direction: -1, Field: "integ_testField1"}
	err = getClient(t).KVStoreService.CreateIndex(model.IndexDescription{Name: testIndex, Collection: testCollection, Namespace: testNamespace, Fields: fields[:]}, testNamespace, testCollection)
	assert.Nil(t, err)

	// Validate if the index was created
	result, err := getClient(t).KVStoreService.GetIndexes(testNamespace, testCollection)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0].Name, testIndex)

	// Delete the test index
	err = getClient(t).KVStoreService.DeleteIndex(testNamespace, testCollection, testIndex)
	assert.Nil(t, err)

	// Validate if the index was deleted
	result, err = getClient(t).KVStoreService.GetIndexes(testNamespace, testCollection)
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
