# ModuleTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Manufacturer** | [**BriefManufacturerRequest**](BriefManufacturerRequest.md) |  | 
**Model** | **string** |  | 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**Airflow** | Pointer to [**NullableModuleTypeRequestAirflow**](ModuleTypeRequestAirflow.md) |  | [optional] 
**Weight** | Pointer to **NullableFloat64** |  | [optional] 
**WeightUnit** | Pointer to [**NullableDeviceTypeRequestWeightUnit**](DeviceTypeRequestWeightUnit.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModuleTypeRequest

`func NewModuleTypeRequest(manufacturer BriefManufacturerRequest, model string, ) *ModuleTypeRequest`

NewModuleTypeRequest instantiates a new ModuleTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleTypeRequestWithDefaults

`func NewModuleTypeRequestWithDefaults() *ModuleTypeRequest`

NewModuleTypeRequestWithDefaults instantiates a new ModuleTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManufacturer

`func (o *ModuleTypeRequest) GetManufacturer() BriefManufacturerRequest`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ModuleTypeRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ModuleTypeRequest) SetManufacturer(v BriefManufacturerRequest)`

SetManufacturer sets Manufacturer field to given value.


### GetModel

`func (o *ModuleTypeRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModuleTypeRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModuleTypeRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetPartNumber

`func (o *ModuleTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *ModuleTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *ModuleTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *ModuleTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetAirflow

`func (o *ModuleTypeRequest) GetAirflow() ModuleTypeRequestAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *ModuleTypeRequest) GetAirflowOk() (*ModuleTypeRequestAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *ModuleTypeRequest) SetAirflow(v ModuleTypeRequestAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *ModuleTypeRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### SetAirflowNil

`func (o *ModuleTypeRequest) SetAirflowNil(b bool)`

 SetAirflowNil sets the value for Airflow to be an explicit nil

### UnsetAirflow
`func (o *ModuleTypeRequest) UnsetAirflow()`

UnsetAirflow ensures that no value is present for Airflow, not even an explicit nil
### GetWeight

`func (o *ModuleTypeRequest) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ModuleTypeRequest) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ModuleTypeRequest) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ModuleTypeRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *ModuleTypeRequest) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *ModuleTypeRequest) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetWeightUnit

`func (o *ModuleTypeRequest) GetWeightUnit() DeviceTypeRequestWeightUnit`

GetWeightUnit returns the WeightUnit field if non-nil, zero value otherwise.

### GetWeightUnitOk

`func (o *ModuleTypeRequest) GetWeightUnitOk() (*DeviceTypeRequestWeightUnit, bool)`

GetWeightUnitOk returns a tuple with the WeightUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightUnit

`func (o *ModuleTypeRequest) SetWeightUnit(v DeviceTypeRequestWeightUnit)`

SetWeightUnit sets WeightUnit field to given value.

### HasWeightUnit

`func (o *ModuleTypeRequest) HasWeightUnit() bool`

HasWeightUnit returns a boolean if a field has been set.

### SetWeightUnitNil

`func (o *ModuleTypeRequest) SetWeightUnitNil(b bool)`

 SetWeightUnitNil sets the value for WeightUnit to be an explicit nil

### UnsetWeightUnit
`func (o *ModuleTypeRequest) UnsetWeightUnit()`

UnsetWeightUnit ensures that no value is present for WeightUnit, not even an explicit nil
### GetDescription

`func (o *ModuleTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModuleTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModuleTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModuleTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *ModuleTypeRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ModuleTypeRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ModuleTypeRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ModuleTypeRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *ModuleTypeRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModuleTypeRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModuleTypeRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModuleTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ModuleTypeRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModuleTypeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModuleTypeRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModuleTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


