package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

// catalog service url prefix
const catalogServicePrefix = "catalog"
const catalogServiceVersion = "v1"

// CatalogService represents catalog service
type CatalogService service

// GetDatasets returns all Datasets
func (c *CatalogService) GetDatasets() ([]model.Dataset, error) {
	var url = c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets")
	response, err := c.client.Get(url)

	var result []model.Dataset
	util.ParseResponse(&result, response, err)

	return result, err
}

// GetDataset returns the Dataset by name
func (c *CatalogService) GetDataset(name string) (*model.Dataset, error) {
	var url = c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", name)
	response, err := c.client.Get(url)

	var result model.Dataset
	util.ParseResponse(&result, response, err)

	return &result, err
}

// CreateDataset creates a new Dataset
// TODO: Can we remove the empty string ("") argument when calling 'BuildURL'?
func (c *CatalogService) CreateDataset(dataset model.Dataset) (*model.Dataset, error) {
	var url = c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", "")
	response, err := c.client.Post(url, dataset)

	var result model.Dataset
	util.ParseResponse(&result, response, err)

	return &result, err
}

// DeleteDataset implements delete Dataset endpoint
func (c *CatalogService) DeleteDataset(datasetName string) error {
	var url = c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", datasetName)
	response, err := c.client.Delete(url)

	return util.ParseError(response, err)
}

// DeleteRule deletes the rule by the given path.
func (c *CatalogService) DeleteRule(rulePath string) error {
	getDeleteURL := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "rules", rulePath)
	response, err := c.client.Delete(getDeleteURL)

	return util.ParseError(response, err)
}

// GetRules returns all the rules.
func (c *CatalogService) GetRules() ([]model.Rule, error) {
	getRuleURL := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "rules")
	response, err := c.client.Get(getRuleURL)
	if err != nil {
		return nil, err
	}

	var result []model.Rule
	util.ParseResponse(&result, response, err)

	return result, err
}

// CreateRule posts a new rule.
func (c *CatalogService) CreateRule(rule model.Rule) (*model.Rule, error) {
	postRuleURL := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "rules")
	response, err := c.client.Post(postRuleURL, rule)
	if err != nil {
		return nil, err
	}

	var result model.Rule
	util.ParseResponse(&result, response, err)

	return &result, err
}
