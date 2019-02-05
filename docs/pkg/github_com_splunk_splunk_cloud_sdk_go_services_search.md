# search
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/search"


## Usage

#### type CreateJobRequest

```go
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
	// Represents parameters on the search job such as 'earliest' and 'latest'.
	QueryParameters *QueryParameters `json:"queryParameters,omitempty"`
}
```

CreateJobRequest defines properties allowed (and possibly required) in fully
constructed Searchjobs in POST payloads and responses

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
func (i *Iterator) Value() (*Results, error)
```
Value returns value in current iteration or error out if iterator is closed

#### type Job

```go
type Job struct {
	// The SPL query string.
	Query string `json:"query"`
	// Determine whether the Search service extracts all available fields in the data, including fields not mentioned in the SPL for the search job.
	// Set to 'false' for better search performance.
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
	QueryParameters QueryParameters `json:"queryParameters,omitempty"`
	// The ID assigned to the search job.
	ID string `json:"sid,omitempty"`
	// The current status of the search job.
	Status JobStatus `json:"status,omitempty"`
	// An estimate of how close the job is to completing.
	PercentComplete float64 `json:"percentComplete,omitempty"`
	// The number of results produced so far for the search job.
	ResultsAvailable int64 `json:"resultsAvailable,omitempty"`
	// Run time messages from Splunkd.
	Messages JobMessages `json:"messages,omitempty"`
}
```

Job represents a fully-constructed search job, including read-only fields.

#### type JobAction

```go
type JobAction string
```

JobAction defines actions to be taken on an existing search job.

```go
const (
	JobCanceled  JobAction = "canceled"
	JobFinalized JobAction = "finalized"
)
```
Define supported job actions

#### type JobMessageType

```go
type JobMessageType string
```

JobMessageType defines type of messages from Splunkd

```go
const (
	InfoType  JobMessageType = "INFO"
	FatalType JobMessageType = "FATAL"
	ErrorType JobMessageType = "ERROR"
	DebugType JobMessageType = "DEBUG"
)
```
Define supported message type

#### type JobMessages

```go
type JobMessages []struct {
	// Enum [INFO, FATAL, ERROR, DEBUG]
	Type string `json:"type"`
	// message text
	Text string `json:"text"`
}
```

JobMessages is used in search results or search job.

#### type JobResultsParams

```go
type JobResultsParams struct {
	Count  int `key:"count"`
	Offset int `key:"offset"`
}
```

JobResultsParams specifies the query params when fetching job results

#### type JobStatus

```go
type JobStatus string
```

JobStatus describes status of a search job

```go
const (
	Queued     JobStatus = "queued"
	Parsing    JobStatus = "parsing"
	Running    JobStatus = "running"
	Finalizing JobStatus = "finalizing"
	Failed     JobStatus = "failed"
	Done       JobStatus = "done"
)
```
Supported JobStatus constants

#### type JobsQuery

```go
type JobsQuery struct {
	//The supported statuses are running, done and failed
	Status string `key:"status"`
}
```

JobsQuery represents Query Parameters that can be provided for ListJobs endpoint

#### type PatchJobResponse

```go
type PatchJobResponse struct {
	// Run time messages from Splunkd.
	Messages JobMessages `json:"messages"`
}
```

PatchJobResponse defines the response from patch endpoint

#### type QueryFunc

```go
type QueryFunc func(step, start int) (*Results, error)
```

QueryFunc is the function to be executed in each Next call of the iterator

#### type QueryParameters

```go
type QueryParameters struct {
	// The earliest time in absolute or relative format to retrieve events (only supported if the query supports time-ranges)
	Earliest string `json:"earliest,omitempty"`
	// The latest time in absolute or relative format to retrieve events (only supported if the query supports time-ranges)
	Latest string `json:"latest,omitempty"`
}
```

QueryParameters is the type representing parameters currently earliest & latest
on search.

#### type Results

```go
type Results struct {
	// Run time messages from Splunkd.
	Messages JobMessages              `json:"messages"`
	Results  []map[string]interface{} `json:"results"`
	Fields   []map[string]interface{} `json:"fields"`
}
```

Results represents results from a search job

#### type ResultsNotReadyResponse

```go
type ResultsNotReadyResponse struct {
	// URL for job results
	NextLink string `json:"nextLink,omitempty"`
	// Number of milliseconds to wait before retrying
	Wait string `json:"wait,omitempty"`
}
```

ResultsNotReadyResponse represents the response when no search results is ready

#### type Service

```go
type Service services.BaseService
```

Service talks to the Splunk Cloud search service

#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new search service client from the given Config

#### func (*Service) CreateJob

```go
func (s *Service) CreateJob(job *CreateJobRequest) (*Job, error)
```
CreateJob creates a new search job

#### func (*Service) GetJob

```go
func (s *Service) GetJob(jobID string) (*Job, error)
```
GetJob retrieves information about the specified search.

#### func (*Service) GetResults

```go
func (s *Service) GetResults(jobID string, count, offset int) (interface{}, error)
```
GetResults Returns the job results with the given `id`. count=0 returns default
number of results from search

#### func (*Service) ListJobs

```go
func (s *Service) ListJobs() ([]Job, error)
```
ListJobs gets the matching list of search jobs

#### func (*Service) ListJobsByQueryParameters

```go
func (s *Service) ListJobsByQueryParameters(query JobsQuery) ([]Job, error)
```
ListJobsByQueryParameters gets the matching list of search jobs filtered by
query parameters specified

#### func (*Service) UpdateJob

```go
func (s *Service) UpdateJob(jobID string, jobStatus JobAction) (*PatchJobResponse, error)
```
UpdateJob updates an existing job with actions and TTL

#### func (*Service) WaitForJob

```go
func (s *Service) WaitForJob(jobID string, pollInterval time.Duration) (interface{}, error)
```
WaitForJob polls the job until it's completed or errors out

#### type Servicer

```go
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
```

Servicer ...
