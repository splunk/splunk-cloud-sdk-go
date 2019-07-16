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

package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHTTPStatusCodeInResponseOKResponse(t *testing.T) {
	httpResp := &http.Response{
		StatusCode: 201,
	}
	if _, err := ParseHTTPStatusCodeInResponse(httpResp); err != nil {
		t.Errorf("ParseHTTPStatusCodeInResponse expected to not return error for good responses, got %v", err)
	}
}

func TestParseHTTPStatusCodeInResponseNilResponse(t *testing.T) {
	if _, err := ParseHTTPStatusCodeInResponse(nil); err != nil {
		t.Errorf("ParseHTTPStatusCodeInResponse expected to not return error for good responses, got %v", err)
	}
}

func TestParseHTTPStatusCodeInResponseBadResponseNilBody(t *testing.T) {
	httpResp := &http.Response{
		StatusCode: 400,
		Status:     "400 Bad Request",
		Body:       nil,
	}
	expectErrMsg := `{"HTTPStatusCode":400,"HTTPStatus":"400 Bad Request"}`
	_, err := ParseHTTPStatusCodeInResponse(httpResp)
	assert.Equal(t, 400, err.(*HTTPError).HTTPStatusCode)
	assert.Equal(t, "400 Bad Request", err.(*HTTPError).HTTPStatus)
	assert.Equal(t, err.Error(), expectErrMsg)
}

func TestParseHTTPStatusCodeInResponseEmptyBody(t *testing.T) {
	httpResp := &http.Response{
		StatusCode: 400,
		Status:     "400 Bad Request",
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{}`))),
	}
	expectErrMsg := `{"HTTPStatusCode":400,"HTTPStatus":"400 Bad Request"}`
	_, err := ParseHTTPStatusCodeInResponse(httpResp)
	assert.Equal(t, 400, err.(*HTTPError).HTTPStatusCode)
	assert.Equal(t, "400 Bad Request", err.(*HTTPError).HTTPStatus)
	assert.Equal(t, err.Error(), expectErrMsg)
}

func TestParseHTTPStatusCodeInResponseBodyMsg(t *testing.T) {
	httpResp := &http.Response{
		StatusCode: 400,
		Status:     "400 Bad Request",
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"code": "1017","message": "Validation Failed"}`))),
	}
	expectErrMsg := `{"HTTPStatusCode":400,"HTTPStatus":"400 Bad Request","message":"Validation Failed","code":"1017"}`
	_, err := ParseHTTPStatusCodeInResponse(httpResp)
	assert.Equal(t, 400, err.(*HTTPError).HTTPStatusCode)
	assert.Equal(t, "400 Bad Request", err.(*HTTPError).HTTPStatus)
	assert.Equal(t, "1017", err.(*HTTPError).Code)
	assert.Equal(t, "Validation Failed", err.(*HTTPError).Message)
	assert.Equal(t, err.Error(), expectErrMsg)
}
