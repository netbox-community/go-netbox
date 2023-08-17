# PatchedWritableL2VPNTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L2vpn** | Pointer to **int32** |  | [optional] 
**AssignedObjectType** | Pointer to **string** |  | [optional] 
**AssignedObjectId** | Pointer to **int64** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableL2VPNTerminationRequest

`func NewPatchedWritableL2VPNTerminationRequest() *PatchedWritableL2VPNTerminationRequest`

NewPatchedWritableL2VPNTerminationRequest instantiates a new PatchedWritableL2VPNTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableL2VPNTerminationRequestWithDefaults

`func NewPatchedWritableL2VPNTerminationRequestWithDefaults() *PatchedWritableL2VPNTerminationRequest`

NewPatchedWritableL2VPNTerminationRequestWithDefaults instantiates a new PatchedWritableL2VPNTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL2vpn

`func (o *PatchedWritableL2VPNTerminationRequest) GetL2vpn() int32`

GetL2vpn returns the L2vpn field if non-nil, zero value otherwise.

### GetL2vpnOk

`func (o *PatchedWritableL2VPNTerminationRequest) GetL2vpnOk() (*int32, bool)`

GetL2vpnOk returns a tuple with the L2vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2vpn

`func (o *PatchedWritableL2VPNTerminationRequest) SetL2vpn(v int32)`

SetL2vpn sets L2vpn field to given value.

### HasL2vpn

`func (o *PatchedWritableL2VPNTerminationRequest) HasL2vpn() bool`

HasL2vpn returns a boolean if a field has been set.

### GetAssignedObjectType

`func (o *PatchedWritableL2VPNTerminationRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *PatchedWritableL2VPNTerminationRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *PatchedWritableL2VPNTerminationRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *PatchedWritableL2VPNTerminationRequest) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### GetAssignedObjectId

`func (o *PatchedWritableL2VPNTerminationRequest) GetAssignedObjectId() int64`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *PatchedWritableL2VPNTerminationRequest) GetAssignedObjectIdOk() (*int64, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *PatchedWritableL2VPNTerminationRequest) SetAssignedObjectId(v int64)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *PatchedWritableL2VPNTerminationRequest) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableL2VPNTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableL2VPNTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableL2VPNTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableL2VPNTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableL2VPNTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableL2VPNTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableL2VPNTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableL2VPNTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


