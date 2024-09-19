# JournalEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**AssignedObjectType** | **string** |  | 
**AssignedObjectId** | **int64** |  | 
**AssignedObject** | **interface{}** |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**CreatedBy** | Pointer to **NullableInt32** |  | [optional] 
**Kind** | Pointer to [**JournalEntryKind**](JournalEntryKind.md) |  | [optional] 
**Comments** | **string** |  | 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewJournalEntry

`func NewJournalEntry(id int32, url string, displayUrl string, display string, assignedObjectType string, assignedObjectId int64, assignedObject interface{}, created NullableTime, comments string, lastUpdated NullableTime, ) *JournalEntry`

NewJournalEntry instantiates a new JournalEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalEntryWithDefaults

`func NewJournalEntryWithDefaults() *JournalEntry`

NewJournalEntryWithDefaults instantiates a new JournalEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JournalEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalEntry) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *JournalEntry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JournalEntry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JournalEntry) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *JournalEntry) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *JournalEntry) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *JournalEntry) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *JournalEntry) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *JournalEntry) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *JournalEntry) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetAssignedObjectType

`func (o *JournalEntry) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *JournalEntry) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *JournalEntry) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObjectId

`func (o *JournalEntry) GetAssignedObjectId() int64`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *JournalEntry) GetAssignedObjectIdOk() (*int64, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *JournalEntry) SetAssignedObjectId(v int64)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetAssignedObject

`func (o *JournalEntry) GetAssignedObject() interface{}`

GetAssignedObject returns the AssignedObject field if non-nil, zero value otherwise.

### GetAssignedObjectOk

`func (o *JournalEntry) GetAssignedObjectOk() (*interface{}, bool)`

GetAssignedObjectOk returns a tuple with the AssignedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObject

`func (o *JournalEntry) SetAssignedObject(v interface{})`

SetAssignedObject sets AssignedObject field to given value.


### SetAssignedObjectNil

`func (o *JournalEntry) SetAssignedObjectNil(b bool)`

 SetAssignedObjectNil sets the value for AssignedObject to be an explicit nil

### UnsetAssignedObject
`func (o *JournalEntry) UnsetAssignedObject()`

UnsetAssignedObject ensures that no value is present for AssignedObject, not even an explicit nil
### GetCreated

`func (o *JournalEntry) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JournalEntry) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JournalEntry) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *JournalEntry) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *JournalEntry) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetCreatedBy

`func (o *JournalEntry) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *JournalEntry) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *JournalEntry) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *JournalEntry) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *JournalEntry) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *JournalEntry) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetKind

`func (o *JournalEntry) GetKind() JournalEntryKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JournalEntry) GetKindOk() (*JournalEntryKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JournalEntry) SetKind(v JournalEntryKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *JournalEntry) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetComments

`func (o *JournalEntry) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *JournalEntry) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *JournalEntry) SetComments(v string)`

SetComments sets Comments field to given value.


### GetTags

`func (o *JournalEntry) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *JournalEntry) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *JournalEntry) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *JournalEntry) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *JournalEntry) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *JournalEntry) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *JournalEntry) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *JournalEntry) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetLastUpdated

`func (o *JournalEntry) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *JournalEntry) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *JournalEntry) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *JournalEntry) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *JournalEntry) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


