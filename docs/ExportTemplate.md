# ExportTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Display** | **string** |  | [readonly] 
**ContentTypes** | **[]string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**TemplateCode** | **string** | Jinja2 template code. The list of objects being exported is passed as a context variable named &lt;code&gt;queryset&lt;/code&gt;. | 
**MimeType** | Pointer to **string** | Defaults to &lt;code&gt;text/plain; charset&#x3D;utf-8&lt;/code&gt; | [optional] 
**FileExtension** | Pointer to **string** | Extension to append to the rendered filename | [optional] 
**AsAttachment** | Pointer to **bool** | Download file as attachment | [optional] 
**DataSource** | Pointer to [**NestedDataSource**](NestedDataSource.md) |  | [optional] 
**DataPath** | **string** | Path to remote file (relative to data source root) | [readonly] 
**DataFile** | [**NestedDataFile**](NestedDataFile.md) |  | [readonly] 
**DataSynced** | **NullableTime** |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewExportTemplate

`func NewExportTemplate(id int32, url string, display string, contentTypes []string, name string, templateCode string, dataPath string, dataFile NestedDataFile, dataSynced NullableTime, created NullableTime, lastUpdated NullableTime, ) *ExportTemplate`

NewExportTemplate instantiates a new ExportTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTemplateWithDefaults

`func NewExportTemplateWithDefaults() *ExportTemplate`

NewExportTemplateWithDefaults instantiates a new ExportTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExportTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExportTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExportTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *ExportTemplate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExportTemplate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExportTemplate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDisplay

`func (o *ExportTemplate) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ExportTemplate) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ExportTemplate) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetContentTypes

`func (o *ExportTemplate) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *ExportTemplate) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *ExportTemplate) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.


### GetName

`func (o *ExportTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExportTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExportTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExportTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExportTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplateCode

`func (o *ExportTemplate) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *ExportTemplate) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *ExportTemplate) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.


### GetMimeType

`func (o *ExportTemplate) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ExportTemplate) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ExportTemplate) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ExportTemplate) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFileExtension

`func (o *ExportTemplate) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *ExportTemplate) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *ExportTemplate) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.

### HasFileExtension

`func (o *ExportTemplate) HasFileExtension() bool`

HasFileExtension returns a boolean if a field has been set.

### GetAsAttachment

`func (o *ExportTemplate) GetAsAttachment() bool`

GetAsAttachment returns the AsAttachment field if non-nil, zero value otherwise.

### GetAsAttachmentOk

`func (o *ExportTemplate) GetAsAttachmentOk() (*bool, bool)`

GetAsAttachmentOk returns a tuple with the AsAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsAttachment

`func (o *ExportTemplate) SetAsAttachment(v bool)`

SetAsAttachment sets AsAttachment field to given value.

### HasAsAttachment

`func (o *ExportTemplate) HasAsAttachment() bool`

HasAsAttachment returns a boolean if a field has been set.

### GetDataSource

`func (o *ExportTemplate) GetDataSource() NestedDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ExportTemplate) GetDataSourceOk() (*NestedDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ExportTemplate) SetDataSource(v NestedDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *ExportTemplate) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataPath

`func (o *ExportTemplate) GetDataPath() string`

GetDataPath returns the DataPath field if non-nil, zero value otherwise.

### GetDataPathOk

`func (o *ExportTemplate) GetDataPathOk() (*string, bool)`

GetDataPathOk returns a tuple with the DataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPath

`func (o *ExportTemplate) SetDataPath(v string)`

SetDataPath sets DataPath field to given value.


### GetDataFile

`func (o *ExportTemplate) GetDataFile() NestedDataFile`

GetDataFile returns the DataFile field if non-nil, zero value otherwise.

### GetDataFileOk

`func (o *ExportTemplate) GetDataFileOk() (*NestedDataFile, bool)`

GetDataFileOk returns a tuple with the DataFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFile

`func (o *ExportTemplate) SetDataFile(v NestedDataFile)`

SetDataFile sets DataFile field to given value.


### GetDataSynced

`func (o *ExportTemplate) GetDataSynced() time.Time`

GetDataSynced returns the DataSynced field if non-nil, zero value otherwise.

### GetDataSyncedOk

`func (o *ExportTemplate) GetDataSyncedOk() (*time.Time, bool)`

GetDataSyncedOk returns a tuple with the DataSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSynced

`func (o *ExportTemplate) SetDataSynced(v time.Time)`

SetDataSynced sets DataSynced field to given value.


### SetDataSyncedNil

`func (o *ExportTemplate) SetDataSyncedNil(b bool)`

 SetDataSyncedNil sets the value for DataSynced to be an explicit nil

### UnsetDataSynced
`func (o *ExportTemplate) UnsetDataSynced()`

UnsetDataSynced ensures that no value is present for DataSynced, not even an explicit nil
### GetCreated

`func (o *ExportTemplate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ExportTemplate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ExportTemplate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ExportTemplate) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ExportTemplate) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ExportTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ExportTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ExportTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ExportTemplate) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ExportTemplate) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


