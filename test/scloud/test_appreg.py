import unittest
import test

SCLOUD_APP = "scloudapp"


def appreg(*args):
    return test.scloud("appreg", *args)


class TestAppregistry(unittest.TestCase):
    def setUp(self):
        # retrieve the selected tenant name
        code, self.tname, err = test.scloud("get", "tenant")
        self.assertEqual(0, code, err)
        self.assertIsNotNone(self.tname)
        type(self).cleanupApp()

    @classmethod
    def tearDownClass(cls):
        cls.cleanupApp()

    @classmethod
    def cleanupApp(cls):
        appreg("delete-app", SCLOUD_APP)
        appreg("delete-subscription", SCLOUD_APP)

    def test_create_get_delete_app(self):
        appName = SCLOUD_APP

        code, result, err = appreg("create-app", appName, "web", "--redirect-urls", "https://redirect1.com", "--title",
                                   appName)
        self.assertEqual(0, code, err)
        self.assertIsNotNone(result)

        # Get-app-Tests
        code, result, err = appreg("get-app", appName)
        self.assertEqual(0, code, err)
        self.assertIsNotNone(result)

        # List-all-app-Tests
        code, result, err = appreg("list-apps")
        self.assertEqual(0, code, err)
        self.assertIsNotNone(result)

        # Update app
        code, result, err = appreg("update-app", appName, "--redirect-urls", "https://redirect2.com , https://mycompany.com", "--title",
                                   appName)
        self.assertEqual(0, code, err)
        self.assertIsNotNone(result)

        # Clean up
        code, result, err = appreg("delete-app", appName)
        self.assertEqual(0, code, err)

    def test_create_get_delete_subscription(self):
        appName = SCLOUD_APP

        # Create-app
        code, result, err = appreg("create-app", appName, "web", "--redirect-urls", "https://redirect1.com", "--title",
                                   appName)
        self.assertEqual(0, code, err)

        # Create-subscription-Tests
        code, result, err = appreg("create-subscription", appName)
        self.assertEqual(0, code, err)

        # Get-subscription-Tests
        code, result, err = appreg("get-subscription", appName)
        self.assertEqual(0, code, err)
        self.assertIsNotNone(result)
        # Get-subscription-of Non-exist app Tests
        code, result, _err = appreg("get-subscription", "nosuchapp")
        self.assertEqual(1, code)

        # List-all-subscriptions-Test
        code, result, err = appreg("list-subscriptions", "web")
        self.assertEqual(0, code, err)
        self.assertIsNotNone(result)
        # Clean up
        code, result, err = appreg("delete-subscription", appName)
        self.assertEqual(0, code, err)
        code, result, err = appreg("delete-app", appName)
        self.assertEqual(0, code, err)

    def test_rotate_secret(self):
        appName = SCLOUD_APP

        code, result, err = appreg("create-app", appName, "web", "--redirect-urls", "https://redirect1.com", "--title",
                                   appName)
        self.assertEqual(0, code, err)
        self.assertIsNotNone(result)

        # rotate app secret
        code, result, err = appreg("rotate-secret", appName)
        self.assertEqual(0, code, err)
        self.assertIsNotNone(result)

        # Clean up
        code, result, err = appreg("delete-app", appName)
        self.assertEqual(0, code, err)

    def test_get_spec_json(self):
        code, result, err = appreg("get-spec-json")
        self.assertEqual(0, code, err)
        self.assertTrue(result)

    def test_get_spec_yaml(self):
        code, result, err = appreg("get-spec-yaml")
        self.assertEqual(0, code, err)
        self.assertTrue(result)
