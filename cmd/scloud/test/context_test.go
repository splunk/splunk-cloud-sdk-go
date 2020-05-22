package test

import (
	"testing"

	utils "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestContextCmdWithoutSubCmd(t *testing.T) {
	command := "context"
	searchString := "Bearer"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t, false)
	assert.Equal(t, false, success)
}

func TestContextListCmd(t *testing.T) {
	command := "context list"
	searchString := "Bearer"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t, false)
	assert.Equal(t, true, success)
}
