# BriefRack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DeviceCount** | **int64** |  | [readonly] 

## Methods

### NewBriefRack

`func NewBriefRack(id int32, url string, display string, name string, deviceCount int64, ) *BriefRack`

NewBriefRack instantiates a new BriefRack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefRackWithDefaults

`func NewBriefRackWithDefaults() *BriefRack`

NewBriefRackWithDefaults instantiates a new BriefRack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefRack) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefRack) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefRack) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefRack) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefRack) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefRack) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefRack) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefRack) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefRack) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefRack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefRack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefRack) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BriefRack) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefRack) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefRack) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefRack) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceCount

`func (o *BriefRack) GetDeviceCount() int64`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *BriefRack) GetDeviceCountOk() (*int64, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *BriefRack) SetDeviceCount(v int64)`

SetDeviceCount sets DeviceCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


