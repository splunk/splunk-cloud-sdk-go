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

package main

import (
	"fmt"
	"net/http"
	"testing"

	test "github.com/splunk/splunk-cloud-sdk-go/v2/test/scloud"

	"net/url"

	"github.com/stretchr/testify/assert"
)

func TestSearchCommand(t *testing.T) {
	const jobID = "12e12ce53c944afd87a0b969a6e37d83_1561581065.169133"
	const version = "v2beta1"

	tests := []struct {
		name string
		args []string
		// Expected values
		xResult     interface{}
		xMethod     string
		xURLPattern string
		xBody       interface{}
		xQueryArgs  interface{}
	}{
		{
			name:        "Get JSON spec",
			args:        []string{"get-spec-json"},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/system/search/specs/%s/openapi.json", test.UnitTestScheme, test.UnitTestHost, version),
		}, {
			name:        "Get YAML spec",
			args:        []string{"get-spec-yaml"},
			xResult:     "", // GetSpecYaml() returns ("", error) instead of (nil, error)
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/system/search/specs/%s/openapi.yaml", test.UnitTestScheme, test.UnitTestHost, version),
		}, {
			name:        "Create job",
			args:        []string{"| from index:main | head 5", "-earliest", "-12h@h", "-latest", "now"},
			xResult:     nil,
			xMethod:     http.MethodPost,
			xURLPattern: fmt.Sprintf("%s://%s/%s/search/%s/jobs", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, version),
			xBody:       "{\"query\":\"| from index:main | head 5\",\"queryParameters\":{\"earliest\":\"-12h@h\",\"latest\":\"now\"}}",
		},
		{
			name:        "List jobs with a status and count",
			args:        []string{"list-jobs", "-count", "1", "-status", "running"},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/%s/search/%s/jobs\\?count=\\d&status=running", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, version),
			xQueryArgs:  url.Values{"count": []string{"0"}, "status": []string{"running"}},
		},
		{
			name:        "List jobs - default",
			args:        []string{"list-jobs"},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/%s/search/%s/jobs\\?count=\\d&status=done", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, version),
		},
		{
			name:        "List Results with a count, offset and field",
			args:        []string{"list-results", jobID, "-count", "5", "-fields", "source", "-offset", "1"},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/%s/search/%s/jobs/%s/results\\?count=\\d&field=([a-z]+)&offset=\\d", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, version, jobID),
		},
		{
			name:        "Update job - cancel job",
			args:        []string{"cancel", jobID},
			xResult:     nil,
			xMethod:     http.MethodPatch,
			xURLPattern: fmt.Sprintf("%s://%s/%s/search/%s/jobs/%s", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, version, jobID),
			xBody:       "{\"status\":\"canceled\"}",
		},
		{
			name:        "Update job - finalize job",
			args:        []string{"finalize", jobID},
			xResult:     nil,
			xMethod:     http.MethodPatch,
			xURLPattern: fmt.Sprintf("%s://%s/%s/search/%s/jobs/%s", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, version, jobID),
			xBody:       "{\"status\":\"finalized\"}",
		},
		{
			name:        "Wait job to reach terminal state - done, failed",
			args:        []string{"wait", jobID},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/%s/search/%s/jobs/%s", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, version, jobID),
		},
	}

	cmd := newSearchCommand(test.GetUnitTestClient())

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := cmd.Dispatch(tc.args)

			if tc.xResult == nil {
				assert.Nil(t, result)
			} else {
				assert.Equal(t, tc.xResult, result)
			}
			assert.NotNil(t, err)

			req := test.ErrorToRequest(t, err)
			assert.Equal(t, tc.xMethod, req.Method)
			assert.Regexp(t, tc.xURLPattern, req.URL.String())
			if tc.xQueryArgs != nil {
				assert.Equal(t, tc.xQueryArgs, req.URL.Query())
			}

			// Avoid the nil-pointer dereference
			if req.Body == nil {
				assert.Nil(t, tc.xBody)
			} else {
				body := test.BodyToString(t, req.Body)
				assert.Equal(t, tc.xBody, body)
			}
		})
	}
}
