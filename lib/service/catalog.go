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
)

const CATALOG_SERVICE_PREFIX string = "/catalog/v1";

type CatalogService service

type DatasetKind string

const (
	VIEW    DatasetKind = "view"
	INDEX   DatasetKind = "index"
	KVSTORE DatasetKind = "kvstore"
	EXTERN  DatasetKind = "extern"
	TOPIC   DatasetKind = "topic"
	CATALOG DatasetKind = "catalog"
)

type Dataset struct {
	Id    string      `json:"id"`
	Name  string      `json:"name"`
	Kind  DatasetKind `json:"kind"`
	Rules []string    `json:"rules"`
	Todo  string      `json:"todo"`
}

type Datasets []Dataset

type ActionKind string

const (
	ALIAS  ActionKind = "ALIAS"
	AUTOKV ActionKind = "AUTOKV"
	REGEX  ActionKind = "REGEX"
	EVAL   ActionKind = "EVAL"
	LOOKUP ActionKind = "LOOKUP"
)

type Rule struct {
	Name        string       `json:"name"`
	Action      []ActionKind `json:"actions"`
	Match       string       `json:"match"`
	Priority    int          `json:"priority"`
	Description string       `json:"description"`
}
type Rules []Rule

// creates a catalog URL //todo: move to client.go or other common files
func (c *CatalogService) BuildURL(prefix string, path string, query string) url.URL {
	return url.URL{
		Scheme:   c.client.Scheme,
		Path:     CATALOG_SERVICE_PREFIX + "/" + path,
		RawQuery: query,
		Host:     c.client.Host,
	}
}

func (c *CatalogService) GetDatasets() (Datasets, error) {
	var url = c.BuildURL(CATALOG_SERVICE_PREFIX, "datasets", "")
	response, err := c.client.Get(url)

	body, err := ioutil.ReadAll(response.Body)

	var result Datasets
	err = json.Unmarshal(body, &result)

	return result, err
}

func (c *CatalogService) GetDataset(name string) (Dataset, error) {
	var url = c.BuildURL(CATALOG_SERVICE_PREFIX, "datasets"+"/"+name, "")
	response, err := c.client.Get(url)
	body, err := ioutil.ReadAll(response.Body)

	var result Dataset
	err = json.Unmarshal(body, &result)

	return result, err
}

func (c *CatalogService) GetRules() (Rules) {
	var ds Rule = Rule{"rule1", []ActionKind{LOOKUP}, "match", 9, "something"};
	return Rules{ds}
}
