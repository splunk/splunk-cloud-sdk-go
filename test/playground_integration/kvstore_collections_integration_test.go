package playgroundintegration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/testutils"
)

// --------------------------------------------------------------------------------
// Collection Endpoints
// --------------------------------------------------------------------------------
// /TENANT_NAME/kvstore/v1/NAMESPACE_NAME/collections/COLLECTION_NAME

// --------
// GET
// --------
func TestListRecordsReturnsEmptyDatasetOnCreation(t *testing.T) {
	createDatasetInfo := model.DatasetInfo{
		Name:         testutils.TestCollection,
		Kind:         model.KVCOLLECTION,
		Owner:        datasetOwner,
		Module:       testutils.TestNamespace,
		Capabilities: datasetCapabilities}

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	datasetInfo, err := getClient(t).CatalogService.CreateDataset(createDatasetInfo)
	assert.NotNil(t, datasetInfo)
	assert.Nil(t, err)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestNamespace, testutils.TestCollection)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)
}

func TestListRecordsReturnsCorrectDatasetAfterSingleInsertRecord(t *testing.T) {
	createDatasetInfo := model.DatasetInfo{
		Name:         testutils.TestCollection,
		Kind:         model.KVCOLLECTION,
		Owner:        datasetOwner,
		Module:       testutils.TestNamespace,
		Capabilities: datasetCapabilities}

	record := map[string]string{
		"TEST_KEY_01": "TEST_VALUE_01",
		"TEST_KEY_02": "TEST_VALUE_02",
		"TEST_KEY_03": "TEST_VALUE_03",
	}

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// Verify the initial dataset has no records
	datasetInfo, err := getClient(t).CatalogService.CreateDataset(createDatasetInfo)
	assert.NotNil(t, datasetInfo)
	assert.Nil(t, err)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestNamespace, testutils.TestCollection)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert a new record into the kvstore
	createRecordResponseMap, err := getClient(t).KVStoreService.InsertRecord(
		testutils.TestNamespace,
		testutils.TestCollection,
		record)

	assert.NotNil(t, createRecordResponseMap)
	assert.Nil(t, err)
	assert.Len(t, createRecordResponseMap, 1)

	// Makes sure that the only value returned is the unique _key
	for key, value := range createRecordResponseMap {
		assert.IsType(t, "string", key)
		assert.Equal(t, "_key", key)

		assert.NotNil(t, value)
		assert.IsType(t, "string", value)
	}

	// Make sure that records return match
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestNamespace, testutils.TestCollection)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)

	for _, element := range recordsAfterInsert {
		assert.NotNil(t, element)
		for key, value := range element {
			assert.IsType(t, "string", key)
			assert.NotNil(t, value)
		}
	}
}

// --------
// POST
// --------
// There is no separation for the testing the insertion of a record when using an incorrect namespace or collections
// because both are required in order to make a dataset on the catalog service
func TestInsertRecordIntoMissingNamespaceAndCollection(t *testing.T) {
	record := map[string]string{
		"TEST_KEY_01": "TEST_VALUE_01",
		"TEST_KEY_02": "TEST_VALUE_02",
		"TEST_KEY_03": "TEST_VALUE_03",
	}

	// Insert a new record into the kvstore
	createRecordResponseMap, err := getClient(t).KVStoreService.InsertRecord(
		testutils.TestNamespace,
		testutils.TestCollection,
		record)

	assert.Nil(t, createRecordResponseMap)
	assert.NotNil(t, err)
}

// Inserts a record into the specified tenant's namespace collection
func TestInsertRecordSuccess(t *testing.T) {
	createDatasetInfo := model.DatasetInfo{
		Name:         testutils.TestCollection,
		Kind:         model.KVCOLLECTION,
		Owner:        datasetOwner,
		Module:       testutils.TestNamespace,
		Capabilities: datasetCapabilities}

	record := map[string]string{
		"TEST_KEY_01": "TEST_VALUE_01",
		"TEST_KEY_02": "TEST_VALUE_02",
		"TEST_KEY_03": "TEST_VALUE_03",
	}

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	// Verify the initial dataset has no records
	datasetInfo, err := getClient(t).CatalogService.CreateDataset(createDatasetInfo)
	assert.NotNil(t, datasetInfo)
	assert.Nil(t, err)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestNamespace, testutils.TestCollection)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert a new record into the kvstore
	createRecordResponseMap, err := getClient(t).KVStoreService.InsertRecord(
		testutils.TestNamespace,
		testutils.TestCollection,
		record)

	assert.NotNil(t, createRecordResponseMap)
	assert.Nil(t, err)
	assert.Len(t, createRecordResponseMap, 1)

	// Makes sure that the only value returned is the unique _key
	for key, value := range createRecordResponseMap {
		assert.IsType(t, "string", key)
		assert.Equal(t, "_key", key)

		assert.NotNil(t, value)
		assert.IsType(t, "string", value)
	}
}
