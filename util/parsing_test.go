package util

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestParseResponseError(t *testing.T) {
	testData := []byte("testbool=true&testfloat32=0.999&testfloat64=0.555&testint=1")
	type TestModel struct {
		TestID string `json:"TestID"`
	}

	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(testData)),
	}
	var testModel TestModel
	err := ParseResponse(&testModel, httpResp, nil)
	if err == nil {
		t.Errorf("ParseResponse expected to raise an error, got %v", err)
	}
}

func TestParseResponse(t *testing.T) {
	testData := []byte(`{"TestID":"1"}`)
	type TestModel struct {
		TestID string `json:"TestID"`
	}
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(testData)),
	}
	var testModel TestModel
	err := ParseResponse(&testModel, httpResp, nil)
	if err != nil {
		t.Errorf("ParseResponse expected to return an error, got %v", err)
	}
}

func TestParseErrorNoError(t *testing.T) {
	testData := []byte(`{"TestID":"1"}`)
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(testData)),
	}
	if err := ParseError(httpResp, nil); err != nil {
		t.Errorf("ParseError expected to not return an error, got %v", err)
	}
}

func TestParseErrorReturnError(t *testing.T) {
	testData := []byte(`{"TestID":"1"}`)
	httpResp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(testData)),
	}
	err := errors.New("TestParseErrorReturnError should return this error")
	if err = ParseError(httpResp, err); err == nil {
		t.Errorf("ParseError expected to return an error, got %v", err)
	}
}
