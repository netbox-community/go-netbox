# PowerOutletStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**PatchedWritablePowerOutletRequestStatus**](PatchedWritablePowerOutletRequestStatus.md) |  | [optional] 
**Label** | Pointer to [**PowerOutletStatusLabel**](PowerOutletStatusLabel.md) |  | [optional] 

## Methods

### NewPowerOutletStatus

`func NewPowerOutletStatus() *PowerOutletStatus`

NewPowerOutletStatus instantiates a new PowerOutletStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerOutletStatusWithDefaults

`func NewPowerOutletStatusWithDefaults() *PowerOutletStatus`

NewPowerOutletStatusWithDefaults instantiates a new PowerOutletStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PowerOutletStatus) GetValue() PatchedWritablePowerOutletRequestStatus`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PowerOutletStatus) GetValueOk() (*PatchedWritablePowerOutletRequestStatus, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PowerOutletStatus) SetValue(v PatchedWritablePowerOutletRequestStatus)`

SetValue sets Value field to given value.

### HasValue

`func (o *PowerOutletStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *PowerOutletStatus) GetLabel() PowerOutletStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PowerOutletStatus) GetLabelOk() (*PowerOutletStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PowerOutletStatus) SetLabel(v PowerOutletStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PowerOutletStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


