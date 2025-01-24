# WritableWirelessLANRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ssid** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**NullableBriefWirelessLANGroupRequest**](BriefWirelessLANGroupRequest.md) |  | [optional] 
**Status** | Pointer to [**PatchedWritableWirelessLANRequestStatus**](PatchedWritableWirelessLANRequestStatus.md) |  | [optional] 
**Vlan** | Pointer to [**NullableBriefVLANRequest**](BriefVLANRequest.md) |  | [optional] 
**ScopeType** | Pointer to **NullableString** |  | [optional] 
**ScopeId** | Pointer to **NullableInt32** |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**AuthType** | Pointer to [**NullableAuthenticationType1**](AuthenticationType1.md) |  | [optional] 
**AuthCipher** | Pointer to [**NullableAuthenticationCipher**](AuthenticationCipher.md) |  | [optional] 
**AuthPsk** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewWritableWirelessLANRequest

`func NewWritableWirelessLANRequest(ssid string, ) *WritableWirelessLANRequest`

NewWritableWirelessLANRequest instantiates a new WritableWirelessLANRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableWirelessLANRequestWithDefaults

`func NewWritableWirelessLANRequestWithDefaults() *WritableWirelessLANRequest`

NewWritableWirelessLANRequestWithDefaults instantiates a new WritableWirelessLANRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsid

`func (o *WritableWirelessLANRequest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WritableWirelessLANRequest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WritableWirelessLANRequest) SetSsid(v string)`

SetSsid sets Ssid field to given value.


### GetDescription

`func (o *WritableWirelessLANRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableWirelessLANRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableWirelessLANRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableWirelessLANRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *WritableWirelessLANRequest) GetGroup() BriefWirelessLANGroupRequest`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *WritableWirelessLANRequest) GetGroupOk() (*BriefWirelessLANGroupRequest, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *WritableWirelessLANRequest) SetGroup(v BriefWirelessLANGroupRequest)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *WritableWirelessLANRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *WritableWirelessLANRequest) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *WritableWirelessLANRequest) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetStatus

`func (o *WritableWirelessLANRequest) GetStatus() PatchedWritableWirelessLANRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableWirelessLANRequest) GetStatusOk() (*PatchedWritableWirelessLANRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableWirelessLANRequest) SetStatus(v PatchedWritableWirelessLANRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableWirelessLANRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVlan

`func (o *WritableWirelessLANRequest) GetVlan() BriefVLANRequest`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *WritableWirelessLANRequest) GetVlanOk() (*BriefVLANRequest, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *WritableWirelessLANRequest) SetVlan(v BriefVLANRequest)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *WritableWirelessLANRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *WritableWirelessLANRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *WritableWirelessLANRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetScopeType

`func (o *WritableWirelessLANRequest) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *WritableWirelessLANRequest) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *WritableWirelessLANRequest) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *WritableWirelessLANRequest) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### SetScopeTypeNil

`func (o *WritableWirelessLANRequest) SetScopeTypeNil(b bool)`

 SetScopeTypeNil sets the value for ScopeType to be an explicit nil

### UnsetScopeType
`func (o *WritableWirelessLANRequest) UnsetScopeType()`

UnsetScopeType ensures that no value is present for ScopeType, not even an explicit nil
### GetScopeId

`func (o *WritableWirelessLANRequest) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *WritableWirelessLANRequest) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *WritableWirelessLANRequest) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *WritableWirelessLANRequest) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### SetScopeIdNil

`func (o *WritableWirelessLANRequest) SetScopeIdNil(b bool)`

 SetScopeIdNil sets the value for ScopeId to be an explicit nil

### UnsetScopeId
`func (o *WritableWirelessLANRequest) UnsetScopeId()`

UnsetScopeId ensures that no value is present for ScopeId, not even an explicit nil
### GetTenant

`func (o *WritableWirelessLANRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableWirelessLANRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableWirelessLANRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableWirelessLANRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableWirelessLANRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableWirelessLANRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetAuthType

`func (o *WritableWirelessLANRequest) GetAuthType() AuthenticationType1`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *WritableWirelessLANRequest) GetAuthTypeOk() (*AuthenticationType1, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *WritableWirelessLANRequest) SetAuthType(v AuthenticationType1)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *WritableWirelessLANRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *WritableWirelessLANRequest) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *WritableWirelessLANRequest) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetAuthCipher

`func (o *WritableWirelessLANRequest) GetAuthCipher() AuthenticationCipher`

GetAuthCipher returns the AuthCipher field if non-nil, zero value otherwise.

### GetAuthCipherOk

`func (o *WritableWirelessLANRequest) GetAuthCipherOk() (*AuthenticationCipher, bool)`

GetAuthCipherOk returns a tuple with the AuthCipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCipher

`func (o *WritableWirelessLANRequest) SetAuthCipher(v AuthenticationCipher)`

SetAuthCipher sets AuthCipher field to given value.

### HasAuthCipher

`func (o *WritableWirelessLANRequest) HasAuthCipher() bool`

HasAuthCipher returns a boolean if a field has been set.

### SetAuthCipherNil

`func (o *WritableWirelessLANRequest) SetAuthCipherNil(b bool)`

 SetAuthCipherNil sets the value for AuthCipher to be an explicit nil

### UnsetAuthCipher
`func (o *WritableWirelessLANRequest) UnsetAuthCipher()`

UnsetAuthCipher ensures that no value is present for AuthCipher, not even an explicit nil
### GetAuthPsk

`func (o *WritableWirelessLANRequest) GetAuthPsk() string`

GetAuthPsk returns the AuthPsk field if non-nil, zero value otherwise.

### GetAuthPskOk

`func (o *WritableWirelessLANRequest) GetAuthPskOk() (*string, bool)`

GetAuthPskOk returns a tuple with the AuthPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPsk

`func (o *WritableWirelessLANRequest) SetAuthPsk(v string)`

SetAuthPsk sets AuthPsk field to given value.

### HasAuthPsk

`func (o *WritableWirelessLANRequest) HasAuthPsk() bool`

HasAuthPsk returns a boolean if a field has been set.

### GetComments

`func (o *WritableWirelessLANRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableWirelessLANRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableWirelessLANRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableWirelessLANRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableWirelessLANRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableWirelessLANRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableWirelessLANRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableWirelessLANRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableWirelessLANRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableWirelessLANRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableWirelessLANRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableWirelessLANRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


