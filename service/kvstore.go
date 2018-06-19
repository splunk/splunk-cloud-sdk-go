package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

const kvStoreServicePrefix = "kvstore"
const kvStoreServiceVersion = "v1"

// KVStoreService talks to kvstore service
type KVStoreService service

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
	return err
}

// GetIndexes retrieves all the indexes in a given namespace and collection
func (c *KVStoreService) GetIndexes(namespace string, collectionName string) ([]model.IndexDescription, error) {
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
	return err
}

// ListRecords - List the records created for the tenant's specified collection
func (c *KVStoreService) ListRecords(namespaceName string, collectionName string) ([]map[string]interface{}, error) {
	listRecordsURL, err := c.client.BuildURL(
		nil,
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
