package util

import "time"

const (
	// TestToken is the auth token used by stubby server
	TestToken = "TEST_AUTH_TOKEN"
	// TestHost is the localhost used for adhoc testing
	TestHost = "https://localhost:8089"
	// TestStubbyHost is the stubby host
	TestStubbyHost = "ssc-sdk-shared-stubby:8882"
	// TestStubbySchme is the stubby scheme
	TestStubbySchme = "http"
	// TestTimeOut is the client timeout used in tests
	TestTimeOut = time.Second * 10
	// TestTenantID is the tenant id used by stubby tests
	TestTenantID = "TEST_TENANT"
)
