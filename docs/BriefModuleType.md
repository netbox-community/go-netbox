# BriefModuleType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Manufacturer** | [**BriefManufacturer**](BriefManufacturer.md) |  | 
**Model** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefModuleType

`func NewBriefModuleType(id int32, url string, display string, manufacturer BriefManufacturer, model string, ) *BriefModuleType`

NewBriefModuleType instantiates a new BriefModuleType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefModuleTypeWithDefaults

`func NewBriefModuleTypeWithDefaults() *BriefModuleType`

NewBriefModuleTypeWithDefaults instantiates a new BriefModuleType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BriefModuleType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BriefModuleType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BriefModuleType) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *BriefModuleType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BriefModuleType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BriefModuleType) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *BriefModuleType) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *BriefModuleType) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *BriefModuleType) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetManufacturer

`func (o *BriefModuleType) GetManufacturer() BriefManufacturer`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BriefModuleType) GetManufacturerOk() (*BriefManufacturer, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BriefModuleType) SetManufacturer(v BriefManufacturer)`

SetManufacturer sets Manufacturer field to given value.


### GetModel

`func (o *BriefModuleType) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BriefModuleType) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BriefModuleType) SetModel(v string)`

SetModel sets Model field to given value.


### GetDescription

`func (o *BriefModuleType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefModuleType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefModuleType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefModuleType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


