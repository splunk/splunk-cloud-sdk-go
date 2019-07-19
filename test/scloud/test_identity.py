import unittest

import test

# The identity tests make use of a predefined "testsclouduser" principal.
TEST_USER = "testsclouduser"


def identity(*args):
    return test.scloud("identity", *args)


class TestIdentity(unittest.TestCase):

    def setUp(self):
        # retrieve the selected tenant name
        code, self.tname, _ = test.scloud("get", "tenant")
        self.assertEqual(0, code)
        self.assertIsNotNone(self.tname)

        # retrieve the current principal name
        code, result, _ = identity("validate-token")
        self.assertEqual(0, code)
        self.pname = result["name"]
        self.assertTrue(self.pname)

        # groups to delete
        self.delete_groups = []

        # roles to delete
        self.delete_roles = []

    def tearDown(self):
        self._delete_resource(self.delete_groups, "get-group", "delete-group")
        self._delete_resource(self.delete_roles, "get-role", "delete-role")

    def _create_group(self):
        gname = "testgroup%s" % test.epoch()
        self.delete_groups.append(gname)
        return gname

    def _create_role(self):
        rname = "testrole%s" % test.epoch()
        self.delete_roles.append(rname)
        return rname

    def _delete_resource(self, resources, get_cmd, delete_cmd):
        for r in resources:
            code, _, _ = identity(get_cmd, r)

            # found
            if code == 0:
                code, _, _ = identity(delete_cmd, r)
                self.assertEqual(0, code)

    def test_groups(self):
        gname = self._create_group()

        code, group, _ = identity("create-group", gname)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, group["tenant"])
        self.assertEqual(gname, group["name"])
        self.assertTrue("createdAt" in group)
        self.assertTrue("createdBy" in group)

        code, groups, _ = identity("list-groups")
        self.assertEqual(0, code)
        self.assertTrue(gname in groups)

        code, group, _ = identity("get-group", gname)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, group["tenant"])
        self.assertEqual(gname, group["name"])
        self.assertTrue("createdAt" in group)
        self.assertTrue("createdBy" in group)

        code, _, err = identity("create-group", gname)
        self.assertEqual(1, code)
        self.assertTrue(test.is409(err))

        code, _, _ = identity("delete-group", gname)
        self.assertEqual(0, code)

        code, groups, _ = identity("list-groups")
        self.assertEqual(0, code)
        self.assertFalse(gname in groups)

        code, group, err = identity("get-group", gname)
        self.assertEqual(1, code)
        self.assertTrue(test.is404(err))

    def test_group_members(self):
        # setup: create test group, add test user to tenant
        gname = self._create_group()

        code, _, _ = identity("create-group", gname)
        self.assertEqual(0, code)
        # add TEST_USER to tenant, ignore error as it may already be a member
        _, _, _ = identity("add-member", TEST_USER)

        # add test user to test group
        code, group_member, _ = identity("add-group-member", gname, TEST_USER)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, group_member["tenant"])
        self.assertEqual(gname, group_member["group"])
        self.assertEqual(TEST_USER, group_member["principal"])
        self.assertTrue("addedAt" in group_member)
        self.assertEqual(self.pname, group_member["addedBy"])

        code, group_members, _ = identity("list-group-members", gname)
        self.assertEqual(0, code)
        self.assertEqual(2, len(group_members))
        self.assertTrue(self.pname in group_members)
        self.assertTrue(TEST_USER in group_members)

        code, group_member, _ = identity("get-group-member", gname, TEST_USER)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, group_member["tenant"])
        self.assertEqual(gname, group_member["group"])
        self.assertEqual(TEST_USER, group_member["principal"])
        self.assertTrue("addedAt" in group_member)
        self.assertEqual(self.pname, group_member["addedBy"])

        # can't add test user again
        code, _, err = identity("add-group-member", gname, TEST_USER)
        self.assertEqual(1, code)
        self.assertTrue(test.is409(err))

        # remove test user from test group
        code, _, _ = identity("remove-group-member", gname, TEST_USER)
        self.assertEqual(0, code)

        code, group_members, _ = identity("list-group-members", gname)
        self.assertEqual(0, code)
        self.assertEqual(1, len(group_members))
        self.assertTrue(self.pname in group_members)

        code, _, err = identity("get-group-member", gname, TEST_USER)
        self.assertEqual(1, code)
        self.assertTrue(test.is404(err))

        # can't remove the user again
        code, _, err = identity("remove-group-member", gname, TEST_USER)
        self.assertEqual(1, code)
        self.assertTrue(test.is404(err))

        # cleanup
        code, _, _ = identity("delete-group", gname)
        self.assertEqual(0, code)
        code, _, _ = identity("remove-member", TEST_USER)
        self.assertEqual(0, code)

    def test_group_roles(self):
        # setup: create test group and test role
        gname = self._create_group()

        code, _, _ = identity("create-group", gname)
        self.assertEqual(0, code)

        rname = self._create_role()
        code, _, _ = identity("create-role", rname)
        self.assertEqual(0, code)

        # add test role to test group
        code, group_role, _ = identity("add-group-role", gname, rname)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, group_role["tenant"])
        self.assertEqual(gname, group_role["group"])
        self.assertEqual(rname, group_role["role"])
        self.assertTrue("addedAt" in group_role)
        self.assertEqual(self.pname, group_role["addedBy"])

        code, group_roles, _ = identity("list-group-roles", gname)
        self.assertEqual(0, code)
        self.assertEqual(1, len(group_roles))
        self.assertTrue(rname in group_roles)

        code, group_role, _ = identity("get-group-role", gname, rname)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, group_role["tenant"])
        self.assertEqual(gname, group_role["group"])
        self.assertEqual(rname, group_role["role"])
        self.assertTrue("addedAt" in group_role)
        self.assertEqual(self.pname, group_role["addedBy"])

        # can't add test role again
        code, _, err = identity("add-group-role", gname, rname)
        self.assertEqual(1, code)
        self.assertTrue(test.is409(err))

        # remove test role from test group
        code, _, _ = identity("remove-group-role", gname, rname)
        self.assertEqual(0, code)

        code, group_roles, _ = identity("list-group-roles", gname)
        self.assertEqual(0, code)
        self.assertEqual(0, len(group_roles))

        # can't remove test role again
        code, _, err = identity("remove-group-role", gname, rname)
        self.assertEqual(1, code)
        self.assertTrue(test.is404(err))

        # cleanup
        code, _, _ = identity("delete-group", gname)
        self.assertEqual(0, code)
        code, _, _ = identity("delete-role", rname)
        self.assertEqual(0, code)

    def test_members(self):
        # make sure we see ourself
        code, members, _ = identity("list-members")
        self.assertEqual(0, code)
        self.assertTrue(self.pname in members)

        code, member, _ = identity("get-member", self.pname)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, member["tenant"])
        self.assertEqual(self.pname, member["name"])
        self.assertTrue("addedAt" in member)
        self.assertTrue("addedBy" in member)

        # remove member from tenant, ignore errors as it may already be removed
        _, _, _ = identity("remove-member", TEST_USER)
        # add the test user
        code, member, _ = identity("add-member", TEST_USER)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, member["tenant"])
        self.assertTrue("addedAt" in member)
        self.assertEqual(self.pname, member["addedBy"])

        code, members, _ = identity("list-members")
        self.assertEqual(0, code)
        self.assertTrue(TEST_USER in members)

        code, member, _ = identity("get-member", TEST_USER)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, member["tenant"])
        self.assertTrue("addedAt" in member)
        self.assertEqual(member["addedBy"], self.pname)

        # remove the test user
        code, _, _ = identity("remove-member", TEST_USER)
        self.assertEqual(0, code)

        code, members, _ = identity("list-members")
        self.assertEqual(0, code)
        self.assertFalse(TEST_USER in members)

        code, _, err = identity("get-member", TEST_USER)
        self.assertEqual(1, code)
        self.assertTrue(test.is404(err))

        # can't add an unknown user
        unknown = "theunknownuser%s" % test.epoch()
        code, _, err = identity("add-member", unknown)
        self.assertEqual(1, code)
        self.assertTrue(test.is404(err))

    def test_principals(self):
        code, principals, _ = identity("list-principals")
        self.assertEqual(0, code)
        self.assertTrue(self.pname in principals)

        code, principal, _ = identity("get-principal", self.pname)
        self.assertEqual(0, code)
        self.assertEqual(self.pname, principal["name"])
        self.assertEqual("user", principal["kind"])
        self.assertTrue("createdAt" in principal)
        self.assertTrue("createdBy" in principal)
        self.assertTrue(self.tname in principal["tenants"])

    def test_roles(self):
        rname = self._create_role()

        code, role, _ = identity("create-role", rname)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, role["tenant"])
        self.assertEqual(rname, role["name"])
        self.assertTrue("createdAt" in role)
        self.assertTrue("createdBy" in role)

        code, roles, _ = identity("list-roles")
        self.assertEqual(0, code)
        self.assertTrue(rname in roles)

        code, role, _ = identity("get-role", rname)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, role["tenant"])
        self.assertEqual(rname, role["name"])
        self.assertTrue("createdAt" in role)
        self.assertTrue("createdBy" in role)

        code, _, err = identity("create-role", rname)
        self.assertEqual(1, code)
        self.assertTrue(test.is409(err))

        code, _, _ = identity("delete-role", rname)
        self.assertEqual(0, code)

        code, roles, _ = identity("list-roles")
        self.assertEqual(0, code)
        self.assertFalse(rname in roles)

        code, _, err = identity("get-role", rname)
        self.assertEqual(1, code)
        self.assertTrue(test.is404(err))

    def test_role_permissions(self):
        # setup: create test role
        rname = self._create_role()

        code, _, _ = identity("create-role", rname)
        self.assertEqual(0, code)

        # add a read permission to the test role
        read = "%s:*:*.read" % self.tname
        code, role_perm, _ = identity("add-role-permission", rname, read)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, role_perm["tenant"])
        self.assertEqual(rname, role_perm["role"])
        self.assertEqual(read, role_perm["permission"])
        self.assertTrue("addedAt" in role_perm)
        self.assertEqual(self.pname, role_perm["addedBy"])

        code, role_perms, _ = identity("list-role-permissions", rname)
        self.assertEqual(0, code)
        self.assertEqual(1, len(role_perms))
        self.assertTrue(read in role_perms)

        code, role_perm, _ = identity("get-role-permission", rname, read)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, role_perm["tenant"])
        self.assertEqual(rname, role_perm["role"])
        self.assertEqual(read, role_perm["permission"])
        self.assertTrue("addedAt" in role_perm)
        self.assertEqual(self.pname, role_perm["addedBy"])

        # can't add the same permisison again
        code, _, err = identity("add-role-permission", rname, read)
        self.assertEqual(1, code)
        self.assertTrue(test.is409(err))

        # add a write permission to the test role
        write = "%s:*:*.write" % self.tname
        code, role_perm, _ = identity("add-role-permission", rname, write)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, role_perm["tenant"])
        self.assertEqual(rname, role_perm["role"])
        self.assertEqual(write, role_perm["permission"])
        self.assertTrue("addedAt" in role_perm)
        self.assertEqual(self.pname, role_perm["addedBy"])

        code, role_perms, _ = identity("list-role-permissions", rname)
        self.assertEqual(0, code)
        self.assertEqual(2, len(role_perms))
        self.assertTrue(write in role_perms)

        code, role_perm, _ = identity("get-role-permission", rname, write)
        self.assertEqual(0, code)
        self.assertEqual(self.tname, role_perm["tenant"])
        self.assertEqual(rname, role_perm["role"])
        self.assertEqual(write, role_perm["permission"])
        self.assertTrue("addedAt" in role_perm)
        self.assertEqual(self.pname, role_perm["addedBy"])

        # can't add the same permisison again
        code, _, err = identity("add-role-permission", rname, write)
        self.assertEqual(1, code)
        self.assertTrue(test.is409(err))

        # remove both permissions from the test role
        code, _, err = identity("remove-role-permission", rname, read)
        self.assertEqual(0, code)
        code, _, err = identity("remove-role-permission", rname, write)
        self.assertEqual(0, code)

        code, role_perms, _ = identity("list-role-permissions", rname)
        self.assertEqual(0, code)
        self.assertEqual(0, len(role_perms))

        code, _, err = identity("get-role-permission", rname, read)
        self.assertEqual(1, code)
        self.assertTrue(test.is404(err))
        code, _, err = identity("get-role-permission", rname, write)
        self.assertEqual(1, code)
        self.assertTrue(test.is404(err))

        # can't remove permissions again
        code, _, err = identity("remove-role-permission", rname, read)
        self.assertEqual(1, code)
        self.assertTrue(test.is404(err))
        code, _, err = identity("remove-role-permission", rname, write)
        self.assertEqual(1, code)
        self.assertTrue(test.is404(err))

        # cleanup
        code, _, _ = identity("delete-role", rname)
        self.assertEqual(0, code)


if __name__ == "__main__":
    unittest.main()
