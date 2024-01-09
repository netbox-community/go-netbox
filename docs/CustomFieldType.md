# CustomFieldType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**CustomFieldTypeValue**](CustomFieldTypeValue.md) |  | [optional] 
**Label** | Pointer to [**CustomFieldTypeLabel**](CustomFieldTypeLabel.md) |  | [optional] 

## Methods

### NewCustomFieldType

`func NewCustomFieldType() *CustomFieldType`

NewCustomFieldType instantiates a new CustomFieldType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldTypeWithDefaults

`func NewCustomFieldTypeWithDefaults() *CustomFieldType`

NewCustomFieldTypeWithDefaults instantiates a new CustomFieldType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CustomFieldType) GetValue() CustomFieldTypeValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldType) GetValueOk() (*CustomFieldTypeValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldType) SetValue(v CustomFieldTypeValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldType) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *CustomFieldType) GetLabel() CustomFieldTypeLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomFieldType) GetLabelOk() (*CustomFieldTypeLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomFieldType) SetLabel(v CustomFieldTypeLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomFieldType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


