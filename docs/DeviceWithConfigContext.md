# DeviceWithConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | Pointer to **NullableString** |  | [optional] 
**DeviceType** | [**BriefDeviceType**](BriefDeviceType.md) |  | 
**Role** | [**BriefDeviceRole**](BriefDeviceRole.md) |  | 
**Tenant** | Pointer to [**NullableBriefTenant**](BriefTenant.md) |  | [optional] 
**Platform** | Pointer to [**NullableBriefPlatform**](BriefPlatform.md) |  | [optional] 
**Serial** | Pointer to **string** | Chassis serial number, assigned by the manufacturer | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Site** | [**BriefSite**](BriefSite.md) |  | 
**Location** | Pointer to [**NullableBriefLocation**](BriefLocation.md) |  | [optional] 
**Rack** | Pointer to [**NullableBriefRack**](BriefRack.md) |  | [optional] 
**Position** | Pointer to **NullableFloat64** |  | [optional] 
**Face** | Pointer to [**DeviceFace**](DeviceFace.md) |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Longitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**ParentDevice** | [**NullableNestedDevice**](NestedDevice.md) |  | [readonly] 
**Status** | Pointer to [**DeviceStatus**](DeviceStatus.md) |  | [optional] 
**Airflow** | Pointer to [**DeviceAirflow**](DeviceAirflow.md) |  | [optional] 
**PrimaryIp** | [**NullableBriefIPAddress**](BriefIPAddress.md) |  | [readonly] 
**PrimaryIp4** | Pointer to [**NullableBriefIPAddress**](BriefIPAddress.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullableBriefIPAddress**](BriefIPAddress.md) |  | [optional] 
**OobIp** | Pointer to [**NullableBriefIPAddress**](BriefIPAddress.md) |  | [optional] 
**Cluster** | Pointer to [**NullableBriefCluster**](BriefCluster.md) |  | [optional] 
**VirtualChassis** | Pointer to [**NullableBriefVirtualChassis**](BriefVirtualChassis.md) |  | [optional] 
**VcPosition** | Pointer to **NullableInt32** |  | [optional] 
**VcPriority** | Pointer to **NullableInt32** | Virtual chassis master election priority | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**ConfigTemplate** | Pointer to [**NullableBriefConfigTemplate**](BriefConfigTemplate.md) |  | [optional] 
**ConfigContext** | **interface{}** |  | [readonly] 
**LocalContextData** | Pointer to **interface{}** | Local config context data takes precedence over source contexts in the final rendered config context | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**ConsolePortCount** | **int32** |  | [readonly] 
**ConsoleServerPortCount** | **int32** |  | [readonly] 
**PowerPortCount** | **int32** |  | [readonly] 
**PowerOutletCount** | **int32** |  | [readonly] 
**InterfaceCount** | **int32** |  | [readonly] 
**FrontPortCount** | **int32** |  | [readonly] 
**RearPortCount** | **int32** |  | [readonly] 
**DeviceBayCount** | **int32** |  | [readonly] 
**ModuleBayCount** | **int32** |  | [readonly] 
**InventoryItemCount** | **int32** |  | [readonly] 

## Methods

### NewDeviceWithConfigContext

`func NewDeviceWithConfigContext(id int32, url string, displayUrl string, display string, deviceType BriefDeviceType, role BriefDeviceRole, site BriefSite, parentDevice NullableNestedDevice, primaryIp NullableBriefIPAddress, configContext interface{}, created NullableTime, lastUpdated NullableTime, consolePortCount int32, consoleServerPortCount int32, powerPortCount int32, powerOutletCount int32, interfaceCount int32, frontPortCount int32, rearPortCount int32, deviceBayCount int32, moduleBayCount int32, inventoryItemCount int32, ) *DeviceWithConfigContext`

NewDeviceWithConfigContext instantiates a new DeviceWithConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithConfigContextWithDefaults

`func NewDeviceWithConfigContextWithDefaults() *DeviceWithConfigContext`

NewDeviceWithConfigContextWithDefaults instantiates a new DeviceWithConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceWithConfigContext) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceWithConfigContext) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceWithConfigContext) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *DeviceWithConfigContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeviceWithConfigContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeviceWithConfigContext) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *DeviceWithConfigContext) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *DeviceWithConfigContext) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *DeviceWithConfigContext) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *DeviceWithConfigContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DeviceWithConfigContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DeviceWithConfigContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *DeviceWithConfigContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceWithConfigContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceWithConfigContext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceWithConfigContext) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeviceWithConfigContext) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeviceWithConfigContext) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeviceType

`func (o *DeviceWithConfigContext) GetDeviceType() BriefDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceWithConfigContext) GetDeviceTypeOk() (*BriefDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceWithConfigContext) SetDeviceType(v BriefDeviceType)`

SetDeviceType sets DeviceType field to given value.


### GetRole

`func (o *DeviceWithConfigContext) GetRole() BriefDeviceRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DeviceWithConfigContext) GetRoleOk() (*BriefDeviceRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DeviceWithConfigContext) SetRole(v BriefDeviceRole)`

SetRole sets Role field to given value.


### GetTenant

`func (o *DeviceWithConfigContext) GetTenant() BriefTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DeviceWithConfigContext) GetTenantOk() (*BriefTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DeviceWithConfigContext) SetTenant(v BriefTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *DeviceWithConfigContext) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *DeviceWithConfigContext) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *DeviceWithConfigContext) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *DeviceWithConfigContext) GetPlatform() BriefPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceWithConfigContext) GetPlatformOk() (*BriefPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceWithConfigContext) SetPlatform(v BriefPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceWithConfigContext) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *DeviceWithConfigContext) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *DeviceWithConfigContext) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *DeviceWithConfigContext) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DeviceWithConfigContext) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DeviceWithConfigContext) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *DeviceWithConfigContext) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *DeviceWithConfigContext) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *DeviceWithConfigContext) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *DeviceWithConfigContext) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *DeviceWithConfigContext) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *DeviceWithConfigContext) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *DeviceWithConfigContext) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *DeviceWithConfigContext) GetSite() BriefSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *DeviceWithConfigContext) GetSiteOk() (*BriefSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *DeviceWithConfigContext) SetSite(v BriefSite)`

SetSite sets Site field to given value.


### GetLocation

`func (o *DeviceWithConfigContext) GetLocation() BriefLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DeviceWithConfigContext) GetLocationOk() (*BriefLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DeviceWithConfigContext) SetLocation(v BriefLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DeviceWithConfigContext) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DeviceWithConfigContext) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DeviceWithConfigContext) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRack

`func (o *DeviceWithConfigContext) GetRack() BriefRack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *DeviceWithConfigContext) GetRackOk() (*BriefRack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *DeviceWithConfigContext) SetRack(v BriefRack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *DeviceWithConfigContext) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *DeviceWithConfigContext) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *DeviceWithConfigContext) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *DeviceWithConfigContext) GetPosition() float64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *DeviceWithConfigContext) GetPositionOk() (*float64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *DeviceWithConfigContext) SetPosition(v float64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *DeviceWithConfigContext) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *DeviceWithConfigContext) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *DeviceWithConfigContext) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *DeviceWithConfigContext) GetFace() DeviceFace`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *DeviceWithConfigContext) GetFaceOk() (*DeviceFace, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *DeviceWithConfigContext) SetFace(v DeviceFace)`

SetFace sets Face field to given value.

### HasFace

`func (o *DeviceWithConfigContext) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetLatitude

`func (o *DeviceWithConfigContext) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *DeviceWithConfigContext) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *DeviceWithConfigContext) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *DeviceWithConfigContext) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *DeviceWithConfigContext) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *DeviceWithConfigContext) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *DeviceWithConfigContext) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *DeviceWithConfigContext) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *DeviceWithConfigContext) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *DeviceWithConfigContext) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *DeviceWithConfigContext) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *DeviceWithConfigContext) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetParentDevice

`func (o *DeviceWithConfigContext) GetParentDevice() NestedDevice`

GetParentDevice returns the ParentDevice field if non-nil, zero value otherwise.

### GetParentDeviceOk

`func (o *DeviceWithConfigContext) GetParentDeviceOk() (*NestedDevice, bool)`

GetParentDeviceOk returns a tuple with the ParentDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDevice

`func (o *DeviceWithConfigContext) SetParentDevice(v NestedDevice)`

SetParentDevice sets ParentDevice field to given value.


### SetParentDeviceNil

`func (o *DeviceWithConfigContext) SetParentDeviceNil(b bool)`

 SetParentDeviceNil sets the value for ParentDevice to be an explicit nil

### UnsetParentDevice
`func (o *DeviceWithConfigContext) UnsetParentDevice()`

UnsetParentDevice ensures that no value is present for ParentDevice, not even an explicit nil
### GetStatus

`func (o *DeviceWithConfigContext) GetStatus() DeviceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceWithConfigContext) GetStatusOk() (*DeviceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceWithConfigContext) SetStatus(v DeviceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceWithConfigContext) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAirflow

`func (o *DeviceWithConfigContext) GetAirflow() DeviceAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *DeviceWithConfigContext) GetAirflowOk() (*DeviceAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *DeviceWithConfigContext) SetAirflow(v DeviceAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *DeviceWithConfigContext) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### GetPrimaryIp

`func (o *DeviceWithConfigContext) GetPrimaryIp() BriefIPAddress`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *DeviceWithConfigContext) GetPrimaryIpOk() (*BriefIPAddress, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *DeviceWithConfigContext) SetPrimaryIp(v BriefIPAddress)`

SetPrimaryIp sets PrimaryIp field to given value.


### SetPrimaryIpNil

`func (o *DeviceWithConfigContext) SetPrimaryIpNil(b bool)`

 SetPrimaryIpNil sets the value for PrimaryIp to be an explicit nil

### UnsetPrimaryIp
`func (o *DeviceWithConfigContext) UnsetPrimaryIp()`

UnsetPrimaryIp ensures that no value is present for PrimaryIp, not even an explicit nil
### GetPrimaryIp4

`func (o *DeviceWithConfigContext) GetPrimaryIp4() BriefIPAddress`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *DeviceWithConfigContext) GetPrimaryIp4Ok() (*BriefIPAddress, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *DeviceWithConfigContext) SetPrimaryIp4(v BriefIPAddress)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *DeviceWithConfigContext) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *DeviceWithConfigContext) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *DeviceWithConfigContext) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *DeviceWithConfigContext) GetPrimaryIp6() BriefIPAddress`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *DeviceWithConfigContext) GetPrimaryIp6Ok() (*BriefIPAddress, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *DeviceWithConfigContext) SetPrimaryIp6(v BriefIPAddress)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *DeviceWithConfigContext) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *DeviceWithConfigContext) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *DeviceWithConfigContext) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetOobIp

`func (o *DeviceWithConfigContext) GetOobIp() BriefIPAddress`

GetOobIp returns the OobIp field if non-nil, zero value otherwise.

### GetOobIpOk

`func (o *DeviceWithConfigContext) GetOobIpOk() (*BriefIPAddress, bool)`

GetOobIpOk returns a tuple with the OobIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIp

`func (o *DeviceWithConfigContext) SetOobIp(v BriefIPAddress)`

SetOobIp sets OobIp field to given value.

### HasOobIp

`func (o *DeviceWithConfigContext) HasOobIp() bool`

HasOobIp returns a boolean if a field has been set.

### SetOobIpNil

`func (o *DeviceWithConfigContext) SetOobIpNil(b bool)`

 SetOobIpNil sets the value for OobIp to be an explicit nil

### UnsetOobIp
`func (o *DeviceWithConfigContext) UnsetOobIp()`

UnsetOobIp ensures that no value is present for OobIp, not even an explicit nil
### GetCluster

`func (o *DeviceWithConfigContext) GetCluster() BriefCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DeviceWithConfigContext) GetClusterOk() (*BriefCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DeviceWithConfigContext) SetCluster(v BriefCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DeviceWithConfigContext) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *DeviceWithConfigContext) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *DeviceWithConfigContext) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *DeviceWithConfigContext) GetVirtualChassis() BriefVirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *DeviceWithConfigContext) GetVirtualChassisOk() (*BriefVirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *DeviceWithConfigContext) SetVirtualChassis(v BriefVirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *DeviceWithConfigContext) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *DeviceWithConfigContext) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *DeviceWithConfigContext) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *DeviceWithConfigContext) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *DeviceWithConfigContext) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *DeviceWithConfigContext) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *DeviceWithConfigContext) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *DeviceWithConfigContext) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *DeviceWithConfigContext) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *DeviceWithConfigContext) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *DeviceWithConfigContext) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *DeviceWithConfigContext) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *DeviceWithConfigContext) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *DeviceWithConfigContext) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *DeviceWithConfigContext) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetDescription

`func (o *DeviceWithConfigContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceWithConfigContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceWithConfigContext) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceWithConfigContext) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *DeviceWithConfigContext) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DeviceWithConfigContext) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DeviceWithConfigContext) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DeviceWithConfigContext) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *DeviceWithConfigContext) GetConfigTemplate() BriefConfigTemplate`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *DeviceWithConfigContext) GetConfigTemplateOk() (*BriefConfigTemplate, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *DeviceWithConfigContext) SetConfigTemplate(v BriefConfigTemplate)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *DeviceWithConfigContext) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *DeviceWithConfigContext) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *DeviceWithConfigContext) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetConfigContext

`func (o *DeviceWithConfigContext) GetConfigContext() interface{}`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *DeviceWithConfigContext) GetConfigContextOk() (*interface{}, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *DeviceWithConfigContext) SetConfigContext(v interface{})`

SetConfigContext sets ConfigContext field to given value.


### SetConfigContextNil

`func (o *DeviceWithConfigContext) SetConfigContextNil(b bool)`

 SetConfigContextNil sets the value for ConfigContext to be an explicit nil

### UnsetConfigContext
`func (o *DeviceWithConfigContext) UnsetConfigContext()`

UnsetConfigContext ensures that no value is present for ConfigContext, not even an explicit nil
### GetLocalContextData

`func (o *DeviceWithConfigContext) GetLocalContextData() interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *DeviceWithConfigContext) GetLocalContextDataOk() (*interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *DeviceWithConfigContext) SetLocalContextData(v interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *DeviceWithConfigContext) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *DeviceWithConfigContext) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *DeviceWithConfigContext) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *DeviceWithConfigContext) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceWithConfigContext) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceWithConfigContext) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceWithConfigContext) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DeviceWithConfigContext) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DeviceWithConfigContext) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DeviceWithConfigContext) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DeviceWithConfigContext) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *DeviceWithConfigContext) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceWithConfigContext) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceWithConfigContext) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *DeviceWithConfigContext) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *DeviceWithConfigContext) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *DeviceWithConfigContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeviceWithConfigContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeviceWithConfigContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *DeviceWithConfigContext) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *DeviceWithConfigContext) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetConsolePortCount

`func (o *DeviceWithConfigContext) GetConsolePortCount() int32`

GetConsolePortCount returns the ConsolePortCount field if non-nil, zero value otherwise.

### GetConsolePortCountOk

`func (o *DeviceWithConfigContext) GetConsolePortCountOk() (*int32, bool)`

GetConsolePortCountOk returns a tuple with the ConsolePortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolePortCount

`func (o *DeviceWithConfigContext) SetConsolePortCount(v int32)`

SetConsolePortCount sets ConsolePortCount field to given value.


### GetConsoleServerPortCount

`func (o *DeviceWithConfigContext) GetConsoleServerPortCount() int32`

GetConsoleServerPortCount returns the ConsoleServerPortCount field if non-nil, zero value otherwise.

### GetConsoleServerPortCountOk

`func (o *DeviceWithConfigContext) GetConsoleServerPortCountOk() (*int32, bool)`

GetConsoleServerPortCountOk returns a tuple with the ConsoleServerPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleServerPortCount

`func (o *DeviceWithConfigContext) SetConsoleServerPortCount(v int32)`

SetConsoleServerPortCount sets ConsoleServerPortCount field to given value.


### GetPowerPortCount

`func (o *DeviceWithConfigContext) GetPowerPortCount() int32`

GetPowerPortCount returns the PowerPortCount field if non-nil, zero value otherwise.

### GetPowerPortCountOk

`func (o *DeviceWithConfigContext) GetPowerPortCountOk() (*int32, bool)`

GetPowerPortCountOk returns a tuple with the PowerPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPortCount

`func (o *DeviceWithConfigContext) SetPowerPortCount(v int32)`

SetPowerPortCount sets PowerPortCount field to given value.


### GetPowerOutletCount

`func (o *DeviceWithConfigContext) GetPowerOutletCount() int32`

GetPowerOutletCount returns the PowerOutletCount field if non-nil, zero value otherwise.

### GetPowerOutletCountOk

`func (o *DeviceWithConfigContext) GetPowerOutletCountOk() (*int32, bool)`

GetPowerOutletCountOk returns a tuple with the PowerOutletCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOutletCount

`func (o *DeviceWithConfigContext) SetPowerOutletCount(v int32)`

SetPowerOutletCount sets PowerOutletCount field to given value.


### GetInterfaceCount

`func (o *DeviceWithConfigContext) GetInterfaceCount() int32`

GetInterfaceCount returns the InterfaceCount field if non-nil, zero value otherwise.

### GetInterfaceCountOk

`func (o *DeviceWithConfigContext) GetInterfaceCountOk() (*int32, bool)`

GetInterfaceCountOk returns a tuple with the InterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceCount

`func (o *DeviceWithConfigContext) SetInterfaceCount(v int32)`

SetInterfaceCount sets InterfaceCount field to given value.


### GetFrontPortCount

`func (o *DeviceWithConfigContext) GetFrontPortCount() int32`

GetFrontPortCount returns the FrontPortCount field if non-nil, zero value otherwise.

### GetFrontPortCountOk

`func (o *DeviceWithConfigContext) GetFrontPortCountOk() (*int32, bool)`

GetFrontPortCountOk returns a tuple with the FrontPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPortCount

`func (o *DeviceWithConfigContext) SetFrontPortCount(v int32)`

SetFrontPortCount sets FrontPortCount field to given value.


### GetRearPortCount

`func (o *DeviceWithConfigContext) GetRearPortCount() int32`

GetRearPortCount returns the RearPortCount field if non-nil, zero value otherwise.

### GetRearPortCountOk

`func (o *DeviceWithConfigContext) GetRearPortCountOk() (*int32, bool)`

GetRearPortCountOk returns a tuple with the RearPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortCount

`func (o *DeviceWithConfigContext) SetRearPortCount(v int32)`

SetRearPortCount sets RearPortCount field to given value.


### GetDeviceBayCount

`func (o *DeviceWithConfigContext) GetDeviceBayCount() int32`

GetDeviceBayCount returns the DeviceBayCount field if non-nil, zero value otherwise.

### GetDeviceBayCountOk

`func (o *DeviceWithConfigContext) GetDeviceBayCountOk() (*int32, bool)`

GetDeviceBayCountOk returns a tuple with the DeviceBayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceBayCount

`func (o *DeviceWithConfigContext) SetDeviceBayCount(v int32)`

SetDeviceBayCount sets DeviceBayCount field to given value.


### GetModuleBayCount

`func (o *DeviceWithConfigContext) GetModuleBayCount() int32`

GetModuleBayCount returns the ModuleBayCount field if non-nil, zero value otherwise.

### GetModuleBayCountOk

`func (o *DeviceWithConfigContext) GetModuleBayCountOk() (*int32, bool)`

GetModuleBayCountOk returns a tuple with the ModuleBayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayCount

`func (o *DeviceWithConfigContext) SetModuleBayCount(v int32)`

SetModuleBayCount sets ModuleBayCount field to given value.


### GetInventoryItemCount

`func (o *DeviceWithConfigContext) GetInventoryItemCount() int32`

GetInventoryItemCount returns the InventoryItemCount field if non-nil, zero value otherwise.

### GetInventoryItemCountOk

`func (o *DeviceWithConfigContext) GetInventoryItemCountOk() (*int32, bool)`

GetInventoryItemCountOk returns a tuple with the InventoryItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemCount

`func (o *DeviceWithConfigContext) SetInventoryItemCount(v int32)`

SetInventoryItemCount sets InventoryItemCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


