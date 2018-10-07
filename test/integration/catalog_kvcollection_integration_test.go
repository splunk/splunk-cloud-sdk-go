// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

// import (
// 	"testing"

// 	"github.com/splunk/splunk-cloud-sdk-go/model"
// 	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
// 	"github.com/stretchr/testify/assert"
// )

// func TestDatastoreKVStoreCreationSuccess(t *testing.T) {
// 	createKVCollectionDataset(t,
// 		testutils.TestNamespace,
// 		testutils.TestCollection,
// 		datasetOwner,
// 		datasetCapabilities)

// 	// Remove the dataset used for testing
// 	defer cleanupDatasets(t)
// }

// func TestDatastoreKVStoreCreationWithMissingCollectionName(t *testing.T) {
// 	createDatasetInfo := model.DatasetCreationPayload{
// 		Kind:         model.KVCOLLECTION,
// 		Owner:        datasetOwner,
// 		Module:       testutils.TestNamespace,
// 		Capabilities: datasetCapabilities,
// 	}

// 	datasetInfo, err := getClient(t).CatalogService.CreateDataset(&createDatasetInfo)

// 	assert.Nil(t, datasetInfo)
// 	assert.NotNil(t, err)
// }
