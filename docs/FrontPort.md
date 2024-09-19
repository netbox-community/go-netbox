# FrontPort

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
**Type** | [**FrontPortType**](FrontPortType.md) |  | 
**Color** | Pointer to **string** |  | [optional] 
**RearPort** | [**FrontPortRearPort**](FrontPortRearPort.md) |  | 
**RearPortPosition** | Pointer to **int32** | Mapped position on corresponding rear port | [optional] [default to 1]
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

### NewFrontPort

`func NewFrontPort(id int32, url string, displayUrl string, display string, device BriefDevice, name string, type_ FrontPortType, rearPort FrontPortRearPort, cable NullableBriefCable, cableEnd string, linkPeers []interface{}, linkPeersType NullableString, created NullableTime, lastUpdated NullableTime, occupied bool, ) *FrontPort`

NewFrontPort instantiates a new FrontPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrontPortWithDefaults

`func NewFrontPortWithDefaults() *FrontPort`

NewFrontPortWithDefaults instantiates a new FrontPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FrontPort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FrontPort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FrontPort) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *FrontPort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FrontPort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FrontPort) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *FrontPort) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *FrontPort) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *FrontPort) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *FrontPort) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *FrontPort) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *FrontPort) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetDevice

`func (o *FrontPort) GetDevice() BriefDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FrontPort) GetDeviceOk() (*BriefDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FrontPort) SetDevice(v BriefDevice)`

SetDevice sets Device field to given value.


### GetModule

`func (o *FrontPort) GetModule() BriefModule`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *FrontPort) GetModuleOk() (*BriefModule, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *FrontPort) SetModule(v BriefModule)`

SetModule sets Module field to given value.

### HasModule

`func (o *FrontPort) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *FrontPort) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *FrontPort) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *FrontPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrontPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrontPort) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *FrontPort) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FrontPort) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FrontPort) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FrontPort) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *FrontPort) GetType() FrontPortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FrontPort) GetTypeOk() (*FrontPortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FrontPort) SetType(v FrontPortType)`

SetType sets Type field to given value.


### GetColor

`func (o *FrontPort) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *FrontPort) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *FrontPort) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *FrontPort) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetRearPort

`func (o *FrontPort) GetRearPort() FrontPortRearPort`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *FrontPort) GetRearPortOk() (*FrontPortRearPort, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *FrontPort) SetRearPort(v FrontPortRearPort)`

SetRearPort sets RearPort field to given value.


### GetRearPortPosition

`func (o *FrontPort) GetRearPortPosition() int32`

GetRearPortPosition returns the RearPortPosition field if non-nil, zero value otherwise.

### GetRearPortPositionOk

`func (o *FrontPort) GetRearPortPositionOk() (*int32, bool)`

GetRearPortPositionOk returns a tuple with the RearPortPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortPosition

`func (o *FrontPort) SetRearPortPosition(v int32)`

SetRearPortPosition sets RearPortPosition field to given value.

### HasRearPortPosition

`func (o *FrontPort) HasRearPortPosition() bool`

HasRearPortPosition returns a boolean if a field has been set.

### GetDescription

`func (o *FrontPort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FrontPort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FrontPort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FrontPort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMarkConnected

`func (o *FrontPort) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *FrontPort) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *FrontPort) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *FrontPort) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetCable

`func (o *FrontPort) GetCable() BriefCable`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *FrontPort) GetCableOk() (*BriefCable, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *FrontPort) SetCable(v BriefCable)`

SetCable sets Cable field to given value.


### SetCableNil

`func (o *FrontPort) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *FrontPort) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetCableEnd

`func (o *FrontPort) GetCableEnd() string`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *FrontPort) GetCableEndOk() (*string, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *FrontPort) SetCableEnd(v string)`

SetCableEnd sets CableEnd field to given value.


### GetLinkPeers

`func (o *FrontPort) GetLinkPeers() []interface{}`

GetLinkPeers returns the LinkPeers field if non-nil, zero value otherwise.

### GetLinkPeersOk

`func (o *FrontPort) GetLinkPeersOk() (*[]interface{}, bool)`

GetLinkPeersOk returns a tuple with the LinkPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeers

`func (o *FrontPort) SetLinkPeers(v []interface{})`

SetLinkPeers sets LinkPeers field to given value.


### GetLinkPeersType

`func (o *FrontPort) GetLinkPeersType() string`

GetLinkPeersType returns the LinkPeersType field if non-nil, zero value otherwise.

### GetLinkPeersTypeOk

`func (o *FrontPort) GetLinkPeersTypeOk() (*string, bool)`

GetLinkPeersTypeOk returns a tuple with the LinkPeersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPeersType

`func (o *FrontPort) SetLinkPeersType(v string)`

SetLinkPeersType sets LinkPeersType field to given value.


### SetLinkPeersTypeNil

`func (o *FrontPort) SetLinkPeersTypeNil(b bool)`

 SetLinkPeersTypeNil sets the value for LinkPeersType to be an explicit nil

### UnsetLinkPeersType
`func (o *FrontPort) UnsetLinkPeersType()`

UnsetLinkPeersType ensures that no value is present for LinkPeersType, not even an explicit nil
### GetTags

`func (o *FrontPort) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FrontPort) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FrontPort) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FrontPort) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *FrontPort) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *FrontPort) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *FrontPort) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *FrontPort) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *FrontPort) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FrontPort) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FrontPort) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *FrontPort) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *FrontPort) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *FrontPort) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *FrontPort) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *FrontPort) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *FrontPort) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *FrontPort) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetOccupied

`func (o *FrontPort) GetOccupied() bool`

GetOccupied returns the Occupied field if non-nil, zero value otherwise.

### GetOccupiedOk

`func (o *FrontPort) GetOccupiedOk() (*bool, bool)`

GetOccupiedOk returns a tuple with the Occupied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupied

`func (o *FrontPort) SetOccupied(v bool)`

SetOccupied sets Occupied field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


