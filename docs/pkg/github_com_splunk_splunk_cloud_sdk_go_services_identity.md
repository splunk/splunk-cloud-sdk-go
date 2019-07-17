# identity
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/identity"


## Usage

#### type AddGroupMemberBody

```go
type AddGroupMemberBody struct {
	Name string `json:"name"`
}
```


#### type AddGroupRoleBody

```go
type AddGroupRoleBody struct {
	Name string `json:"name"`
}
```


#### type AddMemberBody

```go
type AddMemberBody struct {
	Name string `json:"name"`
}
```


#### type AddRolePermissionBody

```go
type AddRolePermissionBody string
```


#### type CreateGroupBody

```go
type CreateGroupBody struct {
	Name string `json:"name"`
}
```


#### type CreateRoleBody

```go
type CreateRoleBody struct {
	Name string `json:"name"`
}
```


#### type Group

```go
type Group struct {
	CreatedAt   string `json:"createdAt"`
	CreatedBy   string `json:"createdBy"`
	MemberCount int32  `json:"memberCount"`
	Name        string `json:"name"`
	RoleCount   int32  `json:"roleCount"`
	Tenant      string `json:"tenant"`
}
```


#### type GroupMember

```go
type GroupMember struct {
	AddedAt   string `json:"addedAt"`
	AddedBy   string `json:"addedBy"`
	Group     string `json:"group"`
	Principal string `json:"principal"`
	Tenant    string `json:"tenant"`
}
```

Represents a member that belongs to a group

#### type GroupRole

```go
type GroupRole struct {
	AddedAt string `json:"addedAt"`
	AddedBy string `json:"addedBy"`
	Group   string `json:"group"`
	Role    string `json:"role"`
	Tenant  string `json:"tenant"`
}
```

Represents a role that is assigned to a group

#### type Member

```go
type Member struct {
	// When the principal was added to the tenant.
	AddedAt string            `json:"addedAt"`
	AddedBy string            `json:"addedBy"`
	Name    string            `json:"name"`
	Tenant  string            `json:"tenant"`
	Profile *PrincipalProfile `json:"profile,omitempty"`
}
```

Represents a member that belongs to a tenant.

#### type Principal

```go
type Principal struct {
	CreatedAt string            `json:"createdAt"`
	CreatedBy string            `json:"createdBy"`
	Kind      PrincipalKind     `json:"kind"`
	Name      string            `json:"name"`
	Tenants   []string          `json:"tenants"`
	UpdatedAt string            `json:"updatedAt"`
	UpdatedBy string            `json:"updatedBy"`
	Profile   *PrincipalProfile `json:"profile,omitempty"`
}
```


#### type PrincipalKind

```go
type PrincipalKind string
```


```go
const (
	PrincipalKindUser           PrincipalKind = "user"
	PrincipalKindServiceAccount PrincipalKind = "service_account"
)
```
List of PrincipalKind

#### type PrincipalProfile

```go
type PrincipalProfile struct {
	Email    *string `json:"email,omitempty"`
	FullName *string `json:"fullName,omitempty"`
}
```

Profile information for a principal

#### type Role

```go
type Role struct {
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	Name      string `json:"name"`
	Tenant    string `json:"tenant"`
}
```


#### type RolePermission

```go
type RolePermission struct {
	AddedAt    string `json:"addedAt"`
	AddedBy    string `json:"addedBy"`
	Permission string `json:"permission"`
	Role       string `json:"role"`
	Tenant     string `json:"tenant"`
}
```


#### type Service

```go
type Service services.BaseService
```


#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new identity service client from the given Config

#### func (*Service) AddGroupMember

```go
func (s *Service) AddGroupMember(group string, addGroupMemberBody AddGroupMemberBody, resp ...*http.Response) (*GroupMember, error)
```
AddGroupMember - identity service endpoint Adds a member to a given group.
Parameters:

    group: The group name.
    addGroupMemberBody: The member to add to a group.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) AddGroupRole

```go
func (s *Service) AddGroupRole(group string, addGroupRoleBody AddGroupRoleBody, resp ...*http.Response) (*GroupRole, error)
```
AddGroupRole - identity service endpoint Adds a role to a given group.
Parameters:

    group: The group name.
    addGroupRoleBody: The role to add to a group.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) AddMember

```go
func (s *Service) AddMember(addMemberBody AddMemberBody, resp ...*http.Response) (*Member, error)
```
AddMember - identity service endpoint Adds a member to a given tenant.
Parameters:

    addMemberBody: The member to associate with a tenant.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) AddRolePermission

```go
func (s *Service) AddRolePermission(role string, body string, resp ...*http.Response) (*RolePermission, error)
```
AddRolePermission - identity service endpoint Adds permissions to a role in a
given tenant. Parameters:

    role: The role name.
    body: The permission to add to a role.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateGroup

```go
func (s *Service) CreateGroup(createGroupBody CreateGroupBody, resp ...*http.Response) (*Group, error)
```
CreateGroup - identity service endpoint Creates a new group in a given tenant.
Parameters:

    createGroupBody: The group definition.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateRole

```go
func (s *Service) CreateRole(createRoleBody CreateRoleBody, resp ...*http.Response) (*Role, error)
```
CreateRole - identity service endpoint Creates a new authorization role in a
given tenant. Parameters:

    createRoleBody: Role definition
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteGroup

```go
func (s *Service) DeleteGroup(group string, resp ...*http.Response) error
```
DeleteGroup - identity service endpoint Deletes a group in a given tenant.
Parameters:

    group: The group name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteRole

```go
func (s *Service) DeleteRole(role string, resp ...*http.Response) error
```
DeleteRole - identity service endpoint Deletes a defined role for a given
tenant. Parameters:

    role: The role name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetGroup

```go
func (s *Service) GetGroup(group string, resp ...*http.Response) (*Group, error)
```
GetGroup - identity service endpoint Returns information about a given group
within a tenant. Parameters:

    group: The group name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetGroupMember

```go
func (s *Service) GetGroupMember(group string, member string, resp ...*http.Response) (*GroupMember, error)
```
GetGroupMember - identity service endpoint Returns information about a given
member within a given group. Parameters:

    group: The group name.
    member: The member name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetGroupRole

```go
func (s *Service) GetGroupRole(group string, role string, resp ...*http.Response) (*GroupRole, error)
```
GetGroupRole - identity service endpoint Returns information about a given role
within a given group. Parameters:

    group: The group name.
    role: The role name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetMember

```go
func (s *Service) GetMember(member string, resp ...*http.Response) (*Member, error)
```
GetMember - identity service endpoint Returns a member of a given tenant.
Parameters:

    member: The member name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetPrincipal

```go
func (s *Service) GetPrincipal(principal string, resp ...*http.Response) (*Principal, error)
```
GetPrincipal - identity service endpoint Returns the details of a principal,
including its tenant membership. Parameters:

    principal: The principal name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetRole

```go
func (s *Service) GetRole(role string, resp ...*http.Response) (*Role, error)
```
GetRole - identity service endpoint Returns a role for a given tenant.
Parameters:

    role: The role name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetRolePermission

```go
func (s *Service) GetRolePermission(role string, permission string, resp ...*http.Response) (*RolePermission, error)
```
GetRolePermission - identity service endpoint Gets a permission for the
specified role. Parameters:

    role: The role name.
    permission: The permission string.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListGroupMembers

```go
func (s *Service) ListGroupMembers(group string, resp ...*http.Response) ([]string, error)
```
ListGroupMembers - identity service endpoint Returns a list of the members
within a given group. Parameters:

    group: The group name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListGroupRoles

```go
func (s *Service) ListGroupRoles(group string, resp ...*http.Response) ([]string, error)
```
ListGroupRoles - identity service endpoint Returns a list of the roles that are
attached to a group within a given tenant. Parameters:

    group: The group name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListGroups

```go
func (s *Service) ListGroups(resp ...*http.Response) ([]string, error)
```
ListGroups - identity service endpoint List the groups that exist in a given
tenant. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListMemberGroups

```go
func (s *Service) ListMemberGroups(member string, resp ...*http.Response) ([]string, error)
```
ListMemberGroups - identity service endpoint Returns a list of groups that a
member belongs to within a tenant. Parameters:

    member: The member name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListMemberPermissions

```go
func (s *Service) ListMemberPermissions(member string, resp ...*http.Response) ([]string, error)
```
ListMemberPermissions - identity service endpoint Returns a set of permissions
granted to the member within the tenant. Parameters:

    member: The member name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListMemberRoles

```go
func (s *Service) ListMemberRoles(member string, resp ...*http.Response) ([]string, error)
```
ListMemberRoles - identity service endpoint Returns a set of roles that a given
member holds within the tenant. Parameters:

    member: The member name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListMembers

```go
func (s *Service) ListMembers(resp ...*http.Response) ([]string, error)
```
ListMembers - identity service endpoint Returns a list of members in a given
tenant. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListPrincipals

```go
func (s *Service) ListPrincipals(resp ...*http.Response) ([]string, error)
```
ListPrincipals - identity service endpoint Returns the list of principals known
to IAC. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListRoleGroups

```go
func (s *Service) ListRoleGroups(role string, resp ...*http.Response) ([]string, error)
```
ListRoleGroups - identity service endpoint Gets a list of groups for a role in a
given tenant. Parameters:

    role: The role name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListRolePermissions

```go
func (s *Service) ListRolePermissions(role string, resp ...*http.Response) ([]string, error)
```
ListRolePermissions - identity service endpoint Gets the permissions for a role
in a given tenant. Parameters:

    role: The role name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListRoles

```go
func (s *Service) ListRoles(resp ...*http.Response) ([]string, error)
```
ListRoles - identity service endpoint Returns all roles for a given tenant.
Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) RemoveGroupMember

```go
func (s *Service) RemoveGroupMember(group string, member string, resp ...*http.Response) error
```
RemoveGroupMember - identity service endpoint Removes the member from a given
group. Parameters:

    group: The group name.
    member: The member name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) RemoveGroupRole

```go
func (s *Service) RemoveGroupRole(group string, role string, resp ...*http.Response) error
```
RemoveGroupRole - identity service endpoint Removes a role from a given group.
Parameters:

    group: The group name.
    role: The role name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) RemoveMember

```go
func (s *Service) RemoveMember(member string, resp ...*http.Response) error
```
RemoveMember - identity service endpoint Removes a member from a given tenant
Parameters:

    member: The member name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) RemoveRolePermission

```go
func (s *Service) RemoveRolePermission(role string, permission string, resp ...*http.Response) error
```
RemoveRolePermission - identity service endpoint Removes a permission from the
role. Parameters:

    role: The role name.
    permission: The permission string.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ValidateToken

```go
func (s *Service) ValidateToken(query *ValidateTokenQueryParams, resp ...*http.Response) (*ValidateInfo, error)
```
ValidateToken - identity service endpoint Validates the access token obtained
from the authorization header and returns the principal name and tenant
memberships. Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### type Servicer

```go
type Servicer interface {
	/*
		AddGroupMember - identity service endpoint
		Adds a member to a given group.
		Parameters:
			group: The group name.
			addGroupMemberBody: The member to add to a group.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	AddGroupMember(group string, addGroupMemberBody AddGroupMemberBody, resp ...*http.Response) (*GroupMember, error)
	/*
		AddGroupRole - identity service endpoint
		Adds a role to a given group.
		Parameters:
			group: The group name.
			addGroupRoleBody: The role to add to a group.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	AddGroupRole(group string, addGroupRoleBody AddGroupRoleBody, resp ...*http.Response) (*GroupRole, error)
	/*
		AddMember - identity service endpoint
		Adds a member to a given tenant.
		Parameters:
			addMemberBody: The member to associate with a tenant.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	AddMember(addMemberBody AddMemberBody, resp ...*http.Response) (*Member, error)
	/*
		AddRolePermission - identity service endpoint
		Adds permissions to a role in a given tenant.
		Parameters:
			role: The role name.
			body: The permission to add to a role.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	AddRolePermission(role string, body string, resp ...*http.Response) (*RolePermission, error)
	/*
		CreateGroup - identity service endpoint
		Creates a new group in a given tenant.
		Parameters:
			createGroupBody: The group definition.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateGroup(createGroupBody CreateGroupBody, resp ...*http.Response) (*Group, error)
	/*
		CreateRole - identity service endpoint
		Creates a new authorization role in a given tenant.
		Parameters:
			createRoleBody: Role definition
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateRole(createRoleBody CreateRoleBody, resp ...*http.Response) (*Role, error)
	/*
		DeleteGroup - identity service endpoint
		Deletes a group in a given tenant.
		Parameters:
			group: The group name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteGroup(group string, resp ...*http.Response) error
	/*
		DeleteRole - identity service endpoint
		Deletes a defined role for a given tenant.
		Parameters:
			role: The role name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteRole(role string, resp ...*http.Response) error
	/*
		GetGroup - identity service endpoint
		Returns information about a given group within a tenant.
		Parameters:
			group: The group name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetGroup(group string, resp ...*http.Response) (*Group, error)
	/*
		GetGroupMember - identity service endpoint
		Returns information about a given member within a given group.
		Parameters:
			group: The group name.
			member: The member name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetGroupMember(group string, member string, resp ...*http.Response) (*GroupMember, error)
	/*
		GetGroupRole - identity service endpoint
		Returns information about a given role within a given group.
		Parameters:
			group: The group name.
			role: The role name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetGroupRole(group string, role string, resp ...*http.Response) (*GroupRole, error)
	/*
		GetMember - identity service endpoint
		Returns a member of a given tenant.
		Parameters:
			member: The member name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetMember(member string, resp ...*http.Response) (*Member, error)
	/*
		GetPrincipal - identity service endpoint
		Returns the details of a principal, including its tenant membership.
		Parameters:
			principal: The principal name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPrincipal(principal string, resp ...*http.Response) (*Principal, error)
	/*
		GetRole - identity service endpoint
		Returns a role for a given tenant.
		Parameters:
			role: The role name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetRole(role string, resp ...*http.Response) (*Role, error)
	/*
		GetRolePermission - identity service endpoint
		Gets a permission for the specified role.
		Parameters:
			role: The role name.
			permission: The permission string.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetRolePermission(role string, permission string, resp ...*http.Response) (*RolePermission, error)
	/*
		ListGroupMembers - identity service endpoint
		Returns a list of the members within a given group.
		Parameters:
			group: The group name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListGroupMembers(group string, resp ...*http.Response) ([]string, error)
	/*
		ListGroupRoles - identity service endpoint
		Returns a list of the roles that are attached to a group within a given tenant.
		Parameters:
			group: The group name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListGroupRoles(group string, resp ...*http.Response) ([]string, error)
	/*
		ListGroups - identity service endpoint
		List the groups that exist in a given tenant.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListGroups(resp ...*http.Response) ([]string, error)
	/*
		ListMemberGroups - identity service endpoint
		Returns a list of groups that a member belongs to within a tenant.
		Parameters:
			member: The member name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListMemberGroups(member string, resp ...*http.Response) ([]string, error)
	/*
		ListMemberPermissions - identity service endpoint
		Returns a set of permissions granted to the member within the tenant.
		Parameters:
			member: The member name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListMemberPermissions(member string, resp ...*http.Response) ([]string, error)
	/*
		ListMemberRoles - identity service endpoint
		Returns a set of roles that a given member holds within the tenant.
		Parameters:
			member: The member name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListMemberRoles(member string, resp ...*http.Response) ([]string, error)
	/*
		ListMembers - identity service endpoint
		Returns a list of members in a given tenant.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListMembers(resp ...*http.Response) ([]string, error)
	/*
		ListPrincipals - identity service endpoint
		Returns the list of principals known to IAC.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListPrincipals(resp ...*http.Response) ([]string, error)
	/*
		ListRoleGroups - identity service endpoint
		Gets a list of groups for a role in a given tenant.
		Parameters:
			role: The role name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListRoleGroups(role string, resp ...*http.Response) ([]string, error)
	/*
		ListRolePermissions - identity service endpoint
		Gets the permissions for a role in a given tenant.
		Parameters:
			role: The role name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListRolePermissions(role string, resp ...*http.Response) ([]string, error)
	/*
		ListRoles - identity service endpoint
		Returns all roles for a given tenant.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListRoles(resp ...*http.Response) ([]string, error)
	/*
		RemoveGroupMember - identity service endpoint
		Removes the member from a given group.
		Parameters:
			group: The group name.
			member: The member name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	RemoveGroupMember(group string, member string, resp ...*http.Response) error
	/*
		RemoveGroupRole - identity service endpoint
		Removes a role from a given group.
		Parameters:
			group: The group name.
			role: The role name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	RemoveGroupRole(group string, role string, resp ...*http.Response) error
	/*
		RemoveMember - identity service endpoint
		Removes a member from a given tenant
		Parameters:
			member: The member name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	RemoveMember(member string, resp ...*http.Response) error
	/*
		RemoveRolePermission - identity service endpoint
		Removes a permission from the role.
		Parameters:
			role: The role name.
			permission: The permission string.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	RemoveRolePermission(role string, permission string, resp ...*http.Response) error
	/*
		ValidateToken - identity service endpoint
		Validates the access token obtained from the authorization header and returns the principal name and tenant memberships.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ValidateToken(query *ValidateTokenQueryParams, resp ...*http.Response) (*ValidateInfo, error)
}
```

Servicer represents the interface for implementing all endpoints for this
service

#### type Tenant

```go
type Tenant struct {
	CreatedAt string       `json:"createdAt"`
	CreatedBy string       `json:"createdBy"`
	Name      string       `json:"name"`
	Status    TenantStatus `json:"status"`
}
```


#### type TenantName

```go
type TenantName string
```


#### type TenantStatus

```go
type TenantStatus string
```


```go
const (
	TenantStatusProvisioning TenantStatus = "provisioning"
	TenantStatusFailed       TenantStatus = "failed"
	TenantStatusReady        TenantStatus = "ready"
	TenantStatusDeleting     TenantStatus = "deleting"
	TenantStatusDeleted      TenantStatus = "deleted"
	TenantStatusSuspended    TenantStatus = "suspended"
)
```
List of TenantStatus

#### type ValidateInfo

```go
type ValidateInfo struct {
	ClientId  string     `json:"clientId"`
	Name      string     `json:"name"`
	Principal *Principal `json:"principal,omitempty"`
	Tenant    *Tenant    `json:"tenant,omitempty"`
}
```


#### type ValidateTokenQueryParams

```go
type ValidateTokenQueryParams struct {
	// Include : Include additional information to return when validating tenant membership. Valid parameters [tenant, principal]
	Include []string `key:"include"`
}
```

ValidateTokenQueryParams represents valid query parameters for the ValidateToken
operation For convenience ValidateTokenQueryParams can be formed in a single
statement, for example:

    `v := ValidateTokenQueryParams{}.SetInclude(...)`

#### func (ValidateTokenQueryParams) SetInclude

```go
func (q ValidateTokenQueryParams) SetInclude(v []string) ValidateTokenQueryParams
```
