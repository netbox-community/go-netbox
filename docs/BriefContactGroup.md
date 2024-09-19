# BriefContactGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ContactCount** | **int32** |  | [readonly] [default to 0]
**Depth** | **int32** |  | [readonly] 

## Methods

### NewBriefContactGroup

`func NewBriefContactGroup(id int32, url string, display string, name string, slug string, contactCount int32, depth int32, ) *BriefContactGroup`

NewBriefContactGroup instantiates a new BriefContactGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefContactGroupWithDefaults

`func NewBriefContactGroupWithDefaults() *BriefContactGroup`

NewBriefContactGroupWithDefaults instantiates a new BriefContactGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefContactGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefContactGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefContactGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefContactGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefContactGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefContactGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefContactGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefContactGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefContactGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefContactGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefContactGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefContactGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefContactGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefContactGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefContactGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefContactGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefContactGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefContactGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefContactGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContactCount

`func (o *BriefContactGroup) GetContactCount() int32`

GetContactCount returns the ContactCount field if non-nil, zero value otherwise.

### GetContactCountOk

`func (o *BriefContactGroup) GetContactCountOk() (*int32, bool)`

GetContactCountOk returns a tuple with the ContactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactCount

`func (o *BriefContactGroup) SetContactCount(v int32)`

SetContactCount sets ContactCount field to given value.


### GetDepth

`func (o *BriefContactGroup) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *BriefContactGroup) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *BriefContactGroup) SetDepth(v int32)`

SetDepth sets Depth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


