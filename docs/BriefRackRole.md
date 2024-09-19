# BriefRackRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**RackCount** | **int64** |  | [readonly] 

## Methods

### NewBriefRackRole

`func NewBriefRackRole(id int32, url string, display string, name string, slug string, rackCount int64, ) *BriefRackRole`

NewBriefRackRole instantiates a new BriefRackRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefRackRoleWithDefaults

`func NewBriefRackRoleWithDefaults() *BriefRackRole`

NewBriefRackRoleWithDefaults instantiates a new BriefRackRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefRackRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefRackRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefRackRole) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefRackRole) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefRackRole) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefRackRole) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefRackRole) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefRackRole) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefRackRole) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefRackRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefRackRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefRackRole) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefRackRole) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefRackRole) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefRackRole) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefRackRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefRackRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefRackRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefRackRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRackCount

`func (o *BriefRackRole) GetRackCount() int64`

GetRackCount returns the RackCount field if non-nil, zero value otherwise.

### GetRackCountOk

`func (o *BriefRackRole) GetRackCountOk() (*int64, bool)`

GetRackCountOk returns a tuple with the RackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackCount

`func (o *BriefRackRole) SetRackCount(v int64)`

SetRackCount sets RackCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


