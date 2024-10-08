# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **NullableString** |  | 
**DeviceType** | Pointer to [**DeviceType**](DeviceType.md) |  | [optional] 
**Role** | Pointer to [**DeviceRole**](DeviceRole.md) |  | [optional] 
**Tenant** | Pointer to [**NullableTenant**](Tenant.md) |  | [optional] 
**Platform** | Pointer to [**NullablePlatform**](Platform.md) |  | [optional] 
**Serial** | Pointer to **string** | Chassis serial number, assigned by the manufacturer | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Site** | Pointer to [**Site**](Site.md) |  | [optional] 
**Location** | Pointer to [**NullableLocation**](Location.md) |  | [optional] 
**Rack** | Pointer to [**NullableRack**](Rack.md) |  | [optional] 
**Position** | Pointer to **NullableFloat64** |  | [optional] 
**Face** | Pointer to [**DeviceFace**](DeviceFace.md) |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Longitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**ParentDevice** | Pointer to [**NullableNestedDevice**](NestedDevice.md) |  | [optional] [readonly] 
**Status** | Pointer to [**DeviceStatus**](DeviceStatus.md) |  | [optional] 
**Airflow** | Pointer to [**DeviceAirflow**](DeviceAirflow.md) |  | [optional] 
**PrimaryIp** | Pointer to [**NullableIPAddress**](IPAddress.md) |  | [optional] [readonly] 
**PrimaryIp4** | Pointer to [**NullableIPAddress**](IPAddress.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullableIPAddress**](IPAddress.md) |  | [optional] 
**OobIp** | Pointer to [**NullableIPAddress**](IPAddress.md) |  | [optional] 
**Cluster** | Pointer to [**NullableCluster**](Cluster.md) |  | [optional] 
**VirtualChassis** | Pointer to [**NullableVirtualChassis**](VirtualChassis.md) |  | [optional] 
**VcPosition** | Pointer to **NullableInt32** |  | [optional] 
**VcPriority** | Pointer to **NullableInt32** | Virtual chassis master election priority | [optional] 
**Description** | **string** |  | 
**Comments** | Pointer to **string** |  | [optional] 
**ConfigTemplate** | Pointer to [**NullableConfigTemplate**](ConfigTemplate.md) |  | [optional] 
**LocalContextData** | Pointer to **interface{}** | Local config context data takes precedence over source contexts in the final rendered config context | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ConsolePortCount** | Pointer to **int32** |  | [optional] [readonly] 
**ConsoleServerPortCount** | Pointer to **int32** |  | [optional] [readonly] 
**PowerPortCount** | Pointer to **int32** |  | [optional] [readonly] 
**PowerOutletCount** | Pointer to **int32** |  | [optional] [readonly] 
**InterfaceCount** | Pointer to **int32** |  | [optional] [readonly] 
**FrontPortCount** | Pointer to **int32** |  | [optional] [readonly] 
**RearPortCount** | Pointer to **int32** |  | [optional] [readonly] 
**DeviceBayCount** | Pointer to **int32** |  | [optional] [readonly] 
**ModuleBayCount** | Pointer to **int32** |  | [optional] [readonly] 
**InventoryItemCount** | Pointer to **int32** |  | [optional] [readonly] 

## Methods

### NewDevice

`func NewDevice(id int32, url string, display string, name NullableString, description string, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *Device) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Device) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Device) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *Device) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Device) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Device) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Device) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Device) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeviceType

`func (o *Device) GetDeviceType() DeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *Device) GetDeviceTypeOk() (*DeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *Device) SetDeviceType(v DeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *Device) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetRole

`func (o *Device) GetRole() DeviceRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Device) GetRoleOk() (*DeviceRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Device) SetRole(v DeviceRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *Device) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTenant

`func (o *Device) GetTenant() Tenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Device) GetTenantOk() (*Tenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Device) SetTenant(v Tenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Device) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *Device) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *Device) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *Device) GetPlatform() Platform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Device) GetPlatformOk() (*Platform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Device) SetPlatform(v Platform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Device) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *Device) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *Device) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *Device) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Device) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Device) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Device) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *Device) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *Device) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *Device) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *Device) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *Device) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *Device) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *Device) GetSite() Site`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Device) GetSiteOk() (*Site, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Device) SetSite(v Site)`

SetSite sets Site field to given value.

### HasSite

`func (o *Device) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetLocation

`func (o *Device) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Device) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Device) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Device) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Device) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Device) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRack

`func (o *Device) GetRack() Rack`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *Device) GetRackOk() (*Rack, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *Device) SetRack(v Rack)`

SetRack sets Rack field to given value.

### HasRack

`func (o *Device) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *Device) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *Device) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *Device) GetPosition() float64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Device) GetPositionOk() (*float64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Device) SetPosition(v float64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Device) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *Device) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *Device) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *Device) GetFace() DeviceFace`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *Device) GetFaceOk() (*DeviceFace, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *Device) SetFace(v DeviceFace)`

SetFace sets Face field to given value.

### HasFace

`func (o *Device) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetLatitude

`func (o *Device) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Device) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Device) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Device) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *Device) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *Device) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *Device) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Device) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Device) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Device) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *Device) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *Device) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetParentDevice

`func (o *Device) GetParentDevice() NestedDevice`

GetParentDevice returns the ParentDevice field if non-nil, zero value otherwise.

### GetParentDeviceOk

`func (o *Device) GetParentDeviceOk() (*NestedDevice, bool)`

GetParentDeviceOk returns a tuple with the ParentDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDevice

`func (o *Device) SetParentDevice(v NestedDevice)`

SetParentDevice sets ParentDevice field to given value.

### HasParentDevice

`func (o *Device) HasParentDevice() bool`

HasParentDevice returns a boolean if a field has been set.

### SetParentDeviceNil

`func (o *Device) SetParentDeviceNil(b bool)`

 SetParentDeviceNil sets the value for ParentDevice to be an explicit nil

### UnsetParentDevice
`func (o *Device) UnsetParentDevice()`

UnsetParentDevice ensures that no value is present for ParentDevice, not even an explicit nil
### GetStatus

`func (o *Device) GetStatus() DeviceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Device) GetStatusOk() (*DeviceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Device) SetStatus(v DeviceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Device) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAirflow

`func (o *Device) GetAirflow() DeviceAirflow`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *Device) GetAirflowOk() (*DeviceAirflow, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *Device) SetAirflow(v DeviceAirflow)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *Device) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### GetPrimaryIp

`func (o *Device) GetPrimaryIp() IPAddress`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *Device) GetPrimaryIpOk() (*IPAddress, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *Device) SetPrimaryIp(v IPAddress)`

SetPrimaryIp sets PrimaryIp field to given value.

### HasPrimaryIp

`func (o *Device) HasPrimaryIp() bool`

HasPrimaryIp returns a boolean if a field has been set.

### SetPrimaryIpNil

`func (o *Device) SetPrimaryIpNil(b bool)`

 SetPrimaryIpNil sets the value for PrimaryIp to be an explicit nil

### UnsetPrimaryIp
`func (o *Device) UnsetPrimaryIp()`

UnsetPrimaryIp ensures that no value is present for PrimaryIp, not even an explicit nil
### GetPrimaryIp4

`func (o *Device) GetPrimaryIp4() IPAddress`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *Device) GetPrimaryIp4Ok() (*IPAddress, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *Device) SetPrimaryIp4(v IPAddress)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *Device) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *Device) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *Device) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *Device) GetPrimaryIp6() IPAddress`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *Device) GetPrimaryIp6Ok() (*IPAddress, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *Device) SetPrimaryIp6(v IPAddress)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *Device) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *Device) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *Device) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetOobIp

`func (o *Device) GetOobIp() IPAddress`

GetOobIp returns the OobIp field if non-nil, zero value otherwise.

### GetOobIpOk

`func (o *Device) GetOobIpOk() (*IPAddress, bool)`

GetOobIpOk returns a tuple with the OobIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIp

`func (o *Device) SetOobIp(v IPAddress)`

SetOobIp sets OobIp field to given value.

### HasOobIp

`func (o *Device) HasOobIp() bool`

HasOobIp returns a boolean if a field has been set.

### SetOobIpNil

`func (o *Device) SetOobIpNil(b bool)`

 SetOobIpNil sets the value for OobIp to be an explicit nil

### UnsetOobIp
`func (o *Device) UnsetOobIp()`

UnsetOobIp ensures that no value is present for OobIp, not even an explicit nil
### GetCluster

`func (o *Device) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Device) GetClusterOk() (*Cluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Device) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Device) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *Device) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *Device) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *Device) GetVirtualChassis() VirtualChassis`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *Device) GetVirtualChassisOk() (*VirtualChassis, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *Device) SetVirtualChassis(v VirtualChassis)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *Device) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *Device) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *Device) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *Device) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *Device) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *Device) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *Device) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *Device) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *Device) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *Device) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *Device) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *Device) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *Device) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *Device) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *Device) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetDescription

`func (o *Device) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Device) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Device) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetComments

`func (o *Device) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *Device) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *Device) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *Device) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *Device) GetConfigTemplate() ConfigTemplate`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *Device) GetConfigTemplateOk() (*ConfigTemplate, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *Device) SetConfigTemplate(v ConfigTemplate)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *Device) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *Device) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *Device) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetLocalContextData

`func (o *Device) GetLocalContextData() interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *Device) GetLocalContextDataOk() (*interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *Device) SetLocalContextData(v interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *Device) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *Device) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *Device) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *Device) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Device) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Device) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Device) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Device) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Device) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Device) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Device) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreated

`func (o *Device) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Device) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Device) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Device) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *Device) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Device) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Device) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Device) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Device) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Device) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *Device) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Device) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetConsolePortCount

`func (o *Device) GetConsolePortCount() int32`

GetConsolePortCount returns the ConsolePortCount field if non-nil, zero value otherwise.

### GetConsolePortCountOk

`func (o *Device) GetConsolePortCountOk() (*int32, bool)`

GetConsolePortCountOk returns a tuple with the ConsolePortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolePortCount

`func (o *Device) SetConsolePortCount(v int32)`

SetConsolePortCount sets ConsolePortCount field to given value.

### HasConsolePortCount

`func (o *Device) HasConsolePortCount() bool`

HasConsolePortCount returns a boolean if a field has been set.

### GetConsoleServerPortCount

`func (o *Device) GetConsoleServerPortCount() int32`

GetConsoleServerPortCount returns the ConsoleServerPortCount field if non-nil, zero value otherwise.

### GetConsoleServerPortCountOk

`func (o *Device) GetConsoleServerPortCountOk() (*int32, bool)`

GetConsoleServerPortCountOk returns a tuple with the ConsoleServerPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleServerPortCount

`func (o *Device) SetConsoleServerPortCount(v int32)`

SetConsoleServerPortCount sets ConsoleServerPortCount field to given value.

### HasConsoleServerPortCount

`func (o *Device) HasConsoleServerPortCount() bool`

HasConsoleServerPortCount returns a boolean if a field has been set.

### GetPowerPortCount

`func (o *Device) GetPowerPortCount() int32`

GetPowerPortCount returns the PowerPortCount field if non-nil, zero value otherwise.

### GetPowerPortCountOk

`func (o *Device) GetPowerPortCountOk() (*int32, bool)`

GetPowerPortCountOk returns a tuple with the PowerPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPortCount

`func (o *Device) SetPowerPortCount(v int32)`

SetPowerPortCount sets PowerPortCount field to given value.

### HasPowerPortCount

`func (o *Device) HasPowerPortCount() bool`

HasPowerPortCount returns a boolean if a field has been set.

### GetPowerOutletCount

`func (o *Device) GetPowerOutletCount() int32`

GetPowerOutletCount returns the PowerOutletCount field if non-nil, zero value otherwise.

### GetPowerOutletCountOk

`func (o *Device) GetPowerOutletCountOk() (*int32, bool)`

GetPowerOutletCountOk returns a tuple with the PowerOutletCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOutletCount

`func (o *Device) SetPowerOutletCount(v int32)`

SetPowerOutletCount sets PowerOutletCount field to given value.

### HasPowerOutletCount

`func (o *Device) HasPowerOutletCount() bool`

HasPowerOutletCount returns a boolean if a field has been set.

### GetInterfaceCount

`func (o *Device) GetInterfaceCount() int32`

GetInterfaceCount returns the InterfaceCount field if non-nil, zero value otherwise.

### GetInterfaceCountOk

`func (o *Device) GetInterfaceCountOk() (*int32, bool)`

GetInterfaceCountOk returns a tuple with the InterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceCount

`func (o *Device) SetInterfaceCount(v int32)`

SetInterfaceCount sets InterfaceCount field to given value.

### HasInterfaceCount

`func (o *Device) HasInterfaceCount() bool`

HasInterfaceCount returns a boolean if a field has been set.

### GetFrontPortCount

`func (o *Device) GetFrontPortCount() int32`

GetFrontPortCount returns the FrontPortCount field if non-nil, zero value otherwise.

### GetFrontPortCountOk

`func (o *Device) GetFrontPortCountOk() (*int32, bool)`

GetFrontPortCountOk returns a tuple with the FrontPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPortCount

`func (o *Device) SetFrontPortCount(v int32)`

SetFrontPortCount sets FrontPortCount field to given value.

### HasFrontPortCount

`func (o *Device) HasFrontPortCount() bool`

HasFrontPortCount returns a boolean if a field has been set.

### GetRearPortCount

`func (o *Device) GetRearPortCount() int32`

GetRearPortCount returns the RearPortCount field if non-nil, zero value otherwise.

### GetRearPortCountOk

`func (o *Device) GetRearPortCountOk() (*int32, bool)`

GetRearPortCountOk returns a tuple with the RearPortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortCount

`func (o *Device) SetRearPortCount(v int32)`

SetRearPortCount sets RearPortCount field to given value.

### HasRearPortCount

`func (o *Device) HasRearPortCount() bool`

HasRearPortCount returns a boolean if a field has been set.

### GetDeviceBayCount

`func (o *Device) GetDeviceBayCount() int32`

GetDeviceBayCount returns the DeviceBayCount field if non-nil, zero value otherwise.

### GetDeviceBayCountOk

`func (o *Device) GetDeviceBayCountOk() (*int32, bool)`

GetDeviceBayCountOk returns a tuple with the DeviceBayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceBayCount

`func (o *Device) SetDeviceBayCount(v int32)`

SetDeviceBayCount sets DeviceBayCount field to given value.

### HasDeviceBayCount

`func (o *Device) HasDeviceBayCount() bool`

HasDeviceBayCount returns a boolean if a field has been set.

### GetModuleBayCount

`func (o *Device) GetModuleBayCount() int32`

GetModuleBayCount returns the ModuleBayCount field if non-nil, zero value otherwise.

### GetModuleBayCountOk

`func (o *Device) GetModuleBayCountOk() (*int32, bool)`

GetModuleBayCountOk returns a tuple with the ModuleBayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleBayCount

`func (o *Device) SetModuleBayCount(v int32)`

SetModuleBayCount sets ModuleBayCount field to given value.

### HasModuleBayCount

`func (o *Device) HasModuleBayCount() bool`

HasModuleBayCount returns a boolean if a field has been set.

### GetInventoryItemCount

`func (o *Device) GetInventoryItemCount() int32`

GetInventoryItemCount returns the InventoryItemCount field if non-nil, zero value otherwise.

### GetInventoryItemCountOk

`func (o *Device) GetInventoryItemCountOk() (*int32, bool)`

GetInventoryItemCountOk returns a tuple with the InventoryItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemCount

`func (o *Device) SetInventoryItemCount(v int32)`

SetInventoryItemCount sets InventoryItemCount field to given value.

### HasInventoryItemCount

`func (o *Device) HasInventoryItemCount() bool`

HasInventoryItemCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


