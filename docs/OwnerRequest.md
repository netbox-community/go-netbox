# OwnerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Group** | Pointer to [**NullableOwnerRequestGroup**](OwnerRequestGroup.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**UserGroups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewOwnerRequest

`func NewOwnerRequest(name string, ) *OwnerRequest`

NewOwnerRequest instantiates a new OwnerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerRequestWithDefaults

`func NewOwnerRequestWithDefaults() *OwnerRequest`

NewOwnerRequestWithDefaults instantiates a new OwnerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OwnerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OwnerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OwnerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetGroup

`func (o *OwnerRequest) GetGroup() OwnerRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *OwnerRequest) GetGroupOk() (*OwnerRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *OwnerRequest) SetGroup(v OwnerRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *OwnerRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *OwnerRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *OwnerRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetDescription

`func (o *OwnerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OwnerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OwnerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OwnerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUserGroups

`func (o *OwnerRequest) GetUserGroups() []int32`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *OwnerRequest) GetUserGroupsOk() (*[]int32, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *OwnerRequest) SetUserGroups(v []int32)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *OwnerRequest) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetUsers

`func (o *OwnerRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *OwnerRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *OwnerRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *OwnerRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


