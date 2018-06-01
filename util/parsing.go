package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

// ParseResponse parses json-formatted http response and decodes it into pre-defined models
func ParseResponse(model interface{}, response *http.Response) error {
	if response == nil {
		return fmt.Errorf("nil response provided")
	}
	return json.NewDecoder(response.Body).Decode(model)
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
