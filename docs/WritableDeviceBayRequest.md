# WritableDeviceBayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | **int32** |  | 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstalledDevice** | Pointer to **NullableInt32** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableDeviceBayRequest

`func NewWritableDeviceBayRequest(device int32, name string, ) *WritableDeviceBayRequest`

NewWritableDeviceBayRequest instantiates a new WritableDeviceBayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableDeviceBayRequestWithDefaults

`func NewWritableDeviceBayRequestWithDefaults() *WritableDeviceBayRequest`

NewWritableDeviceBayRequestWithDefaults instantiates a new WritableDeviceBayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *WritableDeviceBayRequest) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritableDeviceBayRequest) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritableDeviceBayRequest) SetDevice(v int32)`

SetDevice sets Device field to given value.


### GetName

`func (o *WritableDeviceBayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableDeviceBayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableDeviceBayRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritableDeviceBayRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableDeviceBayRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableDeviceBayRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableDeviceBayRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *WritableDeviceBayRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableDeviceBayRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableDeviceBayRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableDeviceBayRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstalledDevice

`func (o *WritableDeviceBayRequest) GetInstalledDevice() int32`

GetInstalledDevice returns the InstalledDevice field if non-nil, zero value otherwise.

### GetInstalledDeviceOk

`func (o *WritableDeviceBayRequest) GetInstalledDeviceOk() (*int32, bool)`

GetInstalledDeviceOk returns a tuple with the InstalledDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledDevice

`func (o *WritableDeviceBayRequest) SetInstalledDevice(v int32)`

SetInstalledDevice sets InstalledDevice field to given value.

### HasInstalledDevice

`func (o *WritableDeviceBayRequest) HasInstalledDevice() bool`

HasInstalledDevice returns a boolean if a field has been set.

### SetInstalledDeviceNil

`func (o *WritableDeviceBayRequest) SetInstalledDeviceNil(b bool)`

 SetInstalledDeviceNil sets the value for InstalledDevice to be an explicit nil

### UnsetInstalledDevice
`func (o *WritableDeviceBayRequest) UnsetInstalledDevice()`

UnsetInstalledDevice ensures that no value is present for InstalledDevice, not even an explicit nil
### GetTags

`func (o *WritableDeviceBayRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableDeviceBayRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableDeviceBayRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableDeviceBayRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableDeviceBayRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableDeviceBayRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableDeviceBayRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableDeviceBayRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


