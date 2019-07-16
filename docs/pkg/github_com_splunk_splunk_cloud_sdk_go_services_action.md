# action
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/action"

This files contains models that can't be auto-generated from codegen

This files contains models that can't be auto-generated from codegen

## Usage

#### type Action

```go
type Action struct {
}
```


#### func  MakeActionFromEmailAction

```go
func MakeActionFromEmailAction(f EmailAction) Action
```
MakeActionFromEmailAction creates a new Action from an instance of EmailAction

#### func  MakeActionFromRawInterface

```go
func MakeActionFromRawInterface(f interface{}) Action
```
MakeActionFromRawInterface creates a new Action from a raw interface{}

#### func  MakeActionFromWebhookAction

```go
func MakeActionFromWebhookAction(f WebhookAction) Action
```
MakeActionFromWebhookAction creates a new Action from an instance of
WebhookAction

#### func (Action) EmailAction

```go
func (m Action) EmailAction() *EmailAction
```
EmailAction returns EmailAction if IsEmailAction() is true, nil otherwise

#### func (Action) IsEmailAction

```go
func (m Action) IsEmailAction() bool
```
IsEmailAction checks if the Action is a EmailAction

#### func (Action) IsRawInterface

```go
func (m Action) IsRawInterface() bool
```
IsRawInterface checks if the Action is an interface{} (unknown type)

#### func (Action) IsWebhookAction

```go
func (m Action) IsWebhookAction() bool
```
IsWebhookAction checks if the Action is a WebhookAction

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

#### func (*Action) UnmarshalJSON

```go
func (m *Action) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals Action using the "kind" property

#### func (Action) WebhookAction

```go
func (m Action) WebhookAction() *WebhookAction
```
WebhookAction returns WebhookAction if IsWebhookAction() is true, nil otherwise

#### type ActionImmutable

```go
type ActionImmutable struct {
	Kind ActionKind `json:"kind"`
	// The name of the action, as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
	Name string `json:"name"`
	// The date and time this action template was created (ISO-8601 date/time with zone).
	CreatedAt *string `json:"createdAt,omitempty"`
	// The principal that created this action template.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The date and time this action template was updated (ISO-8601 date/time with zone).
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The principal that updated this action template.
	UpdatedBy *string `json:"updatedBy,omitempty"`
}
```


#### type ActionKind

```go
type ActionKind string
```


```go
const (
	ActionKindWebhook ActionKind = "webhook"
	ActionKindEmail   ActionKind = "email"
)
```
List of ActionKind

#### type ActionMutable

```go
type ActionMutable struct {
}
```

ActionMutable is EmailActionMutable, WebhookActionMutable, (or interface{} if no
matches are found)

#### func  MakeActionMutableFromEmailActionMutable

```go
func MakeActionMutableFromEmailActionMutable(f EmailActionMutable) ActionMutable
```
MakeActionMutableFromEmailActionMutable creates a new ActionMutable from an
instance of EmailActionMutable

#### func  MakeActionMutableFromRawInterface

```go
func MakeActionMutableFromRawInterface(f interface{}) ActionMutable
```
MakeActionMutableFromRawInterface creates a new ActionMutable from a raw
interface{}

#### func  MakeActionMutableFromWebhookActionMutable

```go
func MakeActionMutableFromWebhookActionMutable(f WebhookActionMutable) ActionMutable
```
MakeActionMutableFromWebhookActionMutable creates a new ActionMutable from an
instance of WebhookActionMutable

#### func (ActionMutable) EmailActionMutable

```go
func (m ActionMutable) EmailActionMutable() *EmailActionMutable
```
EmailActionMutable returns EmailActionMutable if IsEmailActionMutable() is true,
nil otherwise

#### func (ActionMutable) IsEmailActionMutable

```go
func (m ActionMutable) IsEmailActionMutable() bool
```
IsEmailActionMutable checks if the ActionMutable is a EmailActionMutable

#### func (ActionMutable) IsRawInterface

```go
func (m ActionMutable) IsRawInterface() bool
```
IsRawInterface checks if the ActionMutable is an interface{} (unknown type)

#### func (ActionMutable) IsWebhookActionMutable

```go
func (m ActionMutable) IsWebhookActionMutable() bool
```
IsWebhookActionMutable checks if the ActionMutable is a WebhookActionMutable

#### func (ActionMutable) MarshalJSON

```go
func (m ActionMutable) MarshalJSON() ([]byte, error)
```
MarshalJSON marshals ActionMutable using ActionMutable.ActionMutable

#### func (ActionMutable) RawInterface

```go
func (m ActionMutable) RawInterface() interface{}
```
RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil
otherwise

#### func (*ActionMutable) UnmarshalJSON

```go
func (m *ActionMutable) UnmarshalJSON(b []byte) (err error)
```
UnmarshalJSON unmarshals ActionMutable into EmailActionMutable,
WebhookActionMutable, or interface{} if no matches are found

#### func (ActionMutable) WebhookActionMutable

```go
func (m ActionMutable) WebhookActionMutable() *WebhookActionMutable
```
WebhookActionMutable returns WebhookActionMutable if IsWebhookActionMutable() is
true, nil otherwise

#### type ActionResult

```go
type ActionResult struct {
	ActionName string      `json:"actionName"`
	State      StatusState `json:"state"`
	StatusId   string      `json:"statusId"`
	Message    *string     `json:"message,omitempty"`
}
```


#### type ActionResultEmailDetail

```go
type ActionResultEmailDetail struct {
	EmailAddress *string                       `json:"emailAddress,omitempty"`
	State        *ActionResultEmailDetailState `json:"state,omitempty"`
}
```


#### type ActionResultEmailDetailState

```go
type ActionResultEmailDetailState string
```


```go
const (
	ActionResultEmailDetailStatePending             ActionResultEmailDetailState = "PENDING"
	ActionResultEmailDetailStateNotFound            ActionResultEmailDetailState = "NOT_FOUND"
	ActionResultEmailDetailStateSucceeded           ActionResultEmailDetailState = "SUCCEEDED"
	ActionResultEmailDetailStateBounced             ActionResultEmailDetailState = "BOUNCED"
	ActionResultEmailDetailStateRecipientComplained ActionResultEmailDetailState = "RECIPIENT_COMPLAINED"
)
```
List of ActionResultEmailDetailState

#### type EmailAction

```go
type EmailAction struct {
	Kind ActionKind `json:"kind"`
	// The name of the action, as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
	Name string `json:"name"`
	// An array of email addresses to include as recipients. Requires a special permission set for use. Please DO NOT include actual bouncing emails in automated testing.
	Addresses []string `json:"addresses,omitempty"`
	// HTML content to send as the body of the email. You can use a template in this field.
	Body *string `json:"body,omitempty"`
	// Optional text to send as the text/plain part of the email. If this field is not set for an email action, the Action service converts the value from the body field to text and sends that as the text/plain part when triggering the action.
	BodyPlainText *string `json:"bodyPlainText,omitempty"`
	// The date and time this action template was created (ISO-8601 date/time with zone).
	CreatedAt *string `json:"createdAt,omitempty"`
	// The principal that created this action template.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Optional text providing a human-friendly name for the sender. Must be less than or equal to 20 characters.
	FromName *string `json:"fromName,omitempty"`
	// An array of tenant member names, whose profile email addresses will be included as recipients.
	Members []string `json:"members,omitempty"`
	Subject *string  `json:"subject,omitempty"`
	// A human-readable title for the action. Must be less than or equal to 128 characters.
	Title *string `json:"title,omitempty"`
	// The date and time this action template was updated (ISO-8601 date/time with zone).
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The principal that updated this action template.
	UpdatedBy *string `json:"updatedBy,omitempty"`
}
```


#### type EmailActionMutable

```go
type EmailActionMutable struct {
	// An array of email addresses to include as recipients. Requires a special permission set for use. Please DO NOT include actual bouncing emails in automated testing.
	Addresses []string `json:"addresses,omitempty"`
	// HTML content to send as the body of the email. You can use a template in this field.
	Body *string `json:"body,omitempty"`
	// Optional text to send as the text/plain part of the email. If this field is not set for an email action, the Action service converts the value from the body field to text and sends that as the text/plain part when triggering the action.
	BodyPlainText *string `json:"bodyPlainText,omitempty"`
	// Optional text providing a human-friendly name for the sender. Must be less than or equal to 20 characters.
	FromName *string `json:"fromName,omitempty"`
	// An array of tenant member names, whose profile email addresses will be included as recipients.
	Members []string `json:"members,omitempty"`
	Subject *string  `json:"subject,omitempty"`
	// A human-readable title for the action. Must be less than or equal to 128 characters.
	Title *string `json:"title,omitempty"`
}
```


#### type PublicWebhookKey

```go
type PublicWebhookKey struct {
	// A PEM-formatted, ASN.1 DER-encoded PKCS#1 key.
	Key string `json:"key"`
}
```


#### type RawJsonPayload

```go
type RawJsonPayload map[string]interface{}
```


#### type Service

```go
type Service services.BaseService
```


#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new action service client from the given Config

#### func (*Service) CreateAction

```go
func (s *Service) CreateAction(action Action, resp ...*http.Response) (*Action, error)
```
CreateAction - Creates an action template. Parameters:

    action: The action template to create.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteAction

```go
func (s *Service) DeleteAction(actionName string, resp ...*http.Response) error
```
DeleteAction - Removes an action template. Parameters:

    actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetAction

```go
func (s *Service) GetAction(actionName string, resp ...*http.Response) (*Action, error)
```
GetAction - Returns a specific action template. Parameters:

    actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetActionStatus

```go
func (s *Service) GetActionStatus(actionName string, statusId string, resp ...*http.Response) (*ActionResult, error)
```
GetActionStatus - Returns the status of a triggered action. Parameters:

    actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
    statusId: The ID of the action status.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetActionStatusDetails

```go
func (s *Service) GetActionStatusDetails(actionName string, statusId string, resp ...*http.Response) ([]ActionResultEmailDetail, error)
```
GetActionStatusDetails - Returns the status details of the triggered email
action. Parameters:

    actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
    statusId: The ID of the action status.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetPublicWebhookKeys

```go
func (s *Service) GetPublicWebhookKeys(resp ...*http.Response) ([]PublicWebhookKey, error)
```
GetPublicWebhookKeys - Returns an array of one or two webhook keys. The first
key is active. The second key, if present, is expired. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListActions

```go
func (s *Service) ListActions(resp ...*http.Response) ([]Action, error)
```
ListActions - Returns the list of action templates. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) TriggerAction

```go
func (s *Service) TriggerAction(actionName string, triggerEvent TriggerEvent, resp ...*http.Response) error
```
TriggerAction - Triggers an action. Parameters:

    actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
    triggerEvent: The action payload, which should include values for any templated fields.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) TriggerActionWithStatus

```go
func (s *Service) TriggerActionWithStatus(actionName string, triggerEvent TriggerEvent) (*TriggerResponse, error)
```
TriggerActionWithStatus - Trigger an action and return a TriggerResponse with
StatusID

Parameters:

    actionName: The name of the action, as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
    triggerEvent: The action payload, which must include values for any templated fields.

#### func (*Service) UpdateAction

```go
func (s *Service) UpdateAction(actionName string, actionMutable ActionMutable, resp ...*http.Response) (*Action, error)
```
UpdateAction - Modifies an action template. Parameters:

    actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
    actionMutable: Updates to the action template.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### type ServiceError

```go
type ServiceError struct {
	Code     string                 `json:"code"`
	Message  string                 `json:"message"`
	Details  map[string]interface{} `json:"details,omitempty"`
	MoreInfo *string                `json:"moreInfo,omitempty"`
}
```


#### type Servicer

```go
type Servicer interface {
	/*
		CreateAction - Creates an action template.
		Parameters:
			action: The action template to create.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateAction(action Action, resp ...*http.Response) (*Action, error)
	/*
		DeleteAction - Removes an action template.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteAction(actionName string, resp ...*http.Response) error
	/*
		GetAction - Returns a specific action template.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetAction(actionName string, resp ...*http.Response) (*Action, error)
	/*
		GetActionStatus - Returns the status of a triggered action.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			statusId: The ID of the action status.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetActionStatus(actionName string, statusId string, resp ...*http.Response) (*ActionResult, error)
	/*
		GetActionStatusDetails - Returns the status details of the triggered email action.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			statusId: The ID of the action status.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetActionStatusDetails(actionName string, statusId string, resp ...*http.Response) ([]ActionResultEmailDetail, error)
	/*
		GetPublicWebhookKeys - Returns an array of one or two webhook keys. The first key is active. The second key, if present, is expired.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPublicWebhookKeys(resp ...*http.Response) ([]PublicWebhookKey, error)
	/*
		ListActions - Returns the list of action templates.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListActions(resp ...*http.Response) ([]Action, error)
	/*
		TriggerAction - Triggers an action.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			triggerEvent: The action payload, which should include values for any templated fields.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	TriggerAction(actionName string, triggerEvent TriggerEvent, resp ...*http.Response) error
	/*
		UpdateAction - Modifies an action template.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			actionMutable: Updates to the action template.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateAction(actionName string, actionMutable ActionMutable, resp ...*http.Response) (*Action, error)
	/*
		TriggerActionWithStatus - Trigger an action and return a TriggerResponse with StatusID

		Parameters:
			actionName: The name of the action, as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			triggerEvent: The action payload, which must include values for any templated fields.
	*/
	TriggerActionWithStatus(actionName string, triggerEvent TriggerEvent) (*TriggerResponse, error)
}
```

Servicer represents the interface for implementing all endpoints for this
service

#### type StatusState

```go
type StatusState string
```


```go
const (
	StatusStateQueued  StatusState = "QUEUED"
	StatusStateRunning StatusState = "RUNNING"
	StatusStateDone    StatusState = "DONE"
	StatusStateFailed  StatusState = "FAILED"
)
```
List of StatusState

#### type TriggerEvent

```go
type TriggerEvent struct {
	Kind           TriggerEventKind            `json:"kind"`
	ActionMetadata *TriggerEventActionMetadata `json:"actionMetadata,omitempty"`
	// string-ified ISO-8601 date/time with zone.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The principal that generated the trigger event.
	CreatedBy *string `json:"createdBy,omitempty"`
	// A unique identifier for this trigger event. Generated from a hash of all recursively-sorted event field values.
	Id      *string         `json:"id,omitempty"`
	Payload *RawJsonPayload `json:"payload,omitempty"`
	// The tenant within which the trigger event was generated.
	Tenant *string `json:"tenant,omitempty"`
	// A description of the condition that caused the trigger event.
	TriggerCondition *string `json:"triggerCondition,omitempty"`
	// The name of the trigger for which this event was created.
	TriggerName *string `json:"triggerName,omitempty"`
	// A time to live (TTL), expressed as seconds after createdAt, after which the trigger event will no longer be acted upon.
	TtlSeconds *int32 `json:"ttlSeconds,omitempty"`
}
```


#### type TriggerEventActionMetadata

```go
type TriggerEventActionMetadata struct {
	// An array of email addresses to include as recipients. Requires a special permission set for use. Please DO NOT include actual bouncing emails in automated testing.
	Addresses []string `json:"addresses,omitempty"`
	// An array of tenant member names, whose profile email addresses will be included as recipients.
	Members []string `json:"members,omitempty"`
}
```


#### type TriggerEventKind

```go
type TriggerEventKind string
```


```go
const (
	TriggerEventKindTrigger TriggerEventKind = "trigger"
)
```
List of TriggerEventKind

#### type TriggerResponse

```go
type TriggerResponse struct {
	StatusID  *string
	StatusURL *url.URL
}
```

TriggerResponse for returning status url and parsed statusID (if possible)

#### type WebhookAction

```go
type WebhookAction struct {
	Kind ActionKind `json:"kind"`
	// The name of the action, as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
	Name string `json:"name"`
	// The (possibly) templated payload body, which is POSTed to the webhookUrl when triggered.
	WebhookPayload string `json:"webhookPayload"`
	// Only HTTPS is allowed.
	WebhookUrl string `json:"webhookUrl"`
	// The date and time this action template was created (ISO-8601 date/time with zone).
	CreatedAt *string `json:"createdAt,omitempty"`
	// The principal that created this action template.
	CreatedBy *string `json:"createdBy,omitempty"`
	// A human-readable title for the action. Must be less than 128 characters.
	Title *string `json:"title,omitempty"`
	// The date and time this action template was updated (ISO-8601 date/time with zone).
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The principal that updated this action template.
	UpdatedBy      *string             `json:"updatedBy,omitempty"`
	WebhookHeaders map[string][]string `json:"webhookHeaders,omitempty"`
}
```


#### type WebhookActionMutable

```go
type WebhookActionMutable struct {
	// A human-readable title for the action. Must be less than 128 characters.
	Title          *string             `json:"title,omitempty"`
	WebhookHeaders map[string][]string `json:"webhookHeaders,omitempty"`
	// The (possibly) templated payload body, which is POSTed to the webhookUrl when triggered.
	WebhookPayload *string `json:"webhookPayload,omitempty"`
	// Only HTTPS is allowed.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}
```
