# VirtualMachineWithConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Status** | Pointer to [**ModuleStatus**](ModuleStatus.md) |  | [optional] 
**Site** | Pointer to [**NullableNestedSite**](NestedSite.md) |  | [optional] 
**Cluster** | Pointer to [**NullableNestedCluster**](NestedCluster.md) |  | [optional] 
**Device** | Pointer to [**NullableNestedDevice**](NestedDevice.md) |  | [optional] 
**Role** | Pointer to [**NullableNestedDeviceRole**](NestedDeviceRole.md) |  | [optional] 
**Tenant** | Pointer to [**NullableNestedTenant**](NestedTenant.md) |  | [optional] 
**Platform** | Pointer to [**NullableNestedPlatform**](NestedPlatform.md) |  | [optional] 
**PrimaryIp** | [**NestedIPAddress**](NestedIPAddress.md) |  | [readonly] 
**PrimaryIp4** | Pointer to [**NullableNestedIPAddress**](NestedIPAddress.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullableNestedIPAddress**](NestedIPAddress.md) |  | [optional] 
**Vcpus** | Pointer to **NullableFloat64** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Disk** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**LocalContextData** | Pointer to **map[string]interface{}** | Local config context data takes precedence over source contexts in the final rendered config context | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**ConfigContext** | **map[string]interface{}** |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**InterfaceCount** | **int32** |  | [readonly] 

## Methods

### NewVirtualMachineWithConfigContext

`func NewVirtualMachineWithConfigContext(id int32, url string, display string, name string, primaryIp NestedIPAddress, configContext map[string]interface{}, created NullableTime, lastUpdated NullableTime, interfaceCount int32, ) *VirtualMachineWithConfigContext`

NewVirtualMachineWithConfigContext instantiates a new VirtualMachineWithConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualMachineWithConfigContextWithDefaults

`func NewVirtualMachineWithConfigContextWithDefaults() *VirtualMachineWithConfigContext`

NewVirtualMachineWithConfigContextWithDefaults instantiates a new VirtualMachineWithConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualMachineWithConfigContext) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualMachineWithConfigContext) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualMachineWithConfigContext) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *VirtualMachineWithConfigContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VirtualMachineWithConfigContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VirtualMachineWithConfigContext) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *VirtualMachineWithConfigContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VirtualMachineWithConfigContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VirtualMachineWithConfigContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *VirtualMachineWithConfigContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualMachineWithConfigContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualMachineWithConfigContext) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *VirtualMachineWithConfigContext) GetStatus() ModuleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualMachineWithConfigContext) GetStatusOk() (*ModuleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualMachineWithConfigContext) SetStatus(v ModuleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualMachineWithConfigContext) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSite

`func (o *VirtualMachineWithConfigContext) GetSite() NestedSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *VirtualMachineWithConfigContext) GetSiteOk() (*NestedSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *VirtualMachineWithConfigContext) SetSite(v NestedSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *VirtualMachineWithConfigContext) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *VirtualMachineWithConfigContext) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *VirtualMachineWithConfigContext) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetCluster

`func (o *VirtualMachineWithConfigContext) GetCluster() NestedCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualMachineWithConfigContext) GetClusterOk() (*NestedCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualMachineWithConfigContext) SetCluster(v NestedCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualMachineWithConfigContext) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *VirtualMachineWithConfigContext) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *VirtualMachineWithConfigContext) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetDevice

`func (o *VirtualMachineWithConfigContext) GetDevice() NestedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VirtualMachineWithConfigContext) GetDeviceOk() (*NestedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VirtualMachineWithConfigContext) SetDevice(v NestedDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *VirtualMachineWithConfigContext) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *VirtualMachineWithConfigContext) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *VirtualMachineWithConfigContext) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetRole

`func (o *VirtualMachineWithConfigContext) GetRole() NestedDeviceRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VirtualMachineWithConfigContext) GetRoleOk() (*NestedDeviceRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VirtualMachineWithConfigContext) SetRole(v NestedDeviceRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *VirtualMachineWithConfigContext) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VirtualMachineWithConfigContext) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VirtualMachineWithConfigContext) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *VirtualMachineWithConfigContext) GetTenant() NestedTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VirtualMachineWithConfigContext) GetTenantOk() (*NestedTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VirtualMachineWithConfigContext) SetTenant(v NestedTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VirtualMachineWithConfigContext) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VirtualMachineWithConfigContext) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VirtualMachineWithConfigContext) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *VirtualMachineWithConfigContext) GetPlatform() NestedPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *VirtualMachineWithConfigContext) GetPlatformOk() (*NestedPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *VirtualMachineWithConfigContext) SetPlatform(v NestedPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *VirtualMachineWithConfigContext) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *VirtualMachineWithConfigContext) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *VirtualMachineWithConfigContext) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPrimaryIp

`func (o *VirtualMachineWithConfigContext) GetPrimaryIp() NestedIPAddress`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *VirtualMachineWithConfigContext) GetPrimaryIpOk() (*NestedIPAddress, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *VirtualMachineWithConfigContext) SetPrimaryIp(v NestedIPAddress)`

SetPrimaryIp sets PrimaryIp field to given value.


### GetPrimaryIp4

`func (o *VirtualMachineWithConfigContext) GetPrimaryIp4() NestedIPAddress`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *VirtualMachineWithConfigContext) GetPrimaryIp4Ok() (*NestedIPAddress, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *VirtualMachineWithConfigContext) SetPrimaryIp4(v NestedIPAddress)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *VirtualMachineWithConfigContext) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *VirtualMachineWithConfigContext) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *VirtualMachineWithConfigContext) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *VirtualMachineWithConfigContext) GetPrimaryIp6() NestedIPAddress`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *VirtualMachineWithConfigContext) GetPrimaryIp6Ok() (*NestedIPAddress, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *VirtualMachineWithConfigContext) SetPrimaryIp6(v NestedIPAddress)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *VirtualMachineWithConfigContext) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *VirtualMachineWithConfigContext) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *VirtualMachineWithConfigContext) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetVcpus

`func (o *VirtualMachineWithConfigContext) GetVcpus() float64`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *VirtualMachineWithConfigContext) GetVcpusOk() (*float64, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *VirtualMachineWithConfigContext) SetVcpus(v float64)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *VirtualMachineWithConfigContext) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *VirtualMachineWithConfigContext) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *VirtualMachineWithConfigContext) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetMemory

`func (o *VirtualMachineWithConfigContext) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VirtualMachineWithConfigContext) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VirtualMachineWithConfigContext) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VirtualMachineWithConfigContext) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *VirtualMachineWithConfigContext) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *VirtualMachineWithConfigContext) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetDisk

`func (o *VirtualMachineWithConfigContext) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *VirtualMachineWithConfigContext) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *VirtualMachineWithConfigContext) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *VirtualMachineWithConfigContext) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *VirtualMachineWithConfigContext) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *VirtualMachineWithConfigContext) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetDescription

`func (o *VirtualMachineWithConfigContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualMachineWithConfigContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualMachineWithConfigContext) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualMachineWithConfigContext) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComments

`func (o *VirtualMachineWithConfigContext) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VirtualMachineWithConfigContext) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VirtualMachineWithConfigContext) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VirtualMachineWithConfigContext) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLocalContextData

`func (o *VirtualMachineWithConfigContext) GetLocalContextData() map[string]interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *VirtualMachineWithConfigContext) GetLocalContextDataOk() (*map[string]interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *VirtualMachineWithConfigContext) SetLocalContextData(v map[string]interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *VirtualMachineWithConfigContext) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *VirtualMachineWithConfigContext) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *VirtualMachineWithConfigContext) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *VirtualMachineWithConfigContext) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualMachineWithConfigContext) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualMachineWithConfigContext) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualMachineWithConfigContext) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualMachineWithConfigContext) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualMachineWithConfigContext) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualMachineWithConfigContext) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualMachineWithConfigContext) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetConfigContext

`func (o *VirtualMachineWithConfigContext) GetConfigContext() map[string]interface{}`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *VirtualMachineWithConfigContext) GetConfigContextOk() (*map[string]interface{}, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *VirtualMachineWithConfigContext) SetConfigContext(v map[string]interface{})`

SetConfigContext sets ConfigContext field to given value.


### SetConfigContextNil

`func (o *VirtualMachineWithConfigContext) SetConfigContextNil(b bool)`

 SetConfigContextNil sets the value for ConfigContext to be an explicit nil

### UnsetConfigContext
`func (o *VirtualMachineWithConfigContext) UnsetConfigContext()`

UnsetConfigContext ensures that no value is present for ConfigContext, not even an explicit nil
### GetCreated

`func (o *VirtualMachineWithConfigContext) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VirtualMachineWithConfigContext) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VirtualMachineWithConfigContext) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VirtualMachineWithConfigContext) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VirtualMachineWithConfigContext) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VirtualMachineWithConfigContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VirtualMachineWithConfigContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VirtualMachineWithConfigContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VirtualMachineWithConfigContext) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VirtualMachineWithConfigContext) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetInterfaceCount

`func (o *VirtualMachineWithConfigContext) GetInterfaceCount() int32`

GetInterfaceCount returns the InterfaceCount field if non-nil, zero value otherwise.

### GetInterfaceCountOk

`func (o *VirtualMachineWithConfigContext) GetInterfaceCountOk() (*int32, bool)`

GetInterfaceCountOk returns a tuple with the InterfaceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceCount

`func (o *VirtualMachineWithConfigContext) SetInterfaceCount(v int32)`

SetInterfaceCount sets InterfaceCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


