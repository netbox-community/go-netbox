# PowerPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Device** | [**BriefDevice**](BriefDevice.md) |  | 
**Module** | Pointer to [**NullableBriefModule**](BriefModule.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | Pointer to [**NullablePowerPortType**](PowerPortType.md) |  | [optional] 
**MaximumDraw** | Pointer to **NullableInt32** | Maximum power draw (watts) | [optional] 
**AllocatedDraw** | Pointer to **NullableInt32** | Allocated power draw (watts) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**Cable** | [**NullableBriefCable**](BriefCable.md) |  | [readonly] 
**CableEnd** | **string** |  | [readonly] 
**LinkPeers** | **[]interface{}** |  | [readonly] 
**LinkPeersType** | **NullableString** | Return the type of the peer link terminations, or None. | [readonly] 
**ConnectedEndpoints** | **[]interface{}** |  | [readonly] 
**ConnectedEndpointsType** | **NullableString** |  | [readonly] 
**ConnectedEndpointsReachable** | **bool** |  | [readonly] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Occupied** | **bool** |  | [readonly] 

## Methods

### NewPowerPort

`func NewPowerPort(id int32, url string, displayUrl string, display string, device BriefDevice, name string, cable NullableBriefCable, cableEnd string, linkPeers []interface{}, linkPeersType NullableString, connectedEndpoints []interface{}, connectedEndpointsType NullableString, connectedEndpointsReachable bool, created NullableTime, lastUpdated NullableTime, occupied bool, ) *PowerPort`

NewPowerPort instantiates a new PowerPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerPortWithDefaults

`func NewPowerPortWithDefaults() *PowerPort`

NewPowerPortWithDefaults instantiates a new PowerPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerPort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerPort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerPort) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *PowerPort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PowerPort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PowerPort) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *PowerPort) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *PowerPort) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *PowerPort) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *PowerPort) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PowerPort) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PowerPort) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDevice

`func (o *PowerPort) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PowerPort) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PowerPort) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.


### GetModule

`func (o *PowerPort) GetModule() BriefModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *PowerPort) GetModuleOk() (*BriefModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *PowerPort) SetModule(v BriefModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *PowerPort) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *PowerPort) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *PowerPort) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *PowerPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerPort) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *PowerPort) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PowerPort) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PowerPort) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PowerPort) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PowerPort) GetType() PowerPortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PowerPort) GetTypeOk() (*PowerPortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PowerPort) SetType(v PowerPortType)`

SetType sets Type field to given value.

### HasType

`func (o *PowerPort) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PowerPort) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PowerPort) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMaximumDraw

`func (o *PowerPort) GetMaximumDraw() int32`

GetMaximumDraw returns the MaximumDraw field if non-nil, zero value otherwise.

### GetMaximumDrawOk

`func (o *PowerPort) GetMaximumDrawOk() (*int32, bool)`

GetMaximumDrawOk returns a tuple with the MaximumDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDraw

`func (o *PowerPort) SetMaximumDraw(v int32)`

SetMaximumDraw sets MaximumDraw field to given value.

### HasMaximumDraw

`func (o *PowerPort) HasMaximumDraw() bool`

HasMaximumDraw returns a boolean if a field has been set.

### SetMaximumDrawNil

`func (o *PowerPort) SetMaximumDrawNil(b bool)`

 SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil

### UnsetMaximumDraw
`func (o *PowerPort) UnsetMaximumDraw()`

UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
### GetAllocatedDraw

`func (o *PowerPort) GetAllocatedDraw() int32`

GetAllocatedDraw returns the AllocatedDraw field if non-nil, zero value otherwise.

### GetAllocatedDrawOk

`func (o *PowerPort) GetAllocatedDrawOk() (*int32, bool)`

GetAllocatedDrawOk returns a tuple with the AllocatedDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedDraw

`func (o *PowerPort) SetAllocatedDraw(v int32)`

SetAllocatedDraw sets AllocatedDraw field to given value.

### HasAllocatedDraw

`func (o *PowerPort) HasAllocatedDraw() bool`

HasAllocatedDraw returns a boolean if a field has been set.

### SetAllocatedDrawNil

`func (o *PowerPort) SetAllocatedDrawNil(b bool)`

 SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil

### UnsetAllocatedDraw
`func (o *PowerPort) UnsetAllocatedDraw()`

UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
### GetDescription

`func (o *PowerPort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PowerPort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PowerPort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PowerPort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *PowerPort) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *PowerPort) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *PowerPort) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *PowerPort) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *PowerPort) GetCable() BriefCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *PowerPort) GetCableOk() (*BriefCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *PowerPort) SetCable(v BriefCable)`

SetCable sets Cable field to given value.


### SetCableNil

`func (o *PowerPort) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *PowerPort) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetCableEnd

`func (o *PowerPort) GetCableEnd() string`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *PowerPort) GetCableEndOk() (*string, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *PowerPort) SetCableEnd(v string)`

SetCableEnd sets CableEnd field to given value.


### GetLinkPeers

`func (o *PowerPort) GetLinkPeers() []interface{}`

GetLinkPeers returns the LinkPeers field if non-nil, zero value otherwise.

### GetLinkPeersOk

`func (o *PowerPort) GetLinkPeersOk() (*[]interface{}, bool)`

GetLinkPeersOk returns a tuple with the LinkPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeers

`func (o *PowerPort) SetLinkPeers(v []interface{})`

SetLinkPeers sets LinkPeers field to given value.


### GetLinkPeersType

`func (o *PowerPort) GetLinkPeersType() string`

GetLinkPeersType returns the LinkPeersType field if non-nil, zero value otherwise.

### GetLinkPeersTypeOk

`func (o *PowerPort) GetLinkPeersTypeOk() (*string, bool)`

GetLinkPeersTypeOk returns a tuple with the LinkPeersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeersType

`func (o *PowerPort) SetLinkPeersType(v string)`

SetLinkPeersType sets LinkPeersType field to given value.


### SetLinkPeersTypeNil

`func (o *PowerPort) SetLinkPeersTypeNil(b bool)`

 SetLinkPeersTypeNil sets the value for LinkPeersType to be an explicit nil

### UnsetLinkPeersType
`func (o *PowerPort) UnsetLinkPeersType()`

UnsetLinkPeersType ensures that no value is present for LinkPeersType, not even an explicit nil
### GetConnectedEndpoints

`func (o *PowerPort) GetConnectedEndpoints() []interface{}`

GetConnectedEndpoints returns the ConnectedEndpoints field if non-nil, zero value otherwise.

### GetConnectedEndpointsOk

`func (o *PowerPort) GetConnectedEndpointsOk() (*[]interface{}, bool)`

GetConnectedEndpointsOk returns a tuple with the ConnectedEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoints

`func (o *PowerPort) SetConnectedEndpoints(v []interface{})`

SetConnectedEndpoints sets ConnectedEndpoints field to given value.


### SetConnectedEndpointsNil

`func (o *PowerPort) SetConnectedEndpointsNil(b bool)`

 SetConnectedEndpointsNil sets the value for ConnectedEndpoints to be an explicit nil

### UnsetConnectedEndpoints
`func (o *PowerPort) UnsetConnectedEndpoints()`

UnsetConnectedEndpoints ensures that no value is present for ConnectedEndpoints, not even an explicit nil
### GetConnectedEndpointsType

`func (o *PowerPort) GetConnectedEndpointsType() string`

GetConnectedEndpointsType returns the ConnectedEndpointsType field if non-nil, zero value otherwise.

### GetConnectedEndpointsTypeOk

`func (o *PowerPort) GetConnectedEndpointsTypeOk() (*string, bool)`

GetConnectedEndpointsTypeOk returns a tuple with the ConnectedEndpointsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointsType

`func (o *PowerPort) SetConnectedEndpointsType(v string)`

SetConnectedEndpointsType sets ConnectedEndpointsType field to given value.


### SetConnectedEndpointsTypeNil

`func (o *PowerPort) SetConnectedEndpointsTypeNil(b bool)`

 SetConnectedEndpointsTypeNil sets the value for ConnectedEndpointsType to be an explicit nil

### UnsetConnectedEndpointsType
`func (o *PowerPort) UnsetConnectedEndpointsType()`

UnsetConnectedEndpointsType ensures that no value is present for ConnectedEndpointsType, not even an explicit nil
### GetConnectedEndpointsReachable

`func (o *PowerPort) GetConnectedEndpointsReachable() bool`

GetConnectedEndpointsReachable returns the ConnectedEndpointsReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointsReachableOk

`func (o *PowerPort) GetConnectedEndpointsReachableOk() (*bool, bool)`

GetConnectedEndpointsReachableOk returns a tuple with the ConnectedEndpointsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointsReachable

`func (o *PowerPort) SetConnectedEndpointsReachable(v bool)`

SetConnectedEndpointsReachable sets ConnectedEndpointsReachable field to given value.


### GetTags

`func (o *PowerPort) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PowerPort) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PowerPort) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PowerPort) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PowerPort) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PowerPort) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PowerPort) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PowerPort) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *PowerPort) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PowerPort) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PowerPort) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *PowerPort) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *PowerPort) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *PowerPort) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PowerPort) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PowerPort) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *PowerPort) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *PowerPort) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetOccupied

`func (o *PowerPort) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *PowerPort) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *PowerPort) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


