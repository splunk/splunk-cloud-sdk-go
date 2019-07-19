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

package argx

//
// The argx command line parser maps a vector of strings onto a struct
// containing fields annotated with tags indicating how they should be
// interpreted by the parser.
//
// The vector may consists of positional arguments (args) and named arguments
// also known collectively as options or individually as flags.
//
// You can think of the positional arguments as similar to standard positional
// args on a function call, and the options as optional named arguments, for
// example **kwargs in python.
//
// Positional arguments are a fixed number (possibly 0) and the vector must
// match the expected count. If there are extras, the remaining vector will
// be returned as a residual for subsequent processing.
//
// Flags, aka named args, are optional.
//
// Example:
//
//     type MyArgs struct {
//         Name    string `arg:"0"`
//         Kind    string `arg:"1"`
//         Limit   int    `arg:"limit"`
//         Verbose bool   `arg:"v"`
//     }
//
// .. which correspond to a command argument vector that looks like this:
//
//     <command> <name> <kind> [-limit <int>] [-v <bool>]
//
// Additional rules:
//
//   * It is legal to mix named and positional arguments.
//
//   * Will parse until it consumes all inputs, or it finds a positional arg
//     larger than it expects, in which case it will return that argument and
//     everything following it as the residual.
//
//   * Any undeclared positional args between 0 and max will be ignored,
//     think of this as though the argument was declared _, eg:
//
//         foo(arg1, _, arg3)
//
//   * All flags, except for booleans require a compatible argument value.
//     Boolean flags flags take an optional value, and if the value is omitted
//     it defaults to true.
//

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// [items] => head, [tail]
func head(items []string) (string, []string) {
	if items == nil {
		return "", nil
	}
	n := len(items)
	if n == 0 {
		return "", nil
	}
	h := items[0]
	if n == 1 {
		return h, nil
	}
	return h, items[1:]
}

func peek(items []string) string {
	if items == nil {
		return ""
	}
	n := len(items)
	if n == 0 {
		return ""
	}
	return items[0]
}

func push(item string, items []string) []string {
	return append([]string{item}, items...)
}

// Answers if the given string is an option flag.
// A legal flag starts with a '-' or '--' and has a base name that starts with
// an ASCII letter. This ensures that we don't confuse flags with negative
// numeric arguments.
func isFlag(value string) bool {
	if value == "" {
		return false
	}
	b := value[0]
	if b != '-' {
		return false
	}
	value = flagName(value)
	if value == "" {
		return false
	}
	b = value[0]
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// Returns the base name of the given flag.
func flagName(value string) string {
	// assert isFlag(value)
	if len(value) > 0 && value[0] == '-' {
		value = value[1:]
		if len(value) > 0 && value[0] == '-' {
			value = value[1:]
		}
	}
	return value
}

func isBool(slot reflect.Value) bool {
	if slot.Type().Kind() == reflect.Bool {
		return true
	}
	return false
}

// Assign the given value to the given slot, with data type conversion.
func assign(slot reflect.Value, value string) error {
	t := slot.Type()
	switch t.Kind() {
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		slot.SetBool(b)
	case reflect.Int, reflect.Int64:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		slot.SetInt(i)
	case reflect.Int16:
		i, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return err
		}
		slot.SetInt(i)
	case reflect.Int32:
		i, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		slot.SetInt(i)
	case reflect.String:
		slot.SetString(value)
	case reflect.Uint, reflect.Uint64:
		u, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		slot.SetUint(u)
	case reflect.Uint16:
		u, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return err
		}
		slot.SetUint(u)
	case reflect.Uint32:
		u, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		slot.SetUint(u)
	default:
		return fmt.Errorf("bad field type: %s", t.Name())
	}
	return nil
}

// Answers if the given value can be assigned to the given slot.
func canAssign(slot reflect.Value, value string) bool {
	var err error
	if value == "" {
		return false
	}
	t := slot.Type()
	switch t.Kind() {
	case reflect.Bool:
		_, err = strconv.ParseBool(value)
	case reflect.Int, reflect.Int64:
		_, err = strconv.ParseInt(value, 10, 64)
	case reflect.Int16:
		_, err = strconv.ParseInt(value, 10, 16)
	case reflect.Int32:
		_, err = strconv.ParseInt(value, 10, 32)
	case reflect.String:
		// nop
	case reflect.Uint, reflect.Uint64:
		_, err = strconv.ParseUint(value, 10, 64)
	case reflect.Uint16:
		_, err = strconv.ParseUint(value, 10, 16)
	case reflect.Uint32:
		_, err = strconv.ParseUint(value, 10, 32)
	default:
		return false
	}
	if err != nil {
		return false
	}
	return true
}

var errBadResultType = errors.New("bad result type, need *struct")
var errTooFewArgs = errors.New("too few arguments")

func Parse(argv []string, result interface{}) ([]string, error) {
	// validate and prepare arguments
	t := reflect.TypeOf(result)
	if t.Kind() != reflect.Ptr {
		return nil, errBadResultType
	}
	t = t.Elem()
	if t.Kind() != reflect.Struct {
		return nil, errBadResultType
	}
	v := reflect.ValueOf(result).Elem()

	// construct arg and option maps
	// todo: check for redifinition of arg or flag
	maxArg := -1
	args := make(map[int]reflect.Value)
	opts := make(map[string]reflect.Value)
	fnum := t.NumField()
	for i := 0; i < fnum; i++ {
		field := t.Field(i)
		tag := field.Tag.Get("arg")
		if tag == "" || tag == "-" {
			continue // ignore
		}
		slot := v.Field(i)
		argn, err := strconv.Atoi(tag)
		if err == nil {
			args[argn] = slot
			if maxArg < argn {
				maxArg = argn
			}
		} else {
			opts[tag] = slot
		}
		// fmt.Printf("%s %s %s\n", field.Name, field.Type.Name(), tag)
	}

	nextArg := 0 // next arg
	for {
		var arg string
		arg, argv = head(argv)
		if arg == "" {
			break
		}
		if isFlag(arg) {
			flag := flagName(arg)
			slot, ok := opts[flag]
			if !ok {
				return nil, fmt.Errorf("unknown option: %s", arg)
			}
			value := peek(argv)
			if canAssign(slot, value) {
				_, argv = head(argv) // valid arg, accept
			} else if isBool(slot) {
				value = "true" // default
			} else if value == "" {
				return nil, fmt.Errorf("no value given for %s", arg)
			} else {
				return nil, fmt.Errorf("bad value for option %s: %s", arg, value)
			}
			if err := assign(slot, value); err != nil {
				return nil, err
			}
		} else {
			if nextArg > maxArg {
				return push(arg, argv), nil // residual
			}
			if slot, ok := args[nextArg]; ok {
				if err := assign(slot, arg); err != nil {
					return nil, fmt.Errorf("bad value for argument: %s, want %s",
						arg, slot.Type().Name())
				}
			}
			nextArg++
		}
	}
	if nextArg <= maxArg {
		return nil, errTooFewArgs
	}
	return nil, nil
}
