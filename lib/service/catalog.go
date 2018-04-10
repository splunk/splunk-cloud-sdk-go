/*
Implement catalog service endpoints
*/

package service

import (
	"net/url"
	"io/ioutil"
	"encoding/json"
	"github.com/splunk/ssc-client-go/lib/model"
)

const CATALOG_SERVICE_PREFIX string = "/catalog/v1";

type CatalogService service

type dataset_post struct {
	Name  string            `json:"name"`
	Kind  model.DatasetKind `json:"kind"`
	Rules []string          `json:"rules"`
	Todo  string            `json:"todo"`
}

// creates a dataset to post
func (c *CatalogService) CreateDataset(name string, kind model.DatasetKind, rules []string, todo string) dataset_post {
	return dataset_post{
		Name:  name,
		Kind:  kind,
		Rules: rules,
		Todo:  todo,
	}
}

// creates a catalog URL //todo: move to client.go or other common files
func (c *CatalogService) BuildURL(prefix string, path string, query string) url.URL {
	return url.URL{
		Scheme:   c.client.Scheme,
		Path:     CATALOG_SERVICE_PREFIX + "/" + path,
		RawQuery: query,
		Host:     c.client.Host,
	}
}

func (c *CatalogService) GetDatasets() (model.Datasets, error) {
	var url = c.BuildURL(CATALOG_SERVICE_PREFIX, "datasets", "")
	response, err := c.client.Get(url, JSON)

	body, err := ioutil.ReadAll(response.Body)

	var result model.Datasets
	err = json.Unmarshal(body, &result)

	return result, err
}

func (c *CatalogService) GetDataset(name string) (model.Dataset, error) {
	var url = c.BuildURL(CATALOG_SERVICE_PREFIX, "datasets"+"/"+name, "")
	response, err := c.client.Get(url, JSON)
	body, err := ioutil.ReadAll(response.Body)

	var result model.Dataset
	err = json.Unmarshal(body, &result)

	return result, err
}

func (c *CatalogService) PostDataset(dataset dataset_post) (model.Dataset, error) {
	var url = c.BuildURL(CATALOG_SERVICE_PREFIX, "datasets", "")
	response, err := c.client.Post(url, dataset, JSON)

	body, err := ioutil.ReadAll(response.Body)

	var result model.Dataset
	err = json.Unmarshal(body, &result)

	return result, err
}


func (c *CatalogService) DeleteDataset(dataset_name string) (error) {
	var url = c.BuildURL(CATALOG_SERVICE_PREFIX, "datasets"+"/"+dataset_name, "")
	_, err := c.client.Delete(url, JSON)

	return err
}
