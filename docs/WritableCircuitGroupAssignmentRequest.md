# WritableCircuitGroupAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**BriefCircuitGroupRequest**](BriefCircuitGroupRequest.md) |  | 
**Circuit** | [**BriefCircuitRequest**](BriefCircuitRequest.md) |  | 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriorityValue**](BriefCircuitGroupAssignmentSerializerPriorityValue.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewWritableCircuitGroupAssignmentRequest

`func NewWritableCircuitGroupAssignmentRequest(group BriefCircuitGroupRequest, circuit BriefCircuitRequest, ) *WritableCircuitGroupAssignmentRequest`

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

`func (o *WritableCircuitGroupAssignmentRequest) GetGroup() BriefCircuitGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WritableCircuitGroupAssignmentRequest) GetGroupOk() (*BriefCircuitGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WritableCircuitGroupAssignmentRequest) SetGroup(v BriefCircuitGroupRequest)`

SetGroup sets Group field to given value.


### GetCircuit

`func (o *WritableCircuitGroupAssignmentRequest) GetCircuit() BriefCircuitRequest`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *WritableCircuitGroupAssignmentRequest) GetCircuitOk() (*BriefCircuitRequest, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *WritableCircuitGroupAssignmentRequest) SetCircuit(v BriefCircuitRequest)`

SetCircuit sets Circuit field to given value.


### GetPriority

`func (o *WritableCircuitGroupAssignmentRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WritableCircuitGroupAssignmentRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WritableCircuitGroupAssignmentRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WritableCircuitGroupAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

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


