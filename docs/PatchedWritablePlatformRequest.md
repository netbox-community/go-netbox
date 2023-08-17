# PatchedWritablePlatformRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **NullableInt32** | Optionally limit this platform to devices of a certain manufacturer | [optional] 
**ConfigTemplate** | Pointer to **NullableInt32** |  | [optional] 
**NapalmDriver** | Pointer to **string** | The name of the NAPALM driver to use when interacting with devices | [optional] 
**NapalmArgs** | Pointer to **map[string]interface{}** | Additional arguments to pass when initiating the NAPALM driver (JSON format) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritablePlatformRequest

`func NewPatchedWritablePlatformRequest() *PatchedWritablePlatformRequest`

NewPatchedWritablePlatformRequest instantiates a new PatchedWritablePlatformRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritablePlatformRequestWithDefaults

`func NewPatchedWritablePlatformRequestWithDefaults() *PatchedWritablePlatformRequest`

NewPatchedWritablePlatformRequestWithDefaults instantiates a new PatchedWritablePlatformRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritablePlatformRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritablePlatformRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritablePlatformRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritablePlatformRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedWritablePlatformRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedWritablePlatformRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedWritablePlatformRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedWritablePlatformRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetManufacturer

`func (o *PatchedWritablePlatformRequest) GetManufacturer() int32`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *PatchedWritablePlatformRequest) GetManufacturerOk() (*int32, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *PatchedWritablePlatformRequest) SetManufacturer(v int32)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *PatchedWritablePlatformRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *PatchedWritablePlatformRequest) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *PatchedWritablePlatformRequest) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetConfigTemplate

`func (o *PatchedWritablePlatformRequest) GetConfigTemplate() int32`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *PatchedWritablePlatformRequest) GetConfigTemplateOk() (*int32, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *PatchedWritablePlatformRequest) SetConfigTemplate(v int32)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *PatchedWritablePlatformRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *PatchedWritablePlatformRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *PatchedWritablePlatformRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetNapalmDriver

`func (o *PatchedWritablePlatformRequest) GetNapalmDriver() string`

GetNapalmDriver returns the NapalmDriver field if non-nil, zero value otherwise.

### GetNapalmDriverOk

`func (o *PatchedWritablePlatformRequest) GetNapalmDriverOk() (*string, bool)`

GetNapalmDriverOk returns a tuple with the NapalmDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNapalmDriver

`func (o *PatchedWritablePlatformRequest) SetNapalmDriver(v string)`

SetNapalmDriver sets NapalmDriver field to given value.

### HasNapalmDriver

`func (o *PatchedWritablePlatformRequest) HasNapalmDriver() bool`

HasNapalmDriver returns a boolean if a field has been set.

### GetNapalmArgs

`func (o *PatchedWritablePlatformRequest) GetNapalmArgs() map[string]interface{}`

GetNapalmArgs returns the NapalmArgs field if non-nil, zero value otherwise.

### GetNapalmArgsOk

`func (o *PatchedWritablePlatformRequest) GetNapalmArgsOk() (*map[string]interface{}, bool)`

GetNapalmArgsOk returns a tuple with the NapalmArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNapalmArgs

`func (o *PatchedWritablePlatformRequest) SetNapalmArgs(v map[string]interface{})`

SetNapalmArgs sets NapalmArgs field to given value.

### HasNapalmArgs

`func (o *PatchedWritablePlatformRequest) HasNapalmArgs() bool`

HasNapalmArgs returns a boolean if a field has been set.

### SetNapalmArgsNil

`func (o *PatchedWritablePlatformRequest) SetNapalmArgsNil(b bool)`

 SetNapalmArgsNil sets the value for NapalmArgs to be an explicit nil

### UnsetNapalmArgs
`func (o *PatchedWritablePlatformRequest) UnsetNapalmArgs()`

UnsetNapalmArgs ensures that no value is present for NapalmArgs, not even an explicit nil
### GetDescription

`func (o *PatchedWritablePlatformRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritablePlatformRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritablePlatformRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritablePlatformRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritablePlatformRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritablePlatformRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritablePlatformRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritablePlatformRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritablePlatformRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritablePlatformRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritablePlatformRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritablePlatformRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


