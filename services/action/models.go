// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package action

import (
	"net/url"

	"github.com/splunk/splunk-cloud-sdk-go/services/action/generated"
)

// Kind reflects the kinds of actions supported by the Action service
type Kind generated.Kind

const (
	// EmailKind for email actions
	EmailKind = generated.EMAIL
	// WebhookKind for webhook actions
	WebhookKind = generated.WEBHOOK
)

// UpdateFields defines the fields that may be updated for an existing Action
type UpdateFields = generated.ActionMutable

// Action defines the fields for email, sns, and webhooks as one aggregated model
type Action = generated.Action

// NewEmailAction creates a new email kind action
func NewEmailAction(name string, title string, body string, bodyPlainText string, subject string, addresses []string) *Action {
	return &Action{
		Name:          name,
		Kind:          EmailKind,
		Title:         &title,
		Body:          &body,
		BodyPlainText: &bodyPlainText,
		Subject:       &subject,
		Addresses:     &addresses,
	}
}

// NewWebhookAction creates a new webhook kind action
func NewWebhookAction(name string, title string, webhookURL string, payload string) *Action {
	return &Action{
		Name:           name,
		Kind:           WebhookKind,
		Title:          &title,
		WebhookUrl:     &webhookURL,
		WebhookPayload: &payload,
	}
}

// StatusState reflects the status of the action
type StatusState = generated.StatusState

const (
	// StatusQueued status
	StatusQueued = generated.QUEUED
	// StatusRunning status
	StatusRunning = generated.RUNNING
	// StatusDone status
	StatusDone = generated.DONE
	// StatusFailed status
	StatusFailed = generated.FAILED
)

// Status defines the state information
type Status = generated.ActionResult

// TriggerResponse for returning status url and parsed statusID (if possible)
type TriggerResponse struct {
	StatusID  *string
	StatusURL *url.URL
}

// NotificationKind defines the types of notifications
type NotificationKind = generated.NotificationKind

const (
	// SplunkEventKind for splunk event payloads
	SplunkEventKind = generated.SPLUNK_EVENT
	// RawJSONPayloadKind for raw json payloads
	RawJSONPayloadKind = generated.RAW_JSON
)

// Notification defines the action notification format
type Notification = generated.Notification

// Payload is what is sent when the action is triggered
type Payload interface{}

// RawJSONPayload specifies the format for RawJSONPayloadKind Notifications
type RawJSONPayload map[string]interface{}

// SplunkEventPayload is the payload for a notification coming from Splunk
type SplunkEventPayload = generated.SplunkEventPayload
