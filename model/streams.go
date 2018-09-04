package model

// ActivatePipelineRequest contains the request to activate the pipeline
type ActivatePipelineRequest struct {
	IDs []string `json:"ids"`
}

// Pipeline defines a pipeline object
type Pipeline struct {
	ActivatedDate            int64          `json:"ActivatedDate"`
	ActivatedUserID          string         `json:"ActivatedUserId"`
	ActivatedVersion         int64          `json:"ActivatedVersion"`
	CreateDate               int64          `json:"CreateDate"`
	CreateUserID             string         `json:"CreateUserId"`
	CurrentVersion           int64          `json:"CurrentVersion"`
	Data                     UplPipeline    `json:"Data"`
	Description              string         `json:"Description"`
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

// PipelineStatus reflects the status of the pipeline
type PipelineStatus string

const (
	// Created status
	Created ActionStatusState = "CREATED"
	// Activated status
	Activated ActionStatusState = "ACTIVATED"
)

// UplPipeline TODO Description
type UplPipeline struct {
	Edges    []UplEdge `json:"edges"`
	Nodes    []UplNode `json:"nodes"`
	RootNode []string  `json:"root-node"`
	Version  int32     `json:"version"`
}

// UplNode TODO Description
type UplNode struct {
	Attributes struct{} `json:"attributes"`
	ID         string   `json:"id"`
	Op         string   `json:"op"`
}

// UplEdge TODO Description
type UplEdge struct {
	Attributes interface{} `json:"attributes"`
	SourceNode string      `json:"sourceNode"`
	SourcePort string      `json:"sourcePort"`
	TargetNode string      `json:"targetNode"`
	TargetPort string      `json:"targetPort"`
}

// PaginatedPipelineResponse contains the pipeline response
type PaginatedPipelineResponse struct {
	Items []Pipeline `json:"items"`
	Total int64      `json:"total"`
}

// PipelineDeleteResponse contains the response returned as a result of delete pipeline call
type PipelineDeleteResponse struct {
	CouldDeactivate bool `json:"couldDeactivate"`
	Running         bool `json:"running"`
}

// PipelineRequest contains the pipeline data
type PipelineRequest struct {
	BypassValidation         bool        `json:"bypassValidation"`
	CreateUserID             string      `json:"createUserId"`
	Data                     UplPipeline `json:"data"`
	Description              string      `json:"description"`
	Name                     string      `json:"name"`
	StreamingConfigurationID int64       `json:"streamingConfigurationId"`
}

// AdditionalProperties contain the properties in an activate/deactivate response
type AdditionalProperties map[string][]string
