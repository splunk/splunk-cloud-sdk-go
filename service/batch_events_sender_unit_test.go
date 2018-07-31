package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var config = &Config{Token: "EXAMPLE_AUTHENTICATION_TOKEN", URL: "http://example.com", TenantID: "EXAMPLE_TENANT", Timeout: time.Second*5}

func TestBatchEventsSenderInitializationWithZeroBatchSizeAndZeroIntervalParameters(t *testing.T) {
	var client, _ = NewClient(config)
	//var client = getClient()
	_, err := client.NewBatchEventsSender(0, 0)
	assert.EqualError(t, err, "batchSize cannot be 0")
}

func TestBatchEventsSenderInitializationWithZeroBatchSize(t *testing.T) {
	var client, _ = NewClient(config)
	_, err := client.NewBatchEventsSender(0, 1000)
	assert.EqualError(t, err, "batchSize cannot be 0")
}

func TestBatchEventsSenderInitializationWithZeroInterval(t *testing.T) {
	var client, _ = NewClient(config)
	_, err := client.NewBatchEventsSender(5, 0)
	assert.EqualError(t, err, "interval cannot be 0")
}
