# WritableVirtualDiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualMachine** | **int32** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Size** | **int32** |  | 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableVirtualDiskRequest

`func NewWritableVirtualDiskRequest(virtualMachine int32, name string, size int32, ) *WritableVirtualDiskRequest`

NewWritableVirtualDiskRequest instantiates a new WritableVirtualDiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableVirtualDiskRequestWithDefaults

`func NewWritableVirtualDiskRequestWithDefaults() *WritableVirtualDiskRequest`

NewWritableVirtualDiskRequestWithDefaults instantiates a new WritableVirtualDiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualMachine

`func (o *WritableVirtualDiskRequest) GetVirtualMachine() int32`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *WritableVirtualDiskRequest) GetVirtualMachineOk() (*int32, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *WritableVirtualDiskRequest) SetVirtualMachine(v int32)`

SetVirtualMachine sets VirtualMachine field to given value.


### GetName

`func (o *WritableVirtualDiskRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableVirtualDiskRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableVirtualDiskRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WritableVirtualDiskRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableVirtualDiskRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableVirtualDiskRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableVirtualDiskRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSize

`func (o *WritableVirtualDiskRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *WritableVirtualDiskRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *WritableVirtualDiskRequest) SetSize(v int32)`

SetSize sets Size field to given value.


### GetTags

`func (o *WritableVirtualDiskRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableVirtualDiskRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableVirtualDiskRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableVirtualDiskRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableVirtualDiskRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableVirtualDiskRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableVirtualDiskRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableVirtualDiskRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


