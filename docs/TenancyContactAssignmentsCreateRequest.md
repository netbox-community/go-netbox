# TenancyContactAssignmentsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**Contact** | [**ContactAssignmentRequestContact**](ContactAssignmentRequestContact.md) |  | 
**Role** | Pointer to [**NullableContactAssignmentRequestRole**](ContactAssignmentRequestRole.md) |  | [optional] 
**Priority** | Pointer to [**NullablePatchedWritableCircuitGroupAssignmentRequestPriority**](PatchedWritableCircuitGroupAssignmentRequestPriority.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewTenancyContactAssignmentsCreateRequest

`func NewTenancyContactAssignmentsCreateRequest(objectType string, objectId int64, contact ContactAssignmentRequestContact, ) *TenancyContactAssignmentsCreateRequest`

NewTenancyContactAssignmentsCreateRequest instantiates a new TenancyContactAssignmentsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenancyContactAssignmentsCreateRequestWithDefaults

`func NewTenancyContactAssignmentsCreateRequestWithDefaults() *TenancyContactAssignmentsCreateRequest`

NewTenancyContactAssignmentsCreateRequestWithDefaults instantiates a new TenancyContactAssignmentsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *TenancyContactAssignmentsCreateRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TenancyContactAssignmentsCreateRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TenancyContactAssignmentsCreateRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *TenancyContactAssignmentsCreateRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *TenancyContactAssignmentsCreateRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *TenancyContactAssignmentsCreateRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetContact

`func (o *TenancyContactAssignmentsCreateRequest) GetContact() ContactAssignmentRequestContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *TenancyContactAssignmentsCreateRequest) GetContactOk() (*ContactAssignmentRequestContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *TenancyContactAssignmentsCreateRequest) SetContact(v ContactAssignmentRequestContact)`

SetContact sets Contact field to given value.


### GetRole

`func (o *TenancyContactAssignmentsCreateRequest) GetRole() ContactAssignmentRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TenancyContactAssignmentsCreateRequest) GetRoleOk() (*ContactAssignmentRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TenancyContactAssignmentsCreateRequest) SetRole(v ContactAssignmentRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *TenancyContactAssignmentsCreateRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *TenancyContactAssignmentsCreateRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *TenancyContactAssignmentsCreateRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPriority

`func (o *TenancyContactAssignmentsCreateRequest) GetPriority() PatchedWritableCircuitGroupAssignmentRequestPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TenancyContactAssignmentsCreateRequest) GetPriorityOk() (*PatchedWritableCircuitGroupAssignmentRequestPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TenancyContactAssignmentsCreateRequest) SetPriority(v PatchedWritableCircuitGroupAssignmentRequestPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TenancyContactAssignmentsCreateRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *TenancyContactAssignmentsCreateRequest) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *TenancyContactAssignmentsCreateRequest) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetTags

`func (o *TenancyContactAssignmentsCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TenancyContactAssignmentsCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TenancyContactAssignmentsCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TenancyContactAssignmentsCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *TenancyContactAssignmentsCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TenancyContactAssignmentsCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TenancyContactAssignmentsCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TenancyContactAssignmentsCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


