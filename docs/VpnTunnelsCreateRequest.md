# VpnTunnelsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Status** | Pointer to [**PatchedWritableTunnelRequestStatus**](PatchedWritableTunnelRequestStatus.md) |  | [optional] 
**Group** | Pointer to [**NullablePatchedWritableTunnelRequestGroup**](PatchedWritableTunnelRequestGroup.md) |  | [optional] 
**Encapsulation** | [**PatchedWritableTunnelRequestEncapsulation**](PatchedWritableTunnelRequestEncapsulation.md) |  | 
**IpsecProfile** | Pointer to [**NullablePatchedWritableTunnelRequestIpsecProfile**](PatchedWritableTunnelRequestIpsecProfile.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**TunnelId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewVpnTunnelsCreateRequest

`func NewVpnTunnelsCreateRequest(name string, encapsulation PatchedWritableTunnelRequestEncapsulation, ) *VpnTunnelsCreateRequest`

NewVpnTunnelsCreateRequest instantiates a new VpnTunnelsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelsCreateRequestWithDefaults

`func NewVpnTunnelsCreateRequestWithDefaults() *VpnTunnelsCreateRequest`

NewVpnTunnelsCreateRequestWithDefaults instantiates a new VpnTunnelsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VpnTunnelsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnTunnelsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnTunnelsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *VpnTunnelsCreateRequest) GetStatus() PatchedWritableTunnelRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VpnTunnelsCreateRequest) GetStatusOk() (*PatchedWritableTunnelRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VpnTunnelsCreateRequest) SetStatus(v PatchedWritableTunnelRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VpnTunnelsCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGroup

`func (o *VpnTunnelsCreateRequest) GetGroup() PatchedWritableTunnelRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VpnTunnelsCreateRequest) GetGroupOk() (*PatchedWritableTunnelRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VpnTunnelsCreateRequest) SetGroup(v PatchedWritableTunnelRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *VpnTunnelsCreateRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *VpnTunnelsCreateRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *VpnTunnelsCreateRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetEncapsulation

`func (o *VpnTunnelsCreateRequest) GetEncapsulation() PatchedWritableTunnelRequestEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *VpnTunnelsCreateRequest) GetEncapsulationOk() (*PatchedWritableTunnelRequestEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *VpnTunnelsCreateRequest) SetEncapsulation(v PatchedWritableTunnelRequestEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.


### GetIpsecProfile

`func (o *VpnTunnelsCreateRequest) GetIpsecProfile() PatchedWritableTunnelRequestIpsecProfile`

GetIpsecProfile returns the IpsecProfile field if non-nil, zero value otherwise.

### GetIpsecProfileOk

`func (o *VpnTunnelsCreateRequest) GetIpsecProfileOk() (*PatchedWritableTunnelRequestIpsecProfile, bool)`

GetIpsecProfileOk returns a tuple with the IpsecProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecProfile

`func (o *VpnTunnelsCreateRequest) SetIpsecProfile(v PatchedWritableTunnelRequestIpsecProfile)`

SetIpsecProfile sets IpsecProfile field to given value.

### HasIpsecProfile

`func (o *VpnTunnelsCreateRequest) HasIpsecProfile() bool`

HasIpsecProfile returns a boolean if a field has been set.

### SetIpsecProfileNil

`func (o *VpnTunnelsCreateRequest) SetIpsecProfileNil(b bool)`

 SetIpsecProfileNil sets the value for IpsecProfile to be an explicit nil

### UnsetIpsecProfile
`func (o *VpnTunnelsCreateRequest) UnsetIpsecProfile()`

UnsetIpsecProfile ensures that no value is present for IpsecProfile, not even an explicit nil
### GetTenant

`func (o *VpnTunnelsCreateRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VpnTunnelsCreateRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VpnTunnelsCreateRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VpnTunnelsCreateRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VpnTunnelsCreateRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VpnTunnelsCreateRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTunnelId

`func (o *VpnTunnelsCreateRequest) GetTunnelId() int64`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *VpnTunnelsCreateRequest) GetTunnelIdOk() (*int64, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *VpnTunnelsCreateRequest) SetTunnelId(v int64)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *VpnTunnelsCreateRequest) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### SetTunnelIdNil

`func (o *VpnTunnelsCreateRequest) SetTunnelIdNil(b bool)`

 SetTunnelIdNil sets the value for TunnelId to be an explicit nil

### UnsetTunnelId
`func (o *VpnTunnelsCreateRequest) UnsetTunnelId()`

UnsetTunnelId ensures that no value is present for TunnelId, not even an explicit nil
### GetDescription

`func (o *VpnTunnelsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnTunnelsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnTunnelsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpnTunnelsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *VpnTunnelsCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VpnTunnelsCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VpnTunnelsCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VpnTunnelsCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *VpnTunnelsCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *VpnTunnelsCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *VpnTunnelsCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VpnTunnelsCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VpnTunnelsCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VpnTunnelsCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *VpnTunnelsCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpnTunnelsCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpnTunnelsCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpnTunnelsCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VpnTunnelsCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VpnTunnelsCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VpnTunnelsCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VpnTunnelsCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


