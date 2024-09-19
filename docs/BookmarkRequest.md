# BookmarkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**User** | [**BriefUserRequest**](BriefUserRequest.md) |  | 

## Methods

### NewBookmarkRequest

`func NewBookmarkRequest(objectType string, objectId int64, user BriefUserRequest, ) *BookmarkRequest`

NewBookmarkRequest instantiates a new BookmarkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkRequestWithDefaults

`func NewBookmarkRequestWithDefaults() *BookmarkRequest`

NewBookmarkRequestWithDefaults instantiates a new BookmarkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *BookmarkRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BookmarkRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BookmarkRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *BookmarkRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BookmarkRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BookmarkRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetUser

`func (o *BookmarkRequest) GetUser() BriefUserRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BookmarkRequest) GetUserOk() (*BriefUserRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BookmarkRequest) SetUser(v BriefUserRequest)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


