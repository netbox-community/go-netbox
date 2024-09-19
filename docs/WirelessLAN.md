# WirelessLAN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Ssid** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**NullableBriefWirelessLANGroup**](BriefWirelessLANGroup.md) |  | [optional] 
**Status** | Pointer to [**WirelessLANStatus**](WirelessLANStatus.md) |  | [optional] 
**Vlan** | Pointer to [**NullableBriefVLAN**](BriefVLAN.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenant**](BriefTenant.md) |  | [optional] 
**AuthType** | Pointer to [**WirelessLANAuthType**](WirelessLANAuthType.md) |  | [optional] 
**AuthCipher** | Pointer to [**WirelessLANAuthCipher**](WirelessLANAuthCipher.md) |  | [optional] 
**AuthPsk** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewWirelessLAN

`func NewWirelessLAN(id int32, url string, displayUrl string, display string, ssid string, created NullableTime, lastUpdated NullableTime, ) *WirelessLAN`

NewWirelessLAN instantiates a new WirelessLAN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessLANWithDefaults

`func NewWirelessLANWithDefaults() *WirelessLAN`

NewWirelessLANWithDefaults instantiates a new WirelessLAN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WirelessLAN) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WirelessLAN) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WirelessLAN) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *WirelessLAN) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WirelessLAN) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WirelessLAN) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *WirelessLAN) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *WirelessLAN) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *WirelessLAN) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *WirelessLAN) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *WirelessLAN) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *WirelessLAN) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetSsid

`func (o *WirelessLAN) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WirelessLAN) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WirelessLAN) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetDescription

`func (o *WirelessLAN) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WirelessLAN) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WirelessLAN) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WirelessLAN) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *WirelessLAN) GetGroup() BriefWirelessLANGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WirelessLAN) GetGroupOk() (*BriefWirelessLANGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WirelessLAN) SetGroup(v BriefWirelessLANGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *WirelessLAN) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *WirelessLAN) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *WirelessLAN) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetStatus

`func (o *WirelessLAN) GetStatus() WirelessLANStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WirelessLAN) GetStatusOk() (*WirelessLANStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WirelessLAN) SetStatus(v WirelessLANStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WirelessLAN) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVlan

`func (o *WirelessLAN) GetVlan() BriefVLAN`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *WirelessLAN) GetVlanOk() (*BriefVLAN, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *WirelessLAN) SetVlan(v BriefVLAN)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *WirelessLAN) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *WirelessLAN) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *WirelessLAN) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetTenant

`func (o *WirelessLAN) GetTenant() BriefTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WirelessLAN) GetTenantOk() (*BriefTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WirelessLAN) SetTenant(v BriefTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WirelessLAN) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WirelessLAN) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WirelessLAN) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetAuthType

`func (o *WirelessLAN) GetAuthType() WirelessLANAuthType`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *WirelessLAN) GetAuthTypeOk() (*WirelessLANAuthType, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *WirelessLAN) SetAuthType(v WirelessLANAuthType)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *WirelessLAN) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthCipher

`func (o *WirelessLAN) GetAuthCipher() WirelessLANAuthCipher`

GetAuthCipher returns the AuthCipher field if non-nil, zero value otherwise.

### GetAuthCipherOk

`func (o *WirelessLAN) GetAuthCipherOk() (*WirelessLANAuthCipher, bool)`

GetAuthCipherOk returns a tuple with the AuthCipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCipher

`func (o *WirelessLAN) SetAuthCipher(v WirelessLANAuthCipher)`

SetAuthCipher sets AuthCipher field to given value.

### HasAuthCipher

`func (o *WirelessLAN) HasAuthCipher() bool`

HasAuthCipher returns a boolean if a field has been set.

### GetAuthPsk

`func (o *WirelessLAN) GetAuthPsk() string`

GetAuthPsk returns the AuthPsk field if non-nil, zero value otherwise.

### GetAuthPskOk

`func (o *WirelessLAN) GetAuthPskOk() (*string, bool)`

GetAuthPskOk returns a tuple with the AuthPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPsk

`func (o *WirelessLAN) SetAuthPsk(v string)`

SetAuthPsk sets AuthPsk field to given value.

### HasAuthPsk

`func (o *WirelessLAN) HasAuthPsk() bool`

HasAuthPsk returns a boolean if a field has been set.

### GetComments

`func (o *WirelessLAN) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WirelessLAN) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WirelessLAN) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WirelessLAN) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WirelessLAN) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WirelessLAN) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WirelessLAN) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WirelessLAN) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WirelessLAN) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WirelessLAN) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WirelessLAN) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WirelessLAN) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *WirelessLAN) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WirelessLAN) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WirelessLAN) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *WirelessLAN) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *WirelessLAN) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *WirelessLAN) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *WirelessLAN) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *WirelessLAN) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *WirelessLAN) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *WirelessLAN) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


