# PatchedWritableAggregateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** |  | [optional] 
**Rir** | Pointer to [**BriefRIRRequest**](BriefRIRRequest.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**DateAdded** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableAggregateRequest

`func NewPatchedWritableAggregateRequest() *PatchedWritableAggregateRequest`

NewPatchedWritableAggregateRequest instantiates a new PatchedWritableAggregateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableAggregateRequestWithDefaults

`func NewPatchedWritableAggregateRequestWithDefaults() *PatchedWritableAggregateRequest`

NewPatchedWritableAggregateRequestWithDefaults instantiates a new PatchedWritableAggregateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *PatchedWritableAggregateRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PatchedWritableAggregateRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PatchedWritableAggregateRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PatchedWritableAggregateRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRir

`func (o *PatchedWritableAggregateRequest) GetRir() BriefRIRRequest`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *PatchedWritableAggregateRequest) GetRirOk() (*BriefRIRRequest, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *PatchedWritableAggregateRequest) SetRir(v BriefRIRRequest)`

SetRir sets Rir field to given value.

### HasRir

`func (o *PatchedWritableAggregateRequest) HasRir() bool`

HasRir returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableAggregateRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableAggregateRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableAggregateRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableAggregateRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableAggregateRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableAggregateRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDateAdded

`func (o *PatchedWritableAggregateRequest) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *PatchedWritableAggregateRequest) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *PatchedWritableAggregateRequest) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *PatchedWritableAggregateRequest) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### SetDateAddedNil

`func (o *PatchedWritableAggregateRequest) SetDateAddedNil(b bool)`

 SetDateAddedNil sets the value for DateAdded to be an explicit nil

### UnsetDateAdded
`func (o *PatchedWritableAggregateRequest) UnsetDateAdded()`

UnsetDateAdded ensures that no value is present for DateAdded, not even an explicit nil
### GetDescription

`func (o *PatchedWritableAggregateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableAggregateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableAggregateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableAggregateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableAggregateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableAggregateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableAggregateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableAggregateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableAggregateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableAggregateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableAggregateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableAggregateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableAggregateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableAggregateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableAggregateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableAggregateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


