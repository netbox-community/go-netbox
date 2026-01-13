# CircuitTerminationRequestCircuit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** | Unique circuit ID | 
**Provider** | [**BriefCircuitRequestProvider**](BriefCircuitRequestProvider.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCircuitTerminationRequestCircuit

`func NewCircuitTerminationRequestCircuit(cid string, provider BriefCircuitRequestProvider, ) *CircuitTerminationRequestCircuit`

NewCircuitTerminationRequestCircuit instantiates a new CircuitTerminationRequestCircuit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitTerminationRequestCircuitWithDefaults

`func NewCircuitTerminationRequestCircuitWithDefaults() *CircuitTerminationRequestCircuit`

NewCircuitTerminationRequestCircuitWithDefaults instantiates a new CircuitTerminationRequestCircuit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *CircuitTerminationRequestCircuit) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *CircuitTerminationRequestCircuit) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *CircuitTerminationRequestCircuit) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProvider

`func (o *CircuitTerminationRequestCircuit) GetProvider() BriefCircuitRequestProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CircuitTerminationRequestCircuit) GetProviderOk() (*BriefCircuitRequestProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CircuitTerminationRequestCircuit) SetProvider(v BriefCircuitRequestProvider)`

SetProvider sets Provider field to given value.


### GetDescription

`func (o *CircuitTerminationRequestCircuit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CircuitTerminationRequestCircuit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CircuitTerminationRequestCircuit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CircuitTerminationRequestCircuit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


