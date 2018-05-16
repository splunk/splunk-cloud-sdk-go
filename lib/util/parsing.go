package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// ParseResponse parses the http response and unmarshals it into json
func ParseResponse(model interface{}, response *http.Response, err error) error {
	if err == nil {
		defer response.Body.Close()
		b := new(bytes.Buffer)
		b.ReadFrom(response.Body)
		err = json.NewDecoder(b).Decode(model)
	}
	return err
}

// ParseError checks for error and closes the response body
// It can be used when we don't care about the response, but do want to close the response body
func ParseError(response *http.Response, err error) error {
	if err == nil {
		defer response.Body.Close()
	}
	return err
}
