# CustomFieldChoiceSetBaseChoices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;IATA&#x60; - IATA (Airport codes) * &#x60;ISO_3166&#x60; - ISO 3166 (Country codes) * &#x60;UN_LOCODE&#x60; - UN/LOCODE (Location codes) | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomFieldChoiceSetBaseChoices

`func NewCustomFieldChoiceSetBaseChoices() *CustomFieldChoiceSetBaseChoices`

NewCustomFieldChoiceSetBaseChoices instantiates a new CustomFieldChoiceSetBaseChoices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldChoiceSetBaseChoicesWithDefaults

`func NewCustomFieldChoiceSetBaseChoicesWithDefaults() *CustomFieldChoiceSetBaseChoices`

NewCustomFieldChoiceSetBaseChoicesWithDefaults instantiates a new CustomFieldChoiceSetBaseChoices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CustomFieldChoiceSetBaseChoices) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldChoiceSetBaseChoices) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldChoiceSetBaseChoices) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldChoiceSetBaseChoices) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *CustomFieldChoiceSetBaseChoices) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomFieldChoiceSetBaseChoices) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomFieldChoiceSetBaseChoices) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomFieldChoiceSetBaseChoices) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


