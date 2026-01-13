# BriefMACAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**MacAddress** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefMACAddress

`func NewBriefMACAddress(id int32, url string, display string, macAddress string, ) *BriefMACAddress`

NewBriefMACAddress instantiates a new BriefMACAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefMACAddressWithDefaults

`func NewBriefMACAddressWithDefaults() *BriefMACAddress`

NewBriefMACAddressWithDefaults instantiates a new BriefMACAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefMACAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefMACAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefMACAddress) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefMACAddress) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefMACAddress) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefMACAddress) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefMACAddress) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefMACAddress) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefMACAddress) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetMacAddress

`func (o *BriefMACAddress) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BriefMACAddress) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BriefMACAddress) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetDescription

`func (o *BriefMACAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefMACAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefMACAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefMACAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


