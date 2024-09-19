# ObjectPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ObjectTypes** | **[]string** |  | 
**Actions** | **[]string** | The list of actions granted by this permission | 
**Constraints** | Pointer to **interface{}** | Queryset filter matching the applicable objects of the selected type(s) | [optional] 
**Groups** | Pointer to [**[]NestedGroup**](NestedGroup.md) |  | [optional] 
**Users** | Pointer to [**[]NestedUser**](NestedUser.md) |  | [optional] 

## Methods

### NewObjectPermission

`func NewObjectPermission(id int32, url string, displayUrl string, display string, name string, objectTypes []string, actions []string, ) *ObjectPermission`

NewObjectPermission instantiates a new ObjectPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectPermissionWithDefaults

`func NewObjectPermissionWithDefaults() *ObjectPermission`

NewObjectPermissionWithDefaults instantiates a new ObjectPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectPermission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectPermission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectPermission) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ObjectPermission) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ObjectPermission) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ObjectPermission) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *ObjectPermission) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *ObjectPermission) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *ObjectPermission) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *ObjectPermission) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ObjectPermission) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ObjectPermission) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *ObjectPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectPermission) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ObjectPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ObjectPermission) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ObjectPermission) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ObjectPermission) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ObjectPermission) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectTypes

`func (o *ObjectPermission) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *ObjectPermission) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *ObjectPermission) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetActions

`func (o *ObjectPermission) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ObjectPermission) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ObjectPermission) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetConstraints

`func (o *ObjectPermission) GetConstraints() interface{}`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *ObjectPermission) GetConstraintsOk() (*interface{}, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *ObjectPermission) SetConstraints(v interface{})`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *ObjectPermission) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *ObjectPermission) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *ObjectPermission) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetGroups

`func (o *ObjectPermission) GetGroups() []NestedGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ObjectPermission) GetGroupsOk() (*[]NestedGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ObjectPermission) SetGroups(v []NestedGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ObjectPermission) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *ObjectPermission) GetUsers() []NestedUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ObjectPermission) GetUsersOk() (*[]NestedUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ObjectPermission) SetUsers(v []NestedUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ObjectPermission) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


