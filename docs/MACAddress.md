# MACAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | Pointer to **string** |  | [optional] [readonly] 
**Display** | **string** |  | [readonly] 
**MacAddress** | **string** |  | 
**AssignedObjectType** | Pointer to **NullableString** |  | [optional] 
**AssignedObjectId** | Pointer to **NullableInt64** |  | [optional] 
**AssignedObject** | Pointer to **interface{}** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewMACAddress

`func NewMACAddress(id int32, url string, display string, macAddress string, ) *MACAddress`

NewMACAddress instantiates a new MACAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMACAddressWithDefaults

`func NewMACAddressWithDefaults() *MACAddress`

NewMACAddressWithDefaults instantiates a new MACAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MACAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MACAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MACAddress) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *MACAddress) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MACAddress) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MACAddress) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *MACAddress) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *MACAddress) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *MACAddress) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.

### HasDisplayUrl

`func (o *MACAddress) HasDisplayUrl() bool`

HasDisplayUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *MACAddress) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *MACAddress) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *MACAddress) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetMacAddress

`func (o *MACAddress) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *MACAddress) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *MACAddress) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetAssignedObjectType

`func (o *MACAddress) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *MACAddress) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *MACAddress) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *MACAddress) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### SetAssignedObjectTypeNil

`func (o *MACAddress) SetAssignedObjectTypeNil(b bool)`

 SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil

### UnsetAssignedObjectType
`func (o *MACAddress) UnsetAssignedObjectType()`

UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
### GetAssignedObjectId

`func (o *MACAddress) GetAssignedObjectId() int64`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *MACAddress) GetAssignedObjectIdOk() (*int64, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *MACAddress) SetAssignedObjectId(v int64)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *MACAddress) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### SetAssignedObjectIdNil

`func (o *MACAddress) SetAssignedObjectIdNil(b bool)`

 SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil

### UnsetAssignedObjectId
`func (o *MACAddress) UnsetAssignedObjectId()`

UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
### GetAssignedObject

`func (o *MACAddress) GetAssignedObject() interface{}`

GetAssignedObject returns the AssignedObject field if non-nil, zero value otherwise.

### GetAssignedObjectOk

`func (o *MACAddress) GetAssignedObjectOk() (*interface{}, bool)`

GetAssignedObjectOk returns a tuple with the AssignedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObject

`func (o *MACAddress) SetAssignedObject(v interface{})`

SetAssignedObject sets AssignedObject field to given value.

### HasAssignedObject

`func (o *MACAddress) HasAssignedObject() bool`

HasAssignedObject returns a boolean if a field has been set.

### SetAssignedObjectNil

`func (o *MACAddress) SetAssignedObjectNil(b bool)`

 SetAssignedObjectNil sets the value for AssignedObject to be an explicit nil

### UnsetAssignedObject
`func (o *MACAddress) UnsetAssignedObject()`

UnsetAssignedObject ensures that no value is present for AssignedObject, not even an explicit nil
### GetDescription

`func (o *MACAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MACAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MACAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MACAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *MACAddress) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *MACAddress) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *MACAddress) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *MACAddress) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *MACAddress) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MACAddress) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MACAddress) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MACAddress) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *MACAddress) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MACAddress) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MACAddress) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MACAddress) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *MACAddress) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MACAddress) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MACAddress) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MACAddress) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *MACAddress) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *MACAddress) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *MACAddress) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *MACAddress) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *MACAddress) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *MACAddress) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *MACAddress) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *MACAddress) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


