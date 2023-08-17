# DataSourceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;new&#x60; - New * &#x60;queued&#x60; - Queued * &#x60;syncing&#x60; - Syncing * &#x60;completed&#x60; - Completed * &#x60;failed&#x60; - Failed | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewDataSourceStatus

`func NewDataSourceStatus() *DataSourceStatus`

NewDataSourceStatus instantiates a new DataSourceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceStatusWithDefaults

`func NewDataSourceStatusWithDefaults() *DataSourceStatus`

NewDataSourceStatusWithDefaults instantiates a new DataSourceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DataSourceStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataSourceStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataSourceStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DataSourceStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *DataSourceStatus) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DataSourceStatus) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DataSourceStatus) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DataSourceStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


