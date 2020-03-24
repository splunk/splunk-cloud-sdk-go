package test

import (
	"testing"

	test_engine "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/test/utils"
)

func TestCli_action(t *testing.T) {
	test_engine.RunTest("testcases/action_test", t)
}

func TestCli_appreg(t *testing.T) {
	test_engine.RunTest("testcases/appreg_test", t)
}

func TestCli_catalog(t *testing.T) {
	test_engine.RunTest("testcases/catalog_test", t)
}

func TestCli_collect(t *testing.T) {
	test_engine.RunTest("testcases/collect_test", t)
}

func TestCli_forwarders(t *testing.T) {
	test_engine.RunTest("testcases/forwarders_test", t)
}

func TestCli_identity(t *testing.T) {
	test_engine.RunTest("testcases/identity_test", t)
}

func TestCli_ingest(t *testing.T) {
	test_engine.RunTest("testcases/ingest_test", t)
}

func TestCli_kvstore(t *testing.T) {
	test_engine.RunTest("testcases/kvstore_test", t)
}

func TestCli_provisioner(t *testing.T) {
	test_engine.RunTest("testcases/provisioner_test", t)
}

func TestCli_search(t *testing.T) {
	test_engine.RunTest("testcases/search_test", t)
}

func TestCli_streams(t *testing.T) {
	test_engine.RunTest("testcases/streams_test", t)
}
