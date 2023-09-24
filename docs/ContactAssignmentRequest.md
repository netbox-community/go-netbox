# ContactAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  | 
**ObjectId** | **int64** |  | 
**Contact** | [**NestedContactRequest**](NestedContactRequest.md) |  | 
**Role** | Pointer to [**NullableNestedContactRoleRequest**](NestedContactRoleRequest.md) |  | [optional] 
**Priority** | Pointer to **string** | * &#x60;primary&#x60; - Primary * &#x60;secondary&#x60; - Secondary * &#x60;tertiary&#x60; - Tertiary * &#x60;inactive&#x60; - Inactive | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewContactAssignmentRequest

`func NewContactAssignmentRequest(contentType string, objectId int64, contact NestedContactRequest, ) *ContactAssignmentRequest`

NewContactAssignmentRequest instantiates a new ContactAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactAssignmentRequestWithDefaults

`func NewContactAssignmentRequestWithDefaults() *ContactAssignmentRequest`

NewContactAssignmentRequestWithDefaults instantiates a new ContactAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ContactAssignmentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ContactAssignmentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ContactAssignmentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetObjectId

`func (o *ContactAssignmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ContactAssignmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ContactAssignmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetContact

`func (o *ContactAssignmentRequest) GetContact() NestedContactRequest`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ContactAssignmentRequest) GetContactOk() (*NestedContactRequest, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ContactAssignmentRequest) SetContact(v NestedContactRequest)`

SetContact sets Contact field to given value.


### GetRole

`func (o *ContactAssignmentRequest) GetRole() NestedContactRoleRequest`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ContactAssignmentRequest) GetRoleOk() (*NestedContactRoleRequest, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ContactAssignmentRequest) SetRole(v NestedContactRoleRequest)`

SetRole sets Role field to given value.

### HasRole

`func (o *ContactAssignmentRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ContactAssignmentRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ContactAssignmentRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPriority

`func (o *ContactAssignmentRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ContactAssignmentRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ContactAssignmentRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ContactAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTags

`func (o *ContactAssignmentRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContactAssignmentRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContactAssignmentRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContactAssignmentRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


