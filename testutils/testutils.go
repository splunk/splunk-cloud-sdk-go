// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

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

// TestInvalidTestTenantID - the invalid tenant ID that denies permission to make requests against the api
var TestInvalidTestTenantID = "INVALID_TEST_TENANT_ID"

// TestNamespace - A namespace for testing
var TestNamespace = "testnamespace0"

// TestCollection - A collection for testing
var TestCollection = "testcollection0"

// TestTimeOut - the timeout to be used for requests to the api
var TestTimeOut = time.Second * 5
