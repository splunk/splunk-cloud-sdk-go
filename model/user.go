/*
 * Copyright © 2018 Splunk Inc.
 * SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
 * without a valid written license from Splunk Inc. is PROHIBITED.
 *
 */

package model

// User represents a user object
type User struct {
	ID                string   `json:"id"`
	FirstName         string   `json:"firstName,omitempty"`
	LastName          string   `json:"lastName,omitempty"`
	Email             string   `json:"email,omitempty"`
	Locale            string   `json:"locale,omitempty"`
	Name              string   `json:"name,omitempty"`
	TenantMemberships []string `json:"tenantMemberships,omitempty"`
}
