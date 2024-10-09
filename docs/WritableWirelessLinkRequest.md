# WritableWirelessLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceA** | [**BriefInterfaceRequest**](BriefInterfaceRequest.md) |  | 
**InterfaceB** | [**BriefInterfaceRequest**](BriefInterfaceRequest.md) |  | 
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

### NewWritableWirelessLinkRequest

`func NewWritableWirelessLinkRequest(interfaceA BriefInterfaceRequest, interfaceB BriefInterfaceRequest, ) *WritableWirelessLinkRequest`

NewWritableWirelessLinkRequest instantiates a new WritableWirelessLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableWirelessLinkRequestWithDefaults

`func NewWritableWirelessLinkRequestWithDefaults() *WritableWirelessLinkRequest`

NewWritableWirelessLinkRequestWithDefaults instantiates a new WritableWirelessLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceA

`func (o *WritableWirelessLinkRequest) GetInterfaceA() BriefInterfaceRequest`

GetInterfaceA returns the InterfaceA field if non-nil, zero value otherwise.

### GetInterfaceAOk

`func (o *WritableWirelessLinkRequest) GetInterfaceAOk() (*BriefInterfaceRequest, bool)`

GetInterfaceAOk returns a tuple with the InterfaceA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceA

`func (o *WritableWirelessLinkRequest) SetInterfaceA(v BriefInterfaceRequest)`

SetInterfaceA sets InterfaceA field to given value.


### GetInterfaceB

`func (o *WritableWirelessLinkRequest) GetInterfaceB() BriefInterfaceRequest`

GetInterfaceB returns the InterfaceB field if non-nil, zero value otherwise.

### GetInterfaceBOk

`func (o *WritableWirelessLinkRequest) GetInterfaceBOk() (*BriefInterfaceRequest, bool)`

GetInterfaceBOk returns a tuple with the InterfaceB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceB

`func (o *WritableWirelessLinkRequest) SetInterfaceB(v BriefInterfaceRequest)`

SetInterfaceB sets InterfaceB field to given value.


### GetSsid

`func (o *WritableWirelessLinkRequest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WritableWirelessLinkRequest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WritableWirelessLinkRequest) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *WritableWirelessLinkRequest) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetStatus

`func (o *WritableWirelessLinkRequest) GetStatus() CableStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritableWirelessLinkRequest) GetStatusOk() (*CableStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritableWirelessLinkRequest) SetStatus(v CableStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WritableWirelessLinkRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *WritableWirelessLinkRequest) GetTenant() BriefTenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WritableWirelessLinkRequest) GetTenantOk() (*BriefTenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WritableWirelessLinkRequest) SetTenant(v BriefTenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WritableWirelessLinkRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WritableWirelessLinkRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WritableWirelessLinkRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetAuthType

`func (o *WritableWirelessLinkRequest) GetAuthType() AuthenticationType3`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *WritableWirelessLinkRequest) GetAuthTypeOk() (*AuthenticationType3, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *WritableWirelessLinkRequest) SetAuthType(v AuthenticationType3)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *WritableWirelessLinkRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthCipher

`func (o *WritableWirelessLinkRequest) GetAuthCipher() AuthenticationCipher1`

GetAuthCipher returns the AuthCipher field if non-nil, zero value otherwise.

### GetAuthCipherOk

`func (o *WritableWirelessLinkRequest) GetAuthCipherOk() (*AuthenticationCipher1, bool)`

GetAuthCipherOk returns a tuple with the AuthCipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCipher

`func (o *WritableWirelessLinkRequest) SetAuthCipher(v AuthenticationCipher1)`

SetAuthCipher sets AuthCipher field to given value.

### HasAuthCipher

`func (o *WritableWirelessLinkRequest) HasAuthCipher() bool`

HasAuthCipher returns a boolean if a field has been set.

### GetAuthPsk

`func (o *WritableWirelessLinkRequest) GetAuthPsk() string`

GetAuthPsk returns the AuthPsk field if non-nil, zero value otherwise.

### GetAuthPskOk

`func (o *WritableWirelessLinkRequest) GetAuthPskOk() (*string, bool)`

GetAuthPskOk returns a tuple with the AuthPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPsk

`func (o *WritableWirelessLinkRequest) SetAuthPsk(v string)`

SetAuthPsk sets AuthPsk field to given value.

### HasAuthPsk

`func (o *WritableWirelessLinkRequest) HasAuthPsk() bool`

HasAuthPsk returns a boolean if a field has been set.

### GetDistance

`func (o *WritableWirelessLinkRequest) GetDistance() float64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *WritableWirelessLinkRequest) GetDistanceOk() (*float64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *WritableWirelessLinkRequest) SetDistance(v float64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *WritableWirelessLinkRequest) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### SetDistanceNil

`func (o *WritableWirelessLinkRequest) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *WritableWirelessLinkRequest) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil
### GetDistanceUnit

`func (o *WritableWirelessLinkRequest) GetDistanceUnit() PatchedWritableWirelessLinkRequestDistanceUnit`

GetDistanceUnit returns the DistanceUnit field if non-nil, zero value otherwise.

### GetDistanceUnitOk

`func (o *WritableWirelessLinkRequest) GetDistanceUnitOk() (*PatchedWritableWirelessLinkRequestDistanceUnit, bool)`

GetDistanceUnitOk returns a tuple with the DistanceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceUnit

`func (o *WritableWirelessLinkRequest) SetDistanceUnit(v PatchedWritableWirelessLinkRequestDistanceUnit)`

SetDistanceUnit sets DistanceUnit field to given value.

### HasDistanceUnit

`func (o *WritableWirelessLinkRequest) HasDistanceUnit() bool`

HasDistanceUnit returns a boolean if a field has been set.

### GetDescription

`func (o *WritableWirelessLinkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableWirelessLinkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableWirelessLinkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableWirelessLinkRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *WritableWirelessLinkRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritableWirelessLinkRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritableWirelessLinkRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritableWirelessLinkRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WritableWirelessLinkRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableWirelessLinkRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableWirelessLinkRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableWirelessLinkRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritableWirelessLinkRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableWirelessLinkRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableWirelessLinkRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableWirelessLinkRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


