package stubbyintegration

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
