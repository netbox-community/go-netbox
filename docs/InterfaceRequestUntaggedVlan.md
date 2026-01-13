# InterfaceRequestUntaggedVlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vid** | **int32** | Numeric VLAN ID (1-4094) | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewInterfaceRequestUntaggedVlan

`func NewInterfaceRequestUntaggedVlan(vid int32, name string, ) *InterfaceRequestUntaggedVlan`

NewInterfaceRequestUntaggedVlan instantiates a new InterfaceRequestUntaggedVlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceRequestUntaggedVlanWithDefaults

`func NewInterfaceRequestUntaggedVlanWithDefaults() *InterfaceRequestUntaggedVlan`

NewInterfaceRequestUntaggedVlanWithDefaults instantiates a new InterfaceRequestUntaggedVlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVid

`func (o *InterfaceRequestUntaggedVlan) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *InterfaceRequestUntaggedVlan) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *InterfaceRequestUntaggedVlan) SetVid(v int32)`

SetVid sets Vid field to given value.


### GetName

`func (o *InterfaceRequestUntaggedVlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceRequestUntaggedVlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceRequestUntaggedVlan) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *InterfaceRequestUntaggedVlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterfaceRequestUntaggedVlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterfaceRequestUntaggedVlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterfaceRequestUntaggedVlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


