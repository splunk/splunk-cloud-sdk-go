// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"net/url"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --------------------------------------------------------------------------------
// Query Endpoints
// --------------------------------------------------------------------------------
// /TENANT_NAME/kvstore/v2/collections/COLLECTION_NAME/query

// --------
// GET
// --------
func TestKVStoreQueryReturnsEmptyDatasetOnCreation(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqred")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)

	require.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)
}

// --------
// GET ?fields=parameter
// --------
func TestKVStoreQueryReturnsCorrectDatasetAfterSingleInsertRecord(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqasir")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)

	require.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert a new record into the kvstore
	createRecordResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordResponseMap, 1)

	// Make sure that records return match
	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	require.NotNil(t, recordsAfterInsert)
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

func TestKVStoreQueryFieldsValidInclude(t *testing.T) {
	fieldsToFilter := []string{"TEST_KEY_01"}
	filters := map[string][]string{
		"fields": fieldsToFilter,
	}

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqfvi")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)

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
	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, filters)
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

func TestKVStoreQueryFieldsValidExclude(t *testing.T) {
	fieldsToFilter := []string{"TEST_KEY_01:0"}
	filters := map[string][]string{
		"fields": fieldsToFilter,
	}

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqfvex")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)

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
	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, filters)
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

func TestKVStoreQueryFieldsValidIncludeAndExclude(t *testing.T) {
	// From the documenation: A fields value cannot contain both include and exclude specifications except for exclusion
	// of the _key field.
	fieldsToFilter := []string{"TEST_KEY_01,TEST_KEY_02:0"}
	filters := map[string][]string{
		"fields": fieldsToFilter,
	}

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvfvinex")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, filters)
	assert.Nil(t, recordsAfterInsert)
	assert.NotNil(t, err)
}

// --------
// GET ?count=parameter
// --------
func TestKVStoreQueryCountValidInput(t *testing.T) {
	filters := map[string][]string{
		"count": {"1"},
	}

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqcnvin")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)
}

// --------
// GET ?offset=parameter
// --------
func TestKVStoreQueryOffsetValidInput(t *testing.T) {
	filters := map[string][]string{
		"offset": {"1"},
	}

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqoff")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecordOneResponseMap, err := createRecord(t, kvCollection, recordOne)
	assert.Len(t, createRecordOneResponseMap, 1)

	// Insert the second record into the kvstore
	createRecordTwoResponseMap, err := createRecord(t, kvCollection, recordTwo)
	assert.Len(t, createRecordTwoResponseMap, 1)

	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)
}

// --------
// GET ?orderby=parameter
// --------
func TestKVStoreQueryOrderByValidInput(t *testing.T) {
	filters := map[string][]string{
		"orderby": {"TEST_KEY_02"},
	}

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqob")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)

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

	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 3)

	assert.EqualValues(t, "A", recordsAfterInsert[0]["TEST_KEY_02"])
	assert.EqualValues(t, "B", recordsAfterInsert[1]["TEST_KEY_02"])
	assert.EqualValues(t, "C", recordsAfterInsert[2]["TEST_KEY_02"])
}

//--------
//GET ?query=parameter
//--------
func TestKVStoreQueryQueryParameterInput(t *testing.T) {
	filters := url.Values{
		"query": {"{\"TEST_KEY_02\":\"A\"}"},
	}

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqparami")
	defer cleanupDataset(t, kvid)

	// Make sure that the data set is empty
	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
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

	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)

	assert.EqualValues(t, "A", recordsAfterInsert[0]["TEST_KEY_02"])
	assert.EqualValues(t, "B", recordsAfterInsert[0]["TEST_KEY_03"])
	assert.EqualValues(t, "C", recordsAfterInsert[0]["TEST_KEY_01"])
}

// --------
// GET ?fields=count=offset=orderby=parameters
// --------
func TestKVStoreQueryAllParametersSuccess(t *testing.T) {
	fieldsToFilter := []string{"TEST_KEY_01:0"}
	filters := map[string][]string{
		"fields":  fieldsToFilter,
		"count":   {"1"},
		"offset":  {"1"},
		"orderby": {"TEST_KEY_02"},
		"query":   {"{\"TEST_KEY_02\":\"A\"}"},
	}

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqaparam")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
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

	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 0)
}

//--------
//POST
//--------
//There is no separation for the testing the insertion of a record when using an incorrect collections
//because BOTH are required in order to make a dataset of kvcollection via the catalog service
func TestKVStoreQueryInsertRecordIntoMissingCollection(t *testing.T) {
	record := model.Record{
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
func TestKVStoreQueryInsertRecordSuccess(t *testing.T) {
	record := model.Record{
		"TEST_KEY_01": "TEST_VALUE_01",
		"TEST_KEY_02": "TEST_VALUE_02",
		"TEST_KEY_03": "TEST_VALUE_03",
	}

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqirec")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
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
}
