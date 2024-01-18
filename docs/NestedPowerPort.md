# NestedPowerPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Device** | [**NestedDevice**](NestedDevice.md) |  | [readonly] 
**Name** | **string** |  | 
**Cable** | Pointer to **NullableInt32** |  | [optional] 
**Occupied** | **bool** |  | [readonly] 

## Methods

### NewNestedPowerPort

`func NewNestedPowerPort(id int32, url string, display string, device NestedDevice, name string, occupied bool, ) *NestedPowerPort`

NewNestedPowerPort instantiates a new NestedPowerPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedPowerPortWithDefaults

`func NewNestedPowerPortWithDefaults() *NestedPowerPort`

NewNestedPowerPortWithDefaults instantiates a new NestedPowerPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedPowerPort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedPowerPort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedPowerPort) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedPowerPort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedPowerPort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedPowerPort) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *NestedPowerPort) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedPowerPort) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedPowerPort) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDevice

`func (o *NestedPowerPort) GetDevice() NestedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *NestedPowerPort) GetDeviceOk() (*NestedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *NestedPowerPort) SetDevice(v NestedDevice)`

SetDevice sets Device field to given value.


### GetName

`func (o *NestedPowerPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedPowerPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedPowerPort) SetName(v string)`

SetName sets Name field to given value.


### GetCable

`func (o *NestedPowerPort) GetCable() int32`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *NestedPowerPort) GetCableOk() (*int32, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *NestedPowerPort) SetCable(v int32)`

SetCable sets Cable field to given value.

### HasCable

`func (o *NestedPowerPort) HasCable() bool`

HasCable returns a boolean if a field has been set.

### SetCableNil

`func (o *NestedPowerPort) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *NestedPowerPort) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetOccupied

`func (o *NestedPowerPort) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *NestedPowerPort) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *NestedPowerPort) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


