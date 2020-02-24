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

package util

import (
	"fmt"
	"os"

	"github.com/golang/glog"
)

// Prints an info message.
func Info(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	glog.Infof("info: %s\n", msg)
}

// Prints a fatal error message and exits.
func Fatal(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	glog.Exitf("fatal error: %s\n", msg)
}

// Prints a warning message.
func Warning(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	glog.Warningf("warning: %s\n", msg)
}

// Prints an error message.
func Error(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	glog.Errorf("error: %s\n", msg)
}

// Checks if log directory exists. It creates the log directory, if it doesn't exist.
func CreateLogDirectory(logPath string) {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		err = os.Mkdir(logPath, 0700)
		if err != nil {
			fmt.Printf("error creating log folder: %s", err.Error())
		}
	}
}
