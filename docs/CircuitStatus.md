# CircuitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**CircuitStatusValue**](CircuitStatusValue.md) |  | [optional] 
**Label** | Pointer to [**CircuitStatusLabel**](CircuitStatusLabel.md) |  | [optional] 

## Methods

### NewCircuitStatus

`func NewCircuitStatus() *CircuitStatus`

NewCircuitStatus instantiates a new CircuitStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitStatusWithDefaults

`func NewCircuitStatusWithDefaults() *CircuitStatus`

NewCircuitStatusWithDefaults instantiates a new CircuitStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CircuitStatus) GetValue() CircuitStatusValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CircuitStatus) GetValueOk() (*CircuitStatusValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CircuitStatus) SetValue(v CircuitStatusValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *CircuitStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *CircuitStatus) GetLabel() CircuitStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CircuitStatus) GetLabelOk() (*CircuitStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CircuitStatus) SetLabel(v CircuitStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CircuitStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


