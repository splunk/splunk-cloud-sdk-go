// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package util

import (
	"bytes"
	"github.com/splunk/splunk-cloud-sdk-go/.vendor-new/github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
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
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
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
