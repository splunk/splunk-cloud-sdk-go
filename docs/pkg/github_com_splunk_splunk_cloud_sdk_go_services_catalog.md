# catalog
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/catalog"


## Usage

#### type Action

```go
type Action struct {
}
```

Action : A complete action as rendered in POST, PATCH, and GET responses.

#### func  MakeActionFromAliasAction

```go
func MakeActionFromAliasAction(f AliasAction) Action
```
MakeActionFromAliasAction creates a new Action from an instance of AliasAction

#### func  MakeActionFromAutoKvAction

```go
func MakeActionFromAutoKvAction(f AutoKvAction) Action
```
MakeActionFromAutoKvAction creates a new Action from an instance of AutoKvAction

#### func  MakeActionFromEvalAction

```go
func MakeActionFromEvalAction(f EvalAction) Action
```
MakeActionFromEvalAction creates a new Action from an instance of EvalAction

#### func  MakeActionFromLookupAction

```go
func MakeActionFromLookupAction(f LookupAction) Action
```
MakeActionFromLookupAction creates a new Action from an instance of LookupAction

#### func  MakeActionFromRawInterface

```go
func MakeActionFromRawInterface(f interface{}) Action
```
MakeActionFromRawInterface creates a new Action from a raw interface{}

#### func  MakeActionFromRegexAction

```go
func MakeActionFromRegexAction(f RegexAction) Action
```
MakeActionFromRegexAction creates a new Action from an instance of RegexAction

#### func (Action) AliasAction

```go
func (m Action) AliasAction() *AliasAction
```
AliasAction returns AliasAction if IsAliasAction() is true, nil otherwise

#### func (Action) AutoKvAction

```go
func (m Action) AutoKvAction() *AutoKvAction
```
AutoKvAction returns AutoKvAction if IsAutoKvAction() is true, nil otherwise

#### func (Action) EvalAction

```go
func (m Action) EvalAction() *EvalAction
```
EvalAction returns EvalAction if IsEvalAction() is true, nil otherwise

#### func (Action) IsAliasAction

```go
func (m Action) IsAliasAction() bool
```
IsAliasAction checks if the Action is a AliasAction

#### func (Action) IsAutoKvAction

```go
func (m Action) IsAutoKvAction() bool
```
IsAutoKvAction checks if the Action is a AutoKvAction

#### func (Action) IsEvalAction

```go
func (m Action) IsEvalAction() bool
```
IsEvalAction checks if the Action is a EvalAction

#### func (Action) IsLookupAction

```go
func (m Action) IsLookupAction() bool
```
IsLookupAction checks if the Action is a LookupAction

#### func (Action) IsRawInterface

```go
func (m Action) IsRawInterface() bool
```
IsRawInterface checks if the Action is an interface{} (unknown type)

#### func (Action) IsRegexAction

```go
func (m Action) IsRegexAction() bool
```
IsRegexAction checks if the Action is a RegexAction

#### func (Action) LookupAction

```go
func (m Action) LookupAction() *LookupAction
```
LookupAction returns LookupAction if IsLookupAction() is true, nil otherwise

#### func (Action) MarshalJSON

```go
func (m Action) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals Action using the appropriate struct field

#### func (Action) RawInterface

```go
func (m Action) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (Action) RegexAction

```go
func (m Action) RegexAction() *RegexAction
```
RegexAction returns RegexAction if IsRegexAction() is true, nil otherwise

#### func (*Action) UnmarshalJSON

```go
func (m *Action) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals Action using the "kind" property

#### type ActionCommon

```go
type ActionCommon struct {
	// A unique action ID.
	Id *string `json:"id,omitempty"`
	// The rule that this action is part of.
	Ruleid *string `json:"ruleid,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Properties common across all action kinds as rendered in POST, PATCH, and GET
responses and alos for creating a new action using a POST request.

#### type ActionPatch

```go
type ActionPatch struct {
}
```

ActionPatch : Property values to be set in an existing action using a PATCH
request. ActionPatch is AliasActionPatch, AutoKvActionPatch, EvalActionPatch,
LookupActionPatch, RegexActionPatch, (or interface{} if no matches are found)

#### func  MakeActionPatchFromAliasActionPatch

```go
func MakeActionPatchFromAliasActionPatch(f AliasActionPatch) ActionPatch
```
MakeActionPatchFromAliasActionPatch creates a new ActionPatch from an instance
of AliasActionPatch

#### func  MakeActionPatchFromAutoKvActionPatch

```go
func MakeActionPatchFromAutoKvActionPatch(f AutoKvActionPatch) ActionPatch
```
MakeActionPatchFromAutoKvActionPatch creates a new ActionPatch from an instance
of AutoKvActionPatch

#### func  MakeActionPatchFromEvalActionPatch

```go
func MakeActionPatchFromEvalActionPatch(f EvalActionPatch) ActionPatch
```
MakeActionPatchFromEvalActionPatch creates a new ActionPatch from an instance of
EvalActionPatch

#### func  MakeActionPatchFromLookupActionPatch

```go
func MakeActionPatchFromLookupActionPatch(f LookupActionPatch) ActionPatch
```
MakeActionPatchFromLookupActionPatch creates a new ActionPatch from an instance
of LookupActionPatch

#### func  MakeActionPatchFromRawInterface

```go
func MakeActionPatchFromRawInterface(f interface{}) ActionPatch
```
MakeActionPatchFromRawInterface creates a new ActionPatch from a raw interface{}

#### func  MakeActionPatchFromRegexActionPatch

```go
func MakeActionPatchFromRegexActionPatch(f RegexActionPatch) ActionPatch
```
MakeActionPatchFromRegexActionPatch creates a new ActionPatch from an instance
of RegexActionPatch

#### func (ActionPatch) AliasActionPatch

```go
func (m ActionPatch) AliasActionPatch() *AliasActionPatch
```
AliasActionPatch returns AliasActionPatch if IsAliasActionPatch() is true, nil
otherwise

#### func (ActionPatch) AutoKvActionPatch

```go
func (m ActionPatch) AutoKvActionPatch() *AutoKvActionPatch
```
AutoKvActionPatch returns AutoKvActionPatch if IsAutoKvActionPatch() is true,
nil otherwise

#### func (ActionPatch) EvalActionPatch

```go
func (m ActionPatch) EvalActionPatch() *EvalActionPatch
```
EvalActionPatch returns EvalActionPatch if IsEvalActionPatch() is true, nil
otherwise

#### func (ActionPatch) IsAliasActionPatch

```go
func (m ActionPatch) IsAliasActionPatch() bool
```
IsAliasActionPatch checks if the ActionPatch is a AliasActionPatch

#### func (ActionPatch) IsAutoKvActionPatch

```go
func (m ActionPatch) IsAutoKvActionPatch() bool
```
IsAutoKvActionPatch checks if the ActionPatch is a AutoKvActionPatch

#### func (ActionPatch) IsEvalActionPatch

```go
func (m ActionPatch) IsEvalActionPatch() bool
```
IsEvalActionPatch checks if the ActionPatch is a EvalActionPatch

#### func (ActionPatch) IsLookupActionPatch

```go
func (m ActionPatch) IsLookupActionPatch() bool
```
IsLookupActionPatch checks if the ActionPatch is a LookupActionPatch

#### func (ActionPatch) IsRawInterface

```go
func (m ActionPatch) IsRawInterface() bool
```
IsRawInterface checks if the ActionPatch is an interface{} (unknown type)

#### func (ActionPatch) IsRegexActionPatch

```go
func (m ActionPatch) IsRegexActionPatch() bool
```
IsRegexActionPatch checks if the ActionPatch is a RegexActionPatch

#### func (ActionPatch) LookupActionPatch

```go
func (m ActionPatch) LookupActionPatch() *LookupActionPatch
```
LookupActionPatch returns LookupActionPatch if IsLookupActionPatch() is true,
nil otherwise

#### func (ActionPatch) MarshalJSON

```go
func (m ActionPatch) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals ActionPatch using ActionPatch.ActionPatch

#### func (ActionPatch) RawInterface

```go
func (m ActionPatch) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (ActionPatch) RegexActionPatch

```go
func (m ActionPatch) RegexActionPatch() *RegexActionPatch
```
RegexActionPatch returns RegexActionPatch if IsRegexActionPatch() is true, nil
otherwise

#### func (*ActionPatch) UnmarshalJSON

```go
func (m *ActionPatch) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals ActionPatch into AliasActionPatch, AutoKvActionPatch,
EvalActionPatch, LookupActionPatch, RegexActionPatch, or interface{} if no
matches are found

#### type ActionPatchCommon

```go
type ActionPatchCommon struct {
	// The name of the user who owns this action. This value is obtained from the bearer token if not present.
	Owner *string `json:"owner,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Property values common across action kinds for setting existing actions using a
PATCH request.

#### type ActionPost

```go
type ActionPost struct {
}
```

ActionPost : Initial property values for creating a new action using a POST
request.

#### func  MakeActionPostFromAliasActionPost

```go
func MakeActionPostFromAliasActionPost(f AliasActionPost) ActionPost
```
MakeActionPostFromAliasActionPost creates a new ActionPost from an instance of
AliasActionPost

#### func  MakeActionPostFromAutoKvActionPost

```go
func MakeActionPostFromAutoKvActionPost(f AutoKvActionPost) ActionPost
```
MakeActionPostFromAutoKvActionPost creates a new ActionPost from an instance of
AutoKvActionPost

#### func  MakeActionPostFromEvalActionPost

```go
func MakeActionPostFromEvalActionPost(f EvalActionPost) ActionPost
```
MakeActionPostFromEvalActionPost creates a new ActionPost from an instance of
EvalActionPost

#### func  MakeActionPostFromLookupActionPost

```go
func MakeActionPostFromLookupActionPost(f LookupActionPost) ActionPost
```
MakeActionPostFromLookupActionPost creates a new ActionPost from an instance of
LookupActionPost

#### func  MakeActionPostFromRawInterface

```go
func MakeActionPostFromRawInterface(f interface{}) ActionPost
```
MakeActionPostFromRawInterface creates a new ActionPost from a raw interface{}

#### func  MakeActionPostFromRegexActionPost

```go
func MakeActionPostFromRegexActionPost(f RegexActionPost) ActionPost
```
MakeActionPostFromRegexActionPost creates a new ActionPost from an instance of
RegexActionPost

#### func (ActionPost) AliasActionPost

```go
func (m ActionPost) AliasActionPost() *AliasActionPost
```
AliasActionPost returns AliasActionPost if IsAliasActionPost() is true, nil
otherwise

#### func (ActionPost) AutoKvActionPost

```go
func (m ActionPost) AutoKvActionPost() *AutoKvActionPost
```
AutoKvActionPost returns AutoKvActionPost if IsAutoKvActionPost() is true, nil
otherwise

#### func (ActionPost) EvalActionPost

```go
func (m ActionPost) EvalActionPost() *EvalActionPost
```
EvalActionPost returns EvalActionPost if IsEvalActionPost() is true, nil
otherwise

#### func (ActionPost) IsAliasActionPost

```go
func (m ActionPost) IsAliasActionPost() bool
```
IsAliasActionPost checks if the ActionPost is a AliasActionPost

#### func (ActionPost) IsAutoKvActionPost

```go
func (m ActionPost) IsAutoKvActionPost() bool
```
IsAutoKvActionPost checks if the ActionPost is a AutoKvActionPost

#### func (ActionPost) IsEvalActionPost

```go
func (m ActionPost) IsEvalActionPost() bool
```
IsEvalActionPost checks if the ActionPost is a EvalActionPost

#### func (ActionPost) IsLookupActionPost

```go
func (m ActionPost) IsLookupActionPost() bool
```
IsLookupActionPost checks if the ActionPost is a LookupActionPost

#### func (ActionPost) IsRawInterface

```go
func (m ActionPost) IsRawInterface() bool
```
IsRawInterface checks if the ActionPost is an interface{} (unknown type)

#### func (ActionPost) IsRegexActionPost

```go
func (m ActionPost) IsRegexActionPost() bool
```
IsRegexActionPost checks if the ActionPost is a RegexActionPost

#### func (ActionPost) LookupActionPost

```go
func (m ActionPost) LookupActionPost() *LookupActionPost
```
LookupActionPost returns LookupActionPost if IsLookupActionPost() is true, nil
otherwise

#### func (ActionPost) MarshalJSON

```go
func (m ActionPost) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals ActionPost using the appropriate struct field

#### func (ActionPost) RawInterface

```go
func (m ActionPost) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (ActionPost) RegexActionPost

```go
func (m ActionPost) RegexActionPost() *RegexActionPost
```
RegexActionPost returns RegexActionPost if IsRegexActionPost() is true, nil
otherwise

#### func (*ActionPost) UnmarshalJSON

```go
func (m *ActionPost) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals ActionPost using the "kind" property

#### type AliasAction

```go
type AliasAction struct {
	// The alias name.
	Alias string `json:"alias"`
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// The name of the field to be aliased.
	Field string `json:"field"`
	// A unique action ID.
	Id   string          `json:"id"`
	Kind AliasActionKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The rule that this action is part of.
	Ruleid string `json:"ruleid"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete alias action as rendered in POST, PATCH, and GET responses.

#### type AliasActionKind

```go
type AliasActionKind string
```

AliasActionKind : The alias action kind.

```go
const (
	AliasActionKindAlias AliasActionKind = "ALIAS"
)
```
List of AliasActionKind

#### type AliasActionPatch

```go
type AliasActionPatch struct {
	// The alias name.
	Alias *string `json:"alias,omitempty"`
	// The name of the field to be aliased.
	Field *string          `json:"field,omitempty"`
	Kind  *AliasActionKind `json:"kind,omitempty"`
	// The name of the user who owns this action. This value is obtained from the bearer token if not present.
	Owner *string `json:"owner,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Property values for setting existing alias actions using a PATCH request.

#### type AliasActionPost

```go
type AliasActionPost struct {
	// The alias name.
	Alias string `json:"alias"`
	// The name of the field to be aliased.
	Field string          `json:"field"`
	Kind  AliasActionKind `json:"kind"`
	// A unique action ID.
	Id *string `json:"id,omitempty"`
	// The rule that this action is part of.
	Ruleid *string `json:"ruleid,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Initial property values for creating a new alias action using a POST request.

#### type AliasActionProperties

```go
type AliasActionProperties struct {
	// The alias name.
	Alias *string `json:"alias,omitempty"`
	// The name of the field to be aliased.
	Field *string          `json:"field,omitempty"`
	Kind  *AliasActionKind `json:"kind,omitempty"`
}
```

Properties of alias actions which may be read, set, and changed through the API.
Implementation detail of ActionPOST, ActionPOST, and Action, do not use
directly.

#### type Annotation

```go
type Annotation struct {
	// The annotation type ID.
	Annotationtypeid string `json:"annotationtypeid"`
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// A unique annotation ID.
	Id string `json:"id"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The dataset ID. Null if not annotating a dataset.
	Datasetid *string `json:"datasetid,omitempty"`
	// The field ID. Null if not annotating a field.
	Fieldid *string `json:"fieldid,omitempty"`
	// The relationship ID. Null if not annotating a relationship.
	Relationshipid *string `json:"relationshipid,omitempty"`
}
```

A complete annotation as rendered in POST, PATCH, and GET responses. Key:Value
pairs in addition to the defined properties will become annotation tags

#### type AnnotationPost

```go
type AnnotationPost struct {
	// The annotation type ID.
	Annotationtypeid string `json:"annotationtypeid"`
	// Resource name of the annotation type
	Annotationtyperesourcename *string `json:"annotationtyperesourcename,omitempty"`
	// The dataset ID. Null if not annotating a dataset.
	Datasetid *string `json:"datasetid,omitempty"`
	// The field ID. Null if not annotating a field.
	Fieldid *string `json:"fieldid,omitempty"`
	// A unique annotation ID. If not specified, an auto generated ID is created.
	Id *string `json:"id,omitempty"`
	// The relationship ID. Null if not annotating a relationship.
	Relationshipid *string `json:"relationshipid,omitempty"`
}
```

The properties required to create a new annotation using a POST request.
Key:Value pairs in addition to the defined properties will become annotation
tags

#### type AnnotationTypeResourceName

```go
type AnnotationTypeResourceName string
```

The resource name of an AnnotationType. For the default module, the resource
name format is annotationTypeName. Otherwise, the resource name format is
module.annotationTypeName.

#### type AnnotationsProperties

```go
type AnnotationsProperties struct {
	// The annotation type ID.
	Annotationtypeid *string `json:"annotationtypeid,omitempty"`
	// The dataset ID. Null if not annotating a dataset.
	Datasetid *string `json:"datasetid,omitempty"`
	// The field ID. Null if not annotating a field.
	Fieldid *string `json:"fieldid,omitempty"`
	// The relationship ID. Null if not annotating a relationship.
	Relationshipid *string `json:"relationshipid,omitempty"`
}
```

Properties of annotations which are read through the API.

#### type AutoKvAction

```go
type AutoKvAction struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// A unique action ID.
	Id   string           `json:"id"`
	Kind AutoKvActionKind `json:"kind"`
	// The autokv action mode.
	Mode string `json:"mode"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The rule that this action is part of.
	Ruleid string `json:"ruleid"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete autokv action as rendered in POST, PATCH, and GET responses.

#### type AutoKvActionKind

```go
type AutoKvActionKind string
```

AutoKvActionKind : The autokv action kind.

```go
const (
	AutoKvActionKindAutokv AutoKvActionKind = "AUTOKV"
)
```
List of AutoKVActionKind

#### type AutoKvActionPatch

```go
type AutoKvActionPatch struct {
	Kind *AutoKvActionKind `json:"kind,omitempty"`
	// The autokv action mode.
	Mode *string `json:"mode,omitempty"`
	// The name of the user who owns this action. This value is obtained from the bearer token if not present.
	Owner *string `json:"owner,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Property values for setting existing autokv actions using a PATCH request.

#### type AutoKvActionPost

```go
type AutoKvActionPost struct {
	Kind AutoKvActionKind `json:"kind"`
	// The autokv action mode.
	Mode string `json:"mode"`
	// A unique action ID.
	Id *string `json:"id,omitempty"`
	// The rule that this action is part of.
	Ruleid *string `json:"ruleid,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Initial property values for creating a new autokv action using a POST request.

#### type AutoKvActionProperties

```go
type AutoKvActionProperties struct {
	Kind *AutoKvActionKind `json:"kind,omitempty"`
	// The autokv action mode.
	Mode *string `json:"mode,omitempty"`
}
```

Properties of auto kv actions which may be read, set, and changed through the
API. Implementation detail of ActionPOST, ActionPOST, and Action, do not use
directly.

#### type Dashboard

```go
type Dashboard struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// The JSON dashboard definition.
	Definition string `json:"definition"`
	// A unique dashboard ID. Random ID used if not provided.
	Id string `json:"id"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The module that contains the dashboard.
	Module string `json:"module"`
	// The dashboard name. Dashboard names must be unique within each tenant.
	Name string `json:"name"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// Whether the dashboard is active or not.
	Isactive *bool `json:"isactive,omitempty"`
	// The version of the dashboard.
	Version *int32 `json:"version,omitempty"`
}
```

A complete dashboard as rendered in POST, PATCH, and GET responses.

#### type DashboardMutable

```go
type DashboardMutable struct {
	// The JSON dashboard definition.
	Definition *string `json:"definition,omitempty"`
	// Whether the dashboard is active or not.
	Isactive *bool `json:"isactive,omitempty"`
	// The module that contains the dashboard.
	Module *string `json:"module,omitempty"`
	// The dashboard name. Dashboard names must be unique within each tenant.
	Name *string `json:"name,omitempty"`
	// The version of the dashboard.
	Version *int32 `json:"version,omitempty"`
}
```

A list of the mutable dashboard fields.

#### type DashboardPatch

```go
type DashboardPatch struct {
	// The JSON dashboard definition.
	Definition *string `json:"definition,omitempty"`
	// Whether the dashboard is active or not.
	Isactive *bool `json:"isactive,omitempty"`
	// The module that contains the dashboard.
	Module *string `json:"module,omitempty"`
	// The dashboard name. Dashboard names must be unique within each tenant.
	Name *string `json:"name,omitempty"`
	// The version of the dashboard.
	Version *int32 `json:"version,omitempty"`
}
```

Values for updating a dashboard using a PATCH request.

#### type DashboardPost

```go
type DashboardPost struct {
	// The JSON dashboard definition.
	Definition string `json:"definition"`
	// The module that contains the dashboard.
	Module string `json:"module"`
	// The dashboard name. Dashboard names must be unique within each tenant.
	Name string `json:"name"`
	// A unique dashboard ID. Random ID used if not provided.
	Id *string `json:"id,omitempty"`
	// Whether the dashboard is active or not.
	Isactive *bool `json:"isactive,omitempty"`
	// The version of the dashboard.
	Version *int32 `json:"version,omitempty"`
}
```

Initial values for creating a new dashboard using a POST request.

#### type Dataset

```go
type Dataset struct {
}
```

Dataset : A complete dataset as rendered in POST, PATCH, and GET responses.

#### func  MakeDatasetFromImportDataset

```go
func MakeDatasetFromImportDataset(f ImportDataset) Dataset
```
MakeDatasetFromImportDataset creates a new Dataset from an instance of
ImportDataset

#### func  MakeDatasetFromIndexDataset

```go
func MakeDatasetFromIndexDataset(f IndexDataset) Dataset
```
MakeDatasetFromIndexDataset creates a new Dataset from an instance of
IndexDataset

#### func  MakeDatasetFromJobDataset

```go
func MakeDatasetFromJobDataset(f JobDataset) Dataset
```
MakeDatasetFromJobDataset creates a new Dataset from an instance of JobDataset

#### func  MakeDatasetFromKvCollectionDataset

```go
func MakeDatasetFromKvCollectionDataset(f KvCollectionDataset) Dataset
```
MakeDatasetFromKvCollectionDataset creates a new Dataset from an instance of
KvCollectionDataset

#### func  MakeDatasetFromLookupDataset

```go
func MakeDatasetFromLookupDataset(f LookupDataset) Dataset
```
MakeDatasetFromLookupDataset creates a new Dataset from an instance of
LookupDataset

#### func  MakeDatasetFromMetricDataset

```go
func MakeDatasetFromMetricDataset(f MetricDataset) Dataset
```
MakeDatasetFromMetricDataset creates a new Dataset from an instance of
MetricDataset

#### func  MakeDatasetFromRawInterface

```go
func MakeDatasetFromRawInterface(f interface{}) Dataset
```
MakeDatasetFromRawInterface creates a new Dataset from a raw interface{}

#### func  MakeDatasetFromViewDataset

```go
func MakeDatasetFromViewDataset(f ViewDataset) Dataset
```
MakeDatasetFromViewDataset creates a new Dataset from an instance of ViewDataset

#### func (Dataset) ImportDataset

```go
func (m Dataset) ImportDataset() *ImportDataset
```
ImportDataset returns ImportDataset if IsImportDataset() is true, nil otherwise

#### func (Dataset) IndexDataset

```go
func (m Dataset) IndexDataset() *IndexDataset
```
IndexDataset returns IndexDataset if IsIndexDataset() is true, nil otherwise

#### func (Dataset) IsImportDataset

```go
func (m Dataset) IsImportDataset() bool
```
IsImportDataset checks if the Dataset is a ImportDataset

#### func (Dataset) IsIndexDataset

```go
func (m Dataset) IsIndexDataset() bool
```
IsIndexDataset checks if the Dataset is a IndexDataset

#### func (Dataset) IsJobDataset

```go
func (m Dataset) IsJobDataset() bool
```
IsJobDataset checks if the Dataset is a JobDataset

#### func (Dataset) IsKvCollectionDataset

```go
func (m Dataset) IsKvCollectionDataset() bool
```
IsKvCollectionDataset checks if the Dataset is a KvCollectionDataset

#### func (Dataset) IsLookupDataset

```go
func (m Dataset) IsLookupDataset() bool
```
IsLookupDataset checks if the Dataset is a LookupDataset

#### func (Dataset) IsMetricDataset

```go
func (m Dataset) IsMetricDataset() bool
```
IsMetricDataset checks if the Dataset is a MetricDataset

#### func (Dataset) IsRawInterface

```go
func (m Dataset) IsRawInterface() bool
```
IsRawInterface checks if the Dataset is an interface{} (unknown type)

#### func (Dataset) IsViewDataset

```go
func (m Dataset) IsViewDataset() bool
```
IsViewDataset checks if the Dataset is a ViewDataset

#### func (Dataset) JobDataset

```go
func (m Dataset) JobDataset() *JobDataset
```
JobDataset returns JobDataset if IsJobDataset() is true, nil otherwise

#### func (Dataset) KvCollectionDataset

```go
func (m Dataset) KvCollectionDataset() *KvCollectionDataset
```
KvCollectionDataset returns KvCollectionDataset if IsKvCollectionDataset() is
true, nil otherwise

#### func (Dataset) LookupDataset

```go
func (m Dataset) LookupDataset() *LookupDataset
```
LookupDataset returns LookupDataset if IsLookupDataset() is true, nil otherwise

#### func (Dataset) MarshalJSON

```go
func (m Dataset) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals Dataset using the appropriate struct field

#### func (Dataset) MetricDataset

```go
func (m Dataset) MetricDataset() *MetricDataset
```
MetricDataset returns MetricDataset if IsMetricDataset() is true, nil otherwise

#### func (Dataset) RawInterface

```go
func (m Dataset) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (*Dataset) UnmarshalJSON

```go
func (m *Dataset) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals Dataset using the "kind" property

#### func (Dataset) ViewDataset

```go
func (m Dataset) ViewDataset() *ViewDataset
```
ViewDataset returns ViewDataset if IsViewDataset() is true, nil otherwise

#### type DatasetCommon

```go
type DatasetCommon struct {
	// A unique dataset ID.
	Id string `json:"id"`
	// The name of the module that contains the dataset.
	Module string `json:"module"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The dataset name qualified by the module name.
	Resourcename string `json:"resourcename"`
	// Detailed description of the dataset.
	Description *string `json:"description,omitempty"`
	// Summary of the dataset's purpose.
	Summary *string `json:"summary,omitempty"`
	// The title of the dataset.  Does not have to be unique.
	Title *string `json:"title,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Properties that are common across all Dataset kinds rendered in POST, PATCH, and
GET responses.

#### type DatasetImportedBy

```go
type DatasetImportedBy struct {
	// The module that is importing the dataset.
	Module string `json:"module"`
	// The dataset name.
	Name string `json:"name"`
}
```


#### type DatasetPatch

```go
type DatasetPatch struct {
}
```

DatasetPatch : Property values to be set in an existing dataset using a PATCH
request. DatasetPatch is ImportDatasetPatch, IndexDatasetPatch, JobDatasetPatch,
KvCollectionDatasetPatch, LookupDatasetPatch, MetricDatasetPatch,
ViewDatasetPatch, (or interface{} if no matches are found)

#### func  MakeDatasetPatchFromImportDatasetPatch

```go
func MakeDatasetPatchFromImportDatasetPatch(f ImportDatasetPatch) DatasetPatch
```
MakeDatasetPatchFromImportDatasetPatch creates a new DatasetPatch from an
instance of ImportDatasetPatch

#### func  MakeDatasetPatchFromIndexDatasetPatch

```go
func MakeDatasetPatchFromIndexDatasetPatch(f IndexDatasetPatch) DatasetPatch
```
MakeDatasetPatchFromIndexDatasetPatch creates a new DatasetPatch from an
instance of IndexDatasetPatch

#### func  MakeDatasetPatchFromJobDatasetPatch

```go
func MakeDatasetPatchFromJobDatasetPatch(f JobDatasetPatch) DatasetPatch
```
MakeDatasetPatchFromJobDatasetPatch creates a new DatasetPatch from an instance
of JobDatasetPatch

#### func  MakeDatasetPatchFromKvCollectionDatasetPatch

```go
func MakeDatasetPatchFromKvCollectionDatasetPatch(f KvCollectionDatasetPatch) DatasetPatch
```
MakeDatasetPatchFromKvCollectionDatasetPatch creates a new DatasetPatch from an
instance of KvCollectionDatasetPatch

#### func  MakeDatasetPatchFromLookupDatasetPatch

```go
func MakeDatasetPatchFromLookupDatasetPatch(f LookupDatasetPatch) DatasetPatch
```
MakeDatasetPatchFromLookupDatasetPatch creates a new DatasetPatch from an
instance of LookupDatasetPatch

#### func  MakeDatasetPatchFromMetricDatasetPatch

```go
func MakeDatasetPatchFromMetricDatasetPatch(f MetricDatasetPatch) DatasetPatch
```
MakeDatasetPatchFromMetricDatasetPatch creates a new DatasetPatch from an
instance of MetricDatasetPatch

#### func  MakeDatasetPatchFromRawInterface

```go
func MakeDatasetPatchFromRawInterface(f interface{}) DatasetPatch
```
MakeDatasetPatchFromRawInterface creates a new DatasetPatch from a raw
interface{}

#### func  MakeDatasetPatchFromViewDatasetPatch

```go
func MakeDatasetPatchFromViewDatasetPatch(f ViewDatasetPatch) DatasetPatch
```
MakeDatasetPatchFromViewDatasetPatch creates a new DatasetPatch from an instance
of ViewDatasetPatch

#### func (DatasetPatch) ImportDatasetPatch

```go
func (m DatasetPatch) ImportDatasetPatch() *ImportDatasetPatch
```
ImportDatasetPatch returns ImportDatasetPatch if IsImportDatasetPatch() is true,
nil otherwise

#### func (DatasetPatch) IndexDatasetPatch

```go
func (m DatasetPatch) IndexDatasetPatch() *IndexDatasetPatch
```
IndexDatasetPatch returns IndexDatasetPatch if IsIndexDatasetPatch() is true,
nil otherwise

#### func (DatasetPatch) IsImportDatasetPatch

```go
func (m DatasetPatch) IsImportDatasetPatch() bool
```
IsImportDatasetPatch checks if the DatasetPatch is a ImportDatasetPatch

#### func (DatasetPatch) IsIndexDatasetPatch

```go
func (m DatasetPatch) IsIndexDatasetPatch() bool
```
IsIndexDatasetPatch checks if the DatasetPatch is a IndexDatasetPatch

#### func (DatasetPatch) IsJobDatasetPatch

```go
func (m DatasetPatch) IsJobDatasetPatch() bool
```
IsJobDatasetPatch checks if the DatasetPatch is a JobDatasetPatch

#### func (DatasetPatch) IsKvCollectionDatasetPatch

```go
func (m DatasetPatch) IsKvCollectionDatasetPatch() bool
```
IsKvCollectionDatasetPatch checks if the DatasetPatch is a
KvCollectionDatasetPatch

#### func (DatasetPatch) IsLookupDatasetPatch

```go
func (m DatasetPatch) IsLookupDatasetPatch() bool
```
IsLookupDatasetPatch checks if the DatasetPatch is a LookupDatasetPatch

#### func (DatasetPatch) IsMetricDatasetPatch

```go
func (m DatasetPatch) IsMetricDatasetPatch() bool
```
IsMetricDatasetPatch checks if the DatasetPatch is a MetricDatasetPatch

#### func (DatasetPatch) IsRawInterface

```go
func (m DatasetPatch) IsRawInterface() bool
```
IsRawInterface checks if the DatasetPatch is an interface{} (unknown type)

#### func (DatasetPatch) IsViewDatasetPatch

```go
func (m DatasetPatch) IsViewDatasetPatch() bool
```
IsViewDatasetPatch checks if the DatasetPatch is a ViewDatasetPatch

#### func (DatasetPatch) JobDatasetPatch

```go
func (m DatasetPatch) JobDatasetPatch() *JobDatasetPatch
```
JobDatasetPatch returns JobDatasetPatch if IsJobDatasetPatch() is true, nil
otherwise

#### func (DatasetPatch) KvCollectionDatasetPatch

```go
func (m DatasetPatch) KvCollectionDatasetPatch() *KvCollectionDatasetPatch
```
KvCollectionDatasetPatch returns KvCollectionDatasetPatch if
IsKvCollectionDatasetPatch() is true, nil otherwise

#### func (DatasetPatch) LookupDatasetPatch

```go
func (m DatasetPatch) LookupDatasetPatch() *LookupDatasetPatch
```
LookupDatasetPatch returns LookupDatasetPatch if IsLookupDatasetPatch() is true,
nil otherwise

#### func (DatasetPatch) MarshalJSON

```go
func (m DatasetPatch) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals DatasetPatch using DatasetPatch.DatasetPatch

#### func (DatasetPatch) MetricDatasetPatch

```go
func (m DatasetPatch) MetricDatasetPatch() *MetricDatasetPatch
```
MetricDatasetPatch returns MetricDatasetPatch if IsMetricDatasetPatch() is true,
nil otherwise

#### func (DatasetPatch) RawInterface

```go
func (m DatasetPatch) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (*DatasetPatch) UnmarshalJSON

```go
func (m *DatasetPatch) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals DatasetPatch into ImportDatasetPatch,
IndexDatasetPatch, JobDatasetPatch, KvCollectionDatasetPatch,
LookupDatasetPatch, MetricDatasetPatch, ViewDatasetPatch, or interface{} if no
matches are found

#### func (DatasetPatch) ViewDatasetPatch

```go
func (m DatasetPatch) ViewDatasetPatch() *ViewDatasetPatch
```
ViewDatasetPatch returns ViewDatasetPatch if IsViewDatasetPatch() is true, nil
otherwise

#### type DatasetPatchCommon

```go
type DatasetPatchCommon struct {
	// The name of module to reassign dataset into.
	Module *string `json:"module,omitempty"`
	// The dataset name. Dataset names must be unique within each module.
	Name *string `json:"name,omitempty"`
	// The name of the dataset owner. This value is obtained from the bearer token.
	Owner *string `json:"owner,omitempty"`
}
```

Properties that are common across all Dataset kinds to be set in an existing
dataset using a PATCH request.

#### type DatasetPost

```go
type DatasetPost struct {
}
```

DatasetPost : Initial property values for creating a new action using a POST
request.

#### func  MakeDatasetPostFromImportDatasetPost

```go
func MakeDatasetPostFromImportDatasetPost(f ImportDatasetPost) DatasetPost
```
MakeDatasetPostFromImportDatasetPost creates a new DatasetPost from an instance
of ImportDatasetPost

#### func  MakeDatasetPostFromIndexDatasetPost

```go
func MakeDatasetPostFromIndexDatasetPost(f IndexDatasetPost) DatasetPost
```
MakeDatasetPostFromIndexDatasetPost creates a new DatasetPost from an instance
of IndexDatasetPost

#### func  MakeDatasetPostFromJobDatasetPost

```go
func MakeDatasetPostFromJobDatasetPost(f JobDatasetPost) DatasetPost
```
MakeDatasetPostFromJobDatasetPost creates a new DatasetPost from an instance of
JobDatasetPost

#### func  MakeDatasetPostFromKvCollectionDatasetPost

```go
func MakeDatasetPostFromKvCollectionDatasetPost(f KvCollectionDatasetPost) DatasetPost
```
MakeDatasetPostFromKvCollectionDatasetPost creates a new DatasetPost from an
instance of KvCollectionDatasetPost

#### func  MakeDatasetPostFromLookupDatasetPost

```go
func MakeDatasetPostFromLookupDatasetPost(f LookupDatasetPost) DatasetPost
```
MakeDatasetPostFromLookupDatasetPost creates a new DatasetPost from an instance
of LookupDatasetPost

#### func  MakeDatasetPostFromMetricDatasetPost

```go
func MakeDatasetPostFromMetricDatasetPost(f MetricDatasetPost) DatasetPost
```
MakeDatasetPostFromMetricDatasetPost creates a new DatasetPost from an instance
of MetricDatasetPost

#### func  MakeDatasetPostFromRawInterface

```go
func MakeDatasetPostFromRawInterface(f interface{}) DatasetPost
```
MakeDatasetPostFromRawInterface creates a new DatasetPost from a raw interface{}

#### func  MakeDatasetPostFromViewDatasetPost

```go
func MakeDatasetPostFromViewDatasetPost(f ViewDatasetPost) DatasetPost
```
MakeDatasetPostFromViewDatasetPost creates a new DatasetPost from an instance of
ViewDatasetPost

#### func (DatasetPost) ImportDatasetPost

```go
func (m DatasetPost) ImportDatasetPost() *ImportDatasetPost
```
ImportDatasetPost returns ImportDatasetPost if IsImportDatasetPost() is true,
nil otherwise

#### func (DatasetPost) IndexDatasetPost

```go
func (m DatasetPost) IndexDatasetPost() *IndexDatasetPost
```
IndexDatasetPost returns IndexDatasetPost if IsIndexDatasetPost() is true, nil
otherwise

#### func (DatasetPost) IsImportDatasetPost

```go
func (m DatasetPost) IsImportDatasetPost() bool
```
IsImportDatasetPost checks if the DatasetPost is a ImportDatasetPost

#### func (DatasetPost) IsIndexDatasetPost

```go
func (m DatasetPost) IsIndexDatasetPost() bool
```
IsIndexDatasetPost checks if the DatasetPost is a IndexDatasetPost

#### func (DatasetPost) IsJobDatasetPost

```go
func (m DatasetPost) IsJobDatasetPost() bool
```
IsJobDatasetPost checks if the DatasetPost is a JobDatasetPost

#### func (DatasetPost) IsKvCollectionDatasetPost

```go
func (m DatasetPost) IsKvCollectionDatasetPost() bool
```
IsKvCollectionDatasetPost checks if the DatasetPost is a KvCollectionDatasetPost

#### func (DatasetPost) IsLookupDatasetPost

```go
func (m DatasetPost) IsLookupDatasetPost() bool
```
IsLookupDatasetPost checks if the DatasetPost is a LookupDatasetPost

#### func (DatasetPost) IsMetricDatasetPost

```go
func (m DatasetPost) IsMetricDatasetPost() bool
```
IsMetricDatasetPost checks if the DatasetPost is a MetricDatasetPost

#### func (DatasetPost) IsRawInterface

```go
func (m DatasetPost) IsRawInterface() bool
```
IsRawInterface checks if the DatasetPost is an interface{} (unknown type)

#### func (DatasetPost) IsViewDatasetPost

```go
func (m DatasetPost) IsViewDatasetPost() bool
```
IsViewDatasetPost checks if the DatasetPost is a ViewDatasetPost

#### func (DatasetPost) JobDatasetPost

```go
func (m DatasetPost) JobDatasetPost() *JobDatasetPost
```
JobDatasetPost returns JobDatasetPost if IsJobDatasetPost() is true, nil
otherwise

#### func (DatasetPost) KvCollectionDatasetPost

```go
func (m DatasetPost) KvCollectionDatasetPost() *KvCollectionDatasetPost
```
KvCollectionDatasetPost returns KvCollectionDatasetPost if
IsKvCollectionDatasetPost() is true, nil otherwise

#### func (DatasetPost) LookupDatasetPost

```go
func (m DatasetPost) LookupDatasetPost() *LookupDatasetPost
```
LookupDatasetPost returns LookupDatasetPost if IsLookupDatasetPost() is true,
nil otherwise

#### func (DatasetPost) MarshalJSON

```go
func (m DatasetPost) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals DatasetPost using the appropriate struct field

#### func (DatasetPost) MetricDatasetPost

```go
func (m DatasetPost) MetricDatasetPost() *MetricDatasetPost
```
MetricDatasetPost returns MetricDatasetPost if IsMetricDatasetPost() is true,
nil otherwise

#### func (DatasetPost) RawInterface

```go
func (m DatasetPost) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (*DatasetPost) UnmarshalJSON

```go
func (m *DatasetPost) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals DatasetPost using the "kind" property

#### func (DatasetPost) ViewDatasetPost

```go
func (m DatasetPost) ViewDatasetPost() *ViewDatasetPost
```
ViewDatasetPost returns ViewDatasetPost if IsViewDatasetPost() is true, nil
otherwise

#### type DatasetPostCommon

```go
type DatasetPostCommon struct {
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The fields to be associated with this dataset.
	Fields []FieldPost `json:"fields,omitempty"`
	// A unique dataset ID. Random ID used if not provided.
	Id *string `json:"id,omitempty"`
	// The name of the module to create the new dataset in.
	Module *string `json:"module,omitempty"`
}
```

Properties that are common across all Dataset kinds for creating a new dataset
using a POST request.

#### type DateMetadataProperties

```go
type DateMetadataProperties struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The date and time object was modified.
	Modified string `json:"modified"`
}
```

Created and Modified date-time properties for inclusion in other objects.

#### type EvalAction

```go
type EvalAction struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// The EVAL expression that calculates the field.
	Expression string `json:"expression"`
	// The name of the field that is added or modified by the EVAL expression.
	Field string `json:"field"`
	// A unique action ID.
	Id   string         `json:"id"`
	Kind EvalActionKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The rule that this action is part of.
	Ruleid string `json:"ruleid"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete eval action as rendered in POST, PATCH, and GET responses.

#### type EvalActionKind

```go
type EvalActionKind string
```

EvalActionKind : The eval action kind.

```go
const (
	EvalActionKindEval EvalActionKind = "EVAL"
)
```
List of EvalActionKind

#### type EvalActionPatch

```go
type EvalActionPatch struct {
	// The EVAL expression that calculates the field.
	Expression *string `json:"expression,omitempty"`
	// The name of the field that is added or modified by the EVAL expression.
	Field *string         `json:"field,omitempty"`
	Kind  *EvalActionKind `json:"kind,omitempty"`
	// The name of the user who owns this action. This value is obtained from the bearer token if not present.
	Owner *string `json:"owner,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Property values for setting existing eval actions using a PATCH request.

#### type EvalActionPost

```go
type EvalActionPost struct {
	// The EVAL expression that calculates the field.
	Expression string `json:"expression"`
	// The name of the field that is added or modified by the EVAL expression.
	Field string         `json:"field"`
	Kind  EvalActionKind `json:"kind"`
	// A unique action ID.
	Id *string `json:"id,omitempty"`
	// The rule that this action is part of.
	Ruleid *string `json:"ruleid,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Initial property values for creating a new eval action using a POST request.

#### type EvalActionProperties

```go
type EvalActionProperties struct {
	// The EVAL expression that calculates the field.
	Expression *string `json:"expression,omitempty"`
	// The name of the field that is added or modified by the EVAL expression.
	Field *string         `json:"field,omitempty"`
	Kind  *EvalActionKind `json:"kind,omitempty"`
}
```

Properties of eval actions which may be read, set, and changed through the API.
Implementation detail of ActionPOST, ActionPOST, and Action, do not use
directly.

#### type Field

```go
type Field struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The dataset that the field is part of.
	Datasetid string        `json:"datasetid"`
	Datatype  FieldDataType `json:"datatype"`
	Fieldtype FieldType     `json:"fieldtype"`
	// The unique ID of this field.
	Id string `json:"id"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The field name.
	Name       string          `json:"name"`
	Prevalence FieldPrevalence `json:"prevalence"`
	// The field description.
	Description *string `json:"description,omitempty"`
	// Whether or not the field has been indexed.
	Indexed *bool `json:"indexed,omitempty"`
	// The field summary.
	Summary *string `json:"summary,omitempty"`
	// The field title.
	Title *string `json:"title,omitempty"`
}
```

A complete field as rendered in POST, PATCH, and GET responses.

#### type FieldDataType

```go
type FieldDataType string
```

FieldDataType : The type of data in the field. Must be one of the valid values.

```go
const (
	FieldDataTypeDate     FieldDataType = "DATE"
	FieldDataTypeNumber   FieldDataType = "NUMBER"
	FieldDataTypeObjectId FieldDataType = "OBJECT_ID"
	FieldDataTypeString   FieldDataType = "STRING"
	FieldDataTypeUnknown  FieldDataType = "UNKNOWN"
)
```
List of FieldDataType

#### type FieldPatch

```go
type FieldPatch struct {
	Datatype *FieldDataType `json:"datatype,omitempty"`
	// The field description.
	Description *string    `json:"description,omitempty"`
	Fieldtype   *FieldType `json:"fieldtype,omitempty"`
	// Whether or not the field has been indexed.
	Indexed *bool `json:"indexed,omitempty"`
	// The field name.
	Name       *string          `json:"name,omitempty"`
	Prevalence *FieldPrevalence `json:"prevalence,omitempty"`
	// The field summary.
	Summary *string `json:"summary,omitempty"`
	// The field title.
	Title *string `json:"title,omitempty"`
}
```

Property values to be set in an existing field using a PATCH request.

#### type FieldPost

```go
type FieldPost struct {
	// The field name.
	Name     string         `json:"name"`
	Datatype *FieldDataType `json:"datatype,omitempty"`
	// The field description.
	Description *string    `json:"description,omitempty"`
	Fieldtype   *FieldType `json:"fieldtype,omitempty"`
	// Whether or not the field has been indexed.
	Indexed    *bool            `json:"indexed,omitempty"`
	Prevalence *FieldPrevalence `json:"prevalence,omitempty"`
	// The field summary.
	Summary *string `json:"summary,omitempty"`
	// The field title.
	Title *string `json:"title,omitempty"`
}
```

Initial property values for creating a new field using a POST request.

#### type FieldPrevalence

```go
type FieldPrevalence string
```

FieldPrevalence : How frequent the field appears in the dataset. Must be one of
the valid values.

```go
const (
	FieldPrevalenceAll     FieldPrevalence = "ALL"
	FieldPrevalenceSome    FieldPrevalence = "SOME"
	FieldPrevalenceUnknown FieldPrevalence = "UNKNOWN"
)
```
List of FieldPrevalence

#### type FieldProperties

```go
type FieldProperties struct {
	Datatype *FieldDataType `json:"datatype,omitempty"`
	// The field description.
	Description *string    `json:"description,omitempty"`
	Fieldtype   *FieldType `json:"fieldtype,omitempty"`
	// Whether or not the field has been indexed.
	Indexed *bool `json:"indexed,omitempty"`
	// The field name.
	Name       *string          `json:"name,omitempty"`
	Prevalence *FieldPrevalence `json:"prevalence,omitempty"`
	// The field summary.
	Summary *string `json:"summary,omitempty"`
	// The field title.
	Title *string `json:"title,omitempty"`
}
```

Properties of fields which can be read, set, and changed through the API.
Implementation detail of FieldPOST, FieldPATCH, and Field, do not use directly.

#### type FieldType

```go
type FieldType string
```

FieldType : The type of field. Must be one of the valid values.

```go
const (
	FieldTypeDimension FieldType = "DIMENSION"
	FieldTypeMeasure   FieldType = "MEASURE"
	FieldTypeUnknown   FieldType = "UNKNOWN"
)
```
List of FieldType

#### type ImportDataset

```go
type ImportDataset struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// A unique dataset ID.
	Id   string            `json:"id"`
	Kind ImportDatasetKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the module that contains the dataset.
	Module string `json:"module"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The dataset name qualified by the module name.
	Resourcename string `json:"resourcename"`
	// The dataset module being imported.
	SourceModule string `json:"sourceModule"`
	// The dataset name being imported.
	SourceName string `json:"sourceName"`
	// Detailed description of the dataset.
	Description *string `json:"description,omitempty"`
	// Summary of the dataset's purpose.
	Summary *string `json:"summary,omitempty"`
	// The title of the dataset.  Does not have to be unique.
	Title *string `json:"title,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete import dataset as rendered in POST, PATCH, and GET responses.

#### type ImportDatasetByIdPost

```go
type ImportDatasetByIdPost struct {
	Kind ImportDatasetKind `json:"kind"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The dataset ID being imported.
	SourceId string `json:"sourceId"`
	// The fields to be associated with this dataset.
	Fields []FieldPost `json:"fields,omitempty"`
	// A unique dataset ID. Random ID used if not provided.
	Id *string `json:"id,omitempty"`
	// The name of the module to create the new dataset in.
	Module *string `json:"module,omitempty"`
}
```

Initial property values for creating a new import dataset by sourceId using a
POST request.

#### type ImportDatasetByIdProperties

```go
type ImportDatasetByIdProperties struct {
	Kind *ImportDatasetKind `json:"kind,omitempty"`
	// The dataset ID being imported.
	SourceId *string `json:"sourceId,omitempty"`
}
```

Initial property values for creating a new import dataset by sourceId using a
POST request.

#### type ImportDatasetByNamePost

```go
type ImportDatasetByNamePost struct {
	Kind ImportDatasetKind `json:"kind"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The dataset module being imported.
	SourceModule string `json:"sourceModule"`
	// The dataset name being imported.
	SourceName string `json:"sourceName"`
	// The fields to be associated with this dataset.
	Fields []FieldPost `json:"fields,omitempty"`
	// A unique dataset ID. Random ID used if not provided.
	Id *string `json:"id,omitempty"`
	// The name of the module to create the new dataset in.
	Module *string `json:"module,omitempty"`
}
```

Initial property values for creating a new import dataset by sourceName and
sourceModule using a POST request.

#### type ImportDatasetByNameProperties

```go
type ImportDatasetByNameProperties struct {
	Kind *ImportDatasetKind `json:"kind,omitempty"`
	// The dataset module being imported.
	SourceModule *string `json:"sourceModule,omitempty"`
	// The dataset name being imported.
	SourceName *string `json:"sourceName,omitempty"`
}
```

Properties of import datasets which may be read, set, and changed through the
API. Implementation detail of DatasetPOST, DatasetPATCH, and Dataset, do not use
directly.

#### type ImportDatasetKind

```go
type ImportDatasetKind string
```

ImportDatasetKind : The dataset kind.

```go
const (
	ImportDatasetKindModelImport ImportDatasetKind = "import"
)
```
List of ImportDatasetKind

#### type ImportDatasetPatch

```go
type ImportDatasetPatch struct {
	// The name of module to reassign dataset into.
	Module *string `json:"module,omitempty"`
	// The dataset name. Dataset names must be unique within each module.
	Name *string `json:"name,omitempty"`
	// The name of the dataset owner. This value is obtained from the bearer token.
	Owner *string `json:"owner,omitempty"`
}
```

Property values to be set in an existing import dataset using a PATCH request.

#### type ImportDatasetPost

```go
type ImportDatasetPost struct {
}
```

ImportDatasetPost : Initial property values for creating a new import dataset
using a POST request. ImportDatasetPost is ImportDatasetByIdPost,
ImportDatasetByNamePost, (or interface{} if no matches are found)

#### func  MakeImportDatasetPostFromImportDatasetByIdPost

```go
func MakeImportDatasetPostFromImportDatasetByIdPost(f ImportDatasetByIdPost) ImportDatasetPost
```
MakeImportDatasetPostFromImportDatasetByIdPost creates a new ImportDatasetPost
from an instance of ImportDatasetByIdPost

#### func  MakeImportDatasetPostFromImportDatasetByNamePost

```go
func MakeImportDatasetPostFromImportDatasetByNamePost(f ImportDatasetByNamePost) ImportDatasetPost
```
MakeImportDatasetPostFromImportDatasetByNamePost creates a new ImportDatasetPost
from an instance of ImportDatasetByNamePost

#### func  MakeImportDatasetPostFromRawInterface

```go
func MakeImportDatasetPostFromRawInterface(f interface{}) ImportDatasetPost
```
MakeImportDatasetPostFromRawInterface creates a new ImportDatasetPost from a raw
interface{}

#### func (ImportDatasetPost) ImportDatasetByIdPost

```go
func (m ImportDatasetPost) ImportDatasetByIdPost() *ImportDatasetByIdPost
```
ImportDatasetByIdPost returns ImportDatasetByIdPost if IsImportDatasetByIdPost()
is true, nil otherwise

#### func (ImportDatasetPost) ImportDatasetByNamePost

```go
func (m ImportDatasetPost) ImportDatasetByNamePost() *ImportDatasetByNamePost
```
ImportDatasetByNamePost returns ImportDatasetByNamePost if
IsImportDatasetByNamePost() is true, nil otherwise

#### func (ImportDatasetPost) IsImportDatasetByIdPost

```go
func (m ImportDatasetPost) IsImportDatasetByIdPost() bool
```
IsImportDatasetByIdPost checks if the ImportDatasetPost is a
ImportDatasetByIdPost

#### func (ImportDatasetPost) IsImportDatasetByNamePost

```go
func (m ImportDatasetPost) IsImportDatasetByNamePost() bool
```
IsImportDatasetByNamePost checks if the ImportDatasetPost is a
ImportDatasetByNamePost

#### func (ImportDatasetPost) IsRawInterface

```go
func (m ImportDatasetPost) IsRawInterface() bool
```
IsRawInterface checks if the ImportDatasetPost is an interface{} (unknown type)

#### func (ImportDatasetPost) MarshalJSON

```go
func (m ImportDatasetPost) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals ImportDatasetPost using ImportDatasetPost.ImportDatasetPost

#### func (ImportDatasetPost) RawInterface

```go
func (m ImportDatasetPost) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (*ImportDatasetPost) UnmarshalJSON

```go
func (m *ImportDatasetPost) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals ImportDatasetPost into ImportDatasetByIdPost,
ImportDatasetByNamePost, or interface{} if no matches are found

#### type IndexDataset

```go
type IndexDataset struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// Specifies whether or not the Splunk index is disabled.
	Disabled bool `json:"disabled"`
	// A unique dataset ID.
	Id   string           `json:"id"`
	Kind IndexDatasetKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the module that contains the dataset.
	Module string `json:"module"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The dataset name qualified by the module name.
	Resourcename string `json:"resourcename"`
	// Detailed description of the dataset.
	Description *string `json:"description,omitempty"`
	// The frozenTimePeriodInSecs to use for the index
	FrozenTimePeriodInSecs *int32 `json:"frozenTimePeriodInSecs,omitempty"`
	// Summary of the dataset's purpose.
	Summary *string `json:"summary,omitempty"`
	// The title of the dataset.  Does not have to be unique.
	Title *string `json:"title,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete index dataset as rendered in POST, PATCH, and GET responses.

#### type IndexDatasetKind

```go
type IndexDatasetKind string
```

IndexDatasetKind : The dataset kind.

```go
const (
	IndexDatasetKindIndex IndexDatasetKind = "index"
)
```
List of IndexDatasetKind

#### type IndexDatasetPatch

```go
type IndexDatasetPatch struct {
	// Specifies whether or not the Splunk index is disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// The frozenTimePeriodInSecs to use for the index
	FrozenTimePeriodInSecs *int32            `json:"frozenTimePeriodInSecs,omitempty"`
	Kind                   *IndexDatasetKind `json:"kind,omitempty"`
	// The name of module to reassign dataset into.
	Module *string `json:"module,omitempty"`
	// The dataset name. Dataset names must be unique within each module.
	Name *string `json:"name,omitempty"`
	// The name of the dataset owner. This value is obtained from the bearer token.
	Owner *string `json:"owner,omitempty"`
}
```

Property values to be set in an existing index dataset using a PATCH request.

#### type IndexDatasetPost

```go
type IndexDatasetPost struct {
	// Specifies whether or not the Splunk index is disabled.
	Disabled bool             `json:"disabled"`
	Kind     IndexDatasetKind `json:"kind"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The fields to be associated with this dataset.
	Fields []FieldPost `json:"fields,omitempty"`
	// The frozenTimePeriodInSecs to use for the index
	FrozenTimePeriodInSecs *int32 `json:"frozenTimePeriodInSecs,omitempty"`
	// A unique dataset ID. Random ID used if not provided.
	Id *string `json:"id,omitempty"`
	// The name of the module to create the new dataset in.
	Module *string `json:"module,omitempty"`
}
```

Initial property values for creating a new index dataset using a POST request.

#### type IndexDatasetProperties

```go
type IndexDatasetProperties struct {
	// Specifies whether or not the Splunk index is disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// The frozenTimePeriodInSecs to use for the index
	FrozenTimePeriodInSecs *int32            `json:"frozenTimePeriodInSecs,omitempty"`
	Kind                   *IndexDatasetKind `json:"kind,omitempty"`
}
```

Properties of job datasets which may be read, set, and changed through the API.
Implementation detail of DatasetPOST, DatasetPATCH, and Dataset, do not use
directly.

#### type JobDataset

```go
type JobDataset struct {
	// Time that the job was completed
	CompletionTime string `json:"completionTime"`
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// The time the dataset will be available in S3.
	DeleteTime string `json:"deleteTime"`
	// Time that the job was dispatched
	DispatchTime string `json:"dispatchTime"`
	// A unique dataset ID.
	Id   string         `json:"id"`
	Kind JobDatasetKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the module that contains the dataset.
	Module string `json:"module"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// Parameters for the search job, mainly earliest, latest, timezone, and relativeTimeAnchor.
	Parameters map[string]interface{} `json:"parameters"`
	// The SPL query string for the search job.
	Query string `json:"query"`
	// Resolved earliest time for the job
	ResolvedEarliest string `json:"resolvedEarliest"`
	// Resolved latest time for the job
	ResolvedLatest string `json:"resolvedLatest"`
	// The dataset name qualified by the module name.
	Resourcename string `json:"resourcename"`
	// The ID assigned to the search job.
	Sid string `json:"sid"`
	// Was the event summary requested for this searhc job?
	CollectEventSummary *bool `json:"collectEventSummary,omitempty"`
	// Was the field summary requested for this searhc job?
	CollectFieldSummary *bool `json:"collectFieldSummary,omitempty"`
	// Were the time bucketes requested for this searhc job?
	CollectTimeBuckets *bool `json:"collectTimeBuckets,omitempty"`
	// Detailed description of the dataset.
	Description *string `json:"description,omitempty"`
	// The runtime of the search in seconds.
	ExecutionTime *float32 `json:"executionTime,omitempty"`
	// Should the search produce all fields (including those not explicity mentioned in the SPL)?
	ExtractAllFields *bool `json:"extractAllFields,omitempty"`
	// Did the SPL query cause any side effects on a dataset?
	HasSideEffects *bool `json:"hasSideEffects,omitempty"`
	// The maximum number of seconds to run this search before finishing.
	MaxTime *int32 `json:"maxTime,omitempty"`
	// An estimate of how complete the search job is.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// The instantaneous number of results produced by the search job.
	ResultsAvailable *int32 `json:"resultsAvailable,omitempty"`
	// The search head that started this search job.
	SearchHead *string `json:"searchHead,omitempty"`
	// The SPLv2 version of the search job query string.
	Spl *string `json:"spl,omitempty"`
	// The current status of the search job.
	Status *string `json:"status,omitempty"`
	// Summary of the dataset's purpose.
	Summary          *string                               `json:"summary,omitempty"`
	TimelineMetadata *JobDatasetPropertiesTimelineMetadata `json:"timelineMetadata,omitempty"`
	// The title of the dataset.  Does not have to be unique.
	Title *string `json:"title,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete job dataset as rendered in POST, PATCH, and GET responses.

#### type JobDatasetEventSummaryAvailableStatus

```go
type JobDatasetEventSummaryAvailableStatus string
```

JobDatasetEventSummaryAvailableStatus : Availability of event summary.

```go
const (
	JobDatasetEventSummaryAvailableStatusTrue    JobDatasetEventSummaryAvailableStatus = "true"
	JobDatasetEventSummaryAvailableStatusFalse   JobDatasetEventSummaryAvailableStatus = "false"
	JobDatasetEventSummaryAvailableStatusUnknown JobDatasetEventSummaryAvailableStatus = "UNKNOWN"
)
```
List of JobDatasetEventSummaryAvailableStatus

#### type JobDatasetFieldSummaryAvailableStatus

```go
type JobDatasetFieldSummaryAvailableStatus string
```

JobDatasetFieldSummaryAvailableStatus : Availability of field summary.

```go
const (
	JobDatasetFieldSummaryAvailableStatusTrue    JobDatasetFieldSummaryAvailableStatus = "true"
	JobDatasetFieldSummaryAvailableStatusFalse   JobDatasetFieldSummaryAvailableStatus = "false"
	JobDatasetFieldSummaryAvailableStatusUnknown JobDatasetFieldSummaryAvailableStatus = "UNKNOWN"
)
```
List of JobDatasetFieldSummaryAvailableStatus

#### type JobDatasetKind

```go
type JobDatasetKind string
```

JobDatasetKind : The dataset kind.

```go
const (
	JobDatasetKindJob JobDatasetKind = "job"
)
```
List of JobDatasetKind

#### type JobDatasetPatch

```go
type JobDatasetPatch struct {
	// Was the event summary requested for this searhc job?
	CollectEventSummary *bool `json:"collectEventSummary,omitempty"`
	// Was the field summary requested for this searhc job?
	CollectFieldSummary *bool `json:"collectFieldSummary,omitempty"`
	// Were the time bucketes requested for this searhc job?
	CollectTimeBuckets *bool `json:"collectTimeBuckets,omitempty"`
	// Time that the job was completed
	CompletionTime *string `json:"completionTime,omitempty"`
	// The time the dataset will be available in S3.
	DeleteTime *string `json:"deleteTime,omitempty"`
	// Time that the job was dispatched
	DispatchTime *string `json:"dispatchTime,omitempty"`
	// The runtime of the search in seconds.
	ExecutionTime *float32 `json:"executionTime,omitempty"`
	// Should the search produce all fields (including those not explicity mentioned in the SPL)?
	ExtractAllFields *bool `json:"extractAllFields,omitempty"`
	// Did the SPL query cause any side effects on a dataset?
	HasSideEffects *bool           `json:"hasSideEffects,omitempty"`
	Kind           *JobDatasetKind `json:"kind,omitempty"`
	// The maximum number of seconds to run this search before finishing.
	MaxTime *int32 `json:"maxTime,omitempty"`
	// The name of module to reassign dataset into.
	Module *string `json:"module,omitempty"`
	// The dataset name. Dataset names must be unique within each module.
	Name *string `json:"name,omitempty"`
	// The name of the dataset owner. This value is obtained from the bearer token.
	Owner *string `json:"owner,omitempty"`
	// Parameters for the search job, mainly earliest, latest, timezone, and relativeTimeAnchor.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// An estimate of how complete the search job is.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// The SPL query string for the search job.
	Query *string `json:"query,omitempty"`
	// Resolved earliest time for the job
	ResolvedEarliest *string `json:"resolvedEarliest,omitempty"`
	// Resolved latest time for the job
	ResolvedLatest *string `json:"resolvedLatest,omitempty"`
	// The instantaneous number of results produced by the search job.
	ResultsAvailable *int32 `json:"resultsAvailable,omitempty"`
	// The search head that started this search job.
	SearchHead *string `json:"searchHead,omitempty"`
	// The ID assigned to the search job.
	Sid *string `json:"sid,omitempty"`
	// The SPLv2 version of the search job query string.
	Spl *string `json:"spl,omitempty"`
	// The current status of the search job.
	Status           *string                               `json:"status,omitempty"`
	TimelineMetadata *JobDatasetPropertiesTimelineMetadata `json:"timelineMetadata,omitempty"`
}
```

Property values to be set in an existing job dataset using a PATCH request.

#### type JobDatasetPost

```go
type JobDatasetPost struct {
	// Time that the job was completed
	CompletionTime string `json:"completionTime"`
	// The time the dataset will be available in S3.
	DeleteTime string `json:"deleteTime"`
	// Time that the job was dispatched
	DispatchTime string         `json:"dispatchTime"`
	Kind         JobDatasetKind `json:"kind"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// Parameters for the search job, mainly earliest, latest, timezone, and relativeTimeAnchor.
	Parameters map[string]interface{} `json:"parameters"`
	// The SPL query string for the search job.
	Query string `json:"query"`
	// Resolved earliest time for the job
	ResolvedEarliest string `json:"resolvedEarliest"`
	// Resolved latest time for the job
	ResolvedLatest string `json:"resolvedLatest"`
	// The ID assigned to the search job.
	Sid string `json:"sid"`
	// Was the event summary requested for this searhc job?
	CollectEventSummary *bool `json:"collectEventSummary,omitempty"`
	// Was the field summary requested for this searhc job?
	CollectFieldSummary *bool `json:"collectFieldSummary,omitempty"`
	// Were the time bucketes requested for this searhc job?
	CollectTimeBuckets *bool `json:"collectTimeBuckets,omitempty"`
	// The runtime of the search in seconds.
	ExecutionTime *float32 `json:"executionTime,omitempty"`
	// Should the search produce all fields (including those not explicity mentioned in the SPL)?
	ExtractAllFields *bool `json:"extractAllFields,omitempty"`
	// The fields to be associated with this dataset.
	Fields []FieldPost `json:"fields,omitempty"`
	// Did the SPL query cause any side effects on a dataset?
	HasSideEffects *bool `json:"hasSideEffects,omitempty"`
	// A unique dataset ID. Random ID used if not provided.
	Id *string `json:"id,omitempty"`
	// The maximum number of seconds to run this search before finishing.
	MaxTime *int32 `json:"maxTime,omitempty"`
	// The name of the module to create the new dataset in.
	Module *string `json:"module,omitempty"`
	// An estimate of how complete the search job is.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// The instantaneous number of results produced by the search job.
	ResultsAvailable *int32 `json:"resultsAvailable,omitempty"`
	// The search head that started this search job.
	SearchHead *string `json:"searchHead,omitempty"`
	// The SPLv2 version of the search job query string.
	Spl *string `json:"spl,omitempty"`
	// The current status of the search job.
	Status           *string                               `json:"status,omitempty"`
	TimelineMetadata *JobDatasetPropertiesTimelineMetadata `json:"timelineMetadata,omitempty"`
}
```

Initial property values for creating a new job dataset using a POST request.

#### type JobDatasetProperties

```go
type JobDatasetProperties struct {
	// Was the event summary requested for this searhc job?
	CollectEventSummary *bool `json:"collectEventSummary,omitempty"`
	// Was the field summary requested for this searhc job?
	CollectFieldSummary *bool `json:"collectFieldSummary,omitempty"`
	// Were the time bucketes requested for this searhc job?
	CollectTimeBuckets *bool `json:"collectTimeBuckets,omitempty"`
	// Time that the job was completed
	CompletionTime *string `json:"completionTime,omitempty"`
	// The time the dataset will be available in S3.
	DeleteTime *string `json:"deleteTime,omitempty"`
	// Time that the job was dispatched
	DispatchTime *string `json:"dispatchTime,omitempty"`
	// The runtime of the search in seconds.
	ExecutionTime *float32 `json:"executionTime,omitempty"`
	// Should the search produce all fields (including those not explicity mentioned in the SPL)?
	ExtractAllFields *bool `json:"extractAllFields,omitempty"`
	// Did the SPL query cause any side effects on a dataset?
	HasSideEffects *bool           `json:"hasSideEffects,omitempty"`
	Kind           *JobDatasetKind `json:"kind,omitempty"`
	// The maximum number of seconds to run this search before finishing.
	MaxTime *int32 `json:"maxTime,omitempty"`
	// Parameters for the search job, mainly earliest, latest, timezone, and relativeTimeAnchor.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// An estimate of how complete the search job is.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// The SPL query string for the search job.
	Query *string `json:"query,omitempty"`
	// Resolved earliest time for the job
	ResolvedEarliest *string `json:"resolvedEarliest,omitempty"`
	// Resolved latest time for the job
	ResolvedLatest *string `json:"resolvedLatest,omitempty"`
	// The instantaneous number of results produced by the search job.
	ResultsAvailable *int32 `json:"resultsAvailable,omitempty"`
	// The search head that started this search job.
	SearchHead *string `json:"searchHead,omitempty"`
	// The ID assigned to the search job.
	Sid *string `json:"sid,omitempty"`
	// The SPLv2 version of the search job query string.
	Spl *string `json:"spl,omitempty"`
	// The current status of the search job.
	Status           *string                               `json:"status,omitempty"`
	TimelineMetadata *JobDatasetPropertiesTimelineMetadata `json:"timelineMetadata,omitempty"`
}
```

Properties of job datasets which may be read, set, and changed through the API.
Implementation detail of DatasetPOST, DatasetPATCH, and Dataset, do not use
directly.

#### type JobDatasetPropertiesTimelineMetadata

```go
type JobDatasetPropertiesTimelineMetadata struct {
	Auto *JobDatasetPropertiesTimelineMetadataAuto `json:"auto,omitempty"`
}
```

Availability of timeline metadata artifacts.

#### type JobDatasetPropertiesTimelineMetadataAuto

```go
type JobDatasetPropertiesTimelineMetadataAuto struct {
	EventSummaryAvailable *JobDatasetEventSummaryAvailableStatus `json:"eventSummaryAvailable,omitempty"`
	FieldSummaryAvailable *JobDatasetFieldSummaryAvailableStatus `json:"fieldSummaryAvailable,omitempty"`
	TimeBucketsAvailable  *JobDatasetTimeBucketsAvailableStatus  `json:"timeBucketsAvailable,omitempty"`
}
```

Availability of automatic timeline metadata artifacts.

#### type JobDatasetTimeBucketsAvailableStatus

```go
type JobDatasetTimeBucketsAvailableStatus string
```

JobDatasetTimeBucketsAvailableStatus : Availability of time buckets (histogram
of events).

```go
const (
	JobDatasetTimeBucketsAvailableStatusTrue    JobDatasetTimeBucketsAvailableStatus = "true"
	JobDatasetTimeBucketsAvailableStatusFalse   JobDatasetTimeBucketsAvailableStatus = "false"
	JobDatasetTimeBucketsAvailableStatusUnknown JobDatasetTimeBucketsAvailableStatus = "UNKNOWN"
)
```
List of JobDatasetTimeBucketsAvailableStatus

#### type KvCollectionDataset

```go
type KvCollectionDataset struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// A unique dataset ID.
	Id   string                  `json:"id"`
	Kind KvCollectionDatasetKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the module that contains the dataset.
	Module string `json:"module"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The dataset name qualified by the module name.
	Resourcename string `json:"resourcename"`
	// Detailed description of the dataset.
	Description *string `json:"description,omitempty"`
	// Summary of the dataset's purpose.
	Summary *string `json:"summary,omitempty"`
	// The title of the dataset.  Does not have to be unique.
	Title *string `json:"title,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete kvcollection dataset as rendered in POST, PATCH, and GET responses.

#### type KvCollectionDatasetKind

```go
type KvCollectionDatasetKind string
```

KvCollectionDatasetKind : The dataset kind.

```go
const (
	KvCollectionDatasetKindKvcollection KvCollectionDatasetKind = "kvcollection"
)
```
List of KVCollectionDatasetKind

#### type KvCollectionDatasetPatch

```go
type KvCollectionDatasetPatch struct {
	Kind *KvCollectionDatasetKind `json:"kind,omitempty"`
	// The name of module to reassign dataset into.
	Module *string `json:"module,omitempty"`
	// The dataset name. Dataset names must be unique within each module.
	Name *string `json:"name,omitempty"`
	// The name of the dataset owner. This value is obtained from the bearer token.
	Owner *string `json:"owner,omitempty"`
}
```

Property values to be set in an existing kvcollection dataset using a PATCH
request.

#### type KvCollectionDatasetPost

```go
type KvCollectionDatasetPost struct {
	Kind KvCollectionDatasetKind `json:"kind"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The fields to be associated with this dataset.
	Fields []FieldPost `json:"fields,omitempty"`
	// A unique dataset ID. Random ID used if not provided.
	Id *string `json:"id,omitempty"`
	// The name of the module to create the new dataset in.
	Module *string `json:"module,omitempty"`
}
```

Initial property values for creating a new kvcollection dataset using a POST
request.

#### type KvCollectionDatasetProperties

```go
type KvCollectionDatasetProperties struct {
	Kind *KvCollectionDatasetKind `json:"kind,omitempty"`
}
```

Properties of kvcollection datasets which may be read, set, and changed through
the API. Implementation detail of DatasetPOST, DatasetPATCH, and Dataset, do not
use directly.

#### type ListActionsForRuleByIdQueryParams

```go
type ListActionsForRuleByIdQueryParams struct {
	// Filter : A filter query to apply to the actions.
	Filter string `key:"filter"`
}
```

ListActionsForRuleByIdQueryParams represents valid query parameters for the
ListActionsForRuleById operation For convenience
ListActionsForRuleByIdQueryParams can be formed in a single statement, for
example:

    `v := ListActionsForRuleByIdQueryParams{}.SetFilter(...)`

#### func (ListActionsForRuleByIdQueryParams) SetFilter

```go
func (q ListActionsForRuleByIdQueryParams) SetFilter(v string) ListActionsForRuleByIdQueryParams
```

#### type ListActionsForRuleQueryParams

```go
type ListActionsForRuleQueryParams struct {
	// Filter : A filter query to apply to the rule.
	Filter string `key:"filter"`
}
```

ListActionsForRuleQueryParams represents valid query parameters for the
ListActionsForRule operation For convenience ListActionsForRuleQueryParams can
be formed in a single statement, for example:

    `v := ListActionsForRuleQueryParams{}.SetFilter(...)`

#### func (ListActionsForRuleQueryParams) SetFilter

```go
func (q ListActionsForRuleQueryParams) SetFilter(v string) ListActionsForRuleQueryParams
```

#### type ListAnnotationsForDatasetByIdQueryParams

```go
type ListAnnotationsForDatasetByIdQueryParams struct {
	// Filter : A filter query to apply to the annotations.
	Filter string `key:"filter"`
}
```

ListAnnotationsForDatasetByIdQueryParams represents valid query parameters for
the ListAnnotationsForDatasetById operation For convenience
ListAnnotationsForDatasetByIdQueryParams can be formed in a single statement,
for example:

    `v := ListAnnotationsForDatasetByIdQueryParams{}.SetFilter(...)`

#### func (ListAnnotationsForDatasetByIdQueryParams) SetFilter

```go
func (q ListAnnotationsForDatasetByIdQueryParams) SetFilter(v string) ListAnnotationsForDatasetByIdQueryParams
```

#### type ListAnnotationsForDatasetByResourceNameQueryParams

```go
type ListAnnotationsForDatasetByResourceNameQueryParams struct {
	// Filter : A filter query to apply to the annotations.
	Filter string `key:"filter"`
}
```

ListAnnotationsForDatasetByResourceNameQueryParams represents valid query
parameters for the ListAnnotationsForDatasetByResourceName operation For
convenience ListAnnotationsForDatasetByResourceNameQueryParams can be formed in
a single statement, for example:

    `v := ListAnnotationsForDatasetByResourceNameQueryParams{}.SetFilter(...)`

#### func (ListAnnotationsForDatasetByResourceNameQueryParams) SetFilter

```go
func (q ListAnnotationsForDatasetByResourceNameQueryParams) SetFilter(v string) ListAnnotationsForDatasetByResourceNameQueryParams
```

#### type ListDatasetsQueryParams

```go
type ListDatasetsQueryParams struct {
	// Count : The maximum number of results to return.
	Count *int32 `key:"count"`
	// Filter : A filter to apply to the dataset list. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
	// Orderby : A list of fields to order the results by.  You can specify either ascending or descending order using \&quot;&lt;field&gt; asc\&quot; or \&quot;&lt;field&gt; desc.  Ascending order is the default.
	Orderby []string `key:"orderby"`
}
```

ListDatasetsQueryParams represents valid query parameters for the ListDatasets
operation For convenience ListDatasetsQueryParams can be formed in a single
statement, for example:

    `v := ListDatasetsQueryParams{}.SetCount(...).SetFilter(...).SetOrderby(...)`

#### func (ListDatasetsQueryParams) SetCount

```go
func (q ListDatasetsQueryParams) SetCount(v int32) ListDatasetsQueryParams
```

#### func (ListDatasetsQueryParams) SetFilter

```go
func (q ListDatasetsQueryParams) SetFilter(v string) ListDatasetsQueryParams
```

#### func (ListDatasetsQueryParams) SetOrderby

```go
func (q ListDatasetsQueryParams) SetOrderby(v []string) ListDatasetsQueryParams
```

#### type ListFieldsForDatasetByIdQueryParams

```go
type ListFieldsForDatasetByIdQueryParams struct {
	// Filter : A filter to apply to a specified dataset. The filter must be a SPL predicate expression.
	Filter string `key:"filter"`
}
```

ListFieldsForDatasetByIdQueryParams represents valid query parameters for the
ListFieldsForDatasetById operation For convenience
ListFieldsForDatasetByIdQueryParams can be formed in a single statement, for
example:

    `v := ListFieldsForDatasetByIdQueryParams{}.SetFilter(...)`

#### func (ListFieldsForDatasetByIdQueryParams) SetFilter

```go
func (q ListFieldsForDatasetByIdQueryParams) SetFilter(v string) ListFieldsForDatasetByIdQueryParams
```

#### type ListFieldsForDatasetQueryParams

```go
type ListFieldsForDatasetQueryParams struct {
	// Filter : A filter to apply to the dataset.
	Filter string `key:"filter"`
}
```

ListFieldsForDatasetQueryParams represents valid query parameters for the
ListFieldsForDataset operation For convenience ListFieldsForDatasetQueryParams
can be formed in a single statement, for example:

    `v := ListFieldsForDatasetQueryParams{}.SetFilter(...)`

#### func (ListFieldsForDatasetQueryParams) SetFilter

```go
func (q ListFieldsForDatasetQueryParams) SetFilter(v string) ListFieldsForDatasetQueryParams
```

#### type ListFieldsQueryParams

```go
type ListFieldsQueryParams struct {
	// Filter : A filter query to apply to the the field list.
	Filter string `key:"filter"`
}
```

ListFieldsQueryParams represents valid query parameters for the ListFields
operation For convenience ListFieldsQueryParams can be formed in a single
statement, for example:

    `v := ListFieldsQueryParams{}.SetFilter(...)`

#### func (ListFieldsQueryParams) SetFilter

```go
func (q ListFieldsQueryParams) SetFilter(v string) ListFieldsQueryParams
```

#### type ListModulesQueryParams

```go
type ListModulesQueryParams struct {
	// Filter : A filter to apply to the modules.
	Filter string `key:"filter"`
}
```

ListModulesQueryParams represents valid query parameters for the ListModules
operation For convenience ListModulesQueryParams can be formed in a single
statement, for example:

    `v := ListModulesQueryParams{}.SetFilter(...)`

#### func (ListModulesQueryParams) SetFilter

```go
func (q ListModulesQueryParams) SetFilter(v string) ListModulesQueryParams
```

#### type ListRelationshipsQueryParams

```go
type ListRelationshipsQueryParams struct {
	// Filter : A filter to apply to the complete set of relationships.
	Filter string `key:"filter"`
}
```

ListRelationshipsQueryParams represents valid query parameters for the
ListRelationships operation For convenience ListRelationshipsQueryParams can be
formed in a single statement, for example:

    `v := ListRelationshipsQueryParams{}.SetFilter(...)`

#### func (ListRelationshipsQueryParams) SetFilter

```go
func (q ListRelationshipsQueryParams) SetFilter(v string) ListRelationshipsQueryParams
```

#### type ListRulesQueryParams

```go
type ListRulesQueryParams struct {
	// Filter : A query filter to apply to the complete set of rules.
	Filter string `key:"filter"`
}
```

ListRulesQueryParams represents valid query parameters for the ListRules
operation For convenience ListRulesQueryParams can be formed in a single
statement, for example:

    `v := ListRulesQueryParams{}.SetFilter(...)`

#### func (ListRulesQueryParams) SetFilter

```go
func (q ListRulesQueryParams) SetFilter(v string) ListRulesQueryParams
```

#### type ListWorkflowBuildsQueryParams

```go
type ListWorkflowBuildsQueryParams struct {
	// Filter : A filter to apply to the workflow builds. Must be a SPL predicate expression.
	Filter string `key:"filter"`
}
```

ListWorkflowBuildsQueryParams represents valid query parameters for the
ListWorkflowBuilds operation For convenience ListWorkflowBuildsQueryParams can
be formed in a single statement, for example:

    `v := ListWorkflowBuildsQueryParams{}.SetFilter(...)`

#### func (ListWorkflowBuildsQueryParams) SetFilter

```go
func (q ListWorkflowBuildsQueryParams) SetFilter(v string) ListWorkflowBuildsQueryParams
```

#### type ListWorkflowRunsQueryParams

```go
type ListWorkflowRunsQueryParams struct {
	// Filter : A filter to apply to the workflow runs for specified workflow build ID. Must be an SPL predicate expression.
	Filter string `key:"filter"`
}
```

ListWorkflowRunsQueryParams represents valid query parameters for the
ListWorkflowRuns operation For convenience ListWorkflowRunsQueryParams can be
formed in a single statement, for example:

    `v := ListWorkflowRunsQueryParams{}.SetFilter(...)`

#### func (ListWorkflowRunsQueryParams) SetFilter

```go
func (q ListWorkflowRunsQueryParams) SetFilter(v string) ListWorkflowRunsQueryParams
```

#### type ListWorkflowsQueryParams

```go
type ListWorkflowsQueryParams struct {
	// Filter : A filter to apply to the workflows. Must be a SPL predicate expression.
	Filter string `key:"filter"`
}
```

ListWorkflowsQueryParams represents valid query parameters for the ListWorkflows
operation For convenience ListWorkflowsQueryParams can be formed in a single
statement, for example:

    `v := ListWorkflowsQueryParams{}.SetFilter(...)`

#### func (ListWorkflowsQueryParams) SetFilter

```go
func (q ListWorkflowsQueryParams) SetFilter(v string) ListWorkflowsQueryParams
```

#### type LookupAction

```go
type LookupAction struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// The lookup body.
	Expression string `json:"expression"`
	// A unique action ID.
	Id   string           `json:"id"`
	Kind LookupActionKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The rule that this action is part of.
	Ruleid string `json:"ruleid"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete lookup action as rendered in POST, PATCH, and GET responses.

#### type LookupActionKind

```go
type LookupActionKind string
```

LookupActionKind : The lookup action kind.

```go
const (
	LookupActionKindLookup LookupActionKind = "LOOKUP"
)
```
List of LookupActionKind

#### type LookupActionPatch

```go
type LookupActionPatch struct {
	// The lookup body.
	Expression *string           `json:"expression,omitempty"`
	Kind       *LookupActionKind `json:"kind,omitempty"`
	// The name of the user who owns this action. This value is obtained from the bearer token if not present.
	Owner *string `json:"owner,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Property values for setting existing lookup actions using a PATCH request.

#### type LookupActionPost

```go
type LookupActionPost struct {
	// The lookup body.
	Expression string           `json:"expression"`
	Kind       LookupActionKind `json:"kind"`
	// A unique action ID.
	Id *string `json:"id,omitempty"`
	// The rule that this action is part of.
	Ruleid *string `json:"ruleid,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Initial property values for creating a new lookup action using a POST request.

#### type LookupActionProperties

```go
type LookupActionProperties struct {
	// The lookup body.
	Expression *string           `json:"expression,omitempty"`
	Kind       *LookupActionKind `json:"kind,omitempty"`
}
```

Properties of lookup actions which may be read, set, and changed through the
API. Implementation detail of ActionPOST, ActionPOST, and Action, do not use
directly.

#### type LookupDataset

```go
type LookupDataset struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby    string                    `json:"createdby"`
	ExternalKind LookupDatasetExternalKind `json:"externalKind"`
	// The name of the external lookup.
	ExternalName string `json:"externalName"`
	// A unique dataset ID.
	Id   string            `json:"id"`
	Kind LookupDatasetKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the module that contains the dataset.
	Module string `json:"module"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The dataset name qualified by the module name.
	Resourcename string `json:"resourcename"`
	// Match case-sensitively against the lookup.
	CaseSensitiveMatch *bool `json:"caseSensitiveMatch,omitempty"`
	// Detailed description of the dataset.
	Description *string `json:"description,omitempty"`
	// A query that filters results out of the lookup before those results are returned.
	Filter *string `json:"filter,omitempty"`
	// Summary of the dataset's purpose.
	Summary *string `json:"summary,omitempty"`
	// The title of the dataset.  Does not have to be unique.
	Title *string `json:"title,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete lookup dataset as rendered in POST, PATCH, and GET responses.

#### type LookupDatasetExternalKind

```go
type LookupDatasetExternalKind string
```

LookupDatasetExternalKind : The type of the external lookup.

```go
const (
	LookupDatasetExternalKindKvcollection LookupDatasetExternalKind = "kvcollection"
)
```
List of LookupDatasetExternalKind

#### type LookupDatasetKind

```go
type LookupDatasetKind string
```

LookupDatasetKind : The dataset kind.

```go
const (
	LookupDatasetKindLookup LookupDatasetKind = "lookup"
)
```
List of LookupDatasetKind

#### type LookupDatasetPatch

```go
type LookupDatasetPatch struct {
	// Match case-sensitively against the lookup.
	CaseSensitiveMatch *bool                      `json:"caseSensitiveMatch,omitempty"`
	ExternalKind       *LookupDatasetExternalKind `json:"externalKind,omitempty"`
	// The name of the external lookup.
	ExternalName *string `json:"externalName,omitempty"`
	// A query that filters results out of the lookup before those results are returned.
	Filter *string            `json:"filter,omitempty"`
	Kind   *LookupDatasetKind `json:"kind,omitempty"`
	// The name of module to reassign dataset into.
	Module *string `json:"module,omitempty"`
	// The dataset name. Dataset names must be unique within each module.
	Name *string `json:"name,omitempty"`
	// The name of the dataset owner. This value is obtained from the bearer token.
	Owner *string `json:"owner,omitempty"`
}
```

Property values to be set in an existing lookup dataset using a PATCH request.

#### type LookupDatasetPost

```go
type LookupDatasetPost struct {
	ExternalKind LookupDatasetExternalKind `json:"externalKind"`
	// The name of the external lookup.
	ExternalName string            `json:"externalName"`
	Kind         LookupDatasetKind `json:"kind"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// Match case-sensitively against the lookup.
	CaseSensitiveMatch *bool `json:"caseSensitiveMatch,omitempty"`
	// The fields to be associated with this dataset.
	Fields []FieldPost `json:"fields,omitempty"`
	// A query that filters results out of the lookup before those results are returned.
	Filter *string `json:"filter,omitempty"`
	// A unique dataset ID. Random ID used if not provided.
	Id *string `json:"id,omitempty"`
	// The name of the module to create the new dataset in.
	Module *string `json:"module,omitempty"`
}
```

Initial property values for creating a new lookup dataset using a POST request.

#### type LookupDatasetProperties

```go
type LookupDatasetProperties struct {
	// Match case-sensitively against the lookup.
	CaseSensitiveMatch *bool                      `json:"caseSensitiveMatch,omitempty"`
	ExternalKind       *LookupDatasetExternalKind `json:"externalKind,omitempty"`
	// The name of the external lookup.
	ExternalName *string `json:"externalName,omitempty"`
	// A query that filters results out of the lookup before those results are returned.
	Filter *string            `json:"filter,omitempty"`
	Kind   *LookupDatasetKind `json:"kind,omitempty"`
}
```

Properties of lookup datasets which may be read, set, and changed through the
API. Implementation detail of DatasetPOST, DatasetPATCH, and Dataset, do not use
directly.

#### type MetadataProperties

```go
type MetadataProperties struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the object's owner.
	Owner string `json:"owner"`
}
```

Created, createdby, modified, modifiedby, and owner properties for inclusion in
other objects.

#### type MetricDataset

```go
type MetricDataset struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// Specifies whether or not the Splunk index is disabled.
	Disabled bool `json:"disabled"`
	// A unique dataset ID.
	Id   string            `json:"id"`
	Kind MetricDatasetKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the module that contains the dataset.
	Module string `json:"module"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The dataset name qualified by the module name.
	Resourcename string `json:"resourcename"`
	// Detailed description of the dataset.
	Description *string `json:"description,omitempty"`
	// The frozenTimePeriodInSecs to use for the index
	FrozenTimePeriodInSecs *int32 `json:"frozenTimePeriodInSecs,omitempty"`
	// Summary of the dataset's purpose.
	Summary *string `json:"summary,omitempty"`
	// The title of the dataset.  Does not have to be unique.
	Title *string `json:"title,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete metric dataset as rendered in POST, PATCH, and GET responses.

#### type MetricDatasetKind

```go
type MetricDatasetKind string
```

MetricDatasetKind : The dataset kind.

```go
const (
	MetricDatasetKindMetric MetricDatasetKind = "metric"
)
```
List of MetricDatasetKind

#### type MetricDatasetPatch

```go
type MetricDatasetPatch struct {
	// Specifies whether or not the Splunk index is disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// The frozenTimePeriodInSecs to use for the index
	FrozenTimePeriodInSecs *int32             `json:"frozenTimePeriodInSecs,omitempty"`
	Kind                   *MetricDatasetKind `json:"kind,omitempty"`
	// The name of module to reassign dataset into.
	Module *string `json:"module,omitempty"`
	// The dataset name. Dataset names must be unique within each module.
	Name *string `json:"name,omitempty"`
	// The name of the dataset owner. This value is obtained from the bearer token.
	Owner *string `json:"owner,omitempty"`
}
```

Property values to be set in an existing metric dataset using a PATCH request.

#### type MetricDatasetPost

```go
type MetricDatasetPost struct {
	// Specifies whether or not the Splunk index is disabled.
	Disabled bool              `json:"disabled"`
	Kind     MetricDatasetKind `json:"kind"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The fields to be associated with this dataset.
	Fields []FieldPost `json:"fields,omitempty"`
	// The frozenTimePeriodInSecs to use for the index
	FrozenTimePeriodInSecs *int32 `json:"frozenTimePeriodInSecs,omitempty"`
	// A unique dataset ID. Random ID used if not provided.
	Id *string `json:"id,omitempty"`
	// The name of the module to create the new dataset in.
	Module *string `json:"module,omitempty"`
}
```

Initial property values for creating a new metric dataset using a POST request.

#### type MetricDatasetProperties

```go
type MetricDatasetProperties struct {
	// Specifies whether or not the Splunk index is disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// The frozenTimePeriodInSecs to use for the index
	FrozenTimePeriodInSecs *int32             `json:"frozenTimePeriodInSecs,omitempty"`
	Kind                   *MetricDatasetKind `json:"kind,omitempty"`
}
```

Properties of metric datasets which may be read, set, and changed through the
API. Implementation detail of DatasetPOST, DatasetPATCH, and Dataset, do not use
directly.

#### type Module

```go
type Module struct {
	Name *string `json:"name,omitempty"`
}
```

The name of a module.

#### type RegexAction

```go
type RegexAction struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// Name of the field that is matched against the regular expression.
	Field string `json:"field"`
	// A unique action ID.
	Id   string          `json:"id"`
	Kind RegexActionKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// A regular expression that includes named capture groups for the purpose of field extraction.
	Pattern string `json:"pattern"`
	// The rule that this action is part of.
	Ruleid string `json:"ruleid"`
	// The maximum number of times per event to attempt to match fields with the regular expression.
	Limit *int32 `json:"limit,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete regex action as rendered in POST, PATCH, and GET responses.

#### type RegexActionKind

```go
type RegexActionKind string
```

RegexActionKind : The regex action kind.

```go
const (
	RegexActionKindRegex RegexActionKind = "REGEX"
)
```
List of RegexActionKind

#### type RegexActionPatch

```go
type RegexActionPatch struct {
	// Name of the field that is matched against the regular expression.
	Field *string          `json:"field,omitempty"`
	Kind  *RegexActionKind `json:"kind,omitempty"`
	// The maximum number of times per event to attempt to match fields with the regular expression.
	Limit *int32 `json:"limit,omitempty"`
	// The name of the user who owns this action. This value is obtained from the bearer token if not present.
	Owner *string `json:"owner,omitempty"`
	// A regular expression that includes named capture groups for the purpose of field extraction.
	Pattern *string `json:"pattern,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Property values for setting existing regex actions using a PATCH request.

#### type RegexActionPost

```go
type RegexActionPost struct {
	// Name of the field that is matched against the regular expression.
	Field string          `json:"field"`
	Kind  RegexActionKind `json:"kind"`
	// A regular expression that includes named capture groups for the purpose of field extraction.
	Pattern string `json:"pattern"`
	// A unique action ID.
	Id *string `json:"id,omitempty"`
	// The maximum number of times per event to attempt to match fields with the regular expression.
	Limit *int32 `json:"limit,omitempty"`
	// The rule that this action is part of.
	Ruleid *string `json:"ruleid,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Initial property values for creating a new regex action using a POST request.

#### type RegexActionProperties

```go
type RegexActionProperties struct {
	// Name of the field that is matched against the regular expression.
	Field *string          `json:"field,omitempty"`
	Kind  *RegexActionKind `json:"kind,omitempty"`
	// The maximum number of times per event to attempt to match fields with the regular expression.
	Limit *int32 `json:"limit,omitempty"`
	// A regular expression that includes named capture groups for the purpose of field extraction.
	Pattern *string `json:"pattern,omitempty"`
}
```

Properties of regex actions which may be read, set, and changed through the API.
Implementation detail of ActionPOST, ActionPOST, and Action, do not use
directly.

#### type Relationship

```go
type Relationship struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// The relationship fields associated with the relationship.
	Fields []RelationshipField `json:"fields"`
	// A unique relationship ID.
	Id   string           `json:"id"`
	Kind RelationshipKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The module that contains the relationship.
	Module string `json:"module"`
	// The relationship name.
	Name string `json:"name"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// A unique source dataset ID. Either the sourceid or sourceresourcename property must be specified.
	Sourceid string `json:"sourceid"`
	// A unique target dataset ID. Either the targetid or targetresourcename property must be specified.
	Targetid string `json:"targetid"`
	// The Catalog version.
	Version int32 `json:"version"`
	// The source dataset name qualified by module name. Either the sourceid or sourceresourcename property must be specified.
	Sourceresourcename *string `json:"sourceresourcename,omitempty"`
	// The target dataset name qualified by module name. Either the targetid or targetresourcename property must be specified.
	Targetresourcename *string `json:"targetresourcename,omitempty"`
}
```

A complete relationship as rendered in POST, PATCH, and GET responses.

#### type RelationshipField

```go
type RelationshipField struct {
	// The date and time object was created.
	Created string                `json:"created"`
	Kind    RelationshipFieldKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// A unique source dataset ID.
	Sourceid string `json:"sourceid"`
	// A unique target dataset ID.
	Targetid string `json:"targetid"`
	// A unique relationship ID.
	Relationshipid *string `json:"relationshipid,omitempty"`
}
```

A complete relationship field as rendered in POST, PATCH, and GET responses.

#### type RelationshipFieldKind

```go
type RelationshipFieldKind string
```

RelationshipFieldKind : The type of match between the fields. Must be one of the
valid values. The LATEST_BEFORE match type specifies that the datetime field in
one dataset binds to the latest time before the datetime field in another
dataset.

```go
const (
	RelationshipFieldKindExact        RelationshipFieldKind = "EXACT"
	RelationshipFieldKindLatestBefore RelationshipFieldKind = "LATEST_BEFORE"
)
```
List of RelationshipFieldKind

#### type RelationshipFieldPost

```go
type RelationshipFieldPost struct {
	Kind RelationshipFieldKind `json:"kind"`
	// A unique source dataset ID.
	Sourceid string `json:"sourceid"`
	// A unique target dataset ID.
	Targetid string `json:"targetid"`
	// A unique relationship ID.
	Relationshipid *string `json:"relationshipid,omitempty"`
}
```

The properties required to create a new relationship field using a relationship
POST request.

#### type RelationshipFieldProperties

```go
type RelationshipFieldProperties struct {
	Kind *RelationshipFieldKind `json:"kind,omitempty"`
	// A unique relationship ID.
	Relationshipid *string `json:"relationshipid,omitempty"`
	// A unique source dataset ID.
	Sourceid *string `json:"sourceid,omitempty"`
	// A unique target dataset ID.
	Targetid *string `json:"targetid,omitempty"`
}
```

Properties of relationship fields which are read through the API. Implementation
detail of RelationshipFieldPOST. Do not use directly.

#### type RelationshipKind

```go
type RelationshipKind string
```

RelationshipKind : The relationship type. Must be one of the valid values.

```go
const (
	RelationshipKindOne        RelationshipKind = "ONE"
	RelationshipKindMany       RelationshipKind = "MANY"
	RelationshipKindDependency RelationshipKind = "DEPENDENCY"
)
```
List of RelationshipKind

#### type RelationshipPatch

```go
type RelationshipPatch struct {
	// The name of the relationship.
	Name *string `json:"name,omitempty"`
	// The user who is the owner of the relationship.
	Owner *string `json:"owner,omitempty"`
}
```


#### type RelationshipPost

```go
type RelationshipPost struct {
	Kind RelationshipKind `json:"kind"`
	// The relationship name.
	Name string `json:"name"`
	// The fields associated with this relationship.
	Fields []RelationshipFieldPost `json:"fields,omitempty"`
	// A unique relationship ID. If not specified, an auto generated ID is created.
	Id *string `json:"id,omitempty"`
	// The module that contains the relationship.
	Module *string `json:"module,omitempty"`
	// A unique source dataset ID. Either the sourceid or sourceresourcename property must be specified.
	Sourceid *string `json:"sourceid,omitempty"`
	// The source dataset name qualified by module name. Either the sourceid or sourceresourcename property must be specified.
	Sourceresourcename *string `json:"sourceresourcename,omitempty"`
	// A unique target dataset ID. Either the targetid or targetresourcename property must be specified.
	Targetid *string `json:"targetid,omitempty"`
	// The target dataset name qualified by module name. Either the targetid or targetresourcename property must be specified.
	Targetresourcename *string `json:"targetresourcename,omitempty"`
	// The Catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

The properties required to create a new relationship using a POST request.

#### type RelationshipProperties

```go
type RelationshipProperties struct {
	Kind *RelationshipKind `json:"kind,omitempty"`
	// The module that contains the relationship.
	Module *string `json:"module,omitempty"`
	// The relationship name.
	Name *string `json:"name,omitempty"`
	// A unique source dataset ID. Either the sourceid or sourceresourcename property must be specified.
	Sourceid *string `json:"sourceid,omitempty"`
	// The source dataset name qualified by module name. Either the sourceid or sourceresourcename property must be specified.
	Sourceresourcename *string `json:"sourceresourcename,omitempty"`
	// A unique target dataset ID. Either the targetid or targetresourcename property must be specified.
	Targetid *string `json:"targetid,omitempty"`
	// The target dataset name qualified by module name. Either the targetid or targetresourcename property must be specified.
	Targetresourcename *string `json:"targetresourcename,omitempty"`
	// The Catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Properties of relationships which are read through the API. Implementation
detail of RelationshipPOST, RelationshipPATCH and Relationship. Do not use
directly.

#### type Rule

```go
type Rule struct {
	// The actions associated with the rule.
	Actions []Action `json:"actions"`
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// A unique Rule ID.
	Id string `json:"id"`
	// The rule match type.
	Match string `json:"match"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The module containing the rule.
	Module string `json:"module"`
	// The rule name.
	Name string `json:"name"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The rule name qualified by the module name.
	Resourcename string `json:"resourcename"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete rule as rendered in POST, PATCH, and GET responses.

#### type RulePatch

```go
type RulePatch struct {
	// The rule match type.
	Match *string `json:"match,omitempty"`
	// The module containing the rule.
	Module *string `json:"module,omitempty"`
	// The rule name.
	Name *string `json:"name,omitempty"`
	// The name of the user who owns the rule.
	Owner *string `json:"owner,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```


#### type RulePost

```go
type RulePost struct {
	// The rule match type.
	Match string `json:"match"`
	// The rule name.
	Name string `json:"name"`
	// The actions to be associated with this rule.
	Actions []ActionPost `json:"actions,omitempty"`
	// A unique rule ID. The newly created rule object will use this ID value if provided.
	Id *string `json:"id,omitempty"`
	// The module containing the rule.
	Module *string `json:"module,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Initial property values for creating a new rule using a POST request.

#### type RuleProperties

```go
type RuleProperties struct {
	// The rule match type.
	Match *string `json:"match,omitempty"`
	// The module containing the rule.
	Module *string `json:"module,omitempty"`
	// The rule name.
	Name *string `json:"name,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

Properties of rules which may be read, set, and changed through the API.
Implementation detail of RulePOST, RulePATCH, and Rule, do not use directly.

#### type Service

```go
type Service services.BaseService
```


#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new catalog service client from the given Config

#### func (*Service) CreateActionForRule

```go
func (s *Service) CreateActionForRule(ruleresourcename string, actionPost ActionPost, resp ...*http.Response) (*Action, error)
```
CreateActionForRule - catalog service endpoint Create a new action for a rule
associated with a specific resource name. Parameters:

    ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
    actionPost: The JSON representation of the action to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateActionForRuleById

```go
func (s *Service) CreateActionForRuleById(ruleid string, actionPost ActionPost, resp ...*http.Response) (*Action, error)
```
CreateActionForRuleById - catalog service endpoint Create a new action for a
specific rule. Parameters:

    ruleid: ID of a Field.
    actionPost: The JSON representation of the action to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateAnnotationForDatasetById

```go
func (s *Service) CreateAnnotationForDatasetById(datasetid string, annotationPost AnnotationPost, resp ...*http.Response) (*Annotation, error)
```
CreateAnnotationForDatasetById - catalog service endpoint Create a new
annotation for a specific dataset. Parameters:

    datasetid: ID of a Dataset.
    annotationPost: The JSON representation of the annotation to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateAnnotationForDatasetByResourceName

```go
func (s *Service) CreateAnnotationForDatasetByResourceName(datasetresourcename string, annotationPost AnnotationPost, resp ...*http.Response) (*Annotation, error)
```
CreateAnnotationForDatasetByResourceName - catalog service endpoint Create a new
annotation for a specific dataset. Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    annotationPost: The JSON representation of the annotation to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateDashboard

```go
func (s *Service) CreateDashboard(dashboardPost DashboardPost, resp ...*http.Response) (*Dashboard, error)
```
CreateDashboard - catalog service endpoint Create a new dashboard. Parameters:

    dashboardPost: The JSON representation of the Dashboard to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateDataset

```go
func (s *Service) CreateDataset(datasetPost DatasetPost, resp ...*http.Response) (*Dataset, error)
```
CreateDataset - catalog service endpoint Create a new dataset. Parameters:

    datasetPost: JSON representation of the DatasetInfo to be persisted
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateDatasetImport

```go
func (s *Service) CreateDatasetImport(datasetresourcename string, datasetImportedBy DatasetImportedBy, resp ...*http.Response) (*ImportDataset, error)
```
CreateDatasetImport - catalog service endpoint Create a new dataset import.
Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    datasetImportedBy
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateDatasetImportById

```go
func (s *Service) CreateDatasetImportById(datasetid string, datasetImportedBy DatasetImportedBy, resp ...*http.Response) (*ImportDataset, error)
```
CreateDatasetImportById - catalog service endpoint Create a new dataset import.
Parameters:

    datasetid: ID of a Dataset.
    datasetImportedBy
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateFieldForDataset

```go
func (s *Service) CreateFieldForDataset(datasetresourcename string, fieldPost FieldPost, resp ...*http.Response) (*Field, error)
```
CreateFieldForDataset - catalog service endpoint Create a new field on a
specific dataset. Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    fieldPost: The JSON representation of the field to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateFieldForDatasetById

```go
func (s *Service) CreateFieldForDatasetById(datasetid string, fieldPost FieldPost, resp ...*http.Response) (*Field, error)
```
CreateFieldForDatasetById - catalog service endpoint Add a new field to a
dataset. Parameters:

    datasetid: ID of a Dataset.
    fieldPost: The JSON representation of the field to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateRelationship

```go
func (s *Service) CreateRelationship(relationshipPost RelationshipPost, resp ...*http.Response) (*Relationship, error)
```
CreateRelationship - catalog service endpoint Create a new relationship.
Parameters:

    relationshipPost: The JSON representation of the relationship to persist.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateRule

```go
func (s *Service) CreateRule(rulePost RulePost, resp ...*http.Response) (*Rule, error)
```
CreateRule - catalog service endpoint Create a new rule. Parameters:

    rulePost: The JSON representation of the rule to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateWorkflow

```go
func (s *Service) CreateWorkflow(workflowPost WorkflowPost, resp ...*http.Response) (*Workflow, error)
```
CreateWorkflow - catalog service endpoint Create a new workflow configuration.
Parameters:

    workflowPost: The JSON representation of the workflow to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateWorkflowBuild

```go
func (s *Service) CreateWorkflowBuild(workflowid string, workflowBuildPost WorkflowBuildPost, resp ...*http.Response) (*WorkflowBuild, error)
```
CreateWorkflowBuild - catalog service endpoint Create a new workflow build.
Parameters:

    workflowid: ID of a workflow.
    workflowBuildPost: The JSON representation of the workflow build to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateWorkflowRun

```go
func (s *Service) CreateWorkflowRun(workflowid string, workflowbuildid string, workflowRunPost WorkflowRunPost, resp ...*http.Response) (*WorkflowRun, error)
```
CreateWorkflowRun - catalog service endpoint Create a new workflow run for the
specified workflow build ID. Parameters:

    workflowid: ID of a workflow.
    workflowbuildid: ID of a workflow build.
    workflowRunPost: The JSON representation of the workflow run to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteActionByIdForRule

```go
func (s *Service) DeleteActionByIdForRule(ruleresourcename string, actionid string, resp ...*http.Response) error
```
DeleteActionByIdForRule - catalog service endpoint Delete an action on a rule.
Parameters:

    ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
    actionid: ID of an Action.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteActionByIdForRuleById

```go
func (s *Service) DeleteActionByIdForRuleById(ruleid string, actionid string, resp ...*http.Response) error
```
DeleteActionByIdForRuleById - catalog service endpoint Delete an action that is
part of a specific rule. Parameters:

    ruleid: ID of a Field.
    actionid: ID of an Action.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteAnnotationOfDatasetById

```go
func (s *Service) DeleteAnnotationOfDatasetById(datasetid string, annotationid string, resp ...*http.Response) error
```
DeleteAnnotationOfDatasetById - catalog service endpoint Delete a specific
annotation of a dataset. Parameters:

    datasetid: ID of a Dataset.
    annotationid: ID of a annotation.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteAnnotationOfDatasetByResourceName

```go
func (s *Service) DeleteAnnotationOfDatasetByResourceName(datasetresourcename string, annotationid string, resp ...*http.Response) error
```
DeleteAnnotationOfDatasetByResourceName - catalog service endpoint Delete a
specific annotation of a dataset. Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    annotationid: ID of a annotation.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteDashboardById

```go
func (s *Service) DeleteDashboardById(dashboardid string, resp ...*http.Response) error
```
DeleteDashboardById - catalog service endpoint Delete the dashboard with the
specified ID. Parameters:

    dashboardid: ID of a dashboard.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteDashboardByResourceName

```go
func (s *Service) DeleteDashboardByResourceName(dashboardresourcename string, resp ...*http.Response) error
```
DeleteDashboardByResourceName - catalog service endpoint Delete the dashboard
with the specified resource name. Parameters:

    dashboardresourcename: The resource name of a dashvboard. The resource name format is module.dashboardname.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteDataset

```go
func (s *Service) DeleteDataset(datasetresourcename string, resp ...*http.Response) error
```
DeleteDataset - catalog service endpoint Delete the dataset with the specified
resource name, along with its dependencies. For the default module, the resource
name format is datasetName. Otherwise, the resource name format is
module.datasetName. Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteDatasetById

```go
func (s *Service) DeleteDatasetById(datasetid string, resp ...*http.Response) error
```
DeleteDatasetById - catalog service endpoint Delete a specific dataset. Deleting
a dataset also deletes its dependent objects, such as fields. Parameters:

    datasetid: ID of a Dataset.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteFieldByIdForDataset

```go
func (s *Service) DeleteFieldByIdForDataset(datasetresourcename string, fieldid string, resp ...*http.Response) error
```
DeleteFieldByIdForDataset - catalog service endpoint Delete a field that is part
of a specific dataset. Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    fieldid: ID of a Field.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteFieldByIdForDatasetById

```go
func (s *Service) DeleteFieldByIdForDatasetById(datasetid string, fieldid string, resp ...*http.Response) error
```
DeleteFieldByIdForDatasetById - catalog service endpoint Delete a field that is
part of a specific dataset. Parameters:

    datasetid: ID of a Dataset.
    fieldid: ID of a Field.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteRelationshipById

```go
func (s *Service) DeleteRelationshipById(relationshipid string, resp ...*http.Response) error
```
DeleteRelationshipById - catalog service endpoint Delete a specific
relationship. Deleting a relationship also deleletes any objects that are
dependents of that relationship, such as relationship fields. Parameters:

    relationshipid: ID of a relationship.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteRule

```go
func (s *Service) DeleteRule(ruleresourcename string, resp ...*http.Response) error
```
DeleteRule - catalog service endpoint Delete the rule with the specified
resource name and its dependencies. Parameters:

    ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteRuleById

```go
func (s *Service) DeleteRuleById(ruleid string, resp ...*http.Response) error
```
DeleteRuleById - catalog service endpoint Delete a specific rule. Deleting a
rule also deleletes any objects that are dependents of that rule, such as rule
actions. Parameters:

    ruleid: ID of a Field.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteWorkflowBuildById

```go
func (s *Service) DeleteWorkflowBuildById(workflowid string, workflowbuildid string, resp ...*http.Response) error
```
DeleteWorkflowBuildById - catalog service endpoint Delete the workflow build
with the specified workflow build ID. Parameters:

    workflowid: ID of a workflow.
    workflowbuildid: ID of a workflow build.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteWorkflowById

```go
func (s *Service) DeleteWorkflowById(workflowid string, resp ...*http.Response) error
```
DeleteWorkflowById - catalog service endpoint Delete the workflow with the
specified workflow ID. Parameters:

    workflowid: ID of a workflow.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteWorkflowRunById

```go
func (s *Service) DeleteWorkflowRunById(workflowid string, workflowbuildid string, workflowrunid string, resp ...*http.Response) error
```
DeleteWorkflowRunById - catalog service endpoint Delete the workflow run with
the specified workflow run ID. Parameters:

    workflowid: ID of a workflow.
    workflowbuildid: ID of a workflow build.
    workflowrunid: ID of a workflow run.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetActionByIdForRule

```go
func (s *Service) GetActionByIdForRule(ruleresourcename string, actionid string, resp ...*http.Response) (*Action, error)
```
GetActionByIdForRule - catalog service endpoint Return an action that is part of
a specified rule. Parameters:

    ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
    actionid: ID of an Action.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetActionByIdForRuleById

```go
func (s *Service) GetActionByIdForRuleById(ruleid string, actionid string, resp ...*http.Response) (*Action, error)
```
GetActionByIdForRuleById - catalog service endpoint Return information about an
action that is part of a specific rule. Parameters:

    ruleid: ID of a Field.
    actionid: ID of an Action.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetDashboardById

```go
func (s *Service) GetDashboardById(dashboardid string, resp ...*http.Response) (*Dashboard, error)
```
GetDashboardById - catalog service endpoint Return information about a dashboard
with the specified ID. Parameters:

    dashboardid: ID of a dashboard.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetDashboardByResourceName

```go
func (s *Service) GetDashboardByResourceName(dashboardresourcename string, resp ...*http.Response) (*Dashboard, error)
```
GetDashboardByResourceName - catalog service endpoint Return information about a
dashboard with the specified resource name. Parameters:

    dashboardresourcename: The resource name of a dashvboard. The resource name format is module.dashboardname.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetDataset

```go
func (s *Service) GetDataset(datasetresourcename string, resp ...*http.Response) (*Dataset, error)
```
GetDataset - catalog service endpoint Return the dataset with the specified
resource name. For the default module, the resource name format is datasetName.
Otherwise, the resource name format is module.datasetName. Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetDatasetById

```go
func (s *Service) GetDatasetById(datasetid string, resp ...*http.Response) (*Dataset, error)
```
GetDatasetById - catalog service endpoint Return information about the dataset
with the specified ID. Parameters:

    datasetid: ID of a Dataset.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetFieldById

```go
func (s *Service) GetFieldById(fieldid string, resp ...*http.Response) (*Field, error)
```
GetFieldById - catalog service endpoint Get a field that corresponds to a
specific field ID. Parameters:

    fieldid: ID of a Field.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetFieldByIdForDataset

```go
func (s *Service) GetFieldByIdForDataset(datasetresourcename string, fieldid string, resp ...*http.Response) (*Field, error)
```
GetFieldByIdForDataset - catalog service endpoint Return a field that is part of
a specific dataset. Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    fieldid: ID of a Field.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetFieldByIdForDatasetById

```go
func (s *Service) GetFieldByIdForDatasetById(datasetid string, fieldid string, resp ...*http.Response) (*Field, error)
```
GetFieldByIdForDatasetById - catalog service endpoint Return a field that is
part of a specific dataset. Parameters:

    datasetid: ID of a Dataset.
    fieldid: ID of a Field.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetRelationshipById

```go
func (s *Service) GetRelationshipById(relationshipid string, resp ...*http.Response) (*Relationship, error)
```
GetRelationshipById - catalog service endpoint Get a specific relationship.
Parameters:

    relationshipid: ID of a relationship.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetRule

```go
func (s *Service) GetRule(ruleresourcename string, resp ...*http.Response) (*Rule, error)
```
GetRule - catalog service endpoint Get a rule with a specified resource name.
Parameters:

    ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetRuleById

```go
func (s *Service) GetRuleById(ruleid string, resp ...*http.Response) (*Rule, error)
```
GetRuleById - catalog service endpoint Get information about a specific rule.
Parameters:

    ruleid: ID of a Field.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowBuildById

```go
func (s *Service) GetWorkflowBuildById(workflowid string, workflowbuildid string, resp ...*http.Response) (*WorkflowBuild, error)
```
GetWorkflowBuildById - catalog service endpoint Return information about the
workflow build with the specified workflow build ID. Parameters:

    workflowid: ID of a workflow.
    workflowbuildid: ID of a workflow build.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowById

```go
func (s *Service) GetWorkflowById(workflowid string, resp ...*http.Response) (*Workflow, error)
```
GetWorkflowById - catalog service endpoint Return information about a workflow
with the specified workflow ID. Parameters:

    workflowid: ID of a workflow.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetWorkflowRunById

```go
func (s *Service) GetWorkflowRunById(workflowid string, workflowbuildid string, workflowrunid string, resp ...*http.Response) (*WorkflowRun, error)
```
GetWorkflowRunById - catalog service endpoint Return information about the
workflow run with the specified workflow build ID. Parameters:

    workflowid: ID of a workflow.
    workflowbuildid: ID of a workflow build.
    workflowrunid: ID of a workflow run.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListActionsForRule

```go
func (s *Service) ListActionsForRule(ruleresourcename string, query *ListActionsForRuleQueryParams, resp ...*http.Response) ([]Action, error)
```
ListActionsForRule - catalog service endpoint Return the list of actions that
are part of a specified rule. Parameters:

    ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListActionsForRuleById

```go
func (s *Service) ListActionsForRuleById(ruleid string, query *ListActionsForRuleByIdQueryParams, resp ...*http.Response) ([]Action, error)
```
ListActionsForRuleById - catalog service endpoint Return the set of actions that
are part of a rule. Parameters:

    ruleid: ID of a Field.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListAnnotationsForDatasetById

```go
func (s *Service) ListAnnotationsForDatasetById(datasetid string, query *ListAnnotationsForDatasetByIdQueryParams, resp ...*http.Response) ([]Annotation, error)
```
ListAnnotationsForDatasetById - catalog service endpoint Return the set of
annotations that are part of a dataset. Parameters:

    datasetid: ID of a Dataset.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListAnnotationsForDatasetByResourceName

```go
func (s *Service) ListAnnotationsForDatasetByResourceName(datasetresourcename string, query *ListAnnotationsForDatasetByResourceNameQueryParams, resp ...*http.Response) ([]Annotation, error)
```
ListAnnotationsForDatasetByResourceName - catalog service endpoint Return the
set of annotations that are part of a dataset. Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListDashboards

```go
func (s *Service) ListDashboards(resp ...*http.Response) ([]Dashboard, error)
```
ListDashboards - catalog service endpoint Return a list of Dashboards.
Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListDatasets

```go
func (s *Service) ListDatasets(query *ListDatasetsQueryParams, resp ...*http.Response) ([]Dataset, error)
```
ListDatasets - catalog service endpoint Returns a list of all datasets, unless
you specify a filter. Use a filter to return a specific list of datasets.
Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListFields

```go
func (s *Service) ListFields(query *ListFieldsQueryParams, resp ...*http.Response) ([]Field, error)
```
ListFields - catalog service endpoint Get a list of all fields in the Catalog.
Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListFieldsForDataset

```go
func (s *Service) ListFieldsForDataset(datasetresourcename string, query *ListFieldsForDatasetQueryParams, resp ...*http.Response) ([]Field, error)
```
ListFieldsForDataset - catalog service endpoint Return the list of fields that
are part of a specified dataset. Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListFieldsForDatasetById

```go
func (s *Service) ListFieldsForDatasetById(datasetid string, query *ListFieldsForDatasetByIdQueryParams, resp ...*http.Response) ([]Field, error)
```
ListFieldsForDatasetById - catalog service endpoint Return the set of fields for
the specified dataset. Parameters:

    datasetid: ID of a Dataset.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListModules

```go
func (s *Service) ListModules(query *ListModulesQueryParams, resp ...*http.Response) ([]Module, error)
```
ListModules - catalog service endpoint Return a list of all modules, unless you
specify a filter. Use a filter to return a specific list of modules. Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListRelationships

```go
func (s *Service) ListRelationships(query *ListRelationshipsQueryParams, resp ...*http.Response) ([]Relationship, error)
```
ListRelationships - catalog service endpoint Returns a list of all
relationships, unless you specify a filter. Use a filter to return a specific
list of relationships. Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListRules

```go
func (s *Service) ListRules(query *ListRulesQueryParams, resp ...*http.Response) ([]Rule, error)
```
ListRules - catalog service endpoint Return a list of rules that match a filter
query if it is given, otherwise return all rules. Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListWorkflowBuilds

```go
func (s *Service) ListWorkflowBuilds(workflowid string, query *ListWorkflowBuildsQueryParams, resp ...*http.Response) ([]WorkflowBuild, error)
```
ListWorkflowBuilds - catalog service endpoint Return a list of Machine Learning
workflow builds. Parameters:

    workflowid: ID of a workflow.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListWorkflowRuns

```go
func (s *Service) ListWorkflowRuns(workflowid string, workflowbuildid string, query *ListWorkflowRunsQueryParams, resp ...*http.Response) ([]WorkflowRun, error)
```
ListWorkflowRuns - catalog service endpoint Return a list of Machine Learning
workflow runs for specified workflow build ID. Parameters:

    workflowid: ID of a workflow.
    workflowbuildid: ID of a workflow build.
    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListWorkflows

```go
func (s *Service) ListWorkflows(query *ListWorkflowsQueryParams, resp ...*http.Response) ([]Workflow, error)
```
ListWorkflows - catalog service endpoint Return a list of Machine Learning
workflow configurations. Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateActionByIdForRule

```go
func (s *Service) UpdateActionByIdForRule(ruleresourcename string, actionid string, actionPatch ActionPatch, resp ...*http.Response) (*Action, error)
```
UpdateActionByIdForRule - catalog service endpoint Update the Action with the
specified id for the specified Rule Parameters:

    ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
    actionid: ID of an Action.
    actionPatch: The fields to update in the specified action.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateActionByIdForRuleById

```go
func (s *Service) UpdateActionByIdForRuleById(ruleid string, actionid string, actionPatch ActionPatch, resp ...*http.Response) (*Action, error)
```
UpdateActionByIdForRuleById - catalog service endpoint Update an action for a
specific rule. Parameters:

    ruleid: ID of a Field.
    actionid: ID of an Action.
    actionPatch: The properties to update in the specified action.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateDashboardById

```go
func (s *Service) UpdateDashboardById(dashboardid string, dashboardPatch DashboardPatch, resp ...*http.Response) error
```
UpdateDashboardById - catalog service endpoint Update the dashboard with the
specified ID. Parameters:

    dashboardid: ID of a dashboard.
    dashboardPatch: An updated representation of the dashboard to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateDashboardByResourceName

```go
func (s *Service) UpdateDashboardByResourceName(dashboardresourcename string, dashboardPatch DashboardPatch, resp ...*http.Response) error
```
UpdateDashboardByResourceName - catalog service endpoint Update the dashboard
with the specified resource name. Parameters:

    dashboardresourcename: The resource name of a dashvboard. The resource name format is module.dashboardname.
    dashboardPatch: An updated representation of the dashboard to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateDataset

```go
func (s *Service) UpdateDataset(datasetresourcename string, datasetPatch DatasetPatch, resp ...*http.Response) (*Dataset, error)
```
UpdateDataset - catalog service endpoint Update the dataset with the specified
resource name. For the default module, the resource name format is datasetName.
Otherwise, the resource name format is module.datasetName. Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    datasetPatch: An updated representation of the dataset to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateDatasetById

```go
func (s *Service) UpdateDatasetById(datasetid string, datasetPatch DatasetPatch, resp ...*http.Response) (*Dataset, error)
```
UpdateDatasetById - catalog service endpoint Update a specific dataset.
Parameters:

    datasetid: ID of a Dataset.
    datasetPatch: An updated representation of the dataset to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateFieldByIdForDataset

```go
func (s *Service) UpdateFieldByIdForDataset(datasetresourcename string, fieldid string, fieldPatch FieldPatch, resp ...*http.Response) (*Field, error)
```
UpdateFieldByIdForDataset - catalog service endpoint Update a field with a
specified ID for a specified dataset. Parameters:

    datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
    fieldid: ID of a Field.
    fieldPatch: The properties to update in the specified field.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateFieldByIdForDatasetById

```go
func (s *Service) UpdateFieldByIdForDatasetById(datasetid string, fieldid string, fieldPatch FieldPatch, resp ...*http.Response) (*Field, error)
```
UpdateFieldByIdForDatasetById - catalog service endpoint Update a field for a
specific dataset. Parameters:

    datasetid: ID of a Dataset.
    fieldid: ID of a Field.
    fieldPatch: The properties to update in the specified field.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateRelationshipById

```go
func (s *Service) UpdateRelationshipById(relationshipid string, relationshipPatch RelationshipPatch, resp ...*http.Response) (*Relationship, error)
```
UpdateRelationshipById - catalog service endpoint Update a specific
relationship. Parameters:

    relationshipid: ID of a relationship.
    relationshipPatch: The properties to update in the specified relationship.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateRule

```go
func (s *Service) UpdateRule(ruleresourcename string, rulePatch RulePatch, resp ...*http.Response) (*Rule, error)
```
UpdateRule - catalog service endpoint Update the Rule with the specified
resourcename Parameters:

    ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
    rulePatch: The properties to update in the specified rule.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateRuleById

```go
func (s *Service) UpdateRuleById(ruleid string, rulePatch RulePatch, resp ...*http.Response) (*Rule, error)
```
UpdateRuleById - catalog service endpoint Update a specific rule. Parameters:

    ruleid: ID of a Field.
    rulePatch: The properties to update in the specified rule.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateWorkflowBuildById

```go
func (s *Service) UpdateWorkflowBuildById(workflowid string, workflowbuildid string, workflowBuildPatch WorkflowBuildPatch, resp ...*http.Response) error
```
UpdateWorkflowBuildById - catalog service endpoint Update the workflow build
with the specified workflow build ID. Parameters:

    workflowid: ID of a workflow.
    workflowbuildid: ID of a workflow build.
    workflowBuildPatch: An updated representation of the workflow build to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateWorkflowById

```go
func (s *Service) UpdateWorkflowById(workflowid string, workflowPatch WorkflowPatch, resp ...*http.Response) error
```
UpdateWorkflowById - catalog service endpoint Update the workflow with the
specified workflow ID. Parameters:

    workflowid: ID of a workflow.
    workflowPatch: An updated representation of the workflow to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateWorkflowRunById

```go
func (s *Service) UpdateWorkflowRunById(workflowid string, workflowbuildid string, workflowrunid string, workflowRunPatch WorkflowRunPatch, resp ...*http.Response) error
```
UpdateWorkflowRunById - catalog service endpoint Update the workflow run with
the specified workflow run ID. Parameters:

    workflowid: ID of a workflow.
    workflowbuildid: ID of a workflow build.
    workflowrunid: ID of a workflow run.
    workflowRunPatch: An updated representation of the workflow run to be persisted.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### type Servicer

```go
type Servicer interface {
	/*
		CreateActionForRule - catalog service endpoint
		Create a new action for a rule associated with a specific resource name.
		Parameters:
			ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
			actionPost: The JSON representation of the action to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateActionForRule(ruleresourcename string, actionPost ActionPost, resp ...*http.Response) (*Action, error)
	/*
		CreateActionForRuleById - catalog service endpoint
		Create a new action for a specific rule.
		Parameters:
			ruleid: ID of a Field.
			actionPost: The JSON representation of the action to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateActionForRuleById(ruleid string, actionPost ActionPost, resp ...*http.Response) (*Action, error)
	/*
		CreateAnnotationForDatasetById - catalog service endpoint
		Create a new annotation for a specific dataset.
		Parameters:
			datasetid: ID of a Dataset.
			annotationPost: The JSON representation of the annotation to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateAnnotationForDatasetById(datasetid string, annotationPost AnnotationPost, resp ...*http.Response) (*Annotation, error)
	/*
		CreateAnnotationForDatasetByResourceName - catalog service endpoint
		Create a new annotation for a specific dataset.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			annotationPost: The JSON representation of the annotation to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateAnnotationForDatasetByResourceName(datasetresourcename string, annotationPost AnnotationPost, resp ...*http.Response) (*Annotation, error)
	/*
		CreateDashboard - catalog service endpoint
		Create a new dashboard.
		Parameters:
			dashboardPost: The JSON representation of the Dashboard to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateDashboard(dashboardPost DashboardPost, resp ...*http.Response) (*Dashboard, error)
	/*
		CreateDataset - catalog service endpoint
		Create a new dataset.
		Parameters:
			datasetPost: JSON representation of the DatasetInfo to be persisted
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateDataset(datasetPost DatasetPost, resp ...*http.Response) (*Dataset, error)
	/*
		CreateDatasetImport - catalog service endpoint
		Create a new dataset import.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			datasetImportedBy
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateDatasetImport(datasetresourcename string, datasetImportedBy DatasetImportedBy, resp ...*http.Response) (*ImportDataset, error)
	/*
		CreateDatasetImportById - catalog service endpoint
		Create a new dataset import.
		Parameters:
			datasetid: ID of a Dataset.
			datasetImportedBy
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateDatasetImportById(datasetid string, datasetImportedBy DatasetImportedBy, resp ...*http.Response) (*ImportDataset, error)
	/*
		CreateFieldForDataset - catalog service endpoint
		Create a new field on a specific dataset.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			fieldPost: The JSON representation of the field to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateFieldForDataset(datasetresourcename string, fieldPost FieldPost, resp ...*http.Response) (*Field, error)
	/*
		CreateFieldForDatasetById - catalog service endpoint
		Add a new field to a dataset.
		Parameters:
			datasetid: ID of a Dataset.
			fieldPost: The JSON representation of the field to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateFieldForDatasetById(datasetid string, fieldPost FieldPost, resp ...*http.Response) (*Field, error)
	/*
		CreateRelationship - catalog service endpoint
		Create a new relationship.
		Parameters:
			relationshipPost: The JSON representation of the relationship to persist.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateRelationship(relationshipPost RelationshipPost, resp ...*http.Response) (*Relationship, error)
	/*
		CreateRule - catalog service endpoint
		Create a new rule.
		Parameters:
			rulePost: The JSON representation of the rule to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateRule(rulePost RulePost, resp ...*http.Response) (*Rule, error)
	/*
		CreateWorkflow - catalog service endpoint
		Create a new workflow configuration.
		Parameters:
			workflowPost: The JSON representation of the workflow to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflow(workflowPost WorkflowPost, resp ...*http.Response) (*Workflow, error)
	/*
		CreateWorkflowBuild - catalog service endpoint
		Create a new workflow build.
		Parameters:
			workflowid: ID of a workflow.
			workflowBuildPost: The JSON representation of the workflow build to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflowBuild(workflowid string, workflowBuildPost WorkflowBuildPost, resp ...*http.Response) (*WorkflowBuild, error)
	/*
		CreateWorkflowRun - catalog service endpoint
		Create a new workflow run for the specified workflow build ID.
		Parameters:
			workflowid: ID of a workflow.
			workflowbuildid: ID of a workflow build.
			workflowRunPost: The JSON representation of the workflow run to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflowRun(workflowid string, workflowbuildid string, workflowRunPost WorkflowRunPost, resp ...*http.Response) (*WorkflowRun, error)
	/*
		DeleteActionByIdForRule - catalog service endpoint
		Delete an action on a rule.
		Parameters:
			ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
			actionid: ID of an Action.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteActionByIdForRule(ruleresourcename string, actionid string, resp ...*http.Response) error
	/*
		DeleteActionByIdForRuleById - catalog service endpoint
		Delete an action that is part of a specific rule.
		Parameters:
			ruleid: ID of a Field.
			actionid: ID of an Action.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteActionByIdForRuleById(ruleid string, actionid string, resp ...*http.Response) error
	/*
		DeleteAnnotationOfDatasetById - catalog service endpoint
		Delete a specific annotation of a dataset.
		Parameters:
			datasetid: ID of a Dataset.
			annotationid: ID of a annotation.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteAnnotationOfDatasetById(datasetid string, annotationid string, resp ...*http.Response) error
	/*
		DeleteAnnotationOfDatasetByResourceName - catalog service endpoint
		Delete a specific annotation of a dataset.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			annotationid: ID of a annotation.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteAnnotationOfDatasetByResourceName(datasetresourcename string, annotationid string, resp ...*http.Response) error
	/*
		DeleteDashboardById - catalog service endpoint
		Delete the dashboard with the specified ID.
		Parameters:
			dashboardid: ID of a dashboard.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteDashboardById(dashboardid string, resp ...*http.Response) error
	/*
		DeleteDashboardByResourceName - catalog service endpoint
		Delete the dashboard with the specified resource name.
		Parameters:
			dashboardresourcename: The resource name of a dashvboard. The resource name format is module.dashboardname.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteDashboardByResourceName(dashboardresourcename string, resp ...*http.Response) error
	/*
		DeleteDataset - catalog service endpoint
		Delete the dataset with the specified resource name, along with its dependencies. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteDataset(datasetresourcename string, resp ...*http.Response) error
	/*
		DeleteDatasetById - catalog service endpoint
		Delete a specific dataset. Deleting a dataset also deletes its dependent objects, such as fields.
		Parameters:
			datasetid: ID of a Dataset.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteDatasetById(datasetid string, resp ...*http.Response) error
	/*
		DeleteFieldByIdForDataset - catalog service endpoint
		Delete a field that is part of a specific dataset.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			fieldid: ID of a Field.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteFieldByIdForDataset(datasetresourcename string, fieldid string, resp ...*http.Response) error
	/*
		DeleteFieldByIdForDatasetById - catalog service endpoint
		Delete a field that is part of a specific dataset.
		Parameters:
			datasetid: ID of a Dataset.
			fieldid: ID of a Field.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteFieldByIdForDatasetById(datasetid string, fieldid string, resp ...*http.Response) error
	/*
		DeleteRelationshipById - catalog service endpoint
		Delete a specific relationship. Deleting a relationship also deleletes any objects that are dependents of that relationship, such as relationship fields.
		Parameters:
			relationshipid: ID of a relationship.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteRelationshipById(relationshipid string, resp ...*http.Response) error
	/*
		DeleteRule - catalog service endpoint
		Delete the rule with the specified resource name and its dependencies.
		Parameters:
			ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteRule(ruleresourcename string, resp ...*http.Response) error
	/*
		DeleteRuleById - catalog service endpoint
		Delete a specific rule. Deleting a rule also deleletes any objects that are dependents of that rule, such as rule actions.
		Parameters:
			ruleid: ID of a Field.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteRuleById(ruleid string, resp ...*http.Response) error
	/*
		DeleteWorkflowBuildById - catalog service endpoint
		Delete the workflow build with the specified workflow build ID.
		Parameters:
			workflowid: ID of a workflow.
			workflowbuildid: ID of a workflow build.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflowBuildById(workflowid string, workflowbuildid string, resp ...*http.Response) error
	/*
		DeleteWorkflowById - catalog service endpoint
		Delete the workflow with the specified workflow ID.
		Parameters:
			workflowid: ID of a workflow.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflowById(workflowid string, resp ...*http.Response) error
	/*
		DeleteWorkflowRunById - catalog service endpoint
		Delete the workflow run with the specified workflow run ID.
		Parameters:
			workflowid: ID of a workflow.
			workflowbuildid: ID of a workflow build.
			workflowrunid: ID of a workflow run.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflowRunById(workflowid string, workflowbuildid string, workflowrunid string, resp ...*http.Response) error
	/*
		GetActionByIdForRule - catalog service endpoint
		Return an action that is part of a specified rule.
		Parameters:
			ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
			actionid: ID of an Action.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetActionByIdForRule(ruleresourcename string, actionid string, resp ...*http.Response) (*Action, error)
	/*
		GetActionByIdForRuleById - catalog service endpoint
		Return information about an action that is part of a specific rule.
		Parameters:
			ruleid: ID of a Field.
			actionid: ID of an Action.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetActionByIdForRuleById(ruleid string, actionid string, resp ...*http.Response) (*Action, error)
	/*
		GetDashboardById - catalog service endpoint
		Return information about a dashboard with the specified ID.
		Parameters:
			dashboardid: ID of a dashboard.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetDashboardById(dashboardid string, resp ...*http.Response) (*Dashboard, error)
	/*
		GetDashboardByResourceName - catalog service endpoint
		Return information about a dashboard with the specified resource name.
		Parameters:
			dashboardresourcename: The resource name of a dashvboard. The resource name format is module.dashboardname.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetDashboardByResourceName(dashboardresourcename string, resp ...*http.Response) (*Dashboard, error)
	/*
		GetDataset - catalog service endpoint
		Return the dataset with the specified resource name. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetDataset(datasetresourcename string, resp ...*http.Response) (*Dataset, error)
	/*
		GetDatasetById - catalog service endpoint
		Return information about the dataset with the specified ID.
		Parameters:
			datasetid: ID of a Dataset.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetDatasetById(datasetid string, resp ...*http.Response) (*Dataset, error)
	/*
		GetFieldById - catalog service endpoint
		Get a field that corresponds to a specific field ID.
		Parameters:
			fieldid: ID of a Field.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetFieldById(fieldid string, resp ...*http.Response) (*Field, error)
	/*
		GetFieldByIdForDataset - catalog service endpoint
		Return a field that is part of a specific dataset.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			fieldid: ID of a Field.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetFieldByIdForDataset(datasetresourcename string, fieldid string, resp ...*http.Response) (*Field, error)
	/*
		GetFieldByIdForDatasetById - catalog service endpoint
		Return a field that is part of a specific dataset.
		Parameters:
			datasetid: ID of a Dataset.
			fieldid: ID of a Field.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetFieldByIdForDatasetById(datasetid string, fieldid string, resp ...*http.Response) (*Field, error)
	/*
		GetRelationshipById - catalog service endpoint
		Get a specific relationship.
		Parameters:
			relationshipid: ID of a relationship.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetRelationshipById(relationshipid string, resp ...*http.Response) (*Relationship, error)
	/*
		GetRule - catalog service endpoint
		Get a rule with a specified resource name.
		Parameters:
			ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetRule(ruleresourcename string, resp ...*http.Response) (*Rule, error)
	/*
		GetRuleById - catalog service endpoint
		Get information about a specific rule.
		Parameters:
			ruleid: ID of a Field.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetRuleById(ruleid string, resp ...*http.Response) (*Rule, error)
	/*
		GetWorkflowBuildById - catalog service endpoint
		Return information about the workflow build with the specified workflow build ID.
		Parameters:
			workflowid: ID of a workflow.
			workflowbuildid: ID of a workflow build.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowBuildById(workflowid string, workflowbuildid string, resp ...*http.Response) (*WorkflowBuild, error)
	/*
		GetWorkflowById - catalog service endpoint
		Return information about a workflow with the specified workflow ID.
		Parameters:
			workflowid: ID of a workflow.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowById(workflowid string, resp ...*http.Response) (*Workflow, error)
	/*
		GetWorkflowRunById - catalog service endpoint
		Return information about the workflow run with the specified workflow build ID.
		Parameters:
			workflowid: ID of a workflow.
			workflowbuildid: ID of a workflow build.
			workflowrunid: ID of a workflow run.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowRunById(workflowid string, workflowbuildid string, workflowrunid string, resp ...*http.Response) (*WorkflowRun, error)
	/*
		ListActionsForRule - catalog service endpoint
		Return the list of actions that are part of a specified rule.
		Parameters:
			ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListActionsForRule(ruleresourcename string, query *ListActionsForRuleQueryParams, resp ...*http.Response) ([]Action, error)
	/*
		ListActionsForRuleById - catalog service endpoint
		Return the set of actions that are part of a rule.
		Parameters:
			ruleid: ID of a Field.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListActionsForRuleById(ruleid string, query *ListActionsForRuleByIdQueryParams, resp ...*http.Response) ([]Action, error)
	/*
		ListAnnotationsForDatasetById - catalog service endpoint
		Return the set of annotations that are part of a dataset.
		Parameters:
			datasetid: ID of a Dataset.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListAnnotationsForDatasetById(datasetid string, query *ListAnnotationsForDatasetByIdQueryParams, resp ...*http.Response) ([]Annotation, error)
	/*
		ListAnnotationsForDatasetByResourceName - catalog service endpoint
		Return the set of annotations that are part of a dataset.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListAnnotationsForDatasetByResourceName(datasetresourcename string, query *ListAnnotationsForDatasetByResourceNameQueryParams, resp ...*http.Response) ([]Annotation, error)
	/*
		ListDashboards - catalog service endpoint
		Return a list of Dashboards.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListDashboards(resp ...*http.Response) ([]Dashboard, error)
	/*
		ListDatasets - catalog service endpoint
		Returns a list of all datasets, unless you specify a filter. Use a filter to return a specific list of datasets.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListDatasets(query *ListDatasetsQueryParams, resp ...*http.Response) ([]Dataset, error)
	/*
		ListFields - catalog service endpoint
		Get a list of all fields in the Catalog.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListFields(query *ListFieldsQueryParams, resp ...*http.Response) ([]Field, error)
	/*
		ListFieldsForDataset - catalog service endpoint
		Return the list of fields that are part of a specified dataset.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListFieldsForDataset(datasetresourcename string, query *ListFieldsForDatasetQueryParams, resp ...*http.Response) ([]Field, error)
	/*
		ListFieldsForDatasetById - catalog service endpoint
		Return the set of fields for the specified dataset.
		Parameters:
			datasetid: ID of a Dataset.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListFieldsForDatasetById(datasetid string, query *ListFieldsForDatasetByIdQueryParams, resp ...*http.Response) ([]Field, error)
	/*
		ListModules - catalog service endpoint
		Return a list of all modules, unless you specify a filter. Use a filter to return a specific list of modules.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListModules(query *ListModulesQueryParams, resp ...*http.Response) ([]Module, error)
	/*
		ListRelationships - catalog service endpoint
		Returns a list of all relationships, unless you specify a filter. Use a filter to return a specific list of relationships.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListRelationships(query *ListRelationshipsQueryParams, resp ...*http.Response) ([]Relationship, error)
	/*
		ListRules - catalog service endpoint
		Return a list of rules that match a filter query if it is given, otherwise return all rules.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListRules(query *ListRulesQueryParams, resp ...*http.Response) ([]Rule, error)
	/*
		ListWorkflowBuilds - catalog service endpoint
		Return a list of Machine Learning workflow builds.
		Parameters:
			workflowid: ID of a workflow.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListWorkflowBuilds(workflowid string, query *ListWorkflowBuildsQueryParams, resp ...*http.Response) ([]WorkflowBuild, error)
	/*
		ListWorkflowRuns - catalog service endpoint
		Return a list of Machine Learning workflow runs for specified workflow build ID.
		Parameters:
			workflowid: ID of a workflow.
			workflowbuildid: ID of a workflow build.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListWorkflowRuns(workflowid string, workflowbuildid string, query *ListWorkflowRunsQueryParams, resp ...*http.Response) ([]WorkflowRun, error)
	/*
		ListWorkflows - catalog service endpoint
		Return a list of Machine Learning workflow configurations.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListWorkflows(query *ListWorkflowsQueryParams, resp ...*http.Response) ([]Workflow, error)
	/*
		UpdateActionByIdForRule - catalog service endpoint
		Update the Action with the specified id for the specified Rule
		Parameters:
			ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
			actionid: ID of an Action.
			actionPatch: The fields to update in the specified action.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateActionByIdForRule(ruleresourcename string, actionid string, actionPatch ActionPatch, resp ...*http.Response) (*Action, error)
	/*
		UpdateActionByIdForRuleById - catalog service endpoint
		Update an action for a specific rule.
		Parameters:
			ruleid: ID of a Field.
			actionid: ID of an Action.
			actionPatch: The properties to update in the specified action.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateActionByIdForRuleById(ruleid string, actionid string, actionPatch ActionPatch, resp ...*http.Response) (*Action, error)
	/*
		UpdateDashboardById - catalog service endpoint
		Update the dashboard with the specified ID.
		Parameters:
			dashboardid: ID of a dashboard.
			dashboardPatch: An updated representation of the dashboard to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateDashboardById(dashboardid string, dashboardPatch DashboardPatch, resp ...*http.Response) error
	/*
		UpdateDashboardByResourceName - catalog service endpoint
		Update the dashboard with the specified resource name.
		Parameters:
			dashboardresourcename: The resource name of a dashvboard. The resource name format is module.dashboardname.
			dashboardPatch: An updated representation of the dashboard to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateDashboardByResourceName(dashboardresourcename string, dashboardPatch DashboardPatch, resp ...*http.Response) error
	/*
		UpdateDataset - catalog service endpoint
		Update the dataset with the specified resource name. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			datasetPatch: An updated representation of the dataset to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateDataset(datasetresourcename string, datasetPatch DatasetPatch, resp ...*http.Response) (*Dataset, error)
	/*
		UpdateDatasetById - catalog service endpoint
		Update a specific dataset.
		Parameters:
			datasetid: ID of a Dataset.
			datasetPatch: An updated representation of the dataset to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateDatasetById(datasetid string, datasetPatch DatasetPatch, resp ...*http.Response) (*Dataset, error)
	/*
		UpdateFieldByIdForDataset - catalog service endpoint
		Update a field with a specified ID for a specified dataset.
		Parameters:
			datasetresourcename: The resource name of a dataset. For the default module, the resource name format is datasetName. Otherwise, the resource name format is module.datasetName.
			fieldid: ID of a Field.
			fieldPatch: The properties to update in the specified field.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateFieldByIdForDataset(datasetresourcename string, fieldid string, fieldPatch FieldPatch, resp ...*http.Response) (*Field, error)
	/*
		UpdateFieldByIdForDatasetById - catalog service endpoint
		Update a field for a specific dataset.
		Parameters:
			datasetid: ID of a Dataset.
			fieldid: ID of a Field.
			fieldPatch: The properties to update in the specified field.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateFieldByIdForDatasetById(datasetid string, fieldid string, fieldPatch FieldPatch, resp ...*http.Response) (*Field, error)
	/*
		UpdateRelationshipById - catalog service endpoint
		Update a specific relationship.
		Parameters:
			relationshipid: ID of a relationship.
			relationshipPatch: The properties to update in the specified relationship.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateRelationshipById(relationshipid string, relationshipPatch RelationshipPatch, resp ...*http.Response) (*Relationship, error)
	/*
		UpdateRule - catalog service endpoint
		Update the Rule with the specified resourcename
		Parameters:
			ruleresourcename: The resource name of a rule. For the default module, the resource name format is ruleName. Otherwise, the resource name format is module.ruleName.
			rulePatch: The properties to update in the specified rule.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateRule(ruleresourcename string, rulePatch RulePatch, resp ...*http.Response) (*Rule, error)
	/*
		UpdateRuleById - catalog service endpoint
		Update a specific rule.
		Parameters:
			ruleid: ID of a Field.
			rulePatch: The properties to update in the specified rule.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateRuleById(ruleid string, rulePatch RulePatch, resp ...*http.Response) (*Rule, error)
	/*
		UpdateWorkflowBuildById - catalog service endpoint
		Update the workflow build with the specified workflow build ID.
		Parameters:
			workflowid: ID of a workflow.
			workflowbuildid: ID of a workflow build.
			workflowBuildPatch: An updated representation of the workflow build to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateWorkflowBuildById(workflowid string, workflowbuildid string, workflowBuildPatch WorkflowBuildPatch, resp ...*http.Response) error
	/*
		UpdateWorkflowById - catalog service endpoint
		Update the workflow with the specified workflow ID.
		Parameters:
			workflowid: ID of a workflow.
			workflowPatch: An updated representation of the workflow to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateWorkflowById(workflowid string, workflowPatch WorkflowPatch, resp ...*http.Response) error
	/*
		UpdateWorkflowRunById - catalog service endpoint
		Update the workflow run with the specified workflow run ID.
		Parameters:
			workflowid: ID of a workflow.
			workflowbuildid: ID of a workflow build.
			workflowrunid: ID of a workflow run.
			workflowRunPatch: An updated representation of the workflow run to be persisted.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateWorkflowRunById(workflowid string, workflowbuildid string, workflowrunid string, workflowRunPatch WorkflowRunPatch, resp ...*http.Response) error
}
```

Servicer represents the interface for implementing all endpoints for this
service

#### type Task

```go
type Task struct {
	// The task algorithm name.
	Algorithm string `json:"algorithm"`
	// The children tasks of the task.
	Children []string `json:"children"`
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// The evaluation criteria of the task.
	Evaluation []string `json:"evaluation"`
	// The features of the task.
	Features []string `json:"features"`
	// A unique task ID.
	Id string `json:"id"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The task name.
	Name string `json:"name"`
	// The output transformer of the task.
	Outputtransformer string `json:"outputtransformer"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The task parameters.
	Parameters string `json:"parameters"`
	// The parent tasks of the task.
	Parents []string `json:"parents"`
	// The target feature of the task.
	Targetfeature string `json:"targetfeature"`
	// The task type.
	Tasktype string `json:"tasktype"`
	// The timeout secs of the task.
	Timeoutsecs int32 `json:"timeoutsecs"`
	// A unique workflow ID that is associatd with the task.
	Workflowid string `json:"workflowid"`
	// The version of the workflow that is associated with the task.
	Workflowversion int32 `json:"workflowversion"`
}
```

A complete task as rendered in POST, PATCH, and GET responses.

#### type TaskPost

```go
type TaskPost struct {
	// The task algorithm name.
	Algorithm string `json:"algorithm"`
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// The features of the task.
	Features []string `json:"features"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The task parameters.
	Parameters string `json:"parameters"`
	// The task type.
	Tasktype string `json:"tasktype"`
	// The timeout secs of the task.
	Timeoutsecs int32 `json:"timeoutsecs"`
	// The children tasks of the task.
	Children []string `json:"children,omitempty"`
	// The evaluation criteria of the task.
	Evaluation []string `json:"evaluation,omitempty"`
	// A unique task ID.
	Id *string `json:"id,omitempty"`
	// The task name.
	Name *string `json:"name,omitempty"`
	// The output transformer of the task.
	Outputtransformer *string `json:"outputtransformer,omitempty"`
	// The parent tasks of the task.
	Parents []string `json:"parents,omitempty"`
	// The target feature of the task.
	Targetfeature *string `json:"targetfeature,omitempty"`
	// A unique workflow ID that is associated with the task.
	Workflowid *string `json:"workflowid,omitempty"`
	// The version of the workflow that is associated with the task.
	Workflowversion *int32 `json:"workflowversion,omitempty"`
}
```

A complete task as rendered in POST, PATCH, and GET responses.

#### type UserMetadataProperties

```go
type UserMetadataProperties struct {
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the object's owner.
	Owner string `json:"owner"`
}
```

Owner, createdby, and modifiedby user name properties for inclusion in other
objects.

#### type ViewDataset

```go
type ViewDataset struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// A unique dataset ID.
	Id   string          `json:"id"`
	Kind ViewDatasetKind `json:"kind"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The name of the module that contains the dataset.
	Module string `json:"module"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The dataset name qualified by the module name.
	Resourcename string `json:"resourcename"`
	// A valid SPL-defined search.
	Search string `json:"search"`
	// Detailed description of the dataset.
	Description *string `json:"description,omitempty"`
	// Summary of the dataset's purpose.
	Summary *string `json:"summary,omitempty"`
	// The title of the dataset.  Does not have to be unique.
	Title *string `json:"title,omitempty"`
	// The catalog version.
	Version *int32 `json:"version,omitempty"`
}
```

A complete view dataset as rendered in POST, PATCH, and GET responses.

#### type ViewDatasetKind

```go
type ViewDatasetKind string
```

ViewDatasetKind : The dataset kind.

```go
const (
	ViewDatasetKindView ViewDatasetKind = "view"
)
```
List of ViewDatasetKind

#### type ViewDatasetPatch

```go
type ViewDatasetPatch struct {
	Kind *ViewDatasetKind `json:"kind,omitempty"`
	// The name of module to reassign dataset into.
	Module *string `json:"module,omitempty"`
	// The dataset name. Dataset names must be unique within each module.
	Name *string `json:"name,omitempty"`
	// The name of the dataset owner. This value is obtained from the bearer token.
	Owner *string `json:"owner,omitempty"`
	// A valid SPL-defined search.
	Search *string `json:"search,omitempty"`
}
```

Property values to be set in an existing view dataset using a PATCH request.

#### type ViewDatasetPost

```go
type ViewDatasetPost struct {
	Kind ViewDatasetKind `json:"kind"`
	// The dataset name. Dataset names must be unique within each module.
	Name string `json:"name"`
	// A valid SPL-defined search.
	Search string `json:"search"`
	// The fields to be associated with this dataset.
	Fields []FieldPost `json:"fields,omitempty"`
	// A unique dataset ID. Random ID used if not provided.
	Id *string `json:"id,omitempty"`
	// The name of the module to create the new dataset in.
	Module *string `json:"module,omitempty"`
}
```

Initial property values for creating a new view dataset using a POST request.

#### type ViewDatasetProperties

```go
type ViewDatasetProperties struct {
	Kind *ViewDatasetKind `json:"kind,omitempty"`
	// A valid SPL-defined search.
	Search *string `json:"search,omitempty"`
}
```

Properties of job datasets which may be read, set, and changed through the API.
Implementation detail of DatasetPOST, DatasetPATCH, and Dataset, do not use
directly.

#### type Workflow

```go
type Workflow struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// A unique workflow ID.
	Id string `json:"id"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The workflow name. Workflow names must be unique within each tenant.
	Name string `json:"name"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The workflow description.
	Description *string `json:"description,omitempty"`
	// A unique experiment ID that is associated with the workflow.
	Experimentid *string `json:"experimentid,omitempty"`
	Tasks        []Task  `json:"tasks,omitempty"`
	// The version of the workflow.
	Version *int32 `json:"version,omitempty"`
}
```

A complete workflow as rendered in POST, PATCH, and GET responses.

#### type WorkflowBuild

```go
type WorkflowBuild struct {
	// The date and time object was created.
	Created string `json:"created"`
	// The name of the user who created the object. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// A unique workflow build ID.
	Id string `json:"id"`
	// The input data of the workflow build.
	Inputdata []string `json:"inputdata"`
	// The date and time object was modified.
	Modified string `json:"modified"`
	// The name of the user who most recently modified the object.
	Modifiedby string `json:"modifiedby"`
	// The output data of the workflow build.
	Outputdata []string `json:"outputdata"`
	// The name of the object's owner.
	Owner string `json:"owner"`
	// The random state of the workflow build.
	Randomstate int32 `json:"randomstate"`
	// The timeout in seconds of the workflow.
	Timeoutsecs int32 `json:"timeoutsecs"`
	// The train test split of the workflow build.
	Traintestsplit float32 `json:"traintestsplit"`
	// A unique workflow ID that is associated with the workflow build.
	Workflowid string `json:"workflowid"`
	// The version of the workflow that is associated with the workflow build.
	Workflowversion int32 `json:"workflowversion"`
	// The description of the workflow build.
	Description *string `json:"description,omitempty"`
	// The date and time the workflow build ended.
	Ended *string `json:"ended,omitempty"`
	// The evaluation results of the workflow build.
	Evaluationresults []string `json:"evaluationresults,omitempty"`
	// The failure message of the workflow build.
	Failuremessage *string `json:"failuremessage,omitempty"`
	// The number of kfold when the validation type if CrossValidation.
	Kfold *float32 `json:"kfold,omitempty"`
	// The workflow build name.
	Name *string `json:"name,omitempty"`
	// The date and time the workflow build started.
	Started *string `json:"started,omitempty"`
	// The status of the workflow build.
	Status *string `json:"status,omitempty"`
	// Whether data is stratified.
	Stratified *bool `json:"stratified,omitempty"`
	// The type of validation.
	Validationkind *WorkflowBuildValidationkind `json:"validationkind,omitempty"`
	// The version of the workflow.
	Version *int32 `json:"version,omitempty"`
}
```

A complete workflow build as rendered in POST, PATCH, and GET responses.

#### type WorkflowBuildPatch

```go
type WorkflowBuildPatch struct {
	// The workflow build description.
	Description *string `json:"description,omitempty"`
	// The workflow build name.
	Name *string `json:"name,omitempty"`
	// The status of the workflow build.
	Status *string `json:"status,omitempty"`
}
```

Values for updating a workflow build using a PATCH request.

#### type WorkflowBuildPost

```go
type WorkflowBuildPost struct {
	// The input data of the workflow build.
	Inputdata []string `json:"inputdata"`
	// The timeout in seconds of the workflow.
	Timeoutsecs int32 `json:"timeoutsecs"`
	// The description of the workflow build.
	Description *string `json:"description,omitempty"`
	// The date and time the workflow build ended.
	Ended *string `json:"ended,omitempty"`
	// The evaluation results of the workflow build.
	Evaluationresults []string `json:"evaluationresults,omitempty"`
	// The failure message of the workflow build.
	Failuremessage *string `json:"failuremessage,omitempty"`
	// A unique workflow build ID.
	Id *string `json:"id,omitempty"`
	// The workflow build name.
	Name *string `json:"name,omitempty"`
	// The output data of the workflow build.
	Outputdata []string `json:"outputdata,omitempty"`
	// The random state of the workflow build.
	Randomstate *int32 `json:"randomstate,omitempty"`
	// The date and time the workflow build started.
	Started *string `json:"started,omitempty"`
	// The status of the workflow build.
	Status *string `json:"status,omitempty"`
	// The train test split of the workflow build.
	Traintestsplit *float32 `json:"traintestsplit,omitempty"`
	// The version of the workflow.
	Version *int32 `json:"version,omitempty"`
	// A unique workflow ID that is associated with the workflow build.
	Workflowid *string `json:"workflowid,omitempty"`
	// The version of the workflow that is associated with the workflow build.
	Workflowversion *int32 `json:"workflowversion,omitempty"`
}
```

Initial values for creating a new workflow using a POST request.

#### type WorkflowBuildValidationkind

```go
type WorkflowBuildValidationkind string
```

WorkflowBuildValidationkind : The type of validation.

```go
const (
	WorkflowBuildValidationkindTrainTest       WorkflowBuildValidationkind = "TrainTest"
	WorkflowBuildValidationkindCrossValidation WorkflowBuildValidationkind = "CrossValidation"
)
```
List of WorkflowBuildValidationkind

#### type WorkflowPatch

```go
type WorkflowPatch struct {
	// The workflow description.
	Description *string `json:"description,omitempty"`
	// The workflow name.
	Name *string `json:"name,omitempty"`
}
```

Values for updating a workflow using a PATCH request.

#### type WorkflowPost

```go
type WorkflowPost struct {
	Tasks []TaskPost `json:"tasks"`
	// The workflow description.
	Description *string `json:"description,omitempty"`
	// A unique experiment ID that is associate with the workflow.
	Experimentid *string `json:"experimentid,omitempty"`
	// A unique workflow ID. Random ID used if not provided.
	Id *string `json:"id,omitempty"`
	// The dataset name. Dataset names must be unique within each module.
	Name *string `json:"name,omitempty"`
	// The version of the workflow.
	Version *int32 `json:"version,omitempty"`
}
```

Initial values for creating a new workflow using a POST request.

#### type WorkflowRun

```go
type WorkflowRun struct {
	// The date and time the workflow run was created.
	Created string `json:"created"`
	// The name of the user who created the workflow run. This value is obtained from the bearer token and may not be changed.
	Createdby string `json:"createdby"`
	// A unique workflow Run ID.
	Id string `json:"id"`
	// The input data of the workflow run.
	Inputdata []string `json:"inputdata"`
	// The output data of the workflow run.
	Outputdata []string `json:"outputdata"`
	// The name of the workflow run's owner.
	Owner string `json:"owner"`
	// The timeout in seconds of the workflow.
	Timeoutsecs int32 `json:"timeoutsecs"`
	// A unique workflow build ID that is associated with the workflow run.
	Workflowbuildid string `json:"workflowbuildid"`
	// The version of the workflow build that is associated with the workflow run.
	Workflowbuildversion int32 `json:"workflowbuildversion"`
	// The description of the workflow run.
	Description *string `json:"description,omitempty"`
	// The date and time the workflow build ended.
	Ended *string `json:"ended,omitempty"`
	// The failure message of the workflow run.
	Failuremessage *string `json:"failuremessage,omitempty"`
	// The workflow run name.
	Name *string `json:"name,omitempty"`
	// The date and time the workflow build started.
	Started *string `json:"started,omitempty"`
	// The status of the workflow run.
	Status *string `json:"status,omitempty"`
}
```

A complete workflow run as rendered in POST, PATCH, and GET responses.

#### type WorkflowRunPatch

```go
type WorkflowRunPatch struct {
	// The workflow run description.
	Description *string `json:"description,omitempty"`
	// The workflow run name.
	Name *string `json:"name,omitempty"`
	// The status of the workflow run.
	Status *string `json:"status,omitempty"`
}
```

Values for updating a workflow run using a PATCH request.

#### type WorkflowRunPost

```go
type WorkflowRunPost struct {
	// The input data of the workflow run for specified workflow build ID.
	Inputdata []string `json:"inputdata"`
	// The output data of the workflow run for specified workflow build ID.
	Outputdata []string `json:"outputdata"`
	// The timeout in seconds of the workflow run for specified workflow build ID.
	Timeoutsecs int32 `json:"timeoutsecs"`
	// The description of the workflow run.
	Description *string `json:"description,omitempty"`
	// The date and time the workflow run ended for specified workflow build ID.
	Ended *string `json:"ended,omitempty"`
	// The failure message of the workflow run for specified workflow build ID.
	Failuremessage *string `json:"failuremessage,omitempty"`
	// A unique workflow Run ID.
	Id *string `json:"id,omitempty"`
	// The workflow run name.
	Name *string `json:"name,omitempty"`
	// The date and time the workflow run started for specified workflow build ID.
	Started *string `json:"started,omitempty"`
	// The status of the workflow run for specified workflow build ID.
	Status *string `json:"status,omitempty"`
	// A unique workflow build ID that is associated with the workflow run.
	Workflowbuildid *string `json:"workflowbuildid,omitempty"`
	// The version of the workflow build that is assocaited with the workflow run.
	Workflowbuildversion *int32 `json:"workflowbuildversion,omitempty"`
}
```

Initial values for creating a new workflow run using a POST request.
