# ExportTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypes** | **[]string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**TemplateCode** | **string** | Jinja2 template code. The list of objects being exported is passed as a context variable named &lt;code&gt;queryset&lt;/code&gt;. | 
**MimeType** | Pointer to **string** | Defaults to &lt;code&gt;text/plain; charset&#x3D;utf-8&lt;/code&gt; | [optional] 
**FileExtension** | Pointer to **string** | Extension to append to the rendered filename | [optional] 
**AsAttachment** | Pointer to **bool** | Download file as attachment | [optional] 
**DataSource** | Pointer to [**BriefDataSourceRequest**](BriefDataSourceRequest.md) |  | [optional] 

## Methods

### NewExportTemplateRequest

`func NewExportTemplateRequest(objectTypes []string, name string, templateCode string, ) *ExportTemplateRequest`

NewExportTemplateRequest instantiates a new ExportTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTemplateRequestWithDefaults

`func NewExportTemplateRequestWithDefaults() *ExportTemplateRequest`

NewExportTemplateRequestWithDefaults instantiates a new ExportTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypes

`func (o *ExportTemplateRequest) GetObjectTypes() []string`

GetObjectTypes returns the ObjectTypes field if non-nil, zero value otherwise.

### GetObjectTypesOk

`func (o *ExportTemplateRequest) GetObjectTypesOk() (*[]string, bool)`

GetObjectTypesOk returns a tuple with the ObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypes

`func (o *ExportTemplateRequest) SetObjectTypes(v []string)`

SetObjectTypes sets ObjectTypes field to given value.


### GetName

`func (o *ExportTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExportTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExportTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExportTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExportTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplateCode

`func (o *ExportTemplateRequest) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *ExportTemplateRequest) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *ExportTemplateRequest) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.


### GetMimeType

`func (o *ExportTemplateRequest) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ExportTemplateRequest) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ExportTemplateRequest) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ExportTemplateRequest) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFileExtension

`func (o *ExportTemplateRequest) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *ExportTemplateRequest) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *ExportTemplateRequest) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.

### HasFileExtension

`func (o *ExportTemplateRequest) HasFileExtension() bool`

HasFileExtension returns a boolean if a field has been set.

### GetAsAttachment

`func (o *ExportTemplateRequest) GetAsAttachment() bool`

GetAsAttachment returns the AsAttachment field if non-nil, zero value otherwise.

### GetAsAttachmentOk

`func (o *ExportTemplateRequest) GetAsAttachmentOk() (*bool, bool)`

GetAsAttachmentOk returns a tuple with the AsAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsAttachment

`func (o *ExportTemplateRequest) SetAsAttachment(v bool)`

SetAsAttachment sets AsAttachment field to given value.

### HasAsAttachment

`func (o *ExportTemplateRequest) HasAsAttachment() bool`

HasAsAttachment returns a boolean if a field has been set.

### GetDataSource

`func (o *ExportTemplateRequest) GetDataSource() BriefDataSourceRequest`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ExportTemplateRequest) GetDataSourceOk() (*BriefDataSourceRequest, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ExportTemplateRequest) SetDataSource(v BriefDataSourceRequest)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *ExportTemplateRequest) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


