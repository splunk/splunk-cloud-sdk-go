# provisioner
--
    import "github.com/splunk/splunk-cloud-sdk-go/services/provisioner"


## Usage

#### type CreateProvisionJobBody

```go
type CreateProvisionJobBody struct {
	Apps   []string `json:"apps,omitempty"`
	Tenant *string  `json:"tenant,omitempty"`
}
```


#### type Error

```go
type Error struct {
	// Service error code
	Code string `json:"code"`
	// Human readable error message
	Message string `json:"message"`
}
```


#### type ProvisionJobInfo

```go
type ProvisionJobInfo struct {
	Apps      []string                 `json:"apps"`
	CreatedAt string                   `json:"createdAt"`
	CreatedBy string                   `json:"createdBy"`
	Errors    []ProvisionJobInfoErrors `json:"errors"`
	JobID     string                   `json:"jobID"`
	Status    ProvisionJobInfoStatus   `json:"status"`
	Tenant    string                   `json:"tenant"`
}
```


#### type ProvisionJobInfoErrors

```go
type ProvisionJobInfoErrors struct {
	Code     string  `json:"code"`
	JobStage string  `json:"job_stage"`
	Message  string  `json:"message"`
	App      *string `json:"app,omitempty"`
}
```


#### type ProvisionJobInfoStatus

```go
type ProvisionJobInfoStatus string
```


```go
const (
	ProvisionJobInfoStatusCreated   ProvisionJobInfoStatus = "created"
	ProvisionJobInfoStatusRunning   ProvisionJobInfoStatus = "running"
	ProvisionJobInfoStatusCompleted ProvisionJobInfoStatus = "completed"
)
```
List of ProvisionJobInfoStatus

#### type ProvisionJobs

```go
type ProvisionJobs []ProvisionJobInfo
```


#### type Service

```go
type Service services.BaseService
```


#### func  NewService

```go
func NewService(config *services.Config) (*Service, error)
```
NewService creates a new provisioner service client from the given Config

#### func (*Service) CreateProvisionJob

```go
func (s *Service) CreateProvisionJob(createProvisionJobBody CreateProvisionJobBody, resp ...*http.Response) (*ProvisionJobInfo, error)
```
CreateProvisionJob - provisioner service endpoint Creates a new job that
provisions a new tenant and subscribes apps to the tenant Parameters:

    createProvisionJobBody
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetProvisionJob

```go
func (s *Service) GetProvisionJob(jobId string, resp ...*http.Response) (*ProvisionJobInfo, error)
```
GetProvisionJob - provisioner service endpoint Gets details of a specific
provision job Parameters:

    jobId
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) GetTenant

```go
func (s *Service) GetTenant(tenantName string, resp ...*http.Response) (*TenantInfo, error)
```
GetTenant - provisioner service endpoint Gets a specific tenant Parameters:

    tenantName
    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListProvisionJobs

```go
func (s *Service) ListProvisionJobs(resp ...*http.Response) (*ProvisionJobs, error)
```
ListProvisionJobs - provisioner service endpoint Lists all provision jobs
created by the user Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### func (*Service) ListTenants

```go
func (s *Service) ListTenants(resp ...*http.Response) (*Tenants, error)
```
ListTenants - provisioner service endpoint Lists all tenants that the user can
read Parameters:

    resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided

#### type Servicer

```go
type Servicer interface {
	/*
		CreateProvisionJob - provisioner service endpoint
		Creates a new job that provisions a new tenant and subscribes apps to the tenant
		Parameters:
			createProvisionJobBody
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateProvisionJob(createProvisionJobBody CreateProvisionJobBody, resp ...*http.Response) (*ProvisionJobInfo, error)
	/*
		GetProvisionJob - provisioner service endpoint
		Gets details of a specific provision job
		Parameters:
			jobId
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetProvisionJob(jobId string, resp ...*http.Response) (*ProvisionJobInfo, error)
	/*
		GetTenant - provisioner service endpoint
		Gets a specific tenant
		Parameters:
			tenantName
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetTenant(tenantName string, resp ...*http.Response) (*TenantInfo, error)
	/*
		ListProvisionJobs - provisioner service endpoint
		Lists all provision jobs created by the user
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListProvisionJobs(resp ...*http.Response) (*ProvisionJobs, error)
	/*
		ListTenants - provisioner service endpoint
		Lists all tenants that the user can read
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListTenants(resp ...*http.Response) (*Tenants, error)
}
```

Servicer represents the interface for implementing all endpoints for this
service

#### type TenantInfo

```go
type TenantInfo struct {
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}
```


#### type Tenants

```go
type Tenants []TenantInfo
```
