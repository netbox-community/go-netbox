# AvailablePrefix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | **int32** |  | [readonly] 
**Prefix** | **string** |  | [readonly] 
**Vrf** | [**NullableBriefVRF**](BriefVRF.md) |  | [readonly] 

## Methods

### NewAvailablePrefix

`func NewAvailablePrefix(family int32, prefix string, vrf NullableBriefVRF, ) *AvailablePrefix`

NewAvailablePrefix instantiates a new AvailablePrefix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailablePrefixWithDefaults

`func NewAvailablePrefixWithDefaults() *AvailablePrefix`

NewAvailablePrefixWithDefaults instantiates a new AvailablePrefix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *AvailablePrefix) GetFamily() int32`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *AvailablePrefix) GetFamilyOk() (*int32, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *AvailablePrefix) SetFamily(v int32)`

SetFamily sets Family field to given value.


### GetPrefix

`func (o *AvailablePrefix) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AvailablePrefix) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AvailablePrefix) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetVrf

`func (o *AvailablePrefix) GetVrf() BriefVRF`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *AvailablePrefix) GetVrfOk() (*BriefVRF, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *AvailablePrefix) SetVrf(v BriefVRF)`

SetVrf sets Vrf field to given value.


### SetVrfNil

`func (o *AvailablePrefix) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *AvailablePrefix) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


