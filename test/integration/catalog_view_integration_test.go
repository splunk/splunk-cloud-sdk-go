// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestCatalogViewCreationSuccess(t *testing.T) {
	createViewDataset(t, "",
		testutils.TestCollection, "", "",
		"search index='main' | head limit=10 | stats count()",
	)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)
}

func TestCatalogViewCreationWithMissingCollectionName(t *testing.T) {
	createMetricDatasetInfo := model.DatasetCreationPayload{
		Kind:   catalog.View,
		Search: "from index:main | head limit=10 | stats count()",
	}

	datasetInfo, err := getSdkClient(t).CatalogService.CreateDataset(&createMetricDatasetInfo)

	assert.Nil(t, datasetInfo)
	assert.NotNil(t, err)
}
