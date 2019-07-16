/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package integration

import (
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/services/kvstore"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var recordOne = map[string]interface{}{
	"TEST_KEY_01": "A",
	"TEST_KEY_02": "B",
	"TEST_KEY_03": "C",
}

var recordTwo = map[string]interface{}{
	"TEST_KEY_01": "B",
	"TEST_KEY_02": "C",
	"TEST_KEY_03": "A",
}
var recordThree = map[string]interface{}{
	"TEST_KEY_01": "C",
	"TEST_KEY_02": "A",
	"TEST_KEY_03": "B",
}

func createRecord(t *testing.T, collection string, record map[string]interface{}) (kvstore.Key, error) {
	// Insert a new record into the kvstore
	key, err := getClient(t).KVStoreService.InsertRecord(
		collection,
		record)

	assert.NotNil(t, key)
	assert.Nil(t, err)
	assert.NotEmpty(t, (*key).Key)

	return *key, err
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
	kvid, kvCollection := makeCollectionName(t, "kvlr")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	require.NotNil(t, records)
	require.Nil(t, err)
	assert.Len(t, records, 0)
}

// --------
// GET ?fields= parameter
// --------
func TestKVStoreCollectionsListRecordsReturnsCorrectDatasetAfterSingleInsertRecord(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvlrrds")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert a new record into the kvstore
	key, _ := createRecord(t, kvCollection, recordOne)
	assert.NotEmpty(t, key.Key)

	// Make sure that records return match
	time.Sleep(2 * time.Second)
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
}

func TestKVStoreCollectionsListRecordsFieldsValidInclude(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvrfvi")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	key, _ := createRecord(t, kvCollection, recordOne)
	assert.NotEmpty(t, key.Key)

	// Insert the second record into the kvstore
	key, _ = createRecord(t, kvCollection, recordTwo)
	assert.NotEmpty(t, key.Key)

	// Make sure that records return match
	query := kvstore.ListRecordsQueryParams{}.SetFields([]string{"TEST_KEY_01"})
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, &query)
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
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvrfve")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	key, _ := createRecord(t, kvCollection, recordOne)
	assert.NotEmpty(t, key.Key)

	// Insert the second record into the kvstore
	key, _ = createRecord(t, kvCollection, recordTwo)
	assert.NotEmpty(t, key.Key)

	// Make sure that records return match
	query := kvstore.ListRecordsQueryParams{}.SetFields([]string{"TEST_KEY_01:0"})
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, &query)
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
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvrfvx")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)
	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	key, _ := createRecord(t, kvCollection, recordOne)
	assert.NotEmpty(t, key.Key)

	// Insert the second record into the kvstore
	key, _ = createRecord(t, kvCollection, recordTwo)
	assert.NotEmpty(t, key.Key)

	query := kvstore.ListRecordsQueryParams{}.SetFields([]string{"TEST_KEY_01,TEST_KEY_02:0"})
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, &query)
	assert.Nil(t, recordsAfterInsert)
	assert.NotNil(t, err)
}

// --------
// GET ?count= parameter
// --------
func TestKVStoreCollectionsListRecordsCountValidInput(t *testing.T) {

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvlrcvin")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)
	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	key, _ := createRecord(t, kvCollection, recordOne)
	assert.NotEmpty(t, key.Key)
	// Insert the second record into the kvstore
	key, _ = createRecord(t, kvCollection, recordTwo)
	assert.NotEmpty(t, key.Key)

	query := kvstore.ListRecordsQueryParams{}.SetCount(1)
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, &query)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)
}

// --------
// GET ?offset= parameter
// --------
func TestKVStoreCollectionsListRecordsOffsetValidInput(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvlrfvi")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)
	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	key, _ := createRecord(t, kvCollection, recordOne)
	assert.NotEmpty(t, key.Key)

	// Insert the second record into the kvstore
	key, _ = createRecord(t, kvCollection, recordTwo)
	assert.NotEmpty(t, key.Key)

	time.Sleep(2 * time.Second)
	query := kvstore.ListRecordsQueryParams{}.SetOffset(1)
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, &query)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)
}

// --------
// GET ?orderby= parameter
// --------
func TestKVStoreCollectionsListRecordsOrderByValidInput(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvlrorbvi")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	key, _ := createRecord(t, kvCollection, recordOne)
	assert.NotEmpty(t, key.Key)

	// Insert the second record into the kvstore
	key, _ = createRecord(t, kvCollection, recordTwo)
	assert.NotEmpty(t, key.Key)

	// Insert the third record into the kvstore
	key, _ = createRecord(t, kvCollection, recordThree)
	assert.NotEmpty(t, key.Key)

	time.Sleep(2 * time.Second)

	query := kvstore.ListRecordsQueryParams{}.SetFields([]string{"TEST_KEY_02"})
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, &query)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 3)

	assert.EqualValues(t, "B", recordsAfterInsert[0]["TEST_KEY_02"])
	assert.EqualValues(t, "C", recordsAfterInsert[1]["TEST_KEY_02"])
	assert.EqualValues(t, "A", recordsAfterInsert[2]["TEST_KEY_02"])
}

// --------
// GET ?fields=count=offset=orderby= parameters
// --------
func TestKVStoreCollectionsListRecordsAllParametersSuccess(t *testing.T) {
	fieldsToFilter := "TEST_KEY_01:0"
	filters := (&kvstore.ListRecordsQueryParams{}).
		SetFields([]string{fieldsToFilter}).SetCount(1).SetOffset(1).SetOrderby([]string{"TEST_KEY_02"})

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvlrpms")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)
	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	key, _ := createRecord(t, kvCollection, recordOne)
	assert.NotEmpty(t, key.Key)

	// Insert the second record into the kvstore
	key, _ = createRecord(t, kvCollection, recordTwo)
	assert.NotEmpty(t, key.Key)

	// Insert the third record into the kvstore
	key, _ = createRecord(t, kvCollection, recordThree)
	assert.NotEmpty(t, key.Key)

	time.Sleep(2 * time.Second)
	recordsAfterInsert, err := getClient(t).KVStoreService.ListRecords(kvCollection, &filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.Nil(t, err)
	assert.Len(t, recordsAfterInsert, 1)

	// The order before filtering should be: recordThree, recordOne, recordTwo
	// Since we offset by 1, the expect recordOne to be the first element in recordsAfterInsert
	assert.EqualValues(t, recordOne["TEST_KEY_02"].(string), recordsAfterInsert[0]["TEST_KEY_02"])
}

//--------
//POST
//--------
//There is no separation for the testing the insertion of a record when using an incorrect collections
//because BOTH are required in order to make a dataset of kvcollection via the catalog service
func TestKVStoreCollectionsInsertRecordIntoMissingCollection(t *testing.T) {
	record := map[string]interface{}{
		"TEST_KEY_01": "TEST_VALUE_01",
		"TEST_KEY_02": "TEST_VALUE_02",
		"TEST_KEY_03": "TEST_VALUE_03",
	}

	// Insert a new record into the kvstore
	_, err := getClient(t).KVStoreService.InsertRecord(
		"idonotexist",
		record)

	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 404, httpErr.HTTPStatusCode)
}

// Inserts a record into the specified tenant's collection
func TestKVStoreCollectionsInsertRecordSuccess(t *testing.T) {
	record := map[string]interface{}{
		"TEST_KEY_01": "TEST_VALUE_01",
		"TEST_KEY_02": "TEST_VALUE_02",
		"TEST_KEY_03": "TEST_VALUE_03",
	}

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvir")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.ListRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 0)

	// Insert a new record into the kvstore
	key, err := getClient(t).KVStoreService.InsertRecord(
		kvCollection,
		record)

	assert.NotNil(t, key)
	assert.Nil(t, err)
	assert.NotEmpty(t, key.Key)
}
