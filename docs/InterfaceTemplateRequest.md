# InterfaceTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to [**NullableBriefDeviceTypeRequest**](BriefDeviceTypeRequest.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBriefModuleTypeRequest**](BriefModuleTypeRequest.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | [**InterfaceTypeValue**](InterfaceTypeValue.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**MgmtOnly** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Bridge** | Pointer to [**NullableNestedInterfaceTemplateRequest**](NestedInterfaceTemplateRequest.md) |  | [optional] 
**PoeMode** | Pointer to [**NullableInterfaceTemplateRequestPoeMode**](InterfaceTemplateRequestPoeMode.md) |  | [optional] 
**PoeType** | Pointer to [**NullableInterfaceTemplateRequestPoeType**](InterfaceTemplateRequestPoeType.md) |  | [optional] 
**RfRole** | Pointer to [**NullableInterfaceTemplateRequestRfRole**](InterfaceTemplateRequestRfRole.md) |  | [optional] 

## Methods

### NewInterfaceTemplateRequest

`func NewInterfaceTemplateRequest(name string, type_ InterfaceTypeValue, ) *InterfaceTemplateRequest`

NewInterfaceTemplateRequest instantiates a new InterfaceTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceTemplateRequestWithDefaults

`func NewInterfaceTemplateRequestWithDefaults() *InterfaceTemplateRequest`

NewInterfaceTemplateRequestWithDefaults instantiates a new InterfaceTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *InterfaceTemplateRequest) GetDeviceType() BriefDeviceTypeRequest`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *InterfaceTemplateRequest) GetDeviceTypeOk() (*BriefDeviceTypeRequest, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *InterfaceTemplateRequest) SetDeviceType(v BriefDeviceTypeRequest)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *InterfaceTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *InterfaceTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *InterfaceTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *InterfaceTemplateRequest) GetModuleType() BriefModuleTypeRequest`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *InterfaceTemplateRequest) GetModuleTypeOk() (*BriefModuleTypeRequest, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *InterfaceTemplateRequest) SetModuleType(v BriefModuleTypeRequest)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *InterfaceTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *InterfaceTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *InterfaceTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *InterfaceTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *InterfaceTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InterfaceTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InterfaceTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InterfaceTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *InterfaceTemplateRequest) GetType() InterfaceTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterfaceTemplateRequest) GetTypeOk() (*InterfaceTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterfaceTemplateRequest) SetType(v InterfaceTypeValue)`

SetType sets Type field to given value.


### GetEnabled

`func (o *InterfaceTemplateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InterfaceTemplateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InterfaceTemplateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InterfaceTemplateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMgmtOnly

`func (o *InterfaceTemplateRequest) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *InterfaceTemplateRequest) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *InterfaceTemplateRequest) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *InterfaceTemplateRequest) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDescription

`func (o *InterfaceTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterfaceTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterfaceTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterfaceTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBridge

`func (o *InterfaceTemplateRequest) GetBridge() NestedInterfaceTemplateRequest`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *InterfaceTemplateRequest) GetBridgeOk() (*NestedInterfaceTemplateRequest, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *InterfaceTemplateRequest) SetBridge(v NestedInterfaceTemplateRequest)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *InterfaceTemplateRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *InterfaceTemplateRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *InterfaceTemplateRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetPoeMode

`func (o *InterfaceTemplateRequest) GetPoeMode() InterfaceTemplateRequestPoeMode`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *InterfaceTemplateRequest) GetPoeModeOk() (*InterfaceTemplateRequestPoeMode, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *InterfaceTemplateRequest) SetPoeMode(v InterfaceTemplateRequestPoeMode)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *InterfaceTemplateRequest) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### SetPoeModeNil

`func (o *InterfaceTemplateRequest) SetPoeModeNil(b bool)`

 SetPoeModeNil sets the value for PoeMode to be an explicit nil

### UnsetPoeMode
`func (o *InterfaceTemplateRequest) UnsetPoeMode()`

UnsetPoeMode ensures that no value is present for PoeMode, not even an explicit nil
### GetPoeType

`func (o *InterfaceTemplateRequest) GetPoeType() InterfaceTemplateRequestPoeType`

GetPoeType returns the PoeType field if non-nil, zero value otherwise.

### GetPoeTypeOk

`func (o *InterfaceTemplateRequest) GetPoeTypeOk() (*InterfaceTemplateRequestPoeType, bool)`

GetPoeTypeOk returns a tuple with the PoeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeType

`func (o *InterfaceTemplateRequest) SetPoeType(v InterfaceTemplateRequestPoeType)`

SetPoeType sets PoeType field to given value.

### HasPoeType

`func (o *InterfaceTemplateRequest) HasPoeType() bool`

HasPoeType returns a boolean if a field has been set.

### SetPoeTypeNil

`func (o *InterfaceTemplateRequest) SetPoeTypeNil(b bool)`

 SetPoeTypeNil sets the value for PoeType to be an explicit nil

### UnsetPoeType
`func (o *InterfaceTemplateRequest) UnsetPoeType()`

UnsetPoeType ensures that no value is present for PoeType, not even an explicit nil
### GetRfRole

`func (o *InterfaceTemplateRequest) GetRfRole() InterfaceTemplateRequestRfRole`

GetRfRole returns the RfRole field if non-nil, zero value otherwise.

### GetRfRoleOk

`func (o *InterfaceTemplateRequest) GetRfRoleOk() (*InterfaceTemplateRequestRfRole, bool)`

GetRfRoleOk returns a tuple with the RfRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfRole

`func (o *InterfaceTemplateRequest) SetRfRole(v InterfaceTemplateRequestRfRole)`

SetRfRole sets RfRole field to given value.

### HasRfRole

`func (o *InterfaceTemplateRequest) HasRfRole() bool`

HasRfRole returns a boolean if a field has been set.

### SetRfRoleNil

`func (o *InterfaceTemplateRequest) SetRfRoleNil(b bool)`

 SetRfRoleNil sets the value for RfRole to be an explicit nil

### UnsetRfRole
`func (o *InterfaceTemplateRequest) UnsetRfRole()`

UnsetRfRole ensures that no value is present for RfRole, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


