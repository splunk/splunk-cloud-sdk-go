// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

// Package model contains Splunk Cloud models for each service.
//
// Deprecated: v0.6.1 - these models have been moved to their respective services/<service>/models.go files, see below for details for each model.
package model

import (
	"github.com/splunk/splunk-cloud-sdk-go/services/action"
	"github.com/splunk/splunk-cloud-sdk-go/services/catalog"
	"github.com/splunk/splunk-cloud-sdk-go/services/identity"
	"github.com/splunk/splunk-cloud-sdk-go/services/ingest"
	"github.com/splunk/splunk-cloud-sdk-go/services/kvstore"
	"github.com/splunk/splunk-cloud-sdk-go/services/search"
	"github.com/splunk/splunk-cloud-sdk-go/services/streams"
)

//
// Deprecated: Action models
//

// ActionKind is Deprecated: please use services/action.Kind
type ActionKind = action.Kind

const (
	// EmailKind is Deprecated: please use services/action.EmailKind
	EmailKind = action.EmailKind
	// WebhookKind is Deprecated: please use services/action.WebhookKind
	WebhookKind = action.WebhookKind
	// SNSKind is Deprecated: please use services/action.SNSKind
	SNSKind = action.SNSKind
)

// ActionUpdateFields is Deprecated: please use services/action.UpdateFields
type ActionUpdateFields = action.UpdateFields

// Action is Deprecated: please use services/action.Action
type Action = action.Action

// NewEmailAction is Deprecated: please use services/action.NewEmailAction
func NewEmailAction(name string, htmlPart string, subjectPart string, textPart string, templateName string, addresses []string) *Action {
	return action.NewEmailAction(name, htmlPart, subjectPart, textPart, templateName, addresses)
}

// NewSNSAction is Deprecated: please use services/action.NewSNSAction
func NewSNSAction(name string, topic string, message string) *Action {
	return action.NewSNSAction(name, topic, message)
}

// NewWebhookAction is Deprecated: please use services/action.NewWebhookAction
func NewWebhookAction(name string, webhookURL string, message string) *Action {
	return action.NewWebhookAction(name, webhookURL, message)
}

// ActionStatusState is Deprecated: please use services/action.StatusState
type ActionStatusState = action.StatusState

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

// ActionStatus is Deprecated: please use services/action.Status
type ActionStatus = action.Status

// ActionTriggerResponse is Deprecated: please use services/action.TriggerResponse
type ActionTriggerResponse = action.TriggerResponse

// ActionError is Deprecated: please use services/action.Error
type ActionError = action.Error

// ActionNotificationKind is Deprecated: please use services/action.NotificationKind
type ActionNotificationKind = action.NotificationKind

const (
	//SplunkEventKind is Deprecated: please use services/action.SplunkEventKind
	SplunkEventKind = action.SplunkEventKind
	//RawJSONPayloadKind is Deprecated: please use services/action.RawJSONPayloadKind
	RawJSONPayloadKind = action.RawJSONPayloadKind
)

// ActionNotification is Deprecated: please use services/action.Notification
type ActionNotification = action.Notification

// ActionPayload is Deprecated: please use services/action.Payload
type ActionPayload = action.Payload

// RawJSONPayload is Deprecated: please use services/action.RawJSONPayload
type RawJSONPayload = action.RawJSONPayload

// SplunkEventPayload is Deprecated: please use services/action.SplunkEventPayload
type SplunkEventPayload = action.SplunkEventPayload

//
// Deprecated: Catalog models
//

// DatasetInfoKind is Deprecated: please use services/catalog.DatasetInfoKind
type DatasetInfoKind = catalog.DatasetInfoKind

const (
	// LOOKUP is Deprecated: please use services/catalog.Lookup
	LOOKUP DatasetInfoKind = catalog.Lookup
	// KVCOLLECTION is Deprecated: please use services/catalog.KvCollection
	KVCOLLECTION DatasetInfoKind = catalog.KvCollection
	// INDEX is Deprecated: please use services/catalog.Index
	INDEX DatasetInfoKind = catalog.Index
)

// DatasetInfo is Deprecated: please use services/catalog.DatasetInfo
type DatasetInfo = catalog.DatasetInfo

// UpdateDatasetInfoFields is Deprecated: please use services/catalog.UpdateDatasetInfoFields
type UpdateDatasetInfoFields = catalog.UpdateDatasetInfoFields

// Field is Deprecated: please use services/catalog.Field
type Field = catalog.Field

// PrevalenceType is Deprecated: please use services/catalog.PrevalenceType
type PrevalenceType = catalog.PrevalenceType

const (
	// ALL is Deprecated: please use services/catalog.All
	ALL PrevalenceType = catalog.All
	// SOME is Deprecated: please use services/catalog.Some
	SOME PrevalenceType = catalog.Some
	// PREVALANCEUNKNOWN is Deprecated: please use services/catalog.PrevalenceUnknown
	PREVALANCEUNKNOWN PrevalenceType = catalog.PrevalenceUnknown
)

// DataType is Deprecated: please use services/catalog.DataType
type DataType = catalog.DataType

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

// FieldType is Deprecated: please use services/catalog.FieldType
type FieldType = catalog.FieldType

const (
	// DIMENSION is Deprecated: please use services/catalog.Dimension
	DIMENSION FieldType = catalog.Dimension
	// MEASURE is Deprecated: please use services/catalog.Measure
	MEASURE FieldType = catalog.Measure
	// FIELDTYPEUNKNOWN is Deprecated: please use services/catalog.FieldTypeUnknown
	FIELDTYPEUNKNOWN = catalog.FieldTypeUnknown
)

// CatalogActionKind is Deprecated: please use services/catalog.ActionKind
type CatalogActionKind = catalog.ActionKind

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

// Rule is Deprecated: please use services/catalog.Rule
type Rule = catalog.Rule

// RuleUpdateFields is Deprecated: please use services/catalog.RuleUpdateFields
type RuleUpdateFields = catalog.RuleUpdateFields

// CatalogAction is Deprecated: please use services/catalog.Action
type CatalogAction = catalog.Action

// Module is Deprecated: please use services/catalog.Module
type Module = catalog.Module

// NewAliasAction is Deprecated: please use services/catalog.NewAliasAction
func NewAliasAction(field string, alias string, owner string) *CatalogAction {
	return catalog.NewAliasAction(field, alias, owner)
}

// NewAutoKVAction is Deprecated: please use services/catalog.NewAutoKVAction
func NewAutoKVAction(mode string, owner string) *CatalogAction {
	return catalog.NewAutoKVAction(mode, owner)
}

// NewEvalAction is Deprecated: please use services/catalog.NewEvalAction
func NewEvalAction(field string, expression string, owner string) *CatalogAction {
	return catalog.NewEvalAction(field, expression, owner)
}

// NewLookupAction is Deprecated: please use services/catalog.NewLookupAction
func NewLookupAction(expression string, owner string) *CatalogAction {
	return catalog.NewLookupAction(expression, owner)
}

// NewRegexAction is Deprecated: please use services/catalog.NewRegexAction
func NewRegexAction(field string, pattern string, limit *int, owner string) *CatalogAction {
	return catalog.NewRegexAction(field, pattern, limit, owner)
}

// NewUpdateAliasAction is Deprecated: please use services/catalog.NewUpdateAliasAction
func NewUpdateAliasAction(field *string, alias *string) *CatalogAction {
	return catalog.NewUpdateAliasAction(field, alias)
}

// NewUpdateAutoKVAction is Deprecated: please use services/catalog.NewUpdateAutoKVAction
func NewUpdateAutoKVAction(mode *string) *CatalogAction {
	return catalog.NewUpdateAutoKVAction(mode)
}

// NewUpdateEvalAction is Deprecated: please use services/catalog.NewUpdateEvalAction
func NewUpdateEvalAction(field *string, expression *string) *CatalogAction {
	return catalog.NewUpdateEvalAction(field, expression)
}

// NewUpdateLookupAction is Deprecated: please use services/catalog.NewUpdateLookupAction
func NewUpdateLookupAction(expression *string) *CatalogAction {
	return catalog.NewUpdateLookupAction(expression)
}

// NewUpdateRegexAction is Deprecated: please use services/catalog.NewUpdateRegexAction
func NewUpdateRegexAction(field *string, pattern *string, limit *int) *CatalogAction {
	return catalog.NewUpdateRegexAction(field, pattern, limit)
}

//
// Deprecated: Identity models
//

// Tenant is Deprecated: please use services/identity.Tenant
type Tenant = identity.Tenant

// ValidateInfo is Deprecated: please use services/identity.ValidateInfo
type ValidateInfo = identity.ValidateInfo

// Member is Deprecated: please use services/identity.Member
type Member = identity.Member

// Principal is Deprecated: please use services/identity.Principal
type Principal = identity.Principal

// Role is Deprecated: please use services/identity.Role
type Role = identity.Role

// Group is Deprecated: please use services/identity.Group
type Group = identity.Group

// GroupRole is Deprecated: please use services/identity.GroupRole
type GroupRole = identity.GroupRole

// GroupMember is Deprecated: please use services/identity.GroupMember
type GroupMember = identity.GroupMember

// RolePermission is Deprecated: please use services/identity.RolePermission
type RolePermission = identity.RolePermission

//
// Deprecated: Ingest models
//

// Event is Deprecated: please use services/ingest.Event
type Event = ingest.Event

// MetricEvent is Deprecated: please use services/ingest.MetricEvent
type MetricEvent = ingest.MetricEvent

// Metric is Deprecated: please use services/ingest.Metric
type Metric = ingest.Metric

// MetricAttribute is Deprecated: please use services/ingest.MetricAttribute
type MetricAttribute = ingest.MetricAttribute

//
// Deprecated: KVStore models
//

// Error is Deprecated: please use services/kvstore.Error
type Error = kvstore.Error

// AuthError is Deprecated: please use services/kvstore.AuthError
type AuthError = kvstore.AuthError

// PingOKBody is Deprecated: please use services/kvstore.PingOKBody
type PingOKBody = kvstore.PingOKBody

// PingOKBodyStatus is Deprecated: please use services/kvstore.PingOKBodyStatus
type PingOKBodyStatus = kvstore.PingOKBodyStatus

const (
	// PingOKBodyStatusHealthy is Deprecated: please use services/kvstore.PingOKBodyStatusHealthy
	PingOKBodyStatusHealthy PingOKBodyStatus = kvstore.PingOKBodyStatusHealthy

	// PingOKBodyStatusUnhealthy is Deprecated: please use services/kvstore.PingOKBodyStatusUnhealthy
	PingOKBodyStatusUnhealthy PingOKBodyStatus = kvstore.PingOKBodyStatusUnhealthy

	// PingOKBodyStatusUnknown is Deprecated: please use services/kvstore.PingOKBodyStatusUnknown
	PingOKBodyStatusUnknown PingOKBodyStatus = kvstore.PingOKBodyStatusUnknown
)

// IndexFieldDefinition is Deprecated: please use services/kvstore.IndexFieldDefinition
type IndexFieldDefinition = kvstore.IndexFieldDefinition

// IndexDefinition is Deprecated: please use services/kvstore.IndexDefinition
type IndexDefinition = kvstore.IndexDefinition

// IndexDescription is Deprecated: please use services/kvstore.IndexDescription
type IndexDescription = kvstore.IndexDescription

// LookupValue is Deprecated: please use services/kvstore.LookupValue
type LookupValue = kvstore.LookupValue

// Key is Deprecated: please use services/kvstore.Key
type Key = kvstore.Key

// Record is Deprecated: please use services/kvstore.Record
type Record = kvstore.Record

//
// Deprecated: Search models
//

// CreateJobRequest is Deprecated: please use services/search.CreateJobRequest
type CreateJobRequest = search.CreateJobRequest

// QueryParameters is Deprecated: please use services/search.QueryParameters
type QueryParameters = search.QueryParameters

// SearchJobStatus is Deprecated: please use services/search.JobStatus
type SearchJobStatus = search.JobStatus

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

// SearchJob is Deprecated: please use services/search.Job
type SearchJob = search.Job

// JobStatus is Deprecated: please use services/search.JobAction
type JobStatus = search.JobAction

const (
	// JobCanceled is Deprecated: please use services/search.JobCanceled
	JobCanceled JobStatus = search.JobCanceled
	// JobFinalized is Deprecated: please use services/search.JobFinalized
	JobFinalized JobStatus = search.JobFinalized
)

// JobMessageType is Deprecated: please use services/search.JobMessageType
type JobMessageType = search.JobMessageType

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

// SearchJobMessages is Deprecated: please use services/search.JobMessages
type SearchJobMessages = search.JobMessages

// PatchJobResponse is Deprecated: please use services/search.PatchJobResponse
type PatchJobResponse = search.PatchJobResponse

// JobResultsParams is Deprecated: please use services/search.JobResultsParams
type JobResultsParams = search.JobResultsParams

// SearchResults is Deprecated: please use services/search.Results
type SearchResults = search.Results

// ResultsNotReadyResponse is Deprecated: please use services/search.ResultsNotReadyResponse
type ResultsNotReadyResponse = search.ResultsNotReadyResponse

//
// Deprecated: Streams models
//

// ActivatePipelineRequest is Deprecated: please use services/streams.ActivatePipelineRequest
type ActivatePipelineRequest = streams.ActivatePipelineRequest

// AdditionalProperties is Deprecated: please use services/streams.AdditionalProperties
type AdditionalProperties = streams.AdditionalProperties

// DslCompilationRequest is Deprecated: please use services/streams.DslCompilationRequest
type DslCompilationRequest = streams.DslCompilationRequest

// Pipeline is Deprecated: please use services/streams.Pipeline
type Pipeline = streams.Pipeline

// PaginatedPipelineResponse is Deprecated: please use services/streams.PaginatedPipelineResponse
type PaginatedPipelineResponse = streams.PaginatedPipelineResponse

// PipelineDeleteResponse is Deprecated: please use services/streams.PipelineDeleteResponse
type PipelineDeleteResponse = streams.PipelineDeleteResponse

// PipelineQueryParams is Deprecated: please use services/streams.PipelineQueryParams
type PipelineQueryParams = streams.PipelineQueryParams

// PipelineRequest is Deprecated: please use services/streams.PipelineRequest
type PipelineRequest = streams.PipelineRequest

// PipelineStatus is Deprecated: please use services/streams.PipelineStatus
type PipelineStatus = streams.PipelineStatus

const (
	// Created is Deprecated: please use services/streams.Created
	Created PipelineStatus = streams.Created
	// Activated is Deprecated: please use services/streams.Activated
	Activated PipelineStatus = streams.Activated
)

// UplPipeline is Deprecated: please use services/streams.UplPipeline
type UplPipeline = streams.UplPipeline

// UplNode is Deprecated: please use services/streams.UplNode
type UplNode = streams.UplNode

// UplEdge is Deprecated: please use services/streams.UplEdge
type UplEdge = streams.UplEdge
