# CableLengthUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;km&#x60; - Kilometers * &#x60;m&#x60; - Meters * &#x60;cm&#x60; - Centimeters * &#x60;mi&#x60; - Miles * &#x60;ft&#x60; - Feet * &#x60;in&#x60; - Inches | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewCableLengthUnit

`func NewCableLengthUnit() *CableLengthUnit`

NewCableLengthUnit instantiates a new CableLengthUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCableLengthUnitWithDefaults

`func NewCableLengthUnitWithDefaults() *CableLengthUnit`

NewCableLengthUnitWithDefaults instantiates a new CableLengthUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CableLengthUnit) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CableLengthUnit) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CableLengthUnit) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CableLengthUnit) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *CableLengthUnit) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CableLengthUnit) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CableLengthUnit) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CableLengthUnit) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


