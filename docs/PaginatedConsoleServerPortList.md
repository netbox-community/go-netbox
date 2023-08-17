# PaginatedConsoleServerPortList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]ConsoleServerPort**](ConsoleServerPort.md) |  | [optional] 

## Methods

### NewPaginatedConsoleServerPortList

`func NewPaginatedConsoleServerPortList() *PaginatedConsoleServerPortList`

NewPaginatedConsoleServerPortList instantiates a new PaginatedConsoleServerPortList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedConsoleServerPortListWithDefaults

`func NewPaginatedConsoleServerPortListWithDefaults() *PaginatedConsoleServerPortList`

NewPaginatedConsoleServerPortListWithDefaults instantiates a new PaginatedConsoleServerPortList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedConsoleServerPortList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedConsoleServerPortList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedConsoleServerPortList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedConsoleServerPortList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedConsoleServerPortList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedConsoleServerPortList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedConsoleServerPortList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedConsoleServerPortList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedConsoleServerPortList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedConsoleServerPortList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedConsoleServerPortList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedConsoleServerPortList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedConsoleServerPortList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedConsoleServerPortList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedConsoleServerPortList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedConsoleServerPortList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedConsoleServerPortList) GetResults() []ConsoleServerPort`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedConsoleServerPortList) GetResultsOk() (*[]ConsoleServerPort, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedConsoleServerPortList) SetResults(v []ConsoleServerPort)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedConsoleServerPortList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


