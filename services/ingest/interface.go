/*
 * Copyright © 2019 Splunk, Inc.
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

//This file contains interfaces that can't be auto-generated from codegen and interfaces that are auto-generated from codegen

package ingest

import (
	"io"
	"net/http"
)

// Servicer represents the interface for implementing all endpoints for this service
type Servicer interface {
	//interfaces that cannot be auto-generated from codegen
	/*
	   	NewBatchEventsSenderWithMaxAllowedError initializes a BatchEventsSender to collect events and send them as a single
	           batched request when a maximum event batch size, time interval, or maximum payload size is reached. It also
	           validates the user input for BatchEventSender.
	   	Parameters:
	   		batchSize: maximum number of events to reach before sending the batch, default maximum is 500
	   		interval: milliseconds to wait before sending the batch if other conditions have not been met
	   		dataSize: bytes that the overall payload should not exceed before sending, default maximum is 1040000 ~1MiB
	   		maxErrorsAllowed: number of errors after which the BatchEventsSender will stop
	*/
	NewBatchEventsSenderWithMaxAllowedError(batchSize int, interval int64, dataSize int, maxErrorsAllowed int) (*BatchEventsSender, error)
	/*
	   	NewBatchEventsSender initializes a BatchEventsSender to collect events and send them as a single batched
	           request when a maximum event batch size, time interval, or maximum payload size is reached.
	   	Parameters:
	   		batchSize: maximum number of events to reach before sending the batch, default maximum is 500
	   		interval: milliseconds to wait before sending the batch if other conditions have not been met
	   		payLoadSize: bytes that the overall payload should not exceed before sending, default maximum is 1040000 ~1MiB
	*/
	NewBatchEventsSender(batchSize int, interval int64, payLoadSize int) (*BatchEventsSender, error)
	/*
		UploadFilesStream - Upload stream of io.Reader.
		Parameters:
			stream
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UploadFilesStream(stream io.Reader, resp ...*http.Response) error

	//interfaces that are auto-generated in interface_generated.go
	ServicerGenerated
}
