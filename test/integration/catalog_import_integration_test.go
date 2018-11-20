// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"

	"github.com/stretchr/testify/assert"
)

func TestCatalogImportCreationSuccess(t *testing.T) {
	defer cleanupDatasets(t)

	testNamespaceOne := "gotestn0"
	testCollectionOne := "gotestn0"
	testNamespaceTwo := "gotestn1"
	testCollectionTwo := "gotestc1"

	_, _ = createMetricDataset(t,
		testNamespaceOne,
		testCollectionOne,
		datasetOwner,
		datasetCapabilities,
		false)

	// Error checking occurs in the utility function
	_, _ = createImportDataset(t,
		testNamespaceOne,
		testCollectionOne,
		testNamespaceTwo,
		testCollectionTwo,
		datasetCapabilities)
}

func TestCatalogImportCreationWithMissingCollectionName(t *testing.T) {
	defer cleanupDatasets(t)

	testNamespaceOne := "gotestn0"
	testCollectionOne := "gotestn0"
	testCollectionTwo := "gotestc1"

	_, _ = createMetricDataset(t,
		testNamespaceOne,
		testCollectionOne,
		datasetOwner,
		datasetCapabilities,
		false)

	createImportDatasetInfo := catalog.DatasetCreationPayload{
		Capabilities: datasetCapabilities,
		Kind:         catalog.Import,
		Module:       testCollectionTwo,
		SourceName:   testNamespaceOne,
		SourceModule: testCollectionOne,
	}

	datasetInfo, err := getSdkClient(t).CatalogService.CreateDataset(&createImportDatasetInfo)

	assert.NotNil(t, err)
	assert.Nil(t, datasetInfo)
}
