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

package httpx

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"github.com/splunk/splunk-cloud-sdk-go/v2/util"

	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strings"
)

func addrs(ipaddrs []net.IPAddr) string {
	count := len(ipaddrs)
	addrs := make([]string, count)
	for i, addr := range ipaddrs {
		addrs[i] = addr.String()
	}
	return strings.Join(addrs, ",")
}

func newClientTrace() *httptrace.ClientTrace {
	return &httptrace.ClientTrace{
		ConnectDone: func(network, addr string, err error) {
			glog.Infof("ConnectDone network=%s addr=%s error=%v", network, addr, err)
		},
		ConnectStart: func(network, addr string) {
			glog.Infof("ConnectStart network=%s addr=%s ..", network, addr)
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			glog.Infof("DNSDone addrs=[%s] coalesced=%t error=%v",
				addrs(info.Addrs), info.Coalesced, info.Err)
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			glog.Infof("DNSStart host=%s ..", info.Host)
		},
		GetConn: func(hostPort string) {
			glog.Infof("GetConn hostPort=%s", hostPort)
		},
		GotConn: func(info httptrace.GotConnInfo) {
			addr := info.Conn.RemoteAddr()
			glog.Infof("GotConn addr=%s reused=%t wasIdle=%t idleTime=%s",
				addr, info.Reused, info.WasIdle, info.IdleTime)
		},
		TLSHandshakeStart: func() {
			glog.Infof("TLSHandshakeStart ..")
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			glog.Infof("TLSHandshakeDone cipherSuite=%d resumed=%t complete=%t mutual=%t version=%d error=%v",
				state.CipherSuite,
				state.DidResume,
				state.HandshakeComplete,
				state.NegotiatedProtocolIsMutual,
				state.Version,
				err)
		}}
}

// GlogWrapper is used to wrap glog.info() in a Print() function usable by splunk-cloud-sdk-go
type GlogWrapper struct {
}

func (gw *GlogWrapper) Print(v ...interface{}) {
	glog.Info(v...)
}

func NewClient(insecure bool) *http.Client {
	httpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
		Proxy:           http.ProxyFromEnvironment,
	}
	transport := util.NewCustomSdkTransport(&GlogWrapper{}, httpTransport)
	return &http.Client{Transport: transport}
}

// Represents a SSC error response
type HTTPError struct {
	StatusCode int                    `json:"status,omitempty"`
	Code       string                 `json:"code,omitempty"`
	Message    string                 `json:"message,omitempty"`
	Details    map[string]interface{} `json:"details,omitempty"`
	MoreInfo   string                 `json:"moreInfo,omitempty"`
}

func (herr *HTTPError) Error() string {
	b, err := json.Marshal(herr)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// Returns a golang error corresponding to the given http response.
func Error(response *http.Response) error {
	result := &HTTPError{StatusCode: response.StatusCode}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result
	}
	if err = json.Unmarshal(body, result); err != nil {
		result.Message = string(body)
	}
	return result
}

// Encode the given value and return its reader.
func encode(value interface{}) (*strings.Reader, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(data)), nil
}

func Delete(url, token string, insecure bool) (*http.Response, error) {
	request, err := NewDelete(url, token)
	if err != nil {
		return nil, err
	}
	return NewClient(insecure).Do(request)
}

func Get(url, token string, params url.Values, insecure bool) (*http.Response, error) {
	request, err := NewGet(url, token, params)
	if err != nil {
		return nil, err
	}
	return NewClient(insecure).Do(request)
}

func GetText(url, token string, params url.Values, insecure bool) (*http.Response, error) {
	request, err := NewGet(url, token, params)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", "text/html")
	return NewClient(insecure).Do(request)
}

func Patch(url, token string, body interface{}, insecure bool) (*http.Response, error) {
	request, err := NewPatch(url, token, body)
	if err != nil {
		return nil, err
	}
	return NewClient(insecure).Do(request)
}

func Post(url, token string, body interface{}, insecure bool) (*http.Response, error) {
	request, err := NewPost(url, token, body)
	if err != nil {
		return nil, err
	}
	return NewClient(insecure).Do(request)
}

func Put(url, token string, body interface{}, insecure bool) (*http.Response, error) {
	request, err := NewPut(url, token, body)
	if err != nil {
		return nil, err
	}
	return NewClient(insecure).Do(request)
}

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if glog.V(2) {
		trace := newClientTrace()
		context := httptrace.WithClientTrace(request.Context(), trace)
		request = request.WithContext(context)
	}
	return request, nil
}

func NewDelete(url, token string) (*http.Request, error) {
	request, err := NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	bearer := fmt.Sprintf("Bearer %s", token)
	request.Header.Add("Authorization", bearer)
	request.Header.Add("Accept", "application/json")
	return request, nil
}

func NewGet(url, token string, params url.Values) (*http.Request, error) {
	request, err := NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bearer := fmt.Sprintf("Bearer %s", token)
	request.Header.Add("Authorization", bearer)
	request.Header.Add("Accept", "application/json")
	request.URL.RawQuery = params.Encode()
	return request, nil
}

// Create a new PATCH|POST|PUT request.
func newP(method, url, token string, body interface{}) (*http.Request, error) {
	reader, err := encode(body)
	if err != nil {
		return nil, err
	}
	request, err := NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	bearer := fmt.Sprintf("Bearer %s", token)
	request.Header.Add("Authorization", bearer)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	return request, nil
}

func NewPatch(url, token string, body interface{}) (*http.Request, error) {
	return newP("PATCH", url, token, body)
}

func NewPost(url, token string, body interface{}) (*http.Request, error) {
	return newP("POST", url, token, body)
}

func NewPut(url, token string, body interface{}) (*http.Request, error) {
	return newP("PUT", url, token, body)
}
