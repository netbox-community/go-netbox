# ExtrasJournalEntriesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedObjectType** | **string** |  | 
**AssignedObjectId** | **int64** |  | 
**CreatedBy** | Pointer to **NullableInt32** |  | [optional] 
**Kind** | Pointer to [**JournalEntryKindValue**](JournalEntryKindValue.md) |  | [optional] 
**Comments** | **string** |  | 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewExtrasJournalEntriesCreateRequest

`func NewExtrasJournalEntriesCreateRequest(assignedObjectType string, assignedObjectId int64, comments string, ) *ExtrasJournalEntriesCreateRequest`

NewExtrasJournalEntriesCreateRequest instantiates a new ExtrasJournalEntriesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtrasJournalEntriesCreateRequestWithDefaults

`func NewExtrasJournalEntriesCreateRequestWithDefaults() *ExtrasJournalEntriesCreateRequest`

NewExtrasJournalEntriesCreateRequestWithDefaults instantiates a new ExtrasJournalEntriesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedObjectType

`func (o *ExtrasJournalEntriesCreateRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *ExtrasJournalEntriesCreateRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *ExtrasJournalEntriesCreateRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObjectId

`func (o *ExtrasJournalEntriesCreateRequest) GetAssignedObjectId() int64`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *ExtrasJournalEntriesCreateRequest) GetAssignedObjectIdOk() (*int64, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *ExtrasJournalEntriesCreateRequest) SetAssignedObjectId(v int64)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetCreatedBy

`func (o *ExtrasJournalEntriesCreateRequest) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ExtrasJournalEntriesCreateRequest) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ExtrasJournalEntriesCreateRequest) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ExtrasJournalEntriesCreateRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ExtrasJournalEntriesCreateRequest) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ExtrasJournalEntriesCreateRequest) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetKind

`func (o *ExtrasJournalEntriesCreateRequest) GetKind() JournalEntryKindValue`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ExtrasJournalEntriesCreateRequest) GetKindOk() (*JournalEntryKindValue, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ExtrasJournalEntriesCreateRequest) SetKind(v JournalEntryKindValue)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ExtrasJournalEntriesCreateRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetComments

`func (o *ExtrasJournalEntriesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ExtrasJournalEntriesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ExtrasJournalEntriesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.


### GetTags

`func (o *ExtrasJournalEntriesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExtrasJournalEntriesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExtrasJournalEntriesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExtrasJournalEntriesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ExtrasJournalEntriesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ExtrasJournalEntriesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ExtrasJournalEntriesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ExtrasJournalEntriesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


