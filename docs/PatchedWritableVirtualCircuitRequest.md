# PatchedWritableVirtualCircuitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | Pointer to **string** | Unique circuit ID | [optional] 
**ProviderNetwork** | Pointer to [**BriefProviderNetworkRequest**](BriefProviderNetworkRequest.md) |  | [optional] 
**ProviderAccount** | Pointer to [**NullableBriefProviderAccountRequest**](BriefProviderAccountRequest.md) |  | [optional] 
**Type** | Pointer to [**BriefVirtualCircuitTypeRequest**](BriefVirtualCircuitTypeRequest.md) |  | [optional] 
**Status** | Pointer to [**CircuitStatusValue**](CircuitStatusValue.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableVirtualCircuitRequest

`func NewPatchedWritableVirtualCircuitRequest() *PatchedWritableVirtualCircuitRequest`

NewPatchedWritableVirtualCircuitRequest instantiates a new PatchedWritableVirtualCircuitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableVirtualCircuitRequestWithDefaults

`func NewPatchedWritableVirtualCircuitRequestWithDefaults() *PatchedWritableVirtualCircuitRequest`

NewPatchedWritableVirtualCircuitRequestWithDefaults instantiates a new PatchedWritableVirtualCircuitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *PatchedWritableVirtualCircuitRequest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *PatchedWritableVirtualCircuitRequest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *PatchedWritableVirtualCircuitRequest) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *PatchedWritableVirtualCircuitRequest) HasCid() bool`

HasCid returns a boolean if a field has been set.

### GetProviderNetwork

`func (o *PatchedWritableVirtualCircuitRequest) GetProviderNetwork() BriefProviderNetworkRequest`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *PatchedWritableVirtualCircuitRequest) GetProviderNetworkOk() (*BriefProviderNetworkRequest, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *PatchedWritableVirtualCircuitRequest) SetProviderNetwork(v BriefProviderNetworkRequest)`

SetProviderNetwork sets ProviderNetwork field to given value.

### HasProviderNetwork

`func (o *PatchedWritableVirtualCircuitRequest) HasProviderNetwork() bool`

HasProviderNetwork returns a boolean if a field has been set.

### GetProviderAccount

`func (o *PatchedWritableVirtualCircuitRequest) GetProviderAccount() BriefProviderAccountRequest`

GetProviderAccount returns the ProviderAccount field if non-nil, zero value otherwise.

### GetProviderAccountOk

`func (o *PatchedWritableVirtualCircuitRequest) GetProviderAccountOk() (*BriefProviderAccountRequest, bool)`

GetProviderAccountOk returns a tuple with the ProviderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccount

`func (o *PatchedWritableVirtualCircuitRequest) SetProviderAccount(v BriefProviderAccountRequest)`

SetProviderAccount sets ProviderAccount field to given value.

### HasProviderAccount

`func (o *PatchedWritableVirtualCircuitRequest) HasProviderAccount() bool`

HasProviderAccount returns a boolean if a field has been set.

### SetProviderAccountNil

`func (o *PatchedWritableVirtualCircuitRequest) SetProviderAccountNil(b bool)`

 SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil

### UnsetProviderAccount
`func (o *PatchedWritableVirtualCircuitRequest) UnsetProviderAccount()`

UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
### GetType

`func (o *PatchedWritableVirtualCircuitRequest) GetType() BriefVirtualCircuitTypeRequest`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableVirtualCircuitRequest) GetTypeOk() (*BriefVirtualCircuitTypeRequest, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableVirtualCircuitRequest) SetType(v BriefVirtualCircuitTypeRequest)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableVirtualCircuitRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableVirtualCircuitRequest) GetStatus() CircuitStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableVirtualCircuitRequest) GetStatusOk() (*CircuitStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableVirtualCircuitRequest) SetStatus(v CircuitStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableVirtualCircuitRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableVirtualCircuitRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableVirtualCircuitRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableVirtualCircuitRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableVirtualCircuitRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableVirtualCircuitRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableVirtualCircuitRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDescription

`func (o *PatchedWritableVirtualCircuitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableVirtualCircuitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableVirtualCircuitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableVirtualCircuitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableVirtualCircuitRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableVirtualCircuitRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableVirtualCircuitRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableVirtualCircuitRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableVirtualCircuitRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableVirtualCircuitRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableVirtualCircuitRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableVirtualCircuitRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableVirtualCircuitRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableVirtualCircuitRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableVirtualCircuitRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableVirtualCircuitRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


