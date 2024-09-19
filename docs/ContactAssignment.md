# ContactAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**ObjectType** | **string** |  | 
**ObjectId** | **int64** |  | 
**Object** | **map[string]interface{}** |  | [readonly] 
**Contact** | [**BriefContact**](BriefContact.md) |  | 
**Role** | Pointer to [**NullableBriefContactRole**](BriefContactRole.md) |  | [optional] 
**Priority** | Pointer to [**BriefCircuitGroupAssignmentSerializerPriority**](BriefCircuitGroupAssignmentSerializerPriority.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewContactAssignment

`func NewContactAssignment(id int32, url string, display string, objectType string, objectId int64, object map[string]interface{}, contact BriefContact, created NullableTime, lastUpdated NullableTime, ) *ContactAssignment`

NewContactAssignment instantiates a new ContactAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactAssignmentWithDefaults

`func NewContactAssignmentWithDefaults() *ContactAssignment`

NewContactAssignmentWithDefaults instantiates a new ContactAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactAssignment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactAssignment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactAssignment) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ContactAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContactAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContactAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *ContactAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ContactAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ContactAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetObjectType

`func (o *ContactAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ContactAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ContactAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetObjectId

`func (o *ContactAssignment) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ContactAssignment) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ContactAssignment) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.


### GetObject

`func (o *ContactAssignment) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ContactAssignment) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ContactAssignment) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.


### GetContact

`func (o *ContactAssignment) GetContact() BriefContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ContactAssignment) GetContactOk() (*BriefContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ContactAssignment) SetContact(v BriefContact)`

SetContact sets Contact field to given value.


### GetRole

`func (o *ContactAssignment) GetRole() BriefContactRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ContactAssignment) GetRoleOk() (*BriefContactRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ContactAssignment) SetRole(v BriefContactRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ContactAssignment) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ContactAssignment) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ContactAssignment) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetPriority

`func (o *ContactAssignment) GetPriority() BriefCircuitGroupAssignmentSerializerPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ContactAssignment) GetPriorityOk() (*BriefCircuitGroupAssignmentSerializerPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ContactAssignment) SetPriority(v BriefCircuitGroupAssignmentSerializerPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ContactAssignment) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTags

`func (o *ContactAssignment) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContactAssignment) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContactAssignment) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContactAssignment) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ContactAssignment) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ContactAssignment) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ContactAssignment) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ContactAssignment) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *ContactAssignment) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ContactAssignment) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ContactAssignment) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ContactAssignment) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ContactAssignment) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ContactAssignment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ContactAssignment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ContactAssignment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ContactAssignment) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ContactAssignment) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


