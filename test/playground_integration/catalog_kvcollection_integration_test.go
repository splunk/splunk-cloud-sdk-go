// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/testutils"
	"github.com/stretchr/testify/assert"
)

func TestDatastoreKVStoreCreationSuccess(t *testing.T) {
	createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)
}

func TestDatastoreKVStoreCreationWithMissingCollectionName(t *testing.T) {
	createDatasetInfo := model.DatasetCreationPayload{
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
	createDatasetInfo := model.DatasetCreationPayload{
		Name:         testutils.TestCollection,
		Kind:         model.KVCOLLECTION,
		Owner:        datasetOwner,
		Capabilities: datasetCapabilities,
	}

	datasetInfo, err := getClient(t).CatalogService.CreateDataset(createDatasetInfo)

	assert.Nil(t, datasetInfo)
	assert.NotNil(t, err)
}
