# CircuitGroupAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**BriefCircuitGroupRequest**](BriefCircuitGroupRequest.md) |  | 
**Circuit** | [**BriefCircuitRequest**](BriefCircuitRequest.md) |  | 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriorityValue**](BriefCircuitGroupAssignmentSerializerPriorityValue.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewCircuitGroupAssignmentRequest

`func NewCircuitGroupAssignmentRequest(group BriefCircuitGroupRequest, circuit BriefCircuitRequest, ) *CircuitGroupAssignmentRequest`

NewCircuitGroupAssignmentRequest instantiates a new CircuitGroupAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitGroupAssignmentRequestWithDefaults

`func NewCircuitGroupAssignmentRequestWithDefaults() *CircuitGroupAssignmentRequest`

NewCircuitGroupAssignmentRequestWithDefaults instantiates a new CircuitGroupAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *CircuitGroupAssignmentRequest) GetGroup() BriefCircuitGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CircuitGroupAssignmentRequest) GetGroupOk() (*BriefCircuitGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CircuitGroupAssignmentRequest) SetGroup(v BriefCircuitGroupRequest)`

SetGroup sets Group field to given value.


### GetCircuit

`func (o *CircuitGroupAssignmentRequest) GetCircuit() BriefCircuitRequest`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *CircuitGroupAssignmentRequest) GetCircuitOk() (*BriefCircuitRequest, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *CircuitGroupAssignmentRequest) SetCircuit(v BriefCircuitRequest)`

SetCircuit sets Circuit field to given value.


### GetPriority

`func (o *CircuitGroupAssignmentRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CircuitGroupAssignmentRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CircuitGroupAssignmentRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CircuitGroupAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTags

`func (o *CircuitGroupAssignmentRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CircuitGroupAssignmentRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CircuitGroupAssignmentRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CircuitGroupAssignmentRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


