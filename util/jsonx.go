package util

import (
	"encoding/json"
	"reflect"
	"strings"
)

func MarshalByMethod(v interface{}, method string) ([]byte, error) {
	method = strings.ToLower(method)
	fields := flattenStructFields(v)
	for f := 0; f < len(fields); f++ {
		methods := strings.ToLower(fields[f].Tag.Get("methods"))
		if methods == "" {
			continue
		}
		parts := strings.Split(methods, ",")
		found := false
		for m := 0; m < len(parts); m++ {
			if parts[m] == method {
				found = true
				break
			}
		}
		if found {
			// Remove v from marshal
		}
	}
	return json.Marshal(v)
}

// flattenStructFields returns a flattened slice of StructFields, traversing into
// embedded struct fields and following struct pointers
func flattenStructFields(v interface{}) []reflect.StructField {
	fields := make([]reflect.StructField, 0)

	if v == nil {
		return fields
	}

	vvalue := reflect.ValueOf(v)
	vtype := vvalue.Type()

	if vtype.Kind() == reflect.Ptr && vtype.Elem().Kind() == reflect.Struct {
		// For pointer to a struct, follow the pointer
		vvalue = vvalue.Elem()
		vtype = vvalue.Elem().Type()
	} else if vvalue.Kind() != reflect.Struct {
		// Return empty slice for non-structs
		return fields
	}

	for i := 0; i < vtype.NumField(); i++ {
		structfield := vtype.Field(i)
		// Special case: embedded fields (Anonymous == true)
		if structfield.Anonymous {
			fieldval := vvalue.Field(i)
			if fieldval.Kind() == reflect.Struct {
				// For embedded structs, recurse over those fields
				fields = append(fields, flattenStructFields(fieldval.Interface())...)
			} else if fieldval.Kind() == reflect.Ptr && fieldval.Elem().Kind() == reflect.Struct {
				// For embedded pointers to structs, follow pointer and recurse over those fields
				fields = append(fields, flattenStructFields(fieldval.Elem().Interface())...)
			} else {
				// For embedded non-structs, simply add to slice
				fields = append(fields, structfield)
			}
		} else {
			// For non-embedded-type fields, add to slice
			fields = append(fields, structfield)
		}
	}

	return fields
}
