# RackStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**PatchedWritableRackRequestStatus**](PatchedWritableRackRequestStatus.md) |  | [optional] 
**Label** | Pointer to [**RackStatusLabel**](RackStatusLabel.md) |  | [optional] 

## Methods

### NewRackStatus

`func NewRackStatus() *RackStatus`

NewRackStatus instantiates a new RackStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackStatusWithDefaults

`func NewRackStatusWithDefaults() *RackStatus`

NewRackStatusWithDefaults instantiates a new RackStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RackStatus) GetValue() PatchedWritableRackRequestStatus`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RackStatus) GetValueOk() (*PatchedWritableRackRequestStatus, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RackStatus) SetValue(v PatchedWritableRackRequestStatus)`

SetValue sets Value field to given value.

### HasValue

`func (o *RackStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *RackStatus) GetLabel() RackStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RackStatus) GetLabelOk() (*RackStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RackStatus) SetLabel(v RackStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RackStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


