# VRF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Rd** | Pointer to **NullableString** | Unique route distinguisher (as defined in RFC 4364) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PrefixCount** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewVRF

`func NewVRF(id int32, url string, display string, name string, ) *VRF`

NewVRF instantiates a new VRF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVRFWithDefaults

`func NewVRFWithDefaults() *VRF`

NewVRFWithDefaults instantiates a new VRF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VRF) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VRF) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VRF) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *VRF) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VRF) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VRF) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *VRF) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VRF) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VRF) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *VRF) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VRF) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VRF) SetName(v string)`

SetName sets Name field to given value.


### GetRd

`func (o *VRF) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *VRF) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *VRF) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *VRF) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *VRF) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *VRF) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetDescription

`func (o *VRF) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VRF) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VRF) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VRF) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrefixCount

`func (o *VRF) GetPrefixCount() int64`

GetPrefixCount returns the PrefixCount field if non-nil, zero value otherwise.

### GetPrefixCountOk

`func (o *VRF) GetPrefixCountOk() (*int64, bool)`

GetPrefixCountOk returns a tuple with the PrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixCount

`func (o *VRF) SetPrefixCount(v int64)`

SetPrefixCount sets PrefixCount field to given value.

### HasPrefixCount

`func (o *VRF) HasPrefixCount() bool`

HasPrefixCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


