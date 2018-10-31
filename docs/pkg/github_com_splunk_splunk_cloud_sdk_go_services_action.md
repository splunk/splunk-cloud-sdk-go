# action
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/action"


## Usage

#### type Action

```go
type Action struct {
	UpdateFields
	// Common action fields:
	// Name of action, all actions have this field
	Name string `json:"name" binding:"required"`
	// Kind of action (email, webhook, or sns), all actions have this field
	Kind Kind `json:"kind" binding:"required"`
}
```

Action defines the fields for email, sns, and webhooks as one aggregated model

#### func  NewEmailAction

```go
func NewEmailAction(name string, htmlPart string, subjectPart string, textPart string, templateName string, addresses []string) *Action
```
NewEmailAction creates a new email kind action

#### func  NewSNSAction

```go
func NewSNSAction(name string, topic string, message string) *Action
```
NewSNSAction creates a new sns kind action

#### func  NewWebhookAction

```go
func NewWebhookAction(name string, webhookURL string, message string) *Action
```
NewWebhookAction creates a new webhook kind action

#### type Error

```go
type Error struct {
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Details  interface{} `json:"details,omitempty"`
	MoreInfo string      `json:"moreInfo,omitempty"`
}
```

Error defines format for returned errors

#### type Kind

```go
type Kind string
```

Kind reflects the kinds of actions supported by the Action service

```go
const (
	// EmailKind for email actions
	EmailKind Kind = "email"
	// WebhookKind for webhook actions
	WebhookKind Kind = "webhook"
	// SNSKind for SNS actions
	SNSKind Kind = "sns"
)
```

#### type Notification

```go
type Notification struct {
	Kind    NotificationKind `json:"kind" binding:"required"`
	Tenant  string           `json:"tenant" binding:"required"`
	Payload Payload          `json:"payload" binding:"required"`
}
```

Notification defines the action notification format

#### type NotificationKind

```go
type NotificationKind string
```

NotificationKind defines the types of notifications

```go
const (
	//SplunkEventKind for splunk event payloads
	SplunkEventKind NotificationKind = "splunkEvent"
	//RawJSONPayloadKind for raw json payloads
	RawJSONPayloadKind NotificationKind = "rawJSON"
)
```

#### type Payload

```go
type Payload interface{}
```

Payload is what is sent when the action is triggered

#### type RawJSONPayload

```go
type RawJSONPayload map[string]interface{}
```

RawJSONPayload specifies the format for RawJSONPayloadKind Notifications

#### type Service

```go
type Service services.BaseService
```

Service - A service the receives incoming notifications and uses pre-defined
templates to turn those notifications into meaningful actions

#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new action service client from the given Config

#### func (*Service) CreateAction

```go
func (s *Service) CreateAction(action Action) (*Action, error)
```
CreateAction creates an action

#### func (*Service) DeleteAction

```go
func (s *Service) DeleteAction(name string) error
```
DeleteAction deletes an action by name

#### func (*Service) GetAction

```go
func (s *Service) GetAction(name string) (*Action, error)
```
GetAction get an action by name

#### func (*Service) GetActionStatus

```go
func (s *Service) GetActionStatus(name string, statusID string) (*Status, error)
```
GetActionStatus returns an action's status by name

#### func (*Service) GetActions

```go
func (s *Service) GetActions() ([]Action, error)
```
GetActions get all actions

#### func (*Service) TriggerAction

```go
func (s *Service) TriggerAction(name string, notification Notification) (*TriggerResponse, error)
```
TriggerAction triggers an action from a notification

#### func (*Service) UpdateAction

```go
func (s *Service) UpdateAction(name string, action UpdateFields) (*Action, error)
```
UpdateAction updates and action by name

#### type SplunkEventPayload

```go
type SplunkEventPayload struct {
	Event      map[string]interface{} `json:"event" binding:"required"`
	Fields     map[string]string      `json:"fields" binding:"required"`
	Host       string                 `json:"host" binding:"required"`
	Index      string                 `json:"index" binding:"required"`
	Source     string                 `json:"source" binding:"required"`
	Sourcetype string                 `json:"sourcetype" binding:"required"`
	Time       float64                `json:"time" binding:"required"`
}
```

SplunkEventPayload is the payload for a notification coming from Splunk

#### type Status

```go
type Status struct {
	State    StatusState `json:"state"`
	StatusID string      `json:"statusId"`
	Message  string      `json:"message,omitempty"`
}
```

Status defines the state information

#### type StatusState

```go
type StatusState string
```

StatusState reflects the status of the action

```go
const (
	// StatusQueued status
	StatusQueued StatusState = "QUEUED"
	// StatusRunning status
	StatusRunning StatusState = "RUNNING"
	// StatusDone status
	StatusDone StatusState = "DONE"
	// StatusFailed status
	StatusFailed StatusState = "FAILED"
)
```

#### type TriggerResponse

```go
type TriggerResponse struct {
	StatusID  *string
	StatusURL *url.URL
}
```

TriggerResponse for returning status url and parsed statusID (if possible)

#### type UpdateFields

```go
type UpdateFields struct {
	// Email action fields:
	// HTMLPart to send via Email action
	HTMLPart string `json:"htmlPart,omitempty"`
	// SubjectPart to send via Email action
	SubjectPart string `json:"subjectPart,omitempty"`
	// TextPart to send via Email action
	TextPart string `json:"textPart,omitempty"`
	// TemplateName to send via Email action
	TemplateName string `json:"templateName,omitempty"`
	// Addresses to send to when Email action triggered
	Addresses []string `json:"addresses,omitempty"`

	// SNS action fields:
	// Topic to trigger SNS action
	Topic string `json:"topic,omitempty"`
	// Message to send via SNS or Webhook action
	Message string `json:"message,omitempty"`

	// Webhook action fields:
	// WebhookURL to trigger Webhook action
	WebhookURL string `json:"webhookUrl,omitempty"`
}
```

UpdateFields defines the fields that may be updated for an existing Action
