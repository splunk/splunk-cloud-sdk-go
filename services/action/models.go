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
)

// UpdateFields defines the fields that may be updated for an existing Action
type UpdateFields struct {
	// Common action fields:
	// Title is the human readable name title for the action. Optional.
	Title *string `json:"title,omitempty"`

	// Email action fields:
	// Body to send via Email action
	Body string `json:"body,omitempty"`
	// Optional text that will be sent as the text/plain part of this email. If this field is not set
	// for an email action, when triggering that action the Action Service will convert the value from the body
	// field to text and send that as the text/plain part.
	BodyPlainText string `json:"bodyPlainText,omitempty"`
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

// Action defines the fields for email, sns, and webhooks as one aggregated model
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

// NewEmailAction creates a new email kind action
func NewEmailAction(name string, title string, body string, bodyPlainText string, subject string, addresses []string) *Action {
	return &Action{
		Name: name,
		Kind: EmailKind,
		UpdateFields: UpdateFields{
			Title:     &title,
			Body:      body,
			BodyPlainText: bodyPlainText,
			Subject:   subject,
			Addresses: addresses,
		},
	}
}

// NewWebhookAction creates a new webhook kind action
func NewWebhookAction(name string, title string, webhookURL string, payload string) *Action {
	return &Action{
		Name: name,
		Kind: WebhookKind,
		UpdateFields: UpdateFields{
			Title:          &title,
			WebhookURL:     webhookURL,
			WebhookPayload: payload,
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

// NotificationKind defines the types of notifications
type NotificationKind string

const (
	// SplunkEventKind for splunk event payloads
	SplunkEventKind NotificationKind = "splunkEvent"
	// RawJSONPayloadKind for raw json payloads
	RawJSONPayloadKind NotificationKind = "rawJSON"
)

// Notification defines the action notification format
type Notification struct {
	Kind    NotificationKind `json:"kind"`
	Tenant  string           `json:"tenant"`
	Payload Payload          `json:"payload"`
}

// Payload is what is sent when the action is triggered
type Payload interface{}

// RawJSONPayload specifies the format for RawJSONPayloadKind Notifications
type RawJSONPayload map[string]interface{}

// SplunkEventPayload is the payload for a notification coming from Splunk
type SplunkEventPayload struct {
	Event      map[string]interface{} `json:"event"`
	Fields     map[string]string      `json:"fields"`
	Host       string                 `json:"host"`
	Index      string                 `json:"index"`
	Source     string                 `json:"source"`
	Sourcetype string                 `json:"sourcetype"`
	Time       float64                `json:"time"`
}
