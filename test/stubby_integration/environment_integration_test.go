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
