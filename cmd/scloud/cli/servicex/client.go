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

package servicex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/cli/httpx"
)

const (
	defaultPort   = "443"
	defaultScheme = "https"
)

type Client struct {
	Token    string // access token
	Host     string
	Port     string
	Scheme   string
	Insecure bool
}

func (clt *Client) url(path, tenant string, args ...interface{}) string {
	host := clt.Host
	port := clt.Port
	if port == "" {
		port = defaultPort
	}
	scheme := clt.Scheme
	if scheme == "" {
		scheme = defaultScheme
	}
	if len(args) > 0 {
		path = fmt.Sprintf(path, args...)
	}
	return fmt.Sprintf("%s://%s:%s/%s%s", scheme, host, port, tenant, path)
}

// Quiet
func decode(response *http.Response, result interface{}) error {
	return json.NewDecoder(response.Body).Decode(&result)
}

// Verbose
func decodeV(response *http.Response, result interface{}) error { //nolint:deadcode
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}
	return nil
}

func (clt *Client) delete(path, tenant string, args ...interface{}) error {
	return clt.deletex(path, tenant, http.StatusNoContent, args...)
}

func (clt *Client) deletex(path, tenant string, statusCode int, args ...interface{}) error {
	url := clt.url(path, tenant, args...)
	response, err := httpx.Delete(url, clt.Token, clt.Insecure)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != statusCode {
		return httpx.Error(response)
	}
	return nil
}

func (clt *Client) get(url string, result interface{}) error {
	return clt.getWithParams(url, nil, result)
}

func (clt *Client) getWithParams(url string, params url.Values, result interface{}) error {
	response, err := httpx.Get(url, clt.Token, params, clt.Insecure)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return httpx.Error(response)
	}
	if err = decode(response, &result); err != nil {
		return err
	}
	return nil
}

func (clt *Client) getText(url string) (string, error) { //nolint:deadcode
	response, err := httpx.GetText(url, clt.Token, nil, clt.Insecure)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", httpx.Error(response)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (clt *Client) list(path, tenant string, args ...interface{}) ([]string, error) {
	url := clt.url(path, tenant, args...)
	var result []string
	if err := clt.get(url, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (clt *Client) patch(url string, body, result interface{}) error { //nolint:deadcode
	response, err := httpx.Patch(url, clt.Token, body, clt.Insecure)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return httpx.Error(response)
	}
	if result == nil {
		return nil
	}
	return decode(response, &result)
}

func (clt *Client) postx(url string, body, result interface{}, statusCode int) error {
	response, err := httpx.Post(url, clt.Token, body, clt.Insecure)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != statusCode {
		return httpx.Error(response)
	}
	if result == nil {
		return nil
	}
	return decode(response, &result)
}

func (clt *Client) post(url string, body, result interface{}) error {
	return clt.postx(url, body, result, http.StatusCreated)
}

func (clt *Client) postOk(url string, body, result interface{}) error {
	return clt.postx(url, body, result, http.StatusOK)
}

func (clt *Client) put(url string, body, result interface{}) error {
	response, err := httpx.Put(url, clt.Token, body, clt.Insecure)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return httpx.Error(response)
	}
	if result == nil {
		return nil
	}
	return decode(response, &result)
}
