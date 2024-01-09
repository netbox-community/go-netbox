# ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**ClusterStatusValue**](ClusterStatusValue.md) |  | [optional] 
**Label** | Pointer to [**ClusterStatusLabel**](ClusterStatusLabel.md) |  | [optional] 

## Methods

### NewClusterStatus

`func NewClusterStatus() *ClusterStatus`

NewClusterStatus instantiates a new ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusWithDefaults

`func NewClusterStatusWithDefaults() *ClusterStatus`

NewClusterStatusWithDefaults instantiates a new ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ClusterStatus) GetValue() ClusterStatusValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClusterStatus) GetValueOk() (*ClusterStatusValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClusterStatus) SetValue(v ClusterStatusValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ClusterStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *ClusterStatus) GetLabel() ClusterStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ClusterStatus) GetLabelOk() (*ClusterStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ClusterStatus) SetLabel(v ClusterStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ClusterStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


