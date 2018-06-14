package playgroundintegration

import (
	"testing"

	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
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
