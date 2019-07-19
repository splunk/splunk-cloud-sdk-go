import unittest
import os

import test

PEM_FILEPATH = os.path.join(os.path.abspath('.'), 'forwarders.pem')


def forwarders(*args):
    return test.scloud("forwarders", *args)


class TestForwarders(unittest.TestCase):

    def setUp(self):
        # retrieve the selected tenant name
        code, self.tname, _ = test.scloud("get", "tenant")
        self.assertEqual(0, code)
        self.assertIsNotNone(self.tname)

    def tearDown(self):
        code, result, _ = forwarders("delete-certificates")
        self.assertEqual(0, code)

    def test_create_certificates(self):
        code, result, _ = forwarders("create-certificate", PEM_FILEPATH)
        print "\n--Create-Certificate-Tests--\n"
        print "Validate Success Code"
        print code
        self.assertEqual(0, code)
        self.assertTrue("content" in result)

        print "Validate certificate"

        #Create existing certificate - 400
        print "Create existing certificate"
        code, _, err = forwarders("create-certificate", PEM_FILEPATH)
        self.assertTrue(test.is400(err))

        #Clean up
        print "Delete certificate"
        code, result, _ = forwarders("delete-certificates")
        self.assertEqual(0, code)

    def test_get_certificates(self):
        print "\n--Get-Certificate-Tests--\n"
        print "Create-Certificate"
        code, result, _ = forwarders("create-certificate", PEM_FILEPATH)
        self.assertEqual(0, code)
        self.assertTrue("content" in result)
        print "Get-Certificate"
        code, result, _ = forwarders("list-certificates")
        self.assertEqual(0, code)

        print "Delete All certificates"
        code, result, _ = forwarders("delete-certificates")
        self.assertEqual(0, code)

        #Get certificates when no certificates exist
        print "Get-Certificate-None-Exists"
        code, result, _ = forwarders("list-certificates")
        print "Validate no certificates"
        self.assertFalse("content" in result)
        self.assertEqual(0, code)

    def test_delete_certificates(self):
        print "\n--Delete-Certificates-Tests--\n"
        print "Create-Certificate"
        code, result, _ = forwarders("create-certificate", PEM_FILEPATH)
        self.assertEqual(0, code)
        self.assertTrue("content" in result)

        code, result, _ = forwarders("delete-certificates")
        self.assertEqual(0, code)

        code, result, _ = forwarders("list-certificates")
        self.assertEqual(0, code)
        self.assertFalse("content" in result)
        print "Validate all certificates are deleted"

    def test_get_spec_json(self):
        print "\n--Get-Spec-Json--\n"
        code, result, _ = forwarders("get-spec-json")
        self.assertEqual(0, code)
        self.assertTrue(result)

    def test_get_spec_yaml(self):
        print "\n--Get-Spec-Yml--\n"
        code, result, _ = forwarders("get-spec-yaml")
        self.assertEqual(0, code)
        self.assertTrue(result)


if __name__ == "__main__":
    unittest.main()


