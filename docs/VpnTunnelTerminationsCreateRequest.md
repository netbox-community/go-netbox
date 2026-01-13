# VpnTunnelTerminationsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tunnel** | [**PatchedWritableTunnelTerminationRequestTunnel**](PatchedWritableTunnelTerminationRequestTunnel.md) |  | 
**Role** | Pointer to [**PatchedWritableTunnelTerminationRequestRole**](PatchedWritableTunnelTerminationRequestRole.md) |  | [optional] 
**TerminationType** | **string** |  | 
**TerminationId** | Pointer to **NullableInt64** |  | [optional] 
**OutsideIp** | Pointer to [**NullableDeviceWithConfigContextRequestPrimaryIp4**](DeviceWithConfigContextRequestPrimaryIp4.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewVpnTunnelTerminationsCreateRequest

`func NewVpnTunnelTerminationsCreateRequest(tunnel PatchedWritableTunnelTerminationRequestTunnel, terminationType string, ) *VpnTunnelTerminationsCreateRequest`

NewVpnTunnelTerminationsCreateRequest instantiates a new VpnTunnelTerminationsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelTerminationsCreateRequestWithDefaults

`func NewVpnTunnelTerminationsCreateRequestWithDefaults() *VpnTunnelTerminationsCreateRequest`

NewVpnTunnelTerminationsCreateRequestWithDefaults instantiates a new VpnTunnelTerminationsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunnel

`func (o *VpnTunnelTerminationsCreateRequest) GetTunnel() PatchedWritableTunnelTerminationRequestTunnel`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *VpnTunnelTerminationsCreateRequest) GetTunnelOk() (*PatchedWritableTunnelTerminationRequestTunnel, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *VpnTunnelTerminationsCreateRequest) SetTunnel(v PatchedWritableTunnelTerminationRequestTunnel)`

SetTunnel sets Tunnel field to given value.


### GetRole

`func (o *VpnTunnelTerminationsCreateRequest) GetRole() PatchedWritableTunnelTerminationRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VpnTunnelTerminationsCreateRequest) GetRoleOk() (*PatchedWritableTunnelTerminationRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VpnTunnelTerminationsCreateRequest) SetRole(v PatchedWritableTunnelTerminationRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *VpnTunnelTerminationsCreateRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTerminationType

`func (o *VpnTunnelTerminationsCreateRequest) GetTerminationType() string`

GetTerminationType returns the TerminationType field if non-nil, zero value otherwise.

### GetTerminationTypeOk

`func (o *VpnTunnelTerminationsCreateRequest) GetTerminationTypeOk() (*string, bool)`

GetTerminationTypeOk returns a tuple with the TerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationType

`func (o *VpnTunnelTerminationsCreateRequest) SetTerminationType(v string)`

SetTerminationType sets TerminationType field to given value.


### GetTerminationId

`func (o *VpnTunnelTerminationsCreateRequest) GetTerminationId() int64`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *VpnTunnelTerminationsCreateRequest) GetTerminationIdOk() (*int64, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *VpnTunnelTerminationsCreateRequest) SetTerminationId(v int64)`

SetTerminationId sets TerminationId field to given value.

### HasTerminationId

`func (o *VpnTunnelTerminationsCreateRequest) HasTerminationId() bool`

HasTerminationId returns a boolean if a field has been set.

### SetTerminationIdNil

`func (o *VpnTunnelTerminationsCreateRequest) SetTerminationIdNil(b bool)`

 SetTerminationIdNil sets the value for TerminationId to be an explicit nil

### UnsetTerminationId
`func (o *VpnTunnelTerminationsCreateRequest) UnsetTerminationId()`

UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
### GetOutsideIp

`func (o *VpnTunnelTerminationsCreateRequest) GetOutsideIp() DeviceWithConfigContextRequestPrimaryIp4`

GetOutsideIp returns the OutsideIp field if non-nil, zero value otherwise.

### GetOutsideIpOk

`func (o *VpnTunnelTerminationsCreateRequest) GetOutsideIpOk() (*DeviceWithConfigContextRequestPrimaryIp4, bool)`

GetOutsideIpOk returns a tuple with the OutsideIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideIp

`func (o *VpnTunnelTerminationsCreateRequest) SetOutsideIp(v DeviceWithConfigContextRequestPrimaryIp4)`

SetOutsideIp sets OutsideIp field to given value.

### HasOutsideIp

`func (o *VpnTunnelTerminationsCreateRequest) HasOutsideIp() bool`

HasOutsideIp returns a boolean if a field has been set.

### SetOutsideIpNil

`func (o *VpnTunnelTerminationsCreateRequest) SetOutsideIpNil(b bool)`

 SetOutsideIpNil sets the value for OutsideIp to be an explicit nil

### UnsetOutsideIp
`func (o *VpnTunnelTerminationsCreateRequest) UnsetOutsideIp()`

UnsetOutsideIp ensures that no value is present for OutsideIp, not even an explicit nil
### GetTags

`func (o *VpnTunnelTerminationsCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpnTunnelTerminationsCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpnTunnelTerminationsCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpnTunnelTerminationsCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VpnTunnelTerminationsCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VpnTunnelTerminationsCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VpnTunnelTerminationsCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VpnTunnelTerminationsCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


