/*
Package service implements a service client which is used to communicate
with Search Service endpoints
*/
package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/splunk/ssc-client-go/lib/util"
)

// Declare constants for service package
const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodPatch  = "PATCH"
	MethodDelete = "DELETE"
)

// Request parameter format types
const (
	JSON       = "JSON"
	URLEncoded = "URLEncoded"
)

// A Client is used to communicate with service endpoints
type Client struct {
	// Splunk session key
	SessionKey string
	// Basic Auth with username and password
	Auth [2]string
	// Host name
	Host string
	// HTTP Client used to interact with endpoints
	httpClient *http.Client
	// Services designed to talk to different parts of Splunk
	SearchService *SearchService
	//Scheme or http protocol
	Scheme string
}

// service provides the interface between client and services
type service struct {
	client *Client
}

// NewRequest creates a new HTTP Request and set proper header
func (c *Client) NewRequest(httpMethod, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}
	if c.SessionKey != "" {
		request.Header.Add("Authorization", "Splunk "+c.SessionKey)
	} else {
		request.SetBasicAuth(c.Auth[0], c.Auth[1])
	}
	request.Header.Set("Content-Type", "application/json")
	return request, nil
}

// BuildURL creates full Splunk URL
func (c *Client) BuildURL(queryValues url.Values, urlPathParts ...string) url.URL {
	buildPath := ""
	for _, pathPart := range urlPathParts {
		buildPath = path.Join(buildPath, url.PathEscape(pathPart))
	}
	if queryValues == nil {
		queryValues = url.Values{}
	}
	// Always set json as output format for now
	queryValues.Set("output_mode", "json")
	return url.URL{
		Scheme:   c.Scheme,
		Host:     c.Host,
		Path:     buildPath,
		RawQuery: queryValues.Encode(),
	}
}

// Do sends out request and returns HTTP response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Get implements HTTP Get call
func (c *Client) Get(getURL url.URL, format string) (*http.Response, error) {
	return c.DoRequest(MethodGet, getURL, nil, format)
}

// Post implements HTTP POST call
func (c *Client) Post(postURL url.URL, body interface{}, format string) (*http.Response, error) {
	return c.DoRequest(MethodPost, postURL, body, format)
}

// Put implements HTTP PUT call
func (c *Client) Put(putURL url.URL, body interface{}, format string) (*http.Response, error) {
	return c.DoRequest(MethodPut, putURL, body, format)
}

// Delete implements HTTP DELETE call
func (c *Client) Delete(deleteURL url.URL, format string) (*http.Response, error) {
	return c.DoRequest(MethodDelete, deleteURL, nil, format)
}

// Patch implements HTTP Patch call
func (c *Client) Patch(patchURL url.URL, body interface{}, format string) (*http.Response, error) {
	return c.DoRequest(MethodPatch, patchURL, body, format)
}

// DoRequest creates and execute a new request
func (c *Client) DoRequest(method string, requestURL url.URL, body interface{}, format string) (*http.Response, error) {

	// default to JSON
	content, err := c.toJSON(body)

	if format == URLEncoded {
		content, err = c.EncodeRequestBody(body)
	}
	if err != nil {
		return nil, err
	}
	request, err := c.NewRequest(method, requestURL.String(), bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}
	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	return util.ParseHTTPStatusCodeInResponse(response)
}

// EncodeRequestBody takes a json string or object and serializes it to be used in request body
func (c *Client) EncodeRequestBody(content interface{}) ([]byte, error) {
	if content != nil {
		switch value := reflect.ValueOf(content); value.Kind() {
		case reflect.String:
			return []byte(value.String()), nil
		case reflect.Map:
			return c.EncodeObject(value.Interface())
		case reflect.Struct:
			return c.EncodeObject(value.Interface())
		default:
			return nil, &util.HTTPError{Status: 400, Message: "Bad Request"}
		}
	}
	return nil, nil
}

// toJSON takes an object and attempts to convert to a JSON string
func (c *Client) toJSON(data interface{}) ([]byte, error) {
	marshalContent, err := json.Marshal(data)
	return []byte(marshalContent), err
}

// EncodeObject encodes an object into url-encoded string
func (c *Client) EncodeObject(content interface{}) ([]byte, error) {
	URLValues := url.Values{}
	marshalContent, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}
	var valueMap map[string]interface{}
	if err := json.Unmarshal(marshalContent, &valueMap); err != nil {
		return nil, err
	}
	for k, v := range valueMap {
		k = strings.ToLower(k)
		switch val := v.(type) {
		case string:
			URLValues.Set(k, val)
		case bool:
			URLValues.Set(k, strconv.FormatBool(val))
		case int:
			URLValues.Set(k, strconv.FormatInt(int64(val), 10))
		case float32:
			URLValues.Set(k, strconv.FormatFloat(float64(val), 'f', -1, 32))
		case float64:
			URLValues.Set(k, strconv.FormatFloat(float64(val), 'f', -1, 64))
		}
	}
	return []byte(URLValues.Encode()), nil
}

// NewClient creates a Client with custom values passed in
func NewClient(sessionKey string, auth [2]string, host string, scheme string, timeout time.Duration, skipValidateTLS bool) *Client {
	httpClient := newHTTPClient(timeout, skipValidateTLS)
	c := &Client{Auth: auth, Host: host, Scheme: scheme, httpClient: httpClient}

	// TODO(dan): this is here for backward compat, will circle back and refactor after demo.
	if sessionKey != "" {
		c.SessionKey = sessionKey
	}
	// TODO(dan): need to ask Eric why we did this, looks circular
	c.SearchService = &SearchService{client: c}
	return c
}

// NewHTTPClient returns a HTTP Client with timeout and tls validation setup
func newHTTPClient(timeout time.Duration, skipValidateTLS bool) *http.Client {
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: skipValidateTLS},
		},
	}
}
