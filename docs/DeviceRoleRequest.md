# DeviceRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**VmRole** | Pointer to **bool** | Virtual machines may be assigned to this role | [optional] 
**ConfigTemplate** | Pointer to [**NullableBriefConfigTemplateRequest**](BriefConfigTemplateRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeviceRoleRequest

`func NewDeviceRoleRequest(name string, slug string, ) *DeviceRoleRequest`

NewDeviceRoleRequest instantiates a new DeviceRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRoleRequestWithDefaults

`func NewDeviceRoleRequestWithDefaults() *DeviceRoleRequest`

NewDeviceRoleRequestWithDefaults instantiates a new DeviceRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *DeviceRoleRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DeviceRoleRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DeviceRoleRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetColor

`func (o *DeviceRoleRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DeviceRoleRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DeviceRoleRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DeviceRoleRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetVmRole

`func (o *DeviceRoleRequest) GetVmRole() bool`

GetVmRole returns the VmRole field if non-nil, zero value otherwise.

### GetVmRoleOk

`func (o *DeviceRoleRequest) GetVmRoleOk() (*bool, bool)`

GetVmRoleOk returns a tuple with the VmRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmRole

`func (o *DeviceRoleRequest) SetVmRole(v bool)`

SetVmRole sets VmRole field to given value.

### HasVmRole

`func (o *DeviceRoleRequest) HasVmRole() bool`

HasVmRole returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *DeviceRoleRequest) GetConfigTemplate() BriefConfigTemplateRequest`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *DeviceRoleRequest) GetConfigTemplateOk() (*BriefConfigTemplateRequest, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *DeviceRoleRequest) SetConfigTemplate(v BriefConfigTemplateRequest)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *DeviceRoleRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *DeviceRoleRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *DeviceRoleRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetDescription

`func (o *DeviceRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *DeviceRoleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceRoleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceRoleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceRoleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DeviceRoleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DeviceRoleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DeviceRoleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DeviceRoleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


