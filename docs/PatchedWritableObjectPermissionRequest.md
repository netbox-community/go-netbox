# PatchedWritableObjectPermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**Groups** | Pointer to **[]int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 
**Actions** | Pointer to **[]string** | The list of actions granted by this permission | [optional] 
**Constraints** | Pointer to **map[string]interface{}** | Queryset filter matching the applicable objects of the selected type(s) | [optional] 

## Methods

### NewPatchedWritableObjectPermissionRequest

`func NewPatchedWritableObjectPermissionRequest() *PatchedWritableObjectPermissionRequest`

NewPatchedWritableObjectPermissionRequest instantiates a new PatchedWritableObjectPermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableObjectPermissionRequestWithDefaults

`func NewPatchedWritableObjectPermissionRequestWithDefaults() *PatchedWritableObjectPermissionRequest`

NewPatchedWritableObjectPermissionRequestWithDefaults instantiates a new PatchedWritableObjectPermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableObjectPermissionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableObjectPermissionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableObjectPermissionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableObjectPermissionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableObjectPermissionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableObjectPermissionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableObjectPermissionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableObjectPermissionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedWritableObjectPermissionRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedWritableObjectPermissionRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedWritableObjectPermissionRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedWritableObjectPermissionRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetObjectTypes

`func (o *PatchedWritableObjectPermissionRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PatchedWritableObjectPermissionRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PatchedWritableObjectPermissionRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PatchedWritableObjectPermissionRequest) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetGroups

`func (o *PatchedWritableObjectPermissionRequest) GetGroups() []int32`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PatchedWritableObjectPermissionRequest) GetGroupsOk() (*[]int32, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PatchedWritableObjectPermissionRequest) SetGroups(v []int32)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PatchedWritableObjectPermissionRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *PatchedWritableObjectPermissionRequest) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PatchedWritableObjectPermissionRequest) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PatchedWritableObjectPermissionRequest) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PatchedWritableObjectPermissionRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetActions

`func (o *PatchedWritableObjectPermissionRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PatchedWritableObjectPermissionRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PatchedWritableObjectPermissionRequest) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PatchedWritableObjectPermissionRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConstraints

`func (o *PatchedWritableObjectPermissionRequest) GetConstraints() map[string]interface{}`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *PatchedWritableObjectPermissionRequest) GetConstraintsOk() (*map[string]interface{}, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *PatchedWritableObjectPermissionRequest) SetConstraints(v map[string]interface{})`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *PatchedWritableObjectPermissionRequest) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *PatchedWritableObjectPermissionRequest) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *PatchedWritableObjectPermissionRequest) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


