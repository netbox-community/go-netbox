# NotificationGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewNotificationGroup

`func NewNotificationGroup(id int32, url string, display string, displayUrl string, name string, ) *NotificationGroup`

NewNotificationGroup instantiates a new NotificationGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationGroupWithDefaults

`func NewNotificationGroupWithDefaults() *NotificationGroup`

NewNotificationGroupWithDefaults instantiates a new NotificationGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationGroup) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *NotificationGroup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationGroup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationGroup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *NotificationGroup) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *NotificationGroup) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *NotificationGroup) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDisplayUrl

`func (o *NotificationGroup) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *NotificationGroup) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *NotificationGroup) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetName

`func (o *NotificationGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NotificationGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroups

`func (o *NotificationGroup) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *NotificationGroup) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *NotificationGroup) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *NotificationGroup) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *NotificationGroup) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *NotificationGroup) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *NotificationGroup) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *NotificationGroup) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


