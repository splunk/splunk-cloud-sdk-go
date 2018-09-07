// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"errors"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/model"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const searchServicePrefix = "search"
const searchServiceVersion = "v1"

// Search is a wrapper class for convenient search operations
type Search struct {
	sid          string
	svc          *SearchService
	isCancelling bool
}

// Status returns the status of the search job
func (search *Search) Status() (*model.SearchJobContent, error) {
	return search.svc.GetJob(search.sid)
}

// Cancel posts a cancel action to the search job
func (search *Search) Cancel() (*model.JobControlReplyMsg, error) {
	search.isCancelling = true
	return search.svc.PostJobControl(search.sid, &model.JobControlAction{Action: model.CANCEL})
}

// Touch posts a touch action to the search job
func (search *Search) Touch() (*model.JobControlReplyMsg, error) {
	return search.svc.PostJobControl(search.sid, &model.JobControlAction{Action: model.TOUCH})
}

// SetTTL posts a setttl action to the search job
func (search *Search) SetTTL(ttl int) (*model.JobControlReplyMsg, error) {
	return search.svc.PostJobControl(search.sid, &model.JobControlAction{Action: model.SETTTL, TTL: ttl})
}

// Finalize posts a finalize action to the search job
func (search *Search) Finalize() (*model.JobControlReplyMsg, error) {
	return search.svc.PostJobControl(search.sid, &model.JobControlAction{Action: model.FINALIZE})
}

// Save posts a save action to the search job
func (search *Search) Save() (*model.JobControlReplyMsg, error) {
	return search.svc.PostJobControl(search.sid, &model.JobControlAction{Action: model.SAVE})
}

// EnablePreview posts an enablepreview action to the search job
func (search *Search) EnablePreview() (*model.JobControlReplyMsg, error) {
	return search.svc.PostJobControl(search.sid, &model.JobControlAction{Action: model.ENABLEPREVIEW})
}

// DisablePreview posts a disablepreview action to the search job
func (search *Search) DisablePreview() (*model.JobControlReplyMsg, error) {
	return search.svc.PostJobControl(search.sid, &model.JobControlAction{Action: model.DISABLEPREVIEW})
}

// Wait polls the job until it's completed or errors out
func (search *Search) Wait() error {
	err := search.svc.WaitForJob(search.sid, 250*time.Millisecond)
	if search.isCancelling == true {
		return errors.New("search has been cancelled")
	}
	if err != nil && err.(*util.HTTPError).HTTPStatusCode == 404 {
		return errors.New("search has been cancelled")
	}
	return err
}

// GetEvents returns events from the search
func (search *Search) GetEvents(params *model.FetchEventsRequest) (*model.SearchResults, error) {
	err := search.Wait()
	if err != nil {
		return nil, err
	}
	return search.svc.GetJobEvents(search.sid, params)
}

// GetResults returns results from the search
func (search *Search) GetResults(params *model.FetchResultsRequest) (*model.SearchResults, error) {
	err := search.Wait()
	if err != nil {
		return nil, err
	}
	return search.svc.GetJobResults(search.sid, params)
}

// QueryEvents waits for job to complete and returns an iterator. If offset and batchSize are specified,
// the iterator will return that window of results with each Next() call
func (search *Search) QueryEvents(batchSize, offset int, params *model.FetchEventsRequest) (*SearchIterator, error) {
	err := search.Wait()
	if err != nil {
		return nil, err
	}
	jobStatus, err := search.Status()
	if err != nil {
		return nil, err
	}
	iterator := NewSearchIterator(batchSize, offset, jobStatus.EventCount,
		func(count, offset int) (*model.SearchResults, error) {
			params.Count = count
			params.Offset = offset
			return search.GetEvents(params)
		})
	return iterator, nil
}

// QueryResults waits for job to complete and returns an iterator. If offset and batchSize are specified,
// the iterator will return that window of results with each Next() call
func (search *Search) QueryResults(batchSize, offset int, params *model.FetchResultsRequest) (*SearchIterator, error) {
	err := search.Wait()
	if err != nil {
		return nil, err
	}
	jobStatus, err := search.Status()
	if err != nil {
		return nil, err
	}
	if jobStatus.EventCount == 0 {
		return nil, errors.New("no results are retrieved from the search")
	}
	iterator := NewSearchIterator(batchSize, offset, jobStatus.EventCount,
		func(count, offset int) (*model.SearchResults, error) {
			params.Count = count
			params.Offset = offset
			return search.GetResults(params)
		})
	return iterator, nil
}

// SearchService talks to the Splunk Cloud search service
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
	response, err := service.client.Get(RequestParams{URL: jobsURL})
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
	response, err := service.client.Post(RequestParams{URL: jobURL, Body: job})
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
	response, err := service.client.Get(RequestParams{URL: jobURL})
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
	response, err := service.client.Post(RequestParams{URL: jobURL, Body: action})
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
	response, err := service.client.Get(RequestParams{URL: jobURL})
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
	response, err := service.client.Get(RequestParams{URL: jobURL})
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
		if job.DispatchState == model.DONE {
			done = true
		} else if job.DispatchState == model.FAILED {
			return errors.New("job failed")
		} else {
			time.Sleep(pollInterval)
		}
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
		sid:          sid,
		svc:          service,
		isCancelling: false,
	}, nil
}
