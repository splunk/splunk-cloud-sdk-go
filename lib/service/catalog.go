/*
Implement catalog service endpoints
*/

package service

import (
	"encoding/json"
	"github.com/splunk/ssc-client-go/lib/model"
	"io/ioutil"
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

////BuildURL is to create a catalog URL
//func (c *CatalogService) BuildURL(prefix string, path string, query string) url.URL {
//	return c.client.BuildURL(nil, catalogServicePrefix, path)
//}

// GetDatasets implements get Datasets endpoint
func (c *CatalogService) GetDatasets() (model.Datasets, error) {
	var url=c.client.BuildURL(nil, catalogServicePrefix,catalogServiceVersion, "datasets")
	response, err := c.client.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}

	var result model.Datasets
	err = json.Unmarshal(body, &result)

	return result, err
}

// GetDataset implements get Dataset endpoing
func (c *CatalogService) GetDataset(name string) (*model.Dataset, error) {
	var url= c.client.BuildURL(nil, catalogServicePrefix,catalogServiceVersion, "datasets",name)
	response, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}

	var result model.Dataset
	err = json.Unmarshal(body, &result)

	return &result, err
}

// PostDataset implements post Dataset endpoing
func (c *CatalogService) PostDataset(dataset model.Dataset) (*model.Dataset, error) {
	var url = c.client.BuildURL(nil,catalogServicePrefix,catalogServiceVersion, "datasets", "")
	response, err := c.client.Post(url, dataset)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}

	var result model.Dataset
	err = json.Unmarshal(body, &result)

	return &result, err
}

// DeleteDataset implements delete Dataset endpoing
func (c *CatalogService) DeleteDataset(datasetName string) error {
	var url= c.client.BuildURL(nil, catalogServicePrefix,catalogServiceVersion, "datasets", datasetName)
	_, err := c.client.Delete(url)

	return err
}
