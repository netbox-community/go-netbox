# WritableConfigContextRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
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
**Data** | **map[string]interface{}** |  | 

## Methods

### NewWritableConfigContextRequest

`func NewWritableConfigContextRequest(name string, data map[string]interface{}, ) *WritableConfigContextRequest`

NewWritableConfigContextRequest instantiates a new WritableConfigContextRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableConfigContextRequestWithDefaults

`func NewWritableConfigContextRequestWithDefaults() *WritableConfigContextRequest`

NewWritableConfigContextRequestWithDefaults instantiates a new WritableConfigContextRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableConfigContextRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableConfigContextRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableConfigContextRequest) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *WritableConfigContextRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *WritableConfigContextRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *WritableConfigContextRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *WritableConfigContextRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDescription

`func (o *WritableConfigContextRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableConfigContextRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableConfigContextRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableConfigContextRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *WritableConfigContextRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *WritableConfigContextRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *WritableConfigContextRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *WritableConfigContextRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetRegions

`func (o *WritableConfigContextRequest) GetRegions() []int32`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *WritableConfigContextRequest) GetRegionsOk() (*[]int32, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *WritableConfigContextRequest) SetRegions(v []int32)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *WritableConfigContextRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSiteGroups

`func (o *WritableConfigContextRequest) GetSiteGroups() []int32`

GetSiteGroups returns the SiteGroups field if non-nil, zero value otherwise.

### GetSiteGroupsOk

`func (o *WritableConfigContextRequest) GetSiteGroupsOk() (*[]int32, bool)`

GetSiteGroupsOk returns a tuple with the SiteGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteGroups

`func (o *WritableConfigContextRequest) SetSiteGroups(v []int32)`

SetSiteGroups sets SiteGroups field to given value.

### HasSiteGroups

`func (o *WritableConfigContextRequest) HasSiteGroups() bool`

HasSiteGroups returns a boolean if a field has been set.

### GetSites

`func (o *WritableConfigContextRequest) GetSites() []int32`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *WritableConfigContextRequest) GetSitesOk() (*[]int32, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *WritableConfigContextRequest) SetSites(v []int32)`

SetSites sets Sites field to given value.

### HasSites

`func (o *WritableConfigContextRequest) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetLocations

`func (o *WritableConfigContextRequest) GetLocations() []int32`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *WritableConfigContextRequest) GetLocationsOk() (*[]int32, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *WritableConfigContextRequest) SetLocations(v []int32)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *WritableConfigContextRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetDeviceTypes

`func (o *WritableConfigContextRequest) GetDeviceTypes() []int32`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *WritableConfigContextRequest) GetDeviceTypesOk() (*[]int32, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *WritableConfigContextRequest) SetDeviceTypes(v []int32)`

SetDeviceTypes sets DeviceTypes field to given value.

### HasDeviceTypes

`func (o *WritableConfigContextRequest) HasDeviceTypes() bool`

HasDeviceTypes returns a boolean if a field has been set.

### GetRoles

`func (o *WritableConfigContextRequest) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *WritableConfigContextRequest) GetRolesOk() (*[]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *WritableConfigContextRequest) SetRoles(v []int32)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *WritableConfigContextRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPlatforms

`func (o *WritableConfigContextRequest) GetPlatforms() []int32`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *WritableConfigContextRequest) GetPlatformsOk() (*[]int32, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *WritableConfigContextRequest) SetPlatforms(v []int32)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *WritableConfigContextRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetClusterTypes

`func (o *WritableConfigContextRequest) GetClusterTypes() []int32`

GetClusterTypes returns the ClusterTypes field if non-nil, zero value otherwise.

### GetClusterTypesOk

`func (o *WritableConfigContextRequest) GetClusterTypesOk() (*[]int32, bool)`

GetClusterTypesOk returns a tuple with the ClusterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTypes

`func (o *WritableConfigContextRequest) SetClusterTypes(v []int32)`

SetClusterTypes sets ClusterTypes field to given value.

### HasClusterTypes

`func (o *WritableConfigContextRequest) HasClusterTypes() bool`

HasClusterTypes returns a boolean if a field has been set.

### GetClusterGroups

`func (o *WritableConfigContextRequest) GetClusterGroups() []int32`

GetClusterGroups returns the ClusterGroups field if non-nil, zero value otherwise.

### GetClusterGroupsOk

`func (o *WritableConfigContextRequest) GetClusterGroupsOk() (*[]int32, bool)`

GetClusterGroupsOk returns a tuple with the ClusterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroups

`func (o *WritableConfigContextRequest) SetClusterGroups(v []int32)`

SetClusterGroups sets ClusterGroups field to given value.

### HasClusterGroups

`func (o *WritableConfigContextRequest) HasClusterGroups() bool`

HasClusterGroups returns a boolean if a field has been set.

### GetClusters

`func (o *WritableConfigContextRequest) GetClusters() []int32`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *WritableConfigContextRequest) GetClustersOk() (*[]int32, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *WritableConfigContextRequest) SetClusters(v []int32)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *WritableConfigContextRequest) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTenantGroups

`func (o *WritableConfigContextRequest) GetTenantGroups() []int32`

GetTenantGroups returns the TenantGroups field if non-nil, zero value otherwise.

### GetTenantGroupsOk

`func (o *WritableConfigContextRequest) GetTenantGroupsOk() (*[]int32, bool)`

GetTenantGroupsOk returns a tuple with the TenantGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroups

`func (o *WritableConfigContextRequest) SetTenantGroups(v []int32)`

SetTenantGroups sets TenantGroups field to given value.

### HasTenantGroups

`func (o *WritableConfigContextRequest) HasTenantGroups() bool`

HasTenantGroups returns a boolean if a field has been set.

### GetTenants

`func (o *WritableConfigContextRequest) GetTenants() []int32`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *WritableConfigContextRequest) GetTenantsOk() (*[]int32, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *WritableConfigContextRequest) SetTenants(v []int32)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *WritableConfigContextRequest) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetTags

`func (o *WritableConfigContextRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritableConfigContextRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritableConfigContextRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritableConfigContextRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDataSource

`func (o *WritableConfigContextRequest) GetDataSource() int32`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *WritableConfigContextRequest) GetDataSourceOk() (*int32, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *WritableConfigContextRequest) SetDataSource(v int32)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *WritableConfigContextRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### SetDataSourceNil

`func (o *WritableConfigContextRequest) SetDataSourceNil(b bool)`

 SetDataSourceNil sets the value for DataSource to be an explicit nil

### UnsetDataSource
`func (o *WritableConfigContextRequest) UnsetDataSource()`

UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil
### GetData

`func (o *WritableConfigContextRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WritableConfigContextRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WritableConfigContextRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


