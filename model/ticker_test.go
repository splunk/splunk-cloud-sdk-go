package model

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const duration = time.Duration(100) * time.Millisecond

func TestCreateTicker(t *testing.T) {
	ticker := CreateTicker(duration)
	assert.Equal(t, duration, ticker.duration)
	assert.NotNil(t, ticker.ticker)
	assert.False(t, ticker.running)
}

func TestTickerReset(t *testing.T) {
	ticker := CreateTicker(duration)
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
	ticker := CreateTicker(duration)
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
	ticker := CreateTicker(duration)
	assert.False(t, ticker.running)
	ticker.Start()
	assert.True(t, ticker.running)
}

func TestIsRunning(t *testing.T) {
	ticker := CreateTicker(duration)
	assert.False(t, ticker.IsRunning())
}

func TestTickerGetChan(t *testing.T) {
	ticker := CreateTicker(duration)
	assert.Equal(t, reflect.Kind(reflect.Chan), reflect.TypeOf(ticker.GetChan()).Kind())
}
