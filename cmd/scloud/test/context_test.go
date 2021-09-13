package test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	utils "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/test/utils"
	"github.com/stretchr/testify/assert"
)

func TestContextCmdWithoutSubCmd(t *testing.T) {
	loginCmd := "login --use-pkce --tenant " + utils.TestTenant + " --env " + utils.Env1 + " --uid " + utils.Username + " --pwd " + utils.Password
	_, err, _ := utils.ExecuteCmd(loginCmd, t)
	assert.Equal(t, nil, err)

	command := "context"
	searchString := "Bearer"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t, false)
	assert.Equal(t, false, success)
}

func TestContextListCmd(t *testing.T) {
	loginCmd := "login --use-pkce --tenant " + utils.TestTenant + " --env " + utils.Env1 + " --uid " + utils.Username + " --pwd " + utils.Password
	_, err, _ := utils.ExecuteCmd(loginCmd, t)
	assert.Equal(t, nil, err)

	command := "context list"
	searchString := utils.TestTenant
	success := utils.Execute_cmd_with_global_flags(command, searchString, t, false)
	assert.Equal(t, true, success)
}

func TestContextListCmdWithTenant(t *testing.T) {
	loginCmd := "login --use-pkce --tenant " + utils.TestTenant + " --env " + utils.Env1 + " --uid " + utils.Username + " --pwd " + utils.Password
	_, err, _ := utils.ExecuteCmd(loginCmd, t)
	assert.Equal(t, nil, err)

	command := "context list --tenant " + utils.TestTenant
	searchString := "Bearer"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t, false)
	assert.Equal(t, true, success)
}

func TestContextSetCmd(t *testing.T) {
	setUp(t)
	setCommand := "context set --key access_token --value abc123 --tenant " + utils.TestTenant
	results, _, std := utils.ExecuteCmd(setCommand, t)

	assert.Equal(t, "", results)
	assert.Equal(t, "", std)

	listCommand := "context list"
	res, _, std2 := utils.ExecuteCmd(listCommand, t)

	assert.Equal(t, true, strings.Contains(res, "43200"))
	assert.Equal(t, true, strings.Contains(res, "abc123"))
	assert.Equal(t, true, strings.Contains(res, "Bearer"))
	assert.Equal(t, true, strings.Contains(res, "offline_access openid email profile"))
	assert.Equal(t, "", std2)

	tearDown(t)
}

func TestContextSetCmdWithInvalidKey(t *testing.T) {
	setUp(t)
	setCommand := "context set --key accesstoken --value abc123 --tenant " + utils.TestTenant
	results, _, std := utils.ExecuteCmd(setCommand, t)

	assert.Equal(t, "", results)
	assert.Equal(t, true, strings.Contains(std, "Here are the keys you can set"))

	listCommand := "context list --tenant " + utils.TestTenant
	res, _, std2 := utils.ExecuteCmd(listCommand, t)

	assert.Equal(t, false, strings.Contains(res, "43200"))
	assert.Equal(t, false, strings.Contains(res, "abc123"))
	assert.Equal(t, false, strings.Contains(res, "Bearer"))
	assert.Equal(t, false, strings.Contains(res, "offline_access openid email profile"))
	assert.Equal(t, "", std2)

	tearDown(t)
}

func TestContextSetCmdWithoutTenant(t *testing.T) {
	setUp(t)
	setCommand := "context set --key access_token --value abc123"
	results, _, _ := utils.ExecuteCmd(setCommand, t)

	assert.Equal(t, "required flag(s) \"tenant\" not set\n", results)

	tearDown(t)
}

// GetFilename uses reflection to get current filename
func GetFilename() string {
	_, filename, _, _ := runtime.Caller(0)
	return filename
}

func setUp(t *testing.T) {
	envPath := filepath.Join(filepath.Dir(GetFilename()), "fixtures", ".test.context.env")
	overloadErr := godotenv.Overload(envPath)
	if overloadErr != nil {
		message := fmt.Sprintf("Failed to overload %s. Error: %s", envPath, overloadErr)
		t.Log(message)
	}
}

func tearDown(t *testing.T) {
	// remove the test cache file
	scloudTestCachePath := auth.Abspath(os.Getenv("SCLOUD_CACHE_PATH"))
	removeErr := os.Remove(scloudTestCachePath)
	if removeErr != nil {
		message := fmt.Sprintf("Failed to remove %s. Error: %s", scloudTestCachePath, removeErr)
		t.Log(message)
	}

	// override SCLOUD_CACHE_PATH
	envPath := filepath.Join(filepath.Dir(GetFilename()), "..", "..", "..", ".env")
	overloadErr := godotenv.Overload(envPath)
	if overloadErr != nil {
		message := fmt.Sprintf("Failed to overload %s. Error: %s", envPath, overloadErr)
		t.Log(message)
	}
}
