package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
)

// catalog service url prefix
const catalogServicePrefix string = "catalog"
const catalogServiceVersion string = "v1"

// CatalogService represents catalog service
type CatalogService service

// CreateDataset creates a dataset to post
// TODO: ID may be passed in later on
func (c *CatalogService) CreateDataset(name string, kind model.DatasetKind, rules []string, todo string) model.Dataset {
	return model.Dataset{
		ID:    "",
		Name:  name,
		Kind:  kind,
		Rules: rules,
		Todo:  todo,
	}
}

// GetDatasets implements get Datasets endpoint
func (c *CatalogService) GetDatasets() ([]model.Dataset, error) {
	var url= c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets")
	response, err := c.client.Get(url)

	var result []model.Dataset
	util.ParseResponse(&result, response, err)

	return result, err
}

// GetDataset implements get Dataset endpoint
func (c *CatalogService) GetDataset(name string) (*model.Dataset, error) {
	var url= c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", name)
	response, err := c.client.Get(url)

	var result model.Dataset
	util.ParseResponse(&result, response, err)

	return &result, err
}

// PostDataset implements post Dataset endpoint
// TODO: Can we remove the empty string ("") argument when calling 'BuildURL'?
func (c *CatalogService) PostDataset(dataset model.Dataset) (*model.Dataset, error) {
	var url= c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", "")
	response, err := c.client.Post(url, dataset)
	var result model.Dataset
	util.ParseResponse(&result, response, err)

	return &result, err
}

// DeleteDataset implements delete Dataset endpoint
func (c *CatalogService) DeleteDataset(datasetName string) error {
	var url= c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", datasetName)
	_, err := c.client.Delete(url)

	return err
}

// DeleteRule deletes the rule by the given path.
func (c *CatalogService) DeleteRule(rulePath string) (error) {
	getDeleteURL := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "rules", rulePath)
	_, err := c.client.Delete(getDeleteURL)

	return err
}

// GetRules returns all the rules.
func (c *CatalogService) GetRules() ([]model.Rule, error){
	getRuleURL := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "rules")
	response, err := c.client.Get(getRuleURL)
	if err != nil {
		return nil, err
	}

	var result []model.Rule
	util.ParseResponse(&result, response, err)

	return result, err
}


// PostRule posts a new rule.
func (c *CatalogService) PostRule(rule model.Rule) (*model.Rule, error) {
	postRuleURL := c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "rules")
	response, err := c.client.Post(postRuleURL, rule)
	if err != nil {
		return nil, err
	}

	var result model.Rule
	util.ParseResponse(&result, response, err)

	return &result, err
}
