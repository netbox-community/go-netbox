# PaginatedL2VPNTerminationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]L2VPNTermination**](L2VPNTermination.md) |  | 

## Methods

### NewPaginatedL2VPNTerminationList

`func NewPaginatedL2VPNTerminationList(count int32, results []L2VPNTermination, ) *PaginatedL2VPNTerminationList`

NewPaginatedL2VPNTerminationList instantiates a new PaginatedL2VPNTerminationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedL2VPNTerminationListWithDefaults

`func NewPaginatedL2VPNTerminationListWithDefaults() *PaginatedL2VPNTerminationList`

NewPaginatedL2VPNTerminationListWithDefaults instantiates a new PaginatedL2VPNTerminationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedL2VPNTerminationList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedL2VPNTerminationList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedL2VPNTerminationList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedL2VPNTerminationList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedL2VPNTerminationList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedL2VPNTerminationList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedL2VPNTerminationList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedL2VPNTerminationList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedL2VPNTerminationList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedL2VPNTerminationList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedL2VPNTerminationList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedL2VPNTerminationList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedL2VPNTerminationList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedL2VPNTerminationList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedL2VPNTerminationList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedL2VPNTerminationList) GetResults() []L2VPNTermination`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedL2VPNTerminationList) GetResultsOk() (*[]L2VPNTermination, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedL2VPNTerminationList) SetResults(v []L2VPNTermination)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


