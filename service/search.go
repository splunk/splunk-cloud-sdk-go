package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/util"
	"io/ioutil"
	"strconv"
	"errors"
	"net/url"
	"time"
)

const searchServicePrefix = "search"
const searchServiceVersion = "v1"

// Search is a wrapper class for convenient search operations
type Search struct {
	Sid string
	Service *SearchService
}
// Status returns the status of the search job
func (search *Search) Status() (*model.SearchJobContent, error) {
	return search.Service.GetJob(search.Sid)
}
// Cancel posts a cancel action to the search job
func (search *Search) Cancel() (*model.JobControlReplyMsg, error) {
	return search.Service.PostJobControl(search.Sid, &model.JobControlAction{Action: model.CANCEL})
}
// Pause posts a pause action to the search job
func (search *Search) Pause() (*model.JobControlReplyMsg, error) {
	return search.Service.PostJobControl(search.Sid, &model.JobControlAction{Action: model.PAUSE})
}
// Unpause posts an unpause action to the search job
func (search *Search) Unpause() (*model.JobControlReplyMsg, error) {
	return search.Service.PostJobControl(search.Sid, &model.JobControlAction{Action: model.UNPAUSE})
}
// Touch posts a touch action to the search job
func (search *Search) Touch() (*model.JobControlReplyMsg, error) {
	return search.Service.PostJobControl(search.Sid, &model.JobControlAction{Action: model.TOUCH})
}
// SetTTL posts a setttl action to the search job
func (search *Search) SetTTL() (*model.JobControlReplyMsg, error) {
	return search.Service.PostJobControl(search.Sid, &model.JobControlAction{Action: model.SETTTL})
}
// Finalize posts a finalize action to the search job
func (search *Search) Finalize() (*model.JobControlReplyMsg, error) {
	return search.Service.PostJobControl(search.Sid, &model.JobControlAction{Action: model.FINALIZE})
}
// Save posts a save action to the search job
func (search *Search) Save() (*model.JobControlReplyMsg, error) {
	return search.Service.PostJobControl(search.Sid, &model.JobControlAction{Action: model.SAVE})
}
// EnablePreview posts an enablepreview action to the search job
func (search *Search) EnablePreview() (*model.JobControlReplyMsg, error) {
	return search.Service.PostJobControl(search.Sid, &model.JobControlAction{Action: model.ENABLEPREVIEW})
}
// DisablePreview posts a disablepreview action to the search job
func (search *Search) DisablePreview() (*model.JobControlReplyMsg, error) {
	return search.Service.PostJobControl(search.Sid, &model.JobControlAction{Action: model.DISABLEPREVIEW})
}
// Wait polls the job until it's completed or errors out
func (search *Search) Wait() error {
	return search.Service.WaitForJob(search.Sid, 250 * time.Millisecond)
}

// SearchService talks to the SSC search service
type SearchService service

// GetJobs gets details of all current searches.
func (service *SearchService) GetJobs(params *model.JobsRequest) ([]model.SearchJob, error) {
	var jobs []model.SearchJob
	if params == nil {
		params = model.NewDefaultPaginationParams()
	}
	jobsURL, err := service.client.BuildURL(util.ParseURLParams(*params), searchServicePrefix, searchServiceVersion, "jobs")
	if err != nil {
		return jobs, err
	}
	response, err := service.client.Get(jobsURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&jobs, response)
	return jobs, err
}

// CreateJob dispatches a search and returns sid.
func (service *SearchService) CreateJob(job *model.PostJobsRequest) (string, error) {
	// var postJobResponse model.PostJobResponse
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs")
	if err != nil {
		return "", err
	}
	response, err := service.client.Post(jobURL, job)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	jobID, err := strconv.Unquote(string(body))
	if err != nil {
		return jobID, errors.New("unable to parse jobID")
	}
	return jobID, err
}

// GetJob retrieves information about the specified search.
func (service *SearchService) GetJob(jobID string) (*model.SearchJobContent, error) {
	var jobsResponse model.SearchJobContent
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs", jobID)
	response, err := service.client.Get(jobURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&jobsResponse, response)
	return &jobsResponse, err
}

// PostJobControl runs a job control command for the specified search.
func (service *SearchService) PostJobControl(jobID string, action *model.JobControlAction) (*model.JobControlReplyMsg, error) {
	var msg model.JobControlReplyMsg
	jobURL, err := service.client.BuildURL(nil, searchServicePrefix, searchServiceVersion, "jobs", jobID, "control")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Post(jobURL, action)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&msg, response)
	return &msg, err
}

// GetJobResults Returns the job results with the given `id`.
func (service *SearchService) GetJobResults(jobID string, params *model.FetchResultsRequest) (*model.SearchResults, error) {
	var results model.SearchResults
	var query url.Values
	if params != nil {
		query = util.ParseURLParams(*params)
	}
	jobURL, err := service.client.BuildURL(query, searchServicePrefix, searchServiceVersion, "jobs", jobID, "results")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Get(jobURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&results, response)
	return &results, err
}

// GetJobEvents Returns the job events with the given `id`.
func (service *SearchService) GetJobEvents(jobID string, params *model.FetchEventsRequest) (*model.SearchResults, error) {
	var results model.SearchResults
	var query url.Values
	if params != nil {
		query = util.ParseURLParams(*params)
	}
	jobURL, err := service.client.BuildURL(query, searchServicePrefix, searchServiceVersion, "jobs", jobID, "events")
	if err != nil {
		return nil, err
	}
	response, err := service.client.Get(jobURL)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&results, response)
	return &results, err
}

// WaitForJob polls the job until it's completed or errors out
func (service *SearchService) WaitForJob(sid string, pollInterval time.Duration) error {
	var done bool
	for !done {
		job, err := service.GetJob(sid)
		if err != nil {
			return err
		}
		done = job.DispatchState == model.DONE
		time.Sleep(pollInterval)
	}
	return nil
}
// SubmitSearch creates a search job and wraps the response in an object
func (service *SearchService) SubmitSearch(job *model.PostJobsRequest) (*Search, error) {
	sid, err := service.CreateJob(job)
	if err != nil {
		return nil, err
	}
	return &Search{
		Sid: sid,
		Service: service,
	}, nil
}