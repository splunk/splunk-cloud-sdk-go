package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type taggedField struct {
	*reflect.Value
	*reflect.StructField
}

// MethodMarshaler is the interface implemented by types that
// can marshal themselves into valid JSON according to the
// input http method
type MethodMarshaler interface {
	MarshalJSONByMethod(method string) ([]byte, error)
}

// MarshalByMethod marshals any json tagged struct fields matching the method being specified.
//
// If the `methods:` tag is specified for the field, then the field is marshaled if
// the input method is present in the comma-separated list within the tag.
//
// If no `methods:` tag is present then it is presumed that the field is valid for all
// methods, so the field is marshaled.
//
// Examples:
//
func MarshalByMethod(v interface{}, method string) ([]byte, error) {
	fields := getFieldsByTag(v, "json")
	method = strings.ToUpper(method)
	methodsFieldFound := false
	// This will be the map containing json fields that we wish to marshal
	toMarshal := make(map[string]interface{})
	for f := 0; f < len(fields); f++ {
		field := fields[f]
		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" {
			// json tags matching "-" are skipped, see: https://golang.org/pkg/encoding/json/#Marshal
			continue
		}
		methods, methodsExist := field.Tag.Lookup("methods")
		if methodsExist {
			methodsFieldFound = true
			// If method tag exists, check if the method being marshaled is present in the tag
			parts := strings.Split(strings.ToUpper(methods), ",")
			found := false
			// Since m ~ 1-4 linear search will do just fine
			for m := 0; m < len(parts); m++ {
				if parts[m] == method {
					found = true
					break
				}
			}
			if !found {
				// Skip this field if methods tag exists but doesn't match the method being marshaled
				continue
			}
		}
		// Get key for toMarshal map from tag of the form `json:"<mykey>"` or `json:"<mykey>,..."`
		name, opts := parseTag(jsonTag)
		if name == "" {
			return nil, fmt.Errorf("util: jsonx.MarshalByMethod() blank json names are currently unsupported, blank name found for field: %+v", field.StructField)
		}
		if opts.Contains("omitempty") && isEmptyValue(*field.Value) {
			// Omit empty values as json.Marshal does
			continue
		}
		// Add field to our map to be marshaled
		toMarshal[name] = field.Interface()
	}
	if !methodsFieldFound {
		return nil, fmt.Errorf("util: jsonx.MarshalByMethod() should only be used on structs with fields containing at least one `methods:` tag - use json.Marshal() if no such fields exist")
	}
	return json.Marshal(toMarshal)
}

// getFieldsByTag returns a flattened slice of StructFields, traversing into
// embedded struct fields and following struct pointers for fields containing
// the given tag
func getFieldsByTag(v interface{}, tag string) []*taggedField {
	fields := make([]*taggedField, 0)
	vvalue := reflect.ValueOf(v)
	if vvalue.Type().Kind() == reflect.Ptr {
		// If v is a pointer, follow the pointer
		vvalue = reflect.ValueOf(vvalue.Elem().Interface())
	}
	if vvalue.Type().Kind() != reflect.Struct {
		// Only structs are supported since tags only exist on struct fields
		return fields
	}

	for i := 0; i < vvalue.NumField(); i++ {
		fieldval := vvalue.Field(i)
		structfield := vvalue.Type().Field(i)
		//fmt.Printf(" structfield: %+v CanSet(): %v\n\n", structfield, fieldval.CanSet())
		_, hasTag := structfield.Tag.Lookup(tag)
		// Special case: embedded fields (Anonymous == true)
		if structfield.Anonymous {
			if fieldval.Kind() == reflect.Struct {
				// For embedded structs, recurse over those fields
				fields = append(fields, getFieldsByTag(fieldval.Interface(), tag)...)
			} else if fieldval.Kind() == reflect.Ptr && fieldval.Elem().Kind() == reflect.Struct {
				// For embedded pointers to structs, follow pointer and recurse over those fields
				fields = append(fields, getFieldsByTag(fieldval.Interface(), tag)...)
			} else if hasTag {
				// For embedded non-structs fields with tag, simply add to slice
				fields = append(fields, &taggedField{Value: &fieldval, StructField: &structfield})
			}
		} else if hasTag {
			// For non-embedded-type fields with tag, add to slice
			fields = append(fields, &taggedField{Value: &fieldval, StructField: &structfield})
		}
	}

	return fields
}

// From: https://golang.org/src/encoding/json/encode.go
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

// From: https://golang.org/src/encoding/json/tags.go

// tagOptions is the string following a comma in a struct field's "json"
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, tagOptions("")
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var next string
		i := strings.Index(s, ",")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}
		if s == optionName {
			return true
		}
		s = next
	}
	return false
}
