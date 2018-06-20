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

func TestParseHTTPStatusCodeInResponseBadResponse(t *testing.T) {
	httpResp := &http.Response{
		StatusCode: 400,
		Status:     "400 Bad Request",
		Body:       ioutil.NopCloser(bytes.NewBufferString("")),
	}

	expectErrMsg := "Http Error: [400] 400 Bad Request "
	_, err := ParseHTTPStatusCodeInResponse(httpResp)
	httpError := err.(*HTTPError)

	if err == nil || httpError.Status != 400 || httpError.Message != "400 Bad Request" || err.Error() != expectErrMsg {
		t.Errorf("ParseHTTPStatusCodeInResponse expected to return an error for bad responses, got %v", err)
	}
}
