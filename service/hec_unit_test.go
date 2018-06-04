package service

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBuildMultiEventsPayload(t *testing.T) {
	var apiURLProtocol = "http"
	var apiHost = "example.com"
	var apiPort = "8882"
	var apiURL = apiURLProtocol + "://" + apiHost + ":" + apiPort
	var tenant = "EXAMPLE_TENANT"
	var token = "EXAMPLE_AUTHENTICATION_TOKEN"
	var timeout = time.Second * 5
	var client, _ = NewClient(tenant, token, apiURL, timeout)

	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	event3WithEmptyFields := model.HecEvent{Host: "", Event: "test3"}
	payload1, err := client.HecService.buildMultiEventsPayload([]model.HecEvent{event1, event2})
	assert.Nil(t, err)
	assert.Equal(t, `{"host":"host1","event":"test1"}{"host":"host2","event":"test2"}`, string(payload1[:]))
	payload2, err := client.HecService.buildMultiEventsPayload([]model.HecEvent{event1, event3WithEmptyFields})
	assert.Nil(t, err)
	assert.Equal(t, `{"host":"host1","event":"test1"}{"event":"test3"}`, string(payload2[:]))
}
