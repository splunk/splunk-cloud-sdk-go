// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model

// DatasetInfoKind enumerates the kinds of datasets known to the system.
type DatasetInfoKind string

const (
	// LOOKUP represents TODO: Description needed
	LOOKUP DatasetInfoKind = "lookup"
	// KVCOLLECTION represents a key value store, it is used with the kvstore service, but its implementation is separate of kvstore
	KVCOLLECTION DatasetInfoKind = "kvcollection"
	// INDEX represents a Splunk events or metrics index
	INDEX DatasetInfoKind = "index"
)

// DatasetInfo represents the sources of data that can be searched by Splunk
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
	Disabled bool   `json:"disabled"`
}

// DatasetCreationRestricted represents the possible creation fields for datasets - datasets
// are sources of data that can be searched by Splunk
type DatasetCreationRestricted struct {
	ID           string          `json:"id,omitempty"`
	Name         string          `json:"name"`
	Kind         DatasetInfoKind `json:"kind"`
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

// DatasetCreationPayload represents the sources of data that can be searched by Splunk
// TODO: Remove me in favor of DatasetCreationRestricted
type DatasetCreationPayload struct {
	ID           string          `json:"id,omitempty"`
	Name         string          `json:"name"`
	Kind         DatasetInfoKind `json:"kind"`
	Owner        string          `json:"owner"`
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

// UpdateDatasetRestricted represents the updateable fields for datasets - datasets
// are sources of data that can be searched by Splunk
type UpdateDatasetRestricted struct {
	Name         string   `json:"name,omitempty"`
	Owner        string   `json:"owner,omitempty"`
	Capabilities string   `json:"capabilities,omitempty"`
	Version      int      `json:"version,omitempty"`
	Readroles    []string `json:"readroles,omitempty"`
	Writeroles   []string `json:"writeroles,omitempty"`

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

// UpdateDatasetInfoFields represents the updateable fields for datasets - datasets
// are sources of data that can be searched by Splunk
// TODO: Remove me in favor of UpdateDatasetRestricted
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

// Field represents the fields belonging to the specified Dataset
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

// PrevalenceType enumerates the types of prevalance used in fields.
type PrevalenceType string

const (
	// ALL PrevalenceType
	ALL PrevalenceType = "ALL"
	// SOME PrevalenceType
	SOME PrevalenceType = "SOME"
	// PREVALANCEUNKNOWN PrevalenceType
	PREVALANCEUNKNOWN PrevalenceType = "UNKNOWN"
)

// DataType enumerates the kinds of datatypes used in fields.
type DataType string

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

// FieldType enumerates different kinds of fields.
type FieldType string

const (
	// DIMENSION fieldType
	DIMENSION FieldType = "DIMENSION"
	// MEASURE fieldType
	MEASURE FieldType = "MEASURE"
	// FIELDTYPEUNKNOWN fieldType
	FIELDTYPEUNKNOWN FieldType = "UNKNOWN"
)

// CatalogActionKind enumerates the kinds of search time transformation action known by the service.
type CatalogActionKind string

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

// RuleCreationPayload for creating a Rule. Rules are for transforming results at search time.
// A rule consists of a `match` clause and a collection of transformation actions
type RuleCreationPayload struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name"`
	Module string `json:"module,omitempty"`
	Match  string `json:"match"`
	// TODO: Change this to []CatalogActionCreationPayload when RuleCreationPayload when the service
	// is updated to use CreateRule(rule model.RuleCreationPayload)
	Actions []CatalogAction `json:"actions,omitempty"`
	Version int             `json:"version,omitempty"`
}

// Rule for getting Rule information. Rules are for transforming results at search time.
// A rule consists of a `match` clause and a collection of transformation actions
type Rule struct {
	ID         string          `json:"id,omitempty"`
	Name       string          `json:"name"`
	Module     string          `json:"module,omitempty"`
	Match      string          `json:"match"`
	Actions    []CatalogAction `json:"actions,omitempty"`
	Owner      string          `json:"owner"`
	Created    string          `json:"created,omitempty"`
	Modified   string          `json:"modified,omitempty"`
	CreatedBy  string          `json:"createdBy,omitempty"`
	ModifiedBy string          `json:"modifiedBy,omitempty"`
	Version    int             `json:"version,omitempty"`
}

// RuleUpdateFields represents the set of rule properties that can be updated
type RuleUpdateFields struct {
	Name    string `json:"name,omitempty"`
	Module  string `json:"module,omitempty"`
	Match   string `json:"match,omitempty"`
	Owner   string `json:"owner,omitempty"`
	Version int    `json:"version,omitempty"`
}

// CatalogActionCreationRestricted for creating catalog actions. An action represents
// a specific search time transformation action.
type CatalogActionCreationRestricted struct {
	Kind       CatalogActionKind `json:"kind"`
	Version    int               `json:"version,omitempty"`
	Field      string            `json:"field,omitempty"`
	Alias      string            `json:"alias,omitempty"`
	Mode       string            `json:"mode,omitempty"`
	Expression string            `json:"expression,omitempty"`
	Pattern    string            `json:"pattern,omitempty"`
	Limit      *int              `json:"limit,omitempty"`
}

// CatalogActionCreationPayload for creating catalog actions. An action represents
// a specific search time transformation action.
// TODO: Remove me in favor of CatalogActionCreationRestricted
type CatalogActionCreationPayload struct {
	RuleID     string            `json:"ruleid,omitempty"`
	Kind       CatalogActionKind `json:"kind" `
	Owner      string            `json:"owner"`
	Field      string            `json:"field,omitempty"`
	Alias      string            `json:"alias,omitempty"`
	Mode       string            `json:"mode,omitempty"`
	Expression string            `json:"expression,omitempty"`
	Pattern    string            `json:"pattern,omitempty"`
	Limit      *int              `json:"limit,omitempty"`
}

// CatalogAction for getting catalog actions. An action represents a specific
// search time transformation action.
type CatalogAction struct {
	ID         string            `json:"id,omitempty"`
	RuleID     string            `json:"ruleid,omitempty"`
	Kind       CatalogActionKind `json:"kind"`
	Owner      string            `json:"owner"`
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

// CatalogActionUpdateFields for updating catalog actions. An action
// represents a specific search time transformation action.
type CatalogActionUpdateFields struct {
	Owner      string `json:"owner,omitempty"`
	Version    int    `json:"version,omitempty"`
	Field      string `json:"field,omitempty"`
	Alias      string `json:"alias,omitempty"`
	Mode       string `json:"mode,omitempty"`
	Expression string `json:"expression,omitempty"`
	Pattern    string `json:"pattern,omitempty"`
	Limit      *int   `json:"limit,omitempty"`
}

// NewAliasAction creates a new alias kind action
// TODO: remove owner and convert to CatalogActionCreationPayload
func NewAliasAction(field string, alias string, owner string) *CatalogAction {
	return &CatalogAction{
		Kind:  "ALIAS",
		Owner: owner,
		Alias: alias,
		Field: field,
	}
}

// NewAutoKVAction creates a new autokv kind action
// TODO: remove owner and convert to CatalogActionCreationPayload
func NewAutoKVAction(mode string, owner string) *CatalogAction {
	return &CatalogAction{
		Kind:  "AUTOKV",
		Owner: owner,
		Mode:  mode,
	}
}

// NewEvalAction creates a new eval kind action
// TODO: remove owner and convert to CatalogActionCreationPayload
func NewEvalAction(field string, expression string, owner string) *CatalogAction {
	return &CatalogAction{
		Kind:       "EVAL",
		Owner:      owner,
		Field:      field,
		Expression: expression,
	}
}

// NewLookupAction creates a new lookup kind action
// TODO: remove owner and convert to CatalogActionCreationPayload
func NewLookupAction(expression string, owner string) *CatalogAction {
	return &CatalogAction{
		Kind:       "LOOKUP",
		Owner:      owner,
		Expression: expression,
	}
}

// NewRegexAction creates a new regex kind action
// TODO: remove owner and convert to CatalogActionCreationPayload
func NewRegexAction(field string, pattern string, limit *int, owner string) *CatalogAction {
	action := CatalogAction{
		Kind:    "REGEX",
		Owner:   owner,
		Field:   field,
		Pattern: pattern,
		Limit:   limit,
	}

	return &action
}

// NewUpdateAliasAction updates an existing alias kind action
// TODO: Change to return CatalogActionUpdateFields
func NewUpdateAliasAction(field *string, alias *string) *CatalogAction {
	res := CatalogAction{}

	if field != nil {
		res.Field = *field
	}

	if alias != nil {
		res.Alias = *alias
	}

	return &res
}

// NewUpdateAutoKVAction updates an existing autokv kind action
// TODO: Change to return CatalogActionUpdateFields
func NewUpdateAutoKVAction(mode *string) *CatalogAction {
	res := CatalogAction{}

	if mode != nil {
		res.Mode = *mode
	}

	return &res

}

// NewUpdateEvalAction updates an existing eval kind action
// TODO: Change to return CatalogActionUpdateFields
func NewUpdateEvalAction(field *string, expression *string) *CatalogAction {
	res := CatalogAction{}

	if field != nil {
		res.Field = *field
	}

	if expression != nil {
		res.Alias = *expression
	}

	return &res

}

// NewUpdateLookupAction updates an existing lookup kind action
// TODO: Change to return CatalogActionUpdateFields
func NewUpdateLookupAction(expression *string) *CatalogAction {
	res := CatalogAction{}

	if expression != nil {
		res.Expression = *expression
	}

	return &res
}

// NewUpdateRegexAction updates an existing regex kind action
// TODO: Change to return CatalogActionUpdateFields
func NewUpdateRegexAction(field *string, pattern *string, limit *int) *CatalogAction {
	res := CatalogAction{}

	if field != nil {
		res.Field = *field
	}

	if pattern != nil {
		res.Pattern = *pattern
	}

	if limit != nil {
		res.Limit = limit
	}

	return &res
}
