# PatchedASNRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int64** | 16- or 32-bit autonomous system number | [optional] 
**Rir** | Pointer to [**NullableBriefRIRRequest**](BriefRIRRequest.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedASNRequest

`func NewPatchedASNRequest() *PatchedASNRequest`

NewPatchedASNRequest instantiates a new PatchedASNRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedASNRequestWithDefaults

`func NewPatchedASNRequestWithDefaults() *PatchedASNRequest`

NewPatchedASNRequestWithDefaults instantiates a new PatchedASNRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *PatchedASNRequest) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *PatchedASNRequest) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *PatchedASNRequest) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *PatchedASNRequest) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetRir

`func (o *PatchedASNRequest) GetRir() BriefRIRRequest`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *PatchedASNRequest) GetRirOk() (*BriefRIRRequest, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *PatchedASNRequest) SetRir(v BriefRIRRequest)`

SetRir sets Rir field to given value.

### HasRir

`func (o *PatchedASNRequest) HasRir() bool`

HasRir returns a boolean if a field has been set.

### SetRirNil

`func (o *PatchedASNRequest) SetRirNil(b bool)`

 SetRirNil sets the value for Rir to be an explicit nil

### UnsetRir
`func (o *PatchedASNRequest) UnsetRir()`

UnsetRir ensures that no value is present for Rir, not even an explicit nil
### GetTenant

`func (o *PatchedASNRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedASNRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedASNRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedASNRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedASNRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedASNRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDescription

`func (o *PatchedASNRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedASNRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedASNRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedASNRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedASNRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedASNRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedASNRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedASNRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedASNRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedASNRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedASNRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedASNRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedASNRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedASNRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedASNRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedASNRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


