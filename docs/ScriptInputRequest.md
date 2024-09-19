# ScriptInputRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **interface{}** |  | 
**Commit** | **bool** |  | 
**ScheduleAt** | Pointer to **NullableTime** |  | [optional] 
**Interval** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewScriptInputRequest

`func NewScriptInputRequest(data interface{}, commit bool, ) *ScriptInputRequest`

NewScriptInputRequest instantiates a new ScriptInputRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptInputRequestWithDefaults

`func NewScriptInputRequestWithDefaults() *ScriptInputRequest`

NewScriptInputRequestWithDefaults instantiates a new ScriptInputRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ScriptInputRequest) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ScriptInputRequest) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ScriptInputRequest) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *ScriptInputRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ScriptInputRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetCommit

`func (o *ScriptInputRequest) GetCommit() bool`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ScriptInputRequest) GetCommitOk() (*bool, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ScriptInputRequest) SetCommit(v bool)`

SetCommit sets Commit field to given value.


### GetScheduleAt

`func (o *ScriptInputRequest) GetScheduleAt() time.Time`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *ScriptInputRequest) GetScheduleAtOk() (*time.Time, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *ScriptInputRequest) SetScheduleAt(v time.Time)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *ScriptInputRequest) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.

### SetScheduleAtNil

`func (o *ScriptInputRequest) SetScheduleAtNil(b bool)`

 SetScheduleAtNil sets the value for ScheduleAt to be an explicit nil

### UnsetScheduleAt
`func (o *ScriptInputRequest) UnsetScheduleAt()`

UnsetScheduleAt ensures that no value is present for ScheduleAt, not even an explicit nil
### GetInterval

`func (o *ScriptInputRequest) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ScriptInputRequest) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ScriptInputRequest) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ScriptInputRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *ScriptInputRequest) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *ScriptInputRequest) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


