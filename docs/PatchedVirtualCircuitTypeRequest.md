# PatchedVirtualCircuitTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedVirtualCircuitTypeRequest

`func NewPatchedVirtualCircuitTypeRequest() *PatchedVirtualCircuitTypeRequest`

NewPatchedVirtualCircuitTypeRequest instantiates a new PatchedVirtualCircuitTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVirtualCircuitTypeRequestWithDefaults

`func NewPatchedVirtualCircuitTypeRequestWithDefaults() *PatchedVirtualCircuitTypeRequest`

NewPatchedVirtualCircuitTypeRequestWithDefaults instantiates a new PatchedVirtualCircuitTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedVirtualCircuitTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedVirtualCircuitTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedVirtualCircuitTypeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedVirtualCircuitTypeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedVirtualCircuitTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedVirtualCircuitTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedVirtualCircuitTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedVirtualCircuitTypeRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetColor

`func (o *PatchedVirtualCircuitTypeRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchedVirtualCircuitTypeRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchedVirtualCircuitTypeRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchedVirtualCircuitTypeRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedVirtualCircuitTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedVirtualCircuitTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedVirtualCircuitTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedVirtualCircuitTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedVirtualCircuitTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedVirtualCircuitTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedVirtualCircuitTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedVirtualCircuitTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedVirtualCircuitTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedVirtualCircuitTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedVirtualCircuitTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedVirtualCircuitTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


