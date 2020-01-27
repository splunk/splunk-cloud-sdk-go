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

func executeCliCommand(cmd *exec.Cmd) (string, error) {
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
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
	scloud := "../../../bin/scloud_gen"

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
		res, err := executeCliCommand(cmd)

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
