# DcimInterfacesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**BriefInterfaceRequestDevice**](BriefInterfaceRequestDevice.md) |  | 
**Vdcs** | Pointer to **[]int32** |  | [optional] 
**Module** | Pointer to [**NullableConsolePortRequestModule**](ConsolePortRequestModule.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | [**InterfaceTypeValue**](InterfaceTypeValue.md) |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Bridge** | Pointer to **NullableInt32** |  | [optional] 
**Lag** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**PrimaryMacAddress** | Pointer to [**NullableInterfaceRequestPrimaryMacAddress**](InterfaceRequestPrimaryMacAddress.md) |  | [optional] 
**Speed** | Pointer to **NullableInt32** |  | [optional] 
**Duplex** | Pointer to [**NullableInterfaceRequestDuplex**](InterfaceRequestDuplex.md) |  | [optional] 
**Wwn** | Pointer to **NullableString** |  | [optional] 
**MgmtOnly** | Pointer to **bool** | This interface is used only for out-of-band management | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**NullablePatchedWritableInterfaceRequestMode**](PatchedWritableInterfaceRequestMode.md) |  | [optional] 
**RfRole** | Pointer to [**NullableWirelessRole**](WirelessRole.md) |  | [optional] 
**RfChannel** | Pointer to [**NullableWirelessChannel**](WirelessChannel.md) |  | [optional] 
**PoeMode** | Pointer to [**NullableInterfaceTemplateRequestPoeMode**](InterfaceTemplateRequestPoeMode.md) |  | [optional] 
**PoeType** | Pointer to [**NullableInterfaceTemplateRequestPoeType**](InterfaceTemplateRequestPoeType.md) |  | [optional] 
**RfChannelFrequency** | Pointer to **NullableFloat64** | Populated by selected channel (if set) | [optional] 
**RfChannelWidth** | Pointer to **NullableFloat64** | Populated by selected channel (if set) | [optional] 
**TxPower** | Pointer to **NullableInt32** |  | [optional] 
**UntaggedVlan** | Pointer to [**NullableInterfaceRequestUntaggedVlan**](InterfaceRequestUntaggedVlan.md) |  | [optional] 
**TaggedVlans** | Pointer to **[]int32** |  | [optional] 
**QinqSvlan** | Pointer to [**NullableInterfaceRequestUntaggedVlan**](InterfaceRequestUntaggedVlan.md) |  | [optional] 
**VlanTranslationPolicy** | Pointer to [**NullableInterfaceRequestVlanTranslationPolicy**](InterfaceRequestVlanTranslationPolicy.md) |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**WirelessLans** | Pointer to **[]int32** |  | [optional] 
**Vrf** | Pointer to [**NullableIPAddressRequestVrf**](IPAddressRequestVrf.md) |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewDcimInterfacesCreateRequest

`func NewDcimInterfacesCreateRequest(device BriefInterfaceRequestDevice, name string, type_ InterfaceTypeValue, ) *DcimInterfacesCreateRequest`

NewDcimInterfacesCreateRequest instantiates a new DcimInterfacesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDcimInterfacesCreateRequestWithDefaults

`func NewDcimInterfacesCreateRequestWithDefaults() *DcimInterfacesCreateRequest`

NewDcimInterfacesCreateRequestWithDefaults instantiates a new DcimInterfacesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *DcimInterfacesCreateRequest) GetDevice() BriefInterfaceRequestDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DcimInterfacesCreateRequest) GetDeviceOk() (*BriefInterfaceRequestDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DcimInterfacesCreateRequest) SetDevice(v BriefInterfaceRequestDevice)`

SetDevice sets Device field to given value.


### GetVdcs

`func (o *DcimInterfacesCreateRequest) GetVdcs() []int32`

GetVdcs returns the Vdcs field if non-nil, zero value otherwise.

### GetVdcsOk

`func (o *DcimInterfacesCreateRequest) GetVdcsOk() (*[]int32, bool)`

GetVdcsOk returns a tuple with the Vdcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcs

`func (o *DcimInterfacesCreateRequest) SetVdcs(v []int32)`

SetVdcs sets Vdcs field to given value.

### HasVdcs

`func (o *DcimInterfacesCreateRequest) HasVdcs() bool`

HasVdcs returns a boolean if a field has been set.

### GetModule

`func (o *DcimInterfacesCreateRequest) GetModule() ConsolePortRequestModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *DcimInterfacesCreateRequest) GetModuleOk() (*ConsolePortRequestModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *DcimInterfacesCreateRequest) SetModule(v ConsolePortRequestModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *DcimInterfacesCreateRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *DcimInterfacesCreateRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *DcimInterfacesCreateRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *DcimInterfacesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DcimInterfacesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DcimInterfacesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *DcimInterfacesCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DcimInterfacesCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DcimInterfacesCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DcimInterfacesCreateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *DcimInterfacesCreateRequest) GetType() InterfaceTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DcimInterfacesCreateRequest) GetTypeOk() (*InterfaceTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DcimInterfacesCreateRequest) SetType(v InterfaceTypeValue)`

SetType sets Type field to given value.


### GetEnabled

`func (o *DcimInterfacesCreateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DcimInterfacesCreateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DcimInterfacesCreateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DcimInterfacesCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParent

`func (o *DcimInterfacesCreateRequest) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *DcimInterfacesCreateRequest) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *DcimInterfacesCreateRequest) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *DcimInterfacesCreateRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *DcimInterfacesCreateRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *DcimInterfacesCreateRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBridge

`func (o *DcimInterfacesCreateRequest) GetBridge() int32`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *DcimInterfacesCreateRequest) GetBridgeOk() (*int32, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *DcimInterfacesCreateRequest) SetBridge(v int32)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *DcimInterfacesCreateRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *DcimInterfacesCreateRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *DcimInterfacesCreateRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetLag

`func (o *DcimInterfacesCreateRequest) GetLag() int32`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *DcimInterfacesCreateRequest) GetLagOk() (*int32, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *DcimInterfacesCreateRequest) SetLag(v int32)`

SetLag sets Lag field to given value.

### HasLag

`func (o *DcimInterfacesCreateRequest) HasLag() bool`

HasLag returns a boolean if a field has been set.

### SetLagNil

`func (o *DcimInterfacesCreateRequest) SetLagNil(b bool)`

 SetLagNil sets the value for Lag to be an explicit nil

### UnsetLag
`func (o *DcimInterfacesCreateRequest) UnsetLag()`

UnsetLag ensures that no value is present for Lag, not even an explicit nil
### GetMtu

`func (o *DcimInterfacesCreateRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *DcimInterfacesCreateRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *DcimInterfacesCreateRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *DcimInterfacesCreateRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *DcimInterfacesCreateRequest) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *DcimInterfacesCreateRequest) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetPrimaryMacAddress

`func (o *DcimInterfacesCreateRequest) GetPrimaryMacAddress() InterfaceRequestPrimaryMacAddress`

GetPrimaryMacAddress returns the PrimaryMacAddress field if non-nil, zero value otherwise.

### GetPrimaryMacAddressOk

`func (o *DcimInterfacesCreateRequest) GetPrimaryMacAddressOk() (*InterfaceRequestPrimaryMacAddress, bool)`

GetPrimaryMacAddressOk returns a tuple with the PrimaryMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMacAddress

`func (o *DcimInterfacesCreateRequest) SetPrimaryMacAddress(v InterfaceRequestPrimaryMacAddress)`

SetPrimaryMacAddress sets PrimaryMacAddress field to given value.

### HasPrimaryMacAddress

`func (o *DcimInterfacesCreateRequest) HasPrimaryMacAddress() bool`

HasPrimaryMacAddress returns a boolean if a field has been set.

### SetPrimaryMacAddressNil

`func (o *DcimInterfacesCreateRequest) SetPrimaryMacAddressNil(b bool)`

 SetPrimaryMacAddressNil sets the value for PrimaryMacAddress to be an explicit nil

### UnsetPrimaryMacAddress
`func (o *DcimInterfacesCreateRequest) UnsetPrimaryMacAddress()`

UnsetPrimaryMacAddress ensures that no value is present for PrimaryMacAddress, not even an explicit nil
### GetSpeed

`func (o *DcimInterfacesCreateRequest) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *DcimInterfacesCreateRequest) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *DcimInterfacesCreateRequest) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *DcimInterfacesCreateRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *DcimInterfacesCreateRequest) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *DcimInterfacesCreateRequest) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetDuplex

`func (o *DcimInterfacesCreateRequest) GetDuplex() InterfaceRequestDuplex`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *DcimInterfacesCreateRequest) GetDuplexOk() (*InterfaceRequestDuplex, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *DcimInterfacesCreateRequest) SetDuplex(v InterfaceRequestDuplex)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *DcimInterfacesCreateRequest) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### SetDuplexNil

`func (o *DcimInterfacesCreateRequest) SetDuplexNil(b bool)`

 SetDuplexNil sets the value for Duplex to be an explicit nil

### UnsetDuplex
`func (o *DcimInterfacesCreateRequest) UnsetDuplex()`

UnsetDuplex ensures that no value is present for Duplex, not even an explicit nil
### GetWwn

`func (o *DcimInterfacesCreateRequest) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *DcimInterfacesCreateRequest) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *DcimInterfacesCreateRequest) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *DcimInterfacesCreateRequest) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### SetWwnNil

`func (o *DcimInterfacesCreateRequest) SetWwnNil(b bool)`

 SetWwnNil sets the value for Wwn to be an explicit nil

### UnsetWwn
`func (o *DcimInterfacesCreateRequest) UnsetWwn()`

UnsetWwn ensures that no value is present for Wwn, not even an explicit nil
### GetMgmtOnly

`func (o *DcimInterfacesCreateRequest) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *DcimInterfacesCreateRequest) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *DcimInterfacesCreateRequest) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *DcimInterfacesCreateRequest) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDescription

`func (o *DcimInterfacesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DcimInterfacesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DcimInterfacesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DcimInterfacesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *DcimInterfacesCreateRequest) GetMode() PatchedWritableInterfaceRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DcimInterfacesCreateRequest) GetModeOk() (*PatchedWritableInterfaceRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DcimInterfacesCreateRequest) SetMode(v PatchedWritableInterfaceRequestMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *DcimInterfacesCreateRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *DcimInterfacesCreateRequest) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *DcimInterfacesCreateRequest) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetRfRole

`func (o *DcimInterfacesCreateRequest) GetRfRole() WirelessRole`

GetRfRole returns the RfRole field if non-nil, zero value otherwise.

### GetRfRoleOk

`func (o *DcimInterfacesCreateRequest) GetRfRoleOk() (*WirelessRole, bool)`

GetRfRoleOk returns a tuple with the RfRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfRole

`func (o *DcimInterfacesCreateRequest) SetRfRole(v WirelessRole)`

SetRfRole sets RfRole field to given value.

### HasRfRole

`func (o *DcimInterfacesCreateRequest) HasRfRole() bool`

HasRfRole returns a boolean if a field has been set.

### SetRfRoleNil

`func (o *DcimInterfacesCreateRequest) SetRfRoleNil(b bool)`

 SetRfRoleNil sets the value for RfRole to be an explicit nil

### UnsetRfRole
`func (o *DcimInterfacesCreateRequest) UnsetRfRole()`

UnsetRfRole ensures that no value is present for RfRole, not even an explicit nil
### GetRfChannel

`func (o *DcimInterfacesCreateRequest) GetRfChannel() WirelessChannel`

GetRfChannel returns the RfChannel field if non-nil, zero value otherwise.

### GetRfChannelOk

`func (o *DcimInterfacesCreateRequest) GetRfChannelOk() (*WirelessChannel, bool)`

GetRfChannelOk returns a tuple with the RfChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannel

`func (o *DcimInterfacesCreateRequest) SetRfChannel(v WirelessChannel)`

SetRfChannel sets RfChannel field to given value.

### HasRfChannel

`func (o *DcimInterfacesCreateRequest) HasRfChannel() bool`

HasRfChannel returns a boolean if a field has been set.

### SetRfChannelNil

`func (o *DcimInterfacesCreateRequest) SetRfChannelNil(b bool)`

 SetRfChannelNil sets the value for RfChannel to be an explicit nil

### UnsetRfChannel
`func (o *DcimInterfacesCreateRequest) UnsetRfChannel()`

UnsetRfChannel ensures that no value is present for RfChannel, not even an explicit nil
### GetPoeMode

`func (o *DcimInterfacesCreateRequest) GetPoeMode() InterfaceTemplateRequestPoeMode`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *DcimInterfacesCreateRequest) GetPoeModeOk() (*InterfaceTemplateRequestPoeMode, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *DcimInterfacesCreateRequest) SetPoeMode(v InterfaceTemplateRequestPoeMode)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *DcimInterfacesCreateRequest) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### SetPoeModeNil

`func (o *DcimInterfacesCreateRequest) SetPoeModeNil(b bool)`

 SetPoeModeNil sets the value for PoeMode to be an explicit nil

### UnsetPoeMode
`func (o *DcimInterfacesCreateRequest) UnsetPoeMode()`

UnsetPoeMode ensures that no value is present for PoeMode, not even an explicit nil
### GetPoeType

`func (o *DcimInterfacesCreateRequest) GetPoeType() InterfaceTemplateRequestPoeType`

GetPoeType returns the PoeType field if non-nil, zero value otherwise.

### GetPoeTypeOk

`func (o *DcimInterfacesCreateRequest) GetPoeTypeOk() (*InterfaceTemplateRequestPoeType, bool)`

GetPoeTypeOk returns a tuple with the PoeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeType

`func (o *DcimInterfacesCreateRequest) SetPoeType(v InterfaceTemplateRequestPoeType)`

SetPoeType sets PoeType field to given value.

### HasPoeType

`func (o *DcimInterfacesCreateRequest) HasPoeType() bool`

HasPoeType returns a boolean if a field has been set.

### SetPoeTypeNil

`func (o *DcimInterfacesCreateRequest) SetPoeTypeNil(b bool)`

 SetPoeTypeNil sets the value for PoeType to be an explicit nil

### UnsetPoeType
`func (o *DcimInterfacesCreateRequest) UnsetPoeType()`

UnsetPoeType ensures that no value is present for PoeType, not even an explicit nil
### GetRfChannelFrequency

`func (o *DcimInterfacesCreateRequest) GetRfChannelFrequency() float64`

GetRfChannelFrequency returns the RfChannelFrequency field if non-nil, zero value otherwise.

### GetRfChannelFrequencyOk

`func (o *DcimInterfacesCreateRequest) GetRfChannelFrequencyOk() (*float64, bool)`

GetRfChannelFrequencyOk returns a tuple with the RfChannelFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannelFrequency

`func (o *DcimInterfacesCreateRequest) SetRfChannelFrequency(v float64)`

SetRfChannelFrequency sets RfChannelFrequency field to given value.

### HasRfChannelFrequency

`func (o *DcimInterfacesCreateRequest) HasRfChannelFrequency() bool`

HasRfChannelFrequency returns a boolean if a field has been set.

### SetRfChannelFrequencyNil

`func (o *DcimInterfacesCreateRequest) SetRfChannelFrequencyNil(b bool)`

 SetRfChannelFrequencyNil sets the value for RfChannelFrequency to be an explicit nil

### UnsetRfChannelFrequency
`func (o *DcimInterfacesCreateRequest) UnsetRfChannelFrequency()`

UnsetRfChannelFrequency ensures that no value is present for RfChannelFrequency, not even an explicit nil
### GetRfChannelWidth

`func (o *DcimInterfacesCreateRequest) GetRfChannelWidth() float64`

GetRfChannelWidth returns the RfChannelWidth field if non-nil, zero value otherwise.

### GetRfChannelWidthOk

`func (o *DcimInterfacesCreateRequest) GetRfChannelWidthOk() (*float64, bool)`

GetRfChannelWidthOk returns a tuple with the RfChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannelWidth

`func (o *DcimInterfacesCreateRequest) SetRfChannelWidth(v float64)`

SetRfChannelWidth sets RfChannelWidth field to given value.

### HasRfChannelWidth

`func (o *DcimInterfacesCreateRequest) HasRfChannelWidth() bool`

HasRfChannelWidth returns a boolean if a field has been set.

### SetRfChannelWidthNil

`func (o *DcimInterfacesCreateRequest) SetRfChannelWidthNil(b bool)`

 SetRfChannelWidthNil sets the value for RfChannelWidth to be an explicit nil

### UnsetRfChannelWidth
`func (o *DcimInterfacesCreateRequest) UnsetRfChannelWidth()`

UnsetRfChannelWidth ensures that no value is present for RfChannelWidth, not even an explicit nil
### GetTxPower

`func (o *DcimInterfacesCreateRequest) GetTxPower() int32`

GetTxPower returns the TxPower field if non-nil, zero value otherwise.

### GetTxPowerOk

`func (o *DcimInterfacesCreateRequest) GetTxPowerOk() (*int32, bool)`

GetTxPowerOk returns a tuple with the TxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPower

`func (o *DcimInterfacesCreateRequest) SetTxPower(v int32)`

SetTxPower sets TxPower field to given value.

### HasTxPower

`func (o *DcimInterfacesCreateRequest) HasTxPower() bool`

HasTxPower returns a boolean if a field has been set.

### SetTxPowerNil

`func (o *DcimInterfacesCreateRequest) SetTxPowerNil(b bool)`

 SetTxPowerNil sets the value for TxPower to be an explicit nil

### UnsetTxPower
`func (o *DcimInterfacesCreateRequest) UnsetTxPower()`

UnsetTxPower ensures that no value is present for TxPower, not even an explicit nil
### GetUntaggedVlan

`func (o *DcimInterfacesCreateRequest) GetUntaggedVlan() InterfaceRequestUntaggedVlan`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *DcimInterfacesCreateRequest) GetUntaggedVlanOk() (*InterfaceRequestUntaggedVlan, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *DcimInterfacesCreateRequest) SetUntaggedVlan(v InterfaceRequestUntaggedVlan)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *DcimInterfacesCreateRequest) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *DcimInterfacesCreateRequest) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *DcimInterfacesCreateRequest) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetTaggedVlans

`func (o *DcimInterfacesCreateRequest) GetTaggedVlans() []int32`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *DcimInterfacesCreateRequest) GetTaggedVlansOk() (*[]int32, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *DcimInterfacesCreateRequest) SetTaggedVlans(v []int32)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *DcimInterfacesCreateRequest) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetQinqSvlan

`func (o *DcimInterfacesCreateRequest) GetQinqSvlan() InterfaceRequestUntaggedVlan`

GetQinqSvlan returns the QinqSvlan field if non-nil, zero value otherwise.

### GetQinqSvlanOk

`func (o *DcimInterfacesCreateRequest) GetQinqSvlanOk() (*InterfaceRequestUntaggedVlan, bool)`

GetQinqSvlanOk returns a tuple with the QinqSvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqSvlan

`func (o *DcimInterfacesCreateRequest) SetQinqSvlan(v InterfaceRequestUntaggedVlan)`

SetQinqSvlan sets QinqSvlan field to given value.

### HasQinqSvlan

`func (o *DcimInterfacesCreateRequest) HasQinqSvlan() bool`

HasQinqSvlan returns a boolean if a field has been set.

### SetQinqSvlanNil

`func (o *DcimInterfacesCreateRequest) SetQinqSvlanNil(b bool)`

 SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil

### UnsetQinqSvlan
`func (o *DcimInterfacesCreateRequest) UnsetQinqSvlan()`

UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
### GetVlanTranslationPolicy

`func (o *DcimInterfacesCreateRequest) GetVlanTranslationPolicy() InterfaceRequestVlanTranslationPolicy`

GetVlanTranslationPolicy returns the VlanTranslationPolicy field if non-nil, zero value otherwise.

### GetVlanTranslationPolicyOk

`func (o *DcimInterfacesCreateRequest) GetVlanTranslationPolicyOk() (*InterfaceRequestVlanTranslationPolicy, bool)`

GetVlanTranslationPolicyOk returns a tuple with the VlanTranslationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTranslationPolicy

`func (o *DcimInterfacesCreateRequest) SetVlanTranslationPolicy(v InterfaceRequestVlanTranslationPolicy)`

SetVlanTranslationPolicy sets VlanTranslationPolicy field to given value.

### HasVlanTranslationPolicy

`func (o *DcimInterfacesCreateRequest) HasVlanTranslationPolicy() bool`

HasVlanTranslationPolicy returns a boolean if a field has been set.

### SetVlanTranslationPolicyNil

`func (o *DcimInterfacesCreateRequest) SetVlanTranslationPolicyNil(b bool)`

 SetVlanTranslationPolicyNil sets the value for VlanTranslationPolicy to be an explicit nil

### UnsetVlanTranslationPolicy
`func (o *DcimInterfacesCreateRequest) UnsetVlanTranslationPolicy()`

UnsetVlanTranslationPolicy ensures that no value is present for VlanTranslationPolicy, not even an explicit nil
### GetMarkConnected

`func (o *DcimInterfacesCreateRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *DcimInterfacesCreateRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *DcimInterfacesCreateRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *DcimInterfacesCreateRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetWirelessLans

`func (o *DcimInterfacesCreateRequest) GetWirelessLans() []int32`

GetWirelessLans returns the WirelessLans field if non-nil, zero value otherwise.

### GetWirelessLansOk

`func (o *DcimInterfacesCreateRequest) GetWirelessLansOk() (*[]int32, bool)`

GetWirelessLansOk returns a tuple with the WirelessLans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessLans

`func (o *DcimInterfacesCreateRequest) SetWirelessLans(v []int32)`

SetWirelessLans sets WirelessLans field to given value.

### HasWirelessLans

`func (o *DcimInterfacesCreateRequest) HasWirelessLans() bool`

HasWirelessLans returns a boolean if a field has been set.

### GetVrf

`func (o *DcimInterfacesCreateRequest) GetVrf() IPAddressRequestVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *DcimInterfacesCreateRequest) GetVrfOk() (*IPAddressRequestVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *DcimInterfacesCreateRequest) SetVrf(v IPAddressRequestVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *DcimInterfacesCreateRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *DcimInterfacesCreateRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *DcimInterfacesCreateRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetOwner

`func (o *DcimInterfacesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DcimInterfacesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DcimInterfacesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DcimInterfacesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *DcimInterfacesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *DcimInterfacesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *DcimInterfacesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DcimInterfacesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DcimInterfacesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DcimInterfacesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DcimInterfacesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DcimInterfacesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DcimInterfacesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DcimInterfacesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


