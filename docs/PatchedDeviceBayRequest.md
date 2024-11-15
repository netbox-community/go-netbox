# PatchedDeviceBayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**BriefDeviceRequest**](BriefDeviceRequest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstalledDevice** | Pointer to [**NullableBriefDeviceRequest**](BriefDeviceRequest.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedDeviceBayRequest

`func NewPatchedDeviceBayRequest() *PatchedDeviceBayRequest`

NewPatchedDeviceBayRequest instantiates a new PatchedDeviceBayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedDeviceBayRequestWithDefaults

`func NewPatchedDeviceBayRequestWithDefaults() *PatchedDeviceBayRequest`

NewPatchedDeviceBayRequestWithDefaults instantiates a new PatchedDeviceBayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *PatchedDeviceBayRequest) GetDevice() BriefDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedDeviceBayRequest) GetDeviceOk() (*BriefDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedDeviceBayRequest) SetDevice(v BriefDeviceRequest)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedDeviceBayRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetName

`func (o *PatchedDeviceBayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedDeviceBayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedDeviceBayRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedDeviceBayRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedDeviceBayRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedDeviceBayRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedDeviceBayRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedDeviceBayRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedDeviceBayRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedDeviceBayRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedDeviceBayRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedDeviceBayRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstalledDevice

`func (o *PatchedDeviceBayRequest) GetInstalledDevice() BriefDeviceRequest`

GetInstalledDevice returns the InstalledDevice field if non-nil, zero value otherwise.

### GetInstalledDeviceOk

`func (o *PatchedDeviceBayRequest) GetInstalledDeviceOk() (*BriefDeviceRequest, bool)`

GetInstalledDeviceOk returns a tuple with the InstalledDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledDevice

`func (o *PatchedDeviceBayRequest) SetInstalledDevice(v BriefDeviceRequest)`

SetInstalledDevice sets InstalledDevice field to given value.

### HasInstalledDevice

`func (o *PatchedDeviceBayRequest) HasInstalledDevice() bool`

HasInstalledDevice returns a boolean if a field has been set.

### SetInstalledDeviceNil

`func (o *PatchedDeviceBayRequest) SetInstalledDeviceNil(b bool)`

 SetInstalledDeviceNil sets the value for InstalledDevice to be an explicit nil

### UnsetInstalledDevice
`func (o *PatchedDeviceBayRequest) UnsetInstalledDevice()`

UnsetInstalledDevice ensures that no value is present for InstalledDevice, not even an explicit nil
### GetTags

`func (o *PatchedDeviceBayRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedDeviceBayRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedDeviceBayRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedDeviceBayRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedDeviceBayRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedDeviceBayRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedDeviceBayRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedDeviceBayRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


