# BriefRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**SiteCount** | **int32** |  | [readonly] [default to 0]
**Depth** | **int32** |  | [readonly] 

## Methods

### NewBriefRegion

`func NewBriefRegion(id int32, url string, display string, name string, slug string, siteCount int32, depth int32, ) *BriefRegion`

NewBriefRegion instantiates a new BriefRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefRegionWithDefaults

`func NewBriefRegionWithDefaults() *BriefRegion`

NewBriefRegionWithDefaults instantiates a new BriefRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefRegion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefRegion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefRegion) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefRegion) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefRegion) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefRegion) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefRegion) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefRegion) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefRegion) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefRegion) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefRegion) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefRegion) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefRegion) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefRegion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefRegion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefRegion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefRegion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSiteCount

`func (o *BriefRegion) GetSiteCount() int32`

GetSiteCount returns the SiteCount field if non-nil, zero value otherwise.

### GetSiteCountOk

`func (o *BriefRegion) GetSiteCountOk() (*int32, bool)`

GetSiteCountOk returns a tuple with the SiteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteCount

`func (o *BriefRegion) SetSiteCount(v int32)`

SetSiteCount sets SiteCount field to given value.


### GetDepth

`func (o *BriefRegion) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *BriefRegion) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *BriefRegion) SetDepth(v int32)`

SetDepth sets Depth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


