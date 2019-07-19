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
	"io/ioutil"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

func GetSpecJSON(svcCluster, svcVersion, svcName string, client *services.BaseClient) (interface{}, error) {
	var result interface{}
	url, err := client.BuildURLWithTenant("system", nil, svcCluster, svcName, "specs", svcVersion, "openapi.json")
	if err != nil {
		return result, err
	}
	response, err := client.Get(services.RequestParams{URL: url})
	if err != nil {
		return nil, err
	}
	err = util.ParseResponse(&result, response)
	return result, err
}

func GetSpecYaml(svcCluster, svcVersion string, svcName string, client *services.BaseClient) (string, error) {
	url, err := client.BuildURLWithTenant("system", nil, svcCluster, svcName, "specs", svcVersion, "openapi.yaml")
	if err != nil {
		return "", err
	}
	response, err := client.Get(services.RequestParams{URL: url})
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
