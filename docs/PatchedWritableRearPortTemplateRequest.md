# PatchedWritableRearPortTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**NullableBriefDeviceTypeRequest**](BriefDeviceTypeRequest.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBriefModuleTypeRequest**](BriefModuleTypeRequest.md) |  | [optional] 
**Name** | Pointer to **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**FrontPortTypeValue**](FrontPortTypeValue.md) |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Positions** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedWritableRearPortTemplateRequest

`func NewPatchedWritableRearPortTemplateRequest() *PatchedWritableRearPortTemplateRequest`

NewPatchedWritableRearPortTemplateRequest instantiates a new PatchedWritableRearPortTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableRearPortTemplateRequestWithDefaults

`func NewPatchedWritableRearPortTemplateRequestWithDefaults() *PatchedWritableRearPortTemplateRequest`

NewPatchedWritableRearPortTemplateRequestWithDefaults instantiates a new PatchedWritableRearPortTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *PatchedWritableRearPortTemplateRequest) GetDeviceType() BriefDeviceTypeRequest`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritableRearPortTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritableRearPortTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritableRearPortTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *PatchedWritableRearPortTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *PatchedWritableRearPortTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *PatchedWritableRearPortTemplateRequest) GetModuleType() BriefModuleTypeRequest`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *PatchedWritableRearPortTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *PatchedWritableRearPortTemplateRequest) SetModuleType(v BriefModuleTypeRequest)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *PatchedWritableRearPortTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *PatchedWritableRearPortTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *PatchedWritableRearPortTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *PatchedWritableRearPortTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableRearPortTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableRearPortTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableRearPortTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableRearPortTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableRearPortTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableRearPortTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableRearPortTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableRearPortTemplateRequest) GetType() FrontPortTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableRearPortTemplateRequest) GetTypeOk() (*FrontPortTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableRearPortTemplateRequest) SetType(v FrontPortTypeValue)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableRearPortTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetColor

`func (o *PatchedWritableRearPortTemplateRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PatchedWritableRearPortTemplateRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PatchedWritableRearPortTemplateRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *PatchedWritableRearPortTemplateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPositions

`func (o *PatchedWritableRearPortTemplateRequest) GetPositions() int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *PatchedWritableRearPortTemplateRequest) GetPositionsOk() (*int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *PatchedWritableRearPortTemplateRequest) SetPositions(v int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *PatchedWritableRearPortTemplateRequest) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableRearPortTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableRearPortTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableRearPortTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableRearPortTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


