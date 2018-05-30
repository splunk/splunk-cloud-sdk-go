package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBatchEventsSenderInitializationWithZeroBatchSizeAndZeroIntervalParameters(t *testing.T) {
	var client = NewClient("EXAMPLE_TENANT", "EXAMPLE_AUTHENTICATION_TOKEN", "http://example.com", time.Second*5)
	//var client = getClient()
	_, err := client.NewBatchEventsSender(0, 0)
	assert.EqualError(t, err, "batchSize cannot be 0")
}

func TestBatchEventsSenderInitializationWithZeroBatchSize(t *testing.T) {
	var client = NewClient("EXAMPLE_TENANT", "EXAMPLE_AUTHENTICATION_TOKEN", "http://example.com", time.Second*5)
	_, err := client.NewBatchEventsSender(0, 1000)
	assert.EqualError(t, err, "batchSize cannot be 0")
}

func TestBatchEventsSenderInitializationWithZeroInterval(t *testing.T) {
	var client = NewClient("EXAMPLE_TENANT", "EXAMPLE_AUTHENTICATION_TOKEN", "http://example.com", time.Second*5)
	_, err := client.NewBatchEventsSender(5, 0)
	assert.EqualError(t, err, "interval cannot be 0")
}