package stubbyintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/splunk/ssc-client-go/testutils"
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

// Gets records that exist for the tenant's namespace collection
func TestListRecords(t *testing.T) {
	records, err := getClient(t).KVStoreService.ListRecords(
		testutils.TestNamespace,
		testutils.TestCollection)
	assert.NotNil(t, records)
	assert.Nil(t, err)
	assert.Len(t, records, 4)

	for _, element := range records {
		assert.NotNil(t, element)
		for key, value := range element {
			assert.IsType(t, "string", key)
			assert.NotNil(t, value)
		}
	}
}