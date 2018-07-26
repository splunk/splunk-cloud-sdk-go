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

// Action defines the fields for email, sns, and webhooks as one aggregated model
type Action struct {
	// Common action fields:
	// Name of action, all actions have this field
	Name string `json:"name" binding:"required"`
	// Kind of action (email, webhook, or sns), all actions have this field
	Kind ActionKind `json:"kind" binding:"required"`
	// ID of action assigned by action service, all actions have this field
	ID string `json:"id" binding:"omitempty"`
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
	Addresses []string `json:"addresses" binding:"required"`
	// SNS action fields:
	// Topic to trigger SNS action
	Topic string `json:"topic" binding:"required"`
	// Message to send via SNS or Webhook action
	Message string `json:"message" binding:"required"`
	// Webhook action fields:
	// WebhookURL to trigger Webhook action
	WebhookURL string `json:"webhookUrl" binding:"required"`
	// Message string `json:"message" binding:"required"` (defined above)
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
