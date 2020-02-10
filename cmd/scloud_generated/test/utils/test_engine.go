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
	"strings"
	"testing"

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

		line = strings.TrimSuffix(line, "\n")
		line = strings.Trim(line, " ")

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		args := strings.Split(line, " ")
		for index, ele := range args {
			args[index] = strings.Trim(ele, " ")
		}

		args = append([]string{arg}, args...)
		cmd := exec.Command(scloud, args...)
		res, err, _ := executeCliCommand(cmd)

		scloudTestOutput, err := ioutil.ReadFile(".scloudTestOutput")

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

		ret = ret + "#testcase: " + line + "\n" + string(scloudTestOutput) + "\n"

		if err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return ret, errors.New(strings.Join(errs, ","))
	}

	return ret, nil
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
