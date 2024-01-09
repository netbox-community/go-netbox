# PowerFeedStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**PatchedWritablePowerFeedRequestStatus**](PatchedWritablePowerFeedRequestStatus.md) |  | [optional] 
**Label** | Pointer to [**PowerFeedStatusLabel**](PowerFeedStatusLabel.md) |  | [optional] 

## Methods

### NewPowerFeedStatus

`func NewPowerFeedStatus() *PowerFeedStatus`

NewPowerFeedStatus instantiates a new PowerFeedStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerFeedStatusWithDefaults

`func NewPowerFeedStatusWithDefaults() *PowerFeedStatus`

NewPowerFeedStatusWithDefaults instantiates a new PowerFeedStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PowerFeedStatus) GetValue() PatchedWritablePowerFeedRequestStatus`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PowerFeedStatus) GetValueOk() (*PatchedWritablePowerFeedRequestStatus, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PowerFeedStatus) SetValue(v PatchedWritablePowerFeedRequestStatus)`

SetValue sets Value field to given value.

### HasValue

`func (o *PowerFeedStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *PowerFeedStatus) GetLabel() PowerFeedStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PowerFeedStatus) GetLabelOk() (*PowerFeedStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PowerFeedStatus) SetLabel(v PowerFeedStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PowerFeedStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


