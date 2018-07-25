package model

// ActionKind reflects the kinds of actions supported by the Action service
type ActionKind string

const (
	// EmailKind for email actions
	EmailKind ActionKind = "email"
	// WebhookKind for webhook actions
	WebhookKind ActionKind = "webhook"
	// SNSKind for SNS actions
	SNSKind ActionKind = "sns"
)

// ActionBase is the base struct that all actions must inherit
type ActionBase struct {
	Name string     `json:"name" binding:"required"`
	Kind ActionKind `json:"kind" binding:"required"`
	ID   string     `json:"id" binding:"omitempty"`
}

// EmailAction defines email action kinds
type EmailAction struct {
	*ActionBase
	HTMLPart     string   `json:"htmlPart,omitempty"`
	SubjectPart  string   `json:"subjectPart,omitempty"`
	TextPart     string   `json:"textPart,omitempty"`
	TemplateName string   `json:"templateName,omitempty"`
	Addresses    []string `json:"addresses" binding:"required"`
}

// SNSAction defines SNS action kinds
type SNSAction struct {
	*ActionBase
	Topic   string `json:"topic" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// WebhookAction an action to run webhooks
type WebhookAction struct {
	*ActionBase
	WebhookURL string `json:"webhookUrl" binding:"required"`
	Message    string `json:"message" binding:"required"`
}

// ActionStatusState reflects the status of the action
type ActionStatusState string

const (
	// StatusQueued status
	StatusQueued ActionStatusState = "QUEUED"
	// StatusInProgress status
	StatusInProgress ActionStatusState = "IN PROGRESS"
	// StatusDone status
	StatusDone ActionStatusState = "DONE"
	// StatusFailed status
	StatusFailed ActionStatusState = "FAILED"
)

// ActionStatus defines the state information
type ActionStatus struct {
	State   ActionStatusState `json:"state"`
	ID      string            `json:"statusId"`
	Message string            `json:"message,omitempty"`
}

// ActionError defines format for returned errors
type ActionError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ActionNotificationKind defines the types of notifications
type ActionNotificationKind string

const (
	//SplunkEventKind for splunk event payloads
	SplunkEventKind ActionNotificationKind = "splunkEvent"
	//RawJSONPayloadKind for raw json payloads
	RawJSONPayloadKind ActionNotificationKind = "rawJSON"
)

// ActionNotification defines the action notification format
type ActionNotification struct {
	EmailImmediately bool   `json:"emailImmediately,omitempty"`
	Severity         int    `json:"severity" binding:"omitempty,min=0,max=10"`
	Kind             string `json:"kind" binding:"required"`
	Tenant           string `json:"tenant" binding:"required"`
	UserID           string `json:"userId" binding:"required"`
	Payload          interface{}
}

// SplunkEventPayload is the payload for a notification coming from Splunk
type SplunkEventPayload struct {
	Index      string `json:"index" binding:"required"`
	Host       string `json:"host" binding:"required"`
	Source     string `json:"source" binding:"required"`
	Sourcetype string `json:"sourcetype" binding:"required"`
	Raw        string `json:"_raw" binding:"required"`
	Time       string `json:"_time"` // if "required", value can't be 0=
}
