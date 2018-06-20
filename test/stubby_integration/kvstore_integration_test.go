package stubbyintegration

import (
	"encoding/json"
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

var testIndex1 = "TEST_INDEX_01"
var testIndex2 = "TEST_INDEX_02"
var testField1 = "TEST_FIELD_01"
var testField2 = "TEST_FIELD_02"

// Stubby test for GetCollectionStats() kvstore service endpoint
func TestGetCollectionStats(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetCollectionStats(testutils.TestNamespace, testutils.TestCollection)
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, int64(5), result.Count)
	assert.Equal(t, testutils.TestNamespace, result.Ns)
	assert.Equal(t, int64(1), result.Nindexes)
}

// Stubby test for ping/GetServiceHealthStatus() kvstore service endpoint
func TestGetServiceHealthStatus(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetServiceHealthStatus()
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, "healthy", result.Status)
}

// Stubby test for ListIndexes() kvstore service endpoint
func TestGetIndexes(t *testing.T) {
	result, err := getClient(t).KVStoreService.ListIndexes(testutils.TestNamespace, testutils.TestCollection)
	assert.NotNil(t, result)
	assert.Equal(t, len(result), 2)
	assert.Equal(t, len(result[0].Fields), 2)
	assert.Equal(t, result[0].Name, testIndex1)
	assert.Equal(t, result[0].Fields[0].Field, testField1)
	assert.Equal(t, result[0].Fields[0].Direction, int64(1))
	assert.Equal(t, result[0].Fields[1].Field, testField2)
	assert.Equal(t, result[0].Fields[1].Direction, int64(1))
	assert.Nil(t, err)
}

// Stubby test for CreateIndex() kvstore service endpoint
func TestCreateIndex(t *testing.T) {
	var fields [1]model.IndexFieldDefinition
	fields[0] = CreateField(-1, testField1)
	err := getClient(t).KVStoreService.CreateIndex(CreateIndex(testutils.TestCollection, fields[:], testIndex2, testutils.TestNamespace), testutils.TestNamespace, testutils.TestCollection)
	assert.Nil(t, err)
}

// Stubby test for DeleteIndex() kvstore service endpoint
func TestDeleteIndex(t *testing.T) {
	err := getClient(t).KVStoreService.DeleteIndex(testutils.TestNamespace, testutils.TestCollection, testIndex1)
	assert.Nil(t, err)
}

// creates an index to post
func CreateIndex(collection string, fields []model.IndexFieldDefinition, name string, namespace string) model.IndexDescription {
	return model.IndexDescription{
		Collection: collection,
		Fields:     fields,
		Name:       name,
		Namespace:  namespace,
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
	var testRecords = `[{ "capacity_gb": 8, "size": "tiny", "description": "This is a tiny amount of GB", "_raw": ""} ,{"capacity_gb": 16,"size": "small","description": "This is a small amount of GB","_raw": ""}, {"type": "A","name": "test_record","count_of_fields": 3}]`
	var res []model.Record
	err := json.Unmarshal([]byte(testRecords), &res)

	result, err := getClient(t).KVStoreService.InsertRecords(testutils.TestNamespace, testutils.TestCollection, res)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[0], "TEST_RECORD_KEY_01")
}

// Stubby test for QueryRecords() kvstore service endpoint
func TestGetRecordsWithQuery(t *testing.T) {
	query := make(url.Values)
	query.Add("size", "tiny")
	query.Add("capacity_gb", "8")
	result, err := getClient(t).KVStoreService.QueryRecords(query, testutils.TestNamespace, testutils.TestCollection)
	assert.Nil(t, err)

	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0]["_key"], "TEST_RECORD_KEY_01")
	assert.Equal(t, result[0]["capacity_gb"], float64(8))
	assert.Equal(t, result[0]["description"], "This is a tiny amount of GB")
	assert.Equal(t, result[0]["size"], "tiny")
}

// Stubby test for GetRecordsByKey() kvstore service endpoint
func TestGetRecordByKey(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetRecordByKey(testutils.TestNamespace, testutils.TestCollection, "TEST_RECORD_KEY_01")
	assert.Nil(t, err)
	assert.Equal(t, result["_key"], "TEST_RECORD_KEY_01")
	assert.Equal(t, result["capacity_gb"], float64(8))
	assert.Equal(t, result["description"], "This is a tiny amount of GB")
	assert.Equal(t, result["size"], "tiny")
}

// Stubby test for DeleteRecordsByKey() kvstore service endpoint
func TestDeleteRecordByKey(t *testing.T) {
	err := getClient(t).KVStoreService.DeleteRecordByKey(testutils.TestNamespace, testutils.TestCollection, "TEST_RECORD_KEY_01")
	assert.Nil(t, err)
}

// Stubby test for DeleteRecords() kvstore service endpoint based on a query
func TestDeleteRecord(t *testing.T) {
	query := make(url.Values)
	query.Add("size", "tiny")
	query.Add("capacity_gb", "8")
	outerQuery := make(url.Values)
	outerQuery.Add("query", query.Encode())

	err := getClient(t).KVStoreService.DeleteRecords(outerQuery, testutils.TestNamespace, testutils.TestCollection)
	assert.Nil(t, err)
}
