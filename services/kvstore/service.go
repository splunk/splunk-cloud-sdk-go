// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package kvstore

import (
	"net/http"
	"net/url"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const servicePrefix = "kvstore"
const serviceVersion = "v1beta1"
const serviceCluster = "api"

// Service talks to kvstore service
type Service services.BaseService

// NewService creates a new kvstore service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

// GetServiceHealthStatus returns Service Health Status
func (s *Service) GetServiceHealthStatus() (*PingOKBody, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "ping")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result PingOKBody
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateIndex posts a new index to be added to the collection.
func (s *Service) CreateIndex(collectionName string, index IndexDefinition) (*IndexDescription, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "collections", collectionName, "indexes")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: index})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result IndexDescription
	err = util.ParseResponse(&result, response)
	return &result, err
}

// ListIndexes retrieves all the indexes in a given collection
func (s *Service) ListIndexes(collectionName string) ([]IndexDefinition, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "collections", collectionName, "indexes")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result []IndexDefinition
	err = util.ParseResponse(&result, response)
	return result, err
}

// DeleteIndex deletes the specified index in a given collection
func (s *Service) DeleteIndex(collectionName string, indexName string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "collections", collectionName, "indexes", indexName)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// InsertRecords posts new records to the collection.
func (s *Service) InsertRecords(collectionName string, records []Record) ([]string, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "collections", collectionName, "batch")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: records})
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
func (s *Service) QueryRecords(collectionName string, values url.Values) ([]Record, error) {
	url, err := s.Client.BuildURL(values, serviceCluster,
		servicePrefix,
		serviceVersion,
		"collections",
		collectionName,
		"query")

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result []Record
	err = util.ParseResponse(&result, response)

	return result, err
}

// GetRecordByKey queries a particular record present in a given collection based on the key value provided by the user.
func (s *Service) GetRecordByKey(collectionName string, keyValue string) (Record, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "collections", collectionName, "records", keyValue)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Record
	err = util.ParseResponse(&result, response)
	return result, err
}

// DeleteRecords deletes records present in a given collection based on the provided query.
func (s *Service) DeleteRecords(collectionName string, values url.Values) error {
	url, err := s.Client.BuildURL(values, serviceCluster, servicePrefix, serviceVersion, "collections", collectionName, "query")
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// DeleteRecordByKey deletes a particular record present in a given collection based on the key value provided by the user.
func (s *Service) DeleteRecordByKey(collectionName string, keyValue string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "collections", collectionName, "records", keyValue)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return err
	}
	return nil
}

// ListRecords - List the records created for the tenant's specified collection TODO: include count, offset and orderBy
func (s *Service) ListRecords(collectionName string, filters map[string][]string) ([]map[string]interface{}, error) {
	url, err := s.Client.BuildURL(
		filters, serviceCluster,
		servicePrefix,
		serviceVersion,
		"collections",
		collectionName)

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Get(services.RequestParams{URL: url})

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
func (s *Service) InsertRecord(collectionName string, record Record) (map[string]string, error) {
	url, err := s.Client.BuildURL(
		nil, serviceCluster,
		servicePrefix,
		serviceVersion,
		"collections",
		collectionName)

	if err != nil {
		return nil, err
	}

	response, err := s.Client.Post(services.RequestParams{URL: url, Body: record})

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

// PutRecord - Inserts or replaces a record in the tenant's specified collection with the specified key
func (s *Service) PutRecord(collectionName, keyValue string, record Record) (map[string]string, bool, error) {
	url, err := s.Client.BuildURL(
		nil, serviceCluster,
		servicePrefix,
		serviceVersion,
		"collections", collectionName,
		"records", keyValue,
	)
	if err != nil {
		return nil, false, err
	}

	response, err := s.Client.Put(services.RequestParams{URL: url, Body: record})

	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, false, err
	}

	// Should always be a map with one key called "_key"
	var responseMap map[string]string
	err = util.ParseResponse(&responseMap, response)
	if err != nil {
		return nil, false, err
	}

	// returns whether or not the PutRecord was an insert or a replace
	return responseMap, response.StatusCode == http.StatusCreated, nil
}
