# PatchedWritableServiceTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to **[]int32** |  | [optional] 
**Protocol** | Pointer to [**PatchedWritableServiceRequestProtocol**](PatchedWritableServiceRequestProtocol.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableServiceTemplateRequest

`func NewPatchedWritableServiceTemplateRequest() *PatchedWritableServiceTemplateRequest`

NewPatchedWritableServiceTemplateRequest instantiates a new PatchedWritableServiceTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableServiceTemplateRequestWithDefaults

`func NewPatchedWritableServiceTemplateRequestWithDefaults() *PatchedWritableServiceTemplateRequest`

NewPatchedWritableServiceTemplateRequestWithDefaults instantiates a new PatchedWritableServiceTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableServiceTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableServiceTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableServiceTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableServiceTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *PatchedWritableServiceTemplateRequest) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *PatchedWritableServiceTemplateRequest) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *PatchedWritableServiceTemplateRequest) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *PatchedWritableServiceTemplateRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchedWritableServiceTemplateRequest) GetProtocol() PatchedWritableServiceRequestProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchedWritableServiceTemplateRequest) GetProtocolOk() (*PatchedWritableServiceRequestProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchedWritableServiceTemplateRequest) SetProtocol(v PatchedWritableServiceRequestProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchedWritableServiceTemplateRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableServiceTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableServiceTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableServiceTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableServiceTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableServiceTemplateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableServiceTemplateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableServiceTemplateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableServiceTemplateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableServiceTemplateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableServiceTemplateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableServiceTemplateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableServiceTemplateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableServiceTemplateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableServiceTemplateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableServiceTemplateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableServiceTemplateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


