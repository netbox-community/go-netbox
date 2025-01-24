# BriefModuleTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manufacturer** | [**BriefManufacturerRequest**](BriefManufacturerRequest.md) |  | 
**Model** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefModuleTypeRequest

`func NewBriefModuleTypeRequest(manufacturer BriefManufacturerRequest, model string, ) *BriefModuleTypeRequest`

NewBriefModuleTypeRequest instantiates a new BriefModuleTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefModuleTypeRequestWithDefaults

`func NewBriefModuleTypeRequestWithDefaults() *BriefModuleTypeRequest`

NewBriefModuleTypeRequestWithDefaults instantiates a new BriefModuleTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManufacturer

`func (o *BriefModuleTypeRequest) GetManufacturer() BriefManufacturerRequest`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BriefModuleTypeRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BriefModuleTypeRequest) SetManufacturer(v BriefManufacturerRequest)`

SetManufacturer sets Manufacturer field to given value.


### GetModel

`func (o *BriefModuleTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BriefModuleTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BriefModuleTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetDescription

`func (o *BriefModuleTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefModuleTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefModuleTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefModuleTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


