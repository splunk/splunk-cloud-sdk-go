package model

// PostJobsRequest contains params for creating a search job
type PostJobsRequest struct {
	Query   string `json:"query"`   // The SPL query string. (Required)
	Timeout int    `json:"timeout"` // Cancel the search after this many seconds of inactivity. Set to 0 to disable timeout. (Default 30)
	TTL     int    `json:"ttl"`     // The time, in seconds, after the search has completed until the search job expires and results are deleted.
	Limit   int64  `json:"limit"`   // The number of events to process before the job is automatically finalized. Set to 0 to disable automatic finalization.
}

// NewSearchConfig represents the search job post params
type NewSearchConfig struct {

	//sample_ratio this is not documented on docs.splunk.com
	SampleRatio string `json:"sample_ratio"`

	//adhoc_search_level    String        Use one of the following search modes.
	//[ verbose | fast | smart ]
	AdhocSearchLevel string `json:"adhoc_search_level"`

	//auto_cancel    Number    0    If specified, the job automatically cancels after this many seconds of inactivity. (0 means never auto-cancel)
	AutoCancel *uint `json:"auto_cancel"`

	//auto_finalize_ec    Number    0    Auto-finalize the search after at least this many events are processed.
	//Specify 0 to indicate no limit.
	AutoFinalizeEventCount *uint `json:"auto_finalize_ec"`

	//auto_pause    Number    0    If specified, the search job pauses after this many seconds of inactivity. (0 means never auto-pause.)
	//To restart a paused search job, specify unpause as an action to POST search/jobs/{search_id}/control.
	//auto_pause only goes into effect once. Unpausing after auto_pause does not put auto_pause into effect again.
	AutoPause *uint `json:"auto_pause"`

	//custom fields
	CustomFields map[string]interface{}

	//earliest_time    String        Specify a time string. Sets the earliest (inclusive), respectively, time bounds for the search.
	//The time string can be either a UTC time (with fractional seconds), a relative time specifier (to now) or a formatted time string. Refer to Time modifiers for search for information and examples of specifying a time string.
	//Compare to index_earliest parameter. Also see comment for the search_mode parameter.
	EarliestTime string `json:"earliest_time"`

	//enable_lookups    Boolean    true    Indicates whether lookups should be applied to events.
	//Specifying true (the default) may slow searches significantly depending on the nature of the lookups.
	EnableLookUps *bool `json:"enable_lookups"`

	//exec_mode    Enum    normal    Valid values: (blocking | oneshot | normal)
	//If set to normal, runs an asynchronous search.
	//If set to blocking, returns the sid when the job is complete.
	//If set to oneshot, returns results in the same call. In this case, you can specify the format for the output (for example, json output) using the output_mode parameter as described in GET search/jobs/export. Default format for output is xml.
	ExecuteMode string `json:"exec_mode"`

	//id    String        Optional string to specify the search ID (<sid>). If unspecified, a random ID is generated.
	ID string `json:"id"`

	//latest_time    String        Specify a time string. Sets the latest (exclusive), respectively, time bounds for the search.
	//The time string can be either a UTC time (with fractional seconds), a relative time specifier (to now) or a formatted time string.
	//Refer to Time modifiers for search for information and examples of specifying a time string.
	//Compare to index_latest parameter. Also see comment for the search_mode parameter.
	LatestTime string `json:"latest_time"`

	//max_count    Number    10000    The number of events that can be accessible in any given status bucket.
	//Also, in transforming mode, the maximum number of results to store. Specifically, in all calls, codeoffset+count max_count.
	MaxCount *uint `json:"max_count"`

	//max_time    Number    0    The number of seconds to run this search before finalizing. Specify 0 to never finalize.
	MaxTime *uint `json:"max_time"`

	//now    String    current system time    Specify a time string to set the absolute time used for any relative time specifier in the search. Defaults to the current system time.
	//You can specify a relative time modifier for this parameter. For example, specify +2d to specify the current time plus two days.
	//If you specify a relative time modifier both in this parameter and in the search string, the search string modifier takes precedence.
	//Refer to Time modifiers for search for details on specifying relative time modifiers.
	Now string `json:"now"`

	//reduce_freq    Number    0    Determines how frequently to run the MapReduce reduce phase on accumulated map values.
	ReduceFrequency *uint `json:"reduce_freq"`

	//search required    String        The search language string to execute, taking results from the local and remote servers.
	//Examples:
	//"search *"
	//"search * | outputcsv"
	Search string `json:"search" binding:"required"`

	//status_buckets    Number    0    The most status buckets to generate.
	//0 indicates to not generate timeline information.
	StatusBuckets *uint `json:"status_buckets"`

	//time_format    String     %FT%T.%Q%:z    Used to convert a formatted time string from {start,end}_time into UTC seconds. The default value is the ISO-8601 format.
	TimeFormat string `json:"time_format"`

	//timeout    Number    86400    The number of seconds to keep this search after processing has stopped.
	Timeout *uint `json:"timeout"`
}

// PostJobResponse contains a SearchID
type PostJobResponse struct {
	SearchID string `json:"searchId"` // The SearchID returned for the newly created search job.
}

// JobsResponse represents a response that can be unmarshalled from /search/jobs
type JobResponse struct {
	Links     map[string]interface{} `json:"links"`
	Origin    string                 `json:"origin"`
	Updated   string                 `json:"updated"`
	Generator map[string]interface{} `json:"generator"`
	Entry     []SearchJobEntry       `json:"entry"`
	Paging    PagingInfo             `json:"paging"`
}

// SearchJobEntry specifies the fields returned for a /search/jobs entry
type SearchJobEntry struct {
	Name      string                 `json:"name"`
	ID        string                 `json:"id"`
	Updated   string                 `json:"updated"`
	Links     map[string]interface{} `json:"links"`
	Published string                 `json:"published"`
	Author    string                 `json:"author"`
	Content   SearchJobContent       `json:"content"`
	ACL       map[string]interface{} `json:"acl"`
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
	Search string `json:"search"`
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

//TODO: Define supported action
type JobControlAction struct {
	Action string `json:"action"`
	Ttl    int    `json:"ttl"`
}

type JobControlReplyMsg struct {
	Msg []struct {
		TypeKey string `json:"type"`
		Text    string `json:"text"`
	} `json:"messages"`
}

type FetchResultsRequest struct {
	Count      int      `json:"count" form:"count"`
	Offset     int      `json:"offset" form:"offset"`
	Fields     []string `json:"f"`
	OutputMode string   `json:"output_mode"`
	Search     string   `json:"search"`
}

type FetchEventsRequest struct {
	Count            int      `json:"count" form:"count"`
	Offset           int      `json:"offset" form:"offset"`
	EarliestTime     string   `json:"earliest_time"`
	Fields           []string `json:"f"`
	LatestTime       string   `json:"latest_time"`
	MaxLines         *uint    `json:"max_lines"`
	OutputMode       string   `json:"output_mode"`
	TimeFormat       string   `json:"time_format"`
	OutputTimeFormat string   `json:"output_time_format"`
	Search           string   `json:"search"`
	TruncationMode   string   `json:"truncation_mode"`
	Segmentation     string   `json:"segmentation"`
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