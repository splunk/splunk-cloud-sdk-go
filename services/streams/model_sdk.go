package streams

// PipelineStatusQueryParams contains the query parameters that can be provided by the user to fetch specific pipeline job statuses
type PipelineStatusQueryParams struct {
	Offset       *int32  `json:"offset,omitempty"`
	PageSize     *int32  `json:"pageSize,omitempty"`
	SortField    *string `json:"sortField,omitempty"`
	SortDir      *string `json:"sortDir,omitempty"`
	Activated    *bool   `json:"activated,omitempty"`
	CreateUserID *string `json:"createUserId,omitempty"`
	Name         *string `json:"name,omitempty"`
}

// AdditionalProperties contain the properties in an activate/deactivate response
type AdditionalProperties map[string][]string

// PipelineQueryParams contains the query parameters that can be provided by the user to fetch specific pipelines
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

// PartialTemplateRequest contains the template request data for partial update operation
type PartialTemplateRequest struct {
	Data        *Pipeline `json:"data,omitempty"`
	Description *string   `json:"description,omitempty"`
	Name        *string   `json:"name,omitempty"`
}
