# InterfaceTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**DeviceType** | Pointer to [**NullableBriefDeviceType**](BriefDeviceType.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableBriefModuleType**](BriefModuleType.md) |  | [optional] 
**Name** | **string** | {module} is accepted as a substitution for the module bay position when attached to a module type. | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | [**InterfaceType**](InterfaceType.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**MgmtOnly** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Bridge** | Pointer to [**NullableNestedInterfaceTemplate**](NestedInterfaceTemplate.md) |  | [optional] 
**PoeMode** | Pointer to [**NullableInterfaceTemplatePoeMode**](InterfaceTemplatePoeMode.md) |  | [optional] 
**PoeType** | Pointer to [**NullableInterfaceTemplatePoeType**](InterfaceTemplatePoeType.md) |  | [optional] 
**RfRole** | Pointer to [**NullableInterfaceTemplateRfRole**](InterfaceTemplateRfRole.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewInterfaceTemplate

`func NewInterfaceTemplate(id int32, url string, display string, name string, type_ InterfaceType, created NullableTime, lastUpdated NullableTime, ) *InterfaceTemplate`

NewInterfaceTemplate instantiates a new InterfaceTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceTemplateWithDefaults

`func NewInterfaceTemplateWithDefaults() *InterfaceTemplate`

NewInterfaceTemplateWithDefaults instantiates a new InterfaceTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InterfaceTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterfaceTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterfaceTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *InterfaceTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InterfaceTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InterfaceTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *InterfaceTemplate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *InterfaceTemplate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *InterfaceTemplate) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDeviceType

`func (o *InterfaceTemplate) GetDeviceType() BriefDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *InterfaceTemplate) GetDeviceTypeOk() (*BriefDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *InterfaceTemplate) SetDeviceType(v BriefDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *InterfaceTemplate) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *InterfaceTemplate) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *InterfaceTemplate) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *InterfaceTemplate) GetModuleType() BriefModuleType`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *InterfaceTemplate) GetModuleTypeOk() (*BriefModuleType, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *InterfaceTemplate) SetModuleType(v BriefModuleType)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *InterfaceTemplate) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *InterfaceTemplate) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *InterfaceTemplate) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *InterfaceTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *InterfaceTemplate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InterfaceTemplate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InterfaceTemplate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InterfaceTemplate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *InterfaceTemplate) GetType() InterfaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterfaceTemplate) GetTypeOk() (*InterfaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterfaceTemplate) SetType(v InterfaceType)`

SetType sets Type field to given value.


### GetEnabled

`func (o *InterfaceTemplate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InterfaceTemplate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InterfaceTemplate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InterfaceTemplate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMgmtOnly

`func (o *InterfaceTemplate) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *InterfaceTemplate) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *InterfaceTemplate) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *InterfaceTemplate) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDescription

`func (o *InterfaceTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterfaceTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterfaceTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterfaceTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBridge

`func (o *InterfaceTemplate) GetBridge() NestedInterfaceTemplate`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *InterfaceTemplate) GetBridgeOk() (*NestedInterfaceTemplate, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *InterfaceTemplate) SetBridge(v NestedInterfaceTemplate)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *InterfaceTemplate) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *InterfaceTemplate) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *InterfaceTemplate) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetPoeMode

`func (o *InterfaceTemplate) GetPoeMode() InterfaceTemplatePoeMode`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *InterfaceTemplate) GetPoeModeOk() (*InterfaceTemplatePoeMode, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *InterfaceTemplate) SetPoeMode(v InterfaceTemplatePoeMode)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *InterfaceTemplate) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### SetPoeModeNil

`func (o *InterfaceTemplate) SetPoeModeNil(b bool)`

 SetPoeModeNil sets the value for PoeMode to be an explicit nil

### UnsetPoeMode
`func (o *InterfaceTemplate) UnsetPoeMode()`

UnsetPoeMode ensures that no value is present for PoeMode, not even an explicit nil
### GetPoeType

`func (o *InterfaceTemplate) GetPoeType() InterfaceTemplatePoeType`

GetPoeType returns the PoeType field if non-nil, zero value otherwise.

### GetPoeTypeOk

`func (o *InterfaceTemplate) GetPoeTypeOk() (*InterfaceTemplatePoeType, bool)`

GetPoeTypeOk returns a tuple with the PoeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeType

`func (o *InterfaceTemplate) SetPoeType(v InterfaceTemplatePoeType)`

SetPoeType sets PoeType field to given value.

### HasPoeType

`func (o *InterfaceTemplate) HasPoeType() bool`

HasPoeType returns a boolean if a field has been set.

### SetPoeTypeNil

`func (o *InterfaceTemplate) SetPoeTypeNil(b bool)`

 SetPoeTypeNil sets the value for PoeType to be an explicit nil

### UnsetPoeType
`func (o *InterfaceTemplate) UnsetPoeType()`

UnsetPoeType ensures that no value is present for PoeType, not even an explicit nil
### GetRfRole

`func (o *InterfaceTemplate) GetRfRole() InterfaceTemplateRfRole`

GetRfRole returns the RfRole field if non-nil, zero value otherwise.

### GetRfRoleOk

`func (o *InterfaceTemplate) GetRfRoleOk() (*InterfaceTemplateRfRole, bool)`

GetRfRoleOk returns a tuple with the RfRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfRole

`func (o *InterfaceTemplate) SetRfRole(v InterfaceTemplateRfRole)`

SetRfRole sets RfRole field to given value.

### HasRfRole

`func (o *InterfaceTemplate) HasRfRole() bool`

HasRfRole returns a boolean if a field has been set.

### SetRfRoleNil

`func (o *InterfaceTemplate) SetRfRoleNil(b bool)`

 SetRfRoleNil sets the value for RfRole to be an explicit nil

### UnsetRfRole
`func (o *InterfaceTemplate) UnsetRfRole()`

UnsetRfRole ensures that no value is present for RfRole, not even an explicit nil
### GetCreated

`func (o *InterfaceTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InterfaceTemplate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InterfaceTemplate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *InterfaceTemplate) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *InterfaceTemplate) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *InterfaceTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InterfaceTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InterfaceTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *InterfaceTemplate) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *InterfaceTemplate) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


