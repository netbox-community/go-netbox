# WritableObjectPermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ObjectTypes** | **[]string** |  | 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 
**Actions** | **[]string** | The list of actions granted by this permission | 
**Constraints** | Pointer to **map[string]interface{}** | Queryset filter matching the applicable objects of the selected type(s) | [optional] 

## Methods

### NewWritableObjectPermissionRequest

`func NewWritableObjectPermissionRequest(name string, objectTypes []string, actions []string, ) *WritableObjectPermissionRequest`

NewWritableObjectPermissionRequest instantiates a new WritableObjectPermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableObjectPermissionRequestWithDefaults

`func NewWritableObjectPermissionRequestWithDefaults() *WritableObjectPermissionRequest`

NewWritableObjectPermissionRequestWithDefaults instantiates a new WritableObjectPermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableObjectPermissionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableObjectPermissionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableObjectPermissionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WritableObjectPermissionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableObjectPermissionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableObjectPermissionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableObjectPermissionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *WritableObjectPermissionRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WritableObjectPermissionRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WritableObjectPermissionRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WritableObjectPermissionRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectTypes

`func (o *WritableObjectPermissionRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *WritableObjectPermissionRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *WritableObjectPermissionRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetGroups

`func (o *WritableObjectPermissionRequest) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *WritableObjectPermissionRequest) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *WritableObjectPermissionRequest) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *WritableObjectPermissionRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *WritableObjectPermissionRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *WritableObjectPermissionRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *WritableObjectPermissionRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *WritableObjectPermissionRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetActions

`func (o *WritableObjectPermissionRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *WritableObjectPermissionRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *WritableObjectPermissionRequest) SetActions(v []string)`

SetActions sets Actions field to given value.


### GetConstraints

`func (o *WritableObjectPermissionRequest) GetConstraints() map[string]interface{}`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *WritableObjectPermissionRequest) GetConstraintsOk() (*map[string]interface{}, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *WritableObjectPermissionRequest) SetConstraints(v map[string]interface{})`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *WritableObjectPermissionRequest) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *WritableObjectPermissionRequest) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *WritableObjectPermissionRequest) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


