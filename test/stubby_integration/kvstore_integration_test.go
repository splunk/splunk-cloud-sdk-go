// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package stubbyintegration

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testIndex1 = "TEST_INDEX_01"
var testIndex2 = "TEST_INDEX_02"
var testField1 = "TEST_FIELD_01"
var testField2 = "TEST_FIELD_02"

// Stubby test for GetCollectionStats() kvstore service endpoint
func TestGetCollectionStats(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetCollectionStats(testutils.TestCollection)
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, int64(5), result.Count)
	assert.Equal(t, testutils.TestCollection, result.Ns)
	assert.Equal(t, int64(1), result.Nindexes)
}

// Stubby test for ping/GetServiceHealthStatus() kvstore service endpoint
func TestGetServiceHealthStatus(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetServiceHealthStatus()
	require.Empty(t, err)
	require.NotEmpty(t, result)

	assert.Equal(t, model.PingOKBodyStatusHealthy, result.Status)
}

// Stubby test for ListIndexes() kvstore service endpoint
func TestGetIndexes(t *testing.T) {
	result, err := getClient(t).KVStoreService.ListIndexes(testutils.TestCollection)
	require.Nil(t, err)
	require.NotNil(t, result)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, len(result[0].Fields), 2)
	assert.Equal(t, result[0].Name, testIndex1)
	assert.Equal(t, result[0].Fields[0].Field, testField1)
	assert.Equal(t, result[0].Fields[0].Direction, int64(1))
	assert.Equal(t, result[0].Fields[1].Field, testField2)
	assert.Equal(t, result[0].Fields[1].Direction, int64(1))
}

// Stubby test for CreateIndex() kvstore service endpoint
func TestCreateIndex(t *testing.T) {
	var fields [1]model.IndexFieldDefinition
	fields[0] = CreateField(-1, testField1)
	result, err := getClient(t).KVStoreService.CreateIndex(testutils.TestCollection, CreateIndex(fields[:], testIndex2))
	require.Nil(t, err)
	require.NotEmpty(t, result)
	assert.Equal(t, result.Name, testIndex1)
	assert.Equal(t, result.Collection, testutils.TestCollection)
	assert.Equal(t, result.Fields[0].Field, testField1)
	assert.Equal(t, result.Fields[0].Direction, int64(-1))
}

// Stubby test for DeleteIndex() kvstore service endpoint
func TestDeleteIndex(t *testing.T) {
	err := getClient(t).KVStoreService.DeleteIndex(testutils.TestCollection, testIndex1)
	require.Nil(t, err)
}

// creates an index to post
func CreateIndex(fields []model.IndexFieldDefinition, name string) model.IndexDefinition {
	return model.IndexDefinition{
		Fields: fields,
		Name:   name,
	}
}

// creates an indexField to post
func CreateField(direction int64, field string) model.IndexFieldDefinition {
	return model.IndexFieldDefinition{
		Direction: direction,
		Field:     field,
	}
}

// Stubby test for InsertRecords() kvstore service endpoint
func TestCreateRecords(t *testing.T) {
	var testRecords = `[
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
	var res []model.Record
	err := json.Unmarshal([]byte(testRecords), &res)

	result, err := getClient(t).KVStoreService.InsertRecords(testutils.TestCollection, res)
	assert.Nil(t, err)

	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[0], "TEST_RECORD_KEY_01")
}

// Stubby test for QueryRecords() kvstore service endpoint
func TestGetRecords(t *testing.T) {
	result, err := getClient(t).KVStoreService.QueryRecords(testutils.TestCollection, nil)
	require.Nil(t, err)
	require.NotNil(t, result)

	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0]["_key"], "TEST_RECORD_KEY_01")
	assert.Equal(t, result[1]["_key"], "TEST_RECORD_KEY_02")
}

// Stubby test for QueryRecords() kvstore service endpoint based on a query
func TestGetRecordsWithQuery(t *testing.T) {
	query := make(url.Values)

	query.Add("query", "{\"size\": \"tiny\", \"capacity_gb\": 8}")

	result, err := getClient(t).KVStoreService.QueryRecords(testutils.TestCollection, query)
	require.Nil(t, err)

	require.Equal(t, len(result), 1)
	assert.Equal(t, result[0]["_key"], "TEST_RECORD_KEY_01")
	assert.Equal(t, result[0]["capacity_gb"], float64(8))
	assert.Equal(t, result[0]["description"], "This is a tiny amount of GB")
	assert.Equal(t, result[0]["size"], "tiny")
}

// Stubby test for GetRecordsByKey() kvstore service endpoint
func TestGetRecordByKey(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetRecordByKey(testutils.TestCollection, "TEST_RECORD_KEY_01")
	require.Nil(t, err)

	assert.Equal(t, result["_key"], "TEST_RECORD_KEY_01")
	assert.Equal(t, result["capacity_gb"], float64(8))
	assert.Equal(t, result["description"], "This is a tiny amount of GB")
	assert.Equal(t, result["size"], "tiny")
}

// Stubby test for DeleteRecordsByKey() kvstore service endpoint
func TestDeleteRecordByKey(t *testing.T) {
	err := getClient(t).KVStoreService.DeleteRecordByKey(testutils.TestCollection, "TEST_RECORD_KEY_01")
	require.Nil(t, err)
}

// Stubby test for DeleteRecords() kvstore service endpoint
func TestDeleteRecord(t *testing.T) {
	err := getClient(t).KVStoreService.DeleteRecords(testutils.TestCollection, nil)
	require.Nil(t, err)
}

// Stubby test for DeleteRecords() kvstore service endpoint based on a query
func TestDeleteRecordWithQuery(t *testing.T) {
	query := make(url.Values)
	query.Add("query", "{\"size\": \"tiny\", \"capacity_gb\": 8}")

	err := getClient(t).KVStoreService.DeleteRecords(testutils.TestCollection, query)
	require.Nil(t, err)
}

// Gets records that exist for the tenant's namespace collection
func TestListRecords(t *testing.T) {
	records, err := getClient(t).KVStoreService.ListRecords(
		testutils.TestCollection,
		nil)
	require.Nil(t, err)
	assert.NotNil(t, records)
	assert.Len(t, records, 4)

	for _, element := range records {
		assert.NotNil(t, element)
		for key, value := range element {
			assert.IsType(t, "string", key)
			assert.NotNil(t, value)
		}
	}
}

// Inserts a record into the specified tenant's namespace collection
func TestInsertRecord(t *testing.T) {
	record := map[string]string{
		"TEST_KEY_01": "TEST_VALUE_01",
		"TEST_KEY_02": "TEST_VALUE_02",
		"TEST_KEY_03": "TEST_VALUE_03",
	}

	responseMap, err := getClient(t).KVStoreService.InsertRecord(
		testutils.TestCollection,
		record)

	require.Nil(t, err)
	assert.NotNil(t, responseMap)
	assert.Len(t, responseMap, 1)

	for key, value := range responseMap {
		assert.IsType(t, "string", key)
		assert.Equal(t, "_key", key)

		assert.NotNil(t, value)
		assert.IsType(t, "string", value)
	}
}
