import unittest
import test
import os

CERT_FILE = os.path.join(os.path.abspath(os.path.dirname(__file__)), "cert.crt")


class TestCustomHostAndPort(unittest.TestCase):

    def setUp(self):
        # retrieve the selected tenant name
        code, self.tname, _ = test.scloud("get", "tenant")
        self.assertEqual(0, code)
        self.assertIsNotNone(self.tname)

    def tearDown(self):
        code, _, _ = test.scloud("delete", "ca-cert")
        self.assertEqual(0, code)
        code, _, _ = test.scloud("delete", "private-key")
        self.assertEqual(0, code)

    def test_set_certs(self):
        code, _, _ = test.scloud("set", "ca-cert", CERT_FILE)
        self.assertEqual(0, code)

        code, results, _ = test.scloud("get-settings")
        self.assertIsNotNone(results.get("ca-cert"))

    def test_get_spec_with_certs(self):
        """This should pass but not use the certs and run on https
            default insecure=false
            default scheme=https
            set ca-cert CERT_FILE
        """
        code, _, _ = test.scloud("set", "ca-cert", CERT_FILE)
        self.assertEqual(0, code)

        code, results, err = test.scloud("search", "get-spec-json")
        self.assertEqual(0, code)

    def test_get_spec_with_certs_insecure(self):
            """This should pass but not use the certs and run on https
                insecure=true
                default scheme=https
                set ca-cert CERT_FILE
            """
            code, _, _ = test.scloud("set", "insecure", "true")
            self.assertEqual(0, code)

            code, _, _ = test.scloud("set", "ca-cert", CERT_FILE)
            self.assertEqual(0, code)

            code, results, err = test.scloud("search", "get-spec-json")
            self.assertEqual(0, code)

    def test_get_spec_with_cert_flags(self):
        """This should pass but not use the certs and run on https
            default insecure=false
            default scheme=https
            -ca-cert=CERT_FILE
        """
        code, results, err = test.scloud("-ca-cert", CERT_FILE, "search", "get-spec-json")
        self.assertEqual(0, code)

    def test_get_spec_with_cert_flags_host_ip(self):
        """This should fail but not use the certs and run on https
            default insecure=false
            default scheme=https
            -ca-cert=CERT_FILE
            -host=52.13.129.113

            $ scloud -host=52.13.129.113 -insecure=false -logtostderr search get-spec-json
        """
        code, results, err = test.scloud("-host", "127.0.0.1", "-insecure", "false", "-ca-cert", CERT_FILE, "search", "get-spec-json")
        self.assertEqual(1, code)


if __name__ == "__main__":
    unittest.main()
