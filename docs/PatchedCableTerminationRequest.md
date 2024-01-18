# PatchedCableTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cable** | Pointer to **int32** |  | [optional] 
**CableEnd** | Pointer to [**End**](End.md) |  | [optional] 
**TerminationType** | Pointer to **string** |  | [optional] 
**TerminationId** | Pointer to **int64** |  | [optional] 

## Methods

### NewPatchedCableTerminationRequest

`func NewPatchedCableTerminationRequest() *PatchedCableTerminationRequest`

NewPatchedCableTerminationRequest instantiates a new PatchedCableTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCableTerminationRequestWithDefaults

`func NewPatchedCableTerminationRequestWithDefaults() *PatchedCableTerminationRequest`

NewPatchedCableTerminationRequestWithDefaults instantiates a new PatchedCableTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCable

`func (o *PatchedCableTerminationRequest) GetCable() int32`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *PatchedCableTerminationRequest) GetCableOk() (*int32, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *PatchedCableTerminationRequest) SetCable(v int32)`

SetCable sets Cable field to given value.

### HasCable

`func (o *PatchedCableTerminationRequest) HasCable() bool`

HasCable returns a boolean if a field has been set.

### GetCableEnd

`func (o *PatchedCableTerminationRequest) GetCableEnd() End`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *PatchedCableTerminationRequest) GetCableEndOk() (*End, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *PatchedCableTerminationRequest) SetCableEnd(v End)`

SetCableEnd sets CableEnd field to given value.

### HasCableEnd

`func (o *PatchedCableTerminationRequest) HasCableEnd() bool`

HasCableEnd returns a boolean if a field has been set.

### GetTerminationType

`func (o *PatchedCableTerminationRequest) GetTerminationType() string`

GetTerminationType returns the TerminationType field if non-nil, zero value otherwise.

### GetTerminationTypeOk

`func (o *PatchedCableTerminationRequest) GetTerminationTypeOk() (*string, bool)`

GetTerminationTypeOk returns a tuple with the TerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationType

`func (o *PatchedCableTerminationRequest) SetTerminationType(v string)`

SetTerminationType sets TerminationType field to given value.

### HasTerminationType

`func (o *PatchedCableTerminationRequest) HasTerminationType() bool`

HasTerminationType returns a boolean if a field has been set.

### GetTerminationId

`func (o *PatchedCableTerminationRequest) GetTerminationId() int64`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *PatchedCableTerminationRequest) GetTerminationIdOk() (*int64, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *PatchedCableTerminationRequest) SetTerminationId(v int64)`

SetTerminationId sets TerminationId field to given value.

### HasTerminationId

`func (o *PatchedCableTerminationRequest) HasTerminationId() bool`

HasTerminationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


