# BriefRackTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manufacturer** | [**BriefManufacturerRequest**](BriefManufacturerRequest.md) |  | 
**Model** | **string** |  | 
**Slug** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefRackTypeRequest

`func NewBriefRackTypeRequest(manufacturer BriefManufacturerRequest, model string, slug string, ) *BriefRackTypeRequest`

NewBriefRackTypeRequest instantiates a new BriefRackTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefRackTypeRequestWithDefaults

`func NewBriefRackTypeRequestWithDefaults() *BriefRackTypeRequest`

NewBriefRackTypeRequestWithDefaults instantiates a new BriefRackTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManufacturer

`func (o *BriefRackTypeRequest) GetManufacturer() BriefManufacturerRequest`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *BriefRackTypeRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *BriefRackTypeRequest) SetManufacturer(v BriefManufacturerRequest)`

SetManufacturer sets Manufacturer field to given value.


### GetModel

`func (o *BriefRackTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BriefRackTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BriefRackTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetSlug

`func (o *BriefRackTypeRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BriefRackTypeRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BriefRackTypeRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *BriefRackTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefRackTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefRackTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefRackTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


