# PowerFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**PowerPanel** | [**BriefPowerPanel**](BriefPowerPanel.md) |  | 
**Rack** | Pointer to [**NullableBriefRack**](BriefRack.md) |  | [optional] 
**Name** | **string** |  | 
**Status** | Pointer to [**PowerFeedStatus**](PowerFeedStatus.md) |  | [optional] 
**Type** | Pointer to [**PowerFeedType**](PowerFeedType.md) |  | [optional] 
**Supply** | Pointer to [**PowerFeedSupply**](PowerFeedSupply.md) |  | [optional] 
**Phase** | Pointer to [**PowerFeedPhase**](PowerFeedPhase.md) |  | [optional] 
**Voltage** | Pointer to **int32** |  | [optional] 
**Amperage** | Pointer to **int32** |  | [optional] 
**MaxUtilization** | Pointer to **int32** | Maximum permissible draw (percentage) | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Cable** | [**NullableBriefCable**](BriefCable.md) |  | [readonly] 
**CableEnd** | **string** |  | [readonly] 
**LinkPeers** | **[]interface{}** |  | [readonly] 
**LinkPeersType** | **NullableString** | Return the type of the peer link terminations, or None. | [readonly] 
**ConnectedEndpoints** | **[]interface{}** |  | [readonly] 
**ConnectedEndpointsType** | **NullableString** |  | [readonly] 
**ConnectedEndpointsReachable** | **bool** |  | [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**NullableBriefTenant**](BriefTenant.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Occupied** | **bool** |  | [readonly] 

## Methods

### NewPowerFeed

`func NewPowerFeed(id int32, url string, displayUrl string, display string, powerPanel BriefPowerPanel, name string, cable NullableBriefCable, cableEnd string, linkPeers []interface{}, linkPeersType NullableString, connectedEndpoints []interface{}, connectedEndpointsType NullableString, connectedEndpointsReachable bool, created NullableTime, lastUpdated NullableTime, occupied bool, ) *PowerFeed`

NewPowerFeed instantiates a new PowerFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerFeedWithDefaults

`func NewPowerFeedWithDefaults() *PowerFeed`

NewPowerFeedWithDefaults instantiates a new PowerFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerFeed) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerFeed) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerFeed) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *PowerFeed) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PowerFeed) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PowerFeed) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *PowerFeed) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *PowerFeed) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *PowerFeed) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *PowerFeed) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PowerFeed) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PowerFeed) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetPowerPanel

`func (o *PowerFeed) GetPowerPanel() BriefPowerPanel`

GetPowerPanel returns the PowerPanel field if non-nil, zero value otherwise.

### GetPowerPanelOk

`func (o *PowerFeed) GetPowerPanelOk() (*BriefPowerPanel, bool)`

GetPowerPanelOk returns a tuple with the PowerPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPanel

`func (o *PowerFeed) SetPowerPanel(v BriefPowerPanel)`

SetPowerPanel sets PowerPanel field to given value.


### GetRack

`func (o *PowerFeed) GetRack() BriefRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PowerFeed) GetRackOk() (*BriefRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PowerFeed) SetRack(v BriefRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PowerFeed) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *PowerFeed) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *PowerFeed) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetName

`func (o *PowerFeed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerFeed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerFeed) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *PowerFeed) GetStatus() PowerFeedStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PowerFeed) GetStatusOk() (*PowerFeedStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PowerFeed) SetStatus(v PowerFeedStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PowerFeed) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *PowerFeed) GetType() PowerFeedType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PowerFeed) GetTypeOk() (*PowerFeedType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PowerFeed) SetType(v PowerFeedType)`

SetType sets Type field to given value.

### HasType

`func (o *PowerFeed) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSupply

`func (o *PowerFeed) GetSupply() PowerFeedSupply`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *PowerFeed) GetSupplyOk() (*PowerFeedSupply, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *PowerFeed) SetSupply(v PowerFeedSupply)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *PowerFeed) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetPhase

`func (o *PowerFeed) GetPhase() PowerFeedPhase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PowerFeed) GetPhaseOk() (*PowerFeedPhase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PowerFeed) SetPhase(v PowerFeedPhase)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PowerFeed) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetVoltage

`func (o *PowerFeed) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *PowerFeed) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *PowerFeed) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *PowerFeed) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetAmperage

`func (o *PowerFeed) GetAmperage() int32`

GetAmperage returns the Amperage field if non-nil, zero value otherwise.

### GetAmperageOk

`func (o *PowerFeed) GetAmperageOk() (*int32, bool)`

GetAmperageOk returns a tuple with the Amperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmperage

`func (o *PowerFeed) SetAmperage(v int32)`

SetAmperage sets Amperage field to given value.

### HasAmperage

`func (o *PowerFeed) HasAmperage() bool`

HasAmperage returns a boolean if a field has been set.

### GetMaxUtilization

`func (o *PowerFeed) GetMaxUtilization() int32`

GetMaxUtilization returns the MaxUtilization field if non-nil, zero value otherwise.

### GetMaxUtilizationOk

`func (o *PowerFeed) GetMaxUtilizationOk() (*int32, bool)`

GetMaxUtilizationOk returns a tuple with the MaxUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUtilization

`func (o *PowerFeed) SetMaxUtilization(v int32)`

SetMaxUtilization sets MaxUtilization field to given value.

### HasMaxUtilization

`func (o *PowerFeed) HasMaxUtilization() bool`

HasMaxUtilization returns a boolean if a field has been set.

### GetMarkConnected

`func (o *PowerFeed) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *PowerFeed) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *PowerFeed) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *PowerFeed) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *PowerFeed) GetCable() BriefCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *PowerFeed) GetCableOk() (*BriefCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *PowerFeed) SetCable(v BriefCable)`

SetCable sets Cable field to given value.


### SetCableNil

`func (o *PowerFeed) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *PowerFeed) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetCableEnd

`func (o *PowerFeed) GetCableEnd() string`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *PowerFeed) GetCableEndOk() (*string, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *PowerFeed) SetCableEnd(v string)`

SetCableEnd sets CableEnd field to given value.


### GetLinkPeers

`func (o *PowerFeed) GetLinkPeers() []interface{}`

GetLinkPeers returns the LinkPeers field if non-nil, zero value otherwise.

### GetLinkPeersOk

`func (o *PowerFeed) GetLinkPeersOk() (*[]interface{}, bool)`

GetLinkPeersOk returns a tuple with the LinkPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeers

`func (o *PowerFeed) SetLinkPeers(v []interface{})`

SetLinkPeers sets LinkPeers field to given value.


### GetLinkPeersType

`func (o *PowerFeed) GetLinkPeersType() string`

GetLinkPeersType returns the LinkPeersType field if non-nil, zero value otherwise.

### GetLinkPeersTypeOk

`func (o *PowerFeed) GetLinkPeersTypeOk() (*string, bool)`

GetLinkPeersTypeOk returns a tuple with the LinkPeersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeersType

`func (o *PowerFeed) SetLinkPeersType(v string)`

SetLinkPeersType sets LinkPeersType field to given value.


### SetLinkPeersTypeNil

`func (o *PowerFeed) SetLinkPeersTypeNil(b bool)`

 SetLinkPeersTypeNil sets the value for LinkPeersType to be an explicit nil

### UnsetLinkPeersType
`func (o *PowerFeed) UnsetLinkPeersType()`

UnsetLinkPeersType ensures that no value is present for LinkPeersType, not even an explicit nil
### GetConnectedEndpoints

`func (o *PowerFeed) GetConnectedEndpoints() []interface{}`

GetConnectedEndpoints returns the ConnectedEndpoints field if non-nil, zero value otherwise.

### GetConnectedEndpointsOk

`func (o *PowerFeed) GetConnectedEndpointsOk() (*[]interface{}, bool)`

GetConnectedEndpointsOk returns a tuple with the ConnectedEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoints

`func (o *PowerFeed) SetConnectedEndpoints(v []interface{})`

SetConnectedEndpoints sets ConnectedEndpoints field to given value.


### SetConnectedEndpointsNil

`func (o *PowerFeed) SetConnectedEndpointsNil(b bool)`

 SetConnectedEndpointsNil sets the value for ConnectedEndpoints to be an explicit nil

### UnsetConnectedEndpoints
`func (o *PowerFeed) UnsetConnectedEndpoints()`

UnsetConnectedEndpoints ensures that no value is present for ConnectedEndpoints, not even an explicit nil
### GetConnectedEndpointsType

`func (o *PowerFeed) GetConnectedEndpointsType() string`

GetConnectedEndpointsType returns the ConnectedEndpointsType field if non-nil, zero value otherwise.

### GetConnectedEndpointsTypeOk

`func (o *PowerFeed) GetConnectedEndpointsTypeOk() (*string, bool)`

GetConnectedEndpointsTypeOk returns a tuple with the ConnectedEndpointsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointsType

`func (o *PowerFeed) SetConnectedEndpointsType(v string)`

SetConnectedEndpointsType sets ConnectedEndpointsType field to given value.


### SetConnectedEndpointsTypeNil

`func (o *PowerFeed) SetConnectedEndpointsTypeNil(b bool)`

 SetConnectedEndpointsTypeNil sets the value for ConnectedEndpointsType to be an explicit nil

### UnsetConnectedEndpointsType
`func (o *PowerFeed) UnsetConnectedEndpointsType()`

UnsetConnectedEndpointsType ensures that no value is present for ConnectedEndpointsType, not even an explicit nil
### GetConnectedEndpointsReachable

`func (o *PowerFeed) GetConnectedEndpointsReachable() bool`

GetConnectedEndpointsReachable returns the ConnectedEndpointsReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointsReachableOk

`func (o *PowerFeed) GetConnectedEndpointsReachableOk() (*bool, bool)`

GetConnectedEndpointsReachableOk returns a tuple with the ConnectedEndpointsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointsReachable

`func (o *PowerFeed) SetConnectedEndpointsReachable(v bool)`

SetConnectedEndpointsReachable sets ConnectedEndpointsReachable field to given value.


### GetDescription

`func (o *PowerFeed) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PowerFeed) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PowerFeed) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PowerFeed) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTenant

`func (o *PowerFeed) GetTenant() BriefTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PowerFeed) GetTenantOk() (*BriefTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PowerFeed) SetTenant(v BriefTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PowerFeed) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PowerFeed) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PowerFeed) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetComments

`func (o *PowerFeed) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PowerFeed) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PowerFeed) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PowerFeed) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PowerFeed) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PowerFeed) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PowerFeed) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PowerFeed) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PowerFeed) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PowerFeed) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PowerFeed) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PowerFeed) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *PowerFeed) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PowerFeed) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PowerFeed) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *PowerFeed) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *PowerFeed) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *PowerFeed) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PowerFeed) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PowerFeed) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *PowerFeed) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *PowerFeed) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetOccupied

`func (o *PowerFeed) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *PowerFeed) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *PowerFeed) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


