# CableTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cable** | **int32** |  | 
**CableEnd** | [**End1**](End1.md) |  | 
**TerminationType** | **string** |  | 
**TerminationId** | **int64** |  | 

## Methods

### NewCableTerminationRequest

`func NewCableTerminationRequest(cable int32, cableEnd End1, terminationType string, terminationId int64, ) *CableTerminationRequest`

NewCableTerminationRequest instantiates a new CableTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCableTerminationRequestWithDefaults

`func NewCableTerminationRequestWithDefaults() *CableTerminationRequest`

NewCableTerminationRequestWithDefaults instantiates a new CableTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCable

`func (o *CableTerminationRequest) GetCable() int32`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *CableTerminationRequest) GetCableOk() (*int32, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *CableTerminationRequest) SetCable(v int32)`

SetCable sets Cable field to given value.


### GetCableEnd

`func (o *CableTerminationRequest) GetCableEnd() End1`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *CableTerminationRequest) GetCableEndOk() (*End1, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *CableTerminationRequest) SetCableEnd(v End1)`

SetCableEnd sets CableEnd field to given value.


### GetTerminationType

`func (o *CableTerminationRequest) GetTerminationType() string`

GetTerminationType returns the TerminationType field if non-nil, zero value otherwise.

### GetTerminationTypeOk

`func (o *CableTerminationRequest) GetTerminationTypeOk() (*string, bool)`

GetTerminationTypeOk returns a tuple with the TerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationType

`func (o *CableTerminationRequest) SetTerminationType(v string)`

SetTerminationType sets TerminationType field to given value.


### GetTerminationId

`func (o *CableTerminationRequest) GetTerminationId() int64`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *CableTerminationRequest) GetTerminationIdOk() (*int64, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *CableTerminationRequest) SetTerminationId(v int64)`

SetTerminationId sets TerminationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


