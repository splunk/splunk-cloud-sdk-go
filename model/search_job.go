package model

// DispatchState describes dispatchState of a job
type DispatchState string
// Supported DispatchState constants
const (
	QUEUED     DispatchState = "QUEUED"
	PARSING    DispatchState = "PARSING"
	RUNNING    DispatchState = "RUNNING"
	PAUSED     DispatchState = "PAUSED"
	FINALIZING DispatchState = "FINALIZING"
	FAILED     DispatchState = "FAILED"
	DONE       DispatchState = "DONE"
)

// PostJobsRequest represents the search job post params
type PostJobsRequest struct {

	//sample_ratio this is not documented on docs.splunk.com
	SampleRatio string `json:"sampleRatio"`

	//adhoc_search_level    String        Use one of the following search modes.
	//[ verbose | fast | smart ]
	AdhocSearchLevel string `json:"adhocSearchLevel"`

	//auto_cancel    Number    0    If specified, the job automatically cancels after this many seconds of inactivity. (0 means never auto-cancel)
	AutoCancel *uint `json:"autoCancel"`

	//auto_finalize_ec    Number    0    Auto-finalize the search after at least this many events are processed.
	//Specify 0 to indicate no limit.
	AutoFinalizeEventCount *uint `json:"autoFinalizeEc"`

	//auto_pause    Number    0    If specified, the search job pauses after this many seconds of inactivity. (0 means never auto-pause.)
	//To restart a paused search job, specify unpause as an action to POST search/jobs/{search_id}/control.
	//auto_pause only goes into effect once. Unpausing after auto_pause does not put auto_pause into effect again.
	AutoPause *uint `json:"autoPause"`

	//custom fields
	CustomFields map[string]interface{}

	//earliest_time    String        Specify a time string. Sets the earliest (inclusive), respectively, time bounds for the search.
	//The time string can be either a UTC time (with fractional seconds), a relative time specifier (to now) or a formatted time string. Refer to Time modifiers for search for information and examples of specifying a time string.
	//Compare to index_earliest parameter. Also see comment for the search_mode parameter.
	EarliestTime string `json:"earliestTime"`

	//enable_lookups    Boolean    true    Indicates whether lookups should be applied to events.
	//Specifying true (the default) may slow searches significantly depending on the nature of the lookups.
	EnableLookUps *bool `json:"enableLookups"`

	//exec_mode    Enum    normal    Valid values: (blocking | oneshot | normal)
	//If set to normal, runs an asynchronous search.
	//If set to blocking, returns the sid when the job is complete.
	//If set to oneshot, returns results in the same call. In this case, you can specify the format for the output (for example, json output) using the output_mode parameter as described in GET search/jobs/export. Default format for output is xml.
	ExecuteMode string `json:"execMode"`

	//id    String        Optional string to specify the search ID (<sid>). If unspecified, a random ID is generated.
	ID string `json:"id"`

	//latest_time    String        Specify a time string. Sets the latest (exclusive), respectively, time bounds for the search.
	//The time string can be either a UTC time (with fractional seconds), a relative time specifier (to now) or a formatted time string.
	//Refer to Time modifiers for search for information and examples of specifying a time string.
	//Compare to index_latest parameter. Also see comment for the search_mode parameter.
	LatestTime string `json:"latestTime"`

	//limit max events
	Limit int64 `json:"limit"`
	//max_count    Number    10000    The number of events that can be accessible in any given status bucket.
	//Also, in transforming mode, the maximum number of results to store. Specifically, in all calls, codeoffset+count max_count.
	MaxCount *uint `json:"maxCount"`

	//max_time    Number    0    The number of seconds to run this search before finalizing. Specify 0 to never finalize.
	MaxTime *uint `json:"maxTime"`

	//module	String		The Module to run the search in.
	Module string `json:"module"`

	//now    String    current system time    Specify a time string to set the absolute time used for any relative time specifier in the search. Defaults to the current system time.
	//You can specify a relative time modifier for this parameter. For example, specify +2d to specify the current time plus two days.
	//If you specify a relative time modifier both in this parameter and in the search string, the search string modifier takes precedence.
	//Refer to Time modifiers for search for details on specifying relative time modifiers.
	Now string `json:"now"`

	//reduce_freq    Number    0    Determines how frequently to run the MapReduce reduce phase on accumulated map values.
	ReduceFrequency *uint `json:"reduceFreq"`

	//search required    String        The search language string to execute, taking results from the local and remote servers.
	//Examples:
	//"search *"
	//"search * | outputcsv"
	Search string `json:"query"`

	//status_buckets    Number    0    The most status buckets to generate.
	//0 indicates to not generate timeline information.
	StatusBuckets *uint `json:"statusBuckets"`

	//time_format    String     %FT%T.%Q%:z    Used to convert a formatted time string from {start,end}_time into UTC seconds. The default value is the ISO-8601 format.
	TimeFormat string `json:"timeFormat"`

	//timeout    Number    86400    The number of seconds to keep this search after processing has stopped.
	Timeout *uint `json:"timeout"`

	// search timeout
	TTL int `json:"ttl"`
}

// PagingInfo captures fields returned for endpoints supporting paging
type PagingInfo struct {
	Total   float64 `json:"total"`
	PerPage float64 `json:"perPage"`
	Offset  float64 `json:"offset"`
}

// JobsRequest specifies pagination parameters for certain supported requests
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
	DispatchState    DispatchState          `json:"dispatchState"`
	DiskUsage        int64                  `json:"diskUsage"`
	IsFinalized      bool                   `json:"isFinalized"`
	OptimizedSearch  string                 `json:"optimizedSearch"`
	ScanCount        int64                  `json:"scanCount"`
}

// SearchContext specifies the user and app context for a search job
type SearchContext struct {
	User string
	App  string
}
// JobAction defines the action that can be posted on a job
type JobAction string
// Supported JobAction constants
const (
	PAUSE JobAction = "pause"
	UNPAUSE JobAction = "unpause"
	FINALIZE JobAction = "finalize"
	CANCEL JobAction = "cancel"
	TOUCH JobAction = "touch"
	SAVE JobAction = "save"
	SETTTL JobAction = "setttl"
	ENABLEPREVIEW JobAction = "enablepreview"
	DISABLEPREVIEW JobAction = "disablepreview"
)
// JobControlAction specifies the action needs to be taken on a job
type JobControlAction struct {
	Action JobAction `json:"action"`
	TTL    int    `json:"ttl"`
}

// JobControlReplyMsg displays messages returned from taking a job control
type JobControlReplyMsg struct {
	Msg []struct {
		TypeKey string `json:"type"`
		Text    string `json:"text"`
	} `json:"messages"`
}

// FetchResultsRequest specifies the query params when fetching job results
type FetchResultsRequest struct {
	Count      int      `key:"count"`
	Offset     int      `key:"offset"`
	Fields     []string `key:"f"`
	OutputMode string   `key:"outputMode"`
	Search     string   `key:"search"`
}

// FetchEventsRequest specifies the query params when fetching job events
type FetchEventsRequest struct {
	Count            int      `key:"count"`
	Offset           int      `key:"offset"`
	EarliestTime     string   `key:"earliestTime"`
	Fields           []string `key:"f"`
	LatestTime       string   `key:"latestTime"`
	MaxLines         *uint    `key:"maxLines"`
	OutputMode       string   `key:"outputMode"`
	TimeFormat       string   `key:"timeFormat"`
	OutputTimeFormat string   `key:"outputTimeFormat"`
	Search           string   `key:"search"`
	TruncationMode   string   `key:"truncationMode"`
	Segmentation     string   `key:"segmentation"`
}

// SearchResults represents the /search/jobs/{sid}/events or /search/jobs/{sid}/results response
type SearchResults struct {
	Preview     bool                     `json:"preview"`
	InitOffset  int                      `json:"init_offset"`
	Messages    []interface{}            `json:"messages"`
	Results     []*Result
	Fields      []map[string]interface{} `json:"fields"`
	Highlighted map[string]interface{}   `json:"highlighted"`
}

// Result contains information about the search
type Result struct {
	Bkt                string   `json:"_bkt"`
	Cd                 string   `json:"_cd"`
	IndexTime          string   `json:"_indextime"`
	Raw                string   `json:"_raw"`
	Serial             string   `json:"_serial"`
	Si                 []string `json:"_si"`
	OriginalSourceType string   `json:"_sourcetype"`
	Time               string   `json:"_time"`
	Entity             []string `json:"entity"`
	Host               string   `json:"host"`
	Index              string   `json:"index"`
	LineCount          string   `json:"linecount"`
	Log                string   `json:"log"`
	Punct              string   `json:"punct"`
	Source             string   `json:"source"`
	SourceType         string   `json:"sourcetype"`
	SplunkServer       string   `json:"splunk_server"`
}