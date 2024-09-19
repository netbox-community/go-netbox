# Circuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Cid** | **string** | Unique circuit ID | 
**Provider** | [**BriefProvider**](BriefProvider.md) |  | 
**ProviderAccount** | Pointer to [**NullableBriefProviderAccount**](BriefProviderAccount.md) |  | [optional] 
**Type** | [**BriefCircuitType**](BriefCircuitType.md) |  | 
**Status** | Pointer to [**CircuitStatus**](CircuitStatus.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenant**](BriefTenant.md) |  | [optional] 
**InstallDate** | Pointer to **NullableString** |  | [optional] 
**TerminationDate** | Pointer to **NullableString** |  | [optional] 
**CommitRate** | Pointer to **NullableInt32** | Committed rate | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TerminationA** | [**NullableCircuitCircuitTermination**](CircuitCircuitTermination.md) |  | [readonly] 
**TerminationZ** | [**NullableCircuitCircuitTermination**](CircuitCircuitTermination.md) |  | [readonly] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Assignments** | Pointer to [**[]BriefCircuitGroupAssignmentSerializer**](BriefCircuitGroupAssignmentSerializer.md) |  | [optional] 

## Methods

### NewCircuit

`func NewCircuit(id int32, url string, displayUrl string, display string, cid string, provider BriefProvider, type_ BriefCircuitType, terminationA NullableCircuitCircuitTermination, terminationZ NullableCircuitCircuitTermination, created NullableTime, lastUpdated NullableTime, ) *Circuit`

NewCircuit instantiates a new Circuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitWithDefaults

`func NewCircuitWithDefaults() *Circuit`

NewCircuitWithDefaults instantiates a new Circuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Circuit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Circuit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Circuit) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Circuit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Circuit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Circuit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *Circuit) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *Circuit) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *Circuit) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *Circuit) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Circuit) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Circuit) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetCid

`func (o *Circuit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Circuit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Circuit) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProvider

`func (o *Circuit) GetProvider() BriefProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Circuit) GetProviderOk() (*BriefProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Circuit) SetProvider(v BriefProvider)`

SetProvider sets Provider field to given value.


### GetProviderAccount

`func (o *Circuit) GetProviderAccount() BriefProviderAccount`

GetProviderAccount returns the ProviderAccount field if non-nil, zero value otherwise.

### GetProviderAccountOk

`func (o *Circuit) GetProviderAccountOk() (*BriefProviderAccount, bool)`

GetProviderAccountOk returns a tuple with the ProviderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccount

`func (o *Circuit) SetProviderAccount(v BriefProviderAccount)`

SetProviderAccount sets ProviderAccount field to given value.

### HasProviderAccount

`func (o *Circuit) HasProviderAccount() bool`

HasProviderAccount returns a boolean if a field has been set.

### SetProviderAccountNil

`func (o *Circuit) SetProviderAccountNil(b bool)`

 SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil

### UnsetProviderAccount
`func (o *Circuit) UnsetProviderAccount()`

UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
### GetType

`func (o *Circuit) GetType() BriefCircuitType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Circuit) GetTypeOk() (*BriefCircuitType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Circuit) SetType(v BriefCircuitType)`

SetType sets Type field to given value.


### GetStatus

`func (o *Circuit) GetStatus() CircuitStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Circuit) GetStatusOk() (*CircuitStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Circuit) SetStatus(v CircuitStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Circuit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *Circuit) GetTenant() BriefTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Circuit) GetTenantOk() (*BriefTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Circuit) SetTenant(v BriefTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Circuit) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Circuit) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Circuit) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetInstallDate

`func (o *Circuit) GetInstallDate() string`

GetInstallDate returns the InstallDate field if non-nil, zero value otherwise.

### GetInstallDateOk

`func (o *Circuit) GetInstallDateOk() (*string, bool)`

GetInstallDateOk returns a tuple with the InstallDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDate

`func (o *Circuit) SetInstallDate(v string)`

SetInstallDate sets InstallDate field to given value.

### HasInstallDate

`func (o *Circuit) HasInstallDate() bool`

HasInstallDate returns a boolean if a field has been set.

### SetInstallDateNil

`func (o *Circuit) SetInstallDateNil(b bool)`

 SetInstallDateNil sets the value for InstallDate to be an explicit nil

### UnsetInstallDate
`func (o *Circuit) UnsetInstallDate()`

UnsetInstallDate ensures that no value is present for InstallDate, not even an explicit nil
### GetTerminationDate

`func (o *Circuit) GetTerminationDate() string`

GetTerminationDate returns the TerminationDate field if non-nil, zero value otherwise.

### GetTerminationDateOk

`func (o *Circuit) GetTerminationDateOk() (*string, bool)`

GetTerminationDateOk returns a tuple with the TerminationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationDate

`func (o *Circuit) SetTerminationDate(v string)`

SetTerminationDate sets TerminationDate field to given value.

### HasTerminationDate

`func (o *Circuit) HasTerminationDate() bool`

HasTerminationDate returns a boolean if a field has been set.

### SetTerminationDateNil

`func (o *Circuit) SetTerminationDateNil(b bool)`

 SetTerminationDateNil sets the value for TerminationDate to be an explicit nil

### UnsetTerminationDate
`func (o *Circuit) UnsetTerminationDate()`

UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
### GetCommitRate

`func (o *Circuit) GetCommitRate() int32`

GetCommitRate returns the CommitRate field if non-nil, zero value otherwise.

### GetCommitRateOk

`func (o *Circuit) GetCommitRateOk() (*int32, bool)`

GetCommitRateOk returns a tuple with the CommitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitRate

`func (o *Circuit) SetCommitRate(v int32)`

SetCommitRate sets CommitRate field to given value.

### HasCommitRate

`func (o *Circuit) HasCommitRate() bool`

HasCommitRate returns a boolean if a field has been set.

### SetCommitRateNil

`func (o *Circuit) SetCommitRateNil(b bool)`

 SetCommitRateNil sets the value for CommitRate to be an explicit nil

### UnsetCommitRate
`func (o *Circuit) UnsetCommitRate()`

UnsetCommitRate ensures that no value is present for CommitRate, not even an explicit nil
### GetDescription

`func (o *Circuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Circuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Circuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Circuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTerminationA

`func (o *Circuit) GetTerminationA() CircuitCircuitTermination`

GetTerminationA returns the TerminationA field if non-nil, zero value otherwise.

### GetTerminationAOk

`func (o *Circuit) GetTerminationAOk() (*CircuitCircuitTermination, bool)`

GetTerminationAOk returns a tuple with the TerminationA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationA

`func (o *Circuit) SetTerminationA(v CircuitCircuitTermination)`

SetTerminationA sets TerminationA field to given value.


### SetTerminationANil

`func (o *Circuit) SetTerminationANil(b bool)`

 SetTerminationANil sets the value for TerminationA to be an explicit nil

### UnsetTerminationA
`func (o *Circuit) UnsetTerminationA()`

UnsetTerminationA ensures that no value is present for TerminationA, not even an explicit nil
### GetTerminationZ

`func (o *Circuit) GetTerminationZ() CircuitCircuitTermination`

GetTerminationZ returns the TerminationZ field if non-nil, zero value otherwise.

### GetTerminationZOk

`func (o *Circuit) GetTerminationZOk() (*CircuitCircuitTermination, bool)`

GetTerminationZOk returns a tuple with the TerminationZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationZ

`func (o *Circuit) SetTerminationZ(v CircuitCircuitTermination)`

SetTerminationZ sets TerminationZ field to given value.


### SetTerminationZNil

`func (o *Circuit) SetTerminationZNil(b bool)`

 SetTerminationZNil sets the value for TerminationZ to be an explicit nil

### UnsetTerminationZ
`func (o *Circuit) UnsetTerminationZ()`

UnsetTerminationZ ensures that no value is present for TerminationZ, not even an explicit nil
### GetComments

`func (o *Circuit) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Circuit) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Circuit) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Circuit) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *Circuit) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Circuit) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Circuit) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Circuit) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Circuit) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Circuit) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Circuit) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Circuit) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Circuit) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Circuit) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Circuit) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Circuit) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Circuit) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Circuit) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Circuit) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Circuit) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Circuit) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Circuit) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetAssignments

`func (o *Circuit) GetAssignments() []BriefCircuitGroupAssignmentSerializer`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *Circuit) GetAssignmentsOk() (*[]BriefCircuitGroupAssignmentSerializer, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *Circuit) SetAssignments(v []BriefCircuitGroupAssignmentSerializer)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *Circuit) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


