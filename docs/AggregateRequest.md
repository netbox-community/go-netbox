# AggregateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | **string** |  | 
**Rir** | [**BriefRIRRequest**](BriefRIRRequest.md) |  | 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**DateAdded** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAggregateRequest

`func NewAggregateRequest(prefix string, rir BriefRIRRequest, ) *AggregateRequest`

NewAggregateRequest instantiates a new AggregateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateRequestWithDefaults

`func NewAggregateRequestWithDefaults() *AggregateRequest`

NewAggregateRequestWithDefaults instantiates a new AggregateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *AggregateRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AggregateRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AggregateRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetRir

`func (o *AggregateRequest) GetRir() BriefRIRRequest`

GetRir returns the Rir field if non-nil, zero value otherwise.

### GetRirOk

`func (o *AggregateRequest) GetRirOk() (*BriefRIRRequest, bool)`

GetRirOk returns a tuple with the Rir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRir

`func (o *AggregateRequest) SetRir(v BriefRIRRequest)`

SetRir sets Rir field to given value.


### GetTenant

`func (o *AggregateRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AggregateRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AggregateRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AggregateRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *AggregateRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *AggregateRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetDateAdded

`func (o *AggregateRequest) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AggregateRequest) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AggregateRequest) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AggregateRequest) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### SetDateAddedNil

`func (o *AggregateRequest) SetDateAddedNil(b bool)`

 SetDateAddedNil sets the value for DateAdded to be an explicit nil

### UnsetDateAdded
`func (o *AggregateRequest) UnsetDateAdded()`

UnsetDateAdded ensures that no value is present for DateAdded, not even an explicit nil
### GetDescription

`func (o *AggregateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AggregateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AggregateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AggregateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *AggregateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AggregateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AggregateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *AggregateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *AggregateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AggregateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AggregateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AggregateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *AggregateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AggregateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AggregateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AggregateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


