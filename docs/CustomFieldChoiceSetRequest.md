# CustomFieldChoiceSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**BaseChoices** | Pointer to [**CustomFieldChoiceSetBaseChoicesValue**](CustomFieldChoiceSetBaseChoicesValue.md) |  | [optional] 
**ExtraChoices** | **[][]interface{}** |  | 
**OrderAlphabetically** | Pointer to **bool** | Choices are automatically ordered alphabetically | [optional] 

## Methods

### NewCustomFieldChoiceSetRequest

`func NewCustomFieldChoiceSetRequest(name string, extraChoices [][]interface{}, ) *CustomFieldChoiceSetRequest`

NewCustomFieldChoiceSetRequest instantiates a new CustomFieldChoiceSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldChoiceSetRequestWithDefaults

`func NewCustomFieldChoiceSetRequestWithDefaults() *CustomFieldChoiceSetRequest`

NewCustomFieldChoiceSetRequestWithDefaults instantiates a new CustomFieldChoiceSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomFieldChoiceSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFieldChoiceSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFieldChoiceSetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CustomFieldChoiceSetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomFieldChoiceSetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomFieldChoiceSetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomFieldChoiceSetRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBaseChoices

`func (o *CustomFieldChoiceSetRequest) GetBaseChoices() CustomFieldChoiceSetBaseChoicesValue`

GetBaseChoices returns the BaseChoices field if non-nil, zero value otherwise.

### GetBaseChoicesOk

`func (o *CustomFieldChoiceSetRequest) GetBaseChoicesOk() (*CustomFieldChoiceSetBaseChoicesValue, bool)`

GetBaseChoicesOk returns a tuple with the BaseChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseChoices

`func (o *CustomFieldChoiceSetRequest) SetBaseChoices(v CustomFieldChoiceSetBaseChoicesValue)`

SetBaseChoices sets BaseChoices field to given value.

### HasBaseChoices

`func (o *CustomFieldChoiceSetRequest) HasBaseChoices() bool`

HasBaseChoices returns a boolean if a field has been set.

### GetExtraChoices

`func (o *CustomFieldChoiceSetRequest) GetExtraChoices() [][]interface{}`

GetExtraChoices returns the ExtraChoices field if non-nil, zero value otherwise.

### GetExtraChoicesOk

`func (o *CustomFieldChoiceSetRequest) GetExtraChoicesOk() (*[][]interface{}, bool)`

GetExtraChoicesOk returns a tuple with the ExtraChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraChoices

`func (o *CustomFieldChoiceSetRequest) SetExtraChoices(v [][]interface{})`

SetExtraChoices sets ExtraChoices field to given value.


### GetOrderAlphabetically

`func (o *CustomFieldChoiceSetRequest) GetOrderAlphabetically() bool`

GetOrderAlphabetically returns the OrderAlphabetically field if non-nil, zero value otherwise.

### GetOrderAlphabeticallyOk

`func (o *CustomFieldChoiceSetRequest) GetOrderAlphabeticallyOk() (*bool, bool)`

GetOrderAlphabeticallyOk returns a tuple with the OrderAlphabetically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAlphabetically

`func (o *CustomFieldChoiceSetRequest) SetOrderAlphabetically(v bool)`

SetOrderAlphabetically sets OrderAlphabetically field to given value.

### HasOrderAlphabetically

`func (o *CustomFieldChoiceSetRequest) HasOrderAlphabetically() bool`

HasOrderAlphabetically returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


