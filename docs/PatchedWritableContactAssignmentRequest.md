# PatchedWritableContactAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**Contact** | Pointer to [**BriefContactRequest**](BriefContactRequest.md) |  | [optional] 
**Role** | Pointer to [**NullableBriefContactRoleRequest**](BriefContactRoleRequest.md) |  | [optional] 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriorityValue**](BriefCircuitGroupAssignmentSerializerPriorityValue.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableContactAssignmentRequest

`func NewPatchedWritableContactAssignmentRequest() *PatchedWritableContactAssignmentRequest`

NewPatchedWritableContactAssignmentRequest instantiates a new PatchedWritableContactAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableContactAssignmentRequestWithDefaults

`func NewPatchedWritableContactAssignmentRequestWithDefaults() *PatchedWritableContactAssignmentRequest`

NewPatchedWritableContactAssignmentRequestWithDefaults instantiates a new PatchedWritableContactAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *PatchedWritableContactAssignmentRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedWritableContactAssignmentRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedWritableContactAssignmentRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedWritableContactAssignmentRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedWritableContactAssignmentRequest) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedWritableContactAssignmentRequest) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedWritableContactAssignmentRequest) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedWritableContactAssignmentRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetContact

`func (o *PatchedWritableContactAssignmentRequest) GetContact() BriefContactRequest`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PatchedWritableContactAssignmentRequest) GetContactOk() (*BriefContactRequest, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PatchedWritableContactAssignmentRequest) SetContact(v BriefContactRequest)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PatchedWritableContactAssignmentRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableContactAssignmentRequest) GetRole() BriefContactRoleRequest`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableContactAssignmentRequest) GetRoleOk() (*BriefContactRoleRequest, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableContactAssignmentRequest) SetRole(v BriefContactRoleRequest)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableContactAssignmentRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedWritableContactAssignmentRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedWritableContactAssignmentRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPriority

`func (o *PatchedWritableContactAssignmentRequest) GetPriority() BriefCircuitGroupAssignmentSerializerPriorityValue`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchedWritableContactAssignmentRequest) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriorityValue, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchedWritableContactAssignmentRequest) SetPriority(v BriefCircuitGroupAssignmentSerializerPriorityValue)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchedWritableContactAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableContactAssignmentRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableContactAssignmentRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableContactAssignmentRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableContactAssignmentRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableContactAssignmentRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableContactAssignmentRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableContactAssignmentRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableContactAssignmentRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


