// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model

import (
	"time"
)

// Ticker is a wrapper of time.Ticker with additional functionality
type Ticker struct {
	duration time.Duration
	ticker   *time.Ticker
	running  bool
}

// Reset resets ticker
func (t *Ticker) Reset() {
	t.ticker.Stop()
	t.ticker = time.NewTicker(t.duration)
}

// Stop stops ticker and set property running to false
func (t *Ticker) Stop() {
	t.ticker.Stop()
	t.running = false
}

// Start starts a new ticker and set property running to true
func (t *Ticker) Start() {
	t.Reset()
	t.running = true
}

// IsRunning returns bool indicating whether or not ticker is running
func (t *Ticker) IsRunning() bool {
	return t.running == true
}

// GetChan returns the channel from ticker
func (t *Ticker) GetChan() <-chan time.Time {
	return t.ticker.C
}

// NewTicker spits out a pointer to Ticker model. It sets ticker to stop state by default
func NewTicker(duration time.Duration) *Ticker {
	newTicker := time.NewTicker(duration)
	newTicker.Stop()
	return &Ticker{duration: duration, ticker: newTicker, running: false}
}
