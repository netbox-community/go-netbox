# PatchedWritableModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**BriefDeviceRequest**](BriefDeviceRequest.md) |  | [optional] 
**ModuleBay** | Pointer to **int32** |  | [optional] 
**ModuleType** | Pointer to [**BriefModuleTypeRequest**](BriefModuleTypeRequest.md) |  | [optional] 
**Status** | Pointer to [**ModuleStatusValue**](ModuleStatusValue.md) |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableModuleRequest

`func NewPatchedWritableModuleRequest() *PatchedWritableModuleRequest`

NewPatchedWritableModuleRequest instantiates a new PatchedWritableModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableModuleRequestWithDefaults

`func NewPatchedWritableModuleRequestWithDefaults() *PatchedWritableModuleRequest`

NewPatchedWritableModuleRequestWithDefaults instantiates a new PatchedWritableModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *PatchedWritableModuleRequest) GetDevice() BriefDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedWritableModuleRequest) GetDeviceOk() (*BriefDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedWritableModuleRequest) SetDevice(v BriefDeviceRequest)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedWritableModuleRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetModuleBay

`func (o *PatchedWritableModuleRequest) GetModuleBay() int32`

GetModuleBay returns the ModuleBay field if non-nil, zero value otherwise.

### GetModuleBayOk

`func (o *PatchedWritableModuleRequest) GetModuleBayOk() (*int32, bool)`

GetModuleBayOk returns a tuple with the ModuleBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBay

`func (o *PatchedWritableModuleRequest) SetModuleBay(v int32)`

SetModuleBay sets ModuleBay field to given value.

### HasModuleBay

`func (o *PatchedWritableModuleRequest) HasModuleBay() bool`

HasModuleBay returns a boolean if a field has been set.

### GetModuleType

`func (o *PatchedWritableModuleRequest) GetModuleType() BriefModuleTypeRequest`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *PatchedWritableModuleRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *PatchedWritableModuleRequest) SetModuleType(v BriefModuleTypeRequest)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *PatchedWritableModuleRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableModuleRequest) GetStatus() ModuleStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableModuleRequest) GetStatusOk() (*ModuleStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableModuleRequest) SetStatus(v ModuleStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableModuleRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSerial

`func (o *PatchedWritableModuleRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedWritableModuleRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedWritableModuleRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedWritableModuleRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedWritableModuleRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedWritableModuleRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedWritableModuleRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedWritableModuleRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedWritableModuleRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedWritableModuleRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetDescription

`func (o *PatchedWritableModuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableModuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableModuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableModuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableModuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableModuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableModuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableModuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableModuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableModuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableModuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableModuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableModuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableModuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableModuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableModuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


