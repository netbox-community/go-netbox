# PatchedWritableVMInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualMachine** | Pointer to [**BriefVirtualMachineRequest**](BriefVirtualMachineRequest.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to **NullableInt32** |  | [optional] 
**Bridge** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**PrimaryMacAddress** | Pointer to [**NullableBriefMACAddressRequest**](BriefMACAddressRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to [**NullablePatchedWritableInterfaceRequestMode**](PatchedWritableInterfaceRequestMode.md) |  | [optional] 
**UntaggedVlan** | Pointer to [**NullableBriefVLANRequest**](BriefVLANRequest.md) |  | [optional] 
**TaggedVlans** | Pointer to **[]int32** |  | [optional] 
**QinqSvlan** | Pointer to [**NullableBriefVLANRequest**](BriefVLANRequest.md) |  | [optional] 
**VlanTranslationPolicy** | Pointer to [**NullableBriefVLANTranslationPolicyRequest**](BriefVLANTranslationPolicyRequest.md) |  | [optional] 
**Vrf** | Pointer to [**NullableBriefVRFRequest**](BriefVRFRequest.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableVMInterfaceRequest

`func NewPatchedWritableVMInterfaceRequest() *PatchedWritableVMInterfaceRequest`

NewPatchedWritableVMInterfaceRequest instantiates a new PatchedWritableVMInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableVMInterfaceRequestWithDefaults

`func NewPatchedWritableVMInterfaceRequestWithDefaults() *PatchedWritableVMInterfaceRequest`

NewPatchedWritableVMInterfaceRequestWithDefaults instantiates a new PatchedWritableVMInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualMachine

`func (o *PatchedWritableVMInterfaceRequest) GetVirtualMachine() BriefVirtualMachineRequest`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *PatchedWritableVMInterfaceRequest) GetVirtualMachineOk() (*BriefVirtualMachineRequest, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *PatchedWritableVMInterfaceRequest) SetVirtualMachine(v BriefVirtualMachineRequest)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *PatchedWritableVMInterfaceRequest) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritableVMInterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableVMInterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableVMInterfaceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableVMInterfaceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedWritableVMInterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedWritableVMInterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedWritableVMInterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedWritableVMInterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParent

`func (o *PatchedWritableVMInterfaceRequest) GetParent() int32`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedWritableVMInterfaceRequest) GetParentOk() (*int32, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedWritableVMInterfaceRequest) SetParent(v int32)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedWritableVMInterfaceRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedWritableVMInterfaceRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedWritableVMInterfaceRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBridge

`func (o *PatchedWritableVMInterfaceRequest) GetBridge() int32`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *PatchedWritableVMInterfaceRequest) GetBridgeOk() (*int32, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *PatchedWritableVMInterfaceRequest) SetBridge(v int32)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *PatchedWritableVMInterfaceRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *PatchedWritableVMInterfaceRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *PatchedWritableVMInterfaceRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetMtu

`func (o *PatchedWritableVMInterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *PatchedWritableVMInterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *PatchedWritableVMInterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *PatchedWritableVMInterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *PatchedWritableVMInterfaceRequest) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *PatchedWritableVMInterfaceRequest) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetPrimaryMacAddress

`func (o *PatchedWritableVMInterfaceRequest) GetPrimaryMacAddress() BriefMACAddressRequest`

GetPrimaryMacAddress returns the PrimaryMacAddress field if non-nil, zero value otherwise.

### GetPrimaryMacAddressOk

`func (o *PatchedWritableVMInterfaceRequest) GetPrimaryMacAddressOk() (*BriefMACAddressRequest, bool)`

GetPrimaryMacAddressOk returns a tuple with the PrimaryMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMacAddress

`func (o *PatchedWritableVMInterfaceRequest) SetPrimaryMacAddress(v BriefMACAddressRequest)`

SetPrimaryMacAddress sets PrimaryMacAddress field to given value.

### HasPrimaryMacAddress

`func (o *PatchedWritableVMInterfaceRequest) HasPrimaryMacAddress() bool`

HasPrimaryMacAddress returns a boolean if a field has been set.

### SetPrimaryMacAddressNil

`func (o *PatchedWritableVMInterfaceRequest) SetPrimaryMacAddressNil(b bool)`

 SetPrimaryMacAddressNil sets the value for PrimaryMacAddress to be an explicit nil

### UnsetPrimaryMacAddress
`func (o *PatchedWritableVMInterfaceRequest) UnsetPrimaryMacAddress()`

UnsetPrimaryMacAddress ensures that no value is present for PrimaryMacAddress, not even an explicit nil
### GetDescription

`func (o *PatchedWritableVMInterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableVMInterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableVMInterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableVMInterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *PatchedWritableVMInterfaceRequest) GetMode() PatchedWritableInterfaceRequestMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PatchedWritableVMInterfaceRequest) GetModeOk() (*PatchedWritableInterfaceRequestMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PatchedWritableVMInterfaceRequest) SetMode(v PatchedWritableInterfaceRequestMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PatchedWritableVMInterfaceRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *PatchedWritableVMInterfaceRequest) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *PatchedWritableVMInterfaceRequest) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetUntaggedVlan

`func (o *PatchedWritableVMInterfaceRequest) GetUntaggedVlan() BriefVLANRequest`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *PatchedWritableVMInterfaceRequest) GetUntaggedVlanOk() (*BriefVLANRequest, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *PatchedWritableVMInterfaceRequest) SetUntaggedVlan(v BriefVLANRequest)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *PatchedWritableVMInterfaceRequest) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *PatchedWritableVMInterfaceRequest) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *PatchedWritableVMInterfaceRequest) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetTaggedVlans

`func (o *PatchedWritableVMInterfaceRequest) GetTaggedVlans() []int32`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *PatchedWritableVMInterfaceRequest) GetTaggedVlansOk() (*[]int32, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *PatchedWritableVMInterfaceRequest) SetTaggedVlans(v []int32)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *PatchedWritableVMInterfaceRequest) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetQinqSvlan

`func (o *PatchedWritableVMInterfaceRequest) GetQinqSvlan() BriefVLANRequest`

GetQinqSvlan returns the QinqSvlan field if non-nil, zero value otherwise.

### GetQinqSvlanOk

`func (o *PatchedWritableVMInterfaceRequest) GetQinqSvlanOk() (*BriefVLANRequest, bool)`

GetQinqSvlanOk returns a tuple with the QinqSvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqSvlan

`func (o *PatchedWritableVMInterfaceRequest) SetQinqSvlan(v BriefVLANRequest)`

SetQinqSvlan sets QinqSvlan field to given value.

### HasQinqSvlan

`func (o *PatchedWritableVMInterfaceRequest) HasQinqSvlan() bool`

HasQinqSvlan returns a boolean if a field has been set.

### SetQinqSvlanNil

`func (o *PatchedWritableVMInterfaceRequest) SetQinqSvlanNil(b bool)`

 SetQinqSvlanNil sets the value for QinqSvlan to be an explicit nil

### UnsetQinqSvlan
`func (o *PatchedWritableVMInterfaceRequest) UnsetQinqSvlan()`

UnsetQinqSvlan ensures that no value is present for QinqSvlan, not even an explicit nil
### GetVlanTranslationPolicy

`func (o *PatchedWritableVMInterfaceRequest) GetVlanTranslationPolicy() BriefVLANTranslationPolicyRequest`

GetVlanTranslationPolicy returns the VlanTranslationPolicy field if non-nil, zero value otherwise.

### GetVlanTranslationPolicyOk

`func (o *PatchedWritableVMInterfaceRequest) GetVlanTranslationPolicyOk() (*BriefVLANTranslationPolicyRequest, bool)`

GetVlanTranslationPolicyOk returns a tuple with the VlanTranslationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTranslationPolicy

`func (o *PatchedWritableVMInterfaceRequest) SetVlanTranslationPolicy(v BriefVLANTranslationPolicyRequest)`

SetVlanTranslationPolicy sets VlanTranslationPolicy field to given value.

### HasVlanTranslationPolicy

`func (o *PatchedWritableVMInterfaceRequest) HasVlanTranslationPolicy() bool`

HasVlanTranslationPolicy returns a boolean if a field has been set.

### SetVlanTranslationPolicyNil

`func (o *PatchedWritableVMInterfaceRequest) SetVlanTranslationPolicyNil(b bool)`

 SetVlanTranslationPolicyNil sets the value for VlanTranslationPolicy to be an explicit nil

### UnsetVlanTranslationPolicy
`func (o *PatchedWritableVMInterfaceRequest) UnsetVlanTranslationPolicy()`

UnsetVlanTranslationPolicy ensures that no value is present for VlanTranslationPolicy, not even an explicit nil
### GetVrf

`func (o *PatchedWritableVMInterfaceRequest) GetVrf() BriefVRFRequest`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *PatchedWritableVMInterfaceRequest) GetVrfOk() (*BriefVRFRequest, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *PatchedWritableVMInterfaceRequest) SetVrf(v BriefVRFRequest)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *PatchedWritableVMInterfaceRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *PatchedWritableVMInterfaceRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *PatchedWritableVMInterfaceRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTags

`func (o *PatchedWritableVMInterfaceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableVMInterfaceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableVMInterfaceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableVMInterfaceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableVMInterfaceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableVMInterfaceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableVMInterfaceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableVMInterfaceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


