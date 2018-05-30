package model

// DatasetInfoKind enumerates the kinds of datasets known to the system.
type DatasetInfoKind string

const (
	// LOOKUP represents // TODO: Description needed
	LOOKUP DatasetInfoKind = "lookup"
	// INDEX represents a Splunk events or metrics index
	INDEX DatasetInfoKind = "index"
)

// DatasetInfo represents the sources of data that can be serched by Splunk
type DatasetInfo struct {
	ID           string          `json:"id,omitempty"`
	Name         string          `json:"name" binding:"required"`
	Kind         DatasetInfoKind `json:"kind" binding:"required"`
	Owner        string          `json:"owner" binding:"required"`
	Created      string          `json:"created,omitempty"`
	Modified     string          `json:"modified,omitempty"`
	CreatedBy    string          `json:"createdBy,omitempty"`
	ModifiedBy   string          `json:"modifiedBy,omitempty"`
	Capabilities string          `json:"capabilities" binding:"required"`
	Version      int             `json:"version,omitempty"`
	Fields       []Field         `json:"fields,omitempty"`

	ExternalKind       string `json:"externalKind,omitempty"`
	ExternalName       string `json:"externalName,omitempty"`
	CaseSensitiveMatch bool   `json:"caseSensitiveMatch,omitempty"`
	Filter             string `json:"filter,omitempty"`
	MaxMatches         int    `json:"maxMatches,omitempty"`
	MinMatches         int    `json:"minMatches,omitempty"`
	defaultMatch       string `json:"defaultMatch,omitempty"`

	Datatype string `json:"datatype,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

/*// Lookup represents LOOKUP DatasetInfoKind
type Lookup struct {
	DatasetInfoKind    string `json:"datasetInfoKind"`
	ExternalKind       string `json:"externalKind"`
	ExternalName       string `json:"externalName"`
	CaseSensitiveMatch bool   `json:"caseSensitiveMatch"`
	Filter             string `json:"filter"`
	MaxMatches         int    `json:"maxMatches"`
	MinMatches         int    `json:"minMatches"`
	defaultMatch       string `json:"defaultMatch"`
}

// Index represents INDEX DatasetInfoKind
type Index struct {
	Datatype string `json:"datatype"`
	Disabled bool   `json:"disabled"`
}*/

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

	ExternalKind       string `json:"externalKind,omitempty"`
	ExternalName       string `json:"externalName,omitempty"`
	CaseSensitiveMatch bool   `json:"caseSensitiveMatch,omitempty"`
	Filter             string `json:"filter,omitempty"`
	MaxMatches         int    `json:"maxMatches,omitempty"`
	MinMatches         int    `json:"minMatches,omitempty"`
	defaultMatch       string `json:"defaultMatch,omitempty"`

	Datatype string `json:"datatype,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

// Field represents TODO: Description needed
type Field struct {
	ID             string         `json:"id" binding:"required"`
	Name           string         `json:"name" binding:"required"`
	DatasetID      string         `json:"datasetId" binding:"required"` // TODO: Further split
	DataType       DataType       `json:"dataType"`
	FieldType      FieldType      `json:"fieldType"`
	Prevalence     PrevelanceType `json:"prevalence"` // TODO: Further split
	Created        string         `json:"created"`
	Modified       string         `json:"modified"`
	VersionAdded   int            `json:"versionAdded"`
	VersionRemoved int            `json:"versionRemoved"`
	Fields         []Field        `json:"fields"`
	Dataset        DatasetInfo    `json:"dataset"`
}

// PrevelanceType enumerates the types of prevalance used in fields.
type PrevelanceType string

const (
	// ALL PrevelanceType
	ALL PrevelanceType = "ALL"
	// SOME PrevelanceType
	SOME PrevelanceType = "SOME"
	// PREVELANCEUNKNOWN PrevelanceType
	PREVELANCEUNKNOWN PrevelanceType = "UNKNOWN"
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
// A rule consits of a `match` clause and a collection of transformation actions
type Rule struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name"`
	Module     string   `json:"module,omitempty"`
	Match      string   `json:"match"`
	Actions    []Action `json:"actions,omitempty"`
	Owner      string   `json:"owner"`
	Created    string   `json:"created,omitempty"`
	Modified   string   `json:"modified,omitempty"`
	CreatedBy  string   `json:"createdBy,omitempty"`
	ModifiedBy string   `json:"modifiedBy,omitempty"`
	Version    int      `json:"version,omitempty"`
}

// Action represents a specific search time transformation action.
type Action struct {
	ID         string     `json:"id"`
	RuleID     string     `json:"ruleid" binding:"required"`
	Kind       ActionKind `json:"kind" binding:"required"`
	Owner      string     `json:"owner" binding:"required"`
	Created    string     `json:"created"`
	Modified   string     `json:"modified"`
	CreatedBy  string     `json:"createdBy"`
	ModifiedBy string     `json:"modifiedBy"`
	Version    int        `json:"version"`
	Field      string     `json:"field,omitempty"`
	Alias      string     `json:"alias,omitempty"`
	Mode       AutoMode   `json:"mode,omitempty"`
	Expression string     `json:"expression,omitempty"`
	Pattern    string     `json:"pattern,omitempty"`
	Limit      int        `json:"limit,omitempty"`
}

// AutoMode enumerates the automatic key/value extraction modes.
// One of "NONE", "AUTO", "MULTIKV", "XML", "JSON".
type AutoMode string

const (
	// NONE Automode
	NONE AutoMode = "NONE"
	// AUTO Automode
	AUTO AutoMode = "AUTO"
	// MULTIKV Automode
	MULTIKV AutoMode = "MULTIKV"
	// XML Automode
	XML AutoMode = "XML"
	// JSON Automode
	JSON AutoMode = "JSON"
)
