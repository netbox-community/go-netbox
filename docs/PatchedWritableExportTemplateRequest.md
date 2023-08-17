# PatchedWritableExportTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TemplateCode** | Pointer to **string** | Jinja2 template code. The list of objects being exported is passed as a context variable named &lt;code&gt;queryset&lt;/code&gt;. | [optional] 
**MimeType** | Pointer to **string** | Defaults to &lt;code&gt;text/plain; charset&#x3D;utf-8&lt;/code&gt; | [optional] 
**FileExtension** | Pointer to **string** | Extension to append to the rendered filename | [optional] 
**AsAttachment** | Pointer to **bool** | Download file as attachment | [optional] 
**DataSource** | Pointer to **NullableInt32** | Remote data source | [optional] 

## Methods

### NewPatchedWritableExportTemplateRequest

`func NewPatchedWritableExportTemplateRequest() *PatchedWritableExportTemplateRequest`

NewPatchedWritableExportTemplateRequest instantiates a new PatchedWritableExportTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableExportTemplateRequestWithDefaults

`func NewPatchedWritableExportTemplateRequestWithDefaults() *PatchedWritableExportTemplateRequest`

NewPatchedWritableExportTemplateRequestWithDefaults instantiates a new PatchedWritableExportTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTypes

`func (o *PatchedWritableExportTemplateRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *PatchedWritableExportTemplateRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *PatchedWritableExportTemplateRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *PatchedWritableExportTemplateRequest) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritableExportTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableExportTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableExportTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableExportTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableExportTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableExportTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableExportTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableExportTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplateCode

`func (o *PatchedWritableExportTemplateRequest) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *PatchedWritableExportTemplateRequest) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *PatchedWritableExportTemplateRequest) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.

### HasTemplateCode

`func (o *PatchedWritableExportTemplateRequest) HasTemplateCode() bool`

HasTemplateCode returns a boolean if a field has been set.

### GetMimeType

`func (o *PatchedWritableExportTemplateRequest) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *PatchedWritableExportTemplateRequest) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *PatchedWritableExportTemplateRequest) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *PatchedWritableExportTemplateRequest) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFileExtension

`func (o *PatchedWritableExportTemplateRequest) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *PatchedWritableExportTemplateRequest) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *PatchedWritableExportTemplateRequest) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.

### HasFileExtension

`func (o *PatchedWritableExportTemplateRequest) HasFileExtension() bool`

HasFileExtension returns a boolean if a field has been set.

### GetAsAttachment

`func (o *PatchedWritableExportTemplateRequest) GetAsAttachment() bool`

GetAsAttachment returns the AsAttachment field if non-nil, zero value otherwise.

### GetAsAttachmentOk

`func (o *PatchedWritableExportTemplateRequest) GetAsAttachmentOk() (*bool, bool)`

GetAsAttachmentOk returns a tuple with the AsAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsAttachment

`func (o *PatchedWritableExportTemplateRequest) SetAsAttachment(v bool)`

SetAsAttachment sets AsAttachment field to given value.

### HasAsAttachment

`func (o *PatchedWritableExportTemplateRequest) HasAsAttachment() bool`

HasAsAttachment returns a boolean if a field has been set.

### GetDataSource

`func (o *PatchedWritableExportTemplateRequest) GetDataSource() int32`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *PatchedWritableExportTemplateRequest) GetDataSourceOk() (*int32, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *PatchedWritableExportTemplateRequest) SetDataSource(v int32)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *PatchedWritableExportTemplateRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### SetDataSourceNil

`func (o *PatchedWritableExportTemplateRequest) SetDataSourceNil(b bool)`

 SetDataSourceNil sets the value for DataSource to be an explicit nil

### UnsetDataSource
`func (o *PatchedWritableExportTemplateRequest) UnsetDataSource()`

UnsetDataSource ensures that no value is present for DataSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


