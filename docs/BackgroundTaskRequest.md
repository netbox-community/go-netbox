# BackgroundTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Description** | **string** |  | 
**Origin** | **string** |  | 
**FuncName** | **string** |  | 
**Result** | **string** |  | 
**Timeout** | **int32** |  | 
**ResultTtl** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**EnqueuedAt** | **time.Time** |  | 
**StartedAt** | **time.Time** |  | 
**EndedAt** | **time.Time** |  | 
**WorkerName** | **string** |  | 
**Meta** | **map[string]interface{}** |  | 
**LastHeartbeat** | **string** |  | 
**IsFinished** | **bool** |  | 
**IsQueued** | **bool** |  | 
**IsFailed** | **bool** |  | 
**IsStarted** | **bool** |  | 
**IsDeferred** | **bool** |  | 
**IsCanceled** | **bool** |  | 
**IsScheduled** | **bool** |  | 
**IsStopped** | **bool** |  | 

## Methods

### NewBackgroundTaskRequest

`func NewBackgroundTaskRequest(id string, description string, origin string, funcName string, result string, timeout int32, resultTtl int32, createdAt time.Time, enqueuedAt time.Time, startedAt time.Time, endedAt time.Time, workerName string, meta map[string]interface{}, lastHeartbeat string, isFinished bool, isQueued bool, isFailed bool, isStarted bool, isDeferred bool, isCanceled bool, isScheduled bool, isStopped bool, ) *BackgroundTaskRequest`

NewBackgroundTaskRequest instantiates a new BackgroundTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundTaskRequestWithDefaults

`func NewBackgroundTaskRequestWithDefaults() *BackgroundTaskRequest`

NewBackgroundTaskRequestWithDefaults instantiates a new BackgroundTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackgroundTaskRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackgroundTaskRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackgroundTaskRequest) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *BackgroundTaskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackgroundTaskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackgroundTaskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOrigin

`func (o *BackgroundTaskRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *BackgroundTaskRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *BackgroundTaskRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetFuncName

`func (o *BackgroundTaskRequest) GetFuncName() string`

GetFuncName returns the FuncName field if non-nil, zero value otherwise.

### GetFuncNameOk

`func (o *BackgroundTaskRequest) GetFuncNameOk() (*string, bool)`

GetFuncNameOk returns a tuple with the FuncName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncName

`func (o *BackgroundTaskRequest) SetFuncName(v string)`

SetFuncName sets FuncName field to given value.


### GetResult

`func (o *BackgroundTaskRequest) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BackgroundTaskRequest) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BackgroundTaskRequest) SetResult(v string)`

SetResult sets Result field to given value.


### GetTimeout

`func (o *BackgroundTaskRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BackgroundTaskRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BackgroundTaskRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetResultTtl

`func (o *BackgroundTaskRequest) GetResultTtl() int32`

GetResultTtl returns the ResultTtl field if non-nil, zero value otherwise.

### GetResultTtlOk

`func (o *BackgroundTaskRequest) GetResultTtlOk() (*int32, bool)`

GetResultTtlOk returns a tuple with the ResultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTtl

`func (o *BackgroundTaskRequest) SetResultTtl(v int32)`

SetResultTtl sets ResultTtl field to given value.


### GetCreatedAt

`func (o *BackgroundTaskRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackgroundTaskRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackgroundTaskRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnqueuedAt

`func (o *BackgroundTaskRequest) GetEnqueuedAt() time.Time`

GetEnqueuedAt returns the EnqueuedAt field if non-nil, zero value otherwise.

### GetEnqueuedAtOk

`func (o *BackgroundTaskRequest) GetEnqueuedAtOk() (*time.Time, bool)`

GetEnqueuedAtOk returns a tuple with the EnqueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedAt

`func (o *BackgroundTaskRequest) SetEnqueuedAt(v time.Time)`

SetEnqueuedAt sets EnqueuedAt field to given value.


### GetStartedAt

`func (o *BackgroundTaskRequest) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BackgroundTaskRequest) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BackgroundTaskRequest) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetEndedAt

`func (o *BackgroundTaskRequest) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *BackgroundTaskRequest) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *BackgroundTaskRequest) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.


### GetWorkerName

`func (o *BackgroundTaskRequest) GetWorkerName() string`

GetWorkerName returns the WorkerName field if non-nil, zero value otherwise.

### GetWorkerNameOk

`func (o *BackgroundTaskRequest) GetWorkerNameOk() (*string, bool)`

GetWorkerNameOk returns a tuple with the WorkerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerName

`func (o *BackgroundTaskRequest) SetWorkerName(v string)`

SetWorkerName sets WorkerName field to given value.


### GetMeta

`func (o *BackgroundTaskRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BackgroundTaskRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BackgroundTaskRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### GetLastHeartbeat

`func (o *BackgroundTaskRequest) GetLastHeartbeat() string`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *BackgroundTaskRequest) GetLastHeartbeatOk() (*string, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *BackgroundTaskRequest) SetLastHeartbeat(v string)`

SetLastHeartbeat sets LastHeartbeat field to given value.


### GetIsFinished

`func (o *BackgroundTaskRequest) GetIsFinished() bool`

GetIsFinished returns the IsFinished field if non-nil, zero value otherwise.

### GetIsFinishedOk

`func (o *BackgroundTaskRequest) GetIsFinishedOk() (*bool, bool)`

GetIsFinishedOk returns a tuple with the IsFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinished

`func (o *BackgroundTaskRequest) SetIsFinished(v bool)`

SetIsFinished sets IsFinished field to given value.


### GetIsQueued

`func (o *BackgroundTaskRequest) GetIsQueued() bool`

GetIsQueued returns the IsQueued field if non-nil, zero value otherwise.

### GetIsQueuedOk

`func (o *BackgroundTaskRequest) GetIsQueuedOk() (*bool, bool)`

GetIsQueuedOk returns a tuple with the IsQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueued

`func (o *BackgroundTaskRequest) SetIsQueued(v bool)`

SetIsQueued sets IsQueued field to given value.


### GetIsFailed

`func (o *BackgroundTaskRequest) GetIsFailed() bool`

GetIsFailed returns the IsFailed field if non-nil, zero value otherwise.

### GetIsFailedOk

`func (o *BackgroundTaskRequest) GetIsFailedOk() (*bool, bool)`

GetIsFailedOk returns a tuple with the IsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFailed

`func (o *BackgroundTaskRequest) SetIsFailed(v bool)`

SetIsFailed sets IsFailed field to given value.


### GetIsStarted

`func (o *BackgroundTaskRequest) GetIsStarted() bool`

GetIsStarted returns the IsStarted field if non-nil, zero value otherwise.

### GetIsStartedOk

`func (o *BackgroundTaskRequest) GetIsStartedOk() (*bool, bool)`

GetIsStartedOk returns a tuple with the IsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStarted

`func (o *BackgroundTaskRequest) SetIsStarted(v bool)`

SetIsStarted sets IsStarted field to given value.


### GetIsDeferred

`func (o *BackgroundTaskRequest) GetIsDeferred() bool`

GetIsDeferred returns the IsDeferred field if non-nil, zero value otherwise.

### GetIsDeferredOk

`func (o *BackgroundTaskRequest) GetIsDeferredOk() (*bool, bool)`

GetIsDeferredOk returns a tuple with the IsDeferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeferred

`func (o *BackgroundTaskRequest) SetIsDeferred(v bool)`

SetIsDeferred sets IsDeferred field to given value.


### GetIsCanceled

`func (o *BackgroundTaskRequest) GetIsCanceled() bool`

GetIsCanceled returns the IsCanceled field if non-nil, zero value otherwise.

### GetIsCanceledOk

`func (o *BackgroundTaskRequest) GetIsCanceledOk() (*bool, bool)`

GetIsCanceledOk returns a tuple with the IsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCanceled

`func (o *BackgroundTaskRequest) SetIsCanceled(v bool)`

SetIsCanceled sets IsCanceled field to given value.


### GetIsScheduled

`func (o *BackgroundTaskRequest) GetIsScheduled() bool`

GetIsScheduled returns the IsScheduled field if non-nil, zero value otherwise.

### GetIsScheduledOk

`func (o *BackgroundTaskRequest) GetIsScheduledOk() (*bool, bool)`

GetIsScheduledOk returns a tuple with the IsScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScheduled

`func (o *BackgroundTaskRequest) SetIsScheduled(v bool)`

SetIsScheduled sets IsScheduled field to given value.


### GetIsStopped

`func (o *BackgroundTaskRequest) GetIsStopped() bool`

GetIsStopped returns the IsStopped field if non-nil, zero value otherwise.

### GetIsStoppedOk

`func (o *BackgroundTaskRequest) GetIsStoppedOk() (*bool, bool)`

GetIsStoppedOk returns a tuple with the IsStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStopped

`func (o *BackgroundTaskRequest) SetIsStopped(v bool)`

SetIsStopped sets IsStopped field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


