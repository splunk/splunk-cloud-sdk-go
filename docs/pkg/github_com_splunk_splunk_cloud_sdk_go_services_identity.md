# identity
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/identity"


## Usage

#### type Group

```go
type Group struct {

	// created at
	// Required: true
	CreatedAt strfmt.DateTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy string `json:"createdBy"`

	// name
	// Required: true
	Name string `json:"name"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

Group group

#### type GroupMember

```go
type GroupMember struct {

	// added at
	// Required: true
	AddedAt strfmt.DateTime `json:"addedAt"`

	// added by
	// Required: true
	AddedBy string `json:"addedBy"`

	// group
	// Required: true
	Group string `json:"group"`

	// principal
	// Required: true
	Principal string `json:"principal"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

GroupMember Represents a member that belongs to a group

#### type GroupRole

```go
type GroupRole struct {

	// added at
	// Required: true
	AddedAt strfmt.DateTime `json:"addedAt"`

	// added by
	// Required: true
	AddedBy string `json:"addedBy"`

	// group
	// Required: true
	Group string `json:"group"`

	// role
	// Required: true
	Role string `json:"role"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

GroupRole Represents a role that is assigned to a group

#### type Member

```go
type Member struct {

	// When the principal was added to the tenant.
	// Required: true
	AddedAt strfmt.DateTime `json:"addedAt"`

	// added by
	// Required: true
	AddedBy string `json:"addedBy"`

	// name
	// Required: true
	Name string `json:"name"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

Member Represents a member that belongs to a tenant.

#### type Principal

```go
type Principal struct {

	// created at
	// Required: true
	CreatedAt strfmt.DateTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy string `json:"createdBy"`

	// kind
	// Required: true
	Kind string `json:"kind"`

	// name
	// Required: true
	Name string `json:"name"`

	// profile
	Profile interface{} `json:"profile,omitempty"`

	// tenants
	// Required: true
	Tenants []string `json:"tenants"`
}
```

Principal principal

#### type Role

```go
type Role struct {

	// created at
	// Required: true
	CreatedAt strfmt.DateTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy string `json:"createdBy"`

	// name
	// Required: true
	Name string `json:"name"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

Role role

#### type RolePermission

```go
type RolePermission struct {

	// added at
	// Required: true
	// Format: date-time
	AddedAt strfmt.DateTime `json:"addedAt"`

	// added by
	// Required: true
	AddedBy string `json:"addedBy"`

	// permission
	// Required: true
	Permission string `json:"permission"`

	// role
	// Required: true
	Role string `json:"role"`

	// tenant
	// Required: true
	Tenant string `json:"tenant"`
}
```

RolePermission role permission

#### type Service

```go
type Service services.BaseService
```

Service talks to the IAC service

#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new identity service client from the given Config

#### func (*Service) AddMember

```go
func (s *Service) AddMember(name string) (*Member, error)
```
AddMember adds a member to the given tenant

#### func (*Service) AddMemberToGroup

```go
func (s *Service) AddMemberToGroup(groupName string, memberName string) (*GroupMember, error)
```
AddMemberToGroup adds a member to the group

#### func (*Service) AddPermissionToRole

```go
func (s *Service) AddPermissionToRole(roleName string, permissionName string) (*RolePermission, error)
```
AddPermissionToRole Adds permission to a role in this tenant

#### func (*Service) AddRoleToGroup

```go
func (s *Service) AddRoleToGroup(groupName string, roleName string) (*GroupRole, error)
```
AddRoleToGroup adds a role to the group

#### func (*Service) CreateGroup

```go
func (s *Service) CreateGroup(name string) (*Group, error)
```
CreateGroup creates a new group in the given tenant

#### func (*Service) CreatePrincipal

```go
func (s *Service) CreatePrincipal(name string, kind string) (*Principal, error)
```
CreatePrincipal creates a new principal

#### func (*Service) CreateRole

```go
func (s *Service) CreateRole(name string) (*Role, error)
```
CreateRole creates a new authorization role in the given tenant

#### func (*Service) CreateTenant

```go
func (s *Service) CreateTenant(name string) (*Tenant, error)
```
CreateTenant creates a tenant

#### func (*Service) DeleteGroup

```go
func (s *Service) DeleteGroup(name string) error
```
DeleteGroup deletes a group in the given tenant

#### func (*Service) DeletePrincipal

```go
func (s *Service) DeletePrincipal(name string) error
```
DeletePrincipal deletes a principal by name

#### func (*Service) DeleteRole

```go
func (s *Service) DeleteRole(name string) error
```
DeleteRole deletes a defined role for the given tenant

#### func (*Service) DeleteTenant

```go
func (s *Service) DeleteTenant(name string) error
```
DeleteTenant deletes a tenant by name

#### func (*Service) GetGroup

```go
func (s *Service) GetGroup(name string) (*Group, error)
```
GetGroup gets a group in the given tenant

#### func (*Service) GetGroupMember

```go
func (s *Service) GetGroupMember(groupName string, memberName string) (*GroupMember, error)
```
GetGroupMember returns group-member relationship details

#### func (*Service) GetGroupMembers

```go
func (s *Service) GetGroupMembers(groupName string) ([]string, error)
```
GetGroupMembers lists the members attached to the group

#### func (*Service) GetGroupRole

```go
func (s *Service) GetGroupRole(groupName string, roleName string) (*GroupRole, error)
```
GetGroupRole returns group-role relationship details

#### func (*Service) GetGroupRoles

```go
func (s *Service) GetGroupRoles(groupName string) ([]string, error)
```
GetGroupRoles lists the roles attached to the group

#### func (*Service) GetGroups

```go
func (s *Service) GetGroups() ([]string, error)
```
GetGroups list groups that exist int he tenant

#### func (*Service) GetMember

```go
func (s *Service) GetMember(name string) (*Member, error)
```
GetMember gets a member of the given tenant

#### func (*Service) GetMemberGroups

```go
func (s *Service) GetMemberGroups(memberName string) ([]string, error)
```
GetMemberGroups returns the list of groups a member belongs to within a tenant

#### func (*Service) GetMemberPermissions

```go
func (s *Service) GetMemberPermissions(memberName string) ([]string, error)
```
GetMemberPermissions returns the set of permissions granted to the member within
the tenant

#### func (*Service) GetMemberRoles

```go
func (s *Service) GetMemberRoles(memberName string) ([]string, error)
```
GetMemberRoles returns the set of roles thet member posesses within the tenant

#### func (*Service) GetMembers

```go
func (s *Service) GetMembers() ([]string, error)
```
GetMembers returns the list of members in the given tenant

#### func (*Service) GetPrincipal

```go
func (s *Service) GetPrincipal(name string) (*Principal, error)
```
GetPrincipal returns the principal details

#### func (*Service) GetPrincipals

```go
func (s *Service) GetPrincipals() ([]string, error)
```
GetPrincipals returns the list of principals known to IAC

#### func (*Service) GetRole

```go
func (s *Service) GetRole(name string) (*Role, error)
```
GetRole get a role for the given tenant

#### func (*Service) GetRolePermission

```go
func (s *Service) GetRolePermission(roleName string, permissionName string) (*RolePermission, error)
```
GetRolePermission gets permissions for a role in this tenant

#### func (*Service) GetRolePermissions

```go
func (s *Service) GetRolePermissions(roleName string) ([]string, error)
```
GetRolePermissions gets permissions for a role in this tenant

#### func (*Service) GetRoles

```go
func (s *Service) GetRoles() ([]string, error)
```
GetRoles get all roles for the given tenant

#### func (*Service) GetTenant

```go
func (s *Service) GetTenant(name string) (*Tenant, error)
```
GetTenant returns the tenant details

#### func (*Service) GetTenants

```go
func (s *Service) GetTenants() ([]string, error)
```
GetTenants returns the list of tenants in the system

#### func (*Service) RemoveGroupMember

```go
func (s *Service) RemoveGroupMember(groupName string, memberName string) error
```
RemoveGroupMember removes the member from the group

#### func (*Service) RemoveGroupRole

```go
func (s *Service) RemoveGroupRole(groupName string, roleName string) error
```
RemoveGroupRole removes the role from the group

#### func (*Service) RemoveMember

```go
func (s *Service) RemoveMember(name string) error
```
RemoveMember removes a member from the given tenant

#### func (*Service) RemoveRolePermission

```go
func (s *Service) RemoveRolePermission(roleName string, permissionName string) error
```
RemoveRolePermission removes a permission from the role

#### func (*Service) Validate

```go
func (s *Service) Validate() (*ValidateInfo, error)
```
Validate validates the access token obtained from authorization header and
returns the principal name and tenant memberships

#### type Tenant

```go
type Tenant struct {

	// created at
	// Required: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy string `json:"createdBy"`

	// name
	// Required: true
	Name string `json:"name"`

	// status
	// Required: true
	Status string `json:"status"`
}
```

Tenant tenant

#### type ValidateInfo

```go
type ValidateInfo struct {

	// name
	// Required: true
	// Max Length: 36
	// Min Length: 4
	Name string `json:"name"`

	// tenants
	// Required: true
	Tenants []string `json:"tenants"`
}
```

ValidateInfo validate info