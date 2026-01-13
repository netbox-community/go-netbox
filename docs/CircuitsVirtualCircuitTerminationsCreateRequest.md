# CircuitsVirtualCircuitTerminationsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualCircuit** | [**PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit**](PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit.md) |  | 
**Role** | Pointer to [**PatchedWritableTunnelTerminationRequestRole**](PatchedWritableTunnelTerminationRequestRole.md) |  | [optional] 
**Interface** | [**PatchedWritableVirtualCircuitTerminationRequestInterface**](PatchedWritableVirtualCircuitTerminationRequestInterface.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewCircuitsVirtualCircuitTerminationsCreateRequest

`func NewCircuitsVirtualCircuitTerminationsCreateRequest(virtualCircuit PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit, interface_ PatchedWritableVirtualCircuitTerminationRequestInterface, ) *CircuitsVirtualCircuitTerminationsCreateRequest`

NewCircuitsVirtualCircuitTerminationsCreateRequest instantiates a new CircuitsVirtualCircuitTerminationsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitsVirtualCircuitTerminationsCreateRequestWithDefaults

`func NewCircuitsVirtualCircuitTerminationsCreateRequestWithDefaults() *CircuitsVirtualCircuitTerminationsCreateRequest`

NewCircuitsVirtualCircuitTerminationsCreateRequestWithDefaults instantiates a new CircuitsVirtualCircuitTerminationsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualCircuit

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetVirtualCircuit() PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit`

GetVirtualCircuit returns the VirtualCircuit field if non-nil, zero value otherwise.

### GetVirtualCircuitOk

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetVirtualCircuitOk() (*PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit, bool)`

GetVirtualCircuitOk returns a tuple with the VirtualCircuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuit

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) SetVirtualCircuit(v PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit)`

SetVirtualCircuit sets VirtualCircuit field to given value.


### GetRole

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetRole() PatchedWritableTunnelTerminationRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetRoleOk() (*PatchedWritableTunnelTerminationRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) SetRole(v PatchedWritableTunnelTerminationRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetInterface

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetInterface() PatchedWritableVirtualCircuitTerminationRequestInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetInterfaceOk() (*PatchedWritableVirtualCircuitTerminationRequestInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) SetInterface(v PatchedWritableVirtualCircuitTerminationRequestInterface)`

SetInterface sets Interface field to given value.


### GetDescription

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CircuitsVirtualCircuitTerminationsCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


