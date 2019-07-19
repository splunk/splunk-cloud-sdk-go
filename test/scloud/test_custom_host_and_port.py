import unittest

import test

CUSTOM_HOST = "scloudtest.com"
CUSTOM_PORT = "8088"
STAGING_HOST = "api.staging.splunkbeta.com"
STAGING_PORT = "443"

def cleanup_host():
    return test.scloud("delete", "host")

def cleanup_port():
    return test.scloud("delete", "port")

class TestCustomHostAndPort(unittest.TestCase):
    def test_set_host_port(self):
        code, _, _ = test.scloud("set", "host", CUSTOM_HOST)
        self.assertEqual(0, code)
        code, _, _ = test.scloud("set", "port", CUSTOM_PORT)
        self.assertEqual(0, code)
        code, results, _ = test.scloud("get-settings")
        self.assertEqual(CUSTOM_HOST, results.get("host"))
        self.assertEqual(CUSTOM_PORT, results.get("port"))
        cleanup_host()
        cleanup_port()

    def test_get_spec_with_custom_host_port_local_settings(self):
        code, _, _ = test.scloud("set", "host", CUSTOM_HOST)
        self.assertEqual(0, code)
        code, _, _ = test.scloud("set", "port", CUSTOM_PORT)
        self.assertEqual(0, code)   
        code, results, err = test.scloud("search", "get-spec-json")
        self.assertEqual(1, code)
        self.assertIn('{}:{}'.format(CUSTOM_HOST, CUSTOM_PORT), str(err))
        cleanup_host()
        cleanup_port()

    def test_get_spec_with_custom_host_port_flags(self):
        code, results, err = test.scloud("-host", CUSTOM_HOST, "-port", CUSTOM_PORT, "search", "get-spec-json")
        self.assertEqual(1, code)
        self.assertIn('{}:{}'.format(CUSTOM_HOST, CUSTOM_PORT), str(err))
    
    def test_get_spec_with_staging_host_port_local_settings(self):
        code, _, _ = test.scloud("set", "host", STAGING_HOST)
        self.assertEqual(0, code)
        code, _, _ = test.scloud("set", "port", STAGING_PORT)
        self.assertEqual(0, code)   
        code, results, err = test.scloud("search", "get-spec-json")
        self.assertEqual(0, code)
        self.assertIsNotNone(results)
        self.assertIsNone(err)
        cleanup_host()
        cleanup_port()

    def test_get_spec_with_staging_host_port_flags(self):
        code, results, err = test.scloud("-host", STAGING_HOST, "-port", STAGING_PORT, "search", "get-spec-json")
        self.assertEqual(0, code)
        self.assertIsNotNone(results)
        self.assertIsNone(err)


