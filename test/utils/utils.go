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
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/splunk/splunk-cloud-sdk-go/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/services"
)

// GetFilename uses reflection to get current filename
func GetFilename() string {
	_, filename, _, _ := runtime.Caller(0)
	return filename
}

func init() {
	envPath := filepath.Join(filepath.Dir(GetFilename()), "..", "..", ".env")

	err := godotenv.Load(envPath)

	if err != nil {
		log.Println("Error loading .env from ", envPath)
	}

	TestAuthenticationToken = os.Getenv("BEARER_TOKEN")
	TestSplunkCloudHost = os.Getenv("SPLUNK_CLOUD_HOST_TENANT_SCOPED")

	TestTenant = os.Getenv("TEST_TENANT_SCOPED")

	TestRegion = os.Getenv("REGION")

	TestUsername = os.Getenv("BACKEND_CLIENT_ID_TENANT_SCOPED")
	ExpiredAuthenticationToken = os.Getenv("EXPIRED_BEARER_TOKEN")
	PkceClientID = os.Getenv("REFRESH_TOKEN_CLIENT_ID")
	Username = os.Getenv("TEST_USERNAME")
	Password = os.Getenv("TEST_PASSWORD")
	IdpHost = os.Getenv("IDP_HOST_TENANT_SCOPED")
}

// RunSuffix - run instance identifier suffix based on timestamp
var RunSuffix = time.Now().Unix()

// TestSplunkCloudHost - the url for the test api to be used
var TestSplunkCloudHost string

// TestAuthenticationToken - the authentication that gives permission to make requests against the api
var TestAuthenticationToken string

// TestTenant - the tenant to be used for the API
var TestTenant string

// TestUsername - the user running tests on behalf of
var TestUsername string

// TestInvalidTestTenant - the invalid tenant ID that denies permission to make requests against the api
var TestInvalidTestTenant = "INVALID_TENANT_ID"

// ExpiredAuthenticationToken - to test authentication retries
var ExpiredAuthenticationToken string

// TestModule - A namespace for integration testing
var TestModule = fmt.Sprintf("gomod%d", RunSuffix)

// TestModule 2- A namespace for integration testing
var TestModule2 = fmt.Sprintf("gomod2%d", RunSuffix)

// TestCollection - A collection for integration testing
var TestCollection = fmt.Sprintf("gocollection%d", RunSuffix)

// StubbyTestCollection - A collection for stubby testing
var StubbyTestCollection = "testcollection0"

// TestTimeOut - the timeout to be used for requests to the api
var TestTimeOut = time.Second * 30

// TestTimeOut - the timeout to be used for requests to the api in search tests
var LongTestTimeout = time.Second * 600

//Region which correlates to the tenant, used to formulate tenant scoped hostnames
var TestRegion string

// PKCE
const NativeAppRedirectURI = "https://login.splunkbeta.com"

var PkceClientID string
var Username string
var Password string
var IdpHost string

// Get an client without the testing interface
func MakeSdkClient(tr idp.TokenRetriever, tenant string) (*sdk.Client, error) {
	return sdk.NewClient(&services.Config{
		TokenRetriever: tr,
		Host:           TestSplunkCloudHost,
		Tenant:         tenant,
		Timeout:        TestTimeOut,
	})
}
