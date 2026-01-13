# DeviceWithConfigContextRequestVirtualChassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Master** | Pointer to [**NullableNestedDeviceRequest**](NestedDeviceRequest.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceWithConfigContextRequestVirtualChassis

`func NewDeviceWithConfigContextRequestVirtualChassis(name string, ) *DeviceWithConfigContextRequestVirtualChassis`

NewDeviceWithConfigContextRequestVirtualChassis instantiates a new DeviceWithConfigContextRequestVirtualChassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithConfigContextRequestVirtualChassisWithDefaults

`func NewDeviceWithConfigContextRequestVirtualChassisWithDefaults() *DeviceWithConfigContextRequestVirtualChassis`

NewDeviceWithConfigContextRequestVirtualChassisWithDefaults instantiates a new DeviceWithConfigContextRequestVirtualChassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceWithConfigContextRequestVirtualChassis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceWithConfigContextRequestVirtualChassis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceWithConfigContextRequestVirtualChassis) SetName(v string)`

SetName sets Name field to given value.


### GetMaster

`func (o *DeviceWithConfigContextRequestVirtualChassis) GetMaster() NestedDeviceRequest`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *DeviceWithConfigContextRequestVirtualChassis) GetMasterOk() (*NestedDeviceRequest, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *DeviceWithConfigContextRequestVirtualChassis) SetMaster(v NestedDeviceRequest)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *DeviceWithConfigContextRequestVirtualChassis) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### SetMasterNil

`func (o *DeviceWithConfigContextRequestVirtualChassis) SetMasterNil(b bool)`

 SetMasterNil sets the value for Master to be an explicit nil

### UnsetMaster
`func (o *DeviceWithConfigContextRequestVirtualChassis) UnsetMaster()`

UnsetMaster ensures that no value is present for Master, not even an explicit nil
### GetDescription

`func (o *DeviceWithConfigContextRequestVirtualChassis) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceWithConfigContextRequestVirtualChassis) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceWithConfigContextRequestVirtualChassis) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceWithConfigContextRequestVirtualChassis) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


