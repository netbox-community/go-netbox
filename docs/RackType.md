# RackType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;2-post-frame&#x60; - 2-post frame * &#x60;4-post-frame&#x60; - 4-post frame * &#x60;4-post-cabinet&#x60; - 4-post cabinet * &#x60;wall-frame&#x60; - Wall-mounted frame * &#x60;wall-frame-vertical&#x60; - Wall-mounted frame (vertical) * &#x60;wall-cabinet&#x60; - Wall-mounted cabinet * &#x60;wall-cabinet-vertical&#x60; - Wall-mounted cabinet (vertical) | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewRackType

`func NewRackType() *RackType`

NewRackType instantiates a new RackType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRackTypeWithDefaults

`func NewRackTypeWithDefaults() *RackType`

NewRackTypeWithDefaults instantiates a new RackType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RackType) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RackType) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RackType) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RackType) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *RackType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RackType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RackType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RackType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


