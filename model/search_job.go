package model

// PostJobsRequest contains params for creating a search job
type PostJobsRequest struct {
	Query   string `json:"query"`   // The SPL query string. (Required)
	Timeout int    `json:"timeout"` // Cancel the search after this many seconds of inactivity. Set to 0 to disable timeout. (Default 30)
	TTL     int    `json:"ttl"`     // The time, in seconds, after the search has completed until the search job expires and results are deleted.
	Limit   int64  `json:"limit"`   // The number of events to process before the job is automatically finalized. Set to 0 to disable automatic finalization.
}

// PostJobResponse contains a SearchID
type PostJobResponse struct {
	SearchID string `json:"searchId"` // The SearchID returned for the newly created search job.
}

// PaginationParams specifies pagination parameters for certain supported requests
type JobsRequest struct {
	Count  uint `key:"count"`
	Offset uint `key:"offset"`
}

// NewDefaultPaginationParams creates search pagination parameters according to Splunk Enterprise defaults
func NewDefaultPaginationParams() *JobsRequest {
	return &JobsRequest{
		Count:  30,
		Offset: 0,
	}
}
// SearchJob specifies the fields returned for a /search/jobs/ entry for a specific job
type SearchJob struct {
	Sid           string           `json:"sid"`
	Content       SearchJobContent `json:"content"`
	Context       *SearchContext
}
// SearchJobContent represents the content json object from /search/jobs response
type SearchJobContent struct {
	Sid              string                 `json:"sid"`
	EventCount       int                    `json:"eventCount"`
	DispatchState    string                 `json:"dispatchState"`
	DiskUsage        int64                  `json:"diskUsage"`
	IsFinalized      bool                   `json:"isFinalized"`
	OptimizedSearch  string                 `json:"optimizedSearch"`
	ScanCount        int64                  `json:"scanCount"`
	AdditionalFields map[string]interface{} `json:"-"`
}

// SearchContext specifies the user and app context for a search job
type SearchContext struct {
	User string
	App  string
}