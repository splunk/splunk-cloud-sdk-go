package service

import (
	"io/ioutil"
	"fmt"
)


// SearchService implements a new service type
type SearchService service

// NewSearch dispatches a new spl search and returns sid
func (service *SearchService) NewSearch(spl string) (string, error) {
	
	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs")
	response, err := service.client.Post(jobURL, map[string]string{"query": spl})
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	return string(body), err
}
