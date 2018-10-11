package model

import (
	"github.com/splunk/splunk-cloud-sdk-go/services/streams"
)

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
