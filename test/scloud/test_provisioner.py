import unittest

import test


def provisioner(*args):
    return test.scloud("provisioner", *args)


class TestProvisioner(unittest.TestCase):

    def setUp(self):
        # retrieve the selected tenant name
        code, self.tname, _ = test.scloud("get", "tenant")
        self.assertEqual(0, code)
        self.assertIsNotNone(self.tname)

    def test_tenants(self):
        code, tenants, _ = provisioner("list-tenants")
        self.assertEqual(0, code)
        self.assertTrue(any(t["name"] == self.tname for t in tenants))

        code, tenant, _ = provisioner("get-tenant", self.tname)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, tenant["name"])
        self.assertTrue("createdAt" in tenant)
        self.assertTrue("createdBy" in tenant)


if __name__ == "__main__":
    unittest.main()
