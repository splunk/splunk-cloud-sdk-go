package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const duration = time.Duration(1000) * time.Millisecond

func TestCreateTicker(t *testing.T) {
	ticker := CreateTicker(duration)
	assert.Equal(t, duration, ticker.duration)
	assert.NotNil(t, ticker.ticker)
}
