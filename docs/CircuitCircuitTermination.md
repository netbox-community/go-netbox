# CircuitCircuitTermination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | Pointer to **string** |  | [optional] [readonly] 
**Display** | **string** |  | [readonly] 
**TerminationType** | Pointer to **NullableString** |  | [optional] 
**TerminationId** | Pointer to **NullableInt32** |  | [optional] 
**Termination** | Pointer to **interface{}** |  | [optional] [readonly] 
**PortSpeed** | Pointer to **NullableInt32** | Physical circuit speed | [optional] 
**UpstreamSpeed** | Pointer to **NullableInt32** | Upstream speed, if different from port speed | [optional] 
**XconnectId** | Pointer to **string** | ID of the local cross-connect | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCircuitCircuitTermination

`func NewCircuitCircuitTermination(id int32, url string, display string, ) *CircuitCircuitTermination`

NewCircuitCircuitTermination instantiates a new CircuitCircuitTermination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitCircuitTerminationWithDefaults

`func NewCircuitCircuitTerminationWithDefaults() *CircuitCircuitTermination`

NewCircuitCircuitTerminationWithDefaults instantiates a new CircuitCircuitTermination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CircuitCircuitTermination) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CircuitCircuitTermination) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CircuitCircuitTermination) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *CircuitCircuitTermination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CircuitCircuitTermination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CircuitCircuitTermination) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *CircuitCircuitTermination) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *CircuitCircuitTermination) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *CircuitCircuitTermination) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.

### HasDisplayUrl

`func (o *CircuitCircuitTermination) HasDisplayUrl() bool`

HasDisplayUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *CircuitCircuitTermination) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CircuitCircuitTermination) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CircuitCircuitTermination) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetTerminationType

`func (o *CircuitCircuitTermination) GetTerminationType() string`

GetTerminationType returns the TerminationType field if non-nil, zero value otherwise.

### GetTerminationTypeOk

`func (o *CircuitCircuitTermination) GetTerminationTypeOk() (*string, bool)`

GetTerminationTypeOk returns a tuple with the TerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationType

`func (o *CircuitCircuitTermination) SetTerminationType(v string)`

SetTerminationType sets TerminationType field to given value.

### HasTerminationType

`func (o *CircuitCircuitTermination) HasTerminationType() bool`

HasTerminationType returns a boolean if a field has been set.

### SetTerminationTypeNil

`func (o *CircuitCircuitTermination) SetTerminationTypeNil(b bool)`

 SetTerminationTypeNil sets the value for TerminationType to be an explicit nil

### UnsetTerminationType
`func (o *CircuitCircuitTermination) UnsetTerminationType()`

UnsetTerminationType ensures that no value is present for TerminationType, not even an explicit nil
### GetTerminationId

`func (o *CircuitCircuitTermination) GetTerminationId() int32`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *CircuitCircuitTermination) GetTerminationIdOk() (*int32, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *CircuitCircuitTermination) SetTerminationId(v int32)`

SetTerminationId sets TerminationId field to given value.

### HasTerminationId

`func (o *CircuitCircuitTermination) HasTerminationId() bool`

HasTerminationId returns a boolean if a field has been set.

### SetTerminationIdNil

`func (o *CircuitCircuitTermination) SetTerminationIdNil(b bool)`

 SetTerminationIdNil sets the value for TerminationId to be an explicit nil

### UnsetTerminationId
`func (o *CircuitCircuitTermination) UnsetTerminationId()`

UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
### GetTermination

`func (o *CircuitCircuitTermination) GetTermination() interface{}`

GetTermination returns the Termination field if non-nil, zero value otherwise.

### GetTerminationOk

`func (o *CircuitCircuitTermination) GetTerminationOk() (*interface{}, bool)`

GetTerminationOk returns a tuple with the Termination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermination

`func (o *CircuitCircuitTermination) SetTermination(v interface{})`

SetTermination sets Termination field to given value.

### HasTermination

`func (o *CircuitCircuitTermination) HasTermination() bool`

HasTermination returns a boolean if a field has been set.

### SetTerminationNil

`func (o *CircuitCircuitTermination) SetTerminationNil(b bool)`

 SetTerminationNil sets the value for Termination to be an explicit nil

### UnsetTermination
`func (o *CircuitCircuitTermination) UnsetTermination()`

UnsetTermination ensures that no value is present for Termination, not even an explicit nil
### GetPortSpeed

`func (o *CircuitCircuitTermination) GetPortSpeed() int32`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *CircuitCircuitTermination) GetPortSpeedOk() (*int32, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *CircuitCircuitTermination) SetPortSpeed(v int32)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *CircuitCircuitTermination) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### SetPortSpeedNil

`func (o *CircuitCircuitTermination) SetPortSpeedNil(b bool)`

 SetPortSpeedNil sets the value for PortSpeed to be an explicit nil

### UnsetPortSpeed
`func (o *CircuitCircuitTermination) UnsetPortSpeed()`

UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
### GetUpstreamSpeed

`func (o *CircuitCircuitTermination) GetUpstreamSpeed() int32`

GetUpstreamSpeed returns the UpstreamSpeed field if non-nil, zero value otherwise.

### GetUpstreamSpeedOk

`func (o *CircuitCircuitTermination) GetUpstreamSpeedOk() (*int32, bool)`

GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamSpeed

`func (o *CircuitCircuitTermination) SetUpstreamSpeed(v int32)`

SetUpstreamSpeed sets UpstreamSpeed field to given value.

### HasUpstreamSpeed

`func (o *CircuitCircuitTermination) HasUpstreamSpeed() bool`

HasUpstreamSpeed returns a boolean if a field has been set.

### SetUpstreamSpeedNil

`func (o *CircuitCircuitTermination) SetUpstreamSpeedNil(b bool)`

 SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil

### UnsetUpstreamSpeed
`func (o *CircuitCircuitTermination) UnsetUpstreamSpeed()`

UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
### GetXconnectId

`func (o *CircuitCircuitTermination) GetXconnectId() string`

GetXconnectId returns the XconnectId field if non-nil, zero value otherwise.

### GetXconnectIdOk

`func (o *CircuitCircuitTermination) GetXconnectIdOk() (*string, bool)`

GetXconnectIdOk returns a tuple with the XconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXconnectId

`func (o *CircuitCircuitTermination) SetXconnectId(v string)`

SetXconnectId sets XconnectId field to given value.

### HasXconnectId

`func (o *CircuitCircuitTermination) HasXconnectId() bool`

HasXconnectId returns a boolean if a field has been set.

### GetDescription

`func (o *CircuitCircuitTermination) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CircuitCircuitTermination) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CircuitCircuitTermination) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CircuitCircuitTermination) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


