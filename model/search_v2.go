package model

// SearchJobBase defines properties allowed (and possibly required) in fully constructed Searchjobs in POST payloads and responses
type SearchJobBase struct {
	// The SPL query string.
	Query            string `json:"query"`
	// The module to run the search in.
	Module           string `json:"module"`
	// Should SplunkD produce all fields (including those not explicitly mentioned in the SPL)
	ExtractAllFields bool   `json:"extractAllFields"`
	// The number of seconds to run this search before finalizing.
	MaxTime          uint   `json:"maxTime,omitempty"`
	// Used to convert a formatted time string from {start,end}_time into UTC seconds. The default value is the ISO-8601 format.
	TimeFormat   string           `json:"timeFormat,omitempty"`
	// The System time at the time the search job was created. Specify a time string to set
	// the absolute time used for any relative time specifier in the search.
	// Defaults to the current system time when the Job is created.
	TimeOfSearch string           `json:"timeOfSearch,omitempty"`
	// Parameters for the search job, mainly earliest and latest.
	Parameters   SearchParameters `json:"parameters,omitempty"`
}

// SearchParameters is the type representing parameters currently earliest & latest on search.
type SearchParameters struct {
	// The earliest time in absolute or relative format to retrieve events (only supported if the query supports time-ranges)
	Earliest string `json:"earliest,omitempty"`
	// The latest time in absolute or relative format to retrieve events (only supported if the query supports time-ranges)
	Latest   string `json:"latest,omitempty"`
}

// JobStatus describes status of a search job
type SearchJobStatus string

// Supported SearchJobStatus constants
const (
	Queued     SearchJobStatus = "queued"
	Parsing    SearchJobStatus = "parsing"
	Running    SearchJobStatus = "running"
	Finalizing SearchJobStatus = "finalizing"
	Failed     SearchJobStatus = "failed"
	Done       SearchJobStatus = "done"
)

// JobStatus is a single job status common info which is used by CatalogJobStatus and PostJobResponse
type JobStatus struct {
	// The id assigned to the search job
	Id    string `json:"sid"`
	// The current status of the job
	Status SearchJobStatus `json:"status"`
	*SearchJobBase
	// An estimate of how far through the job is complete
	PercentComplete  float64 `json:"percentComplete"`
	// The number of results Splunkd produced so far
	ResultsAvailable int64   `json:"resultsAvailable"`
}

// SearchJob represents fully constructed search job including readonly fields.
type SearchJob struct {
	*JobStatus
	// Run time messages from Splunkd.
	Messages SearchJobMessages `json:"messages"`
}

// PatchJobRequest represents the request payload to the patch jobs endpoint
type PatchJobRequest struct {
	Action JobAction `json:"action"`
	TTL    uint   `json:"ttl"`
}

// JobAction defines actions to be taken on an existing search job.
type JobAction string
const (
	Cancel JobAction = "cancel"
	Finalize JobAction = "finalize"
	Touch JobAction = "touch"
	Save JobAction = "save"
)

// SearchJobMessages is used in search results or search job.
type SearchJobMessages []struct {
	// Enum [INFO, FATAL, ERROR, DEBUG]
	Type string `json:"type"`
	// Message text
	Text string `json:"text"`
}

// PatchJobResponse defines the response from patch endpoint
type PatchJobResponse struct {
	// Run time messages from Splunkd.
	Messages SearchJobMessages `json:"messages"`
}

// JobResultsParams specifies the query params when fetching job results
type JobResultsParams struct {
	Count      int      `key:"count"`
	Offset     int      `key:"offset"`
}

// SearchResults represents results from a search job
type SearchResults struct {
	Preview     bool                     `json:"preview"`
	InitOffset  int                      `json:"init_offset"`
	// Run time messages from Splunkd.
	Messages    SearchJobMessages        `json:"messages"`
	Results     []map[string]interface{} `json:"results"`
	Fields      []map[string]interface{} `json:"fields"`
}