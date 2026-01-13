# WirelessWirelessLinksCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceA** | [**PatchedWritableVirtualCircuitTerminationRequestInterface**](PatchedWritableVirtualCircuitTerminationRequestInterface.md) |  | 
**InterfaceB** | [**PatchedWritableVirtualCircuitTerminationRequestInterface**](PatchedWritableVirtualCircuitTerminationRequestInterface.md) |  | 
**Ssid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**CableStatusValue**](CableStatusValue.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**AuthType** | Pointer to [**NullableAuthenticationType1**](AuthenticationType1.md) |  | [optional] 
**AuthCipher** | Pointer to [**NullableAuthenticationCipher**](AuthenticationCipher.md) |  | [optional] 
**AuthPsk** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to **NullableFloat64** |  | [optional] 
**DistanceUnit** | Pointer to [**NullableCircuitRequestDistanceUnit**](CircuitRequestDistanceUnit.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewWirelessWirelessLinksCreateRequest

`func NewWirelessWirelessLinksCreateRequest(interfaceA PatchedWritableVirtualCircuitTerminationRequestInterface, interfaceB PatchedWritableVirtualCircuitTerminationRequestInterface, ) *WirelessWirelessLinksCreateRequest`

NewWirelessWirelessLinksCreateRequest instantiates a new WirelessWirelessLinksCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessWirelessLinksCreateRequestWithDefaults

`func NewWirelessWirelessLinksCreateRequestWithDefaults() *WirelessWirelessLinksCreateRequest`

NewWirelessWirelessLinksCreateRequestWithDefaults instantiates a new WirelessWirelessLinksCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceA

`func (o *WirelessWirelessLinksCreateRequest) GetInterfaceA() PatchedWritableVirtualCircuitTerminationRequestInterface`

GetInterfaceA returns the InterfaceA field if non-nil, zero value otherwise.

### GetInterfaceAOk

`func (o *WirelessWirelessLinksCreateRequest) GetInterfaceAOk() (*PatchedWritableVirtualCircuitTerminationRequestInterface, bool)`

GetInterfaceAOk returns a tuple with the InterfaceA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceA

`func (o *WirelessWirelessLinksCreateRequest) SetInterfaceA(v PatchedWritableVirtualCircuitTerminationRequestInterface)`

SetInterfaceA sets InterfaceA field to given value.


### GetInterfaceB

`func (o *WirelessWirelessLinksCreateRequest) GetInterfaceB() PatchedWritableVirtualCircuitTerminationRequestInterface`

GetInterfaceB returns the InterfaceB field if non-nil, zero value otherwise.

### GetInterfaceBOk

`func (o *WirelessWirelessLinksCreateRequest) GetInterfaceBOk() (*PatchedWritableVirtualCircuitTerminationRequestInterface, bool)`

GetInterfaceBOk returns a tuple with the InterfaceB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceB

`func (o *WirelessWirelessLinksCreateRequest) SetInterfaceB(v PatchedWritableVirtualCircuitTerminationRequestInterface)`

SetInterfaceB sets InterfaceB field to given value.


### GetSsid

`func (o *WirelessWirelessLinksCreateRequest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *WirelessWirelessLinksCreateRequest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *WirelessWirelessLinksCreateRequest) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *WirelessWirelessLinksCreateRequest) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetStatus

`func (o *WirelessWirelessLinksCreateRequest) GetStatus() CableStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WirelessWirelessLinksCreateRequest) GetStatusOk() (*CableStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WirelessWirelessLinksCreateRequest) SetStatus(v CableStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WirelessWirelessLinksCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenant

`func (o *WirelessWirelessLinksCreateRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *WirelessWirelessLinksCreateRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *WirelessWirelessLinksCreateRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *WirelessWirelessLinksCreateRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *WirelessWirelessLinksCreateRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *WirelessWirelessLinksCreateRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetAuthType

`func (o *WirelessWirelessLinksCreateRequest) GetAuthType() AuthenticationType1`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *WirelessWirelessLinksCreateRequest) GetAuthTypeOk() (*AuthenticationType1, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *WirelessWirelessLinksCreateRequest) SetAuthType(v AuthenticationType1)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *WirelessWirelessLinksCreateRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *WirelessWirelessLinksCreateRequest) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *WirelessWirelessLinksCreateRequest) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetAuthCipher

`func (o *WirelessWirelessLinksCreateRequest) GetAuthCipher() AuthenticationCipher`

GetAuthCipher returns the AuthCipher field if non-nil, zero value otherwise.

### GetAuthCipherOk

`func (o *WirelessWirelessLinksCreateRequest) GetAuthCipherOk() (*AuthenticationCipher, bool)`

GetAuthCipherOk returns a tuple with the AuthCipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCipher

`func (o *WirelessWirelessLinksCreateRequest) SetAuthCipher(v AuthenticationCipher)`

SetAuthCipher sets AuthCipher field to given value.

### HasAuthCipher

`func (o *WirelessWirelessLinksCreateRequest) HasAuthCipher() bool`

HasAuthCipher returns a boolean if a field has been set.

### SetAuthCipherNil

`func (o *WirelessWirelessLinksCreateRequest) SetAuthCipherNil(b bool)`

 SetAuthCipherNil sets the value for AuthCipher to be an explicit nil

### UnsetAuthCipher
`func (o *WirelessWirelessLinksCreateRequest) UnsetAuthCipher()`

UnsetAuthCipher ensures that no value is present for AuthCipher, not even an explicit nil
### GetAuthPsk

`func (o *WirelessWirelessLinksCreateRequest) GetAuthPsk() string`

GetAuthPsk returns the AuthPsk field if non-nil, zero value otherwise.

### GetAuthPskOk

`func (o *WirelessWirelessLinksCreateRequest) GetAuthPskOk() (*string, bool)`

GetAuthPskOk returns a tuple with the AuthPsk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPsk

`func (o *WirelessWirelessLinksCreateRequest) SetAuthPsk(v string)`

SetAuthPsk sets AuthPsk field to given value.

### HasAuthPsk

`func (o *WirelessWirelessLinksCreateRequest) HasAuthPsk() bool`

HasAuthPsk returns a boolean if a field has been set.

### GetDistance

`func (o *WirelessWirelessLinksCreateRequest) GetDistance() float64`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *WirelessWirelessLinksCreateRequest) GetDistanceOk() (*float64, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *WirelessWirelessLinksCreateRequest) SetDistance(v float64)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *WirelessWirelessLinksCreateRequest) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### SetDistanceNil

`func (o *WirelessWirelessLinksCreateRequest) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *WirelessWirelessLinksCreateRequest) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil
### GetDistanceUnit

`func (o *WirelessWirelessLinksCreateRequest) GetDistanceUnit() CircuitRequestDistanceUnit`

GetDistanceUnit returns the DistanceUnit field if non-nil, zero value otherwise.

### GetDistanceUnitOk

`func (o *WirelessWirelessLinksCreateRequest) GetDistanceUnitOk() (*CircuitRequestDistanceUnit, bool)`

GetDistanceUnitOk returns a tuple with the DistanceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceUnit

`func (o *WirelessWirelessLinksCreateRequest) SetDistanceUnit(v CircuitRequestDistanceUnit)`

SetDistanceUnit sets DistanceUnit field to given value.

### HasDistanceUnit

`func (o *WirelessWirelessLinksCreateRequest) HasDistanceUnit() bool`

HasDistanceUnit returns a boolean if a field has been set.

### SetDistanceUnitNil

`func (o *WirelessWirelessLinksCreateRequest) SetDistanceUnitNil(b bool)`

 SetDistanceUnitNil sets the value for DistanceUnit to be an explicit nil

### UnsetDistanceUnit
`func (o *WirelessWirelessLinksCreateRequest) UnsetDistanceUnit()`

UnsetDistanceUnit ensures that no value is present for DistanceUnit, not even an explicit nil
### GetDescription

`func (o *WirelessWirelessLinksCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WirelessWirelessLinksCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WirelessWirelessLinksCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WirelessWirelessLinksCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *WirelessWirelessLinksCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WirelessWirelessLinksCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WirelessWirelessLinksCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WirelessWirelessLinksCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WirelessWirelessLinksCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WirelessWirelessLinksCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *WirelessWirelessLinksCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WirelessWirelessLinksCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WirelessWirelessLinksCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WirelessWirelessLinksCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *WirelessWirelessLinksCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WirelessWirelessLinksCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WirelessWirelessLinksCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WirelessWirelessLinksCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WirelessWirelessLinksCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WirelessWirelessLinksCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WirelessWirelessLinksCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WirelessWirelessLinksCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


