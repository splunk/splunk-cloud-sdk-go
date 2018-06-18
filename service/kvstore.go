package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
	"net/url"
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

// CreateRecords posts new records to the collection.
func (c *KVStoreService) CreateRecords(namespace string, collectionName string, records []model.Record) ([]string, error) {
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

// GetRecords queries records present in a given collection.
func (c *KVStoreService) GetRecords(values url.Values, namespace string, collectionName string) ([]model.Record, error) {
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

// TODO: Delete record based on query - currently this is deleting all the records.
// DeleteRecords deletes all the records present in a given collection.
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
	return err
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
	return err
}