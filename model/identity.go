// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model
import (
	"github.com/go-openapi/strfmt"

)

// Tenant tenant
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


// ValidateInfo validate info
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

// Member Represents a member that belongs to a tenant.
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

// Principal principal
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

// Role role
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

// Group group
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

// GroupRole Represents a role that is assigned to a group
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

// GroupMember Represents a member that belongs to a group
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

// RolePermission role permission
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