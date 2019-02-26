// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package appregistry

import "time"

// AppKind : The kind of application.
type AppKind string

// List of AppKind
const (
	WEB     AppKind = "web"
	NATIVE  AppKind = "native"
	SERVICE AppKind = "service"
)

type AppResponse struct {
	// The kind of application.
	Kind AppKind `json:"kind"`
	// Application name that is unique within the SDC platform.
	Name string `json:"name"`
	// Short paragraph describing the application.
	Description string `json:"description,omitempty"`
	// The URL used to login to the app.
	LoginURL string `json:"loginUrl,omitempty"`
	// The URL used to display the application's logo.
	LogoURL string `json:"logoUrl,omitempty"`
	// Human readable title for the application.
	Title string `json:"title"`
	// Array of permissions that are granted to the app principal when a tenant subscribes.
	AppPrincipalPermissions []string `json:"appPrincipalPermissions,omitempty"`
	// Array of Permission Filter Templates that are used to intersect with a Users permissions when using the app.
	UserPermissionsFilter []string `json:"userPermissionsFilter,omitempty"`
	// Array of URLs that can be used for redirect after logging into the app
	RedirectURLs []string `json:"redirectUrls"`
	// URL to redirect to after a subscription is created
	SetupURL string `json:"setupUrl,omitempty"`
	// URL that webhook events will be sent to.
	WebhookURL string `json:"webhookUrl,omitempty"`
	// OAuth 2.0 Client ID.
	ClientID string `json:"clientId"`
	// OAuth 2.0 Client Secret string (used for confidential clients).
	ClientSecret string `json:"clientSecret,omitempty"`
	// Time at which the Client Secret will expire, or 0 if it does not expire (unix epoch seconds).
	ClientSecretExpiresAt int64 `json:"clientSecretExpiresAt,omitempty"`
	// The date that the application was created.
	CreatedAt time.Time `json:"createdAt"`
	// The principal who created this application.
	CreatedBy string `json:"createdBy"`
}

type CreateAppRequest struct {
	Kind AppKind `json:"kind"`
	// Application name that is unique within the SDC platform.
	Name string `json:"name"`
	// Short paragraph describing the application.
	Description *string `json:"description,omitempty"`
	// The URL used to login to the app.
	LoginURL *string `json:"loginUrl,omitempty"`
	// The URL used to display the application's logo.
	LogoURL *string `json:"logoUrl,omitempty"`
	// Human readable title for the application.
	Title *string `json:"title,omitempty"`
	// Array of permissions that are granted to the app principal when a tenant subscribes
	AppPrincipalPermissions *[]string `json:"appPrincipalPermissions,omitempty"`
	// Array of Permission Filter Templates that are used to intersect with a Users permissions when using the app.
	UserPermissionsFilter *[]string `json:"userPermissionsFilter,omitempty"`
	// Array of URLs that can be used for redirect after logging into the app
	RedirectURLs []string `json:"redirectUrls"`
	// URL to redirect to after a subscription is created
	SetupURL *string `json:"setupUrl,omitempty"`
	// URL that webhook events will be sent to.
	WebhookURL *string `json:"webhookUrl,omitempty"`
}

type Subscription struct {
	// Application Name.
	AppName string `json:"appName"`
	// Time of subscription creation (UTC).
	CreatedAt time.Time `json:"createdAt"`
	// The user who created the subscription.
	CreatedBy string `json:"createdBy"`
	// Short paragraph describing the application.
	Description string `json:"description"`
	// URL used to login to the app.
	LoginURL string `json:"loginUrl"`
	// The URL used to display the application's logo.
	LogoURL string `json:"logoUrl"`
	// The tenant that is subscribed to the app.
	Tenant *string `json:"tenant,omitempty"`
	// Human-readable name for the application.
	Title string `json:"title"`
}

type UpdateAppRequest struct {
	// Short paragraph describing the application.
	Description *string `json:"description,omitempty"`
	// The URL used to login to the app.
	LoginURL *string `json:"loginUrl,omitempty"`
	// The URL used to display the application's logo.
	LogoURL *string `json:"logoUrl,omitempty"`
	// Human readable title for the application.
	Title *string `json:"title,omitempty"`
	// Array of permissions that are granted to the app principal when a tenant subscribes
	AppPrincipalPermissions *[]string `json:"appPrincipalPermissions,omitempty"`
	// Array of Permission Filter Templates that are used to intersect with a Users permissions when using the app.
	UserPermissionsFilter *[]string `json:"userPermissionsFilter,omitempty"`
	// Array of URLs that can be used for redirect after logging into the app
	RedirectURLs *[]string `json:"redirectUrls,omitempty"`
	// URL to redirect to after a subscription is created
	SetupURL *string `json:"setupUrl,omitempty"`
	// URL that webhook events will be sent to.
	WebhookURL *string `json:"webhookUrl,omitempty"`
}
