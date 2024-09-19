# PatchedSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**User** | Pointer to [**BriefUserRequest**](BriefUserRequest.md) |  | [optional] 

## Methods

### NewPatchedSubscriptionRequest

`func NewPatchedSubscriptionRequest() *PatchedSubscriptionRequest`

NewPatchedSubscriptionRequest instantiates a new PatchedSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSubscriptionRequestWithDefaults

`func NewPatchedSubscriptionRequestWithDefaults() *PatchedSubscriptionRequest`

NewPatchedSubscriptionRequestWithDefaults instantiates a new PatchedSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *PatchedSubscriptionRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedSubscriptionRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedSubscriptionRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedSubscriptionRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedSubscriptionRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedSubscriptionRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedSubscriptionRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedSubscriptionRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetUser

`func (o *PatchedSubscriptionRequest) GetUser() BriefUserRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedSubscriptionRequest) GetUserOk() (*BriefUserRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedSubscriptionRequest) SetUser(v BriefUserRequest)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedSubscriptionRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


