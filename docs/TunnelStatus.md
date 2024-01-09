# TunnelStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**PatchedWritableTunnelRequestStatus**](PatchedWritableTunnelRequestStatus.md) |  | [optional] 
**Label** | Pointer to [**TunnelStatusLabel**](TunnelStatusLabel.md) |  | [optional] 

## Methods

### NewTunnelStatus

`func NewTunnelStatus() *TunnelStatus`

NewTunnelStatus instantiates a new TunnelStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelStatusWithDefaults

`func NewTunnelStatusWithDefaults() *TunnelStatus`

NewTunnelStatusWithDefaults instantiates a new TunnelStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *TunnelStatus) GetValue() PatchedWritableTunnelRequestStatus`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TunnelStatus) GetValueOk() (*PatchedWritableTunnelRequestStatus, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TunnelStatus) SetValue(v PatchedWritableTunnelRequestStatus)`

SetValue sets Value field to given value.

### HasValue

`func (o *TunnelStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *TunnelStatus) GetLabel() TunnelStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TunnelStatus) GetLabelOk() (*TunnelStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TunnelStatus) SetLabel(v TunnelStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TunnelStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


