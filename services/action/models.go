// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package action

import (
	"net/url"
)

// Kind reflects the kinds of actions supported by the Action service
type Kind string

const (
	// EmailKind for email actions
	EmailKind Kind = "email"
	// WebhookKind for webhook actions
	WebhookKind Kind = "webhook"
	// SNSKind for SNS actions
	SNSKind Kind = "sns"
)

// UpdateFields defines the fields that may be updated for an existing Action
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
	// Message string `json:"message"`(defined above)
}

// Action defines the fields for email, sns, and webhooks as one aggregated model
type Action struct {
	ActionUpdateFields UpdateFields
	// Common action fields:
	// Name of action, all actions have this field
	Name string `json:"name" binding:"required"`
	// Kind of action (email, webhook, or sns), all actions have this field
	Kind Kind `json:"kind" binding:"required"`
}

// NewEmailAction creates a new email kind action
func NewEmailAction(name string, htmlPart string, subjectPart string, textPart string, templateName string, addresses []string) *Action {
	return &Action{
		Name: name,
		Kind: EmailKind,
		ActionUpdateFields: UpdateFields{
			HTMLPart:     htmlPart,
			SubjectPart:  subjectPart,
			TextPart:     textPart,
			TemplateName: templateName,
			Addresses:    addresses,
		},
	}
}

// NewSNSAction creates a new sns kind action
func NewSNSAction(name string, topic string, message string) *Action {
	return &Action{
		Name: name,
		Kind: SNSKind,
		ActionUpdateFields: UpdateFields{
			Topic:   topic,
			Message: message,
		},
	}
}

// NewWebhookAction creates a new webhook kind action
func NewWebhookAction(name string, webhookURL string, message string) *Action {
	return &Action{
		Name: name,
		Kind: WebhookKind,
		ActionUpdateFields: UpdateFields{
			WebhookURL: webhookURL,
			Message:    message,
		},
	}
}

// StatusState reflects the status of the action
type StatusState string

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

// Status defines the state information
type Status struct {
	State    StatusState `json:"state"`
	StatusID string      `json:"statusId"`
	Message  string      `json:"message,omitempty"`
}

// TriggerResponse for returning status url and parsed statusID (if possible)
type TriggerResponse struct {
	StatusID  *string
	StatusURL *url.URL
}

// Error defines format for returned errors
type Error struct {
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Details  interface{} `json:"details,omitempty"`
	MoreInfo string      `json:"moreInfo,omitempty"`
}

// NotificationKind defines the types of notifications
type NotificationKind string

const (
	//SplunkEventKind for splunk event payloads
	SplunkEventKind NotificationKind = "splunkEvent"
	//RawJSONPayloadKind for raw json payloads
	RawJSONPayloadKind NotificationKind = "rawJSON"
)

// Notification defines the action notification format
type Notification struct {
	Kind    NotificationKind `json:"kind" binding:"required"`
	Tenant  string           `json:"tenant" binding:"required"`
	Payload Payload          `json:"payload" binding:"required"`
}

// Payload is what is sent when the action is triggered
type Payload interface{}

// RawJSONPayload specifies the format for RawJSONPayloadKind Notifications
type RawJSONPayload map[string]interface{}

// SplunkEventPayload is the payload for a notification coming from Splunk
type SplunkEventPayload struct {
	Event      map[string]interface{} `json:"event" binding:"required"`
	Fields     map[string]string      `json:"fields" binding:"required"`
	Host       string                 `json:"host" binding:"required"`
	Index      string                 `json:"index" binding:"required"`
	Source     string                 `json:"source" binding:"required"`
	Sourcetype string                 `json:"sourcetype" binding:"required"`
	Time       float64                `json:"time" binding:"required"`
}
