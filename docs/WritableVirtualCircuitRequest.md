# WritableVirtualCircuitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** | Unique circuit ID | 
**ProviderNetwork** | [**BriefProviderNetworkRequest**](BriefProviderNetworkRequest.md) |  | 
**ProviderAccount** | Pointer to [**NullableBriefProviderAccountRequest**](BriefProviderAccountRequest.md) |  | [optional] 
**Type** | [**BriefVirtualCircuitTypeRequest**](BriefVirtualCircuitTypeRequest.md) |  | 
**Status** | Pointer to [**CircuitStatusValue**](CircuitStatusValue.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableVirtualCircuitRequest

`func NewWritableVirtualCircuitRequest(cid string, providerNetwork BriefProviderNetworkRequest, type_ BriefVirtualCircuitTypeRequest, ) *WritableVirtualCircuitRequest`

NewWritableVirtualCircuitRequest instantiates a new WritableVirtualCircuitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableVirtualCircuitRequestWithDefaults

`func NewWritableVirtualCircuitRequestWithDefaults() *WritableVirtualCircuitRequest`

NewWritableVirtualCircuitRequestWithDefaults instantiates a new WritableVirtualCircuitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *WritableVirtualCircuitRequest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *WritableVirtualCircuitRequest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *WritableVirtualCircuitRequest) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProviderNetwork

`func (o *WritableVirtualCircuitRequest) GetProviderNetwork() BriefProviderNetworkRequest`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *WritableVirtualCircuitRequest) GetProviderNetworkOk() (*BriefProviderNetworkRequest, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *WritableVirtualCircuitRequest) SetProviderNetwork(v BriefProviderNetworkRequest)`

SetProviderNetwork sets ProviderNetwork field to given value.


### GetProviderAccount

`func (o *WritableVirtualCircuitRequest) GetProviderAccount() BriefProviderAccountRequest`

GetProviderAccount returns the ProviderAccount field if non-nil, zero value otherwise.

### GetProviderAccountOk

`func (o *WritableVirtualCircuitRequest) GetProviderAccountOk() (*BriefProviderAccountRequest, bool)`

GetProviderAccountOk returns a tuple with the ProviderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccount

`func (o *WritableVirtualCircuitRequest) SetProviderAccount(v BriefProviderAccountRequest)`

SetProviderAccount sets ProviderAccount field to given value.

### HasProviderAccount

`func (o *WritableVirtualCircuitRequest) HasProviderAccount() bool`

HasProviderAccount returns a boolean if a field has been set.

### SetProviderAccountNil

`func (o *WritableVirtualCircuitRequest) SetProviderAccountNil(b bool)`

 SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil

### UnsetProviderAccount
`func (o *WritableVirtualCircuitRequest) UnsetProviderAccount()`

UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
### GetType

`func (o *WritableVirtualCircuitRequest) GetType() BriefVirtualCircuitTypeRequest`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableVirtualCircuitRequest) GetTypeOk() (*BriefVirtualCircuitTypeRequest, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableVirtualCircuitRequest) SetType(v BriefVirtualCircuitTypeRequest)`

SetType sets Type field to given value.


### GetStatus

`func (o *WritableVirtualCircuitRequest) GetStatus() CircuitStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableVirtualCircuitRequest) GetStatusOk() (*CircuitStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableVirtualCircuitRequest) SetStatus(v CircuitStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableVirtualCircuitRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *WritableVirtualCircuitRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableVirtualCircuitRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableVirtualCircuitRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableVirtualCircuitRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableVirtualCircuitRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableVirtualCircuitRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDescription

`func (o *WritableVirtualCircuitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableVirtualCircuitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableVirtualCircuitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableVirtualCircuitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *WritableVirtualCircuitRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableVirtualCircuitRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableVirtualCircuitRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableVirtualCircuitRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableVirtualCircuitRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableVirtualCircuitRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableVirtualCircuitRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableVirtualCircuitRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableVirtualCircuitRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableVirtualCircuitRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableVirtualCircuitRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableVirtualCircuitRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


