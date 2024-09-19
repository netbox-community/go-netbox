# PatchedWritableTunnelTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tunnel** | Pointer to [**BriefTunnelRequest**](BriefTunnelRequest.md) |  | [optional] 
**Role** | Pointer to [**PatchedWritableTunnelTerminationRequestRole**](PatchedWritableTunnelTerminationRequestRole.md) |  | [optional] 
**TerminationType** | Pointer to **string** |  | [optional] 
**TerminationId** | Pointer to **NullableInt64** |  | [optional] 
**OutsideIp** | Pointer to [**NullableBriefIPAddressRequest**](BriefIPAddressRequest.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableTunnelTerminationRequest

`func NewPatchedWritableTunnelTerminationRequest() *PatchedWritableTunnelTerminationRequest`

NewPatchedWritableTunnelTerminationRequest instantiates a new PatchedWritableTunnelTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableTunnelTerminationRequestWithDefaults

`func NewPatchedWritableTunnelTerminationRequestWithDefaults() *PatchedWritableTunnelTerminationRequest`

NewPatchedWritableTunnelTerminationRequestWithDefaults instantiates a new PatchedWritableTunnelTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunnel

`func (o *PatchedWritableTunnelTerminationRequest) GetTunnel() BriefTunnelRequest`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *PatchedWritableTunnelTerminationRequest) GetTunnelOk() (*BriefTunnelRequest, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *PatchedWritableTunnelTerminationRequest) SetTunnel(v BriefTunnelRequest)`

SetTunnel sets Tunnel field to given value.

### HasTunnel

`func (o *PatchedWritableTunnelTerminationRequest) HasTunnel() bool`

HasTunnel returns a boolean if a field has been set.

### GetRole

`func (o *PatchedWritableTunnelTerminationRequest) GetRole() PatchedWritableTunnelTerminationRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedWritableTunnelTerminationRequest) GetRoleOk() (*PatchedWritableTunnelTerminationRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedWritableTunnelTerminationRequest) SetRole(v PatchedWritableTunnelTerminationRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedWritableTunnelTerminationRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTerminationType

`func (o *PatchedWritableTunnelTerminationRequest) GetTerminationType() string`

GetTerminationType returns the TerminationType field if non-nil, zero value otherwise.

### GetTerminationTypeOk

`func (o *PatchedWritableTunnelTerminationRequest) GetTerminationTypeOk() (*string, bool)`

GetTerminationTypeOk returns a tuple with the TerminationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationType

`func (o *PatchedWritableTunnelTerminationRequest) SetTerminationType(v string)`

SetTerminationType sets TerminationType field to given value.

### HasTerminationType

`func (o *PatchedWritableTunnelTerminationRequest) HasTerminationType() bool`

HasTerminationType returns a boolean if a field has been set.

### GetTerminationId

`func (o *PatchedWritableTunnelTerminationRequest) GetTerminationId() int64`

GetTerminationId returns the TerminationId field if non-nil, zero value otherwise.

### GetTerminationIdOk

`func (o *PatchedWritableTunnelTerminationRequest) GetTerminationIdOk() (*int64, bool)`

GetTerminationIdOk returns a tuple with the TerminationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationId

`func (o *PatchedWritableTunnelTerminationRequest) SetTerminationId(v int64)`

SetTerminationId sets TerminationId field to given value.

### HasTerminationId

`func (o *PatchedWritableTunnelTerminationRequest) HasTerminationId() bool`

HasTerminationId returns a boolean if a field has been set.

### SetTerminationIdNil

`func (o *PatchedWritableTunnelTerminationRequest) SetTerminationIdNil(b bool)`

 SetTerminationIdNil sets the value for TerminationId to be an explicit nil

### UnsetTerminationId
`func (o *PatchedWritableTunnelTerminationRequest) UnsetTerminationId()`

UnsetTerminationId ensures that no value is present for TerminationId, not even an explicit nil
### GetOutsideIp

`func (o *PatchedWritableTunnelTerminationRequest) GetOutsideIp() BriefIPAddressRequest`

GetOutsideIp returns the OutsideIp field if non-nil, zero value otherwise.

### GetOutsideIpOk

`func (o *PatchedWritableTunnelTerminationRequest) GetOutsideIpOk() (*BriefIPAddressRequest, bool)`

GetOutsideIpOk returns a tuple with the OutsideIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideIp

`func (o *PatchedWritableTunnelTerminationRequest) SetOutsideIp(v BriefIPAddressRequest)`

SetOutsideIp sets OutsideIp field to given value.

### HasOutsideIp

`func (o *PatchedWritableTunnelTerminationRequest) HasOutsideIp() bool`

HasOutsideIp returns a boolean if a field has been set.

### SetOutsideIpNil

`func (o *PatchedWritableTunnelTerminationRequest) SetOutsideIpNil(b bool)`

 SetOutsideIpNil sets the value for OutsideIp to be an explicit nil

### UnsetOutsideIp
`func (o *PatchedWritableTunnelTerminationRequest) UnsetOutsideIp()`

UnsetOutsideIp ensures that no value is present for OutsideIp, not even an explicit nil
### GetTags

`func (o *PatchedWritableTunnelTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableTunnelTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableTunnelTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableTunnelTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableTunnelTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableTunnelTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableTunnelTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableTunnelTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


