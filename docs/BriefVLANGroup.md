# BriefVLANGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**VlanCount** | **int64** |  | [readonly] 

## Methods

### NewBriefVLANGroup

`func NewBriefVLANGroup(id int32, url string, display string, name string, slug string, vlanCount int64, ) *BriefVLANGroup`

NewBriefVLANGroup instantiates a new BriefVLANGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefVLANGroupWithDefaults

`func NewBriefVLANGroupWithDefaults() *BriefVLANGroup`

NewBriefVLANGroupWithDefaults instantiates a new BriefVLANGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefVLANGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefVLANGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefVLANGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefVLANGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefVLANGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefVLANGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefVLANGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefVLANGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefVLANGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefVLANGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefVLANGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefVLANGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefVLANGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefVLANGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefVLANGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefVLANGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefVLANGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefVLANGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefVLANGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVlanCount

`func (o *BriefVLANGroup) GetVlanCount() int64`

GetVlanCount returns the VlanCount field if non-nil, zero value otherwise.

### GetVlanCountOk

`func (o *BriefVLANGroup) GetVlanCountOk() (*int64, bool)`

GetVlanCountOk returns a tuple with the VlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCount

`func (o *BriefVLANGroup) SetVlanCount(v int64)`

SetVlanCount sets VlanCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


