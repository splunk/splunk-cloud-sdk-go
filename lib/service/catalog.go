/*
Package service implements a service client which is used to communicate
with Splunkd endpoints as well as a collection of services that group
logically related Splunkd endpoints.
*/
package service

import (
	//"bytes"
	//"crypto/tls"
	//"encoding/json"
	//"io"
	//"net/http"
	//"net/url"
	"path"
	//"reflect"
	//"strconv"
	//"strings"
	//"time"
	//
	//"github.com/splunk/ssc-client-go/lib/util"
	//"github.com/go-openapi/runtime/client"
	"net/url"
	//"path"
	"github.com/splunk/ssc-client-go/lib/util"
	"fmt"
)

const SEARCH_SERVICE_PREFIX string = "/search/v1";
const EVENT_SERVICE_PREFIX string = "/v1";
const CATALOG_SERVICE_PREFIX string = "/catalog/v1";

type CatalogService service

type DatasetKind string
const (
	VIEW DatasetKind ="view"
	INDEX DatasetKind ="index"
	KVSTORE DatasetKind ="kvstore"
)

type Dataset struct {
	kind DatasetKind
	todo string
}

type Datasets []Dataset

type ActionKind string
const (
	ALIAS ActionKind ="ALIAS"
	AUTOKV ActionKind ="AUTOKV"
	REGEX ActionKind ="REGEX"
	EVAL ActionKind ="EVAL"
	LOOKUP ActionKind ="LOOKUP"
)

type Rule struct {
	name string
	action []ActionKind
	match string
	priority int
	description string

}
type Rules []Rule


// BuildSplunkdURL creates full Splunkd URL
func (c *CatalogService) BuildURL(prefix string, path string ) url.URL {
	return url.URL{
		Scheme:   defaultScheme,
		Path:     CATALOG_SERVICE_PREFIX,
		RawQuery: path,
		Host: "localhost:32769",
	}
}

func (c *CatalogService) GetDatasets() (Datasets) {
	var ds Dataset = Dataset{VIEW, "dfdsf"};
	//var url = c.BuildURL(CATALOG_SERVICE_PREFIX,"datasets")
	//response, err := c.client.Get(url)


	//fmt.Print(err)
	//fmt.Print(response)
	return Datasets{ds}
}

func (c *CatalogService) GetRules() (Rules) {
	var ds Rule = Rule{"rule1", []ActionKind{LOOKUP},"match",9,"something"};
	return Rules{ds}
}


/**
 * Delete the rule by the given path.
 * @param {string} rulePath
 */
func (c *CatalogService) DeleteRule(rulePath string) error {
	buildPath := ""
	buildPath = path.Join(buildPath, "/rules", rulePath)
	deleteJobUrl := c.BuildURL(CATALOG_SERVICE_PREFIX, buildPath)
	fmt.Println(deleteJobUrl)
	response, err := c.client.Delete(deleteJobUrl)
	return util.ParseError(response, err)
}

/*func (c *CatalogService) PostRules() error {

}*/


