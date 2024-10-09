# PatchedWritableTunnelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PatchedWritableTunnelRequestStatus**](PatchedWritableTunnelRequestStatus.md) |  | [optional] 
**Group** | Pointer to [**NullableBriefTunnelGroupRequest**](BriefTunnelGroupRequest.md) |  | [optional] 
**Encapsulation** | Pointer to [**PatchedWritableTunnelRequestEncapsulation**](PatchedWritableTunnelRequestEncapsulation.md) |  | [optional] 
**IpsecProfile** | Pointer to [**NullableBriefIPSecProfileRequest**](BriefIPSecProfileRequest.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**TunnelId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableTunnelRequest

`func NewPatchedWritableTunnelRequest() *PatchedWritableTunnelRequest`

NewPatchedWritableTunnelRequest instantiates a new PatchedWritableTunnelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableTunnelRequestWithDefaults

`func NewPatchedWritableTunnelRequestWithDefaults() *PatchedWritableTunnelRequest`

NewPatchedWritableTunnelRequestWithDefaults instantiates a new PatchedWritableTunnelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableTunnelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableTunnelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableTunnelRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableTunnelRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableTunnelRequest) GetStatus() PatchedWritableTunnelRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableTunnelRequest) GetStatusOk() (*PatchedWritableTunnelRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableTunnelRequest) SetStatus(v PatchedWritableTunnelRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableTunnelRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedWritableTunnelRequest) GetGroup() BriefTunnelGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedWritableTunnelRequest) GetGroupOk() (*BriefTunnelGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedWritableTunnelRequest) SetGroup(v BriefTunnelGroupRequest)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedWritableTunnelRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedWritableTunnelRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedWritableTunnelRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetEncapsulation

`func (o *PatchedWritableTunnelRequest) GetEncapsulation() PatchedWritableTunnelRequestEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *PatchedWritableTunnelRequest) GetEncapsulationOk() (*PatchedWritableTunnelRequestEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *PatchedWritableTunnelRequest) SetEncapsulation(v PatchedWritableTunnelRequestEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.

### HasEncapsulation

`func (o *PatchedWritableTunnelRequest) HasEncapsulation() bool`

HasEncapsulation returns a boolean if a field has been set.

### GetIpsecProfile

`func (o *PatchedWritableTunnelRequest) GetIpsecProfile() BriefIPSecProfileRequest`

GetIpsecProfile returns the IpsecProfile field if non-nil, zero value otherwise.

### GetIpsecProfileOk

`func (o *PatchedWritableTunnelRequest) GetIpsecProfileOk() (*BriefIPSecProfileRequest, bool)`

GetIpsecProfileOk returns a tuple with the IpsecProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecProfile

`func (o *PatchedWritableTunnelRequest) SetIpsecProfile(v BriefIPSecProfileRequest)`

SetIpsecProfile sets IpsecProfile field to given value.

### HasIpsecProfile

`func (o *PatchedWritableTunnelRequest) HasIpsecProfile() bool`

HasIpsecProfile returns a boolean if a field has been set.

### SetIpsecProfileNil

`func (o *PatchedWritableTunnelRequest) SetIpsecProfileNil(b bool)`

 SetIpsecProfileNil sets the value for IpsecProfile to be an explicit nil

### UnsetIpsecProfile
`func (o *PatchedWritableTunnelRequest) UnsetIpsecProfile()`

UnsetIpsecProfile ensures that no value is present for IpsecProfile, not even an explicit nil
### GetTenant

`func (o *PatchedWritableTunnelRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableTunnelRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableTunnelRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableTunnelRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableTunnelRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableTunnelRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTunnelId

`func (o *PatchedWritableTunnelRequest) GetTunnelId() int64`

GetTunnelId returns the TunnelId field if non-nil, zero value otherwise.

### GetTunnelIdOk

`func (o *PatchedWritableTunnelRequest) GetTunnelIdOk() (*int64, bool)`

GetTunnelIdOk returns a tuple with the TunnelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelId

`func (o *PatchedWritableTunnelRequest) SetTunnelId(v int64)`

SetTunnelId sets TunnelId field to given value.

### HasTunnelId

`func (o *PatchedWritableTunnelRequest) HasTunnelId() bool`

HasTunnelId returns a boolean if a field has been set.

### SetTunnelIdNil

`func (o *PatchedWritableTunnelRequest) SetTunnelIdNil(b bool)`

 SetTunnelIdNil sets the value for TunnelId to be an explicit nil

### UnsetTunnelId
`func (o *PatchedWritableTunnelRequest) UnsetTunnelId()`

UnsetTunnelId ensures that no value is present for TunnelId, not even an explicit nil
### GetDescription

`func (o *PatchedWritableTunnelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableTunnelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableTunnelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableTunnelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableTunnelRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableTunnelRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableTunnelRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableTunnelRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableTunnelRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableTunnelRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableTunnelRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableTunnelRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableTunnelRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableTunnelRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableTunnelRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableTunnelRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


