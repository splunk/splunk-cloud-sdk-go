// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package util

import (
	"bytes"
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
	expectErrMsg := "Http Error: [400] 400 Bad Request "
	if _, err := ParseHTTPStatusCodeInResponse(httpResp); err == nil || err.(*HTTPError).HTTPStatusCode != 400 || err.Error() != expectErrMsg {
		t.Errorf("ParseHTTPStatusCodeInResponse expected to return an error for bad responses, got %v", err)
	}
}

func TestParseHTTPStatusCodeInResponseEmptyBody(t *testing.T) {
	httpResp := &http.Response{
		StatusCode: 400,
		Status:     "400 Bad Request",
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
	}
	expectErrMsg := "Http Error: [400] 400 Bad Request "
	if _, err := ParseHTTPStatusCodeInResponse(httpResp); err == nil || err.(*HTTPError).HTTPStatusCode != 400 || err.Error() != expectErrMsg {
		t.Errorf("ParseHTTPStatusCodeInResponse expected to return an error with body, got %v", err)
	}
}

func TestParseHTTPStatusCodeInResponseBodyMsg(t *testing.T) {
	httpResp := &http.Response{
		StatusCode: 400,
		Status:     "400 Bad Request",
		Body:       ioutil.NopCloser(bytes.NewReader([]byte("unknown sid"))),
	}
	expectErrMsg := "Http Error: [400] 400 Bad Request unknown sid"
	if _, err := ParseHTTPStatusCodeInResponse(httpResp); err == nil || err.(*HTTPError).HTTPStatusCode != 400 || err.Error() != expectErrMsg {
		t.Errorf("ParseHTTPStatusCodeInResponse expected to return an error with body, got %v", err)
	}
}
