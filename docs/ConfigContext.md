# ConfigContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Weight** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Regions** | Pointer to [**[]Region**](Region.md) |  | [optional] 
**SiteGroups** | Pointer to [**[]SiteGroup**](SiteGroup.md) |  | [optional] 
**Sites** | Pointer to [**[]Site**](Site.md) |  | [optional] 
**Locations** | Pointer to [**[]Location**](Location.md) |  | [optional] 
**DeviceTypes** | Pointer to [**[]DeviceType**](DeviceType.md) |  | [optional] 
**Roles** | Pointer to [**[]DeviceRole**](DeviceRole.md) |  | [optional] 
**Platforms** | Pointer to [**[]Platform**](Platform.md) |  | [optional] 
**ClusterTypes** | Pointer to [**[]ClusterType**](ClusterType.md) |  | [optional] 
**ClusterGroups** | Pointer to [**[]ClusterGroup**](ClusterGroup.md) |  | [optional] 
**Clusters** | Pointer to [**[]Cluster**](Cluster.md) |  | [optional] 
**TenantGroups** | Pointer to [**[]TenantGroup**](TenantGroup.md) |  | [optional] 
**Tenants** | Pointer to [**[]Tenant**](Tenant.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**DataSource** | Pointer to [**BriefDataSource**](BriefDataSource.md) |  | [optional] 
**DataPath** | **string** | Path to remote file (relative to data source root) | [readonly] 
**DataFile** | [**BriefDataFile**](BriefDataFile.md) |  | [readonly] 
**DataSynced** | **NullableTime** |  | [readonly] 
**Data** | **interface{}** |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewConfigContext

`func NewConfigContext(id int32, url string, displayUrl string, display string, name string, dataPath string, dataFile BriefDataFile, dataSynced NullableTime, data interface{}, created NullableTime, lastUpdated NullableTime, ) *ConfigContext`

NewConfigContext instantiates a new ConfigContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigContextWithDefaults

`func NewConfigContextWithDefaults() *ConfigContext`

NewConfigContextWithDefaults instantiates a new ConfigContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigContext) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigContext) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigContext) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ConfigContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigContext) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *ConfigContext) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *ConfigContext) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *ConfigContext) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.


### GetDisplay

`func (o *ConfigContext) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConfigContext) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConfigContext) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *ConfigContext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigContext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigContext) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *ConfigContext) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ConfigContext) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ConfigContext) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ConfigContext) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDescription

`func (o *ConfigContext) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigContext) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigContext) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigContext) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsActive

`func (o *ConfigContext) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ConfigContext) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ConfigContext) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ConfigContext) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetRegions

`func (o *ConfigContext) GetRegions() []Region`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ConfigContext) GetRegionsOk() (*[]Region, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ConfigContext) SetRegions(v []Region)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *ConfigContext) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSiteGroups

`func (o *ConfigContext) GetSiteGroups() []SiteGroup`

GetSiteGroups returns the SiteGroups field if non-nil, zero value otherwise.

### GetSiteGroupsOk

`func (o *ConfigContext) GetSiteGroupsOk() (*[]SiteGroup, bool)`

GetSiteGroupsOk returns a tuple with the SiteGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteGroups

`func (o *ConfigContext) SetSiteGroups(v []SiteGroup)`

SetSiteGroups sets SiteGroups field to given value.

### HasSiteGroups

`func (o *ConfigContext) HasSiteGroups() bool`

HasSiteGroups returns a boolean if a field has been set.

### GetSites

`func (o *ConfigContext) GetSites() []Site`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ConfigContext) GetSitesOk() (*[]Site, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ConfigContext) SetSites(v []Site)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ConfigContext) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetLocations

`func (o *ConfigContext) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ConfigContext) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ConfigContext) SetLocations(v []Location)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ConfigContext) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetDeviceTypes

`func (o *ConfigContext) GetDeviceTypes() []DeviceType`

GetDeviceTypes returns the DeviceTypes field if non-nil, zero value otherwise.

### GetDeviceTypesOk

`func (o *ConfigContext) GetDeviceTypesOk() (*[]DeviceType, bool)`

GetDeviceTypesOk returns a tuple with the DeviceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTypes

`func (o *ConfigContext) SetDeviceTypes(v []DeviceType)`

SetDeviceTypes sets DeviceTypes field to given value.

### HasDeviceTypes

`func (o *ConfigContext) HasDeviceTypes() bool`

HasDeviceTypes returns a boolean if a field has been set.

### GetRoles

`func (o *ConfigContext) GetRoles() []DeviceRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ConfigContext) GetRolesOk() (*[]DeviceRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ConfigContext) SetRoles(v []DeviceRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ConfigContext) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPlatforms

`func (o *ConfigContext) GetPlatforms() []Platform`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ConfigContext) GetPlatformsOk() (*[]Platform, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ConfigContext) SetPlatforms(v []Platform)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *ConfigContext) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetClusterTypes

`func (o *ConfigContext) GetClusterTypes() []ClusterType`

GetClusterTypes returns the ClusterTypes field if non-nil, zero value otherwise.

### GetClusterTypesOk

`func (o *ConfigContext) GetClusterTypesOk() (*[]ClusterType, bool)`

GetClusterTypesOk returns a tuple with the ClusterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTypes

`func (o *ConfigContext) SetClusterTypes(v []ClusterType)`

SetClusterTypes sets ClusterTypes field to given value.

### HasClusterTypes

`func (o *ConfigContext) HasClusterTypes() bool`

HasClusterTypes returns a boolean if a field has been set.

### GetClusterGroups

`func (o *ConfigContext) GetClusterGroups() []ClusterGroup`

GetClusterGroups returns the ClusterGroups field if non-nil, zero value otherwise.

### GetClusterGroupsOk

`func (o *ConfigContext) GetClusterGroupsOk() (*[]ClusterGroup, bool)`

GetClusterGroupsOk returns a tuple with the ClusterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGroups

`func (o *ConfigContext) SetClusterGroups(v []ClusterGroup)`

SetClusterGroups sets ClusterGroups field to given value.

### HasClusterGroups

`func (o *ConfigContext) HasClusterGroups() bool`

HasClusterGroups returns a boolean if a field has been set.

### GetClusters

`func (o *ConfigContext) GetClusters() []Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ConfigContext) GetClustersOk() (*[]Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ConfigContext) SetClusters(v []Cluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *ConfigContext) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetTenantGroups

`func (o *ConfigContext) GetTenantGroups() []TenantGroup`

GetTenantGroups returns the TenantGroups field if non-nil, zero value otherwise.

### GetTenantGroupsOk

`func (o *ConfigContext) GetTenantGroupsOk() (*[]TenantGroup, bool)`

GetTenantGroupsOk returns a tuple with the TenantGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantGroups

`func (o *ConfigContext) SetTenantGroups(v []TenantGroup)`

SetTenantGroups sets TenantGroups field to given value.

### HasTenantGroups

`func (o *ConfigContext) HasTenantGroups() bool`

HasTenantGroups returns a boolean if a field has been set.

### GetTenants

`func (o *ConfigContext) GetTenants() []Tenant`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ConfigContext) GetTenantsOk() (*[]Tenant, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ConfigContext) SetTenants(v []Tenant)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ConfigContext) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetTags

`func (o *ConfigContext) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigContext) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigContext) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigContext) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDataSource

`func (o *ConfigContext) GetDataSource() BriefDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ConfigContext) GetDataSourceOk() (*BriefDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ConfigContext) SetDataSource(v BriefDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *ConfigContext) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataPath

`func (o *ConfigContext) GetDataPath() string`

GetDataPath returns the DataPath field if non-nil, zero value otherwise.

### GetDataPathOk

`func (o *ConfigContext) GetDataPathOk() (*string, bool)`

GetDataPathOk returns a tuple with the DataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPath

`func (o *ConfigContext) SetDataPath(v string)`

SetDataPath sets DataPath field to given value.


### GetDataFile

`func (o *ConfigContext) GetDataFile() BriefDataFile`

GetDataFile returns the DataFile field if non-nil, zero value otherwise.

### GetDataFileOk

`func (o *ConfigContext) GetDataFileOk() (*BriefDataFile, bool)`

GetDataFileOk returns a tuple with the DataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFile

`func (o *ConfigContext) SetDataFile(v BriefDataFile)`

SetDataFile sets DataFile field to given value.


### GetDataSynced

`func (o *ConfigContext) GetDataSynced() time.Time`

GetDataSynced returns the DataSynced field if non-nil, zero value otherwise.

### GetDataSyncedOk

`func (o *ConfigContext) GetDataSyncedOk() (*time.Time, bool)`

GetDataSyncedOk returns a tuple with the DataSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSynced

`func (o *ConfigContext) SetDataSynced(v time.Time)`

SetDataSynced sets DataSynced field to given value.


### SetDataSyncedNil

`func (o *ConfigContext) SetDataSyncedNil(b bool)`

 SetDataSyncedNil sets the value for DataSynced to be an explicit nil

### UnsetDataSynced
`func (o *ConfigContext) UnsetDataSynced()`

UnsetDataSynced ensures that no value is present for DataSynced, not even an explicit nil
### GetData

`func (o *ConfigContext) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConfigContext) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConfigContext) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *ConfigContext) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ConfigContext) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetCreated

`func (o *ConfigContext) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConfigContext) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConfigContext) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ConfigContext) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ConfigContext) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ConfigContext) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ConfigContext) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ConfigContext) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ConfigContext) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ConfigContext) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


