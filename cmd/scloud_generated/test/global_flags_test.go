package test

import (
	"os"
	"testing"

	utils "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/test/utils"
	"github.com/stretchr/testify/assert"
)

func Test_global_flag_host_url(t *testing.T) {
	command := "provisioner list-tenants --host-url " + utils.TestSplunkCloudHost
	searchString := "createdAt"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t)

	if !success {
		assert.Equal(t, success, true)
	}
}

func Test_global_flag_invalid_host_url(t *testing.T) {
	command := "provisioner list-tenants --host-url " + utils.InvalidHostUrl
	searchString := "no such host"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t)

	if !success {
		assert.Equal(t, success, true)
	}
}

func Test_global_flag_auth_url(t *testing.T) {
	password := "\"" + utils.Password + "\""
	command := "login --pwd " + password + " --auth-url " + utils.IdpHost
	searchString := ""
	success := utils.Execute_cmd_with_global_flags(command, searchString, t)

	if !success {
		assert.Equal(t, success, true)
	}

	command = "provisioner list-tenants --host-url " + utils.TestSplunkCloudHost
	searchString = "createdAt"
	success = utils.Execute_cmd_with_global_flags(command, searchString, t)

	if !success {
		assert.Equal(t, success, true)
	}
}

func Test_global_flag_invalid_auth_url(t *testing.T) {
	password := "\"" + utils.Password + "\""
	command := "login --pwd " + password + " --auth-url " + utils.InvalidAuthUrl
	searchString := "failed to get valid response from csrfToken endpoint"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t)

	if !success {
		assert.Equal(t, success, true)
	}
}

func Test_global_flag_tenant(t *testing.T) {
	command := "action list-actions --tenant " + utils.TestTenant
	searchString := "kind"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t)

	if !success {
		assert.Equal(t, success, true)
	}
}

func Test_global_flag_invalid_tenant(t *testing.T) {
	command := "action list-actions --tenant invalidTenant"
	searchString := "\"HTTPStatusCode\":404,\"HTTPStatus\":\"404 Not Found\",\"code\":\"path-not-found\""
	success := utils.Execute_cmd_with_global_flags(command, searchString, t)

	if !success {
		assert.Equal(t, success, true)
	}
}

func Test_global_flag_env(t *testing.T) {
	command := "action list-actions --env " + utils.Env1
	searchString := "kind"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t)

	if !success {
		assert.Equal(t, success, true)
	}
}

func Test_global_flag_invalid_env(t *testing.T) {
	command := "action list-actions --env invalidEnv"
	searchString := "Environment specified does not exist"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t)

	if !success {
		assert.Equal(t, success, true)
	}
}

func Test_global_flag_incorrect_env(t *testing.T) {
	command := "action list-actions --env " + utils.Env2
	searchString := "\"HTTPStatusCode\":401,\"HTTPStatus\":\"401 Unauthorized\",\"code\":\"unauthorized\""
	success := utils.Execute_cmd_with_global_flags(command, searchString, t)

	if !success {
		assert.Equal(t, success, true)
	}
}

func Test_global_flag_insecure_enabled(t *testing.T) {
	command := "provisioner list-tenants --insecure true"
	searchString := "createdAt"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t)

	assert.Equal(t, success, true)
}

func Test_global_flag_insecure_cacert(t *testing.T) {
	currDir, err := os.Getwd()
	if err != nil {
		assert.Fail(t, "Cannot get current directory")
	}
	fileName := currDir + "/testcases/certfile.crt"
	command := "provisioner list-tenants --insecure false --ca-cert " + fileName
	searchString := "createdAt"
	success := utils.Execute_cmd_with_global_flags(command, searchString, t)

	assert.Equal(t, success, true)

}
