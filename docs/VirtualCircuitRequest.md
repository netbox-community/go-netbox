# VirtualCircuitRequest

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

### NewVirtualCircuitRequest

`func NewVirtualCircuitRequest(cid string, providerNetwork BriefProviderNetworkRequest, type_ BriefVirtualCircuitTypeRequest, ) *VirtualCircuitRequest`

NewVirtualCircuitRequest instantiates a new VirtualCircuitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCircuitRequestWithDefaults

`func NewVirtualCircuitRequestWithDefaults() *VirtualCircuitRequest`

NewVirtualCircuitRequestWithDefaults instantiates a new VirtualCircuitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *VirtualCircuitRequest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *VirtualCircuitRequest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *VirtualCircuitRequest) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProviderNetwork

`func (o *VirtualCircuitRequest) GetProviderNetwork() BriefProviderNetworkRequest`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *VirtualCircuitRequest) GetProviderNetworkOk() (*BriefProviderNetworkRequest, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *VirtualCircuitRequest) SetProviderNetwork(v BriefProviderNetworkRequest)`

SetProviderNetwork sets ProviderNetwork field to given value.


### GetProviderAccount

`func (o *VirtualCircuitRequest) GetProviderAccount() BriefProviderAccountRequest`

GetProviderAccount returns the ProviderAccount field if non-nil, zero value otherwise.

### GetProviderAccountOk

`func (o *VirtualCircuitRequest) GetProviderAccountOk() (*BriefProviderAccountRequest, bool)`

GetProviderAccountOk returns a tuple with the ProviderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccount

`func (o *VirtualCircuitRequest) SetProviderAccount(v BriefProviderAccountRequest)`

SetProviderAccount sets ProviderAccount field to given value.

### HasProviderAccount

`func (o *VirtualCircuitRequest) HasProviderAccount() bool`

HasProviderAccount returns a boolean if a field has been set.

### SetProviderAccountNil

`func (o *VirtualCircuitRequest) SetProviderAccountNil(b bool)`

 SetProviderAccountNil sets the value for ProviderAccount to be an explicit nil

### UnsetProviderAccount
`func (o *VirtualCircuitRequest) UnsetProviderAccount()`

UnsetProviderAccount ensures that no value is present for ProviderAccount, not even an explicit nil
### GetType

`func (o *VirtualCircuitRequest) GetType() BriefVirtualCircuitTypeRequest`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualCircuitRequest) GetTypeOk() (*BriefVirtualCircuitTypeRequest, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualCircuitRequest) SetType(v BriefVirtualCircuitTypeRequest)`

SetType sets Type field to given value.


### GetStatus

`func (o *VirtualCircuitRequest) GetStatus() CircuitStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualCircuitRequest) GetStatusOk() (*CircuitStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualCircuitRequest) SetStatus(v CircuitStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualCircuitRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *VirtualCircuitRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VirtualCircuitRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VirtualCircuitRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VirtualCircuitRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VirtualCircuitRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VirtualCircuitRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDescription

`func (o *VirtualCircuitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualCircuitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualCircuitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualCircuitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *VirtualCircuitRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VirtualCircuitRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VirtualCircuitRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VirtualCircuitRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *VirtualCircuitRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualCircuitRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualCircuitRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualCircuitRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualCircuitRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualCircuitRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualCircuitRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualCircuitRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


