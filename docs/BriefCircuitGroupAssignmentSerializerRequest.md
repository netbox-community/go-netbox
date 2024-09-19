# BriefCircuitGroupAssignmentSerializerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**BriefCircuitGroupRequest**](BriefCircuitGroupRequest.md) |  | 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriorityValue**](BriefCircuitGroupAssignmentSerializerPriorityValue.md) |  | [optional] 

## Methods

### NewBriefCircuitGroupAssignmentSerializerRequest

`func NewBriefCircuitGroupAssignmentSerializerRequest(group BriefCircuitGroupRequest, ) *BriefCircuitGroupAssignmentSerializerRequest`

NewBriefCircuitGroupAssignmentSerializerRequest instantiates a new BriefCircuitGroupAssignmentSerializerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefCircuitGroupAssignmentSerializerRequestWithDefaults

`func NewBriefCircuitGroupAssignmentSerializerRequestWithDefaults() *BriefCircuitGroupAssignmentSerializerRequest`

NewBriefCircuitGroupAssignmentSerializerRequestWithDefaults instantiates a new BriefCircuitGroupAssignmentSerializerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *BriefCircuitGroupAssignmentSerializerRequest) GetGroup() BriefCircuitGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BriefCircuitGroupAssignmentSerializerRequest) GetGroupOk() (*BriefCircuitGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BriefCircuitGroupAssignmentSerializerRequest) SetGroup(v BriefCircuitGroupRequest)`

SetGroup sets Group field to given value.


### GetPriority

`func (o *BriefCircuitGroupAssignmentSerializerRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BriefCircuitGroupAssignmentSerializerRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BriefCircuitGroupAssignmentSerializerRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BriefCircuitGroupAssignmentSerializerRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


