# DataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**DataSourceTypeValue**](DataSourceTypeValue.md) |  | 
**SourceUrl** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **interface{}** |  | [optional] 
**IgnoreRules** | Pointer to **string** | Patterns (one per line) matching files to ignore when syncing | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDataSourceRequest

`func NewDataSourceRequest(name string, type_ DataSourceTypeValue, sourceUrl string, ) *DataSourceRequest`

NewDataSourceRequest instantiates a new DataSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceRequestWithDefaults

`func NewDataSourceRequestWithDefaults() *DataSourceRequest`

NewDataSourceRequestWithDefaults instantiates a new DataSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSourceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DataSourceRequest) GetType() DataSourceTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataSourceRequest) GetTypeOk() (*DataSourceTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataSourceRequest) SetType(v DataSourceTypeValue)`

SetType sets Type field to given value.


### GetSourceUrl

`func (o *DataSourceRequest) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *DataSourceRequest) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *DataSourceRequest) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetEnabled

`func (o *DataSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DataSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DataSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DataSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *DataSourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataSourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataSourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataSourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameters

`func (o *DataSourceRequest) GetParameters() interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DataSourceRequest) GetParametersOk() (*interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DataSourceRequest) SetParameters(v interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DataSourceRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *DataSourceRequest) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *DataSourceRequest) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetIgnoreRules

`func (o *DataSourceRequest) GetIgnoreRules() string`

GetIgnoreRules returns the IgnoreRules field if non-nil, zero value otherwise.

### GetIgnoreRulesOk

`func (o *DataSourceRequest) GetIgnoreRulesOk() (*string, bool)`

GetIgnoreRulesOk returns a tuple with the IgnoreRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreRules

`func (o *DataSourceRequest) SetIgnoreRules(v string)`

SetIgnoreRules sets IgnoreRules field to given value.

### HasIgnoreRules

`func (o *DataSourceRequest) HasIgnoreRules() bool`

HasIgnoreRules returns a boolean if a field has been set.

### GetComments

`func (o *DataSourceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DataSourceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DataSourceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DataSourceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCustomFields

`func (o *DataSourceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DataSourceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DataSourceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DataSourceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


