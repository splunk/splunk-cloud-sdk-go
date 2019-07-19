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
	"net/url"
	"testing"

	test "github.com/splunk/splunk-cloud-sdk-go/test/scloud"

	"github.com/stretchr/testify/assert"
)

func TestKVStoreCommand(t *testing.T) {
	const collection = "myCollection"
	const index = "myIndex"
	const key = "myKey"

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
			name:        "Create index",
			args:        []string{"create-index", collection, "-index", index},
			xResult:     nil,
			xMethod:     http.MethodPost,
			xURLPattern: fmt.Sprintf("%s://%s/%s/kvstore/v1beta1/collections/%s/indexes", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, collection),
			xBody:       "{\"fields\":null,\"name\":\"myIndex\"}",
		}, {
			name:        "Delete record",
			args:        []string{"delete-record", collection, key},
			xResult:     nil,
			xMethod:     http.MethodDelete,
			xURLPattern: fmt.Sprintf("%s://%s/%s/kvstore/v1beta1/collections/%s/records/%s", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, collection, key),
		}, {
			name:        "Delete records",
			args:        []string{"delete-records", collection, "-query", "{\"potato\": 1}"},
			xResult:     nil,
			xMethod:     http.MethodDelete,
			xURLPattern: fmt.Sprintf("%s://%s/%s/kvstore/v1beta1/collections/%s/query\\?query=", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, collection),
			xQueryArgs:  url.Values{"query": []string{"{\"potato\": 1}"}},
		}, {
			name:        "Get record",
			args:        []string{"get-record", collection, key},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/%s/kvstore/v1beta1/collections/%s/records/%s", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, collection, key),
		}, {
			name:        "Get health status",
			args:        []string{"get-health-status"},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/%s/kvstore/v1beta1/ping", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant),
		}, {
			name:        "Get JSON spec",
			args:        []string{"get-spec-json"},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/system/kvstore/specs/v1beta1/openapi.json", test.UnitTestScheme, test.UnitTestHost),
		}, {
			name:        "Get YAML spec",
			args:        []string{"get-spec-yaml"},
			xResult:     "", // GetSpecYaml() returns ("", error) instead of (nil, error)
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/system/kvstore/specs/v1beta1/openapi.ya?ml", test.UnitTestScheme, test.UnitTestHost),
		}, {
			//	name:        "Insert batch records", // TODO: probably won't be unit testable since it reads from stdin
			//	args:        []string{"insert-batch-records", collection},
			//	xResult:     nil,
			//	xMethod:     http.MethodPost,
			//	xURLPattern: fmt.Sprintf("%s://%s/%s/kvstore/v1beta1/collections/%s", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, collection),
			//  xBody: "",
			//}, {
			name:        "Insert record",
			args:        []string{"insert-record", collection, "{\"some\": \"data\"}"},
			xResult:     nil,
			xMethod:     http.MethodPost,
			xURLPattern: fmt.Sprintf("%s://%s/%s/kvstore/v1beta1/collections/%s", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, collection),
			xBody:       "{\"some\":\"data\"}",
		}, {
			name:        "List indexes",
			args:        []string{"list-indexes", collection},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/%s/kvstore/v1beta1/collections/%s/indexes", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, collection),
		}, {
			name:        "List records",
			args:        []string{"list-records", collection},
			xResult:     nil,
			xMethod:     http.MethodGet,
			xURLPattern: fmt.Sprintf("%s://%s/%s/kvstore/v1beta1/collections/%s", test.UnitTestScheme, test.UnitTestHost, test.UnitTestTenant, collection),
		},
	}

	cmd := newKVStoreCommand(test.GetUnitTestClient())

	// TODO: Commenting out for now until help() consumers are refactored so fmt.println() doesn't happen in that method
	//res, helpText := Dispatch(kv, []string{"help"})
	//assert.Nil(t, res)
	//assert.Equal(t, helpText, help("kvstore.txt"))

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
