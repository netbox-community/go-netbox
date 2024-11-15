# WritablePowerPortTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**NullableBriefDeviceTypeRequest**](BriefDeviceTypeRequest.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBriefModuleTypeRequest**](BriefModuleTypeRequest.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**PatchedWritablePowerPortTemplateRequestType**](PatchedWritablePowerPortTemplateRequestType.md) |  | [optional] 
**MaximumDraw** | Pointer to **NullableInt32** | Maximum power draw (watts) | [optional] 
**AllocatedDraw** | Pointer to **NullableInt32** | Allocated power draw (watts) | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewWritablePowerPortTemplateRequest

`func NewWritablePowerPortTemplateRequest(name string, ) *WritablePowerPortTemplateRequest`

NewWritablePowerPortTemplateRequest instantiates a new WritablePowerPortTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePowerPortTemplateRequestWithDefaults

`func NewWritablePowerPortTemplateRequestWithDefaults() *WritablePowerPortTemplateRequest`

NewWritablePowerPortTemplateRequestWithDefaults instantiates a new WritablePowerPortTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *WritablePowerPortTemplateRequest) GetDeviceType() BriefDeviceTypeRequest`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WritablePowerPortTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WritablePowerPortTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *WritablePowerPortTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *WritablePowerPortTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *WritablePowerPortTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *WritablePowerPortTemplateRequest) GetModuleType() BriefModuleTypeRequest`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *WritablePowerPortTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *WritablePowerPortTemplateRequest) SetModuleType(v BriefModuleTypeRequest)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *WritablePowerPortTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *WritablePowerPortTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *WritablePowerPortTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *WritablePowerPortTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritablePowerPortTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritablePowerPortTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritablePowerPortTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritablePowerPortTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritablePowerPortTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritablePowerPortTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *WritablePowerPortTemplateRequest) GetType() PatchedWritablePowerPortTemplateRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritablePowerPortTemplateRequest) GetTypeOk() (*PatchedWritablePowerPortTemplateRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritablePowerPortTemplateRequest) SetType(v PatchedWritablePowerPortTemplateRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *WritablePowerPortTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMaximumDraw

`func (o *WritablePowerPortTemplateRequest) GetMaximumDraw() int32`

GetMaximumDraw returns the MaximumDraw field if non-nil, zero value otherwise.

### GetMaximumDrawOk

`func (o *WritablePowerPortTemplateRequest) GetMaximumDrawOk() (*int32, bool)`

GetMaximumDrawOk returns a tuple with the MaximumDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDraw

`func (o *WritablePowerPortTemplateRequest) SetMaximumDraw(v int32)`

SetMaximumDraw sets MaximumDraw field to given value.

### HasMaximumDraw

`func (o *WritablePowerPortTemplateRequest) HasMaximumDraw() bool`

HasMaximumDraw returns a boolean if a field has been set.

### SetMaximumDrawNil

`func (o *WritablePowerPortTemplateRequest) SetMaximumDrawNil(b bool)`

 SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil

### UnsetMaximumDraw
`func (o *WritablePowerPortTemplateRequest) UnsetMaximumDraw()`

UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
### GetAllocatedDraw

`func (o *WritablePowerPortTemplateRequest) GetAllocatedDraw() int32`

GetAllocatedDraw returns the AllocatedDraw field if non-nil, zero value otherwise.

### GetAllocatedDrawOk

`func (o *WritablePowerPortTemplateRequest) GetAllocatedDrawOk() (*int32, bool)`

GetAllocatedDrawOk returns a tuple with the AllocatedDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedDraw

`func (o *WritablePowerPortTemplateRequest) SetAllocatedDraw(v int32)`

SetAllocatedDraw sets AllocatedDraw field to given value.

### HasAllocatedDraw

`func (o *WritablePowerPortTemplateRequest) HasAllocatedDraw() bool`

HasAllocatedDraw returns a boolean if a field has been set.

### SetAllocatedDrawNil

`func (o *WritablePowerPortTemplateRequest) SetAllocatedDrawNil(b bool)`

 SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil

### UnsetAllocatedDraw
`func (o *WritablePowerPortTemplateRequest) UnsetAllocatedDraw()`

UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
### GetDescription

`func (o *WritablePowerPortTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritablePowerPortTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritablePowerPortTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritablePowerPortTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


