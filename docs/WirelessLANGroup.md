# WirelessLANGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**WirelesslanCount** | **int32** |  | [readonly] [default to 0]
**Depth** | **int32** |  | [readonly] 

## Methods

### NewWirelessLANGroup

`func NewWirelessLANGroup(id int32, url string, display string, name string, slug string, wirelesslanCount int32, depth int32, ) *WirelessLANGroup`

NewWirelessLANGroup instantiates a new WirelessLANGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessLANGroupWithDefaults

`func NewWirelessLANGroupWithDefaults() *WirelessLANGroup`

NewWirelessLANGroupWithDefaults instantiates a new WirelessLANGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WirelessLANGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WirelessLANGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WirelessLANGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *WirelessLANGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WirelessLANGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WirelessLANGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *WirelessLANGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WirelessLANGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WirelessLANGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *WirelessLANGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WirelessLANGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WirelessLANGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WirelessLANGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WirelessLANGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WirelessLANGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *WirelessLANGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WirelessLANGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WirelessLANGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WirelessLANGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWirelesslanCount

`func (o *WirelessLANGroup) GetWirelesslanCount() int32`

GetWirelesslanCount returns the WirelesslanCount field if non-nil, zero value otherwise.

### GetWirelesslanCountOk

`func (o *WirelessLANGroup) GetWirelesslanCountOk() (*int32, bool)`

GetWirelesslanCountOk returns a tuple with the WirelesslanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelesslanCount

`func (o *WirelessLANGroup) SetWirelesslanCount(v int32)`

SetWirelesslanCount sets WirelesslanCount field to given value.


### GetDepth

`func (o *WirelessLANGroup) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *WirelessLANGroup) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *WirelessLANGroup) SetDepth(v int32)`

SetDepth sets Depth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


