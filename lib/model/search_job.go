package model

// PostJobsRequest contains params for creating a search job
type PostJobsRequest struct {
	Query    string `json:"query"`   // The SPL query string. (Required)
	Timeout  int    `json:"timeout"` // Cancel the search after this many seconds of inactivity. Set to 0 to disable timeout. (Default 30)
	TTL      int    `json:"ttl"`     // The time, in seconds, after the search has completed until the search job expires and results are deleted.
	Limit    int64  `json:"limit"`   // The number of events to process before the job is automatically finalized. Set to 0 to disable automatic finalization.
}
