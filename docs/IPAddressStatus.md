# IPAddressStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**IPAddressStatusValue**](IPAddressStatusValue.md) |  | [optional] 
**Label** | Pointer to [**IPAddressStatusLabel**](IPAddressStatusLabel.md) |  | [optional] 

## Methods

### NewIPAddressStatus

`func NewIPAddressStatus() *IPAddressStatus`

NewIPAddressStatus instantiates a new IPAddressStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressStatusWithDefaults

`func NewIPAddressStatusWithDefaults() *IPAddressStatus`

NewIPAddressStatusWithDefaults instantiates a new IPAddressStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *IPAddressStatus) GetValue() IPAddressStatusValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IPAddressStatus) GetValueOk() (*IPAddressStatusValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IPAddressStatus) SetValue(v IPAddressStatusValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *IPAddressStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *IPAddressStatus) GetLabel() IPAddressStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IPAddressStatus) GetLabelOk() (*IPAddressStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IPAddressStatus) SetLabel(v IPAddressStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *IPAddressStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


