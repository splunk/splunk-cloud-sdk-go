# search
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/search"


## Usage

#### type FieldsSummary

```go
type FieldsSummary struct {
	// The amount of time, in seconds, that a time bucket spans from the earliest to the latest time.
	Duration *float32 `json:"duration,omitempty"`
	// If specified, the earliest timestamp, in UTC format, of the events to process.
	EarliestTime *string `json:"earliestTime,omitempty"`
	// The total number of events for all fields returned in the time range (earliestTime and latestTime) specified.
	EventCount *int32 `json:"eventCount,omitempty"`
	// A map of the fields in the time range specified.
	Fields map[string]SingleFieldSummary `json:"fields,omitempty"`
	// If specified, the latest timestamp, in UTC format, of the events to process.
	LatestTime *string `json:"latestTime,omitempty"`
}
```

Fields statistics summary model of the events to-date, for search ID (sid).

#### type Iterator

```go
type Iterator struct {
}
```

Iterator is the result of a search query. Its cursor starts at 0 index of the
result set. Use Next() to advance through the rows:

     search, _ := client.SearchService.SubmitSearch(&PostJobsRequest{Search: "search index=main | head 5"})
    	pages, _ := search.QueryResults(2, 0, &FetchResultsRequest{Count: 5})
    	defer pages.Close()
    	for pages.Next() {
    		values, err := pages.Value()
         ...

    	}
    	err := pages.Err() // get any error encountered during iteration
     ...

#### func  NewIterator

```go
func NewIterator(batch, offset, max int, fn QueryFunc) *Iterator
```
NewIterator creates a new reference to the iterator object

#### func (*Iterator) Close

```go
func (i *Iterator) Close()
```
Close checks the status and closes iterator if it's not already. After Close, no
results can be retrieved

#### func (*Iterator) Err

```go
func (i *Iterator) Err() error
```
Err returns error encountered during iteration

#### func (*Iterator) Next

```go
func (i *Iterator) Next() bool
```
Next prepares the next result set for reading with the Value method. It returns
true on success, or false if there is no next result row or an error occurred
while preparing it.

Every call to Value, even the first one, must be preceded by a call to Next.

#### func (*Iterator) Value

```go
func (i *Iterator) Value() (*ListSearchResultsResponse, error)
```
Value returns value in current iteration or error out if iterator is closed

#### type ListEventsSummaryQueryParams

```go
type ListEventsSummaryQueryParams struct {
	// Count : The maximum number of entries to return. Set to 0 to return all available entries.
	Count *float32 `key:"count"`
	// Earliest : The earliest time filter, in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Earliest string `key:"earliest"`
	// Field : A field to return for the result set. You can specify multiple fields of comma-separated values if multiple fields are required.
	Field string `key:"field"`
	// Latest : The latest time filter in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Latest string `key:"latest"`
	// Offset : Index of first item to return.
	Offset *float32 `key:"offset"`
}
```

ListEventsSummaryQueryParams represents valid query parameters for the
ListEventsSummary operation For convenience ListEventsSummaryQueryParams can be
formed in a single statement, for example:

    `v := ListEventsSummaryQueryParams{}.SetCount(...).SetEarliest(...).SetField(...).SetLatest(...).SetOffset(...)`

#### func (ListEventsSummaryQueryParams) SetCount

```go
func (q ListEventsSummaryQueryParams) SetCount(v float32) ListEventsSummaryQueryParams
```

#### func (ListEventsSummaryQueryParams) SetEarliest

```go
func (q ListEventsSummaryQueryParams) SetEarliest(v string) ListEventsSummaryQueryParams
```

#### func (ListEventsSummaryQueryParams) SetField

```go
func (q ListEventsSummaryQueryParams) SetField(v string) ListEventsSummaryQueryParams
```

#### func (ListEventsSummaryQueryParams) SetLatest

```go
func (q ListEventsSummaryQueryParams) SetLatest(v string) ListEventsSummaryQueryParams
```

#### func (ListEventsSummaryQueryParams) SetOffset

```go
func (q ListEventsSummaryQueryParams) SetOffset(v float32) ListEventsSummaryQueryParams
```

#### type ListFieldsSummaryQueryParams

```go
type ListFieldsSummaryQueryParams struct {
	// Earliest : The earliest time filter, in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Earliest string `key:"earliest"`
	// Latest : The latest time filter in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Latest string `key:"latest"`
}
```

ListFieldsSummaryQueryParams represents valid query parameters for the
ListFieldsSummary operation For convenience ListFieldsSummaryQueryParams can be
formed in a single statement, for example:

    `v := ListFieldsSummaryQueryParams{}.SetEarliest(...).SetLatest(...)`

#### func (ListFieldsSummaryQueryParams) SetEarliest

```go
func (q ListFieldsSummaryQueryParams) SetEarliest(v string) ListFieldsSummaryQueryParams
```

#### func (ListFieldsSummaryQueryParams) SetLatest

```go
func (q ListFieldsSummaryQueryParams) SetLatest(v string) ListFieldsSummaryQueryParams
```

#### type ListJobsQueryParams

```go
type ListJobsQueryParams struct {
	// Count : The maximum number of jobs that you want to return the status entries for.
	Count *float32 `key:"count"`
	// Status : Filter the list of jobs by status. Valid status values are &#39;running&#39;, &#39;done&#39;, &#39;canceled&#39;, or &#39;failed&#39;.
	Status *SearchStatus `key:"status"`
}
```

ListJobsQueryParams represents valid query parameters for the ListJobs operation
For convenience ListJobsQueryParams can be formed in a single statement, for
example:

    `v := ListJobsQueryParams{}.SetCount(...).SetStatus(...)`

#### func (ListJobsQueryParams) SetCount

```go
func (q ListJobsQueryParams) SetCount(v float32) ListJobsQueryParams
```

#### func (ListJobsQueryParams) SetStatus

```go
func (q ListJobsQueryParams) SetStatus(v SearchStatus) ListJobsQueryParams
```

#### type ListResultsQueryParams

```go
type ListResultsQueryParams struct {
	// Count : The maximum number of entries to return. Set to 0 to return all available entries.
	Count *float32 `key:"count"`
	// Field : A field to return for the result set. You can specify multiple fields of comma-separated values if multiple fields are required.
	Field string `key:"field"`
	// Offset : Index of first item to return.
	Offset *float32 `key:"offset"`
}
```

ListResultsQueryParams represents valid query parameters for the ListResults
operation For convenience ListResultsQueryParams can be formed in a single
statement, for example:

    `v := ListResultsQueryParams{}.SetCount(...).SetField(...).SetOffset(...)`

#### func (ListResultsQueryParams) SetCount

```go
func (q ListResultsQueryParams) SetCount(v float32) ListResultsQueryParams
```

#### func (ListResultsQueryParams) SetField

```go
func (q ListResultsQueryParams) SetField(v string) ListResultsQueryParams
```

#### func (ListResultsQueryParams) SetOffset

```go
func (q ListResultsQueryParams) SetOffset(v float32) ListResultsQueryParams
```

#### type ListSearchResultsResponse

```go
type ListSearchResultsResponse struct {
	Results  []map[string]interface{}          `json:"results"`
	Fields   []ListSearchResultsResponseFields `json:"fields,omitempty"`
	Messages []Message                         `json:"messages,omitempty"`
	NextLink *string                           `json:"nextLink,omitempty"`
	Wait     *string                           `json:"wait,omitempty"`
}
```

The structure of the search results or events metadata that is returned for the
job with the specified search ID (SID). When search is running, it might return
incomplete or truncated search results. Incomplete search results occur when a
search has not completed. Wait until search completes for full result set.
Truncated search results occur because the number of requested results exceeds
the page limit. Follow the 'nextLink' URL to retrieve the next page of results.

#### type ListSearchResultsResponseFields

```go
type ListSearchResultsResponseFields struct {
	Name           string  `json:"name"`
	DataSource     *string `json:"dataSource,omitempty"`
	GroupbyRank    *string `json:"groupbyRank,omitempty"`
	SplitField     *string `json:"splitField,omitempty"`
	SplitValue     *string `json:"splitValue,omitempty"`
	SplitbySpecial *string `json:"splitbySpecial,omitempty"`
	TypeSpecial    *string `json:"typeSpecial,omitempty"`
}
```


#### type Message

```go
type Message struct {
	Text *string      `json:"text,omitempty"`
	Type *MessageType `json:"type,omitempty"`
}
```

The message field in search results or search jobs. The types of messages are
INFO, DEBUG, FATAL, and ERROR.

#### type MessageType

```go
type MessageType string
```


```go
const (
	MessageTypeInfo  MessageType = "INFO"
	MessageTypeDebug MessageType = "DEBUG"
	MessageTypeFatal MessageType = "FATAL"
	MessageTypeError MessageType = "ERROR"
)
```
List of MessageType

#### type QueryFunc

```go
type QueryFunc func(step, start int) (*ListSearchResultsResponse, error)
```

QueryFunc is the function to be executed in each Next call of the iterator

#### type QueryParameters

```go
type QueryParameters struct {
	// The earliest time, in absolute or relative format, to retrieve events.  When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Earliest *string `json:"earliest,omitempty"`
	// The latest time, in absolute or relative format, to retrieve events.  When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format.  For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Latest *string `json:"latest,omitempty"`
	// Relative values for the 'earliest' and 'latest' parameters snap to the unit that you specify.  For example, if 'earliest' is set to -d@d, the unit is day. If the 'relativeTimeAnchor' is is set to '1994-11-05T13:15:30Z'  then 'resolvedEarliest' is snapped to '1994-11-05T00:00:00Z', which is the day. Hours, minutes, and seconds are dropped.  If no 'relativeTimeAnchor' is specified, the default value is set to the time the search job was created.
	RelativeTimeAnchor *string `json:"relativeTimeAnchor,omitempty"`
	// The timezone that relative time specifiers are based off of. Timezone only applies to relative time literals  for 'earliest' and 'latest'. If UNIX time or UTC format is used for 'earliest' and 'latest', this field is ignored. For the list of supported timezone formats, see https://docs.splunk.com/Documentation/Splunk/latest/Data/Applytimezoneoffsetstotimestamps#zoneinfo_.28TZ.29_database type: string default: \"GMT\"
	Timezone interface{} `json:"timezone,omitempty"`
}
```

Represents parameters on the search job such as 'earliest' and 'latest'.

#### type SearchJob

```go
type SearchJob struct {
	// The SPL search string.
	Query string `json:"query"`
	// Specifies whether a search that contains commands with side effects (with possible security risks) is allowed to run. type: boolean default: false
	AllowSideEffects interface{} `json:"allowSideEffects,omitempty"`
	// Specified whether a search is allowed to collect events summary during the run time.
	CollectEventSummary *bool `json:"collectEventSummary,omitempty"`
	// Specified whether a search is allowed to collect Fields summary during the run time.
	CollectFieldSummary *bool `json:"collectFieldSummary,omitempty"`
	// Specified whether a search is allowed to collect Timeline Buckets summary during the run time.
	CollectTimeBuckets *bool `json:"collectTimeBuckets,omitempty"`
	// The time, in GMT, that the search job is finished. Empty if the search job has not completed.
	CompletionTime *string `json:"completionTime,omitempty"`
	// The time, in GMT, that the search job is dispatched.
	DispatchTime *string `json:"dispatchTime,omitempty"`
	// Specifies whether the Search service should extract all of the available fields in the data, including fields not mentioned in the SPL for the search job. Set to 'false' for better search peformance.
	ExtractAllFields *bool `json:"extractAllFields,omitempty"`
	// The number of seconds to run the search before finalizing the search. The maximum value is 21600 seconds (6 hours).
	MaxTime  *float32  `json:"maxTime,omitempty"`
	Messages []Message `json:"messages,omitempty"`
	// The module to run the search in. The default module is used if a module is not specified.
	Module *string `json:"module,omitempty"`
	// The name of the created search job.
	Name *string `json:"name,omitempty"`
	// An estimate of the percent of time remaining before the job completes.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// Represents parameters on the search job such as 'earliest' and 'latest'.
	QueryParameters *QueryParameters `json:"queryParameters,omitempty"`
	// The earliest time speciifed as an absolute value in GMT. The time is computed based on the values you specify for the 'timezone' and 'earliest' queryParameters.
	ResolvedEarliest *string `json:"resolvedEarliest,omitempty"`
	// The latest time specified as an absolute value in GMT. The time is computed based on the values you specify for the 'timezone' and 'earliest' queryParameters.
	ResolvedLatest *string `json:"resolvedLatest,omitempty"`
	// The number of results produced so far for the search job.
	ResultsAvailable *int32 `json:"resultsAvailable,omitempty"`
	// The ID assigned to the search job.
	Sid    *string       `json:"sid,omitempty"`
	Status *SearchStatus `json:"status,omitempty"`
}
```

A fully-constructed search job, including read-only fields.

#### type SearchStatus

```go
type SearchStatus string
```

SearchStatus : The current status of the search job. The valid status values are
'running', 'done', 'canceled', and 'failed'.

```go
const (
	SearchStatusRunning  SearchStatus = "running"
	SearchStatusDone     SearchStatus = "done"
	SearchStatusCanceled SearchStatus = "canceled"
	SearchStatusFailed   SearchStatus = "failed"
)
```
List of SearchStatus

#### type Service

```go
type Service services.BaseService
```


#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new search service client from the given Config

#### func (*Service) CreateJob

```go
func (s *Service) CreateJob(searchJob SearchJob, resp ...*http.Response) (*SearchJob, error)
```
CreateJob - search service endpoint Creates a search job. Parameters:

    searchJob
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetJob

```go
func (s *Service) GetJob(sid string, resp ...*http.Response) (*SearchJob, error)
```
GetJob - search service endpoint Return the search job with the specified search
ID (SID). Parameters:

    sid: The search ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListEventsSummary

```go
func (s *Service) ListEventsSummary(sid string, query *ListEventsSummaryQueryParams, resp ...*http.Response) (*ListSearchResultsResponse, error)
```
ListEventsSummary - search service endpoint Return events summary, for search ID
(SID) search. Parameters:

    sid: The search ID.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListFieldsSummary

```go
func (s *Service) ListFieldsSummary(sid string, query *ListFieldsSummaryQueryParams, resp ...*http.Response) (*FieldsSummary, error)
```
ListFieldsSummary - search service endpoint Return fields stats summary of the
events to-date, for search ID (SID) search. Parameters:

    sid: The search ID.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListJobs

```go
func (s *Service) ListJobs(query *ListJobsQueryParams, resp ...*http.Response) ([]SearchJob, error)
```
ListJobs - search service endpoint Return the matching list of search jobs.
Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListResults

```go
func (s *Service) ListResults(sid string, query *ListResultsQueryParams, resp ...*http.Response) (*ListSearchResultsResponse, error)
```
ListResults - search service endpoint Return the search results for the job with
the specified search ID (SID). Parameters:

    sid: The search ID.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListTimeBuckets

```go
func (s *Service) ListTimeBuckets(sid string, resp ...*http.Response) (*TimeBucketsSummary, error)
```
ListTimeBuckets - search service endpoint Return event distribution over time of
the untransformed events read to-date, for search ID(SID) search. Parameters:

    sid: The search ID.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateJob

```go
func (s *Service) UpdateJob(sid string, updateJob UpdateJob, resp ...*http.Response) (*SearchJob, error)
```
UpdateJob - search service endpoint Update the search job with the specified
search ID (SID) with an action. Parameters:

    sid: The search ID.
    updateJob
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) WaitForJob

```go
func (s *Service) WaitForJob(jobID string, pollInterval time.Duration) (interface{}, error)
```
WaitForJob polls the job until it's completed or errors out

#### type Servicer

```go
type Servicer interface {
	// WaitForJob polls the job until it's completed or errors out
	WaitForJob(jobID string, pollInterval time.Duration) (interface{}, error)
	/*
		CreateJob - search service endpoint
		Creates a search job.
		Parameters:
			searchJob
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateJob(searchJob SearchJob, resp ...*http.Response) (*SearchJob, error)
	/*
		GetJob - search service endpoint
		Return the search job with the specified search ID (SID).
		Parameters:
			sid: The search ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetJob(sid string, resp ...*http.Response) (*SearchJob, error)
	/*
		ListEventsSummary - search service endpoint
		Return events summary, for search ID (SID) search.
		Parameters:
			sid: The search ID.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListEventsSummary(sid string, query *ListEventsSummaryQueryParams, resp ...*http.Response) (*ListSearchResultsResponse, error)
	/*
		ListFieldsSummary - search service endpoint
		Return fields stats summary of the events to-date, for search ID (SID) search.
		Parameters:
			sid: The search ID.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListFieldsSummary(sid string, query *ListFieldsSummaryQueryParams, resp ...*http.Response) (*FieldsSummary, error)
	/*
		ListJobs - search service endpoint
		Return the matching list of search jobs.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListJobs(query *ListJobsQueryParams, resp ...*http.Response) ([]SearchJob, error)
	/*
		ListResults - search service endpoint
		Return the search results for the job with the specified search ID (SID).
		Parameters:
			sid: The search ID.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListResults(sid string, query *ListResultsQueryParams, resp ...*http.Response) (*ListSearchResultsResponse, error)
	/*
		ListTimeBuckets - search service endpoint
		Return event distribution over time of the untransformed events read to-date, for search ID(SID) search.
		Parameters:
			sid: The search ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListTimeBuckets(sid string, resp ...*http.Response) (*TimeBucketsSummary, error)
	/*
		UpdateJob - search service endpoint
		Update the search job with the specified search ID (SID) with an action.
		Parameters:
			sid: The search ID.
			updateJob
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateJob(sid string, updateJob UpdateJob, resp ...*http.Response) (*SearchJob, error)
}
```

Servicer represents the interface for implementing all endpoints for this
service

#### type SingleFieldSummary

```go
type SingleFieldSummary struct {
	// The total number of events that contain the field.
	Count *int32 `json:"count,omitempty"`
	// The total number of unique values in the field.
	DistictCount *int32 `json:"distictCount,omitempty"`
	// Specifies if the distinctCount is accurate. The isExact property is FALSE when the distinctCount exceeds the maximum count and an exact count is not available.
	IsExact *bool `json:"isExact,omitempty"`
	// The maximum numeric values in the field.
	Max *string `json:"max,omitempty"`
	// The mean (average) for the numeric values in the field.
	Mean *float32 `json:"mean,omitempty"`
	// The minimum numeric values in the field.
	Min *string `json:"min,omitempty"`
	// An array of the values in the field.
	Modes []SingleValueMode `json:"modes,omitempty"`
	// The count of the numeric values in the field.
	NumericCount *int32 `json:"numericCount,omitempty"`
	// Specifies if the field was added or changed by the search.
	Relevant *bool `json:"relevant,omitempty"`
	// The standard deviation for the numeric values in the field.
	Stddev *float32 `json:"stddev,omitempty"`
}
```

Summary of each field.

#### type SingleTimeBucket

```go
type SingleTimeBucket struct {
	// Count of available events. Not all events in a bucket are retrievable. Typically this count is capped at 10000.
	AvailableCount *int32   `json:"availableCount,omitempty"`
	Duration       *float32 `json:"duration,omitempty"`
	// The timestamp of the earliest event in the current bucket, in UNIX format. This is the same time as 'earliestTimeStrfTime' in UNIX format.
	EarliestTime *float32 `json:"earliestTime,omitempty"`
	// The timestamp of the earliest event in the current bucket, in UTC format with seconds. For example 2019-01-25T13:15:30Z, which follows the ISO-8601 (%FT%T.%Q) format.
	EarliestTimeStrfTime *string `json:"earliestTimeStrfTime,omitempty"`
	// Specifies if all of the events in the current bucket have been finalized.
	IsFinalized *bool `json:"isFinalized,omitempty"`
	// The total count of the events in the current bucket.
	TotalCount *int32 `json:"totalCount,omitempty"`
}
```

Events summary in single time bucket.

#### type SingleValueMode

```go
type SingleValueMode struct {
	// The number of occurences that the value appears in a field.
	Count *int32 `json:"count,omitempty"`
	// Specifies if the count is accurate. The isExact property is FALSE when the count exceeds the maximum count and an exact count is not available.
	IsExact *bool `json:"isExact,omitempty"`
	// The value in the field.
	Value *string `json:"value,omitempty"`
}
```

Single value summary of the field.

#### type TimeBucketsSummary

```go
type TimeBucketsSummary struct {
	// Specifies if the events are returned in time order.
	IsTimeCursored *bool              `json:"IsTimeCursored,omitempty"`
	Buckets        []SingleTimeBucket `json:"buckets,omitempty"`
	// Identifies where the cursor is in processing the events. The cursorTime is a timestamp specified in UNIX time.
	CursorTime *float32 `json:"cursorTime,omitempty"`
	// The number of events processed at the cursorTime.
	EventCount *int32 `json:"eventCount,omitempty"`
}
```

A timeline metadata model of the event distribution. The model shows the
untransformed events that are read to date for a specific for search ID (sid).

#### type UpdateJob

```go
type UpdateJob struct {
	// The status to PATCH to an existing search job. The only status values you can PATCH are 'canceled' and 'finalized'. You can PATCH the 'canceled' status only to a search job that is running.
	Status UpdateJobStatus `json:"status"`
}
```

Update a search job with a status.

#### type UpdateJobStatus

```go
type UpdateJobStatus string
```

UpdateJobStatus : The status to PATCH to an existing search job. The only status
values you can PATCH are 'canceled' and 'finalized'. You can PATCH the
'canceled' status only to a search job that is running.

```go
const (
	UpdateJobStatusCanceled  UpdateJobStatus = "canceled"
	UpdateJobStatusFinalized UpdateJobStatus = "finalized"
)
```
List of UpdateJobStatus
