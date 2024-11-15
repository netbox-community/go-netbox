# BriefVRF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Rd** | Pointer to **NullableString** | Unique route distinguisher (as defined in RFC 4364) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PrefixCount** | **int64** |  | [readonly] 

## Methods

### NewBriefVRF

`func NewBriefVRF(id int32, url string, display string, name string, prefixCount int64, ) *BriefVRF`

NewBriefVRF instantiates a new BriefVRF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefVRFWithDefaults

`func NewBriefVRFWithDefaults() *BriefVRF`

NewBriefVRFWithDefaults instantiates a new BriefVRF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefVRF) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefVRF) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefVRF) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefVRF) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefVRF) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefVRF) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefVRF) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefVRF) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefVRF) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *BriefVRF) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BriefVRF) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BriefVRF) SetName(v string)`

SetName sets Name field to given value.


### GetRd

`func (o *BriefVRF) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *BriefVRF) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *BriefVRF) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *BriefVRF) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *BriefVRF) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *BriefVRF) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetDescription

`func (o *BriefVRF) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefVRF) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefVRF) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefVRF) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrefixCount

`func (o *BriefVRF) GetPrefixCount() int64`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *BriefVRF) GetPrefixCountOk() (*int64, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *BriefVRF) SetPrefixCount(v int64)`

SetPrefixCount sets PrefixCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


