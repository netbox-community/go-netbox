# PatchedScriptInputRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Commit** | Pointer to **bool** |  | [optional] 
**ScheduleAt** | Pointer to **NullableTime** |  | [optional] 
**Interval** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPatchedScriptInputRequest

`func NewPatchedScriptInputRequest() *PatchedScriptInputRequest`

NewPatchedScriptInputRequest instantiates a new PatchedScriptInputRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedScriptInputRequestWithDefaults

`func NewPatchedScriptInputRequestWithDefaults() *PatchedScriptInputRequest`

NewPatchedScriptInputRequestWithDefaults instantiates a new PatchedScriptInputRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PatchedScriptInputRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatchedScriptInputRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatchedScriptInputRequest) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *PatchedScriptInputRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PatchedScriptInputRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PatchedScriptInputRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetCommit

`func (o *PatchedScriptInputRequest) GetCommit() bool`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PatchedScriptInputRequest) GetCommitOk() (*bool, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PatchedScriptInputRequest) SetCommit(v bool)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *PatchedScriptInputRequest) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetScheduleAt

`func (o *PatchedScriptInputRequest) GetScheduleAt() time.Time`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *PatchedScriptInputRequest) GetScheduleAtOk() (*time.Time, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *PatchedScriptInputRequest) SetScheduleAt(v time.Time)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *PatchedScriptInputRequest) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.

### SetScheduleAtNil

`func (o *PatchedScriptInputRequest) SetScheduleAtNil(b bool)`

 SetScheduleAtNil sets the value for ScheduleAt to be an explicit nil

### UnsetScheduleAt
`func (o *PatchedScriptInputRequest) UnsetScheduleAt()`

UnsetScheduleAt ensures that no value is present for ScheduleAt, not even an explicit nil
### GetInterval

`func (o *PatchedScriptInputRequest) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PatchedScriptInputRequest) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PatchedScriptInputRequest) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PatchedScriptInputRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *PatchedScriptInputRequest) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *PatchedScriptInputRequest) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


