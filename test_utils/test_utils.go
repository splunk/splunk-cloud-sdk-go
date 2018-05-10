package test_utils

import (
	"time"

	"github.com/splunk/ssc-client-go/service"
)

const (
	// TestToken is the auth token used by stubby server
	TestToken = "TEST_AUTH_TOKEN"
	// TestStubbyHost is the stubby host
	TestStubbyHost = "ssc-sdk-shared-stubby:8882"
	// TestSubbyLocalHost is the stubby localhost
	TestSubbyLocalHost = "localhost:8882"
	// TestStubbySchme is the stubby scheme
	TestStubbySchme = "http"
	// TestTimeOut is the client timeout used in tests
	TestTimeOut = time.Second * 5
	// TestTenantID is the tenant id used by stubby tests
	TestTenantID = "TEST_TENANT"
)

func GetSplunkClient(local ...bool) *service.Client {
	if len(local) > 0 {
		return service.NewClient(TestTenantID, TestToken, TestStubbySchme+"://"+TestSubbyLocalHost, TestTimeOut)
	}
	return service.NewClient(TestTenantID, TestToken, TestStubbySchme+"://"+TestStubbyHost, TestTimeOut)
}
