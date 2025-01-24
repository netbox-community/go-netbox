# VirtualCircuitTypeRequest

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

### NewVirtualCircuitTypeRequest

`func NewVirtualCircuitTypeRequest(name string, slug string, ) *VirtualCircuitTypeRequest`

NewVirtualCircuitTypeRequest instantiates a new VirtualCircuitTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCircuitTypeRequestWithDefaults

`func NewVirtualCircuitTypeRequestWithDefaults() *VirtualCircuitTypeRequest`

NewVirtualCircuitTypeRequestWithDefaults instantiates a new VirtualCircuitTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualCircuitTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCircuitTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCircuitTypeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *VirtualCircuitTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *VirtualCircuitTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *VirtualCircuitTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetColor

`func (o *VirtualCircuitTypeRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *VirtualCircuitTypeRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *VirtualCircuitTypeRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *VirtualCircuitTypeRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualCircuitTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualCircuitTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualCircuitTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualCircuitTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *VirtualCircuitTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualCircuitTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualCircuitTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualCircuitTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualCircuitTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualCircuitTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualCircuitTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualCircuitTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


