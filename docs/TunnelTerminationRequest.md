# TunnelTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tunnel** | [**BriefTunnelRequest**](BriefTunnelRequest.md) |  | 
**Role** | [**PatchedWritableTunnelTerminationRequestRole**](PatchedWritableTunnelTerminationRequestRole.md) |  | 
**TerminationType** | **string** |  | 
**TerminationId** | **NullableInt64** |  | 
**OutsideIp** | Pointer to [**NullableBriefIPAddressRequest**](BriefIPAddressRequest.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTunnelTerminationRequest

`func NewTunnelTerminationRequest(tunnel BriefTunnelRequest, role PatchedWritableTunnelTerminationRequestRole, terminationType string, terminationId NullableInt64, ) *TunnelTerminationRequest`

NewTunnelTerminationRequest instantiates a new TunnelTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunnelTerminationRequestWithDefaults

`func NewTunnelTerminationRequestWithDefaults() *TunnelTerminationRequest`

NewTunnelTerminationRequestWithDefaults instantiates a new TunnelTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunnel

`func (o *TunnelTerminationRequest) GetTunnel() BriefTunnelRequest`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *TunnelTerminationRequest) GetTunnelOk() (*BriefTunnelRequest, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *TunnelTerminationRequest) SetTunnel(v BriefTunnelRequest)`

SetTunnel sets Tunnel field to given value.


### GetRole

`func (o *TunnelTerminationRequest) GetRole() PatchedWritableTunnelTerminationRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TunnelTerminationRequest) GetRoleOk() (*PatchedWritableTunnelTerminationRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TunnelTerminationRequest) SetRole(v PatchedWritableTunnelTerminationRequestRole)`

SetRole sets Role field to given value.


### GetTerminationType

`func (o *TunnelTerminationRequest) GetTerminationType() string`

GetTerminationType returns the TerminationType field if non-nil, zero value otherwise.

### GetTerminationTypeOk

`func (o *TunnelTerminationRequest) GetTerminationTypeOk() (*string, bool)`

GetTerminationTypeOk returns a tuple with the TerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationType

`func (o *TunnelTerminationRequest) SetTerminationType(v string)`

SetTerminationType sets TerminationType field to given value.


### GetTerminationId

`func (o *TunnelTerminationRequest) GetTerminationId() int64`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *TunnelTerminationRequest) GetTerminationIdOk() (*int64, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *TunnelTerminationRequest) SetTerminationId(v int64)`

SetTerminationId sets TerminationId field to given value.


### SetTerminationIdNil

`func (o *TunnelTerminationRequest) SetTerminationIdNil(b bool)`

 SetTerminationIdNil sets the value for TerminationId to be an explicit nil

### UnsetTerminationId
`func (o *TunnelTerminationRequest) UnsetTerminationId()`

UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
### GetOutsideIp

`func (o *TunnelTerminationRequest) GetOutsideIp() BriefIPAddressRequest`

GetOutsideIp returns the OutsideIp field if non-nil, zero value otherwise.

### GetOutsideIpOk

`func (o *TunnelTerminationRequest) GetOutsideIpOk() (*BriefIPAddressRequest, bool)`

GetOutsideIpOk returns a tuple with the OutsideIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideIp

`func (o *TunnelTerminationRequest) SetOutsideIp(v BriefIPAddressRequest)`

SetOutsideIp sets OutsideIp field to given value.

### HasOutsideIp

`func (o *TunnelTerminationRequest) HasOutsideIp() bool`

HasOutsideIp returns a boolean if a field has been set.

### SetOutsideIpNil

`func (o *TunnelTerminationRequest) SetOutsideIpNil(b bool)`

 SetOutsideIpNil sets the value for OutsideIp to be an explicit nil

### UnsetOutsideIp
`func (o *TunnelTerminationRequest) UnsetOutsideIp()`

UnsetOutsideIp ensures that no value is present for OutsideIp, not even an explicit nil
### GetTags

`func (o *TunnelTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TunnelTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TunnelTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TunnelTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *TunnelTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TunnelTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TunnelTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TunnelTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


