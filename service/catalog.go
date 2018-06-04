package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

// catalog service url prefix
const catalogServicePrefix = "catalog"
const catalogServiceVersion = "v1"

// CatalogService talks to the SSC catalog service
type CatalogService service

// GetDatasets returns all Datasets
func (c *CatalogService) GetDatasets() ([]model.DatasetInfo, error) {
	url, err := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets")
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
	var result []model.DatasetInfo
	err = util.ParseResponse(&result, response)
	return result, err
}

// GetDataset returns the Dataset by name
func (c *CatalogService) GetDataset(id string) (*model.DatasetInfo, error) {
	url, err := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", id)
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
	var result model.DatasetInfo
	err = util.ParseResponse(&result, response)
	return &result, err
}

// CreateDataset creates a new Dataset
func (c *CatalogService) CreateDataset(dataset model.DatasetInfo) (*model.DatasetInfo, error) {
	url, err := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(url, dataset)
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

// UpdateDataset updates an existing Dataset
func (c *CatalogService) UpdateDataset(dataset model.PartialDatasetInfo, datasetID string) (*model.PartialDatasetInfo, error) {
	url, err := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", datasetID)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Patch(url, dataset)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	var result model.PartialDatasetInfo
	err = util.ParseResponse(&result, response)
	return &result, err
}

// DeleteDataset implements delete Dataset endpoint
func (c *CatalogService) DeleteDataset(datasetID string) error {
	url, err := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", datasetID)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(url)
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// DeleteRule deletes the rule by the given path.
func (c *CatalogService) DeleteRule(ruleID string) error {
	getDeleteURL, err := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "rules", ruleID)
	if err != nil {
		return err
	}
	response, err := c.client.Delete(getDeleteURL)
	if response != nil {
		defer response.Body.Close()
	}
	return err
}

// GetRules returns all the rules.
func (c *CatalogService) GetRules() ([]model.Rule, error) {
	getRuleURL, err := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "rules")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(getRuleURL)
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

// GetRule returns rule by an ID.
func (c *CatalogService) GetRule(ruleID string) (*model.Rule, error) {
	getRuleURL, err := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "rules", ruleID)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Get(getRuleURL)
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
	postRuleURL, err := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "rules")
	if err != nil {
		return nil, err
	}
	response, err := c.client.Post(postRuleURL, rule)
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
