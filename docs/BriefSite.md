# BriefSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** | Full name of the site | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefSite

`func NewBriefSite(id int32, url string, display string, name string, slug string, ) *BriefSite`

NewBriefSite instantiates a new BriefSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefSiteWithDefaults

`func NewBriefSiteWithDefaults() *BriefSite`

NewBriefSiteWithDefaults instantiates a new BriefSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefSite) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefSite) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefSite) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefSite) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefSite) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefSite) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefSite) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefSite) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefSite) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefSite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefSite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefSite) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefSite) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefSite) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefSite) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefSite) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefSite) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefSite) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefSite) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


