# PatchedNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**User** | Pointer to [**BriefUserRequest**](BriefUserRequest.md) |  | [optional] 
**Read** | Pointer to **NullableTime** |  | [optional] 
**EventType** | Pointer to [**Event**](Event.md) |  | [optional] 

## Methods

### NewPatchedNotificationRequest

`func NewPatchedNotificationRequest() *PatchedNotificationRequest`

NewPatchedNotificationRequest instantiates a new PatchedNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedNotificationRequestWithDefaults

`func NewPatchedNotificationRequestWithDefaults() *PatchedNotificationRequest`

NewPatchedNotificationRequestWithDefaults instantiates a new PatchedNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *PatchedNotificationRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedNotificationRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedNotificationRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedNotificationRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedNotificationRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedNotificationRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedNotificationRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedNotificationRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetUser

`func (o *PatchedNotificationRequest) GetUser() BriefUserRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedNotificationRequest) GetUserOk() (*BriefUserRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedNotificationRequest) SetUser(v BriefUserRequest)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedNotificationRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRead

`func (o *PatchedNotificationRequest) GetRead() time.Time`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *PatchedNotificationRequest) GetReadOk() (*time.Time, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *PatchedNotificationRequest) SetRead(v time.Time)`

SetRead sets Read field to given value.

### HasRead

`func (o *PatchedNotificationRequest) HasRead() bool`

HasRead returns a boolean if a field has been set.

### SetReadNil

`func (o *PatchedNotificationRequest) SetReadNil(b bool)`

 SetReadNil sets the value for Read to be an explicit nil

### UnsetRead
`func (o *PatchedNotificationRequest) UnsetRead()`

UnsetRead ensures that no value is present for Read, not even an explicit nil
### GetEventType

`func (o *PatchedNotificationRequest) GetEventType() Event`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PatchedNotificationRequest) GetEventTypeOk() (*Event, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PatchedNotificationRequest) SetEventType(v Event)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *PatchedNotificationRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


