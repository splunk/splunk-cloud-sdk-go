/*
 * Copyright © 2018 Splunk Inc.
 * SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
 * without a valid written license from Splunk Inc. is PROHIBITED.
 *
 */

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
	var client, _ = NewClient(&Config{token, apiURL, tenant, timeout})

	event1 := model.Event{Host: "host1", Event: "test1"}
	event2 := model.Event{Host: "host2", Event: "test2"}
	event3WithEmptyFields := model.Event{Host: "", Event: "test3"}
	payload1, err := client.IngestService.buildMultiEventsPayload([]model.Event{event1, event2})
	assert.Nil(t, err)
	assert.Equal(t, `{"host":"host1","event":"test1"}{"host":"host2","event":"test2"}`, string(payload1[:]))
	payload2, err := client.IngestService.buildMultiEventsPayload([]model.Event{event1, event3WithEmptyFields})
	assert.Nil(t, err)
	assert.Equal(t, `{"host":"host1","event":"test1"}{"event":"test3"}`, string(payload2[:]))
}
