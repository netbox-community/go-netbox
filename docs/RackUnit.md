# RackUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** |  | [readonly] 
**Name** | **string** |  | [readonly] 
**Face** | [**RackUnitFace**](RackUnitFace.md) |  | 
**Device** | [**BriefDevice**](BriefDevice.md) |  | [readonly] 
**Occupied** | **bool** |  | [readonly] 
**Display** | **string** |  | [readonly] 

## Methods

### NewRackUnit

`func NewRackUnit(id float64, name string, face RackUnitFace, device BriefDevice, occupied bool, display string, ) *RackUnit`

NewRackUnit instantiates a new RackUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackUnitWithDefaults

`func NewRackUnitWithDefaults() *RackUnit`

NewRackUnitWithDefaults instantiates a new RackUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RackUnit) GetId() float64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RackUnit) GetIdOk() (*float64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RackUnit) SetId(v float64)`

SetId sets Id field to given value.


### GetName

`func (o *RackUnit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RackUnit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RackUnit) SetName(v string)`

SetName sets Name field to given value.


### GetFace

`func (o *RackUnit) GetFace() RackUnitFace`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *RackUnit) GetFaceOk() (*RackUnitFace, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *RackUnit) SetFace(v RackUnitFace)`

SetFace sets Face field to given value.


### GetDevice

`func (o *RackUnit) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *RackUnit) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *RackUnit) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.


### GetOccupied

`func (o *RackUnit) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *RackUnit) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *RackUnit) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.


### GetDisplay

`func (o *RackUnit) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RackUnit) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RackUnit) SetDisplay(v string)`

SetDisplay sets Display field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


