# PatchedSavedFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Parameters** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPatchedSavedFilterRequest

`func NewPatchedSavedFilterRequest() *PatchedSavedFilterRequest`

NewPatchedSavedFilterRequest instantiates a new PatchedSavedFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSavedFilterRequestWithDefaults

`func NewPatchedSavedFilterRequestWithDefaults() *PatchedSavedFilterRequest`

NewPatchedSavedFilterRequestWithDefaults instantiates a new PatchedSavedFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypes

`func (o *PatchedSavedFilterRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *PatchedSavedFilterRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *PatchedSavedFilterRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.

### HasObjectTypes

`func (o *PatchedSavedFilterRequest) HasObjectTypes() bool`

HasObjectTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedSavedFilterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSavedFilterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSavedFilterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSavedFilterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PatchedSavedFilterRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PatchedSavedFilterRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PatchedSavedFilterRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PatchedSavedFilterRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedSavedFilterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedSavedFilterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedSavedFilterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedSavedFilterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUser

`func (o *PatchedSavedFilterRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedSavedFilterRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedSavedFilterRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedSavedFilterRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PatchedSavedFilterRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PatchedSavedFilterRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetWeight

`func (o *PatchedSavedFilterRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedSavedFilterRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedSavedFilterRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedSavedFilterRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedSavedFilterRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedSavedFilterRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedSavedFilterRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedSavedFilterRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetShared

`func (o *PatchedSavedFilterRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *PatchedSavedFilterRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *PatchedSavedFilterRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *PatchedSavedFilterRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetParameters

`func (o *PatchedSavedFilterRequest) GetParameters() interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PatchedSavedFilterRequest) GetParametersOk() (*interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PatchedSavedFilterRequest) SetParameters(v interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PatchedSavedFilterRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *PatchedSavedFilterRequest) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *PatchedSavedFilterRequest) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


