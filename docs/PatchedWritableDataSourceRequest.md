# PatchedWritableDataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **interface{}** |  | [optional] 
**IgnoreRules** | Pointer to **string** | Patterns (one per line) matching files to ignore when syncing | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableDataSourceRequest

`func NewPatchedWritableDataSourceRequest() *PatchedWritableDataSourceRequest`

NewPatchedWritableDataSourceRequest instantiates a new PatchedWritableDataSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableDataSourceRequestWithDefaults

`func NewPatchedWritableDataSourceRequestWithDefaults() *PatchedWritableDataSourceRequest`

NewPatchedWritableDataSourceRequestWithDefaults instantiates a new PatchedWritableDataSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableDataSourceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableDataSourceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableDataSourceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableDataSourceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableDataSourceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableDataSourceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableDataSourceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableDataSourceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourceUrl

`func (o *PatchedWritableDataSourceRequest) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *PatchedWritableDataSourceRequest) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *PatchedWritableDataSourceRequest) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *PatchedWritableDataSourceRequest) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedWritableDataSourceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedWritableDataSourceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedWritableDataSourceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedWritableDataSourceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableDataSourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableDataSourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableDataSourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableDataSourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameters

`func (o *PatchedWritableDataSourceRequest) GetParameters() interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *PatchedWritableDataSourceRequest) GetParametersOk() (*interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *PatchedWritableDataSourceRequest) SetParameters(v interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *PatchedWritableDataSourceRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *PatchedWritableDataSourceRequest) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *PatchedWritableDataSourceRequest) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetIgnoreRules

`func (o *PatchedWritableDataSourceRequest) GetIgnoreRules() string`

GetIgnoreRules returns the IgnoreRules field if non-nil, zero value otherwise.

### GetIgnoreRulesOk

`func (o *PatchedWritableDataSourceRequest) GetIgnoreRulesOk() (*string, bool)`

GetIgnoreRulesOk returns a tuple with the IgnoreRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreRules

`func (o *PatchedWritableDataSourceRequest) SetIgnoreRules(v string)`

SetIgnoreRules sets IgnoreRules field to given value.

### HasIgnoreRules

`func (o *PatchedWritableDataSourceRequest) HasIgnoreRules() bool`

HasIgnoreRules returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableDataSourceRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableDataSourceRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableDataSourceRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableDataSourceRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableDataSourceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableDataSourceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableDataSourceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableDataSourceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


