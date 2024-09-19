# CircuitTermination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Circuit** | [**BriefCircuit**](BriefCircuit.md) |  | 
**TermSide** | [**Termination1**](Termination1.md) |  | 
**Site** | Pointer to [**NullableBriefSite**](BriefSite.md) |  | [optional] 
**ProviderNetwork** | Pointer to [**NullableBriefProviderNetwork**](BriefProviderNetwork.md) |  | [optional] 
**PortSpeed** | Pointer to **NullableInt32** | Physical circuit speed | [optional] 
**UpstreamSpeed** | Pointer to **NullableInt32** | Upstream speed, if different from port speed | [optional] 
**XconnectId** | Pointer to **string** | ID of the local cross-connect | [optional] 
**PpInfo** | Pointer to **string** | Patch panel ID and port number(s) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Cable** | [**NullableBriefCable**](BriefCable.md) |  | [readonly] 
**CableEnd** | **string** |  | [readonly] 
**LinkPeers** | **[]interface{}** |  | [readonly] 
**LinkPeersType** | **NullableString** | Return the type of the peer link terminations, or None. | [readonly] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Occupied** | **bool** |  | [readonly] 

## Methods

### NewCircuitTermination

`func NewCircuitTermination(id int32, url string, displayUrl string, display string, circuit BriefCircuit, termSide Termination1, cable NullableBriefCable, cableEnd string, linkPeers []interface{}, linkPeersType NullableString, created NullableTime, lastUpdated NullableTime, occupied bool, ) *CircuitTermination`

NewCircuitTermination instantiates a new CircuitTermination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitTerminationWithDefaults

`func NewCircuitTerminationWithDefaults() *CircuitTermination`

NewCircuitTerminationWithDefaults instantiates a new CircuitTermination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CircuitTermination) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CircuitTermination) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CircuitTermination) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *CircuitTermination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CircuitTermination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CircuitTermination) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *CircuitTermination) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *CircuitTermination) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *CircuitTermination) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *CircuitTermination) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CircuitTermination) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CircuitTermination) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetCircuit

`func (o *CircuitTermination) GetCircuit() BriefCircuit`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *CircuitTermination) GetCircuitOk() (*BriefCircuit, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *CircuitTermination) SetCircuit(v BriefCircuit)`

SetCircuit sets Circuit field to given value.


### GetTermSide

`func (o *CircuitTermination) GetTermSide() Termination1`

GetTermSide returns the TermSide field if non-nil, zero value otherwise.

### GetTermSideOk

`func (o *CircuitTermination) GetTermSideOk() (*Termination1, bool)`

GetTermSideOk returns a tuple with the TermSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSide

`func (o *CircuitTermination) SetTermSide(v Termination1)`

SetTermSide sets TermSide field to given value.


### GetSite

`func (o *CircuitTermination) GetSite() BriefSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CircuitTermination) GetSiteOk() (*BriefSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CircuitTermination) SetSite(v BriefSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *CircuitTermination) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *CircuitTermination) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *CircuitTermination) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetProviderNetwork

`func (o *CircuitTermination) GetProviderNetwork() BriefProviderNetwork`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *CircuitTermination) GetProviderNetworkOk() (*BriefProviderNetwork, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *CircuitTermination) SetProviderNetwork(v BriefProviderNetwork)`

SetProviderNetwork sets ProviderNetwork field to given value.

### HasProviderNetwork

`func (o *CircuitTermination) HasProviderNetwork() bool`

HasProviderNetwork returns a boolean if a field has been set.

### SetProviderNetworkNil

`func (o *CircuitTermination) SetProviderNetworkNil(b bool)`

 SetProviderNetworkNil sets the value for ProviderNetwork to be an explicit nil

### UnsetProviderNetwork
`func (o *CircuitTermination) UnsetProviderNetwork()`

UnsetProviderNetwork ensures that no value is present for ProviderNetwork, not even an explicit nil
### GetPortSpeed

`func (o *CircuitTermination) GetPortSpeed() int32`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *CircuitTermination) GetPortSpeedOk() (*int32, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *CircuitTermination) SetPortSpeed(v int32)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *CircuitTermination) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### SetPortSpeedNil

`func (o *CircuitTermination) SetPortSpeedNil(b bool)`

 SetPortSpeedNil sets the value for PortSpeed to be an explicit nil

### UnsetPortSpeed
`func (o *CircuitTermination) UnsetPortSpeed()`

UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
### GetUpstreamSpeed

`func (o *CircuitTermination) GetUpstreamSpeed() int32`

GetUpstreamSpeed returns the UpstreamSpeed field if non-nil, zero value otherwise.

### GetUpstreamSpeedOk

`func (o *CircuitTermination) GetUpstreamSpeedOk() (*int32, bool)`

GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamSpeed

`func (o *CircuitTermination) SetUpstreamSpeed(v int32)`

SetUpstreamSpeed sets UpstreamSpeed field to given value.

### HasUpstreamSpeed

`func (o *CircuitTermination) HasUpstreamSpeed() bool`

HasUpstreamSpeed returns a boolean if a field has been set.

### SetUpstreamSpeedNil

`func (o *CircuitTermination) SetUpstreamSpeedNil(b bool)`

 SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil

### UnsetUpstreamSpeed
`func (o *CircuitTermination) UnsetUpstreamSpeed()`

UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
### GetXconnectId

`func (o *CircuitTermination) GetXconnectId() string`

GetXconnectId returns the XconnectId field if non-nil, zero value otherwise.

### GetXconnectIdOk

`func (o *CircuitTermination) GetXconnectIdOk() (*string, bool)`

GetXconnectIdOk returns a tuple with the XconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXconnectId

`func (o *CircuitTermination) SetXconnectId(v string)`

SetXconnectId sets XconnectId field to given value.

### HasXconnectId

`func (o *CircuitTermination) HasXconnectId() bool`

HasXconnectId returns a boolean if a field has been set.

### GetPpInfo

`func (o *CircuitTermination) GetPpInfo() string`

GetPpInfo returns the PpInfo field if non-nil, zero value otherwise.

### GetPpInfoOk

`func (o *CircuitTermination) GetPpInfoOk() (*string, bool)`

GetPpInfoOk returns a tuple with the PpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpInfo

`func (o *CircuitTermination) SetPpInfo(v string)`

SetPpInfo sets PpInfo field to given value.

### HasPpInfo

`func (o *CircuitTermination) HasPpInfo() bool`

HasPpInfo returns a boolean if a field has been set.

### GetDescription

`func (o *CircuitTermination) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CircuitTermination) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CircuitTermination) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CircuitTermination) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *CircuitTermination) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *CircuitTermination) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *CircuitTermination) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *CircuitTermination) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *CircuitTermination) GetCable() BriefCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *CircuitTermination) GetCableOk() (*BriefCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *CircuitTermination) SetCable(v BriefCable)`

SetCable sets Cable field to given value.


### SetCableNil

`func (o *CircuitTermination) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *CircuitTermination) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetCableEnd

`func (o *CircuitTermination) GetCableEnd() string`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *CircuitTermination) GetCableEndOk() (*string, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *CircuitTermination) SetCableEnd(v string)`

SetCableEnd sets CableEnd field to given value.


### GetLinkPeers

`func (o *CircuitTermination) GetLinkPeers() []interface{}`

GetLinkPeers returns the LinkPeers field if non-nil, zero value otherwise.

### GetLinkPeersOk

`func (o *CircuitTermination) GetLinkPeersOk() (*[]interface{}, bool)`

GetLinkPeersOk returns a tuple with the LinkPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeers

`func (o *CircuitTermination) SetLinkPeers(v []interface{})`

SetLinkPeers sets LinkPeers field to given value.


### GetLinkPeersType

`func (o *CircuitTermination) GetLinkPeersType() string`

GetLinkPeersType returns the LinkPeersType field if non-nil, zero value otherwise.

### GetLinkPeersTypeOk

`func (o *CircuitTermination) GetLinkPeersTypeOk() (*string, bool)`

GetLinkPeersTypeOk returns a tuple with the LinkPeersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeersType

`func (o *CircuitTermination) SetLinkPeersType(v string)`

SetLinkPeersType sets LinkPeersType field to given value.


### SetLinkPeersTypeNil

`func (o *CircuitTermination) SetLinkPeersTypeNil(b bool)`

 SetLinkPeersTypeNil sets the value for LinkPeersType to be an explicit nil

### UnsetLinkPeersType
`func (o *CircuitTermination) UnsetLinkPeersType()`

UnsetLinkPeersType ensures that no value is present for LinkPeersType, not even an explicit nil
### GetTags

`func (o *CircuitTermination) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CircuitTermination) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CircuitTermination) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CircuitTermination) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *CircuitTermination) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CircuitTermination) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CircuitTermination) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CircuitTermination) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *CircuitTermination) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CircuitTermination) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CircuitTermination) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *CircuitTermination) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CircuitTermination) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *CircuitTermination) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CircuitTermination) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CircuitTermination) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *CircuitTermination) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CircuitTermination) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetOccupied

`func (o *CircuitTermination) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *CircuitTermination) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *CircuitTermination) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


