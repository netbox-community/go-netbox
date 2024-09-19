# WritableIPRangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAddress** | **string** |  | 
**EndAddress** | **string** |  | 
**Vrf** | Pointer to [**NullableBriefVRFRequest**](BriefVRFRequest.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Status** | Pointer to [**PatchedWritableIPRangeRequestStatus**](PatchedWritableIPRangeRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableBriefRoleRequest**](BriefRoleRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**MarkUtilized** | Pointer to **bool** | Treat as fully utilized | [optional] 

## Methods

### NewWritableIPRangeRequest

`func NewWritableIPRangeRequest(startAddress string, endAddress string, ) *WritableIPRangeRequest`

NewWritableIPRangeRequest instantiates a new WritableIPRangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableIPRangeRequestWithDefaults

`func NewWritableIPRangeRequestWithDefaults() *WritableIPRangeRequest`

NewWritableIPRangeRequestWithDefaults instantiates a new WritableIPRangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAddress

`func (o *WritableIPRangeRequest) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *WritableIPRangeRequest) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *WritableIPRangeRequest) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.


### GetEndAddress

`func (o *WritableIPRangeRequest) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *WritableIPRangeRequest) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *WritableIPRangeRequest) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.


### GetVrf

`func (o *WritableIPRangeRequest) GetVrf() BriefVRFRequest`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *WritableIPRangeRequest) GetVrfOk() (*BriefVRFRequest, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *WritableIPRangeRequest) SetVrf(v BriefVRFRequest)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *WritableIPRangeRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *WritableIPRangeRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *WritableIPRangeRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTenant

`func (o *WritableIPRangeRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableIPRangeRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableIPRangeRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableIPRangeRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableIPRangeRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableIPRangeRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *WritableIPRangeRequest) GetStatus() PatchedWritableIPRangeRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableIPRangeRequest) GetStatusOk() (*PatchedWritableIPRangeRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableIPRangeRequest) SetStatus(v PatchedWritableIPRangeRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableIPRangeRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *WritableIPRangeRequest) GetRole() BriefRoleRequest`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *WritableIPRangeRequest) GetRoleOk() (*BriefRoleRequest, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *WritableIPRangeRequest) SetRole(v BriefRoleRequest)`

SetRole sets Role field to given value.

### HasRole

`func (o *WritableIPRangeRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *WritableIPRangeRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *WritableIPRangeRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetDescription

`func (o *WritableIPRangeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableIPRangeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableIPRangeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableIPRangeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *WritableIPRangeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableIPRangeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableIPRangeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableIPRangeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableIPRangeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableIPRangeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableIPRangeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableIPRangeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableIPRangeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableIPRangeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableIPRangeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableIPRangeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetMarkUtilized

`func (o *WritableIPRangeRequest) GetMarkUtilized() bool`

GetMarkUtilized returns the MarkUtilized field if non-nil, zero value otherwise.

### GetMarkUtilizedOk

`func (o *WritableIPRangeRequest) GetMarkUtilizedOk() (*bool, bool)`

GetMarkUtilizedOk returns a tuple with the MarkUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkUtilized

`func (o *WritableIPRangeRequest) SetMarkUtilized(v bool)`

SetMarkUtilized sets MarkUtilized field to given value.

### HasMarkUtilized

`func (o *WritableIPRangeRequest) HasMarkUtilized() bool`

HasMarkUtilized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


