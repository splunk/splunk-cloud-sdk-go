/*
 * Copyright 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

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
