# WritableContactAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  | 
**ObjectId** | **int64** |  | 
**Contact** | **int32** |  | 
**Role** | **int32** |  | 
**Priority** | Pointer to **string** | * &#x60;primary&#x60; - Primary * &#x60;secondary&#x60; - Secondary * &#x60;tertiary&#x60; - Tertiary * &#x60;inactive&#x60; - Inactive | [optional] 

## Methods

### NewWritableContactAssignmentRequest

`func NewWritableContactAssignmentRequest(contentType string, objectId int64, contact int32, role int32, ) *WritableContactAssignmentRequest`

NewWritableContactAssignmentRequest instantiates a new WritableContactAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableContactAssignmentRequestWithDefaults

`func NewWritableContactAssignmentRequestWithDefaults() *WritableContactAssignmentRequest`

NewWritableContactAssignmentRequestWithDefaults instantiates a new WritableContactAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *WritableContactAssignmentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WritableContactAssignmentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WritableContactAssignmentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetObjectId

`func (o *WritableContactAssignmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *WritableContactAssignmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *WritableContactAssignmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetContact

`func (o *WritableContactAssignmentRequest) GetContact() int32`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *WritableContactAssignmentRequest) GetContactOk() (*int32, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *WritableContactAssignmentRequest) SetContact(v int32)`

SetContact sets Contact field to given value.


### GetRole

`func (o *WritableContactAssignmentRequest) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *WritableContactAssignmentRequest) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *WritableContactAssignmentRequest) SetRole(v int32)`

SetRole sets Role field to given value.


### GetPriority

`func (o *WritableContactAssignmentRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *WritableContactAssignmentRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *WritableContactAssignmentRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *WritableContactAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


