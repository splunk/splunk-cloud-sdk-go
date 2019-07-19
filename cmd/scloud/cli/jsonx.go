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
	"encoding/json"
	"strings"
)

// Dumps the given value to a string.
func dumps(value interface{}) string { //nolint:deadcode
	result, err := json.Marshal(value)
	if err != nil {
		return "[can't marshal value]"
	}
	return string(result[:])
}

// Loads a json "object" from the given string.
func loads(value string) (interface{}, error) {
	var result interface{}
	r := strings.NewReader(value)
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
