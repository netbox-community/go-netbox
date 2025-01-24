# TunnelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Status** | [**PatchedWritableTunnelRequestStatus**](PatchedWritableTunnelRequestStatus.md) |  | 
**Group** | Pointer to [**NullableBriefTunnelGroupRequest**](BriefTunnelGroupRequest.md) |  | [optional] 
**Encapsulation** | [**PatchedWritableTunnelRequestEncapsulation**](PatchedWritableTunnelRequestEncapsulation.md) |  | 
**IpsecProfile** | Pointer to [**NullableBriefIPSecProfileRequest**](BriefIPSecProfileRequest.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**TunnelId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTunnelRequest

`func NewTunnelRequest(name string, status PatchedWritableTunnelRequestStatus, encapsulation PatchedWritableTunnelRequestEncapsulation, ) *TunnelRequest`

NewTunnelRequest instantiates a new TunnelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelRequestWithDefaults

`func NewTunnelRequestWithDefaults() *TunnelRequest`

NewTunnelRequestWithDefaults instantiates a new TunnelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TunnelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TunnelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TunnelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *TunnelRequest) GetStatus() PatchedWritableTunnelRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TunnelRequest) GetStatusOk() (*PatchedWritableTunnelRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TunnelRequest) SetStatus(v PatchedWritableTunnelRequestStatus)`

SetStatus sets Status field to given value.


### GetGroup

`func (o *TunnelRequest) GetGroup() BriefTunnelGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *TunnelRequest) GetGroupOk() (*BriefTunnelGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *TunnelRequest) SetGroup(v BriefTunnelGroupRequest)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *TunnelRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *TunnelRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *TunnelRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetEncapsulation

`func (o *TunnelRequest) GetEncapsulation() PatchedWritableTunnelRequestEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *TunnelRequest) GetEncapsulationOk() (*PatchedWritableTunnelRequestEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *TunnelRequest) SetEncapsulation(v PatchedWritableTunnelRequestEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.


### GetIpsecProfile

`func (o *TunnelRequest) GetIpsecProfile() BriefIPSecProfileRequest`

GetIpsecProfile returns the IpsecProfile field if non-nil, zero value otherwise.

### GetIpsecProfileOk

`func (o *TunnelRequest) GetIpsecProfileOk() (*BriefIPSecProfileRequest, bool)`

GetIpsecProfileOk returns a tuple with the IpsecProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecProfile

`func (o *TunnelRequest) SetIpsecProfile(v BriefIPSecProfileRequest)`

SetIpsecProfile sets IpsecProfile field to given value.

### HasIpsecProfile

`func (o *TunnelRequest) HasIpsecProfile() bool`

HasIpsecProfile returns a boolean if a field has been set.

### SetIpsecProfileNil

`func (o *TunnelRequest) SetIpsecProfileNil(b bool)`

 SetIpsecProfileNil sets the value for IpsecProfile to be an explicit nil

### UnsetIpsecProfile
`func (o *TunnelRequest) UnsetIpsecProfile()`

UnsetIpsecProfile ensures that no value is present for IpsecProfile, not even an explicit nil
### GetTenant

`func (o *TunnelRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TunnelRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TunnelRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TunnelRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *TunnelRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *TunnelRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTunnelId

`func (o *TunnelRequest) GetTunnelId() int64`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *TunnelRequest) GetTunnelIdOk() (*int64, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *TunnelRequest) SetTunnelId(v int64)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *TunnelRequest) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### SetTunnelIdNil

`func (o *TunnelRequest) SetTunnelIdNil(b bool)`

 SetTunnelIdNil sets the value for TunnelId to be an explicit nil

### UnsetTunnelId
`func (o *TunnelRequest) UnsetTunnelId()`

UnsetTunnelId ensures that no value is present for TunnelId, not even an explicit nil
### GetDescription

`func (o *TunnelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TunnelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TunnelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TunnelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *TunnelRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *TunnelRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *TunnelRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *TunnelRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *TunnelRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TunnelRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TunnelRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TunnelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *TunnelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TunnelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TunnelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TunnelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


