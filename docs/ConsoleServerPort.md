# ConsoleServerPort

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
**Type** | Pointer to [**ConsolePortType**](ConsolePortType.md) |  | [optional] 
**Speed** | Pointer to [**NullableConsolePortSpeed**](ConsolePortSpeed.md) |  | [optional] 
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

### NewConsoleServerPort

`func NewConsoleServerPort(id int32, url string, displayUrl string, display string, device BriefDevice, name string, cable NullableBriefCable, cableEnd string, linkPeers []interface{}, linkPeersType NullableString, connectedEndpoints []interface{}, connectedEndpointsType NullableString, connectedEndpointsReachable bool, created NullableTime, lastUpdated NullableTime, occupied bool, ) *ConsoleServerPort`

NewConsoleServerPort instantiates a new ConsoleServerPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleServerPortWithDefaults

`func NewConsoleServerPortWithDefaults() *ConsoleServerPort`

NewConsoleServerPortWithDefaults instantiates a new ConsoleServerPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleServerPort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleServerPort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleServerPort) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ConsoleServerPort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConsoleServerPort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConsoleServerPort) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *ConsoleServerPort) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *ConsoleServerPort) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *ConsoleServerPort) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *ConsoleServerPort) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConsoleServerPort) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConsoleServerPort) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDevice

`func (o *ConsoleServerPort) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ConsoleServerPort) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ConsoleServerPort) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.


### GetModule

`func (o *ConsoleServerPort) GetModule() BriefModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ConsoleServerPort) GetModuleOk() (*BriefModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ConsoleServerPort) SetModule(v BriefModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *ConsoleServerPort) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *ConsoleServerPort) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *ConsoleServerPort) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *ConsoleServerPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleServerPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleServerPort) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ConsoleServerPort) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConsoleServerPort) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConsoleServerPort) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConsoleServerPort) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *ConsoleServerPort) GetType() ConsolePortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsoleServerPort) GetTypeOk() (*ConsolePortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsoleServerPort) SetType(v ConsolePortType)`

SetType sets Type field to given value.

### HasType

`func (o *ConsoleServerPort) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSpeed

`func (o *ConsoleServerPort) GetSpeed() ConsolePortSpeed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ConsoleServerPort) GetSpeedOk() (*ConsolePortSpeed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ConsoleServerPort) SetSpeed(v ConsolePortSpeed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ConsoleServerPort) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *ConsoleServerPort) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *ConsoleServerPort) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetDescription

`func (o *ConsoleServerPort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleServerPort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleServerPort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleServerPort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *ConsoleServerPort) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *ConsoleServerPort) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *ConsoleServerPort) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *ConsoleServerPort) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *ConsoleServerPort) GetCable() BriefCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *ConsoleServerPort) GetCableOk() (*BriefCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *ConsoleServerPort) SetCable(v BriefCable)`

SetCable sets Cable field to given value.


### SetCableNil

`func (o *ConsoleServerPort) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *ConsoleServerPort) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetCableEnd

`func (o *ConsoleServerPort) GetCableEnd() string`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *ConsoleServerPort) GetCableEndOk() (*string, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *ConsoleServerPort) SetCableEnd(v string)`

SetCableEnd sets CableEnd field to given value.


### GetLinkPeers

`func (o *ConsoleServerPort) GetLinkPeers() []interface{}`

GetLinkPeers returns the LinkPeers field if non-nil, zero value otherwise.

### GetLinkPeersOk

`func (o *ConsoleServerPort) GetLinkPeersOk() (*[]interface{}, bool)`

GetLinkPeersOk returns a tuple with the LinkPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeers

`func (o *ConsoleServerPort) SetLinkPeers(v []interface{})`

SetLinkPeers sets LinkPeers field to given value.


### GetLinkPeersType

`func (o *ConsoleServerPort) GetLinkPeersType() string`

GetLinkPeersType returns the LinkPeersType field if non-nil, zero value otherwise.

### GetLinkPeersTypeOk

`func (o *ConsoleServerPort) GetLinkPeersTypeOk() (*string, bool)`

GetLinkPeersTypeOk returns a tuple with the LinkPeersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeersType

`func (o *ConsoleServerPort) SetLinkPeersType(v string)`

SetLinkPeersType sets LinkPeersType field to given value.


### SetLinkPeersTypeNil

`func (o *ConsoleServerPort) SetLinkPeersTypeNil(b bool)`

 SetLinkPeersTypeNil sets the value for LinkPeersType to be an explicit nil

### UnsetLinkPeersType
`func (o *ConsoleServerPort) UnsetLinkPeersType()`

UnsetLinkPeersType ensures that no value is present for LinkPeersType, not even an explicit nil
### GetConnectedEndpoints

`func (o *ConsoleServerPort) GetConnectedEndpoints() []interface{}`

GetConnectedEndpoints returns the ConnectedEndpoints field if non-nil, zero value otherwise.

### GetConnectedEndpointsOk

`func (o *ConsoleServerPort) GetConnectedEndpointsOk() (*[]interface{}, bool)`

GetConnectedEndpointsOk returns a tuple with the ConnectedEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoints

`func (o *ConsoleServerPort) SetConnectedEndpoints(v []interface{})`

SetConnectedEndpoints sets ConnectedEndpoints field to given value.


### SetConnectedEndpointsNil

`func (o *ConsoleServerPort) SetConnectedEndpointsNil(b bool)`

 SetConnectedEndpointsNil sets the value for ConnectedEndpoints to be an explicit nil

### UnsetConnectedEndpoints
`func (o *ConsoleServerPort) UnsetConnectedEndpoints()`

UnsetConnectedEndpoints ensures that no value is present for ConnectedEndpoints, not even an explicit nil
### GetConnectedEndpointsType

`func (o *ConsoleServerPort) GetConnectedEndpointsType() string`

GetConnectedEndpointsType returns the ConnectedEndpointsType field if non-nil, zero value otherwise.

### GetConnectedEndpointsTypeOk

`func (o *ConsoleServerPort) GetConnectedEndpointsTypeOk() (*string, bool)`

GetConnectedEndpointsTypeOk returns a tuple with the ConnectedEndpointsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointsType

`func (o *ConsoleServerPort) SetConnectedEndpointsType(v string)`

SetConnectedEndpointsType sets ConnectedEndpointsType field to given value.


### SetConnectedEndpointsTypeNil

`func (o *ConsoleServerPort) SetConnectedEndpointsTypeNil(b bool)`

 SetConnectedEndpointsTypeNil sets the value for ConnectedEndpointsType to be an explicit nil

### UnsetConnectedEndpointsType
`func (o *ConsoleServerPort) UnsetConnectedEndpointsType()`

UnsetConnectedEndpointsType ensures that no value is present for ConnectedEndpointsType, not even an explicit nil
### GetConnectedEndpointsReachable

`func (o *ConsoleServerPort) GetConnectedEndpointsReachable() bool`

GetConnectedEndpointsReachable returns the ConnectedEndpointsReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointsReachableOk

`func (o *ConsoleServerPort) GetConnectedEndpointsReachableOk() (*bool, bool)`

GetConnectedEndpointsReachableOk returns a tuple with the ConnectedEndpointsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointsReachable

`func (o *ConsoleServerPort) SetConnectedEndpointsReachable(v bool)`

SetConnectedEndpointsReachable sets ConnectedEndpointsReachable field to given value.


### GetTags

`func (o *ConsoleServerPort) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConsoleServerPort) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConsoleServerPort) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConsoleServerPort) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ConsoleServerPort) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ConsoleServerPort) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ConsoleServerPort) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ConsoleServerPort) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *ConsoleServerPort) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConsoleServerPort) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConsoleServerPort) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ConsoleServerPort) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ConsoleServerPort) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ConsoleServerPort) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ConsoleServerPort) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ConsoleServerPort) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ConsoleServerPort) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ConsoleServerPort) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetOccupied

`func (o *ConsoleServerPort) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *ConsoleServerPort) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *ConsoleServerPort) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


