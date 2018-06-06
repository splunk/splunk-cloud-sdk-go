package stubbyintegration

import (
	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
	"github.com/splunk/ssc-client-go/testutils"
	"github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateEventSuccess(t *testing.T) {
	timeValue := float64(1523637597)
	err := getClient(t).HecService.CreateEvent(
		model.HecEvent{Host: getClient(t).URL.RequestURI(), Index: "main", Event: "test", Sourcetype: "sourcetype:eventgen", Source: "manual-events", Time: &timeValue, Fields: map[string]string{"testKey": "testValue"}})
	assert.Empty(t, err)
}

func TestCreateRawEventSuccess(t *testing.T) {
	err := getClient(t).HecService.CreateRawEvent(
		model.HecEvent{Event: "test"})
	assert.Empty(t, err)
}

func TestHecEventFail(t *testing.T) {
	client, _ := service.NewClient(testutils.TestTenantID, "wrongToken", testutils.TestURLProtocol+"://"+testutils.TestSSCHost, time.Second*5)
	err := client.HecService.CreateEvent(model.HecEvent{Event: "failed test"})
	assert.NotEmpty(t, err)
	assert.Equal(t, 401, err.(*util.HTTPError).Status)
	assert.Equal(t, "401 Unauthorized", err.(*util.HTTPError).Message)
}

func TestCreateEvents(t *testing.T) {
	event1 := model.HecEvent{Host: "host1", Event: "test1"}
	event2 := model.HecEvent{Host: "host2", Event: "test2"}
	err := getClient(t).HecService.CreateEvents([]model.HecEvent{event1, event2})
	assert.Empty(t, err)
}
