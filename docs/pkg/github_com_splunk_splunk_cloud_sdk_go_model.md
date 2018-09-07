# model
--
    import "github.com/splunk/splunk-cloud-sdk-go/model"


## Usage

#### type Action

```go
type Action struct {
	// Common action fields:
	// Name of action, all actions have this field
	Name string `json:"name" binding:"required"`
	// Kind of action (email, webhook, or sns), all actions have this field
	Kind ActionKind `json:"kind" binding:"required"`
	ActionUpdateFields
}
```

Action defines the fields for email, sns, and webhooks as one aggregated model

#### func  NewEmailAction

```go
func NewEmailAction(name string, htmlPart string, subjectPart string, textPart string, templateName string, addresses []string) *Action
```
NewEmailAction creates a new email kind action

#### func  NewSNSAction

```go
func NewSNSAction(name string, topic string, message string) *Action
```
NewSNSAction creates a new sns kind action

#### func  NewWebhookAction

```go
func NewWebhookAction(name string, webhookURL string, message string) *Action
```
NewWebhookAction creates a new webhook kind action

#### type ActionError

```go
type ActionError struct {
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Details  interface{} `json:"details,omitempty"`
	MoreInfo string      `json:"moreInfo,omitempty"`
}
```

ActionError defines format for returned errors

#### type ActionKind

```go
type ActionKind string
```

ActionKind reflects the kinds of actions supported by the Action service

```go
const (
	// EmailKind for email actions
	EmailKind ActionKind = "email"
	// WebhookKind for webhook actions
	WebhookKind ActionKind = "webhook"
	// SNSKind for SNS actions
	SNSKind ActionKind = "sns"
)
```

#### type ActionNotification

```go
type ActionNotification struct {
	Kind    ActionNotificationKind `json:"kind" binding:"required"`
	Tenant  string                 `json:"tenant" binding:"required"`
	Payload ActionPayload          `json:"payload" binding:"required"`
}
```

ActionNotification defines the action notification format

#### type ActionNotificationKind

```go
type ActionNotificationKind string
```

ActionNotificationKind defines the types of notifications

```go
const (
	//SplunkEventKind for splunk event payloads
	SplunkEventKind ActionNotificationKind = "splunkEvent"
	//RawJSONPayloadKind for raw json payloads
	RawJSONPayloadKind ActionNotificationKind = "rawJSON"
)
```

#### type ActionPayload

```go
type ActionPayload interface{}
```

ActionPayload is what is sent when the action is triggered

#### type ActionStatus

```go
type ActionStatus struct {
	State    ActionStatusState `json:"state"`
	StatusID string            `json:"statusId"`
	Message  string            `json:"message,omitempty"`
}
```

ActionStatus defines the state information

#### type ActionStatusState

```go
type ActionStatusState string
```

ActionStatusState reflects the status of the action

```go
const (
	// StatusQueued status
	StatusQueued ActionStatusState = "QUEUED"
	// StatusRunning status
	StatusRunning ActionStatusState = "RUNNING"
	// StatusDone status
	StatusDone ActionStatusState = "DONE"
	// StatusFailed status
	StatusFailed ActionStatusState = "FAILED"
)
```

#### type ActionTriggerResponse

```go
type ActionTriggerResponse struct {
	StatusID  *string
	StatusURL *url.URL
}
```

ActionTriggerResponse for returning status url and parsed statusID (if possible)

#### type ActionUpdateFields

```go
type ActionUpdateFields struct {
	// Email action fields:
	// HTMLPart to send via Email action
	HTMLPart string `json:"htmlPart,omitempty"`
	// SubjectPart to send via Email action
	SubjectPart string `json:"subjectPart,omitempty"`
	// TextPart to send via Email action
	TextPart string `json:"textPart,omitempty"`
	// TemplateName to send via Email action
	TemplateName string `json:"templateName,omitempty"`
	// Addresses to send to when Email action triggered
	Addresses []string `json:"addresses" binding:"required"`

	// SNS action fields:
	// Topic to trigger SNS action
	Topic string `json:"topic" binding:"required"`
	// Message to send via SNS or Webhook action
	Message string `json:"message" binding:"required"`

	// Webhook action fields:
	// WebhookURL to trigger Webhook action
	WebhookURL string `json:"webhookUrl" binding:"required"`
}
```

ActionUpdateFields defines the fields that may be updated for an existing Action

#### type AuthError

```go
type AuthError struct {

	// The reason of the auth error
	Reason string `json:"reason"`
}
```

AuthError auth error reason

#### type AutoMode

```go
type AutoMode string
```

AutoMode enumerates the automatic key/value extraction modes. One of "NONE",
"AUTO", "MULTIKV", "XML", "JSON".

```go
const (
	// NONE Automode
	NONE AutoMode = "NONE"
	// AUTO Automode
	AUTO AutoMode = "AUTO"
	// MULTIKV Automode
	MULTIKV AutoMode = "MULTIKV"
	// XML Automode
	XML AutoMode = "XML"
	// JSON Automode
	JSON AutoMode = "JSON"
)
```

#### type CatalogAction

```go
type CatalogAction struct {
	ID         string            `json:"id,omitempty"`
	RuleID     string            `json:"ruleid,omitempty"`
	Kind       CatalogActionKind `json:"kind" binding:"required"`
	Owner      string            `json:"owner" binding:"required"`
	Created    string            `json:"created,omitempty"`
	Modified   string            `json:"modified,omitempty"`
	CreatedBy  string            `json:"createdBy,omitempty"`
	ModifiedBy string            `json:"modifiedBy,omitempty"`
	Version    int               `json:"version,omitempty"`
	Field      string            `json:"field,omitempty"`
	Alias      string            `json:"alias,omitempty"`
	Mode       AutoMode          `json:"mode,omitempty"`
	Expression string            `json:"expression,omitempty"`
	Pattern    string            `json:"pattern,omitempty"`
	Limit      int               `json:"limit,omitempty"`
}
```

CatalogAction represents a specific search time transformation action.

#### type CatalogActionKind

```go
type CatalogActionKind string
```

CatalogActionKind enumerates the kinds of search time transformation action
known by the service.

```go
const (
	// ALIAS action
	ALIAS CatalogActionKind = "ALIAS"
	// AUTOKV action
	AUTOKV CatalogActionKind = "AUTOKV"
	// REGEX action
	REGEX CatalogActionKind = "REGEX"
	// EVAL action
	EVAL CatalogActionKind = "EVAL"
	// LOOKUPACTION action
	LOOKUPACTION CatalogActionKind = "LOOKUP"
)
```

#### type CollectionDefinition

```go
type CollectionDefinition struct {

	// The collection name
	// Max Length: 45
	// Min Length: 1
	Collection string `json:"collection"`
}
```

CollectionDefinition collection definition

#### type CollectionStats

```go
type CollectionStats struct {

	// Number of records in collection
	Count int64 `json:"count"`

	// Map of index name to index size in bytes
	IndexSizes interface{} `json:"indexSizes"`

	// Number of indexes on collection
	Nindexes int64 `json:"nindexes"`

	// Collection name
	Ns string `json:"ns"`

	// Size in bytes of collection, not including indexes
	Size int64 `json:"size"`

	// Total size of indexes
	TotalIndexSize int64 `json:"totalIndexSize"`
}
```

CollectionStats collection stats

#### type DataType

```go
type DataType string
```

DataType enumerates the kinds of datatypes used in fields.

```go
const (
	// DATE DataType
	DATE DataType = "DATE"
	// NUMBER DataType
	NUMBER DataType = "NUMBER"
	// OBJECTID DataType
	OBJECTID DataType = "OBJECT_ID"
	// STRING DataType
	STRING DataType = "STRING"
	// DATATYPEUNKNOWN DataType
	DATATYPEUNKNOWN DataType = "UNKNOWN"
)
```

#### type DatasetInfo

```go
type DatasetInfo struct {
	ID           string          `json:"id,omitempty"`
	Name         string          `json:"name"`
	Kind         DatasetInfoKind `json:"kind"`
	Owner        string          `json:"owner"`
	Module       string          `json:"module,omitempty"`
	Created      string          `json:"created,omitempty"`
	Modified     string          `json:"modified,omitempty"`
	CreatedBy    string          `json:"createdBy,omitempty"`
	ModifiedBy   string          `json:"modifiedBy,omitempty"`
	Capabilities string          `json:"capabilities"`
	Version      int             `json:"version,omitempty"`
	Fields       []Field         `json:"fields,omitempty"`
	Readroles    []string        `json:"readroles,omitempty"`
	Writeroles   []string        `json:"writeroles,omitempty"`

	ExternalKind       string `json:"externalKind,omitempty"`
	ExternalName       string `json:"externalName,omitempty"`
	CaseSensitiveMatch bool   `json:"caseSensitiveMatch,omitempty"`
	Filter             string `json:"filter,omitempty"`
	MaxMatches         int    `json:"maxMatches,omitempty"`
	MinMatches         int    `json:"minMatches,omitempty"`
	DefaultMatch       string `json:"defaultMatch,omitempty"`

	Datatype string `json:"datatype,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}
```

DatasetInfo represents the sources of data that can be serched by Splunk

#### type DatasetInfoKind

```go
type DatasetInfoKind string
```

DatasetInfoKind enumerates the kinds of datasets known to the system.

```go
const (
	// LOOKUP represents TODO: Description needed
	LOOKUP DatasetInfoKind = "lookup"
	// KVCOLLECTION represents a key value store, it is used with the kvstore service, but its implementation is separate of kvstore
	KVCOLLECTION DatasetInfoKind = "kvcollection"
	// INDEX represents a Splunk events or metrics index
	INDEX DatasetInfoKind = "index"
)
```

#### type DispatchState

```go
type DispatchState string
```

DispatchState describes dispatchState of a job

```go
const (
	QUEUED     DispatchState = "QUEUED"
	PARSING    DispatchState = "PARSING"
	RUNNING    DispatchState = "RUNNING"
	FINALIZING DispatchState = "FINALIZING"
	FAILED     DispatchState = "FAILED"
	DONE       DispatchState = "DONE"
)
```
Supported DispatchState constants

#### type Error

```go
type Error struct {

	// The reason of the error
	Code int64 `json:"code"`
	// Error message
	Message string `json:"message"`
	// State Storage error code
	SsCode int64 `json:"ssCode"`
}
```

Error error reason

#### type Event

```go
type Event struct {
	Host       string            `json:"host,omitempty" key:"host"`
	Index      string            `json:"index,omitempty" key:"index"`
	Sourcetype string            `json:"sourcetype,omitempty" key:"sourcetype"`
	Source     string            `json:"source,omitempty" key:"source"`
	Time       *float64          `json:"time,omitempty" key:"time"`
	Event      interface{}       `json:"event"`
	Fields     map[string]string `json:"fields,omitempty"`
}
```

Event contains metadata about the event

#### type ExportCollectionContentType

```go
type ExportCollectionContentType string
```

ExportCollectionContentType used to specify the export collection file content
type

```go
const (
	// CSV captures enum value "csv"
	CSV ExportCollectionContentType = "csv"

	// GZIP captures enum value "gzip"
	GZIP ExportCollectionContentType = "gzip"
)
```

#### type FetchEventsRequest

```go
type FetchEventsRequest struct {
	Count            int      `key:"count"`
	Offset           int      `key:"offset"`
	EarliestTime     string   `key:"earliestTime"`
	Fields           []string `key:"f"`
	LatestTime       string   `key:"latestTime"`
	MaxLines         *uint    `key:"maxLines"`
	TimeFormat       string   `key:"timeFormat"`
	OutputTimeFormat string   `key:"outputTimeFormat"`
	Search           string   `key:"search"`
	TruncationMode   string   `key:"truncationMode"`
	Segmentation     string   `key:"segmentation"`
}
```

FetchEventsRequest specifies the query params when fetching job events

#### type FetchResultsRequest

```go
type FetchResultsRequest struct {
	Count  int      `key:"count"`
	Offset int      `key:"offset"`
	Fields []string `key:"f"`
	Search string   `key:"search"`
}
```

FetchResultsRequest specifies the query params when fetching job results

#### type Field

```go
type Field struct {
	ID         string         `json:"id,omitempty"`
	Name       string         `json:"name,omitempty"`
	DatasetID  string         `json:"datasetid,omitempty"`
	DataType   DataType       `json:"datatype,omitempty"`
	FieldType  FieldType      `json:"fieldtype,omitempty"`
	Prevalence PrevalenceType `json:"prevalence,omitempty"`
	Created    string         `json:"created,omitempty"`
	Modified   string         `json:"modified,omitempty"`
}
```

Field represents the fields belonging to the specified Dataset

#### type FieldType

```go
type FieldType string
```

FieldType enumerates different kinds of fields.

```go
const (
	// DIMENSION fieldType
	DIMENSION FieldType = "DIMENSION"
	// MEASURE fieldType
	MEASURE FieldType = "MEASURE"
	// FIELDTYPEUNKNOWN fieldType
	FIELDTYPEUNKNOWN FieldType = "UNKNOWN"
)
```

#### type Group

```go
type Group struct {

	// created at
	// Required: true
	CreatedAt strfmt.DateTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy string `json:"createdBy"`

	// name
	// Required: true
	Name string `json:"name"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

Group group

#### type GroupMember

```go
type GroupMember struct {

	// added at
	// Required: true
	AddedAt strfmt.DateTime `json:"addedAt"`

	// added by
	// Required: true
	AddedBy string `json:"addedBy"`

	// group
	// Required: true
	Group string `json:"group"`

	// principal
	// Required: true
	Principal string `json:"principal"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

GroupMember Represents a member that belongs to a group

#### type GroupRole

```go
type GroupRole struct {

	// added at
	// Required: true
	AddedAt strfmt.DateTime `json:"addedAt"`

	// added by
	// Required: true
	AddedBy string `json:"addedBy"`

	// group
	// Required: true
	Group string `json:"group"`

	// role
	// Required: true
	Role string `json:"role"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

GroupRole Represents a role that is assigned to a group

#### type IndexDefinition

```go
type IndexDefinition struct {

	// The name of the index
	Name string `json:"name,omitempty"`

	// fields
	Fields []IndexFieldDefinition `json:"fields"`
}
```

IndexDefinition index field definition

#### type IndexDescription

```go
type IndexDescription struct {

	// The collection name
	Collection string `json:"collection,omitempty"`

	// fields
	Fields []IndexFieldDefinition `json:"fields"`

	// The name of the index
	Name string `json:"name,omitempty"`
}
```

IndexDescription index description

#### type IndexFieldDefinition

```go
type IndexFieldDefinition struct {

	// The sort direction for the indexed field
	Direction int64 `json:"direction"`

	// The name of the field to index
	Field string `json:"field"`
}
```

IndexFieldDefinition index field definition

#### type JobAction

```go
type JobAction string
```

JobAction defines the action that can be posted on a job

```go
const (
	FINALIZE       JobAction = "finalize"
	CANCEL         JobAction = "cancel"
	TOUCH          JobAction = "touch"
	SAVE           JobAction = "save"
	SETTTL         JobAction = "setttl"
	ENABLEPREVIEW  JobAction = "enablepreview"
	DISABLEPREVIEW JobAction = "disablepreview"
)
```
Supported JobAction constants

#### type JobControlAction

```go
type JobControlAction struct {
	Action JobAction `json:"action"`
	TTL    int       `json:"ttl"`
}
```

JobControlAction specifies the action needs to be taken on a job

#### type JobControlReplyMsg

```go
type JobControlReplyMsg struct {
	Msg []struct {
		TypeKey string `json:"type"`
		Text    string `json:"text"`
	} `json:"messages"`
}
```

JobControlReplyMsg displays messages returned from taking a job control

#### type JobsRequest

```go
type JobsRequest struct {
	Count  uint `key:"count"`
	Offset uint `key:"offset"`
}
```

JobsRequest specifies pagination parameters for certain supported requests

#### func  NewDefaultPaginationParams

```go
func NewDefaultPaginationParams() *JobsRequest
```
NewDefaultPaginationParams creates search pagination parameters according to
Splunk Enterprise defaults

#### type Key

```go
type Key struct {
	Key string `json:"_key"`
}
```

Key to identify a record in a collection

#### type LookupValue

```go
type LookupValue []interface{}
```

LookupValue Value tuple used for lookup

#### type Member

```go
type Member struct {

	// When the principal was added to the tenant.
	// Required: true
	AddedAt strfmt.DateTime `json:"addedAt"`

	// added by
	// Required: true
	AddedBy string `json:"addedBy"`

	// name
	// Required: true
	Name string `json:"name"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

Member Represents a member that belongs to a tenant.

#### type Metric

```go
type Metric struct {
	// Name of the metric e.g. CPU, Memory etc.
	Name string `json:"name"`
	// Value of the metric.
	Value float64 `json:"value"`
	// Dimensions allow metrics to be classified e.g. {"Server":"nginx", "Region":"us-west-1", ...}
	Dimensions map[string]string `json:"dimensions"`
	// Type of metric. Default is g for gauge.
	Type string `json:"type"`
	// Unit of the metric e.g. percent, megabytes, seconds etc.
	Unit string `json:"unit"`
}
```

Metric defines individual metric data.

#### type MetricAttribute

```go
type MetricAttribute struct {
	// Optional. If set, individual Metrics will inherit these dimensions and can override any/all of them.
	DefaultDimensions map[string]string `json:"defaultDimensions"`
	// Optional. If set, individual Metrics will inherit this type and can optionally override.
	DefaultType string `json:"defaultType"`
	// Optional. If set, individual Metrics will inherit this unit and can optionally override.
	DefaultUnit string `json:"defaultUnit"`
}
```

MetricAttribute defines default attributes for the metric.

#### type MetricEvent

```go
type MetricEvent struct {
	// Specify multiple related metrics e.g. Memory, CPU etc.
	Body []Metric `json:"body"`
	// Epoch time in milliseconds.
	Timestamp int64 `json:"timestamp"`
	// Optional nanoseconds part of the timestamp.
	Nanos int32 `json:"nanos"`
	// The source value to assign to the event data. For example, if you're sending data from an app you're developing,
	// you could set this key to the name of the app.
	Source string `json:"source"`
	// The sourcetype value to assign to the event data.
	Sourcetype string `json:"sourcetype"`
	// The host value to assign to the event data. This is typically the hostname of the client from which you're sending data.
	Host string `json:"host"`
	// Optional ID uniquely identifies the metric data. It is used to deduplicate the data if same data is set multiple times.
	// If ID is not specified, it will be assigned by the system.
	ID string `json:"id"`
	// Default attributes for the metric data.
	Attributes MetricAttribute `json:"attributes"`
}
```

MetricEvent define event send to metric endpoint

#### type PagingInfo

```go
type PagingInfo struct {
	Total   float64 `json:"total"`
	PerPage float64 `json:"perPage"`
	Offset  float64 `json:"offset"`
}
```

PagingInfo captures fields returned for endpoints supporting paging

#### type PartialDatasetInfo

```go
type PartialDatasetInfo struct {
	Name         string          `json:"name,omitempty"`
	Kind         DatasetInfoKind `json:"kind,omitempty"`
	Owner        string          `json:"owner,omitempty"`
	Created      string          `json:"created,omitempty"`
	Modified     string          `json:"modified,omitempty"`
	CreatedBy    string          `json:"createdBy,omitempty"`
	ModifiedBy   string          `json:"modifiedBy,omitempty"`
	Capabilities string          `json:"capabilities,omitempty"`
	Version      int             `json:"version,omitempty"`
	Readroles    []string        `json:"readroles,omitempty"`
	Writeroles   []string        `json:"writeroles,omitempty"`

	ExternalKind       string `json:"externalKind,omitempty"`
	ExternalName       string `json:"externalName,omitempty"`
	CaseSensitiveMatch bool   `json:"caseSensitiveMatch,omitempty"`
	Filter             string `json:"filter,omitempty"`
	MaxMatches         int    `json:"maxMatches,omitempty"`
	MinMatches         int    `json:"minMatches,omitempty"`
	DefaultMatch       string `json:"defaultMatch,omitempty"`

	Datatype string `json:"datatype,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}
```

PartialDatasetInfo represents the sources of data that can be updated by Splunk,
same structure as DatasetInfo

#### type PingOKBody

```go
type PingOKBody struct {

	// If database is not healthy, detailed error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Database status
	// Enum: [healthy unhealthy unknown]
	Status PingOKBodyStatus `json:"status"`
}
```

PingOKBody ping ok body

#### type PingOKBodyStatus

```go
type PingOKBodyStatus string
```

PingOKBodyStatus used to force type expectation for KVStore Ping endpoint
response

```go
const (
	// PingOKBodyStatusHealthy captures enum value "healthy"
	PingOKBodyStatusHealthy PingOKBodyStatus = "healthy"

	// PingOKBodyStatusUnhealthy captures enum value "unhealthy"
	PingOKBodyStatusUnhealthy PingOKBodyStatus = "unhealthy"

	// PingOKBodyStatusUnknown captures enum value "unknown"
	PingOKBodyStatusUnknown PingOKBodyStatus = "unknown"
)
```

#### type PostJobsRequest

```go
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
```

PostJobsRequest represents the search job post params

#### type PrevalenceType

```go
type PrevalenceType string
```

PrevalenceType enumerates the types of prevalance used in fields.

```go
const (
	// ALL PrevalenceType
	ALL PrevalenceType = "ALL"
	// SOME PrevalenceType
	SOME PrevalenceType = "SOME"
	// PREVALANCEUNKNOWN PrevalenceType
	PREVALANCEUNKNOWN PrevalenceType = "UNKNOWN"
)
```

#### type Principal

```go
type Principal struct {

	// created at
	// Required: true
	CreatedAt strfmt.DateTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy string `json:"createdBy"`

	// kind
	// Required: true
	Kind string `json:"kind"`

	// name
	// Required: true
	Name string `json:"name"`

	// profile
	Profile interface{} `json:"profile,omitempty"`

	// tenants
	// Required: true
	Tenants []string `json:"tenants"`
}
```

Principal principal

#### type RawJSONPayload

```go
type RawJSONPayload map[string]interface{}
```

RawJSONPayload specifies the format for RawJSONPayloadKind ActionNotifications

#### type Record

```go
type Record map[string]interface{}
```

Record is a JSON document entity contained in collections

#### type Role

```go
type Role struct {

	// created at
	// Required: true
	CreatedAt strfmt.DateTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy string `json:"createdBy"`

	// name
	// Required: true
	Name string `json:"name"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

Role role

#### type RolePermission

```go
type RolePermission struct {

	// added at
	// Required: true
	// Format: date-time
	AddedAt strfmt.DateTime `json:"addedAt"`

	// added by
	// Required: true
	AddedBy string `json:"addedBy"`

	// permission
	// Required: true
	Permission string `json:"permission"`

	// role
	// Required: true
	Role string `json:"role"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

RolePermission role permission

#### type Rule

```go
type Rule struct {
	ID         string          `json:"id,omitempty"`
	Name       string          `json:"name" binding:"required"`
	Module     string          `json:"module,omitempty"`
	Match      string          `json:"match" binding:"required"`
	Actions    []CatalogAction `json:"actions,omitempty"`
	Owner      string          `json:"owner" binding:"required"`
	Created    string          `json:"created,omitempty"`
	Modified   string          `json:"modified,omitempty"`
	CreatedBy  string          `json:"createdBy,omitempty"`
	ModifiedBy string          `json:"modifiedBy,omitempty"`
	Version    int             `json:"version,omitempty"`
}
```

Rule represents a rule for transforming results at search time. A rule consits
of a `match` clause and a collection of transformation actions

#### type SearchContext

```go
type SearchContext struct {
	User string
	App  string
}
```

SearchContext specifies the user and app context for a search job

#### type SearchJob

```go
type SearchJob struct {
	Sid     string           `json:"sid"`
	Content SearchJobContent `json:"content"`
	Context *SearchContext
}
```

SearchJob specifies the fields returned for a /search/jobs/ entry for a specific
job

#### type SearchJobContent

```go
type SearchJobContent struct {
	Sid             string        `json:"sid"`
	EventCount      int           `json:"eventCount"`
	DispatchState   DispatchState `json:"dispatchState"`
	DiskUsage       int64         `json:"diskUsage"`
	IsFinalized     bool          `json:"isFinalized"`
	OptimizedSearch string        `json:"optimizedSearch"`
	ScanCount       int64         `json:"scanCount"`
}
```

SearchJobContent represents the content json object from /search/jobs response

#### type SearchResults

```go
type SearchResults struct {
	Preview     bool                     `json:"preview"`
	InitOffset  int                      `json:"init_offset"`
	Messages    []interface{}            `json:"messages"`
	Results     []map[string]interface{} `json:"results"`
	Fields      []map[string]interface{} `json:"fields"`
	Highlighted map[string]interface{}   `json:"highlighted"`
}
```

SearchResults represents the /search/jobs/{sid}/events or
/search/jobs/{sid}/results response

#### type SplunkEventPayload

```go
type SplunkEventPayload struct {
	Event      map[string]interface{} `json:"event" binding:"required"`
	Fields     map[string]string      `json:"fields" binding:"required"`
	Host       string                 `json:"host" binding:"required"`
	Index      string                 `json:"index" binding:"required"`
	Source     string                 `json:"source" binding:"required"`
	Sourcetype string                 `json:"sourcetype" binding:"required"`
	Time       float64                `json:"time" binding:"required"`
}
```

SplunkEventPayload is the payload for a notification coming from Splunk

#### type Tenant

```go
type Tenant struct {

	// created at
	// Required: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy string `json:"createdBy"`

	// name
	// Required: true
	Name string `json:"name"`

	// status
	// Required: true
	Status string `json:"status"`
}
```

Tenant tenant

#### type Ticker

```go
type Ticker struct {
}
```

Ticker is a wrapper of time.Ticker with additional functionality

#### func  NewTicker

```go
func NewTicker(duration time.Duration) *Ticker
```
NewTicker spits out a pointer to Ticker model. It sets ticker to stop state by
default

#### func (*Ticker) GetChan

```go
func (t *Ticker) GetChan() <-chan time.Time
```
GetChan returns the channel from ticker

#### func (*Ticker) IsRunning

```go
func (t *Ticker) IsRunning() bool
```
IsRunning returns bool indicating whether or not ticker is running

#### func (*Ticker) Reset

```go
func (t *Ticker) Reset()
```
Reset resets ticker

#### func (*Ticker) Start

```go
func (t *Ticker) Start()
```
Start starts a new ticker and set property running to true

#### func (*Ticker) Stop

```go
func (t *Ticker) Stop()
```
Stop stops ticker and set property running to false

#### type User

```go
type User struct {
	ID                string   `json:"id"`
	FirstName         string   `json:"firstName,omitempty"`
	LastName          string   `json:"lastName,omitempty"`
	Email             string   `json:"email,omitempty"`
	Locale            string   `json:"locale,omitempty"`
	Name              string   `json:"name,omitempty"`
	TenantMemberships []string `json:"tenantMemberships,omitempty"`
}
```

User represents a user object

#### type ValidateInfo

```go
type ValidateInfo struct {

	// name
	// Required: true
	// Max Length: 36
	// Min Length: 4
	Name string `json:"name"`

	// tenants
	// Required: true
	Tenants []string `json:"tenants"`
}
```

ValidateInfo validate info
