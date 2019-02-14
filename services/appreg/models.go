// Copyright © 2018 Splunk Inc.
// SPLUNK CONFIDENTIAL – Use or disclosure of this material in whole or in part
// without a valid written license from Splunk Inc. is PROHIBITED.
//

package appreg

import "time"

type App struct {
	Client OidcClient `json:"client"`
	// The date that the application was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user who created this application.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Short paragraph describing the application.
	Description *string `json:"description,omitempty"`
	// The URL used to display the application's logo.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Application name.
	Name string `json:"name"`
	// Human readable title for the application.
	Title *string `json:"title,omitempty"`
}

// AppKind : The kind of application.
type AppKind string

// List of AppKind
const (
	WEB AppKind = "web"
	NATIVE AppKind = "native"
	SERVICE AppKind = "service"
)


type AppMetadataInternal struct {
	// OAuth 2.0 Client ID.
	ClientId string `json:"clientId"`
	// OAuth 2.0 Client Secret string (used for confidential clients).
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Time at which the Client Secret will expire, or 0 if it does not expire (unix epoch seconds).
	ClientSecretExpiresAt *int64 `json:"clientSecretExpiresAt,omitempty"`
	// The date that the application was created.
	CreatedAt time.Time `json:"createdAt"`
	// The principal who created this application.
	CreatedBy string `json:"createdBy"`
}


type AppMetadataPrivate struct {
	// Array of permissions that are granted to the app principal when a tenant subscribes
	PrincipalPermissions *[]string `json:"principalPermissions,omitempty"`
	// Array of URLs that can be used for redirect after logging into the app
	RedirectUrls *[]string `json:"redirectUrls,omitempty"`
	// URL to redirect to after a subscription is created
	SetupUrl *string `json:"setupUrl,omitempty"`
	// URL that webhook events will be sent to.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}


type AppMetadataPublic struct {
	// Short paragraph describing the application.
	Description *string `json:"description,omitempty"`
	// The URL used to login to the app.
	LoginUrl *string `json:"loginUrl,omitempty"`
	// The URL used to display the application's logo.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Human readable title for the application.
	Title *string `json:"title,omitempty"`
}


type AppName struct {
	AppName string `json:"appName"`
}


type AppResource struct {
	Kind AppKind `json:"kind"`
	// Application name that is unique within the SDC platform.
	Name string `json:"name"`
}



type AppResponse struct {
	Kind AppKind `json:"kind"`
	// Application name that is unique within the SDC platform.
	Name string `json:"name"`
	// Short paragraph describing the application.
	Description *string `json:"description,omitempty"`
	// The URL used to login to the app.
	LoginUrl *string `json:"loginUrl,omitempty"`
	// The URL used to display the application's logo.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Human readable title for the application.
	Title *string `json:"title,omitempty"`
	// Array of permissions that are granted to the app principal when a tenant subscribes
	PrincipalPermissions *[]string `json:"principalPermissions,omitempty"`
	// Array of URLs that can be used for redirect after logging into the app
	RedirectUrls *[]string `json:"redirectUrls,omitempty"`
	// URL to redirect to after a subscription is created
	SetupUrl *string `json:"setupUrl,omitempty"`
	// URL that webhook events will be sent to.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
	// OAuth 2.0 Client ID.
	ClientId string `json:"clientId"`
	// OAuth 2.0 Client Secret string (used for confidential clients).
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Time at which the Client Secret will expire, or 0 if it does not expire (unix epoch seconds).
	ClientSecretExpiresAt *int64 `json:"clientSecretExpiresAt,omitempty"`
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
	LoginUrl *string `json:"loginUrl,omitempty"`
	// The URL used to display the application's logo.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Human readable title for the application.
	Title *string `json:"title,omitempty"`
	// Array of permissions that are granted to the app principal when a tenant subscribes
	PrincipalPermissions *[]string `json:"principalPermissions,omitempty"`
	// Array of URLs that can be used for redirect after logging into the app
	RedirectUrls *[]string `json:"redirectUrls,omitempty"`
	// URL to redirect to after a subscription is created
	SetupUrl *string `json:"setupUrl,omitempty"`
	// URL that webhook events will be sent to.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}


// The request was processed successfully.
type Created struct {
}


// The request was processed successfully.
type NoContent struct {
}

type OidcAppType string

// List of OidcAppType
const (
	appregWEB OidcAppType = "web"
	appregNATIVE OidcAppType = "native"
	appregBROWSER OidcAppType = "browser"
	appregSERVICE OidcAppType = "service"
)

// OIDC client metadata.
type OidcClient struct {
	ApplicationType OidcAppType `json:"applicationType"`
	// OAuth 2.0 Client ID.
	ClientId *string `json:"clientId,omitempty"`
	// OAuth 2.0 Client Secret string (used for confidential clients).
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Time at which the Client Secret will expire, or 0 if it does not expire (unix epoch seconds).
	ClientSecretExpiresAt *int64 `json:"clientSecretExpiresAt,omitempty"`
	// OAuth 2.0 grant type strings. Default value: authorization_code.
	GrantTypes *[]OidcGrantType `json:"grantTypes,omitempty"`
	// The URL used to login to the app.
	LoginUrl *string `json:"loginUrl,omitempty"`
	// Array of URLs that the user can be redirected to during logout.
	LogoutRedirectUrls *[]string `json:"logoutRedirectUrls,omitempty"`
	// Array of redirection URL strings for use in redirect-based flows.
	RedirectUrls *[]string `json:"redirectUrls,omitempty"`
	// OAuth 2.0 response type strings. Default value: code.
	ResponseTypes *[]OidcResponseType `json:"responseTypes,omitempty"`
	TokenEndpointAuthMethod *OidcTokenEndpointAuthMethod `json:"tokenEndpointAuthMethod,omitempty"`
}

type OidcGrantType string

// List of OidcGrantType
const (
	AUTHORIZATION_CODE OidcGrantType = "authorization_code"
	IMPLICIT OidcGrantType = "implicit"
	REFRESH_TOKEN OidcGrantType = "refresh_token"
	CLIENT_CREDENTIALS OidcGrantType = "client_credentials"
)
// OidcResponseType : OAuth 2.0 response type.
type OidcResponseType string

// List of OidcResponseType
const (
	CODE OidcResponseType = "code"
	TOKEN OidcResponseType = "token"
	ID_TOKEN OidcResponseType = "id_token"
)
type OidcTokenEndpointAuthMethod string

// List of OidcTokenEndpointAuthMethod
const (
	NONE OidcTokenEndpointAuthMethod = "none"
	CLIENT_SECRET_POST OidcTokenEndpointAuthMethod = "client_secret_post"
	CLIENT_SECRET_BASIC OidcTokenEndpointAuthMethod = "client_secret_basic"
	CLIENT_SECRET_JWT OidcTokenEndpointAuthMethod = "client_secret_jwt"
)


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
	LoginUrl string `json:"loginUrl"`
	// The URL used to display the application's logo.
	LogoUrl string `json:"logoUrl"`
	// The tenant that is subscribed to the app.
	Tenant *string `json:"tenant,omitempty"`
	// Human-readable name for the application.
	Title string `json:"title"`
}



type TrustedOrigin struct {
	// The date that the application was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The user who created this application.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The name of the Trusted Origin.
	Name string `json:"name"`
	// Origin URL.
	Origin string `json:"origin"`
	Scopes []TrustedOriginScope `json:"scopes"`
}

// TrustedOriginScope : Defines the scenarios to which this Trusted Origin applies.
type TrustedOriginScope string

// List of TrustedOriginScope
const (
	CORS TrustedOriginScope = "CORS"
)

type UpdateAppRequest struct {
	// Short paragraph describing the application.
	Description *string `json:"description,omitempty"`
	// The URL used to login to the app.
	LoginUrl *string `json:"loginUrl,omitempty"`
	// The URL used to display the application's logo.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Human readable title for the application.
	Title *string `json:"title,omitempty"`
	// Array of permissions that are granted to the app principal when a tenant subscribes
	PrincipalPermissions *[]string `json:"principalPermissions,omitempty"`
	// Array of URLs that can be used for redirect after logging into the app
	RedirectUrls *[]string `json:"redirectUrls,omitempty"`
	// URL to redirect to after a subscription is created
	SetupUrl *string `json:"setupUrl,omitempty"`
	// URL that webhook events will be sent to.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}
