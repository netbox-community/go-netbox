# CircuitTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCircuitTypeRequest

`func NewCircuitTypeRequest(name string, slug string, ) *CircuitTypeRequest`

NewCircuitTypeRequest instantiates a new CircuitTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitTypeRequestWithDefaults

`func NewCircuitTypeRequestWithDefaults() *CircuitTypeRequest`

NewCircuitTypeRequestWithDefaults instantiates a new CircuitTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CircuitTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CircuitTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CircuitTypeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *CircuitTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CircuitTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CircuitTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetColor

`func (o *CircuitTypeRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CircuitTypeRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CircuitTypeRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CircuitTypeRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *CircuitTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CircuitTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CircuitTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CircuitTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *CircuitTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CircuitTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CircuitTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CircuitTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *CircuitTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CircuitTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CircuitTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CircuitTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


