# BriefDeviceRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DeviceCount** | Pointer to **int64** |  | [optional] [readonly] 
**VirtualmachineCount** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewBriefDeviceRole

`func NewBriefDeviceRole(id int32, url string, display string, name string, slug string, ) *BriefDeviceRole`

NewBriefDeviceRole instantiates a new BriefDeviceRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefDeviceRoleWithDefaults

`func NewBriefDeviceRoleWithDefaults() *BriefDeviceRole`

NewBriefDeviceRoleWithDefaults instantiates a new BriefDeviceRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefDeviceRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefDeviceRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefDeviceRole) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefDeviceRole) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefDeviceRole) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefDeviceRole) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefDeviceRole) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefDeviceRole) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefDeviceRole) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefDeviceRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefDeviceRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefDeviceRole) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefDeviceRole) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefDeviceRole) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefDeviceRole) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefDeviceRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefDeviceRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefDeviceRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefDeviceRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceCount

`func (o *BriefDeviceRole) GetDeviceCount() int64`

GetDeviceCount returns the DeviceCount field if non-nil, zero value otherwise.

### GetDeviceCountOk

`func (o *BriefDeviceRole) GetDeviceCountOk() (*int64, bool)`

GetDeviceCountOk returns a tuple with the DeviceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCount

`func (o *BriefDeviceRole) SetDeviceCount(v int64)`

SetDeviceCount sets DeviceCount field to given value.

### HasDeviceCount

`func (o *BriefDeviceRole) HasDeviceCount() bool`

HasDeviceCount returns a boolean if a field has been set.

### GetVirtualmachineCount

`func (o *BriefDeviceRole) GetVirtualmachineCount() int64`

GetVirtualmachineCount returns the VirtualmachineCount field if non-nil, zero value otherwise.

### GetVirtualmachineCountOk

`func (o *BriefDeviceRole) GetVirtualmachineCountOk() (*int64, bool)`

GetVirtualmachineCountOk returns a tuple with the VirtualmachineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualmachineCount

`func (o *BriefDeviceRole) SetVirtualmachineCount(v int64)`

SetVirtualmachineCount sets VirtualmachineCount field to given value.

### HasVirtualmachineCount

`func (o *BriefDeviceRole) HasVirtualmachineCount() bool`

HasVirtualmachineCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


