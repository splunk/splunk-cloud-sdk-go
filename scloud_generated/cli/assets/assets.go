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

package assets

// Helper functions for accessing static assets.

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/rakyll/statik/fs"

	// Import needed to register files with fs
	_ "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/cli/statik"
)

// Open the named static file asset.
func Open(fileName string) (io.Reader, error) {
	statikFs, err := fs.New()
	if err != nil {
		return nil, fmt.Errorf("assets.go: err calling fs.New() %v", err)
	}
	filePath := "/" + fileName
	httpFs := http.FileSystem(statikFs)
	b, err := fs.ReadFile(httpFs, filePath)
	if err != nil {
		return nil, fmt.Errorf("assets.go: err opening %s %v", filePath, err)
	}
	return bytes.NewReader(b), nil

}

// Read the named static file asset.
func Read(fileName string) ([]byte, error) {
	statikFs, err := fs.New()
	if err != nil {
		return nil, fmt.Errorf("assets.go: err calling fs.New() %v", err)
	}
	filePath := "/" + fileName
	httpFs := http.FileSystem(statikFs)
	b, err := fs.ReadFile(httpFs, filePath)
	if err != nil {
		return nil, fmt.Errorf("assets.go: err reading %s %v", filePath, err)
	}
	return b, nil
}
