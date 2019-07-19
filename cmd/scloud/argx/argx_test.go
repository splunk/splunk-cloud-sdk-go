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

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BoolArgs struct {
	P0 bool `arg:"0"`
	P1 bool `arg:"1"`
	N0 bool `arg:"n0"`
	N1 bool `arg:"n1"`
}

type StringArgs struct {
	P0 string `arg:"0"`
	P1 string `arg:"1"`
	N0 string `arg:"n0"`
	N1 string `arg:"n1"`
}

type MiscArgs struct {
	Bool   bool   `arg:"bool"`
	String string `arg:"string"`
	Int    int    `arg:"int"`
	Int16  int16  `arg:"int16"`
	Int32  int32  `arg:"int32"`
	Int64  int64  `arg:"int64"`
	UInt   uint   `arg:"uint"`
	UInt16 uint16 `arg:"uint16"`
	UInt32 uint32 `arg:"uint32"`
	UInt64 uint64 `arg:"uint64"`
}

func TestResultType(t *testing.T) {
	argv := []string{"..."}

	args := &StringArgs{}
	_, err := Parse(argv, &args)
	assert.Equal(t, errBadResultType, err)

	var sz string
	_, err = Parse(argv, &sz)
	assert.Equal(t, errBadResultType, err)

	var iface interface{}
	_, err = Parse(argv, &iface)
	assert.Equal(t, errBadResultType, err)

	var lst []interface{}
	_, err = Parse(argv, &lst)
	assert.Equal(t, errBadResultType, err)
}

func TestParseString(t *testing.T) {
	// positional args
	argv := []string{"zero", "one"}
	args := &StringArgs{}
	r, err := Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, "zero", args.P0)
	assert.Equal(t, "one", args.P1)
	assert.Equal(t, "", args.N0)
	assert.Equal(t, "", args.N1)

	// positional & named args
	argv = []string{"zero", "one", "-n0", "narg0", "-n1", "narg1"}
	args = &StringArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, "zero", args.P0)
	assert.Equal(t, "one", args.P1)
	assert.Equal(t, "narg0", args.N0)
	assert.Equal(t, "narg1", args.N1)

	// positional & named, with residual
	argv = []string{"zero", "one", "-n0", "narg0", "-n1", "narg1", "extra", "..."}
	args = &StringArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(r))
	assert.Equal(t, []string{"extra", "..."}, r)
	assert.Equal(t, "zero", args.P0)
	assert.Equal(t, "one", args.P1)
	assert.Equal(t, "narg0", args.N0)
	assert.Equal(t, "narg1", args.N1)

	argv = []string{"zero", "-n0", "narg0", "-n1", "narg1"}
	args = &StringArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, errTooFewArgs, err)
	assert.Equal(t, 0, len(r))
}

func TestParseBool(t *testing.T) {
	argv := []string{"0", "1", "-n0", "0", "-n1", "1"}
	args := &BoolArgs{}
	r, err := Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, false, args.P0)
	assert.Equal(t, true, args.P1)
	assert.Equal(t, false, args.N0)
	assert.Equal(t, true, args.N1)

	argv = []string{"f", "t", "-n0", "f", "-n1", "t"}
	args = &BoolArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, false, args.P0)
	assert.Equal(t, true, args.P1)
	assert.Equal(t, false, args.N0)
	assert.Equal(t, true, args.N1)

	argv = []string{"F", "T", "-n0", "F", "-n1", "T"}
	args = &BoolArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, false, args.P0)
	assert.Equal(t, true, args.P1)
	assert.Equal(t, false, args.N0)
	assert.Equal(t, true, args.N1)

	argv = []string{"false", "true", "-n0", "false", "-n1", "true"}
	args = &BoolArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, false, args.P0)
	assert.Equal(t, true, args.P1)
	assert.Equal(t, false, args.N0)
	assert.Equal(t, true, args.N1)

	argv = []string{"False", "True", "-n0", "False", "-n1", "True"}
	args = &BoolArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, false, args.P0)
	assert.Equal(t, true, args.P1)
	assert.Equal(t, false, args.N0)
	assert.Equal(t, true, args.N1)

	argv = []string{"FALSE", "TRUE", "-n0", "FALSE", "-n1", "TRUE"}
	args = &BoolArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, false, args.P0)
	assert.Equal(t, true, args.P1)
	assert.Equal(t, false, args.N0)
	assert.Equal(t, true, args.N1)

	// incompatible bool value interpreted as residual
	argv = []string{"-bool", "xyzzy"}
	misc := &MiscArgs{}
	r, err = Parse(argv, misc)
	assert.Equal(t, nil, err)
	assert.Equal(t, []string{"xyzzy"}, r)
}

func isBadArgValue(err error) bool {
	return strings.HasPrefix(err.Error(), "bad value for argument")
}

func isBadFlagValue(err error) bool {
	return strings.HasPrefix(err.Error(), "bad value for option")
}

func TestParseBoolError(t *testing.T) {
	argv := []string{"hello", "world"}
	args := &BoolArgs{}
	_, err := Parse(argv, args)
	assert.True(t, isBadArgValue(err))

	argv = []string{"1", "42"}
	args = &BoolArgs{}
	_, err = Parse(argv, args)
	assert.True(t, isBadArgValue(err))

	argv = []string{"-1", "-42"}
	args = &BoolArgs{}
	_, err = Parse(argv, args)
	assert.True(t, isBadArgValue(err))
}

func TestParseInt(t *testing.T) {
	argv := []string{"-int", "42"}
	args := &MiscArgs{}
	r, err := Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, 42, args.Int)

	argv = []string{"-int", "-42"}
	args = &MiscArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, -42, args.Int)

	argv = []string{"-int16", "42"}
	args = &MiscArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, int16(42), args.Int16)

	argv = []string{"-int16", "-42"}
	args = &MiscArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, int16(-42), args.Int16)

	argv = []string{"-int32", "42"}
	args = &MiscArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, int32(42), args.Int32)

	argv = []string{"-int32", "-42"}
	args = &MiscArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, int32(-42), args.Int32)

	argv = []string{"-int64", "42"}
	args = &MiscArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, int64(42), args.Int64)

	argv = []string{"-int64", "-42"}
	args = &MiscArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, int64(-42), args.Int64)
}

func TestParseIntError(t *testing.T) {
	argv := []string{"-int", "xyzzy"}
	args := &MiscArgs{}
	_, err := Parse(argv, args)
	assert.True(t, isBadFlagValue(err))

	argv = []string{"-int16", "xyzzy"}
	args = &MiscArgs{}
	_, err = Parse(argv, args)
	assert.True(t, isBadFlagValue(err))

	argv = []string{"-int32", "xyzzy"}
	args = &MiscArgs{}
	_, err = Parse(argv, args)
	assert.True(t, isBadFlagValue(err))

	argv = []string{"-int64", "xyzzy"}
	args = &MiscArgs{}
	_, err = Parse(argv, args)
	assert.True(t, isBadFlagValue(err))
}

func TestParseUInt(t *testing.T) {
	argv := []string{"-uint", "42"}
	args := &MiscArgs{}
	r, err := Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, uint(42), args.UInt)

	argv = []string{"-uint16", "42"}
	args = &MiscArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, uint16(42), args.UInt16)

	argv = []string{"-uint32", "42"}
	args = &MiscArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, uint32(42), args.UInt32)

	argv = []string{"-uint64", "42"}
	args = &MiscArgs{}
	r, err = Parse(argv, args)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, len(r))
	assert.Equal(t, uint64(42), args.UInt64)
}

func TestParseUIntError(t *testing.T) {
	argv := []string{"-uint", "-42"}
	args := &MiscArgs{}
	_, err := Parse(argv, args)
	assert.True(t, isBadFlagValue(err))

	argv = []string{"-uint16", "-42"}
	args = &MiscArgs{}
	_, err = Parse(argv, args)
	assert.True(t, isBadFlagValue(err))

	argv = []string{"-uint32", "-42"}
	args = &MiscArgs{}
	_, err = Parse(argv, args)
	assert.True(t, isBadFlagValue(err))

	argv = []string{"-uint64", "-42"}
	args = &MiscArgs{}
	_, err = Parse(argv, args)
	assert.True(t, isBadFlagValue(err))
}
