# Manufacturer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DevicetypeCount** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewManufacturer

`func NewManufacturer(id int32, url string, display string, name string, slug string, ) *Manufacturer`

NewManufacturer instantiates a new Manufacturer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManufacturerWithDefaults

`func NewManufacturerWithDefaults() *Manufacturer`

NewManufacturerWithDefaults instantiates a new Manufacturer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Manufacturer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Manufacturer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Manufacturer) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Manufacturer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Manufacturer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Manufacturer) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *Manufacturer) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Manufacturer) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Manufacturer) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *Manufacturer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Manufacturer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Manufacturer) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Manufacturer) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Manufacturer) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Manufacturer) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *Manufacturer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Manufacturer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Manufacturer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Manufacturer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevicetypeCount

`func (o *Manufacturer) GetDevicetypeCount() int64`

GetDevicetypeCount returns the DevicetypeCount field if non-nil, zero value otherwise.

### GetDevicetypeCountOk

`func (o *Manufacturer) GetDevicetypeCountOk() (*int64, bool)`

GetDevicetypeCountOk returns a tuple with the DevicetypeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicetypeCount

`func (o *Manufacturer) SetDevicetypeCount(v int64)`

SetDevicetypeCount sets DevicetypeCount field to given value.

### HasDevicetypeCount

`func (o *Manufacturer) HasDevicetypeCount() bool`

HasDevicetypeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


