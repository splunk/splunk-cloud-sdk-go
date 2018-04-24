package model

// User represents a user object
type User struct {
	ID                string   `json:"id"`
	FirstName         string   `json:"firstName,omitempty"`
	LastName          string   `json:"lastName,omitempty"`
	Email             string   `json:"email,omitempty"`
	Locale            string   `json:"locale,omitempty"`
	Name              string   `json:"name,omitempty"`
	TenantMemberships []string `json:"tenantMemberships,omitempty"`
}
