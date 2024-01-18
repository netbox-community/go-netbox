# WritableCustomFieldChoiceSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**BaseChoices** | Pointer to [**PatchedWritableCustomFieldChoiceSetRequestBaseChoices**](PatchedWritableCustomFieldChoiceSetRequestBaseChoices.md) |  | [optional] 
**ExtraChoices** | Pointer to **[][]string** |  | [optional] 
**OrderAlphabetically** | Pointer to **bool** | Choices are automatically ordered alphabetically | [optional] 

## Methods

### NewWritableCustomFieldChoiceSetRequest

`func NewWritableCustomFieldChoiceSetRequest(name string, ) *WritableCustomFieldChoiceSetRequest`

NewWritableCustomFieldChoiceSetRequest instantiates a new WritableCustomFieldChoiceSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCustomFieldChoiceSetRequestWithDefaults

`func NewWritableCustomFieldChoiceSetRequestWithDefaults() *WritableCustomFieldChoiceSetRequest`

NewWritableCustomFieldChoiceSetRequestWithDefaults instantiates a new WritableCustomFieldChoiceSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableCustomFieldChoiceSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableCustomFieldChoiceSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableCustomFieldChoiceSetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WritableCustomFieldChoiceSetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableCustomFieldChoiceSetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableCustomFieldChoiceSetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableCustomFieldChoiceSetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseChoices

`func (o *WritableCustomFieldChoiceSetRequest) GetBaseChoices() PatchedWritableCustomFieldChoiceSetRequestBaseChoices`

GetBaseChoices returns the BaseChoices field if non-nil, zero value otherwise.

### GetBaseChoicesOk

`func (o *WritableCustomFieldChoiceSetRequest) GetBaseChoicesOk() (*PatchedWritableCustomFieldChoiceSetRequestBaseChoices, bool)`

GetBaseChoicesOk returns a tuple with the BaseChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseChoices

`func (o *WritableCustomFieldChoiceSetRequest) SetBaseChoices(v PatchedWritableCustomFieldChoiceSetRequestBaseChoices)`

SetBaseChoices sets BaseChoices field to given value.

### HasBaseChoices

`func (o *WritableCustomFieldChoiceSetRequest) HasBaseChoices() bool`

HasBaseChoices returns a boolean if a field has been set.

### GetExtraChoices

`func (o *WritableCustomFieldChoiceSetRequest) GetExtraChoices() [][]string`

GetExtraChoices returns the ExtraChoices field if non-nil, zero value otherwise.

### GetExtraChoicesOk

`func (o *WritableCustomFieldChoiceSetRequest) GetExtraChoicesOk() (*[][]string, bool)`

GetExtraChoicesOk returns a tuple with the ExtraChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraChoices

`func (o *WritableCustomFieldChoiceSetRequest) SetExtraChoices(v [][]string)`

SetExtraChoices sets ExtraChoices field to given value.

### HasExtraChoices

`func (o *WritableCustomFieldChoiceSetRequest) HasExtraChoices() bool`

HasExtraChoices returns a boolean if a field has been set.

### SetExtraChoicesNil

`func (o *WritableCustomFieldChoiceSetRequest) SetExtraChoicesNil(b bool)`

 SetExtraChoicesNil sets the value for ExtraChoices to be an explicit nil

### UnsetExtraChoices
`func (o *WritableCustomFieldChoiceSetRequest) UnsetExtraChoices()`

UnsetExtraChoices ensures that no value is present for ExtraChoices, not even an explicit nil
### GetOrderAlphabetically

`func (o *WritableCustomFieldChoiceSetRequest) GetOrderAlphabetically() bool`

GetOrderAlphabetically returns the OrderAlphabetically field if non-nil, zero value otherwise.

### GetOrderAlphabeticallyOk

`func (o *WritableCustomFieldChoiceSetRequest) GetOrderAlphabeticallyOk() (*bool, bool)`

GetOrderAlphabeticallyOk returns a tuple with the OrderAlphabetically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAlphabetically

`func (o *WritableCustomFieldChoiceSetRequest) SetOrderAlphabetically(v bool)`

SetOrderAlphabetically sets OrderAlphabetically field to given value.

### HasOrderAlphabetically

`func (o *WritableCustomFieldChoiceSetRequest) HasOrderAlphabetically() bool`

HasOrderAlphabetically returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


