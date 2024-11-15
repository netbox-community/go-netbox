# FrontPortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**BriefDeviceRequest**](BriefDeviceRequest.md) |  | 
**Module** | Pointer to [**NullableBriefModuleRequest**](BriefModuleRequest.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | [**FrontPortTypeValue**](FrontPortTypeValue.md) |  | 
**Color** | Pointer to **string** |  | [optional] 
**RearPort** | [**FrontPortRearPortRequest**](FrontPortRearPortRequest.md) |  | 
**RearPortPosition** | Pointer to **int32** | Mapped position on corresponding rear port | [optional] [default to 1]
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewFrontPortRequest

`func NewFrontPortRequest(device BriefDeviceRequest, name string, type_ FrontPortTypeValue, rearPort FrontPortRearPortRequest, ) *FrontPortRequest`

NewFrontPortRequest instantiates a new FrontPortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrontPortRequestWithDefaults

`func NewFrontPortRequestWithDefaults() *FrontPortRequest`

NewFrontPortRequestWithDefaults instantiates a new FrontPortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *FrontPortRequest) GetDevice() BriefDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FrontPortRequest) GetDeviceOk() (*BriefDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FrontPortRequest) SetDevice(v BriefDeviceRequest)`

SetDevice sets Device field to given value.


### GetModule

`func (o *FrontPortRequest) GetModule() BriefModuleRequest`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *FrontPortRequest) GetModuleOk() (*BriefModuleRequest, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *FrontPortRequest) SetModule(v BriefModuleRequest)`

SetModule sets Module field to given value.

### HasModule

`func (o *FrontPortRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *FrontPortRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *FrontPortRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *FrontPortRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrontPortRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrontPortRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *FrontPortRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FrontPortRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FrontPortRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FrontPortRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *FrontPortRequest) GetType() FrontPortTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FrontPortRequest) GetTypeOk() (*FrontPortTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FrontPortRequest) SetType(v FrontPortTypeValue)`

SetType sets Type field to given value.


### GetColor

`func (o *FrontPortRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *FrontPortRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *FrontPortRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *FrontPortRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetRearPort

`func (o *FrontPortRequest) GetRearPort() FrontPortRearPortRequest`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *FrontPortRequest) GetRearPortOk() (*FrontPortRearPortRequest, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *FrontPortRequest) SetRearPort(v FrontPortRearPortRequest)`

SetRearPort sets RearPort field to given value.


### GetRearPortPosition

`func (o *FrontPortRequest) GetRearPortPosition() int32`

GetRearPortPosition returns the RearPortPosition field if non-nil, zero value otherwise.

### GetRearPortPositionOk

`func (o *FrontPortRequest) GetRearPortPositionOk() (*int32, bool)`

GetRearPortPositionOk returns a tuple with the RearPortPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortPosition

`func (o *FrontPortRequest) SetRearPortPosition(v int32)`

SetRearPortPosition sets RearPortPosition field to given value.

### HasRearPortPosition

`func (o *FrontPortRequest) HasRearPortPosition() bool`

HasRearPortPosition returns a boolean if a field has been set.

### GetDescription

`func (o *FrontPortRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FrontPortRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FrontPortRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FrontPortRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *FrontPortRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *FrontPortRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *FrontPortRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *FrontPortRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetTags

`func (o *FrontPortRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FrontPortRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FrontPortRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FrontPortRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *FrontPortRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *FrontPortRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *FrontPortRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *FrontPortRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


