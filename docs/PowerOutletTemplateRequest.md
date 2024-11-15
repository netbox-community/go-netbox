# PowerOutletTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**NullableBriefDeviceTypeRequest**](BriefDeviceTypeRequest.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBriefModuleTypeRequest**](BriefModuleTypeRequest.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**NullablePowerOutletRequestType**](PowerOutletRequestType.md) |  | [optional] 
**PowerPort** | Pointer to [**NullableBriefPowerPortTemplateRequest**](BriefPowerPortTemplateRequest.md) |  | [optional] 
**FeedLeg** | Pointer to [**NullablePowerOutletRequestFeedLeg**](PowerOutletRequestFeedLeg.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPowerOutletTemplateRequest

`func NewPowerOutletTemplateRequest(name string, ) *PowerOutletTemplateRequest`

NewPowerOutletTemplateRequest instantiates a new PowerOutletTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerOutletTemplateRequestWithDefaults

`func NewPowerOutletTemplateRequestWithDefaults() *PowerOutletTemplateRequest`

NewPowerOutletTemplateRequestWithDefaults instantiates a new PowerOutletTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *PowerOutletTemplateRequest) GetDeviceType() BriefDeviceTypeRequest`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PowerOutletTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PowerOutletTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PowerOutletTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *PowerOutletTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *PowerOutletTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *PowerOutletTemplateRequest) GetModuleType() BriefModuleTypeRequest`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *PowerOutletTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *PowerOutletTemplateRequest) SetModuleType(v BriefModuleTypeRequest)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *PowerOutletTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *PowerOutletTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *PowerOutletTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *PowerOutletTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerOutletTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerOutletTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *PowerOutletTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PowerOutletTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PowerOutletTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PowerOutletTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PowerOutletTemplateRequest) GetType() PowerOutletRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PowerOutletTemplateRequest) GetTypeOk() (*PowerOutletRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PowerOutletTemplateRequest) SetType(v PowerOutletRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PowerOutletTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PowerOutletTemplateRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PowerOutletTemplateRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPowerPort

`func (o *PowerOutletTemplateRequest) GetPowerPort() BriefPowerPortTemplateRequest`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *PowerOutletTemplateRequest) GetPowerPortOk() (*BriefPowerPortTemplateRequest, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *PowerOutletTemplateRequest) SetPowerPort(v BriefPowerPortTemplateRequest)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *PowerOutletTemplateRequest) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *PowerOutletTemplateRequest) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *PowerOutletTemplateRequest) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetFeedLeg

`func (o *PowerOutletTemplateRequest) GetFeedLeg() PowerOutletRequestFeedLeg`

GetFeedLeg returns the FeedLeg field if non-nil, zero value otherwise.

### GetFeedLegOk

`func (o *PowerOutletTemplateRequest) GetFeedLegOk() (*PowerOutletRequestFeedLeg, bool)`

GetFeedLegOk returns a tuple with the FeedLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedLeg

`func (o *PowerOutletTemplateRequest) SetFeedLeg(v PowerOutletRequestFeedLeg)`

SetFeedLeg sets FeedLeg field to given value.

### HasFeedLeg

`func (o *PowerOutletTemplateRequest) HasFeedLeg() bool`

HasFeedLeg returns a boolean if a field has been set.

### SetFeedLegNil

`func (o *PowerOutletTemplateRequest) SetFeedLegNil(b bool)`

 SetFeedLegNil sets the value for FeedLeg to be an explicit nil

### UnsetFeedLeg
`func (o *PowerOutletTemplateRequest) UnsetFeedLeg()`

UnsetFeedLeg ensures that no value is present for FeedLeg, not even an explicit nil
### GetDescription

`func (o *PowerOutletTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PowerOutletTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PowerOutletTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PowerOutletTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


