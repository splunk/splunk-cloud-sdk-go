# catalog
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/catalog"


## Usage

#### type Action

```go
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
```

Action represents a specific search time transformation action. This struct
should NOT be directly used to construct object, use the NewXXXAction() instead

#### func  NewAliasAction

```go
func NewAliasAction(field string, alias string, owner string) *Action
```
NewAliasAction creates a new alias kind action

#### func  NewAutoKVAction

```go
func NewAutoKVAction(mode string, owner string) *Action
```
NewAutoKVAction creates a new autokv kind action

#### func  NewEvalAction

```go
func NewEvalAction(field string, expression string, owner string) *Action
```
NewEvalAction creates a new eval kind action

#### func  NewLookupAction

```go
func NewLookupAction(expression string, owner string) *Action
```
NewLookupAction creates a new lookup kind action

#### func  NewRegexAction

```go
func NewRegexAction(field string, pattern string, limit *int, owner string) *Action
```
NewRegexAction creates a new regex kind action

#### func  NewUpdateAliasAction

```go
func NewUpdateAliasAction(field *string, alias *string) *Action
```
NewUpdateAliasAction updates an existing alias kind action

#### func  NewUpdateAutoKVAction

```go
func NewUpdateAutoKVAction(mode *string) *Action
```
NewUpdateAutoKVAction updates an existing autokv kind action

#### func  NewUpdateEvalAction

```go
func NewUpdateEvalAction(field *string, expression *string) *Action
```
NewUpdateEvalAction updates an existing eval kind action

#### func  NewUpdateLookupAction

```go
func NewUpdateLookupAction(expression *string) *Action
```
NewUpdateLookupAction updates an existing lookup kind action

#### func  NewUpdateRegexAction

```go
func NewUpdateRegexAction(field *string, pattern *string, limit *int) *Action
```
NewUpdateRegexAction updates an existing regex kind action

#### type ActionKind

```go
type ActionKind string
```

ActionKind enumerates the kinds of search time transformation action known by
the service.

```go
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
```

#### type DataType

```go
type DataType string
```

DataType enumerates the kinds of datatypes used in fields.

```go
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
```

#### type Dataset

```go
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
```

Dataset represents the sources of data that can be searched by Splunk

#### func  ParseRawDataset

```go
func ParseRawDataset(dataset interface{}) (Dataset, error)
```
ParseRawDataset parses a raw interface{} type into specific Dataset subtype
based on 'kind'

#### type DatasetCreationPayload

```go
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
```

DatasetCreationPayload is Deprecated: 0.7.2 please use *Dataset for each kind

#### type DatasetInfoKind

```go
type DatasetInfoKind string
```

DatasetInfoKind enumerates the kinds of datasets known to the system.

```go
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
```

#### type Field

```go
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
```

Field represents the fields belonging to the specified Dataset

#### type FieldType

```go
type FieldType string
```

FieldType enumerates different kinds of fields.

```go
const (
	// Dimension fieldType
	Dimension FieldType = "DIMENSION"
	// Measure fieldType
	Measure FieldType = "MEASURE"
	// FieldTypeUnknown fieldType
	FieldTypeUnknown FieldType = "UNKNOWN"
)
```

#### type ImportDataset

```go
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
```

ImportDataset represents a fully-constructed import dataset

#### func (*ImportDataset) GetID

```go
func (ds *ImportDataset) GetID() string
```
GetID returns a unique dataset ID. Random ID used if not provided

#### func (*ImportDataset) GetKind

```go
func (ds *ImportDataset) GetKind() string
```
GetKind returns the dataset kind.

#### func (*ImportDataset) GetModule

```go
func (ds *ImportDataset) GetModule() string
```
GetModule returns the name of the module associated with the dataset.

#### func (*ImportDataset) GetName

```go
func (ds *ImportDataset) GetName() string
```
GetName the dataset name. Dataset names must be unique within each module.

#### func (*ImportDataset) MarshalJSONByMethod

```go
func (ds *ImportDataset) MarshalJSONByMethod(method string) ([]byte, error)
```
MarshalJSONByMethod implements the util.MethodMarshaler interface

#### type IndexDataset

```go
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
```

IndexDataset represents a fully-constructed index dataset

#### func (*IndexDataset) GetID

```go
func (ds *IndexDataset) GetID() string
```
GetID returns a unique dataset ID. Random ID used if not provided

#### func (*IndexDataset) GetKind

```go
func (ds *IndexDataset) GetKind() string
```
GetKind returns the dataset kind.

#### func (*IndexDataset) GetModule

```go
func (ds *IndexDataset) GetModule() string
```
GetModule returns the name of the module associated with the dataset.

#### func (*IndexDataset) GetName

```go
func (ds *IndexDataset) GetName() string
```
GetName the dataset name. Dataset names must be unique within each module.

#### func (*IndexDataset) MarshalJSONByMethod

```go
func (ds *IndexDataset) MarshalJSONByMethod(method string) ([]byte, error)
```
MarshalJSONByMethod implements the util.MethodMarshaler interface

#### type JobDataset

```go
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
```

JobDataset represents a fully-constructed job dataset NOTE: POST is not
supported for Job datasets, please use the search service to create jobs NOTE:
only Name, Module, Owner, and Status are supported for PATCH

#### func (*JobDataset) GetID

```go
func (ds *JobDataset) GetID() string
```
GetID returns a unique dataset ID. Random ID used if not provided

#### func (*JobDataset) GetKind

```go
func (ds *JobDataset) GetKind() string
```
GetKind returns the dataset kind.

#### func (*JobDataset) GetModule

```go
func (ds *JobDataset) GetModule() string
```
GetModule returns the name of the module associated with the dataset.

#### func (*JobDataset) GetName

```go
func (ds *JobDataset) GetName() string
```
GetName the dataset name. Dataset names must be unique within each module.

#### func (*JobDataset) MarshalJSONByMethod

```go
func (ds *JobDataset) MarshalJSONByMethod(method string) ([]byte, error)
```
MarshalJSONByMethod implements the util.MethodMarshaler interface

#### type KVCollectionDataset

```go
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
}
```

KVCollectionDataset represents a fully-constructed kvcollection dataset NOTE:
Only GET, POST, and DELETE are supported for KVCollection datasets

#### func (*KVCollectionDataset) GetID

```go
func (ds *KVCollectionDataset) GetID() string
```
GetID returns a unique dataset ID. Random ID used if not provided

#### func (*KVCollectionDataset) GetKind

```go
func (ds *KVCollectionDataset) GetKind() string
```
GetKind returns the dataset kind.

#### func (*KVCollectionDataset) GetModule

```go
func (ds *KVCollectionDataset) GetModule() string
```
GetModule returns the name of the module associated with the dataset.

#### func (*KVCollectionDataset) GetName

```go
func (ds *KVCollectionDataset) GetName() string
```
GetName the dataset name. Dataset names must be unique within each module.

#### func (*KVCollectionDataset) MarshalJSONByMethod

```go
func (ds *KVCollectionDataset) MarshalJSONByMethod(method string) ([]byte, error)
```
MarshalJSONByMethod implements the util.MethodMarshaler interface

#### type LookupDataset

```go
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
```

LookupDataset represents a fully-constructed lookup dataset

#### func (*LookupDataset) GetID

```go
func (ds *LookupDataset) GetID() string
```
GetID returns a unique dataset ID. Random ID used if not provided

#### func (*LookupDataset) GetKind

```go
func (ds *LookupDataset) GetKind() string
```
GetKind returns the dataset kind.

#### func (*LookupDataset) GetModule

```go
func (ds *LookupDataset) GetModule() string
```
GetModule returns the name of the module associated with the dataset.

#### func (*LookupDataset) GetName

```go
func (ds *LookupDataset) GetName() string
```
GetName the dataset name. Dataset names must be unique within each module.

#### func (*LookupDataset) MarshalJSONByMethod

```go
func (ds *LookupDataset) MarshalJSONByMethod(method string) ([]byte, error)
```
MarshalJSONByMethod implements the util.MethodMarshaler interface

#### type MetricDataset

```go
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
```

MetricDataset represents a fully-constructed index dataset

#### func (*MetricDataset) GetID

```go
func (ds *MetricDataset) GetID() string
```
GetID returns a unique dataset ID. Random ID used if not provided

#### func (*MetricDataset) GetKind

```go
func (ds *MetricDataset) GetKind() string
```
GetKind returns the dataset kind.

#### func (*MetricDataset) GetModule

```go
func (ds *MetricDataset) GetModule() string
```
GetModule returns the name of the module associated with the dataset.

#### func (*MetricDataset) GetName

```go
func (ds *MetricDataset) GetName() string
```
GetName the dataset name. Dataset names must be unique within each module.

#### func (*MetricDataset) MarshalJSONByMethod

```go
func (ds *MetricDataset) MarshalJSONByMethod(method string) ([]byte, error)
```
MarshalJSONByMethod implements the util.MethodMarshaler interface

#### type Module

```go
type Module struct {
	Name string `json:"name"`
}
```

Module represents catalog module

#### type OtherDataset

```go
type OtherDataset map[string]interface{}
```

OtherDataset represents a dataset of unknown kind

#### func (*OtherDataset) GetID

```go
func (ds *OtherDataset) GetID() string
```
GetID returns a unique dataset ID. Random ID used if not provided

#### func (*OtherDataset) GetKind

```go
func (ds *OtherDataset) GetKind() string
```
GetKind returns the dataset kind.

#### func (*OtherDataset) GetModule

```go
func (ds *OtherDataset) GetModule() string
```
GetModule returns the name of the module associated with the dataset.

#### func (*OtherDataset) GetName

```go
func (ds *OtherDataset) GetName() string
```
GetName the dataset name. Dataset names must be unique within each module.

#### type PrevalenceType

```go
type PrevalenceType string
```

PrevalenceType enumerates the types of prevalance used in fields.

```go
const (
	// All PrevalenceType
	All PrevalenceType = "ALL"
	// Some PrevalenceType
	Some PrevalenceType = "SOME"
	// PrevalenceUnknown PrevalenceType
	PrevalenceUnknown PrevalenceType = "UNKNOWN"
)
```

#### type Rule

```go
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
```

Rule represents a rule for transforming results at search time. A rule consists
of a `match` clause and a collection of transformation actions

#### type RuleUpdateFields

```go
type RuleUpdateFields struct {
	Name    string `json:"name,omitempty"`
	Module  string `json:"module,omitempty"`
	Match   string `json:"match,omitempty"`
	Owner   string `json:"owner,omitempty"`
	Version int    `json:"version,omitempty"`
}
```

RuleUpdateFields represents the set of rule properties that can be updated

#### type Service

```go
type Service services.BaseService
```

Service talks to the Splunk Cloud catalog service

#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new catalog service client from the given Config

#### func (*Service) CreateDataset

```go
func (s *Service) CreateDataset(dataset interface{}) (Dataset, error)
```
CreateDataset creates a new Dataset

#### func (*Service) CreateDatasetField

```go
func (s *Service) CreateDatasetField(datasetID string, datasetField *Field) (*Field, error)
```
CreateDatasetField creates a new field in the specified dataset

#### func (*Service) CreateImportDataset

```go
func (s *Service) CreateImportDataset(importDataset *ImportDataset) (*ImportDataset, error)
```
CreateImportDataset creates an import Dataset

#### func (*Service) CreateIndexDataset

```go
func (s *Service) CreateIndexDataset(indexDataset *IndexDataset) (*IndexDataset, error)
```
CreateIndexDataset creates an index Dataset

#### func (*Service) CreateKVCollectionDataset

```go
func (s *Service) CreateKVCollectionDataset(kvDataset *KVCollectionDataset) (*KVCollectionDataset, error)
```
CreateKVCollectionDataset creates a KVCollection Dataset

#### func (*Service) CreateLookupDataset

```go
func (s *Service) CreateLookupDataset(lookupDataset *LookupDataset) (*LookupDataset, error)
```
CreateLookupDataset creates a lookup Dataset

#### func (*Service) CreateMetricDataset

```go
func (s *Service) CreateMetricDataset(metricDataset *MetricDataset) (*MetricDataset, error)
```
CreateMetricDataset creates a metric Dataset

#### func (*Service) CreateRule

```go
func (s *Service) CreateRule(rule Rule) (*Rule, error)
```
CreateRule posts a new rule.

#### func (*Service) CreateRuleAction

```go
func (s *Service) CreateRuleAction(ruleID string, action *Action) (*Action, error)
```
CreateRuleAction creates a new Action on the rule specified

#### func (*Service) CreateViewDataset

```go
func (s *Service) CreateViewDataset(viewDataset *ViewDataset) (*ViewDataset, error)
```
CreateViewDataset creates a view Dataset

#### func (*Service) DeleteDataset

```go
func (s *Service) DeleteDataset(resourceNameOrID string) error
```
DeleteDataset implements delete Dataset endpoint with the specified resourceName
or ID

#### func (*Service) DeleteDatasetField

```go
func (s *Service) DeleteDatasetField(datasetID string, datasetFieldID string) error
```
DeleteDatasetField deletes the field belonging to the specified dataset with the
id datasetFieldID

#### func (*Service) DeleteRule

```go
func (s *Service) DeleteRule(resourceNameOrID string) error
```
DeleteRule deletes the rule and its dependencies with the specified rule id or
resourceName

#### func (*Service) DeleteRuleAction

```go
func (s *Service) DeleteRuleAction(ruleID string, actionID string) error
```
DeleteRuleAction deletes the action of specified belonging to the specified rule

#### func (*Service) GetDataset

```go
func (s *Service) GetDataset(resourceNameOrID string) (Dataset, error)
```
GetDataset returns the Dataset by resourceName or ID

#### func (*Service) GetDatasetField

```go
func (s *Service) GetDatasetField(datasetID string, datasetFieldID string) (*Field, error)
```
GetDatasetField returns the field belonging to the specified dataset with the id
datasetFieldID

#### func (*Service) GetDatasetFields

```go
func (s *Service) GetDatasetFields(datasetID string, values url.Values) ([]Field, error)
```
GetDatasetFields returns all the fields belonging to the specified dataset

#### func (*Service) GetDatasets

```go
func (s *Service) GetDatasets() ([]Dataset, error)
```
GetDatasets returns all Datasets Deprecated: v0.6.1 - Use ListDatasets instead

#### func (*Service) GetField

```go
func (s *Service) GetField(fieldID string) (*Field, error)
```
GetField returns the Field corresponding to fieldid

#### func (*Service) GetFields

```go
func (s *Service) GetFields() ([]Field, error)
```
GetFields returns a list of all Fields on Catalog

#### func (*Service) GetModules

```go
func (s *Service) GetModules(filter url.Values) ([]Module, error)
```
GetModules returns a list of a list of modules that match a filter query if it
is given, otherwise return all modules

#### func (*Service) GetRule

```go
func (s *Service) GetRule(resourceNameOrID string) (*Rule, error)
```
GetRule returns rule by the specified resourceName or ID.

#### func (*Service) GetRuleAction

```go
func (s *Service) GetRuleAction(ruleID string, actionID string) (*Action, error)
```
GetRuleAction returns the action of specified belonging to the specified rule

#### func (*Service) GetRuleActions

```go
func (s *Service) GetRuleActions(ruleID string) ([]Action, error)
```
GetRuleActions returns a list of all actions belonging to the specified rule

#### func (*Service) GetRules

```go
func (s *Service) GetRules() ([]Rule, error)
```
GetRules returns all the rules.

#### func (*Service) ListDatasets

```go
func (s *Service) ListDatasets(values url.Values) ([]Dataset, error)
```
ListDatasets returns all Datasets with optional filter, count, or orderby params

#### func (*Service) UpdateDataset

```go
func (s *Service) UpdateDataset(dataset interface{}, resourceNameOrID string) (Dataset, error)
```
UpdateDataset updates an existing Dataset with the specified resourceName or ID

#### func (*Service) UpdateDatasetField

```go
func (s *Service) UpdateDatasetField(datasetID string, datasetFieldID string, datasetField *Field) (*Field, error)
```
UpdateDatasetField updates an already existing field in the specified dataset

#### func (*Service) UpdateIndexDataset

```go
func (s *Service) UpdateIndexDataset(indexDataset *IndexDataset, id string) (*IndexDataset, error)
```
UpdateIndexDataset updates an existing index Dataset with the specified
resourceName or ID

#### func (*Service) UpdateJobDataset

```go
func (s *Service) UpdateJobDataset(jobDataset *JobDataset, id string) (*JobDataset, error)
```
UpdateJobDataset updates an existing job Dataset with the specified resourceName
or ID

#### func (*Service) UpdateLookupDataset

```go
func (s *Service) UpdateLookupDataset(lookupDataset *LookupDataset, id string) (*LookupDataset, error)
```
UpdateLookupDataset updates an existing lookup Dataset with the specified
resourceName or ID

#### func (*Service) UpdateMetricDataset

```go
func (s *Service) UpdateMetricDataset(metricDataset *MetricDataset, id string) (*MetricDataset, error)
```
UpdateMetricDataset updates an existing metric Dataset with the specified
resourceName or ID

#### func (*Service) UpdateRule

```go
func (s *Service) UpdateRule(resourceNameOrID string, rule *RuleUpdateFields) (*Rule, error)
```
UpdateRule updates the rule with the specified resourceName or ID

#### func (*Service) UpdateRuleAction

```go
func (s *Service) UpdateRuleAction(ruleID string, actionID string, action *Action) (*Action, error)
```
UpdateRuleAction updates the action with the specified id for the specified Rule

#### func (*Service) UpdateViewDataset

```go
func (s *Service) UpdateViewDataset(viewDataset *ViewDataset, id string) (*ViewDataset, error)
```
UpdateViewDataset updates an existing view Dataset with the specified
resourceName or ID

#### type Servicer

```go
type Servicer interface {
	// ListDatasets returns all Datasets with optional filter, count, or orderby params
	ListDatasets(values url.Values) ([]Dataset, error)
	// GetDatasets returns all Datasets
	// Deprecated: v0.6.1 - Use ListDatasets instead
	GetDatasets() ([]Dataset, error)
	// GetDataset returns the Dataset by resourceName or ID
	GetDataset(resourceNameOrID string) (Dataset, error)
	// CreateDataset creates a new Dataset
	CreateDataset(dataset interface{}) (Dataset, error)
	// CreateIndexDataset creates an index Dataset
	CreateIndexDataset(indexDataset *IndexDataset) (*IndexDataset, error)
	// CreateLookupDataset creates a lookup Dataset
	CreateLookupDataset(lookupDataset *LookupDataset) (*LookupDataset, error)
	// CreateViewDataset creates a view Dataset
	CreateViewDataset(viewDataset *ViewDataset) (*ViewDataset, error)
	// CreateKVCollectionDataset creates a KVCollection Dataset
	CreateKVCollectionDataset(kvDataset *KVCollectionDataset) (*KVCollectionDataset, error)
	// CreateImportDataset creates an import Dataset
	CreateImportDataset(importDataset *ImportDataset) (*ImportDataset, error)
	// CreateMetricDataset creates a metric Dataset
	CreateMetricDataset(metricDataset *MetricDataset) (*MetricDataset, error)
	// UpdateDataset updates an existing Dataset with the specified resourceName or ID
	UpdateDataset(dataset interface{}, resourceNameOrID string) (Dataset, error)
	// UpdateIndexDataset updates an existing index Dataset with the specified resourceName or ID
	UpdateIndexDataset(indexDataset *IndexDataset, id string) (*IndexDataset, error)
	// UpdateJobDataset updates an existing job Dataset with the specified resourceName or ID
	UpdateJobDataset(jobDataset *JobDataset, id string) (*JobDataset, error)
	// UpdateLookupDataset updates an existing lookup Dataset with the specified resourceName or ID
	UpdateLookupDataset(lookupDataset *LookupDataset, id string) (*LookupDataset, error)
	// UpdateViewDataset updates an existing view Dataset with the specified resourceName or ID
	UpdateViewDataset(viewDataset *ViewDataset, id string) (*ViewDataset, error)
	// UpdateMetricDataset updates an existing metric Dataset with the specified resourceName or ID
	UpdateMetricDataset(metricDataset *MetricDataset, id string) (*MetricDataset, error)
	// DeleteDataset implements delete Dataset endpoint with the specified resourceName or ID
	DeleteDataset(resourceNameOrID string) error
	// DeleteRule deletes the rule and its dependencies with the specified rule id or resourceName
	DeleteRule(resourceNameOrID string) error
	// GetRules returns all the rules.
	GetRules() ([]Rule, error)
	// GetRule returns rule by the specified resourceName or ID.
	GetRule(resourceNameOrID string) (*Rule, error)
	// CreateRule posts a new rule.
	CreateRule(rule Rule) (*Rule, error)
	// UpdateRule updates the rule with the specified resourceName or ID
	UpdateRule(resourceNameOrID string, rule *RuleUpdateFields) (*Rule, error)
	// GetDatasetFields returns all the fields belonging to the specified dataset
	GetDatasetFields(datasetID string, values url.Values) ([]Field, error)
	// GetDatasetField returns the field belonging to the specified dataset with the id datasetFieldID
	GetDatasetField(datasetID string, datasetFieldID string) (*Field, error)
	// CreateDatasetField creates a new field in the specified dataset
	CreateDatasetField(datasetID string, datasetField *Field) (*Field, error)
	// UpdateDatasetField updates an already existing field in the specified dataset
	UpdateDatasetField(datasetID string, datasetFieldID string, datasetField *Field) (*Field, error)
	// DeleteDatasetField deletes the field belonging to the specified dataset with the id datasetFieldID
	DeleteDatasetField(datasetID string, datasetFieldID string) error
	// GetFields returns a list of all Fields on Catalog
	GetFields() ([]Field, error)
	// GetField returns the Field corresponding to fieldid
	GetField(fieldID string) (*Field, error)
	// CreateRuleAction creates a new Action on the rule specified
	CreateRuleAction(ruleID string, action *Action) (*Action, error)
	// GetRuleActions returns a list of all actions belonging to the specified rule
	GetRuleActions(ruleID string) ([]Action, error)
	// GetRuleAction returns the action of specified belonging to the specified rule
	GetRuleAction(ruleID string, actionID string) (*Action, error)
	// DeleteRuleAction deletes the action of specified belonging to the specified rule
	DeleteRuleAction(ruleID string, actionID string) error
	// UpdateRuleAction updates the action with the specified id for the specified Rule
	UpdateRuleAction(ruleID string, actionID string, action *Action) (*Action, error)
	// GetModules returns a list of a list of modules that match a filter query if it is given, otherwise return all modules
	GetModules(filter url.Values) ([]Module, error)
}
```

Servicer ...

#### type UpdateDatasetInfoFields

```go
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
```

UpdateDatasetInfoFields is Deprecated: 0.7.2 please use *Dataset for each kind

#### type ViewDataset

```go
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
```

ViewDataset represents a fully-constructed view dataset

#### func (*ViewDataset) GetID

```go
func (ds *ViewDataset) GetID() string
```
GetID returns a unique dataset ID. Random ID used if not provided

#### func (*ViewDataset) GetKind

```go
func (ds *ViewDataset) GetKind() string
```
GetKind returns the dataset kind.

#### func (*ViewDataset) GetModule

```go
func (ds *ViewDataset) GetModule() string
```
GetModule returns the name of the module associated with the dataset.

#### func (*ViewDataset) GetName

```go
func (ds *ViewDataset) GetName() string
```
GetName the dataset name. Dataset names must be unique within each module.

#### func (*ViewDataset) MarshalJSONByMethod

```go
func (ds *ViewDataset) MarshalJSONByMethod(method string) ([]byte, error)
```
MarshalJSONByMethod implements the util.MethodMarshaler interface
