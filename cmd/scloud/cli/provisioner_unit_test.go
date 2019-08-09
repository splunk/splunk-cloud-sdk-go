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

	test "github.com/splunk/splunk-cloud-sdk-go/test/scloud"

	"github.com/stretchr/testify/assert"
)

func TestProvisionerCommand(t *testing.T) {
	const jobID = "jobID"
	const tenantName = "tenantname"

	tests := []struct {
		name string
		args []string
		// Expected values
		xResult     interface{}
		xMethod     string
		xURLPattern string
		xQueryArgs  interface{}
		xBody       interface{}
	}{
		{
			name:        "Create provision job",
			args:        []string{"create-provision-job", "-tenant", tenantName, "-app", "app1", "-app", "app2"},
			xResult:     nil,
			xMethod:     http.MethodPost,
			xURLPattern: fmt.Sprintf("%s://%s/system/provisioner/v1beta1/jobs/tenants/provision", test.UnitTestScheme, test.UnitTestHost),
			xBody:       `{"apps":["app1","app2"],"tenant":"` + tenantName + `"}`,
		}, {
			name:        "Get provision job",
			args:        []string{"get-provision-job", jobID},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/system/provisioner/v1beta1/jobs/tenants/provision/%s", test.UnitTestScheme, test.UnitTestHost, jobID),
		}, {
			name:        "Get spec JSON",
			args:        []string{"get-spec-json"},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/system/provisioner/specs/v1beta1/openapi.json", test.UnitTestScheme, test.UnitTestHost),
		}, {
			name:        "Get spec YAML",
			args:        []string{"get-spec-yaml"},
			xResult:     "", // GetSpecYaml() returns ("", error) instead of (nil, error)
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/system/provisioner/specs/v1beta1/openapi.yaml", test.UnitTestScheme, test.UnitTestHost),
		}, {
			name:        "Get tenant",
			args:        []string{"get-tenant", tenantName},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/system/provisioner/v1beta1/tenants/%s", test.UnitTestScheme, test.UnitTestHost, tenantName),
		}, {
			name:        "List provision jobs",
			args:        []string{"list-provision-jobs"},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/system/provisioner/v1beta1/jobs/tenants/provision", test.UnitTestScheme, test.UnitTestHost),
		}, {
			name:        "List tenants",
			args:        []string{"list-tenants"},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/system/provisioner/v1beta1/tenants", test.UnitTestScheme, test.UnitTestHost),
		},
	}

	cmd := newProvisionerCommand(test.GetUnitTestClient())

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
