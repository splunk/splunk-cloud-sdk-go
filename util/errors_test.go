package util

import (
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
	}
	expectErrMsg := "Http Error: [400] 400 Bad Request"
	if _, err := ParseHTTPStatusCodeInResponse(httpResp); err == nil || err.(*HTTPError).Status != 400 || err.Error() != expectErrMsg {
		t.Errorf("ParseHTTPStatusCodeInResponse expected to return an error for bad responses, got %v", err)
	}
}
