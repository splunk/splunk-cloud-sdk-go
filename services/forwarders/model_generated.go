/*
 * Copyright © 2020 Splunk, Inc.
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
 *
 * Splunk Forwarder Service
 *
 * Send data from a Splunk forwarder to the Splunk Forwarder service in Splunk Cloud Services.
 *
 * API version: v2beta1.1 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package forwarders

type Certificate struct {
	Pem *string `json:"pem,omitempty"`
}

type CertificateInfo struct {
	Content    *string `json:"content,omitempty"`
	Hash       *string `json:"hash,omitempty"`
	Issuer     *string `json:"issuer,omitempty"`
	LastUpdate *string `json:"lastUpdate,omitempty"`
	NotAfter   *string `json:"notAfter,omitempty"`
	NotBefore  *string `json:"notBefore,omitempty"`
	Slot       *int64  `json:"slot,omitempty"`
	Subject    *string `json:"subject,omitempty"`
}

type Error struct {
	Code    *string                `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Message *string                `json:"message,omitempty"`
}
