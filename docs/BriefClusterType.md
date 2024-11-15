# BriefClusterType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ClusterCount** | **int64** |  | [readonly] 

## Methods

### NewBriefClusterType

`func NewBriefClusterType(id int32, url string, display string, name string, slug string, clusterCount int64, ) *BriefClusterType`

NewBriefClusterType instantiates a new BriefClusterType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefClusterTypeWithDefaults

`func NewBriefClusterTypeWithDefaults() *BriefClusterType`

NewBriefClusterTypeWithDefaults instantiates a new BriefClusterType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefClusterType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefClusterType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefClusterType) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefClusterType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefClusterType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefClusterType) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefClusterType) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefClusterType) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefClusterType) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefClusterType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefClusterType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefClusterType) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefClusterType) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefClusterType) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefClusterType) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefClusterType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefClusterType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefClusterType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefClusterType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetClusterCount

`func (o *BriefClusterType) GetClusterCount() int64`

GetClusterCount returns the ClusterCount field if non-nil, zero value otherwise.

### GetClusterCountOk

`func (o *BriefClusterType) GetClusterCountOk() (*int64, bool)`

GetClusterCountOk returns a tuple with the ClusterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCount

`func (o *BriefClusterType) SetClusterCount(v int64)`

SetClusterCount sets ClusterCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


