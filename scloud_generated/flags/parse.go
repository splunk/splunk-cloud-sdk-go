package flags

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

func ParseFlag(flags *pflag.FlagSet, name string, out interface{}) error {
	flag := flags.Lookup(name)
	if flag == nil {
		return fmt.Errorf("flags.ParseFlag: no flag defined, flag was nil")
	}
	val := reflect.ValueOf(out)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("flags.ParseFlag: must accept a pointer to value")
	}
	// Follow the pointer
	deref := val.Elem()
	if !flag.Changed && deref.Kind() == reflect.Ptr {
		// If no value set by user set to nil for optional flags (pointers)
		if !deref.CanSet() {
			return fmt.Errorf("flags.ParseFlag: value must be settable")
		}
		deref.Set(reflect.Zero(deref.Type()))
		return nil
	}
	if deref.Kind() == reflect.Ptr {
		// Optional flags will be a pointer to a pointer, deref again
		deref = deref.Elem()
	}
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
		sbval, err := flags.GetString(name)
		if err != nil {
			return fmt.Errorf(`flags.ParseFlag: error retrieving flag flags.GetBool("%s") err: %s`, name, err)
		}
		bval, err := strconv.ParseBool(sbval)
		if err != nil {
			return fmt.Errorf(`flags.ParseFlag: error converting --%s flag to bool using strconv.ParseBool("%s") err: %s`, name, sbval, err)
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
		case reflect.Int64:
			slval, err := getInt64Slice(name, flags)
			if err != nil {
				return fmt.Errorf(`flags.ParseFlag: error retrieving flag getInt64Slice("%s", flags) err: %s`, name, err)
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

// GetIntSlice return the []int64 value of a flag with the given name
func getInt64Slice(name string, flags *pflag.FlagSet) ([]int64, error) {
	flag := flags.Lookup(name)
	if flag == nil {
		err := fmt.Errorf("flag accessed but not defined: %s", name)
		return nil, err
	}

	if flag.Value.Type() != "intSlice" {
		err := fmt.Errorf("trying to get %s value of flag of type %s", "intSlice", flag.Value.Type())
		return nil, err
	}
	sval := flag.Value.String()
	result, err := int64SliceConv(sval)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return []int64{}, err
	}
	return result.([]int64), nil
}

func int64SliceConv(val string) (interface{}, error) {
	val = strings.Trim(val, "[]")
	// Empty string would cause a slice with one (empty) entry
	if len(val) == 0 {
		return []int{}, nil
	}
	ss := strings.Split(val, ",")
	out := make([]int64, len(ss))
	for i, d := range ss {
		var err error
		j, err := strconv.Atoi(d)
		out[i] = int64(j)
		if err != nil {
			return nil, err
		}

	}
	return out, nil
}
