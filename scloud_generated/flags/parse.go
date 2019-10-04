package flags

import (
	"encoding/json"
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
	if !deref.CanSet() {
		return fmt.Errorf("flags.ParseFlag: value must be settable")
	}
	kind := deref.Kind()
	outtype := deref.Type().String()
	strval := flag.Value.String() // the actual string supplied by the user
	// Determine the type of out and inject a value into it
	switch kind {
	case reflect.String:
		deref.SetString(strval)
		return nil
	default:
		err := json.Unmarshal([]byte(strval), out)
		if err != nil {
			return fmt.Errorf("flags.ParseFlag: failure to unmarshal to type %s err: %s", outtype, err)
		}
		return nil
	}
}
