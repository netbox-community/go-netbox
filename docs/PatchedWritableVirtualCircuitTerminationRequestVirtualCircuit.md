# PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** | Unique circuit ID | 
**ProviderNetwork** | [**BriefVirtualCircuitRequestProviderNetwork**](BriefVirtualCircuitRequestProviderNetwork.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedWritableVirtualCircuitTerminationRequestVirtualCircuit

`func NewPatchedWritableVirtualCircuitTerminationRequestVirtualCircuit(cid string, providerNetwork BriefVirtualCircuitRequestProviderNetwork, ) *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit`

NewPatchedWritableVirtualCircuitTerminationRequestVirtualCircuit instantiates a new PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableVirtualCircuitTerminationRequestVirtualCircuitWithDefaults

`func NewPatchedWritableVirtualCircuitTerminationRequestVirtualCircuitWithDefaults() *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit`

NewPatchedWritableVirtualCircuitTerminationRequestVirtualCircuitWithDefaults instantiates a new PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProviderNetwork

`func (o *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit) GetProviderNetwork() BriefVirtualCircuitRequestProviderNetwork`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit) GetProviderNetworkOk() (*BriefVirtualCircuitRequestProviderNetwork, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit) SetProviderNetwork(v BriefVirtualCircuitRequestProviderNetwork)`

SetProviderNetwork sets ProviderNetwork field to given value.


### GetDescription

`func (o *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableVirtualCircuitTerminationRequestVirtualCircuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


