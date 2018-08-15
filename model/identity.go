// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package model
import (
	"github.com/go-openapi/strfmt"

	"time"
)
// Member Represents a member that belongs to a tenant.
// swagger:model Member
type Member struct {

	// When the principal was added to the tenant.
	// Required: true
	AddedAt *strfmt.DateTime `json:"addedAt"`

	// added by
	// Required: true
	AddedBy *string `json:"addedBy"`

	// name
	// Required: true
	Name *string `json:"name"`

	// tenant
	// Required: true
	Tenant *string `json:"tenant"`
}

// Principal principal
// swagger:model Principal
type Principal struct {

	// created at
	// Required: true
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy *string `json:"createdBy"`

	// kind
	// Required: true
	Kind *string `json:"kind"`

	// name
	// Required: true
	Name *string `json:"name"`

	// profile
	Profile interface{} `json:"profile,omitempty"`

	// tenants
	// Required: true
	Tenants []string `json:"tenants"`
}

// Role role
// swagger:model Role
type Role struct {

	// created at
	// Required: true
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy *string `json:"createdBy"`

	// name
	// Required: true
	Name *string `json:"name"`

	// tenant
	// Required: true
	Tenant *string `json:"tenant"`
}

// Group group
// swagger:model Group
type Group struct {

	// created at
	// Required: true
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy *string `json:"createdBy"`

	// name
	// Required: true
	Name *string `json:"name"`

	// tenant
	// Required: true
	Tenant *string `json:"tenant"`
}

// GroupRole Represents a role that is assigned to a group
// swagger:model GroupRole
type GroupRole struct {

	// added at
	// Required: true
	AddedAt *strfmt.DateTime `json:"addedAt"`

	// added by
	// Required: true
	AddedBy *string `json:"addedBy"`

	// group
	// Required: true
	Group *string `json:"group"`

	// role
	// Required: true
	Role *string `json:"role"`

	// tenant
	// Required: true
	Tenant *string `json:"tenant"`
}

// Represents a permission assigned to a role.
type RolePermission struct {
	Tenant     string
	Role       string
	Permission string
	AddedAt    time.Time
	AddedBy    string
}

//
//// TenantDetails tenant details
//// swagger:model TenantDetails
//type TenantDetails struct {
//
//	// created at
//	// Required: true
//	CreatedAt *strfmt.DateTime `json:"createdAt"`
//
//	// created by
//	// Required: true
//	CreatedBy *string `json:"createdBy"`
//
//	// modified at
//	// Required: true
//	ModifiedAt *strfmt.DateTime `json:"modifiedAt"`
//
//	// modified by
//	// Required: true
//	ModifiedBy *string `json:"modifiedBy"`
//
//	// status
//	// Required: true
//	Status TenantStatusValue `json:"status"`
//
//	// tenant Id
//	// Required: true
//	// Pattern: [a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*
//	TenantID *string `json:"tenantId"`
//}
