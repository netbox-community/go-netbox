# PatchedOwnerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**NullableOwnerRequestGroup**](OwnerRequestGroup.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**UserGroups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewPatchedOwnerRequest

`func NewPatchedOwnerRequest() *PatchedOwnerRequest`

NewPatchedOwnerRequest instantiates a new PatchedOwnerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedOwnerRequestWithDefaults

`func NewPatchedOwnerRequestWithDefaults() *PatchedOwnerRequest`

NewPatchedOwnerRequestWithDefaults instantiates a new PatchedOwnerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedOwnerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedOwnerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedOwnerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedOwnerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGroup

`func (o *PatchedOwnerRequest) GetGroup() OwnerRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PatchedOwnerRequest) GetGroupOk() (*OwnerRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PatchedOwnerRequest) SetGroup(v OwnerRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PatchedOwnerRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *PatchedOwnerRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *PatchedOwnerRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetDescription

`func (o *PatchedOwnerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedOwnerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedOwnerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedOwnerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUserGroups

`func (o *PatchedOwnerRequest) GetUserGroups() []int32`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *PatchedOwnerRequest) GetUserGroupsOk() (*[]int32, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *PatchedOwnerRequest) SetUserGroups(v []int32)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *PatchedOwnerRequest) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetUsers

`func (o *PatchedOwnerRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PatchedOwnerRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PatchedOwnerRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PatchedOwnerRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


