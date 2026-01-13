# DataSourceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**DataSourceTypeValue**](DataSourceTypeValue.md) |  | [optional] 
**Label** | Pointer to [**DataSourceTypeLabel**](DataSourceTypeLabel.md) |  | [optional] 

## Methods

### NewDataSourceType

`func NewDataSourceType() *DataSourceType`

NewDataSourceType instantiates a new DataSourceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceTypeWithDefaults

`func NewDataSourceTypeWithDefaults() *DataSourceType`

NewDataSourceTypeWithDefaults instantiates a new DataSourceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DataSourceType) GetValue() DataSourceTypeValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataSourceType) GetValueOk() (*DataSourceTypeValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataSourceType) SetValue(v DataSourceTypeValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *DataSourceType) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *DataSourceType) GetLabel() DataSourceTypeLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DataSourceType) GetLabelOk() (*DataSourceTypeLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DataSourceType) SetLabel(v DataSourceTypeLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DataSourceType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


