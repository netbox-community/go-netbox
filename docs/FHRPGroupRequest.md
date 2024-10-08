# FHRPGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | [**FHRPGroupProtocol**](FHRPGroupProtocol.md) |  | 
**GroupId** | **int32** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewFHRPGroupRequest

`func NewFHRPGroupRequest(protocol FHRPGroupProtocol, groupId int32, ) *FHRPGroupRequest`

NewFHRPGroupRequest instantiates a new FHRPGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFHRPGroupRequestWithDefaults

`func NewFHRPGroupRequestWithDefaults() *FHRPGroupRequest`

NewFHRPGroupRequestWithDefaults instantiates a new FHRPGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *FHRPGroupRequest) GetProtocol() FHRPGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FHRPGroupRequest) GetProtocolOk() (*FHRPGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FHRPGroupRequest) SetProtocol(v FHRPGroupProtocol)`

SetProtocol sets Protocol field to given value.


### GetGroupId

`func (o *FHRPGroupRequest) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *FHRPGroupRequest) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *FHRPGroupRequest) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetDescription

`func (o *FHRPGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FHRPGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FHRPGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FHRPGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


