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

package utils

import (
	"fmt"

	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cli/assets"
)

// Deprecated: Print the named help file. Usages should be replaced with getHelp(), see kvstore & forwarders as examples
func Help(fileName string) error {
	ret, err := getHelp(fileName)
	if err == nil {
		fmt.Println(ret)
	}
	return err
}

// Get contents of the named help file.
func getHelp(fileName string) (string, error) {
	filePath := fmt.Sprintf("help/%s", fileName) // TODO: replace with os-agnostic path formation
	b, err := assets.Read(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
