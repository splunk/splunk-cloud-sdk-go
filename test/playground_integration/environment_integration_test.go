// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package playgroundintegration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/ssc-client-go/testutils"
)

func TestIntegrationEnvironment(t *testing.T) {
	assert.NotEmpty(t, testutils.TestAuthenticationToken)
	assert.NotEmpty(t, testutils.TestSSCHost)
	assert.NotEmpty(t, testutils.TestTenantID)
	assert.NotEmpty(t, testutils.TestURLProtocol)
}
