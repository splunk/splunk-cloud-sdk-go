package stubbyintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testNamespace = "test_namespace"
var testCollection = "test_collection"
var testIndex1 = "test_index1"
var testIndex2 = "test_index2"
var testField1 = "test_field1"
var testField2 = "test_field2"

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