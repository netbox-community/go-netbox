# PatchedWritableInterfaceTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **NullableInt32** |  | [optional] 
**ModuleType** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |          {module} is accepted as a substitution for the module bay position when attached to a module type.          | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to **string** | * &#x60;virtual&#x60; - Virtual * &#x60;bridge&#x60; - Bridge * &#x60;lag&#x60; - Link Aggregation Group (LAG) * &#x60;100base-fx&#x60; - 100BASE-FX (10/100ME FIBER) * &#x60;100base-lfx&#x60; - 100BASE-LFX (10/100ME FIBER) * &#x60;100base-tx&#x60; - 100BASE-TX (10/100ME) * &#x60;100base-t1&#x60; - 100BASE-T1 (10/100ME Single Pair) * &#x60;1000base-t&#x60; - 1000BASE-T (1GE) * &#x60;2.5gbase-t&#x60; - 2.5GBASE-T (2.5GE) * &#x60;5gbase-t&#x60; - 5GBASE-T (5GE) * &#x60;10gbase-t&#x60; - 10GBASE-T (10GE) * &#x60;10gbase-cx4&#x60; - 10GBASE-CX4 (10GE) * &#x60;1000base-x-gbic&#x60; - GBIC (1GE) * &#x60;1000base-x-sfp&#x60; - SFP (1GE) * &#x60;10gbase-x-sfpp&#x60; - SFP+ (10GE) * &#x60;10gbase-x-xfp&#x60; - XFP (10GE) * &#x60;10gbase-x-xenpak&#x60; - XENPAK (10GE) * &#x60;10gbase-x-x2&#x60; - X2 (10GE) * &#x60;25gbase-x-sfp28&#x60; - SFP28 (25GE) * &#x60;50gbase-x-sfp56&#x60; - SFP56 (50GE) * &#x60;40gbase-x-qsfpp&#x60; - QSFP+ (40GE) * &#x60;50gbase-x-sfp28&#x60; - QSFP28 (50GE) * &#x60;100gbase-x-cfp&#x60; - CFP (100GE) * &#x60;100gbase-x-cfp2&#x60; - CFP2 (100GE) * &#x60;200gbase-x-cfp2&#x60; - CFP2 (200GE) * &#x60;400gbase-x-cfp2&#x60; - CFP2 (400GE) * &#x60;100gbase-x-cfp4&#x60; - CFP4 (100GE) * &#x60;100gbase-x-cxp&#x60; - CXP (100GE) * &#x60;100gbase-x-cpak&#x60; - Cisco CPAK (100GE) * &#x60;100gbase-x-dsfp&#x60; - DSFP (100GE) * &#x60;100gbase-x-sfpdd&#x60; - SFP-DD (100GE) * &#x60;100gbase-x-qsfp28&#x60; - QSFP28 (100GE) * &#x60;100gbase-x-qsfpdd&#x60; - QSFP-DD (100GE) * &#x60;200gbase-x-qsfp56&#x60; - QSFP56 (200GE) * &#x60;200gbase-x-qsfpdd&#x60; - QSFP-DD (200GE) * &#x60;400gbase-x-qsfpdd&#x60; - QSFP-DD (400GE) * &#x60;400gbase-x-osfp&#x60; - OSFP (400GE) * &#x60;400gbase-x-cdfp&#x60; - CDFP (400GE) * &#x60;400gbase-x-cfp8&#x60; - CPF8 (400GE) * &#x60;800gbase-x-qsfpdd&#x60; - QSFP-DD (800GE) * &#x60;800gbase-x-osfp&#x60; - OSFP (800GE) * &#x60;1000base-kx&#x60; - 1000BASE-KX (1GE) * &#x60;10gbase-kr&#x60; - 10GBASE-KR (10GE) * &#x60;10gbase-kx4&#x60; - 10GBASE-KX4 (10GE) * &#x60;25gbase-kr&#x60; - 25GBASE-KR (25GE) * &#x60;40gbase-kr4&#x60; - 40GBASE-KR4 (40GE) * &#x60;50gbase-kr&#x60; - 50GBASE-KR (50GE) * &#x60;100gbase-kp4&#x60; - 100GBASE-KP4 (100GE) * &#x60;100gbase-kr2&#x60; - 100GBASE-KR2 (100GE) * &#x60;100gbase-kr4&#x60; - 100GBASE-KR4 (100GE) * &#x60;ieee802.11a&#x60; - IEEE 802.11a * &#x60;ieee802.11g&#x60; - IEEE 802.11b/g * &#x60;ieee802.11n&#x60; - IEEE 802.11n * &#x60;ieee802.11ac&#x60; - IEEE 802.11ac * &#x60;ieee802.11ad&#x60; - IEEE 802.11ad * &#x60;ieee802.11ax&#x60; - IEEE 802.11ax * &#x60;ieee802.11ay&#x60; - IEEE 802.11ay * &#x60;ieee802.15.1&#x60; - IEEE 802.15.1 (Bluetooth) * &#x60;other-wireless&#x60; - Other (Wireless) * &#x60;gsm&#x60; - GSM * &#x60;cdma&#x60; - CDMA * &#x60;lte&#x60; - LTE * &#x60;sonet-oc3&#x60; - OC-3/STM-1 * &#x60;sonet-oc12&#x60; - OC-12/STM-4 * &#x60;sonet-oc48&#x60; - OC-48/STM-16 * &#x60;sonet-oc192&#x60; - OC-192/STM-64 * &#x60;sonet-oc768&#x60; - OC-768/STM-256 * &#x60;sonet-oc1920&#x60; - OC-1920/STM-640 * &#x60;sonet-oc3840&#x60; - OC-3840/STM-1234 * &#x60;1gfc-sfp&#x60; - SFP (1GFC) * &#x60;2gfc-sfp&#x60; - SFP (2GFC) * &#x60;4gfc-sfp&#x60; - SFP (4GFC) * &#x60;8gfc-sfpp&#x60; - SFP+ (8GFC) * &#x60;16gfc-sfpp&#x60; - SFP+ (16GFC) * &#x60;32gfc-sfp28&#x60; - SFP28 (32GFC) * &#x60;64gfc-qsfpp&#x60; - QSFP+ (64GFC) * &#x60;128gfc-qsfp28&#x60; - QSFP28 (128GFC) * &#x60;infiniband-sdr&#x60; - SDR (2 Gbps) * &#x60;infiniband-ddr&#x60; - DDR (4 Gbps) * &#x60;infiniband-qdr&#x60; - QDR (8 Gbps) * &#x60;infiniband-fdr10&#x60; - FDR10 (10 Gbps) * &#x60;infiniband-fdr&#x60; - FDR (13.5 Gbps) * &#x60;infiniband-edr&#x60; - EDR (25 Gbps) * &#x60;infiniband-hdr&#x60; - HDR (50 Gbps) * &#x60;infiniband-ndr&#x60; - NDR (100 Gbps) * &#x60;infiniband-xdr&#x60; - XDR (250 Gbps) * &#x60;t1&#x60; - T1 (1.544 Mbps) * &#x60;e1&#x60; - E1 (2.048 Mbps) * &#x60;t3&#x60; - T3 (45 Mbps) * &#x60;e3&#x60; - E3 (34 Mbps) * &#x60;xdsl&#x60; - xDSL * &#x60;docsis&#x60; - DOCSIS * &#x60;gpon&#x60; - GPON (2.5 Gbps / 1.25 Gps) * &#x60;xg-pon&#x60; - XG-PON (10 Gbps / 2.5 Gbps) * &#x60;xgs-pon&#x60; - XGS-PON (10 Gbps) * &#x60;ng-pon2&#x60; - NG-PON2 (TWDM-PON) (4x10 Gbps) * &#x60;epon&#x60; - EPON (1 Gbps) * &#x60;10g-epon&#x60; - 10G-EPON (10 Gbps) * &#x60;cisco-stackwise&#x60; - Cisco StackWise * &#x60;cisco-stackwise-plus&#x60; - Cisco StackWise Plus * &#x60;cisco-flexstack&#x60; - Cisco FlexStack * &#x60;cisco-flexstack-plus&#x60; - Cisco FlexStack Plus * &#x60;cisco-stackwise-80&#x60; - Cisco StackWise-80 * &#x60;cisco-stackwise-160&#x60; - Cisco StackWise-160 * &#x60;cisco-stackwise-320&#x60; - Cisco StackWise-320 * &#x60;cisco-stackwise-480&#x60; - Cisco StackWise-480 * &#x60;cisco-stackwise-1t&#x60; - Cisco StackWise-1T * &#x60;juniper-vcp&#x60; - Juniper VCP * &#x60;extreme-summitstack&#x60; - Extreme SummitStack * &#x60;extreme-summitstack-128&#x60; - Extreme SummitStack-128 * &#x60;extreme-summitstack-256&#x60; - Extreme SummitStack-256 * &#x60;extreme-summitstack-512&#x60; - Extreme SummitStack-512 * &#x60;other&#x60; - Other | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MgmtOnly** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Bridge** | Pointer to **NullableInt32** |  | [optional] 
**PoeMode** | Pointer to **string** | * &#x60;pd&#x60; - PD * &#x60;pse&#x60; - PSE | [optional] 
**PoeType** | Pointer to **string** | * &#x60;type1-ieee802.3af&#x60; - 802.3af (Type 1) * &#x60;type2-ieee802.3at&#x60; - 802.3at (Type 2) * &#x60;type3-ieee802.3bt&#x60; - 802.3bt (Type 3) * &#x60;type4-ieee802.3bt&#x60; - 802.3bt (Type 4) * &#x60;passive-24v-2pair&#x60; - Passive 24V (2-pair) * &#x60;passive-24v-4pair&#x60; - Passive 24V (4-pair) * &#x60;passive-48v-2pair&#x60; - Passive 48V (2-pair) * &#x60;passive-48v-4pair&#x60; - Passive 48V (4-pair) | [optional] 

## Methods

### NewPatchedWritableInterfaceTemplateRequest

`func NewPatchedWritableInterfaceTemplateRequest() *PatchedWritableInterfaceTemplateRequest`

NewPatchedWritableInterfaceTemplateRequest instantiates a new PatchedWritableInterfaceTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableInterfaceTemplateRequestWithDefaults

`func NewPatchedWritableInterfaceTemplateRequestWithDefaults() *PatchedWritableInterfaceTemplateRequest`

NewPatchedWritableInterfaceTemplateRequestWithDefaults instantiates a new PatchedWritableInterfaceTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *PatchedWritableInterfaceTemplateRequest) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritableInterfaceTemplateRequest) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritableInterfaceTemplateRequest) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritableInterfaceTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *PatchedWritableInterfaceTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *PatchedWritableInterfaceTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *PatchedWritableInterfaceTemplateRequest) GetModuleType() int32`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *PatchedWritableInterfaceTemplateRequest) GetModuleTypeOk() (*int32, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *PatchedWritableInterfaceTemplateRequest) SetModuleType(v int32)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *PatchedWritableInterfaceTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *PatchedWritableInterfaceTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *PatchedWritableInterfaceTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetName

`func (o *PatchedWritableInterfaceTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableInterfaceTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableInterfaceTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableInterfaceTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableInterfaceTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableInterfaceTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableInterfaceTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableInterfaceTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableInterfaceTemplateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableInterfaceTemplateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableInterfaceTemplateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableInterfaceTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedWritableInterfaceTemplateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedWritableInterfaceTemplateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedWritableInterfaceTemplateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedWritableInterfaceTemplateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMgmtOnly

`func (o *PatchedWritableInterfaceTemplateRequest) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *PatchedWritableInterfaceTemplateRequest) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *PatchedWritableInterfaceTemplateRequest) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *PatchedWritableInterfaceTemplateRequest) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableInterfaceTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableInterfaceTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableInterfaceTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableInterfaceTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBridge

`func (o *PatchedWritableInterfaceTemplateRequest) GetBridge() int32`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *PatchedWritableInterfaceTemplateRequest) GetBridgeOk() (*int32, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *PatchedWritableInterfaceTemplateRequest) SetBridge(v int32)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *PatchedWritableInterfaceTemplateRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *PatchedWritableInterfaceTemplateRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *PatchedWritableInterfaceTemplateRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetPoeMode

`func (o *PatchedWritableInterfaceTemplateRequest) GetPoeMode() string`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *PatchedWritableInterfaceTemplateRequest) GetPoeModeOk() (*string, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *PatchedWritableInterfaceTemplateRequest) SetPoeMode(v string)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *PatchedWritableInterfaceTemplateRequest) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### GetPoeType

`func (o *PatchedWritableInterfaceTemplateRequest) GetPoeType() string`

GetPoeType returns the PoeType field if non-nil, zero value otherwise.

### GetPoeTypeOk

`func (o *PatchedWritableInterfaceTemplateRequest) GetPoeTypeOk() (*string, bool)`

GetPoeTypeOk returns a tuple with the PoeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeType

`func (o *PatchedWritableInterfaceTemplateRequest) SetPoeType(v string)`

SetPoeType sets PoeType field to given value.

### HasPoeType

`func (o *PatchedWritableInterfaceTemplateRequest) HasPoeType() bool`

HasPoeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


