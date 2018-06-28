package service

import (
	"net/url"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

const kvStoreServicePrefix = "kvstore"
const kvStoreServiceVersion = "v1"
const kvStoreCollectionsResource = "collections"

// KVStoreService talks to kvstore service
type KVStoreService service

// GetCollectionStats returns Collection Stats for the collection
func (c *KVStoreService) GetCollectionStats(namespace string, collection string) (*model.CollectionStats, error) {
	url, err := c.client.BuildURL(nil, kvStoreServicePrefix, kvStoreServiceVersion, namespace, kvStoreCollectionsResource, collection, "stats")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(url)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.CollectionStats
	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetServiceHealthStatus returns Service Health Status
func (c *KVStoreService) GetServiceHealthStatus() (*model.PingOKBody, error) {
	url, err := c.client.BuildURL(nil, kvStoreServicePrefix, kvStoreServiceVersion, "ping")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(url)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.PingOKBody
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateIndex posts a new index to be added to the collection.
func (c *KVStoreService) CreateIndex(index model.IndexDescription, namespace string, collectionName string) error {
	postIndexURL, err := c.client.BuildURL(nil, kvStoreServicePrefix, kvStoreServiceVersion, namespace, "collections", collectionName, "indexes")
	if err != nil {
		return err
	}
	response, err := c.client.Post(postIndexURL, index)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// ListIndexes retrieves all the indexes in a given namespace and collection
func (c *KVStoreService) ListIndexes(namespace string, collectionName string) ([]model.IndexDescription, error) {
	getIndexURL, err := c.client.BuildURL(nil, kvStoreServicePrefix, kvStoreServiceVersion, namespace, "collections", collectionName, "indexes")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(getIndexURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result []model.IndexDescription
	err = util.ParseResponse(&result, response)
	return result, err
}

// DeleteIndex deletes the specified index in a given namespace and collection
func (c *KVStoreService) DeleteIndex(namespace string, collectionName string, indexName string) error {
	deleteIndexURL, err := c.client.BuildURL(nil, kvStoreServicePrefix, kvStoreServiceVersion, namespace, "collections", collectionName, "indexes", indexName)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(deleteIndexURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// InsertRecords posts new records to the collection.
func (c *KVStoreService) InsertRecords(namespace string, collectionName string, records []model.Record) ([]string, error) {
	postRecordURL, err := c.client.BuildURL(nil, kvStoreServicePrefix, kvStoreServiceVersion, namespace, "collections", collectionName, "batch")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(postRecordURL, records)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result []string
	err = util.ParseResponse(&result, response)
	return result, err
}

// QueryRecords queries records present in a given collection.
func (c *KVStoreService) QueryRecords(values url.Values, namespace string, collectionName string) ([]model.Record, error) {
	getRecordURL, err := c.client.BuildURL(values, kvStoreServicePrefix, kvStoreServiceVersion, namespace, "collections", collectionName)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(getRecordURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result []model.Record
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetRecordByKey queries a particular record present in a given collection based on the key value provided by the user.
func (c *KVStoreService) GetRecordByKey(namespace string, collectionName string, keyValue string) (model.Record, error) {
	getRecordURL, err := c.client.BuildURL(nil, kvStoreServicePrefix, kvStoreServiceVersion, namespace, "collections", collectionName, keyValue)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(getRecordURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.Record
	err = util.ParseResponse(&result, response)
	return result, err
}

// DeleteRecords deletes records present in a given collection based on the provided query.
func (c *KVStoreService) DeleteRecords(values url.Values, namespace string, collectionName string) error {
	deleteRecordURL, err := c.client.BuildURL(values, kvStoreServicePrefix, kvStoreServiceVersion, namespace, "collections", collectionName, "query")
	if err != nil {
		return err
	}
	response, err := c.client.Delete(deleteRecordURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// DeleteRecordByKey deletes a particular record present in a given collection based on the key value provided by the user.
func (c *KVStoreService) DeleteRecordByKey(namespace string, collectionName string, keyValue string) error {
	deleteRecordURL, err := c.client.BuildURL(nil, kvStoreServicePrefix, kvStoreServiceVersion, namespace, "collections", collectionName, keyValue)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(deleteRecordURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// ListRecords - List the records created for the tenant's specified collection
func (c *KVStoreService) ListRecords(namespaceName string, collectionName string, filters map[string][]string) ([]map[string]interface{}, error) {
	listRecordsURL, err := c.client.BuildURL(
		filters,
		kvStoreServicePrefix,
		kvStoreServiceVersion,
		namespaceName,
		"collections",
		collectionName)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(listRecordsURL)

	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	var records []map[string]interface{}
	err = util.ParseResponse(&records, response)

	return records, err
}

// InsertRecord - Create a new record in the tenant's specified collection
func (c *KVStoreService) InsertRecord(namespaceName string, collectionName string, record map[string]string) (map[string]string, error) {
	insertRecordURL, err := c.client.BuildURL(
		nil,
		kvStoreServicePrefix,
		kvStoreServiceVersion,
		namespaceName,
		"collections",
		collectionName)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Post(insertRecordURL, record)

	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	// Should always be a map with one key called "_key"
	var responseMap map[string]string
	err = util.ParseResponse(&responseMap, response)

	return responseMap, err
}
