package test_engine

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

var TEST_DEVICE_CODE = "deviceCode"
var TEST_USER_CODE = "userCode"
var TEST_VERIFICATION_URL = "https://auth.testing.splunk.com/verify"
var DEVICE_HANDLER_AUTHORIZED = true
var TOKEN_HANDLER_AUTHORIZED = true

// Create a moocked identity service and add endpoints handlers
func MockedIdentityServer() *httptest.Server {
	handler := http.NewServeMux()

	// Add Endpoints and endpoint handlers
	handler.HandleFunc(getEndpointWithTenant("device"), deviceHandler)
	handler.HandleFunc(getEndpointWithTenant("token"), tokenHandler)

	server := httptest.NewServer(handler)

	return server
}

// return /<tenantName>/<endpointName>
func getEndpointWithTenant(endpointName string) string {
	return fmt.Sprintf("/%v/%v", TestTenant, endpointName)
}

func deviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response string

	if DEVICE_HANDLER_AUTHORIZED {
		w.WriteHeader(http.StatusOK)
		response = fmt.Sprintf("{\"device_code\":\"%v\",\"expires_in\":600,\"interval\":5,\"user_code\":\"%v\",\"verification_uri\":\"%v\"}", TEST_DEVICE_CODE, TEST_USER_CODE, TEST_VERIFICATION_URL)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		response = "Unauthorized"
	}

	w.Write([]byte(response))
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response string

	if TOKEN_HANDLER_AUTHORIZED {
		w.WriteHeader(http.StatusOK)
		response = fmt.Sprintf("{\"device\": \"device\",\"scope\": \"offline_access email profile\",\"tenant\": \"%v\",\"user_code\": \"%v\"}", TestTenant, TEST_USER_CODE)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		response = "Unauthorized"
	}
	w.Write([]byte(response))
}
