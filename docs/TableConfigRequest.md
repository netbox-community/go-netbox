# TableConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | **string** |  | 
**Table** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**User** | Pointer to **NullableInt32** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Columns** | **[]string** |  | 
**Ordering** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTableConfigRequest

`func NewTableConfigRequest(objectType string, table string, name string, columns []string, ) *TableConfigRequest`

NewTableConfigRequest instantiates a new TableConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableConfigRequestWithDefaults

`func NewTableConfigRequestWithDefaults() *TableConfigRequest`

NewTableConfigRequestWithDefaults instantiates a new TableConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *TableConfigRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TableConfigRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TableConfigRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTable

`func (o *TableConfigRequest) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *TableConfigRequest) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *TableConfigRequest) SetTable(v string)`

SetTable sets Table field to given value.


### GetName

`func (o *TableConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TableConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TableConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TableConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TableConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TableConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TableConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUser

`func (o *TableConfigRequest) GetUser() int32`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TableConfigRequest) GetUserOk() (*int32, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TableConfigRequest) SetUser(v int32)`

SetUser sets User field to given value.

### HasUser

`func (o *TableConfigRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *TableConfigRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *TableConfigRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetWeight

`func (o *TableConfigRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *TableConfigRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *TableConfigRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *TableConfigRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetEnabled

`func (o *TableConfigRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TableConfigRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TableConfigRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TableConfigRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetShared

`func (o *TableConfigRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *TableConfigRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *TableConfigRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *TableConfigRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetColumns

`func (o *TableConfigRequest) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TableConfigRequest) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TableConfigRequest) SetColumns(v []string)`

SetColumns sets Columns field to given value.


### GetOrdering

`func (o *TableConfigRequest) GetOrdering() []string`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *TableConfigRequest) GetOrderingOk() (*[]string, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *TableConfigRequest) SetOrdering(v []string)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *TableConfigRequest) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.

### SetOrderingNil

`func (o *TableConfigRequest) SetOrderingNil(b bool)`

 SetOrderingNil sets the value for Ordering to be an explicit nil

### UnsetOrdering
`func (o *TableConfigRequest) UnsetOrdering()`

UnsetOrdering ensures that no value is present for Ordering, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


