# BriefProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** | Full name of the provider | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**CircuitCount** | **int64** |  | [readonly] 

## Methods

### NewBriefProvider

`func NewBriefProvider(id int32, url string, display string, name string, slug string, circuitCount int64, ) *BriefProvider`

NewBriefProvider instantiates a new BriefProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefProviderWithDefaults

`func NewBriefProviderWithDefaults() *BriefProvider`

NewBriefProviderWithDefaults instantiates a new BriefProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefProvider) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefProvider) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefProvider) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefProvider) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefProvider) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefProvider) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefProvider) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefProvider) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefProvider) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefProvider) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *BriefProvider) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefProvider) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefProvider) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCircuitCount

`func (o *BriefProvider) GetCircuitCount() int64`

GetCircuitCount returns the CircuitCount field if non-nil, zero value otherwise.

### GetCircuitCountOk

`func (o *BriefProvider) GetCircuitCountOk() (*int64, bool)`

GetCircuitCountOk returns a tuple with the CircuitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitCount

`func (o *BriefProvider) SetCircuitCount(v int64)`

SetCircuitCount sets CircuitCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


