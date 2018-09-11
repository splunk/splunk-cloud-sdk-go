package model

// CreateJobRequest defines properties allowed (and possibly required) in fully constructed Searchjobs in POST payloads and responses
type CreateJobRequest struct {
	// The SPL query string.
	Query string `json:"query"`
	// The module to run the search in.
	Module string `json:"module"`
	// Should SplunkD produce all fields (including those not explicitly mentioned in the SPL)
	ExtractAllFields bool `json:"extractAllFields"`
	// The number of seconds to run this search before finalizing.
	MaxTime uint `json:"maxTime,omitempty"`
	// Used to convert a formatted time string from {start,end}_time into UTC seconds. The default value is the ISO-8601 format.
	TimeFormat string `json:"timeFormat,omitempty"`
	// The System time at the time the search job was created. Specify a time string to set
	// the absolute time used for any relative time specifier in the search.
	// Defaults to the current system time when the Job is created.
	TimeOfSearch string `json:"timeOfSearch,omitempty"`
	// The earliest time in absolute or relative format to retrieve events (only supported if the query supports time-ranges)
	// Default is last 24 hour
	Earliest string `json:"earliest,omitempty"`
	// The latest time in absolute or relative format to retrieve events (only supported if the query supports time-ranges)
	// Default is now
	Latest string `json:"latest,omitempty"`
}

// SearchParameters is the type representing parameters currently earliest & latest on search.
type SearchParameters struct {
	// The earliest time in absolute or relative format to retrieve events (only supported if the query supports time-ranges)
	Earliest string `json:"earliest,omitempty"`
	// The latest time in absolute or relative format to retrieve events (only supported if the query supports time-ranges)
	Latest string `json:"latest,omitempty"`
}

// SearchJobStatus describes status of a search job
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

// SearchJob represents fully constructed search job including readonly fields.
type SearchJob struct {
	// The SPL query string.
	Query string `json:"query"`
	// Determine whether the Search service extracts all available fields in the data, including fields not mentioned in the SPL for the search job.
	// Set to 'false' for better search performance. Defaults to 'false'.
	ExtractAllFields bool `json:"extractAllFields"`
	// Converts a formatted time string from {start,end}_time into UTC seconds. The default value is the ISO-8601 format.
	TimeFormat string `json:"timeFormat,omitempty"`
	// The module to run the search in.
	Module string `json:"module,omitempty"`
	// The number of seconds to run this search before finalizing.
	MaxTime uint `json:"maxTime,omitempty"`
	// The System time at the time the search job was created. Specify a time string to set the absolute time used for any relative time specifier in the search.
	// Defaults to the current system time when the search job is created.
	TimeOfSearch string `json:"timeOfSearch,omitempty"`
	// Represents parameters on the search job such as 'earliest' and 'latest'.
	Parameters SearchParameters `json:"parameters,omitempty"`
	// The id assigned to the search job
	Id string `json:"sid,omitempty"`
	// The current status of the job
	Status SearchJobStatus `json:"status,omitempty"`
	// An estimate of how close the job is to completing.
	PercentComplete float64 `json:"percentComplete,omitempty"`
	// The number of results produced so far for the search job.
	ResultsAvailable int64 `json:"resultsAvailable,omitempty"`
	// Run time messages from Splunkd.
	Messages SearchJobMessages `json:"messages,omitempty"`
}

// JobAction defines actions to be taken on an existing search job.
type JobAction string

const (
	CancelJob   JobAction = "cancel"
	FinalizeJob JobAction = "finalize"
	TouchJob    JobAction = "touch"
	SaveJob     JobAction = "save"
)

type JobMessageType string

const (
	InfoType  JobMessageType = "INFO"
	FatalType JobMessageType = "FATAL"
	ErrorType JobMessageType = "ERROR"
	DebugType JobMessageType = "DEBUG"
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
	Count  int `key:"count"`
	Offset int `key:"offset"`
}

// SearchResults represents results from a search job
type SearchResults struct {
	Preview    bool `json:"preview"`
	InitOffset int  `json:"init_offset"`
	// Run time messages from Splunkd.
	Messages SearchJobMessages        `json:"messages"`
	Results  []map[string]interface{} `json:"results"`
	Fields   []map[string]interface{} `json:"fields"`
}

// ResultsNotReadyResponse represents the response when no search results is ready
type ResultsNotReadyResponse struct {
	// URL for job results
	NextLink string `json:"nextLink,omitempty"`
	// Number of milliseconds to wait before retrying
	Wait string `json:"wait,omitempty"`
}
