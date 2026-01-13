# PatchedTableConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **string** |  | [optional] 
**Table** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Columns** | Pointer to **[]string** |  | [optional] 
**Ordering** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPatchedTableConfigRequest

`func NewPatchedTableConfigRequest() *PatchedTableConfigRequest`

NewPatchedTableConfigRequest instantiates a new PatchedTableConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedTableConfigRequestWithDefaults

`func NewPatchedTableConfigRequestWithDefaults() *PatchedTableConfigRequest`

NewPatchedTableConfigRequestWithDefaults instantiates a new PatchedTableConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *PatchedTableConfigRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PatchedTableConfigRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PatchedTableConfigRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *PatchedTableConfigRequest) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetTable

`func (o *PatchedTableConfigRequest) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *PatchedTableConfigRequest) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *PatchedTableConfigRequest) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *PatchedTableConfigRequest) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetName

`func (o *PatchedTableConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedTableConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedTableConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedTableConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedTableConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedTableConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedTableConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedTableConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUser

`func (o *PatchedTableConfigRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedTableConfigRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedTableConfigRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedTableConfigRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PatchedTableConfigRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PatchedTableConfigRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetWeight

`func (o *PatchedTableConfigRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedTableConfigRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedTableConfigRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedTableConfigRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedTableConfigRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedTableConfigRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedTableConfigRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedTableConfigRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetShared

`func (o *PatchedTableConfigRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *PatchedTableConfigRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *PatchedTableConfigRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *PatchedTableConfigRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetColumns

`func (o *PatchedTableConfigRequest) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *PatchedTableConfigRequest) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *PatchedTableConfigRequest) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *PatchedTableConfigRequest) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetOrdering

`func (o *PatchedTableConfigRequest) GetOrdering() []string`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *PatchedTableConfigRequest) GetOrderingOk() (*[]string, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *PatchedTableConfigRequest) SetOrdering(v []string)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *PatchedTableConfigRequest) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.

### SetOrderingNil

`func (o *PatchedTableConfigRequest) SetOrderingNil(b bool)`

 SetOrderingNil sets the value for Ordering to be an explicit nil

### UnsetOrdering
`func (o *PatchedTableConfigRequest) UnsetOrdering()`

UnsetOrdering ensures that no value is present for Ordering, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


