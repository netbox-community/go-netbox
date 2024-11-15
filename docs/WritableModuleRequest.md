# WritableModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**BriefDeviceRequest**](BriefDeviceRequest.md) |  | 
**ModuleBay** | **int32** |  | 
**ModuleType** | [**BriefModuleTypeRequest**](BriefModuleTypeRequest.md) |  | 
**Status** | Pointer to [**ModuleStatusValue**](ModuleStatusValue.md) |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableModuleRequest

`func NewWritableModuleRequest(device BriefDeviceRequest, moduleBay int32, moduleType BriefModuleTypeRequest, ) *WritableModuleRequest`

NewWritableModuleRequest instantiates a new WritableModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableModuleRequestWithDefaults

`func NewWritableModuleRequestWithDefaults() *WritableModuleRequest`

NewWritableModuleRequestWithDefaults instantiates a new WritableModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *WritableModuleRequest) GetDevice() BriefDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritableModuleRequest) GetDeviceOk() (*BriefDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritableModuleRequest) SetDevice(v BriefDeviceRequest)`

SetDevice sets Device field to given value.


### GetModuleBay

`func (o *WritableModuleRequest) GetModuleBay() int32`

GetModuleBay returns the ModuleBay field if non-nil, zero value otherwise.

### GetModuleBayOk

`func (o *WritableModuleRequest) GetModuleBayOk() (*int32, bool)`

GetModuleBayOk returns a tuple with the ModuleBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBay

`func (o *WritableModuleRequest) SetModuleBay(v int32)`

SetModuleBay sets ModuleBay field to given value.


### GetModuleType

`func (o *WritableModuleRequest) GetModuleType() BriefModuleTypeRequest`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *WritableModuleRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *WritableModuleRequest) SetModuleType(v BriefModuleTypeRequest)`

SetModuleType sets ModuleType field to given value.


### GetStatus

`func (o *WritableModuleRequest) GetStatus() ModuleStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableModuleRequest) GetStatusOk() (*ModuleStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableModuleRequest) SetStatus(v ModuleStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableModuleRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSerial

`func (o *WritableModuleRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *WritableModuleRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *WritableModuleRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *WritableModuleRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *WritableModuleRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *WritableModuleRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *WritableModuleRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *WritableModuleRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *WritableModuleRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *WritableModuleRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetDescription

`func (o *WritableModuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableModuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableModuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableModuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *WritableModuleRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableModuleRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableModuleRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableModuleRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableModuleRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableModuleRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableModuleRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableModuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableModuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableModuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableModuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableModuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


