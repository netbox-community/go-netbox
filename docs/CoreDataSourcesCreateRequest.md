# CoreDataSourcesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**SourceUrl** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SyncInterval** | Pointer to [**NullableDataSourceSyncInterval**](DataSourceSyncInterval.md) |  | [optional] 
**Parameters** | Pointer to **interface{}** |  | [optional] 
**IgnoreRules** | Pointer to **string** | Patterns (one per line) matching files to ignore when syncing | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewCoreDataSourcesCreateRequest

`func NewCoreDataSourcesCreateRequest(name string, type_ string, sourceUrl string, ) *CoreDataSourcesCreateRequest`

NewCoreDataSourcesCreateRequest instantiates a new CoreDataSourcesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreDataSourcesCreateRequestWithDefaults

`func NewCoreDataSourcesCreateRequestWithDefaults() *CoreDataSourcesCreateRequest`

NewCoreDataSourcesCreateRequestWithDefaults instantiates a new CoreDataSourcesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CoreDataSourcesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreDataSourcesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreDataSourcesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CoreDataSourcesCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CoreDataSourcesCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CoreDataSourcesCreateRequest) SetType(v string)`

SetType sets Type field to given value.


### GetSourceUrl

`func (o *CoreDataSourcesCreateRequest) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *CoreDataSourcesCreateRequest) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *CoreDataSourcesCreateRequest) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetEnabled

`func (o *CoreDataSourcesCreateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CoreDataSourcesCreateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CoreDataSourcesCreateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CoreDataSourcesCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *CoreDataSourcesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreDataSourcesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreDataSourcesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreDataSourcesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSyncInterval

`func (o *CoreDataSourcesCreateRequest) GetSyncInterval() DataSourceSyncInterval`

GetSyncInterval returns the SyncInterval field if non-nil, zero value otherwise.

### GetSyncIntervalOk

`func (o *CoreDataSourcesCreateRequest) GetSyncIntervalOk() (*DataSourceSyncInterval, bool)`

GetSyncIntervalOk returns a tuple with the SyncInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncInterval

`func (o *CoreDataSourcesCreateRequest) SetSyncInterval(v DataSourceSyncInterval)`

SetSyncInterval sets SyncInterval field to given value.

### HasSyncInterval

`func (o *CoreDataSourcesCreateRequest) HasSyncInterval() bool`

HasSyncInterval returns a boolean if a field has been set.

### SetSyncIntervalNil

`func (o *CoreDataSourcesCreateRequest) SetSyncIntervalNil(b bool)`

 SetSyncIntervalNil sets the value for SyncInterval to be an explicit nil

### UnsetSyncInterval
`func (o *CoreDataSourcesCreateRequest) UnsetSyncInterval()`

UnsetSyncInterval ensures that no value is present for SyncInterval, not even an explicit nil
### GetParameters

`func (o *CoreDataSourcesCreateRequest) GetParameters() interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CoreDataSourcesCreateRequest) GetParametersOk() (*interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CoreDataSourcesCreateRequest) SetParameters(v interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CoreDataSourcesCreateRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *CoreDataSourcesCreateRequest) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *CoreDataSourcesCreateRequest) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetIgnoreRules

`func (o *CoreDataSourcesCreateRequest) GetIgnoreRules() string`

GetIgnoreRules returns the IgnoreRules field if non-nil, zero value otherwise.

### GetIgnoreRulesOk

`func (o *CoreDataSourcesCreateRequest) GetIgnoreRulesOk() (*string, bool)`

GetIgnoreRulesOk returns a tuple with the IgnoreRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreRules

`func (o *CoreDataSourcesCreateRequest) SetIgnoreRules(v string)`

SetIgnoreRules sets IgnoreRules field to given value.

### HasIgnoreRules

`func (o *CoreDataSourcesCreateRequest) HasIgnoreRules() bool`

HasIgnoreRules returns a boolean if a field has been set.

### GetOwner

`func (o *CoreDataSourcesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CoreDataSourcesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CoreDataSourcesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CoreDataSourcesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *CoreDataSourcesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *CoreDataSourcesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *CoreDataSourcesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CoreDataSourcesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CoreDataSourcesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CoreDataSourcesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCustomFields

`func (o *CoreDataSourcesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CoreDataSourcesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CoreDataSourcesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CoreDataSourcesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


