# CustomFieldUiVisibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;read-write&#x60; - Read/Write * &#x60;read-only&#x60; - Read-only * &#x60;hidden&#x60; - Hidden * &#x60;hidden-ifunset&#x60; - Hidden (if unset) | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomFieldUiVisibility

`func NewCustomFieldUiVisibility() *CustomFieldUiVisibility`

NewCustomFieldUiVisibility instantiates a new CustomFieldUiVisibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldUiVisibilityWithDefaults

`func NewCustomFieldUiVisibilityWithDefaults() *CustomFieldUiVisibility`

NewCustomFieldUiVisibilityWithDefaults instantiates a new CustomFieldUiVisibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CustomFieldUiVisibility) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldUiVisibility) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldUiVisibility) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldUiVisibility) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *CustomFieldUiVisibility) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomFieldUiVisibility) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomFieldUiVisibility) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CustomFieldUiVisibility) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


