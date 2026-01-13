# ExtrasCustomFieldChoiceSetsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**BaseChoices** | Pointer to [**NullablePatchedWritableCustomFieldChoiceSetRequestBaseChoices**](PatchedWritableCustomFieldChoiceSetRequestBaseChoices.md) |  | [optional] 
**ExtraChoices** | **[][]map[string]interface{}** |  | 
**OrderAlphabetically** | Pointer to **bool** | Choices are automatically ordered alphabetically | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 

## Methods

### NewExtrasCustomFieldChoiceSetsCreateRequest

`func NewExtrasCustomFieldChoiceSetsCreateRequest(name string, extraChoices [][]map[string]interface{}, ) *ExtrasCustomFieldChoiceSetsCreateRequest`

NewExtrasCustomFieldChoiceSetsCreateRequest instantiates a new ExtrasCustomFieldChoiceSetsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtrasCustomFieldChoiceSetsCreateRequestWithDefaults

`func NewExtrasCustomFieldChoiceSetsCreateRequestWithDefaults() *ExtrasCustomFieldChoiceSetsCreateRequest`

NewExtrasCustomFieldChoiceSetsCreateRequestWithDefaults instantiates a new ExtrasCustomFieldChoiceSetsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseChoices

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetBaseChoices() PatchedWritableCustomFieldChoiceSetRequestBaseChoices`

GetBaseChoices returns the BaseChoices field if non-nil, zero value otherwise.

### GetBaseChoicesOk

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetBaseChoicesOk() (*PatchedWritableCustomFieldChoiceSetRequestBaseChoices, bool)`

GetBaseChoicesOk returns a tuple with the BaseChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseChoices

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) SetBaseChoices(v PatchedWritableCustomFieldChoiceSetRequestBaseChoices)`

SetBaseChoices sets BaseChoices field to given value.

### HasBaseChoices

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) HasBaseChoices() bool`

HasBaseChoices returns a boolean if a field has been set.

### SetBaseChoicesNil

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) SetBaseChoicesNil(b bool)`

 SetBaseChoicesNil sets the value for BaseChoices to be an explicit nil

### UnsetBaseChoices
`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) UnsetBaseChoices()`

UnsetBaseChoices ensures that no value is present for BaseChoices, not even an explicit nil
### GetExtraChoices

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetExtraChoices() [][]map[string]interface{}`

GetExtraChoices returns the ExtraChoices field if non-nil, zero value otherwise.

### GetExtraChoicesOk

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetExtraChoicesOk() (*[][]map[string]interface{}, bool)`

GetExtraChoicesOk returns a tuple with the ExtraChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraChoices

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) SetExtraChoices(v [][]map[string]interface{})`

SetExtraChoices sets ExtraChoices field to given value.


### GetOrderAlphabetically

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetOrderAlphabetically() bool`

GetOrderAlphabetically returns the OrderAlphabetically field if non-nil, zero value otherwise.

### GetOrderAlphabeticallyOk

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetOrderAlphabeticallyOk() (*bool, bool)`

GetOrderAlphabeticallyOk returns a tuple with the OrderAlphabetically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAlphabetically

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) SetOrderAlphabetically(v bool)`

SetOrderAlphabetically sets OrderAlphabetically field to given value.

### HasOrderAlphabetically

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) HasOrderAlphabetically() bool`

HasOrderAlphabetically returns a boolean if a field has been set.

### GetOwner

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ExtrasCustomFieldChoiceSetsCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


