# WritableServiceTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Ports** | **[]int32** |  | 
**Protocol** | [**PatchedWritableServiceRequestProtocol**](PatchedWritableServiceRequestProtocol.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableServiceTemplateRequest

`func NewWritableServiceTemplateRequest(name string, ports []int32, protocol PatchedWritableServiceRequestProtocol, ) *WritableServiceTemplateRequest`

NewWritableServiceTemplateRequest instantiates a new WritableServiceTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableServiceTemplateRequestWithDefaults

`func NewWritableServiceTemplateRequestWithDefaults() *WritableServiceTemplateRequest`

NewWritableServiceTemplateRequestWithDefaults instantiates a new WritableServiceTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableServiceTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableServiceTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableServiceTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPorts

`func (o *WritableServiceTemplateRequest) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *WritableServiceTemplateRequest) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *WritableServiceTemplateRequest) SetPorts(v []int32)`

SetPorts sets Ports field to given value.


### GetProtocol

`func (o *WritableServiceTemplateRequest) GetProtocol() PatchedWritableServiceRequestProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *WritableServiceTemplateRequest) GetProtocolOk() (*PatchedWritableServiceRequestProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *WritableServiceTemplateRequest) SetProtocol(v PatchedWritableServiceRequestProtocol)`

SetProtocol sets Protocol field to given value.


### GetDescription

`func (o *WritableServiceTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableServiceTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableServiceTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableServiceTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *WritableServiceTemplateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableServiceTemplateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableServiceTemplateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableServiceTemplateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableServiceTemplateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableServiceTemplateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableServiceTemplateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableServiceTemplateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableServiceTemplateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableServiceTemplateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableServiceTemplateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableServiceTemplateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


