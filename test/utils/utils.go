/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
)

// TimeSec denotes the current timestamp
var TimeSec = time.Now().Unix()

// TestSplunkCloudHost - the url for the test api to be used
var TestSplunkCloudHost = os.Getenv("SPLUNK_CLOUD_HOST")

// TestAuthenticationToken - the authentication that gives permission to make requests against the api
var TestAuthenticationToken = os.Getenv("BEARER_TOKEN")

// TestTenant - the tenant to be used for the API
var TestTenant = os.Getenv("TENANT_ID")

// TestTenant - the tenant to be used for the API
var TestMLTenant = os.Getenv("ML_TENANT_ID")

// TestTenant - the tenant to be used for the API
var TestProvisionerTenant = os.Getenv("TENANT_PROVISIONER_ID")

// TestUsername - the user running tests on behalf of
var TestUsername = os.Getenv("BACKEND_CLIENT_ID")

// TestInvalidTestTenant - the invalid tenant ID that denies permission to make requests against the api
var TestInvalidTestTenant = "INVALID_TENANT_ID"

// ExpiredAuthenticationToken - to test authentication retries
var ExpiredAuthenticationToken = os.Getenv("EXPIRED_BEARER_TOKEN")

// TestModule - A namespace for integration testing
var TestModule = fmt.Sprintf("gomod%d", TimeSec)

// TestModule 2- A namespace for integration testing
var TestModule2 = fmt.Sprintf("gomod2%d", TimeSec)

// TestCollection - A collection for integration testing
var TestCollection = fmt.Sprintf("gocollection%d", TimeSec)

// StubbyTestCollection - A collection for stubby testing
var StubbyTestCollection = "testcollection0"

// TestTimeOut - the timeout to be used for requests to the api
var TestTimeOut = time.Second * 5

// TenantCreationOn specifies whether tenants should be created on the fly for identity service /tenant CRUD testing
var TenantCreationOn = os.Getenv("TENANT_CREATION") == "1"

// PKCE
const NativeAppRedirectURI = "https://login.splunkbeta.com"

var PkceClientID = os.Getenv("REFRESH_TOKEN_CLIENT_ID")
var Username = os.Getenv("TEST_USERNAME")
var Password = os.Getenv("TEST_PASSWORD")
var IdpHost = os.Getenv("IDP_HOST")

// Get an client without the testing interface
func MakeSdkClient(tr idp.TokenRetriever, tenant string) (*sdk.Client, error) {
	return sdk.NewClient(&services.Config{
		TokenRetriever: tr,
		Host:           TestSplunkCloudHost,
		Tenant:         tenant,
		Timeout:        TestTimeOut,
	})
}
