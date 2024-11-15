# PatchedDeviceRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**VmRole** | Pointer to **bool** | Virtual machines may be assigned to this role | [optional] 
**ConfigTemplate** | Pointer to [**NullableBriefConfigTemplateRequest**](BriefConfigTemplateRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedDeviceRoleRequest

`func NewPatchedDeviceRoleRequest() *PatchedDeviceRoleRequest`

NewPatchedDeviceRoleRequest instantiates a new PatchedDeviceRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDeviceRoleRequestWithDefaults

`func NewPatchedDeviceRoleRequestWithDefaults() *PatchedDeviceRoleRequest`

NewPatchedDeviceRoleRequestWithDefaults instantiates a new PatchedDeviceRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedDeviceRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDeviceRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDeviceRoleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDeviceRoleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedDeviceRoleRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedDeviceRoleRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedDeviceRoleRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedDeviceRoleRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetColor

`func (o *PatchedDeviceRoleRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchedDeviceRoleRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchedDeviceRoleRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchedDeviceRoleRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetVmRole

`func (o *PatchedDeviceRoleRequest) GetVmRole() bool`

GetVmRole returns the VmRole field if non-nil, zero value otherwise.

### GetVmRoleOk

`func (o *PatchedDeviceRoleRequest) GetVmRoleOk() (*bool, bool)`

GetVmRoleOk returns a tuple with the VmRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRole

`func (o *PatchedDeviceRoleRequest) SetVmRole(v bool)`

SetVmRole sets VmRole field to given value.

### HasVmRole

`func (o *PatchedDeviceRoleRequest) HasVmRole() bool`

HasVmRole returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *PatchedDeviceRoleRequest) GetConfigTemplate() BriefConfigTemplateRequest`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *PatchedDeviceRoleRequest) GetConfigTemplateOk() (*BriefConfigTemplateRequest, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *PatchedDeviceRoleRequest) SetConfigTemplate(v BriefConfigTemplateRequest)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *PatchedDeviceRoleRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *PatchedDeviceRoleRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *PatchedDeviceRoleRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetDescription

`func (o *PatchedDeviceRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedDeviceRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedDeviceRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedDeviceRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedDeviceRoleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedDeviceRoleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedDeviceRoleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedDeviceRoleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedDeviceRoleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedDeviceRoleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedDeviceRoleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedDeviceRoleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


