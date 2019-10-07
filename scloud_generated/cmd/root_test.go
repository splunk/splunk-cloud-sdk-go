package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Setup
	// Now run all of our tests
	os.Exit(m.Run())
}

func TestCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOutput(buf)
	rootCmd.SetArgs([]string{
		"help",
	})
	err := rootCmd.Execute()
	assert.Nil(t, err)
}
