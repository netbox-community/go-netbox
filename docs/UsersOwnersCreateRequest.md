# UsersOwnersCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Group** | Pointer to [**NullableOwnerRequestGroup**](OwnerRequestGroup.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**UserGroups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewUsersOwnersCreateRequest

`func NewUsersOwnersCreateRequest(name string, ) *UsersOwnersCreateRequest`

NewUsersOwnersCreateRequest instantiates a new UsersOwnersCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersOwnersCreateRequestWithDefaults

`func NewUsersOwnersCreateRequestWithDefaults() *UsersOwnersCreateRequest`

NewUsersOwnersCreateRequestWithDefaults instantiates a new UsersOwnersCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UsersOwnersCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsersOwnersCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsersOwnersCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetGroup

`func (o *UsersOwnersCreateRequest) GetGroup() OwnerRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UsersOwnersCreateRequest) GetGroupOk() (*OwnerRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UsersOwnersCreateRequest) SetGroup(v OwnerRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UsersOwnersCreateRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *UsersOwnersCreateRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *UsersOwnersCreateRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetDescription

`func (o *UsersOwnersCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UsersOwnersCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UsersOwnersCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UsersOwnersCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUserGroups

`func (o *UsersOwnersCreateRequest) GetUserGroups() []int32`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *UsersOwnersCreateRequest) GetUserGroupsOk() (*[]int32, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *UsersOwnersCreateRequest) SetUserGroups(v []int32)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *UsersOwnersCreateRequest) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetUsers

`func (o *UsersOwnersCreateRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersOwnersCreateRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersOwnersCreateRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UsersOwnersCreateRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


