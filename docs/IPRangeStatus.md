# IPRangeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;active&#x60; - Active * &#x60;reserved&#x60; - Reserved * &#x60;deprecated&#x60; - Deprecated | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewIPRangeStatus

`func NewIPRangeStatus() *IPRangeStatus`

NewIPRangeStatus instantiates a new IPRangeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPRangeStatusWithDefaults

`func NewIPRangeStatusWithDefaults() *IPRangeStatus`

NewIPRangeStatusWithDefaults instantiates a new IPRangeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *IPRangeStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IPRangeStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IPRangeStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IPRangeStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *IPRangeStatus) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IPRangeStatus) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IPRangeStatus) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IPRangeStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


