package model

import "time"

type Ticker struct {
	duration time.Duration
	ticker   *time.Ticker
	running  bool
}

func (t *Ticker) Reset() {
	t.ticker = time.NewTicker(t.duration)
}

func (t *Ticker) Stop() {
	t.ticker.Stop()
	t.running = false
}

func (t *Ticker) Start() {
	t.Reset()
	t.running = true
}

func (t *Ticker) IsRunning() bool {
	return t.running == true
}

func CreateTicker(duration time.Duration) *Ticker {
	return &Ticker{duration: duration, ticker: time.NewTicker(duration), running: false}
}
