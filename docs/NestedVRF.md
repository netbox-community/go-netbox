# NestedVRF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Rd** | Pointer to **NullableString** | Unique route distinguisher (as defined in RFC 4364) | [optional] 

## Methods

### NewNestedVRF

`func NewNestedVRF(id int32, url string, display string, name string, ) *NestedVRF`

NewNestedVRF instantiates a new NestedVRF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedVRFWithDefaults

`func NewNestedVRFWithDefaults() *NestedVRF`

NewNestedVRFWithDefaults instantiates a new NestedVRF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NestedVRF) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NestedVRF) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NestedVRF) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *NestedVRF) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NestedVRF) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NestedVRF) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *NestedVRF) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NestedVRF) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NestedVRF) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *NestedVRF) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NestedVRF) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NestedVRF) SetName(v string)`

SetName sets Name field to given value.


### GetRd

`func (o *NestedVRF) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *NestedVRF) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *NestedVRF) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *NestedVRF) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *NestedVRF) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *NestedVRF) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


