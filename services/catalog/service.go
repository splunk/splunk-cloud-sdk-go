// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package catalog

import (
	"github.com/splunk/splunk-cloud-sdk-go/services"
)

// catalog service url prefix
const catalogServicePrefix = "catalog"
const catalogServiceVersion = "v1beta1"

// Service talks to the Splunk Cloud catalog service
type Service services.BaseService

// NewService creates a new catalog service with client
func NewService(client *services.Client) *Service {
	return &Service{Client: client}
}

/*
// GetDatasets returns all Datasets
func (c *Service) GetDatasets() ([]DatasetInfo, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets")
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
	var result []DatasetInfo
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetDataset returns the Dataset by resourceName or ID
func (c *Service) GetDataset(resourceNameOrID string) (*DatasetInfo, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets", resourceNameOrID)
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
	var result DatasetInfo
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateDataset creates a new Dataset
func (c *Service) CreateDataset(dataset *DatasetCreationPayload) (*DatasetInfo, error) {
	// TODO: remove this from DatasetCreationPayload
	dataset.Owner = ""
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: dataset})
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

// UpdateDataset updates an existing Dataset with the specified resourceName or ID
func (c *Service) UpdateDataset(dataset *UpdateDatasetInfoFields, resourceNameOrID string) (*DatasetInfo, error) {
	// TODO: remove these from UpdateDatasetInfoFields
	dataset.Created = ""
	dataset.CreatedBy = ""
	dataset.Kind = ""
	dataset.Modified = ""
	dataset.ModifiedBy = ""
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets", resourceNameOrID)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Patch(RequestParams{URL: url, Body: dataset})
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

// DeleteDataset implements delete Dataset endpoint with the specified resourceName or ID
func (c *Service) DeleteDataset(resourceNameOrID string) error {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets", resourceNameOrID)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// DeleteRule deletes the rule and its dependencies with the specified rule id or resourceName
func (c *Service) DeleteRule(resourceNameOrID string) error {
	getDeleteURL, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", resourceNameOrID)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: getDeleteURL})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetRules returns all the rules.
func (c *Service) GetRules() ([]Rule, error) {
	getRuleURL, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(RequestParams{URL: getRuleURL})
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
func (c *Service) GetRule(resourceNameOrID string) (*Rule, error) {
	getRuleURL, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", resourceNameOrID)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(RequestParams{URL: getRuleURL})
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
func (c *Service) CreateRule(rule Rule) (*Rule, error) {
	// TODO: make a new RuleCreationPayload that omits these:
	rule.Created = ""
	rule.CreatedBy = ""
	rule.Modified = ""
	rule.ModifiedBy = ""
	rule.Owner = ""
	postRuleURL, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: postRuleURL, Body: rule})
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
func (c *Service) UpdateRule(resourceNameOrID string, rule *RuleUpdateFields) (*Rule, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", resourceNameOrID)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Patch(RequestParams{URL: url, Body: rule})
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
func (c *Service) GetDatasetFields(datasetID string, values url.Values) ([]Field, error) {
	url, err := c.client.BuildURL(values, catalogServicePrefix, catalogServiceVersion, "datasets", datasetID, "fields")
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
	var result []Field
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetDatasetField returns the field belonging to the specified dataset with the id datasetFieldID
func (c *Service) GetDatasetField(datasetID string, datasetFieldID string) (*Field, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets", datasetID, "fields", datasetFieldID)
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
	var result Field
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateDatasetField creates a new field in the specified dataset
func (c *Service) CreateDatasetField(datasetID string, datasetField *Field) (*Field, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets", datasetID, "fields")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: datasetField})
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
func (c *Service) UpdateDatasetField(datasetID string, datasetFieldID string, datasetField *Field) (*Field, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets", datasetID, "fields", datasetFieldID)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Patch(RequestParams{URL: url, Body: datasetField})
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
func (c *Service) DeleteDatasetField(datasetID string, datasetFieldID string) error {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets", datasetID, "fields", datasetFieldID)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetFields returns a list of all Fields on Catalog
func (c *Service) GetFields() ([]Field, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "fields")
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
	var result []Field
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetField returns the Field corresponding to fieldid
func (c *Service) GetField(fieldID string) (*Field, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "fields", fieldID)
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
	var result Field
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateRuleAction creates a new Action on the rule specified
func (c *Service) CreateRuleAction(ruleID string, action *CatalogAction) (*CatalogAction, error) {
	// TODO: create a new CatalogActionCreationPayload that omits these:
	action.Created = ""
	action.CreatedBy = ""
	action.Modified = ""
	action.ModifiedBy = ""
	action.Owner = ""
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", ruleID, "actions")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(RequestParams{URL: url, Body: action})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var result CatalogAction
	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetRuleActions returns a list of all actions belonging to the specified rule
func (c *Service) GetRuleActions(ruleID string) ([]CatalogAction, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", ruleID, "actions")
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
	var result []CatalogAction
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetRuleAction returns the action of specified belonging to the specified rule
func (c *Service) GetRuleAction(ruleID string, actionID string) (*CatalogAction, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", ruleID, "actions", actionID)
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
	var result CatalogAction

	err = util.ParseResponse(&result, response)
	return &result, err
}

// DeleteRuleAction deletes the action of specified belonging to the specified rule
func (c *Service) DeleteRuleAction(ruleID string, actionID string) error {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", ruleID, "actions", actionID)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(RequestParams{URL: url})
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// UpdateRuleAction updates the action with the specified id for the specified Rule
func (c *Service) UpdateRuleAction(ruleID string, actionID string, action *CatalogAction) (*CatalogAction, error) {
	// TODO: create a new CatalogActionUpdateFields that omits these:
	action.Created = ""
	action.CreatedBy = ""
	action.Kind = ""
	action.Modified = ""
	action.ModifiedBy = ""
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", ruleID, "actions", actionID)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Patch(RequestParams{URL: url, Body: action})
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result CatalogAction
	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetModules returns a list of a list of modules that match a filter query if it is given, otherwise return all modules
func (c *Service) GetModules(filter url.Values) ([]Module, error) {
	url, err := c.client.BuildURL(filter, catalogServicePrefix, catalogServiceVersion, "modules")
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
	var result []Module
	err = util.ParseResponse(&result, response)
	return result, err
}
*/
