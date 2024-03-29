// +build ignore

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
	"strings"
	"text/template"
	"time"
)

var versionTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package version

const (
    BuildBranch = "{{ .Branch }}"
    BuildTime int64 = {{ .EpochSec }}
    Version = ScloudVersion
)
`))

type templateArgs struct {
	Branch   string
	EpochSec int64
}

// Prints an error message and exits.
func fatal(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	fmt.Fprintf(os.Stderr, "gen_version.go: %s\n", msg)
	os.Exit(1)
}

func main() {
	// check the git status
	cmd := exec.Command("git", "status", "-s")
	out, err := cmd.Output()
	if err != nil {
		fatal("error running git status: %v", err)
	}
	status := strings.Trim(fmt.Sprintf("%s", out), "\n\r")
	fmt.Printf("The git status is: \n%s\n", status)

	cmd = exec.Command("git", "symbolic-ref", "--short", "HEAD")
	out, err = cmd.Output()
	branch := "detatched-head"
	if err == nil {
		branch = strings.Trim(fmt.Sprintf("%s", out), "\n\r")
	}

	_ = os.Remove("version/version.go")

	f, err := os.Create("version/version.go")
	if err != nil {
		fatal("error writing version.go: %v", err)
	}
	defer f.Close()

	err = versionTemplate.Execute(f, templateArgs{
		Branch:   branch,
		EpochSec: time.Now().Unix(),
	})
	if err != nil {
		fatal("error executing template to write version.go: %v", err)
	}
}
