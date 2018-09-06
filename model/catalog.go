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

// DatasetInfo represents the sources of data that can be serched by Splunk
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

// PartialDatasetInfo represents the sources of data that can be updated by Splunk, same structure as DatasetInfo
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

// Rule represents a rule for transforming results at search time.
// A rule consits of a `match` clause and a collection of transformation actions
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

// Rule represents a rule for transforming results at search time.
type PartialRule struct {
	Name       string          `json:"name,omitempty" `
	Module     string          `json:"module,omitempty"`
	Match      string          `json:"match,omitempty" `
	Owner      string          `json:"owner,omitempty" `
	ModifiedBy string          `json:"modifiedBy,omitempty"`
	Version    int             `json:"version,omitempty"`
}

// CatalogAction represents a specific search time transformation action.
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
	Mode       string          `json:"mode,omitempty"`
	Expression string            `json:"expression,omitempty"`
	Pattern    string            `json:"pattern,omitempty"`
	Limit      int               `json:"limit,omitempty"`
}

//// AutoMode enumerates the automatic key/value extraction modes.
//// One of "NONE", "AUTO", "MULTIKV", "XML", "JSON".
//type AutoMode string
//
//const (
//	// NONE Automode
//	NONE AutoMode = "NONE"
//	// AUTO Automode
//	AUTO AutoMode = "AUTO"
//	// MULTIKV Automode
//	MULTIKV AutoMode = "MULTIKV"
//	// XML Automode
//	XML AutoMode = "XML"
//	// JSON Automode
//	JSON AutoMode = "JSON"
//)

// DatasetImportPayload represents the dataset import payload
type DatasetImportPayload struct {
	Modeule string `json:"module"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
}


// NewAliasAction creates a new alias kind action
func NewAliasAction(field string, alias string,  owner string) *CatalogAction {
	return &CatalogAction{
		Kind:"ALIAS",
		Owner:owner,
		Alias:alias,
		Field:field,
	}
}

// NewAutoKVAction creates a new autokv kind action
func NewAutoKVAction(mode string, owner string) *CatalogAction {
	return &CatalogAction{
		Kind:"AUTOKV",
		Owner:owner,
		Mode:mode,
	}
}

// NewEvalAction creates a new autokv kind action
func NewEvalAction(field string, expression string, owner string) *CatalogAction {
	return &CatalogAction{
		Kind:"EVAL",
		Owner:owner,
		Field:field,
		Expression:expression,
	}
}

// NewLookupAction creates a new autokv kind action
func NewLookupAction(expression string, owner string) *CatalogAction {
	return &CatalogAction{
		Kind:"LOOKUP",
		Owner:owner,
		Expression:expression,
	}
}

// NewLookupAction creates a new autokv kind action
func NewRegexAction(field string, pattern string, limit *int, owner string) *CatalogAction {
	action := CatalogAction{
		Kind:"REGEX",
		Owner:owner,
		Field:field,
		Pattern:pattern,
	}

	if limit!=nil {
		action.Limit=*limit
	}

	return &action
}