// Package identity -- generated by scloudgen
// !! DO NOT EDIT !! 
// 
package identity

import (
	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/impl/identity"
	"github.com/splunk/splunk-cloud-sdk-go/scloud_generated/utils"
)

// AddGroupMember -- impl
func AddGroupMember(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.AddGroupMember(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// AddGroupRole -- impl
func AddGroupRole(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.AddGroupRole(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// AddMember -- impl
func AddMember(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.AddMember(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// AddRolePermission -- impl
func AddRolePermission(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.AddRolePermission(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// CreateGroup -- impl
func CreateGroup(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.CreateGroup(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// CreateRole -- impl
func CreateRole(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.CreateRole(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// DeleteGroup -- impl
func DeleteGroup(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.DeleteGroup(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// DeleteRole -- impl
func DeleteRole(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.DeleteRole(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// GetGroup -- impl
func GetGroup(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.GetGroup(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// GetGroupMember -- impl
func GetGroupMember(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.GetGroupMember(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// GetGroupRole -- impl
func GetGroupRole(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.GetGroupRole(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// GetMember -- impl
func GetMember(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.GetMember(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// GetPrincipal -- impl
func GetPrincipal(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.GetPrincipal(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// GetRole -- impl
func GetRole(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.GetRole(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// GetRolePermission -- impl
func GetRolePermission(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.GetRolePermission(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ListGroupMembers -- impl
func ListGroupMembers(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ListGroupMembers(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ListGroupRoles -- impl
func ListGroupRoles(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ListGroupRoles(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ListGroups -- impl
func ListGroups(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ListGroups(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ListMemberGroups -- impl
func ListMemberGroups(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ListMemberGroups(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ListMemberPermissions -- impl
func ListMemberPermissions(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ListMemberPermissions(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ListMemberRoles -- impl
func ListMemberRoles(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ListMemberRoles(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ListMembers -- impl
func ListMembers(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ListMembers(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ListPrincipals -- impl
func ListPrincipals(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ListPrincipals(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ListRoleGroups -- impl
func ListRoleGroups(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ListRoleGroups(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ListRolePermissions -- impl
func ListRolePermissions(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ListRolePermissions(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ListRoles -- impl
func ListRoles(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ListRoles(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// RemoveGroupMember -- impl
func RemoveGroupMember(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.RemoveGroupMember(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// RemoveGroupRole -- impl
func RemoveGroupRole(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.RemoveGroupRole(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// RemoveMember -- impl
func RemoveMember(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.RemoveMember(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// RemoveRolePermission -- impl
func RemoveRolePermission(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.RemoveRolePermission(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

// ValidateToken -- impl
func ValidateToken(cmd *cobra.Command, args []string) error {
	client, err := utils.GetClient()
	if err != nil {
		return err
	}

	client_system, err := utils.GetClientSystemTenant()
	if err != nil {
		return err
	}

	ret, err := identity.ValidateToken(client,client_system, args)
	if err != nil {
		return err
	}

	utils.Pprint(ret)

	return nil
}

