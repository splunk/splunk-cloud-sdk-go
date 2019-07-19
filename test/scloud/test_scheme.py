import unittest
import test


class TestCustomHostAndPort(unittest.TestCase):

    def setUp(self):
        # retrieve the selected tenant name
        code, self.tname, _ = test.scloud("get", "tenant")
        self.assertEqual(0, code)
        self.assertIsNotNone(self.tname)

    def tearDown(self):
        code, _, _ = test.scloud("delete", "scheme")
        self.assertEqual(0, code)

    def test_set_scheme(self):
        code, _, _ = test.scloud("set", "scheme", "http")
        self.assertEqual(0, code)

        code, results, _ = test.scloud("get-settings")
        self.assertEqual("http", results.get("scheme"))

    def test_get_spec_with_http(self):
        """The endpoint system/search/openapi.json
        Does not exists.
        """
        code, _, _ = test.scloud("set", "scheme", "http")
        self.assertEqual(0, code)

        code, results, err = test.scloud("search", "get-spec-json")
        self.assertEqual(1, code)

    def test_get_spec_with_scheme_flag_http(self):
        """This should fail since SDC staging will not allow this."""
        code, results, err = test.scloud("-scheme", "http", "search", "get-spec-json")
        self.assertEqual(1, code)
        print str(err)
        self.assertIn('http:', str(err))

    def test_get_spec_with_scheme_flag_https(self):
        """This should pass since SDC staging only allows this."""
        code, results, err = test.scloud("-scheme", "https", "search", "get-spec-json")
        self.assertEqual(0, code)


if __name__ == "__main__":
    unittest.main()
