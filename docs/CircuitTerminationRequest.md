# CircuitTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Circuit** | [**BriefCircuitRequest**](BriefCircuitRequest.md) |  | 
**TermSide** | [**Termination1**](Termination1.md) |  | 
**Site** | Pointer to [**NullableBriefSiteRequest**](BriefSiteRequest.md) |  | [optional] 
**ProviderNetwork** | Pointer to [**NullableBriefProviderNetworkRequest**](BriefProviderNetworkRequest.md) |  | [optional] 
**PortSpeed** | Pointer to **NullableInt32** | Physical circuit speed | [optional] 
**UpstreamSpeed** | Pointer to **NullableInt32** | Upstream speed, if different from port speed | [optional] 
**XconnectId** | Pointer to **string** | ID of the local cross-connect | [optional] 
**PpInfo** | Pointer to **string** | Patch panel ID and port number(s) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCircuitTerminationRequest

`func NewCircuitTerminationRequest(circuit BriefCircuitRequest, termSide Termination1, ) *CircuitTerminationRequest`

NewCircuitTerminationRequest instantiates a new CircuitTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitTerminationRequestWithDefaults

`func NewCircuitTerminationRequestWithDefaults() *CircuitTerminationRequest`

NewCircuitTerminationRequestWithDefaults instantiates a new CircuitTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuit

`func (o *CircuitTerminationRequest) GetCircuit() BriefCircuitRequest`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *CircuitTerminationRequest) GetCircuitOk() (*BriefCircuitRequest, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *CircuitTerminationRequest) SetCircuit(v BriefCircuitRequest)`

SetCircuit sets Circuit field to given value.


### GetTermSide

`func (o *CircuitTerminationRequest) GetTermSide() Termination1`

GetTermSide returns the TermSide field if non-nil, zero value otherwise.

### GetTermSideOk

`func (o *CircuitTerminationRequest) GetTermSideOk() (*Termination1, bool)`

GetTermSideOk returns a tuple with the TermSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSide

`func (o *CircuitTerminationRequest) SetTermSide(v Termination1)`

SetTermSide sets TermSide field to given value.


### GetSite

`func (o *CircuitTerminationRequest) GetSite() BriefSiteRequest`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CircuitTerminationRequest) GetSiteOk() (*BriefSiteRequest, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CircuitTerminationRequest) SetSite(v BriefSiteRequest)`

SetSite sets Site field to given value.

### HasSite

`func (o *CircuitTerminationRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *CircuitTerminationRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *CircuitTerminationRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetProviderNetwork

`func (o *CircuitTerminationRequest) GetProviderNetwork() BriefProviderNetworkRequest`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *CircuitTerminationRequest) GetProviderNetworkOk() (*BriefProviderNetworkRequest, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *CircuitTerminationRequest) SetProviderNetwork(v BriefProviderNetworkRequest)`

SetProviderNetwork sets ProviderNetwork field to given value.

### HasProviderNetwork

`func (o *CircuitTerminationRequest) HasProviderNetwork() bool`

HasProviderNetwork returns a boolean if a field has been set.

### SetProviderNetworkNil

`func (o *CircuitTerminationRequest) SetProviderNetworkNil(b bool)`

 SetProviderNetworkNil sets the value for ProviderNetwork to be an explicit nil

### UnsetProviderNetwork
`func (o *CircuitTerminationRequest) UnsetProviderNetwork()`

UnsetProviderNetwork ensures that no value is present for ProviderNetwork, not even an explicit nil
### GetPortSpeed

`func (o *CircuitTerminationRequest) GetPortSpeed() int32`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *CircuitTerminationRequest) GetPortSpeedOk() (*int32, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *CircuitTerminationRequest) SetPortSpeed(v int32)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *CircuitTerminationRequest) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### SetPortSpeedNil

`func (o *CircuitTerminationRequest) SetPortSpeedNil(b bool)`

 SetPortSpeedNil sets the value for PortSpeed to be an explicit nil

### UnsetPortSpeed
`func (o *CircuitTerminationRequest) UnsetPortSpeed()`

UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
### GetUpstreamSpeed

`func (o *CircuitTerminationRequest) GetUpstreamSpeed() int32`

GetUpstreamSpeed returns the UpstreamSpeed field if non-nil, zero value otherwise.

### GetUpstreamSpeedOk

`func (o *CircuitTerminationRequest) GetUpstreamSpeedOk() (*int32, bool)`

GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamSpeed

`func (o *CircuitTerminationRequest) SetUpstreamSpeed(v int32)`

SetUpstreamSpeed sets UpstreamSpeed field to given value.

### HasUpstreamSpeed

`func (o *CircuitTerminationRequest) HasUpstreamSpeed() bool`

HasUpstreamSpeed returns a boolean if a field has been set.

### SetUpstreamSpeedNil

`func (o *CircuitTerminationRequest) SetUpstreamSpeedNil(b bool)`

 SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil

### UnsetUpstreamSpeed
`func (o *CircuitTerminationRequest) UnsetUpstreamSpeed()`

UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
### GetXconnectId

`func (o *CircuitTerminationRequest) GetXconnectId() string`

GetXconnectId returns the XconnectId field if non-nil, zero value otherwise.

### GetXconnectIdOk

`func (o *CircuitTerminationRequest) GetXconnectIdOk() (*string, bool)`

GetXconnectIdOk returns a tuple with the XconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXconnectId

`func (o *CircuitTerminationRequest) SetXconnectId(v string)`

SetXconnectId sets XconnectId field to given value.

### HasXconnectId

`func (o *CircuitTerminationRequest) HasXconnectId() bool`

HasXconnectId returns a boolean if a field has been set.

### GetPpInfo

`func (o *CircuitTerminationRequest) GetPpInfo() string`

GetPpInfo returns the PpInfo field if non-nil, zero value otherwise.

### GetPpInfoOk

`func (o *CircuitTerminationRequest) GetPpInfoOk() (*string, bool)`

GetPpInfoOk returns a tuple with the PpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpInfo

`func (o *CircuitTerminationRequest) SetPpInfo(v string)`

SetPpInfo sets PpInfo field to given value.

### HasPpInfo

`func (o *CircuitTerminationRequest) HasPpInfo() bool`

HasPpInfo returns a boolean if a field has been set.

### GetDescription

`func (o *CircuitTerminationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CircuitTerminationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CircuitTerminationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CircuitTerminationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *CircuitTerminationRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *CircuitTerminationRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *CircuitTerminationRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *CircuitTerminationRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetTags

`func (o *CircuitTerminationRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CircuitTerminationRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CircuitTerminationRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CircuitTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *CircuitTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CircuitTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CircuitTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CircuitTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


