# DeviceWithConfigContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**DeviceType** | [**DeviceTypeRequest**](DeviceTypeRequest.md) |  | 
**Role** | [**DeviceRoleRequest**](DeviceRoleRequest.md) |  | 
**Tenant** | Pointer to [**NullableTenantRequest**](TenantRequest.md) |  | [optional] 
**Platform** | Pointer to [**NullablePlatformRequest**](PlatformRequest.md) |  | [optional] 
**Serial** | Pointer to **string** | Chassis serial number, assigned by the manufacturer | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Site** | [**SiteRequest**](SiteRequest.md) |  | 
**Location** | Pointer to [**NullableLocationRequest**](LocationRequest.md) |  | [optional] 
**Rack** | Pointer to [**NullableRackRequest**](RackRequest.md) |  | [optional] 
**Position** | Pointer to **NullableFloat64** |  | [optional] 
**Face** | Pointer to [**DeviceFaceValue**](DeviceFaceValue.md) |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Longitude** | Pointer to **NullableFloat64** | GPS coordinate in decimal format (xx.yyyyyy) | [optional] 
**Status** | Pointer to [**DeviceStatusValue**](DeviceStatusValue.md) |  | [optional] 
**Airflow** | Pointer to [**DeviceAirflowValue**](DeviceAirflowValue.md) |  | [optional] 
**PrimaryIp4** | Pointer to [**NullableIPAddressRequest**](IPAddressRequest.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullableIPAddressRequest**](IPAddressRequest.md) |  | [optional] 
**OobIp** | Pointer to [**NullableIPAddressRequest**](IPAddressRequest.md) |  | [optional] 
**Cluster** | Pointer to [**NullableClusterRequest**](ClusterRequest.md) |  | [optional] 
**VirtualChassis** | Pointer to [**NullableVirtualChassisRequest**](VirtualChassisRequest.md) |  | [optional] 
**VcPosition** | Pointer to **NullableInt32** |  | [optional] 
**VcPriority** | Pointer to **NullableInt32** | Virtual chassis master election priority | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**ConfigTemplate** | Pointer to [**NullableConfigTemplateRequest**](ConfigTemplateRequest.md) |  | [optional] 
**LocalContextData** | Pointer to **interface{}** | Local config context data takes precedence over source contexts in the final rendered config context | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeviceWithConfigContextRequest

`func NewDeviceWithConfigContextRequest(deviceType DeviceTypeRequest, role DeviceRoleRequest, site SiteRequest, ) *DeviceWithConfigContextRequest`

NewDeviceWithConfigContextRequest instantiates a new DeviceWithConfigContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithConfigContextRequestWithDefaults

`func NewDeviceWithConfigContextRequestWithDefaults() *DeviceWithConfigContextRequest`

NewDeviceWithConfigContextRequestWithDefaults instantiates a new DeviceWithConfigContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceWithConfigContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceWithConfigContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceWithConfigContextRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceWithConfigContextRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeviceWithConfigContextRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeviceWithConfigContextRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeviceType

`func (o *DeviceWithConfigContextRequest) GetDeviceType() DeviceTypeRequest`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *DeviceWithConfigContextRequest) GetDeviceTypeOk() (*DeviceTypeRequest, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *DeviceWithConfigContextRequest) SetDeviceType(v DeviceTypeRequest)`

SetDeviceType sets DeviceType field to given value.


### GetRole

`func (o *DeviceWithConfigContextRequest) GetRole() DeviceRoleRequest`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DeviceWithConfigContextRequest) GetRoleOk() (*DeviceRoleRequest, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DeviceWithConfigContextRequest) SetRole(v DeviceRoleRequest)`

SetRole sets Role field to given value.


### GetTenant

`func (o *DeviceWithConfigContextRequest) GetTenant() TenantRequest`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DeviceWithConfigContextRequest) GetTenantOk() (*TenantRequest, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DeviceWithConfigContextRequest) SetTenant(v TenantRequest)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *DeviceWithConfigContextRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *DeviceWithConfigContextRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *DeviceWithConfigContextRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *DeviceWithConfigContextRequest) GetPlatform() PlatformRequest`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceWithConfigContextRequest) GetPlatformOk() (*PlatformRequest, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceWithConfigContextRequest) SetPlatform(v PlatformRequest)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceWithConfigContextRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *DeviceWithConfigContextRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *DeviceWithConfigContextRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *DeviceWithConfigContextRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DeviceWithConfigContextRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DeviceWithConfigContextRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *DeviceWithConfigContextRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *DeviceWithConfigContextRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *DeviceWithConfigContextRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *DeviceWithConfigContextRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *DeviceWithConfigContextRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *DeviceWithConfigContextRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *DeviceWithConfigContextRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *DeviceWithConfigContextRequest) GetSite() SiteRequest`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *DeviceWithConfigContextRequest) GetSiteOk() (*SiteRequest, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *DeviceWithConfigContextRequest) SetSite(v SiteRequest)`

SetSite sets Site field to given value.


### GetLocation

`func (o *DeviceWithConfigContextRequest) GetLocation() LocationRequest`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DeviceWithConfigContextRequest) GetLocationOk() (*LocationRequest, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DeviceWithConfigContextRequest) SetLocation(v LocationRequest)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DeviceWithConfigContextRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DeviceWithConfigContextRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DeviceWithConfigContextRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRack

`func (o *DeviceWithConfigContextRequest) GetRack() RackRequest`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *DeviceWithConfigContextRequest) GetRackOk() (*RackRequest, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *DeviceWithConfigContextRequest) SetRack(v RackRequest)`

SetRack sets Rack field to given value.

### HasRack

`func (o *DeviceWithConfigContextRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *DeviceWithConfigContextRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *DeviceWithConfigContextRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *DeviceWithConfigContextRequest) GetPosition() float64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *DeviceWithConfigContextRequest) GetPositionOk() (*float64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *DeviceWithConfigContextRequest) SetPosition(v float64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *DeviceWithConfigContextRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *DeviceWithConfigContextRequest) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *DeviceWithConfigContextRequest) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *DeviceWithConfigContextRequest) GetFace() DeviceFaceValue`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *DeviceWithConfigContextRequest) GetFaceOk() (*DeviceFaceValue, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *DeviceWithConfigContextRequest) SetFace(v DeviceFaceValue)`

SetFace sets Face field to given value.

### HasFace

`func (o *DeviceWithConfigContextRequest) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetLatitude

`func (o *DeviceWithConfigContextRequest) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *DeviceWithConfigContextRequest) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *DeviceWithConfigContextRequest) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *DeviceWithConfigContextRequest) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *DeviceWithConfigContextRequest) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *DeviceWithConfigContextRequest) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *DeviceWithConfigContextRequest) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *DeviceWithConfigContextRequest) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *DeviceWithConfigContextRequest) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *DeviceWithConfigContextRequest) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *DeviceWithConfigContextRequest) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *DeviceWithConfigContextRequest) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetStatus

`func (o *DeviceWithConfigContextRequest) GetStatus() DeviceStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceWithConfigContextRequest) GetStatusOk() (*DeviceStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceWithConfigContextRequest) SetStatus(v DeviceStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceWithConfigContextRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAirflow

`func (o *DeviceWithConfigContextRequest) GetAirflow() DeviceAirflowValue`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *DeviceWithConfigContextRequest) GetAirflowOk() (*DeviceAirflowValue, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *DeviceWithConfigContextRequest) SetAirflow(v DeviceAirflowValue)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *DeviceWithConfigContextRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### GetPrimaryIp4

`func (o *DeviceWithConfigContextRequest) GetPrimaryIp4() IPAddressRequest`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *DeviceWithConfigContextRequest) GetPrimaryIp4Ok() (*IPAddressRequest, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *DeviceWithConfigContextRequest) SetPrimaryIp4(v IPAddressRequest)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *DeviceWithConfigContextRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *DeviceWithConfigContextRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *DeviceWithConfigContextRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *DeviceWithConfigContextRequest) GetPrimaryIp6() IPAddressRequest`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *DeviceWithConfigContextRequest) GetPrimaryIp6Ok() (*IPAddressRequest, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *DeviceWithConfigContextRequest) SetPrimaryIp6(v IPAddressRequest)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *DeviceWithConfigContextRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *DeviceWithConfigContextRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *DeviceWithConfigContextRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetOobIp

`func (o *DeviceWithConfigContextRequest) GetOobIp() IPAddressRequest`

GetOobIp returns the OobIp field if non-nil, zero value otherwise.

### GetOobIpOk

`func (o *DeviceWithConfigContextRequest) GetOobIpOk() (*IPAddressRequest, bool)`

GetOobIpOk returns a tuple with the OobIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOobIp

`func (o *DeviceWithConfigContextRequest) SetOobIp(v IPAddressRequest)`

SetOobIp sets OobIp field to given value.

### HasOobIp

`func (o *DeviceWithConfigContextRequest) HasOobIp() bool`

HasOobIp returns a boolean if a field has been set.

### SetOobIpNil

`func (o *DeviceWithConfigContextRequest) SetOobIpNil(b bool)`

 SetOobIpNil sets the value for OobIp to be an explicit nil

### UnsetOobIp
`func (o *DeviceWithConfigContextRequest) UnsetOobIp()`

UnsetOobIp ensures that no value is present for OobIp, not even an explicit nil
### GetCluster

`func (o *DeviceWithConfigContextRequest) GetCluster() ClusterRequest`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *DeviceWithConfigContextRequest) GetClusterOk() (*ClusterRequest, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *DeviceWithConfigContextRequest) SetCluster(v ClusterRequest)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *DeviceWithConfigContextRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *DeviceWithConfigContextRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *DeviceWithConfigContextRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *DeviceWithConfigContextRequest) GetVirtualChassis() VirtualChassisRequest`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *DeviceWithConfigContextRequest) GetVirtualChassisOk() (*VirtualChassisRequest, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *DeviceWithConfigContextRequest) SetVirtualChassis(v VirtualChassisRequest)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *DeviceWithConfigContextRequest) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *DeviceWithConfigContextRequest) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *DeviceWithConfigContextRequest) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *DeviceWithConfigContextRequest) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *DeviceWithConfigContextRequest) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *DeviceWithConfigContextRequest) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *DeviceWithConfigContextRequest) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *DeviceWithConfigContextRequest) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *DeviceWithConfigContextRequest) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *DeviceWithConfigContextRequest) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *DeviceWithConfigContextRequest) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *DeviceWithConfigContextRequest) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *DeviceWithConfigContextRequest) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *DeviceWithConfigContextRequest) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *DeviceWithConfigContextRequest) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetDescription

`func (o *DeviceWithConfigContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceWithConfigContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceWithConfigContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceWithConfigContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *DeviceWithConfigContextRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *DeviceWithConfigContextRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *DeviceWithConfigContextRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *DeviceWithConfigContextRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *DeviceWithConfigContextRequest) GetConfigTemplate() ConfigTemplateRequest`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *DeviceWithConfigContextRequest) GetConfigTemplateOk() (*ConfigTemplateRequest, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *DeviceWithConfigContextRequest) SetConfigTemplate(v ConfigTemplateRequest)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *DeviceWithConfigContextRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *DeviceWithConfigContextRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *DeviceWithConfigContextRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetLocalContextData

`func (o *DeviceWithConfigContextRequest) GetLocalContextData() interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *DeviceWithConfigContextRequest) GetLocalContextDataOk() (*interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *DeviceWithConfigContextRequest) SetLocalContextData(v interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *DeviceWithConfigContextRequest) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *DeviceWithConfigContextRequest) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *DeviceWithConfigContextRequest) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *DeviceWithConfigContextRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceWithConfigContextRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceWithConfigContextRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceWithConfigContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *DeviceWithConfigContextRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *DeviceWithConfigContextRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *DeviceWithConfigContextRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *DeviceWithConfigContextRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


