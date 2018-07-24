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
