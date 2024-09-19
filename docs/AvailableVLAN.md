# AvailableVLAN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vid** | **int32** |  | [readonly] 
**Group** | [**NullableBriefVLANGroup**](BriefVLANGroup.md) |  | [readonly] 

## Methods

### NewAvailableVLAN

`func NewAvailableVLAN(vid int32, group NullableBriefVLANGroup, ) *AvailableVLAN`

NewAvailableVLAN instantiates a new AvailableVLAN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableVLANWithDefaults

`func NewAvailableVLANWithDefaults() *AvailableVLAN`

NewAvailableVLANWithDefaults instantiates a new AvailableVLAN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVid

`func (o *AvailableVLAN) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *AvailableVLAN) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *AvailableVLAN) SetVid(v int32)`

SetVid sets Vid field to given value.


### GetGroup

`func (o *AvailableVLAN) GetGroup() BriefVLANGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AvailableVLAN) GetGroupOk() (*BriefVLANGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AvailableVLAN) SetGroup(v BriefVLANGroup)`

SetGroup sets Group field to given value.


### SetGroupNil

`func (o *AvailableVLAN) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *AvailableVLAN) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


