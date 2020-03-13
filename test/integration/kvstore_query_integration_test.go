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
	assert.NoError(t, err)
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
	assert.NoError(t, err)
	assert.Len(t, records, 0)

	// Insert a new record into the kvstore
	createRecord(t, kvCollection, recordOne)

	// Make sure that records return match
	time.Sleep(2 * time.Second)
	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	require.NotNil(t, recordsAfterInsert)
	assert.NoError(t, err)
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

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqfvi")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.NoError(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecord(t, kvCollection, recordOne)

	// Insert the second record into the kvstore
	createRecord(t, kvCollection, recordTwo)

	// Make sure that records return match
	query := kvstore.QueryRecordsQueryParams{}.SetFields([]string{"TEST_KEY_01"})
	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, &query)
	assert.NotNil(t, recordsAfterInsert)
	assert.NoError(t, err)
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

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqfvex")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.NoError(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecord(t, kvCollection, recordOne)

	// Insert the second record into the kvstore
	createRecord(t, kvCollection, recordTwo)

	// Make sure that records return match
	time.Sleep(2 * time.Second)

	query := kvstore.QueryRecordsQueryParams{}.SetFields([]string{"TEST_KEY_01:0"})
	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, &query)
	assert.NotNil(t, recordsAfterInsert)
	assert.NoError(t, err)
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

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvfvinex")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	assert.NotNil(t, records)
	assert.NoError(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecord(t, kvCollection, recordOne)

	// Insert the second record into the kvstore
	createRecord(t, kvCollection, recordTwo)

	query := kvstore.QueryRecordsQueryParams{}.SetFields([]string{"TEST_KEY_01,TEST_KEY_02:0"})
	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, &query)
	assert.Nil(t, recordsAfterInsert)
	assert.NotNil(t, err)
}

// --------
// GET ?count=parameter
// --------
func TestKVStoreQueryCountValidInput(t *testing.T) {

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqcnvin")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.NoError(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecord(t, kvCollection, recordOne)

	// Insert the second record into the kvstore
	createRecord(t, kvCollection, recordTwo)

	query := kvstore.QueryRecordsQueryParams{}.SetCount(1)
	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, &query)
	assert.NotNil(t, recordsAfterInsert)
	assert.NoError(t, err)
	assert.Len(t, recordsAfterInsert, 1)
}

// --------
// GET ?offset=parameter
// --------
func TestKVStoreQueryOffsetValidInput(t *testing.T) {

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqoff")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	assert.NotNil(t, records)
	assert.NoError(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecord(t, kvCollection, recordOne)

	// Insert the second record into the kvstore
	createRecord(t, kvCollection, recordTwo)

	time.Sleep(2 * time.Second)
	query := kvstore.QueryRecordsQueryParams{}.SetOffset(1)
	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, &query)
	assert.NotNil(t, recordsAfterInsert)
	assert.NoError(t, err)
	assert.Len(t, recordsAfterInsert, 1)
}

// --------
// GET ?orderby=parameter
// --------
func TestKVStoreQueryOrderByValidInput(t *testing.T) {

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqob")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)

	assert.NotNil(t, records)
	assert.NoError(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecord(t, kvCollection, recordOne)

	// Insert the second record into the kvstore
	createRecord(t, kvCollection, recordTwo)

	// Insert the third record into the kvstore
	createRecord(t, kvCollection, recordThree)

	time.Sleep(2 * time.Second)

	query := kvstore.QueryRecordsQueryParams{}.SetOrderby([]string{"TEST_KEY_02"})
	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, &query)
	assert.NotNil(t, recordsAfterInsert)
	assert.NoError(t, err)
	assert.Len(t, recordsAfterInsert, 3)

	assert.EqualValues(t, "A", recordsAfterInsert[0]["TEST_KEY_02"])
	assert.EqualValues(t, "B", recordsAfterInsert[1]["TEST_KEY_02"])
	assert.EqualValues(t, "C", recordsAfterInsert[2]["TEST_KEY_02"])
}

//--------
//GET ?query=parameter
//--------
func TestKVStoreQueryQueryParameterInput(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqparami")
	defer cleanupDataset(t, kvid)

	// Make sure that the data set is empty
	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	assert.NotNil(t, records)
	assert.NoError(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecord(t, kvCollection, recordOne)

	// Insert the second record into the kvstore
	createRecord(t, kvCollection, recordTwo)

	// Insert the third record into the kvstore
	createRecord(t, kvCollection, recordThree)

	time.Sleep(2 * time.Second)

	query := kvstore.QueryRecordsQueryParams{}.SetQuery("{\"TEST_KEY_02\":\"A\"}")
	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, &query)
	assert.NotNil(t, recordsAfterInsert)
	assert.NoError(t, err)
	assert.Len(t, recordsAfterInsert, 1)

	// We expect the first result to be recordThree because we're soring by TEST_KEY_02
	assert.EqualValues(t, recordThree["TEST_KEY_02"].(string), recordsAfterInsert[0]["TEST_KEY_02"])
	assert.EqualValues(t, recordThree["TEST_KEY_03"].(string), recordsAfterInsert[0]["TEST_KEY_03"])
	assert.EqualValues(t, recordThree["TEST_KEY_01"].(string), recordsAfterInsert[0]["TEST_KEY_01"])
}

// --------
// GET ?fields=count=offset=orderby=parameters
// --------
func TestKVStoreQueryAllParametersSuccess(t *testing.T) {
	filters := (&kvstore.QueryRecordsQueryParams{}).
		SetFields([]string{"TEST_KEY_01:0"}).
		SetCount(1).SetOffset(1).
		SetOrderby([]string{"TEST_KEY_02"}).
		SetQuery("{\"TEST_KEY_02\":\"A\"}")

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqaparam")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	assert.NotNil(t, records)
	assert.NoError(t, err)
	assert.Len(t, records, 0)

	// Insert the first record into the kvstore
	createRecord(t, kvCollection, recordOne)

	// Insert the second record into the kvstore
	createRecord(t, kvCollection, recordTwo)

	// Insert the third record into the kvstore
	createRecord(t, kvCollection, recordThree)

	recordsAfterInsert, err := getClient(t).KVStoreService.QueryRecords(kvCollection, &filters)
	assert.NotNil(t, recordsAfterInsert)
	assert.NoError(t, err)
	assert.Len(t, recordsAfterInsert, 0)
}

//--------
//POST
//--------
//There is no separation for the testing the insertion of a record when using an incorrect collections
//because BOTH are required in order to make a dataset of kvcollection via the catalog service
func TestKVStoreQueryInsertRecordIntoMissingCollection(t *testing.T) {
	record := map[string]interface{}{
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
	record := map[string]interface{}{
		"TEST_KEY_01": "TEST_VALUE_01",
		"TEST_KEY_02": "TEST_VALUE_02",
		"TEST_KEY_03": "TEST_VALUE_03",
	}

	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvqirec")
	defer cleanupDataset(t, kvid)

	records, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	assert.NotNil(t, records)
	assert.NoError(t, err)
	assert.Len(t, records, 0)

	// Insert a new record into the kvstore
	key, err := getClient(t).KVStoreService.InsertRecord(
		kvCollection,
		record)

	assert.NotNil(t, key)
	assert.NoError(t, err)
	assert.NotEmpty(t, (*key).Key)
}
