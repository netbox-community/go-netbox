# WritableServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**NullableBriefDeviceRequest**](BriefDeviceRequest.md) |  | [optional] 
**VirtualMachine** | Pointer to [**NullableBriefVirtualMachineRequest**](BriefVirtualMachineRequest.md) |  | [optional] 
**Name** | **string** |  | 
**Protocol** | [**PatchedWritableServiceRequestProtocol**](PatchedWritableServiceRequestProtocol.md) |  | 
**Ports** | **[]int32** |  | 
**Ipaddresses** | Pointer to **[]int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableServiceRequest

`func NewWritableServiceRequest(name string, protocol PatchedWritableServiceRequestProtocol, ports []int32, ) *WritableServiceRequest`

NewWritableServiceRequest instantiates a new WritableServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableServiceRequestWithDefaults

`func NewWritableServiceRequestWithDefaults() *WritableServiceRequest`

NewWritableServiceRequestWithDefaults instantiates a new WritableServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *WritableServiceRequest) GetDevice() BriefDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *WritableServiceRequest) GetDeviceOk() (*BriefDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *WritableServiceRequest) SetDevice(v BriefDeviceRequest)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *WritableServiceRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *WritableServiceRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *WritableServiceRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetVirtualMachine

`func (o *WritableServiceRequest) GetVirtualMachine() BriefVirtualMachineRequest`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *WritableServiceRequest) GetVirtualMachineOk() (*BriefVirtualMachineRequest, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *WritableServiceRequest) SetVirtualMachine(v BriefVirtualMachineRequest)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *WritableServiceRequest) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### SetVirtualMachineNil

`func (o *WritableServiceRequest) SetVirtualMachineNil(b bool)`

 SetVirtualMachineNil sets the value for VirtualMachine to be an explicit nil

### UnsetVirtualMachine
`func (o *WritableServiceRequest) UnsetVirtualMachine()`

UnsetVirtualMachine ensures that no value is present for VirtualMachine, not even an explicit nil
### GetName

`func (o *WritableServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableServiceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *WritableServiceRequest) GetProtocol() PatchedWritableServiceRequestProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *WritableServiceRequest) GetProtocolOk() (*PatchedWritableServiceRequestProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *WritableServiceRequest) SetProtocol(v PatchedWritableServiceRequestProtocol)`

SetProtocol sets Protocol field to given value.


### GetPorts

`func (o *WritableServiceRequest) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *WritableServiceRequest) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *WritableServiceRequest) SetPorts(v []int32)`

SetPorts sets Ports field to given value.


### GetIpaddresses

`func (o *WritableServiceRequest) GetIpaddresses() []int32`

GetIpaddresses returns the Ipaddresses field if non-nil, zero value otherwise.

### GetIpaddressesOk

`func (o *WritableServiceRequest) GetIpaddressesOk() (*[]int32, bool)`

GetIpaddressesOk returns a tuple with the Ipaddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddresses

`func (o *WritableServiceRequest) SetIpaddresses(v []int32)`

SetIpaddresses sets Ipaddresses field to given value.

### HasIpaddresses

`func (o *WritableServiceRequest) HasIpaddresses() bool`

HasIpaddresses returns a boolean if a field has been set.

### GetDescription

`func (o *WritableServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *WritableServiceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableServiceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableServiceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableServiceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableServiceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableServiceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableServiceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableServiceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableServiceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableServiceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableServiceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableServiceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


