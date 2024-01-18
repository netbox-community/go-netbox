# NestedVirtualChassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Master** | [**NestedDevice**](NestedDevice.md) |  | 

## Methods

### NewNestedVirtualChassis

`func NewNestedVirtualChassis(id int32, url string, display string, name string, master NestedDevice, ) *NestedVirtualChassis`

NewNestedVirtualChassis instantiates a new NestedVirtualChassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedVirtualChassisWithDefaults

`func NewNestedVirtualChassisWithDefaults() *NestedVirtualChassis`

NewNestedVirtualChassisWithDefaults instantiates a new NestedVirtualChassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedVirtualChassis) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedVirtualChassis) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedVirtualChassis) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedVirtualChassis) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedVirtualChassis) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedVirtualChassis) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *NestedVirtualChassis) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedVirtualChassis) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedVirtualChassis) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *NestedVirtualChassis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedVirtualChassis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedVirtualChassis) SetName(v string)`

SetName sets Name field to given value.


### GetMaster

`func (o *NestedVirtualChassis) GetMaster() NestedDevice`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *NestedVirtualChassis) GetMasterOk() (*NestedDevice, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *NestedVirtualChassis) SetMaster(v NestedDevice)`

SetMaster sets Master field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


