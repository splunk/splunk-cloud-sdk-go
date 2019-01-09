// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package catalog

import (
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

// DatasetInfoKind enumerates the kinds of datasets known to the system.
type DatasetInfoKind string

const (
	// Lookup allows a dataset to be joined with search results.
	Lookup DatasetInfoKind = "lookup"
	// KvCollection represents a key value store, it is used with the kvstore service, but its implementation is separate of kvstore
	KvCollection DatasetInfoKind = "kvcollection"
	// Index represents a Splunk events or metrics index
	Index DatasetInfoKind = "index"
	// Metric represents metrics dataset
	Metric DatasetInfoKind = "metric"
	// View represents a dataset defined by a search, similar to a 'saved search' in Splunk Enterprise
	View DatasetInfoKind = "view"
	// Import represents a dataset that points to an existing dataset in order to support usage in a different module.
	Import DatasetInfoKind = "import"
	// Job represents a dataset relating to a search job
	Job DatasetInfoKind = "job"
)

// Dataset represents the sources of data that can be searched by Splunk
type Dataset interface {
	// GetName the dataset name. Dataset names must be unique within each module.
	GetName() string
	// GetID returns a unique dataset ID.
	GetID() string
	// GetModule returns the name of the module associated with the dataset.
	GetModule() string
	// GetKind returns the dataset kind.
	GetKind() string
}

// ImportDataset represents a fully-constructed import dataset
type ImportDataset struct {
	// Common dataset properties:
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name,omitempty"`
	// The dataset kind.
	Kind DatasetInfoKind `json:"kind,omitempty" methods:"GET,POST"`
	// A unique dataset ID. Random ID used if not provided. Not valid for PATCH method.
	ID string `json:"id,omitempty" methods:"GET,POST"`
	// The name of the module to create the new dataset in. The default module is "".
	Module *string `json:"module,omitempty"`
	// The catalog version.
	Version int `json:"version,omitempty" methods:"GET"`
	// The date and time object was created.
	Created string `json:"created,omitempty" methods:"GET"`
	// The date and time object was modified.
	Modified string `json:"modified,omitempty" methods:"GET"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	CreatedBy string `json:"createdby,omitempty" methods:"GET"`
	// The name of the user who most recently modified the object.
	ModifiedBy string `json:"modifiedby,omitempty" methods:"GET"`
	// The name of the object's owner.
	Owner string `json:"owner,omitempty" methods:"GET,PATCH"`
	// The dataset name qualified by the module name.
	ResourceName string `json:"resourcename,omitempty" methods:"GET"`

	// Import-specific properties:
	// The dataset module being imported.
	SourceModule *string `json:"sourceModule,omitempty" methods:"GET,PATCH"`
	// The dataset name being imported.
	SourceName *string `json:"sourceName,omitempty" methods:"GET,PATCH"`
	// The dataset ID being imported.
	OriginalDatasetID *string `json:"originalDatasetId,omitempty" methods:"GET"`
	// The dataset ID being imported.
	// TODO: sourceId is only used for POST operations, and is returned as originalDatasetId.
	// This should be fixed by the Catalog service in the future.
	SourceID *string `json:"sourceId,omitempty" methods:"POST"`
}

// GetName the dataset name. Dataset names must be unique within each module.
func (ds ImportDataset) GetName() string {
	return ds.Name
}

// GetID returns a unique dataset ID. Random ID used if not provided
func (ds ImportDataset) GetID() string {
	return ds.ID
}

// GetModule returns the name of the module associated with the dataset.
func (ds ImportDataset) GetModule() string {
	return *ds.Module
}

// GetKind returns the dataset kind.
func (ds ImportDataset) GetKind() string {
	return string(ds.Kind)
}

// MarshalJSONByMethod implements the util.MethodMarshaler interface
func (ds *ImportDataset) MarshalJSONByMethod(method string) ([]byte, error) {
	return util.MarshalByMethod(*ds, method)
}

// MetricDataset represents a fully-constructed index dataset
type MetricDataset struct {
	// Common dataset properties:
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name,omitempty"`
	// The dataset kind.
	Kind DatasetInfoKind `json:"kind,omitempty" methods:"GET,POST"`
	// A unique dataset ID. Random ID used if not provided. Not valid for PATCH method.
	ID string `json:"id,omitempty" methods:"GET,POST"`
	// The name of the module to create the new dataset in. The default module is "".
	Module *string `json:"module,omitempty"`
	// The catalog version.
	Version int `json:"version,omitempty" methods:"GET"`
	// The date and time object was created.
	Created string `json:"created,omitempty" methods:"GET"`
	// The date and time object was modified.
	Modified string `json:"modified,omitempty" methods:"GET"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	CreatedBy string `json:"createdby,omitempty" methods:"GET"`
	// The name of the user who most recently modified the object.
	ModifiedBy string `json:"modifiedby,omitempty" methods:"GET"`
	// The name of the object's owner.
	Owner string `json:"owner,omitempty" methods:"GET,PATCH"`
	// The dataset name qualified by the module name.
	ResourceName string `json:"resourcename,omitempty" methods:"GET"`

	// Metric-specific properties:
	// Specifies whether or not the Splunk index is disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// The frozenTimePeriodInSecs to use for the index
	FrozenTimePeriodInSecs *int `json:"frozenTimePeriodInSecs,omitempty"`
}

// GetName the dataset name. Dataset names must be unique within each module.
func (ds MetricDataset) GetName() string {
	return ds.Name
}

// GetID returns a unique dataset ID. Random ID used if not provided
func (ds MetricDataset) GetID() string {
	return ds.ID
}

// GetModule returns the name of the module associated with the dataset.
func (ds MetricDataset) GetModule() string {
	return *ds.Module
}

// GetKind returns the dataset kind.
func (ds MetricDataset) GetKind() string {
	return string(ds.Kind)
}

// MarshalJSONByMethod implements the util.MethodMarshaler interface
func (ds *MetricDataset) MarshalJSONByMethod(method string) ([]byte, error) {
	return util.MarshalByMethod(*ds, method)
}

// IndexDataset represents a fully-constructed index dataset
type IndexDataset struct {
	// Common dataset properties:
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name,omitempty"`
	// The dataset kind.
	Kind DatasetInfoKind `json:"kind,omitempty" methods:"GET,POST"`
	// A unique dataset ID. Random ID used if not provided. Not valid for PATCH method.
	ID string `json:"id,omitempty" methods:"GET,POST"`
	// The name of the module to create the new dataset in. The default module is "".
	Module *string `json:"module,omitempty"`
	// The catalog version.
	Version int `json:"version,omitempty" methods:"GET"`
	// The date and time object was created.
	Created string `json:"created,omitempty" methods:"GET"`
	// The date and time object was modified.
	Modified string `json:"modified,omitempty" methods:"GET"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	CreatedBy string `json:"createdby,omitempty" methods:"GET"`
	// The name of the user who most recently modified the object.
	ModifiedBy string `json:"modifiedby,omitempty" methods:"GET"`
	// The name of the object's owner.
	Owner string `json:"owner,omitempty" methods:"GET,PATCH"`
	// The dataset name qualified by the module name.
	ResourceName string `json:"resourcename,omitempty" methods:"GET"`

	// Index-specific properties:
	// Specifies whether or not the Splunk index is disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// The frozenTimePeriodInSecs to use for the index
	FrozenTimePeriodInSecs *int `json:"frozenTimePeriodInSecs,omitempty"`
}

// GetName the dataset name. Dataset names must be unique within each module.
func (ds IndexDataset) GetName() string {
	return ds.Name
}

// GetID returns a unique dataset ID. Random ID used if not provided
func (ds IndexDataset) GetID() string {
	return ds.ID
}

// GetModule returns the name of the module associated with the dataset.
func (ds IndexDataset) GetModule() string {
	return *ds.Module
}

// GetKind returns the dataset kind.
func (ds IndexDataset) GetKind() string {
	return string(ds.Kind)
}

// MarshalJSONByMethod implements the util.MethodMarshaler interface
func (ds *IndexDataset) MarshalJSONByMethod(method string) ([]byte, error) {
	return util.MarshalByMethod(*ds, method)
}

// JobDataset represents a fully-constructed job dataset
// NOTE: POST is not supported for Job datasets, please use the search service to create jobs
// NOTE: only Name, Module, Owner, and Status are supported for PATCH
type JobDataset struct {
	// Common dataset properties:
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name,omitempty" methods:"GET,PATCH"`
	// The dataset kind.
	Kind DatasetInfoKind `json:"kind,omitempty" methods:"GET"`
	// A unique dataset ID. Random ID used if not provided. Not valid for PATCH method.
	ID string `json:"id,omitempty" methods:"GET"`
	// The name of the module to create the new dataset in. The default module is "".
	Module *string `json:"module,omitempty" methods:"GET,PATCH"`
	// The catalog version.
	Version int `json:"version,omitempty" methods:"GET"`
	// The date and time object was created.
	Created string `json:"created,omitempty" methods:"GET"`
	// The date and time object was modified.
	Modified string `json:"modified,omitempty" methods:"GET"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	CreatedBy string `json:"createdby,omitempty" methods:"GET"`
	// The name of the user who most recently modified the object.
	ModifiedBy string `json:"modifiedby,omitempty" methods:"GET"`
	// The name of the object's owner.
	Owner string `json:"owner,omitempty" methods:"GET,PATCH"`
	// The dataset name qualified by the module name.
	ResourceName string `json:"resourcename,omitempty" methods:"GET"`

	// Job-specific properties:
	// The time the dataset will be available.
	DeleteTime *string `json:"deleteTime,omitempty" methods:"GET"`
	// Should the search produce all fields (including those not explicity mentioned in the SPL)?
	ExtractAllFields *bool `json:"extractAllFields,omitempty" methods:"GET"`
	// The number of seconds to run this search before finishing.
	MaxTime *int `json:"maxTime,omitempty" methods:"GET"`
	// Parameters for the search job, mainly earliest and latest.
	Parameters interface{} `json:"parameters,omitempty" methods:"GET"`
	// An estimate of how complete the search job is.
	PercentComplete *int `json:"percentComplete,omitempty" methods:"GET"`
	// The SPL query string for the search job.
	Query *string `json:"query,omitempty" methods:"GET"`
	// The instantaneous number of results produced by the search job.
	ResultsAvailable *int `json:"resultsAvailable,omitempty" methods:"GET"`
	// The ID assigned to the search job.
	SID *string `json:"sid,omitempty" methods:"GET"`
	// The SPLv2 version of the search job query string.
	SPL *string `json:"spl,omitempty" methods:"GET"`
	// The current status of the search job.
	Status *string `json:"status,omitempty" methods:"GET,PATCH"`
	// Converts a formatted time string from into UTC seconds.
	TimeFormat *string `json:"timeFormat,omitempty" methods:"GET"`
	// The system time at the time the search job was created
	TimeOfSearch *string `json:"timeOfSearch,omitempty" methods:"GET"`
}

// GetName the dataset name. Dataset names must be unique within each module.
func (ds JobDataset) GetName() string {
	return ds.Name
}

// GetID returns a unique dataset ID. Random ID used if not provided
func (ds JobDataset) GetID() string {
	return ds.ID
}

// GetModule returns the name of the module associated with the dataset.
func (ds JobDataset) GetModule() string {
	return *ds.Module
}

// GetKind returns the dataset kind.
func (ds JobDataset) GetKind() string {
	return string(ds.Kind)
}

// MarshalJSONByMethod implements the util.MethodMarshaler interface
func (ds *JobDataset) MarshalJSONByMethod(method string) ([]byte, error) {
	return util.MarshalByMethod(*ds, method)
}

// ViewDataset represents a fully-constructed view dataset
type ViewDataset struct {
	// Common dataset properties:
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name,omitempty"`
	// The dataset kind.
	Kind DatasetInfoKind `json:"kind,omitempty" methods:"GET,POST"`
	// A unique dataset ID. Random ID used if not provided. Not valid for PATCH method.
	ID string `json:"id,omitempty" methods:"GET,POST"`
	// The name of the module to create the new dataset in. The default module is "".
	Module *string `json:"module,omitempty"`
	// The catalog version.
	Version int `json:"version,omitempty" methods:"GET"`
	// The date and time object was created.
	Created string `json:"created,omitempty" methods:"GET"`
	// The date and time object was modified.
	Modified string `json:"modified,omitempty" methods:"GET"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	CreatedBy string `json:"createdby,omitempty" methods:"GET"`
	// The name of the user who most recently modified the object.
	ModifiedBy string `json:"modifiedby,omitempty" methods:"GET"`
	// The name of the object's owner.
	Owner string `json:"owner,omitempty" methods:"GET,PATCH"`
	// The dataset name qualified by the module name.
	ResourceName string `json:"resourcename,omitempty" methods:"GET"`

	// View-specific properties:
	// A valid SPL-defined search.
	Search *string `json:"search,omitempty"`
}

// GetName the dataset name. Dataset names must be unique within each module.
func (ds ViewDataset) GetName() string {
	return ds.Name
}

// GetID returns a unique dataset ID. Random ID used if not provided
func (ds ViewDataset) GetID() string {
	return ds.ID
}

// GetModule returns the name of the module associated with the dataset.
func (ds ViewDataset) GetModule() string {
	return *ds.Module
}

// GetKind returns the dataset kind.
func (ds ViewDataset) GetKind() string {
	return string(ds.Kind)
}

// MarshalJSONByMethod implements the util.MethodMarshaler interface
func (ds *ViewDataset) MarshalJSONByMethod(method string) ([]byte, error) {
	return util.MarshalByMethod(*ds, method)
}

// LookupDataset represents a fully-constructed lookup dataset
type LookupDataset struct {
	// Common dataset properties:
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name,omitempty"`
	// The dataset kind.
	Kind DatasetInfoKind `json:"kind,omitempty" methods:"GET,POST"`
	// A unique dataset ID. Random ID used if not provided. Not valid for PATCH method.
	ID string `json:"id,omitempty" methods:"GET,POST"`
	// The name of the module to create the new dataset in. The default module is "".
	Module *string `json:"module,omitempty"`
	// The catalog version.
	Version int `json:"version,omitempty" methods:"GET"`
	// The date and time object was created.
	Created string `json:"created,omitempty" methods:"GET"`
	// The date and time object was modified.
	Modified string `json:"modified,omitempty" methods:"GET"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	CreatedBy string `json:"createdby,omitempty" methods:"GET"`
	// The name of the user who most recently modified the object.
	ModifiedBy string `json:"modifiedby,omitempty" methods:"GET"`
	// The name of the object's owner.
	Owner string `json:"owner,omitempty" methods:"GET,PATCH"`
	// The dataset name qualified by the module name.
	ResourceName string `json:"resourcename,omitempty" methods:"GET"`

	// Lookup-specific properties:
	// Match case-sensitively against the lookup.
	CaseSensitiveMatch *bool `json:"caseSensitiveMatch,omitempty"`
	// The type of the external lookup, this should always be `kvcollection`
	ExternalKind string `json:"externalKind,omitempty"`
	// The name of the external lookup.
	ExternalName *string `json:"externalName,omitempty"`
	// A query that filters results out of the lookup before those results are returned.
	Filter string `json:"filter,omitempty"`
}

// GetName the dataset name. Dataset names must be unique within each module.
func (ds LookupDataset) GetName() string {
	return ds.Name
}

// GetID returns a unique dataset ID. Random ID used if not provided
func (ds LookupDataset) GetID() string {
	return ds.ID
}

// GetModule returns the name of the module associated with the dataset.
func (ds LookupDataset) GetModule() string {
	return *ds.Module
}

// GetKind returns the dataset kind.
func (ds LookupDataset) GetKind() string {
	return string(ds.Kind)
}

// MarshalJSONByMethod implements the util.MethodMarshaler interface
func (ds *LookupDataset) MarshalJSONByMethod(method string) ([]byte, error) {
	return util.MarshalByMethod(*ds, method)
}

// KVCollectionDataset represents a fully-constructed kvcollection dataset
// NOTE: Only GET, POST, and DELETE are supported for KVCollection datasets
type KVCollectionDataset struct {
	// Common dataset properties:
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name,omitempty" methods:"GET,POST"`
	// The dataset kind.
	Kind DatasetInfoKind `json:"kind,omitempty" methods:"GET,POST"`
	// A unique dataset ID. Random ID used if not provided. Not valid for PATCH method.
	ID string `json:"id,omitempty" methods:"GET,POST"`
	// The name of the module to create the new dataset in. The default module is "".
	Module *string `json:"module,omitempty" methods:"GET,POST"`
	// The catalog version.
	Version int `json:"version,omitempty" methods:"GET"`
	// The date and time object was created.
	Created string `json:"created,omitempty" methods:"GET"`
	// The date and time object was modified.
	Modified string `json:"modified,omitempty" methods:"GET"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	CreatedBy string `json:"createdby,omitempty" methods:"GET"`
	// The name of the user who most recently modified the object.
	ModifiedBy string `json:"modifiedby,omitempty" methods:"GET"`
	// The name of the object's owner.
	Owner string `json:"owner,omitempty" methods:"GET"`
	// The dataset name qualified by the module name.
	ResourceName string `json:"resourcename,omitempty" methods:"GET"`

	// Note: there are no KVCollection-specific properties.
}

// GetName the dataset name. Dataset names must be unique within each module.
func (ds KVCollectionDataset) GetName() string {
	return ds.Name
}

// GetID returns a unique dataset ID. Random ID used if not provided
func (ds KVCollectionDataset) GetID() string {
	return ds.ID
}

// GetModule returns the name of the module associated with the dataset.
func (ds KVCollectionDataset) GetModule() string {
	return *ds.Module
}

// GetKind returns the dataset kind.
func (ds KVCollectionDataset) GetKind() string {
	return string(ds.Kind)
}

// MarshalJSONByMethod implements the util.MethodMarshaler interface
func (ds *KVCollectionDataset) MarshalJSONByMethod(method string) ([]byte, error) {
	return util.MarshalByMethod(*ds, method)
}

// GenericDataset represents a dataset of unknown kind
type GenericDataset map[string]interface{}

// GetName the dataset name. Dataset names must be unique within each module.
func (ds GenericDataset) GetName() string {
	val, ok := ds["name"]
	if !ok {
		return ""
	}
	strval, _ := val.(string)
	return strval
}

// GetID returns a unique dataset ID. Random ID used if not provided
func (ds GenericDataset) GetID() string {
	val, ok := ds["id"]
	if !ok {
		return ""
	}
	strval, _ := val.(string)
	return strval
}

// GetModule returns the name of the module associated with the dataset.
func (ds GenericDataset) GetModule() string {
	val, ok := ds["module"]
	if !ok {
		return ""
	}
	strval, _ := val.(string)
	return strval
}

// GetKind returns the dataset kind.
func (ds GenericDataset) GetKind() string {
	val, ok := ds["kind"]
	if !ok {
		return ""
	}
	strval, _ := val.(string)
	return strval
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
	// All PrevalenceType
	All PrevalenceType = "ALL"
	// Some PrevalenceType
	Some PrevalenceType = "SOME"
	// PrevalenceUnknown PrevalenceType
	PrevalenceUnknown PrevalenceType = "UNKNOWN"
)

// DataType enumerates the kinds of datatypes used in fields.
type DataType string

const (
	// Date DataType
	Date DataType = "DATE"
	// Number DataType
	Number DataType = "NUMBER"
	// ObjectID DataType
	ObjectID DataType = "OBJECT_ID"
	// String DataType
	String DataType = "STRING"
	// DataTypeUnknown DataType
	DataTypeUnknown DataType = "UNKNOWN"
)

// FieldType enumerates different kinds of fields.
type FieldType string

const (
	// Dimension fieldType
	Dimension FieldType = "DIMENSION"
	// Measure fieldType
	Measure FieldType = "MEASURE"
	// FieldTypeUnknown fieldType
	FieldTypeUnknown FieldType = "UNKNOWN"
)

// ActionKind enumerates the kinds of search time transformation action known by the service.
type ActionKind string

const (
	// Alias action
	Alias ActionKind = "ALIAS"
	// AutoKV action
	AutoKV ActionKind = "AUTOKV"
	// Regex action
	Regex ActionKind = "REGEX"
	// Eval action
	Eval ActionKind = "EVAL"
	// LookupAction action
	LookupAction ActionKind = "LOOKUP"
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
		Kind:  Alias,
		Owner: owner,
		Alias: alias,
		Field: field,
	}
}

// NewAutoKVAction creates a new autokv kind action
func NewAutoKVAction(mode string, owner string) *Action {
	return &Action{
		Kind:  AutoKV,
		Owner: owner,
		Mode:  mode,
	}
}

// NewEvalAction creates a new eval kind action
func NewEvalAction(field string, expression string, owner string) *Action {
	return &Action{
		Kind:       Eval,
		Owner:      owner,
		Field:      field,
		Expression: expression,
	}
}

// NewLookupAction creates a new lookup kind action
func NewLookupAction(expression string, owner string) *Action {
	return &Action{
		Kind:       LookupAction,
		Owner:      owner,
		Expression: expression,
	}
}

// NewRegexAction creates a new regex kind action
func NewRegexAction(field string, pattern string, limit *int, owner string) *Action {
	action := Action{
		Kind:    Regex,
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

// DEPRECATED: The following fields have been deprecated, see comments for more details

// DatasetInfo is Deprecated: 0.7.2 please use DatasetBase and *Dataset for each kind
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

// DatasetCreationPayload is Deprecated: 0.7.2 please use CreateDatasetBase and Create*Dataset for each kind
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

	Search                 string `json:"search,omitempty"`
	FrozenTimePeriodInSecs *int   `json:"frozenTimePeriodInSecs,omitempty"`
}

// UpdateDatasetInfoFields is Deprecated: 0.7.2 please use UpdateDatasetBase and Update*Dataset for each kind
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
