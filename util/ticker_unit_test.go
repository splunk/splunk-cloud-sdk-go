// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package util

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const duration = time.Duration(100) * time.Millisecond

func TestCreateTicker(t *testing.T) {
	ticker := NewTicker(duration)
	assert.Equal(t, duration, ticker.duration)
	assert.NotNil(t, ticker.ticker)
	assert.False(t, ticker.running)
}

func TestTickerReset(t *testing.T) {
	ticker := NewTicker(duration)
	ticker.Start()
	tickerStatusBefore := ticker.running
	time.Sleep(duration)
	assert.Equal(t, 1, len(ticker.GetChan()))
	ticker.Reset()
	tickerStatusAfter := ticker.running
	assert.Equal(t, 0, len(ticker.GetChan()))
	assert.Equal(t, tickerStatusBefore, tickerStatusAfter)
}

func TestTickerStop(t *testing.T) {
	ticker := NewTicker(duration)
	ticker.Start()
	time.Sleep(duration)
	assert.Equal(t, 1, len(ticker.GetChan()))
	<-ticker.GetChan()
	assert.Equal(t, 0, len(ticker.GetChan()))
	ticker.Stop()
	time.Sleep(duration)
	assert.Equal(t, 0, len(ticker.GetChan()))
	assert.False(t, ticker.running)
}

func TestTickerStart(t *testing.T) {
	ticker := NewTicker(duration)
	assert.False(t, ticker.running)
	ticker.Start()
	assert.True(t, ticker.running)
}

func TestIsRunning(t *testing.T) {
	ticker := NewTicker(duration)
	assert.False(t, ticker.IsRunning())
}

func TestTickerGetChan(t *testing.T) {
	ticker := NewTicker(duration)
	assert.Equal(t, reflect.Kind(reflect.Chan), reflect.TypeOf(ticker.GetChan()).Kind())
}
