package test

import (
	"testing"

	test_engine "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/test/utils"
)

func TestCli_action(t *testing.T) {
	test_engine.RunTest("testcases/action_test", t)
}

//func TestCli_identit(t *testing.T) {
//	test_util.RunTest("testcases/identity_test", t)
//}

func TestCli_ingest(t *testing.T) {
	test_engine.RunTest("testcases/ingest_test", t)
}
