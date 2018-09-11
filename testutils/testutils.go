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
var TestURLProtocol = "https"

// TestSplunkCloudHost - the url for the test api to be used
var TestSplunkCloudHost ="api.playground.splunkbeta.com"

// TestAuthenticationToken - the authentication that gives permission to make requests against the api
var TestAuthenticationToken = "eyJraWQiOiJIR1RMbXJGUWNsSGVTRUZzdTIzQ1k4cTZ3S3pjR2JwUGtvT014R2hQVVBVIiwiYWxnIjoiUlMyNTYifQ.eyJ2ZXIiOjEsImp0aSI6IkFULmhCcjRiekVYUGlsb2F2MHQ3dTJWU3RZeGNtQ0RwRTZuR2hQT29jWjN6ZVkiLCJpc3MiOiJodHRwczovL3NwbHVuay1jaWFtLm9rdGEuY29tL29hdXRoMi9kZWZhdWx0IiwiYXVkIjoiYXBpOi8vZGVmYXVsdCIsImlhdCI6MTUzNjYwMTc3OSwiZXhwIjoxNTM2NjQ0OTc5LCJjaWQiOiIwb2FwYmcyem1MYW1wV2daNDJwNiIsInVpZCI6IjAwdTEwa20yMWhiT3BUbzdTMnA3Iiwic2NwIjpbInByb2ZpbGUiLCJlbWFpbCIsIm9wZW5pZCJdLCJzdWIiOiJ0ZXN0MUBzcGx1bmsuY29tIn0.XzITtjDgws7AFzYv39eoL_5QmC5LR_U0rTL1egupWSq_crDGt3JqRk-kyGe8YPkIJdTsQ97bdVXR0mEsxxixdHWCw6IRtBPPIMENkr6Qc6ZtGK7ESaQPPFEUuyU9fMdfJ2Gh-Z4YDhCo758yptwFItVdOJwXrA1TU6R4XHvZKXaoVNjUhKNSXtkEIwXHIaGeYk_d88q-L_XYD7OIZZdlVIgxCkAMLaWwqHX_rlBbvoN4aNQl7BrtZ0ofS41isNbPxQSFJBwD27f6ZTuQTGjhfBeTyH3146Qt5TlXlfkJ6WigeV3VnzmWB9QNqOZN_8LriB06INYw6gmtqHYah7wQKg"

// TestInvalidAuthenticationToken - the invalid access token that denies permission to make requests against the api
var TestInvalidAuthenticationToken = "INVALID_TOKEN"

// TestTenantID - the tenant to be used for the API
var TestTenantID = "ljiang6"

// TestInvalidTestTenantID - the invalid tenant ID that denies permission to make requests against the api
var TestInvalidTestTenantID = "INVALID_TEST_TENANT_ID"

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


