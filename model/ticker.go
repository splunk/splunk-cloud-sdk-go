package model

import "time"

type Ticker struct {
	duration time.Duration
	ticker *time.Ticker
}

func (t *Ticker) Reset() {
	t.ticker = time.NewTicker(t.duration)
}

func (t *Ticker) Stop() {
	t.ticker.Stop()
}

func CreateTicker(duration time.Duration) *Ticker {
	return &Ticker{duration: duration, ticker: time.NewTicker(duration)}
}