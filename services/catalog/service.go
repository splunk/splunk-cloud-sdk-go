// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package catalog

import (
	"net/url"

	"encoding/json"
	"fmt"
	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
	"net/http"
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
func (s *Service) ListDatasets(values url.Values) ([]DatasetInfo, error) {
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
	var result []DatasetInfo
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetDatasets returns all Datasets
// Deprecated: v0.6.1 - Use ListDatasets instead
func (s *Service) GetDatasets() ([]DatasetInfo, error) {
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
func (s *Service) CreateDataset(dataset Dataset) (Dataset, error) {
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

// UpdateDataset updates an existing Dataset with the specified resourceName or ID
func (s *Service) UpdateDataset(dataset *UpdateDatasetInfoFields, resourceNameOrID string) (*DatasetInfo, error) {
	// TODO: remove these from UpdateDatasetInfoFields
	dataset.Created = ""
	dataset.CreatedBy = ""
	dataset.Kind = ""
	dataset.Modified = ""
	dataset.ModifiedBy = ""
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
	var result DatasetInfo
	err = util.ParseResponse(&result, response)
	return &result, err
}

// UpdateIndexDataset updates an existing index Dataset with the specified resourceName or ID  TODO: No module in Update?
func (s *Service) UpdateIndexDataset(indexDataset *UpdateIndex, id string) (*DatasetInfo, error) {
	return s.UpdateDataset(&UpdateDatasetInfoFields{
		Name:                   indexDataset.UpdateDataset.Name,
		Kind:                   indexDataset.UpdateDataset.Kind,
		Owner:                  indexDataset.UpdateDataset.Owner,
		Capabilities:           indexDataset.UpdateDataset.Capabilities,
		Version:                indexDataset.UpdateDataset.Version,
		FrozenTimePeriodInSecs: indexDataset.FrozenTimePeriodInSecs,
		Disabled:               indexDataset.Disabled}, id)
}

// UpdateLookupDataset updates an existing lookup Dataset with the specified resourceName or ID
func (s *Service) UpdateLookupDataset(lookupDataset *UpdateLookup, id string) (*DatasetInfo, error) {
	return s.UpdateDataset(&UpdateDatasetInfoFields{
		Name:               lookupDataset.UpdateDataset.Name,
		Kind:               lookupDataset.UpdateDataset.Kind,
		Owner:              lookupDataset.UpdateDataset.Owner,
		Capabilities:       lookupDataset.UpdateDataset.Capabilities,
		Version:            lookupDataset.UpdateDataset.Version,
		ExternalKind:       lookupDataset.ExternalKind,
		ExternalName:       lookupDataset.ExternalName,
		CaseSensitiveMatch: lookupDataset.CaseSensitiveMatch,
		Filter:             lookupDataset.Filter}, id)
}

// UpdateViewDataset updates an existing view Dataset with the specified resourceName or ID
func (s *Service) UpdateViewDataset(viewDataset *UpdateView, id string) (*DatasetInfo, error) {
	return s.UpdateDataset(&UpdateDatasetInfoFields{
		Name:         viewDataset.UpdateDataset.Name,
		Kind:         viewDataset.UpdateDataset.Kind,
		Owner:        viewDataset.UpdateDataset.Owner,
		Capabilities: viewDataset.UpdateDataset.Capabilities,
		Version:      viewDataset.UpdateDataset.Version,
		Search:       viewDataset.Search}, id)
}

// UpdateImportDataset updates an existing import Dataset with the specified resourceName or ID
func (s *Service) UpdateImportDataset(importDataset *UpdateImport, id string) (*DatasetInfo, error) {
	return s.UpdateDataset(&UpdateDatasetInfoFields{
		Name:         importDataset.UpdateDataset.Name,
		Kind:         importDataset.UpdateDataset.Kind,
		Owner:        importDataset.UpdateDataset.Owner,
		Capabilities: importDataset.UpdateDataset.Capabilities,
		Version:      importDataset.UpdateDataset.Version,
		SourceName:   importDataset.SourceName,
		SourceModule: importDataset.SourceModule}, id)
}

// UpdateMetricDataset updates an existing metric Dataset with the specified resourceName or ID
func (s *Service) UpdateMetricDataset(metricDataset *UpdateMetric, id string) (*DatasetInfo, error) {
	return s.UpdateDataset(&UpdateDatasetInfoFields{
		Name:         metricDataset.UpdateDataset.Name,
		Kind:         metricDataset.UpdateDataset.Kind,
		Owner:        metricDataset.UpdateDataset.Owner,
		Capabilities: metricDataset.UpdateDataset.Capabilities,
		Version:      metricDataset.UpdateDataset.Version,
		Disabled:     metricDataset.Disabled}, id)
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

// parseDatasetResponse parses incoming http response into specific Dataset subytpe based on 'Kind'
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

	datasetMap, _ := datasetInterface.(map[string]interface{})
	switch datasetMap["kind"] {
	case "index":
		var datasetResult IndexDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	case "view":
		var datasetResult ViewDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	case "lookup":
		var datasetResult LookupDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	case "import":
		var datasetResult ImportDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	case "metric":
		var datasetResult MetricDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	case "kvcollection":
		var datasetResult KVCollectionDataset
		err = json.Unmarshal(datasetByte, &datasetResult)
		return datasetResult, err
	default:
		fmt.Printf("ERROR: Unknown Dataset Type")
	}

	return nil, err
}
