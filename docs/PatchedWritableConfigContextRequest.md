# PatchedWritableConfigContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Regions** | Pointer to **[]int32** |  | [optional] 
**SiteGroups** | Pointer to **[]int32** |  | [optional] 
**Sites** | Pointer to **[]int32** |  | [optional] 
**Locations** | Pointer to **[]int32** |  | [optional] 
**DeviceTypes** | Pointer to **[]int32** |  | [optional] 
**Roles** | Pointer to **[]int32** |  | [optional] 
**Platforms** | Pointer to **[]int32** |  | [optional] 
**ClusterTypes** | Pointer to **[]int32** |  | [optional] 
**ClusterGroups** | Pointer to **[]int32** |  | [optional] 
**Clusters** | Pointer to **[]int32** |  | [optional] 
**TenantGroups** | Pointer to **[]int32** |  | [optional] 
**Tenants** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**DataSource** | Pointer to **NullableInt32** | Remote data source | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPatchedWritableConfigContextRequest

`func NewPatchedWritableConfigContextRequest() *PatchedWritableConfigContextRequest`

NewPatchedWritableConfigContextRequest instantiates a new PatchedWritableConfigContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableConfigContextRequestWithDefaults

`func NewPatchedWritableConfigContextRequestWithDefaults() *PatchedWritableConfigContextRequest`

NewPatchedWritableConfigContextRequestWithDefaults instantiates a new PatchedWritableConfigContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableConfigContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableConfigContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableConfigContextRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableConfigContextRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedWritableConfigContextRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedWritableConfigContextRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedWritableConfigContextRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedWritableConfigContextRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableConfigContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableConfigContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableConfigContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableConfigContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *PatchedWritableConfigContextRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PatchedWritableConfigContextRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PatchedWritableConfigContextRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PatchedWritableConfigContextRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetRegions

`func (o *PatchedWritableConfigContextRequest) GetRegions() []int32`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PatchedWritableConfigContextRequest) GetRegionsOk() (*[]int32, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PatchedWritableConfigContextRequest) SetRegions(v []int32)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PatchedWritableConfigContextRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSiteGroups

`func (o *PatchedWritableConfigContextRequest) GetSiteGroups() []int32`

GetSiteGroups returns the SiteGroups field if non-nil, zero value otherwise.

### GetSiteGroupsOk

`func (o *PatchedWritableConfigContextRequest) GetSiteGroupsOk() (*[]int32, bool)`

GetSiteGroupsOk returns a tuple with the SiteGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteGroups

`func (o *PatchedWritableConfigContextRequest) SetSiteGroups(v []int32)`

SetSiteGroups sets SiteGroups field to given value.

### HasSiteGroups

`func (o *PatchedWritableConfigContextRequest) HasSiteGroups() bool`

HasSiteGroups returns a boolean if a field has been set.

### GetSites

`func (o *PatchedWritableConfigContextRequest) GetSites() []int32`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *PatchedWritableConfigContextRequest) GetSitesOk() (*[]int32, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *PatchedWritableConfigContextRequest) SetSites(v []int32)`

SetSites sets Sites field to given value.

### HasSites

`func (o *PatchedWritableConfigContextRequest) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetLocations

`func (o *PatchedWritableConfigContextRequest) GetLocations() []int32`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *PatchedWritableConfigContextRequest) GetLocationsOk() (*[]int32, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *PatchedWritableConfigContextRequest) SetLocations(v []int32)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *PatchedWritableConfigContextRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetDeviceTypes

`func (o *PatchedWritableConfigContextRequest) GetDeviceTypes() []int32`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *PatchedWritableConfigContextRequest) GetDeviceTypesOk() (*[]int32, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *PatchedWritableConfigContextRequest) SetDeviceTypes(v []int32)`

SetDeviceTypes sets DeviceTypes field to given value.

### HasDeviceTypes

`func (o *PatchedWritableConfigContextRequest) HasDeviceTypes() bool`

HasDeviceTypes returns a boolean if a field has been set.

### GetRoles

`func (o *PatchedWritableConfigContextRequest) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PatchedWritableConfigContextRequest) GetRolesOk() (*[]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PatchedWritableConfigContextRequest) SetRoles(v []int32)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PatchedWritableConfigContextRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPlatforms

`func (o *PatchedWritableConfigContextRequest) GetPlatforms() []int32`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *PatchedWritableConfigContextRequest) GetPlatformsOk() (*[]int32, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *PatchedWritableConfigContextRequest) SetPlatforms(v []int32)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *PatchedWritableConfigContextRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetClusterTypes

`func (o *PatchedWritableConfigContextRequest) GetClusterTypes() []int32`

GetClusterTypes returns the ClusterTypes field if non-nil, zero value otherwise.

### GetClusterTypesOk

`func (o *PatchedWritableConfigContextRequest) GetClusterTypesOk() (*[]int32, bool)`

GetClusterTypesOk returns a tuple with the ClusterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTypes

`func (o *PatchedWritableConfigContextRequest) SetClusterTypes(v []int32)`

SetClusterTypes sets ClusterTypes field to given value.

### HasClusterTypes

`func (o *PatchedWritableConfigContextRequest) HasClusterTypes() bool`

HasClusterTypes returns a boolean if a field has been set.

### GetClusterGroups

`func (o *PatchedWritableConfigContextRequest) GetClusterGroups() []int32`

GetClusterGroups returns the ClusterGroups field if non-nil, zero value otherwise.

### GetClusterGroupsOk

`func (o *PatchedWritableConfigContextRequest) GetClusterGroupsOk() (*[]int32, bool)`

GetClusterGroupsOk returns a tuple with the ClusterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroups

`func (o *PatchedWritableConfigContextRequest) SetClusterGroups(v []int32)`

SetClusterGroups sets ClusterGroups field to given value.

### HasClusterGroups

`func (o *PatchedWritableConfigContextRequest) HasClusterGroups() bool`

HasClusterGroups returns a boolean if a field has been set.

### GetClusters

`func (o *PatchedWritableConfigContextRequest) GetClusters() []int32`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *PatchedWritableConfigContextRequest) GetClustersOk() (*[]int32, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *PatchedWritableConfigContextRequest) SetClusters(v []int32)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *PatchedWritableConfigContextRequest) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTenantGroups

`func (o *PatchedWritableConfigContextRequest) GetTenantGroups() []int32`

GetTenantGroups returns the TenantGroups field if non-nil, zero value otherwise.

### GetTenantGroupsOk

`func (o *PatchedWritableConfigContextRequest) GetTenantGroupsOk() (*[]int32, bool)`

GetTenantGroupsOk returns a tuple with the TenantGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroups

`func (o *PatchedWritableConfigContextRequest) SetTenantGroups(v []int32)`

SetTenantGroups sets TenantGroups field to given value.

### HasTenantGroups

`func (o *PatchedWritableConfigContextRequest) HasTenantGroups() bool`

HasTenantGroups returns a boolean if a field has been set.

### GetTenants

`func (o *PatchedWritableConfigContextRequest) GetTenants() []int32`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *PatchedWritableConfigContextRequest) GetTenantsOk() (*[]int32, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *PatchedWritableConfigContextRequest) SetTenants(v []int32)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *PatchedWritableConfigContextRequest) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableConfigContextRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableConfigContextRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableConfigContextRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableConfigContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDataSource

`func (o *PatchedWritableConfigContextRequest) GetDataSource() int32`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *PatchedWritableConfigContextRequest) GetDataSourceOk() (*int32, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *PatchedWritableConfigContextRequest) SetDataSource(v int32)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *PatchedWritableConfigContextRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### SetDataSourceNil

`func (o *PatchedWritableConfigContextRequest) SetDataSourceNil(b bool)`

 SetDataSourceNil sets the value for DataSource to be an explicit nil

### UnsetDataSource
`func (o *PatchedWritableConfigContextRequest) UnsetDataSource()`

UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
### GetData

`func (o *PatchedWritableConfigContextRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PatchedWritableConfigContextRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PatchedWritableConfigContextRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *PatchedWritableConfigContextRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


