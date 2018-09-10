// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"io/ioutil"
	"net/url"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const kvStoreServicePrefix = "kvstore"
const kvStoreServiceVersion = "v1"
const kvStoreCollectionsResource = "collections"

// KVStoreService talks to kvstore service
type KVStoreService service

// GetCollectionStats returns Collection Stats for the collection
func (c *KVStoreService) GetCollectionStats(collection string) (*model.CollectionStats, error) {
	url, err := c.client.BuildURLDefaultTenant(nil, kvStoreServicePrefix, kvStoreServiceVersion, kvStoreCollectionsResource, collection, "stats")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(RequestParams{URL: url})
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
	url, err := c.client.BuildURLDefaultTenant(nil, kvStoreServicePrefix, kvStoreServiceVersion, "ping")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(RequestParams{URL: url})
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

// GetCollections gets all the collections
func (c *KVStoreService) GetCollections() ([]model.CollectionDefinition, error) {
	url, err := c.client.BuildURLDefaultTenant(nil, kvStoreServicePrefix, kvStoreServiceVersion, kvStoreCollectionsResource)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result []model.CollectionDefinition
	err = util.ParseResponse(&result, response)
	return result, err
}

// ExportCollection exports the specified collection records to an external file
func (c *KVStoreService) ExportCollection(collectionName string, contentType model.ExportCollectionContentType) (string, error) {
	url, err := c.client.BuildURLDefaultTenant(nil, kvStoreServicePrefix, kvStoreServiceVersion, kvStoreCollectionsResource, collectionName, "export")
	if err != nil {
		return "", err
	}

	var acceptType string
	if contentType == model.CSV {
		acceptType = "text/csv"
	} else {
		acceptType = "application/gzip"
	}

	headers := map[string]string{
		"Accept": acceptType,
	}
	response, err := c.client.Get(RequestParams{URL: url, Headers: headers})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return "", err
	}

	records, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(records), nil
}

// CreateIndex posts a new index to be added to the collection.
func (c *KVStoreService) CreateIndex(collectionName string, index model.IndexDefinition) (*model.IndexDescription, error) {
	url, err := c.client.BuildURLDefaultTenant(nil, kvStoreServicePrefix, kvStoreServiceVersion, "collections", collectionName, "indexes")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: index})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.IndexDescription
	err = util.ParseResponse(&result, response)
	return &result, err
}

// ListIndexes retrieves all the indexes in a given collection
func (c *KVStoreService) ListIndexes(collectionName string) ([]model.IndexDefinition, error) {
	url, err := c.client.BuildURLDefaultTenant(nil, kvStoreServicePrefix, kvStoreServiceVersion, "collections", collectionName, "indexes")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result []model.IndexDefinition
	err = util.ParseResponse(&result, response)
	return result, err
}

// DeleteIndex deletes the specified index in a given collection
func (c *KVStoreService) DeleteIndex(collectionName string, indexName string) error {
	url, err := c.client.BuildURLDefaultTenant(nil, kvStoreServicePrefix, kvStoreServiceVersion, "collections", collectionName, "indexes", indexName)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// InsertRecords posts new records to the collection.
func (c *KVStoreService) InsertRecords(collectionName string, records []model.Record) ([]string, error) {
	url, err := c.client.BuildURLDefaultTenant(nil, kvStoreServicePrefix, kvStoreServiceVersion, "collections", collectionName, "batch")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: records})
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
func (c *KVStoreService) QueryRecords(collectionName string, values url.Values) ([]model.Record, error) {
	url, err := c.client.BuildURLDefaultTenant(values,
		kvStoreServicePrefix,
		kvStoreServiceVersion,
		"collections",
		collectionName,
		"query")

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *KVStoreService) GetRecordByKey(collectionName string, keyValue string) (model.Record, error) {
	url, err := c.client.BuildURLDefaultTenant(nil, kvStoreServicePrefix, kvStoreServiceVersion, "collections", collectionName, keyValue)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(RequestParams{URL: url})
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
func (c *KVStoreService) DeleteRecords(collectionName string, values url.Values) error {
	url, err := c.client.BuildURLDefaultTenant(values, kvStoreServicePrefix, kvStoreServiceVersion, "collections", collectionName, "query")
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// DeleteRecordByKey deletes a particular record present in a given collection based on the key value provided by the user.
func (c *KVStoreService) DeleteRecordByKey(collectionName string, keyValue string) error {
	url, err := c.client.BuildURLDefaultTenant(nil, kvStoreServicePrefix, kvStoreServiceVersion, "collections", collectionName, keyValue)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// ListRecords - List the records created for the tenant's specified collection TODO: include count, offset and orderBy
func (c *KVStoreService) ListRecords(collectionName string, filters map[string][]string) ([]map[string]interface{}, error) {
	url, err := c.client.BuildURLDefaultTenant(
		filters,
		kvStoreServicePrefix,
		kvStoreServiceVersion,
		"collections",
		collectionName)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Get(RequestParams{URL: url})

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
func (c *KVStoreService) InsertRecord(collectionName string, record map[string]string) (map[string]string, error) {
	url, err := c.client.BuildURLDefaultTenant(
		nil,
		kvStoreServicePrefix,
		kvStoreServiceVersion,
		"collections",
		collectionName)

	if err != nil {
		return nil, err
	}

	response, err := c.client.Post(RequestParams{URL: url, Body: record})

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
