# PatchedWritableCustomFieldChoiceSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BaseChoices** | Pointer to [**PatchedWritableCustomFieldChoiceSetRequestBaseChoices**](PatchedWritableCustomFieldChoiceSetRequestBaseChoices.md) |  | [optional] 
**ExtraChoices** | Pointer to **[][]interface{}** |  | [optional] 
**OrderAlphabetically** | Pointer to **bool** | Choices are automatically ordered alphabetically | [optional] 

## Methods

### NewPatchedWritableCustomFieldChoiceSetRequest

`func NewPatchedWritableCustomFieldChoiceSetRequest() *PatchedWritableCustomFieldChoiceSetRequest`

NewPatchedWritableCustomFieldChoiceSetRequest instantiates a new PatchedWritableCustomFieldChoiceSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableCustomFieldChoiceSetRequestWithDefaults

`func NewPatchedWritableCustomFieldChoiceSetRequestWithDefaults() *PatchedWritableCustomFieldChoiceSetRequest`

NewPatchedWritableCustomFieldChoiceSetRequestWithDefaults instantiates a new PatchedWritableCustomFieldChoiceSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableCustomFieldChoiceSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableCustomFieldChoiceSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableCustomFieldChoiceSetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableCustomFieldChoiceSetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableCustomFieldChoiceSetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableCustomFieldChoiceSetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableCustomFieldChoiceSetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableCustomFieldChoiceSetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseChoices

`func (o *PatchedWritableCustomFieldChoiceSetRequest) GetBaseChoices() PatchedWritableCustomFieldChoiceSetRequestBaseChoices`

GetBaseChoices returns the BaseChoices field if non-nil, zero value otherwise.

### GetBaseChoicesOk

`func (o *PatchedWritableCustomFieldChoiceSetRequest) GetBaseChoicesOk() (*PatchedWritableCustomFieldChoiceSetRequestBaseChoices, bool)`

GetBaseChoicesOk returns a tuple with the BaseChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseChoices

`func (o *PatchedWritableCustomFieldChoiceSetRequest) SetBaseChoices(v PatchedWritableCustomFieldChoiceSetRequestBaseChoices)`

SetBaseChoices sets BaseChoices field to given value.

### HasBaseChoices

`func (o *PatchedWritableCustomFieldChoiceSetRequest) HasBaseChoices() bool`

HasBaseChoices returns a boolean if a field has been set.

### GetExtraChoices

`func (o *PatchedWritableCustomFieldChoiceSetRequest) GetExtraChoices() [][]interface{}`

GetExtraChoices returns the ExtraChoices field if non-nil, zero value otherwise.

### GetExtraChoicesOk

`func (o *PatchedWritableCustomFieldChoiceSetRequest) GetExtraChoicesOk() (*[][]interface{}, bool)`

GetExtraChoicesOk returns a tuple with the ExtraChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraChoices

`func (o *PatchedWritableCustomFieldChoiceSetRequest) SetExtraChoices(v [][]interface{})`

SetExtraChoices sets ExtraChoices field to given value.

### HasExtraChoices

`func (o *PatchedWritableCustomFieldChoiceSetRequest) HasExtraChoices() bool`

HasExtraChoices returns a boolean if a field has been set.

### GetOrderAlphabetically

`func (o *PatchedWritableCustomFieldChoiceSetRequest) GetOrderAlphabetically() bool`

GetOrderAlphabetically returns the OrderAlphabetically field if non-nil, zero value otherwise.

### GetOrderAlphabeticallyOk

`func (o *PatchedWritableCustomFieldChoiceSetRequest) GetOrderAlphabeticallyOk() (*bool, bool)`

GetOrderAlphabeticallyOk returns a tuple with the OrderAlphabetically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAlphabetically

`func (o *PatchedWritableCustomFieldChoiceSetRequest) SetOrderAlphabetically(v bool)`

SetOrderAlphabetically sets OrderAlphabetically field to given value.

### HasOrderAlphabetically

`func (o *PatchedWritableCustomFieldChoiceSetRequest) HasOrderAlphabetically() bool`

HasOrderAlphabetically returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


