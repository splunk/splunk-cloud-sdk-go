# model
--
    import "github.com/splunk/splunk-cloud-sdk-go/model"

Package model contains Splunk Cloud models for each service.

Deprecated: v0.6.1 - these models have been moved to their respective
services/<service>/models.go files, see below for details for each model.

## Usage

```go
const (
	// EmailKind is Deprecated: please use services/action.EmailKind
	EmailKind = action.EmailKind
	// WebhookKind is Deprecated: please use services/action.WebhookKind
	WebhookKind = action.WebhookKind
	// SNSKind is Deprecated: please use services/action.SNSKind
	SNSKind = action.SNSKind
)
```

```go
const (
	// StatusQueued is Deprecated: please use services/action.StatusQueued
	StatusQueued = action.StatusQueued
	// StatusRunning is Deprecated: please use services/action.StatusRunning
	StatusRunning = action.StatusRunning
	// StatusDone is Deprecated: please use services/action.StatusDone
	StatusDone = action.StatusDone
	// StatusFailed is Deprecated: please use services/action.StatusFailed
	StatusFailed = action.StatusFailed
)
```

```go
const (
	//SplunkEventKind is Deprecated: please use services/action.SplunkEventKind
	SplunkEventKind = action.SplunkEventKind
	//RawJSONPayloadKind is Deprecated: please use services/action.RawJSONPayloadKind
	RawJSONPayloadKind = action.RawJSONPayloadKind
)
```

#### type Action

```go
type Action = action.Action
```

Action is Deprecated: please use services/action.Action

#### func  NewEmailAction

```go
func NewEmailAction(name string, htmlPart string, subjectPart string, textPart string, templateName string, addresses []string) *Action
```
NewEmailAction is Deprecated: please use services/action.NewEmailAction

#### func  NewSNSAction

```go
func NewSNSAction(name string, topic string, message string) *Action
```
NewSNSAction is Deprecated: please use services/action.NewSNSAction

#### func  NewWebhookAction

```go
func NewWebhookAction(name string, webhookURL string, message string) *Action
```
NewWebhookAction is Deprecated: please use services/action.NewWebhookAction

#### type ActionError

```go
type ActionError = action.Error
```

ActionError is Deprecated: please use services/action.Error

#### type ActionKind

```go
type ActionKind = action.Kind
```

ActionKind is Deprecated: please use services/action.Kind

#### type ActionNotification

```go
type ActionNotification = action.Notification
```

ActionNotification is Deprecated: please use services/action.Notification

#### type ActionNotificationKind

```go
type ActionNotificationKind = action.NotificationKind
```

ActionNotificationKind is Deprecated: please use
services/action.NotificationKind

#### type ActionPayload

```go
type ActionPayload = action.Payload
```

ActionPayload is Deprecated: please use services/action.Payload

#### type ActionStatus

```go
type ActionStatus = action.Status
```

ActionStatus is Deprecated: please use services/action.Status

#### type ActionStatusState

```go
type ActionStatusState = action.StatusState
```

ActionStatusState is Deprecated: please use services/action.StatusState

#### type ActionTriggerResponse

```go
type ActionTriggerResponse = action.TriggerResponse
```

ActionTriggerResponse is Deprecated: please use services/action.TriggerResponse

#### type ActionUpdateFields

```go
type ActionUpdateFields = action.UpdateFields
```

ActionUpdateFields is Deprecated: please use services/action.UpdateFields

#### type ActivatePipelineRequest

```go
type ActivatePipelineRequest = streams.ActivatePipelineRequest
```

ActivatePipelineRequest is Deprecated: please use
services/streams.ActivatePipelineRequest

#### type AdditionalProperties

```go
type AdditionalProperties = streams.AdditionalProperties
```

AdditionalProperties is Deprecated: please use
services/streams.AdditionalProperties

#### type AuthError

```go
type AuthError = kvstore.AuthError
```

AuthError is Deprecated: please use services/kvstore.AuthError

#### type CatalogAction

```go
type CatalogAction = catalog.Action
```

CatalogAction is Deprecated: please use services/catalog.Action

#### func  NewAliasAction

```go
func NewAliasAction(field string, alias string, owner string) *CatalogAction
```
NewAliasAction is Deprecated: please use services/catalog.NewAliasAction

#### func  NewAutoKVAction

```go
func NewAutoKVAction(mode string, owner string) *CatalogAction
```
NewAutoKVAction is Deprecated: please use services/catalog.NewAutoKVAction

#### func  NewEvalAction

```go
func NewEvalAction(field string, expression string, owner string) *CatalogAction
```
NewEvalAction is Deprecated: please use services/catalog.NewEvalAction

#### func  NewLookupAction

```go
func NewLookupAction(expression string, owner string) *CatalogAction
```
NewLookupAction is Deprecated: please use services/catalog.NewLookupAction

#### func  NewRegexAction

```go
func NewRegexAction(field string, pattern string, limit *int, owner string) *CatalogAction
```
NewRegexAction is Deprecated: please use services/catalog.NewRegexAction

#### func  NewUpdateAliasAction

```go
func NewUpdateAliasAction(field *string, alias *string) *CatalogAction
```
NewUpdateAliasAction is Deprecated: please use
services/catalog.NewUpdateAliasAction

#### func  NewUpdateAutoKVAction

```go
func NewUpdateAutoKVAction(mode *string) *CatalogAction
```
NewUpdateAutoKVAction is Deprecated: please use
services/catalog.NewUpdateAutoKVAction

#### func  NewUpdateEvalAction

```go
func NewUpdateEvalAction(field *string, expression *string) *CatalogAction
```
NewUpdateEvalAction is Deprecated: please use
services/catalog.NewUpdateEvalAction

#### func  NewUpdateLookupAction

```go
func NewUpdateLookupAction(expression *string) *CatalogAction
```
NewUpdateLookupAction is Deprecated: please use
services/catalog.NewUpdateLookupAction

#### func  NewUpdateRegexAction

```go
func NewUpdateRegexAction(field *string, pattern *string, limit *int) *CatalogAction
```
NewUpdateRegexAction is Deprecated: please use
services/catalog.NewUpdateRegexAction

#### type CatalogActionKind

```go
type CatalogActionKind = catalog.ActionKind
```

CatalogActionKind is Deprecated: please use services/catalog.ActionKind

```go
const (
	// ALIAS is Deprecated: please use services/catalog.Alias
	ALIAS CatalogActionKind = catalog.Alias
	// AUTOKV is Deprecated: please use services/catalog.AutoKV
	AUTOKV CatalogActionKind = catalog.AutoKV
	// REGEX is Deprecated: please use services/catalog.Regex
	REGEX CatalogActionKind = catalog.Regex
	// EVAL is Deprecated: please use services/catalog.Eval
	EVAL CatalogActionKind = catalog.Eval
	// LOOKUPACTION is Deprecated: please use services/catalog.LookupAction
	LOOKUPACTION CatalogActionKind = catalog.LookupAction
)
```

#### type CreateJobRequest

```go
type CreateJobRequest = search.CreateJobRequest
```

CreateJobRequest is Deprecated: please use services/search.CreateJobRequest

#### type DataType

```go
type DataType = catalog.DataType
```

DataType is Deprecated: please use services/catalog.DataType

```go
const (
	// DATE is Deprecated: please use services/catalog.Date
	DATE DataType = catalog.Date
	// NUMBER is Deprecated: please use services/catalog.Number
	NUMBER DataType = catalog.Number
	// OBJECTID is Deprecated: please use services/catalog.ObjectID
	OBJECTID DataType = catalog.ObjectID
	// STRING is Deprecated: please use services/catalog.String
	STRING DataType = catalog.String
	// DATATYPEUNKNOWN is Deprecated: please use services/catalog.DataTypeUnknown
	DATATYPEUNKNOWN DataType = catalog.DataTypeUnknown
)
```

#### type DatasetCreationPayload

```go
type DatasetCreationPayload = catalog.DatasetCreationPayload
```

DatasetCreationPayload is Deprecated: please use
services/catalog.DatasetCreationPayload

#### type DatasetInfo

```go
type DatasetInfo = catalog.DatasetInfo
```

DatasetInfo is Deprecated: please use services/catalog.DatasetInfo

#### type DatasetInfoKind

```go
type DatasetInfoKind = catalog.DatasetInfoKind
```

DatasetInfoKind is Deprecated: please use services/catalog.DatasetInfoKind

```go
const (
	// LOOKUP is Deprecated: please use services/catalog.Lookup
	LOOKUP DatasetInfoKind = catalog.Lookup
	// KVCOLLECTION is Deprecated: please use services/catalog.KvCollection
	KVCOLLECTION DatasetInfoKind = catalog.KvCollection
	// INDEX is Deprecated: please use services/catalog.Index
	INDEX DatasetInfoKind = catalog.Index
)
```

#### type DslCompilationRequest

```go
type DslCompilationRequest = streams.DslCompilationRequest
```

DslCompilationRequest is Deprecated: please use
services/streams.DslCompilationRequest

#### type Error

```go
type Error = kvstore.Error
```

Error is Deprecated: please use services/kvstore.Error

#### type Event

```go
type Event = ingest.Event
```

Event is Deprecated: please use services/ingest.Event

#### type Field

```go
type Field = catalog.Field
```

Field is Deprecated: please use services/catalog.Field

#### type FieldType

```go
type FieldType = catalog.FieldType
```

FieldType is Deprecated: please use services/catalog.FieldType

```go
const (
	// DIMENSION is Deprecated: please use services/catalog.Dimension
	DIMENSION FieldType = catalog.Dimension
	// MEASURE is Deprecated: please use services/catalog.Measure
	MEASURE FieldType = catalog.Measure
	// FIELDTYPEUNKNOWN is Deprecated: please use services/catalog.FieldTypeUnknown
	FIELDTYPEUNKNOWN = catalog.FieldTypeUnknown
)
```

#### type Group

```go
type Group = identity.Group
```

Group is Deprecated: please use services/identity.Group

#### type GroupMember

```go
type GroupMember = identity.GroupMember
```

GroupMember is Deprecated: please use services/identity.GroupMember

#### type GroupRole

```go
type GroupRole = identity.GroupRole
```

GroupRole is Deprecated: please use services/identity.GroupRole

#### type IndexDefinition

```go
type IndexDefinition = kvstore.IndexDefinition
```

IndexDefinition is Deprecated: please use services/kvstore.IndexDefinition

#### type IndexDescription

```go
type IndexDescription = kvstore.IndexDescription
```

IndexDescription is Deprecated: please use services/kvstore.IndexDescription

#### type IndexFieldDefinition

```go
type IndexFieldDefinition = kvstore.IndexFieldDefinition
```

IndexFieldDefinition is Deprecated: please use
services/kvstore.IndexFieldDefinition

#### type JobMessageType

```go
type JobMessageType = search.JobMessageType
```

JobMessageType is Deprecated: please use services/search.JobMessageType

```go
const (
	// InfoType is Deprecated: please use services/search.InfoType
	InfoType JobMessageType = search.InfoType
	// FatalType is Deprecated: please use services/search.FatalType
	FatalType JobMessageType = search.FatalType
	// ErrorType is Deprecated: please use services/search.ErrorType
	ErrorType JobMessageType = search.ErrorType
	// DebugType is Deprecated: please use services/search.DebugType
	DebugType JobMessageType = search.DebugType
)
```

#### type JobResultsParams

```go
type JobResultsParams = search.JobResultsParams
```

JobResultsParams is Deprecated: please use services/search.JobResultsParams

#### type JobStatus

```go
type JobStatus = search.JobAction
```

JobStatus is Deprecated: please use services/search.JobAction

```go
const (
	// JobCanceled is Deprecated: please use services/search.JobCanceled
	JobCanceled JobStatus = search.JobCanceled
	// JobFinalized is Deprecated: please use services/search.JobFinalized
	JobFinalized JobStatus = search.JobFinalized
)
```

#### type Key

```go
type Key = kvstore.Key
```

Key is Deprecated: please use services/kvstore.Key

#### type LookupValue

```go
type LookupValue = kvstore.LookupValue
```

LookupValue is Deprecated: please use services/kvstore.LookupValue

#### type Member

```go
type Member = identity.Member
```

Member is Deprecated: please use services/identity.Member

#### type Metric

```go
type Metric = ingest.Metric
```

Metric is Deprecated: please use services/ingest.Metric

#### type MetricAttribute

```go
type MetricAttribute = ingest.MetricAttribute
```

MetricAttribute is Deprecated: please use services/ingest.MetricAttribute

#### type MetricEvent

```go
type MetricEvent = ingest.MetricEvent
```

MetricEvent is Deprecated: please use services/ingest.MetricEvent

#### type Module

```go
type Module = catalog.Module
```

Module is Deprecated: please use services/catalog.Module

#### type PaginatedPipelineResponse

```go
type PaginatedPipelineResponse = streams.PaginatedPipelineResponse
```

PaginatedPipelineResponse is Deprecated: please use
services/streams.PaginatedPipelineResponse

#### type PatchJobResponse

```go
type PatchJobResponse = search.PatchJobResponse
```

PatchJobResponse is Deprecated: please use services/search.PatchJobResponse

#### type PingOKBody

```go
type PingOKBody = kvstore.PingOKBody
```

PingOKBody is Deprecated: please use services/kvstore.PingOKBody

#### type PingOKBodyStatus

```go
type PingOKBodyStatus = kvstore.PingOKBodyStatus
```

PingOKBodyStatus is Deprecated: please use services/kvstore.PingOKBodyStatus

```go
const (
	// PingOKBodyStatusHealthy is Deprecated: please use services/kvstore.PingOKBodyStatusHealthy
	PingOKBodyStatusHealthy PingOKBodyStatus = kvstore.PingOKBodyStatusHealthy

	// PingOKBodyStatusUnhealthy is Deprecated: please use services/kvstore.PingOKBodyStatusUnhealthy
	PingOKBodyStatusUnhealthy PingOKBodyStatus = kvstore.PingOKBodyStatusUnhealthy

	// PingOKBodyStatusUnknown is Deprecated: please use services/kvstore.PingOKBodyStatusUnknown
	PingOKBodyStatusUnknown PingOKBodyStatus = kvstore.PingOKBodyStatusUnknown
)
```

#### type Pipeline

```go
type Pipeline = streams.Pipeline
```

Pipeline is Deprecated: please use services/streams.Pipeline

#### type PipelineDeleteResponse

```go
type PipelineDeleteResponse = streams.PipelineDeleteResponse
```

PipelineDeleteResponse is Deprecated: please use
services/streams.PipelineDeleteResponse

#### type PipelineQueryParams

```go
type PipelineQueryParams = streams.PipelineQueryParams
```

PipelineQueryParams is Deprecated: please use
services/streams.PipelineQueryParams

#### type PipelineRequest

```go
type PipelineRequest = streams.PipelineRequest
```

PipelineRequest is Deprecated: please use services/streams.PipelineRequest

#### type PipelineStatus

```go
type PipelineStatus = streams.PipelineStatus
```

PipelineStatus is Deprecated: please use services/streams.PipelineStatus

```go
const (
	// Created is Deprecated: please use services/streams.Created
	Created PipelineStatus = streams.Created
	// Activated is Deprecated: please use services/streams.Activated
	Activated PipelineStatus = streams.Activated
)
```

#### type PrevalenceType

```go
type PrevalenceType = catalog.PrevalenceType
```

PrevalenceType is Deprecated: please use services/catalog.PrevalenceType

```go
const (
	// ALL is Deprecated: please use services/catalog.All
	ALL PrevalenceType = catalog.All
	// SOME is Deprecated: please use services/catalog.Some
	SOME PrevalenceType = catalog.Some
	// PREVALANCEUNKNOWN is Deprecated: please use services/catalog.PrevalenceUnknown
	PREVALANCEUNKNOWN PrevalenceType = catalog.PrevalenceUnknown
)
```

#### type Principal

```go
type Principal = identity.Principal
```

Principal is Deprecated: please use services/identity.Principal

#### type QueryParameters

```go
type QueryParameters = search.QueryParameters
```

QueryParameters is Deprecated: please use services/search.QueryParameters

#### type RawJSONPayload

```go
type RawJSONPayload = action.RawJSONPayload
```

RawJSONPayload is Deprecated: please use services/action.RawJSONPayload

#### type Record

```go
type Record = kvstore.Record
```

Record is Deprecated: please use services/kvstore.Record

#### type ResultsNotReadyResponse

```go
type ResultsNotReadyResponse = search.ResultsNotReadyResponse
```

ResultsNotReadyResponse is Deprecated: please use
services/search.ResultsNotReadyResponse

#### type Role

```go
type Role = identity.Role
```

Role is Deprecated: please use services/identity.Role

#### type RolePermission

```go
type RolePermission = identity.RolePermission
```

RolePermission is Deprecated: please use services/identity.RolePermission

#### type Rule

```go
type Rule = catalog.Rule
```

Rule is Deprecated: please use services/catalog.Rule

#### type RuleUpdateFields

```go
type RuleUpdateFields = catalog.RuleUpdateFields
```

RuleUpdateFields is Deprecated: please use services/catalog.RuleUpdateFields

#### type SearchJob

```go
type SearchJob = search.Job
```

SearchJob is Deprecated: please use services/search.Job

#### type SearchJobMessages

```go
type SearchJobMessages = search.JobMessages
```

SearchJobMessages is Deprecated: please use services/search.JobMessages

#### type SearchJobStatus

```go
type SearchJobStatus = search.JobStatus
```

SearchJobStatus is Deprecated: please use services/search.JobStatus

```go
const (
	// Queued is Deprecated: please use services/search.Queued
	Queued SearchJobStatus = search.Queued
	// Parsing is Deprecated: please use services/search.Parsing
	Parsing SearchJobStatus = search.Parsing
	// Running is Deprecated: please use services/search.Running
	Running SearchJobStatus = search.Running
	// Finalizing is Deprecated: please use services/search.Finalizing
	Finalizing SearchJobStatus = search.Finalizing
	// Failed is Deprecated: please use services/search.Failed
	Failed SearchJobStatus = search.Failed
	// Done is Deprecated: please use services/search.Done
	Done SearchJobStatus = search.Done
)
```

#### type SearchResults

```go
type SearchResults = search.Results
```

SearchResults is Deprecated: please use services/search.Results

#### type SplunkEventPayload

```go
type SplunkEventPayload = action.SplunkEventPayload
```

SplunkEventPayload is Deprecated: please use services/action.SplunkEventPayload

#### type Tenant

```go
type Tenant = identity.Tenant
```

Tenant is Deprecated: please use services/identity.Tenant

#### type UpdateDatasetInfoFields

```go
type UpdateDatasetInfoFields = catalog.UpdateDatasetInfoFields
```

UpdateDatasetInfoFields is Deprecated: please use
services/catalog.UpdateDatasetInfoFields

#### type UplEdge

```go
type UplEdge = streams.UplEdge
```

UplEdge is Deprecated: please use services/streams.UplEdge

#### type UplNode

```go
type UplNode = streams.UplNode
```

UplNode is Deprecated: please use services/streams.UplNode

#### type UplPipeline

```go
type UplPipeline = streams.UplPipeline
```

UplPipeline is Deprecated: please use services/streams.UplPipeline

#### type ValidateInfo

```go
type ValidateInfo = identity.ValidateInfo
```

ValidateInfo is Deprecated: please use services/identity.ValidateInfo
