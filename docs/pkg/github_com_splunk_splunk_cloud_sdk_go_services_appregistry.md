# appregistry
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/appregistry"


## Usage

#### type AppKind

```go
type AppKind string
```

AppKind : The kind of application.

```go
const (
	WEB     AppKind = "web"
	NATIVE  AppKind = "native"
	SERVICE AppKind = "service"
)
```
List of AppKind

#### type AppResponse

```go
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
```


#### type CreateAppRequest

```go
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
	Title string `json:"title"`
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
```


#### type Service

```go
type Service services.BaseService
```

Service - A service the receives incoming notifications and uses pre-defined
templates to turn those notifications into meaningful actions

#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new action service client from the given Config

#### func (*Service) CreateApp

```go
func (s *Service) CreateApp(createAppRequest *CreateAppRequest) (*AppResponse, error)
```
CreateApp Create a new application. * @param createAppRequest Create a new
application. @return AppResponse

#### func (*Service) CreateSubscription

```go
func (s *Service) CreateSubscription(appName string) error
```
CreateSubscription Create a subscription. * @param appName Subscribe a tenant to
an app.

#### func (*Service) DeleteApp

```go
func (s *Service) DeleteApp(appName string) error
```
DeleteApp Delete an application. * @param appName Application name.

#### func (*Service) DeleteSubscription

```go
func (s *Service) DeleteSubscription(appName string) error
```
DeleteSubscription Delete a subscription. * @param appName Application name.

#### func (*Service) GetApp

```go
func (s *Service) GetApp(appName string) (*AppResponse, error)
```
GetApp Retrieve the metadata of an application. * @param appName Application
name. @return AppResponse

#### func (*Service) GetAppSubscriptions

```go
func (s *Service) GetAppSubscriptions(appName string) ([]Subscription, error)
```
GetAppSubscriptions Retrieve the collection of subscriptions to an app. * @param
appName Application name. @return []Subscription

#### func (*Service) GetSubscription

```go
func (s *Service) GetSubscription(appName string) (*Subscription, error)
```
GetSubscription Retrieve or validate a subscription. * @param appName
Application name. @return Subscription

#### func (*Service) ListApps

```go
func (s *Service) ListApps() ([]AppResponse, error)
```
ListApps List applications. @return []AppResponse

#### func (*Service) ListSubscriptions

```go
func (s *Service) ListSubscriptions() ([]Subscription, error)
```
ListSubscriptions Retrieve this tenant&#39;s subscriptions. @return
[]Subscription

#### func (*Service) RotateSecret

```go
func (s *Service) RotateSecret(appName string) (*AppResponse, error)
```
RotateSecret Rotate the client secret for the application. * @param appName
Application name. @return AppResponse

#### func (*Service) UpdateApp

```go
func (s *Service) UpdateApp(appName string, updateAppRequest *UpdateAppRequest) (*AppResponse, error)
```
UpdateApp Update an application. * @param appName Application name. * @param
updateAppRequest Updated app contents. @return AppResponse

#### type Servicer

```go
type Servicer interface {
	/*
	   CreateApp
	   Create a new application.
	   * @param createAppRequest Create a new application.
	   @return AppResponse
	*/
	CreateApp(createAppRequest *CreateAppRequest) (*AppResponse, error)
	/*
	   DeleteApp
	   Delete an application.
	   * @param appName Application name.
	*/
	DeleteApp(appName string) error
	/*
	   GetApp
	   Retrieve the metadata of an application.
	   * @param appName Application name.
	   @return AppResponse
	*/
	GetApp(appName string) (*AppResponse, error)
	/*
	   GetAppSubscriptions
	   Retrieve the collection of subscriptions to an app.
	   * @param appName Application name.
	   @return []Subscription
	*/
	GetAppSubscriptions(appName string) ([]Subscription, error)
	/*
	   ListApps
	   List applications.
	   @return []AppResponse
	*/
	ListApps() ([]AppResponse, error)
	/*
	   RotateSecret
	   Rotate the client secret for the application.
	   * @param appName Application name.
	   @return AppResponse
	*/
	RotateSecret(appName string) (*AppResponse, error)
	/*
	   UpdateApp
	   Update an application.
	   * @param appName Application name.
	   * @param updateAppRequest Updated app contents.
	   @return AppResponse
	*/
	UpdateApp(appName string, updateAppRequest *UpdateAppRequest) (*AppResponse, error)
	/*
	   CreateSubscription
	   Create a subscription.
	   * @param appName Subscribe a tenant to an app.
	*/
	CreateSubscription(appName string) error
	/*
	   DeleteSubscription
	   Delete a subscription.
	   * @param appName Application name.
	*/
	DeleteSubscription(appName string) error
	/*
	   GetSubscription
	   Retrieve or validate a subscription.
	   * @param appName Application name.
	   @return Subscription
	*/
	GetSubscription(appName string) (*Subscription, error)
	/*
	   ListSubscriptions
	   Retrieve this tenant&#39;s subscriptions.
	   @return []Subscription
	*/
	ListSubscriptions() ([]Subscription, error)
}
```

Servicer ...

#### type Subscription

```go
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
```


#### type UpdateAppRequest

```go
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
```
