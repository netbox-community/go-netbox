# CableStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**CableStatusValue**](CableStatusValue.md) |  | [optional] 
**Label** | Pointer to [**CableStatusLabel**](CableStatusLabel.md) |  | [optional] 

## Methods

### NewCableStatus

`func NewCableStatus() *CableStatus`

NewCableStatus instantiates a new CableStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCableStatusWithDefaults

`func NewCableStatusWithDefaults() *CableStatus`

NewCableStatusWithDefaults instantiates a new CableStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CableStatus) GetValue() CableStatusValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CableStatus) GetValueOk() (*CableStatusValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CableStatus) SetValue(v CableStatusValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *CableStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *CableStatus) GetLabel() CableStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CableStatus) GetLabelOk() (*CableStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CableStatus) SetLabel(v CableStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CableStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


