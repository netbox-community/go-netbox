# VLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to [**NullableBriefSiteRequest**](BriefSiteRequest.md) |  | [optional] 
**Group** | Pointer to [**NullableBriefVLANGroupRequest**](BriefVLANGroupRequest.md) |  | [optional] 
**Vid** | **int32** | Numeric VLAN ID (1-4094) | 
**Name** | **string** |  | 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Status** | Pointer to [**IPRangeStatusValue**](IPRangeStatusValue.md) |  | [optional] 
**Role** | Pointer to [**NullableBriefRoleRequest**](BriefRoleRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVLANRequest

`func NewVLANRequest(vid int32, name string, ) *VLANRequest`

NewVLANRequest instantiates a new VLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVLANRequestWithDefaults

`func NewVLANRequestWithDefaults() *VLANRequest`

NewVLANRequestWithDefaults instantiates a new VLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *VLANRequest) GetSite() BriefSiteRequest`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *VLANRequest) GetSiteOk() (*BriefSiteRequest, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *VLANRequest) SetSite(v BriefSiteRequest)`

SetSite sets Site field to given value.

### HasSite

`func (o *VLANRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *VLANRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *VLANRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetGroup

`func (o *VLANRequest) GetGroup() BriefVLANGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VLANRequest) GetGroupOk() (*BriefVLANGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VLANRequest) SetGroup(v BriefVLANGroupRequest)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *VLANRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *VLANRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *VLANRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetVid

`func (o *VLANRequest) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *VLANRequest) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *VLANRequest) SetVid(v int32)`

SetVid sets Vid field to given value.


### GetName

`func (o *VLANRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VLANRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VLANRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTenant

`func (o *VLANRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VLANRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VLANRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VLANRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VLANRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VLANRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *VLANRequest) GetStatus() IPRangeStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VLANRequest) GetStatusOk() (*IPRangeStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VLANRequest) SetStatus(v IPRangeStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VLANRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *VLANRequest) GetRole() BriefRoleRequest`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VLANRequest) GetRoleOk() (*BriefRoleRequest, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VLANRequest) SetRole(v BriefRoleRequest)`

SetRole sets Role field to given value.

### HasRole

`func (o *VLANRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VLANRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VLANRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetDescription

`func (o *VLANRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VLANRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VLANRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VLANRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *VLANRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VLANRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VLANRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VLANRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *VLANRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VLANRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VLANRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VLANRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VLANRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VLANRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VLANRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VLANRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


