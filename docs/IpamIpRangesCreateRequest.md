# IpamIpRangesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAddress** | **string** |  | 
**EndAddress** | **string** |  | 
**Vrf** | Pointer to [**NullableIPAddressRequestVrf**](IPAddressRequestVrf.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Status** | Pointer to [**PatchedWritableIPRangeRequestStatus**](PatchedWritableIPRangeRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableIPRangeRequestRole**](IPRangeRequestRole.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**MarkPopulated** | Pointer to **bool** | Prevent the creation of IP addresses within this range | [optional] 
**MarkUtilized** | Pointer to **bool** | Report space as fully utilized | [optional] 

## Methods

### NewIpamIpRangesCreateRequest

`func NewIpamIpRangesCreateRequest(startAddress string, endAddress string, ) *IpamIpRangesCreateRequest`

NewIpamIpRangesCreateRequest instantiates a new IpamIpRangesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamIpRangesCreateRequestWithDefaults

`func NewIpamIpRangesCreateRequestWithDefaults() *IpamIpRangesCreateRequest`

NewIpamIpRangesCreateRequestWithDefaults instantiates a new IpamIpRangesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAddress

`func (o *IpamIpRangesCreateRequest) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *IpamIpRangesCreateRequest) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *IpamIpRangesCreateRequest) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.


### GetEndAddress

`func (o *IpamIpRangesCreateRequest) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *IpamIpRangesCreateRequest) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *IpamIpRangesCreateRequest) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.


### GetVrf

`func (o *IpamIpRangesCreateRequest) GetVrf() IPAddressRequestVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IpamIpRangesCreateRequest) GetVrfOk() (*IPAddressRequestVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IpamIpRangesCreateRequest) SetVrf(v IPAddressRequestVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IpamIpRangesCreateRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *IpamIpRangesCreateRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *IpamIpRangesCreateRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTenant

`func (o *IpamIpRangesCreateRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IpamIpRangesCreateRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IpamIpRangesCreateRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IpamIpRangesCreateRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *IpamIpRangesCreateRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *IpamIpRangesCreateRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *IpamIpRangesCreateRequest) GetStatus() PatchedWritableIPRangeRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IpamIpRangesCreateRequest) GetStatusOk() (*PatchedWritableIPRangeRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IpamIpRangesCreateRequest) SetStatus(v PatchedWritableIPRangeRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IpamIpRangesCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *IpamIpRangesCreateRequest) GetRole() IPRangeRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IpamIpRangesCreateRequest) GetRoleOk() (*IPRangeRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IpamIpRangesCreateRequest) SetRole(v IPRangeRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *IpamIpRangesCreateRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *IpamIpRangesCreateRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *IpamIpRangesCreateRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetDescription

`func (o *IpamIpRangesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpamIpRangesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpamIpRangesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpamIpRangesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *IpamIpRangesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IpamIpRangesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IpamIpRangesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IpamIpRangesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *IpamIpRangesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *IpamIpRangesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *IpamIpRangesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *IpamIpRangesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *IpamIpRangesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *IpamIpRangesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *IpamIpRangesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamIpRangesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamIpRangesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamIpRangesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IpamIpRangesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IpamIpRangesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IpamIpRangesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IpamIpRangesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetMarkPopulated

`func (o *IpamIpRangesCreateRequest) GetMarkPopulated() bool`

GetMarkPopulated returns the MarkPopulated field if non-nil, zero value otherwise.

### GetMarkPopulatedOk

`func (o *IpamIpRangesCreateRequest) GetMarkPopulatedOk() (*bool, bool)`

GetMarkPopulatedOk returns a tuple with the MarkPopulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPopulated

`func (o *IpamIpRangesCreateRequest) SetMarkPopulated(v bool)`

SetMarkPopulated sets MarkPopulated field to given value.

### HasMarkPopulated

`func (o *IpamIpRangesCreateRequest) HasMarkPopulated() bool`

HasMarkPopulated returns a boolean if a field has been set.

### GetMarkUtilized

`func (o *IpamIpRangesCreateRequest) GetMarkUtilized() bool`

GetMarkUtilized returns the MarkUtilized field if non-nil, zero value otherwise.

### GetMarkUtilizedOk

`func (o *IpamIpRangesCreateRequest) GetMarkUtilizedOk() (*bool, bool)`

GetMarkUtilizedOk returns a tuple with the MarkUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkUtilized

`func (o *IpamIpRangesCreateRequest) SetMarkUtilized(v bool)`

SetMarkUtilized sets MarkUtilized field to given value.

### HasMarkUtilized

`func (o *IpamIpRangesCreateRequest) HasMarkUtilized() bool`

HasMarkUtilized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


