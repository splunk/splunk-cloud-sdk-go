// Code generated by gen-interface.go. DO NOT EDIT.
package search

import (
	"time"
)

type Servicer interface {
	// ListJobs gets the matching list of search jobs
	ListJobs() ([]Job, error)
	// ListJobsByQueryParameters gets the matching list of search jobs filtered by query parameters specified
	ListJobsByQueryParameters(query JobsQuery) ([]Job, error)
	// CreateJob creates a new search job
	CreateJob(job *CreateJobRequest) (*Job, error)
	// GetJob retrieves information about the specified search.
	GetJob(jobID string) (*Job, error)
	// UpdateJob updates an existing job with actions and TTL
	UpdateJob(jobID string, jobStatus JobAction) (*PatchJobResponse, error)
	// GetResults Returns the job results with the given `id`. count=0 returns default number of results from search
	GetResults(jobID string, count, offset int) (interface{}, error)
	// WaitForJob polls the job until it's completed or errors out
	WaitForJob(jobID string, pollInterval time.Duration) (interface{}, error)
}