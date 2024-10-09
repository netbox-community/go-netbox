# PatchedWritableWirelessLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceA** | Pointer to [**BriefInterfaceRequest**](BriefInterfaceRequest.md) |  | [optional] 
**InterfaceB** | Pointer to [**BriefInterfaceRequest**](BriefInterfaceRequest.md) |  | [optional] 
**Ssid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**CableStatusValue**](CableStatusValue.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenantRequest**](BriefTenantRequest.md) |  | [optional] 
**AuthType** | Pointer to [**AuthenticationType3**](AuthenticationType3.md) |  | [optional] 
**AuthCipher** | Pointer to [**AuthenticationCipher1**](AuthenticationCipher1.md) |  | [optional] 
**AuthPsk** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to **NullableFloat64** |  | [optional] 
**DistanceUnit** | Pointer to [**PatchedWritableWirelessLinkRequestDistanceUnit**](PatchedWritableWirelessLinkRequestDistanceUnit.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableWirelessLinkRequest

`func NewPatchedWritableWirelessLinkRequest() *PatchedWritableWirelessLinkRequest`

NewPatchedWritableWirelessLinkRequest instantiates a new PatchedWritableWirelessLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableWirelessLinkRequestWithDefaults

`func NewPatchedWritableWirelessLinkRequestWithDefaults() *PatchedWritableWirelessLinkRequest`

NewPatchedWritableWirelessLinkRequestWithDefaults instantiates a new PatchedWritableWirelessLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceA

`func (o *PatchedWritableWirelessLinkRequest) GetInterfaceA() BriefInterfaceRequest`

GetInterfaceA returns the InterfaceA field if non-nil, zero value otherwise.

### GetInterfaceAOk

`func (o *PatchedWritableWirelessLinkRequest) GetInterfaceAOk() (*BriefInterfaceRequest, bool)`

GetInterfaceAOk returns a tuple with the InterfaceA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceA

`func (o *PatchedWritableWirelessLinkRequest) SetInterfaceA(v BriefInterfaceRequest)`

SetInterfaceA sets InterfaceA field to given value.

### HasInterfaceA

`func (o *PatchedWritableWirelessLinkRequest) HasInterfaceA() bool`

HasInterfaceA returns a boolean if a field has been set.

### GetInterfaceB

`func (o *PatchedWritableWirelessLinkRequest) GetInterfaceB() BriefInterfaceRequest`

GetInterfaceB returns the InterfaceB field if non-nil, zero value otherwise.

### GetInterfaceBOk

`func (o *PatchedWritableWirelessLinkRequest) GetInterfaceBOk() (*BriefInterfaceRequest, bool)`

GetInterfaceBOk returns a tuple with the InterfaceB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceB

`func (o *PatchedWritableWirelessLinkRequest) SetInterfaceB(v BriefInterfaceRequest)`

SetInterfaceB sets InterfaceB field to given value.

### HasInterfaceB

`func (o *PatchedWritableWirelessLinkRequest) HasInterfaceB() bool`

HasInterfaceB returns a boolean if a field has been set.

### GetSsid

`func (o *PatchedWritableWirelessLinkRequest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *PatchedWritableWirelessLinkRequest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *PatchedWritableWirelessLinkRequest) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *PatchedWritableWirelessLinkRequest) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableWirelessLinkRequest) GetStatus() CableStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableWirelessLinkRequest) GetStatusOk() (*CableStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableWirelessLinkRequest) SetStatus(v CableStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableWirelessLinkRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableWirelessLinkRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableWirelessLinkRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableWirelessLinkRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableWirelessLinkRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableWirelessLinkRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableWirelessLinkRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetAuthType

`func (o *PatchedWritableWirelessLinkRequest) GetAuthType() AuthenticationType3`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *PatchedWritableWirelessLinkRequest) GetAuthTypeOk() (*AuthenticationType3, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *PatchedWritableWirelessLinkRequest) SetAuthType(v AuthenticationType3)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *PatchedWritableWirelessLinkRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthCipher

`func (o *PatchedWritableWirelessLinkRequest) GetAuthCipher() AuthenticationCipher1`

GetAuthCipher returns the AuthCipher field if non-nil, zero value otherwise.

### GetAuthCipherOk

`func (o *PatchedWritableWirelessLinkRequest) GetAuthCipherOk() (*AuthenticationCipher1, bool)`

GetAuthCipherOk returns a tuple with the AuthCipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCipher

`func (o *PatchedWritableWirelessLinkRequest) SetAuthCipher(v AuthenticationCipher1)`

SetAuthCipher sets AuthCipher field to given value.

### HasAuthCipher

`func (o *PatchedWritableWirelessLinkRequest) HasAuthCipher() bool`

HasAuthCipher returns a boolean if a field has been set.

### GetAuthPsk

`func (o *PatchedWritableWirelessLinkRequest) GetAuthPsk() string`

GetAuthPsk returns the AuthPsk field if non-nil, zero value otherwise.

### GetAuthPskOk

`func (o *PatchedWritableWirelessLinkRequest) GetAuthPskOk() (*string, bool)`

GetAuthPskOk returns a tuple with the AuthPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPsk

`func (o *PatchedWritableWirelessLinkRequest) SetAuthPsk(v string)`

SetAuthPsk sets AuthPsk field to given value.

### HasAuthPsk

`func (o *PatchedWritableWirelessLinkRequest) HasAuthPsk() bool`

HasAuthPsk returns a boolean if a field has been set.

### GetDistance

`func (o *PatchedWritableWirelessLinkRequest) GetDistance() float64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *PatchedWritableWirelessLinkRequest) GetDistanceOk() (*float64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *PatchedWritableWirelessLinkRequest) SetDistance(v float64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *PatchedWritableWirelessLinkRequest) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### SetDistanceNil

`func (o *PatchedWritableWirelessLinkRequest) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *PatchedWritableWirelessLinkRequest) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil
### GetDistanceUnit

`func (o *PatchedWritableWirelessLinkRequest) GetDistanceUnit() PatchedWritableWirelessLinkRequestDistanceUnit`

GetDistanceUnit returns the DistanceUnit field if non-nil, zero value otherwise.

### GetDistanceUnitOk

`func (o *PatchedWritableWirelessLinkRequest) GetDistanceUnitOk() (*PatchedWritableWirelessLinkRequestDistanceUnit, bool)`

GetDistanceUnitOk returns a tuple with the DistanceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceUnit

`func (o *PatchedWritableWirelessLinkRequest) SetDistanceUnit(v PatchedWritableWirelessLinkRequestDistanceUnit)`

SetDistanceUnit sets DistanceUnit field to given value.

### HasDistanceUnit

`func (o *PatchedWritableWirelessLinkRequest) HasDistanceUnit() bool`

HasDistanceUnit returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableWirelessLinkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableWirelessLinkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableWirelessLinkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableWirelessLinkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableWirelessLinkRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableWirelessLinkRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableWirelessLinkRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableWirelessLinkRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableWirelessLinkRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableWirelessLinkRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableWirelessLinkRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableWirelessLinkRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableWirelessLinkRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableWirelessLinkRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableWirelessLinkRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableWirelessLinkRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


