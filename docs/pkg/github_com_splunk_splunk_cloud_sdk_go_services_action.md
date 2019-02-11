# action
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/action"


## Usage

#### type Action

```go
type Action struct {
	UpdateFields
	// Common action fields:
	// Name of the action. The name is one or more identifier strings separated by dots. Each
	// identifier string consists of lower case letters, digits, and underscores, and cannot start
	// with a digit.
	Name string `json:"name"`
	// Kind of action (email or webhook)
	Kind Kind `json:"kind"`
}
```

Action defines the fields for email, sns, and webhooks as one aggregated model

#### func  NewEmailAction

```go
func NewEmailAction(name string, title string, body string, subject string, addresses []string) *Action
```
NewEmailAction creates a new email kind action

#### func  NewWebhookAction

```go
func NewWebhookAction(name string, title string, webhookURL string, payload string) *Action
```
NewWebhookAction creates a new webhook kind action

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
)
```

#### type Notification

```go
type Notification struct {
	Kind    NotificationKind `json:"kind"`
	Tenant  string           `json:"tenant"`
	Payload Payload          `json:"payload"`
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
	// SplunkEventKind for splunk event payloads
	SplunkEventKind NotificationKind = "splunkEvent"
	// RawJSONPayloadKind for raw json payloads
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

#### type Servicer

```go
type Servicer interface {
	// GetActions get all actions
	GetActions() ([]Action, error)
	// CreateAction creates an action
	CreateAction(action Action) (*Action, error)
	// GetAction get an action by name
	GetAction(name string) (*Action, error)
	// TriggerAction triggers an action from a notification
	TriggerAction(name string, notification Notification) (*TriggerResponse, error)
	// UpdateAction updates and action by name
	UpdateAction(name string, action UpdateFields) (*Action, error)
	// DeleteAction deletes an action by name
	DeleteAction(name string) error
	// GetActionStatus returns an action's status by name
	GetActionStatus(name string, statusID string) (*Status, error)
}
```

Servicer ...

#### type SplunkEventPayload

```go
type SplunkEventPayload struct {
	Event      map[string]interface{} `json:"event"`
	Fields     map[string]string      `json:"fields"`
	Host       string                 `json:"host"`
	Index      string                 `json:"index"`
	Source     string                 `json:"source"`
	Sourcetype string                 `json:"sourcetype"`
	Time       float64                `json:"time"`
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
	// Common action fields:
	// Title is the human readable name title for the action. Optional.
	Title *string `json:"title,omitempty"`

	// Email action fields:
	// Body to send via Email action
	Body string `json:"body,omitempty"`
	// Subject to send via Email action
	Subject string `json:"subject,omitempty"`
	// Addresses to send to when Email action triggered (required for Email actions)
	Addresses []string `json:"addresses,omitempty"`

	// Webhook action fields:
	// WebhookPayload is the (possibly) templated payload body which will be POSTed to the webhookUrl when triggered
	WebhookPayload string `json:"webhookPayload,omitempty"`
	// WebhookURL to trigger Webhook action, only allows the HTTPS scheme
	WebhookURL string `json:"webhookUrl,omitempty"`
}
```

UpdateFields defines the fields that may be updated for an existing Action
