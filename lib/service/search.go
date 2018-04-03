package service

import (
	"github.com/splunk/ssc-client-go/lib/model"
	"github.com/splunk/ssc-client-go/lib/util"
)


// SearchService implements a new service type
type SearchService service

// NewSearch dispatches a new spl search and returns sid
func (service *SearchService) NewSearch(spl string) (*model.SearchJob, error) {
	var job = model.NewSearchJob(service)

	jobURL := service.client.BuildSplunkdURL(nil, "search", "v1", "jobs")
	response, err := service.client.Post(jobURL, map[string]string{"query": spl})
	err = util.ParseResponse(&job, response, err)
	return &job, err
}
