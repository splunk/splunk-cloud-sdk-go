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

package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestScloudBinaryWithCoverage(t *testing.T) {
	if os.Getenv("SCLOUD_TEST_DIR") == "" {
		t.Skip("Skipping TestScloudBinaryWithCoverage, set $SCLOUD_TEST_DIR in order to run this test ...")
	}
	// This is a trick to get coverage reports for the scloud binary -
	// it will run the dummy command "scloud version" by calling main directly
	// in order to generate coverage for tests below
	// see: https://blog.cloudflare.com/go-coverage-with-external-tests/
	os.Args[1] = "version"
	go main()

	// Follow closely, we're running 'go test' on this file ...
	// in order to call python tests ...
	// which call the scloud binary built from go code ...
	//                 ¯\_(ツ)_/¯
	dir := os.Getenv("SCLOUD_TEST_DIR")
	fmt.Println("Running `sh run.sh` from within scloud_integration_test.go for coverage ...")
	cmd := exec.Command("sh", "run.sh")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		panic(fmt.Sprintf("`sh run.sh` finished with error: %v\n\noutput:\n%s", err, out))
	}
	fmt.Printf("`sh run.sh` succeeded with output:\n\n%s", out)
}
