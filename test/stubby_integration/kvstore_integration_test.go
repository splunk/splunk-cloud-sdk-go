package stubbyintegration

import (
	"encoding/json"
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

var namespaceName = "test_namespace"
var collectionName = "test_collection"

// Stubby test for GetCollectionStats() kvstore service endpoint
func TestGetCollectionStats(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetCollectionStats("test_namespace", "test_collection")
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, int64(5), result.Count)
	assert.Equal(t, "test_namespace", result.Ns)
	assert.Equal(t, int64(1), result.Nindexes)
}

// Stubby test for ping/GetServiceHealthStatus() kvstore service endpoint
func TestGetServiceHealthStatus(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetServiceHealthStatus()
	assert.Empty(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, "healthy", result.Status)
}

// Stubby test for CreateRecords() kvstore service endpoint
func TestCreateRecords(t *testing.T) {
	var obj1 = `[{ "capacity_gb": 8, "size": "tiny", "description": "This is a tiny amount of GB", "_raw": ""} ,{"capacity_gb": 16,"size": "small","description": "This is a small amount of GB","_raw": ""}, {"type": "A","name": "test_record","count_of_fields": 3}]`
	var res []model.Record
	err := json.Unmarshal([]byte(obj1), &res)

	result, err := getClient(t).KVStoreService.CreateRecords(namespaceName, collectionName, res)
	assert.Nil(t, err)
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[0], "5b22a2c74e2ed20001b345fb")
}

// Stubby test for GetRecords() kvstore service endpoint based on a query
func TestGetRecordsWithQuery(t *testing.T) {
	query := make(url.Values)
	query.Add("size", "tiny")
	query.Add("capacity_gb", "8")
	result, err := getClient(t).KVStoreService.GetRecords(query, namespaceName, collectionName)
	assert.Nil(t, err)

	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0]["_key"], "5b22a2c74e2ed20001b345fb")
	assert.Equal(t, result[0]["capacity_gb"], float64(8))
	assert.Equal(t, result[0]["description"], "This is a tiny amount of GB")
	assert.Equal(t, result[0]["size"], "tiny")
}

/*func TestGetAllRecords(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetRecords(nil, namespaceName, collectionName)
	fmt.Println(result)
	fmt.Println(err)
	// [map[size:tiny _key:5b22a2c74e2ed20001b345fb _raw: capacity_gb:8 description:This is a tiny amount of GB] map[_key:5b22a2c74e2ed20001b346fb _raw: capacity_gb:16 description:This is a small amount of GB size:small] map[count_of_fields:3 _key:5b22a2c74e2ed20001b347fb type:A name:test_record]]
}*/

// Stubby test for GetRecords() kvstore service endpoint based on a key
func TestGetRecordByKey(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetRecordByKey(namespaceName, collectionName, "5b22a2c74e2ed20001b345fb")
	assert.Nil(t, err)
	assert.Equal(t, result["_key"], "5b22a2c74e2ed20001b345fb")
	assert.Equal(t, result["capacity_gb"], float64(8))
	assert.Equal(t, result["description"], "This is a tiny amount of GB")
	assert.Equal(t, result["size"], "tiny")
}

// Stubby test for DeleteRecords() kvstore service endpoint based on a key
func TestDeleteRecordByKey(t *testing.T) {
	err := getClient(t).KVStoreService.DeleteRecordByKey(namespaceName, collectionName, "5b22a2c74e2ed20001b345fb")
	assert.Nil(t, err)
}

// Stubby test for Delete Records() kvstore service endpoint based on a query TODO: Delete record based on a query doesn't work currently
func TestDeleteRecord(t *testing.T) {
	query := make(url.Values)
	query.Add("size", "tiny")
	query.Add("capacity_gb", "8")
	outerQuery := make(url.Values)
	outerQuery.Add("query", query.Encode())

	err := getClient(t).KVStoreService.DeleteRecords(outerQuery, namespaceName, collectionName)
	assert.Nil(t, err)
}
