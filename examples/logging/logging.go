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

// This example demonstrates how to setup logging of requests/responses with the sdk using the standard Go "log" library.
//
// By default, this example logs to stout (for INFO level logs) and stderr (for ERROR level logs):
//    ```$ go run -v ./examples/logging/logging.go```
//
// To log INFO and ERROR to a single log file use:
//    ```$ go run -v ./examples/logging/logging.go -logfile=<path to logfile>```
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/services/identity"
	testutils "github.com/splunk/splunk-cloud-sdk-go/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

var logInfo *log.Logger
var logErr *log.Logger

func main() {
	// Setup logging to stdout and stderr by default
	logInfo, logErr = createStdLoggers()

	// If log file is specified, log there instead
	logFileArg := flag.String("logfile", "", "If non-empty, write log files in this file")
	flag.Parse()

	if logFileArg != nil && *logFileArg != "" {
		logFile, err := os.OpenFile(*logFileArg, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		exitOnError(err)
		defer logFile.Close()
		logInfo, logErr = createFileLoggers(logFile)
	}

	// Get client
	logInfo.Print("Creating identity service client")

	client, err := identity.NewService(&services.Config{
		Token:        testutils.TestAuthenticationToken,
		Host:         testutils.TestSplunkCloudHost,
		Tenant:       testutils.TestTenant,
		RoundTripper: util.NewSdkTransport(logInfo),
	})
	exitOnError(err)
	logInfo.Print("Validating token")
	input := identity.ValidateTokenQueryParams{Include: []identity.ValidateTokenincludeEnum{"principal", "tenant"}}

	info, err := client.ValidateToken(&input)
	exitOnError(err)
	logInfo.Print(fmt.Sprintf("Success! Info: %+v", info))
}

func exitOnError(err error) {
	if err != nil {
		if logErr != nil {
			logErr.Print(err)
		}
		os.Exit(1)
	}
}

func createStdLoggers() (infoLogger *log.Logger, errLogger *log.Logger) {
	return log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func createFileLoggers(logFile *os.File) (infoLogger *log.Logger, errLogger *log.Logger) {
	return log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
