// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package catalog

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

// DatasetCreationPayload represents the sources of data that can be searched by Splunk
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

// UpdateDatasetInfoFields represents the sources of data that can be updated by Splunk, same structure as DatasetInfo
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

// ActionKind enumerates the kinds of search time transformation action known by the service.
type ActionKind string

const (
	// ALIAS action
	ALIAS ActionKind = "ALIAS"
	// AUTOKV action
	AUTOKV ActionKind = "AUTOKV"
	// REGEX action
	REGEX ActionKind = "REGEX"
	// EVAL action
	EVAL ActionKind = "EVAL"
	// LOOKUPACTION action
	LOOKUPACTION ActionKind = "LOOKUP"
)

// Rule represents a rule for transforming results at search time.
// A rule consists of a `match` clause and a collection of transformation actions
type Rule struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name"`
	Module     string   `json:"module,omitempty"`
	Match      string   `json:"match"`
	Actions    []Action `json:"actions,omitempty"`
	Owner      string   `json:"owner,omitempty"`
	Created    string   `json:"created,omitempty"`
	Modified   string   `json:"modified,omitempty"`
	CreatedBy  string   `json:"createdBy,omitempty"`
	ModifiedBy string   `json:"modifiedBy,omitempty"`
	Version    int      `json:"version,omitempty"`
}

// RuleUpdateFields represents the set of rule properties that can be updated
type RuleUpdateFields struct {
	Name    string `json:"name,omitempty"`
	Module  string `json:"module,omitempty"`
	Match   string `json:"match,omitempty"`
	Owner   string `json:"owner,omitempty"`
	Version int    `json:"version,omitempty"`
}

// Action represents a specific search time transformation action.
// This struct should NOT be directly used to construct object, use the NewXXXAction() instead
type Action struct {
	ID         string     `json:"id,omitempty"`
	RuleID     string     `json:"ruleid,omitempty"`
	Kind       ActionKind `json:"kind,omitempty"`
	Owner      string     `json:"owner,omitempty"`
	Created    string     `json:"created,omitempty"`
	Modified   string     `json:"modified,omitempty"`
	CreatedBy  string     `json:"createdBy,omitempty"`
	ModifiedBy string     `json:"modifiedBy,omitempty"`
	Version    int        `json:"version,omitempty"`
	Field      string     `json:"field,omitempty"`
	Alias      string     `json:"alias,omitempty"`
	Mode       string     `json:"mode,omitempty"`
	Expression string     `json:"expression,omitempty"`
	Pattern    string     `json:"pattern,omitempty"`
	Limit      *int       `json:"limit,omitempty"`
}

// Module represents catalog module
type Module struct {
	Name string `json:"name"`
}

// NewAliasAction creates a new alias kind action
func NewAliasAction(field string, alias string, owner string) *Action {
	return &Action{
		Kind:  "ALIAS",
		Owner: owner,
		Alias: alias,
		Field: field,
	}
}

// NewAutoKVAction creates a new autokv kind action
func NewAutoKVAction(mode string, owner string) *Action {
	return &Action{
		Kind:  "AUTOKV",
		Owner: owner,
		Mode:  mode,
	}
}

// NewEvalAction creates a new eval kind action
func NewEvalAction(field string, expression string, owner string) *Action {
	return &Action{
		Kind:       "EVAL",
		Owner:      owner,
		Field:      field,
		Expression: expression,
	}
}

// NewLookupAction creates a new lookup kind action
func NewLookupAction(expression string, owner string) *Action {
	return &Action{
		Kind:       "LOOKUP",
		Owner:      owner,
		Expression: expression,
	}
}

// NewRegexAction creates a new regex kind action
func NewRegexAction(field string, pattern string, limit *int, owner string) *Action {
	action := Action{
		Kind:    "REGEX",
		Owner:   owner,
		Field:   field,
		Pattern: pattern,
		Limit:   limit,
	}

	return &action
}

// NewUpdateAliasAction updates an existing alias kind action
func NewUpdateAliasAction(field *string, alias *string) *Action {
	res := Action{}

	if field != nil {
		res.Field = *field
	}

	if alias != nil {
		res.Alias = *alias
	}

	return &res
}

// NewUpdateAutoKVAction updates an existing autokv kind action
func NewUpdateAutoKVAction(mode *string) *Action {
	res := Action{}

	if mode != nil {
		res.Mode = *mode
	}

	return &res

}

// NewUpdateEvalAction updates an existing eval kind action
func NewUpdateEvalAction(field *string, expression *string) *Action {
	res := Action{}

	if field != nil {
		res.Field = *field
	}

	if expression != nil {
		res.Alias = *expression
	}

	return &res

}

// NewUpdateLookupAction updates an existing lookup kind action
func NewUpdateLookupAction(expression *string) *Action {
	res := Action{}

	if expression != nil {
		res.Expression = *expression
	}

	return &res
}

// NewUpdateRegexAction updates an existing regex kind action
func NewUpdateRegexAction(field *string, pattern *string, limit *int) *Action {
	res := Action{}

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
