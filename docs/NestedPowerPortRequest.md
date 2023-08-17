# NestedPowerPortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Cable** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNestedPowerPortRequest

`func NewNestedPowerPortRequest(name string, ) *NestedPowerPortRequest`

NewNestedPowerPortRequest instantiates a new NestedPowerPortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedPowerPortRequestWithDefaults

`func NewNestedPowerPortRequestWithDefaults() *NestedPowerPortRequest`

NewNestedPowerPortRequestWithDefaults instantiates a new NestedPowerPortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NestedPowerPortRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedPowerPortRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedPowerPortRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCable

`func (o *NestedPowerPortRequest) GetCable() int32`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *NestedPowerPortRequest) GetCableOk() (*int32, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *NestedPowerPortRequest) SetCable(v int32)`

SetCable sets Cable field to given value.

### HasCable

`func (o *NestedPowerPortRequest) HasCable() bool`

HasCable returns a boolean if a field has been set.

### SetCableNil

`func (o *NestedPowerPortRequest) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *NestedPowerPortRequest) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


