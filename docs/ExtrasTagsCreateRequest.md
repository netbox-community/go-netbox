# ExtrasTagsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**ObjectTypes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewExtrasTagsCreateRequest

`func NewExtrasTagsCreateRequest(name string, slug string, ) *ExtrasTagsCreateRequest`

NewExtrasTagsCreateRequest instantiates a new ExtrasTagsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtrasTagsCreateRequestWithDefaults

`func NewExtrasTagsCreateRequestWithDefaults() *ExtrasTagsCreateRequest`

NewExtrasTagsCreateRequestWithDefaults instantiates a new ExtrasTagsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtrasTagsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtrasTagsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtrasTagsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *ExtrasTagsCreateRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ExtrasTagsCreateRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ExtrasTagsCreateRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetColor

`func (o *ExtrasTagsCreateRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ExtrasTagsCreateRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ExtrasTagsCreateRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ExtrasTagsCreateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *ExtrasTagsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtrasTagsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtrasTagsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtrasTagsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWeight

`func (o *ExtrasTagsCreateRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ExtrasTagsCreateRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ExtrasTagsCreateRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ExtrasTagsCreateRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetObjectTypes

`func (o *ExtrasTagsCreateRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *ExtrasTagsCreateRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *ExtrasTagsCreateRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *ExtrasTagsCreateRequest) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


