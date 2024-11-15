# BriefRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**PrefixCount** | **int64** |  | [readonly] 
**VlanCount** | **int64** |  | [readonly] 

## Methods

### NewBriefRole

`func NewBriefRole(id int32, url string, display string, name string, slug string, prefixCount int64, vlanCount int64, ) *BriefRole`

NewBriefRole instantiates a new BriefRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefRoleWithDefaults

`func NewBriefRoleWithDefaults() *BriefRole`

NewBriefRoleWithDefaults instantiates a new BriefRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefRole) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefRole) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefRole) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefRole) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefRole) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefRole) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefRole) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefRole) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefRole) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefRole) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefRole) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrefixCount

`func (o *BriefRole) GetPrefixCount() int64`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *BriefRole) GetPrefixCountOk() (*int64, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *BriefRole) SetPrefixCount(v int64)`

SetPrefixCount sets PrefixCount field to given value.


### GetVlanCount

`func (o *BriefRole) GetVlanCount() int64`

GetVlanCount returns the VlanCount field if non-nil, zero value otherwise.

### GetVlanCountOk

`func (o *BriefRole) GetVlanCountOk() (*int64, bool)`

GetVlanCountOk returns a tuple with the VlanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanCount

`func (o *BriefRole) SetVlanCount(v int64)`

SetVlanCount sets VlanCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


