# ingest
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/ingest"


## Usage

#### type BatchEventsSender

```go
type BatchEventsSender struct {
	PayLoadBytes int
	BatchSize    int
	EventsChan   chan Event
	EventsQueue  []Event
	QuitChan     chan struct{}
	EventService *Service
	IngestTicker *util.Ticker
	WaitGroup    *sync.WaitGroup
	ErrorChan    chan struct{}
	IsRunning    bool

	Errors []ingestError
}
```

BatchEventsSender sends events in batches or periodically if batch is not full
to Splunk Cloud ingest service endpoints

#### func (*BatchEventsSender) AddEvent

```go
func (b *BatchEventsSender) AddEvent(event Event) error
```
AddEvent pushes a single event into EventsChan

#### func (*BatchEventsSender) ResetQueue

```go
func (b *BatchEventsSender) ResetQueue()
```
ResetQueue sets b.EventsQueue to empty, but keep memory allocated for underlying
array

#### func (*BatchEventsSender) Restart

```go
func (b *BatchEventsSender) Restart()
```
Restart will reset batch event sender, clear up queue, error msg, timer etc.

#### func (*BatchEventsSender) Run

```go
func (b *BatchEventsSender) Run()
```
Run sets up ticker and starts a new goroutine

#### func (*BatchEventsSender) SetCallbackHandler

```go
func (b *BatchEventsSender) SetCallbackHandler(callback UserErrHandler)
```
SetCallbackHandler allows users to pass their own callback function

#### func (*BatchEventsSender) Stop

```go
func (b *BatchEventsSender) Stop()
```
Stop sends a signal to QuitChan, wait for all registered goroutines to finish,
stop ticker and clear queue

#### type Error

```go
type Error struct {
	Code    *string                `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Message *string                `json:"message,omitempty"`
}
```


#### type Event

```go
type Event struct {
	// JSON object for the event.
	Body interface{} `json:"body"`
	// Specifies a JSON object that contains explicit custom fields to be defined at index time.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// The host value assigned to the event data. Typically, this is the hostname of the client from which you are sending data.
	Host *string `json:"host,omitempty"`
	// An optional ID that uniquely identifies the event data. It is used to deduplicate the data if same data is set multiple times. If ID is not specified, it will be assigned by the system.
	Id *string `json:"id,omitempty"`
	// Optional nanoseconds part of the timestamp.
	Nanos *int32 `json:"nanos,omitempty"`
	// The source value to assign to the event data. For example, if you are sending data from an app that you are developing, set this key to the name of the app.
	Source *string `json:"source,omitempty"`
	// The sourcetype value assigned to the event data.
	Sourcetype *string `json:"sourcetype,omitempty"`
	// Epoch time in milliseconds.
	Timestamp *int64 `json:"timestamp,omitempty"`
}
```


#### type HttpResponse

```go
type HttpResponse struct {
	Code    *string                `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Message *string                `json:"message,omitempty"`
}
```


#### type Metric

```go
type Metric struct {
	// Name of the metric e.g. CPU, Memory etc.
	Name string `json:"name"`
	// Dimensions allow metrics to be classified e.g. {\"Server\":\"nginx\", \"Region\":\"us-west-1\", ...}
	Dimensions map[string]string `json:"dimensions,omitempty"`
	// Type of metric. Default is g for gauge.
	Type *string `json:"type,omitempty"`
	// Unit of the metric e.g. percent, megabytes, seconds etc.
	Unit *string `json:"unit,omitempty"`
	// Value of the metric. If not specified, it will be defaulted to 0.
	Value *float64 `json:"value,omitempty"`
}
```


#### type MetricAttribute

```go
type MetricAttribute struct {
	// Optional. If set, individual metrics inherit these dimensions and can override any and/or all of them.
	DefaultDimensions map[string]string `json:"defaultDimensions,omitempty"`
	// Optional. If set, individual metrics inherit this type and can optionally override.
	DefaultType *string `json:"defaultType,omitempty"`
	// Optional. If set, individual metrics inherit this unit and can optionally override.
	DefaultUnit *string `json:"defaultUnit,omitempty"`
}
```


#### type MetricEvent

```go
type MetricEvent struct {
	// Specifies multiple related metrics e.g. Memory, CPU etc.
	Body       []Metric         `json:"body"`
	Attributes *MetricAttribute `json:"attributes,omitempty"`
	// The host value assigned to the event data. Typically, this is the hostname of the client from which you are sending data.
	Host *string `json:"host,omitempty"`
	// An optional ID that uniquely identifies the metric data. It is used to deduplicate the data if same data is set multiple times. If ID is not specified, it will be assigned by the system.
	Id *string `json:"id,omitempty"`
	// Optional nanoseconds part of the timestamp.
	Nanos *int32 `json:"nanos,omitempty"`
	// The source value to assign to the event data. For example, if you are sending data from an app that you are developing, set this key to the name of the app.
	Source *string `json:"source,omitempty"`
	// The sourcetype value assigned to the event data.
	Sourcetype *string `json:"sourcetype,omitempty"`
	// Epoch time in milliseconds.
	Timestamp *int64 `json:"timestamp,omitempty"`
}
```


#### type Service

```go
type Service services.BaseService
```


#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new ingest service client from the given Config

#### func (*Service) NewBatchEventsSender

```go
func (s *Service) NewBatchEventsSender(batchSize int, interval int64, payLoadSize int) (*BatchEventsSender, error)
```
NewBatchEventsSender used to initialize dependencies and set values

#### func (*Service) NewBatchEventsSenderWithMaxAllowedError

```go
func (s *Service) NewBatchEventsSenderWithMaxAllowedError(batchSize int, interval int64, dataSize int, maxErrorsAllowed int) (*BatchEventsSender, error)
```

#### func (*Service) PostEvents

```go
func (s *Service) PostEvents(event []Event, resp ...*http.Response) (*HttpResponse, error)
```
PostEvents - Sends events. Parameters:

    event
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) PostMetrics

```go
func (s *Service) PostMetrics(metricEvent []MetricEvent, resp ...*http.Response) (*HttpResponse, error)
```
PostMetrics - Sends metric events. Parameters:

    metricEvent
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### type Servicer

```go
type Servicer interface {
	NewBatchEventsSenderWithMaxAllowedError(batchSize int, interval int64, dataSize int, maxErrorsAllowed int) (*BatchEventsSender, error)
	// NewBatchEventsSender used to initialize dependencies and set values
	NewBatchEventsSender(batchSize int, interval int64, payLoadSize int) (*BatchEventsSender, error)
	/*
		PostEvents - Sends events.
		Parameters:
			event
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PostEvents(event []Event, resp ...*http.Response) (*HttpResponse, error)
	/*
		PostMetrics - Sends metric events.
		Parameters:
			metricEvent
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PostMetrics(metricEvent []MetricEvent, resp ...*http.Response) (*HttpResponse, error)
}
```

Servicer represents the interface for implementing all endpoints for this
service

#### type UserErrHandler

```go
type UserErrHandler func(*BatchEventsSender)
```

UserErrHandler defines the type of user callback function for batchEventSender
