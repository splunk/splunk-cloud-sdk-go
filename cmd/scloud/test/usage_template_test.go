package test

import (
	"strings"
	"testing"

	utils "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestServiceCmdShouldReturnShortUsage(t *testing.T) {
	command := "action"
	_, err, std := utils.ExecuteCmd(command, t)

	assert.Equal(t, nil, err)

	isShortUsage := !strings.Contains(std, "Global Flags:")
	assert.Equal(t, true, isShortUsage)
}

func TestRootCmdShouldReturnShortUsage(t *testing.T) {
	command := ""
	_, err, std := utils.ExecuteCmd(command, t)

	assert.Equal(t, nil, err)

	isShortUsage := !strings.Contains(std, "Global Flags:")
	assert.Equal(t, true, isShortUsage)
}

func TestHelpFlagShouldLongUsage(t *testing.T) {
	command := "action --help"
	res, err, _ := utils.ExecuteCmd(command, t)

	assert.Equal(t, nil, err)

	isLongUsage := strings.Contains(res, "Global Flags:")
	assert.Equal(t, true, isLongUsage)
}
