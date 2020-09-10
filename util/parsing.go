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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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

type OpenAPIParameterStyle string

const (
	// StyleForm corresponds to style: form which is the default for query parameters
	StyleForm OpenAPIParameterStyle = "form"
	// Add support for other styles here if services utilize them in the future
)

// ParseURLParams parses a struct into url params based on its "key" tag
// It parses basic values and slices, and will parse structs recursively
// see https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#style-examples
// for details about style and explode formatting
func ParseURLParams(model interface{}) url.Values {
	values := url.Values{}
	if model == nil {
		return nil
	}
	indirect := reflect.Indirect(reflect.ValueOf(model))
	// must be a struct, this should always be a *QueryParams model
	if indirect.Kind() != reflect.Struct {
		return values
	}
	return toURLValues("", reflect.ValueOf(model), "true", StyleForm)
}

func toURLValues(keyName string, value reflect.Value, explode string, style OpenAPIParameterStyle) url.Values {
	values := url.Values{}
	if style != StyleForm {
		// At the moment only style: "form" is supported, which is the default for query parameters
		return values
	}
	// If pointer or interface, follow the pointer
	if value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		if value.IsNil() {
			return values
		}
		value = value.Elem()
	}
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		// explode=true: repeat k=v pairs e.g. ?k1=v1&k1=v2&k1=v3&...
		// explode=false (default): comma separate values e.g. ?k1=v1,v2,v3,...
		// TODO: we are using comma-separated values by default although
		// in OpenAPI3.0.x explode=true for query parameters by default
		// use `explode:"true"` to repeat kv pairs for now
		vals := []string{}
		for i := 0; i < value.Len(); i++ {
			uvals := toURLValues(keyName, value.Index(i), "true", StyleForm)
			for uk, vslice := range uvals {
				// don't support a slice of maps
				if uk != keyName {
					continue
				}
				for _, v := range vslice {
					if explode == "true" {
						values.Add(uk, v)
					} else {
						vals = append(vals, v)
					}
				}
			}
		}
		if explode != "true" && len(vals) != 0 {
			values.Set(keyName, strings.Join(vals, ","))
		}
	case reflect.Struct:
		// explode=true (default): ignore keyName and add k=v pairs for fields e.g. ?k1=v1&k2=v2&k3=v3&...
		// explode=false: comma separate keys and values e.g. ?keyName=k1,v1,k2,v2,...
		keyvals := []string{}
		for f := 0; f < value.Type().NumField(); f++ {
			field := value.Type().Field(f)
			fval := value.FieldByName(field.Name)
			// If explode is not specified, default is true
			fx, _ := field.Tag.Lookup("explode")
			// If style is specified, use that otherwise default is "form"
			fs := StyleForm
			if v, ok := field.Tag.Lookup("form"); ok {
				fs = OpenAPIParameterStyle(v)
			}
			// If `key:` tag is specified, use it otherwise default to field name
			fkey := field.Name
			if v, ok := field.Tag.Lookup("key"); ok {
				fkey = v
			}
			uvals := toURLValues(fkey, fval, fx, fs)
			for uk, vslice := range uvals {
				for _, v := range vslice {
					if explode != "false" {
						values.Add(uk, v)
					} else {
						keyvals = append(keyvals, uk)
						keyvals = append(keyvals, v)
					}
				}
			}
		}
		if explode == "false" && len(keyvals) != 0 {
			values.Set(keyName, strings.Join(keyvals, ","))
		}
	case reflect.Map:
		// explode=true (default): ignore keyName and add k=v pairs for fields e.g. ?k1=v1&k2=v2&k3=v3&...
		// explode=false: comma separate keys and values e.g. ?keyName=k1,v1,k2,v2,...
		keyvals := []string{}
		for _, k := range value.MapKeys() {
			key := fmt.Sprintf("%v", k.Interface())
			if len(key) == 0 {
				continue
			}
			mv := value.MapIndex(k)
			uvals := toURLValues(key, mv, "true", StyleForm)
			for uk, vslice := range uvals {
				for _, v := range vslice {
					if explode != "false" {
						values.Add(uk, v)
					} else {
						keyvals = append(keyvals, uk)
						keyvals = append(keyvals, v)
					}
				}
			}
		}
		if explode == "false" && len(keyvals) != 0 {
			values.Set(keyName, strings.Join(keyvals, ","))
		}
	default:
		if val := fmt.Sprintf("%v", value.Interface()); len(val) > 0 {
			values.Set(keyName, val)
		}
	}
	return values
}
