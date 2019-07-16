# appregistry
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/appregistry"


## Usage

#### type AppMetadataInternal

```go
type AppMetadataInternal struct {
	// OAuth 2.0 Client ID.
	ClientId string `json:"clientId"`
	// The date that the app was created.
	CreatedAt string `json:"createdAt"`
	// The principal who created this app.
	CreatedBy string `json:"createdBy"`
}
```


#### type AppMetadataPrivate

```go
type AppMetadataPrivate struct {
	// Array of URLs that can be used for redirect after logging into the app.
	RedirectUrls []string `json:"redirectUrls,omitempty"`
	// URL to redirect to after a subscription is created.
	SetupUrl *string `json:"setupUrl,omitempty"`
	// URL that webhook events are sent to.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}
```


#### type AppMetadataPublic

```go
type AppMetadataPublic struct {
	// Human-readable title for the app.
	Title string `json:"title"`
	// Array of permission templates that are used to grant permission to the app principal when a tenant subscribes.
	AppPrincipalPermissions []string `json:"appPrincipalPermissions,omitempty"`
	// Short paragraph describing the app.
	Description *string `json:"description,omitempty"`
	// The URL used to log in to the app.
	LoginUrl *string `json:"loginUrl,omitempty"`
	// The URL used to display the app's logo.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Array of permission filter templates that are used to intersect with a user's permissions when using the app.
	UserPermissionsFilter []string `json:"userPermissionsFilter,omitempty"`
}
```


#### type AppMetadataSecret

```go
type AppMetadataSecret struct {
	// OAuth 2.0 Client Secret string (used for confidential clients).
	ClientSecret string `json:"clientSecret"`
}
```


#### type AppName

```go
type AppName struct {
	AppName string `json:"appName"`
}
```


#### type AppResource

```go
type AppResource struct {
	Kind AppResourceKind `json:"kind"`
	// App name that is unique within Splunk Cloud Platform.
	Name string `json:"name"`
}
```


#### type AppResourceKind

```go
type AppResourceKind string
```


```go
const (
	AppResourceKindWeb     AppResourceKind = "web"
	AppResourceKindNative  AppResourceKind = "native"
	AppResourceKindService AppResourceKind = "service"
)
```
List of AppResourceKind

#### type AppResponseCreateUpdate

```go
type AppResponseCreateUpdate struct {
	// OAuth 2.0 Client ID.
	ClientId string `json:"clientId"`
	// OAuth 2.0 Client Secret string (used for confidential clients).
	ClientSecret string `json:"clientSecret"`
	// The date that the app was created.
	CreatedAt string `json:"createdAt"`
	// The principal who created this app.
	CreatedBy string          `json:"createdBy"`
	Kind      AppResourceKind `json:"kind"`
	// App name that is unique within Splunk Cloud Platform.
	Name string `json:"name"`
	// Human-readable title for the app.
	Title string `json:"title"`
	// Array of permission templates that are used to grant permission to the app principal when a tenant subscribes.
	AppPrincipalPermissions []string `json:"appPrincipalPermissions,omitempty"`
	// Short paragraph describing the app.
	Description *string `json:"description,omitempty"`
	// The URL used to log in to the app.
	LoginUrl *string `json:"loginUrl,omitempty"`
	// The URL used to display the app's logo.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Array of URLs that can be used for redirect after logging into the app.
	RedirectUrls []string `json:"redirectUrls,omitempty"`
	// URL to redirect to after a subscription is created.
	SetupUrl *string `json:"setupUrl,omitempty"`
	// Array of permission filter templates that are used to intersect with a user's permissions when using the app.
	UserPermissionsFilter []string `json:"userPermissionsFilter,omitempty"`
	// URL that webhook events are sent to.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}
```


#### type AppResponseGetList

```go
type AppResponseGetList struct {
	// OAuth 2.0 Client ID.
	ClientId string `json:"clientId"`
	// The date that the app was created.
	CreatedAt string `json:"createdAt"`
	// The principal who created this app.
	CreatedBy string          `json:"createdBy"`
	Kind      AppResourceKind `json:"kind"`
	// App name that is unique within Splunk Cloud Platform.
	Name string `json:"name"`
	// Human-readable title for the app.
	Title string `json:"title"`
	// Array of permission templates that are used to grant permission to the app principal when a tenant subscribes.
	AppPrincipalPermissions []string `json:"appPrincipalPermissions,omitempty"`
	// Short paragraph describing the app.
	Description *string `json:"description,omitempty"`
	// The URL used to log in to the app.
	LoginUrl *string `json:"loginUrl,omitempty"`
	// The URL used to display the app's logo.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Array of URLs that can be used for redirect after logging into the app.
	RedirectUrls []string `json:"redirectUrls,omitempty"`
	// URL to redirect to after a subscription is created.
	SetupUrl *string `json:"setupUrl,omitempty"`
	// Array of permission filter templates that are used to intersect with a user's permissions when using the app.
	UserPermissionsFilter []string `json:"userPermissionsFilter,omitempty"`
	// URL that webhook events are sent to.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}
```


#### type CreateAppRequest

```go
type CreateAppRequest struct {
	Kind AppResourceKind `json:"kind"`
	// App name that is unique within Splunk Cloud Platform.
	Name string `json:"name"`
	// Human-readable title for the app.
	Title string `json:"title"`
	// Array of permission templates that are used to grant permission to the app principal when a tenant subscribes.
	AppPrincipalPermissions []string `json:"appPrincipalPermissions,omitempty"`
	// Short paragraph describing the app.
	Description *string `json:"description,omitempty"`
	// The URL used to log in to the app.
	LoginUrl *string `json:"loginUrl,omitempty"`
	// The URL used to display the app's logo.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Array of URLs that can be used for redirect after logging into the app.
	RedirectUrls []string `json:"redirectUrls,omitempty"`
	// URL to redirect to after a subscription is created.
	SetupUrl *string `json:"setupUrl,omitempty"`
	// Array of permission filter templates that are used to intersect with a user's permissions when using the app.
	UserPermissionsFilter []string `json:"userPermissionsFilter,omitempty"`
	// URL that webhook events are sent to.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}
```


#### type Error

```go
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
```


#### type Key

```go
type Key struct {
	// Public key used for verifying signed webhook requests.
	Key *string `json:"key,omitempty"`
}
```

Public Key

#### type ListSubscriptionsQueryParams

```go
type ListSubscriptionsQueryParams struct {
	// Kind : The kind of application.
	Kind *AppResourceKind `key:"kind"`
}
```

ListSubscriptionsQueryParams represents valid query parameters for the
ListSubscriptions operation For convenience ListSubscriptionsQueryParams can be
formed in a single statement, for example:

    `v := ListSubscriptionsQueryParams{}.SetKind(...)`

#### func (ListSubscriptionsQueryParams) SetKind

```go
func (q ListSubscriptionsQueryParams) SetKind(v AppResourceKind) ListSubscriptionsQueryParams
```

#### type Service

```go
type Service services.BaseService
```


#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new appregistry service client from the given Config

#### func (*Service) CreateApp

```go
func (s *Service) CreateApp(createAppRequest CreateAppRequest, resp ...*http.Response) (*AppResponseCreateUpdate, error)
```
CreateApp - appregistry service endpoint Creates an app. Parameters:

    createAppRequest: Creates a new app.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) CreateSubscription

```go
func (s *Service) CreateSubscription(appName AppName, resp ...*http.Response) error
```
CreateSubscription - appregistry service endpoint Creates a subscription.
Parameters:

    appName: Creates a subscription between a tenant and an app.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteApp

```go
func (s *Service) DeleteApp(appName string, resp ...*http.Response) error
```
DeleteApp - appregistry service endpoint Removes an app. Parameters:

    appName: App name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) DeleteSubscription

```go
func (s *Service) DeleteSubscription(appName string, resp ...*http.Response) error
```
DeleteSubscription - appregistry service endpoint Removes a subscription.
Parameters:

    appName: App name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetApp

```go
func (s *Service) GetApp(appName string, resp ...*http.Response) (*AppResponseGetList, error)
```
GetApp - appregistry service endpoint Returns the metadata of an app.
Parameters:

    appName: App name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetKeys

```go
func (s *Service) GetKeys(resp ...*http.Response) ([]Key, error)
```
GetKeys - appregistry service endpoint Returns a list of the public keys used
for verifying signed webhook requests. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetSubscription

```go
func (s *Service) GetSubscription(appName string, resp ...*http.Response) (*Subscription, error)
```
GetSubscription - appregistry service endpoint Returns or validates a
subscription. Parameters:

    appName: App name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListAppSubscriptions

```go
func (s *Service) ListAppSubscriptions(appName string, resp ...*http.Response) ([]Subscription, error)
```
ListAppSubscriptions - appregistry service endpoint Returns the collection of
subscriptions to an app. Parameters:

    appName: App name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListApps

```go
func (s *Service) ListApps(resp ...*http.Response) ([]AppResponseGetList, error)
```
ListApps - appregistry service endpoint Returns a list of apps. Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListSubscriptions

```go
func (s *Service) ListSubscriptions(query *ListSubscriptionsQueryParams, resp ...*http.Response) ([]Subscription, error)
```
ListSubscriptions - appregistry service endpoint Returns the tenant
subscriptions. Parameters:

    query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) RotateSecret

```go
func (s *Service) RotateSecret(appName string, resp ...*http.Response) (*AppResponseCreateUpdate, error)
```
RotateSecret - appregistry service endpoint Rotates the client secret for an
app. Parameters:

    appName: App name.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) UpdateApp

```go
func (s *Service) UpdateApp(appName string, updateAppRequest UpdateAppRequest, resp ...*http.Response) (*AppResponseCreateUpdate, error)
```
UpdateApp - appregistry service endpoint Updates an app. Parameters:

    appName: App name.
    updateAppRequest: Updates app contents.
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### type Servicer

```go
type Servicer interface {
	/*
		CreateApp - appregistry service endpoint
		Creates an app.
		Parameters:
			createAppRequest: Creates a new app.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateApp(createAppRequest CreateAppRequest, resp ...*http.Response) (*AppResponseCreateUpdate, error)
	/*
		CreateSubscription - appregistry service endpoint
		Creates a subscription.
		Parameters:
			appName: Creates a subscription between a tenant and an app.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateSubscription(appName AppName, resp ...*http.Response) error
	/*
		DeleteApp - appregistry service endpoint
		Removes an app.
		Parameters:
			appName: App name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteApp(appName string, resp ...*http.Response) error
	/*
		DeleteSubscription - appregistry service endpoint
		Removes a subscription.
		Parameters:
			appName: App name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteSubscription(appName string, resp ...*http.Response) error
	/*
		GetApp - appregistry service endpoint
		Returns the metadata of an app.
		Parameters:
			appName: App name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetApp(appName string, resp ...*http.Response) (*AppResponseGetList, error)
	/*
		GetKeys - appregistry service endpoint
		Returns a list of the public keys used for verifying signed webhook requests.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetKeys(resp ...*http.Response) ([]Key, error)
	/*
		GetSubscription - appregistry service endpoint
		Returns or validates a subscription.
		Parameters:
			appName: App name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetSubscription(appName string, resp ...*http.Response) (*Subscription, error)
	/*
		ListAppSubscriptions - appregistry service endpoint
		Returns the collection of subscriptions to an app.
		Parameters:
			appName: App name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListAppSubscriptions(appName string, resp ...*http.Response) ([]Subscription, error)
	/*
		ListApps - appregistry service endpoint
		Returns a list of apps.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListApps(resp ...*http.Response) ([]AppResponseGetList, error)
	/*
		ListSubscriptions - appregistry service endpoint
		Returns the tenant subscriptions.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListSubscriptions(query *ListSubscriptionsQueryParams, resp ...*http.Response) ([]Subscription, error)
	/*
		RotateSecret - appregistry service endpoint
		Rotates the client secret for an app.
		Parameters:
			appName: App name.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	RotateSecret(appName string, resp ...*http.Response) (*AppResponseCreateUpdate, error)
	/*
		UpdateApp - appregistry service endpoint
		Updates an app.
		Parameters:
			appName: App name.
			updateAppRequest: Updates app contents.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateApp(appName string, updateAppRequest UpdateAppRequest, resp ...*http.Response) (*AppResponseCreateUpdate, error)
}
```

Servicer represents the interface for implementing all endpoints for this
service

#### type Subscription

```go
type Subscription struct {
	// App name.
	AppName string `json:"appName"`
	// Time of subscription creation (UTC).
	CreatedAt string `json:"createdAt"`
	// The user who created the subscription.
	CreatedBy string `json:"createdBy"`
	// Short paragraph describing the app.
	Description string `json:"description"`
	// URL used to login to the app.
	LoginUrl string `json:"loginUrl"`
	// The URL used to display the app's logo.
	LogoUrl string `json:"logoUrl"`
	// Human-readable name for the application.
	Title string `json:"title"`
	// The tenant that is subscribed to the app.
	Tenant *string `json:"tenant,omitempty"`
}
```


#### type UpdateAppRequest

```go
type UpdateAppRequest struct {
	// Human-readable title for the app.
	Title string `json:"title"`
	// Array of permission templates that are used to grant permission to the app principal when a tenant subscribes.
	AppPrincipalPermissions []string `json:"appPrincipalPermissions,omitempty"`
	// Short paragraph describing the app.
	Description *string `json:"description,omitempty"`
	// The URL used to log in to the app.
	LoginUrl *string `json:"loginUrl,omitempty"`
	// The URL used to display the app's logo.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// Array of URLs that can be used for redirect after logging into the app.
	RedirectUrls []string `json:"redirectUrls,omitempty"`
	// URL to redirect to after a subscription is created.
	SetupUrl *string `json:"setupUrl,omitempty"`
	// Array of permission filter templates that are used to intersect with a user's permissions when using the app.
	UserPermissionsFilter []string `json:"userPermissionsFilter,omitempty"`
	// URL that webhook events are sent to.
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}
```
