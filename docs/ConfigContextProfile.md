# ConfigContextProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**DisplayUrl** | Pointer to **string** |  | [optional] [readonly] 
**Display** | **string** |  | [readonly] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Schema** | Pointer to **interface{}** | A JSON schema specifying the structure of the context data for this profile | [optional] 
**Tags** | Pointer to [**[]NestedTag**](NestedTag.md) |  | [optional] 
**Owner** | Pointer to [**NullableBriefOwner**](BriefOwner.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**DataSource** | Pointer to [**BriefDataSource**](BriefDataSource.md) |  | [optional] 
**DataPath** | **string** | Path to remote file (relative to data source root) | [readonly] 
**DataFile** | [**BriefDataFile**](BriefDataFile.md) |  | [readonly] 
**DataSynced** | Pointer to **NullableTime** |  | [optional] [readonly] 
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] [readonly] 

## Methods

### NewConfigContextProfile

`func NewConfigContextProfile(id int32, url string, display string, name string, dataPath string, dataFile BriefDataFile, ) *ConfigContextProfile`

NewConfigContextProfile instantiates a new ConfigContextProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigContextProfileWithDefaults

`func NewConfigContextProfileWithDefaults() *ConfigContextProfile`

NewConfigContextProfileWithDefaults instantiates a new ConfigContextProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigContextProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigContextProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigContextProfile) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ConfigContextProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConfigContextProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConfigContextProfile) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplayUrl

`func (o *ConfigContextProfile) GetDisplayUrl() string`

GetDisplayUrl returns the DisplayUrl field if non-nil, zero value otherwise.

### GetDisplayUrlOk

`func (o *ConfigContextProfile) GetDisplayUrlOk() (*string, bool)`

GetDisplayUrlOk returns a tuple with the DisplayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUrl

`func (o *ConfigContextProfile) SetDisplayUrl(v string)`

SetDisplayUrl sets DisplayUrl field to given value.

### HasDisplayUrl

`func (o *ConfigContextProfile) HasDisplayUrl() bool`

HasDisplayUrl returns a boolean if a field has been set.

### GetDisplay

`func (o *ConfigContextProfile) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConfigContextProfile) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConfigContextProfile) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetName

`func (o *ConfigContextProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigContextProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigContextProfile) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConfigContextProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigContextProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigContextProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigContextProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSchema

`func (o *ConfigContextProfile) GetSchema() interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ConfigContextProfile) GetSchemaOk() (*interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ConfigContextProfile) SetSchema(v interface{})`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ConfigContextProfile) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *ConfigContextProfile) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *ConfigContextProfile) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil
### GetTags

`func (o *ConfigContextProfile) GetTags() []NestedTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigContextProfile) GetTagsOk() (*[]NestedTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigContextProfile) SetTags(v []NestedTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigContextProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOwner

`func (o *ConfigContextProfile) GetOwner() BriefOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ConfigContextProfile) GetOwnerOk() (*BriefOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ConfigContextProfile) SetOwner(v BriefOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ConfigContextProfile) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ConfigContextProfile) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ConfigContextProfile) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetComments

`func (o *ConfigContextProfile) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ConfigContextProfile) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ConfigContextProfile) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ConfigContextProfile) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetDataSource

`func (o *ConfigContextProfile) GetDataSource() BriefDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ConfigContextProfile) GetDataSourceOk() (*BriefDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ConfigContextProfile) SetDataSource(v BriefDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *ConfigContextProfile) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataPath

`func (o *ConfigContextProfile) GetDataPath() string`

GetDataPath returns the DataPath field if non-nil, zero value otherwise.

### GetDataPathOk

`func (o *ConfigContextProfile) GetDataPathOk() (*string, bool)`

GetDataPathOk returns a tuple with the DataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPath

`func (o *ConfigContextProfile) SetDataPath(v string)`

SetDataPath sets DataPath field to given value.


### GetDataFile

`func (o *ConfigContextProfile) GetDataFile() BriefDataFile`

GetDataFile returns the DataFile field if non-nil, zero value otherwise.

### GetDataFileOk

`func (o *ConfigContextProfile) GetDataFileOk() (*BriefDataFile, bool)`

GetDataFileOk returns a tuple with the DataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFile

`func (o *ConfigContextProfile) SetDataFile(v BriefDataFile)`

SetDataFile sets DataFile field to given value.


### GetDataSynced

`func (o *ConfigContextProfile) GetDataSynced() time.Time`

GetDataSynced returns the DataSynced field if non-nil, zero value otherwise.

### GetDataSyncedOk

`func (o *ConfigContextProfile) GetDataSyncedOk() (*time.Time, bool)`

GetDataSyncedOk returns a tuple with the DataSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSynced

`func (o *ConfigContextProfile) SetDataSynced(v time.Time)`

SetDataSynced sets DataSynced field to given value.

### HasDataSynced

`func (o *ConfigContextProfile) HasDataSynced() bool`

HasDataSynced returns a boolean if a field has been set.

### SetDataSyncedNil

`func (o *ConfigContextProfile) SetDataSyncedNil(b bool)`

 SetDataSyncedNil sets the value for DataSynced to be an explicit nil

### UnsetDataSynced
`func (o *ConfigContextProfile) UnsetDataSynced()`

UnsetDataSynced ensures that no value is present for DataSynced, not even an explicit nil
### GetCreated

`func (o *ConfigContextProfile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConfigContextProfile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConfigContextProfile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ConfigContextProfile) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *ConfigContextProfile) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ConfigContextProfile) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ConfigContextProfile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ConfigContextProfile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ConfigContextProfile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ConfigContextProfile) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *ConfigContextProfile) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ConfigContextProfile) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


