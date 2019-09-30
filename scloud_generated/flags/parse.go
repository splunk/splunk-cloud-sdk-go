package flags

import (
	"fmt"
	"reflect"

	"github.com/spf13/pflag"
)

func ParseFlag(flag *pflag.Flag, out interface{}) error {
	if flag == nil {
		return fmt.Errorf("flags.ParseFlag: flag was nil")
	}
	//typ := reflect.TypeOf(out)
	val := reflect.ValueOf(out)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("flags.ParseFlag: must accept a pointer to value")
	}
	// Follow the pointer
	deref := val.Elem()
	outtype := deref.Type().String()
	// Determine the type of out and inject a value into it
	if outtype == "string" {
		sout, ok := out.(*string)
		if !ok {
			return fmt.Errorf("flags.ParseFlag: unexpected type assertion failure for type: %s", outtype)
		}
		*sout = flag.Value.String()
		return nil
	}
	return fmt.Errorf("flags.ParseFlag: only string is accepted at the moment, found type: %s", outtype)
}
