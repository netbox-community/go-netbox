# AggregateFamily

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int32** | * &#x60;4&#x60; - IPv4 * &#x60;6&#x60; - IPv6 | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewAggregateFamily

`func NewAggregateFamily() *AggregateFamily`

NewAggregateFamily instantiates a new AggregateFamily object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateFamilyWithDefaults

`func NewAggregateFamilyWithDefaults() *AggregateFamily`

NewAggregateFamilyWithDefaults instantiates a new AggregateFamily object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *AggregateFamily) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AggregateFamily) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AggregateFamily) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *AggregateFamily) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *AggregateFamily) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AggregateFamily) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AggregateFamily) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AggregateFamily) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


