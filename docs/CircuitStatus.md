# CircuitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;planned&#x60; - Planned * &#x60;provisioning&#x60; - Provisioning * &#x60;active&#x60; - Active * &#x60;offline&#x60; - Offline * &#x60;deprovisioning&#x60; - Deprovisioning * &#x60;decommissioned&#x60; - Decommissioned | [optional] 
**Label** | Pointer to **string** |  | [optional] 

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

`func (o *CircuitStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CircuitStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CircuitStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CircuitStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *CircuitStatus) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CircuitStatus) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CircuitStatus) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CircuitStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


