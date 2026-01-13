# PatchedVLANTranslationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **int32** |  | [optional] 
**LocalVid** | Pointer to **int32** | Numeric VLAN ID (1-4094) | [optional] 
**RemoteVid** | Pointer to **int32** | Numeric VLAN ID (1-4094) | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedVLANTranslationRuleRequest

`func NewPatchedVLANTranslationRuleRequest() *PatchedVLANTranslationRuleRequest`

NewPatchedVLANTranslationRuleRequest instantiates a new PatchedVLANTranslationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVLANTranslationRuleRequestWithDefaults

`func NewPatchedVLANTranslationRuleRequestWithDefaults() *PatchedVLANTranslationRuleRequest`

NewPatchedVLANTranslationRuleRequestWithDefaults instantiates a new PatchedVLANTranslationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *PatchedVLANTranslationRuleRequest) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PatchedVLANTranslationRuleRequest) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PatchedVLANTranslationRuleRequest) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PatchedVLANTranslationRuleRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetLocalVid

`func (o *PatchedVLANTranslationRuleRequest) GetLocalVid() int32`

GetLocalVid returns the LocalVid field if non-nil, zero value otherwise.

### GetLocalVidOk

`func (o *PatchedVLANTranslationRuleRequest) GetLocalVidOk() (*int32, bool)`

GetLocalVidOk returns a tuple with the LocalVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVid

`func (o *PatchedVLANTranslationRuleRequest) SetLocalVid(v int32)`

SetLocalVid sets LocalVid field to given value.

### HasLocalVid

`func (o *PatchedVLANTranslationRuleRequest) HasLocalVid() bool`

HasLocalVid returns a boolean if a field has been set.

### GetRemoteVid

`func (o *PatchedVLANTranslationRuleRequest) GetRemoteVid() int32`

GetRemoteVid returns the RemoteVid field if non-nil, zero value otherwise.

### GetRemoteVidOk

`func (o *PatchedVLANTranslationRuleRequest) GetRemoteVidOk() (*int32, bool)`

GetRemoteVidOk returns a tuple with the RemoteVid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVid

`func (o *PatchedVLANTranslationRuleRequest) SetRemoteVid(v int32)`

SetRemoteVid sets RemoteVid field to given value.

### HasRemoteVid

`func (o *PatchedVLANTranslationRuleRequest) HasRemoteVid() bool`

HasRemoteVid returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedVLANTranslationRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedVLANTranslationRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedVLANTranslationRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedVLANTranslationRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


