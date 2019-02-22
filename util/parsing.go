// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

// ParseResponse parses json-formatted http response and decodes it into pre-defined models
func ParseResponse(model interface{}, response *http.Response) error {
	if reflect.ValueOf(model).Kind() != reflect.Ptr {
		return errors.New("model parameter to ParseResponse() must be a pointer")
	}

	if response == nil {
		return errors.New("nil response provided")
	}
	if response.StatusCode == 204 {
		return nil
	}
	err := json.NewDecoder(response.Body).Decode(model)
	return err
}

// ParseURLParams parses a struct into url params based on its "key" tag
// It parses basic values and slices, and will parse structs recursively
func ParseURLParams(model interface{}) url.Values {
	values := url.Values{}
	if model == nil {
		return values
	}
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
