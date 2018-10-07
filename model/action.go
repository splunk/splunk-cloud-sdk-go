// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model

import (
	"github.com/splunk/splunk-cloud-sdk-go/services/action"
)

// ActionKind is DEPRECATED, please use services/action.Kind
type ActionKind = action.Kind

const (
	// EmailKind is DEPRECATED, please use services/action.EmailKind
	EmailKind = action.EmailKind
	// WebhookKind is DEPRECATED, please use services/action.WebhookKind
	WebhookKind = action.WebhookKind
	// SNSKind is DEPRECATED, please use services/action.SNSKind
	SNSKind = action.SNSKind
)

// ActionUpdateFields is DEPRECATED, please use services/action.UpdateFields
type ActionUpdateFields = action.UpdateFields

// Action is DEPRECATED, please use services/action.Action
type Action = action.Action

// NewEmailAction is DEPRECATED, please use services/action.NewEmailAction
func NewEmailAction(name string, htmlPart string, subjectPart string, textPart string, templateName string, addresses []string) *Action {
	return action.NewEmailAction(name, htmlPart, subjectPart, textPart, templateName, addresses)
}

// NewSNSAction is DEPRECATED, please use services/action.NewSNSAction
func NewSNSAction(name string, topic string, message string) *Action {
	return action.NewSNSAction(name, topic, message)
}

// NewWebhookAction is DEPRECATED, please use services/action.NewWebhookAction
func NewWebhookAction(name string, webhookURL string, message string) *Action {
	return action.NewWebhookAction(name, webhookURL, message)
}

// ActionStatusState is DEPRECATED, please use services/action.StatusState
type ActionStatusState = action.StatusState

const (
	// StatusQueued is DEPRECATED, please use services/action.StatusQueued
	StatusQueued = action.StatusQueued
	// StatusRunning is DEPRECATED, please use services/action.StatusRunning
	StatusRunning = action.StatusRunning
	// StatusDone is DEPRECATED, please use services/action.StatusDone
	StatusDone = action.StatusDone
	// StatusFailed is DEPRECATED, please use services/action.StatusFailed
	StatusFailed = action.StatusFailed
)

// ActionStatus is DEPRECATED, please use services/action.Status
type ActionStatus = action.Status

// ActionTriggerResponse is DEPRECATED, please use services/action.TriggerResponse
type ActionTriggerResponse = action.TriggerResponse

// ActionError is DEPRECATED, please use services/action.Error
type ActionError = action.Error

// ActionNotificationKind is DEPRECATED, please use services/action.NotificationKind
type ActionNotificationKind = action.NotificationKind

const (
	//SplunkEventKind is DEPRECATED, please use services/action.SplunkEventKind
	SplunkEventKind = action.SplunkEventKind
	//RawJSONPayloadKind is DEPRECATED, please use services/action.RawJSONPayloadKind
	RawJSONPayloadKind = action.RawJSONPayloadKind
)

// ActionNotification is DEPRECATED, please use services/action.Notification
type ActionNotification = action.Notification

// ActionPayload is DEPRECATED, please use services/action.Payload
type ActionPayload = action.Payload

// RawJSONPayload is DEPRECATED, please use services/action.RawJSONPayload
type RawJSONPayload = action.RawJSONPayload

// SplunkEventPayload is DEPRECATED, please use services/action.SplunkEventPayload
type SplunkEventPayload = action.SplunkEventPayload
