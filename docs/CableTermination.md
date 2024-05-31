# CableTermination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Cable** | **int32** |  | 
**CableEnd** | [**End1**](End1.md) |  | 
**TerminationType** | **string** |  | 
**TerminationId** | **int64** |  | 
**Termination** | **interface{}** |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewCableTermination

`func NewCableTermination(id int32, url string, display string, cable int32, cableEnd End1, terminationType string, terminationId int64, termination interface{}, created NullableTime, lastUpdated NullableTime, ) *CableTermination`

NewCableTermination instantiates a new CableTermination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCableTerminationWithDefaults

`func NewCableTerminationWithDefaults() *CableTermination`

NewCableTerminationWithDefaults instantiates a new CableTermination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CableTermination) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CableTermination) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CableTermination) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *CableTermination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CableTermination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CableTermination) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *CableTermination) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CableTermination) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CableTermination) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetCable

`func (o *CableTermination) GetCable() int32`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *CableTermination) GetCableOk() (*int32, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *CableTermination) SetCable(v int32)`

SetCable sets Cable field to given value.


### GetCableEnd

`func (o *CableTermination) GetCableEnd() End1`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *CableTermination) GetCableEndOk() (*End1, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *CableTermination) SetCableEnd(v End1)`

SetCableEnd sets CableEnd field to given value.


### GetTerminationType

`func (o *CableTermination) GetTerminationType() string`

GetTerminationType returns the TerminationType field if non-nil, zero value otherwise.

### GetTerminationTypeOk

`func (o *CableTermination) GetTerminationTypeOk() (*string, bool)`

GetTerminationTypeOk returns a tuple with the TerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationType

`func (o *CableTermination) SetTerminationType(v string)`

SetTerminationType sets TerminationType field to given value.


### GetTerminationId

`func (o *CableTermination) GetTerminationId() int64`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *CableTermination) GetTerminationIdOk() (*int64, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *CableTermination) SetTerminationId(v int64)`

SetTerminationId sets TerminationId field to given value.


### GetTermination

`func (o *CableTermination) GetTermination() interface{}`

GetTermination returns the Termination field if non-nil, zero value otherwise.

### GetTerminationOk

`func (o *CableTermination) GetTerminationOk() (*interface{}, bool)`

GetTerminationOk returns a tuple with the Termination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermination

`func (o *CableTermination) SetTermination(v interface{})`

SetTermination sets Termination field to given value.


### SetTerminationNil

`func (o *CableTermination) SetTerminationNil(b bool)`

 SetTerminationNil sets the value for Termination to be an explicit nil

### UnsetTermination
`func (o *CableTermination) UnsetTermination()`

UnsetTermination ensures that no value is present for Termination, not even an explicit nil
### GetCreated

`func (o *CableTermination) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CableTermination) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CableTermination) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *CableTermination) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CableTermination) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *CableTermination) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CableTermination) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CableTermination) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *CableTermination) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CableTermination) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


