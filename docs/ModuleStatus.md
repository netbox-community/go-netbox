# ModuleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**ModuleStatusValue**](ModuleStatusValue.md) |  | [optional] 
**Label** | Pointer to [**ModuleStatusLabel**](ModuleStatusLabel.md) |  | [optional] 

## Methods

### NewModuleStatus

`func NewModuleStatus() *ModuleStatus`

NewModuleStatus instantiates a new ModuleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleStatusWithDefaults

`func NewModuleStatusWithDefaults() *ModuleStatus`

NewModuleStatusWithDefaults instantiates a new ModuleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ModuleStatus) GetValue() ModuleStatusValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModuleStatus) GetValueOk() (*ModuleStatusValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModuleStatus) SetValue(v ModuleStatusValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModuleStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *ModuleStatus) GetLabel() ModuleStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ModuleStatus) GetLabelOk() (*ModuleStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ModuleStatus) SetLabel(v ModuleStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ModuleStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


