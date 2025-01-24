# NotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**User** | [**BriefUserRequest**](BriefUserRequest.md) |  | 
**Read** | Pointer to **NullableTime** |  | [optional] 
**EventType** | [**Event**](Event.md) |  | 

## Methods

### NewNotificationRequest

`func NewNotificationRequest(objectType string, objectId int64, user BriefUserRequest, eventType Event, ) *NotificationRequest`

NewNotificationRequest instantiates a new NotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRequestWithDefaults

`func NewNotificationRequestWithDefaults() *NotificationRequest`

NewNotificationRequestWithDefaults instantiates a new NotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *NotificationRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NotificationRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NotificationRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *NotificationRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *NotificationRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *NotificationRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetUser

`func (o *NotificationRequest) GetUser() BriefUserRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NotificationRequest) GetUserOk() (*BriefUserRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NotificationRequest) SetUser(v BriefUserRequest)`

SetUser sets User field to given value.


### GetRead

`func (o *NotificationRequest) GetRead() time.Time`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *NotificationRequest) GetReadOk() (*time.Time, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *NotificationRequest) SetRead(v time.Time)`

SetRead sets Read field to given value.

### HasRead

`func (o *NotificationRequest) HasRead() bool`

HasRead returns a boolean if a field has been set.

### SetReadNil

`func (o *NotificationRequest) SetReadNil(b bool)`

 SetReadNil sets the value for Read to be an explicit nil

### UnsetRead
`func (o *NotificationRequest) UnsetRead()`

UnsetRead ensures that no value is present for Read, not even an explicit nil
### GetEventType

`func (o *NotificationRequest) GetEventType() Event`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *NotificationRequest) GetEventTypeOk() (*Event, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *NotificationRequest) SetEventType(v Event)`

SetEventType sets EventType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


