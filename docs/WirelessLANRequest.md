# WirelessLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ssid** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**NullableBriefWirelessLANGroupRequest**](BriefWirelessLANGroupRequest.md) |  | [optional] 
**Status** | Pointer to [**WirelessLANStatusValue**](WirelessLANStatusValue.md) |  | [optional] 
**Vlan** | Pointer to [**NullableBriefVLANRequest**](BriefVLANRequest.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**AuthType** | Pointer to [**WirelessLANAuthTypeValue**](WirelessLANAuthTypeValue.md) |  | [optional] 
**AuthCipher** | Pointer to [**WirelessLANAuthCipherValue**](WirelessLANAuthCipherValue.md) |  | [optional] 
**AuthPsk** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWirelessLANRequest

`func NewWirelessLANRequest(ssid string, ) *WirelessLANRequest`

NewWirelessLANRequest instantiates a new WirelessLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessLANRequestWithDefaults

`func NewWirelessLANRequestWithDefaults() *WirelessLANRequest`

NewWirelessLANRequestWithDefaults instantiates a new WirelessLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsid

`func (o *WirelessLANRequest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WirelessLANRequest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WirelessLANRequest) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetDescription

`func (o *WirelessLANRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WirelessLANRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WirelessLANRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WirelessLANRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *WirelessLANRequest) GetGroup() BriefWirelessLANGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WirelessLANRequest) GetGroupOk() (*BriefWirelessLANGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WirelessLANRequest) SetGroup(v BriefWirelessLANGroupRequest)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *WirelessLANRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *WirelessLANRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *WirelessLANRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetStatus

`func (o *WirelessLANRequest) GetStatus() WirelessLANStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WirelessLANRequest) GetStatusOk() (*WirelessLANStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WirelessLANRequest) SetStatus(v WirelessLANStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WirelessLANRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVlan

`func (o *WirelessLANRequest) GetVlan() BriefVLANRequest`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *WirelessLANRequest) GetVlanOk() (*BriefVLANRequest, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *WirelessLANRequest) SetVlan(v BriefVLANRequest)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *WirelessLANRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *WirelessLANRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *WirelessLANRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetTenant

`func (o *WirelessLANRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WirelessLANRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WirelessLANRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WirelessLANRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WirelessLANRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WirelessLANRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetAuthType

`func (o *WirelessLANRequest) GetAuthType() WirelessLANAuthTypeValue`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *WirelessLANRequest) GetAuthTypeOk() (*WirelessLANAuthTypeValue, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *WirelessLANRequest) SetAuthType(v WirelessLANAuthTypeValue)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *WirelessLANRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthCipher

`func (o *WirelessLANRequest) GetAuthCipher() WirelessLANAuthCipherValue`

GetAuthCipher returns the AuthCipher field if non-nil, zero value otherwise.

### GetAuthCipherOk

`func (o *WirelessLANRequest) GetAuthCipherOk() (*WirelessLANAuthCipherValue, bool)`

GetAuthCipherOk returns a tuple with the AuthCipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCipher

`func (o *WirelessLANRequest) SetAuthCipher(v WirelessLANAuthCipherValue)`

SetAuthCipher sets AuthCipher field to given value.

### HasAuthCipher

`func (o *WirelessLANRequest) HasAuthCipher() bool`

HasAuthCipher returns a boolean if a field has been set.

### GetAuthPsk

`func (o *WirelessLANRequest) GetAuthPsk() string`

GetAuthPsk returns the AuthPsk field if non-nil, zero value otherwise.

### GetAuthPskOk

`func (o *WirelessLANRequest) GetAuthPskOk() (*string, bool)`

GetAuthPskOk returns a tuple with the AuthPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPsk

`func (o *WirelessLANRequest) SetAuthPsk(v string)`

SetAuthPsk sets AuthPsk field to given value.

### HasAuthPsk

`func (o *WirelessLANRequest) HasAuthPsk() bool`

HasAuthPsk returns a boolean if a field has been set.

### GetComments

`func (o *WirelessLANRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WirelessLANRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WirelessLANRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WirelessLANRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WirelessLANRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WirelessLANRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WirelessLANRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WirelessLANRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WirelessLANRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WirelessLANRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WirelessLANRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WirelessLANRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


