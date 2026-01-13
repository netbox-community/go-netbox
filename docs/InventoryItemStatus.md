# InventoryItemStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**InventoryItemStatusValue**](InventoryItemStatusValue.md) |  | [optional] 
**Label** | Pointer to [**InventoryItemStatusLabel**](InventoryItemStatusLabel.md) |  | [optional] 

## Methods

### NewInventoryItemStatus

`func NewInventoryItemStatus() *InventoryItemStatus`

NewInventoryItemStatus instantiates a new InventoryItemStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryItemStatusWithDefaults

`func NewInventoryItemStatusWithDefaults() *InventoryItemStatus`

NewInventoryItemStatusWithDefaults instantiates a new InventoryItemStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *InventoryItemStatus) GetValue() InventoryItemStatusValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InventoryItemStatus) GetValueOk() (*InventoryItemStatusValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InventoryItemStatus) SetValue(v InventoryItemStatusValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *InventoryItemStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *InventoryItemStatus) GetLabel() InventoryItemStatusLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InventoryItemStatus) GetLabelOk() (*InventoryItemStatusLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InventoryItemStatus) SetLabel(v InventoryItemStatusLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InventoryItemStatus) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


