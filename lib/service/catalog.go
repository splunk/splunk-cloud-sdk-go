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
	"path"
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

func (c *CatalogService) PostDataset(dataset dataset_post) (model.Dataset, error) {
	var url = c.BuildURL(CATALOG_SERVICE_PREFIX, "datasets", "")
	response, err := c.client.Post(url, dataset)

	body, err := ioutil.ReadAll(response.Body)

	var result model.Dataset
	err = json.Unmarshal(body, &result)

	return result, err

}

/**
 * Delete the rule by the given path.
 * @param {string} rulePath
 */
func (c *CatalogService) deleteRule(rulePath string) (model.Rule, error) {
	buildPath := ""
	buildPath = path.Join(buildPath, "rules", rulePath)
	getDeleteUrl := c.BuildURL(CATALOG_SERVICE_PREFIX, buildPath, "")
	response, err := c.client.Delete(getDeleteUrl)

	body, err := ioutil.ReadAll(response.Body)

	var result model.Rule
	err = json.Unmarshal(body, &result)

	return result, err
}

/**
 * Returns the rule identified by the given path.
 * The path must be fully qualified, if the path is a prefix the request returns 404
 * because it does identify a rule resource.
 * @param {string} rulePath
 * @return {Promise<CatalogProxy~Rule>}
 */
func (c *CatalogService) GetRules() ([]model.Rule, error){
	getRuleUrl := c.BuildURL(CATALOG_SERVICE_PREFIX, "rules", "")
	response, err := c.client.Get(getRuleUrl)

	body, err := ioutil.ReadAll(response.Body)

	var result []model.Rule
	err = json.Unmarshal(body, &result)

	return result, err
}

/**
 * Post a new rule by the given path.
 * @param {string} rule
 */
func (c *CatalogService) PostRule(rule model.Rule) (model.Rule, error) {
	postRuleUrl := c.BuildURL(CATALOG_SERVICE_PREFIX, "rules", "")
	response, err := c.client.Post(postRuleUrl, rule)

	body, err := ioutil.ReadAll(response.Body)

	var result model.Rule
	err = json.Unmarshal(body, &result)

	return result, err
}