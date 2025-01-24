# BriefFHRPGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**BriefFHRPGroupProtocol**](BriefFHRPGroupProtocol.md) |  | 
**GroupId** | **int32** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBriefFHRPGroupRequest

`func NewBriefFHRPGroupRequest(protocol BriefFHRPGroupProtocol, groupId int32, ) *BriefFHRPGroupRequest`

NewBriefFHRPGroupRequest instantiates a new BriefFHRPGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBriefFHRPGroupRequestWithDefaults

`func NewBriefFHRPGroupRequestWithDefaults() *BriefFHRPGroupRequest`

NewBriefFHRPGroupRequestWithDefaults instantiates a new BriefFHRPGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *BriefFHRPGroupRequest) GetProtocol() BriefFHRPGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BriefFHRPGroupRequest) GetProtocolOk() (*BriefFHRPGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BriefFHRPGroupRequest) SetProtocol(v BriefFHRPGroupProtocol)`

SetProtocol sets Protocol field to given value.


### GetGroupId

`func (o *BriefFHRPGroupRequest) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BriefFHRPGroupRequest) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BriefFHRPGroupRequest) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetDescription

`func (o *BriefFHRPGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BriefFHRPGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BriefFHRPGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BriefFHRPGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


