package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

// ParseResponse parses the http response and unmarshals it into json
func ParseResponse(model interface{}, response *http.Response, err error) error {

	// The request body could not be serialized or the http request is invalid
	if response == nil {
		return err
	}
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

// ParseURLParams parses a struct into url params based on its "key" tag
// It parses basic values and slices, and will parse structs recursively
func ParseURLParams(model interface{}) url.Values {
	values := url.Values{}
	typ := reflect.TypeOf(model)
	indirect := reflect.Indirect(reflect.ValueOf(model))

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		indirectField := indirect.FieldByName(field.Name)

		if keyName, ok := field.Tag.Lookup("key"); ok {
			switch indirectField.Kind() {
			case reflect.String, reflect.Int, reflect.Uint, reflect.Float32, reflect.Float64, reflect.Bool:
				if val := fmt.Sprintf("%v", indirectField.Interface()); len(val) > 0 {
					values.Set(keyName, val)
				}
			case reflect.Slice:
				for i := 0; i < indirectField.Len(); i++ {
					if val := fmt.Sprintf("%v", indirectField.Index(i)); len(val) > 0 {
						values.Add(keyName, val)
					}
				}
				//TODO should log or warn about incorrect params
			}
		} else {
			if indirectField.Kind() == reflect.Struct {
				for k, vList := range ParseURLParams(indirectField.Interface()) {
					for _, v := range vList {
						values.Add(k, v)
					}
				}
			}
		}
	}

	return values
}
