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
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)
}

// --------
// GET ?fields= parameter
// --------
func TestKVStoreCollectionsListRecordsReturnsCorrectDatasetAfterSingleInsertRecord(t *testing.T) {
	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert a new record into the kvstore
	createRecordResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordResponseMap, 1)

	// Make sure that records return match
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)
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

func TestKVStoreCollectionsListRecordsFieldsValidInclude(t *testing.T) {
	fieldsToFilter := []string{"TEST_KEY_01"}
	filters := map[string][]string{
		"fields": fieldsToFilter,
	}

	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	// Make sure that records return match
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)
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
}

func TestKVStoreCollectionsListRecordsFieldsValidExclude(t *testing.T) {
	fieldsToFilter := []string{"TEST_KEY_01:0"}
	filters := map[string][]string{
		"fields": fieldsToFilter,
	}

	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	// Make sure that records return match
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)
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
}

func TestKVStoreCollectionsListRecordsFieldsValidIncludeAndExclude(t *testing.T) {
	// From the documenation: A fields value cannot contain both include and exclude specifications except for exclusion
	// of the _key field.
	fieldsToFilter := []string{"TEST_KEY_01,TEST_KEY_02:0"}
	filters := map[string][]string{
		"fields": fieldsToFilter,
	}

	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)
	assert.Nil(t, recordsAfterInsert)
	assert.NotNil(t, err)
}

// --------
// GET ?count= parameter
// --------
func TestKVStoreCollectionsListRecordsCountValidInput(t *testing.T) {
	filters := map[string][]string{
		"count": {"1"},
	}

	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)
}

func TestKVStoreCollectionsListRecordsCountNegativeOutOfBoundsInput(t *testing.T) {
	filters := map[string][]string{
		"count": {"-1"},
	}

	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)

	assert.Nil(t, recordsAfterInsert)
	assert.NotNil(t, err)
}

func TestKVStoreCollectionsListRecordsCountPositiveOutOfBoundsInput(t *testing.T) {
	filters := map[string][]string{
		"count": {"10000"},
	}

	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 2)
}

// --------
// GET ?offset= parameter
// --------
func TestKVStoreCollectionsListRecordsOffsetValidInput(t *testing.T) {
	filters := map[string][]string{
		"offset": {"1"},
	}

	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)
}

func TestKVStoreCollectionsListRecordsOffsetNegativeOutOfBoundsInput(t *testing.T) {
	filters := map[string][]string{
		"offset": {"-1"},
	}

	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)

	assert.Nil(t, recordsAfterInsert)
	assert.NotNil(t, err)
}

func TestKVStoreCollectionsListRecordsOffsetPositiveOutOfBoundsInput(t *testing.T) {
	filters := map[string][]string{
		"offset": {"10000"},
	}

	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 0)
}

// --------
// GET ?orderby= parameter
// --------
func TestKVStoreCollectionsListRecordsOrderByValidInput(t *testing.T) {
	filters := map[string][]string{
		"orderby": {"TEST_KEY_02"},
	}

	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	// Insert the third record into the kvstore
	createRecordThreeResponseMap, err := createRecord(t, testutils.TestCollection, recordThree)
	assert.Len(t, createRecordThreeResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 3)

	assert.EqualValues(t, "A", recordsAfterInsert[0]["TEST_KEY_02"])
	assert.EqualValues(t, "B", recordsAfterInsert[1]["TEST_KEY_02"])
	assert.EqualValues(t, "C", recordsAfterInsert[2]["TEST_KEY_02"])
}

func TestKVStoreCollectionsListRecordsOrderByNonExisentInput(t *testing.T) {
	filters := map[string][]string{
		"orderby": {"thisdoesntexistasakey"},
	}

	// Create the test collection
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	// Insert the third record into the kvstore
	createRecordThreeResponseMap, err := createRecord(t, testutils.TestCollection, recordThree)
	assert.Len(t, createRecordThreeResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 3)

	assert.EqualValues(t, "A", recordsAfterInsert[0]["TEST_KEY_01"])
	assert.EqualValues(t, "B", recordsAfterInsert[1]["TEST_KEY_01"])
	assert.EqualValues(t, "C", recordsAfterInsert[2]["TEST_KEY_01"])
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
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, testutils.TestCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, testutils.TestCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	// Insert the third record into the kvstore
	createRecordThreeResponseMap, err := createRecord(t, testutils.TestCollection, recordThree)
	assert.Len(t, createRecordThreeResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)

	assert.EqualValues(t, "B", recordsAfterInsert[0]["TEST_KEY_02"])
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
		testutils.TestCollection,
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
	createKVCollection(t, testutils.TestCollection)

	// Remove the test collection used for testing
	defer cleanupCollection(t)

	records, err := getClient(t).KVStoreService.ListRecords(testutils.TestCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert a new record into the kvstore
	createRecordResponseMap, err := getClient(t).KVStoreService.InsertRecord(
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
