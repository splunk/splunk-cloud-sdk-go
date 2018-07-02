package playgroundintegration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/testutils"
)

// createDatastoreKVCollection - Helper function for creating a valid KV Collection in Catalog
func createDatastoreKVCollection(t *testing.T, namespaceName string, collectionName string, datasetOwner string, capabilities string) (*model.DatasetInfo, error) {
	createDatasetInfo := model.DatasetInfo{
		Name:         testutils.TestCollection,
		Kind:         model.KVCOLLECTION,
		Owner:        datasetOwner,
		Module:       testutils.TestNamespace,
		Capabilities: datasetCapabilities}

	datasetInfo, err := getClient(t).CatalogService.CreateDataset(createDatasetInfo)
	assert.NotNil(t, datasetInfo)
	assert.Nil(t, err)

	return datasetInfo, err
}

func TestDatastoreKVStoreCreationSuccess(t *testing.T) {
	createDatastoreKVCollection(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	cleanupDatasets(t)
}

func TestDatastoreKVStoreCreationWithMissingCollectionName(t *testing.T) {
	createDatasetInfo := model.DatasetInfo{
		Kind:         model.KVCOLLECTION,
		Owner:        datasetOwner,
		Module:       testutils.TestNamespace,
		Capabilities: datasetCapabilities,
	}

	datasetInfo, err := getClient(t).CatalogService.CreateDataset(createDatasetInfo)

	assert.Nil(t, datasetInfo)
	assert.NotNil(t, err)
}

func TestDatastoreKVStoreCreationWithMissingNamespace(t *testing.T) {
	createDatasetInfo := model.DatasetInfo{
		Name:         testutils.TestCollection,
		Kind:         model.KVCOLLECTION,
		Owner:        datasetOwner,
		Capabilities: datasetCapabilities,
	}

	datasetInfo, err := getClient(t).CatalogService.CreateDataset(createDatasetInfo)

	assert.Nil(t, datasetInfo)
	assert.NotNil(t, err)
}
