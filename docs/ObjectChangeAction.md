# ObjectChangeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**ObjectChangeActionValue**](ObjectChangeActionValue.md) |  | [optional] 
**Label** | Pointer to [**ObjectChangeActionLabel**](ObjectChangeActionLabel.md) |  | [optional] 

## Methods

### NewObjectChangeAction

`func NewObjectChangeAction() *ObjectChangeAction`

NewObjectChangeAction instantiates a new ObjectChangeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectChangeActionWithDefaults

`func NewObjectChangeActionWithDefaults() *ObjectChangeAction`

NewObjectChangeActionWithDefaults instantiates a new ObjectChangeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ObjectChangeAction) GetValue() ObjectChangeActionValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ObjectChangeAction) GetValueOk() (*ObjectChangeActionValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ObjectChangeAction) SetValue(v ObjectChangeActionValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ObjectChangeAction) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *ObjectChangeAction) GetLabel() ObjectChangeActionLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ObjectChangeAction) GetLabelOk() (*ObjectChangeActionLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ObjectChangeAction) SetLabel(v ObjectChangeActionLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ObjectChangeAction) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


