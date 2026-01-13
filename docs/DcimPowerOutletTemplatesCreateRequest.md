# DcimPowerOutletTemplatesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**NullableConsolePortTemplateRequestDeviceType**](ConsolePortTemplateRequestDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableConsolePortTemplateRequestModuleType**](ConsolePortTemplateRequestModuleType.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**NullablePatchedWritablePowerOutletTemplateRequestType**](PatchedWritablePowerOutletTemplateRequestType.md) |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**PowerPort** | Pointer to [**NullablePatchedWritablePowerOutletTemplateRequestPowerPort**](PatchedWritablePowerOutletTemplateRequestPowerPort.md) |  | [optional] 
**FeedLeg** | Pointer to [**NullablePatchedWritablePowerOutletRequestFeedLeg**](PatchedWritablePowerOutletRequestFeedLeg.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewDcimPowerOutletTemplatesCreateRequest

`func NewDcimPowerOutletTemplatesCreateRequest(name string, ) *DcimPowerOutletTemplatesCreateRequest`

NewDcimPowerOutletTemplatesCreateRequest instantiates a new DcimPowerOutletTemplatesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimPowerOutletTemplatesCreateRequestWithDefaults

`func NewDcimPowerOutletTemplatesCreateRequestWithDefaults() *DcimPowerOutletTemplatesCreateRequest`

NewDcimPowerOutletTemplatesCreateRequestWithDefaults instantiates a new DcimPowerOutletTemplatesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *DcimPowerOutletTemplatesCreateRequest) GetDeviceType() ConsolePortTemplateRequestDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DcimPowerOutletTemplatesCreateRequest) GetDeviceTypeOk() (*ConsolePortTemplateRequestDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DcimPowerOutletTemplatesCreateRequest) SetDeviceType(v ConsolePortTemplateRequestDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *DcimPowerOutletTemplatesCreateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *DcimPowerOutletTemplatesCreateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *DcimPowerOutletTemplatesCreateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *DcimPowerOutletTemplatesCreateRequest) GetModuleType() ConsolePortTemplateRequestModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *DcimPowerOutletTemplatesCreateRequest) GetModuleTypeOk() (*ConsolePortTemplateRequestModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *DcimPowerOutletTemplatesCreateRequest) SetModuleType(v ConsolePortTemplateRequestModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *DcimPowerOutletTemplatesCreateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *DcimPowerOutletTemplatesCreateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *DcimPowerOutletTemplatesCreateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *DcimPowerOutletTemplatesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcimPowerOutletTemplatesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcimPowerOutletTemplatesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *DcimPowerOutletTemplatesCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DcimPowerOutletTemplatesCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DcimPowerOutletTemplatesCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DcimPowerOutletTemplatesCreateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *DcimPowerOutletTemplatesCreateRequest) GetType() PatchedWritablePowerOutletTemplateRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DcimPowerOutletTemplatesCreateRequest) GetTypeOk() (*PatchedWritablePowerOutletTemplateRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DcimPowerOutletTemplatesCreateRequest) SetType(v PatchedWritablePowerOutletTemplateRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *DcimPowerOutletTemplatesCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DcimPowerOutletTemplatesCreateRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DcimPowerOutletTemplatesCreateRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetColor

`func (o *DcimPowerOutletTemplatesCreateRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DcimPowerOutletTemplatesCreateRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DcimPowerOutletTemplatesCreateRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DcimPowerOutletTemplatesCreateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPowerPort

`func (o *DcimPowerOutletTemplatesCreateRequest) GetPowerPort() PatchedWritablePowerOutletTemplateRequestPowerPort`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *DcimPowerOutletTemplatesCreateRequest) GetPowerPortOk() (*PatchedWritablePowerOutletTemplateRequestPowerPort, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *DcimPowerOutletTemplatesCreateRequest) SetPowerPort(v PatchedWritablePowerOutletTemplateRequestPowerPort)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *DcimPowerOutletTemplatesCreateRequest) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *DcimPowerOutletTemplatesCreateRequest) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *DcimPowerOutletTemplatesCreateRequest) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetFeedLeg

`func (o *DcimPowerOutletTemplatesCreateRequest) GetFeedLeg() PatchedWritablePowerOutletRequestFeedLeg`

GetFeedLeg returns the FeedLeg field if non-nil, zero value otherwise.

### GetFeedLegOk

`func (o *DcimPowerOutletTemplatesCreateRequest) GetFeedLegOk() (*PatchedWritablePowerOutletRequestFeedLeg, bool)`

GetFeedLegOk returns a tuple with the FeedLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedLeg

`func (o *DcimPowerOutletTemplatesCreateRequest) SetFeedLeg(v PatchedWritablePowerOutletRequestFeedLeg)`

SetFeedLeg sets FeedLeg field to given value.

### HasFeedLeg

`func (o *DcimPowerOutletTemplatesCreateRequest) HasFeedLeg() bool`

HasFeedLeg returns a boolean if a field has been set.

### SetFeedLegNil

`func (o *DcimPowerOutletTemplatesCreateRequest) SetFeedLegNil(b bool)`

 SetFeedLegNil sets the value for FeedLeg to be an explicit nil

### UnsetFeedLeg
`func (o *DcimPowerOutletTemplatesCreateRequest) UnsetFeedLeg()`

UnsetFeedLeg ensures that no value is present for FeedLeg, not even an explicit nil
### GetDescription

`func (o *DcimPowerOutletTemplatesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DcimPowerOutletTemplatesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DcimPowerOutletTemplatesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DcimPowerOutletTemplatesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


