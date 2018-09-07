# service
--
    import "github.com/splunk/splunk-cloud-sdk-go/service"

Package service implements a service client which is used to communicate with
Search Service endpoints

## Usage

```go
const (
	AuthorizationType = "Bearer"
)
```
Declare constants for service package

```go
const (
	// DefaultMaxAuthnAttempts defines the maximum number of retries that will be performed for a request encountering an authentication issue
	DefaultMaxAuthnAttempts = 1
)
```

```go
const UserAgent = "client-go"
```
UserAgent SDK Client Identifier

```go
const Version = "0.0.4-218-g77dabca7e9e8"
```
Version the released version of the SDK

#### type ActionService

```go
type ActionService service
```

ActionService - A service the receives incoming notifications and uses
pre-defined templates to turn those notifications into meaningful actions

#### func (*ActionService) CreateAction

```go
func (c *ActionService) CreateAction(action model.Action) (*model.Action, error)
```
CreateAction creates an action

#### func (*ActionService) DeleteAction

```go
func (c *ActionService) DeleteAction(name string) error
```
DeleteAction deletes an action by name

#### func (*ActionService) GetAction

```go
func (c *ActionService) GetAction(name string) (*model.Action, error)
```
GetAction get an action by name

#### func (*ActionService) GetActionStatus

```go
func (c *ActionService) GetActionStatus(name string, statusID string) (*model.ActionStatus, error)
```
GetActionStatus returns an action's status by name

#### func (*ActionService) GetActions

```go
func (c *ActionService) GetActions() ([]model.Action, error)
```
GetActions get all actions

#### func (*ActionService) TriggerAction

```go
func (c *ActionService) TriggerAction(name string, notification model.ActionNotification) (*model.ActionTriggerResponse, error)
```
TriggerAction triggers an action from a notification

#### func (*ActionService) UpdateAction

```go
func (c *ActionService) UpdateAction(name string, action model.ActionUpdateFields) (*model.Action, error)
```
UpdateAction updates and action by name

#### type AuthnResponseHandler

```go
type AuthnResponseHandler struct {
	TokenRetriever idp.TokenRetriever
}
```

AuthnResponseHandler handles logic for updating the client access token in
response to 401 errors

#### func (AuthnResponseHandler) HandleResponse

```go
func (rh AuthnResponseHandler) HandleResponse(client *Client, request *Request, response *http.Response) (*http.Response, error)
```
HandleResponse will retry a request once after re-authenticating if a 401
response code is encountered

#### type BatchEventsSender

```go
type BatchEventsSender struct {
	BatchSize    int
	EventsChan   chan model.Event
	EventsQueue  []model.Event
	QuitChan     chan struct{}
	EventService *IngestService
	IngestTicker *model.Ticker
	WaitGroup    *sync.WaitGroup
	ErrorChan    chan string

	IsRunning bool
}
```

BatchEventsSender sends events in batches or periodically if batch is not full
to Splunk HTTP Event Collector endpoint

#### func (*BatchEventsSender) AddEvent

```go
func (b *BatchEventsSender) AddEvent(event model.Event) error
```
AddEvent pushes a single event into EventsChan

#### func (*BatchEventsSender) GetErrors

```go
func (b *BatchEventsSender) GetErrors() []string
```
GetErrors return all the error messages as an array

#### func (*BatchEventsSender) ResetQueue

```go
func (b *BatchEventsSender) ResetQueue()
```
ResetQueue sets b.EventsQueue to empty, but keep memory allocated for underlying
array

#### func (*BatchEventsSender) Restart

```go
func (b *BatchEventsSender) Restart()
```
Restart will reset batch event sender, clear up queue, error msg, timer etc.

#### func (*BatchEventsSender) Run

```go
func (b *BatchEventsSender) Run()
```
Run sets up ticker and starts a new goroutine

#### func (*BatchEventsSender) SetCallbackHandler

```go
func (b *BatchEventsSender) SetCallbackHandler(callback UserErrHandler)
```
SetCallbackHandler allows users to pass their own callback function

#### func (*BatchEventsSender) Stop

```go
func (b *BatchEventsSender) Stop()
```
Stop sends a signal to QuitChan, wait for all registered goroutines to finish,
stop ticker and clear queue

#### type CatalogService

```go
type CatalogService service
```

CatalogService talks to the Splunk Cloud catalog service

#### func (*CatalogService) CreateDataset

```go
func (c *CatalogService) CreateDataset(dataset model.DatasetInfo) (*model.DatasetInfo, error)
```
CreateDataset creates a new Dataset

#### func (*CatalogService) CreateRule

```go
func (c *CatalogService) CreateRule(rule model.Rule) (*model.Rule, error)
```
CreateRule posts a new rule.

#### func (*CatalogService) DeleteDataset

```go
func (c *CatalogService) DeleteDataset(datasetID string) error
```
DeleteDataset implements delete Dataset endpoint

#### func (*CatalogService) DeleteDatasetField

```go
func (c *CatalogService) DeleteDatasetField(datasetID string, datasetFieldID string) error
```
DeleteDatasetField deletes the field belonging to the specified dataset with the
id datasetFieldID

#### func (*CatalogService) DeleteRule

```go
func (c *CatalogService) DeleteRule(ruleID string) error
```
DeleteRule deletes the rule by the given path.

#### func (*CatalogService) GetDataset

```go
func (c *CatalogService) GetDataset(id string) (*model.DatasetInfo, error)
```
GetDataset returns the Dataset by name

#### func (*CatalogService) GetDatasetField

```go
func (c *CatalogService) GetDatasetField(datasetID string, datasetFieldID string) (*model.Field, error)
```
GetDatasetField returns the field belonging to the specified dataset with the id
datasetFieldID

#### func (*CatalogService) GetDatasetFields

```go
func (c *CatalogService) GetDatasetFields(datasetID string, values url.Values) ([]model.Field, error)
```
GetDatasetFields returns all the fields belonging to the specified dataset

#### func (*CatalogService) GetDatasets

```go
func (c *CatalogService) GetDatasets() ([]model.DatasetInfo, error)
```
GetDatasets returns all Datasets

#### func (*CatalogService) GetRule

```go
func (c *CatalogService) GetRule(ruleID string) (*model.Rule, error)
```
GetRule returns rule by an ID.

#### func (*CatalogService) GetRules

```go
func (c *CatalogService) GetRules() ([]model.Rule, error)
```
GetRules returns all the rules.

#### func (*CatalogService) PatchDatasetField

```go
func (c *CatalogService) PatchDatasetField(datasetID string, datasetFieldID string, datasetField model.Field) (*model.Field, error)
```
PatchDatasetField updates an already existing field in the specified dataset

#### func (*CatalogService) PostDatasetField

```go
func (c *CatalogService) PostDatasetField(datasetID string, datasetField model.Field) (*model.Field, error)
```
PostDatasetField creates a new field in the specified dataset

#### func (*CatalogService) UpdateDataset

```go
func (c *CatalogService) UpdateDataset(dataset model.PartialDatasetInfo, datasetID string) (*model.DatasetInfo, error)
```
UpdateDataset updates an existing Dataset

#### type Client

```go
type Client struct {

	// SearchService talks to the Splunk Cloud search service
	SearchService *SearchService
	// CatalogService talks to the Splunk Cloud catalog service
	CatalogService *CatalogService
	// IngestService talks to the Splunk Cloud ingest service
	IngestService *IngestService
	// IdentityService talks to Splunk Cloud IAC service
	IdentityService *IdentityService
	// KVStoreService talks to Splunk Cloud kvstore service
	KVStoreService *KVStoreService
	// ActionService talks to Splunk Cloud action service
	ActionService *ActionService
}
```

A Client is used to communicate with service endpoints

#### func  NewClient

```go
func NewClient(config *Config) (*Client, error)
```
NewClient creates a Client with config values passed in

#### func (*Client) BuildURL

```go
func (c *Client) BuildURL(queryValues url.Values, urlPathParts ...string) (url.URL, error)
```
BuildURL creates full Splunk Cloud URL with the client cached tenantID

#### func (*Client) BuildURLWithTenantID

```go
func (c *Client) BuildURLWithTenantID(tenantID string, queryValues url.Values, urlPathParts ...string) (url.URL, error)
```
BuildURLWithTenantID creates full Splunk Cloud URL with tenantID

#### func (*Client) Delete

```go
func (c *Client) Delete(requestParams RequestParams) (*http.Response, error)
```
Delete implements HTTP DELETE call RFC2616 does not explicitly forbid it but in
practice some versions of server implementations (tomcat, netty etc) ignore
bodies in DELETE requests

#### func (*Client) Do

```go
func (c *Client) Do(req *Request) (*http.Response, error)
```
Do sends out request and returns HTTP response

#### func (*Client) DoRequest

```go
func (c *Client) DoRequest(requestParams RequestParams) (*http.Response, error)
```
DoRequest creates and execute a new request

#### func (*Client) Get

```go
func (c *Client) Get(requestParams RequestParams) (*http.Response, error)
```
Get implements HTTP Get call

#### func (*Client) GetURL

```go
func (c *Client) GetURL() (*url.URL, error)
```
GetURL returns the client config url string as a url.URL

#### func (*Client) NewBatchEventsSender

```go
func (c *Client) NewBatchEventsSender(batchSize int, interval int64) (*BatchEventsSender, error)
```
NewBatchEventsSender used to initialize dependencies and set values

#### func (*Client) NewBatchEventsSenderWithMaxAllowedError

```go
func (c *Client) NewBatchEventsSenderWithMaxAllowedError(batchSize int, interval int64, maxErrorsAllowed int) (*BatchEventsSender, error)
```
NewBatchEventsSenderWithMaxAllowedError used to initialize dependencies and set
values, the maxErrorsAllowed is the max number of errors allowed before the
eventsender quit

#### func (*Client) NewRequest

```go
func (c *Client) NewRequest(httpMethod, url string, body io.Reader, headers map[string]string) (*Request, error)
```
NewRequest creates a new HTTP Request and set proper header

#### func (*Client) Patch

```go
func (c *Client) Patch(requestParams RequestParams) (*http.Response, error)
```
Patch implements HTTP Patch call

#### func (*Client) Post

```go
func (c *Client) Post(requestParams RequestParams) (*http.Response, error)
```
Post implements HTTP POST call

#### func (*Client) Put

```go
func (c *Client) Put(requestParams RequestParams) (*http.Response, error)
```
Put implements HTTP PUT call

#### func (*Client) UpdateTokenContext

```go
func (c *Client) UpdateTokenContext(ctx *idp.Context)
```
UpdateTokenContext the access token in the Authorization: Bearer header and
retains related context information

#### type Config

```go
type Config struct {
	// TokenRetriever to gather access tokens to be sent in the Authorization: Bearer header on client initialization and upon encountering a 401 response
	TokenRetriever idp.TokenRetriever
	// Token to be sent in the Authorization: Bearer header (not required if TokenRetriever is specified)
	Token string
	// Url string
	URL string
	// TenantID
	TenantID string
	// Timeout
	Timeout time.Duration
	// ResponseHandlers is a slice of handlers to call after a response has been received in the client
	ResponseHandlers []ResponseHandler
}
```

Config is used to set the client specific attributes

#### type IdentityService

```go
type IdentityService service
```

IdentityService talks to the IAC service

#### func (*IdentityService) AddMember

```go
func (c *IdentityService) AddMember(name string) (*model.Member, error)
```
AddMember adds a member to the given tenant

#### func (*IdentityService) AddMemberToGroup

```go
func (c *IdentityService) AddMemberToGroup(groupName string, memberName string) (*model.GroupMember, error)
```
AddMemberToGroup adds a member to the group

#### func (*IdentityService) AddPermissionToRole

```go
func (c *IdentityService) AddPermissionToRole(roleName string, permissionName string) (*model.RolePermission, error)
```
AddPermissionToRole Adds permission to a role in this tenant

#### func (*IdentityService) AddRoleToGroup

```go
func (c *IdentityService) AddRoleToGroup(groupName string, roleName string) (*model.GroupRole, error)
```
AddRoleToGroup adds a role to the group

#### func (*IdentityService) CreateGroup

```go
func (c *IdentityService) CreateGroup(name string) (*model.Group, error)
```
CreateGroup creates a new group in the given tenant

#### func (*IdentityService) CreatePrincipal

```go
func (c *IdentityService) CreatePrincipal(name string, kind string) (*model.Principal, error)
```
CreatePrincipal creates a new principal

#### func (*IdentityService) CreateRole

```go
func (c *IdentityService) CreateRole(name string) (*model.Role, error)
```
CreateRole creates a new authorization role in the given tenant

#### func (*IdentityService) CreateTenant

```go
func (c *IdentityService) CreateTenant(name string) (*model.Tenant, error)
```
CreateTenant creates a tenant

#### func (*IdentityService) DeleteGroup

```go
func (c *IdentityService) DeleteGroup(name string) error
```
DeleteGroup deletes a group in the given tenant

#### func (*IdentityService) DeletePrincipal

```go
func (c *IdentityService) DeletePrincipal(name string) error
```
DeletePrincipal deletes a principal by name

#### func (*IdentityService) DeleteRole

```go
func (c *IdentityService) DeleteRole(name string) error
```
DeleteRole deletes a defined role for the given tenant

#### func (*IdentityService) DeleteTenant

```go
func (c *IdentityService) DeleteTenant(name string) error
```
DeleteTenant deletes a tenant by name

#### func (*IdentityService) GetGroup

```go
func (c *IdentityService) GetGroup(name string) (*model.Group, error)
```
GetGroup gets a group in the given tenant

#### func (*IdentityService) GetGroupMember

```go
func (c *IdentityService) GetGroupMember(groupName string, memberName string) (*model.GroupMember, error)
```
GetGroupMember returns group-member relationship details

#### func (*IdentityService) GetGroupMembers

```go
func (c *IdentityService) GetGroupMembers(groupName string) ([]string, error)
```
GetGroupMembers lists the members attached to the group

#### func (*IdentityService) GetGroupRole

```go
func (c *IdentityService) GetGroupRole(groupName string, roleName string) (*model.GroupRole, error)
```
GetGroupRole returns group-role relationship details

#### func (*IdentityService) GetGroupRoles

```go
func (c *IdentityService) GetGroupRoles(groupName string) ([]string, error)
```
GetGroupRoles lists the roles attached to the group

#### func (*IdentityService) GetGroups

```go
func (c *IdentityService) GetGroups() ([]string, error)
```
GetGroups list groups that exist int he tenant

#### func (*IdentityService) GetMember

```go
func (c *IdentityService) GetMember(name string) (*model.Member, error)
```
GetMember gets a member of the given tenant

#### func (*IdentityService) GetMemberGroups

```go
func (c *IdentityService) GetMemberGroups(memberName string) ([]string, error)
```
GetMemberGroups returns the list of groups a member belongs to within a tenant

#### func (*IdentityService) GetMemberPermissions

```go
func (c *IdentityService) GetMemberPermissions(memberName string) ([]string, error)
```
GetMemberPermissions returns the set of permissions granted to the member within
the tenant

#### func (*IdentityService) GetMemberRoles

```go
func (c *IdentityService) GetMemberRoles(memberName string) ([]string, error)
```
GetMemberRoles returns the set of roles thet member posesses within the tenant

#### func (*IdentityService) GetMembers

```go
func (c *IdentityService) GetMembers() ([]string, error)
```
GetMembers returns the list of members in the given tenant

#### func (*IdentityService) GetPrincipal

```go
func (c *IdentityService) GetPrincipal(name string) (*model.Principal, error)
```
GetPrincipal returns the principal details

#### func (*IdentityService) GetPrincipals

```go
func (c *IdentityService) GetPrincipals() ([]string, error)
```
GetPrincipals returns the list of principals known to IAC

#### func (*IdentityService) GetRole

```go
func (c *IdentityService) GetRole(name string) (*model.Role, error)
```
GetRole get a role for the given tenant

#### func (*IdentityService) GetRolePermission

```go
func (c *IdentityService) GetRolePermission(roleName string, permissionName string) (*model.RolePermission, error)
```
GetRolePermission gets permissions for a role in this tenant

#### func (*IdentityService) GetRolePermissions

```go
func (c *IdentityService) GetRolePermissions(roleName string) ([]string, error)
```
GetRolePermissions gets permissions for a role in this tenant

#### func (*IdentityService) GetRoles

```go
func (c *IdentityService) GetRoles() ([]string, error)
```
GetRoles get all roles for the given tenant

#### func (*IdentityService) GetTenant

```go
func (c *IdentityService) GetTenant(name string) (*model.Tenant, error)
```
GetTenant returns the tenant details

#### func (*IdentityService) GetTenants

```go
func (c *IdentityService) GetTenants() ([]string, error)
```
GetTenants returns the list of tenants in the system

#### func (*IdentityService) RemoveGroupMember

```go
func (c *IdentityService) RemoveGroupMember(groupName string, memberName string) error
```
RemoveGroupMember removes the member from the group

#### func (*IdentityService) RemoveGroupRole

```go
func (c *IdentityService) RemoveGroupRole(groupName string, roleName string) error
```
RemoveGroupRole removes the role from the group

#### func (*IdentityService) RemoveMember

```go
func (c *IdentityService) RemoveMember(name string) error
```
RemoveMember removes a member from the given tenant

#### func (*IdentityService) RemoveRolePermission

```go
func (c *IdentityService) RemoveRolePermission(roleName string, permissionName string) error
```
RemoveRolePermission removes a permission from the role

#### func (*IdentityService) Validate

```go
func (c *IdentityService) Validate() (*model.ValidateInfo, error)
```
Validate validates the access token obtained from authorization header and
returns the principal name and tenant memberships

#### type IngestService

```go
type IngestService service
```

IngestService talks to the Splunk Cloud ingest service

#### func (*IngestService) CreateEvent

```go
func (h *IngestService) CreateEvent(event model.Event) error
```
CreateEvent implements Ingest event endpoint

#### func (*IngestService) CreateEvents

```go
func (h *IngestService) CreateEvents(events []model.Event) error
```
CreateEvents post multiple events in one payload

#### func (*IngestService) CreateMetricEvent

```go
func (h *IngestService) CreateMetricEvent(event model.MetricEvent) error
```
CreateMetricEvent implements Ingest metrics endpoint to send one metric event

#### func (*IngestService) CreateMetricEvents

```go
func (h *IngestService) CreateMetricEvents(events []model.MetricEvent) error
```
CreateMetricEvents implements Ingest metrics endpoint to send multipe metric
events

#### func (*IngestService) CreateRawEvent

```go
func (h *IngestService) CreateRawEvent(event model.Event) error
```
CreateRawEvent implements Ingest raw endpoint

#### type KVStoreService

```go
type KVStoreService service
```

KVStoreService talks to kvstore service

#### func (*KVStoreService) CreateIndex

```go
func (c *KVStoreService) CreateIndex(collectionName string, index model.IndexDefinition) (*model.IndexDescription, error)
```
CreateIndex posts a new index to be added to the collection.

#### func (*KVStoreService) DeleteIndex

```go
func (c *KVStoreService) DeleteIndex(collectionName string, indexName string) error
```
DeleteIndex deletes the specified index in a given collection

#### func (*KVStoreService) DeleteRecordByKey

```go
func (c *KVStoreService) DeleteRecordByKey(collectionName string, keyValue string) error
```
DeleteRecordByKey deletes a particular record present in a given collection
based on the key value provided by the user.

#### func (*KVStoreService) DeleteRecords

```go
func (c *KVStoreService) DeleteRecords(collectionName string, values url.Values) error
```
DeleteRecords deletes records present in a given collection based on the
provided query.

#### func (*KVStoreService) ExportCollection

```go
func (c *KVStoreService) ExportCollection(collectionName string, contentType model.ExportCollectionContentType) (string, error)
```
ExportCollection exports the specified collection records to an external file

#### func (*KVStoreService) GetCollectionStats

```go
func (c *KVStoreService) GetCollectionStats(collection string) (*model.CollectionStats, error)
```
GetCollectionStats returns Collection Stats for the collection

#### func (*KVStoreService) GetCollections

```go
func (c *KVStoreService) GetCollections() ([]model.CollectionDefinition, error)
```
GetCollections gets all the collections

#### func (*KVStoreService) GetRecordByKey

```go
func (c *KVStoreService) GetRecordByKey(collectionName string, keyValue string) (model.Record, error)
```
GetRecordByKey queries a particular record present in a given collection based
on the key value provided by the user.

#### func (*KVStoreService) GetServiceHealthStatus

```go
func (c *KVStoreService) GetServiceHealthStatus() (*model.PingOKBody, error)
```
GetServiceHealthStatus returns Service Health Status

#### func (*KVStoreService) InsertRecord

```go
func (c *KVStoreService) InsertRecord(collectionName string, record map[string]string) (map[string]string, error)
```
InsertRecord - Create a new record in the tenant's specified collection

#### func (*KVStoreService) InsertRecords

```go
func (c *KVStoreService) InsertRecords(collectionName string, records []model.Record) ([]string, error)
```
InsertRecords posts new records to the collection.

#### func (*KVStoreService) ListIndexes

```go
func (c *KVStoreService) ListIndexes(collectionName string) ([]model.IndexDefinition, error)
```
ListIndexes retrieves all the indexes in a given collection

#### func (*KVStoreService) ListRecords

```go
func (c *KVStoreService) ListRecords(collectionName string, filters map[string][]string) ([]map[string]interface{}, error)
```
ListRecords - List the records created for the tenant's specified collection
TODO: include count, offset and orderBy

#### func (*KVStoreService) QueryRecords

```go
func (c *KVStoreService) QueryRecords(collectionName string, values url.Values) ([]model.Record, error)
```
QueryRecords queries records present in a given collection.

#### type Request

```go
type Request struct {
	*http.Request
	NumAttempts     uint
	NumErrorsByType map[string]uint
}
```

Request extends net/http.Request to track number of total attempts and error
counts by type of error

#### func (*Request) GetNumErrorsByResponseCode

```go
func (r *Request) GetNumErrorsByResponseCode(respCode int) uint
```
GetNumErrorsByResponseCode returns number of attemps for a given response code
>= 400

#### func (*Request) UpdateToken

```go
func (r *Request) UpdateToken(accessToken string)
```
UpdateToken replaces the access token in the `Authorization: Bearer` header

#### type RequestParams

```go
type RequestParams struct {
	// Http method name
	Method string
	// Http url
	URL url.URL
	// Body parameter
	Body interface{}
	// Additional headers
	Headers map[string]string
}
```

RequestParams contains all the optional request URL parameters

#### type ResponseHandler

```go
type ResponseHandler interface {
	HandleResponse(client *Client, request *Request, response *http.Response) (*http.Response, error)
}
```

ResponseHandler defines the interface for implementing custom response handling
logic

#### type Search

```go
type Search struct {
}
```

Search is a wrapper class for convenient search operations

#### func (*Search) Cancel

```go
func (search *Search) Cancel() (*model.JobControlReplyMsg, error)
```
Cancel posts a cancel action to the search job

#### func (*Search) DisablePreview

```go
func (search *Search) DisablePreview() (*model.JobControlReplyMsg, error)
```
DisablePreview posts a disablepreview action to the search job

#### func (*Search) EnablePreview

```go
func (search *Search) EnablePreview() (*model.JobControlReplyMsg, error)
```
EnablePreview posts an enablepreview action to the search job

#### func (*Search) Finalize

```go
func (search *Search) Finalize() (*model.JobControlReplyMsg, error)
```
Finalize posts a finalize action to the search job

#### func (*Search) GetEvents

```go
func (search *Search) GetEvents(params *model.FetchEventsRequest) (*model.SearchResults, error)
```
GetEvents returns events from the search

#### func (*Search) GetResults

```go
func (search *Search) GetResults(params *model.FetchResultsRequest) (*model.SearchResults, error)
```
GetResults returns results from the search

#### func (*Search) QueryEvents

```go
func (search *Search) QueryEvents(batchSize, offset int, params *model.FetchEventsRequest) (*SearchIterator, error)
```
QueryEvents waits for job to complete and returns an iterator. If offset and
batchSize are specified, the iterator will return that window of results with
each Next() call

#### func (*Search) QueryResults

```go
func (search *Search) QueryResults(batchSize, offset int, params *model.FetchResultsRequest) (*SearchIterator, error)
```
QueryResults waits for job to complete and returns an iterator. If offset and
batchSize are specified, the iterator will return that window of results with
each Next() call

#### func (*Search) Save

```go
func (search *Search) Save() (*model.JobControlReplyMsg, error)
```
Save posts a save action to the search job

#### func (*Search) SetTTL

```go
func (search *Search) SetTTL() (*model.JobControlReplyMsg, error)
```
SetTTL posts a setttl action to the search job

#### func (*Search) Status

```go
func (search *Search) Status() (*model.SearchJobContent, error)
```
Status returns the status of the search job

#### func (*Search) Touch

```go
func (search *Search) Touch() (*model.JobControlReplyMsg, error)
```
Touch posts a touch action to the search job

#### func (*Search) Wait

```go
func (search *Search) Wait() error
```
Wait polls the job until it's completed or errors out

#### type SearchIterator

```go
type SearchIterator struct {
}
```

SearchIterator is the result of a search query. Its cursor starts at 0 index of
the result set. Use Next() to advance through the rows:

     search, _ := client.SearchService.SubmitSearch(&model.PostJobsRequest{Search: "search index=main | head 5"})
    	pages, _ := search.QueryResults(2, 0, &model.FetchResultsRequest{Count: 5})
    	defer pages.Close()
    	for pages.Next() {
    		values, err := pages.Value()
         ...

    	}
    	err := pages.Err() // get any error encountered during iteration
     ...

#### func  NewSearchIterator

```go
func NewSearchIterator(batch, offset, max int, fn queryFunc) *SearchIterator
```
NewSearchIterator creates a new reference to the iterator object

#### func (*SearchIterator) Close

```go
func (i *SearchIterator) Close()
```
Close checks the status and closes iterator if it's not already. After Close, no
results can be retrieved

#### func (*SearchIterator) Err

```go
func (i *SearchIterator) Err() error
```
Err returns error encountered during iteration

#### func (*SearchIterator) Next

```go
func (i *SearchIterator) Next() bool
```
Next prepares the next result set for reading with the Value method. It returns
true on success, or false if there is no next result row or an error occurred
while preparing it.

Every call to Value, even the first one, must be preceded by a call to Next.

#### func (*SearchIterator) Value

```go
func (i *SearchIterator) Value() (*model.SearchResults, error)
```
Value returns value in current iteration or error out if iterator is closed

#### type SearchService

```go
type SearchService service
```

SearchService talks to the Splunk Cloud search service

#### func (*SearchService) CreateJob

```go
func (service *SearchService) CreateJob(job *model.PostJobsRequest) (string, error)
```
CreateJob dispatches a search and returns sid.

#### func (*SearchService) GetJob

```go
func (service *SearchService) GetJob(jobID string) (*model.SearchJobContent, error)
```
GetJob retrieves information about the specified search.

#### func (*SearchService) GetJobEvents

```go
func (service *SearchService) GetJobEvents(jobID string, params *model.FetchEventsRequest) (*model.SearchResults, error)
```
GetJobEvents Returns the job events with the given `id`.

#### func (*SearchService) GetJobResults

```go
func (service *SearchService) GetJobResults(jobID string, params *model.FetchResultsRequest) (*model.SearchResults, error)
```
GetJobResults Returns the job results with the given `id`.

#### func (*SearchService) GetJobs

```go
func (service *SearchService) GetJobs(params *model.JobsRequest) ([]model.SearchJob, error)
```
GetJobs gets details of all current searches.

#### func (*SearchService) PostJobControl

```go
func (service *SearchService) PostJobControl(jobID string, action *model.JobControlAction) (*model.JobControlReplyMsg, error)
```
PostJobControl runs a job control command for the specified search.

#### func (*SearchService) SubmitSearch

```go
func (service *SearchService) SubmitSearch(job *model.PostJobsRequest) (*Search, error)
```
SubmitSearch creates a search job and wraps the response in an object

#### func (*SearchService) WaitForJob

```go
func (service *SearchService) WaitForJob(sid string, pollInterval time.Duration) error
```
WaitForJob polls the job until it's completed or errors out

#### type UserErrHandler

```go
type UserErrHandler func(*BatchEventsSender)
```

UserErrHandler defines the type of user callback function for batchEventSender