// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"net/url"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

// catalog service url prefix
const catalogServicePrefix = "catalog"
const catalogServiceVersion = "v1"

// CatalogService talks to the Splunk Cloud catalog service
type CatalogService service

// GetDatasets returns all Datasets
func (c *CatalogService) GetDatasets() ([]model.DatasetInfo, error) {
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
	var result []model.DatasetInfo
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetDataset returns the Dataset by name or ID
func (c *CatalogService) GetDataset(resourceNameOrID string) (*model.DatasetInfo, error) {
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
	var result model.DatasetInfo
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateDataset creates a new Dataset
func (c *CatalogService) CreateDataset(dataset model.DatasetCreationPayload) (*model.DatasetInfo, error) {
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
	var result model.DatasetInfo
	err = util.ParseResponse(&result, response)
	return &result, err
}

// UpdateDataset updates an existing Dataset with the specified name or ID
func (c *CatalogService) UpdateDataset(dataset model.PartialDatasetInfo, resourceNameOrID string) (*model.DatasetInfo, error) {
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
	var result model.DatasetInfo
	err = util.ParseResponse(&result, response)
	return &result, err
}

// DeleteDataset implements delete Dataset endpoint with the specified name or ID
func (c *CatalogService) DeleteDataset(resourceNameOrID string) error {
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

// DeleteRule deletes the rule and its dependencies with the specified rule id or name
func (c *CatalogService) DeleteRule(resourceNameOrID string) error {
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
func (c *CatalogService) GetRules() ([]model.Rule, error) {
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
	var result []model.Rule
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetRule returns rule by an with the specified rule name or ID.
func (c *CatalogService) GetRule( resourceNameOrID string) (*model.Rule, error) {
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
	var result model.Rule
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateRule posts a new rule.
func (c *CatalogService) CreateRule(rule model.Rule) (*model.Rule, error) {
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
	var result model.Rule
	err = util.ParseResponse(&result, response)
	return &result, err
}

// UpdateRule updates the rule with the specified name or ID
func (c *CatalogService) UpdateRule(dataset model.PartialRule, resourceNameOrID string) (*model.Rule, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", resourceNameOrID)
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
	var result model.Rule
	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetDatasetFields returns all the fields belonging to the specified dataset
func (c *CatalogService) GetDatasetFields(datasetID string, values url.Values) ([]model.Field, error) {
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
	var result []model.Field
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetDatasetField returns the field belonging to the specified dataset with the id datasetFieldID
func (c *CatalogService) GetDatasetField(datasetID string, datasetFieldID string) (*model.Field, error) {
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
	var result model.Field
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateDatasetField creates a new field in the specified dataset
func (c *CatalogService) CreateDatasetField(datasetID string, datasetField model.Field) (*model.Field, error) {
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
	var result model.Field
	err = util.ParseResponse(&result, response)
	return &result, err
}

// UpdateDatasetField updates an already existing field in the specified dataset
func (c *CatalogService) UpdateDatasetField(datasetID string, datasetFieldID string, datasetField model.Field) (*model.Field, error) {
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
	var result model.Field
	err = util.ParseResponse(&result, response)
	return &result, err
}

// DeleteDatasetField deletes the field belonging to the specified dataset with the id datasetFieldID
func (c *CatalogService) DeleteDatasetField(datasetID string, datasetFieldID string) error {
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
func (c *CatalogService) GetFields() ([]model.Field, error) {
	url, err := c.client.BuildURL(nil,catalogServicePrefix, catalogServiceVersion, "fields")
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
	var result []model.Field
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetField returns he Field corresponding to fieldid
func (c *CatalogService) GetField(fieldID string) (*model.Field, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion,  "fields", fieldID)
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
	var result model.Field
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateRuleAction creates a new Action on the rule specified
func (c *CatalogService) CreateRuleAction(resourceNameOrRuleID string, action *model.CatalogAction) (*model.CatalogAction, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", resourceNameOrRuleID, "actions")
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

	var result model.CatalogAction
	err = util.ParseResponse(&result, response)
	return &result, err
}

// GetRuleActions returns a list of all actions belonging to the specified rule
func (c *CatalogService) GetRuleActions(ruleID string) ([]model.CatalogAction, error) {
	url, err := c.client.BuildURL(nil,catalogServicePrefix, catalogServiceVersion, "rules",ruleID,"actions")
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
	var result []model.CatalogAction
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetRuleAction returns the action of specified belonging to the specified rule
func (c *CatalogService) GetRuleAction(ruleID string, actionID string) (*model.CatalogAction, error) {
	url, err := c.client.BuildURL(nil,catalogServicePrefix, catalogServiceVersion, "rules",ruleID,"actions",actionID)
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
	var result model.CatalogAction
	err = util.ParseResponse(&result, response)
	return &result, err
}

// DeleteRuleAction deletes the action of specified belonging to the specified rule
func (c *CatalogService) DeleteRuleAction(ruleID string, actionID string) error {
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
func (c *CatalogService) UpdateRuleAction(ruleID string, actionID string, action *model.CatalogAction) (*model.Field, error) {
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
	var result model.Field
	err = util.ParseResponse(&result, response)
	return &result, err
}
