# PatchedWritableConsolePortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**BriefDeviceRequest**](BriefDeviceRequest.md) |  | [optional] 
**Module** | Pointer to [**NullableBriefModuleRequest**](BriefModuleRequest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**PatchedWritableConsolePortRequestType**](PatchedWritableConsolePortRequestType.md) |  | [optional] 
**Speed** | Pointer to [**NullablePatchedWritableConsolePortRequestSpeed**](PatchedWritableConsolePortRequestSpeed.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableConsolePortRequest

`func NewPatchedWritableConsolePortRequest() *PatchedWritableConsolePortRequest`

NewPatchedWritableConsolePortRequest instantiates a new PatchedWritableConsolePortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableConsolePortRequestWithDefaults

`func NewPatchedWritableConsolePortRequestWithDefaults() *PatchedWritableConsolePortRequest`

NewPatchedWritableConsolePortRequestWithDefaults instantiates a new PatchedWritableConsolePortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *PatchedWritableConsolePortRequest) GetDevice() BriefDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedWritableConsolePortRequest) GetDeviceOk() (*BriefDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedWritableConsolePortRequest) SetDevice(v BriefDeviceRequest)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedWritableConsolePortRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetModule

`func (o *PatchedWritableConsolePortRequest) GetModule() BriefModuleRequest`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *PatchedWritableConsolePortRequest) GetModuleOk() (*BriefModuleRequest, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *PatchedWritableConsolePortRequest) SetModule(v BriefModuleRequest)`

SetModule sets Module field to given value.

### HasModule

`func (o *PatchedWritableConsolePortRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *PatchedWritableConsolePortRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *PatchedWritableConsolePortRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *PatchedWritableConsolePortRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableConsolePortRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableConsolePortRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableConsolePortRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableConsolePortRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableConsolePortRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableConsolePortRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableConsolePortRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableConsolePortRequest) GetType() PatchedWritableConsolePortRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableConsolePortRequest) GetTypeOk() (*PatchedWritableConsolePortRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableConsolePortRequest) SetType(v PatchedWritableConsolePortRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableConsolePortRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSpeed

`func (o *PatchedWritableConsolePortRequest) GetSpeed() PatchedWritableConsolePortRequestSpeed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *PatchedWritableConsolePortRequest) GetSpeedOk() (*PatchedWritableConsolePortRequestSpeed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *PatchedWritableConsolePortRequest) SetSpeed(v PatchedWritableConsolePortRequestSpeed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *PatchedWritableConsolePortRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *PatchedWritableConsolePortRequest) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *PatchedWritableConsolePortRequest) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetDescription

`func (o *PatchedWritableConsolePortRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableConsolePortRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableConsolePortRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableConsolePortRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *PatchedWritableConsolePortRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *PatchedWritableConsolePortRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *PatchedWritableConsolePortRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *PatchedWritableConsolePortRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableConsolePortRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableConsolePortRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableConsolePortRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableConsolePortRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableConsolePortRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableConsolePortRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableConsolePortRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableConsolePortRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


