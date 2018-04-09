/*
Package service implements a service client which is used to communicate
with Splunkd endpoints as well as a collection of services that group
logically related Splunkd endpoints.
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
	response, err := c.client.Get(url)

	body, err := ioutil.ReadAll(response.Body)

	var result model.Datasets
	err = json.Unmarshal(body, &result)

	return result, err
}

func (c *CatalogService) GetDataset(name string) (model.Dataset, error) {
	var url = c.BuildURL(CATALOG_SERVICE_PREFIX, "datasets"+"/"+name, "")
	response, err := c.client.Get(url)
	body, err := ioutil.ReadAll(response.Body)

	var result model.Dataset
	err = json.Unmarshal(body, &result)

	return result, err
}

func (c *CatalogService) GetRules() (model.Rules) {
	var ds model.Rule = model.Rule{"rule1", []model.ActionKind{model.LOOKUP}, "match", 9, "something"};
	return model.Rules{ds}
}
