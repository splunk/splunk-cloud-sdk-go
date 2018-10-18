// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"fmt"
	"testing"
	"time"
)

func TestCatalogImportCreationSuccess(t *testing.T) {

	timeSec := time.Now().Unix()
	var testNamespaceOne = fmt.Sprintf("go1%d", timeSec)
	var testCollectionOne = fmt.Sprintf("go1%d", timeSec)
	var testNamespaceTwo = fmt.Sprintf("gometric2%d", timeSec)
	var testCollectionTwo = fmt.Sprintf("gomodule2%d", timeSec)

	metricDatasetInfoOne, _ := createMetricDataset(t,
		testNamespaceOne,
		testCollectionOne,
		datasetOwner,
		datasetCapabilities,
		false)

	metricDatasetInfoTwo, _ := createMetricDataset(t,
		testNamespaceTwo,
		testCollectionTwo,
		datasetOwner,
		datasetCapabilities,
		false)

	createImportDataset(t,
		metricDatasetInfoOne.Name,
		metricDatasetInfoOne.Module,
		metricDatasetInfoTwo.Name,
		metricDatasetInfoTwo.Module,
	)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)
}

//func TestCatalogImportCreationWithMissingCollectionName(t *testing.T) {
//  createMetricDatasetInfo := model.DatasetCreationPayload{
//      Kind:         catalog.Import,
//      Owner:        datasetOwner,
//      Module:       testutils.TestNamespace,
//      Capabilities: datasetCapabilities,
//  }
//
//  datasetInfo, err := getSdkClient(t).CatalogService.CreateDataset(&createMetricDatasetInfo)
//
//  assert.Nil(t, datasetInfo)
//  assert.NotNil(t, err)
//}
