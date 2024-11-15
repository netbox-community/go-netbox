# PatchedWritableClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**BriefClusterTypeRequest**](BriefClusterTypeRequest.md) |  | [optional] 
**Group** | Pointer to [**NullableBriefClusterGroupRequest**](BriefClusterGroupRequest.md) |  | [optional] 
**Status** | Pointer to [**ClusterStatusValue**](ClusterStatusValue.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Site** | Pointer to [**NullableBriefSiteRequest**](BriefSiteRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableClusterRequest

`func NewPatchedWritableClusterRequest() *PatchedWritableClusterRequest`

NewPatchedWritableClusterRequest instantiates a new PatchedWritableClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableClusterRequestWithDefaults

`func NewPatchedWritableClusterRequestWithDefaults() *PatchedWritableClusterRequest`

NewPatchedWritableClusterRequestWithDefaults instantiates a new PatchedWritableClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableClusterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableClusterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableClusterRequest) GetType() BriefClusterTypeRequest`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableClusterRequest) GetTypeOk() (*BriefClusterTypeRequest, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableClusterRequest) SetType(v BriefClusterTypeRequest)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableClusterRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedWritableClusterRequest) GetGroup() BriefClusterGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedWritableClusterRequest) GetGroupOk() (*BriefClusterGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedWritableClusterRequest) SetGroup(v BriefClusterGroupRequest)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedWritableClusterRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedWritableClusterRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedWritableClusterRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetStatus

`func (o *PatchedWritableClusterRequest) GetStatus() ClusterStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableClusterRequest) GetStatusOk() (*ClusterStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableClusterRequest) SetStatus(v ClusterStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableClusterRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableClusterRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableClusterRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableClusterRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableClusterRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableClusterRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableClusterRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetSite

`func (o *PatchedWritableClusterRequest) GetSite() BriefSiteRequest`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableClusterRequest) GetSiteOk() (*BriefSiteRequest, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableClusterRequest) SetSite(v BriefSiteRequest)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableClusterRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *PatchedWritableClusterRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *PatchedWritableClusterRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetDescription

`func (o *PatchedWritableClusterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableClusterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableClusterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableClusterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableClusterRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableClusterRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableClusterRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableClusterRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableClusterRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableClusterRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableClusterRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableClusterRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableClusterRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableClusterRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableClusterRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableClusterRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


