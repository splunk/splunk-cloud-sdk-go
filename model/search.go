package model

import (
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
)

// CreateJobRequest is Deprecated: please use services/search.CreateJobRequest
type CreateJobRequest = search.CreateJobRequest

// QueryParameters is Deprecated: please use services/search.QueryParameters
type QueryParameters = search.QueryParameters

// SearchJobStatus is Deprecated: please use services/search.JobStatus
type SearchJobStatus = search.JobStatus

const (
	// Queued is Deprecated: please use services/search.Queued
	Queued SearchJobStatus = search.Queued
	// Parsing is Deprecated: please use services/search.Parsing
	Parsing SearchJobStatus = search.Parsing
	// Running is Deprecated: please use services/search.Running
	Running SearchJobStatus = search.Running
	// Finalizing is Deprecated: please use services/search.Finalizing
	Finalizing SearchJobStatus = search.Finalizing
	// Failed is Deprecated: please use services/search.Failed
	Failed SearchJobStatus = search.Failed
	// Done is Deprecated: please use services/search.Done
	Done SearchJobStatus = search.Done
)

// SearchJob is Deprecated: please use services/search.Job
type SearchJob = search.Job

// JobStatus is Deprecated: please use services/search.JobAction
type JobStatus = search.JobAction

const (
	// JobCanceled is Deprecated: please use services/search.JobCanceled
	JobCanceled JobStatus = search.JobCanceled
	// JobFinalized is Deprecated: please use services/search.JobFinalized
	JobFinalized JobStatus = search.JobFinalized
)

// JobMessageType is Deprecated: please use services/search.JobMessageType
type JobMessageType = search.JobMessageType

const (
	// InfoType is Deprecated: please use services/search.InfoType
	InfoType JobMessageType = search.InfoType
	// FatalType is Deprecated: please use services/search.FatalType
	FatalType JobMessageType = search.FatalType
	// ErrorType is Deprecated: please use services/search.ErrorType
	ErrorType JobMessageType = search.ErrorType
	// DebugType is Deprecated: please use services/search.DebugType
	DebugType JobMessageType = search.DebugType
)

// SearchJobMessages is Deprecated: please use services/search.JobMessages
type SearchJobMessages = search.JobMessages

// PatchJobResponse is Deprecated: please use services/search.PatchJobResponse
type PatchJobResponse = search.PatchJobResponse

// JobResultsParams is Deprecated: please use services/search.JobResultsParams
type JobResultsParams = search.JobResultsParams

// SearchResults is Deprecated: please use services/search.Results
type SearchResults = search.Results

// ResultsNotReadyResponse is Deprecated: please use services/search.ResultsNotReadyResponse
type ResultsNotReadyResponse = search.ResultsNotReadyResponse
