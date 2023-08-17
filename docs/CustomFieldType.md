# CustomFieldType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;text&#x60; - Text * &#x60;longtext&#x60; - Text (long) * &#x60;integer&#x60; - Integer * &#x60;decimal&#x60; - Decimal * &#x60;boolean&#x60; - Boolean (true/false) * &#x60;date&#x60; - Date * &#x60;datetime&#x60; - Date &amp; time * &#x60;url&#x60; - URL * &#x60;json&#x60; - JSON * &#x60;select&#x60; - Selection * &#x60;multiselect&#x60; - Multiple selection * &#x60;object&#x60; - Object * &#x60;multiobject&#x60; - Multiple objects | [optional] 
**Label** | Pointer to **string** |  | [optional] 

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

`func (o *CustomFieldType) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldType) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldType) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldType) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *CustomFieldType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomFieldType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomFieldType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomFieldType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


