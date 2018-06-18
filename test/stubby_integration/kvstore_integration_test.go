package stubbyintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"net/url"
	"encoding/json"
)

var testNamespace = "TEST_NAMESPACE"
var testCollection = "TEST_COLLECTION"
var testIndex1 = "TEST_INDEX_01"
var testIndex2 = "TEST_INDEX_02"
var testField1 = "TEST_FIELD_01"
var testField2 = "TEST_FIELD_02"

// Stubby test for GetCollectionStats() kvstore service endpoint
func TestGetCollectionStats(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetCollectionStats(testNamespace, testCollection)
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, int64(5), result.Count)
	assert.Equal(t, testNamespace, result.Ns)
	assert.Equal(t, int64(1), result.Nindexes)
}

// Stubby test for ping/GetServiceHealthStatus() kvstore service endpoint
func TestGetServiceHealthStatus(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetServiceHealthStatus()
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, "healthy", result.Status)
}

// Stubby test for GetIndexes() kvstore service endpoint
func TestGetIndexes(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetIndexes(testNamespace, testCollection)
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
	err := getClient(t).KVStoreService.CreateIndex(CreateIndex(testCollection, fields[:], testIndex2, testNamespace), testNamespace, testCollection)
	assert.Nil(t, err)
}

// Stubby test for DeleteIndex() kvstore service endpoint
func TestDeleteIndex(t *testing.T) {
	err := getClient(t).KVStoreService.DeleteIndex(testNamespace, testCollection, testIndex1)
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

// Stubby test for CreateRecords() kvstore service endpoint
func TestCreateRecords(t *testing.T) {
	var obj1 = `[{ "capacity_gb": 8, "size": "tiny", "description": "This is a tiny amount of GB", "_raw": ""} ,{"capacity_gb": 16,"size": "small","description": "This is a small amount of GB","_raw": ""}, {"type": "A","name": "test_record","count_of_fields": 3}]`
	var res []model.Record
	err := json.Unmarshal([]byte(obj1), &res)

	result, err := getClient(t).KVStoreService.CreateRecords(testNamespace, testCollection, res)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[0], "TEST_RECORD_KEY_01")
}

// Stubby test for GetRecords() kvstore service endpoint based on a query
func TestGetRecordsWithQuery(t *testing.T) {
	query := make(url.Values)
	query.Add("size", "tiny")
	query.Add("capacity_gb", "8")
	result, err := getClient(t).KVStoreService.GetRecords(query, testNamespace, testCollection)
	assert.Nil(t, err)

	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0]["_key"], "TEST_RECORD_KEY_01")
	assert.Equal(t, result[0]["capacity_gb"], float64(8))
	assert.Equal(t, result[0]["description"], "This is a tiny amount of GB")
	assert.Equal(t, result[0]["size"], "tiny")
}

/*func TestGetAllRecords(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetRecords(nil, testNamespace, testCollection)
	fmt.Println(result)
	fmt.Println(err)
	// [map[size:tiny _key:5b22a2c74e2ed20001b345fb _raw: capacity_gb:8 description:This is a tiny amount of GB] map[_key:5b22a2c74e2ed20001b346fb _raw: capacity_gb:16 description:This is a small amount of GB size:small] map[count_of_fields:3 _key:5b22a2c74e2ed20001b347fb type:A name:test_record]]
}*/

// Stubby test for GetRecords() kvstore service endpoint based on a key
func TestGetRecordByKey(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetRecordByKey(testNamespace, testCollection, "TEST_RECORD_KEY_01")
	assert.Nil(t, err)
	assert.Equal(t, result["_key"], "TEST_RECORD_KEY_01")
	assert.Equal(t, result["capacity_gb"], float64(8))
	assert.Equal(t, result["description"], "This is a tiny amount of GB")
	assert.Equal(t, result["size"], "tiny")
}

// Stubby test for DeleteRecords() kvstore service endpoint based on a key
func TestDeleteRecordByKey(t *testing.T) {
	err := getClient(t).KVStoreService.DeleteRecordByKey(testNamespace, testCollection, "TEST_RECORD_KEY_01")
	assert.Nil(t, err)
}

// Stubby test for Delete Records() kvstore service endpoint based on a query
func TestDeleteRecord(t *testing.T) {
	query := make(url.Values)
	query.Add("size", "tiny")
	query.Add("capacity_gb", "8")
	outerQuery := make(url.Values)
	outerQuery.Add("query", query.Encode())

	err := getClient(t).KVStoreService.DeleteRecords(outerQuery, testNamespace, testCollection)
	assert.Nil(t, err)
}
