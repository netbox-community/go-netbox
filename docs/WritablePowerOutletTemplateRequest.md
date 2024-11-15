# WritablePowerOutletTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**NullableBriefDeviceTypeRequest**](BriefDeviceTypeRequest.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBriefModuleTypeRequest**](BriefModuleTypeRequest.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**PatchedWritablePowerOutletTemplateRequestType**](PatchedWritablePowerOutletTemplateRequestType.md) |  | [optional] 
**PowerPort** | Pointer to [**NullableBriefPowerPortTemplateRequest**](BriefPowerPortTemplateRequest.md) |  | [optional] 
**FeedLeg** | Pointer to [**PatchedWritablePowerOutletRequestFeedLeg**](PatchedWritablePowerOutletRequestFeedLeg.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewWritablePowerOutletTemplateRequest

`func NewWritablePowerOutletTemplateRequest(name string, ) *WritablePowerOutletTemplateRequest`

NewWritablePowerOutletTemplateRequest instantiates a new WritablePowerOutletTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePowerOutletTemplateRequestWithDefaults

`func NewWritablePowerOutletTemplateRequestWithDefaults() *WritablePowerOutletTemplateRequest`

NewWritablePowerOutletTemplateRequestWithDefaults instantiates a new WritablePowerOutletTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *WritablePowerOutletTemplateRequest) GetDeviceType() BriefDeviceTypeRequest`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WritablePowerOutletTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WritablePowerOutletTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *WritablePowerOutletTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *WritablePowerOutletTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *WritablePowerOutletTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *WritablePowerOutletTemplateRequest) GetModuleType() BriefModuleTypeRequest`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *WritablePowerOutletTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *WritablePowerOutletTemplateRequest) SetModuleType(v BriefModuleTypeRequest)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *WritablePowerOutletTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *WritablePowerOutletTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *WritablePowerOutletTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *WritablePowerOutletTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritablePowerOutletTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritablePowerOutletTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritablePowerOutletTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritablePowerOutletTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritablePowerOutletTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritablePowerOutletTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *WritablePowerOutletTemplateRequest) GetType() PatchedWritablePowerOutletTemplateRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritablePowerOutletTemplateRequest) GetTypeOk() (*PatchedWritablePowerOutletTemplateRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritablePowerOutletTemplateRequest) SetType(v PatchedWritablePowerOutletTemplateRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *WritablePowerOutletTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPowerPort

`func (o *WritablePowerOutletTemplateRequest) GetPowerPort() BriefPowerPortTemplateRequest`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *WritablePowerOutletTemplateRequest) GetPowerPortOk() (*BriefPowerPortTemplateRequest, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *WritablePowerOutletTemplateRequest) SetPowerPort(v BriefPowerPortTemplateRequest)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *WritablePowerOutletTemplateRequest) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *WritablePowerOutletTemplateRequest) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *WritablePowerOutletTemplateRequest) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetFeedLeg

`func (o *WritablePowerOutletTemplateRequest) GetFeedLeg() PatchedWritablePowerOutletRequestFeedLeg`

GetFeedLeg returns the FeedLeg field if non-nil, zero value otherwise.

### GetFeedLegOk

`func (o *WritablePowerOutletTemplateRequest) GetFeedLegOk() (*PatchedWritablePowerOutletRequestFeedLeg, bool)`

GetFeedLegOk returns a tuple with the FeedLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedLeg

`func (o *WritablePowerOutletTemplateRequest) SetFeedLeg(v PatchedWritablePowerOutletRequestFeedLeg)`

SetFeedLeg sets FeedLeg field to given value.

### HasFeedLeg

`func (o *WritablePowerOutletTemplateRequest) HasFeedLeg() bool`

HasFeedLeg returns a boolean if a field has been set.

### GetDescription

`func (o *WritablePowerOutletTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritablePowerOutletTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritablePowerOutletTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritablePowerOutletTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


