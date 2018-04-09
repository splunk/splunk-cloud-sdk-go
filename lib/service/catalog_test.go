/*
Package service implements a service client which is used to communicate
with Splunkd endpoints as well as a collection of services that group
logically related Splunkd endpoints.
*/
package service

import (

	"fmt"
	"testing"
)


/*func Test_dataset(t *testing.T) {
	var splunkClient = NewSplunkdClient("",
		[2]string{"admin", "changeme"},
		"localhost:8889", "https",nil)

	//client := &Client{Host:"http://localhost:8882", Auth: [2]string{"admin", "changeme"} }
	url := splunkClient.CatalogService.GetDatasets()
	fmt.Print(url)
	//t.Error("hahah")

}*/


func TestDeleteDataRule(t *testing.T) {
	var splunkClient = NewSplunkdClient("",
		[2]string{"admin", "changeme"},
		"localhost:32769", "https",nil)

	urlDeleteErr := splunkClient.CatalogService.DeleteRule("rule1")
	fmt.Println(urlDeleteErr)
}
