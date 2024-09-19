# PatchedWritableCircuitGroupAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**BriefCircuitGroupRequest**](BriefCircuitGroupRequest.md) |  | [optional] 
**Circuit** | Pointer to [**BriefCircuitRequest**](BriefCircuitRequest.md) |  | [optional] 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriorityValue**](BriefCircuitGroupAssignmentSerializerPriorityValue.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewPatchedWritableCircuitGroupAssignmentRequest

`func NewPatchedWritableCircuitGroupAssignmentRequest() *PatchedWritableCircuitGroupAssignmentRequest`

NewPatchedWritableCircuitGroupAssignmentRequest instantiates a new PatchedWritableCircuitGroupAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableCircuitGroupAssignmentRequestWithDefaults

`func NewPatchedWritableCircuitGroupAssignmentRequestWithDefaults() *PatchedWritableCircuitGroupAssignmentRequest`

NewPatchedWritableCircuitGroupAssignmentRequestWithDefaults instantiates a new PatchedWritableCircuitGroupAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetGroup() BriefCircuitGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetGroupOk() (*BriefCircuitGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedWritableCircuitGroupAssignmentRequest) SetGroup(v BriefCircuitGroupRequest)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedWritableCircuitGroupAssignmentRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCircuit

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetCircuit() BriefCircuitRequest`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetCircuitOk() (*BriefCircuitRequest, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *PatchedWritableCircuitGroupAssignmentRequest) SetCircuit(v BriefCircuitRequest)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *PatchedWritableCircuitGroupAssignmentRequest) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetPriority

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchedWritableCircuitGroupAssignmentRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchedWritableCircuitGroupAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableCircuitGroupAssignmentRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableCircuitGroupAssignmentRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


