# PatchedWritableCircuitGroupAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**BriefCircuitGroupAssignmentSerializerRequestGroup**](BriefCircuitGroupAssignmentSerializerRequestGroup.md) |  | [optional] 
**MemberType** | Pointer to **string** |  | [optional] 
**MemberId** | Pointer to **int64** |  | [optional] 
**Priority** | Pointer to [**NullablePatchedWritableCircuitGroupAssignmentRequestPriority**](PatchedWritableCircuitGroupAssignmentRequestPriority.md) |  | [optional] 
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

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetGroup() BriefCircuitGroupAssignmentSerializerRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetGroupOk() (*BriefCircuitGroupAssignmentSerializerRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedWritableCircuitGroupAssignmentRequest) SetGroup(v BriefCircuitGroupAssignmentSerializerRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedWritableCircuitGroupAssignmentRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMemberType

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetMemberType() string`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetMemberTypeOk() (*string, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *PatchedWritableCircuitGroupAssignmentRequest) SetMemberType(v string)`

SetMemberType sets MemberType field to given value.

### HasMemberType

`func (o *PatchedWritableCircuitGroupAssignmentRequest) HasMemberType() bool`

HasMemberType returns a boolean if a field has been set.

### GetMemberId

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *PatchedWritableCircuitGroupAssignmentRequest) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *PatchedWritableCircuitGroupAssignmentRequest) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetPriority

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetPriority() PatchedWritableCircuitGroupAssignmentRequestPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchedWritableCircuitGroupAssignmentRequest) GetPriorityOk() (*PatchedWritableCircuitGroupAssignmentRequestPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchedWritableCircuitGroupAssignmentRequest) SetPriority(v PatchedWritableCircuitGroupAssignmentRequestPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchedWritableCircuitGroupAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *PatchedWritableCircuitGroupAssignmentRequest) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *PatchedWritableCircuitGroupAssignmentRequest) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
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


