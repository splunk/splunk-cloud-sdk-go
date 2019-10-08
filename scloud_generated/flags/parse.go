package flags

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/spf13/pflag"
)

func ParseFlag(flags *pflag.FlagSet, name string, out interface{}) error {
	flag := flags.Lookup(name)
	if flag == nil {
		return fmt.Errorf("flags.ParseFlag: no flag defined, flag was nil")
	}
	if !flag.Changed {
		// If no value set by user, do not parse it here (use Go's default)
		// TODO: support defaults if specified in the flag
		return nil
	}
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
	// Determine the kind of out and inject a value into it
	switch kind {
	case reflect.String:
		strval, err := flags.GetString(name)
		if err != nil {
			return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetString("%s") err: %s`, name, err)
		}
		deref.SetString(strval)
		return nil
	case reflect.Bool:
		bval, err := flags.GetBool(name)
		if err != nil {
			return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetBool("%s") err: %s`, name, err)
		}
		deref.SetBool(bval)
		return nil
	case reflect.Int:
		ival, err := flags.GetInt(name)
		if err != nil {
			return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetInt("%s") err: %s`, name, err)
		}
		deref.SetInt(int64(ival))
		return nil
	case reflect.Int32:
		ival, err := flags.GetInt32(name)
		if err != nil {
			return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetInt32("%s") err: %s`, name, err)
		}
		deref.SetInt(int64(ival))
		return nil
	case reflect.Int64:
		ival, err := flags.GetInt64(name)
		if err != nil {
			return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetInt64("%s") err: %s`, name, err)
		}
		deref.SetInt(ival)
		return nil
	case reflect.Float32:
		flval, err := flags.GetFloat32(name)
		if err != nil {
			return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetFloat32("%s") err: %s`, name, err)
		}
		deref.SetFloat(float64(flval))
		return nil
	case reflect.Float64:
		flval, err := flags.GetFloat64(name)
		if err != nil {
			return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetFloat64("%s") err: %s`, name, err)
		}
		deref.SetFloat(flval)
		return nil
	case reflect.Slice:
		// Switch on the kind of slice
		switch deref.Type().Elem().Kind() {
		case reflect.String:
			slval, err := flags.GetStringSlice(name)
			if err != nil {
				return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetStringSlice("%s") err: %s`, name, err)
			}
			deref.Set(reflect.ValueOf(slval))
			return nil
		case reflect.Int:
			slval, err := flags.GetIntSlice(name)
			if err != nil {
				return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetIntSlice("%s") err: %s`, name, err)
			}
			deref.Set(reflect.ValueOf(slval))
			return nil
		case reflect.Bool:
			slval, err := flags.GetBoolSlice(name)
			if err != nil {
				return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetBoolSlice("%s") err: %s`, name, err)
			}
			deref.Set(reflect.ValueOf(slval))
			return nil
		}
	}
	// If complex kind then attempt to unmarshal as json string
	strval, err := flags.GetString(name)
	if err != nil {
		return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetString("%s") err: %s`, name, err)
	}
	err = json.Unmarshal([]byte(strval), out)
	if err != nil {
		return fmt.Errorf("flags.ParseFlag: failure to unmarshal to type %s err: %s", outtype, err)
	}
	return nil
}
