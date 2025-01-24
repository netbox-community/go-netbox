# FHRPGroupAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**BriefFHRPGroupRequest**](BriefFHRPGroupRequest.md) |  | 
**InterfaceType** | **string** |  | 
**InterfaceId** | **int64** |  | 
**Priority** | **int32** |  | 

## Methods

### NewFHRPGroupAssignmentRequest

`func NewFHRPGroupAssignmentRequest(group BriefFHRPGroupRequest, interfaceType string, interfaceId int64, priority int32, ) *FHRPGroupAssignmentRequest`

NewFHRPGroupAssignmentRequest instantiates a new FHRPGroupAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFHRPGroupAssignmentRequestWithDefaults

`func NewFHRPGroupAssignmentRequestWithDefaults() *FHRPGroupAssignmentRequest`

NewFHRPGroupAssignmentRequestWithDefaults instantiates a new FHRPGroupAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *FHRPGroupAssignmentRequest) GetGroup() BriefFHRPGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FHRPGroupAssignmentRequest) GetGroupOk() (*BriefFHRPGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FHRPGroupAssignmentRequest) SetGroup(v BriefFHRPGroupRequest)`

SetGroup sets Group field to given value.


### GetInterfaceType

`func (o *FHRPGroupAssignmentRequest) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *FHRPGroupAssignmentRequest) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *FHRPGroupAssignmentRequest) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.


### GetInterfaceId

`func (o *FHRPGroupAssignmentRequest) GetInterfaceId() int64`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *FHRPGroupAssignmentRequest) GetInterfaceIdOk() (*int64, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *FHRPGroupAssignmentRequest) SetInterfaceId(v int64)`

SetInterfaceId sets InterfaceId field to given value.


### GetPriority

`func (o *FHRPGroupAssignmentRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *FHRPGroupAssignmentRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *FHRPGroupAssignmentRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


