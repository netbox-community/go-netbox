# DeviceBayTemplateRequestDeviceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manufacturer** | [**BriefDeviceTypeRequestManufacturer**](BriefDeviceTypeRequestManufacturer.md) |  | 
**Model** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewDeviceBayTemplateRequestDeviceType

`func NewDeviceBayTemplateRequestDeviceType(manufacturer BriefDeviceTypeRequestManufacturer, model string, slug string, ) *DeviceBayTemplateRequestDeviceType`

NewDeviceBayTemplateRequestDeviceType instantiates a new DeviceBayTemplateRequestDeviceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceBayTemplateRequestDeviceTypeWithDefaults

`func NewDeviceBayTemplateRequestDeviceTypeWithDefaults() *DeviceBayTemplateRequestDeviceType`

NewDeviceBayTemplateRequestDeviceTypeWithDefaults instantiates a new DeviceBayTemplateRequestDeviceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManufacturer

`func (o *DeviceBayTemplateRequestDeviceType) GetManufacturer() BriefDeviceTypeRequestManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *DeviceBayTemplateRequestDeviceType) GetManufacturerOk() (*BriefDeviceTypeRequestManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *DeviceBayTemplateRequestDeviceType) SetManufacturer(v BriefDeviceTypeRequestManufacturer)`

SetManufacturer sets Manufacturer field to given value.


### GetModel

`func (o *DeviceBayTemplateRequestDeviceType) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceBayTemplateRequestDeviceType) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceBayTemplateRequestDeviceType) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *DeviceBayTemplateRequestDeviceType) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DeviceBayTemplateRequestDeviceType) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DeviceBayTemplateRequestDeviceType) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *DeviceBayTemplateRequestDeviceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceBayTemplateRequestDeviceType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceBayTemplateRequestDeviceType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceBayTemplateRequestDeviceType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


