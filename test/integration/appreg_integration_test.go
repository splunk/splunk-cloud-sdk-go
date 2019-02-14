// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test GetActions which returns the list of all actions for the tenant
func TestIntegrationGetApps(t *testing.T) {
	client := getSdkClient(t)

	// Get Actions
	actions, err := client.AppRegService.ListApps()
	require.Nil(t, err)
	assert.True(t, len(actions) >= 0)
}
