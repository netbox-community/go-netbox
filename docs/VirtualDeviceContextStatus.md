# VirtualDeviceContextStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**PatchedWritableVirtualDeviceContextRequestStatus**](PatchedWritableVirtualDeviceContextRequestStatus.md) |  | [optional] 
**Label** | Pointer to [**VirtualDeviceContextStatusLabel**](VirtualDeviceContextStatusLabel.md) |  | [optional] 

## Methods

### NewVirtualDeviceContextStatus

`func NewVirtualDeviceContextStatus() *VirtualDeviceContextStatus`

NewVirtualDeviceContextStatus instantiates a new VirtualDeviceContextStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDeviceContextStatusWithDefaults

`func NewVirtualDeviceContextStatusWithDefaults() *VirtualDeviceContextStatus`

NewVirtualDeviceContextStatusWithDefaults instantiates a new VirtualDeviceContextStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *VirtualDeviceContextStatus) GetValue() PatchedWritableVirtualDeviceContextRequestStatus`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VirtualDeviceContextStatus) GetValueOk() (*PatchedWritableVirtualDeviceContextRequestStatus, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VirtualDeviceContextStatus) SetValue(v PatchedWritableVirtualDeviceContextRequestStatus)`

SetValue sets Value field to given value.

### HasValue

`func (o *VirtualDeviceContextStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *VirtualDeviceContextStatus) GetLabel() VirtualDeviceContextStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VirtualDeviceContextStatus) GetLabelOk() (*VirtualDeviceContextStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VirtualDeviceContextStatus) SetLabel(v VirtualDeviceContextStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VirtualDeviceContextStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


