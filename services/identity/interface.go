// DO NOT EDIT
package identity

type IdentityIface interface {
	// CreateTenant creates a tenant
	CreateTenant(name string) (*Tenant, error)
	// GetTenants returns the list of tenants in the system
	GetTenants() ([]string, error)
	// GetTenant returns  the tenant details
	GetTenant(name string) (*Tenant, error)
	// Validate validates the access token obtained from authorization header and returns the principal name and tenant memberships
	Validate() (*ValidateInfo, error)
	// DeleteTenant deletes a tenant by name
	DeleteTenant(name string) error
	// GetPrincipals returns the list of principals known to IAC
	GetPrincipals() ([]string, error)
	// GetPrincipal returns the principal details
	GetPrincipal(name string) (*Principal, error)
	// DeletePrincipal deletes a principal by name
	DeletePrincipal(name string) error
	// GetMembers returns the list of members in the given tenant
	GetMembers() ([]string, error)
	// GetMember gets a member of the given tenant
	GetMember(name string) (*Member, error)
	// GetMemberGroups returns the list of groups a member belongs to within a tenant
	GetMemberGroups(memberName string) ([]string, error)
	// GetMemberRoles returns the set of roles thet member posesses within the tenant
	GetMemberRoles(memberName string) ([]string, error)
	// GetMemberPermissions returns the set of permissions granted to the member within the tenant
	GetMemberPermissions(memberName string) ([]string, error)
	// GetRoles get all roles for the given tenant
	GetRoles() ([]string, error)
	// GetRole get a role for the given tenant
	GetRole(name string) (*Role, error)
	// GetRolePermissions gets permissions for a role in this tenant
	GetRolePermissions(roleName string) ([]string, error)
	// GetRolePermission gets permissions for a role in this tenant
	GetRolePermission(roleName string, permissionName string) (*RolePermission, error)
	// GetGroups list groups that exist int he tenant
	GetGroups() ([]string, error)
	// GetGroup gets a group in the given tenant
	GetGroup(name string) (*Group, error)
	// GetGroupRoles lists the roles attached to the group
	GetGroupRoles(groupName string) ([]string, error)
	// GetGroupRole returns group-role relationship details
	GetGroupRole(groupName string, roleName string) (*GroupRole, error)
	// GetGroupMembers lists the members attached to the group
	GetGroupMembers(groupName string) ([]string, error)
	// GetGroupMember returns group-member relationship details
	GetGroupMember(groupName string, memberName string) (*GroupMember, error)
	// CreatePrincipal creates a new principal
	CreatePrincipal(name string, kind string) (*Principal, error)
	// AddMember adds a member to the given tenant
	AddMember(name string) (*Member, error)
	// CreateRole creates a new authorization role in the given tenant
	CreateRole(name string) (*Role, error)
	// AddPermissionToRole Adds permission to a role in this tenant
	AddPermissionToRole(roleName string, permissionName string) (*RolePermission, error)
	// CreateGroup creates a new group in the given tenant
	CreateGroup(name string) (*Group, error)
	// AddRoleToGroup adds a role to the group
	AddRoleToGroup(groupName string, roleName string) (*GroupRole, error)
	// AddMemberToGroup adds a member to the group
	AddMemberToGroup(groupName string, memberName string) (*GroupMember, error)
	// RemoveMember removes a member from the given tenant
	RemoveMember(name string) error
	// DeleteRole deletes a defined role for the given tenant
	DeleteRole(name string) error
	// RemoveRolePermission removes a permission from the role
	RemoveRolePermission(roleName string, permissionName string) error
	// DeleteGroup deletes a group in the given tenant
	DeleteGroup(name string) error
	// RemoveGroupRole removes the role from the group
	RemoveGroupRole(groupName string, roleName string) error
	// RemoveGroupMember removes the member from the group
	RemoveGroupMember(groupName string, memberName string) error
}
