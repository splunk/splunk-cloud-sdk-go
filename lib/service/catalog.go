/*
Implement catalog service endpoints
*/

package service

import (
	"github.com/splunk/ssc-client-go/lib/model"
	"github.com/splunk/ssc-client-go/lib/util"
)

// catalog service url prefix
const catalogServicePrefix string = "catalog"
const catalogServiceVersion string = "v1"

// CatalogService represents catalog service
type CatalogService service

// CreateDataset creates a dataset to post
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

// GetDataset implements get Dataset endpoing
func (c *CatalogService) GetDataset(name string) (*model.Dataset, error) {
	var url= c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", name)
	response, err := c.client.Get(url)

	var result model.Dataset
	util.ParseResponse(&result, response, err)

	return &result, err
}

// PostDataset implements post Dataset endpoing
func (c *CatalogService) PostDataset(dataset model.Dataset) (*model.Dataset, error) {
	var url= c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", "")
	response, err := c.client.Post(url, dataset)
	var result model.Dataset
	util.ParseResponse(&result, response, err)

	return &result, err
}

// DeleteDataset implements delete Dataset endpoing
func (c *CatalogService) DeleteDataset(datasetName string) error {
	var url= c.client.BuildURL(catalogServicePrefix, catalogServiceVersion, "datasets", datasetName)
	_, err := c.client.Delete(url)

	return err
}
