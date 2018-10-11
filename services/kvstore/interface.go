// DO NOT EDIT
package kvstore

import (
	"net/url"
)

type KVstoreIface interface {
	// GetServiceHealthStatus returns Service Health Status
	GetServiceHealthStatus() (*PingOKBody, error)
	// CreateIndex posts a new index to be added to the collection.
	CreateIndex(collectionName string, index IndexDefinition) (*IndexDescription, error)
	// ListIndexes retrieves all the indexes in a given collection
	ListIndexes(collectionName string) ([]IndexDefinition, error)
	// DeleteIndex deletes the specified index in a given collection
	DeleteIndex(collectionName string, indexName string) error
	// InsertRecords posts new records to the collection.
	InsertRecords(collectionName string, records []Record) ([]string, error)
	// QueryRecords queries records present in a given collection.
	QueryRecords(collectionName string, values url.Values) ([]Record, error)
	// GetRecordByKey queries a particular record present in a given collection based on the key value provided by the user.
	GetRecordByKey(collectionName string, keyValue string) (Record, error)
	// DeleteRecords deletes records present in a given collection based on the provided query.
	DeleteRecords(collectionName string, values url.Values) error
	// DeleteRecordByKey deletes a particular record present in a given collection based on the key value provided by the user.
	DeleteRecordByKey(collectionName string, keyValue string) error
	// ListRecords - List the records created for the tenant's specified collection TODO: include count, offset and orderBy
	ListRecords(collectionName string, filters map[string][]string) ([]map[string]interface{}, error)
	// InsertRecord - Create a new record in the tenant's specified collection
	InsertRecord(collectionName string, record Record) (map[string]string, error)
}
