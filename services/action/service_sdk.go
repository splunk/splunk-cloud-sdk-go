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

// This files contains models that can't be auto-generated from codegen
package action

import (
	"fmt"
	"net/http"

	"github.com/splunk/splunk-cloud-sdk-go/v2/util"
)

/*
	TriggerActionWithStatus - Trigger an action and return a TriggerResponse with StatusID

	Parameters:
		actionName: The name of the action, as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
		triggerEvent: The action payload, which must include values for any templated fields.
*/
func (s *Service) TriggerActionWithStatus(actionName string, triggerEvent TriggerEvent) (*TriggerResponse, error) {
	var resp http.Response
	err := s.TriggerAction(actionName, triggerEvent, &resp)
	if err != nil {
		return nil, err
	}
	u, err := resp.Location()
	if err != nil {
		return nil, err
	}
	template := fmt.Sprintf(`/{tenant}/action/v1beta2/actions/{action_name}/status/{status_id}`)
	params, err := util.ParseTemplatedPath(template, u.Path)
	if err == nil {
		id := params["status_id"]
		return &TriggerResponse{StatusURL: u, StatusID: &id}, nil
	}
	// If format doesn't match what we expect just return url
	return &TriggerResponse{StatusURL: u}, nil
}
