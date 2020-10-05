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

package action

// Servicer represents the interface for implementing all endpoints for this service
type Servicer interface {
	//interfaces that cannot be auto-generated from codegen
	/*
		TriggerActionWithStatus - Trigger an action and return a TriggerResponse with StatusID

		Parameters:
			actionName: The name of the action, as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			triggerEvent: The action payload, which must include values for any templated fields.
	*/
	TriggerActionWithStatus(actionName string, triggerEvent TriggerEvent) (*TriggerResponse, error)

	//interfaces that are auto-generated in interface_generated.go
	ServicerGenerated
}
