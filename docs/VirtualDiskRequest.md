# VirtualDiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualMachine** | [**BriefVirtualMachineRequest**](BriefVirtualMachineRequest.md) |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Size** | **int32** |  | 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVirtualDiskRequest

`func NewVirtualDiskRequest(virtualMachine BriefVirtualMachineRequest, name string, size int32, ) *VirtualDiskRequest`

NewVirtualDiskRequest instantiates a new VirtualDiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDiskRequestWithDefaults

`func NewVirtualDiskRequestWithDefaults() *VirtualDiskRequest`

NewVirtualDiskRequestWithDefaults instantiates a new VirtualDiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualMachine

`func (o *VirtualDiskRequest) GetVirtualMachine() BriefVirtualMachineRequest`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VirtualDiskRequest) GetVirtualMachineOk() (*BriefVirtualMachineRequest, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VirtualDiskRequest) SetVirtualMachine(v BriefVirtualMachineRequest)`

SetVirtualMachine sets VirtualMachine field to given value.


### GetName

`func (o *VirtualDiskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualDiskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualDiskRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VirtualDiskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualDiskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualDiskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualDiskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSize

`func (o *VirtualDiskRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VirtualDiskRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VirtualDiskRequest) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTags

`func (o *VirtualDiskRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualDiskRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualDiskRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualDiskRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualDiskRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualDiskRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualDiskRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualDiskRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


