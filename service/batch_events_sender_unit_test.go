// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var config = &Config{Token: "EXAMPLE_AUTHENTICATION_TOKEN"}

func TestBatchEventsSenderInitializationWithZeroBatchSizeAndZeroIntervalParameters(t *testing.T) {
	var client, _ = NewClient(config)
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
