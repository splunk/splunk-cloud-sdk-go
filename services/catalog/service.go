// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package catalog

import (
	"net/url"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

// catalog service url prefix
const servicePrefix = "catalog"
const serviceVersion = "v1beta1"
const serviceCluster = "api"

// Service talks to the Splunk Cloud catalog service
type Service services.BaseService

// NewService creates a new catalog service client from the given Config
func NewService(config *services.Config) (*Service, error) {
	baseClient, err := services.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Service{Client: baseClient}, nil
}

// ListDatasets returns all Datasets with optional filter, count, or orderby params
func (s *Service) ListDatasets(values url.Values) ([]interface{}, error) {
	url, err := s.Client.BuildURL(values, serviceCluster, servicePrefix, serviceVersion, "datasets")
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
	var result []interface{}
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetDatasets returns all Datasets
// Deprecated: v0.6.1 - Use ListDatasets instead
func (s *Service) GetDatasets() ([]interface{}, error) {
	return s.ListDatasets(nil)
}

// GetDataset returns the Dataset by resourceName or ID
func (s *Service) GetDataset(resourceNameOrID string) (Dataset, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "datasets", resourceNameOrID)
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

	result, err := parseDatasetResponse(response)
	return result, err
}

// CreateDataset creates a new Dataset
func (s *Service) CreateDataset(dataset interface{}) (Dataset, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "datasets")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: dataset})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	result, err := parseDatasetResponse(response)
	return result, err
}

// CreateIndexDataset creates an index Dataset
func (s *Service) CreateIndexDataset(indexDataset *CreateIndexDataset) (*IndexDataset, error) {
	ds, err := s.CreateDataset(indexDataset)
	if err != nil {
		return nil, err
	}
	ids, ok := ds.(IndexDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: CreateDataset response did not match expected kind: %s", Index)
	}
	return &ids, nil
}

// CreateLookupDataset creates a lookup Dataset
func (s *Service) CreateLookupDataset(lookupDataset *CreateLookupDataset) (*LookupDataset, error) {
	ds, err := s.CreateDataset(lookupDataset)
	if err != nil {
		return nil, err
	}
	lds, ok := ds.(LookupDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: CreateDataset response did not match expected kind: %s", Lookup)
	}
	return &lds, nil
}

// CreateViewDataset creates a view Dataset
func (s *Service) CreateViewDataset(viewDataset *CreateViewDataset) (*ViewDataset, error) {
	ds, err := s.CreateDataset(viewDataset)
	if err != nil {
		return nil, err
	}
	vds, ok := ds.(ViewDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: CreateDataset response did not match expected kind: %s", View)
	}
	return &vds, nil
}

// CreateKVCollectionDataset creates a KVCollection Dataset
func (s *Service) CreateKVCollectionDataset(kvDataset *CreateKVCollectionDataset) (*KVCollectionDataset, error) {
	ds, err := s.CreateDataset(kvDataset)
	if err != nil {
		return nil, err
	}
	kvds, ok := ds.(KVCollectionDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: CreateDataset response did not match expected kind: %s", KvCollection)
	}
	return &kvds, nil
}

// CreateImportDataset creates an import Dataset
func (s *Service) CreateImportDataset(importDataset *CreateImportDataset) (*ImportDataset, error) {
	ds, err := s.CreateDataset(importDataset)
	if err != nil {
		return nil, err
	}
	ids, ok := ds.(ImportDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: CreateDataset response did not match expected kind: %s", Import)
	}
	return &ids, nil
}

// CreateMetricDataset creates a metric Dataset
func (s *Service) CreateMetricDataset(metricDataset *CreateMetricDataset) (*MetricDataset, error) {
	ds, err := s.CreateDataset(metricDataset)
	if err != nil {
		return nil, err
	}
	mds, ok := ds.(MetricDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: CreateDataset response did not match expected kind: %s", Metric)
	}
	return &mds, nil
}

// UpdateDataset updates an existing Dataset with the specified resourceName or ID
func (s *Service) UpdateDataset(dataset interface{}, resourceNameOrID string) (Dataset, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "datasets", resourceNameOrID)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: url, Body: dataset})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	result, err := parseDatasetResponse(response)
	return result, err
}

// UpdateIndexDataset updates an existing index Dataset with the specified resourceName or ID
func (s *Service) UpdateIndexDataset(indexDataset *UpdateIndexDataset, id string) (*IndexDataset, error) {
	ds, err := s.UpdateDataset(indexDataset, id)
	if err != nil {
		return nil, err
	}
	ids, ok := ds.(IndexDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: UpdateDataset response did not match expected kind: %s", Index)
	}
	return &ids, nil
}

// UpdateLookupDataset updates an existing lookup Dataset with the specified resourceName or ID
func (s *Service) UpdateLookupDataset(lookupDataset *UpdateLookupDataset, id string) (*LookupDataset, error) {
	ds, err := s.UpdateDataset(lookupDataset, id)
	if err != nil {
		return nil, err
	}
	lds, ok := ds.(LookupDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: UpdateDataset response did not match expected kind: %s", Lookup)
	}
	return &lds, nil
}

// UpdateViewDataset updates an existing view Dataset with the specified resourceName or ID
func (s *Service) UpdateViewDataset(viewDataset *UpdateViewDataset, id string) (*ViewDataset, error) {
	ds, err := s.UpdateDataset(viewDataset, id)
	if err != nil {
		return nil, err
	}
	vds, ok := ds.(ViewDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: UpdateDataset response did not match expected kind: %s", View)
	}
	return &vds, nil
}

// UpdateKVCollectionDataset updates an existing KVCollection Dataset with the specified resourceName or ID
func (s *Service) UpdateKVCollectionDataset(kvDataset *UpdateKVCollectionDataset, id string) (*KVCollectionDataset, error) {
	ds, err := s.UpdateDataset(kvDataset, id)
	if err != nil {
		return nil, err
	}
	kvds, ok := ds.(KVCollectionDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: UpdateDataset response did not match expected kind: %s", KvCollection)
	}
	return &kvds, nil
}

// UpdateImportDataset updates an existing import Dataset with the specified resourceName or ID
func (s *Service) UpdateImportDataset(importDataset *UpdateImportDataset, id string) (*ImportDataset, error) {
	ds, err := s.UpdateDataset(importDataset, id)
	if err != nil {
		return nil, err
	}
	ids, ok := ds.(ImportDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: UpdateDataset response did not match expected kind: %s", Import)
	}
	return &ids, nil
}

// UpdateMetricDataset updates an existing metric Dataset with the specified resourceName or ID
func (s *Service) UpdateMetricDataset(metricDataset *UpdateMetricDataset, id string) (*MetricDataset, error) {
	ds, err := s.UpdateDataset(metricDataset, id)
	if err != nil {
		return nil, err
	}
	mds, ok := ds.(MetricDataset)
	if !ok {
		return nil, fmt.Errorf("catalog: UpdateDataset response did not match expected kind: %s", Metric)
	}
	return &mds, nil
}

// DeleteDataset implements delete Dataset endpoint with the specified resourceName or ID
func (s *Service) DeleteDataset(resourceNameOrID string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "datasets", resourceNameOrID)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// DeleteRule deletes the rule and its dependencies with the specified rule id or resourceName
func (s *Service) DeleteRule(resourceNameOrID string) error {
	getDeleteURL, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "rules", resourceNameOrID)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: getDeleteURL})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetRules returns all the rules.
func (s *Service) GetRules() ([]Rule, error) {
	getRuleURL, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "rules")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: getRuleURL})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result []Rule
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetRule returns rule by the specified resourceName or ID.
func (s *Service) GetRule(resourceNameOrID string) (*Rule, error) {
	getRuleURL, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "rules", resourceNameOrID)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: getRuleURL})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Rule
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateRule posts a new rule.
func (s *Service) CreateRule(rule Rule) (*Rule, error) {
	// TODO: make a new RuleCreationPayload that omits these:
	rule.Created = ""
	rule.CreatedBy = ""
	rule.Modified = ""
	rule.ModifiedBy = ""
	rule.Owner = ""
	postRuleURL, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "rules")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: postRuleURL, Body: rule})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Rule
	err = util.ParseResponse(&result, response)
	return &result, err
}

// UpdateRule updates the rule with the specified resourceName or ID
func (s *Service) UpdateRule(resourceNameOrID string, rule *RuleUpdateFields) (*Rule, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "rules", resourceNameOrID)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: url, Body: rule})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Rule
	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetDatasetFields returns all the fields belonging to the specified dataset
func (s *Service) GetDatasetFields(datasetID string, values url.Values) ([]Field, error) {
	url, err := s.Client.BuildURL(values, serviceCluster, servicePrefix, serviceVersion, "datasets", datasetID, "fields")
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
	var result []Field
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetDatasetField returns the field belonging to the specified dataset with the id datasetFieldID
func (s *Service) GetDatasetField(datasetID string, datasetFieldID string) (*Field, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "datasets", datasetID, "fields", datasetFieldID)
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
	var result Field
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateDatasetField creates a new field in the specified dataset
func (s *Service) CreateDatasetField(datasetID string, datasetField *Field) (*Field, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "datasets", datasetID, "fields")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: datasetField})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Field
	err = util.ParseResponse(&result, response)
	return &result, err
}

// UpdateDatasetField updates an already existing field in the specified dataset
func (s *Service) UpdateDatasetField(datasetID string, datasetFieldID string, datasetField *Field) (*Field, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "datasets", datasetID, "fields", datasetFieldID)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: url, Body: datasetField})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Field
	err = util.ParseResponse(&result, response)
	return &result, err
}

// DeleteDatasetField deletes the field belonging to the specified dataset with the id datasetFieldID
func (s *Service) DeleteDatasetField(datasetID string, datasetFieldID string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "datasets", datasetID, "fields", datasetFieldID)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetFields returns a list of all Fields on Catalog
func (s *Service) GetFields() ([]Field, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "fields")
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
	var result []Field
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetField returns the Field corresponding to fieldid
func (s *Service) GetField(fieldID string) (*Field, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "fields", fieldID)
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
	var result Field
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateRuleAction creates a new Action on the rule specified
func (s *Service) CreateRuleAction(ruleID string, action *Action) (*Action, error) {
	// TODO: create a new ActionCreationPayload that omits these:
	action.Created = ""
	action.CreatedBy = ""
	action.Modified = ""
	action.ModifiedBy = ""
	action.Owner = ""
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "rules", ruleID, "actions")
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: url, Body: action})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result Action
	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetRuleActions returns a list of all actions belonging to the specified rule
func (s *Service) GetRuleActions(ruleID string) ([]Action, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "rules", ruleID, "actions")
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
	var result []Action
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetRuleAction returns the action of specified belonging to the specified rule
func (s *Service) GetRuleAction(ruleID string, actionID string) (*Action, error) {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "rules", ruleID, "actions", actionID)
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
	var result Action

	err = util.ParseResponse(&result, response)
	return &result, err
}

// DeleteRuleAction deletes the action of specified belonging to the specified rule
func (s *Service) DeleteRuleAction(ruleID string, actionID string) error {
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "rules", ruleID, "actions", actionID)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// UpdateRuleAction updates the action with the specified id for the specified Rule
func (s *Service) UpdateRuleAction(ruleID string, actionID string, action *Action) (*Action, error) {
	// TODO: create a new ActionUpdateFields that omits these:
	action.Created = ""
	action.CreatedBy = ""
	action.Kind = ""
	action.Modified = ""
	action.ModifiedBy = ""
	url, err := s.Client.BuildURL(nil, serviceCluster, servicePrefix, serviceVersion, "rules", ruleID, "actions", actionID)
	if err != nil {
		return nil, err
	}

	response, err := s.Client.Patch(services.RequestParams{URL: url, Body: action})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result Action
	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetModules returns a list of a list of modules that match a filter query if it is given, otherwise return all modules
func (s *Service) GetModules(filter url.Values) ([]Module, error) {
	url, err := s.Client.BuildURL(filter, serviceCluster, servicePrefix, serviceVersion, "modules")
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
	var result []Module
	err = util.ParseResponse(&result, response)
	return result, err
}

// parseDatasetResponse parses incoming http response into specific Dataset subytpe based on 'kind'
func parseDatasetResponse(response *http.Response) (Dataset, error) {
	var datasetInterface interface{}
	err := util.ParseResponse(&datasetInterface, response)
	if err != nil {
		return nil, err
	}

	datasetByte, err := json.Marshal(datasetInterface)
	if err != nil {
		return nil, err
	}

	datasetMap, ok := datasetInterface.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("catalog: response was not of type DatasetBase")
	}
	kind, ok := datasetMap["kind"].(string)
	if !ok {
		return nil, fmt.Errorf("catalog: dataset response did not contain key 'kind' with string value in contents")
	}

	switch kind {
	case string(Index):
		var datasetResult IndexDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	case string(View):
		var datasetResult ViewDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	case string(Lookup):
		var datasetResult LookupDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	case string(Import):
		var datasetResult ImportDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	case string(Metric):
		var datasetResult MetricDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	case string(KvCollection):
		var datasetResult KVCollectionDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	default:
		return nil, fmt.Errorf("catalog: unknown dataset kind: %s", kind)
	}
}
