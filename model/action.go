// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model

import (
	"net/url"
)

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

// ActionUpdateFields defines the fields that may be updated for an existing Action
type ActionUpdateFields struct {
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

// Action defines the fields for email, sns, and webhooks as one aggregated model
type Action struct {
	// Common action fields:
	// Name of action, all actions have this field
	Name string `json:"name" binding:"required"`
	// Kind of action (email, webhook, or sns), all actions have this field
	Kind ActionKind `json:"kind" binding:"required"`
	ActionUpdateFields
}

// NewEmailAction creates a new email kind action
func NewEmailAction(name string, htmlPart string, subjectPart string, textPart string, templateName string, addresses []string) *Action {
	return &Action{
		Name: name,
		Kind: EmailKind,
		ActionUpdateFields: ActionUpdateFields{
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
		ActionUpdateFields: ActionUpdateFields{
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
		ActionUpdateFields: ActionUpdateFields{
			WebhookURL: webhookURL,
			Message:    message,
		},
	}
}

// ActionStatusState reflects the status of the action
type ActionStatusState string

const (
	// StatusQueued status
	StatusQueued ActionStatusState = "QUEUED"
	// StatusRunning status
	StatusRunning ActionStatusState = "RUNNING"
	// StatusDone status
	StatusDone ActionStatusState = "DONE"
	// StatusFailed status
	StatusFailed ActionStatusState = "FAILED"
)

// ActionStatus defines the state information
type ActionStatus struct {
	State    ActionStatusState `json:"state"`
	StatusID string            `json:"statusId"`
	Message  string            `json:"message,omitempty"`
}

// ActionTriggerResponse for returning status url and parsed statusID (if possible)
type ActionTriggerResponse struct {
	StatusID  *string
	StatusURL *url.URL
}

// ActionError defines format for returned errors
type ActionError struct {
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Details  interface{} `json:"details,omitempty"`
	MoreInfo string      `json:"moreInfo,omitempty"`
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
	Kind    ActionNotificationKind `json:"kind" binding:"required"`
	Tenant  string                 `json:"tenant" binding:"required"`
	Payload ActionPayload          `json:"payload" binding:"required"`
}

// ActionPayload is what is sent when the action is triggered
type ActionPayload interface{}

// RawJSONPayload specifies the format for RawJSONPayloadKind ActionNotifications
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
