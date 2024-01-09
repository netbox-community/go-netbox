# PatchedWritableCircuitTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Circuit** | Pointer to **int32** |  | [optional] 
**TermSide** | Pointer to [**Termination**](Termination.md) |  | [optional] 
**Site** | Pointer to **NullableInt32** |  | [optional] 
**ProviderNetwork** | Pointer to **NullableInt32** |  | [optional] 
**PortSpeed** | Pointer to **NullableInt32** | Physical circuit speed | [optional] 
**UpstreamSpeed** | Pointer to **NullableInt32** | Upstream speed, if different from port speed | [optional] 
**XconnectId** | Pointer to **string** | ID of the local cross-connect | [optional] 
**PpInfo** | Pointer to **string** | Patch panel ID and port number(s) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableCircuitTerminationRequest

`func NewPatchedWritableCircuitTerminationRequest() *PatchedWritableCircuitTerminationRequest`

NewPatchedWritableCircuitTerminationRequest instantiates a new PatchedWritableCircuitTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableCircuitTerminationRequestWithDefaults

`func NewPatchedWritableCircuitTerminationRequestWithDefaults() *PatchedWritableCircuitTerminationRequest`

NewPatchedWritableCircuitTerminationRequestWithDefaults instantiates a new PatchedWritableCircuitTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuit

`func (o *PatchedWritableCircuitTerminationRequest) GetCircuit() int32`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *PatchedWritableCircuitTerminationRequest) GetCircuitOk() (*int32, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *PatchedWritableCircuitTerminationRequest) SetCircuit(v int32)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *PatchedWritableCircuitTerminationRequest) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetTermSide

`func (o *PatchedWritableCircuitTerminationRequest) GetTermSide() Termination`

GetTermSide returns the TermSide field if non-nil, zero value otherwise.

### GetTermSideOk

`func (o *PatchedWritableCircuitTerminationRequest) GetTermSideOk() (*Termination, bool)`

GetTermSideOk returns a tuple with the TermSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSide

`func (o *PatchedWritableCircuitTerminationRequest) SetTermSide(v Termination)`

SetTermSide sets TermSide field to given value.

### HasTermSide

`func (o *PatchedWritableCircuitTerminationRequest) HasTermSide() bool`

HasTermSide returns a boolean if a field has been set.

### GetSite

`func (o *PatchedWritableCircuitTerminationRequest) GetSite() int32`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableCircuitTerminationRequest) GetSiteOk() (*int32, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableCircuitTerminationRequest) SetSite(v int32)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableCircuitTerminationRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *PatchedWritableCircuitTerminationRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *PatchedWritableCircuitTerminationRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetProviderNetwork

`func (o *PatchedWritableCircuitTerminationRequest) GetProviderNetwork() int32`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *PatchedWritableCircuitTerminationRequest) GetProviderNetworkOk() (*int32, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *PatchedWritableCircuitTerminationRequest) SetProviderNetwork(v int32)`

SetProviderNetwork sets ProviderNetwork field to given value.

### HasProviderNetwork

`func (o *PatchedWritableCircuitTerminationRequest) HasProviderNetwork() bool`

HasProviderNetwork returns a boolean if a field has been set.

### SetProviderNetworkNil

`func (o *PatchedWritableCircuitTerminationRequest) SetProviderNetworkNil(b bool)`

 SetProviderNetworkNil sets the value for ProviderNetwork to be an explicit nil

### UnsetProviderNetwork
`func (o *PatchedWritableCircuitTerminationRequest) UnsetProviderNetwork()`

UnsetProviderNetwork ensures that no value is present for ProviderNetwork, not even an explicit nil
### GetPortSpeed

`func (o *PatchedWritableCircuitTerminationRequest) GetPortSpeed() int32`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *PatchedWritableCircuitTerminationRequest) GetPortSpeedOk() (*int32, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *PatchedWritableCircuitTerminationRequest) SetPortSpeed(v int32)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *PatchedWritableCircuitTerminationRequest) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### SetPortSpeedNil

`func (o *PatchedWritableCircuitTerminationRequest) SetPortSpeedNil(b bool)`

 SetPortSpeedNil sets the value for PortSpeed to be an explicit nil

### UnsetPortSpeed
`func (o *PatchedWritableCircuitTerminationRequest) UnsetPortSpeed()`

UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
### GetUpstreamSpeed

`func (o *PatchedWritableCircuitTerminationRequest) GetUpstreamSpeed() int32`

GetUpstreamSpeed returns the UpstreamSpeed field if non-nil, zero value otherwise.

### GetUpstreamSpeedOk

`func (o *PatchedWritableCircuitTerminationRequest) GetUpstreamSpeedOk() (*int32, bool)`

GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamSpeed

`func (o *PatchedWritableCircuitTerminationRequest) SetUpstreamSpeed(v int32)`

SetUpstreamSpeed sets UpstreamSpeed field to given value.

### HasUpstreamSpeed

`func (o *PatchedWritableCircuitTerminationRequest) HasUpstreamSpeed() bool`

HasUpstreamSpeed returns a boolean if a field has been set.

### SetUpstreamSpeedNil

`func (o *PatchedWritableCircuitTerminationRequest) SetUpstreamSpeedNil(b bool)`

 SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil

### UnsetUpstreamSpeed
`func (o *PatchedWritableCircuitTerminationRequest) UnsetUpstreamSpeed()`

UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
### GetXconnectId

`func (o *PatchedWritableCircuitTerminationRequest) GetXconnectId() string`

GetXconnectId returns the XconnectId field if non-nil, zero value otherwise.

### GetXconnectIdOk

`func (o *PatchedWritableCircuitTerminationRequest) GetXconnectIdOk() (*string, bool)`

GetXconnectIdOk returns a tuple with the XconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXconnectId

`func (o *PatchedWritableCircuitTerminationRequest) SetXconnectId(v string)`

SetXconnectId sets XconnectId field to given value.

### HasXconnectId

`func (o *PatchedWritableCircuitTerminationRequest) HasXconnectId() bool`

HasXconnectId returns a boolean if a field has been set.

### GetPpInfo

`func (o *PatchedWritableCircuitTerminationRequest) GetPpInfo() string`

GetPpInfo returns the PpInfo field if non-nil, zero value otherwise.

### GetPpInfoOk

`func (o *PatchedWritableCircuitTerminationRequest) GetPpInfoOk() (*string, bool)`

GetPpInfoOk returns a tuple with the PpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpInfo

`func (o *PatchedWritableCircuitTerminationRequest) SetPpInfo(v string)`

SetPpInfo sets PpInfo field to given value.

### HasPpInfo

`func (o *PatchedWritableCircuitTerminationRequest) HasPpInfo() bool`

HasPpInfo returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableCircuitTerminationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableCircuitTerminationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableCircuitTerminationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableCircuitTerminationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *PatchedWritableCircuitTerminationRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *PatchedWritableCircuitTerminationRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *PatchedWritableCircuitTerminationRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *PatchedWritableCircuitTerminationRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableCircuitTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableCircuitTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableCircuitTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableCircuitTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableCircuitTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableCircuitTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableCircuitTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableCircuitTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


