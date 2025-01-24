# PatchedVLANGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**ScopeType** | Pointer to **NullableString** |  | [optional] 
**ScopeId** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedVLANGroupRequest

`func NewPatchedVLANGroupRequest() *PatchedVLANGroupRequest`

NewPatchedVLANGroupRequest instantiates a new PatchedVLANGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVLANGroupRequestWithDefaults

`func NewPatchedVLANGroupRequestWithDefaults() *PatchedVLANGroupRequest`

NewPatchedVLANGroupRequestWithDefaults instantiates a new PatchedVLANGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedVLANGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedVLANGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedVLANGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedVLANGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedVLANGroupRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedVLANGroupRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedVLANGroupRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedVLANGroupRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetScopeType

`func (o *PatchedVLANGroupRequest) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *PatchedVLANGroupRequest) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *PatchedVLANGroupRequest) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *PatchedVLANGroupRequest) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### SetScopeTypeNil

`func (o *PatchedVLANGroupRequest) SetScopeTypeNil(b bool)`

 SetScopeTypeNil sets the value for ScopeType to be an explicit nil

### UnsetScopeType
`func (o *PatchedVLANGroupRequest) UnsetScopeType()`

UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
### GetScopeId

`func (o *PatchedVLANGroupRequest) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *PatchedVLANGroupRequest) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *PatchedVLANGroupRequest) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *PatchedVLANGroupRequest) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### SetScopeIdNil

`func (o *PatchedVLANGroupRequest) SetScopeIdNil(b bool)`

 SetScopeIdNil sets the value for ScopeId to be an explicit nil

### UnsetScopeId
`func (o *PatchedVLANGroupRequest) UnsetScopeId()`

UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
### GetDescription

`func (o *PatchedVLANGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedVLANGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedVLANGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedVLANGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedVLANGroupRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedVLANGroupRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedVLANGroupRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedVLANGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedVLANGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedVLANGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedVLANGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedVLANGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


