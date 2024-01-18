# NestedFHRPGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**FHRPGroupProtocol**](FHRPGroupProtocol.md) |  | 
**GroupId** | **int32** |  | 

## Methods

### NewNestedFHRPGroupRequest

`func NewNestedFHRPGroupRequest(protocol FHRPGroupProtocol, groupId int32, ) *NestedFHRPGroupRequest`

NewNestedFHRPGroupRequest instantiates a new NestedFHRPGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNestedFHRPGroupRequestWithDefaults

`func NewNestedFHRPGroupRequestWithDefaults() *NestedFHRPGroupRequest`

NewNestedFHRPGroupRequestWithDefaults instantiates a new NestedFHRPGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *NestedFHRPGroupRequest) GetProtocol() FHRPGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NestedFHRPGroupRequest) GetProtocolOk() (*FHRPGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NestedFHRPGroupRequest) SetProtocol(v FHRPGroupProtocol)`

SetProtocol sets Protocol field to given value.


### GetGroupId

`func (o *NestedFHRPGroupRequest) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NestedFHRPGroupRequest) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NestedFHRPGroupRequest) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


