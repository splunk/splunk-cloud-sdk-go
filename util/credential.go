// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package util

// Credential is a simple string whose value is redacted when converted to a
// string via the Stringer interface in order to prevent accidental logging or
// other unintentional disclosure - value is retrieved using ClearText() method
type Credential struct {
	string
}

// NewCredential creates a Credential from a simple string
func NewCredential(s string) *Credential {
	return &Credential{s}
}

// String returns a redacted string
func (c *Credential) String() string {
	return "XXXXX"
}

// ClearText returns the actual cleartext string value
func (c *Credential) ClearText() string {
	return c.string
}
