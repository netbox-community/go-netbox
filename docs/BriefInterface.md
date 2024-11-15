# BriefInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Device** | [**BriefDevice**](BriefDevice.md) |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Cable** | [**NullableBriefCable**](BriefCable.md) |  | [readonly] 
**Occupied** | **bool** |  | [readonly] 

## Methods

### NewBriefInterface

`func NewBriefInterface(id int32, url string, display string, device BriefDevice, name string, cable NullableBriefCable, occupied bool, ) *BriefInterface`

NewBriefInterface instantiates a new BriefInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefInterfaceWithDefaults

`func NewBriefInterfaceWithDefaults() *BriefInterface`

NewBriefInterfaceWithDefaults instantiates a new BriefInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefInterface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefInterface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefInterface) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefInterface) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefInterface) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefInterface) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefInterface) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefInterface) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefInterface) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDevice

`func (o *BriefInterface) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *BriefInterface) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *BriefInterface) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.


### GetName

`func (o *BriefInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefInterface) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BriefInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCable

`func (o *BriefInterface) GetCable() BriefCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *BriefInterface) GetCableOk() (*BriefCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *BriefInterface) SetCable(v BriefCable)`

SetCable sets Cable field to given value.


### SetCableNil

`func (o *BriefInterface) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *BriefInterface) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetOccupied

`func (o *BriefInterface) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *BriefInterface) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *BriefInterface) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


