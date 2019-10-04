package cmd

import (
	"bytes"
	"os"
	"strings"
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
	args := "appreg list-subscriptions --kind service"
	rootCmd.SetArgs(strings.Split(args, " "))
	err := rootCmd.Execute()
	assert.Nil(t, err)
}
