# AvailableIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | **int32** |  | [readonly] 
**Address** | **string** |  | [readonly] 
**Vrf** | [**NullableBriefVRF**](BriefVRF.md) |  | [readonly] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAvailableIP

`func NewAvailableIP(family int32, address string, vrf NullableBriefVRF, ) *AvailableIP`

NewAvailableIP instantiates a new AvailableIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableIPWithDefaults

`func NewAvailableIPWithDefaults() *AvailableIP`

NewAvailableIPWithDefaults instantiates a new AvailableIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *AvailableIP) GetFamily() int32`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *AvailableIP) GetFamilyOk() (*int32, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *AvailableIP) SetFamily(v int32)`

SetFamily sets Family field to given value.


### GetAddress

`func (o *AvailableIP) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AvailableIP) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AvailableIP) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetVrf

`func (o *AvailableIP) GetVrf() BriefVRF`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *AvailableIP) GetVrfOk() (*BriefVRF, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *AvailableIP) SetVrf(v BriefVRF)`

SetVrf sets Vrf field to given value.


### SetVrfNil

`func (o *AvailableIP) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *AvailableIP) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetDescription

`func (o *AvailableIP) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AvailableIP) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AvailableIP) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AvailableIP) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


