/*
 * Copyright © 2020 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * Action Service
 *
 * With the Action service in Splunk Cloud Services, you can receive incoming trigger events and use pre-defined action templates to turn these events into meaningful actions.
 *
 * API version: v1beta2.11 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package action

import (
	"bytes"
	"encoding/json"
)

type Action struct {
	emailAction   *EmailAction
	webhookAction *WebhookAction
	raw           interface{}
}

// MakeActionFromEmailAction creates a new Action from an instance of EmailAction
func MakeActionFromEmailAction(f EmailAction) Action {
	return Action{emailAction: &f}
}

// IsEmailAction checks if the Action is a EmailAction
func (m Action) IsEmailAction() bool {
	return m.emailAction != nil
}

// EmailAction returns EmailAction if IsEmailAction() is true, nil otherwise
func (m Action) EmailAction() *EmailAction {
	return m.emailAction
}

// MakeActionFromWebhookAction creates a new Action from an instance of WebhookAction
func MakeActionFromWebhookAction(f WebhookAction) Action {
	return Action{webhookAction: &f}
}

// IsWebhookAction checks if the Action is a WebhookAction
func (m Action) IsWebhookAction() bool {
	return m.webhookAction != nil
}

// WebhookAction returns WebhookAction if IsWebhookAction() is true, nil otherwise
func (m Action) WebhookAction() *WebhookAction {
	return m.webhookAction
}

// MakeActionFromRawInterface creates a new Action from a raw interface{}
func MakeActionFromRawInterface(f interface{}) Action {
	return Action{raw: f}
}

// IsRawInterface checks if the Action is an interface{} (unknown type)
func (m Action) IsRawInterface() bool {
	return m.raw != nil
}

// RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil otherwise
func (m Action) RawInterface() interface{} {
	return m.raw
}

// UnmarshalJSON unmarshals Action using the "kind" property
func (m *Action) UnmarshalJSON(b []byte) (err error) {
	type discriminator struct {
		Kind string `json:"kind"`
	}
	var d discriminator
	err = json.Unmarshal(b, &d)
	if err != nil {
		return err
	}
	// Resolve into respective struct based on the discriminator value
	switch d.Kind {
	case "email":
		m.emailAction = &EmailAction{}
		return json.Unmarshal(b, m.emailAction)
	case "webhook":
		m.webhookAction = &WebhookAction{}
		return json.Unmarshal(b, m.webhookAction)
	}
	// Unknown discriminator value (this type may not yet be supported)
	// unmarhsal to raw interface
	var raw interface{}
	err = json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}
	m.raw = raw
	return nil
}

// MarshalJSON marshals Action using the appropriate struct field
func (m Action) MarshalJSON() ([]byte, error) {
	if m.IsEmailAction() {
		return json.Marshal(m.emailAction)
	} else if m.IsWebhookAction() {
		return json.Marshal(m.webhookAction)
	}
	// None of the structs are populated, send raw
	return json.Marshal(m.raw)
}

type ActionImmutable struct {
	Kind ActionKind `json:"kind"`
	// The name of the action, as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
	Name string `json:"name"`
	// The date and time this action template was created (ISO-8601 date/time with zone).
	CreatedAt *string `json:"createdAt,omitempty"`
	// The principal that created this action template.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The date and time this action template was updated (ISO-8601 date/time with zone).
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The principal that updated this action template.
	UpdatedBy *string `json:"updatedBy,omitempty"`
}

// ActionKind :
type ActionKind string

// List of ActionKind
const (
	ActionKindWebhook ActionKind = "webhook"
	ActionKindEmail   ActionKind = "email"
)

// ActionMutable is EmailActionMutable, WebhookActionMutable, (or interface{} if no matches are found)
type ActionMutable struct {
	actionMutable interface{}
	isRaw         bool
}

// UnmarshalJSON unmarshals ActionMutable into EmailActionMutable, WebhookActionMutable, or interface{} if no matches are found
func (m *ActionMutable) UnmarshalJSON(b []byte) (err error) {
	reader := bytes.NewReader(b)
	d := json.NewDecoder(reader)
	d.DisallowUnknownFields()
	// Attempt to unmarshal to each oneOf, if unknown fields then move to next
	attempt := func(m interface{}) error {
		_, err = reader.Seek(0, 0)
		if err != nil {
			return err
		}
		return d.Decode(m)
	}
	var testEmailActionMutable EmailActionMutable
	if err = attempt(&testEmailActionMutable); err == nil {
		m.actionMutable = testEmailActionMutable
		return nil
	}
	var testWebhookActionMutable WebhookActionMutable
	if err = attempt(&testWebhookActionMutable); err == nil {
		m.actionMutable = testWebhookActionMutable
		return nil
	}
	// If no matches, decode model to raw interface
	var raw interface{}
	err = attempt(&raw)
	if err != nil {
		return err
	}
	m.isRaw = true
	m.actionMutable = raw
	return nil
}

// MarshalJSON marshals ActionMutable using ActionMutable.ActionMutable
func (m ActionMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.actionMutable)
}

// MakeActionMutableFromEmailActionMutable creates a new ActionMutable from an instance of EmailActionMutable
func MakeActionMutableFromEmailActionMutable(f EmailActionMutable) ActionMutable {
	return ActionMutable{actionMutable: f}
}

// IsEmailActionMutable checks if the ActionMutable is a EmailActionMutable
func (m ActionMutable) IsEmailActionMutable() bool {
	_, ok := m.actionMutable.(EmailActionMutable)
	return ok
}

// EmailActionMutable returns EmailActionMutable if IsEmailActionMutable() is true, nil otherwise
func (m ActionMutable) EmailActionMutable() *EmailActionMutable {
	if v, ok := m.actionMutable.(EmailActionMutable); ok {
		return &v
	}
	return nil
}

// MakeActionMutableFromWebhookActionMutable creates a new ActionMutable from an instance of WebhookActionMutable
func MakeActionMutableFromWebhookActionMutable(f WebhookActionMutable) ActionMutable {
	return ActionMutable{actionMutable: f}
}

// IsWebhookActionMutable checks if the ActionMutable is a WebhookActionMutable
func (m ActionMutable) IsWebhookActionMutable() bool {
	_, ok := m.actionMutable.(WebhookActionMutable)
	return ok
}

// WebhookActionMutable returns WebhookActionMutable if IsWebhookActionMutable() is true, nil otherwise
func (m ActionMutable) WebhookActionMutable() *WebhookActionMutable {
	if v, ok := m.actionMutable.(WebhookActionMutable); ok {
		return &v
	}
	return nil
}

// MakeActionMutableFromRawInterface creates a new ActionMutable from a raw interface{}
func MakeActionMutableFromRawInterface(f interface{}) ActionMutable {
	return ActionMutable{
		actionMutable: f,
		isRaw:         true,
	}
}

// IsRawInterface checks if the ActionMutable is an interface{} (unknown type)
func (m ActionMutable) IsRawInterface() bool {
	return m.isRaw
}

// RawInterface returns interface{} if IsRawInterface() is true (unknown type), nil otherwise
func (m ActionMutable) RawInterface() interface{} {
	if !m.IsRawInterface() {
		return nil
	}
	return m.actionMutable
}

type ActionResult struct {
	ActionName string      `json:"actionName"`
	State      StatusState `json:"state"`
	StatusId   string      `json:"statusId"`
	Message    *string     `json:"message,omitempty"`
}

type ActionResultEmailDetail struct {
	EmailAddress *string                       `json:"emailAddress,omitempty"`
	State        *ActionResultEmailDetailState `json:"state,omitempty"`
}

type ActionResultEmailDetailState string

// List of ActionResultEmailDetailState
const (
	ActionResultEmailDetailStatePending             ActionResultEmailDetailState = "PENDING"
	ActionResultEmailDetailStateNotFound            ActionResultEmailDetailState = "NOT_FOUND"
	ActionResultEmailDetailStateSucceeded           ActionResultEmailDetailState = "SUCCEEDED"
	ActionResultEmailDetailStateBounced             ActionResultEmailDetailState = "BOUNCED"
	ActionResultEmailDetailStateRecipientComplained ActionResultEmailDetailState = "RECIPIENT_COMPLAINED"
)

type EmailAction struct {
	Kind ActionKind `json:"kind"`
	// The name of the action, as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
	Name string `json:"name"`
	// An array of email addresses to include as recipients. Requires a special permission set for use. Please DO NOT include actual bouncing emails in automated testing.
	Addresses []string `json:"addresses,omitempty"`
	// HTML content to send as the body of the email. You can use a template in this field.
	Body *string `json:"body,omitempty"`
	// Optional text to send as the text/plain part of the email. If this field is not set for an email action, the Action service converts the value from the body field to text and sends that as the text/plain part when invoking the action. You can use a template in this field.
	BodyPlainText *string `json:"bodyPlainText,omitempty"`
	// The date and time this action template was created (ISO-8601 date/time with zone).
	CreatedAt *string `json:"createdAt,omitempty"`
	// The principal that created this action template.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Optional text providing a human-friendly name for the sender. Must be less than or equal to 81 characters. You can use a template in this field.
	FromName *string `json:"fromName,omitempty"`
	// An array of tenant member names, whose profile email addresses will be included as recipients.
	Members []string `json:"members,omitempty"`
	// The subject of the email. You can use a template in this field.
	Subject *string `json:"subject,omitempty"`
	// A human-readable title for the action. Must be less than or equal to 128 characters.
	Title *string `json:"title,omitempty"`
	// The date and time this action template was updated (ISO-8601 date/time with zone).
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The principal that updated this action template.
	UpdatedBy *string `json:"updatedBy,omitempty"`
}

type EmailActionMutable struct {
	// An array of email addresses to include as recipients. Requires a special permission set for use. Please DO NOT include actual bouncing emails in automated testing.
	Addresses []string `json:"addresses,omitempty"`
	// HTML content to send as the body of the email. You can use a template in this field.
	Body *string `json:"body,omitempty"`
	// Optional text to send as the text/plain part of the email. If this field is not set for an email action, the Action service converts the value from the body field to text and sends that as the text/plain part when invoking the action. You can use a template in this field.
	BodyPlainText *string `json:"bodyPlainText,omitempty"`
	// Optional text providing a human-friendly name for the sender. Must be less than or equal to 81 characters. You can use a template in this field.
	FromName *string `json:"fromName,omitempty"`
	// An array of tenant member names, whose profile email addresses will be included as recipients.
	Members []string `json:"members,omitempty"`
	// The subject of the email. You can use a template in this field.
	Subject *string `json:"subject,omitempty"`
	// A human-readable title for the action. Must be less than or equal to 128 characters.
	Title *string `json:"title,omitempty"`
}

type PublicWebhookKey struct {
	// A PEM-formatted, ASN.1 DER-encoded PKCS#1 key.
	Key string `json:"key"`
}

type RawJsonPayload map[string]interface{}

type ServiceError struct {
	Code     string                 `json:"code"`
	Message  string                 `json:"message"`
	Details  map[string]interface{} `json:"details,omitempty"`
	MoreInfo *string                `json:"moreInfo,omitempty"`
}

// StatusState :
type StatusState string

// List of StatusState
const (
	StatusStateQueued  StatusState = "QUEUED"
	StatusStateRunning StatusState = "RUNNING"
	StatusStateDone    StatusState = "DONE"
	StatusStateFailed  StatusState = "FAILED"
)

type TriggerEvent struct {
	ActionMetadata *TriggerEventActionMetadata `json:"actionMetadata,omitempty"`
	// string-ified ISO-8601 date/time with zone.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The principal that generated the trigger event.
	CreatedBy *string `json:"createdBy,omitempty"`
	// A unique identifier for this trigger event. Generated from a hash of all recursively-sorted event field values.
	Id      *string           `json:"id,omitempty"`
	Kind    *TriggerEventKind `json:"kind,omitempty"`
	Payload *RawJsonPayload   `json:"payload,omitempty"`
	// The tenant within which the trigger event was generated.
	Tenant *string `json:"tenant,omitempty"`
	// A description of the condition that caused the trigger event.
	TriggerCondition *string `json:"triggerCondition,omitempty"`
	// The name of the trigger for which this event was created.
	TriggerName *string `json:"triggerName,omitempty"`
	// A time to live (TTL), expressed as seconds after createdAt, after which the trigger event will no longer be acted upon.
	TtlSeconds *int32 `json:"ttlSeconds,omitempty"`
}

type TriggerEventActionMetadata struct {
	// An array of email addresses to include as recipients. Requires a special permission set for use. Please DO NOT include actual bouncing emails in automated testing.
	Addresses []string `json:"addresses,omitempty"`
	// An array of tenant member names, whose profile email addresses will be included as recipients.
	Members []string `json:"members,omitempty"`
}

// TriggerEventKind :
type TriggerEventKind string

// List of TriggerEventKind
const (
	TriggerEventKindTrigger TriggerEventKind = "trigger"
)

type WebhookAction struct {
	Kind ActionKind `json:"kind"`
	// The name of the action, as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
	Name string `json:"name"`
	// The (possibly) templated payload body, which is POSTed to the webhookUrl when invoked.
	WebhookPayload string `json:"webhookPayload"`
	// Only HTTPS is allowed.
	WebhookUrl string `json:"webhookUrl"`
	// The date and time this action template was created (ISO-8601 date/time with zone).
	CreatedAt *string `json:"createdAt,omitempty"`
	// The principal that created this action template.
	CreatedBy *string `json:"createdBy,omitempty"`
	// A human-readable title for the action. Must be less than 128 characters.
	Title *string `json:"title,omitempty"`
	// The date and time this action template was updated (ISO-8601 date/time with zone).
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The principal that updated this action template.
	UpdatedBy      *string             `json:"updatedBy,omitempty"`
	WebhookHeaders map[string][]string `json:"webhookHeaders,omitempty"`
}

type WebhookActionMutable struct {
	// A human-readable title for the action. Must be less than 128 characters.
	Title          *string             `json:"title,omitempty"`
	WebhookHeaders map[string][]string `json:"webhookHeaders,omitempty"`
	// The (possibly) templated payload body, which is POSTed to the webhookUrl when invoked.
	WebhookPayload *string `json:"webhookPayload,omitempty"`
	// Only HTTPS is allowed.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}
