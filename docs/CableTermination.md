# CableTermination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Cable** | **int32** |  | [readonly] 
**CableEnd** | Pointer to [**End1**](End1.md) |  | [optional] 
**TerminationType** | **string** |  | [readonly] 
**TerminationId** | **int32** |  | [readonly] 
**Termination** | Pointer to **interface{}** |  | [optional] [readonly] 
**Connector** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**Positions** | Pointer to **[]int32** |  | [optional] [readonly] 
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewCableTermination

`func NewCableTermination(id int32, url string, display string, cable int32, terminationType string, terminationId int32, ) *CableTermination`

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

### HasCableEnd

`func (o *CableTermination) HasCableEnd() bool`

HasCableEnd returns a boolean if a field has been set.

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

`func (o *CableTermination) GetTerminationId() int32`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *CableTermination) GetTerminationIdOk() (*int32, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *CableTermination) SetTerminationId(v int32)`

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

### HasTermination

`func (o *CableTermination) HasTermination() bool`

HasTermination returns a boolean if a field has been set.

### SetTerminationNil

`func (o *CableTermination) SetTerminationNil(b bool)`

 SetTerminationNil sets the value for Termination to be an explicit nil

### UnsetTermination
`func (o *CableTermination) UnsetTermination()`

UnsetTermination ensures that no value is present for Termination, not even an explicit nil
### GetConnector

`func (o *CableTermination) GetConnector() int32`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *CableTermination) GetConnectorOk() (*int32, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *CableTermination) SetConnector(v int32)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *CableTermination) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### SetConnectorNil

`func (o *CableTermination) SetConnectorNil(b bool)`

 SetConnectorNil sets the value for Connector to be an explicit nil

### UnsetConnector
`func (o *CableTermination) UnsetConnector()`

UnsetConnector ensures that no value is present for Connector, not even an explicit nil
### GetPositions

`func (o *CableTermination) GetPositions() []int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *CableTermination) GetPositionsOk() (*[]int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *CableTermination) SetPositions(v []int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *CableTermination) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### SetPositionsNil

`func (o *CableTermination) SetPositionsNil(b bool)`

 SetPositionsNil sets the value for Positions to be an explicit nil

### UnsetPositions
`func (o *CableTermination) UnsetPositions()`

UnsetPositions ensures that no value is present for Positions, not even an explicit nil
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

### HasCreated

`func (o *CableTermination) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

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

### HasLastUpdated

`func (o *CableTermination) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *CableTermination) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CableTermination) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


