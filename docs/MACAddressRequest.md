# MACAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacAddress** | **string** |  | 
**AssignedObjectType** | Pointer to **NullableString** |  | [optional] 
**AssignedObjectId** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMACAddressRequest

`func NewMACAddressRequest(macAddress string, ) *MACAddressRequest`

NewMACAddressRequest instantiates a new MACAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMACAddressRequestWithDefaults

`func NewMACAddressRequestWithDefaults() *MACAddressRequest`

NewMACAddressRequestWithDefaults instantiates a new MACAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacAddress

`func (o *MACAddressRequest) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *MACAddressRequest) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *MACAddressRequest) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetAssignedObjectType

`func (o *MACAddressRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *MACAddressRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *MACAddressRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *MACAddressRequest) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### SetAssignedObjectTypeNil

`func (o *MACAddressRequest) SetAssignedObjectTypeNil(b bool)`

 SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil

### UnsetAssignedObjectType
`func (o *MACAddressRequest) UnsetAssignedObjectType()`

UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
### GetAssignedObjectId

`func (o *MACAddressRequest) GetAssignedObjectId() int64`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *MACAddressRequest) GetAssignedObjectIdOk() (*int64, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *MACAddressRequest) SetAssignedObjectId(v int64)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *MACAddressRequest) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### SetAssignedObjectIdNil

`func (o *MACAddressRequest) SetAssignedObjectIdNil(b bool)`

 SetAssignedObjectIdNil sets the value for AssignedObjectId to be an explicit nil

### UnsetAssignedObjectId
`func (o *MACAddressRequest) UnsetAssignedObjectId()`

UnsetAssignedObjectId ensures that no value is present for AssignedObjectId, not even an explicit nil
### GetDescription

`func (o *MACAddressRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MACAddressRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MACAddressRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MACAddressRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *MACAddressRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *MACAddressRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *MACAddressRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *MACAddressRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *MACAddressRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MACAddressRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MACAddressRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MACAddressRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *MACAddressRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *MACAddressRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *MACAddressRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *MACAddressRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


