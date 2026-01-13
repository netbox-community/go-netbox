# PaginatedVirtualMachineWithConfigContextList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]VirtualMachineWithConfigContext**](VirtualMachineWithConfigContext.md) |  | 

## Methods

### NewPaginatedVirtualMachineWithConfigContextList

`func NewPaginatedVirtualMachineWithConfigContextList(count int32, results []VirtualMachineWithConfigContext, ) *PaginatedVirtualMachineWithConfigContextList`

NewPaginatedVirtualMachineWithConfigContextList instantiates a new PaginatedVirtualMachineWithConfigContextList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedVirtualMachineWithConfigContextListWithDefaults

`func NewPaginatedVirtualMachineWithConfigContextListWithDefaults() *PaginatedVirtualMachineWithConfigContextList`

NewPaginatedVirtualMachineWithConfigContextListWithDefaults instantiates a new PaginatedVirtualMachineWithConfigContextList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedVirtualMachineWithConfigContextList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedVirtualMachineWithConfigContextList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedVirtualMachineWithConfigContextList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedVirtualMachineWithConfigContextList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedVirtualMachineWithConfigContextList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedVirtualMachineWithConfigContextList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedVirtualMachineWithConfigContextList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedVirtualMachineWithConfigContextList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedVirtualMachineWithConfigContextList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedVirtualMachineWithConfigContextList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedVirtualMachineWithConfigContextList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedVirtualMachineWithConfigContextList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedVirtualMachineWithConfigContextList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedVirtualMachineWithConfigContextList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedVirtualMachineWithConfigContextList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedVirtualMachineWithConfigContextList) GetResults() []VirtualMachineWithConfigContext`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedVirtualMachineWithConfigContextList) GetResultsOk() (*[]VirtualMachineWithConfigContext, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedVirtualMachineWithConfigContextList) SetResults(v []VirtualMachineWithConfigContext)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


