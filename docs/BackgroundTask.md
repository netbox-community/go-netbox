# BackgroundTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | [readonly] 
**Description** | **string** |  | 
**Origin** | **string** |  | 
**FuncName** | **string** |  | 
**Args** | **[]interface{}** |  | [readonly] 
**Kwargs** | **map[string]interface{}** |  | [readonly] 
**Result** | **string** |  | 
**Timeout** | **int32** |  | 
**ResultTtl** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**EnqueuedAt** | **time.Time** |  | 
**StartedAt** | **time.Time** |  | 
**EndedAt** | **time.Time** |  | 
**WorkerName** | **string** |  | 
**Position** | **int32** |  | [readonly] 
**Status** | **string** |  | [readonly] 
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

### NewBackgroundTask

`func NewBackgroundTask(id string, url string, description string, origin string, funcName string, args []interface{}, kwargs map[string]interface{}, result string, timeout int32, resultTtl int32, createdAt time.Time, enqueuedAt time.Time, startedAt time.Time, endedAt time.Time, workerName string, position int32, status string, meta map[string]interface{}, lastHeartbeat string, isFinished bool, isQueued bool, isFailed bool, isStarted bool, isDeferred bool, isCanceled bool, isScheduled bool, isStopped bool, ) *BackgroundTask`

NewBackgroundTask instantiates a new BackgroundTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundTaskWithDefaults

`func NewBackgroundTaskWithDefaults() *BackgroundTask`

NewBackgroundTaskWithDefaults instantiates a new BackgroundTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackgroundTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackgroundTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackgroundTask) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *BackgroundTask) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BackgroundTask) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BackgroundTask) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *BackgroundTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BackgroundTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BackgroundTask) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOrigin

`func (o *BackgroundTask) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *BackgroundTask) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *BackgroundTask) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetFuncName

`func (o *BackgroundTask) GetFuncName() string`

GetFuncName returns the FuncName field if non-nil, zero value otherwise.

### GetFuncNameOk

`func (o *BackgroundTask) GetFuncNameOk() (*string, bool)`

GetFuncNameOk returns a tuple with the FuncName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncName

`func (o *BackgroundTask) SetFuncName(v string)`

SetFuncName sets FuncName field to given value.


### GetArgs

`func (o *BackgroundTask) GetArgs() []interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *BackgroundTask) GetArgsOk() (*[]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *BackgroundTask) SetArgs(v []interface{})`

SetArgs sets Args field to given value.


### GetKwargs

`func (o *BackgroundTask) GetKwargs() map[string]interface{}`

GetKwargs returns the Kwargs field if non-nil, zero value otherwise.

### GetKwargsOk

`func (o *BackgroundTask) GetKwargsOk() (*map[string]interface{}, bool)`

GetKwargsOk returns a tuple with the Kwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKwargs

`func (o *BackgroundTask) SetKwargs(v map[string]interface{})`

SetKwargs sets Kwargs field to given value.


### GetResult

`func (o *BackgroundTask) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BackgroundTask) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BackgroundTask) SetResult(v string)`

SetResult sets Result field to given value.


### GetTimeout

`func (o *BackgroundTask) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BackgroundTask) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BackgroundTask) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetResultTtl

`func (o *BackgroundTask) GetResultTtl() int32`

GetResultTtl returns the ResultTtl field if non-nil, zero value otherwise.

### GetResultTtlOk

`func (o *BackgroundTask) GetResultTtlOk() (*int32, bool)`

GetResultTtlOk returns a tuple with the ResultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultTtl

`func (o *BackgroundTask) SetResultTtl(v int32)`

SetResultTtl sets ResultTtl field to given value.


### GetCreatedAt

`func (o *BackgroundTask) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackgroundTask) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackgroundTask) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnqueuedAt

`func (o *BackgroundTask) GetEnqueuedAt() time.Time`

GetEnqueuedAt returns the EnqueuedAt field if non-nil, zero value otherwise.

### GetEnqueuedAtOk

`func (o *BackgroundTask) GetEnqueuedAtOk() (*time.Time, bool)`

GetEnqueuedAtOk returns a tuple with the EnqueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedAt

`func (o *BackgroundTask) SetEnqueuedAt(v time.Time)`

SetEnqueuedAt sets EnqueuedAt field to given value.


### GetStartedAt

`func (o *BackgroundTask) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BackgroundTask) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BackgroundTask) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetEndedAt

`func (o *BackgroundTask) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *BackgroundTask) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *BackgroundTask) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.


### GetWorkerName

`func (o *BackgroundTask) GetWorkerName() string`

GetWorkerName returns the WorkerName field if non-nil, zero value otherwise.

### GetWorkerNameOk

`func (o *BackgroundTask) GetWorkerNameOk() (*string, bool)`

GetWorkerNameOk returns a tuple with the WorkerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerName

`func (o *BackgroundTask) SetWorkerName(v string)`

SetWorkerName sets WorkerName field to given value.


### GetPosition

`func (o *BackgroundTask) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BackgroundTask) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BackgroundTask) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetStatus

`func (o *BackgroundTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackgroundTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackgroundTask) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMeta

`func (o *BackgroundTask) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BackgroundTask) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BackgroundTask) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.


### GetLastHeartbeat

`func (o *BackgroundTask) GetLastHeartbeat() string`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *BackgroundTask) GetLastHeartbeatOk() (*string, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *BackgroundTask) SetLastHeartbeat(v string)`

SetLastHeartbeat sets LastHeartbeat field to given value.


### GetIsFinished

`func (o *BackgroundTask) GetIsFinished() bool`

GetIsFinished returns the IsFinished field if non-nil, zero value otherwise.

### GetIsFinishedOk

`func (o *BackgroundTask) GetIsFinishedOk() (*bool, bool)`

GetIsFinishedOk returns a tuple with the IsFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinished

`func (o *BackgroundTask) SetIsFinished(v bool)`

SetIsFinished sets IsFinished field to given value.


### GetIsQueued

`func (o *BackgroundTask) GetIsQueued() bool`

GetIsQueued returns the IsQueued field if non-nil, zero value otherwise.

### GetIsQueuedOk

`func (o *BackgroundTask) GetIsQueuedOk() (*bool, bool)`

GetIsQueuedOk returns a tuple with the IsQueued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueued

`func (o *BackgroundTask) SetIsQueued(v bool)`

SetIsQueued sets IsQueued field to given value.


### GetIsFailed

`func (o *BackgroundTask) GetIsFailed() bool`

GetIsFailed returns the IsFailed field if non-nil, zero value otherwise.

### GetIsFailedOk

`func (o *BackgroundTask) GetIsFailedOk() (*bool, bool)`

GetIsFailedOk returns a tuple with the IsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFailed

`func (o *BackgroundTask) SetIsFailed(v bool)`

SetIsFailed sets IsFailed field to given value.


### GetIsStarted

`func (o *BackgroundTask) GetIsStarted() bool`

GetIsStarted returns the IsStarted field if non-nil, zero value otherwise.

### GetIsStartedOk

`func (o *BackgroundTask) GetIsStartedOk() (*bool, bool)`

GetIsStartedOk returns a tuple with the IsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStarted

`func (o *BackgroundTask) SetIsStarted(v bool)`

SetIsStarted sets IsStarted field to given value.


### GetIsDeferred

`func (o *BackgroundTask) GetIsDeferred() bool`

GetIsDeferred returns the IsDeferred field if non-nil, zero value otherwise.

### GetIsDeferredOk

`func (o *BackgroundTask) GetIsDeferredOk() (*bool, bool)`

GetIsDeferredOk returns a tuple with the IsDeferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeferred

`func (o *BackgroundTask) SetIsDeferred(v bool)`

SetIsDeferred sets IsDeferred field to given value.


### GetIsCanceled

`func (o *BackgroundTask) GetIsCanceled() bool`

GetIsCanceled returns the IsCanceled field if non-nil, zero value otherwise.

### GetIsCanceledOk

`func (o *BackgroundTask) GetIsCanceledOk() (*bool, bool)`

GetIsCanceledOk returns a tuple with the IsCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCanceled

`func (o *BackgroundTask) SetIsCanceled(v bool)`

SetIsCanceled sets IsCanceled field to given value.


### GetIsScheduled

`func (o *BackgroundTask) GetIsScheduled() bool`

GetIsScheduled returns the IsScheduled field if non-nil, zero value otherwise.

### GetIsScheduledOk

`func (o *BackgroundTask) GetIsScheduledOk() (*bool, bool)`

GetIsScheduledOk returns a tuple with the IsScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScheduled

`func (o *BackgroundTask) SetIsScheduled(v bool)`

SetIsScheduled sets IsScheduled field to given value.


### GetIsStopped

`func (o *BackgroundTask) GetIsStopped() bool`

GetIsStopped returns the IsStopped field if non-nil, zero value otherwise.

### GetIsStoppedOk

`func (o *BackgroundTask) GetIsStoppedOk() (*bool, bool)`

GetIsStoppedOk returns a tuple with the IsStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStopped

`func (o *BackgroundTask) SetIsStopped(v bool)`

SetIsStopped sets IsStopped field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


