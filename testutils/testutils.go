// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package testutils

import (
	"fmt"
	"os"
	"time"
)

// timeSec denotes the current timestamp
var timeSec = time.Now().Unix()

// TestURLProtocol - the url protocol for the test api to be used
var TestURLProtocol = os.Getenv("TEST_URL_PROTOCOL")

// TestSplunkCloudHost - the url for the test api to be used
var TestSplunkCloudHost = os.Getenv("TEST_SPLUNK_CLOUD_HOST")

// TestAuthenticationToken - the authentication that gives permission to make requests against the api
var TestAuthenticationToken = os.Getenv("TEST_BEARER_TOKEN")

// TestInvalidAuthenticationToken - the invalid access token that denies permission to make requests against the api
var TestInvalidAuthenticationToken = "INVALID_TOKEN"

// TestTenantID - the tenant to be used for the API
var TestTenantID = os.Getenv("TEST_TENANT_ID")

// TestInvalidTestTenantID - the invalid tenant ID that denies permission to make requests against the api
var TestInvalidTestTenantID = "INVALID_TEST_TENANT_ID"

// ExpiredAuthenticationToken - to test authentication retries
var ExpiredAuthenticationToken = os.Getenv("EXPIRED_BEARER_TOKEN")

// TestNamespace - A namespace for integration testing
var TestNamespace = fmt.Sprintf("gonamespace%d", timeSec)

// TestCollection - A collection for integration testing
var TestCollection = fmt.Sprintf("gocollection%d", timeSec)

// StubbyTestCollection - A collection for stubby testing
var StubbyTestCollection = "testcollection0"

// TestTimeOut - the timeout to be used for requests to the api
var TestTimeOut = time.Second * 5

// TenantCreationOn specifies whether tenants should be created on the fly for identity service /tenant CRUD testing
var TenantCreationOn = (os.Getenv("TENANT_CREATION") == "1")
