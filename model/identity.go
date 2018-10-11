// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model

import (
	"github.com/splunk/splunk-cloud-sdk-go/services/identity"
)

// Tenant is Deprecated: please use services/identity.Tenant
type Tenant = identity.Tenant

// ValidateInfo is Deprecated: please use services/identity.ValidateInfo
type ValidateInfo = identity.ValidateInfo

// Member is Deprecated: please use services/identity.Member
type Member = identity.Member

// Principal is Deprecated: please use services/identity.Principal
type Principal = identity.Principal

// Role is Deprecated: please use services/identity.Role
type Role = identity.Role

// Group is Deprecated: please use services/identity.Group
type Group = identity.Group

// GroupRole is Deprecated: please use services/identity.GroupRole
type GroupRole = identity.GroupRole

// GroupMember is Deprecated: please use services/identity.GroupMember
type GroupMember = identity.GroupMember

// RolePermission is Deprecated: please use services/identity.RolePermission
type RolePermission = identity.RolePermission
