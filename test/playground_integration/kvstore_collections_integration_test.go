package playgroundintegration

import (
	"testing"

	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
)

var recordOne = map[string]string{
	"TEST_KEY_01": "A",
	"TEST_KEY_02": "B",
	"TEST_KEY_03": "C",
}
var recordTwo = map[string]string{
	"TEST_KEY_01": "B",
	"TEST_KEY_02": "C",
	"TEST_KEY_03": "A",
}
var recordThree = map[string]string{
	"TEST_KEY_01": "C",
	"TEST_KEY_02": "A",
	"TEST_KEY_03": "B",
}

func createRecord(t *testing.T, collection string, record map[string]string) (map[string]string, error) {
	// Insert a new record into the kvstore
	createRecordResponseMap, err := getClient(t).KVStoreService.InsertRecord(
		collection,
		record)

	assert.NotNil(t, createRecordResponseMap)
	assert.Nil(t, err)

	for key, value := range createRecordResponseMap {
		assert.IsType(t, "string", key)
		assert.Equal(t, "_key", key)

		assert.NotNil(t, value)
		assert.IsType(t, "string", value)
	}

	return createRecordResponseMap, err
}

// --------------------------------------------------------------------------------
// Collection Endpoints
// --------------------------------------------------------------------------------
// /TENANT_NAME/kvstore/v2/collections/COLLECTION_NAME

// --------
// GET
// --------
func TestKVStoreCollectionsListRecordsReturnsEmptyDatasetOnCreation(t *testing.T) {
	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// --------
// GET ?fields= parameter
// --------
func TestKVStoreCollectionsListRecordsReturnsCorrectDatasetAfterSingleInsertRecord(t *testing.T) {
	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert a new record into the kvstore
	createRecordResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordResponseMap, 1)

	// Make sure that records return match
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)
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

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

func TestKVStoreCollectionsListRecordsFieldsValidInclude(t *testing.T) {
	fieldsToFilter := []string{"TEST_KEY_01"}
	filters := map[string][]string{
		"fields": fieldsToFilter,
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	// Make sure that records return match
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 2)

	for _, element := range recordsAfterInsert {
		assert.NotNil(t, element)
		for key, value := range element {
			assert.IsType(t, "string", key)
			assert.NotNil(t, value)
			assert.EqualValues(t, "TEST_KEY_01", key)
		}
	}

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

func TestKVStoreCollectionsListRecordsFieldsValidExclude(t *testing.T) {
	fieldsToFilter := []string{"TEST_KEY_01:0"}
	filters := map[string][]string{
		"fields": fieldsToFilter,
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	// Make sure that records return match
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 2)

	for _, element := range recordsAfterInsert {
		assert.NotNil(t, element)
		for key, value := range element {
			assert.IsType(t, "string", key)
			assert.NotNil(t, value)
			assert.NotEqual(t, "TEST_KEY_01", key)
		}
	}

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

func TestKVStoreCollectionsListRecordsFieldsValidIncludeAndExclude(t *testing.T) {
	// From the documenation: A fields value cannot contain both include and exclude specifications except for exclusion
	// of the _key field.
	fieldsToFilter := []string{"TEST_KEY_01,TEST_KEY_02:0"}
	filters := map[string][]string{
		"fields": fieldsToFilter,
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)
	assert.Nil(t, recordsAfterInsert)
	assert.NotNil(t, err)

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// --------
// GET ?count= parameter
// --------
func TestKVStoreCollectionsListRecordsCountValidInput(t *testing.T) {
	filters := map[string][]string{
		"count": {"1"},
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

func TestKVStoreCollectionsListRecordsCountNegativeOutOfBoundsInput(t *testing.T) {
	filters := map[string][]string{
		"count": {"-1"},
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)

	assert.Nil(t, recordsAfterInsert)
	assert.NotNil(t, err)

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

func TestKVStoreCollectionsListRecordsCountPositiveOutOfBoundsInput(t *testing.T) {
	filters := map[string][]string{
		"count": {"10000"},
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 2)

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// --------
// GET ?offset= parameter
// --------
func TestKVStoreCollectionsListRecordsOffsetValidInput(t *testing.T) {
	filters := map[string][]string{
		"offset": {"1"},
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

func TestKVStoreCollectionsListRecordsOffsetNegativeOutOfBoundsInput(t *testing.T) {
	filters := map[string][]string{
		"offset": {"-1"},
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)

	assert.Nil(t, recordsAfterInsert)
	assert.NotNil(t, err)

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

func TestKVStoreCollectionsListRecordsOffsetPositiveOutOfBoundsInput(t *testing.T) {
	filters := map[string][]string{
		"offset": {"10000"},
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 0)

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// --------
// GET ?orderby= parameter
// --------
func TestKVStoreCollectionsListRecordsOrderByValidInput(t *testing.T) {
	filters := map[string][]string{
		"orderby": {"TEST_KEY_02"},
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	// Insert the third record into the kvstore
	createRecordThreeResponseMap, err := createRecord(t, kvCollection, recordThree)
	assert.Len(t, createRecordThreeResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 3)

	assert.EqualValues(t, "A", recordsAfterInsert[0]["TEST_KEY_02"])
	assert.EqualValues(t, "B", recordsAfterInsert[1]["TEST_KEY_02"])
	assert.EqualValues(t, "C", recordsAfterInsert[2]["TEST_KEY_02"])

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

func TestKVStoreCollectionsListRecordsOrderByNonExisentInput(t *testing.T) {
	filters := map[string][]string{
		"orderby": {"thisdoesntexistasakey"},
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	// Insert the third record into the kvstore
	createRecordThreeResponseMap, err := createRecord(t, kvCollection, recordThree)
	assert.Len(t, createRecordThreeResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 3)

	assert.EqualValues(t, "A", recordsAfterInsert[0]["TEST_KEY_01"])
	assert.EqualValues(t, "B", recordsAfterInsert[1]["TEST_KEY_01"])
	assert.EqualValues(t, "C", recordsAfterInsert[2]["TEST_KEY_01"])

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

// --------
// GET ?fields=count=offset=orderby= parameters
// --------
func TestKVStoreCollectionsListRecordsAllParametersSuccess(t *testing.T) {
	fieldsToFilter := []string{"TEST_KEY_01:0"}
	filters := map[string][]string{
		"fields":  fieldsToFilter,
		"count":   {"1"},
		"offset":  {"1"},
		"orderby": {"TEST_KEY_02"},
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	// Insert the third record into the kvstore
	createRecordThreeResponseMap, err := createRecord(t, kvCollection, recordThree)
	assert.Len(t, createRecordThreeResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)

	assert.EqualValues(t, "B", recordsAfterInsert[0]["TEST_KEY_02"])

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}

//--------
//POST
//--------
//There is no separation for the testing the insertion of a record when using an incorrect collections
//because BOTH are required in order to make a dataset of kvcollection via the catalog service
func TestKVStoreCollectionsInsertRecordIntoMissingCollection(t *testing.T) {
	record := map[string]string{
		"TEST_KEY_01": "TEST_VALUE_01",
		"TEST_KEY_02": "TEST_VALUE_02",
		"TEST_KEY_03": "TEST_VALUE_03",
	}

	// Insert a new record into the kvstore
	createRecordResponseMap, err := getClient(t).KVStoreService.InsertRecord(
		kvCollection,
		record)

	assert.Nil(t, createRecordResponseMap)
	assert.NotNil(t, err)
}

// Inserts a record into the specified tenant's collection
func TestKVStoreCollectionsInsertRecordSuccess(t *testing.T) {
	record := map[string]string{
		"TEST_KEY_01": "TEST_VALUE_01",
		"TEST_KEY_02": "TEST_VALUE_02",
		"TEST_KEY_03": "TEST_VALUE_03",
	}

	// Create the test collection
	dataset, err := createKVCollectionDataset(t,
		testutils.TestNamespace,
		testutils.TestCollection,
		datasetOwner,
		datasetCapabilities)

	// Remove the dataset used for testing
	defer cleanupDatasets(t)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert a new record into the kvstore
	createRecordResponseMap, err := getClient(t).KVStoreService.InsertRecord(
		kvCollection,
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

	// Delete the test collection
	err = getClient(t).CatalogService.DeleteDataset(dataset.ID)
	assert.Nil(t, err)
}
