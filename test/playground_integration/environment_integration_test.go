// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/testutils"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationEnvironment(t *testing.T) {
	assert.NotEmpty(t, testutils.TestAuthenticationToken)
	assert.NotEmpty(t, testutils.TestSplunkCloudHost)
	assert.NotEmpty(t, testutils.TestTenant)
	assert.NotEmpty(t, testutils.TestURLProtocol)
}
