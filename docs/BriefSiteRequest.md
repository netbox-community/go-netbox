# BriefSiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Full name of the site | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefSiteRequest

`func NewBriefSiteRequest(name string, slug string, ) *BriefSiteRequest`

NewBriefSiteRequest instantiates a new BriefSiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefSiteRequestWithDefaults

`func NewBriefSiteRequestWithDefaults() *BriefSiteRequest`

NewBriefSiteRequestWithDefaults instantiates a new BriefSiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BriefSiteRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefSiteRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefSiteRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefSiteRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefSiteRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefSiteRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefSiteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefSiteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefSiteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefSiteRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


