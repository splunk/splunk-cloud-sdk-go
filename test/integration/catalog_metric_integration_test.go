// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
)

func TestCatalogMetricCreationSuccess(t *testing.T) {
	createMetricDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities,
		false)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)
}

func TestCatalogMetricCreationWithMissingCollectionName(t *testing.T) {
	createMetricDatasetInfo := catalog.MetricDataset{
		Kind:         "metric",
		Module:       testutils.TestNamespace,
		Capabilities: datasetCapabilities,
	}

	datasetInfo, err := getSdkClient(t).CatalogService.CreateDataset(createMetricDatasetInfo)

	assert.Nil(t, datasetInfo)
	assert.NotNil(t, err)
}
