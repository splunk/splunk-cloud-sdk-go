package playgroundintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var test_collection = "integ_test_collection"
var test_namespace = "integ_test_namespace"
var test_index = "integ_testIndex2"

// Test CreateIndex, GetIndexes and DeleteIndex kvstore endpoints
func TestIntegrationIndexEndpoints(t *testing.T) {
	// Create the test collection and test namespace for all the index operations
	dataset, err := getClient(t).CatalogService.CreateDataset(model.DatasetInfo{Name: test_collection, Kind: "kvcollection", Owner: "integ_test", Module: test_namespace, Capabilities: "1100-11111:00000"})

	// Create Index
	var fields [1]model.IndexFieldDefinition
	fields[0] = model.IndexFieldDefinition{Direction: -1, Field: "integ_testField1"}
	err = getClient(t).KVStoreService.CreateIndex(model.IndexDescription{Name: test_index, Collection: test_collection, Namespace: test_namespace, Fields: fields[:]}, test_namespace, test_collection)
	assert.Nil(t, err)

	// Validate if the index was created
	result, err := getClient(t).KVStoreService.GetIndexes(test_namespace, test_collection)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0].Name, test_index)

	// Delete the test index
	err = getClient(t).KVStoreService.DeleteIndex(test_namespace, test_collection, test_index)
	assert.Nil(t, err)

	// Validate if the index was deleted
	result, err = getClient(t).KVStoreService.GetIndexes(test_namespace, test_collection)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 0)

	// Delete the test collection and test namespace
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// Test CreateIndex, GetIndexes and DeleteIndex kvstore endpoints - Error scenarios
func TestIntegrationIndexEndpointsErrorScenarios(t *testing.T) {
	// 422 (If fields null or not present in Index)
	// 404 (If deleting an index, not present)
	// 500 Internal server error (If the collection or namespace where the index is being created is not present)
}
