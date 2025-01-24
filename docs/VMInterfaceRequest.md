# VMInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualMachine** | [**BriefVirtualMachineRequest**](BriefVirtualMachineRequest.md) |  | 
**Name** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to [**NullableNestedVMInterfaceRequest**](NestedVMInterfaceRequest.md) |  | [optional] 
**Bridge** | Pointer to [**NullableNestedVMInterfaceRequest**](NestedVMInterfaceRequest.md) |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**PrimaryMacAddress** | Pointer to [**NullableBriefMACAddressRequest**](BriefMACAddressRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**InterfaceModeValue**](InterfaceModeValue.md) |  | [optional] 
**UntaggedVlan** | Pointer to [**NullableBriefVLANRequest**](BriefVLANRequest.md) |  | [optional] 
**TaggedVlans** | Pointer to **[]int32** |  | [optional] 
**QinqSvlan** | Pointer to [**NullableBriefVLANRequest**](BriefVLANRequest.md) |  | [optional] 
**VlanTranslationPolicy** | Pointer to [**NullableBriefVLANTranslationPolicyRequest**](BriefVLANTranslationPolicyRequest.md) |  | [optional] 
**Vrf** | Pointer to [**NullableBriefVRFRequest**](BriefVRFRequest.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVMInterfaceRequest

`func NewVMInterfaceRequest(virtualMachine BriefVirtualMachineRequest, name string, ) *VMInterfaceRequest`

NewVMInterfaceRequest instantiates a new VMInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInterfaceRequestWithDefaults

`func NewVMInterfaceRequestWithDefaults() *VMInterfaceRequest`

NewVMInterfaceRequestWithDefaults instantiates a new VMInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualMachine

`func (o *VMInterfaceRequest) GetVirtualMachine() BriefVirtualMachineRequest`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VMInterfaceRequest) GetVirtualMachineOk() (*BriefVirtualMachineRequest, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VMInterfaceRequest) SetVirtualMachine(v BriefVirtualMachineRequest)`

SetVirtualMachine sets VirtualMachine field to given value.


### GetName

`func (o *VMInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *VMInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VMInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VMInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VMInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParent

`func (o *VMInterfaceRequest) GetParent() NestedVMInterfaceRequest`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *VMInterfaceRequest) GetParentOk() (*NestedVMInterfaceRequest, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *VMInterfaceRequest) SetParent(v NestedVMInterfaceRequest)`

SetParent sets Parent field to given value.

### HasParent

`func (o *VMInterfaceRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *VMInterfaceRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *VMInterfaceRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBridge

`func (o *VMInterfaceRequest) GetBridge() NestedVMInterfaceRequest`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *VMInterfaceRequest) GetBridgeOk() (*NestedVMInterfaceRequest, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *VMInterfaceRequest) SetBridge(v NestedVMInterfaceRequest)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *VMInterfaceRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *VMInterfaceRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *VMInterfaceRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetMtu

`func (o *VMInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *VMInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *VMInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *VMInterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *VMInterfaceRequest) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *VMInterfaceRequest) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetPrimaryMacAddress

`func (o *VMInterfaceRequest) GetPrimaryMacAddress() BriefMACAddressRequest`

GetPrimaryMacAddress returns the PrimaryMacAddress field if non-nil, zero value otherwise.

### GetPrimaryMacAddressOk

`func (o *VMInterfaceRequest) GetPrimaryMacAddressOk() (*BriefMACAddressRequest, bool)`

GetPrimaryMacAddressOk returns a tuple with the PrimaryMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMacAddress

`func (o *VMInterfaceRequest) SetPrimaryMacAddress(v BriefMACAddressRequest)`

SetPrimaryMacAddress sets PrimaryMacAddress field to given value.

### HasPrimaryMacAddress

`func (o *VMInterfaceRequest) HasPrimaryMacAddress() bool`

HasPrimaryMacAddress returns a boolean if a field has been set.

### SetPrimaryMacAddressNil

`func (o *VMInterfaceRequest) SetPrimaryMacAddressNil(b bool)`

 SetPrimaryMacAddressNil sets the value for PrimaryMacAddress to be an explicit nil

### UnsetPrimaryMacAddress
`func (o *VMInterfaceRequest) UnsetPrimaryMacAddress()`

UnsetPrimaryMacAddress ensures that no value is present for PrimaryMacAddress, not even an explicit nil
### GetDescription

`func (o *VMInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VMInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VMInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VMInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *VMInterfaceRequest) GetMode() InterfaceModeValue`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VMInterfaceRequest) GetModeOk() (*InterfaceModeValue, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VMInterfaceRequest) SetMode(v InterfaceModeValue)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VMInterfaceRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetUntaggedVlan

`func (o *VMInterfaceRequest) GetUntaggedVlan() BriefVLANRequest`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *VMInterfaceRequest) GetUntaggedVlanOk() (*BriefVLANRequest, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *VMInterfaceRequest) SetUntaggedVlan(v BriefVLANRequest)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *VMInterfaceRequest) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *VMInterfaceRequest) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *VMInterfaceRequest) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetTaggedVlans

`func (o *VMInterfaceRequest) GetTaggedVlans() []int32`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *VMInterfaceRequest) GetTaggedVlansOk() (*[]int32, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *VMInterfaceRequest) SetTaggedVlans(v []int32)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *VMInterfaceRequest) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetQinqSvlan

`func (o *VMInterfaceRequest) GetQinqSvlan() BriefVLANRequest`

GetQinqSvlan returns the QinqSvlan field if non-nil, zero value otherwise.

### GetQinqSvlanOk

`func (o *VMInterfaceRequest) GetQinqSvlanOk() (*BriefVLANRequest, bool)`

GetQinqSvlanOk returns a tuple with the QinqSvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqSvlan

`func (o *VMInterfaceRequest) SetQinqSvlan(v BriefVLANRequest)`

SetQinqSvlan sets QinqSvlan field to given value.

### HasQinqSvlan

`func (o *VMInterfaceRequest) HasQinqSvlan() bool`

HasQinqSvlan returns a boolean if a field has been set.

### SetQinqSvlanNil

`func (o *VMInterfaceRequest) SetQinqSvlanNil(b bool)`

 SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil

### UnsetQinqSvlan
`func (o *VMInterfaceRequest) UnsetQinqSvlan()`

UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
### GetVlanTranslationPolicy

`func (o *VMInterfaceRequest) GetVlanTranslationPolicy() BriefVLANTranslationPolicyRequest`

GetVlanTranslationPolicy returns the VlanTranslationPolicy field if non-nil, zero value otherwise.

### GetVlanTranslationPolicyOk

`func (o *VMInterfaceRequest) GetVlanTranslationPolicyOk() (*BriefVLANTranslationPolicyRequest, bool)`

GetVlanTranslationPolicyOk returns a tuple with the VlanTranslationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTranslationPolicy

`func (o *VMInterfaceRequest) SetVlanTranslationPolicy(v BriefVLANTranslationPolicyRequest)`

SetVlanTranslationPolicy sets VlanTranslationPolicy field to given value.

### HasVlanTranslationPolicy

`func (o *VMInterfaceRequest) HasVlanTranslationPolicy() bool`

HasVlanTranslationPolicy returns a boolean if a field has been set.

### SetVlanTranslationPolicyNil

`func (o *VMInterfaceRequest) SetVlanTranslationPolicyNil(b bool)`

 SetVlanTranslationPolicyNil sets the value for VlanTranslationPolicy to be an explicit nil

### UnsetVlanTranslationPolicy
`func (o *VMInterfaceRequest) UnsetVlanTranslationPolicy()`

UnsetVlanTranslationPolicy ensures that no value is present for VlanTranslationPolicy, not even an explicit nil
### GetVrf

`func (o *VMInterfaceRequest) GetVrf() BriefVRFRequest`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *VMInterfaceRequest) GetVrfOk() (*BriefVRFRequest, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *VMInterfaceRequest) SetVrf(v BriefVRFRequest)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *VMInterfaceRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *VMInterfaceRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *VMInterfaceRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTags

`func (o *VMInterfaceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VMInterfaceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VMInterfaceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VMInterfaceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VMInterfaceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VMInterfaceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VMInterfaceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VMInterfaceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


