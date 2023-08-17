# PaginatedConsolePortTemplateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]ConsolePortTemplate**](ConsolePortTemplate.md) |  | [optional] 

## Methods

### NewPaginatedConsolePortTemplateList

`func NewPaginatedConsolePortTemplateList() *PaginatedConsolePortTemplateList`

NewPaginatedConsolePortTemplateList instantiates a new PaginatedConsolePortTemplateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedConsolePortTemplateListWithDefaults

`func NewPaginatedConsolePortTemplateListWithDefaults() *PaginatedConsolePortTemplateList`

NewPaginatedConsolePortTemplateListWithDefaults instantiates a new PaginatedConsolePortTemplateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedConsolePortTemplateList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedConsolePortTemplateList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedConsolePortTemplateList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaginatedConsolePortTemplateList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedConsolePortTemplateList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedConsolePortTemplateList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedConsolePortTemplateList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedConsolePortTemplateList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedConsolePortTemplateList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedConsolePortTemplateList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedConsolePortTemplateList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedConsolePortTemplateList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedConsolePortTemplateList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedConsolePortTemplateList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedConsolePortTemplateList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedConsolePortTemplateList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedConsolePortTemplateList) GetResults() []ConsolePortTemplate`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedConsolePortTemplateList) GetResultsOk() (*[]ConsolePortTemplate, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedConsolePortTemplateList) SetResults(v []ConsolePortTemplate)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedConsolePortTemplateList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


