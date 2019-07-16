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
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/services/kvstore"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test variables
var (
	testIndex = "integtestindex"
)

func makeCollectionName(t *testing.T, ctx string) (id, kvCollection string) {
	dsname := makeDSName(ctx)
	kvds, err := createKVCollectionDataset(t, dsname)
	require.Nil(t, err)
	return kvds.KvCollectionDataset().Id, fmt.Sprintf("%s.%s", kvds.KvCollectionDataset().Module, kvds.KvCollectionDataset().Name)
}

// --------------------------------------------------------------------------------
// Admin Endpoints
// --------------------------------------------------------------------------------

// Test GetServiceHealthStatus
func TestIntegrationGetServiceHealth(t *testing.T) {
	response, err := getClient(t).KVStoreService.Ping()
	require.Empty(t, err)
	assert.NotEmpty(t, response)
	assert.Equal(t, kvstore.PingResponseStatusHealthy, response.Status)
}

// --------------------------------------------------------------------------------
// Index Endpoints
// --------------------------------------------------------------------------------
// /TENANT_NAME/kvstore/v2/collections/COLLECTION_NAME/indexes

// Test CreateIndex, ListIndexes and DeleteIndex kvstore endpoints
func TestIntegrationIndexEndpoints(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvidx")
	defer cleanupDataset(t, kvid)

	// Create Index
	var fields [1]kvstore.IndexFieldDefinition
	fields[0] = kvstore.IndexFieldDefinition{Direction: -1, Field: "integ_testField1"}
	indexDescription, err := getClient(t).KVStoreService.CreateIndex(kvCollection,
		kvstore.IndexDefinition{
			Name:   testIndex,
			Fields: fields[:]})
	require.Nil(t, err)
	require.NotEmpty(t, indexDescription)
	assert.Equal(t, *indexDescription.Collection, kvCollection)

	// Validate if the index was created
	time.Sleep(2 * time.Second)
	indexes, err := getClient(t).KVStoreService.ListIndexes(kvCollection)
	require.Nil(t, err)
	require.NotNil(t, indexes)
	assert.Equal(t, len(indexes), 1)
	assert.Equal(t, indexes[0].Name, testIndex)

	// Delete the test index
	err = getClient(t).KVStoreService.DeleteIndex(kvCollection, testIndex)
	assert.Nil(t, err)

	// Validate if the index was deleted
	result, err := getClient(t).KVStoreService.ListIndexes(kvCollection)
	assert.Nil(t, err)
	require.NotNil(t, result)
	assert.Equal(t, len(result), 0)
}

// Test CreateIndex for 422 Unprocessable Entity error
func TestIntegrationCreateIndexUnprocessableEntityError(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvidx422")
	defer cleanupDataset(t, kvid)

	// Create Index
	_, err := getClient(t).KVStoreService.CreateIndex(kvCollection, kvstore.IndexDefinition{Name: testIndex, Fields: nil})
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.Equal(t, 422, httpErr.HTTPStatusCode)
	assert.Equal(t, "fields in body is required", httpErr.Message)
}

// Test CreateIndex for 404 Not Found error TODO: Change name of non existing collection
func TestIntegrationCreateIndexNonExistingCollection(t *testing.T) {
	// Create the test collection
	kvid, _ := makeCollectionName(t, "kvidx404")
	defer cleanupDataset(t, kvid)

	// Create Index
	var fields [1]kvstore.IndexFieldDefinition
	fields[0] = kvstore.IndexFieldDefinition{Direction: -1, Field: "integ_testField1"}
	_, err := getClient(t).KVStoreService.CreateIndex(testutils.TestCollection, kvstore.IndexDefinition{Name: testIndex, Fields: fields[:]})
	require.NotNil(t, err)
	httpErr, ok := err.(*util.HTTPError)
	require.True(t, ok)
	assert.EqualValues(t, 404, httpErr.HTTPStatusCode)
	// Known bug: should actually provide collection name - see https://jira.splunk.com/browse/SSC-5084
	assert.EqualValues(t, "collection not found", httpErr.Message)
}

// Test DeleteIndex for 404 Index not found error
func TestIntegrationDeleteNonExistingIndex(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvidx404")
	defer cleanupDataset(t, kvid)

	// DeleteIndex
	err := getClient(t).KVStoreService.DeleteIndex(kvCollection, testIndex)
	require.Nil(t, err)
}

// --------------------------------------------------------------------------------
// Record Endpoints
// --------------------------------------------------------------------------------
// /TENANT_NAME/kvstore/v2/collections/COLLECTION_NAME

// Test InsertRecords() kvstore service endpoint
func TestCreateRecords(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvcrr")
	defer cleanupDataset(t, kvid)

	createTestRecord(t, kvCollection)
}

// Test InsertRecords() kvstore service endpoint
func TestPutRecords(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvpr")
	defer cleanupDataset(t, kvid)

	keys := createTestRecord(t, kvCollection)

	record := `
	{
		"capacity_gb": 8,
		"size": "notbig",
		"description": "this is a notbig amount of GB",
		"_raw": ""
	}`

	var res map[string]interface{}
	err := json.Unmarshal([]byte(record), &res)
	require.Nil(t, err)

	// test replace record
	key, err := getClient(t).KVStoreService.PutRecord(kvCollection, keys[0], res)
	require.Nil(t, err)
	require.NotNil(t, key)
	assert.Equal(t, (*key).Key, keys[0])

	// test insert record
	recordID := "recordID"
	key, err = getClient(t).KVStoreService.PutRecord(kvCollection, recordID, res)
	require.Nil(t, err)
	require.NotNil(t, key)
	assert.Equal(t, (*key).Key, recordID)
}

// Test getRecordByKey() kvstore service endpoint
func TestGetRecordByKey(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvgbk")
	defer cleanupDataset(t, kvid)

	keys := createTestRecord(t, kvCollection)

	result, err := getClient(t).KVStoreService.GetRecordByKey(kvCollection, keys[0])

	require.Nil(t, err)
	require.NotNil(t, result)
	assert.NotNil(t, (*result)["_key"])
	assert.Equal(t, (*result)["capacity_gb"], float64(8))
	assert.Equal(t, (*result)["description"], "This is a tiny amount of GB")
	assert.Equal(t, (*result)["size"], "tiny")
}

// Test DeleteRecords() kvstore service endpoint based on a key
func TestDeleteRecordByKey(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvdrbk")
	defer cleanupDataset(t, kvid)

	keys := createTestRecord(t, kvCollection)

	// Delete record by key
	err := getClient(t).KVStoreService.DeleteRecordByKey(kvCollection, keys[0])
	require.Nil(t, err)
	time.Sleep(2 * time.Second)

	// Validate that the record has been deleted
	retrievedRecordsByKey, err := getClient(t).KVStoreService.GetRecordByKey(kvCollection, keys[0])
	assert.Nil(t, retrievedRecordsByKey)

	retrievedRecords, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	require.NotNil(t, retrievedRecords)
	assert.Equal(t, len(retrievedRecords), 2)
}

// Test DeleteRecords() kvstore service endpoint based on a query
func TestDeleteRecord(t *testing.T) {
	// Create the test collection
	kvid, kvCollection := makeCollectionName(t, "kvdrec")
	defer cleanupDataset(t, kvid)

	// Create records
	createTestRecord(t, kvCollection)

	// Create query to test delete operation
	query := kvstore.DeleteRecordsQueryParams{}.SetQuery(`{"capacity_gb": 16}`)
	err := getClient(t).KVStoreService.DeleteRecords(kvCollection, &query)
	require.Nil(t, err)

	// Validate that the record has been deleted
	time.Sleep(2 * time.Second)
	retrievedRecords, err := getClient(t).KVStoreService.QueryRecords(kvCollection, nil)
	require.NotNil(t, retrievedRecords)
	assert.Equal(t, len(retrievedRecords), 2)
}

// Create test record
func createTestRecord(t *testing.T, kvCollection string) []string {
	var integrationTestRecord = `[
         {
          "capacity_gb": 8,
          "size": "tiny",
          "description": "This is a tiny amount of GB",
          "_raw": ""
         },
         {
          "capacity_gb": 16,
          "size": "small",
          "description": "This is a small amount of GB",
          "_raw": ""
         },
         {
          "type": "A",
          "name": "test_record",
          "count_of_fields": 3
         }
        ]`
	var res []map[string]interface{}
	err := json.Unmarshal([]byte(integrationTestRecord), &res)
	require.Nil(t, err)

	keys, err := getClient(t).KVStoreService.InsertRecords(kvCollection, res)
	require.Nil(t, err)
	require.NotNil(t, keys)
	assert.Equal(t, len(keys), 3)

	return keys
}
