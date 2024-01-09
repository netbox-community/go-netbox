# LocationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**LocationStatusValue**](LocationStatusValue.md) |  | [optional] 
**Label** | Pointer to [**LocationStatusLabel**](LocationStatusLabel.md) |  | [optional] 

## Methods

### NewLocationStatus

`func NewLocationStatus() *LocationStatus`

NewLocationStatus instantiates a new LocationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationStatusWithDefaults

`func NewLocationStatusWithDefaults() *LocationStatus`

NewLocationStatusWithDefaults instantiates a new LocationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *LocationStatus) GetValue() LocationStatusValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LocationStatus) GetValueOk() (*LocationStatusValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LocationStatus) SetValue(v LocationStatusValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *LocationStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *LocationStatus) GetLabel() LocationStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LocationStatus) GetLabelOk() (*LocationStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LocationStatus) SetLabel(v LocationStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LocationStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


