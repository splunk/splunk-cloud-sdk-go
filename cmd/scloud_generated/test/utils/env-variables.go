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

package test_engine

import "os"

// TestSplunkCloudHost - the url for the test api to be used
var TestSplunkCloudHost = os.Getenv("SPLUNK_CLOUD_HOST")

//Invalid Host Url
var InvalidHostUrl = "https://InvalidHostUrl:443"

// TestTenant - the tenant to be used for the API
var TestTenant = os.Getenv("TENANT_ID")

//Username to be used in scloud login flow
var Username = os.Getenv("TEST_USERNAME")

var Password = os.Getenv("TEST_PASSWORD")

//Auth url to provide token for request validation
var IdpHost = os.Getenv("IDP_HOST")

//Invalid Auth Url
var InvalidAuthUrl = "https://InvalidAuthUrl:443"

//Valid test env1
var Env1 = os.Getenv("TEST_ENVIRONMENT_1")

//Valid test env2
var Env2 = os.Getenv("TEST_ENVIRONMENT_2")
