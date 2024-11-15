# ASNRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | **int64** | 16- or 32-bit autonomous system number | 
**Rir** | Pointer to [**NullableBriefRIRRequest**](BriefRIRRequest.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewASNRequest

`func NewASNRequest(asn int64, ) *ASNRequest`

NewASNRequest instantiates a new ASNRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASNRequestWithDefaults

`func NewASNRequestWithDefaults() *ASNRequest`

NewASNRequestWithDefaults instantiates a new ASNRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *ASNRequest) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ASNRequest) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ASNRequest) SetAsn(v int64)`

SetAsn sets Asn field to given value.


### GetRir

`func (o *ASNRequest) GetRir() BriefRIRRequest`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *ASNRequest) GetRirOk() (*BriefRIRRequest, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *ASNRequest) SetRir(v BriefRIRRequest)`

SetRir sets Rir field to given value.

### HasRir

`func (o *ASNRequest) HasRir() bool`

HasRir returns a boolean if a field has been set.

### SetRirNil

`func (o *ASNRequest) SetRirNil(b bool)`

 SetRirNil sets the value for Rir to be an explicit nil

### UnsetRir
`func (o *ASNRequest) UnsetRir()`

UnsetRir ensures that no value is present for Rir, not even an explicit nil
### GetTenant

`func (o *ASNRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ASNRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ASNRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ASNRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *ASNRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *ASNRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDescription

`func (o *ASNRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ASNRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ASNRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ASNRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *ASNRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ASNRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ASNRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ASNRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *ASNRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ASNRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ASNRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ASNRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ASNRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ASNRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ASNRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ASNRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


