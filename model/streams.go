package model

// ActivatePipelineRequest contains the request to activate the pipeline
type ActivatePipelineRequest struct {
	IDs []string `json:"ids"`
}

// Pipeline defines a pipeline object
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

// PipelineStatus reflects the status of the pipeline
type PipelineStatus string

const (
	// Created status
	Created PipelineStatus = "CREATED"
	// Activated status
	Activated PipelineStatus = "ACTIVATED"
)

// UplPipeline TODO Description
type UplPipeline struct {
	Edges    []UplEdge `json:"edges"`
	Nodes    []UplNode `json:"nodes"`
	RootNode []string  `json:"root-node"`
	Version  int32     `json:"version"`
}

// UplNode defines the nodes forming a pipeline
type UplNode interface {
}

// UplNodeCommon contains fields common to all the node implementations
type UplNodeCommon struct {
	Attributes map[string]interface{} `json:"attributes"`
	ID         string                 `json:"id"`
	Op         string                 `json:"op"`
}

/*func main() {
	t1 := type1{msg: "Hey! I'm type1", gen2: gen2{x: "12", y: "13"}}

	tl := []generic{t1}

	for _, i := range tl {
		switch v := i.(type) {
		case type1:
			fmt.Println(v.msg)
			//case type2:
			//fmt.Println(v.say())
			// fmt.Println(v.msg) Uncomment to see that type2 has no msg attribute.
		}
	}
}*/

// KafkaReader contains fields of the read-kafka pipeline node
type KafkaReader struct {
	UplNodeCommon
	Brokers            string   `json:"brokers"`
	Topic              string   `json:"topic"`
	ConsumerProperties struct{} `json:"consumer-properties,omitempty"` // TODO: Take opinion on this. Is struct{} a good choice?
}

// KafkaWriter contains fields of the write-kafka pipeline node
type KafkaWriter struct {
	UplNodeCommon
	Brokers string `json:"brokers"`
	// Input map[string]interface{} `json:"input"`
	ProducerProperties struct{} `json:"producer-properties"`
}

// SerializeEvents contains fields of the serialize-events pipeline node
type SerializeEvents struct {
	UplNodeCommon
	Topic string `json:"topic"`
}

// UplEdge TODO Description
type UplEdge struct {
	Attributes map[string]interface{} `json:"attributes"`
	SourceNode string                 `json:"sourceNode"`
	SourcePort string                 `json:"sourcePort"`
	TargetNode string                 `json:"targetNode"`
	TargetPort string                 `json:"targetPort"`
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
	StreamingConfigurationID *int64      `json:"streamingConfigurationId,omitempty"`
}

// AdditionalProperties contain the properties in an activate/deactivate response
type AdditionalProperties map[string][]string

// TODO: Include other BLAM nodes as base child of UplNode after initial code review.
