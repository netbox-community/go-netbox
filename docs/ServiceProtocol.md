# ServiceProtocol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;tcp&#x60; - TCP * &#x60;udp&#x60; - UDP * &#x60;sctp&#x60; - SCTP | [optional] 
**Label** | Pointer to **string** |  | [optional] 

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

`func (o *ServiceProtocol) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ServiceProtocol) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ServiceProtocol) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ServiceProtocol) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *ServiceProtocol) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServiceProtocol) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServiceProtocol) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServiceProtocol) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


