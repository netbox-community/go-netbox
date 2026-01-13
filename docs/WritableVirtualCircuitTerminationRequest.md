# WritableVirtualCircuitTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualCircuit** | [**PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit**](PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit.md) |  | 
**Role** | Pointer to [**PatchedWritableTunnelTerminationRequestRole**](PatchedWritableTunnelTerminationRequestRole.md) |  | [optional] 
**Interface** | [**PatchedWritableVirtualCircuitTerminationRequestInterface**](PatchedWritableVirtualCircuitTerminationRequestInterface.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableVirtualCircuitTerminationRequest

`func NewWritableVirtualCircuitTerminationRequest(virtualCircuit PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit, interface_ PatchedWritableVirtualCircuitTerminationRequestInterface, ) *WritableVirtualCircuitTerminationRequest`

NewWritableVirtualCircuitTerminationRequest instantiates a new WritableVirtualCircuitTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableVirtualCircuitTerminationRequestWithDefaults

`func NewWritableVirtualCircuitTerminationRequestWithDefaults() *WritableVirtualCircuitTerminationRequest`

NewWritableVirtualCircuitTerminationRequestWithDefaults instantiates a new WritableVirtualCircuitTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualCircuit

`func (o *WritableVirtualCircuitTerminationRequest) GetVirtualCircuit() PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit`

GetVirtualCircuit returns the VirtualCircuit field if non-nil, zero value otherwise.

### GetVirtualCircuitOk

`func (o *WritableVirtualCircuitTerminationRequest) GetVirtualCircuitOk() (*PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit, bool)`

GetVirtualCircuitOk returns a tuple with the VirtualCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuit

`func (o *WritableVirtualCircuitTerminationRequest) SetVirtualCircuit(v PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit)`

SetVirtualCircuit sets VirtualCircuit field to given value.


### GetRole

`func (o *WritableVirtualCircuitTerminationRequest) GetRole() PatchedWritableTunnelTerminationRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *WritableVirtualCircuitTerminationRequest) GetRoleOk() (*PatchedWritableTunnelTerminationRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *WritableVirtualCircuitTerminationRequest) SetRole(v PatchedWritableTunnelTerminationRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *WritableVirtualCircuitTerminationRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetInterface

`func (o *WritableVirtualCircuitTerminationRequest) GetInterface() PatchedWritableVirtualCircuitTerminationRequestInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *WritableVirtualCircuitTerminationRequest) GetInterfaceOk() (*PatchedWritableVirtualCircuitTerminationRequestInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *WritableVirtualCircuitTerminationRequest) SetInterface(v PatchedWritableVirtualCircuitTerminationRequestInterface)`

SetInterface sets Interface field to given value.


### GetDescription

`func (o *WritableVirtualCircuitTerminationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableVirtualCircuitTerminationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableVirtualCircuitTerminationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableVirtualCircuitTerminationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *WritableVirtualCircuitTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableVirtualCircuitTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableVirtualCircuitTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableVirtualCircuitTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableVirtualCircuitTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableVirtualCircuitTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableVirtualCircuitTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableVirtualCircuitTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


