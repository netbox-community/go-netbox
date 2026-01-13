# PatchedWritableVirtualCircuitTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualCircuit** | Pointer to [**PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit**](PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit.md) |  | [optional] 
**Role** | Pointer to [**PatchedWritableTunnelTerminationRequestRole**](PatchedWritableTunnelTerminationRequestRole.md) |  | [optional] 
**Interface** | Pointer to [**PatchedWritableVirtualCircuitTerminationRequestInterface**](PatchedWritableVirtualCircuitTerminationRequestInterface.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableVirtualCircuitTerminationRequest

`func NewPatchedWritableVirtualCircuitTerminationRequest() *PatchedWritableVirtualCircuitTerminationRequest`

NewPatchedWritableVirtualCircuitTerminationRequest instantiates a new PatchedWritableVirtualCircuitTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableVirtualCircuitTerminationRequestWithDefaults

`func NewPatchedWritableVirtualCircuitTerminationRequestWithDefaults() *PatchedWritableVirtualCircuitTerminationRequest`

NewPatchedWritableVirtualCircuitTerminationRequestWithDefaults instantiates a new PatchedWritableVirtualCircuitTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualCircuit

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetVirtualCircuit() PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit`

GetVirtualCircuit returns the VirtualCircuit field if non-nil, zero value otherwise.

### GetVirtualCircuitOk

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetVirtualCircuitOk() (*PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit, bool)`

GetVirtualCircuitOk returns a tuple with the VirtualCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuit

`func (o *PatchedWritableVirtualCircuitTerminationRequest) SetVirtualCircuit(v PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit)`

SetVirtualCircuit sets VirtualCircuit field to given value.

### HasVirtualCircuit

`func (o *PatchedWritableVirtualCircuitTerminationRequest) HasVirtualCircuit() bool`

HasVirtualCircuit returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetRole() PatchedWritableTunnelTerminationRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetRoleOk() (*PatchedWritableTunnelTerminationRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableVirtualCircuitTerminationRequest) SetRole(v PatchedWritableTunnelTerminationRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableVirtualCircuitTerminationRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetInterface

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetInterface() PatchedWritableVirtualCircuitTerminationRequestInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetInterfaceOk() (*PatchedWritableVirtualCircuitTerminationRequestInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PatchedWritableVirtualCircuitTerminationRequest) SetInterface(v PatchedWritableVirtualCircuitTerminationRequestInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PatchedWritableVirtualCircuitTerminationRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableVirtualCircuitTerminationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableVirtualCircuitTerminationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableVirtualCircuitTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableVirtualCircuitTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableVirtualCircuitTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableVirtualCircuitTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableVirtualCircuitTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


