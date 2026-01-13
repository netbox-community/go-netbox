# ExtrasNotificationGroupsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewExtrasNotificationGroupsCreateRequest

`func NewExtrasNotificationGroupsCreateRequest(name string, ) *ExtrasNotificationGroupsCreateRequest`

NewExtrasNotificationGroupsCreateRequest instantiates a new ExtrasNotificationGroupsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtrasNotificationGroupsCreateRequestWithDefaults

`func NewExtrasNotificationGroupsCreateRequestWithDefaults() *ExtrasNotificationGroupsCreateRequest`

NewExtrasNotificationGroupsCreateRequestWithDefaults instantiates a new ExtrasNotificationGroupsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtrasNotificationGroupsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtrasNotificationGroupsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtrasNotificationGroupsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExtrasNotificationGroupsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtrasNotificationGroupsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtrasNotificationGroupsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtrasNotificationGroupsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroups

`func (o *ExtrasNotificationGroupsCreateRequest) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ExtrasNotificationGroupsCreateRequest) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ExtrasNotificationGroupsCreateRequest) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ExtrasNotificationGroupsCreateRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *ExtrasNotificationGroupsCreateRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ExtrasNotificationGroupsCreateRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ExtrasNotificationGroupsCreateRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ExtrasNotificationGroupsCreateRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


