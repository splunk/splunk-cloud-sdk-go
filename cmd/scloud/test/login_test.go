package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	utils "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/test/utils"
	"github.com/splunk/splunk-cloud-sdk-go/idp"
	"github.com/stretchr/testify/assert"
)

//login without username specified should return prompt asking for input of username
func TestLoginWithNoUsername(t *testing.T) {
	tearDown(t)

	//set username to be empty
	configCommand := "config set --key \"username\" --value \"\""
	result1, err, std := utils.ExecuteCmd(configCommand, t)
	fmt.Println("===================" + configCommand)
	fmt.Println(result1)
	fmt.Println(err)
	fmt.Println(std)

	configCommand = "config list "
	result1, err, std = utils.ExecuteCmd(configCommand, t)
	fmt.Println("===================" + configCommand)
	fmt.Println(result1)
	fmt.Println(err)
	fmt.Println(std)

	configCommand = "config reset"
	result1, err, std = utils.ExecuteCmd(configCommand, t)
	fmt.Println("===================" + configCommand)
	fmt.Println(result1)
	fmt.Println(err)
	fmt.Println(std)

	configCommand = "config list "
	result1, err, std = utils.ExecuteCmd(configCommand, t)
	fmt.Println("===================" + configCommand)
	fmt.Println(result1)
	fmt.Println(err)
	fmt.Println(std)

	configCommand = "config get --key \"username\""
	result1, err, std = utils.ExecuteCmd(configCommand, t)
	fmt.Println("===================" + configCommand)
	fmt.Println(result1)
	fmt.Println(err)
	fmt.Println(std)

	configCommand = "config list "
	result1, err, std = utils.ExecuteCmd(configCommand, t)
	fmt.Println("===================" + configCommand)
	fmt.Println(result1)
	fmt.Println(err)
	fmt.Println(std)

	configCommand = "context list "
	result1, err, std = utils.ExecuteCmd(configCommand, t)
	fmt.Println("===================" + configCommand)
	fmt.Println(result1)
	fmt.Println(err)
	fmt.Println(std)

	//execute a command, and this should prompt to input username
	configCommand = "action list-actions"
	results, err, std := utils.ExecuteCmd(configCommand, t)
	fmt.Println("===================" + configCommand)
	fmt.Println(result1)
	fmt.Println(err)
	fmt.Println(std)
	assert.Equal(t, "Username: ", results)
}

func TestRefreshLoginFlow(t *testing.T) {
	loginCommand := "login  --pwd " + utils.Password + " --use-refresh-token"
	results, err, std := utils.ExecuteCmd(loginCommand, t)

	assert.Equal(t, "", results)
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)

	identityCommand := "identity list-groups"
	identityResults, err, std := utils.ExecuteCmd(identityCommand, t)
	assert.NotEmpty(t, identityResults)
	assert.False(t, strings.Contains(std, "401 Unauthorized"))
	assert.False(t, strings.Contains(std, "404 Not Found"))
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)
}

func TestKCLoginFlow(t *testing.T) {
	//set config of the user enabled with Keycloak
	setConfig("username", utils.Username2, t)
	setConfig("env", utils.Env3, t)
	setConfig("tenant", utils.TestTenant2, t)

	//validate login
	loginCommand := "login  --pwd " + utils.Password2
	results, err, std := utils.ExecuteCmd(loginCommand, t)
	assert.Equal(t, "", results)
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)

	//reset config settings
	setConfig("username", utils.Username, t)
	setConfig("env", utils.Env1, t)
	setConfig("tenant", utils.TestTenant, t)

	//reset login
	loginCommand = "login  --pwd " + utils.Password
	results, err, std = utils.ExecuteCmd(loginCommand, t)

	assert.Equal(t, "", results)
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)
}

func TestDefaultLoginFlow(t *testing.T) {
	loginCommand := "login --pwd " + utils.Password
	results, err, std := utils.ExecuteCmd(loginCommand, t)

	assert.Equal(t, "", results)
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)

	identityCommand := "identity list-groups"
	identityResults, err, std := utils.ExecuteCmd(identityCommand, t)

	assert.NotEmpty(t, identityResults)
	assert.False(t, strings.Contains(std, "401 Unauthorized"))
	assert.False(t, strings.Contains(std, "404 Not Found"))
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)
}

func TestRefreshLoginFlowWithVerbose(t *testing.T) {
	loginCommand := "login  --pwd " + utils.Password + " --use-refresh-token --verbose"
	results, err, std := utils.ExecuteCmd(loginCommand, t)

	assert.True(t, strings.Contains(results, "access_token"))
	assert.True(t, strings.Contains(results, "token_type"))
	assert.True(t, strings.Contains(results, "scope"))
	assert.True(t, strings.Contains(results, "expires_in"))
	assert.True(t, strings.Contains(results, "refresh_token"))
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)
}

func TestDefaultLoginFlowWithVerbose(t *testing.T) {
	command := "login --pwd " + utils.Password + " --verbose"
	results, err, std := utils.ExecuteCmd(command, t)

	assert.True(t, strings.Contains(results, "access_token"))
	assert.True(t, strings.Contains(results, "token_type"))
	assert.True(t, strings.Contains(results, "scope"))
	assert.True(t, strings.Contains(results, "expires_in"))
	assert.True(t, strings.Contains(results, "refresh_token"))
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)
}

func TestLoginWithUidShouldNotRequireUsername(t *testing.T) {
	command := "login --uid " + utils.Username + " --pwd " + utils.Password + " --verbose"
	searchString := "username"
	resultsContainSearchString := utils.Execute_cmd_with_global_flags(command, searchString, t, false)
	assert.Equal(t, false, resultsContainSearchString)
}

func TestAuthLogin(t *testing.T) {
	setUp(t)
	auth.LoginSetUp()
	context, err := auth.Login(nil, mockedAuthFlow)

	assert.Nil(t, err)
	assert.Equal(t, "TestingAccessToken", context.AccessToken)
	assert.Equal(t, 0, context.ExpiresIn)
	assert.Equal(t, "TestingTokenType", context.TokenType)
	assert.Equal(t, "TestingIDToken", context.IDToken)
	assert.Equal(t, "TestingRefreshToken", context.RefreshToken)
	assert.Equal(t, "TestingScope", context.Scope)
	assert.Equal(t, int64(0), context.StartTime)

	tearDown(t)
}

func setConfig(key string, value string, t *testing.T) {
	configCommand := "config set --key " + key + " --value " + strings.TrimSpace(value)
	results, err, std := utils.ExecuteCmd(configCommand, t)
	assert.Equal(t, "", results)
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)
}

func mockedAuthFlow(profile map[string]string, cmd *cobra.Command) (*idp.Context, error) {
	return &idp.Context{
		AccessToken:  "TestingAccessToken",
		ExpiresIn:    0,
		TokenType:    "TestingTokenType",
		Scope:        "TestingScope",
		StartTime:    0,
		RefreshToken: "TestingRefreshToken",
		IDToken:      "TestingIDToken",
	}, nil
}

func TestDeviceFlow(t *testing.T) {
	// Step 1: Setup test env file + mock server
	setUp(t)
	server := utils.MockedIdentityServer()
	defer server.Close()

	// Step 2: Override auth-url and execute device flow login
	loginCommand := "login  --use-device --auth-url " + server.URL + " --tenant " + utils.TestTenant
	results, err, std := utils.ExecuteCmd(loginCommand, t)

	// Step 3: Assert results
	assert.True(t, strings.Contains(results, "Verification URL: "+utils.TEST_VERIFICATION_URL))
	assert.True(t, strings.Contains(results, "User Code: "+utils.TEST_USER_CODE))
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)

	// Step 4 Tear down testing environment
	tearDown(t)
}

func TestDeviceFlowDeviceEndpointError(t *testing.T) {
	// Step 1: Setup test env file + mock server
	setUp(t)
	server := utils.MockedIdentityServer()
	defer server.Close()

	// Step 2: Set authorized to be false
	utils.DEVICE_HANDLER_AUTHORIZED = false
	utils.TOKEN_HANDLER_AUTHORIZED = true

	// Step 3: Override auth-url and execute device flow login
	loginCommand := "login  --use-device --auth-url " + server.URL + " --tenant " + utils.TestTenant
	results, err, std := utils.ExecuteCmd(loginCommand, t)

	// Step 4: Assert results
	assert.False(t, strings.Contains(results, "Verification URL: "+utils.TEST_VERIFICATION_URL))
	assert.False(t, strings.Contains(results, "User Code: "+utils.TEST_USER_CODE))
	assert.True(t, strings.Contains(results, "401"))
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)

	// Step 5: Tear down testing environment
	tearDown(t)
}

func TestDeviceFlowTokenEndpointError(t *testing.T) {
	// Step 1: Setup test env file + mock server
	setUp(t)
	server := utils.MockedIdentityServer()
	defer server.Close()

	// Step 2: Set authorized to be false
	utils.DEVICE_HANDLER_AUTHORIZED = true
	utils.TOKEN_HANDLER_AUTHORIZED = false

	// Step 3: Override auth-url and execute device flow login
	loginCommand := "login  --use-device --auth-url " + server.URL + " --tenant " + utils.TestTenant
	results, err, std := utils.ExecuteCmd(loginCommand, t)

	// Step 4: Assert results
	assert.True(t, strings.Contains(results, "Verification URL: "+utils.TEST_VERIFICATION_URL))
	assert.True(t, strings.Contains(results, "User Code: "+utils.TEST_USER_CODE))
	assert.True(t, strings.Contains(results, "401"))
	assert.Equal(t, nil, err)
	assert.Equal(t, "", std)

	// Step 5: Tear down testing environment
	tearDown(t)
}
