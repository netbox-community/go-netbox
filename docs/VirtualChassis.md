# VirtualChassis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Master** | Pointer to [**NullableNestedDevice**](NestedDevice.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MemberCount** | **int32** |  | [readonly] 

## Methods

### NewVirtualChassis

`func NewVirtualChassis(id int32, url string, display string, name string, memberCount int32, ) *VirtualChassis`

NewVirtualChassis instantiates a new VirtualChassis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualChassisWithDefaults

`func NewVirtualChassisWithDefaults() *VirtualChassis`

NewVirtualChassisWithDefaults instantiates a new VirtualChassis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualChassis) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualChassis) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualChassis) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *VirtualChassis) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualChassis) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualChassis) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *VirtualChassis) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VirtualChassis) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VirtualChassis) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *VirtualChassis) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualChassis) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualChassis) SetName(v string)`

SetName sets Name field to given value.


### GetMaster

`func (o *VirtualChassis) GetMaster() NestedDevice`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *VirtualChassis) GetMasterOk() (*NestedDevice, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *VirtualChassis) SetMaster(v NestedDevice)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *VirtualChassis) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### SetMasterNil

`func (o *VirtualChassis) SetMasterNil(b bool)`

 SetMasterNil sets the value for Master to be an explicit nil

### UnsetMaster
`func (o *VirtualChassis) UnsetMaster()`

UnsetMaster ensures that no value is present for Master, not even an explicit nil
### GetDescription

`func (o *VirtualChassis) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualChassis) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualChassis) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualChassis) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMemberCount

`func (o *VirtualChassis) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *VirtualChassis) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *VirtualChassis) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


