// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package stubbyintegration

import (
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegrationEnvironment(t *testing.T) {
	assert.NotEmpty(t, testutils.TestAuthenticationToken)
	assert.NotEmpty(t, testutils.TestSSCHost)
	assert.NotEmpty(t, testutils.TestTenantID)
	assert.NotEmpty(t, testutils.TestURLProtocol)
}
