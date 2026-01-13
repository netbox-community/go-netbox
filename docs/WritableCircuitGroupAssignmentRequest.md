# WritableCircuitGroupAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**BriefCircuitGroupAssignmentSerializerRequestGroup**](BriefCircuitGroupAssignmentSerializerRequestGroup.md) |  | 
**MemberType** | **string** |  | 
**MemberId** | **int64** |  | 
**Priority** | Pointer to [**NullablePatchedWritableCircuitGroupAssignmentRequestPriority**](PatchedWritableCircuitGroupAssignmentRequestPriority.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewWritableCircuitGroupAssignmentRequest

`func NewWritableCircuitGroupAssignmentRequest(group BriefCircuitGroupAssignmentSerializerRequestGroup, memberType string, memberId int64, ) *WritableCircuitGroupAssignmentRequest`

NewWritableCircuitGroupAssignmentRequest instantiates a new WritableCircuitGroupAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableCircuitGroupAssignmentRequestWithDefaults

`func NewWritableCircuitGroupAssignmentRequestWithDefaults() *WritableCircuitGroupAssignmentRequest`

NewWritableCircuitGroupAssignmentRequestWithDefaults instantiates a new WritableCircuitGroupAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *WritableCircuitGroupAssignmentRequest) GetGroup() BriefCircuitGroupAssignmentSerializerRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WritableCircuitGroupAssignmentRequest) GetGroupOk() (*BriefCircuitGroupAssignmentSerializerRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WritableCircuitGroupAssignmentRequest) SetGroup(v BriefCircuitGroupAssignmentSerializerRequestGroup)`

SetGroup sets Group field to given value.


### GetMemberType

`func (o *WritableCircuitGroupAssignmentRequest) GetMemberType() string`

GetMemberType returns the MemberType field if non-nil, zero value otherwise.

### GetMemberTypeOk

`func (o *WritableCircuitGroupAssignmentRequest) GetMemberTypeOk() (*string, bool)`

GetMemberTypeOk returns a tuple with the MemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberType

`func (o *WritableCircuitGroupAssignmentRequest) SetMemberType(v string)`

SetMemberType sets MemberType field to given value.


### GetMemberId

`func (o *WritableCircuitGroupAssignmentRequest) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *WritableCircuitGroupAssignmentRequest) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *WritableCircuitGroupAssignmentRequest) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.


### GetPriority

`func (o *WritableCircuitGroupAssignmentRequest) GetPriority() PatchedWritableCircuitGroupAssignmentRequestPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WritableCircuitGroupAssignmentRequest) GetPriorityOk() (*PatchedWritableCircuitGroupAssignmentRequestPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WritableCircuitGroupAssignmentRequest) SetPriority(v PatchedWritableCircuitGroupAssignmentRequestPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WritableCircuitGroupAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *WritableCircuitGroupAssignmentRequest) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *WritableCircuitGroupAssignmentRequest) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetTags

`func (o *WritableCircuitGroupAssignmentRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableCircuitGroupAssignmentRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableCircuitGroupAssignmentRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableCircuitGroupAssignmentRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


