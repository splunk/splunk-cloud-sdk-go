package testutils

import (
	"os"
	"time"
)

// TestURLProtocol - the url protocol for the test api to be used
var TestURLProtocol = os.Getenv("TEST_URL_PROTOCOL")

// TestSSCHost - the url for the test api to be used
var TestSSCHost = os.Getenv("TEST_SSC_HOST")

// TestAuthenticationToken - the authentication that gives permission to make requests against the api
var TestAuthenticationToken = os.Getenv("TEST_BEARER_TOKEN")

// TestTenantID - the tenant to be used for the API
var TestTenantID = os.Getenv("TEST_TENANT_ID")

// TestInvalidAuthenticationToken - the invalid authentication that denies permission to make requests against the api
var TestInvalidAuthenticationToken = "INVALID_TEST_AUTH_TOKEN"

// TestTimeOut - the timeout to be used for requests to the api
var TestTimeOut = time.Second * 5
