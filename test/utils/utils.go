// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package utils

import (
	"fmt"
	"os"
	"time"
)

// TimeSec denotes the current timestamp
var TimeSec = time.Now().Unix()

// TestSplunkCloudHost - the url for the test api to be used
var TestSplunkCloudHost = os.Getenv("SPLUNK_CLOUD_HOST")

// TestAuthenticationToken - the authentication that gives permission to make requests against the api
var TestAuthenticationToken = os.Getenv("BEARER_TOKEN")

// TestTenant - the tenant to be used for the API
var TestTenant = os.Getenv("TENANT_ID")

// TestUsername - the user running tests on behalf of
var TestUsername = os.Getenv("TEST_USERNAME")

// TestInvalidTestTenant - the invalid tenant ID that denies permission to make requests against the api
var TestInvalidTestTenant = "INVALID_TENANT_ID"

// ExpiredAuthenticationToken - to test authentication retries
var ExpiredAuthenticationToken = os.Getenv("EXPIRED_BEARER_TOKEN")

// TestModule - A namespace for integration testing
var TestModule = fmt.Sprintf("gomod%d", TimeSec)

// TestCollection - A collection for integration testing
var TestCollection = fmt.Sprintf("gocollection%d", TimeSec)

// StubbyTestCollection - A collection for stubby testing
var StubbyTestCollection = "testcollection0"

// TestTimeOut - the timeout to be used for requests to the api
var TestTimeOut = time.Second * 5

// TenantCreationOn specifies whether tenants should be created on the fly for identity service /tenant CRUD testing
var TenantCreationOn = os.Getenv("TENANT_CREATION") == "1"
