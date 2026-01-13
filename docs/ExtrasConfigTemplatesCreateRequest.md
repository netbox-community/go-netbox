# ExtrasConfigTemplatesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentParams** | Pointer to **interface{}** | Any &lt;a href&#x3D;\&quot;https://jinja.palletsprojects.com/en/stable/api/#jinja2.Environment\&quot;&gt;additional parameters&lt;/a&gt; to pass when constructing the Jinja environment | [optional] 
**TemplateCode** | **string** | Jinja template code. | 
**MimeType** | Pointer to **string** | Defaults to &lt;code&gt;text/plain; charset&#x3D;utf-8&lt;/code&gt; | [optional] 
**FileName** | Pointer to **string** | Filename to give to the rendered export file | [optional] 
**FileExtension** | Pointer to **string** | Extension to append to the rendered filename | [optional] 
**AsAttachment** | Pointer to **bool** | Download file as attachment | [optional] 
**DataSource** | Pointer to [**ConfigContextProfileRequestDataSource**](ConfigContextProfileRequestDataSource.md) |  | [optional] 
**AutoSyncEnabled** | Pointer to **bool** | Enable automatic synchronization of data when the data file is updated | [optional] 
**Owner** | Pointer to [**NullableASNRangeRequestOwner**](ASNRangeRequestOwner.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 

## Methods

### NewExtrasConfigTemplatesCreateRequest

`func NewExtrasConfigTemplatesCreateRequest(name string, templateCode string, ) *ExtrasConfigTemplatesCreateRequest`

NewExtrasConfigTemplatesCreateRequest instantiates a new ExtrasConfigTemplatesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtrasConfigTemplatesCreateRequestWithDefaults

`func NewExtrasConfigTemplatesCreateRequestWithDefaults() *ExtrasConfigTemplatesCreateRequest`

NewExtrasConfigTemplatesCreateRequestWithDefaults instantiates a new ExtrasConfigTemplatesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtrasConfigTemplatesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtrasConfigTemplatesCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExtrasConfigTemplatesCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtrasConfigTemplatesCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtrasConfigTemplatesCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentParams

`func (o *ExtrasConfigTemplatesCreateRequest) GetEnvironmentParams() interface{}`

GetEnvironmentParams returns the EnvironmentParams field if non-nil, zero value otherwise.

### GetEnvironmentParamsOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetEnvironmentParamsOk() (*interface{}, bool)`

GetEnvironmentParamsOk returns a tuple with the EnvironmentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentParams

`func (o *ExtrasConfigTemplatesCreateRequest) SetEnvironmentParams(v interface{})`

SetEnvironmentParams sets EnvironmentParams field to given value.

### HasEnvironmentParams

`func (o *ExtrasConfigTemplatesCreateRequest) HasEnvironmentParams() bool`

HasEnvironmentParams returns a boolean if a field has been set.

### SetEnvironmentParamsNil

`func (o *ExtrasConfigTemplatesCreateRequest) SetEnvironmentParamsNil(b bool)`

 SetEnvironmentParamsNil sets the value for EnvironmentParams to be an explicit nil

### UnsetEnvironmentParams
`func (o *ExtrasConfigTemplatesCreateRequest) UnsetEnvironmentParams()`

UnsetEnvironmentParams ensures that no value is present for EnvironmentParams, not even an explicit nil
### GetTemplateCode

`func (o *ExtrasConfigTemplatesCreateRequest) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *ExtrasConfigTemplatesCreateRequest) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.


### GetMimeType

`func (o *ExtrasConfigTemplatesCreateRequest) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ExtrasConfigTemplatesCreateRequest) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ExtrasConfigTemplatesCreateRequest) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFileName

`func (o *ExtrasConfigTemplatesCreateRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ExtrasConfigTemplatesCreateRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ExtrasConfigTemplatesCreateRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileExtension

`func (o *ExtrasConfigTemplatesCreateRequest) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *ExtrasConfigTemplatesCreateRequest) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.

### HasFileExtension

`func (o *ExtrasConfigTemplatesCreateRequest) HasFileExtension() bool`

HasFileExtension returns a boolean if a field has been set.

### GetAsAttachment

`func (o *ExtrasConfigTemplatesCreateRequest) GetAsAttachment() bool`

GetAsAttachment returns the AsAttachment field if non-nil, zero value otherwise.

### GetAsAttachmentOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetAsAttachmentOk() (*bool, bool)`

GetAsAttachmentOk returns a tuple with the AsAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsAttachment

`func (o *ExtrasConfigTemplatesCreateRequest) SetAsAttachment(v bool)`

SetAsAttachment sets AsAttachment field to given value.

### HasAsAttachment

`func (o *ExtrasConfigTemplatesCreateRequest) HasAsAttachment() bool`

HasAsAttachment returns a boolean if a field has been set.

### GetDataSource

`func (o *ExtrasConfigTemplatesCreateRequest) GetDataSource() ConfigContextProfileRequestDataSource`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetDataSourceOk() (*ConfigContextProfileRequestDataSource, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ExtrasConfigTemplatesCreateRequest) SetDataSource(v ConfigContextProfileRequestDataSource)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *ExtrasConfigTemplatesCreateRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetAutoSyncEnabled

`func (o *ExtrasConfigTemplatesCreateRequest) GetAutoSyncEnabled() bool`

GetAutoSyncEnabled returns the AutoSyncEnabled field if non-nil, zero value otherwise.

### GetAutoSyncEnabledOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetAutoSyncEnabledOk() (*bool, bool)`

GetAutoSyncEnabledOk returns a tuple with the AutoSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSyncEnabled

`func (o *ExtrasConfigTemplatesCreateRequest) SetAutoSyncEnabled(v bool)`

SetAutoSyncEnabled sets AutoSyncEnabled field to given value.

### HasAutoSyncEnabled

`func (o *ExtrasConfigTemplatesCreateRequest) HasAutoSyncEnabled() bool`

HasAutoSyncEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *ExtrasConfigTemplatesCreateRequest) GetOwner() ASNRangeRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetOwnerOk() (*ASNRangeRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ExtrasConfigTemplatesCreateRequest) SetOwner(v ASNRangeRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ExtrasConfigTemplatesCreateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ExtrasConfigTemplatesCreateRequest) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ExtrasConfigTemplatesCreateRequest) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetTags

`func (o *ExtrasConfigTemplatesCreateRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExtrasConfigTemplatesCreateRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExtrasConfigTemplatesCreateRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ExtrasConfigTemplatesCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


