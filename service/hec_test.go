package service

import (
	"testing"
	"time"

	"github.com/splunk/ssc-client-go/model"
	. "github.com/splunk/ssc-client-go/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateEventSuccess(t *testing.T) {
	err := getSplunkClient().HecService.CreateEvent(
		model.HecEvent{Host: "http://ssc-sdk-shared-stubby:8882", Event: "test", Source: "manual-events", Sourcetype: "sourcetype:eventgen", Time: 1523637597})
	assert.Empty(t, err)
}

func TestCreateEventFail(t *testing.T) {
	client := NewClient("", TestStubbySchme+"://"+TestStubbyHost, time.Second*5, true)
	err := client.HecService.CreateEvent(
		model.HecEvent{Host: "http://ssc-sdk-shared-stubby:8882", Event: "", Source: "manual-events", Sourcetype: "sourcetype:eventgen", Time: 1523637597})
	assert.Empty(t, err)
}

func TestCreateRawEventSuccess(t *testing.T) {
	client := getSplunkClient()
	var url = client.BuildURL(hecServicePrefix, hecServiceVersion, "raw")
	fields := map[string]string{
		"test": "value",
	}
	event := model.HecEvent{Host: "http://ssc-sdk-shared-stubby:8882", Event: "test", Source: "manual-events", Sourcetype: "sourcetype:eventgen", Time: 1523637597, Fields: fields}
	url.RawQuery = ParseURLParams(event).Encode()
	t.Errorf("%v", url.String())
}
