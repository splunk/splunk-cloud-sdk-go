package scloud

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/splunk/splunk-cloud-sdk-go/v2/sdk"
	"github.com/splunk/splunk-cloud-sdk-go/v2/services"
)

const UnitTestScheme = "https"
const UnitTestHost = "testing.splunk.com"
const UnitTestToken = "faketoken"
const UnitTestTenant = "faketenant"

func GetUnitTestClient() *sdk.Client {
	config := &services.Config{
		Token:        UnitTestToken,
		OverrideHost: UnitTestHost,
		Scheme:       UnitTestScheme,
		Tenant:       UnitTestTenant,
		RoundTripper: &UnitTestTransport{}}
	client, err := sdk.NewClient(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return client
}

// UnitTestError wraps an http request
type UnitTestError struct {
	Req http.Request
}

// This function is just here to implement the error interface
func (self *UnitTestError) Error() string {
	return "unused"
}

// UnitTestTransport is only used for GetUnitTestClient() to abort requests immediately
type UnitTestTransport struct {
	_ *http.Request // Unused, but needed to match the transport interface
}

func (self *UnitTestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Embed the request in an error object
	return nil, &UnitTestError{Req: *req}
}

// Test helpers

// The Go http.client do() function wraps any error from a roundtripper inside a url.Error under the Err field
func ErrorToRequest(t *testing.T, err error) http.Request {
	assert.IsType(t, &url.Error{}, err)
	uErr := err.(*url.Error)
	assert.IsType(t, &UnitTestError{}, (*uErr).Err)
	unitTestError := uErr.Err.(*UnitTestError)
	return unitTestError.Req
}

// Tries to read a http.Request.Body and return its contents as a string
func BodyToString(t *testing.T, requestBody io.ReadCloser) string {
	body, err := ioutil.ReadAll(requestBody)
	assert.Nil(t, err, "Unexpected error %s", err)
	return string(body)
}
