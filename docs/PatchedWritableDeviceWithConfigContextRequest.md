# PatchedWritableDeviceWithConfigContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**DeviceType** | Pointer to **int32** |  | [optional] 
**DeviceRole** | Pointer to **int32** | The function this device serves | [optional] 
**Tenant** | Pointer to **NullableInt32** |  | [optional] 
**Platform** | Pointer to **NullableInt32** |  | [optional] 
**Serial** | Pointer to **string** | Chassis serial number, assigned by the manufacturer | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this device | [optional] 
**Site** | Pointer to **int32** |  | [optional] 
**Location** | Pointer to **NullableInt32** |  | [optional] 
**Rack** | Pointer to **NullableInt32** |  | [optional] 
**Position** | Pointer to **NullableFloat64** |  | [optional] 
**Face** | Pointer to **string** | * &#x60;front&#x60; - Front * &#x60;rear&#x60; - Rear | [optional] 
**Status** | Pointer to **string** | * &#x60;offline&#x60; - Offline * &#x60;active&#x60; - Active * &#x60;planned&#x60; - Planned * &#x60;staged&#x60; - Staged * &#x60;failed&#x60; - Failed * &#x60;inventory&#x60; - Inventory * &#x60;decommissioning&#x60; - Decommissioning | [optional] 
**Airflow** | Pointer to **string** | * &#x60;front-to-rear&#x60; - Front to rear * &#x60;rear-to-front&#x60; - Rear to front * &#x60;left-to-right&#x60; - Left to right * &#x60;right-to-left&#x60; - Right to left * &#x60;side-to-rear&#x60; - Side to rear * &#x60;passive&#x60; - Passive * &#x60;mixed&#x60; - Mixed | [optional] 
**PrimaryIp4** | Pointer to **NullableInt32** |  | [optional] 
**PrimaryIp6** | Pointer to **NullableInt32** |  | [optional] 
**Cluster** | Pointer to **NullableInt32** |  | [optional] 
**VirtualChassis** | Pointer to **NullableInt32** |  | [optional] 
**VcPosition** | Pointer to **NullableInt32** |  | [optional] 
**VcPriority** | Pointer to **NullableInt32** | Virtual chassis master election priority | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**LocalContextData** | Pointer to **map[string]interface{}** | Local config context data takes precedence over source contexts in the final rendered config context | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ConfigTemplate** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPatchedWritableDeviceWithConfigContextRequest

`func NewPatchedWritableDeviceWithConfigContextRequest() *PatchedWritableDeviceWithConfigContextRequest`

NewPatchedWritableDeviceWithConfigContextRequest instantiates a new PatchedWritableDeviceWithConfigContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableDeviceWithConfigContextRequestWithDefaults

`func NewPatchedWritableDeviceWithConfigContextRequestWithDefaults() *PatchedWritableDeviceWithConfigContextRequest`

NewPatchedWritableDeviceWithConfigContextRequestWithDefaults instantiates a new PatchedWritableDeviceWithConfigContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDeviceType

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceRole

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetDeviceRole() int32`

GetDeviceRole returns the DeviceRole field if non-nil, zero value otherwise.

### GetDeviceRoleOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetDeviceRoleOk() (*int32, bool)`

GetDeviceRoleOk returns a tuple with the DeviceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRole

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetDeviceRole(v int32)`

SetDeviceRole sets DeviceRole field to given value.

### HasDeviceRole

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasDeviceRole() bool`

HasDeviceRole returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetTenant() int32`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetTenantOk() (*int32, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetTenant(v int32)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPlatform() int32`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPlatformOk() (*int32, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPlatform(v int32)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSerial

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetAssetTag

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetSite

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetSite() int32`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetSiteOk() (*int32, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetSite(v int32)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLocation() int32`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLocationOk() (*int32, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLocation(v int32)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetRack

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetRack() int32`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetRackOk() (*int32, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetRack(v int32)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPosition() float64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPositionOk() (*float64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPosition(v float64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetFace

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetFace() string`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetFaceOk() (*string, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetFace(v string)`

SetFace sets Face field to given value.

### HasFace

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAirflow

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetAirflow() string`

GetAirflow returns the Airflow field if non-nil, zero value otherwise.

### GetAirflowOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetAirflowOk() (*string, bool)`

GetAirflowOk returns a tuple with the Airflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirflow

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetAirflow(v string)`

SetAirflow sets Airflow field to given value.

### HasAirflow

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasAirflow() bool`

HasAirflow returns a boolean if a field has been set.

### GetPrimaryIp4

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPrimaryIp4() int32`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPrimaryIp4Ok() (*int32, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPrimaryIp4(v int32)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPrimaryIp6() int32`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetPrimaryIp6Ok() (*int32, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPrimaryIp6(v int32)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetCluster

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetCluster() int32`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetClusterOk() (*int32, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetCluster(v int32)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetVirtualChassis

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVirtualChassis() int32`

GetVirtualChassis returns the VirtualChassis field if non-nil, zero value otherwise.

### GetVirtualChassisOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVirtualChassisOk() (*int32, bool)`

GetVirtualChassisOk returns a tuple with the VirtualChassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualChassis

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVirtualChassis(v int32)`

SetVirtualChassis sets VirtualChassis field to given value.

### HasVirtualChassis

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasVirtualChassis() bool`

HasVirtualChassis returns a boolean if a field has been set.

### SetVirtualChassisNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVirtualChassisNil(b bool)`

 SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil

### UnsetVirtualChassis
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetVirtualChassis()`

UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
### GetVcPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVcPosition() int32`

GetVcPosition returns the VcPosition field if non-nil, zero value otherwise.

### GetVcPositionOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVcPositionOk() (*int32, bool)`

GetVcPositionOk returns a tuple with the VcPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVcPosition(v int32)`

SetVcPosition sets VcPosition field to given value.

### HasVcPosition

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasVcPosition() bool`

HasVcPosition returns a boolean if a field has been set.

### SetVcPositionNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVcPositionNil(b bool)`

 SetVcPositionNil sets the value for VcPosition to be an explicit nil

### UnsetVcPosition
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetVcPosition()`

UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
### GetVcPriority

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVcPriority() int32`

GetVcPriority returns the VcPriority field if non-nil, zero value otherwise.

### GetVcPriorityOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetVcPriorityOk() (*int32, bool)`

GetVcPriorityOk returns a tuple with the VcPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcPriority

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVcPriority(v int32)`

SetVcPriority sets VcPriority field to given value.

### HasVcPriority

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasVcPriority() bool`

HasVcPriority returns a boolean if a field has been set.

### SetVcPriorityNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetVcPriorityNil(b bool)`

 SetVcPriorityNil sets the value for VcPriority to be an explicit nil

### UnsetVcPriority
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetVcPriority()`

UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
### GetDescription

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalContextData

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLocalContextData() map[string]interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetLocalContextDataOk() (*map[string]interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLocalContextData(v map[string]interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetConfigTemplate() int32`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *PatchedWritableDeviceWithConfigContextRequest) GetConfigTemplateOk() (*int32, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetConfigTemplate(v int32)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *PatchedWritableDeviceWithConfigContextRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *PatchedWritableDeviceWithConfigContextRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *PatchedWritableDeviceWithConfigContextRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


