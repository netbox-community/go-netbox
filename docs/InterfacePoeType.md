# InterfacePoeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | * &#x60;type1-ieee802.3af&#x60; - 802.3af (Type 1) * &#x60;type2-ieee802.3at&#x60; - 802.3at (Type 2) * &#x60;type3-ieee802.3bt&#x60; - 802.3bt (Type 3) * &#x60;type4-ieee802.3bt&#x60; - 802.3bt (Type 4) * &#x60;passive-24v-2pair&#x60; - Passive 24V (2-pair) * &#x60;passive-24v-4pair&#x60; - Passive 24V (4-pair) * &#x60;passive-48v-2pair&#x60; - Passive 48V (2-pair) * &#x60;passive-48v-4pair&#x60; - Passive 48V (4-pair) | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewInterfacePoeType

`func NewInterfacePoeType() *InterfacePoeType`

NewInterfacePoeType instantiates a new InterfacePoeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfacePoeTypeWithDefaults

`func NewInterfacePoeTypeWithDefaults() *InterfacePoeType`

NewInterfacePoeTypeWithDefaults instantiates a new InterfacePoeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *InterfacePoeType) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InterfacePoeType) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InterfacePoeType) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InterfacePoeType) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *InterfacePoeType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InterfacePoeType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InterfacePoeType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InterfacePoeType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


