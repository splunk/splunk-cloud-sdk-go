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

DatasetCreationPayload represents the sources of data that can be searched by
Splunk

#### type DatasetInfo

```go
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
```

DatasetInfo represents the sources of data that can be searched by Splunk

#### type DatasetInfoKind

```go
type DatasetInfoKind string
```

DatasetInfoKind enumerates the kinds of datasets known to the system.

```go
const (
	// Lookup represents TODO: Description needed
	Lookup DatasetInfoKind = "lookup"
	// KvCollection represents a key value store, it is used with the kvstore service, but its implementation is separate of kvstore
	KvCollection DatasetInfoKind = "kvcollection"
	// Index represents a Splunk events or metrics index
	Index DatasetInfoKind = "index"
	// Metric represents TODO: Description needed
	Metric DatasetInfoKind = "metric"
	// View represents TODO: Description needed
	View DatasetInfoKind = "view"
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

#### type Module

```go
type Module struct {
	Name string `json:"name"`
}
```

Module represents catalog module

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
func (s *Service) CreateDataset(dataset *DatasetCreationPayload) (*DatasetInfo, error)
```
CreateDataset creates a new Dataset

#### func (*Service) CreateDatasetField

```go
func (s *Service) CreateDatasetField(datasetID string, datasetField *Field) (*Field, error)
```
CreateDatasetField creates a new field in the specified dataset

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
func (s *Service) GetDataset(resourceNameOrID string) (*DatasetInfo, error)
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
func (s *Service) GetDatasets() ([]DatasetInfo, error)
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
func (s *Service) ListDatasets(values url.Values) ([]DatasetInfo, error)
```
ListDatasets returns all Datasets with optional filter, count, or orderby params

#### func (*Service) UpdateDataset

```go
func (s *Service) UpdateDataset(dataset *UpdateDatasetInfoFields, resourceNameOrID string) (*DatasetInfo, error)
```
UpdateDataset updates an existing Dataset with the specified resourceName or ID

#### func (*Service) UpdateDatasetField

```go
func (s *Service) UpdateDatasetField(datasetID string, datasetFieldID string, datasetField *Field) (*Field, error)
```
UpdateDatasetField updates an already existing field in the specified dataset

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

UpdateDatasetInfoFields represents the sources of data that can be updated by
Splunk, same structure as DatasetInfo