// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCredentialOutputRedactedString(t *testing.T) {
	pw := NewCredential("mypassword")
	assert.Equal(t, fmt.Sprintf("%s", pw), "XXXXX")
	assert.Equal(t, fmt.Sprintf("%s", pw.ClearText()), "mypassword")
}
