# RackReservationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**PatchedWritableRackReservationRequestStatus**](PatchedWritableRackReservationRequestStatus.md) |  | [optional] 
**Label** | Pointer to [**RackReservationStatusLabel**](RackReservationStatusLabel.md) |  | [optional] 

## Methods

### NewRackReservationStatus

`func NewRackReservationStatus() *RackReservationStatus`

NewRackReservationStatus instantiates a new RackReservationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackReservationStatusWithDefaults

`func NewRackReservationStatusWithDefaults() *RackReservationStatus`

NewRackReservationStatusWithDefaults instantiates a new RackReservationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RackReservationStatus) GetValue() PatchedWritableRackReservationRequestStatus`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RackReservationStatus) GetValueOk() (*PatchedWritableRackReservationRequestStatus, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RackReservationStatus) SetValue(v PatchedWritableRackReservationRequestStatus)`

SetValue sets Value field to given value.

### HasValue

`func (o *RackReservationStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *RackReservationStatus) GetLabel() RackReservationStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RackReservationStatus) GetLabelOk() (*RackReservationStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RackReservationStatus) SetLabel(v RackReservationStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RackReservationStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


