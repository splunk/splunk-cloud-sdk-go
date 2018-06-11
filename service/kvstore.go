package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)


const kvStoreServicePrefix = "kvstore"
const kvStoreServiceVersion = "v1"
const kvStoreCollectionsResource = "collections"

// KVStoreService talks to kvstore service
type KVStoreService service


// GetCollectionStats returns Collection Stats for the collection
func (c *KVStoreService) GetCollectionStats(namespace string, collection string) ([]model.CollectionStats, error) {
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
	var result []model.CollectionStats
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetPingStatus returns Service Health Status
func (c *KVStoreService) GetServiceHealthStatus() ([]model.ServiceHealthStatus, error) {
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
	var result []model.ServiceHealthStatus
	err = util.ParseResponse(&result, response)
	return result, err
}