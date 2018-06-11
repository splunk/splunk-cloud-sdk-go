package playgroundintegration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Rule variables
var namespaceName = "namespace100"
var collectionName = "collection100"

func TestIntegrationGetCollectionStatus(t *testing.T) {
	client := getClient(t)

	//ToDo: To implement as a part of another APPLAT-1205
	//CreateNamespace(namespaceName)
	//CreateCollection(collectionName)

	response, err := client.KVStoreService.GetCollectionStats("ns100", "collection100")
	assert.Empty(t, err)
	assert.NotEmpty(t, response)
}

func TestIntegrationGetServiceHealth(t *testing.T) {
	client := getClient(t)

	//CreateNamespace(namespaceName)
	//CreateCollection(collectionName)

	response, err := client.KVStoreService.GetServiceHealthStatus()
	assert.Empty(t, err)
	assert.NotEmpty(t, response)
}