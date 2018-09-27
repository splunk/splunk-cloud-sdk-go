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
	Addresses []string `json:"addresses,omitempty"`

	// SNS action fields:
	// Topic to trigger SNS action
	Topic string `json:"topic,omitempty"`
	// Message to send via SNS or Webhook action
	Message string `json:"message,omitempty"`

	// Webhook action fields:
	// WebhookURL to trigger Webhook action
	WebhookURL string `json:"webhookUrl,omitempty"`
}
```

ActionUpdateFields defines the fields that may be updated for an existing Action

#### type ActivatePipelineRequest

```go
type ActivatePipelineRequest struct {
	IDs           []string `json:"ids"`
	SkipSavePoint bool     `json:"skipSavepoint"`
}
```

ActivatePipelineRequest contains the request to activate the pipeline

#### type AdditionalProperties

```go
type AdditionalProperties map[string][]string
```

AdditionalProperties contain the properties in an activate/deactivate response

#### type AuthError

```go
type AuthError struct {

	// The reason of the auth error
	Reason string `json:"reason"`
}
```

AuthError auth error reason

#### type CatalogAction

```go
type CatalogAction struct {
	ID         string            `json:"id,omitempty"`
	RuleID     string            `json:"ruleid,omitempty"`
	Kind       CatalogActionKind `json:"kind,omitempty"`
	Owner      string            `json:"owner,omitempty"`
	Created    string            `json:"created,omitempty"`
	Modified   string            `json:"modified,omitempty"`
	CreatedBy  string            `json:"createdBy,omitempty"`
	ModifiedBy string            `json:"modifiedBy,omitempty"`
	Version    int               `json:"version,omitempty"`
	Field      string            `json:"field,omitempty"`
	Alias      string            `json:"alias,omitempty"`
	Mode       string            `json:"mode,omitempty"`
	Expression string            `json:"expression,omitempty"`
	Pattern    string            `json:"pattern,omitempty"`
	Limit      *int              `json:"limit,omitempty"`
}
```

CatalogAction represents a specific search time transformation action. This
struct should NOT be directly used to construct object, use the NewXXXAction()
instead

#### func  NewAliasAction

```go
func NewAliasAction(field string, alias string, owner string) *CatalogAction
```
NewAliasAction creates a new alias kind action

#### func  NewAutoKVAction

```go
func NewAutoKVAction(mode string, owner string) *CatalogAction
```
NewAutoKVAction creates a new autokv kind action

#### func  NewEvalAction

```go
func NewEvalAction(field string, expression string, owner string) *CatalogAction
```
NewEvalAction creates a new eval kind action

#### func  NewLookupAction

```go
func NewLookupAction(expression string, owner string) *CatalogAction
```
NewLookupAction creates a new lookup kind action

#### func  NewRegexAction

```go
func NewRegexAction(field string, pattern string, limit *int, owner string) *CatalogAction
```
NewRegexAction creates a new regex kind action

#### func  NewUpdateAliasAction

```go
func NewUpdateAliasAction(field *string, alias *string) *CatalogAction
```
NewUpdateAliasAction updates an existing alias kind action

#### func  NewUpdateAutoKVAction

```go
func NewUpdateAutoKVAction(mode *string) *CatalogAction
```
NewUpdateAutoKVAction updates an existing autokv kind action

#### func  NewUpdateEvalAction

```go
func NewUpdateEvalAction(field *string, expression *string) *CatalogAction
```
NewUpdateEvalAction updates an existing eval kind action

#### func  NewUpdateLookupAction

```go
func NewUpdateLookupAction(expression *string) *CatalogAction
```
NewUpdateLookupAction updates an existing lookup kind action

#### func  NewUpdateRegexAction

```go
func NewUpdateRegexAction(field *string, pattern *string, limit *int) *CatalogAction
```
NewUpdateRegexAction updates an existing regex kind action

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

#### type DatasetCreationPayload

```go
type DatasetCreationPayload struct {
	ID           string          `json:"id,omitempty"`
	Name         string          `json:"name"`
	Kind         DatasetInfoKind `json:"kind"`
	Owner        string          `json:"owner,omitempty"`
	Module       string          `json:"module,omitempty"`
	Capabilities string          `json:"capabilities"`
	Fields       []Field         `json:"fields,omitempty"`
	Readroles    []string        `json:"readroles,omitempty"`
	Writeroles   []string        `json:"writeroles,omitempty"`

	ExternalKind       string `json:"externalKind,omitempty"`
	ExternalName       string `json:"externalName,omitempty"`
	CaseSensitiveMatch *bool  `json:"caseSensitiveMatch,omitempty"`
	Filter             string `json:"filter,omitempty"`
	MaxMatches         *int   `json:"maxMatches,omitempty"`
	MinMatches         *int   `json:"minMatches,omitempty"`
	DefaultMatch       string `json:"defaultMatch,omitempty"`

	Datatype string `json:"datatype,omitempty"`
	Disabled *bool  `json:"disabled,omitempty"`
}
```

DatasetCreationPayload represents the sources of data that can be searched by
Splunk

#### type DatasetInfo

```go
type DatasetInfo struct {
	ID           string          `json:"id,omitempty"`
	Name         string          `json:"name"`
	Kind         DatasetInfoKind `json:"kind"`
	Owner        string          `json:"owner,omitempty"`
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
	Disabled bool   `json:"disabled"`
}
```

DatasetInfo represents the sources of data that can be searched by Splunk

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

#### type DslCompilationRequest

```go
type DslCompilationRequest struct {
	Dsl string `json:"dsl"`
}
```

DslCompilationRequest contains the DSL that needs to be compiled into a valid
UPL JSON

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
	// Specifies a JSON object that contains explicit custom fields to be defined at index time.
	Attributes map[string]interface{} `json:"attributes"`
	// JSON object for the event.
	Body interface{} `json:"body"`
	// Epoch time in milliseconds.
	Timestamp int64 `json:"timestamp"`
	// Optional nanoseconds part of the timestamp.
	Nanos int32 `json:"nanos"`
	// The source value to assign to the event data. For example, if you are sending data from an app that you are developing,
	// set this key to the name of the app.
	Source string `json:"source"`
	// The sourcetype value assigned to the event data.
	Sourcetype string `json:"sourcetype"`
	// The host value assigned to the event data. Typically, this is the hostname of the client from which you are sending data.
	Host string `json:"host"`
	// An optional ID that uniquely identifies the metric data. It is used to deduplicate the data if same data is set multiple times.
	// If ID is not specified, it will be assigned by the system.
	ID string `json:"id"`
}
```

Event defines raw event send to event endpoint

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

JobStatus defines actions to be taken on an existing search job.

```go
const (
	JobCanceled  JobStatus = "canceled"
	JobFinalized JobStatus = "finalized"
)
```
Define supported job actions

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

#### type Module

```go
type Module struct {
	Name string `json:"name"`
}
```

Module represents catalog module

#### type PaginatedPipelineResponse

```go
type PaginatedPipelineResponse struct {
	Items []Pipeline `json:"items"`
	Total int64      `json:"total"`
}
```

PaginatedPipelineResponse contains the pipeline response

#### type PatchJobResponse

```go
type PatchJobResponse struct {
	// Run time messages from Splunkd.
	Messages SearchJobMessages `json:"messages"`
}
```

PatchJobResponse defines the response from patch endpoint

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

#### type Pipeline

```go
type Pipeline struct {
	ActivatedDate            int64          `json:"activatedDate"`
	ActivatedUserID          string         `json:"activatedUserId"`
	ActivatedVersion         int64          `json:"activatedVersion"`
	CreateDate               int64          `json:"createDate"`
	CreateUserID             string         `json:"createUserId"`
	CurrentVersion           int64          `json:"currentVersion"`
	Data                     UplPipeline    `json:"data"`
	Description              string         `json:"description"`
	ID                       string         `json:"id"`
	JobID                    string         `json:"jobId"`
	LastUpdateDate           int64          `json:"lastUpdateDate"`
	LastUpdateUserID         string         `json:"lastUpdateUserId"`
	Name                     string         `json:"name"`
	Status                   PipelineStatus `json:"status"`
	StatusMessage            string         `json:"statusMessage"`
	StreamingConfigurationID int64          `json:"streamingConfigurationId"`
	TenantID                 string         `json:"tenantId"`
	ValidationMessages       []string       `json:"validationMessages"`
	Version                  int64          `json:"version"`
}
```

Pipeline defines a pipeline object

#### type PipelineDeleteResponse

```go
type PipelineDeleteResponse struct {
	CouldDeactivate bool `json:"couldDeactivate"`
	Running         bool `json:"running"`
}
```

PipelineDeleteResponse contains the response returned as a result of a delete
pipeline call

#### type PipelineQueryParams

```go
type PipelineQueryParams struct {
	Offset       *int32  `json:"offset,omitempty"`
	PageSize     *int32  `json:"pageSize,omitempty"`
	SortField    *string `json:"sortField,omitempty"`
	SortDir      *string `json:"sortDir,omitempty"`
	Activated    *bool   `json:"activated,omitempty"`
	CreateUserID *string `json:"createUserId,omitempty"`
	Name         *string `json:"name,omitempty"`
	IncludeData  *bool   `json:"includeData,omitempty"`
}
```

PipelineQueryParams contains the query parameters that can be provided by the
user to fetch specific pipelines

#### type PipelineRequest

```go
type PipelineRequest struct {
	BypassValidation         bool         `json:"bypassValidation"`
	CreateUserID             string       `json:"createUserId"`
	Data                     *UplPipeline `json:"data"`
	Description              string       `json:"description"`
	Name                     string       `json:"name"`
	StreamingConfigurationID *int64       `json:"streamingConfigurationId,omitempty"`
}
```

PipelineRequest contains the pipeline data

#### type PipelineStatus

```go
type PipelineStatus string
```

PipelineStatus reflects the status of a pipeline

```go
const (
	// Created status
	Created PipelineStatus = "CREATED"
	// Activated status
	Activated PipelineStatus = "ACTIVATED"
)
```

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
	Name       string          `json:"name"`
	Module     string          `json:"module,omitempty"`
	Match      string          `json:"match"`
	Actions    []CatalogAction `json:"actions,omitempty"`
	Owner      string          `json:"owner,omitempty"`
	Created    string          `json:"created,omitempty"`
	Modified   string          `json:"modified,omitempty"`
	CreatedBy  string          `json:"createdBy,omitempty"`
	ModifiedBy string          `json:"modifiedBy,omitempty"`
	Version    int             `json:"version,omitempty"`
}
```

Rule represents a rule for transforming results at search time. A rule consists
of a `match` clause and a collection of transformation actions

#### type RuleUpdateFields

```go
type RuleUpdateFields struct {
	Name    string `json:"name,omitempty"`
	Module  string `json:"module,omitempty"`
	Match   string `json:"match,omitempty"`
	Owner   string `json:"owner,omitempty"`
	Version int    `json:"version,omitempty"`
}
```

RuleUpdateFields represents the set of rule properties that can be updated

#### type SearchJob

```go
type SearchJob struct {
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
	Status SearchJobStatus `json:"status,omitempty"`
	// An estimate of how close the job is to completing.
	PercentComplete float64 `json:"percentComplete,omitempty"`
	// The number of results produced so far for the search job.
	ResultsAvailable int64 `json:"resultsAvailable,omitempty"`
	// Run time messages from Splunkd.
	Messages SearchJobMessages `json:"messages,omitempty"`
}
```

SearchJob represents a fully-constructed search job, including read-only fields.

#### type SearchJobMessages

```go
type SearchJobMessages []struct {
	// Enum [INFO, FATAL, ERROR, DEBUG]
	Type string `json:"type"`
	// message text
	Text string `json:"text"`
}
```

SearchJobMessages is used in search results or search job.

#### type SearchJobStatus

```go
type SearchJobStatus string
```

SearchJobStatus describes status of a search job

```go
const (
	Queued     SearchJobStatus = "queued"
	Parsing    SearchJobStatus = "parsing"
	Running    SearchJobStatus = "running"
	Finalizing SearchJobStatus = "finalizing"
	Failed     SearchJobStatus = "failed"
	Done       SearchJobStatus = "done"
)
```
Supported SearchJobStatus constants

#### type SearchResults

```go
type SearchResults struct {
	// Run time messages from Splunkd.
	Messages SearchJobMessages        `json:"messages"`
	Results  []map[string]interface{} `json:"results"`
	Fields   []map[string]interface{} `json:"fields"`
}
```

SearchResults represents results from a search job

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

#### type UpdateDatasetInfoFields

```go
type UpdateDatasetInfoFields struct {
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
	Disabled *bool  `json:"disabled,omitempty"`
}
```

UpdateDatasetInfoFields represents the sources of data that can be updated by
Splunk, same structure as DatasetInfo

#### type UplEdge

```go
type UplEdge struct {
	Attributes interface{} `json:"attributes"`
	SourceNode string      `json:"sourceNode"`
	SourcePort string      `json:"sourcePort"`
	TargetNode string      `json:"targetNode"`
	TargetPort string      `json:"targetPort"`
}
```

UplEdge contains information on the edges between two pipeline nodes

#### type UplNode

```go
type UplNode interface{}
```

UplNode defines the nodes forming a pipeline

#### type UplPipeline

```go
type UplPipeline struct {
	Edges    []UplEdge `json:"edges"`
	Nodes    []UplNode `json:"nodes"`
	RootNode []string  `json:"root-node"`
	Version  int32     `json:"version"`
}
```

UplPipeline contains the pipeline data

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
