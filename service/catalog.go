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
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets")
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
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets", id)
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
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets")
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
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets", datasetID)
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
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets", datasetID)
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
	getDeleteURL, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", ruleID)
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
	getRuleURL, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules")
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
	getRuleURL, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules", ruleID)
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
	postRuleURL, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "rules")
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

// GetDatasetFields returns all the fields belonging to the specified dataset
func (c *CatalogService) GetDatasetFields(datasetID string) ([]model.Field, error) {
	url, err := c.client.BuildURL(nil, catalogServicePrefix, catalogServiceVersion, "datasets", datasetID, "fields")
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
	response, err := c.client.Get(url)
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
	response, err := c.client.Delete(url)
	if response != nil {
		defer response.Body.Close()
	}
	return err
}
