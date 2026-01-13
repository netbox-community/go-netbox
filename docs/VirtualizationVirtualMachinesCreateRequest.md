# VirtualizationVirtualMachinesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Status** | Pointer to [**PatchedWritableVirtualMachineWithConfigContextRequestStatus**](PatchedWritableVirtualMachineWithConfigContextRequestStatus.md) |  | [optional] 
**StartOnBoot** | Pointer to [**PatchedWritableVirtualMachineWithConfigContextRequestStartOnBoot**](PatchedWritableVirtualMachineWithConfigContextRequestStartOnBoot.md) |  | [optional] 
**Site** | Pointer to [**NullablePatchedWritableVLANRequestSite**](PatchedWritableVLANRequestSite.md) |  | [optional] 
**Cluster** | Pointer to [**NullableDeviceWithConfigContextRequestCluster**](DeviceWithConfigContextRequestCluster.md) |  | [optional] 
**Device** | Pointer to [**NullableDeviceBayRequestInstalledDevice**](DeviceBayRequestInstalledDevice.md) |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**NullablePatchedWritableVirtualMachineWithConfigContextRequestRole**](PatchedWritableVirtualMachineWithConfigContextRequestRole.md) |  | [optional] 
**Tenant** | Pointer to [**NullableASNRangeRequestTenant**](ASNRangeRequestTenant.md) |  | [optional] 
**Platform** | Pointer to [**NullableDeviceTypeRequestDefaultPlatform**](DeviceTypeRequestDefaultPlatform.md) |  | [optional] 
**PrimaryIp4** | Pointer to [**NullableDeviceWithConfigContextRequestPrimaryIp4**](DeviceWithConfigContextRequestPrimaryIp4.md) |  | [optional] 
**PrimaryIp6** | Pointer to [**NullableDeviceWithConfigContextRequestPrimaryIp4**](DeviceWithConfigContextRequestPrimaryIp4.md) |  | [optional] 
**Vcpus** | Pointer to **NullableFloat64** |  | [optional] 
**Memory** | Pointer to **NullableInt32** |  | [optional] 
**Disk** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**ConfigTemplate** | Pointer to [**NullableDeviceRoleRequestConfigTemplate**](DeviceRoleRequestConfigTemplate.md) |  | [optional] 
**LocalContextData** | Pointer to **interface{}** | Local config context data takes precedence over source contexts in the final rendered config context | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewVirtualizationVirtualMachinesCreateRequest

`func NewVirtualizationVirtualMachinesCreateRequest(name string, ) *VirtualizationVirtualMachinesCreateRequest`

NewVirtualizationVirtualMachinesCreateRequest instantiates a new VirtualizationVirtualMachinesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualMachinesCreateRequestWithDefaults

`func NewVirtualizationVirtualMachinesCreateRequestWithDefaults() *VirtualizationVirtualMachinesCreateRequest`

NewVirtualizationVirtualMachinesCreateRequestWithDefaults instantiates a new VirtualizationVirtualMachinesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualizationVirtualMachinesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVirtualMachinesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *VirtualizationVirtualMachinesCreateRequest) GetStatus() PatchedWritableVirtualMachineWithConfigContextRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetStatusOk() (*PatchedWritableVirtualMachineWithConfigContextRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationVirtualMachinesCreateRequest) SetStatus(v PatchedWritableVirtualMachineWithConfigContextRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationVirtualMachinesCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartOnBoot

`func (o *VirtualizationVirtualMachinesCreateRequest) GetStartOnBoot() PatchedWritableVirtualMachineWithConfigContextRequestStartOnBoot`

GetStartOnBoot returns the StartOnBoot field if non-nil, zero value otherwise.

### GetStartOnBootOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetStartOnBootOk() (*PatchedWritableVirtualMachineWithConfigContextRequestStartOnBoot, bool)`

GetStartOnBootOk returns a tuple with the StartOnBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOnBoot

`func (o *VirtualizationVirtualMachinesCreateRequest) SetStartOnBoot(v PatchedWritableVirtualMachineWithConfigContextRequestStartOnBoot)`

SetStartOnBoot sets StartOnBoot field to given value.

### HasStartOnBoot

`func (o *VirtualizationVirtualMachinesCreateRequest) HasStartOnBoot() bool`

HasStartOnBoot returns a boolean if a field has been set.

### GetSite

`func (o *VirtualizationVirtualMachinesCreateRequest) GetSite() PatchedWritableVLANRequestSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetSiteOk() (*PatchedWritableVLANRequestSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *VirtualizationVirtualMachinesCreateRequest) SetSite(v PatchedWritableVLANRequestSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *VirtualizationVirtualMachinesCreateRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetCluster

`func (o *VirtualizationVirtualMachinesCreateRequest) GetCluster() DeviceWithConfigContextRequestCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetClusterOk() (*DeviceWithConfigContextRequestCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVirtualMachinesCreateRequest) SetCluster(v DeviceWithConfigContextRequestCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVirtualMachinesCreateRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetDevice

`func (o *VirtualizationVirtualMachinesCreateRequest) GetDevice() DeviceBayRequestInstalledDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetDeviceOk() (*DeviceBayRequestInstalledDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VirtualizationVirtualMachinesCreateRequest) SetDevice(v DeviceBayRequestInstalledDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *VirtualizationVirtualMachinesCreateRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetSerial

`func (o *VirtualizationVirtualMachinesCreateRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *VirtualizationVirtualMachinesCreateRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *VirtualizationVirtualMachinesCreateRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetRole

`func (o *VirtualizationVirtualMachinesCreateRequest) GetRole() PatchedWritableVirtualMachineWithConfigContextRequestRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetRoleOk() (*PatchedWritableVirtualMachineWithConfigContextRequestRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VirtualizationVirtualMachinesCreateRequest) SetRole(v PatchedWritableVirtualMachineWithConfigContextRequestRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *VirtualizationVirtualMachinesCreateRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *VirtualizationVirtualMachinesCreateRequest) GetTenant() ASNRangeRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetTenantOk() (*ASNRangeRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VirtualizationVirtualMachinesCreateRequest) SetTenant(v ASNRangeRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VirtualizationVirtualMachinesCreateRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetPlatform

`func (o *VirtualizationVirtualMachinesCreateRequest) GetPlatform() DeviceTypeRequestDefaultPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetPlatformOk() (*DeviceTypeRequestDefaultPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *VirtualizationVirtualMachinesCreateRequest) SetPlatform(v DeviceTypeRequestDefaultPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *VirtualizationVirtualMachinesCreateRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPrimaryIp4

`func (o *VirtualizationVirtualMachinesCreateRequest) GetPrimaryIp4() DeviceWithConfigContextRequestPrimaryIp4`

GetPrimaryIp4 returns the PrimaryIp4 field if non-nil, zero value otherwise.

### GetPrimaryIp4Ok

`func (o *VirtualizationVirtualMachinesCreateRequest) GetPrimaryIp4Ok() (*DeviceWithConfigContextRequestPrimaryIp4, bool)`

GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp4

`func (o *VirtualizationVirtualMachinesCreateRequest) SetPrimaryIp4(v DeviceWithConfigContextRequestPrimaryIp4)`

SetPrimaryIp4 sets PrimaryIp4 field to given value.

### HasPrimaryIp4

`func (o *VirtualizationVirtualMachinesCreateRequest) HasPrimaryIp4() bool`

HasPrimaryIp4 returns a boolean if a field has been set.

### SetPrimaryIp4Nil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetPrimaryIp4Nil(b bool)`

 SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil

### UnsetPrimaryIp4
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetPrimaryIp4()`

UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
### GetPrimaryIp6

`func (o *VirtualizationVirtualMachinesCreateRequest) GetPrimaryIp6() DeviceWithConfigContextRequestPrimaryIp4`

GetPrimaryIp6 returns the PrimaryIp6 field if non-nil, zero value otherwise.

### GetPrimaryIp6Ok

`func (o *VirtualizationVirtualMachinesCreateRequest) GetPrimaryIp6Ok() (*DeviceWithConfigContextRequestPrimaryIp4, bool)`

GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp6

`func (o *VirtualizationVirtualMachinesCreateRequest) SetPrimaryIp6(v DeviceWithConfigContextRequestPrimaryIp4)`

SetPrimaryIp6 sets PrimaryIp6 field to given value.

### HasPrimaryIp6

`func (o *VirtualizationVirtualMachinesCreateRequest) HasPrimaryIp6() bool`

HasPrimaryIp6 returns a boolean if a field has been set.

### SetPrimaryIp6Nil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetPrimaryIp6Nil(b bool)`

 SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil

### UnsetPrimaryIp6
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetPrimaryIp6()`

UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
### GetVcpus

`func (o *VirtualizationVirtualMachinesCreateRequest) GetVcpus() float64`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetVcpusOk() (*float64, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *VirtualizationVirtualMachinesCreateRequest) SetVcpus(v float64)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *VirtualizationVirtualMachinesCreateRequest) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### SetVcpusNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetVcpusNil(b bool)`

 SetVcpusNil sets the value for Vcpus to be an explicit nil

### UnsetVcpus
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetVcpus()`

UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
### GetMemory

`func (o *VirtualizationVirtualMachinesCreateRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VirtualizationVirtualMachinesCreateRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VirtualizationVirtualMachinesCreateRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### SetMemoryNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetMemoryNil(b bool)`

 SetMemoryNil sets the value for Memory to be an explicit nil

### UnsetMemory
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetMemory()`

UnsetMemory ensures that no value is present for Memory, not even an explicit nil
### GetDisk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *VirtualizationVirtualMachinesCreateRequest) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *VirtualizationVirtualMachinesCreateRequest) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### SetDiskNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetDiskNil(b bool)`

 SetDiskNil sets the value for Disk to be an explicit nil

### UnsetDisk
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetDisk()`

UnsetDisk ensures that no value is present for Disk, not even an explicit nil
### GetDescription

`func (o *VirtualizationVirtualMachinesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualizationVirtualMachinesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualizationVirtualMachinesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *VirtualizationVirtualMachinesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VirtualizationVirtualMachinesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VirtualizationVirtualMachinesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *VirtualizationVirtualMachinesCreateRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *VirtualizationVirtualMachinesCreateRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *VirtualizationVirtualMachinesCreateRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *VirtualizationVirtualMachinesCreateRequest) GetConfigTemplate() DeviceRoleRequestConfigTemplate`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetConfigTemplateOk() (*DeviceRoleRequestConfigTemplate, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *VirtualizationVirtualMachinesCreateRequest) SetConfigTemplate(v DeviceRoleRequestConfigTemplate)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *VirtualizationVirtualMachinesCreateRequest) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### SetConfigTemplateNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetConfigTemplateNil(b bool)`

 SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil

### UnsetConfigTemplate
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetConfigTemplate()`

UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
### GetLocalContextData

`func (o *VirtualizationVirtualMachinesCreateRequest) GetLocalContextData() interface{}`

GetLocalContextData returns the LocalContextData field if non-nil, zero value otherwise.

### GetLocalContextDataOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetLocalContextDataOk() (*interface{}, bool)`

GetLocalContextDataOk returns a tuple with the LocalContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalContextData

`func (o *VirtualizationVirtualMachinesCreateRequest) SetLocalContextData(v interface{})`

SetLocalContextData sets LocalContextData field to given value.

### HasLocalContextData

`func (o *VirtualizationVirtualMachinesCreateRequest) HasLocalContextData() bool`

HasLocalContextData returns a boolean if a field has been set.

### SetLocalContextDataNil

`func (o *VirtualizationVirtualMachinesCreateRequest) SetLocalContextDataNil(b bool)`

 SetLocalContextDataNil sets the value for LocalContextData to be an explicit nil

### UnsetLocalContextData
`func (o *VirtualizationVirtualMachinesCreateRequest) UnsetLocalContextData()`

UnsetLocalContextData ensures that no value is present for LocalContextData, not even an explicit nil
### GetTags

`func (o *VirtualizationVirtualMachinesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualizationVirtualMachinesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualizationVirtualMachinesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *VirtualizationVirtualMachinesCreateRequest) GetCustomFields() map[string]map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VirtualizationVirtualMachinesCreateRequest) GetCustomFieldsOk() (*map[string]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VirtualizationVirtualMachinesCreateRequest) SetCustomFields(v map[string]map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VirtualizationVirtualMachinesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


