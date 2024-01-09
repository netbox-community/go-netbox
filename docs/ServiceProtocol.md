# ServiceProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**PatchedWritableServiceRequestProtocol**](PatchedWritableServiceRequestProtocol.md) |  | [optional] 
**Label** | Pointer to [**ServiceProtocolLabel**](ServiceProtocolLabel.md) |  | [optional] 

## Methods

### NewServiceProtocol

`func NewServiceProtocol() *ServiceProtocol`

NewServiceProtocol instantiates a new ServiceProtocol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProtocolWithDefaults

`func NewServiceProtocolWithDefaults() *ServiceProtocol`

NewServiceProtocolWithDefaults instantiates a new ServiceProtocol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ServiceProtocol) GetValue() PatchedWritableServiceRequestProtocol`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServiceProtocol) GetValueOk() (*PatchedWritableServiceRequestProtocol, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServiceProtocol) SetValue(v PatchedWritableServiceRequestProtocol)`

SetValue sets Value field to given value.

### HasValue

`func (o *ServiceProtocol) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *ServiceProtocol) GetLabel() ServiceProtocolLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServiceProtocol) GetLabelOk() (*ServiceProtocolLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServiceProtocol) SetLabel(v ServiceProtocolLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServiceProtocol) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


