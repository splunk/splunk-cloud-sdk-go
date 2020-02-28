package test_engine

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"

	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/auth"
	"github.com/stretchr/testify/assert"
)

var scloud = "../../../bin/scloud_gen"

func executeCliCommand(cmd *exec.Cmd) (string, error, string) {
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	err := cmd.Run()

	return out.String(), err, stderr.String()
}

func getTestcaseResultFilePath(input string) string {

	parts := strings.Split(input, "/")
	filepath := parts[0]
	parts = strings.Split(parts[1], ".test")
	filepath += "/" + parts[0] + ".expected"

	return filepath

}

func getTestCasesAndExecuteCliCommands(filepath string, testarg string) (string, error) {
	ret := ""

	arg := testarg

	// read in testcases line by line
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var line string
	var errs []string
	for {
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		// If a testcase has stdin input file specified, the command and the stdin input are separated for processing here based on the '<'  delimiter.
		testComponents := strings.Split(line, "<")
		var stdinFileName string
		if len(testComponents) > 1 {
			line = testComponents[0]
			stdinFileName = testComponents[1]
			stdinFileName = formatInputForCommandExecution(stdinFileName)
		}
		line = formatInputForCommandExecution(line)

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		args := splitArgs(line)
		for index, ele := range args {
			args[index] = strings.Trim(ele, " ")
		}

		os.Remove(auth.Abspath(".scloudTestOutput"))
		scloudTestOutput := ""
		args = append([]string{arg}, args...)
		cmd := exec.Command(scloud, args...)

		if stdinFileName != "" {
			setCommandStdin(cmd, stdinFileName)
		}

		res, err, _ := executeCliCommand(cmd)
		if !strings.Contains(res, auth.ScloudTestDone) {
			scloudTestOutput = res
		} else {
			outcache, err := ioutil.ReadFile(auth.Abspath(".scloudTestOutput"))
			scloudTestOutput = string(outcache)

			if testarg == "--testhook" {
				if err != nil {
					resStr := ""
					if len(res) > 0 {
						resStr = ", response:" + res
					}
					fmt.Printf("-- Execute \"%s\" Failed: %s%s\n", line, err.Error(), resStr)
				} else {
					fmt.Printf("-- Execute \"%s\" Succeed\n", line)
				}
			}
		}

		if strings.Contains(scloudTestOutput, "Content-Type: application/octet-stream") {
			scloudTestOutput = replaceBoundaryParameter(scloudTestOutput)
		}
		ret = ret + "#testcase: " + line + "\n" + scloudTestOutput + "\n"

		if err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return ret, errors.New(strings.Join(errs, ","))
	}

	return ret, nil
}

func splitArgs(line string) []string {

	line = strings.Trim(line, " ")
	var results []string
	chars := strings.Split(line, "")

	//find all flags
	pattern := " --[a-z\\-]+ "

	re := regexp.MustCompile(pattern)
	found := re.FindAllStringIndex(line, -1)

	if len(found) == 0 {
		return strings.Split(line, " ")
	}

	i := 0
	for index, ele := range found {
		before := strings.Join(chars[i:ele[0]], "")
		if index == 0 {
			args := strings.Split(strings.Trim(before, " "), " ")
			results = append(results, args...)
		} else {
			results = append(results, before)
		}

		flag := strings.Join(chars[ele[0]:ele[1]], "")
		results = append(results, flag)
		i = ele[1]
	}

	if i < len(line) {
		str := strings.Join(chars[i:len(line)], "")
		results = append(results, str)

	}

	return results
}

// Replaces the boundary parameter random string (generated) with a fixed name 'BOUNDARY_PARAMETER'
func replaceBoundaryParameter(scloudTestOutput string) string {
	scloudTestOutputParts := strings.Split(scloudTestOutput, "REQUEST BODY:{--")
	if len(scloudTestOutputParts) > 1 {
		scloudTestOutputParts = strings.Split(scloudTestOutputParts[1], "Content-Disposition")
		if len(scloudTestOutputParts) > 1 {
			boundaryParameter := formatInputForCommandExecution(scloudTestOutputParts[0])
			boundaryParameter = strings.TrimSuffix(boundaryParameter, "\r")
			scloudTestOutput = strings.ReplaceAll(scloudTestOutput, boundaryParameter, "BOUNDARY_PARAMETER")
		}
	}

	return scloudTestOutput
}

// Set stdin value of the command to be executed
func setCommandStdin(cmd *exec.Cmd, stdinFileName string) {
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Printf("Error creating stdin pipe: %s", err.Error())
	}
	file, err := os.Open(stdinFileName)
	if err != nil {
		fmt.Printf("Error opening the stdin input file: %s", err.Error())
	}
	_, err = io.Copy(stdin, file)
	if err != nil {
		fmt.Printf("Error copying stdin input file contents to cmd.Stdin: %s", err.Error())
	}
	err = stdin.Close()
	if err != nil {
		fmt.Printf("Error closing the stdin pipe: %s", err.Error())
	}
}

// Removes whitespaces and new line from test command
func formatInputForCommandExecution(input string) string {
	input = strings.TrimSuffix(input, "\n")
	input = strings.Trim(input, " ")

	return input
}

//Set config and Login
func setConfigurationAndLogin() (string, error) {

	//Config set username, env, tenant and login
	confArgs := []string{"config set --key username --value " + Username,
		"config set --key env --value " + Env1,
		"config set --key tenant --value " + TestTenant,
		"login --pwd " + Password,
	}

	for _, arg := range confArgs {
		args := strings.Split(arg, " ")
		for index, ele := range args {
			args[index] = strings.Trim(ele, " ")
		}
		comnd := exec.Command(scloud, args...)
		res, err, _ := executeCliCommand(comnd)
		if res != "" {
			//error that config set or login didnt work and terminate tests
			return res, err
		}
	}
	return "", nil
}

func RunTest(filepath string, t *testing.T) {
	arg := "--testhook-dryrun"
	results, _ := getTestCasesAndExecuteCliCommands(filepath, arg)

	//read expected result file
	expectedResults, err := ioutil.ReadFile(getTestcaseResultFilePath(filepath))
	fmt.Print(string(expectedResults))

	assert.Nil(t, err)

	// verify result
	assert.Equal(t, string(expectedResults), results)
}

func Record_test_result(filepath string, testhook_arg string, t *testing.T) {

	results, err := getTestCasesAndExecuteCliCommands(filepath, testhook_arg)
	// write results string to file
	f, err := os.Create(getTestcaseResultFilePath(filepath))
	assert.Nil(t, err)

	defer f.Close()

	count, err := f.WriteString(results)
	fmt.Printf("wrote %d bytes\n", count)

	assert.Nil(t, err)
}

//Execute a global flag test case
func Execute_cmd_with_global_flags(command string, searchString string, t *testing.T) bool {

	stderr := ""

	//Set a default env, tenant and username to begin with and login
	res, err := setConfigurationAndLogin()
	assert.Empty(t, res)
	assert.Nil(t, err)

	args := strings.Split(command, " ")
	for index, ele := range args {
		args[index] = strings.Trim(ele, " ")
	}

	comnd := exec.Command(scloud, args...)
	//execute testcase
	res, err, stderr = executeCliCommand(comnd)
	//Validate if response output contains either expected results or an an expected error string
	if strings.Contains(res, searchString) == false && strings.Contains(stderr, searchString) == false {
		return false
	}

	return true
}
