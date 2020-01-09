/*
 * Copyright 2020 Splunk, Inc.
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

package flags

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testFlagName = "my-flag"

// testStringRequired is a helper that accepts arguments (args), a flagDefault, an expected flag
// value (flagExpVal) after parsing the FlagSet and an expected value (expVal) after parsing with
// ParseFlag to get the expected value of a given type (in this case string)
func testStringRequired(t *testing.T, args []string, flagDefault string, flagExpVal string, expVal string) {
	fs := pflag.FlagSet{}
	var d string
	fs.StringVar(&d, testFlagName, flagDefault, "Test flag for a string")
	err := fs.Parse(args)
	require.Nil(t, err)
	assert.Equal(t, flagExpVal, d)
	v := ""
	err = ParseFlag(&fs, testFlagName, &v)
	require.Nil(t, err)
	assert.Equal(t, expVal, v)
}

// TestStringDefaultRequired Test defaults for required flags where value is a string
func TestStringDefaultRequired(t *testing.T) {
	// Check that defaults are parsed properly if flag not specified
	testStringRequired(t, []string{}, "", "", "")
	testStringRequired(t, []string{}, "my_default", "my_default", "my_default")
	testStringRequired(t, []string{}, "0009.2", "0009.2", "0009.2")
	testStringRequired(t, []string{}, "1", "1", "1")
	testStringRequired(t, []string{}, "true", "true", "true")
	testStringRequired(t, []string{}, `{"foo":true}`, `{"foo":true}`, `{"foo":true}`)
}

// TestStringRequired Test required flags where value is a string
func TestStringRequired(t *testing.T) {
	testStringRequired(t, []string{"--" + testFlagName, ""}, "default", "", "")
	testStringRequired(t, []string{"--" + testFlagName, "my_string"}, "", "my_string", "my_string")
	testStringRequired(t, []string{"--" + testFlagName, "0009.2"}, "", "0009.2", "0009.2")
	testStringRequired(t, []string{"--" + testFlagName, "1"}, "", "1", "1")
	testStringRequired(t, []string{"--" + testFlagName, "true"}, "", "true", "true")
	testStringRequired(t, []string{"--" + testFlagName, `{"foo":true}`}, "default", `{"foo":true}`, `{"foo":true}`)
}

// testStringOptional is a helper that accepts arguments (args), a flagDefault, an expected flag
// value (flagExpVal) after parsing the FlagSet and an expected value (expVal) after parsing with
// ParseFlag to get the expected value of a given type (in this case a pointer to a string)
func testStringOptional(t *testing.T, args []string, flagDefault string, flagExpVal string, expVal *string) {
	fs := pflag.FlagSet{}
	var d string
	fs.StringVar(&d, testFlagName, flagDefault, "Test flag for a string")
	err := fs.Parse(args)
	require.Nil(t, err)
	assert.Equal(t, flagExpVal, d)
	vv := ""
	v := &vv
	err = ParseFlag(&fs, testFlagName, &v)
	require.Nil(t, err)
	assert.Equal(t, expVal, v)
}

// TestStringDefaultOptional Test defaults for optional flags where value is a pointer to a string
func TestStringDefaultOptional(t *testing.T) {
	// Check that defaults are parsed properly if flag not specified (should be nil)
	testStringOptional(t, []string{}, "", "", nil)
	testStringOptional(t, []string{}, "my_default", "my_default", nil)
	testStringOptional(t, []string{}, "0009.2", "0009.2", nil)
	testStringOptional(t, []string{}, "1", "1", nil)
	testStringOptional(t, []string{}, "true", "true", nil)
	testStringOptional(t, []string{}, `{"foo":true}`, `{"foo":true}`, nil)
}

// TestStringOptional Test optional flags where value is a pointer to a string
func TestStringOptional(t *testing.T) {
	s1 := ""
	testStringOptional(t, []string{"--" + testFlagName, ""}, "my_default", "", &s1)
	s2 := "my_string"
	testStringOptional(t, []string{"--" + testFlagName, "my_string"}, "", "my_string", &s2)
	s3 := "0009.2"
	testStringOptional(t, []string{"--" + testFlagName, "0009.2"}, "", "0009.2", &s3)
	s4 := "1"
	testStringOptional(t, []string{"--" + testFlagName, "1"}, "", "1", &s4)
	s5 := "true"
	testStringOptional(t, []string{"--" + testFlagName, "true"}, "", "true", &s5)
	s6 := `{"foo":true}`
	testStringOptional(t, []string{"--" + testFlagName, `{"foo":true}`}, "", `{"foo":true}`, &s6)
}

// testStringBoolRequired is a helper that accepts arguments (args), a flagDefault, an expected flag
// value (flagExpVal) after parsing the FlagSet and an expected value (expVal) after parsing with
// ParseFlag to get the expected value of a given type (in this case a bool)
func testStringBoolRequired(t *testing.T, args []string, flagDefault string, flagExpVal string, expVal bool) {
	fs := pflag.FlagSet{}
	var d string
	fs.StringVar(&d, testFlagName, flagDefault, "Test flag for a bool")
	err := fs.Parse(args)
	require.Nil(t, err)
	assert.Equal(t, flagExpVal, d)
	v := !expVal
	err = ParseFlag(&fs, testFlagName, &v)
	require.Nil(t, err)
	assert.Equal(t, expVal, v)
}

// TestStringBoolDefaults Test defaults for required flags where value is a bool
func TestStringBoolDefaultRequired(t *testing.T) {
	// Check that defaults are parsed properly if flag not specified
	testStringBoolRequired(t, []string{}, "true", "true", true)
	testStringBoolRequired(t, []string{}, "false", "false", false)
	// Bool-like defaults should also be parsed according to strconv.ParseBool
	testStringBoolRequired(t, []string{}, "TRUE", "TRUE", true)
	testStringBoolRequired(t, []string{}, "t", "t", true)
	testStringBoolRequired(t, []string{}, "1", "1", true)
	testStringBoolRequired(t, []string{}, "FALSE", "FALSE", false)
	testStringBoolRequired(t, []string{}, "f", "f", false)
	testStringBoolRequired(t, []string{}, "0", "0", false)
}

// testStringBoolOptional is a helper that accepts arguments (args), a flagDefault, an expected flag
// value (flagExpVal) after parsing the FlagSet and an expected value (expVal) after parsing with
// ParseFlag to get the expected value of a given type (in this case a pointer to a bool)
func testStringBoolOptional(t *testing.T, args []string, flagDefault string, flagExpVal string, expVal *bool) {
	fs := pflag.FlagSet{}
	var d string
	fs.StringVar(&d, testFlagName, flagDefault, "Test flag for a bool")
	err := fs.Parse(args)
	require.Nil(t, err)
	assert.Equal(t, flagExpVal, d)
	vv := false
	v := &vv
	err = ParseFlag(&fs, testFlagName, &v)
	require.Nil(t, err)
	assert.Equal(t, expVal, v)
}

// TestStringBoolDefaults Test defaults for optional flags where value is a pointer to a bool
func TestStringBoolDefaultOptional(t *testing.T) {
	// Check that defaults are parsed properly if flag not specified
	testStringBoolOptional(t, []string{}, "true", "true", nil)
	testStringBoolOptional(t, []string{}, "false", "false", nil)
}

// TestStringBoolTrue tests inputs where value should always parse to true
func TestStringBoolTrue(t *testing.T) {
	expVal := true
	testStringBoolRequired(t, []string{"--" + testFlagName, "true"}, "false", "true", expVal)
	testStringBoolRequired(t, []string{"--" + testFlagName, "TRUE"}, "false", "TRUE", expVal)
	testStringBoolRequired(t, []string{"--" + testFlagName, "t"}, "false", "t", expVal)
	testStringBoolRequired(t, []string{"--" + testFlagName, "1"}, "false", "1", expVal)
	testStringBoolOptional(t, []string{"--" + testFlagName, "true"}, "false", "true", &expVal)
	testStringBoolOptional(t, []string{"--" + testFlagName, "TRUE"}, "false", "TRUE", &expVal)
	testStringBoolOptional(t, []string{"--" + testFlagName, "t"}, "false", "t", &expVal)
	testStringBoolOptional(t, []string{"--" + testFlagName, "1"}, "false", "1", &expVal)
}

// TestStringBoolFalse tests inputs where value should always parse to false
func TestStringBoolFalse(t *testing.T) {
	expVal := false
	testStringBoolRequired(t, []string{"--" + testFlagName, "false"}, "false", "false", expVal)
	testStringBoolRequired(t, []string{"--" + testFlagName, "FALSE"}, "false", "FALSE", expVal)
	testStringBoolRequired(t, []string{"--" + testFlagName, "f"}, "false", "f", expVal)
	testStringBoolRequired(t, []string{"--" + testFlagName, "0"}, "false", "0", expVal)
	testStringBoolOptional(t, []string{"--" + testFlagName, "false"}, "false", "false", &expVal)
	testStringBoolOptional(t, []string{"--" + testFlagName, "FALSE"}, "false", "FALSE", &expVal)
	testStringBoolOptional(t, []string{"--" + testFlagName, "f"}, "false", "f", &expVal)
	testStringBoolOptional(t, []string{"--" + testFlagName, "0"}, "false", "0", &expVal)
}

// TestIntsDefaultRequired tests inputs where default value is an int
func TestIntsDefaultRequired(t *testing.T) {
	fs := pflag.FlagSet{}
	var (
		i   int
		i32 int32
		i64 int64
	)
	id := int(1234)
	fs.IntVar(&i, "my-int", id, "Test flag for an int")
	id32 := int32(12345)
	fs.Int32Var(&i32, "my-int32", id32, "Test flag for an int")
	id64 := int64(123456)
	fs.Int64Var(&i64, "my-int64", id64, "Test flag for an int")
	err := fs.Parse([]string{})
	require.Nil(t, err)
	assert.Equal(t, id, i)
	assert.Equal(t, id32, i32)
	assert.Equal(t, id64, i64)
	var (
		vi   int
		vi32 int32
		vi64 int64
	)
	err = ParseFlag(&fs, "my-int", &vi)
	require.Nil(t, err)
	err = ParseFlag(&fs, "my-int32", &vi32)
	require.Nil(t, err)
	err = ParseFlag(&fs, "my-int64", &vi64)
	require.Nil(t, err)
	assert.Equal(t, id, vi)
	assert.Equal(t, id32, vi32)
	assert.Equal(t, id64, vi64)
}

// TestIntsDefaultOptional tests inputs where default value is a pointer to an int,
// if there are no args the values should parse as nil
func TestIntsDefaultOptional(t *testing.T) {
	fs := pflag.FlagSet{}
	var (
		i   int
		i32 int32
		i64 int64
	)
	id := int(1234)
	fs.IntVar(&i, "my-int", id, "Test flag for an int")
	id32 := int32(12345)
	fs.Int32Var(&i32, "my-int32", id32, "Test flag for an int")
	id64 := int64(123456)
	fs.Int64Var(&i64, "my-int64", id64, "Test flag for an int")
	err := fs.Parse([]string{})
	require.Nil(t, err)
	assert.Equal(t, id, i)
	assert.Equal(t, id32, i32)
	assert.Equal(t, id64, i64)
	var (
		vi   *int
		vi32 *int32
		vi64 *int64
	)
	err = ParseFlag(&fs, "my-int", &vi)
	require.Nil(t, err)
	err = ParseFlag(&fs, "my-int32", &vi32)
	require.Nil(t, err)
	err = ParseFlag(&fs, "my-int64", &vi64)
	require.Nil(t, err)
	assert.Nil(t, vi)
	assert.Nil(t, vi32)
	assert.Nil(t, vi64)
}

// TestIntsRequired tests inputs where value is an int
func TestIntsRequired(t *testing.T) {
	fs := pflag.FlagSet{}
	var (
		i   int
		i32 int32
		i64 int64
	)
	fs.IntVar(&i, "my-int", 0, "Test flag for an int")
	fs.Int32Var(&i32, "my-int32", 0, "Test flag for an int")
	fs.Int64Var(&i64, "my-int64", 0, "Test flag for an int")
	err := fs.Parse([]string{"--my-int64", "123456", "--my-int32", "12345", "--my-int", "1234"})
	require.Nil(t, err)
	iv := int(1234)
	iv32 := int32(12345)
	iv64 := int64(123456)
	assert.Equal(t, iv, i)
	assert.Equal(t, iv32, i32)
	assert.Equal(t, iv64, i64)
	var (
		vi   int
		vi32 int32
		vi64 int64
	)
	err = ParseFlag(&fs, "my-int", &vi)
	require.Nil(t, err)
	err = ParseFlag(&fs, "my-int32", &vi32)
	require.Nil(t, err)
	err = ParseFlag(&fs, "my-int64", &vi64)
	require.Nil(t, err)
	assert.Equal(t, iv, vi)
	assert.Equal(t, iv32, vi32)
	assert.Equal(t, iv64, vi64)
}

// TestIntsOptional tests inputs where value is a pointer to an int
func TestIntsOptional(t *testing.T) {
	fs := pflag.FlagSet{}
	var (
		i   int
		i32 int32
		i64 int64
	)
	fs.IntVar(&i, "my-int", 0, "Test flag for an int")
	fs.Int32Var(&i32, "my-int32", 0, "Test flag for an int")
	fs.Int64Var(&i64, "my-int64", 0, "Test flag for an int")
	err := fs.Parse([]string{"--my-int64", "123456", "--my-int32", "12345", "--my-int", "1234"})
	require.Nil(t, err)
	iv := int(1234)
	iv32 := int32(12345)
	iv64 := int64(123456)
	assert.Equal(t, iv, i)
	assert.Equal(t, iv32, i32)
	assert.Equal(t, iv64, i64)
	var (
		dvi   int   = 0
		dvi32 int32 = 0
		dvi64 int64 = 0
		vi          = &dvi
		vi32        = &dvi32
		vi64        = &dvi64
	)
	err = ParseFlag(&fs, "my-int", &vi)
	require.Nil(t, err)
	err = ParseFlag(&fs, "my-int32", &vi32)
	require.Nil(t, err)
	err = ParseFlag(&fs, "my-int64", &vi64)
	require.Nil(t, err)
	assert.Equal(t, iv, *vi)
	assert.Equal(t, iv32, *vi32)
	assert.Equal(t, iv64, *vi64)
}

// TestFloatsRequired tests inputs where value is a float
func TestFloatsRequired(t *testing.T) {
	fs := pflag.FlagSet{}
	var (
		f32 float32
		f64 float64
	)
	fs.Float32Var(&f32, "my-float32", 0, "Test flag for a float")
	fs.Float64Var(&f64, "my-float64", 0, "Test flag for a float")
	err := fs.Parse([]string{"--my-float64", "1234.56", "--my-float32", "1.2345"})
	require.Nil(t, err)
	fv32 := float32(1.2345)
	fv64 := float64(1234.56)
	assert.Equal(t, fv32, f32)
	assert.Equal(t, fv64, f64)
	var (
		vf32 float32
		vf64 float64
	)
	err = ParseFlag(&fs, "my-float32", &vf32)
	require.Nil(t, err)
	err = ParseFlag(&fs, "my-float64", &vf64)
	require.Nil(t, err)
	assert.Equal(t, fv32, vf32)
	assert.Equal(t, fv64, vf64)
}

// TestFloatsOptional tests inputs where value is a pointer to a float
func TestFloatsOptional(t *testing.T) {
	fs := pflag.FlagSet{}
	var (
		f32 float32
		f64 float64
	)
	fs.Float32Var(&f32, "my-float32", 0, "Test flag for a float")
	fs.Float64Var(&f64, "my-float64", 0, "Test flag for a float")
	err := fs.Parse([]string{"--my-float64", "1234.56", "--my-float32", "1.2345"})
	require.Nil(t, err)
	fv32 := float32(1.2345)
	fv64 := float64(1234.56)
	assert.Equal(t, fv32, f32)
	assert.Equal(t, fv64, f64)
	var (
		dvf32 float32 = 0
		dvf64 float64 = 0
		vf32          = &dvf32
		vf64          = &dvf64
	)
	err = ParseFlag(&fs, "my-float32", &vf32)
	require.Nil(t, err)
	err = ParseFlag(&fs, "my-float64", &vf64)
	require.Nil(t, err)
	assert.Equal(t, fv32, *vf32)
	assert.Equal(t, fv64, *vf64)
}

// TestStringSliceMultipleArgs tests StringSliceVar flags where the expected type is a simple
// slice of strings ([]string) and the input arguments are of the form --arg1 val1 --arg1 val2
// which is supported by cobra by default
func TestStringSliceMultipleArgs(t *testing.T) {
	fs := pflag.FlagSet{}
	// this var is provided in our generated cmd files but never used
	var s []string
	fs.StringSliceVar(&s, "addresses", nil, "Test flag for a slice of strings")
	const (
		addr1 = "one@example.com"
		addr2 = "two@example.com"
	)
	err := fs.Parse([]string{"--addresses", addr1, "--addresses", addr2})
	require.Nil(t, err)
	require.Equal(t, 2, len(s))
	assert.Equal(t, addr1, s[0])
	assert.Equal(t, addr2, s[1])
	// this var is provided in our generated pkg files and what is actually parsed
	var slice []string
	err = ParseFlag(&fs, "addresses", &slice)
	require.Nil(t, err)
	require.Equal(t, 2, len(slice))
	assert.Equal(t, addr1, slice[0])
	assert.Equal(t, addr2, slice[1])
}

// TestStringSliceCommaDelim tests StringSliceVar flags where the expected type is a simple
// slice of strings ([]string) and the input arguments are of the form --arg1 val1,val2
// which is supported by cobra by default
func TestStringSliceCommaDelim(t *testing.T) {
	fs := pflag.FlagSet{}
	// this var is provided in our generated cmd files but never used
	var s []string
	fs.StringSliceVar(&s, "addresses", nil, "Test flag for a slice of strings")
	const (
		addr1 = "one@example.com"
		addr2 = "two@example.com"
	)
	err := fs.Parse([]string{"--addresses", addr1 + "," + addr2})
	require.Nil(t, err)
	require.Equal(t, 2, len(s))
	assert.Equal(t, addr1, s[0])
	assert.Equal(t, addr2, s[1])
	// this var is provided in our generated pkg files and what is actually parsed
	var slice []string
	err = ParseFlag(&fs, "addresses", &slice)
	require.Nil(t, err)
	require.Equal(t, 2, len(slice))
	assert.Equal(t, addr1, slice[0])
	assert.Equal(t, addr2, slice[1])
}

// TestIntSliceMultipleArgs tests IntSliceVar flags where the expected type is a simple
// slice of ints ([]int or []int64) and the input arguments are of the form --arg1 val1 --arg1 val2
// which is supported by cobra by default
func TestIntSliceMultipleArgs(t *testing.T) {
	fs := pflag.FlagSet{}
	// this var is provided in our generated cmd files but never used
	var s []int
	fs.IntSliceVar(&s, "slots", nil, "Test flag for a slice of ints")
	const (
		slot1 = int(11)
		slot2 = int(222)
	)
	err := fs.Parse([]string{"--slots", "11", "--slots", "222"})
	require.Nil(t, err)
	require.Equal(t, 2, len(s))
	assert.Equal(t, slot1, s[0])
	assert.Equal(t, slot2, s[1])
	// this var is provided in our generated pkg files and what is actually parsed
	var slice []int
	err = ParseFlag(&fs, "slots", &slice)
	require.Nil(t, err)
	require.Equal(t, 2, len(slice))
	assert.Equal(t, slot1, slice[0])
	assert.Equal(t, slot2, slice[1])
	var slice64 []int64
	err = ParseFlag(&fs, "slots", &slice64)
	require.Nil(t, err)
	require.Equal(t, 2, len(slice64))
	assert.Equal(t, int64(slot1), slice64[0])
	assert.Equal(t, int64(slot2), slice64[1])
}

// TestIntSliceCommaDelim tests IntSliceVar flags where the expected type is a simple
// slice of ints ([]int or []int64) and the input arguments are of the form --arg1 val1,val2
// which is supported by cobra by default
func TestIntSliceCommaDelim(t *testing.T) {
	fs := pflag.FlagSet{}
	// this var is provided in our generated cmd files but never used
	var s []int
	fs.IntSliceVar(&s, "slots", nil, "Test flag for a slice of ints")
	const (
		slot1 = int(11)
		slot2 = int(222)
	)
	err := fs.Parse([]string{"--slots", "11,222"})
	require.Nil(t, err)
	require.Equal(t, 2, len(s))
	assert.Equal(t, slot1, s[0])
	assert.Equal(t, slot2, s[1])
	// this var is provided in our generated pkg files and what is actually parsed
	var slice []int
	err = ParseFlag(&fs, "slots", &slice)
	require.Nil(t, err)
	require.Equal(t, 2, len(slice))
	assert.Equal(t, slot1, slice[0])
	assert.Equal(t, slot2, slice[1])
	var slice64 []int64
	err = ParseFlag(&fs, "slots", &slice64)
	require.Nil(t, err)
	require.Equal(t, 2, len(slice64))
	assert.Equal(t, int64(slot1), slice64[0])
	assert.Equal(t, int64(slot2), slice64[1])
}

// TestBoolSliceMultipleArgs tests BoolSliceVar flags where the expected type is a simple
// slice of bools ([]bool) and the input arguments are of the form --arg1 val1 --arg1 val2
// which is supported by cobra by default
func TestBoolSliceMultipleArgs(t *testing.T) {
	fs := pflag.FlagSet{}
	// this var is provided in our generated cmd files but never used
	var s []bool
	fs.BoolSliceVar(&s, "truths", nil, "Test flag for a slice of bools")
	const (
		t1 = true
		t2 = false
		t3 = false
		t4 = true
	)
	err := fs.Parse([]string{"--truths", "true", "--truths", "false", "--truths", "0", "--truths", "T"})
	require.Nil(t, err)
	require.Equal(t, 4, len(s))
	assert.Equal(t, t1, s[0])
	assert.Equal(t, t2, s[1])
	assert.Equal(t, t3, s[2])
	assert.Equal(t, t4, s[3])
	// this var is provided in our generated pkg files and what is actually parsed
	var slice []bool
	err = ParseFlag(&fs, "truths", &slice)
	require.Nil(t, err)
	require.Equal(t, 4, len(slice))
	assert.Equal(t, t1, slice[0])
	assert.Equal(t, t2, slice[1])
	assert.Equal(t, t3, slice[2])
	assert.Equal(t, t4, slice[3])
}

// TestBoolSliceCommaDelim tests BoolSliceVar flags where the expected type is a simple
// slice of bools ([]bool) and the input arguments are of the form --arg1 val1,val2
// which is supported by cobra by default
func TestBoolSliceCommaDelim(t *testing.T) {
	fs := pflag.FlagSet{}
	// this var is provided in our generated cmd files but never used
	var s []bool
	fs.BoolSliceVar(&s, "truths", nil, "Test flag for a slice of bools")
	const (
		t1 = true
		t2 = false
		t3 = false
		t4 = true
	)
	err := fs.Parse([]string{"--truths", "true,false,0,T"})
	require.Nil(t, err)
	require.Equal(t, 4, len(s))
	assert.Equal(t, t1, s[0])
	assert.Equal(t, t2, s[1])
	assert.Equal(t, t3, s[2])
	assert.Equal(t, t4, s[3])
	// this var is provided in our generated pkg files and what is actually parsed
	var slice []bool
	err = ParseFlag(&fs, "truths", &slice)
	require.Nil(t, err)
	require.Equal(t, 4, len(slice))
	assert.Equal(t, t1, slice[0])
	assert.Equal(t, t2, slice[1])
	assert.Equal(t, t3, slice[2])
	assert.Equal(t, t4, slice[3])
}
